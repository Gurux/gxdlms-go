package enums

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
// Gurux Device Framework is Open Source software; you can redistribute it
// and/or modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation; version 2 of the License.
// Gurux Device Framework is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
// See the GNU General Public License for more details.
//
// More information of Gurux products: https://www.gurux.org
//
// This code is licensed under the GNU General Public License v2.
// Full text may be retrieved at http://www.gnu.org/licenses/gpl-2.0.txt
//---------------------------------------------------------------------------

import (
	"fmt"
	"strings"

	"github.com/Gurux/gxcommon-go"
)

type PkcsObjectIdentifier int

const (
	PkcsObjectIdentifierNone PkcsObjectIdentifier = iota
	PkcsObjectIdentifierRsaEncryption
	PkcsObjectIdentifierMD2WithRsaEncryption
	PkcsObjectIdentifierMD4WithRsaEncryption
	PkcsObjectIdentifierMD5WithRsaEncryption
	PkcsObjectIdentifierSha1WithRsaEncryption
	PkcsObjectIdentifierSrsaOaepEncryptionSet
	PkcsObjectIdentifierIdRsaesOaep
	PkcsObjectIdentifierIdMgf1
	PkcsObjectIdentifierIdPSpecified
	PkcsObjectIdentifierIdRsassaPss
	PkcsObjectIdentifierSha256WithRsaEncryption
	PkcsObjectIdentifierSha384WithRsaEncryption
	PkcsObjectIdentifierSha512WithRsaEncryption
	PkcsObjectIdentifierSha224WithRsaEncryption
	PkcsObjectIdentifierDhKeyAgree1ment
	PkcsObjectIdentifierPbeWithMD2AndDesCbc
	PkcsObjectIdentifierPbeWithMD2AndRC2Cbc
	PkcsObjectIdentifierPbeWithMD5AndDesCbc
	PkcsObjectIdentifierPbeWithMD5AndRC2Cbc
	PkcsObjectIdentifierPbeWithSha1AndDesCbc
	PkcsObjectIdentifierPbeWithSha1AndRC2Cbc
	PkcsObjectIdentifierIdPbeS2
	PkcsObjectIdentifierIdPbkdf2
	PkcsObjectIdentifierDesEde3Cbc
	PkcsObjectIdentifierRC2Cbc
	PkcsObjectIdentifierMD2
	PkcsObjectIdentifierMD4
	PkcsObjectIdentifierMD5
	PkcsObjectIdentifierIdHmacWithSha1
	PkcsObjectIdentifierIdHmacWithSha224
	PkcsObjectIdentifierIdHmacWithSha256
	PkcsObjectIdentifierIdHmacWithSha384
	PkcsObjectIdentifierIdHmacWithSha512
	PkcsObjectIdentifierData
	PkcsObjectIdentifierSignedData
	PkcsObjectIdentifierEnvelopedData
	PkcsObjectIdentifierSignedAndEnvelopedData
	PkcsObjectIdentifierDigestedData
	PkcsObjectIdentifierEncryptedData
	PkcsObjectIdentifierPkcs9AtEmailAddress
	PkcsObjectIdentifierPkcs9AtUnstructuredName
	PkcsObjectIdentifierPkcs9AtContentType
	PkcsObjectIdentifierPkcs9AtMessageDigest
	PkcsObjectIdentifierPkcs9AtSigningTime
	PkcsObjectIdentifierPkcs9AtCounterSignature
	PkcsObjectIdentifierPkcs9AtChallengePassword
	PkcsObjectIdentifierPkcs9AtUnstructuredAddress
	PkcsObjectIdentifierPkcs9AtExtendedCertificateAttributes
	PkcsObjectIdentifierPkcs9AtSigningDescription
	PkcsObjectIdentifierPkcs9AtExtensionRequest
	PkcsObjectIdentifierPkcs9AtSmimeCapabilities
	PkcsObjectIdentifierIdSmime
	PkcsObjectIdentifierPkcs9AtFriendlyName
	PkcsObjectIdentifierPkcs9AtLocalKeyID
	PkcsObjectIdentifierX509Certificate
	PkcsObjectIdentifierSdsiCertificate
	PkcsObjectIdentifierX509Crl
	PkcsObjectIdentifierIdAlg
	PkcsObjectIdentifierIdAlgEsdh
	PkcsObjectIdentifierIdAlgCms3DesWrap
	PkcsObjectIdentifierIdAlgCmsRC2Wrap
	PkcsObjectIdentifierIdAlgPwriKek
	PkcsObjectIdentifierIdAlgSsdh
	PkcsObjectIdentifierIdRsaKem
	PkcsObjectIdentifierPreferSignedData
	PkcsObjectIdentifierCannotDecryptAny
	PkcsObjectIdentifierSmimeCapabilitiesVersions
	PkcsObjectIdentifierIdAAReceiptRequest
	PkcsObjectIdentifierIdCTAuthData
	PkcsObjectIdentifierIdCTTstInfo
	PkcsObjectIdentifierIdCTCompressedData
	PkcsObjectIdentifierIdCTAuthEnvelopedData
	PkcsObjectIdentifierIdCTTimestampedData
	PkcsObjectIdentifierIdCtiEtsProofOfOrigin
	PkcsObjectIdentifierIdCtiEtsProofOfReceipt
	PkcsObjectIdentifierIdCtiEtsProofOfDelivery
	PkcsObjectIdentifierIdCtiEtsProofOfSender
	PkcsObjectIdentifierIdCtiEtsProofOfApproval
	PkcsObjectIdentifierIdCtiEtsProofOfCreation
	PkcsObjectIdentifierIdAAContentHint
	PkcsObjectIdentifierIdAAMsgSigDigest
	PkcsObjectIdentifierIdAAContentReference
	PkcsObjectIdentifierIdAAEncrypKeyPref
	PkcsObjectIdentifierIdAASigningCertificate
	PkcsObjectIdentifierIdAASigningCertificateV2
	PkcsObjectIdentifierIdAAContentIdentifier
	PkcsObjectIdentifierIdAASignatureTimeStampToken
	PkcsObjectIdentifierIdAAEtsSigPolicyID
	PkcsObjectIdentifierIdAAEtsCommitmentType
	PkcsObjectIdentifierIdAAEtsSignerLocation
	PkcsObjectIdentifierIdAAEtsSignerAttr
	PkcsObjectIdentifierIdAAEtsOtherSigCert
	PkcsObjectIdentifierIdAAEtsContentTimestamp
	PkcsObjectIdentifierIdAAEtsCertificateRefs
	PkcsObjectIdentifierIdAAEtsRevocationRefs
	PkcsObjectIdentifierIdAAEtsCertValues
	PkcsObjectIdentifierIdAAEtsRevocationValues
	PkcsObjectIdentifierIdAAEtsEscTimeStamp
	PkcsObjectIdentifierIdAAEtsCertCrlTimestamp
	PkcsObjectIdentifierIdAAEtsArchiveTimestamp
	PkcsObjectIdentifierIdSpqEtsUri
	PkcsObjectIdentifierIdSpqEtsUNotice
	PkcsObjectIdentifierKeyBag
	PkcsObjectIdentifierPkcs8ShroudedKeyBag
	PkcsObjectIdentifierCertBag
	PkcsObjectIdentifierCrlBag
	PkcsObjectIdentifierSecretBag
	PkcsObjectIdentifierSafeContentsBag
	PkcsObjectIdentifierPbeWithShaAnd128BitRC4
	PkcsObjectIdentifierPbeWithShaAnd40BitRC4
	PkcsObjectIdentifierPbeWithShaAnd3KeyTripleDesCbc
	PkcsObjectIdentifierPbeWithShaAnd2KeyTripleDesCbc
	PkcsObjectIdentifierPbeWithShaAnd128BitRC2Cbc
	PkcsObjectIdentifierPbewithShaAnd40BitRC2Cbc
)

// PkcsObjectIdentifierParse converts the given string into a PkcsObjectIdentifier value.
//
// It returns the corresponding PkcsObjectIdentifier constant if the string matches
// a known level name, or an error if the input is invalid.
func PkcsObjectIdentifierParse(value string) (PkcsObjectIdentifier, error) {
	var ret PkcsObjectIdentifier
	var err error
	switch strings.ToUpper(value) {
	case "NONE":
		ret = PkcsObjectIdentifierNone
	case "RSAENCRYPTION":
		ret = PkcsObjectIdentifierRsaEncryption
	case "MD2WITHRSAENCRYPTION":
		ret = PkcsObjectIdentifierMD2WithRsaEncryption
	case "MD4WITHRSAENCRYPTION":
		ret = PkcsObjectIdentifierMD4WithRsaEncryption
	case "MD5WITHRSAENCRYPTION":
		ret = PkcsObjectIdentifierMD5WithRsaEncryption
	case "SHA1WITHRSAENCRYPTION":
		ret = PkcsObjectIdentifierSha1WithRsaEncryption
	case "SRSAOAEPENCRYPTIONSET":
		ret = PkcsObjectIdentifierSrsaOaepEncryptionSet
	case "IDRSAESOAEP":
		ret = PkcsObjectIdentifierIdRsaesOaep
	case "IDMGF1":
		ret = PkcsObjectIdentifierIdMgf1
	case "IDPSPECIFIED":
		ret = PkcsObjectIdentifierIdPSpecified
	case "IDRSASSAPSS":
		ret = PkcsObjectIdentifierIdRsassaPss
	case "SHA256WITHRSAENCRYPTION":
		ret = PkcsObjectIdentifierSha256WithRsaEncryption
	case "SHA384WITHRSAENCRYPTION":
		ret = PkcsObjectIdentifierSha384WithRsaEncryption
	case "SHA512WITHRSAENCRYPTION":
		ret = PkcsObjectIdentifierSha512WithRsaEncryption
	case "SHA224WITHRSAENCRYPTION":
		ret = PkcsObjectIdentifierSha224WithRsaEncryption
	case "DHKEYAGREE1MENT":
		ret = PkcsObjectIdentifierDhKeyAgree1ment
	case "PBEWITHMD2ANDDESCBC":
		ret = PkcsObjectIdentifierPbeWithMD2AndDesCbc
	case "PBEWITHMD2ANDRC2CBC":
		ret = PkcsObjectIdentifierPbeWithMD2AndRC2Cbc
	case "PBEWITHMD5ANDDESCBC":
		ret = PkcsObjectIdentifierPbeWithMD5AndDesCbc
	case "PBEWITHMD5ANDRC2CBC":
		ret = PkcsObjectIdentifierPbeWithMD5AndRC2Cbc
	case "PBEWITHSHA1ANDDESCBC":
		ret = PkcsObjectIdentifierPbeWithSha1AndDesCbc
	case "PBEWITHSHA1ANDRC2CBC":
		ret = PkcsObjectIdentifierPbeWithSha1AndRC2Cbc
	case "IDPBES2":
		ret = PkcsObjectIdentifierIdPbeS2
	case "IDPBKDF2":
		ret = PkcsObjectIdentifierIdPbkdf2
	case "DESEDE3CBC":
		ret = PkcsObjectIdentifierDesEde3Cbc
	case "RC2CBC":
		ret = PkcsObjectIdentifierRC2Cbc
	case "MD2":
		ret = PkcsObjectIdentifierMD2
	case "MD4":
		ret = PkcsObjectIdentifierMD4
	case "MD5":
		ret = PkcsObjectIdentifierMD5
	case "IDHMACWITHSHA1":
		ret = PkcsObjectIdentifierIdHmacWithSha1
	case "IDHMACWITHSHA224":
		ret = PkcsObjectIdentifierIdHmacWithSha224
	case "IDHMACWITHSHA256":
		ret = PkcsObjectIdentifierIdHmacWithSha256
	case "IDHMACWITHSHA384":
		ret = PkcsObjectIdentifierIdHmacWithSha384
	case "IDHMACWITHSHA512":
		ret = PkcsObjectIdentifierIdHmacWithSha512
	case "DATA":
		ret = PkcsObjectIdentifierData
	case "SIGNEDDATA":
		ret = PkcsObjectIdentifierSignedData
	case "ENVELOPEDDATA":
		ret = PkcsObjectIdentifierEnvelopedData
	case "SIGNEDANDENVELOPEDDATA":
		ret = PkcsObjectIdentifierSignedAndEnvelopedData
	case "DIGESTEDDATA":
		ret = PkcsObjectIdentifierDigestedData
	case "ENCRYPTEDDATA":
		ret = PkcsObjectIdentifierEncryptedData
	case "PKCS9ATEMAILADDRESS":
		ret = PkcsObjectIdentifierPkcs9AtEmailAddress
	case "PKCS9ATUNSTRUCTUREDNAME":
		ret = PkcsObjectIdentifierPkcs9AtUnstructuredName
	case "PKCS9ATCONTENTTYPE":
		ret = PkcsObjectIdentifierPkcs9AtContentType
	case "PKCS9ATMESSAGEDIGEST":
		ret = PkcsObjectIdentifierPkcs9AtMessageDigest
	case "PKCS9ATSIGNINGTIME":
		ret = PkcsObjectIdentifierPkcs9AtSigningTime
	case "PKCS9ATCOUNTERSIGNATURE":
		ret = PkcsObjectIdentifierPkcs9AtCounterSignature
	case "PKCS9ATCHALLENGEPASSWORD":
		ret = PkcsObjectIdentifierPkcs9AtChallengePassword
	case "PKCS9ATUNSTRUCTUREDADDRESS":
		ret = PkcsObjectIdentifierPkcs9AtUnstructuredAddress
	case "PKCS9ATEXTENDEDCERTIFICATEATTRIBUTES":
		ret = PkcsObjectIdentifierPkcs9AtExtendedCertificateAttributes
	case "PKCS9ATSIGNINGDESCRIPTION":
		ret = PkcsObjectIdentifierPkcs9AtSigningDescription
	case "PKCS9ATEXTENSIONREQUEST":
		ret = PkcsObjectIdentifierPkcs9AtExtensionRequest
	case "PKCS9ATSMIMECAPABILITIES":
		ret = PkcsObjectIdentifierPkcs9AtSmimeCapabilities
	case "IDSMIME":
		ret = PkcsObjectIdentifierIdSmime
	case "PKCS9ATFRIENDLYNAME":
		ret = PkcsObjectIdentifierPkcs9AtFriendlyName
	case "PKCS9ATLOCALKEYID":
		ret = PkcsObjectIdentifierPkcs9AtLocalKeyID
	case "X509CERTIFICATE":
		ret = PkcsObjectIdentifierX509Certificate
	case "SDSICERTIFICATE":
		ret = PkcsObjectIdentifierSdsiCertificate
	case "X509CRL":
		ret = PkcsObjectIdentifierX509Crl
	case "IDALG":
		ret = PkcsObjectIdentifierIdAlg
	case "IDALGESDH":
		ret = PkcsObjectIdentifierIdAlgEsdh
	case "IDALGCMS3DESWRAP":
		ret = PkcsObjectIdentifierIdAlgCms3DesWrap
	case "IDALGCMSRC2WRAP":
		ret = PkcsObjectIdentifierIdAlgCmsRC2Wrap
	case "IDALGPWRIKEK":
		ret = PkcsObjectIdentifierIdAlgPwriKek
	case "IDALGSSDH":
		ret = PkcsObjectIdentifierIdAlgSsdh
	case "IDRSAKEM":
		ret = PkcsObjectIdentifierIdRsaKem
	case "PREFERSIGNEDDATA":
		ret = PkcsObjectIdentifierPreferSignedData
	case "CANNOTDECRYPTANY":
		ret = PkcsObjectIdentifierCannotDecryptAny
	case "SMIMECAPABILITIESVERSIONS":
		ret = PkcsObjectIdentifierSmimeCapabilitiesVersions
	case "IDAARECEIPTREQUEST":
		ret = PkcsObjectIdentifierIdAAReceiptRequest
	case "IDCTAUTHDATA":
		ret = PkcsObjectIdentifierIdCTAuthData
	case "IDCTTSTINFO":
		ret = PkcsObjectIdentifierIdCTTstInfo
	case "IDCTCOMPRESSEDDATA":
		ret = PkcsObjectIdentifierIdCTCompressedData
	case "IDCTAUTHENVELOPEDDATA":
		ret = PkcsObjectIdentifierIdCTAuthEnvelopedData
	case "IDCTTIMESTAMPEDDATA":
		ret = PkcsObjectIdentifierIdCTTimestampedData
	case "IDCTIETSPROOFOFORIGIN":
		ret = PkcsObjectIdentifierIdCtiEtsProofOfOrigin
	case "IDCTIETSPROOFOFRECEIPT":
		ret = PkcsObjectIdentifierIdCtiEtsProofOfReceipt
	case "IDCTIETSPROOFOFDELIVERY":
		ret = PkcsObjectIdentifierIdCtiEtsProofOfDelivery
	case "IDCTIETSPROOFOFSENDER":
		ret = PkcsObjectIdentifierIdCtiEtsProofOfSender
	case "IDCTIETSPROOFOFAPPROVAL":
		ret = PkcsObjectIdentifierIdCtiEtsProofOfApproval
	case "IDCTIETSPROOFOFCREATION":
		ret = PkcsObjectIdentifierIdCtiEtsProofOfCreation
	case "IDAACONTENTHINT":
		ret = PkcsObjectIdentifierIdAAContentHint
	case "IDAAMSGSIGDIGEST":
		ret = PkcsObjectIdentifierIdAAMsgSigDigest
	case "IDAACONTENTREFERENCE":
		ret = PkcsObjectIdentifierIdAAContentReference
	case "IDAAENCRYPKEYPREF":
		ret = PkcsObjectIdentifierIdAAEncrypKeyPref
	case "IDAASIGNINGCERTIFICATE":
		ret = PkcsObjectIdentifierIdAASigningCertificate
	case "IDAASIGNINGCERTIFICATEV2":
		ret = PkcsObjectIdentifierIdAASigningCertificateV2
	case "IDAACONTENTIDENTIFIER":
		ret = PkcsObjectIdentifierIdAAContentIdentifier
	case "IDAASIGNATURETIMESTAMPTOKEN":
		ret = PkcsObjectIdentifierIdAASignatureTimeStampToken
	case "IDAAETSSIGPOLICYID":
		ret = PkcsObjectIdentifierIdAAEtsSigPolicyID
	case "IDAAETSCOMMITMENTTYPE":
		ret = PkcsObjectIdentifierIdAAEtsCommitmentType
	case "IDAAETSSIGNERLOCATION":
		ret = PkcsObjectIdentifierIdAAEtsSignerLocation
	case "IDAAETSSIGNERATTR":
		ret = PkcsObjectIdentifierIdAAEtsSignerAttr
	case "IDAAETSOTHERSIGCERT":
		ret = PkcsObjectIdentifierIdAAEtsOtherSigCert
	case "IDAAETSCONTENTTIMESTAMP":
		ret = PkcsObjectIdentifierIdAAEtsContentTimestamp
	case "IDAAETSCERTIFICATEREFS":
		ret = PkcsObjectIdentifierIdAAEtsCertificateRefs
	case "IDAAETSREVOCATIONREFS":
		ret = PkcsObjectIdentifierIdAAEtsRevocationRefs
	case "IDAAETSCERTVALUES":
		ret = PkcsObjectIdentifierIdAAEtsCertValues
	case "IDAAETSREVOCATIONVALUES":
		ret = PkcsObjectIdentifierIdAAEtsRevocationValues
	case "IDAAETSESCTIMESTAMP":
		ret = PkcsObjectIdentifierIdAAEtsEscTimeStamp
	case "IDAAETSCERTCRLTIMESTAMP":
		ret = PkcsObjectIdentifierIdAAEtsCertCrlTimestamp
	case "IDAAETSARCHIVETIMESTAMP":
		ret = PkcsObjectIdentifierIdAAEtsArchiveTimestamp
	case "IDSPQETSURI":
		ret = PkcsObjectIdentifierIdSpqEtsUri
	case "IDSPQETSUNOTICE":
		ret = PkcsObjectIdentifierIdSpqEtsUNotice
	case "KEYBAG":
		ret = PkcsObjectIdentifierKeyBag
	case "PKCS8SHROUDEDKEYBAG":
		ret = PkcsObjectIdentifierPkcs8ShroudedKeyBag
	case "CERTBAG":
		ret = PkcsObjectIdentifierCertBag
	case "CRLBAG":
		ret = PkcsObjectIdentifierCrlBag
	case "SECRETBAG":
		ret = PkcsObjectIdentifierSecretBag
	case "SAFECONTENTSBAG":
		ret = PkcsObjectIdentifierSafeContentsBag
	case "PBEWITHSHAAND128BITRC4":
		ret = PkcsObjectIdentifierPbeWithShaAnd128BitRC4
	case "PBEWITHSHAAND40BITRC4":
		ret = PkcsObjectIdentifierPbeWithShaAnd40BitRC4
	case "PBEWITHSHAAND3KEYTRIPLEDESCBC":
		ret = PkcsObjectIdentifierPbeWithShaAnd3KeyTripleDesCbc
	case "PBEWITHSHAAND2KEYTRIPLEDESCBC":
		ret = PkcsObjectIdentifierPbeWithShaAnd2KeyTripleDesCbc
	case "PBEWITHSHAAND128BITRC2CBC":
		ret = PkcsObjectIdentifierPbeWithShaAnd128BitRC2Cbc
	case "PBEWITHSHAAND40BITRC2CBC":
		ret = PkcsObjectIdentifierPbewithShaAnd40BitRC2Cbc
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the PkcsObjectIdentifier.
// It satisfies fmt.Stringer.
func (g PkcsObjectIdentifier) String() string {
	var ret string
	switch g {
	case PkcsObjectIdentifierNone:
		ret = "NONE"
	case PkcsObjectIdentifierRsaEncryption:
		ret = "RSAENCRYPTION"
	case PkcsObjectIdentifierMD2WithRsaEncryption:
		ret = "MD2WITHRSAENCRYPTION"
	case PkcsObjectIdentifierMD4WithRsaEncryption:
		ret = "MD4WITHRSAENCRYPTION"
	case PkcsObjectIdentifierMD5WithRsaEncryption:
		ret = "MD5WITHRSAENCRYPTION"
	case PkcsObjectIdentifierSha1WithRsaEncryption:
		ret = "SHA1WITHRSAENCRYPTION"
	case PkcsObjectIdentifierSrsaOaepEncryptionSet:
		ret = "SRSAOAEPENCRYPTIONSET"
	case PkcsObjectIdentifierIdRsaesOaep:
		ret = "IDRSAESOAEP"
	case PkcsObjectIdentifierIdMgf1:
		ret = "IDMGF1"
	case PkcsObjectIdentifierIdPSpecified:
		ret = "IDPSPECIFIED"
	case PkcsObjectIdentifierIdRsassaPss:
		ret = "IDRSASSAPSS"
	case PkcsObjectIdentifierSha256WithRsaEncryption:
		ret = "SHA256WITHRSAENCRYPTION"
	case PkcsObjectIdentifierSha384WithRsaEncryption:
		ret = "SHA384WITHRSAENCRYPTION"
	case PkcsObjectIdentifierSha512WithRsaEncryption:
		ret = "SHA512WITHRSAENCRYPTION"
	case PkcsObjectIdentifierSha224WithRsaEncryption:
		ret = "SHA224WITHRSAENCRYPTION"
	case PkcsObjectIdentifierDhKeyAgree1ment:
		ret = "DHKEYAGREE1MENT"
	case PkcsObjectIdentifierPbeWithMD2AndDesCbc:
		ret = "PBEWITHMD2ANDDESCBC"
	case PkcsObjectIdentifierPbeWithMD2AndRC2Cbc:
		ret = "PBEWITHMD2ANDRC2CBC"
	case PkcsObjectIdentifierPbeWithMD5AndDesCbc:
		ret = "PBEWITHMD5ANDDESCBC"
	case PkcsObjectIdentifierPbeWithMD5AndRC2Cbc:
		ret = "PBEWITHMD5ANDRC2CBC"
	case PkcsObjectIdentifierPbeWithSha1AndDesCbc:
		ret = "PBEWITHSHA1ANDDESCBC"
	case PkcsObjectIdentifierPbeWithSha1AndRC2Cbc:
		ret = "PBEWITHSHA1ANDRC2CBC"
	case PkcsObjectIdentifierIdPbeS2:
		ret = "IDPBES2"
	case PkcsObjectIdentifierIdPbkdf2:
		ret = "IDPBKDF2"
	case PkcsObjectIdentifierDesEde3Cbc:
		ret = "DESEDE3CBC"
	case PkcsObjectIdentifierRC2Cbc:
		ret = "RC2CBC"
	case PkcsObjectIdentifierMD2:
		ret = "MD2"
	case PkcsObjectIdentifierMD4:
		ret = "MD4"
	case PkcsObjectIdentifierMD5:
		ret = "MD5"
	case PkcsObjectIdentifierIdHmacWithSha1:
		ret = "IDHMACWITHSHA1"
	case PkcsObjectIdentifierIdHmacWithSha224:
		ret = "IDHMACWITHSHA224"
	case PkcsObjectIdentifierIdHmacWithSha256:
		ret = "IDHMACWITHSHA256"
	case PkcsObjectIdentifierIdHmacWithSha384:
		ret = "IDHMACWITHSHA384"
	case PkcsObjectIdentifierIdHmacWithSha512:
		ret = "IDHMACWITHSHA512"
	case PkcsObjectIdentifierData:
		ret = "DATA"
	case PkcsObjectIdentifierSignedData:
		ret = "SIGNEDDATA"
	case PkcsObjectIdentifierEnvelopedData:
		ret = "ENVELOPEDDATA"
	case PkcsObjectIdentifierSignedAndEnvelopedData:
		ret = "SIGNEDANDENVELOPEDDATA"
	case PkcsObjectIdentifierDigestedData:
		ret = "DIGESTEDDATA"
	case PkcsObjectIdentifierEncryptedData:
		ret = "ENCRYPTEDDATA"
	case PkcsObjectIdentifierPkcs9AtEmailAddress:
		ret = "PKCS9ATEMAILADDRESS"
	case PkcsObjectIdentifierPkcs9AtUnstructuredName:
		ret = "PKCS9ATUNSTRUCTUREDNAME"
	case PkcsObjectIdentifierPkcs9AtContentType:
		ret = "PKCS9ATCONTENTTYPE"
	case PkcsObjectIdentifierPkcs9AtMessageDigest:
		ret = "PKCS9ATMESSAGEDIGEST"
	case PkcsObjectIdentifierPkcs9AtSigningTime:
		ret = "PKCS9ATSIGNINGTIME"
	case PkcsObjectIdentifierPkcs9AtCounterSignature:
		ret = "PKCS9ATCOUNTERSIGNATURE"
	case PkcsObjectIdentifierPkcs9AtChallengePassword:
		ret = "PKCS9ATCHALLENGEPASSWORD"
	case PkcsObjectIdentifierPkcs9AtUnstructuredAddress:
		ret = "PKCS9ATUNSTRUCTUREDADDRESS"
	case PkcsObjectIdentifierPkcs9AtExtendedCertificateAttributes:
		ret = "PKCS9ATEXTENDEDCERTIFICATEATTRIBUTES"
	case PkcsObjectIdentifierPkcs9AtSigningDescription:
		ret = "PKCS9ATSIGNINGDESCRIPTION"
	case PkcsObjectIdentifierPkcs9AtExtensionRequest:
		ret = "PKCS9ATEXTENSIONREQUEST"
	case PkcsObjectIdentifierPkcs9AtSmimeCapabilities:
		ret = "PKCS9ATSMIMECAPABILITIES"
	case PkcsObjectIdentifierIdSmime:
		ret = "IDSMIME"
	case PkcsObjectIdentifierPkcs9AtFriendlyName:
		ret = "PKCS9ATFRIENDLYNAME"
	case PkcsObjectIdentifierPkcs9AtLocalKeyID:
		ret = "PKCS9ATLOCALKEYID"
	case PkcsObjectIdentifierX509Certificate:
		ret = "X509CERTIFICATE"
	case PkcsObjectIdentifierSdsiCertificate:
		ret = "SDSICERTIFICATE"
	case PkcsObjectIdentifierX509Crl:
		ret = "X509CRL"
	case PkcsObjectIdentifierIdAlg:
		ret = "IDALG"
	case PkcsObjectIdentifierIdAlgEsdh:
		ret = "IDALGESDH"
	case PkcsObjectIdentifierIdAlgCms3DesWrap:
		ret = "IDALGCMS3DESWRAP"
	case PkcsObjectIdentifierIdAlgCmsRC2Wrap:
		ret = "IDALGCMSRC2WRAP"
	case PkcsObjectIdentifierIdAlgPwriKek:
		ret = "IDALGPWRIKEK"
	case PkcsObjectIdentifierIdAlgSsdh:
		ret = "IDALGSSDH"
	case PkcsObjectIdentifierIdRsaKem:
		ret = "IDRSAKEM"
	case PkcsObjectIdentifierPreferSignedData:
		ret = "PREFERSIGNEDDATA"
	case PkcsObjectIdentifierCannotDecryptAny:
		ret = "CANNOTDECRYPTANY"
	case PkcsObjectIdentifierSmimeCapabilitiesVersions:
		ret = "SMIMECAPABILITIESVERSIONS"
	case PkcsObjectIdentifierIdAAReceiptRequest:
		ret = "IDAARECEIPTREQUEST"
	case PkcsObjectIdentifierIdCTAuthData:
		ret = "IDCTAUTHDATA"
	case PkcsObjectIdentifierIdCTTstInfo:
		ret = "IDCTTSTINFO"
	case PkcsObjectIdentifierIdCTCompressedData:
		ret = "IDCTCOMPRESSEDDATA"
	case PkcsObjectIdentifierIdCTAuthEnvelopedData:
		ret = "IDCTAUTHENVELOPEDDATA"
	case PkcsObjectIdentifierIdCTTimestampedData:
		ret = "IDCTTIMESTAMPEDDATA"
	case PkcsObjectIdentifierIdCtiEtsProofOfOrigin:
		ret = "IDCTIETSPROOFOFORIGIN"
	case PkcsObjectIdentifierIdCtiEtsProofOfReceipt:
		ret = "IDCTIETSPROOFOFRECEIPT"
	case PkcsObjectIdentifierIdCtiEtsProofOfDelivery:
		ret = "IDCTIETSPROOFOFDELIVERY"
	case PkcsObjectIdentifierIdCtiEtsProofOfSender:
		ret = "IDCTIETSPROOFOFSENDER"
	case PkcsObjectIdentifierIdCtiEtsProofOfApproval:
		ret = "IDCTIETSPROOFOFAPPROVAL"
	case PkcsObjectIdentifierIdCtiEtsProofOfCreation:
		ret = "IDCTIETSPROOFOFCREATION"
	case PkcsObjectIdentifierIdAAContentHint:
		ret = "IDAACONTENTHINT"
	case PkcsObjectIdentifierIdAAMsgSigDigest:
		ret = "IDAAMSGSIGDIGEST"
	case PkcsObjectIdentifierIdAAContentReference:
		ret = "IDAACONTENTREFERENCE"
	case PkcsObjectIdentifierIdAAEncrypKeyPref:
		ret = "IDAAENCRYPKEYPREF"
	case PkcsObjectIdentifierIdAASigningCertificate:
		ret = "IDAASIGNINGCERTIFICATE"
	case PkcsObjectIdentifierIdAASigningCertificateV2:
		ret = "IDAASIGNINGCERTIFICATEV2"
	case PkcsObjectIdentifierIdAAContentIdentifier:
		ret = "IDAACONTENTIDENTIFIER"
	case PkcsObjectIdentifierIdAASignatureTimeStampToken:
		ret = "IDAASIGNATURETIMESTAMPTOKEN"
	case PkcsObjectIdentifierIdAAEtsSigPolicyID:
		ret = "IDAAETSSIGPOLICYID"
	case PkcsObjectIdentifierIdAAEtsCommitmentType:
		ret = "IDAAETSCOMMITMENTTYPE"
	case PkcsObjectIdentifierIdAAEtsSignerLocation:
		ret = "IDAAETSSIGNERLOCATION"
	case PkcsObjectIdentifierIdAAEtsSignerAttr:
		ret = "IDAAETSSIGNERATTR"
	case PkcsObjectIdentifierIdAAEtsOtherSigCert:
		ret = "IDAAETSOTHERSIGCERT"
	case PkcsObjectIdentifierIdAAEtsContentTimestamp:
		ret = "IDAAETSCONTENTTIMESTAMP"
	case PkcsObjectIdentifierIdAAEtsCertificateRefs:
		ret = "IDAAETSCERTIFICATEREFS"
	case PkcsObjectIdentifierIdAAEtsRevocationRefs:
		ret = "IDAAETSREVOCATIONREFS"
	case PkcsObjectIdentifierIdAAEtsCertValues:
		ret = "IDAAETSCERTVALUES"
	case PkcsObjectIdentifierIdAAEtsRevocationValues:
		ret = "IDAAETSREVOCATIONVALUES"
	case PkcsObjectIdentifierIdAAEtsEscTimeStamp:
		ret = "IDAAETSESCTIMESTAMP"
	case PkcsObjectIdentifierIdAAEtsCertCrlTimestamp:
		ret = "IDAAETSCERTCRLTIMESTAMP"
	case PkcsObjectIdentifierIdAAEtsArchiveTimestamp:
		ret = "IDAAETSARCHIVETIMESTAMP"
	case PkcsObjectIdentifierIdSpqEtsUri:
		ret = "IDSPQETSURI"
	case PkcsObjectIdentifierIdSpqEtsUNotice:
		ret = "IDSPQETSUNOTICE"
	case PkcsObjectIdentifierKeyBag:
		ret = "KEYBAG"
	case PkcsObjectIdentifierPkcs8ShroudedKeyBag:
		ret = "PKCS8SHROUDEDKEYBAG"
	case PkcsObjectIdentifierCertBag:
		ret = "CERTBAG"
	case PkcsObjectIdentifierCrlBag:
		ret = "CRLBAG"
	case PkcsObjectIdentifierSecretBag:
		ret = "SECRETBAG"
	case PkcsObjectIdentifierSafeContentsBag:
		ret = "SAFECONTENTSBAG"
	case PkcsObjectIdentifierPbeWithShaAnd128BitRC4:
		ret = "PBEWITHSHAAND128BITRC4"
	case PkcsObjectIdentifierPbeWithShaAnd40BitRC4:
		ret = "PBEWITHSHAAND40BITRC4"
	case PkcsObjectIdentifierPbeWithShaAnd3KeyTripleDesCbc:
		ret = "PBEWITHSHAAND3KEYTRIPLEDESCBC"
	case PkcsObjectIdentifierPbeWithShaAnd2KeyTripleDesCbc:
		ret = "PBEWITHSHAAND2KEYTRIPLEDESCBC"
	case PkcsObjectIdentifierPbeWithShaAnd128BitRC2Cbc:
		ret = "PBEWITHSHAAND128BITRC2CBC"
	case PkcsObjectIdentifierPbewithShaAnd40BitRC2Cbc:
		ret = "PBEWITHSHAAND40BITRC2CBC"
	}
	return ret
}
