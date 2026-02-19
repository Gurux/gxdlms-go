package objects

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
	"errors"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/Gurux/gxcommon-go"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/buffer"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// Online help:
// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSSecuritySetup
type GXDLMSSecuritySetup struct {
	GXDLMSObject

	guek           []byte
	gbek           []byte
	gak            []byte
	securityPolicy enums.SecurityPolicy

	// Signing key of the server.
	signingKey *types.GXKeyValuePair[*types.GXPublicKey, *types.GXPrivateKey]

	// Key agreement key of the server.
	keyAgreementKey *types.GXKeyValuePair[*types.GXPublicKey, *types.GXPrivateKey]

	// TLS pair of the server.
	tlsKey *types.GXKeyValuePair[*types.GXPublicKey, *types.GXPrivateKey]

	serverCertificates types.GXx509CertificateCollection

	// Master key.
	Kek []byte

	// Security suite.
	SecuritySuite enums.SecuritySuite

	// Client system title.
	ClientSystemTitle []byte

	// Server system title.
	ServerSystemTitle []byte

	// Available certificates.
	Certificates GXDLMSCertificateCollection
}

// base returns the base GXDLMSObject of the object.
func (g *GXDLMSSecuritySetup) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

// Guek returns the block cipher key.
func (g *GXDLMSSecuritySetup) Guek() []byte {
	return g.guek
}

func validateKeyLength(suite enums.SecuritySuite, value []byte) error {
	if suite == enums.SecuritySuite2 && len(value) != 32 && len(value) != 0 {
		return fmt.Errorf("Invalid key length. Key length must be 32 bytes for suite 2.")
	}
	if suite != enums.SecuritySuite2 && len(value) != 16 && len(value) != 0 {
		return fmt.Errorf("Invalid key length. Key length must be 16 bytes for suite 0 or suite 1.")
	}
	return nil
}

// SetGuek sets the block cipher key.
func (g *GXDLMSSecuritySetup) SetGuek(value []byte) error {
	err := validateKeyLength(g.SecuritySuite, value)
	if err != nil {
		return err
	}
	g.guek = value
	return nil
}

// Gbek returns the broadcast block cipher key.
func (g *GXDLMSSecuritySetup) Gbek() []byte {
	return g.gbek
}

// SetGbek sets the broadcast block cipher key.
func (g *GXDLMSSecuritySetup) SetGbek(value []byte) error {
	err := validateKeyLength(g.SecuritySuite, value)
	if err != nil {
		return err
	}
	g.gbek = value
	return nil
}

// Gak returns the authentication key.
func (g *GXDLMSSecuritySetup) Gak() []byte {
	return g.gak
}

// SetGak sets the authentication key.
func (g *GXDLMSSecuritySetup) SetGak(value []byte) error {
	err := validateKeyLength(g.SecuritySuite, value)
	if err != nil {
		return err
	}
	g.gak = value
	return nil
}

// SecurityPolicy returns the security policy for Version 0 and 1.
func (g *GXDLMSSecuritySetup) SecurityPolicy() enums.SecurityPolicy {
	return g.securityPolicy
}

// SetSecurityPolicy sets the security policy for Version 0 and 1.
func (g *GXDLMSSecuritySetup) SetSecurityPolicy(value enums.SecurityPolicy) error {
	if g.Version == 0 {
		switch value {
		case enums.SecurityPolicyNone:
		case enums.SecurityPolicyAuthenticated:
		case enums.SecurityPolicyEncrypted:
		case enums.SecurityPolicyAuthenticatedEncrypted:
			break
		default:
			return fmt.Errorf("Invalid security policy value %v for Version 0.", value)
		}
	} else if g.Version == 1 {
		if (int(value) & 0x3) != 0 {
			return fmt.Errorf("Invalid security policy value %v for Version 1.", value)
		}
	}
	g.securityPolicy = value
	return nil
}

// Parameters:
//
//	security: Security level.
//
// Returns:
//
//	Integer value of security level.
func (g *GXDLMSSecuritySetup) getSecurityValue(security enums.Security) (int, error) {
	var value int
	switch security {
	case enums.SecurityNone:
		value = 0
	case enums.SecurityAuthentication:
		value = 1
	case enums.SecurityEncryption:
		value = 2
	case enums.SecurityAuthenticationEncryption:
		value = 3
	default:
		return 0, fmt.Errorf("Invalid Security enum.")
	}
	return value, nil
}

// VerifyIssuer returns the verify that issuer is in types.Asn1 format.
//
// Parameters:
//
//	issuer: Certificate issuer.
func (g *GXDLMSSecuritySetup) VerifyIssuer(issuer []byte) error {
	if len(issuer) == 0 {
		return gxcommon.ErrInvalidArgument
	}
	_, err := types.Asn1FromByteArray(issuer)
	return err
}

func certificateTypeToKeyUsage(type_ enums.CertificateType) enums.KeyUsage {
	var k enums.KeyUsage
	switch type_ {
	case enums.CertificateTypeDigitalSignature:
		k = enums.KeyUsageDigitalSignature
	case enums.CertificateTypeKeyAgreement:
		k = enums.KeyUsageKeyAgreement
	case enums.CertificateTypeTLS:
		k = enums.KeyUsageDigitalSignature | enums.KeyUsageKeyAgreement
	case enums.CertificateTypeOther:
		k = enums.KeyUsageCrlSign
	default:
		k = enums.KeyUsageNone
	}
	return k
}

// FindCertificateByEntity returns the find certificate using entity information.
//
// Parameters:
//
//	certificates: Certificate collection.
//	entity: Certificate entity type.
//	type_: Certificate type.
//	systemtitle: System title.
func (g *GXDLMSSecuritySetup) FindCertificateByEntity(certificates types.GXx509CertificateCollection,
	entity enums.CertificateEntity,
	type_ enums.CertificateType,
	systemtitle []byte) *types.GXx509Certificate {
	subject := types.Asn1SystemTitleToSubject(systemtitle)
	k := certificateTypeToKeyUsage(type_)
	for _, it := range certificates {
		if (it.KeyUsage&k) != 0 && strings.Contains(it.Subject, subject) {
			return it
		}
	}
	return nil
}

// FindCertificateBySerial returns the find certificate using serial information.
//
// Parameters:
//
//	certificates: Certificate collection.
//	serialNumber: Serial number.
//	issuer: Issuer.
func (g *GXDLMSSecuritySetup) FindCertificateBySerial(certificates types.GXx509CertificateCollection,
	serialNumber []byte,
	issuer string) *types.GXx509Certificate {
	for _, it := range certificates {
		if types.ToHex(it.SerialNumber.Bytes(), false) == types.ToHex(serialNumber, false) && it.Issuer == issuer {
			return it
		}
	}
	return nil
}

func getEcc(suite enums.SecuritySuite) enums.Ecc {
	if suite == enums.SecuritySuite1 {
		return enums.EccP256
	}
	return enums.EccP384
}

// Invoke returns the invokes method.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Invoke parameters.
func (g *GXDLMSSecuritySetup) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	var err error
	if g.SecuritySuite == enums.SecuritySuite0 && e.Index > 3 {
		return nil, fmt.Errorf("Invalid security suite version.")
	}
	switch e.Index {
	case 1:
		g.securityPolicy = enums.SecurityPolicy(e.Parameters.(int))
	case 2:
		err = g.keyTransfer(settings, e)
	case 3:
		return g.invokeKeyAgreement(settings, e)
	case 4:
		err = g.generateKeyPair(e)
	case 5:
		return g.generateCertificateRequest(settings, e)
	case 6:
		g.importCertificate(settings, e)
	case 7:
		return g.exportCertificate(e)
	case 8:
		err = g.removeCertificate(e)
	default:
		e.Error = enums.ErrorCodeInconsistentClass
	}
	if err != nil {
		e.Error = enums.ErrorCodeInconsistentClass
	}
	// Return standard reply.
	return nil, err
}

func (g *GXDLMSSecuritySetup) removeCertificate(e *internal.ValueEventArgs) error {
	var err error
	tmp := e.Parameters.([]any)
	type_ := tmp[0].(byte)
	tmp = tmp[1].([]any)
	var cert *types.GXx509Certificate
	switch type_ {
	case 0:
		cert = g.FindCertificateByEntity(g.serverCertificates, enums.CertificateEntity(tmp[1].(int)), enums.CertificateType(tmp[2].(int)), tmp[3].([]byte))
	case 1:
		//TODO:  buffer.reverse(tmp[0].([]byte))
		sn := new(big.Int).SetBytes(tmp[0].([]byte))
		cert = g.serverCertificates.FindBySerial(sn, string(tmp[1].([]byte)))
	}
	if cert == nil {
		e.Error = enums.ErrorCodeInconsistentClass
	} else {
		internal.Remove(g.serverCertificates, cert)
	}
	return err
}

func (g *GXDLMSSecuritySetup) importCertificate(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	cert, err := types.NewGXx509Certificate(e.Parameters.([]byte))
	if err != nil {
		return err
	}
	st := g.ServerSystemTitle
	if st == nil {
		st = settings.Cipher.SystemTitle()
	}
	serverSubject := types.Asn1SystemTitleToSubject(st)
	// If server certification is added.
	isServerCert := strings.Contains(cert.Subject, serverSubject)
	if cert.KeyUsage != enums.KeyUsageKeyAgreement && cert.KeyUsage != enums.KeyUsageDigitalSignature && cert.KeyUsage != (enums.KeyUsageKeyAgreement|enums.KeyUsageDigitalSignature) {
		e.Error = enums.ErrorCodeInconsistentClass
	} else {
		// Remove old certificate if it exists.
		list := g.serverCertificates.GetCertificates(cert.KeyUsage)
		for _, it := range list {
			isServer := strings.Contains(it.Subject, serverSubject)
			if isServer == isServerCert {
				internal.Remove(g.serverCertificates, it)
			}
		}
	}
	g.serverCertificates = append(g.serverCertificates, cert)
	log.Printf("New certificate imported: %s.\n", cert.SerialNumber.String())
	return nil
}

func (g *GXDLMSSecuritySetup) exportCertificate(e *internal.ValueEventArgs) ([]byte, error) {
	tmp := e.Parameters.([]any)
	type_ := tmp[0].(byte)
	tmp = tmp[1].([]any)
	var cert *types.GXx509Certificate
	switch type_ {
	case 0:
		cert = g.FindCertificateByEntity(g.serverCertificates, enums.CertificateEntity(tmp[0].(int)), enums.CertificateType(tmp[1].(int)), tmp[2].([]byte))
	case 1:
		issuer, err := getStringFromAsn1(tmp[1].([]byte))
		if err != nil {
			return nil, err
		}
		sn, err := types.Asn1FromByteArray(tmp[0].([]byte))
		if err != nil {
			return nil, err
		}
		tmp := sn.(*types.GXAsn1Integer)
		cert = g.serverCertificates.FindBySerial(tmp.ToBigInteger(), issuer)
	}
	if cert == nil {
		e.Error = enums.ErrorCodeInconsistentClass
	} else {
		log.Printf("Export certificate: %s", cert.SerialNumber.String())
		return cert.Encoded()
	}
	return nil, nil
}

func (g *GXDLMSSecuritySetup) generateKeyPair(e *internal.ValueEventArgs) error {
	if g.SecuritySuite == enums.SecuritySuite0 {
		return fmt.Errorf("Invalid security suite version.")
	}
	key := enums.CertificateType(e.Parameters.(int))
	value, err := types.GXEcdsaGenerateKeyPair(getEcc(g.SecuritySuite))
	if err != nil {
		return err
	}
	switch key {
	case enums.CertificateTypeDigitalSignature:
		g.signingKey = value
	case enums.CertificateTypeKeyAgreement:
		g.keyAgreementKey = value
	default:
		g.tlsKey = value
	}
	return nil
}

func (g *GXDLMSSecuritySetup) generateCertificateRequest(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	key := enums.CertificateType(e.Parameters.(int))
	var kp *types.GXKeyValuePair[*types.GXPublicKey, *types.GXPrivateKey]
	st := g.ServerSystemTitle
	if st == nil {
		st = settings.Cipher.SystemTitle()
	}
	switch key {
	case enums.CertificateTypeDigitalSignature:
		kp = g.signingKey
	case enums.CertificateTypeKeyAgreement:
		kp = g.keyAgreementKey
	case enums.CertificateTypeTLS:
		kp = g.tlsKey
	default:
	}
	if kp.Key != nil {
		pkc10, err := types.Pkcs10CreateCertificateSigningRequest(kp, types.Asn1SystemTitleToSubject(st))
		if err != nil {
			return nil, err
		}
		return pkc10.Encoded()
	} else {
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil, nil
}

func (g *GXDLMSSecuritySetup) keyTransfer(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	for _, item := range e.Parameters.([]any) {
		type_ := enums.GlobalKeyType(item.([]any)[0].(int))
		data := item.([]any)[1].([]byte)
		ret, err := internal.Decrypt(settings.Kek, data)
		if err != nil {
			return err
		}
		switch type_ {
		case enums.GlobalKeyTypeUnicastEncryption:
			g.guek = ret
		case enums.GlobalKeyTypeBroadcastEncryption:
			g.gbek = ret
		case enums.GlobalKeyTypeAuthentication:
			g.gak = ret
		case enums.GlobalKeyTypeKek:
			g.Kek = ret
		default:
			e.Error = enums.ErrorCodeInconsistentClass
		}
	}
	return nil
}

func (g *GXDLMSSecuritySetup) invokeKeyAgreement(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	var err error
	tmp := e.Parameters.([]any)[0].([]any)
	keyId := tmp[0].(byte)
	if keyId != 0 {
		e.Error = enums.ErrorCodeInconsistentClass
	} else {
		data := tmp[0].([]byte)
		// ephemeral public key
		data2 := types.NewGXByteBufferWithCapacity(65)
		err = data2.SetUint8(keyId)
		if err != nil {
			return nil, err
		}
		err = data2.SetAt(data, 0, 64)
		if err != nil {
			return nil, err
		}
		sign := types.GXByteBuffer{}
		err = sign.SetAt(data, 64, 64)
		if err != nil {
			return nil, err
		}
		var pk *types.GXPublicKey
		subject := types.Asn1SystemTitleToSubject(settings.SourceSystemTitle())
		for _, it := range settings.Cipher.Certificates() {
			if (it.KeyUsage&enums.KeyUsageDigitalSignature) != 0 && it.Subject == subject {
				pk = it.PublicKey
				break
			}
		}
		if pk == nil {
			e.Error = enums.ErrorCodeInconsistentClass
			settings.TargetEphemeralKey = nil
		} else {
			ret, err := data2.SubArray(1, 64)
			if err != nil {
				return nil, err
			}
			settings.TargetEphemeralKey, err = types.PublicKeyFromRawBytes(ret)
			if err != nil {
				return nil, err
			}
			// Generate ephemeral keys.
			if settings.Cipher.EphemeralKeyPair() == nil {
				ret, err := types.GXEcdsaGenerateKeyPair(getEcc(g.SecuritySuite))
				if err == nil {
					settings.Cipher.SetEphemeralKeyPair(ret)
				}
			}
		}
	}
	return nil, err
}

// GetAttributeIndexToRead returns the collection of attributes to read.
// If attribute is static and already read or device is returned HW error it is not returned.
//
// Parameters:
//
//	all: All items are returned even if they are read already.
//
// Returns:
//
//	Collection of attributes to read.
func (g *GXDLMSSecuritySetup) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	// LN is static and read only once.
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	// SecurityPolicy
	if all || g.CanRead(2) {
		attributes = append(attributes, 2)
	}
	// SecuritySuite
	if all || g.CanRead(3) {
		attributes = append(attributes, 3)
	}
	// ClientSystemTitle
	if all || g.CanRead(4) {
		attributes = append(attributes, 4)
	}
	// ServerSystemTitle
	if all || g.CanRead(5) {
		attributes = append(attributes, 5)
	}
	if g.Version > 0 {
		// Certificates
		if all || g.CanRead(6) {
			attributes = append(attributes, 6)
		}
	}
	return attributes
}

// GetNames returns the names of attribute indexes.
func (g *GXDLMSSecuritySetup) GetNames() []string {
	names := []string{"Logical Name", "Security Policy", "Security Suite", "Client System Title", "Server System Title"}
	if g.Version > 0 {
		names = append(names, "Certificates")
	}
	return names
}

// GetMethodNames returns the names of method indexes.
func (g *GXDLMSSecuritySetup) GetMethodNames() []string {
	names := []string{"Security activate", "Key transfer"}
	if g.Version > 0 {
		names = append(names, "Key agreement")
		names = append(names, "Generate key pair")
		names = append(names, "Generate certificate request")
		names = append(names, "Import certificate")
		names = append(names, "Export certificate")
		names = append(names, "Remove certificate")
	}
	return names
}

// GetAttributeCount returns the amount of attributes.
//
// Returns:
//
//	Count of attributes.
func (g *GXDLMSSecuritySetup) GetAttributeCount() int {
	if g.Version == 0 {
		return 5
	}
	return 6
}

// GetMethodCount returns the amount of methods.
func (g *GXDLMSSecuritySetup) GetMethodCount() int {
	if g.Version == 0 {
		return 2
	}
	return 8
}

// getSertificates returns the get sertificates as byte buffer.
func (g *GXDLMSSecuritySetup) getSertificates() ([]byte, error) {
	bb := types.GXByteBuffer{}
	err := bb.SetUint8(uint8(enums.DataTypeArray))
	if err != nil {
		return nil, err
	}
	types.SetObjectCount(len(g.serverCertificates), &bb)
	for _, it := range g.serverCertificates {
		ret, err := types.Asn1FromByteArray(it.RawData())
		if err != nil {
			return nil, err
		}
		seq := ret.(types.GXAsn1Sequence)
		err = bb.SetUint8(uint8(enums.DataTypeStructure))
		if err != nil {
			return nil, err
		}
		types.SetObjectCount(6, &bb)
		err = bb.SetUint8(uint8(enums.DataTypeEnum))
		if err != nil {
			return nil, err
		}
		if it.BasicConstraints {
			err = bb.SetUint8(uint8(enums.CertificateEntityCertificationAuthority))
			if err != nil {
				return nil, err
			}
		} else if strings.Contains((it.Subject), types.Asn1SystemTitleToSubject(g.ServerSystemTitle)) {
			err = bb.SetUint8(uint8(enums.CertificateEntityServer))
			if err != nil {
				return nil, err
			}
		} else {
			err = bb.SetUint8(uint8(enums.CertificateEntityClient))
			if err != nil {
				return nil, err
			}
		}
		err = bb.SetUint8(uint8(enums.DataTypeEnum))
		if err != nil {
			return nil, err
		}
		switch it.KeyUsage {
		case enums.KeyUsageDigitalSignature | enums.KeyUsageKeyAgreement:
			err = bb.SetUint8(uint8(enums.CertificateTypeTLS))
			if err != nil {
				return nil, err
			}
		case enums.KeyUsageDigitalSignature:
			err = bb.SetUint8(uint8(enums.CertificateTypeDigitalSignature))
			if err != nil {
				return nil, err
			}
		case enums.KeyUsageKeyAgreement:
			err = bb.SetUint8(uint8(enums.CertificateTypeKeyAgreement))
			if err != nil {
				return nil, err
			}
		default:
			err = bb.SetUint8(uint8(enums.CertificateTypeOther))
			if err != nil {
				return nil, err
			}
		}
		// Get serial number.
		reqInfo := seq[0].(types.GXAsn1Sequence)
		tmp, err := types.Asn1ToByteArray(reqInfo[1])
		if err != nil {
			return nil, err
		}
		err = bb.SetUint8(uint8(enums.DataTypeOctetString))
		if err != nil {
			return nil, err
		}
		err = bb.SetUint8(uint8(len(tmp)))
		if err != nil {
			return nil, err
		}
		err = bb.Set(tmp)
		if err != nil {
			return nil, err
		}
		tmp, err = types.Asn1ToByteArray(reqInfo[3])
		if err != nil {
			return nil, err
		}
		err = bb.SetUint8(uint8(enums.DataTypeOctetString))
		if err != nil {
			return nil, err
		}
		err = bb.SetUint8(uint8(len(tmp)))
		if err != nil {
			return nil, err
		}
		err = bb.Set(tmp)
		if err != nil {
			return nil, err
		}
		tmp, err = types.Asn1ToByteArray(reqInfo[5])
		if err != nil {
			return nil, err
		}
		err = bb.SetUint8(uint8(enums.DataTypeOctetString))
		if err != nil {
			return nil, err
		}
		err = bb.SetUint8(uint8(len(tmp)))
		if err != nil {
			return nil, err
		}
		err = bb.Set(tmp)
		if err != nil {
			return nil, err
		}
		tmp = make([]byte, 0)
		if len(reqInfo) > 7 {
			for _, s := range (reqInfo[7]).(*types.GXAsn1Context).Items[0].(types.GXAsn1Sequence) {
				tmp2 := s.(types.GXAsn1Sequence)
				id := tmp2[0].(*types.GXAsn1ObjectIdentifier)
				t := types.X509CertificateTypeFromString(id.String())
				if t == enums.X509CertificateTypeSubjectAlternativeName {
					tmp, err = types.Asn1ToByteArray(s)
					if err != nil {
						return nil, err
					}
					break
				}
			}
		}
		err = bb.SetUint8(enums.DataTypeOctetString)
		if err != nil {
			return nil, err
		}
		err = bb.SetUint8(uint8(len(tmp)))
		if err != nil {
			return nil, err
		}
		err = bb.Set(tmp)
		if err != nil {
			return nil, err
		}
	}
	return bb.Array(), nil
}

// GetValue returns the value of given attribute.
// When raw parameter us not used example register multiplies value by scalar.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Get parameters.
//
// Returns:
//
//	Value of the attribute index.
func (g *GXDLMSSecuritySetup) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	if e.Index == 1 {
		v, err := helpers.LogicalNameToBytes(g.LogicalName())
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return v, err
	}
	if e.Index == 2 {
		return g.SecurityPolicy, nil
	}
	if e.Index == 3 {
		return g.SecuritySuite, nil
	}
	if e.Index == 4 {
		return g.ClientSystemTitle, nil
	}
	if e.Index == 5 {
		return g.ServerSystemTitle, nil
	}
	if e.Index == 6 {
		return g.getSertificates()
	}
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, nil
}

func getStringFromAsn1(value []byte) (string, error) {
	issuer, err := types.Asn1FromByteArray(value)
	if err != nil {
		return "", err
	}
	if _, ok := issuer.(*types.GXAsn1Sequence); ok {
		return types.Asn1GetSubject(issuer.(*types.GXAsn1Sequence)), nil
	}
	return fmt.Sprint(issuer), nil
}

func (g *GXDLMSSecuritySetup) updateSertificates(list []any) error {
	var err error
	g.Certificates.Clear()
	for _, tmp := range list {
		it := tmp.([]any)
		info := &GXDLMSCertificateInfo{}
		info.Entity = enums.CertificateEntity(it[0].(int))
		info.Type = enums.CertificateType(it[1].(int))
		ret, err := types.Asn1FromByteArray(it[2].([]byte))
		if err != nil {
			return err
		}
		value := ret.(*types.GXAsn1Integer)
		serial := make([]byte, len(value.Value()))
		copy(serial, value.Value())
		buffer.Reverse(serial)
		info.SerialNumber = big.NewInt(0).SetBytes(serial)
		info.issuerRaw = []byte(it[3].([]byte))
		info.Issuer, err = getStringFromAsn1(info.issuerRaw)
		if err != nil {
			return err
		}
		info.subjectRaw = []byte(it[4].([]byte))
		info.Subject, err = getStringFromAsn1(info.subjectRaw)
		if err != nil {
			return err
		}
		info.subjectAltNameRaw = []byte(it[5].([]byte))
		info.SubjectAltName, err = getStringFromAsn1(info.subjectAltNameRaw)
		if err != nil {
			return err
		}
		g.Certificates = append(g.Certificates, info)
	}
	return err
}

// SetValue returns the set value of given attribute.
// When raw parameter us not used example register multiplies value by scalar.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Set parameters.
func (g *GXDLMSSecuritySetup) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	var err error
	switch e.Index {
	case 1:
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return g.SetLogicalName(ln)
	case 2:
		g.securityPolicy = enums.SecurityPolicy(e.Value.(types.GXEnum).Value)
	case 3:
		g.SecuritySuite = enums.SecuritySuite(e.Value.(types.GXEnum).Value)
	case 4:
		g.ClientSystemTitle = e.Value.([]byte)
	case 5:
		g.ServerSystemTitle = e.Value.([]byte)
	case 6:
		err = g.updateSertificates(e.Value.([]any))
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return err
}

// Constructor.
// ln: Logical Name of the object.
// sn: Short Name of the object.
func NewGXDLMSSecuritySetup(ln string, sn int16) (*GXDLMSSecuritySetup, error) {
	err := ValidateLogicalName(ln)
	if err != nil {
		return nil, err
	}
	return &GXDLMSSecuritySetup{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeSecuritySetup,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}

// Load returns the load object content from XML.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSSecuritySetup) Load(reader *GXXmlReader) error {
	ret, err := reader.ReadElementContentAsInt("SecurityPolicy", 0)
	if err != nil {
		return err
	}
	g.securityPolicy = enums.SecurityPolicy(ret)
	ret, err = reader.ReadElementContentAsInt("SecuritySuite", 0)
	if err != nil {
		return err
	}
	g.SecuritySuite = enums.SecuritySuite(ret)
	if g.SecuritySuite != enums.SecuritySuite0 && g.Version == 0 {
		return fmt.Errorf("Security suite %s is not available for Suite 0", g.SecuritySuite.String())
	}
	str, err := reader.ReadElementContentAsString("ClientSystemTitle", "")
	if err != nil {
		return err
	}
	if str == "" {
		g.ClientSystemTitle = nil
	} else {
		g.ClientSystemTitle = types.HexToBytes(str)
	}
	str, err = reader.ReadElementContentAsString("ServerSystemTitle", "")
	if err != nil {
		return err
	}
	if str == "" {
		g.ServerSystemTitle = nil
	} else {
		g.ServerSystemTitle = types.HexToBytes(str)
	}
	g.Certificates.Clear()
	if reader.isStartElementNamed2("Certificates", true) {
		for reader.isStartElementNamed2("Item", true) {
			it := &GXDLMSCertificateInfo{}
			g.Certificates = append(g.Certificates, it)
			ret, err := reader.ReadElementContentAsInt("Entity", 0)
			if err != nil {
				return err
			}
			it.Entity = enums.CertificateEntity(ret)
			ret, err = reader.ReadElementContentAsInt("Type", 0)
			if err != nil {
				return err
			}
			it.Type = enums.CertificateType(ret)
			ret2, err := reader.ReadElementContentAsString("SerialNumber", "")
			if ret2 != "" {
				ret, ok := new(big.Int).SetString(ret2, 10)
				if !ok {
					return fmt.Errorf("failed to parse SerialNumber: %s", ret2)
				}
				it.SerialNumber = ret
			}
			it.Issuer, err = reader.ReadElementContentAsString("Issuer", "")
			if err != nil {
				return err
			}
			it.Subject, err = reader.ReadElementContentAsString("Subject", "")
			if err != nil {
				return err
			}
			it.SubjectAltName, err = reader.ReadElementContentAsString("SubjectAltName", "")
			if err != nil {
				return err
			}
		}
		reader.ReadEndElement("Certificates")
	}
	str, err = reader.ReadElementContentAsString("SigningKey", "")
	if err != nil {
		return err
	}
	if str == "" {
		g.signingKey = nil
	} else {
		pk, err := types.Pkcs8FromDer(str)
		if err != nil {
			return err
		}
		g.signingKey = types.NewGXKeyValuePair(pk.PublicKey(), pk.PrivateKey())
		log.Printf("Signing Private Key: %s", pk.PrivateKey().ToHex())
		log.Printf("Signing Public Key: %s", pk.PublicKey().ToHex())
	}
	str, err = reader.ReadElementContentAsString("KeyAgreement", "")
	if err != nil {
		return err
	}
	if str == "" {
		g.keyAgreementKey = nil
	} else {
		pk, err := types.Pkcs8FromDer(str)
		if err != nil {
			return err
		}
		g.keyAgreementKey = types.NewGXKeyValuePair(pk.PublicKey(), pk.PrivateKey())
	}
	str, err = reader.ReadElementContentAsString("TLS", "")
	if err != nil {
		return err
	}
	if str == "" {
		g.tlsKey = nil
	} else {
		pk, err := types.Pkcs8FromDer(str)
		if err != nil {
			return err
		}
		g.tlsKey = types.NewGXKeyValuePair(pk.PublicKey(), pk.PrivateKey())
	}
	g.serverCertificates.Clear()
	if reader.isStartElementNamed2("ServerCertificates", true) {
		for reader.isStartElementNamed2("Cert", false) {
			ret, err := reader.ReadElementContentAsString("Cert", "")
			if err != nil {
				return err
			}
			cert, err := types.X509CertificateFromDer(ret)
			if err != nil {
				return err
			}
			g.serverCertificates = append(g.serverCertificates, cert)
		}
		err = reader.ReadEndElement("ServerCertificates")
		if err != nil {
			return err
		}
	}
	str, err = reader.ReadElementContentAsString("Guek", "")
	if err != nil {
		return err
	}
	if str != "" {
		g.guek = types.HexToBytes(str)
	}
	str, err = reader.ReadElementContentAsString("Gbek", "")
	if err != nil {
		return err
	}
	if str != "" {
		g.gbek = types.HexToBytes(str)
	}
	str, err = reader.ReadElementContentAsString("Gak", "")
	if err != nil {
		return err
	}
	if str != "" {
		g.gak = types.HexToBytes(str)
	}
	str, err = reader.ReadElementContentAsString("Kek", "")
	if err != nil {
		return err
	}
	if str != "" {
		g.Kek = types.HexToBytes(str)
	}
	return nil
}

// Save returns the save object content to XML.
//
// Parameters:
//
//	writer: XML writer.
func (g *GXDLMSSecuritySetup) Save(writer *GXXmlWriter) error {
	err := writer.WriteElementString("SecurityPolicy", uint8(g.securityPolicy))
	if err != nil {
		return err
	}
	err = writer.WriteElementString("SecuritySuite", int(g.SecuritySuite))
	if err != nil {
		return err
	}
	err = writer.WriteElementString("ClientSystemTitle", types.ToHex(g.ClientSystemTitle, false))
	if err != nil {
		return err
	}
	err = writer.WriteElementString("ServerSystemTitle", types.ToHex(g.ServerSystemTitle, false))
	if err != nil {
		return err
	}
	if g.Certificates != nil {
		writer.WriteStartElement("Certificates")
		for _, it := range g.Certificates {
			writer.WriteStartElement("Item")
			err = writer.WriteElementString("Entity", int(it.Entity))
			if err != nil {
				return err
			}
			err = writer.WriteElementString("Type", int(it.Type))
			if err != nil {
				return err
			}
			err = writer.WriteElementString("SerialNumber", fmt.Sprint(it.SerialNumber))
			if err != nil {
				return err
			}
			err = writer.WriteElementString("Issuer", it.Issuer)
			if err != nil {
				return err
			}
			err = writer.WriteElementString("Subject", it.Subject)
			if err != nil {
				return err
			}
			err = writer.WriteElementString("SubjectAltName", it.SubjectAltName)
			if err != nil {
				return err
			}
			writer.WriteEndElement()
		}
		writer.WriteEndElement()
	}
	if g.signingKey != nil {
		kp, err := types.NewGXPkcs8FromKeys(g.signingKey)
		if err != nil {
			return err
		}
		err = writer.WriteElementString("SigningKey", kp.ToDer())
		if err != nil {
			return err
		}
	}
	if g.keyAgreementKey != nil {
		kp, err := types.NewGXPkcs8FromKeys(g.keyAgreementKey)
		if err != nil {
			return err
		}
		err = writer.WriteElementString("KeyAgreement", kp.ToDer())
		if err != nil {
			return err
		}
	}
	if g.tlsKey != nil {
		kp, err := types.NewGXPkcs8FromKeys(g.tlsKey)
		if err != nil {
			return err
		}
		err = writer.WriteElementString("TLS", kp.ToDer())
		if err != nil {
			return err
		}
	}
	if len(g.serverCertificates) != 0 {
		writer.WriteStartElement("ServerCertificates")
		for _, it := range g.serverCertificates {
			err = writer.WriteElementString("Cert", it.ToDer())
			if err != nil {
				return err
			}
		}
		writer.WriteEndElement()
	}
	if g.guek != nil {
		err = writer.WriteElementString("Guek", types.ToHex(g.guek, false))
		if err != nil {
			return err
		}
	}
	if g.gbek != nil {
		err = writer.WriteElementString("Gbek", types.ToHex(g.gbek, false))
		if err != nil {
			return err
		}
	}
	if g.gak != nil {
		err = writer.WriteElementString("Gak", types.ToHex(g.gak, false))
		if err != nil {
			return err
		}
	}
	if g.Kek != nil {
		err = writer.WriteElementString("Kek", types.ToHex(g.Kek, false))
		if err != nil {
			return err
		}
	}
	return err
}

// PostLoad returns the handle actions after Load.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSSecuritySetup) PostLoad(reader *GXXmlReader) error {
	return nil
}

// ApplyKeys returns the start to use new keys after reply is generated.
//
// Parameters:
//
//	settings: DLMS settings.
func (g *GXDLMSSecuritySetup) ApplyKeys(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	for _, t := range e.Parameters.([]any) {
		item := t.([]any)
		type_ := enums.GlobalKeyType(item[0].(int))
		data := item[1].([]byte)
		key, err := internal.Decrypt(settings.Kek, data)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return err
		}
		switch type_ {
		case enums.GlobalKeyTypeUnicastEncryption:
			err = settings.Cipher.SetBlockCipherKey(key)
		case enums.GlobalKeyTypeBroadcastEncryption:
			err = settings.Cipher.SetBroadcastBlockCipherKey(key)
		case enums.GlobalKeyTypeAuthentication:
			err = settings.Cipher.SetAuthenticationKey(key)
		case enums.GlobalKeyTypeKek:
			settings.Kek = key
		default:
			e.Error = enums.ErrorCodeReadWriteDenied
		}
	}
	return nil
}

// GetValues returns the an array containing the COSEM object's attribute values.
func (g *GXDLMSSecuritySetup) GetValues() []any {
	return []any{g.LogicalName(), g.securityPolicy, g.SecuritySuite, g.ClientSystemTitle, g.ServerSystemTitle, g.Certificates}
}

// Activate returns the activates and strengthens the security policy.
//
// Parameters:
//
//	client: DLMS client that is used to generate action.
//	security: New security level.
//
// Returns:
//
//	Generated action.
func (g *GXDLMSSecuritySetup) Activate(client IGXDLMSClient, policy enums.SecurityPolicy) ([][]byte, error) {
	return client.Method(g, 1, uint8(policy), enums.DataTypeEnum)
}

// GlobalKeyTransfer returns the updates one or more global keys.
//
// Parameters:
//
//	client: DLMS client that is used to generate action.
//	kek: Master key, also known as Key Encrypting Key.
//	list: List of Global key types and keys.
//
// Returns:
//
//	Generated action.
func (g *GXDLMSSecuritySetup) GlobalKeyTransfer(client IGXDLMSClient, kek []byte, list []types.GXKeyValuePair[enums.GlobalKeyType, []byte]) ([][]byte, error) {
	var err error
	if len(list) == 0 {
		return nil, errors.New("Invalid list. It is empty.")
	}
	bb := types.GXByteBuffer{}
	err = bb.SetUint8(enums.DataTypeArray)
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(uint8(len(list)))
	if err != nil {
		return nil, err
	}
	var tmp []byte
	for _, it := range list {
		err = bb.SetUint8(enums.DataTypeStructure)
		if err != nil {
			return nil, err
		}
		err = bb.SetUint8(2)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(client.Settings(), &bb, enums.DataTypeEnum, it.Key)
		if err != nil {
			return nil, err
		}
		tmp, err = internal.Encrypt(kek, it.Value)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(client.Settings(), &bb, enums.DataTypeOctetString, tmp)
		if err != nil {
			return nil, err
		}
	}
	return client.Method(g, 2, bb.Array(), enums.DataTypeArray)
}

// KeyAgreement returns the agree on one or more symmetric keys using the key agreement algorithm.
//
// Parameters:
//
//	client: DLMS client that is used to generate action.
//	list: List of keys.
//
// Returns:
//
//	Generated action.
func (g *GXDLMSSecuritySetup) KeyAgreement(client IGXDLMSClient,
	list []types.GXKeyValuePair[enums.GlobalKeyType, []byte]) ([][]byte, error) {
	var err error
	if len(list) == 0 {
		return nil, errors.New("Invalid list. It is empty.")
	}
	bb := types.GXByteBuffer{}
	err = bb.SetUint8(enums.DataTypeArray)
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(uint8(len(list)))
	if err != nil {
		return nil, err
	}
	for _, it := range list {
		err = bb.SetUint8(enums.DataTypeStructure)
		if err != nil {
			return nil, err
		}
		err = bb.SetUint8(2)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(client.Settings(), &bb, enums.DataTypeEnum, it.Key)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(client.Settings(), &bb, enums.DataTypeOctetString, it.Value)
		if err != nil {
			return nil, err
		}
	}
	return client.Method(g, 3, bb.Array(), enums.DataTypeArray)
}

func getSettings(client IGXDLMSClient) *settings.GXDLMSSettings {

	return client.Settings().(*settings.GXDLMSSettings)
}

func getCipheting(client IGXDLMSClient) settings.GXICipher {

	return getSettings(client).Cipher
}

// KeyAgreement returns the agree on global unicast encryption key.
//
// Parameters:
//
//	client: DLMS client that is used to generate action.
//	key: List of keys.
//
// Returns:
//
//	Generated action.
func (g *GXDLMSSecuritySetup) KeyAgreementFromClient(client IGXDLMSClient) ([][]byte, error) {
	var err error
	if g.Version == 0 {
		return nil, errors.New("Public and private key isn't implemented for version 0.")
	}
	if getCipheting(client).EphemeralKeyPair().Value == nil {
		return nil, errors.New("Invalid Ephemeral key.")
	}
	if getCipheting(client).SigningKeyPair().Value == nil {
		return nil, errors.New("Invalid Signiture key.")
	}
	bb := types.GXByteBuffer{}
	ek := getCipheting(client).EphemeralKeyPair()
	err = bb.SetAt(ek.Value.RawValue(), 1, len(ek.Value.RawValue())-1)
	if err != nil {
		return nil, err
	}
	sign, err := settings.GetEphemeralPublicKeySignature(0, ek.Key, getCipheting(client).SigningKeyPair().Value)
	if err != nil {
		return nil, err
	}
	err = bb.Set(sign)
	if err != nil {
		return nil, err
	}
	list := []types.GXKeyValuePair[enums.GlobalKeyType, []byte]{}
	list = append(list, *types.NewGXKeyValuePair(enums.GlobalKeyTypeUnicastEncryption, bb.Array()))
	return g.KeyAgreement(client, list)
}

// GenerateKeyPair returns the generates an asymmetric key pair as required by the security suite.
//
// Parameters:
//
//	client: DLMS client that is used to generate action.
//	type_: New certificate type_.
//
// Returns:
//
//	Generated action.
func (g *GXDLMSSecuritySetup) GenerateKeyPair(client IGXDLMSClient, type_ enums.CertificateType) ([][]byte, error) {
	return client.Method(g, 4, int(type_), enums.DataTypeEnum)
}

// GenerateCertificate returns the ask Server sends the Certificate Signing Request (CSR) data.
//
// Parameters:
//
//	client: DLMS client that is used to generate action.
//	type_: identifies the key pair for which the certificate will be requested.
//
// Returns:
//
//	Generated action.
func (g *GXDLMSSecuritySetup) GenerateCertificate(client IGXDLMSClient, type_ enums.CertificateType) ([][]byte, error) {
	return client.Method(g, 5, type_, enums.DataTypeEnum)
}

// ImportCertificate returns the imports an X.509 v3 certificate of a public key.
//
// Parameters:
//
//	client: DLMS client that is used to generate action.
//	key: Public key.
//
// Returns:
//
//	Generated action.
func (g *GXDLMSSecuritySetup) ImportCertificate(client IGXDLMSClient, certificate *types.GXx509Certificate) ([][]byte, error) {
	ret, err := certificate.Encoded()
	if err != nil {
		return nil, err
	}
	return g.ImportCertificateFromBytes(client, ret)
}

// ImportCertificateFromBytes returns the imports an X.509 v3 certificate of a public key.
//
// Parameters:
//
//	client: DLMS client that is used to generate action.
//	key: Public key.
//
// Returns:
//
//	Generated action.
func (g *GXDLMSSecuritySetup) ImportCertificateFromBytes(client IGXDLMSClient, key []byte) ([][]byte, error) {
	return client.Method(g, 6, key, enums.DataTypeOctetString)
}

// ExportCertificateByEntity returns the exports an X.509 v3 certificate from the server using entity information.
//
// Parameters:
//
//	client: DLMS client that is used to generate action.
//	entity: Certificate entity.
//	type_: Certificate type_.
//	systemTitle: System title.
//
// Returns:
//
//	Generated action.
func (g *GXDLMSSecuritySetup) ExportCertificateByEntity(client IGXDLMSClient,
	entity enums.CertificateEntity,
	type_ enums.CertificateType,
	systemTitle []byte) ([][]byte, error) {
	var err error
	bb := types.GXByteBuffer{}
	err = bb.SetUint8(enums.DataTypeStructure)
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(2)
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(enums.DataTypeEnum)
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(0)
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(enums.DataTypeStructure)
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(3)
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(enums.DataTypeEnum)
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(uint8(entity))
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(enums.DataTypeEnum)
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(uint8(type_))
	if err != nil {
		return nil, err
	}
	err = internal.SetData(client.Settings(), &bb, enums.DataTypeOctetString, systemTitle)
	if err != nil {
		return nil, err
	}
	return client.Method(g, 7, bb.Array(), enums.DataTypeStructure)
}

// ExportCertificateBySerial returns the exports an X.509 v3 certificate from the server using serial information.
//
// Parameters:
//
//	client: DLMS client that is used to generate action.
//	serialNumber: Serial number.
//	issuer: Issuer
//
// Returns:
//
//	Generated action.
func (g *GXDLMSSecuritySetup) ExportCertificateBySerial(client IGXDLMSClient, serialNumber big.Int, issuer []byte) ([][]byte, error) {
	g.VerifyIssuer(issuer)
	bb := types.GXByteBuffer{}
	err := bb.SetUint8(enums.DataTypeStructure)
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(2)
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(enums.DataTypeEnum)
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(1)
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(enums.DataTypeStructure)
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(2)
	if err != nil {
		return nil, err
	}
	sn, err := types.Asn1ToByteArray(types.NewGXAsn1IntegerFromBigInteger(serialNumber))
	if err != nil {
		return nil, err
	}
	err = internal.SetData(client.Settings(), &bb, enums.DataTypeOctetString, sn)
	if err != nil {
		return nil, err
	}
	err = internal.SetData(client.Settings(), &bb, enums.DataTypeOctetString, issuer)
	if err != nil {
		return nil, err
	}
	return client.Method(g, 7, bb.Array(), enums.DataTypeStructure)
}

// RemoveCertificateByEntity returns the removes X.509 v3 certificate from the server using entity.
//
// Parameters:
//
//	client: DLMS client that is used to generate action.
//	entity: Certificate entity type_.
//	type_: Certificate type_.
//	systemTitle: System title.
//
// Returns:
//
//	Generated action.
func (g *GXDLMSSecuritySetup) RemoveCertificateByEntity(client IGXDLMSClient,
	entity enums.CertificateEntity,
	type_ enums.CertificateType,
	systemTitle []byte) ([][]byte, error) {
	bb := types.GXByteBuffer{}
	err := bb.SetUint8(enums.DataTypeStructure)
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(2)
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(enums.DataTypeEnum)
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(0)
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(enums.DataTypeStructure)
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(3)
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(enums.DataTypeEnum)
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(uint8(entity))
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(enums.DataTypeEnum)
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(uint8(type_))
	if err != nil {
		return nil, err
	}
	err = internal.SetData(client.Settings(), &bb, enums.DataTypeOctetString, systemTitle)
	if err != nil {
		return nil, err
	}
	return client.Method(g, 8, bb.Array(), enums.DataTypeStructure)
}

// RemoveCertificateBySerial returns the removes X.509 v3 certificate from the server using serial number.
//
// Parameters:
//
//	client: DLMS client that is used to generate action.
//	serialNumber: Serial number.
//	issuer: Issuer.
//
// Returns:
//
//	Generated action.
func (g *GXDLMSSecuritySetup) RemoveCertificateBySerial(client IGXDLMSClient,
	serialNumber big.Int,
	issuer []byte) ([][]byte, error) {
	g.VerifyIssuer(issuer)
	bb := types.GXByteBuffer{}
	err := bb.SetUint8(enums.DataTypeStructure)
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(2)
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(enums.DataTypeEnum)
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(1)
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(enums.DataTypeStructure)
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(2)
	if err != nil {
		return nil, err
	}
	sn, err := types.Asn1ToByteArray(types.NewGXAsn1IntegerFromBigInteger(serialNumber))
	if err != nil {
		return nil, err
	}
	err = internal.SetData(client.Settings(), &bb, enums.DataTypeOctetString, sn)
	if err != nil {
		return nil, err
	}
	err = internal.SetData(client.Settings(), &bb, enums.DataTypeOctetString, issuer)
	if err != nil {
		return nil, err
	}
	return client.Method(g, 8, bb.Array(), enums.DataTypeStructure)
}

// UpdateEphemeralKeys returns the update ephemeral keys.
//
// Returns:
//
//	List of Parsed key id and GUAK. This is for debugging purpose.
func (g *GXDLMSSecuritySetup) UpdateEphemeralKeys(client IGXDLMSClient, value []byte) ([]*types.GXKeyValuePair[enums.GlobalKeyType, []byte], error) {
	bb := types.NewGXByteBufferWithData(value)
	return g.UpdateEphemeralKeysFromByteBuffer(client, bb)
}

// UpdateEphemeralKeys returns the update ephemeral keys.
//
// Parameters:
//
//	client: DLMS Client.
//	value: Received reply from the server.
//
// Returns:
//
//	List of Parsed key id and GUAK. This is for debugging purpose.
func (g *GXDLMSSecuritySetup) UpdateEphemeralKeysFromByteBuffer(client IGXDLMSClient,
	value *types.GXByteBuffer) ([]*types.GXKeyValuePair[enums.GlobalKeyType, []byte], error) {
	if client == nil {
		return nil, errors.New("client")
	}
	ret, err := value.Uint8()
	if err != nil {
		return nil, err
	}
	if ret != uint8(enums.DataTypeArray) {
		return nil, errors.New("Invalid tag.")
	}
	chipering := getCipheting(client)
	c, err := types.NewGXEcdsaFromPublicKey(chipering.EphemeralKeyPair().Key)
	if err != nil {
		return nil, err
	}
	count, err := types.GetObjectCount(value)
	if err != nil {
		return nil, err
	}
	list := []*types.GXKeyValuePair[enums.GlobalKeyType, []byte]{}
	for pos := 0; pos != count; pos++ {
		ret, err := value.Uint8()
		if err != nil {
			return nil, err
		}
		if ret != uint8(enums.DataTypeStructure) {
			return nil, errors.New("Invalid tag.")
		}
		ret, err = value.Uint8()
		if err != nil {
			return nil, err
		}
		if ret != 2 {
			return nil, errors.New("Invalid length.")
		}
		ret, err = value.Uint8()
		if err != nil {
			return nil, err
		}
		if ret != uint8(enums.DataTypeEnum) {
			return nil, errors.New("Invalid key id data type.")
		}
		keyId, err := value.Uint8()
		if err != nil {
			return nil, err
		}
		if keyId > 4 {
			return nil, errors.New("Invalid key type.")
		}
		ret, err = value.Uint8()
		if err != nil {
			return nil, err
		}
		if ret != uint8(enums.DataTypeOctetString) {
			return nil, errors.New("Invalid tag.")
		}
		len_, err := types.GetObjectCount(value)
		if err != nil {
			return nil, err
		}
		if len_ != 128 {
			return nil, errors.New("Invalid length.")
		}
		// Get ephemeral public key server.
		key := types.GXByteBuffer{}
		err = key.SetUint8(4)
		if err != nil {
			return nil, err
		}
		err = key.SetByteBufferByCount(value, 64)
		if err != nil {
			return nil, err
		}
		targetEphemeralKey, err := types.PublicKeyFromRawBytes(key.Array())
		if err != nil {
			return nil, err
		}
		// Get ephemeral public key signature server.
		signature := make([]byte, 64)
		err = value.Get(signature)
		if err != nil {
			return nil, err
		}
		err = key.SetUint8At(0, uint8(keyId))
		if err != nil {
			return nil, err
		}
		ret2, err := settings.ValidateEphemeralPublicKeySignature(key.Array(), signature, chipering.SigningKeyPair().Key)
		if err != nil {
			return nil, err
		}
		if !ret2 {
			return nil, errors.New("Invalid signature.")
		}
		z, err := c.GenerateSecret(targetEphemeralKey)
		if err != nil {
			return nil, err
		}
		log.Printf("Shared secret: %s\n", types.ToHex(z, true))
		kdf := types.GXByteBuffer{}
		ret3, err := settings.GenerateKDFWithInfo(chipering.SecuritySuite(), z, enums.AlgorithmIdAesGcm128, chipering.SystemTitle(), getSettings(client).SourceSystemTitle(), nil, nil)
		err = kdf.Set(ret3)
		if err != nil {
			return nil, err
		}
		log.Printf("Shared secret: %s\n", types.ToHex(z, true))
		kdf = types.GXByteBuffer{}
		ret3, err = settings.GenerateKDFWithInfo(chipering.SecuritySuite(), z, enums.AlgorithmIdAesGcm128, chipering.SystemTitle(), getSettings(client).SourceSystemTitle(), nil, nil)
		if err != nil {
			return nil, err
		}
		err = kdf.Set(ret3)
		if err != nil {
			return nil, err
		}
		log.Println("KDF:" + kdf.String())
		ret3, err = kdf.SubArray(0, 16)
		if err != nil {
			return nil, err
		}
		list = append(list, types.NewGXKeyValuePair(enums.GlobalKeyType(keyId), ret3))
	}
	s := getSettings(client)
	for _, v := range list {
		switch v.Key {
		case enums.GlobalKeyTypeUnicastEncryption:
			s.EphemeralBlockCipherKey = v.Value
		case enums.GlobalKeyTypeBroadcastEncryption:
			s.EphemeralBroadcastBlockCipherKey = v.Value
		case enums.GlobalKeyTypeAuthentication:
			s.EphemeralAuthenticationKey = v.Value
		case enums.GlobalKeyTypeKek:
			s.EphemeralKek = v.Value
		}
	}
	return list, nil
}

// GetDataType returns the device data type of selected attribute index.
//
// Parameters:
//
//	index: Attribute index of the object.
//
// Returns:
//
//	Device data type_ of the object.
func (g *GXDLMSSecuritySetup) GetDataType(index int) (enums.DataType, error) {
	if index == 1 {
		return enums.DataTypeOctetString, nil
	}
	if index == 2 {
		return enums.DataTypeEnum, nil
	}
	if index == 3 {
		return enums.DataTypeEnum, nil
	}
	if index == 4 {
		return enums.DataTypeOctetString, nil
	}
	if index == 5 {
		return enums.DataTypeOctetString, nil
	}
	if g.Version > 0 && index == 6 {
		return enums.DataTypeArray, nil
	}
	return enums.DataTypeNone, errors.New("GetDataType failed. Invalid attribute index.")
}
