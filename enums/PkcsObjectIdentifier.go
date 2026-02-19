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
	switch {
	case strings.EqualFold(value, "None"):
		ret = PkcsObjectIdentifierNone
	case strings.EqualFold(value, "RsaEncryption"):
		ret = PkcsObjectIdentifierRsaEncryption
	case strings.EqualFold(value, "MD2WithRsaEncryption"):
		ret = PkcsObjectIdentifierMD2WithRsaEncryption
	case strings.EqualFold(value, "MD4WithRsaEncryption"):
		ret = PkcsObjectIdentifierMD4WithRsaEncryption
	case strings.EqualFold(value, "MD5WithRsaEncryption"):
		ret = PkcsObjectIdentifierMD5WithRsaEncryption
	case strings.EqualFold(value, "Sha1WithRsaEncryption"):
		ret = PkcsObjectIdentifierSha1WithRsaEncryption
	case strings.EqualFold(value, "SrsaOaepEncryptionSet"):
		ret = PkcsObjectIdentifierSrsaOaepEncryptionSet
	case strings.EqualFold(value, "IdRsaesOaep"):
		ret = PkcsObjectIdentifierIdRsaesOaep
	case strings.EqualFold(value, "IdMgf1"):
		ret = PkcsObjectIdentifierIdMgf1
	case strings.EqualFold(value, "IdPSpecified"):
		ret = PkcsObjectIdentifierIdPSpecified
	case strings.EqualFold(value, "IdRsassaPss"):
		ret = PkcsObjectIdentifierIdRsassaPss
	case strings.EqualFold(value, "Sha256WithRsaEncryption"):
		ret = PkcsObjectIdentifierSha256WithRsaEncryption
	case strings.EqualFold(value, "Sha384WithRsaEncryption"):
		ret = PkcsObjectIdentifierSha384WithRsaEncryption
	case strings.EqualFold(value, "Sha512WithRsaEncryption"):
		ret = PkcsObjectIdentifierSha512WithRsaEncryption
	case strings.EqualFold(value, "Sha224WithRsaEncryption"):
		ret = PkcsObjectIdentifierSha224WithRsaEncryption
	case strings.EqualFold(value, "DhKeyAgree1ment"):
		ret = PkcsObjectIdentifierDhKeyAgree1ment
	case strings.EqualFold(value, "PbeWithMD2AndDesCbc"):
		ret = PkcsObjectIdentifierPbeWithMD2AndDesCbc
	case strings.EqualFold(value, "PbeWithMD2AndRC2Cbc"):
		ret = PkcsObjectIdentifierPbeWithMD2AndRC2Cbc
	case strings.EqualFold(value, "PbeWithMD5AndDesCbc"):
		ret = PkcsObjectIdentifierPbeWithMD5AndDesCbc
	case strings.EqualFold(value, "PbeWithMD5AndRC2Cbc"):
		ret = PkcsObjectIdentifierPbeWithMD5AndRC2Cbc
	case strings.EqualFold(value, "PbeWithSha1AndDesCbc"):
		ret = PkcsObjectIdentifierPbeWithSha1AndDesCbc
	case strings.EqualFold(value, "PbeWithSha1AndRC2Cbc"):
		ret = PkcsObjectIdentifierPbeWithSha1AndRC2Cbc
	case strings.EqualFold(value, "IdPbeS2"):
		ret = PkcsObjectIdentifierIdPbeS2
	case strings.EqualFold(value, "IdPbkdf2"):
		ret = PkcsObjectIdentifierIdPbkdf2
	case strings.EqualFold(value, "DesEde3Cbc"):
		ret = PkcsObjectIdentifierDesEde3Cbc
	case strings.EqualFold(value, "RC2Cbc"):
		ret = PkcsObjectIdentifierRC2Cbc
	case strings.EqualFold(value, "MD2"):
		ret = PkcsObjectIdentifierMD2
	case strings.EqualFold(value, "MD4"):
		ret = PkcsObjectIdentifierMD4
	case strings.EqualFold(value, "MD5"):
		ret = PkcsObjectIdentifierMD5
	case strings.EqualFold(value, "IdHmacWithSha1"):
		ret = PkcsObjectIdentifierIdHmacWithSha1
	case strings.EqualFold(value, "IdHmacWithSha224"):
		ret = PkcsObjectIdentifierIdHmacWithSha224
	case strings.EqualFold(value, "IdHmacWithSha256"):
		ret = PkcsObjectIdentifierIdHmacWithSha256
	case strings.EqualFold(value, "IdHmacWithSha384"):
		ret = PkcsObjectIdentifierIdHmacWithSha384
	case strings.EqualFold(value, "IdHmacWithSha512"):
		ret = PkcsObjectIdentifierIdHmacWithSha512
	case strings.EqualFold(value, "Data"):
		ret = PkcsObjectIdentifierData
	case strings.EqualFold(value, "SignedData"):
		ret = PkcsObjectIdentifierSignedData
	case strings.EqualFold(value, "EnvelopedData"):
		ret = PkcsObjectIdentifierEnvelopedData
	case strings.EqualFold(value, "SignedAndEnvelopedData"):
		ret = PkcsObjectIdentifierSignedAndEnvelopedData
	case strings.EqualFold(value, "DigestedData"):
		ret = PkcsObjectIdentifierDigestedData
	case strings.EqualFold(value, "EncryptedData"):
		ret = PkcsObjectIdentifierEncryptedData
	case strings.EqualFold(value, "Pkcs9AtEmailAddress"):
		ret = PkcsObjectIdentifierPkcs9AtEmailAddress
	case strings.EqualFold(value, "Pkcs9AtUnstructuredName"):
		ret = PkcsObjectIdentifierPkcs9AtUnstructuredName
	case strings.EqualFold(value, "Pkcs9AtContentType"):
		ret = PkcsObjectIdentifierPkcs9AtContentType
	case strings.EqualFold(value, "Pkcs9AtMessageDigest"):
		ret = PkcsObjectIdentifierPkcs9AtMessageDigest
	case strings.EqualFold(value, "Pkcs9AtSigningTime"):
		ret = PkcsObjectIdentifierPkcs9AtSigningTime
	case strings.EqualFold(value, "Pkcs9AtCounterSignature"):
		ret = PkcsObjectIdentifierPkcs9AtCounterSignature
	case strings.EqualFold(value, "Pkcs9AtChallengePassword"):
		ret = PkcsObjectIdentifierPkcs9AtChallengePassword
	case strings.EqualFold(value, "Pkcs9AtUnstructuredAddress"):
		ret = PkcsObjectIdentifierPkcs9AtUnstructuredAddress
	case strings.EqualFold(value, "Pkcs9AtExtendedCertificateAttributes"):
		ret = PkcsObjectIdentifierPkcs9AtExtendedCertificateAttributes
	case strings.EqualFold(value, "Pkcs9AtSigningDescription"):
		ret = PkcsObjectIdentifierPkcs9AtSigningDescription
	case strings.EqualFold(value, "Pkcs9AtExtensionRequest"):
		ret = PkcsObjectIdentifierPkcs9AtExtensionRequest
	case strings.EqualFold(value, "Pkcs9AtSmimeCapabilities"):
		ret = PkcsObjectIdentifierPkcs9AtSmimeCapabilities
	case strings.EqualFold(value, "IdSmime"):
		ret = PkcsObjectIdentifierIdSmime
	case strings.EqualFold(value, "Pkcs9AtFriendlyName"):
		ret = PkcsObjectIdentifierPkcs9AtFriendlyName
	case strings.EqualFold(value, "Pkcs9AtLocalKeyID"):
		ret = PkcsObjectIdentifierPkcs9AtLocalKeyID
	case strings.EqualFold(value, "X509Certificate"):
		ret = PkcsObjectIdentifierX509Certificate
	case strings.EqualFold(value, "SdsiCertificate"):
		ret = PkcsObjectIdentifierSdsiCertificate
	case strings.EqualFold(value, "X509Crl"):
		ret = PkcsObjectIdentifierX509Crl
	case strings.EqualFold(value, "IdAlg"):
		ret = PkcsObjectIdentifierIdAlg
	case strings.EqualFold(value, "IdAlgEsdh"):
		ret = PkcsObjectIdentifierIdAlgEsdh
	case strings.EqualFold(value, "IdAlgCms3DesWrap"):
		ret = PkcsObjectIdentifierIdAlgCms3DesWrap
	case strings.EqualFold(value, "IdAlgCmsRC2Wrap"):
		ret = PkcsObjectIdentifierIdAlgCmsRC2Wrap
	case strings.EqualFold(value, "IdAlgPwriKek"):
		ret = PkcsObjectIdentifierIdAlgPwriKek
	case strings.EqualFold(value, "IdAlgSsdh"):
		ret = PkcsObjectIdentifierIdAlgSsdh
	case strings.EqualFold(value, "IdRsaKem"):
		ret = PkcsObjectIdentifierIdRsaKem
	case strings.EqualFold(value, "PreferSignedData"):
		ret = PkcsObjectIdentifierPreferSignedData
	case strings.EqualFold(value, "CannotDecryptAny"):
		ret = PkcsObjectIdentifierCannotDecryptAny
	case strings.EqualFold(value, "SmimeCapabilitiesVersions"):
		ret = PkcsObjectIdentifierSmimeCapabilitiesVersions
	case strings.EqualFold(value, "IdAAReceiptRequest"):
		ret = PkcsObjectIdentifierIdAAReceiptRequest
	case strings.EqualFold(value, "IdCTAuthData"):
		ret = PkcsObjectIdentifierIdCTAuthData
	case strings.EqualFold(value, "IdCTTstInfo"):
		ret = PkcsObjectIdentifierIdCTTstInfo
	case strings.EqualFold(value, "IdCTCompressedData"):
		ret = PkcsObjectIdentifierIdCTCompressedData
	case strings.EqualFold(value, "IdCTAuthEnvelopedData"):
		ret = PkcsObjectIdentifierIdCTAuthEnvelopedData
	case strings.EqualFold(value, "IdCTTimestampedData"):
		ret = PkcsObjectIdentifierIdCTTimestampedData
	case strings.EqualFold(value, "IdCtiEtsProofOfOrigin"):
		ret = PkcsObjectIdentifierIdCtiEtsProofOfOrigin
	case strings.EqualFold(value, "IdCtiEtsProofOfReceipt"):
		ret = PkcsObjectIdentifierIdCtiEtsProofOfReceipt
	case strings.EqualFold(value, "IdCtiEtsProofOfDelivery"):
		ret = PkcsObjectIdentifierIdCtiEtsProofOfDelivery
	case strings.EqualFold(value, "IdCtiEtsProofOfSender"):
		ret = PkcsObjectIdentifierIdCtiEtsProofOfSender
	case strings.EqualFold(value, "IdCtiEtsProofOfApproval"):
		ret = PkcsObjectIdentifierIdCtiEtsProofOfApproval
	case strings.EqualFold(value, "IdCtiEtsProofOfCreation"):
		ret = PkcsObjectIdentifierIdCtiEtsProofOfCreation
	case strings.EqualFold(value, "IdAAContentHint"):
		ret = PkcsObjectIdentifierIdAAContentHint
	case strings.EqualFold(value, "IdAAMsgSigDigest"):
		ret = PkcsObjectIdentifierIdAAMsgSigDigest
	case strings.EqualFold(value, "IdAAContentReference"):
		ret = PkcsObjectIdentifierIdAAContentReference
	case strings.EqualFold(value, "IdAAEncrypKeyPref"):
		ret = PkcsObjectIdentifierIdAAEncrypKeyPref
	case strings.EqualFold(value, "IdAASigningCertificate"):
		ret = PkcsObjectIdentifierIdAASigningCertificate
	case strings.EqualFold(value, "IdAASigningCertificateV2"):
		ret = PkcsObjectIdentifierIdAASigningCertificateV2
	case strings.EqualFold(value, "IdAAContentIdentifier"):
		ret = PkcsObjectIdentifierIdAAContentIdentifier
	case strings.EqualFold(value, "IdAASignatureTimeStampToken"):
		ret = PkcsObjectIdentifierIdAASignatureTimeStampToken
	case strings.EqualFold(value, "IdAAEtsSigPolicyID"):
		ret = PkcsObjectIdentifierIdAAEtsSigPolicyID
	case strings.EqualFold(value, "IdAAEtsCommitmentType"):
		ret = PkcsObjectIdentifierIdAAEtsCommitmentType
	case strings.EqualFold(value, "IdAAEtsSignerLocation"):
		ret = PkcsObjectIdentifierIdAAEtsSignerLocation
	case strings.EqualFold(value, "IdAAEtsSignerAttr"):
		ret = PkcsObjectIdentifierIdAAEtsSignerAttr
	case strings.EqualFold(value, "IdAAEtsOtherSigCert"):
		ret = PkcsObjectIdentifierIdAAEtsOtherSigCert
	case strings.EqualFold(value, "IdAAEtsContentTimestamp"):
		ret = PkcsObjectIdentifierIdAAEtsContentTimestamp
	case strings.EqualFold(value, "IdAAEtsCertificateRefs"):
		ret = PkcsObjectIdentifierIdAAEtsCertificateRefs
	case strings.EqualFold(value, "IdAAEtsRevocationRefs"):
		ret = PkcsObjectIdentifierIdAAEtsRevocationRefs
	case strings.EqualFold(value, "IdAAEtsCertValues"):
		ret = PkcsObjectIdentifierIdAAEtsCertValues
	case strings.EqualFold(value, "IdAAEtsRevocationValues"):
		ret = PkcsObjectIdentifierIdAAEtsRevocationValues
	case strings.EqualFold(value, "IdAAEtsEscTimeStamp"):
		ret = PkcsObjectIdentifierIdAAEtsEscTimeStamp
	case strings.EqualFold(value, "IdAAEtsCertCrlTimestamp"):
		ret = PkcsObjectIdentifierIdAAEtsCertCrlTimestamp
	case strings.EqualFold(value, "IdAAEtsArchiveTimestamp"):
		ret = PkcsObjectIdentifierIdAAEtsArchiveTimestamp
	case strings.EqualFold(value, "IdSpqEtsUri"):
		ret = PkcsObjectIdentifierIdSpqEtsUri
	case strings.EqualFold(value, "IdSpqEtsUNotice"):
		ret = PkcsObjectIdentifierIdSpqEtsUNotice
	case strings.EqualFold(value, "KeyBag"):
		ret = PkcsObjectIdentifierKeyBag
	case strings.EqualFold(value, "Pkcs8ShroudedKeyBag"):
		ret = PkcsObjectIdentifierPkcs8ShroudedKeyBag
	case strings.EqualFold(value, "CertBag"):
		ret = PkcsObjectIdentifierCertBag
	case strings.EqualFold(value, "CrlBag"):
		ret = PkcsObjectIdentifierCrlBag
	case strings.EqualFold(value, "SecretBag"):
		ret = PkcsObjectIdentifierSecretBag
	case strings.EqualFold(value, "SafeContentsBag"):
		ret = PkcsObjectIdentifierSafeContentsBag
	case strings.EqualFold(value, "PbeWithShaAnd128BitRC4"):
		ret = PkcsObjectIdentifierPbeWithShaAnd128BitRC4
	case strings.EqualFold(value, "PbeWithShaAnd40BitRC4"):
		ret = PkcsObjectIdentifierPbeWithShaAnd40BitRC4
	case strings.EqualFold(value, "PbeWithShaAnd3KeyTripleDesCbc"):
		ret = PkcsObjectIdentifierPbeWithShaAnd3KeyTripleDesCbc
	case strings.EqualFold(value, "PbeWithShaAnd2KeyTripleDesCbc"):
		ret = PkcsObjectIdentifierPbeWithShaAnd2KeyTripleDesCbc
	case strings.EqualFold(value, "PbeWithShaAnd128BitRC2Cbc"):
		ret = PkcsObjectIdentifierPbeWithShaAnd128BitRC2Cbc
	case strings.EqualFold(value, "PbewithShaAnd40BitRC2Cbc"):
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
		ret = "None"
	case PkcsObjectIdentifierRsaEncryption:
		ret = "RsaEncryption"
	case PkcsObjectIdentifierMD2WithRsaEncryption:
		ret = "MD2WithRsaEncryption"
	case PkcsObjectIdentifierMD4WithRsaEncryption:
		ret = "MD4WithRsaEncryption"
	case PkcsObjectIdentifierMD5WithRsaEncryption:
		ret = "MD5WithRsaEncryption"
	case PkcsObjectIdentifierSha1WithRsaEncryption:
		ret = "Sha1WithRsaEncryption"
	case PkcsObjectIdentifierSrsaOaepEncryptionSet:
		ret = "SrsaOaepEncryptionSet"
	case PkcsObjectIdentifierIdRsaesOaep:
		ret = "IdRsaesOaep"
	case PkcsObjectIdentifierIdMgf1:
		ret = "IdMgf1"
	case PkcsObjectIdentifierIdPSpecified:
		ret = "IdPSpecified"
	case PkcsObjectIdentifierIdRsassaPss:
		ret = "IdRsassaPss"
	case PkcsObjectIdentifierSha256WithRsaEncryption:
		ret = "Sha256WithRsaEncryption"
	case PkcsObjectIdentifierSha384WithRsaEncryption:
		ret = "Sha384WithRsaEncryption"
	case PkcsObjectIdentifierSha512WithRsaEncryption:
		ret = "Sha512WithRsaEncryption"
	case PkcsObjectIdentifierSha224WithRsaEncryption:
		ret = "Sha224WithRsaEncryption"
	case PkcsObjectIdentifierDhKeyAgree1ment:
		ret = "DhKeyAgree1ment"
	case PkcsObjectIdentifierPbeWithMD2AndDesCbc:
		ret = "PbeWithMD2AndDesCbc"
	case PkcsObjectIdentifierPbeWithMD2AndRC2Cbc:
		ret = "PbeWithMD2AndRC2Cbc"
	case PkcsObjectIdentifierPbeWithMD5AndDesCbc:
		ret = "PbeWithMD5AndDesCbc"
	case PkcsObjectIdentifierPbeWithMD5AndRC2Cbc:
		ret = "PbeWithMD5AndRC2Cbc"
	case PkcsObjectIdentifierPbeWithSha1AndDesCbc:
		ret = "PbeWithSha1AndDesCbc"
	case PkcsObjectIdentifierPbeWithSha1AndRC2Cbc:
		ret = "PbeWithSha1AndRC2Cbc"
	case PkcsObjectIdentifierIdPbeS2:
		ret = "IdPbeS2"
	case PkcsObjectIdentifierIdPbkdf2:
		ret = "IdPbkdf2"
	case PkcsObjectIdentifierDesEde3Cbc:
		ret = "DesEde3Cbc"
	case PkcsObjectIdentifierRC2Cbc:
		ret = "RC2Cbc"
	case PkcsObjectIdentifierMD2:
		ret = "MD2"
	case PkcsObjectIdentifierMD4:
		ret = "MD4"
	case PkcsObjectIdentifierMD5:
		ret = "MD5"
	case PkcsObjectIdentifierIdHmacWithSha1:
		ret = "IdHmacWithSha1"
	case PkcsObjectIdentifierIdHmacWithSha224:
		ret = "IdHmacWithSha224"
	case PkcsObjectIdentifierIdHmacWithSha256:
		ret = "IdHmacWithSha256"
	case PkcsObjectIdentifierIdHmacWithSha384:
		ret = "IdHmacWithSha384"
	case PkcsObjectIdentifierIdHmacWithSha512:
		ret = "IdHmacWithSha512"
	case PkcsObjectIdentifierData:
		ret = "Data"
	case PkcsObjectIdentifierSignedData:
		ret = "SignedData"
	case PkcsObjectIdentifierEnvelopedData:
		ret = "EnvelopedData"
	case PkcsObjectIdentifierSignedAndEnvelopedData:
		ret = "SignedAndEnvelopedData"
	case PkcsObjectIdentifierDigestedData:
		ret = "DigestedData"
	case PkcsObjectIdentifierEncryptedData:
		ret = "EncryptedData"
	case PkcsObjectIdentifierPkcs9AtEmailAddress:
		ret = "Pkcs9AtEmailAddress"
	case PkcsObjectIdentifierPkcs9AtUnstructuredName:
		ret = "Pkcs9AtUnstructuredName"
	case PkcsObjectIdentifierPkcs9AtContentType:
		ret = "Pkcs9AtContentType"
	case PkcsObjectIdentifierPkcs9AtMessageDigest:
		ret = "Pkcs9AtMessageDigest"
	case PkcsObjectIdentifierPkcs9AtSigningTime:
		ret = "Pkcs9AtSigningTime"
	case PkcsObjectIdentifierPkcs9AtCounterSignature:
		ret = "Pkcs9AtCounterSignature"
	case PkcsObjectIdentifierPkcs9AtChallengePassword:
		ret = "Pkcs9AtChallengePassword"
	case PkcsObjectIdentifierPkcs9AtUnstructuredAddress:
		ret = "Pkcs9AtUnstructuredAddress"
	case PkcsObjectIdentifierPkcs9AtExtendedCertificateAttributes:
		ret = "Pkcs9AtExtendedCertificateAttributes"
	case PkcsObjectIdentifierPkcs9AtSigningDescription:
		ret = "Pkcs9AtSigningDescription"
	case PkcsObjectIdentifierPkcs9AtExtensionRequest:
		ret = "Pkcs9AtExtensionRequest"
	case PkcsObjectIdentifierPkcs9AtSmimeCapabilities:
		ret = "Pkcs9AtSmimeCapabilities"
	case PkcsObjectIdentifierIdSmime:
		ret = "IdSmime"
	case PkcsObjectIdentifierPkcs9AtFriendlyName:
		ret = "Pkcs9AtFriendlyName"
	case PkcsObjectIdentifierPkcs9AtLocalKeyID:
		ret = "Pkcs9AtLocalKeyID"
	case PkcsObjectIdentifierX509Certificate:
		ret = "X509Certificate"
	case PkcsObjectIdentifierSdsiCertificate:
		ret = "SdsiCertificate"
	case PkcsObjectIdentifierX509Crl:
		ret = "X509Crl"
	case PkcsObjectIdentifierIdAlg:
		ret = "IdAlg"
	case PkcsObjectIdentifierIdAlgEsdh:
		ret = "IdAlgEsdh"
	case PkcsObjectIdentifierIdAlgCms3DesWrap:
		ret = "IdAlgCms3DesWrap"
	case PkcsObjectIdentifierIdAlgCmsRC2Wrap:
		ret = "IdAlgCmsRC2Wrap"
	case PkcsObjectIdentifierIdAlgPwriKek:
		ret = "IdAlgPwriKek"
	case PkcsObjectIdentifierIdAlgSsdh:
		ret = "IdAlgSsdh"
	case PkcsObjectIdentifierIdRsaKem:
		ret = "IdRsaKem"
	case PkcsObjectIdentifierPreferSignedData:
		ret = "PreferSignedData"
	case PkcsObjectIdentifierCannotDecryptAny:
		ret = "CannotDecryptAny"
	case PkcsObjectIdentifierSmimeCapabilitiesVersions:
		ret = "SmimeCapabilitiesVersions"
	case PkcsObjectIdentifierIdAAReceiptRequest:
		ret = "IdAAReceiptRequest"
	case PkcsObjectIdentifierIdCTAuthData:
		ret = "IdCTAuthData"
	case PkcsObjectIdentifierIdCTTstInfo:
		ret = "IdCTTstInfo"
	case PkcsObjectIdentifierIdCTCompressedData:
		ret = "IdCTCompressedData"
	case PkcsObjectIdentifierIdCTAuthEnvelopedData:
		ret = "IdCTAuthEnvelopedData"
	case PkcsObjectIdentifierIdCTTimestampedData:
		ret = "IdCTTimestampedData"
	case PkcsObjectIdentifierIdCtiEtsProofOfOrigin:
		ret = "IdCtiEtsProofOfOrigin"
	case PkcsObjectIdentifierIdCtiEtsProofOfReceipt:
		ret = "IdCtiEtsProofOfReceipt"
	case PkcsObjectIdentifierIdCtiEtsProofOfDelivery:
		ret = "IdCtiEtsProofOfDelivery"
	case PkcsObjectIdentifierIdCtiEtsProofOfSender:
		ret = "IdCtiEtsProofOfSender"
	case PkcsObjectIdentifierIdCtiEtsProofOfApproval:
		ret = "IdCtiEtsProofOfApproval"
	case PkcsObjectIdentifierIdCtiEtsProofOfCreation:
		ret = "IdCtiEtsProofOfCreation"
	case PkcsObjectIdentifierIdAAContentHint:
		ret = "IdAAContentHint"
	case PkcsObjectIdentifierIdAAMsgSigDigest:
		ret = "IdAAMsgSigDigest"
	case PkcsObjectIdentifierIdAAContentReference:
		ret = "IdAAContentReference"
	case PkcsObjectIdentifierIdAAEncrypKeyPref:
		ret = "IdAAEncrypKeyPref"
	case PkcsObjectIdentifierIdAASigningCertificate:
		ret = "IdAASigningCertificate"
	case PkcsObjectIdentifierIdAASigningCertificateV2:
		ret = "IdAASigningCertificateV2"
	case PkcsObjectIdentifierIdAAContentIdentifier:
		ret = "IdAAContentIdentifier"
	case PkcsObjectIdentifierIdAASignatureTimeStampToken:
		ret = "IdAASignatureTimeStampToken"
	case PkcsObjectIdentifierIdAAEtsSigPolicyID:
		ret = "IdAAEtsSigPolicyID"
	case PkcsObjectIdentifierIdAAEtsCommitmentType:
		ret = "IdAAEtsCommitmentType"
	case PkcsObjectIdentifierIdAAEtsSignerLocation:
		ret = "IdAAEtsSignerLocation"
	case PkcsObjectIdentifierIdAAEtsSignerAttr:
		ret = "IdAAEtsSignerAttr"
	case PkcsObjectIdentifierIdAAEtsOtherSigCert:
		ret = "IdAAEtsOtherSigCert"
	case PkcsObjectIdentifierIdAAEtsContentTimestamp:
		ret = "IdAAEtsContentTimestamp"
	case PkcsObjectIdentifierIdAAEtsCertificateRefs:
		ret = "IdAAEtsCertificateRefs"
	case PkcsObjectIdentifierIdAAEtsRevocationRefs:
		ret = "IdAAEtsRevocationRefs"
	case PkcsObjectIdentifierIdAAEtsCertValues:
		ret = "IdAAEtsCertValues"
	case PkcsObjectIdentifierIdAAEtsRevocationValues:
		ret = "IdAAEtsRevocationValues"
	case PkcsObjectIdentifierIdAAEtsEscTimeStamp:
		ret = "IdAAEtsEscTimeStamp"
	case PkcsObjectIdentifierIdAAEtsCertCrlTimestamp:
		ret = "IdAAEtsCertCrlTimestamp"
	case PkcsObjectIdentifierIdAAEtsArchiveTimestamp:
		ret = "IdAAEtsArchiveTimestamp"
	case PkcsObjectIdentifierIdSpqEtsUri:
		ret = "IdSpqEtsUri"
	case PkcsObjectIdentifierIdSpqEtsUNotice:
		ret = "IdSpqEtsUNotice"
	case PkcsObjectIdentifierKeyBag:
		ret = "KeyBag"
	case PkcsObjectIdentifierPkcs8ShroudedKeyBag:
		ret = "Pkcs8ShroudedKeyBag"
	case PkcsObjectIdentifierCertBag:
		ret = "CertBag"
	case PkcsObjectIdentifierCrlBag:
		ret = "CrlBag"
	case PkcsObjectIdentifierSecretBag:
		ret = "SecretBag"
	case PkcsObjectIdentifierSafeContentsBag:
		ret = "SafeContentsBag"
	case PkcsObjectIdentifierPbeWithShaAnd128BitRC4:
		ret = "PbeWithShaAnd128BitRC4"
	case PkcsObjectIdentifierPbeWithShaAnd40BitRC4:
		ret = "PbeWithShaAnd40BitRC4"
	case PkcsObjectIdentifierPbeWithShaAnd3KeyTripleDesCbc:
		ret = "PbeWithShaAnd3KeyTripleDesCbc"
	case PkcsObjectIdentifierPbeWithShaAnd2KeyTripleDesCbc:
		ret = "PbeWithShaAnd2KeyTripleDesCbc"
	case PkcsObjectIdentifierPbeWithShaAnd128BitRC2Cbc:
		ret = "PbeWithShaAnd128BitRC2Cbc"
	case PkcsObjectIdentifierPbewithShaAnd40BitRC2Cbc:
		ret = "PbewithShaAnd40BitRC2Cbc"
	}
	return ret
}

// AllPkcsObjectIdentifier returns a slice containing all defined PkcsObjectIdentifier values.
func AllPkcsObjectIdentifier() []PkcsObjectIdentifier {
	return []PkcsObjectIdentifier{
	PkcsObjectIdentifierNone,
	PkcsObjectIdentifierRsaEncryption,
	PkcsObjectIdentifierMD2WithRsaEncryption,
	PkcsObjectIdentifierMD4WithRsaEncryption,
	PkcsObjectIdentifierMD5WithRsaEncryption,
	PkcsObjectIdentifierSha1WithRsaEncryption,
	PkcsObjectIdentifierSrsaOaepEncryptionSet,
	PkcsObjectIdentifierIdRsaesOaep,
	PkcsObjectIdentifierIdMgf1,
	PkcsObjectIdentifierIdPSpecified,
	PkcsObjectIdentifierIdRsassaPss,
	PkcsObjectIdentifierSha256WithRsaEncryption,
	PkcsObjectIdentifierSha384WithRsaEncryption,
	PkcsObjectIdentifierSha512WithRsaEncryption,
	PkcsObjectIdentifierSha224WithRsaEncryption,
	PkcsObjectIdentifierDhKeyAgree1ment,
	PkcsObjectIdentifierPbeWithMD2AndDesCbc,
	PkcsObjectIdentifierPbeWithMD2AndRC2Cbc,
	PkcsObjectIdentifierPbeWithMD5AndDesCbc,
	PkcsObjectIdentifierPbeWithMD5AndRC2Cbc,
	PkcsObjectIdentifierPbeWithSha1AndDesCbc,
	PkcsObjectIdentifierPbeWithSha1AndRC2Cbc,
	PkcsObjectIdentifierIdPbeS2,
	PkcsObjectIdentifierIdPbkdf2,
	PkcsObjectIdentifierDesEde3Cbc,
	PkcsObjectIdentifierRC2Cbc,
	PkcsObjectIdentifierMD2,
	PkcsObjectIdentifierMD4,
	PkcsObjectIdentifierMD5,
	PkcsObjectIdentifierIdHmacWithSha1,
	PkcsObjectIdentifierIdHmacWithSha224,
	PkcsObjectIdentifierIdHmacWithSha256,
	PkcsObjectIdentifierIdHmacWithSha384,
	PkcsObjectIdentifierIdHmacWithSha512,
	PkcsObjectIdentifierData,
	PkcsObjectIdentifierSignedData,
	PkcsObjectIdentifierEnvelopedData,
	PkcsObjectIdentifierSignedAndEnvelopedData,
	PkcsObjectIdentifierDigestedData,
	PkcsObjectIdentifierEncryptedData,
	PkcsObjectIdentifierPkcs9AtEmailAddress,
	PkcsObjectIdentifierPkcs9AtUnstructuredName,
	PkcsObjectIdentifierPkcs9AtContentType,
	PkcsObjectIdentifierPkcs9AtMessageDigest,
	PkcsObjectIdentifierPkcs9AtSigningTime,
	PkcsObjectIdentifierPkcs9AtCounterSignature,
	PkcsObjectIdentifierPkcs9AtChallengePassword,
	PkcsObjectIdentifierPkcs9AtUnstructuredAddress,
	PkcsObjectIdentifierPkcs9AtExtendedCertificateAttributes,
	PkcsObjectIdentifierPkcs9AtSigningDescription,
	PkcsObjectIdentifierPkcs9AtExtensionRequest,
	PkcsObjectIdentifierPkcs9AtSmimeCapabilities,
	PkcsObjectIdentifierIdSmime,
	PkcsObjectIdentifierPkcs9AtFriendlyName,
	PkcsObjectIdentifierPkcs9AtLocalKeyID,
	PkcsObjectIdentifierX509Certificate,
	PkcsObjectIdentifierSdsiCertificate,
	PkcsObjectIdentifierX509Crl,
	PkcsObjectIdentifierIdAlg,
	PkcsObjectIdentifierIdAlgEsdh,
	PkcsObjectIdentifierIdAlgCms3DesWrap,
	PkcsObjectIdentifierIdAlgCmsRC2Wrap,
	PkcsObjectIdentifierIdAlgPwriKek,
	PkcsObjectIdentifierIdAlgSsdh,
	PkcsObjectIdentifierIdRsaKem,
	PkcsObjectIdentifierPreferSignedData,
	PkcsObjectIdentifierCannotDecryptAny,
	PkcsObjectIdentifierSmimeCapabilitiesVersions,
	PkcsObjectIdentifierIdAAReceiptRequest,
	PkcsObjectIdentifierIdCTAuthData,
	PkcsObjectIdentifierIdCTTstInfo,
	PkcsObjectIdentifierIdCTCompressedData,
	PkcsObjectIdentifierIdCTAuthEnvelopedData,
	PkcsObjectIdentifierIdCTTimestampedData,
	PkcsObjectIdentifierIdCtiEtsProofOfOrigin,
	PkcsObjectIdentifierIdCtiEtsProofOfReceipt,
	PkcsObjectIdentifierIdCtiEtsProofOfDelivery,
	PkcsObjectIdentifierIdCtiEtsProofOfSender,
	PkcsObjectIdentifierIdCtiEtsProofOfApproval,
	PkcsObjectIdentifierIdCtiEtsProofOfCreation,
	PkcsObjectIdentifierIdAAContentHint,
	PkcsObjectIdentifierIdAAMsgSigDigest,
	PkcsObjectIdentifierIdAAContentReference,
	PkcsObjectIdentifierIdAAEncrypKeyPref,
	PkcsObjectIdentifierIdAASigningCertificate,
	PkcsObjectIdentifierIdAASigningCertificateV2,
	PkcsObjectIdentifierIdAAContentIdentifier,
	PkcsObjectIdentifierIdAASignatureTimeStampToken,
	PkcsObjectIdentifierIdAAEtsSigPolicyID,
	PkcsObjectIdentifierIdAAEtsCommitmentType,
	PkcsObjectIdentifierIdAAEtsSignerLocation,
	PkcsObjectIdentifierIdAAEtsSignerAttr,
	PkcsObjectIdentifierIdAAEtsOtherSigCert,
	PkcsObjectIdentifierIdAAEtsContentTimestamp,
	PkcsObjectIdentifierIdAAEtsCertificateRefs,
	PkcsObjectIdentifierIdAAEtsRevocationRefs,
	PkcsObjectIdentifierIdAAEtsCertValues,
	PkcsObjectIdentifierIdAAEtsRevocationValues,
	PkcsObjectIdentifierIdAAEtsEscTimeStamp,
	PkcsObjectIdentifierIdAAEtsCertCrlTimestamp,
	PkcsObjectIdentifierIdAAEtsArchiveTimestamp,
	PkcsObjectIdentifierIdSpqEtsUri,
	PkcsObjectIdentifierIdSpqEtsUNotice,
	PkcsObjectIdentifierKeyBag,
	PkcsObjectIdentifierPkcs8ShroudedKeyBag,
	PkcsObjectIdentifierCertBag,
	PkcsObjectIdentifierCrlBag,
	PkcsObjectIdentifierSecretBag,
	PkcsObjectIdentifierSafeContentsBag,
	PkcsObjectIdentifierPbeWithShaAnd128BitRC4,
	PkcsObjectIdentifierPbeWithShaAnd40BitRC4,
	PkcsObjectIdentifierPbeWithShaAnd3KeyTripleDesCbc,
	PkcsObjectIdentifierPbeWithShaAnd2KeyTripleDesCbc,
	PkcsObjectIdentifierPbeWithShaAnd128BitRC2Cbc,
	PkcsObjectIdentifierPbewithShaAnd40BitRC2Cbc,
	}
}
