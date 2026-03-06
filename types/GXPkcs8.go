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

// GXPkcs8 represents a PKCS #8 private key container.
//
// It holds the raw encoded form as well as decoded private/public key objects.
// This type supports parsing from DER/PEM and serializing back to those formats.
//
// See: https://tools.ietf.org/html/rfc5208
type GXPkcs8 struct {
	// Loaded PKCS #8 certificate as raw bytes.
	_rawData []byte

	// Algorithm holds the key algorithm identifier.
	algorithm enums.X9ObjectIdentifier

	// Private key.
	privateKey *GXPrivateKey

	// Public key.
	publicKey *GXPublicKey

	// Description is optional metadata stored in the PEM comment header.
	Description string

	// Version is the PKCS #8 version value.
	Version enums.CertificateVersion
}

// NewGXPkcs8 parses a PKCS #8 blob and returns a GXPkcs8 instance.
func NewGXPkcs8(data []byte) (*GXPkcs8, error) {
	ret := &GXPkcs8{}
	err := ret.init(data)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// NewGXPkcs8FromKeys creates a GXPkcs8 object from an existing public/private key pair.
func NewGXPkcs8FromKeys(keyValuePair *GXKeyValuePair[*GXPublicKey, *GXPrivateKey]) (*GXPkcs8, error) {
	ret := &GXPkcs8{}
	ret.publicKey = keyValuePair.Key
	ret.privateKey = keyValuePair.Value
	return ret, nil
}

// PrivateKey returns the wrapped private key.
func (g *GXPkcs8) PrivateKey() *GXPrivateKey {
	return g.privateKey
}

// PublicKey returns the wrapped public key.
func (g *GXPkcs8) PublicKey() *GXPublicKey {
	return g.publicKey
}

// encoded returns the PKCS #8 DER-encoded byte array for this instance.
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

func (g *GXPkcs8) init(data []byte) error {
	g._rawData = data
	ret, err := Asn1FromByteArray(data)
	if err != nil {
		return err
	}
	seq := *ret.(*GXAsn1Sequence)
	var tmp GXAsn1Sequence
	var tmp2 []any
	if len(seq) < 3 {
		return errors.New("Wrong number of elements in sequence.")
	}
	if _, ok := seq[0].(int8); !ok {
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
	tmp = *seq[1].(*GXAsn1Sequence)
	oi := tmp[0].(*GXAsn1ObjectIdentifier)
	g.algorithm = X9ObjectIdentifierFromString(oi.String())
	s := *seq[2].(*GXAsn1Sequence)
	g.privateKey, err = PrivateKeyFromRawBytes(s[1].([]byte))
	if err != nil {
		return err
	}
	if g.privateKey == nil {
		return errors.New("Invalid private key.")
	}
	// If public key is not included.
	tmp2 = ([]any)(*seq[2].(*GXAsn1Sequence))
	if len(tmp2) > 2 {
		bs := tmp2[2].(*GXAsn1Context).Items[0].(*GXBitString)
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

// GetFilePathFromSystemTitle returns the default file path for storing a PKCS #8 key.
//
// The returned path includes a prefix based on certificateType and the systemTitle.
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

// GetFilePath returns the default file path for storing the PKCS #8 key.
//
// The returned path differs based on certificate type and the key's unique identifier.
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
// Pkcs8FromPem parses a PEM-encoded PKCS #8 key and returns a GXPkcs8 instance.
//
// The input should include the "-----BEGIN PRIVATE KEY-----" header.
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
// FromHexString parses a hex-encoded PKCS #8 value and returns a GXPkcs8.
func (g *GXPkcs8) FromHexString(data string) GXPkcs8 {
	cert := GXPkcs8{}
	cert.init(HexToBytes(data))
	return cert
}

// Pkcs8FromDer parses a Base64-encoded DER PKCS #8 key and returns a GXPkcs8.
//
// The input should be the Base64 encoding of an ASN.1 DER PKCS #8 structure.
func Pkcs8FromDer(der string) (*GXPkcs8, error) {
	der = strings.ReplaceAll(der, "\r\n", "")
	der = strings.ReplaceAll(der, "\n", "")
	key, err := base64.StdEncoding.DecodeString(der)
	if err != nil {
		return nil, err
	}
	cert := GXPkcs8{}
	err = cert.init(key)
	if err != nil {
		return nil, err
	}
	return &cert, nil
}

// String implements fmt.Stringer and returns a compact description of the PKCS #8 key.
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

// Pkcs8Load reads a PEM file and returns a parsed GXPkcs8 instance.
func Pkcs8Load(path string) (*GXPkcs8, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return Pkcs8FromPem(string(data))
}

// Save writes the PKCS #8 key to a PEM file at the given path.
func (g *GXPkcs8) Save(path string) error {
	ret, err := g.ToPem()
	if err != nil {
		return err
	}
	return os.WriteFile(path, []byte(ret), 0644)
}

// ToPem returns the private key encoded as a PEM string.
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

// ToDer returns the PKCS #8 key as a Base64-encoded DER string.
func (g *GXPkcs8) ToDer() string {
	ret, err := g.encoded()
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(ret)
}

// Import parses a PKCS #8 key from a string in PEM, DER (base64), or raw private-key formats.
//
// The method attempts multiple parsing strategies until one succeeds.
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

// Equals reports whether the specified object represents the same PKCS #8 key.
func (g *GXPkcs8) Equals(obj any) bool {
	if o, ok := obj.(GXPkcs8); ok {
		return g.privateKey.Equals(o.PrivateKey())
	}
	return false
}
