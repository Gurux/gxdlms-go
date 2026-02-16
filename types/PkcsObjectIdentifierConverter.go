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

	"github.com/Gurux/gxdlms-go/enums"
)

// PkcsObjectIdentifierToString converts the given PkcsObjectIdentifier value into a string.
func PkcsObjectIdentifierToString(value enums.PkcsObjectIdentifier) (string, error) {
	var ret string
	switch value {
	case enums.PkcsObjectIdentifierRsaEncryption:
		ret = "1.2.840.113549.1.1.1"
	case enums.PkcsObjectIdentifierMD2WithRsaEncryption:
		ret = "1.2.840.113549.1.1.2"
	case enums.PkcsObjectIdentifierMD4WithRsaEncryption:
		ret = "1.2.840.113549.1.1.3"
	case enums.PkcsObjectIdentifierMD5WithRsaEncryption:
		ret = "1.2.840.113549.1.1.4"
	case enums.PkcsObjectIdentifierSha1WithRsaEncryption:
		ret = "1.2.840.113549.1.1.5"
	case enums.PkcsObjectIdentifierSrsaOaepEncryptionSet:
		ret = "1.2.840.113549.1.1.6"
	case enums.PkcsObjectIdentifierIdRsaesOaep:
		ret = "1.2.840.113549.1.1.7"
	case enums.PkcsObjectIdentifierIdMgf1:
		ret = "1.2.840.113549.1.1.8"
	case enums.PkcsObjectIdentifierIdPSpecified:
		ret = "1.2.840.113549.1.1.9"
	case enums.PkcsObjectIdentifierIdRsassaPss:
		ret = "1.2.840.113549.1.1.10"
	case enums.PkcsObjectIdentifierSha256WithRsaEncryption:
		ret = "1.2.840.113549.1.1.11"
	case enums.PkcsObjectIdentifierSha384WithRsaEncryption:
		ret = "1.2.840.113549.1.1.12"
	case enums.PkcsObjectIdentifierSha512WithRsaEncryption:
		ret = "1.2.840.113549.1.1.13"
	case enums.PkcsObjectIdentifierSha224WithRsaEncryption:
		ret = "1.2.840.113549.1.1.14"
	case enums.PkcsObjectIdentifierDhKeyAgree1ment:
		ret = "1.2.840.113549.1.3.1"
	case enums.PkcsObjectIdentifierPbeWithMD2AndDesCbc:
		ret = "1.2.840.113549.1.5.1"
	case enums.PkcsObjectIdentifierPbeWithMD2AndRC2Cbc:
		ret = "1.2.840.113549.1.5.4"
	case enums.PkcsObjectIdentifierPbeWithMD5AndDesCbc:
		ret = "1.2.840.113549.1.5.3"
	case enums.PkcsObjectIdentifierPbeWithMD5AndRC2Cbc:
		ret = "1.2.840.113549.1.5.6"
	case enums.PkcsObjectIdentifierPbeWithSha1AndDesCbc:
		ret = "1.2.840.113549.1.5.10"
	case enums.PkcsObjectIdentifierPbeWithSha1AndRC2Cbc:
		ret = "1.2.840.113549.1.5.11"
	case enums.PkcsObjectIdentifierIdPbeS2:
		ret = "1.2.840.113549.1.5.13"
	case enums.PkcsObjectIdentifierIdPbkdf2:
		ret = "1.2.840.113549.1.5.12"
	case enums.PkcsObjectIdentifierDesEde3Cbc:
		ret = "1.2.840.113549.3.7"
	case enums.PkcsObjectIdentifierRC2Cbc:
		ret = "1.2.840.113549.3.2"
	case enums.PkcsObjectIdentifierMD2:
		ret = "1.2.840.113549.2.2"
	case enums.PkcsObjectIdentifierMD4:
		ret = "1.2.840.113549.2.4"
	case enums.PkcsObjectIdentifierMD5:
		ret = "1.2.840.113549.2.5"
	case enums.PkcsObjectIdentifierIdHmacWithSha1:
		ret = "1.2.840.113549.2.7"
	case enums.PkcsObjectIdentifierIdHmacWithSha224:
		ret = "1.2.840.113549.2.8"
	case enums.PkcsObjectIdentifierIdHmacWithSha256:
		ret = "1.2.840.113549.2.9"
	case enums.PkcsObjectIdentifierIdHmacWithSha384:
		ret = "1.2.840.113549.2.10"
	case enums.PkcsObjectIdentifierIdHmacWithSha512:
		ret = "1.2.840.113549.2.11"
	case enums.PkcsObjectIdentifierData:
		ret = "1.2.840.113549.1.7.1"
	case enums.PkcsObjectIdentifierSignedData:
		ret = "1.2.840.113549.1.7.2"
	case enums.PkcsObjectIdentifierEnvelopedData:
		ret = "1.2.840.113549.1.7.3"
	case enums.PkcsObjectIdentifierSignedAndEnvelopedData:
		ret = "1.2.840.113549.1.7.4"
	case enums.PkcsObjectIdentifierDigestedData:
		ret = "1.2.840.113549.1.7.5"
	case enums.PkcsObjectIdentifierEncryptedData:
		ret = "1.2.840.113549.1.7.6"
	case enums.PkcsObjectIdentifierPkcs9AtEmailAddress:
		ret = "1.2.840.113549.1.9.1"
	case enums.PkcsObjectIdentifierPkcs9AtUnstructuredName:
		ret = "1.2.840.113549.1.9.2"
	case enums.PkcsObjectIdentifierPkcs9AtContentType:
		ret = "1.2.840.113549.1.9.3"
	case enums.PkcsObjectIdentifierPkcs9AtMessageDigest:
		ret = "1.2.840.113549.1.9.4"
	case enums.PkcsObjectIdentifierPkcs9AtSigningTime:
		ret = "1.2.840.113549.1.9.5"
	case enums.PkcsObjectIdentifierPkcs9AtCounterSignature:
		ret = "1.2.840.113549.1.9.6"
	case enums.PkcsObjectIdentifierPkcs9AtChallengePassword:
		ret = "1.2.840.113549.1.9.7"
	case enums.PkcsObjectIdentifierPkcs9AtUnstructuredAddress:
		ret = "1.2.840.113549.1.9.8"
	case enums.PkcsObjectIdentifierPkcs9AtExtendedCertificateAttributes:
		ret = "1.2.840.113549.1.9.9"
	case enums.PkcsObjectIdentifierPkcs9AtSigningDescription:
		ret = "1.2.840.113549.1.9.13"
	case enums.PkcsObjectIdentifierPkcs9AtExtensionRequest:
		ret = "1.2.840.113549.1.9.14"
	case enums.PkcsObjectIdentifierPkcs9AtSmimeCapabilities:
		ret = "1.2.840.113549.1.9.15"
	case enums.PkcsObjectIdentifierIdSmime:
		ret = "1.2.840.113549.1.9.16"
	case enums.PkcsObjectIdentifierPkcs9AtFriendlyName:
		ret = "1.2.840.113549.1.9.20"
	case enums.PkcsObjectIdentifierPkcs9AtLocalKeyID:
		ret = "1.2.840.113549.1.9.21"
	case enums.PkcsObjectIdentifierX509Certificate:
		ret = "1.2.840.113549.1.9.22.1"
	case enums.PkcsObjectIdentifierSdsiCertificate:
		ret = "1.2.840.113549.1.9.22.2"
	case enums.PkcsObjectIdentifierX509Crl:
		ret = "1.2.840.113549.1.9.23.1"
	case enums.PkcsObjectIdentifierIdAlg:
		ret = "1.2.840.113549.1.9.16.3"
	case enums.PkcsObjectIdentifierIdAlgEsdh:
		ret = "1.2.840.113549.1.9.16.3.5"
	case enums.PkcsObjectIdentifierIdAlgCms3DesWrap:
		ret = "1.2.840.113549.1.9.16.3.6"
	case enums.PkcsObjectIdentifierIdAlgCmsRC2Wrap:
		ret = "1.2.840.113549.1.9.16.3.7"
	case enums.PkcsObjectIdentifierIdAlgPwriKek:
		ret = "1.2.840.113549.1.9.16.3.9"
	case enums.PkcsObjectIdentifierIdAlgSsdh:
		ret = "1.2.840.113549.1.9.16.3.10"
	case enums.PkcsObjectIdentifierIdRsaKem:
		ret = "1.2.840.113549.1.9.16.3.14"
	case enums.PkcsObjectIdentifierPreferSignedData:
		ret = "1.2.840.113549.1.9.15.1"
	case enums.PkcsObjectIdentifierCannotDecryptAny:
		ret = "1.2.840.113549.1.9.15.2"
	case enums.PkcsObjectIdentifierSmimeCapabilitiesVersions:
		ret = "1.2.840.113549.1.9.15.3"
	case enums.PkcsObjectIdentifierIdAAReceiptRequest:
		ret = "1.2.840.113549.1.9.16.2.1"
	case enums.PkcsObjectIdentifierIdCTAuthData:
		ret = "1.2.840.113549.1.9.16.1.2"
	case enums.PkcsObjectIdentifierIdCTTstInfo:
		ret = "1.2.840.113549.1.9.16.1.4"
	case enums.PkcsObjectIdentifierIdCTCompressedData:
		ret = "1.2.840.113549.1.9.16.1.9"
	case enums.PkcsObjectIdentifierIdCTAuthEnvelopedData:
		ret = "1.2.840.113549.1.9.16.1.23"
	case enums.PkcsObjectIdentifierIdCTTimestampedData:
		ret = "1.2.840.113549.1.9.16.1.31"
	case enums.PkcsObjectIdentifierIdCtiEtsProofOfOrigin:
		ret = "1.2.840.113549.1.9.16.6.1"
	case enums.PkcsObjectIdentifierIdCtiEtsProofOfReceipt:
		ret = "1.2.840.113549.1.9.16.6.2"
	case enums.PkcsObjectIdentifierIdCtiEtsProofOfDelivery:
		ret = "1.2.840.113549.1.9.16.6.3"
	case enums.PkcsObjectIdentifierIdCtiEtsProofOfSender:
		ret = "1.2.840.113549.1.9.16.6.4"
	case enums.PkcsObjectIdentifierIdCtiEtsProofOfApproval:
		ret = "1.2.840.113549.1.9.16.6.5"
	case enums.PkcsObjectIdentifierIdCtiEtsProofOfCreation:
		ret = "1.2.840.113549.1.9.16.6.6"
	case enums.PkcsObjectIdentifierIdAAContentHint:
		ret = "1.2.840.113549.1.9.16.2.4"
	case enums.PkcsObjectIdentifierIdAAMsgSigDigest:
		ret = "1.2.840.113549.1.9.16.2.5"
	case enums.PkcsObjectIdentifierIdAAContentReference:
		ret = "1.2.840.113549.1.9.16.2.10"
	case enums.PkcsObjectIdentifierIdAAEncrypKeyPref:
		ret = "1.2.840.113549.1.9.16.2.11"
	case enums.PkcsObjectIdentifierIdAASigningCertificate:
		ret = "1.2.840.113549.1.9.16.2.12"
	case enums.PkcsObjectIdentifierIdAASigningCertificateV2:
		ret = "1.2.840.113549.1.9.16.2.47"
	case enums.PkcsObjectIdentifierIdAAContentIdentifier:
		ret = "1.2.840.113549.1.9.16.2.7"
	case enums.PkcsObjectIdentifierIdAASignatureTimeStampToken:
		ret = "1.2.840.113549.1.9.16.2.14"
	case enums.PkcsObjectIdentifierIdAAEtsSigPolicyID:
		ret = "1.2.840.113549.1.9.16.2.15"
	case enums.PkcsObjectIdentifierIdAAEtsCommitmentType:
		ret = "1.2.840.113549.1.9.16.2.16"
	case enums.PkcsObjectIdentifierIdAAEtsSignerLocation:
		ret = "1.2.840.113549.1.9.16.2.17"
	case enums.PkcsObjectIdentifierIdAAEtsSignerAttr:
		ret = "1.2.840.113549.1.9.16.2.18"
	case enums.PkcsObjectIdentifierIdAAEtsOtherSigCert:
		ret = "1.2.840.113549.1.9.16.2.19"
	case enums.PkcsObjectIdentifierIdAAEtsContentTimestamp:
		ret = "1.2.840.113549.1.9.16.2.20"
	case enums.PkcsObjectIdentifierIdAAEtsCertificateRefs:
		ret = "1.2.840.113549.1.9.16.2.21"
	case enums.PkcsObjectIdentifierIdAAEtsRevocationRefs:
		ret = "1.2.840.113549.1.9.16.2.22"
	case enums.PkcsObjectIdentifierIdAAEtsCertValues:
		ret = "1.2.840.113549.1.9.16.2.23"
	case enums.PkcsObjectIdentifierIdAAEtsRevocationValues:
		ret = "1.2.840.113549.1.9.16.2.24"
	case enums.PkcsObjectIdentifierIdAAEtsEscTimeStamp:
		ret = "1.2.840.113549.1.9.16.2.25"
	case enums.PkcsObjectIdentifierIdAAEtsCertCrlTimestamp:
		ret = "1.2.840.113549.1.9.16.2.26"
	case enums.PkcsObjectIdentifierIdAAEtsArchiveTimestamp:
		ret = "1.2.840.113549.1.9.16.2.27"
	case enums.PkcsObjectIdentifierIdSpqEtsUri:
		ret = "1.2.840.113549.1.9.16.5.1"
	case enums.PkcsObjectIdentifierIdSpqEtsUNotice:
		ret = "1.2.840.113549.1.9.16.5.2"
	case enums.PkcsObjectIdentifierKeyBag:
		ret = "1.2.840.113549.1.12.10.1.1"
	case enums.PkcsObjectIdentifierPkcs8ShroudedKeyBag:
		ret = "1.2.840.113549.1.12.10.1.2"
	case enums.PkcsObjectIdentifierCertBag:
		ret = "1.2.840.113549.1.12.10.1.3"
	case enums.PkcsObjectIdentifierCrlBag:
		ret = "1.2.840.113549.1.12.10.1.4"
	case enums.PkcsObjectIdentifierSecretBag:
		ret = "1.2.840.113549.1.12.10.1.5"
	case enums.PkcsObjectIdentifierSafeContentsBag:
		ret = "1.2.840.113549.1.12.10.1.6"
	case enums.PkcsObjectIdentifierPbeWithShaAnd128BitRC4:
		ret = "1.2.840.113549.1.12.1.1"
	case enums.PkcsObjectIdentifierPbeWithShaAnd40BitRC4:
		ret = "1.2.840.113549.1.12.1.2"
	case enums.PkcsObjectIdentifierPbeWithShaAnd3KeyTripleDesCbc:
		ret = "1.2.840.113549.1.12.1.3"
	case enums.PkcsObjectIdentifierPbeWithShaAnd2KeyTripleDesCbc:
		ret = "1.2.840.113549.1.12.1.4"
	case enums.PkcsObjectIdentifierPbeWithShaAnd128BitRC2Cbc:
		ret = "1.2.840.113549.1.12.1.5"
	case enums.PkcsObjectIdentifierPbewithShaAnd40BitRC2Cbc:
		ret = "1.2.840.113549.1.12.1.6"
	default:
		return "", fmt.Errorf("Invalid PKCS Object Identifier. %s ", value)
	}
	return ret, nil
}

// PkcsObjectIdentifierFromString converts the given string value into a PkcsObjectIdentifier enum.
func PkcsObjectIdentifierFromString(value string) enums.PkcsObjectIdentifier {
	var ret enums.PkcsObjectIdentifier
	switch value {
	case "1.2.840.113549.1.1.1":
		ret = enums.PkcsObjectIdentifierRsaEncryption
	case "1.2.840.113549.1.1.2":
		ret = enums.PkcsObjectIdentifierMD2WithRsaEncryption
	case "1.2.840.113549.1.1.3":
		ret = enums.PkcsObjectIdentifierMD4WithRsaEncryption
	case "1.2.840.113549.1.1.4":
		ret = enums.PkcsObjectIdentifierMD5WithRsaEncryption
	case "1.2.840.113549.1.1.5":
		ret = enums.PkcsObjectIdentifierSha1WithRsaEncryption
	case "1.2.840.113549.1.1.6":
		ret = enums.PkcsObjectIdentifierSrsaOaepEncryptionSet
	case "1.2.840.113549.1.1.7":
		ret = enums.PkcsObjectIdentifierIdRsaesOaep
	case "1.2.840.113549.1.1.8":
		ret = enums.PkcsObjectIdentifierIdMgf1
	case "1.2.840.113549.1.1.9":
		ret = enums.PkcsObjectIdentifierIdPSpecified
	case "1.2.840.113549.1.1.10":
		ret = enums.PkcsObjectIdentifierIdRsassaPss
	case "1.2.840.113549.1.1.11":
		ret = enums.PkcsObjectIdentifierSha256WithRsaEncryption
	case "1.2.840.113549.1.1.12":
		ret = enums.PkcsObjectIdentifierSha384WithRsaEncryption
	case "1.2.840.113549.1.1.13":
		ret = enums.PkcsObjectIdentifierSha512WithRsaEncryption
	case "1.2.840.113549.1.1.14":
		ret = enums.PkcsObjectIdentifierSha224WithRsaEncryption
	case "1.2.840.113549.1.3.1":
		ret = enums.PkcsObjectIdentifierDhKeyAgree1ment
	case "1.2.840.113549.1.5.1":
		ret = enums.PkcsObjectIdentifierPbeWithMD2AndDesCbc
	case "1.2.840.113549.1.5.4":
		ret = enums.PkcsObjectIdentifierPbeWithMD2AndRC2Cbc
	case "1.2.840.113549.1.5.3":
		ret = enums.PkcsObjectIdentifierPbeWithMD5AndDesCbc
	case "1.2.840.113549.1.5.6":
		ret = enums.PkcsObjectIdentifierPbeWithMD5AndRC2Cbc
	case "1.2.840.113549.1.5.10":
		ret = enums.PkcsObjectIdentifierPbeWithSha1AndDesCbc
	case "1.2.840.113549.1.5.11":
		ret = enums.PkcsObjectIdentifierPbeWithSha1AndRC2Cbc
	case "1.2.840.113549.1.5.13":
		ret = enums.PkcsObjectIdentifierIdPbeS2
	case "1.2.840.113549.1.5.12":
		ret = enums.PkcsObjectIdentifierIdPbkdf2
	case "1.2.840.113549.3.7":
		ret = enums.PkcsObjectIdentifierDesEde3Cbc
	case "1.2.840.113549.3.2":
		ret = enums.PkcsObjectIdentifierRC2Cbc
	case "1.2.840.113549.2.2":
		ret = enums.PkcsObjectIdentifierMD2
	case "1.2.840.113549.2.4":
		ret = enums.PkcsObjectIdentifierMD4
	case "1.2.840.113549.2.5":
		ret = enums.PkcsObjectIdentifierMD5
	case "1.2.840.113549.2.7":
		ret = enums.PkcsObjectIdentifierIdHmacWithSha1
	case "1.2.840.113549.2.8":
		ret = enums.PkcsObjectIdentifierIdHmacWithSha224
	case "1.2.840.113549.2.9":
		ret = enums.PkcsObjectIdentifierIdHmacWithSha256
	case "1.2.840.113549.2.10":
		ret = enums.PkcsObjectIdentifierIdHmacWithSha384
	case "1.2.840.113549.2.11":
		ret = enums.PkcsObjectIdentifierIdHmacWithSha512
	case "1.2.840.113549.1.7.1":
		ret = enums.PkcsObjectIdentifierData
	case "1.2.840.113549.1.7.2":
		ret = enums.PkcsObjectIdentifierSignedData
	case "1.2.840.113549.1.7.3":
		ret = enums.PkcsObjectIdentifierEnvelopedData
	case "1.2.840.113549.1.7.4":
		ret = enums.PkcsObjectIdentifierSignedAndEnvelopedData
	case "1.2.840.113549.1.7.5":
		ret = enums.PkcsObjectIdentifierDigestedData
	case "1.2.840.113549.1.7.6":
		ret = enums.PkcsObjectIdentifierEncryptedData
	case "1.2.840.113549.1.9.1":
		ret = enums.PkcsObjectIdentifierPkcs9AtEmailAddress
	case "1.2.840.113549.1.9.2":
		ret = enums.PkcsObjectIdentifierPkcs9AtUnstructuredName
	case "1.2.840.113549.1.9.3":
		ret = enums.PkcsObjectIdentifierPkcs9AtContentType
	case "1.2.840.113549.1.9.4":
		ret = enums.PkcsObjectIdentifierPkcs9AtMessageDigest
	case "1.2.840.113549.1.9.5":
		ret = enums.PkcsObjectIdentifierPkcs9AtSigningTime
	case "1.2.840.113549.1.9.6":
		ret = enums.PkcsObjectIdentifierPkcs9AtCounterSignature
	case "1.2.840.113549.1.9.7":
		ret = enums.PkcsObjectIdentifierPkcs9AtChallengePassword
	case "1.2.840.113549.1.9.8":
		ret = enums.PkcsObjectIdentifierPkcs9AtUnstructuredAddress
	case "1.2.840.113549.1.9.9":
		ret = enums.PkcsObjectIdentifierPkcs9AtExtendedCertificateAttributes
	case "1.2.840.113549.1.9.13":
		ret = enums.PkcsObjectIdentifierPkcs9AtSigningDescription
	case "1.2.840.113549.1.9.14":
		ret = enums.PkcsObjectIdentifierPkcs9AtExtensionRequest
	case "1.2.840.113549.1.9.15":
		ret = enums.PkcsObjectIdentifierPkcs9AtSmimeCapabilities
	case "1.2.840.113549.1.9.16":
		ret = enums.PkcsObjectIdentifierIdSmime
	case "1.2.840.113549.1.9.20":
		ret = enums.PkcsObjectIdentifierPkcs9AtFriendlyName
	case "1.2.840.113549.1.9.21":
		ret = enums.PkcsObjectIdentifierPkcs9AtLocalKeyID
	case "1.2.840.113549.1.9.22.1":
		ret = enums.PkcsObjectIdentifierX509Certificate
	case "1.2.840.113549.1.9.22.2":
		ret = enums.PkcsObjectIdentifierSdsiCertificate
	case "1.2.840.113549.1.9.23.1":
		ret = enums.PkcsObjectIdentifierX509Crl
	case "1.2.840.113549.1.9.16.3":
		ret = enums.PkcsObjectIdentifierIdAlg
	case "1.2.840.113549.1.9.16.3.5":
		ret = enums.PkcsObjectIdentifierIdAlgEsdh
	case "1.2.840.113549.1.9.16.3.6":
		ret = enums.PkcsObjectIdentifierIdAlgCms3DesWrap
	case "1.2.840.113549.1.9.16.3.7":
		ret = enums.PkcsObjectIdentifierIdAlgCmsRC2Wrap
	case "1.2.840.113549.1.9.16.3.9":
		ret = enums.PkcsObjectIdentifierIdAlgPwriKek
	case "1.2.840.113549.1.9.16.3.10":
		ret = enums.PkcsObjectIdentifierIdAlgSsdh
	case "1.2.840.113549.1.9.16.3.14":
		ret = enums.PkcsObjectIdentifierIdRsaKem
	case "1.2.840.113549.1.9.15.1":
		ret = enums.PkcsObjectIdentifierPreferSignedData
	case "1.2.840.113549.1.9.15.2":
		ret = enums.PkcsObjectIdentifierCannotDecryptAny
	case "1.2.840.113549.1.9.15.3":
		ret = enums.PkcsObjectIdentifierSmimeCapabilitiesVersions
	case "1.2.840.113549.1.9.16.2.1":
		ret = enums.PkcsObjectIdentifierIdAAReceiptRequest
	case "1.2.840.113549.1.9.16.1.2":
		ret = enums.PkcsObjectIdentifierIdCTAuthData
	case "1.2.840.113549.1.9.16.1.4":
		ret = enums.PkcsObjectIdentifierIdCTTstInfo
	case "1.2.840.113549.1.9.16.1.9":
		ret = enums.PkcsObjectIdentifierIdCTCompressedData
	case "1.2.840.113549.1.9.16.1.23":
		ret = enums.PkcsObjectIdentifierIdCTAuthEnvelopedData
	case "1.2.840.113549.1.9.16.1.31":
		ret = enums.PkcsObjectIdentifierIdCTTimestampedData
	case "1.2.840.113549.1.9.16.6.1":
		ret = enums.PkcsObjectIdentifierIdCtiEtsProofOfOrigin
	case "1.2.840.113549.1.9.16.6.2":
		ret = enums.PkcsObjectIdentifierIdCtiEtsProofOfReceipt
	case "1.2.840.113549.1.9.16.6.3":
		ret = enums.PkcsObjectIdentifierIdCtiEtsProofOfDelivery
	case "1.2.840.113549.1.9.16.6.4":
		ret = enums.PkcsObjectIdentifierIdCtiEtsProofOfSender
	case "1.2.840.113549.1.9.16.6.5":
		ret = enums.PkcsObjectIdentifierIdCtiEtsProofOfApproval
	case "1.2.840.113549.1.9.16.6.6":
		ret = enums.PkcsObjectIdentifierIdCtiEtsProofOfCreation
	case "1.2.840.113549.1.9.16.2.4":
		ret = enums.PkcsObjectIdentifierIdAAContentHint
	case "1.2.840.113549.1.9.16.2.5":
		ret = enums.PkcsObjectIdentifierIdAAMsgSigDigest
	case "1.2.840.113549.1.9.16.2.10":
		ret = enums.PkcsObjectIdentifierIdAAContentReference
	case "1.2.840.113549.1.9.16.2.11":
		ret = enums.PkcsObjectIdentifierIdAAEncrypKeyPref
	case "1.2.840.113549.1.9.16.2.12":
		ret = enums.PkcsObjectIdentifierIdAASigningCertificate
	case "1.2.840.113549.1.9.16.2.47":
		ret = enums.PkcsObjectIdentifierIdAASigningCertificateV2
	case "1.2.840.113549.1.9.16.2.7":
		ret = enums.PkcsObjectIdentifierIdAAContentIdentifier
	case "1.2.840.113549.1.9.16.2.14":
		ret = enums.PkcsObjectIdentifierIdAASignatureTimeStampToken
	case "1.2.840.113549.1.9.16.2.15":
		ret = enums.PkcsObjectIdentifierIdAAEtsSigPolicyID
	case "1.2.840.113549.1.9.16.2.16":
		ret = enums.PkcsObjectIdentifierIdAAEtsCommitmentType
	case "1.2.840.113549.1.9.16.2.17":
		ret = enums.PkcsObjectIdentifierIdAAEtsSignerLocation
	case "1.2.840.113549.1.9.16.2.18":
		ret = enums.PkcsObjectIdentifierIdAAEtsSignerAttr
	case "1.2.840.113549.1.9.16.2.19":
		ret = enums.PkcsObjectIdentifierIdAAEtsOtherSigCert
	case "1.2.840.113549.1.9.16.2.20":
		ret = enums.PkcsObjectIdentifierIdAAEtsContentTimestamp
	case "1.2.840.113549.1.9.16.2.21":
		ret = enums.PkcsObjectIdentifierIdAAEtsCertificateRefs
	case "1.2.840.113549.1.9.16.2.22":
		ret = enums.PkcsObjectIdentifierIdAAEtsRevocationRefs
	case "1.2.840.113549.1.9.16.2.23":
		ret = enums.PkcsObjectIdentifierIdAAEtsCertValues
	case "1.2.840.113549.1.9.16.2.24":
		ret = enums.PkcsObjectIdentifierIdAAEtsRevocationValues
	case "1.2.840.113549.1.9.16.2.25":
		ret = enums.PkcsObjectIdentifierIdAAEtsEscTimeStamp
	case "1.2.840.113549.1.9.16.2.26":
		ret = enums.PkcsObjectIdentifierIdAAEtsCertCrlTimestamp
	case "1.2.840.113549.1.9.16.2.27":
		ret = enums.PkcsObjectIdentifierIdAAEtsArchiveTimestamp
	case "1.2.840.113549.1.9.16.5.1":
		ret = enums.PkcsObjectIdentifierIdSpqEtsUri
	case "1.2.840.113549.1.9.16.5.2":
		ret = enums.PkcsObjectIdentifierIdSpqEtsUNotice
	case "1.2.840.113549.1.12.10.1.1":
		ret = enums.PkcsObjectIdentifierKeyBag
	case "1.2.840.113549.1.12.10.1.2":
		ret = enums.PkcsObjectIdentifierPkcs8ShroudedKeyBag
	case "1.2.840.113549.1.12.10.1.3":
		ret = enums.PkcsObjectIdentifierCertBag
	case "1.2.840.113549.1.12.10.1.4":
		ret = enums.PkcsObjectIdentifierCrlBag
	case "1.2.840.113549.1.12.10.1.5":
		ret = enums.PkcsObjectIdentifierSecretBag
	case "1.2.840.113549.1.12.10.1.6":
		ret = enums.PkcsObjectIdentifierSafeContentsBag
	case "1.2.840.113549.1.12.1.1":
		ret = enums.PkcsObjectIdentifierPbeWithShaAnd128BitRC4
	case "1.2.840.113549.1.12.1.2":
		ret = enums.PkcsObjectIdentifierPbeWithShaAnd40BitRC4
	case "1.2.840.113549.1.12.1.3":
		ret = enums.PkcsObjectIdentifierPbeWithShaAnd3KeyTripleDesCbc
	case "1.2.840.113549.1.12.1.4":
		ret = enums.PkcsObjectIdentifierPbeWithShaAnd2KeyTripleDesCbc
	case "1.2.840.113549.1.12.1.5":
		ret = enums.PkcsObjectIdentifierPbeWithShaAnd128BitRC2Cbc
	case "1.2.840.113549.1.12.1.6":
		ret = enums.PkcsObjectIdentifierPbewithShaAnd40BitRC2Cbc
	default:
		ret = enums.PkcsObjectIdentifierNone
	}
	return ret
}
