package types

//
// --------------------------------------------------------------------------
//  Gurux Ltd
//
//
//
// Filename:        $HeadURL$
//
// Version:         $Revision$,
//                  $Date$
//                  $Author$
//
// Copyright (c) Gurux Ltd
//
//---------------------------------------------------------------------------
//
//  DESCRIPTION
//
// This file is a part of Gurux Device Framework.
//
// Gurux Device Framework is Open Source software you can redistribute it
// and/or modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation version 2 of the License.
// Gurux Device Framework is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
// See the GNU General Public License for more details.
//
// More information of Gurux products: https://www.gurux.org
//
// This code is licensed under the GNU General Public License v2.
// Full text may be retrieved at http://www.gnu.org/licenses/gpl-2.0.txt
//---------------------------------------------------------------------------

import (
	"encoding/base64"
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/Gurux/gxdlms-go/enums"
)

// Pkcs8 certification request. Private key is saved using this format.
// https://tools.ietf.org/html/rfc5208
type GXPkcs8 struct {
	// Loaded PKCS #8 certificate as a raw data.
	_rawData []byte

	// Algorithm.
	algorithm enums.X9ObjectIdentifier

	// Private key.
	privateKey *GXPrivateKey

	// Public key.
	publicKey *GXPublicKey

	// Description is extra metadata that is saved to PEM file.
	Description string

	// Private key version.
	Version enums.CertificateVersion
}

// Constructor from byte array.
func NewGXPkcs8(data []byte) (*GXPkcs8, error) {
	ret := &GXPkcs8{}
	err := ret.Init(data)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// Constructor from byte array.
func NewGXPkcs8FromKeys(keyValuePair *GXKeyValuePair[*GXPublicKey, *GXPrivateKey]) (*GXPkcs8, error) {
	ret := &GXPkcs8{}
	ret.publicKey = keyValuePair.Key
	ret.privateKey = keyValuePair.Value
	return ret, nil
}

// Private key.
func (g *GXPkcs8) PrivateKey() *GXPrivateKey {
	return g.privateKey
}

// Public key.
func (g *GXPkcs8) PublicKey() *GXPublicKey {
	return g.publicKey
}

// / encoded returns the PKCS #8 encoded byte array.
func (g *GXPkcs8) encoded() ([]byte, error) {
	if g._rawData != nil {
		return g._rawData, nil
	}
	d := GXAsn1Sequence{}
	d = append(d, int8(g.Version))
	d1 := GXAsn1Sequence{}
	s, err := X9ObjectIdentifierToString(g.algorithm)
	if err != nil {
		return nil, err
	}
	d1 = append(d1, NewGXAsn1ObjectIdentifier(s))
	var alg GXAsn1ObjectIdentifier
	if g.publicKey.Scheme() == enums.EccP256 {
		alg = *NewGXAsn1ObjectIdentifier("1.2.840.10045.3.1.7")
	} else {
		alg = *NewGXAsn1ObjectIdentifier("1.3.132.0.34")
	}
	d1 = append(d1, alg)
	d = append(d, d1)
	d2 := GXAsn1Sequence{}
	d2 = append(d2, int8(1))
	d2 = append(d2, g.privateKey.RawValue())
	d3 := GXAsn1Context{}
	d3.Index = 1
	bs, err := NewGXBitString(g.publicKey.RawValue(), 0)
	if err != nil {
		return nil, err
	}
	d3.Items = append(d3.Items, bs)
	d2 = append(d2, d3)
	ret, err := Asn1ToByteArray(d2)
	if err != nil {
		return nil, err
	}
	d = append(d, ret)
	ret, err = Asn1ToByteArray(d)
	if err != nil {
		return nil, err
	}
	g._rawData = ret
	return g._rawData, nil
}

func (g *GXPkcs8) Init(data []byte) error {
	g._rawData = data
	ret, err := Asn1FromByteArray(data)
	if err != nil {
		return err
	}
	seq := ret.(GXAsn1Sequence)
	var tmp GXAsn1Sequence
	var tmp2 []any
	if len(seq) < 3 {
		return errors.New("Wrong number of elements in sequence.")
	}
	if _, ok := seq[0].(int8); ok {
		type_, err := Asn1GetCertificateType(data, seq)
		if err != nil {
			return err
		}
		switch type_ {
		case enums.PkcsTypePkcs10:
			return errors.New("Invalid Certificate. This is PKCS 10 certification requests, not PKCS 8.")
		case enums.PkcsTypex509Certificate:
			return errors.New("Invalid Certificate. This is PKCS x509 certificate, not PKCS 8.")
		}
		return errors.New("Invalid Certificate Version.")
	}
	g.Version = enums.CertificateVersion(seq[0].(int8))
	if _, ok := seq[1].([]byte); ok {
		return errors.New("Invalid Certificate. This looks more like private key, not PKCS 8.")
	}
	tmp = seq[1].(GXAsn1Sequence)
	g.algorithm = X9ObjectIdentifierFromString(tmp[0].(string))
	s := seq[2].(GXAsn1Sequence)
	g.privateKey, err = PrivateKeyFromRawBytes(s[1].([]byte))
	if err != nil {
		return err
	}
	if g.privateKey == nil {
		return errors.New("Invalid private key.")
	}
	// If public key is not included.
	tmp2 = ([]any)(seq[2].(GXAsn1Sequence))
	if len(tmp2) > 2 {
		bs := tmp2[2].(GXAsn1Context).Items[0].(GXBitString)
		g.publicKey, err = PublicKeyFromRawBytes(bs.Value())
		if err != nil {
			return err
		}
		err = EcdsaValidate(g.publicKey)
	} else {
		g.publicKey, err = g.privateKey.GetPublicKey()
		if err != nil {
			return err
		}
	}
	return err
}

// GetFilePath returns the default file path.
//
// Parameters:
//
//	scheme: Used scheme.
//	certificateType: Certificate type.
//	systemTitle: System title.
//
// Returns:
//
//	File path.
func (g *GXPkcs8) GetFilePathFromSystemTitle(scheme enums.Ecc, certificateType enums.CertificateType, systemTitle []byte) (string, error) {
	var path string
	switch certificateType {
	case enums.CertificateTypeDigitalSignature:
		path = "D"
	case enums.CertificateTypeKeyAgreement:
		path = "A"
	case enums.CertificateTypeTLS:
		path = "T"
	default:
		return "", errors.New("Unknown certificate type.")
	}
	path = path + ToHex(systemTitle, false) + ".pem"
	if scheme == enums.EccP256 {
		path = filepath.Join("Keys", path)
	} else {
		path = filepath.Join("Keys384", path)
	}
	return path, nil
}

// GetFilePath returns the default file path.
//
// Parameters:
//
//	scheme: Used scheme.
//	certificateType: Certificate type.
func (g *GXPkcs8) GetFilePath(scheme enums.Ecc, certificateType enums.CertificateType) (string, error) {
	var path string
	switch certificateType {
	case enums.CertificateTypeDigitalSignature:
		path = "D"
	case enums.CertificateTypeKeyAgreement:
		path = "A"
	case enums.CertificateTypeTLS:
		path = "T"
	default:
		return "", errors.New("Unknown certificate type.")
	}
	path = path + g.privateKey.String()[1:] + ".pem"
	if scheme == enums.EccP256 {
		path = filepath.Join("Certificates", path)
	} else {
		path = filepath.Join("Certificates384", path)
	}
	return path, nil
}

// FromPem returns the create PKCS #8 from PEM string.
//
// Parameters:
//
//	data: PEM string.
func Pkcs8FromPem(data string) (*GXPkcs8, error) {
	const START = "PRIVATE KEY-----"
	const END = "-----END"
	data = strings.ReplaceAll(data, "\r\n", "\n")
	start := strings.Index(data, START)
	if start == -1 {
		return nil, errors.New("Invalid PEM file.")
	}
	data = data[start+len(START):]
	end := strings.Index(data, END)
	if end == -1 {
		return nil, errors.New("Invalid PEM file.")
	}
	return Pkcs8FromDer(data[0:end])
}

// FromHexString returns the create PKCS 8 from hex string.
//
// Parameters:
//
//	data: Hex string.
//
// Returns:
//
//	PKCS 8
func (g *GXPkcs8) FromHexString(data string) GXPkcs8 {
	cert := GXPkcs8{}
	cert.Init(HexToBytes(data))
	return cert
}

// Pkcs8FromDer returns the create PKCS #8 from DER Base64 encoded string.
//
// Parameters:
//
//	der: Base64 DER string.
func Pkcs8FromDer(der string) (*GXPkcs8, error) {
	der = strings.ReplaceAll(der, "\r\n", "")
	der = strings.ReplaceAll(der, "\n", "")
	key, err := base64.StdEncoding.DecodeString(der)
	if err != nil {
		return nil, err
	}
	cert := GXPkcs8{}
	err = cert.Init(key)
	if err != nil {
		return nil, err
	}
	return &cert, nil
}

func (g *GXPkcs8) String() string {
	bb := strings.Builder{}
	bb.WriteString("PKCS #8:")
	bb.WriteString("Version: ")
	bb.WriteString(g.Version.String())
	bb.WriteString("Algorithm: ")
	bb.WriteString(g.algorithm.String())
	bb.WriteString("PrivateKey: ")
	bb.WriteString(g.privateKey.ToHex())
	bb.WriteString("PublicKey: ")
	bb.WriteString(g.publicKey.String())
	return bb.String()
}

// Load returns the load private key from the PEM file.
//
// Parameters:
//
//	path: File path.
//
// Returns:
//
//	Created GXPkcs8 object.
func Pkcs8Load(path string) (*GXPkcs8, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return Pkcs8FromPem(string(data))
}

// Save returns the save private key to PEM file.
//
// Parameters:
//
//	path: File path.
func (g *GXPkcs8) Save(path string) error {
	ret, err := g.ToPem()
	if err != nil {
		return err
	}
	return os.WriteFile(path, []byte(ret), 0644)
}

// ToPem returns the private key in PEM format.
//
// Returns:
//
//	Private key as in PEM string.
func (g *GXPkcs8) ToPem() (string, error) {
	sb := strings.Builder{}
	if g.privateKey == nil {
		return "", errors.New("Public or private key is not set.")
	}
	sb.WriteString("-----BEGIN PRIVATE KEY----- \n")
	sb.WriteString(g.ToDer())
	sb.WriteString("\n-----END PRIVATE KEY-----\n")
	return sb.String(), nil
}

// ToDer returns the private key in DER format.
//
// Returns:
//
//	Private key as in DER string.
func (g *GXPkcs8) ToDer() string {
	ret, err := g.encoded()
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(ret)
}

// Import returns the import certificate from string.
func (g *GXPkcs8) Import(value string) (*GXPkcs8, error) {
	ret, err := Pkcs8FromPem(value)
	if err == nil {
		return ret, nil
	}
	ret, err = Pkcs8FromDer(value)
	if err == nil {
		return ret, nil
	}
	pk, err := PrivateKeyFromPem(value)
	if err == nil {
		pub, err := pk.GetPublicKey()
		if err != nil {
			return nil, err
		}
		ret = &GXPkcs8{privateKey: pk, publicKey: pub}
		return ret, nil
	}
	pk, err = PrivateKeyFromDer(value)
	if err == nil {
		pub, err := pk.GetPublicKey()
		if err != nil {
			return nil, err
		}
		ret = &GXPkcs8{privateKey: pk, publicKey: pub}
		return ret, nil
	}
	ret = &GXPkcs8{}
	ret.privateKey, err = PrivateKeyFromRawBytes(HexToBytes(value))
	if err == nil {
		return ret, nil
	}
	ret.publicKey, err = pk.GetPublicKey()
	if err != nil {
		return nil, err
	}
	return nil, errors.New("Invalid private key format.")
}

// Equals determines whether the specified object is equal to the current GXPkcs8 instance.
func (g *GXPkcs8) Equals(obj any) bool {
	if o, ok := obj.(GXPkcs8); ok {
		return g.privateKey.Equals(o.PrivateKey())
	}
	return false
}
