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

// X509CertificateType x509 Certificate.
type X509CertificateType int

const (
	X509CertificateTypeNone X509CertificateType = iota
	X509CertificateTypeOldAuthorityKeyIdentifier
	X509CertificateTypeOldPrimaryKeyAttributes
	X509CertificateTypeCertificatePolicies
	X509CertificateTypeOrimaryKeyUsageRestriction
	X509CertificateTypeSubjectDirectoryAttributes
	X509CertificateTypeSubjectKeyIdentifier
	X509CertificateTypeKeyUsage
	X509CertificateTypePrivateKeyUsagePeriod
	X509CertificateTypeSubjectAlternativeName
	X509CertificateTypeIssuerAlternativeName
	X509CertificateTypeBasicConstraints
	X509CertificateTypeCrlNumber
	X509CertificateTypeReasonCode
	X509CertificateTypeHoldInstructionCode
	X509CertificateTypeInvalidityDate
	X509CertificateTypeDeltaCrlIndicator
	X509CertificateTypeIssuingDistributionPoint
	X509CertificateTypeCertificateIssuer
	X509CertificateTypeNameConstraints
	X509CertificateTypeCrlDistributionPoints
	X509CertificateTypeCertificatePolicies2
	X509CertificateTypePolicyMappings
	X509CertificateTypeAuthorityKeyIdentifier
	X509CertificateTypePolicyConstraints
	X509CertificateTypeExtendedKeyUsage
	X509CertificateTypeFreshestCrl
)

// X509CertificateTypeParse converts the given string into a X509CertificateType value.
//
// It returns the corresponding X509CertificateType constant if the string matches
// a known level name, or an error if the input is invalid.
func X509CertificateTypeParse(value string) (X509CertificateType, error) {
	var ret X509CertificateType
	var err error
	switch strings.ToUpper(value) {
	case "NONE":
		ret = X509CertificateTypeNone
	case "OLDAUTHORITYKEYIDENTIFIER":
		ret = X509CertificateTypeOldAuthorityKeyIdentifier
	case "OLDPRIMARYKEYATTRIBUTES":
		ret = X509CertificateTypeOldPrimaryKeyAttributes
	case "CERTIFICATEPOLICIES":
		ret = X509CertificateTypeCertificatePolicies
	case "ORIMARYKEYUSAGERESTRICTION":
		ret = X509CertificateTypeOrimaryKeyUsageRestriction
	case "SUBJECTDIRECTORYATTRIBUTES":
		ret = X509CertificateTypeSubjectDirectoryAttributes
	case "SUBJECTKEYIDENTIFIER":
		ret = X509CertificateTypeSubjectKeyIdentifier
	case "KEYUSAGE":
		ret = X509CertificateTypeKeyUsage
	case "PRIVATEKEYUSAGEPERIOD":
		ret = X509CertificateTypePrivateKeyUsagePeriod
	case "SUBJECTALTERNATIVENAME":
		ret = X509CertificateTypeSubjectAlternativeName
	case "ISSUERALTERNATIVENAME":
		ret = X509CertificateTypeIssuerAlternativeName
	case "BASICCONSTRAINTS":
		ret = X509CertificateTypeBasicConstraints
	case "CRLNUMBER":
		ret = X509CertificateTypeCrlNumber
	case "REASONCODE":
		ret = X509CertificateTypeReasonCode
	case "HOLDINSTRUCTIONCODE":
		ret = X509CertificateTypeHoldInstructionCode
	case "INVALIDITYDATE":
		ret = X509CertificateTypeInvalidityDate
	case "DELTACRLINDICATOR":
		ret = X509CertificateTypeDeltaCrlIndicator
	case "ISSUINGDISTRIBUTIONPOINT":
		ret = X509CertificateTypeIssuingDistributionPoint
	case "CERTIFICATEISSUER":
		ret = X509CertificateTypeCertificateIssuer
	case "NAMECONSTRAINTS":
		ret = X509CertificateTypeNameConstraints
	case "CRLDISTRIBUTIONPOINTS":
		ret = X509CertificateTypeCrlDistributionPoints
	case "CERTIFICATEPOLICIES2":
		ret = X509CertificateTypeCertificatePolicies2
	case "POLICYMAPPINGS":
		ret = X509CertificateTypePolicyMappings
	case "AUTHORITYKEYIDENTIFIER":
		ret = X509CertificateTypeAuthorityKeyIdentifier
	case "POLICYCONSTRAINTS":
		ret = X509CertificateTypePolicyConstraints
	case "EXTENDEDKEYUSAGE":
		ret = X509CertificateTypeExtendedKeyUsage
	case "FRESHESTCRL":
		ret = X509CertificateTypeFreshestCrl
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the X509CertificateType.
// It satisfies fmt.Stringer.
func (g X509CertificateType) String() string {
	var ret string
	switch g {
	case X509CertificateTypeNone:
		ret = "NONE"
	case X509CertificateTypeOldAuthorityKeyIdentifier:
		ret = "OLDAUTHORITYKEYIDENTIFIER"
	case X509CertificateTypeOldPrimaryKeyAttributes:
		ret = "OLDPRIMARYKEYATTRIBUTES"
	case X509CertificateTypeCertificatePolicies:
		ret = "CERTIFICATEPOLICIES"
	case X509CertificateTypeOrimaryKeyUsageRestriction:
		ret = "ORIMARYKEYUSAGERESTRICTION"
	case X509CertificateTypeSubjectDirectoryAttributes:
		ret = "SUBJECTDIRECTORYATTRIBUTES"
	case X509CertificateTypeSubjectKeyIdentifier:
		ret = "SUBJECTKEYIDENTIFIER"
	case X509CertificateTypeKeyUsage:
		ret = "KEYUSAGE"
	case X509CertificateTypePrivateKeyUsagePeriod:
		ret = "PRIVATEKEYUSAGEPERIOD"
	case X509CertificateTypeSubjectAlternativeName:
		ret = "SUBJECTALTERNATIVENAME"
	case X509CertificateTypeIssuerAlternativeName:
		ret = "ISSUERALTERNATIVENAME"
	case X509CertificateTypeBasicConstraints:
		ret = "BASICCONSTRAINTS"
	case X509CertificateTypeCrlNumber:
		ret = "CRLNUMBER"
	case X509CertificateTypeReasonCode:
		ret = "REASONCODE"
	case X509CertificateTypeHoldInstructionCode:
		ret = "HOLDINSTRUCTIONCODE"
	case X509CertificateTypeInvalidityDate:
		ret = "INVALIDITYDATE"
	case X509CertificateTypeDeltaCrlIndicator:
		ret = "DELTACRLINDICATOR"
	case X509CertificateTypeIssuingDistributionPoint:
		ret = "ISSUINGDISTRIBUTIONPOINT"
	case X509CertificateTypeCertificateIssuer:
		ret = "CERTIFICATEISSUER"
	case X509CertificateTypeNameConstraints:
		ret = "NAMECONSTRAINTS"
	case X509CertificateTypeCrlDistributionPoints:
		ret = "CRLDISTRIBUTIONPOINTS"
	case X509CertificateTypeCertificatePolicies2:
		ret = "CERTIFICATEPOLICIES2"
	case X509CertificateTypePolicyMappings:
		ret = "POLICYMAPPINGS"
	case X509CertificateTypeAuthorityKeyIdentifier:
		ret = "AUTHORITYKEYIDENTIFIER"
	case X509CertificateTypePolicyConstraints:
		ret = "POLICYCONSTRAINTS"
	case X509CertificateTypeExtendedKeyUsage:
		ret = "EXTENDEDKEYUSAGE"
	case X509CertificateTypeFreshestCrl:
		ret = "FRESHESTCRL"
	}
	return ret
}
