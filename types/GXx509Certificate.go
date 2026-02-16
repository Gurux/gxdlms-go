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
	"log"
	"math/big"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal/buffer"
	"github.com/Gurux/gxdlms-go/internal/constants"
)

// x509 Certificate.
// https://tools.ietf.org/html/rfc5280
type GXx509Certificate struct {
	// Loaded x509Certificate as raw data.
	rawData []byte

	// Raw Issuer in ASN1 format.
	issuerRaw []byte

	// Version.
	// Version is read-only because DLMS supports only v3.
	version enums.CertificateVersion

	// Description is extra metadata that is saved to PEM file.
	Description string

	// This extension identifies the public key being certified.
	SubjectKeyIdentifier []byte

	// May be used either as a certificate or CRL extension. It identifies the
	// public key to be used to verify the signature on this certificate or CRL.
	// It enables distinct keys used by the same CA to be distinguished.
	AuthorityKeyIdentifier []byte

	// Authority certification serial number.
	AuthorityCertificationSerialNumber big.Int

	// Indicates if the Subject may act as a CA.
	BasicConstraints bool

	// Indicates that a certificate can be used as an TLS server or client certificate.
	ExtendedKeyUsage enums.ExtendedKeyUsage

	// Signature algorithm.
	SignatureAlgorithm enums.HashAlgorithm

	// Signature Parameters.
	SignatureParameters any

	// Public key.
	PublicKey *GXPublicKey

	// Public Key algorithm.
	PublicKeyAlgorithm enums.HashAlgorithm

	// Parameters.
	PublicKeyParameters any

	// Signature.
	Signature []byte

	// Subject. Example: "CN=Test, O=Gurux, L=Tampere, C=FI".
	Subject string

	// Subject Alternative Name.
	SubjectAlternativeName string

	// Issuer. Example: "CN=Test O=Gurux, L=Tampere, C=FI".
	Issuer string

	// Authority Cert Issuer. Example: "CN=Test O=Gurux, L=Tampere, C=FI".
	AuthorityCertIssuer string

	// Serial number.
	SerialNumber *big.Int

	// Validity from.
	ValidFrom time.Time

	// Validity to.
	ValidTo time.Time

	// Indicates the purpose for which the certified public key is used.
	KeyUsage enums.KeyUsage
}

// Constructor from byte array.
func NewGXx509Certificate(data []byte) (*GXx509Certificate, error) {
	ret := &GXx509Certificate{}
	err := ret.Init(data)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// Loaded x509Certificate as raw data.
func (g *GXx509Certificate) RawData() []byte {
	return g.rawData
}

// Raw Issuer in ASN1 format.
func (g *GXx509Certificate) IssuerRaw() []byte {
	return g.issuerRaw
}

// Version.
// Version is read-only because DLMS supports only v3.
func (g *GXx509Certificate) Version() enums.CertificateVersion {
	return g.version
}

func (g *GXx509Certificate) Encoded() ([]byte, error) {
	if g.rawData != nil {
		return g.rawData, nil
	}
	ret, err := HashAlgorithmToString(g.SignatureAlgorithm)
	if err != nil {
		return nil, err
	}
	tmp := []any{*NewGXAsn1ObjectIdentifier(ret)}
	l, err := g.GetDataList()
	if err != nil {
		return nil, err
	}
	s, err := NewGXBitString(g.Signature, 0)
	if err != nil {
		return nil, err
	}
	list := []any{l, tmp, s}
	g.rawData, err = Asn1ToByteArray(list)
	if err != nil {
		return nil, err
	}
	return g.rawData, nil
}

func (g *GXx509Certificate) Init(data []byte) error {
	g.rawData = data
	ret, err := Asn1FromByteArray(data)
	if err != nil {
		return err
	}
	seq, ok := asn1Sequence(ret)
	if !ok {
		return errors.New("Invalid Certificate Version.")
	}
	if len(*seq) != 3 {
		return errors.New("Invalid Certificate Version. Wrong number of elements in sequence.")
	}
	if _, ok := asn1Sequence((*seq)[0]); !ok {
		type_, err := Asn1GetCertificateType(data, *seq)
		if err != nil {
			return err
		}
		switch type_ {
		case enums.PkcsTypePkcs8:
			return errors.New("Invalid Certificate. This is PKCS 8 private key, not x509 certificate.")
		case enums.PkcsTypePkcs10:
			return errors.New("Invalid Certificate. This is PKCS 10 certification requests, not x509 certificate.")
		}
		return errors.New("Invalid Certificate Version.")
	}
	reqInfo, _ := asn1Sequence((*seq)[0])
	if reqInfo == nil {
		return errors.New("Invalid Certificate Version.")
	}
	if _, ok := (*reqInfo)[0].(*GXAsn1Integer); ok {
		return errors.New("Invalid Certificate. DLMS certificate version number must be integer.")
	}
	ctx, ok := (*reqInfo)[0].(*GXAsn1Context)
	if !ok || len(ctx.Items) == 0 {
		return errors.New("Invalid Certificate. DLMS certificate version number must be integer.")
	}
	switch v := ctx.Items[0].(type) {
	case int8:
		g.version = enums.CertificateVersion(v)
	case int16:
		g.version = enums.CertificateVersion(v)
	case int32:
		g.version = enums.CertificateVersion(v)
	case int:
		g.version = enums.CertificateVersion(v)
	case *GXAsn1Integer:
		val, err := v.ToInt()
		if err != nil {
			return err
		}
		g.version = enums.CertificateVersion(val)
	default:
		return errors.New("Invalid Certificate. DLMS certificate version number must be integer.")
	}
	switch v := (*reqInfo)[1].(type) {
	case int8:
		g.SerialNumber = big.NewInt(int64(v))
	case int16:
		g.SerialNumber = big.NewInt(int64(v))
	case int32:
		g.SerialNumber = big.NewInt(int64(v))
	case int64:
		g.SerialNumber = big.NewInt(int64(v))
	case int:
		g.SerialNumber = big.NewInt(int64(v))
	case *GXAsn1Integer:
		g.SerialNumber = v.ToBigInteger()
	default:
		return errors.New("Invalid serial number.")
	}
	signSeq, ok := asn1Sequence((*reqInfo)[2])
	if !ok || len(*signSeq) == 0 {
		return errors.New("Invalid signature algorithm.")
	}
	signOID, ok := asn1OidString((*signSeq)[0])
	if !ok {
		return errors.New("Invalid signature algorithm.")
	}
	g.SignatureAlgorithm = HashAlgorithmFromString(signOID)
	if g.SignatureAlgorithm != enums.HashAlgorithmSha256WithEcdsa &&
		g.SignatureAlgorithm != enums.HashAlgorithmSha384WithEcdsa {
		return errors.New("DLMS certificate must be signed with ECDSA with SHA256 or SHA384.")
	}
	if len(*signSeq) > 1 {
		g.SignatureParameters = (*signSeq)[1]
	}
	g.issuerRaw, err = Asn1ToByteArray((*reqInfo)[3])
	if err != nil {
		return err
	}
	issuerSeq, ok := asn1Sequence((*reqInfo)[3])
	if !ok {
		return errors.New("Invalid issuer.")
	}
	g.Issuer = Asn1GetSubject(issuerSeq)
	validitySeq, ok := asn1Sequence((*reqInfo)[4])
	if !ok || len(*validitySeq) < 2 {
		return errors.New("Invalid validity.")
	}
	if t, ok := (*validitySeq)[0].(time.Time); ok {
		g.ValidFrom = t
	} else {
		return errors.New("ValidFrom")
	}
	if t, ok := (*validitySeq)[1].(time.Time); ok {
		g.ValidTo = t
	} else {
		return errors.New("ValidTo")
	}
	subjectSeq, ok := asn1Sequence((*reqInfo)[5])
	if !ok {
		return errors.New("Invalid subject.")
	}
	g.Subject = Asn1GetSubject(subjectSeq)
	cn, err := X509NameToString(enums.X509NameCN)
	if err != nil {
		return err
	}
	subjectPKInfo, ok := asn1Sequence((*reqInfo)[6])
	if !ok || len(*subjectPKInfo) < 2 {
		return errors.New("Invalid subject public key info.")
	}
	bs, ok := (*subjectPKInfo)[1].(*GXBitString)
	if !ok {
		return errors.New("Invalid subject public key.")
	}
	g.PublicKey, err = PublicKeyFromRawBytes(bs.Value())
	if err != nil {
		return err
	}
	if err := EcdsaValidate(g.PublicKey); err != nil {
		return err
	}

	basicConstraintsExists := false
	if len(*reqInfo) > 7 {
		ctx, ok := (*reqInfo)[7].(*GXAsn1Context)
		if ok && len(ctx.Items) > 0 {
			exts, ok := asn1Sequence(ctx.Items[0])
			if ok {
				for _, extItem := range *exts {
					s, ok := asn1Sequence(extItem)
					if !ok || len(*s) < 2 {
						continue
					}
					id, ok := asn1ObjectIdentifier((*s)[0])
					if !ok {
						continue
					}
					value := (*s)[1]
					switch X509CertificateTypeFromString(id.String()) {
					case enums.X509CertificateTypeSubjectKeyIdentifier:
						if b, ok := value.([]byte); ok {
							g.SubjectKeyIdentifier = b
						}
					case enums.X509CertificateTypeSubjectAlternativeName:
						alt, ok := asn1Sequence(value)
						if !ok {
							return errors.New("Invalid subject alternative name.")
						}
						var sb strings.Builder
						for _, it := range *alt {
							ctx, ok := it.(*GXAsn1Context)
							if !ok || len(ctx.Items) == 0 {
								continue
							}
							if b, ok := ctx.Items[0].([]byte); ok {
								if sb.Len() != 0 {
									sb.WriteString(", ")
								}
								sb.WriteString("DNS:")
								sb.Write(b)
							}
						}
						g.SubjectAlternativeName = sb.String()
					case enums.X509CertificateTypeAuthorityKeyIdentifier:
						aks, ok := asn1Sequence(value)
						if !ok {
							return errors.New("Invalid authority key identifier.")
						}
						for _, it := range *aks {
							ctx, ok := it.(*GXAsn1Context)
							if !ok || len(ctx.Items) == 0 {
								continue
							}
							switch ctx.Index {
							case 0:
								if b, ok := ctx.Items[0].([]byte); ok {
									g.AuthorityKeyIdentifier = b
								}
							case 1:
								issuerCtx, ok := ctx.Items[0].(*GXAsn1Context)
								if !ok || len(issuerCtx.Items) == 0 {
									continue
								}
								issuerSeq, ok := asn1Sequence(issuerCtx.Items[0])
								if !ok {
									continue
								}
								var sb strings.Builder
								for _, it2 := range *issuerSeq {
									var key any
									var val any
									switch p := it2.(type) {
									case GXKeyValuePair[any, any]:
										key = p.Key
										val = p.Value
									case *GXKeyValuePair[any, any]:
										key = p.Key
										val = p.Value
									default:
										continue
									}
									if sb.Len() != 0 {
										sb.WriteString(", ")
									}
									if oid, ok := asn1OidString(key); ok {
										name := X509NameFromString(oid)
										sb.WriteString(name.String())
									} else {
										sb.WriteString(fmt.Sprintf("%v", key))
									}
									sb.WriteString("=")
									sb.WriteString(fmt.Sprintf("%v", val))
								}
								g.AuthorityCertIssuer = sb.String()
							case 2:
								if b, ok := ctx.Items[0].([]byte); ok {
									tmp := make([]byte, len(b))
									copy(tmp, b)
									buffer.Reverse(tmp)
									g.AuthorityCertificationSerialNumber = *new(big.Int).SetBytes(tmp)
								}
							default:
								return fmt.Errorf("Invalid context.%d", ctx.Index)
							}
						}
					case enums.X509CertificateTypeKeyUsage:
						if bs, ok := value.(*GXBitString); ok {
							g.KeyUsage = enums.KeyUsage(bs.ToInteger())
						} else if _, ok := value.(bool); ok {
							if len(*s) < 3 {
								return errors.New("Invalid key usage.")
							}
							if bs, ok := (*s)[2].(*GXBitString); ok {
								g.KeyUsage = enums.KeyUsage(bs.ToInteger())
							} else {
								return errors.New("Invalid key usage.")
							}
						} else {
							return errors.New("Invalid key usage.")
						}
					case enums.X509CertificateTypeExtendedKeyUsage:
						ext, ok := asn1Sequence(value)
						if !ok {
							return errors.New("Invalid extended key usage.")
						}
						for _, it := range *ext {
							oid, ok := asn1OidString(it)
							if !ok {
								return errors.New("Invalid extended key usage.")
							}
							switch oid {
							case "1.3.6.1.5.5.7.3.1":
								g.ExtendedKeyUsage |= enums.ExtendedKeyUsageServerAuth
							case "1.3.6.1.5.5.7.3.2":
								g.ExtendedKeyUsage |= enums.ExtendedKeyUsageClientAuth
							default:
								return errors.New("Invalid extended key usage.")
							}
						}
					case enums.X509CertificateTypeBasicConstraints:
						basicConstraintsExists = true
						if seq, ok := asn1Sequence(value); ok {
							if len(*seq) != 0 {
								if b, ok := (*seq)[0].(bool); ok {
									g.BasicConstraints = b
								}
							}
						} else if b, ok := value.(bool); ok {
							g.BasicConstraints = b
						} else {
							return errors.New("Invalid key usage.")
						}
					default:
						// Unknown extension.
					}
				}
			}
		}
	}
	if !basicConstraintsExists {
		commonNameFound := false
		for _, it := range *subjectSeq {
			var key any
			var val any
			switch p := it.(type) {
			case GXKeyValuePair[any, any]:
				key = p.Key
				val = p.Value
			case *GXKeyValuePair[any, any]:
				key = p.Key
				val = p.Value
			default:
				continue
			}
			keyStr, ok := asn1OidString(key)
			if !ok {
				continue
			}
			if keyStr == cn {
				value := fmt.Sprintf("%v", val)
				if len(value) != 16 {
					return errors.New("System title is not included in Common Name.")
				}
				commonNameFound = true
				break
			}
		}
		if !commonNameFound {
			return errors.New("Common name doesn't exist.")
		}
	}
	if g.KeyUsage == enums.KeyUsageNone {
		return errors.New("Key usage not present. It's mandotory.")
	}
	if (g.KeyUsage&(enums.KeyUsageKeyCertSign|enums.KeyUsageCrlSign)) != 0 && !basicConstraintsExists {
		return errors.New("Basic Constraints value not present. It's mandotory.")
	}
	if g.KeyUsage == (enums.KeyUsageDigitalSignature|enums.KeyUsageKeyAgreement) &&
		g.ExtendedKeyUsage == enums.ExtendedKeyUsageNone {
		return errors.New("Extended key usage not present. It's mandotory for TLS.")
	}
	if g.ExtendedKeyUsage != enums.ExtendedKeyUsageNone &&
		g.KeyUsage != (enums.KeyUsageDigitalSignature|enums.KeyUsageKeyAgreement) {
		return errors.New("Extended key usage present. It's used only for TLS.")
	}
	pubAlgSeq, ok := asn1Sequence((*seq)[1])
	if !ok || len(*pubAlgSeq) == 0 {
		return errors.New("Invalid public key algorithm.")
	}
	pubAlgOID, ok := asn1OidString((*pubAlgSeq)[0])
	if !ok {
		return errors.New("Invalid public key algorithm.")
	}
	g.PublicKeyAlgorithm = HashAlgorithmFromString(pubAlgOID)
	if g.PublicKeyAlgorithm != enums.HashAlgorithmSha256WithEcdsa &&
		g.PublicKeyAlgorithm != enums.HashAlgorithmSha384WithEcdsa {
		return errors.New("DLMS certificate must be signed with ECDSA with SHA256 or SHA384.")
	}
	if len(*pubAlgSeq) > 1 {
		g.PublicKeyParameters = (*pubAlgSeq)[1]
	}
	sig, ok := (*seq)[2].(*GXBitString)
	if !ok {
		return errors.New("Invalid signature.")
	}
	g.Signature = sig.Value()
	return nil
}

func asn1Sequence(value any) (*GXAsn1Sequence, bool) {
	switch v := value.(type) {
	case *GXAsn1Sequence:
		return v, true
	case GXAsn1Sequence:
		return &v, true
	default:
		return nil, false
	}
}

func asn1ObjectIdentifier(value any) (*GXAsn1ObjectIdentifier, bool) {
	switch v := value.(type) {
	case *GXAsn1ObjectIdentifier:
		return v, true
	case GXAsn1ObjectIdentifier:
		return &v, true
	default:
		return nil, false
	}
}

func asn1OidString(value any) (string, bool) {
	switch v := value.(type) {
	case string:
		return v, true
	case *GXAsn1ObjectIdentifier:
		return v.String(), true
	case GXAsn1ObjectIdentifier:
		return v.String(), true
	default:
		return "", false
	}
}

func (g *GXx509Certificate) GetDataList() ([]any, error) {
	if g.Issuer == "" {
		return nil, errors.New("Issuer is empty.")
	}
	if g.Subject == "" {
		return nil, errors.New("Subject is empty.")
	}
	encodeSubject := func(value string) (GXAsn1Sequence, error) {
		pairs, err := Asn1EncodeSubject(value)
		if err != nil {
			return nil, err
		}
		seq := GXAsn1Sequence{}
		for _, pair := range pairs {
			seq = append(seq, NewGXKeyValuePair[any, any](pair.Key, pair.Value))
		}
		return seq, nil
	}

	oid, err := HashAlgorithmToString(g.SignatureAlgorithm)
	if err != nil {
		return nil, err
	}
	a := NewGXAsn1ObjectIdentifier(oid)
	p := NewGXAsn1Context()
	p.Items = append(p.Items, int8(g.version))
	extensions := GXAsn1Sequence{}

	if len(g.SubjectKeyIdentifier) != 0 {
		s1 := GXAsn1Sequence{}
		oid, err := X509CertificateTypeToString(enums.X509CertificateTypeSubjectKeyIdentifier)
		if err != nil {
			return nil, err
		}
		s1 = append(s1, NewGXAsn1ObjectIdentifier(oid))
		bb := GXByteBuffer{}
		if err := bb.SetUint8(byte(constants.BerTypeOctetString)); err != nil {
			return nil, err
		}
		if err := SetObjectCount(len(g.SubjectKeyIdentifier), &bb); err != nil {
			return nil, err
		}
		if err := bb.Set(g.SubjectKeyIdentifier); err != nil {
			return nil, err
		}
		s1 = append(s1, bb.Array())
		extensions = append(extensions, s1)
	}

	if len(g.AuthorityKeyIdentifier) != 0 || g.AuthorityCertIssuer != "" || g.AuthorityCertificationSerialNumber.BitLen() != 0 {
		s1 := GXAsn1Sequence{}
		oid, err := X509CertificateTypeToString(enums.X509CertificateTypeAuthorityKeyIdentifier)
		if err != nil {
			return nil, err
		}
		s1 = append(s1, NewGXAsn1ObjectIdentifier(oid))
		c1 := GXAsn1Sequence{}
		if len(g.AuthorityKeyIdentifier) != 0 {
			c4 := NewGXAsn1Context()
			c4.Constructed = false
			c4.Index = 0
			c4.Items = append(c4.Items, g.AuthorityKeyIdentifier)
			c1 = append(c1, c4)
		}
		if g.AuthorityCertIssuer != "" {
			c2 := NewGXAsn1Context()
			c2.Index = 1
			c3 := NewGXAsn1Context()
			c3.Index = 4
			issuer, err := encodeSubject(g.AuthorityCertIssuer)
			if err != nil {
				return nil, err
			}
			c3.Items = append(c3.Items, issuer)
			c2.Items = append(c2.Items, c3)
			c1 = append(c1, c2)
		}
		if g.AuthorityCertificationSerialNumber.BitLen() != 0 {
			c4 := NewGXAsn1Context()
			c4.Constructed = false
			c4.Index = 2
			tmp := g.AuthorityCertificationSerialNumber.Bytes()
			if len(tmp) != 0 {
				tmp2 := make([]byte, len(tmp))
				copy(tmp2, tmp)
				buffer.Reverse(tmp2)
				c4.Items = append(c4.Items, tmp2)
			}
			c1 = append(c1, c4)
		}
		if len(c1) != 0 {
			ext, err := Asn1ToByteArray(&c1)
			if err != nil {
				return nil, err
			}
			s1 = append(s1, ext)
		}
		extensions = append(extensions, s1)
	}

	// BasicConstraints
	s1 := GXAsn1Sequence{}
	oid, err = X509CertificateTypeToString(enums.X509CertificateTypeBasicConstraints)
	if err != nil {
		return nil, err
	}
	s1 = append(s1, NewGXAsn1ObjectIdentifier(oid))
	seq := GXAsn1Sequence{}
	if g.BasicConstraints {
		// BasicConstraints is critical if it exists.
		s1 = append(s1, g.BasicConstraints)
	} else if g.KeyUsage == enums.KeyUsageNone {
		return nil, errors.New("Key usage not present.")
	}
	ext, err := Asn1ToByteArray(&seq)
	if err != nil {
		return nil, err
	}
	s1 = append(s1, ext)
	extensions = append(extensions, s1)

	// KeyUsage
	s1 = GXAsn1Sequence{}
	oid, err = X509CertificateTypeToString(enums.X509CertificateTypeKeyUsage)
	if err != nil {
		return nil, err
	}
	s1 = append(s1, NewGXAsn1ObjectIdentifier(oid))
	value := byte(0)
	min := 255
	keyUsage := SwapBits(byte(g.KeyUsage))
	keyUsages := []enums.KeyUsage{
		enums.KeyUsageDigitalSignature,
		enums.KeyUsageNonRepudiation,
		enums.KeyUsageKeyEncipherment,
		enums.KeyUsageDataEncipherment,
		enums.KeyUsageKeyAgreement,
		enums.KeyUsageKeyCertSign,
		enums.KeyUsageCrlSign,
		enums.KeyUsageEncipherOnly,
		enums.KeyUsageDecipherOnly,
	}
	for _, it := range keyUsages {
		if (byte(it) & keyUsage) != 0 {
			val := byte(it)
			value |= val
			if int(val) < min {
				min = int(val)
			}
		}
	}
	ignore := 0
	for tmpMin := min; (tmpMin >> 1) != 0; {
		tmpMin >>= 1
		ignore++
	}
	bs, err := NewGXBitString([]byte{value}, ignore%8)
	tmp, err := Asn1ToByteArray(bs)
	if err != nil {
		return nil, err
	}
	s1 = append(s1, tmp)
	extensions = append(extensions, s1)

	// ExtendedKeyUsage
	if g.ExtendedKeyUsage != enums.ExtendedKeyUsageNone {
		s1 = GXAsn1Sequence{}
		oid, err = X509CertificateTypeToString(enums.X509CertificateTypeExtendedKeyUsage)
		if err != nil {
			return nil, err
		}
		s1 = append(s1, NewGXAsn1ObjectIdentifier(oid))
		s2 := GXAsn1Sequence{}
		if (g.ExtendedKeyUsage & enums.ExtendedKeyUsageServerAuth) != 0 {
			s2 = append(s2, NewGXAsn1ObjectIdentifier("1.3.6.1.5.5.7.3.1"))
		}
		if (g.ExtendedKeyUsage & enums.ExtendedKeyUsageClientAuth) != 0 {
			s2 = append(s2, NewGXAsn1ObjectIdentifier("1.3.6.1.5.5.7.3.2"))
		}
		tmp, err := Asn1ToByteArray(&s2)
		if err != nil {
			return nil, err
		}
		s1 = append(s1, tmp)
		extensions = append(extensions, s1)
	}

	valid := GXAsn1Sequence{}
	valid = append(valid, g.ValidFrom)
	valid = append(valid, g.ValidTo)

	var alg *GXAsn1ObjectIdentifier
	if g.PublicKey.Scheme() == enums.EccP256 {
		alg = NewGXAsn1ObjectIdentifier("1.2.840.10045.3.1.7")
	} else {
		alg = NewGXAsn1ObjectIdentifier("1.3.132.0.34")
	}

	tmp3 := []any{NewGXAsn1ObjectIdentifier("1.2.840.10045.2.1"), alg}
	tmp4 := NewGXAsn1Context()
	tmp4.Index = 3
	tmp4.Items = append(tmp4.Items, extensions)
	bs, err = NewGXBitString(g.PublicKey.RawValue(), 0)
	if err != nil {
		return nil, err
	}
	tmp2 := []any{tmp3, bs}
	var p2 []any
	if g.SignatureParameters == nil {
		p2 = []any{a}
	} else {
		p2 = []any{a, g.SignatureParameters}
	}

	issuer, err := encodeSubject(g.Issuer)
	if err != nil {
		return nil, err
	}
	subject, err := encodeSubject(g.Subject)
	if err != nil {
		return nil, err
	}
	if g.SerialNumber == nil {
		return nil, errors.New("SerialNumber is nil.")
	}
	list := []any{p, NewGXAsn1IntegerFromBigInteger(*g.SerialNumber), p2, issuer, valid, subject, tmp2, tmp4}
	return list, nil
}

// GetFilePath returns the default file path.
func (g *GXx509Certificate) GetFilePath(cert *GXx509Certificate) (string, error) {
	var path string
	switch cert.KeyUsage {
	case enums.KeyUsageDigitalSignature:
		path = "D"
	case enums.KeyUsageKeyAgreement:
		path = "A"
	case enums.KeyUsageDigitalSignature | enums.KeyUsageKeyAgreement:
		path = "T"
	default:
		return "", errors.New("Unknown certificate type.")
	}
	ret, err := HexSystemTitleFromSubject(strings.TrimSpace(cert.Subject))
	if err != nil {
		return "", err
	}
	path = filepath.Join(path, ret+".pem")
	if g.PublicKey.Scheme() == enums.EccP256 {
		path = filepath.Join("Certificates", path)
	} else {
		path = filepath.Join("Certificates384", path)
	}
	return path, nil
}

// FromHexString returns the create x509Certificate from hex string.
//
// Parameters:
//
//	data: Hex string.
//
// Returns:
//
//	x509 certificate
func (g *GXx509Certificate) FromHexString(data string) GXx509Certificate {
	cert := GXx509Certificate{}
	cert.Init(buffer.HexToBytes(data))
	return cert
}

// x509CertificateFromPem returns the create x509Certificate from PEM string.
//
// Parameters:
//
//	data: PEM string.
//
// Returns:
//
//	x509 certificate
func X509CertificateFromPem(data string) (*GXx509Certificate, error) {
	data = strings.ReplaceAll(data, "\r\n", "\n")
	const START = "CERTIFICATE-----"
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
	return X509CertificateFromDer(data[0:end])
}

// x509CertificateFromDer returns the create x509Certificate from DER Base64 encoded string.
//
// Parameters:
//
//	der: Base64 DER string.
//
// Returns:
//
//	x509 certificate
func X509CertificateFromDer(der string) (*GXx509Certificate, error) {
	der = strings.ReplaceAll(der, "\r\n", "")
	der = strings.ReplaceAll(der, "\n", "")
	key, err := base64.StdEncoding.DecodeString(der)
	if err != nil {
		return nil, err
	}
	cert := GXx509Certificate{}
	err = cert.Init(key)
	if err != nil {
		return nil, err
	}
	return &cert, nil
}

// GetData returns the get data as byte array.
func (g *GXx509Certificate) GetData() []byte {
	ret, err := g.GetDataList()
	if err != nil {
		return nil
	}
	ret2, err := Asn1ToByteArray(ret)
	if err != nil {
		return nil
	}
	return ret2
}

// String returns the string representation of the x509 certificate.
func (g *GXx509Certificate) String() string {
	bb := strings.Builder{}
	switch g.ExtendedKeyUsage {
	case enums.ExtendedKeyUsageServerAuth:
		bb.WriteString("Server certificate")
	case enums.ExtendedKeyUsageClientAuth:
		bb.WriteString("Client certificate")
	case enums.ExtendedKeyUsageServerAuth | enums.ExtendedKeyUsageClientAuth:
		bb.WriteString("TLS certificate")
	}
	bb.WriteString("")
	bb.WriteString("Version: ")
	bb.WriteString(g.version.String())
	bb.WriteString("SerialNumber: ")
	bb.WriteString(fmt.Sprint(g.SerialNumber))
	bb.WriteString("Signature: ")
	bb.WriteString(g.SignatureAlgorithm.String())
	bb.WriteString(", OID = ")
	ret, err := HashAlgorithmToString(g.SignatureAlgorithm)
	if err != nil {
		log.Println("Error in String(): " + err.Error())
		return ""
	}
	bb.WriteString(ret)
	bb.WriteString("\n")
	bb.WriteString("Issuer: ")
	bb.WriteString(g.Issuer)
	bb.WriteString("\n")
	bb.WriteString("Validity: [From: ")
	bb.WriteString(g.ValidFrom.String())
	bb.WriteString(" GMT To: ")
	bb.WriteString(g.ValidTo.String())
	bb.WriteString(" GMT]\n")
	bb.WriteString("Subject: ")
	bb.WriteString(g.Subject)
	bb.WriteString("\n")
	bb.WriteString("Public Key Algorithm: ")
	bb.WriteString(g.PublicKeyAlgorithm.String())
	bb.WriteString("\n")
	bb.WriteString("Key: ")
	bb.WriteString(g.PublicKey.ToHex())
	bb.WriteString("\n")
	if g.PublicKey.Scheme() == enums.EccP256 {
		bb.WriteString("ASN1 OID: prime256v1\n")
		bb.WriteString("NIST CURVE: P-256")
	} else if g.PublicKey.Scheme() == enums.EccP384 {
		bb.WriteString("ASN1 OID: prime384v1\n")
		bb.WriteString("\n")
		bb.WriteString("NIST CURVE: P-384")
	}
	bb.WriteString("\n")
	bb.WriteString("Basic constraints: ")
	bb.WriteString(strconv.FormatBool(g.BasicConstraints))
	bb.WriteString("\n")
	bb.WriteString("ExtendedKeyUsage: ")
	ExtendedKeyUsage := ""
	if (g.ExtendedKeyUsage & enums.ExtendedKeyUsageServerAuth) != 0 {
		ExtendedKeyUsage += "ServerAuth "
	}
	if (g.ExtendedKeyUsage & enums.ExtendedKeyUsageClientAuth) != 0 {
		ExtendedKeyUsage += "ClientAuth "
	}
	bb.WriteString(ExtendedKeyUsage)
	bb.WriteString("\n")
	bb.WriteString("SubjectKeyIdentifier: ")
	bb.WriteString(buffer.ToHex(g.SubjectKeyIdentifier, true))
	bb.WriteString("\n")
	bb.WriteString("KeyUsage: ")
	bb.WriteString(g.KeyUsage.String())
	bb.WriteString("\n")
	bb.WriteString("Signature Algorithm: ")
	bb.WriteString(g.SignatureAlgorithm.String())
	bb.WriteString("\n")
	bb.WriteString("Signature: ")
	bb.WriteString(buffer.ToHex(g.Signature, false))
	bb.WriteString("\n")
	return bb.String()
}

// Load returns the load x509 certificate from the PEM file.
//
// Parameters:
//
//	path: File path.
//
// Returns:
//
//	Created GXx509Certificate object.
func X509CertificateLoad(path string) (*GXx509Certificate, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return X509CertificateFromPem(string(data))
}

// Save returns the x509 certificate to PEM file.
//
// Parameters:
//
//	path: File path.
func (g *GXx509Certificate) Save(path string) error {
	ret, err := g.ToPem()
	if err != nil {
		return err
	}
	return os.WriteFile(path, []byte(ret), 0644)
}

// ToPem returns the x509 certificate in PEM format.
//
// Returns:
//
//	Public key as in PEM string.
func (g *GXx509Certificate) ToPem() (string, error) {
	sb := strings.Builder{}
	if g.PublicKey == nil {
		return "", errors.New("Public or private key is not set.")
	}
	sb.WriteString("-----BEGIN CERTIFICATE-----\n")
	sb.WriteString(g.ToDer())
	sb.WriteString("-----END CERTIFICATE-----\n")
	return sb.String(), nil
}

// ToDer returns the x509 certificate in DER format.
//
// Returns:
//
//	Public key as in PEM string.
func (g *GXx509Certificate) ToDer() string {
	ret, err := g.Encoded()
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(ret)
}

// Equals determines whether the specified object is equal to the current GXx509Certificate instance.
func (g *GXx509Certificate) Equals(obj any) bool {
	if o, ok := obj.(GXx509Certificate); ok {
		return g.SerialNumber.Cmp(o.SerialNumber) == 0
	}
	return false
}

// IsCertified returns the test is x509 file certified by the certifier.
//
// Parameters:
//
//	certifier: Public key of the certifier.
//
// Returns:
//
//	True, if certifier has certified the certificate.
func (g *GXx509Certificate) IsCertified(certifier *GXPublicKey) (bool, error) {
	if certifier == nil {
		return false, errors.New("certifier")
	}
	// Get raw data
	tmp2 := GXByteBuffer{}
	encoded, error := g.Encoded()
	if error != nil {
		return false, error
	}
	err := tmp2.Set(encoded)
	if err != nil {
		return false, err
	}
	_, err = Asn1GetNext(&tmp2)
	if err != nil {
		return false, err
	}
	tmp2.SetSize(tmp2.Position())
	tmp2.SetPosition(1)
	GetObjectCount(&tmp2)
	e, err := NewGXEcdsaFromPublicKey(certifier)
	if err != nil {
		return false, err
	}
	v, err := Asn1FromByteArray(g.Signature)
	if err != nil {
		return false, err
	}
	tmp3 := v.(GXAsn1Sequence)
	bb := GXByteBuffer{}
	size := 0
	if g.SignatureAlgorithm == enums.HashAlgorithmSha256WithEcdsa {
		size = 32
	} else {
		size = 48
	}
	v2 := tmp3[0].(*GXAsn1Integer).Value()
	a := 0
	if len(v2) != size {
		a = 1
	}
	bb.SetAt(v2, a, size)

	v2 = tmp3[1].(*GXAsn1Integer).Value()
	a = 0
	if len(v2) != size {
		a = 1
	}
	bb.SetAt(v2, a, size)
	tmp4, err := tmp2.SubArray(tmp2.Position(), tmp2.Available())
	if err != nil {
		return false, err
	}
	return e.Verify(bb.Array(), tmp4)
}

// X509CertificateSearch returns the search x509 Certificate from the PEM file in given folder.
//
// Parameters:
//
//	folder: Folder to search.
//	type: Certificate type.
//
// Returns:
//
//	Created GXPkcs8 object.
func X509CertificateSearch(folder string, type_ enums.CertificateType, systemtitle []byte) []GXx509Certificate {
	var usage enums.KeyUsage
	if type_ == enums.CertificateTypeDigitalSignature {
		usage = enums.KeyUsageDigitalSignature
	} else if type_ == enums.CertificateTypeKeyAgreement {
		usage = enums.KeyUsageKeyAgreement
	} else if type_ == enums.CertificateTypeTLS {
		usage = enums.KeyUsageDigitalSignature | enums.KeyUsageKeyAgreement
	} else {
		usage = enums.KeyUsageNone
	}
	subject := Asn1SystemTitleToSubject(systemtitle)
	certificates := []GXx509Certificate{}
	entries, err := os.ReadDir(folder)
	if err != nil {
		log.Printf("Failed to read directory %s: %v", folder, err)
		return nil
	}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		full := filepath.Join(folder, e.Name())
		ext := strings.ToLower(filepath.Ext(full))
		if ext == ".pem" || ext == ".cer" {
			cert, err := X509CertificateLoad(full)
			if err != nil {
				log.Printf("Failed to load PKCS #8 certificate. %s: %v", full, err)
				continue
			}
			if (usage == enums.KeyUsageNone || cert.KeyUsage == usage) && strings.Contains(cert.Subject, subject) {
				certificates = append(certificates, *cert)
			}
		}
	}
	return certificates
}

// GetSystemTitle returns the system title from the certificate.
func (g *GXx509Certificate) GetSystemTitle() ([]byte, error) {
	if g.Subject == "" {
		return nil, nil
	}
	ret, err := HexSystemTitleFromSubject(g.Subject)
	if err != nil {
		return nil, err
	}
	return buffer.HexToBytes(ret), nil
}
