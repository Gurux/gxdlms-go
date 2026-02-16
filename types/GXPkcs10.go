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
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal/buffer"
	"github.com/Gurux/gxdlms-go/internal/helpers"
)

// Pkcs10 Certificate Signing Request.
// https://tools.ietf.org/html/rfc2986
type GXPkcs10 struct {
	// Loaded PKCS #10 certificate as a raw data.
	rawData []byte

	// Certificate version.
	version enums.CertificateVersion

	// Subject.
	subject string

	// Collection of attributes providing additional information about the subject of the certificate.
	attributes []GXKeyValuePair[enums.PkcsObjectIdentifier, []any]

	// Algorithm.
	algorithm enums.X9ObjectIdentifier

	// Subject public key.
	publicKey *GXPublicKey

	// Signature algorithm.
	signatureAlgorithm enums.HashAlgorithm

	// Signature parameters.
	signatureParameters any

	// Signature.
	signature []byte
}

// Certificate version.
func (g *GXPkcs10) Version() enums.CertificateVersion {
	return g.version
}

// Subject.
func (g *GXPkcs10) Subject() string {
	return g.subject
}

// Collection of attributes providing additional information about the subject of the certificate.
func (g *GXPkcs10) Attributes() []GXKeyValuePair[enums.PkcsObjectIdentifier, []any] {
	return g.attributes
}

// Algorithm.
func (g *GXPkcs10) Algorithm() enums.X9ObjectIdentifier {
	return g.algorithm
}

// Subject public key.
func (g *GXPkcs10) PublicKey() *GXPublicKey {
	return g.publicKey
}

// Signature algorithm.
func (g *GXPkcs10) SignatureAlgorithm() enums.HashAlgorithm {
	return g.signatureAlgorithm
}

// Signature parameters.
func (g *GXPkcs10) SignatureParameters() any {
	return g.signatureParameters
}

// Signature.
func (g *GXPkcs10) Signature() []byte {
	return g.signature
}

func (g *GXPkcs10) Encoded() ([]byte, error) {
	if g.rawData != nil {
		return g.rawData, nil
	}
	if g.signature == nil {
		return nil, errors.New("Sign first.")
	}
	// Certification request info. subject Public key info.
	ret, err := HashAlgorithmToString(g.signatureAlgorithm)
	if err != nil {
		return nil, err
	}
	sa := NewGXAsn1ObjectIdentifier(ret)
	bs, err := NewGXBitString(g.signature, 0)
	if err != nil {
		return nil, err
	}
	list := []any{g.getData(), []any{sa}, bs}
	return Asn1ToByteArray(list)
}

func (g *GXPkcs10) init(data []byte) error {
	g.rawData = data
	ret, err := Asn1FromByteArray(data)
	if err != nil {
		return err
	}
	seq := ret.(GXAsn1Sequence)
	if len(seq) < 3 {
		return errors.New("Wrong number of elements in sequence.")
	}
	if _, ok := seq[0].(GXAsn1Sequence); ok {
		ret, err := Asn1GetCertificateType(data, seq)
		if err != nil {
			return err
		}
		type_ := ret
		switch type_ {
		case enums.PkcsTypePkcs8:
			return errors.New("Invalid Certificate. This is PKCS 8, not PKCS 10.")
		case enums.PkcsTypex509Certificate:
			return errors.New("Invalid Certificate. This is PKCS x509 certificate, not PKCS 10.")
		}
		return errors.New("Invalid Certificate Version.")
	}
	reqInfo := seq[0].(GXAsn1Sequence)
	g.version = enums.CertificateVersion(reqInfo[0].(int8))
	g.subject = Asn1GetSubject(reqInfo[1].(*GXAsn1Sequence))
	// subject Public key info.
	subjectPKInfo := reqInfo[2].(GXAsn1Sequence)
	if len(reqInfo) > 3 {
		for _, i := range reqInfo[3].(GXAsn1Context).Items {
			it := i.(GXAsn1Sequence)
			values := []any{}
			t := it[1].(GXKeyValuePair[any, any])
			for _, v := range t.Key.([]any) {
				values = append(values, v)
			}
			id_ := PkcsObjectIdentifierFromString(fmt.Sprintf("%v", it[0]))
			g.attributes = append(g.attributes, *NewGXKeyValuePair(id_, values))
		}
	}
	tmp := subjectPKInfo[0].(GXAsn1Sequence)
	g.algorithm = X9ObjectIdentifierFromString(tmp[0].(string))
	if g.algorithm != enums.X9ObjectIdentifierIdECPublicKey {
		var algorithm int
		algorithm = int(g.algorithm)
		if g.algorithm == enums.X9ObjectIdentifierNone {
			algorithm = int(PkcsObjectIdentifierFromString(tmp[0].(string)))
			if enums.PkcsObjectIdentifier(algorithm) == enums.PkcsObjectIdentifierNone {
				algorithm = tmp[0].(int)
			}
		}
		return fmt.Errorf("Invalid PKCS #10 certificate algorithm. %s", g.algorithm)
	}
	pub, err := PublicKeyFromRawBytes(subjectPKInfo[1].(*GXBitString).Value())
	if err != nil {
		return err
	}
	g.publicKey = pub
	err = EcdsaValidate(g.publicKey)
	if err != nil {
		return err
	}
	// signatureAlgorithm
	sign := seq[1].(GXAsn1Sequence)
	g.signatureAlgorithm = HashAlgorithmFromString(sign[0].(string))
	if g.signatureAlgorithm != enums.HashAlgorithmSha256WithEcdsa && g.signatureAlgorithm != enums.HashAlgorithmSha384WithEcdsa {
		return fmt.Errorf("Invalid signature algorithm. %s", sign[0].(string))
	}
	if len(sign) != 1 {
		g.signatureParameters = sign[1]
	}
	// signatureGet raw data
	tmp2 := GXByteBuffer{}
	err = tmp2.Set(data)
	if err != nil {
		return err
	}
	Asn1GetNext(&tmp2)
	tmp2.SetSize(tmp2.Position())
	tmp2.SetPosition(1)
	GetObjectCount(&tmp2)
	g.signature = seq[2].(*GXBitString).Value()
	e, err := NewGXEcdsaFromPublicKey(g.publicKey)
	if err != nil {
		return err
	}
	ret, err = Asn1FromByteArray(g.signature)
	tmp3 := ret.(GXAsn1Sequence)
	bb := GXByteBuffer{}
	var size int
	var add int
	if g.signatureAlgorithm == enums.HashAlgorithmSha256WithEcdsa {
		size = 32
	} else {
		size = 48
	}
	v := tmp3[0].(*GXAsn1Integer).Value()
	if len(v) == size {
		add = 0
	} else {
		add = 1
	}
	bb.SetAt(v, add, size)
	v = tmp3[1].(*GXAsn1Integer).Value()
	if len(v) == size {
		add = 0
	} else {
		add = 1
	}
	bb.SetAt(v, add, size)
	tmp4, err := tmp2.SubArray(tmp2.Position(), tmp2.Available())
	if err != nil {
		return err
	}
	ret2, _ := e.Verify(bb.Array(), tmp4)
	if !ret2 {
		return errors.New("Invalid Signature.")
	}
	return nil
}

func (g *GXPkcs10) getData() []any {
	var alg *GXAsn1ObjectIdentifier
	if g.publicKey.Scheme() == enums.EccP256 {
		alg = NewGXAsn1ObjectIdentifier("1.2.840.10045.3.1.7")
	} else {
		alg = NewGXAsn1ObjectIdentifier("1.3.132.0.34")
	}
	subjectPKInfo, err := NewGXBitString(g.publicKey.RawValue(), 0)
	if err != nil {
		return nil
	}
	tmp := []any{NewGXAsn1ObjectIdentifier("1.2.840.10045.2.1"), alg}
	attributes := GXAsn1Context{}
	for _, v := range g.attributes {
		s := GXAsn1Sequence{}
		ret, err := PkcsObjectIdentifierToString(v.Key)
		if err != nil {
			return nil
		}
		s = append(s, NewGXAsn1ObjectIdentifier(ret))
		values := []any{}
		for _, v2 := range v.Value {
			values = append(values, v2)
		}
		s = append(s, NewGXKeyValuePair[any, any](values, nil))
		attributes.Items = append(attributes.Items, s)
	}
	ret, err := Asn1EncodeSubject(g.subject)
	if err != nil {
		return nil
	}
	return []any{int8(g.version), ret, []any{tmp, subjectPKInfo}, attributes}
}

// Sign returns the sign
//
// Parameters:
//
//	key: Private key.
//	hashAlgorithm: Used algorithm for signing.
func (g *GXPkcs10) Sign(key *GXPrivateKey, hashAlgorithm enums.HashAlgorithm) error {
	data, err := Asn1ToByteArray(g.getData())
	if err != nil {
		return err
	}
	e, err := NewGXEcdsaFromPrivateKey(key)
	if err != nil {
		return err
	}
	g.signatureAlgorithm = hashAlgorithm
	bb := GXByteBuffer{}
	ret, err := e.Sign(data)
	err = bb.Set(ret)
	if err != nil {
		return err
	}
	var size int
	if g.signatureAlgorithm == enums.HashAlgorithmSha256WithEcdsa {
		size = 32
	} else {
		size = 48
	}
	v1, err := bb.SubArray(0, size)
	if err != nil {
		return err
	}
	v2, err := bb.SubArray(size, size)
	if err != nil {
		return err
	}
	tmp := []any{NewGXAsn1Integer(v1), NewGXAsn1Integer(v2)}
	g.signature, err = Asn1ToByteArray(tmp)
	if err != nil {
		return err
	}
	return nil
}

// FromHexString returns the create PKCS 10 from hex string.
//
// Parameters:
//
//	data: Hex string.
//
// Returns:
//
//	PKCS 10
func NewGXPkcs10(data []byte) (*GXPkcs10, error) {
	cert := GXPkcs10{}
	err := cert.init(data)
	return &cert, err
}

// FromHexString returns the create PKCS 10 from hex string.
//
// Parameters:
//
//	data: Hex string.
//
// Returns:
//
//	PKCS 10
func Pkcs10FromHexString(data string) (*GXPkcs10, error) {
	cert := GXPkcs10{}
	err := cert.init(HexToBytes(data))
	return &cert, err
}

// FromPem returns the create x509Certificate from PEM string.
//
// Parameters:
//
//	data: PEM string.
func Pkcs10FromPem(data string) (*GXPkcs10, error) {
	const START = "CERTIFICATE REQUEST-----"
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
	return Pkcs10FromDer(strings.TrimSpace(data[0:end]))
}

// FromDer returns the create x509Certificate from DER Base64 encoded string.
//
// Parameters:
//
//	der: Base64 DER string.
func Pkcs10FromDer(der string) (*GXPkcs10, error) {
	der = strings.ReplaceAll(der, "\r\n", "")
	der = strings.ReplaceAll(der, "\n", "")
	key, err := base64.StdEncoding.DecodeString(der)
	if err != nil {
		return nil, err
	}
	cert := GXPkcs10{}
	err = cert.init(key)
	return &cert, err
}

// String returns the PKCS#10 certificate request as a string.
func (g *GXPkcs10) String() string {
	var bb strings.Builder
	bb.WriteString("PKCS #10 certificate request:\n")
	bb.WriteString("Version: ")
	bb.WriteString(g.version.String())
	bb.WriteString("\n")

	bb.WriteString("Subject: ")
	bb.WriteString(g.subject)
	bb.WriteString("\n")

	bb.WriteString("Algorithm: ")
	bb.WriteString(g.algorithm.String())
	bb.WriteString("\n")
	bb.WriteString("Public Key: ")
	if g.publicKey != nil {
		bb.WriteString(g.publicKey.String())
	}
	bb.WriteString("\n")
	bb.WriteString("Signature algorithm: ")
	bb.WriteString(g.signatureAlgorithm.String())
	bb.WriteString("\n")
	if g.signatureParameters != nil {
		bb.WriteString("Signature parameters: ")
		bb.WriteString(fmt.Sprint(g.signatureParameters))
		bb.WriteString("\n")
	}
	bb.WriteString("Signature: ")
	bb.WriteString(buffer.ToHex(g.signature, false))
	bb.WriteString("\n")
	return bb.String()
}

// Pkcs10CreateCertificateSigningRequest returns the create Certificate Signing Request.
//
// Parameters:
//
//	kp: KeyPair
//	subject: Subject.
//
// Returns:
//
//	Created GXPkcs10.
func Pkcs10CreateCertificateSigningRequest(kp *GXKeyValuePair[*GXPublicKey, *GXPrivateKey], subject string) (*GXPkcs10, error) {
	if !strings.Contains(subject, "CN=") {
		return nil, errors.New("subject")
	}
	pkc10 := GXPkcs10{}
	pkc10.algorithm = enums.X9ObjectIdentifierIdECPublicKey
	pkc10.publicKey = kp.Key
	pkc10.subject = subject
	var alg enums.HashAlgorithm
	if kp.Key.Scheme() == enums.EccP256 {
		alg = enums.HashAlgorithmSha256WithEcdsa
	} else {
		alg = enums.HashAlgorithmSha384WithEcdsa
	}
	err := pkc10.Sign(kp.Value, alg)
	if err != nil {
		return nil, err
	}
	return &pkc10, nil
}

// GetCertificate requests the Gurux certificate server to generate new certificates.
//
// Parameters:
//
//	address: Certificate server address.
//	certifications: List of certification requests.
//
// Returns:
//
//	Generated certificate(s) or error if request fails.
func (g *GXPkcs10) GetCertificate(address string, certifications []GXCertificateRequest) ([]GXx509Certificate, error) {
	var usage strings.Builder
	for _, it := range certifications {
		if usage.Len() != 0 {
			usage.WriteString(", ")
		}
		usage.WriteString("{\"KeyUsage\":")

		var keyUsageValue int
		switch it.CertificateType {
		case enums.CertificateTypeDigitalSignature:
			keyUsageValue = int(enums.KeyUsageDigitalSignature)
		case enums.CertificateTypeKeyAgreement:
			keyUsageValue = int(enums.KeyUsageKeyAgreement)
		case enums.CertificateTypeTLS:
			keyUsageValue = int(enums.KeyUsageDigitalSignature) | int(enums.KeyUsageKeyAgreement)
		default:
			return nil, fmt.Errorf("invalid certificate type")
		}
		usage.WriteString(fmt.Sprint(keyUsageValue))

		if it.ExtendedKeyUsage != enums.ExtendedKeyUsageNone {
			usage.WriteString(", \"ExtendedKeyUsage\":")
			usage.WriteString(fmt.Sprint(int(it.ExtendedKeyUsage)))
		}
		usage.WriteString(", \"CSR\":\"")
		csr := it.Certificate.ToDer()
		usage.WriteString(csr)
		usage.WriteString("\"}")
	}

	requestBody := "{\"Certificates\":[" + usage.String() + "]}"

	req, err := http.NewRequest("POST", address, strings.NewReader(requestBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned status code %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	str := string(body)
	pos := strings.Index(str, "[")
	if pos == -1 {
		return nil, fmt.Errorf("certificates are missing")
	}

	str = str[pos+2:]
	pos = strings.Index(str, "]")
	if pos == -1 {
		return nil, fmt.Errorf("certificates are missing")
	}

	str = str[:pos-1]
	parts := strings.FieldsFunc(str, func(r rune) bool {
		return r == '"' || r == ','
	})

	var certs []GXx509Certificate
	for i, certStr := range parts {
		if i >= len(certifications) {
			break
		}
		x509, err := X509CertificateFromDer(certStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse certificate: %w", err)
		}
		if !helpers.Compare(certifications[i].Certificate.PublicKey().RawValue(), x509.PublicKey.RawValue()) {
			return nil, fmt.Errorf("certificate signing request generated wrong public key")
		}
		certs = append(certs, *x509)
	}

	return certs, nil
}

// Load returns the load Pkcs10 Certificate Signing Request from the PEM (.csr) file.
//
// Parameters:
//
//	path: File path.
//
// Returns:
//
//	Created GXPkcs10 object.
func (g *GXPkcs10) Load(path string) (*GXPkcs10, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return Pkcs10FromPem(string(data))
}

// Save returns the save Pkcs #10 Certificate Signing Request to PEM file.
//
// Parameters:
//
//	path: File path.
func (g *GXPkcs10) Save(path string) error {
	ret, err := g.ToPem()
	if err != nil {
		return err
	}
	return os.WriteFile(path, []byte(ret), 0644)
}

// ToPem returns the pkcs #10 Certificate Signing Request in DER format.
//
// Returns:
//
//	Public key as in PEM string.
func (g *GXPkcs10) ToPem() (string, error) {
	sb := strings.Builder{}
	sb.WriteString("-----BEGIN CERTIFICATE REQUEST-----")
	sb.WriteString(g.ToDer())
	sb.WriteString("-----END CERTIFICATE REQUEST-----")
	return sb.String(), nil
}

// ToDer returns the pkcs #10 Certificate Signing Request in DER format.
//
// Returns:
//
//	Public key as in PEM string.
func (g *GXPkcs10) ToDer() string {
	if g.rawData != nil {
		return base64.StdEncoding.EncodeToString([]byte(g.rawData))
	}
	ret, err := g.Encoded()
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(ret)
}
