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

// X509CertificateType enumerates all x509 certificate types.
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
	switch {
	case strings.EqualFold(value, "None"):
		ret = X509CertificateTypeNone
	case strings.EqualFold(value, "OldAuthorityKeyIdentifier"):
		ret = X509CertificateTypeOldAuthorityKeyIdentifier
	case strings.EqualFold(value, "OldPrimaryKeyAttributes"):
		ret = X509CertificateTypeOldPrimaryKeyAttributes
	case strings.EqualFold(value, "CertificatePolicies"):
		ret = X509CertificateTypeCertificatePolicies
	case strings.EqualFold(value, "OrimaryKeyUsageRestriction"):
		ret = X509CertificateTypeOrimaryKeyUsageRestriction
	case strings.EqualFold(value, "SubjectDirectoryAttributes"):
		ret = X509CertificateTypeSubjectDirectoryAttributes
	case strings.EqualFold(value, "SubjectKeyIdentifier"):
		ret = X509CertificateTypeSubjectKeyIdentifier
	case strings.EqualFold(value, "KeyUsage"):
		ret = X509CertificateTypeKeyUsage
	case strings.EqualFold(value, "PrivateKeyUsagePeriod"):
		ret = X509CertificateTypePrivateKeyUsagePeriod
	case strings.EqualFold(value, "SubjectAlternativeName"):
		ret = X509CertificateTypeSubjectAlternativeName
	case strings.EqualFold(value, "IssuerAlternativeName"):
		ret = X509CertificateTypeIssuerAlternativeName
	case strings.EqualFold(value, "BasicConstraints"):
		ret = X509CertificateTypeBasicConstraints
	case strings.EqualFold(value, "CrlNumber"):
		ret = X509CertificateTypeCrlNumber
	case strings.EqualFold(value, "ReasonCode"):
		ret = X509CertificateTypeReasonCode
	case strings.EqualFold(value, "HoldInstructionCode"):
		ret = X509CertificateTypeHoldInstructionCode
	case strings.EqualFold(value, "InvalidityDate"):
		ret = X509CertificateTypeInvalidityDate
	case strings.EqualFold(value, "DeltaCrlIndicator"):
		ret = X509CertificateTypeDeltaCrlIndicator
	case strings.EqualFold(value, "IssuingDistributionPoint"):
		ret = X509CertificateTypeIssuingDistributionPoint
	case strings.EqualFold(value, "CertificateIssuer"):
		ret = X509CertificateTypeCertificateIssuer
	case strings.EqualFold(value, "NameConstraints"):
		ret = X509CertificateTypeNameConstraints
	case strings.EqualFold(value, "CrlDistributionPoints"):
		ret = X509CertificateTypeCrlDistributionPoints
	case strings.EqualFold(value, "CertificatePolicies2"):
		ret = X509CertificateTypeCertificatePolicies2
	case strings.EqualFold(value, "PolicyMappings"):
		ret = X509CertificateTypePolicyMappings
	case strings.EqualFold(value, "AuthorityKeyIdentifier"):
		ret = X509CertificateTypeAuthorityKeyIdentifier
	case strings.EqualFold(value, "PolicyConstraints"):
		ret = X509CertificateTypePolicyConstraints
	case strings.EqualFold(value, "ExtendedKeyUsage"):
		ret = X509CertificateTypeExtendedKeyUsage
	case strings.EqualFold(value, "FreshestCrl"):
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
		ret = "None"
	case X509CertificateTypeOldAuthorityKeyIdentifier:
		ret = "OldAuthorityKeyIdentifier"
	case X509CertificateTypeOldPrimaryKeyAttributes:
		ret = "OldPrimaryKeyAttributes"
	case X509CertificateTypeCertificatePolicies:
		ret = "CertificatePolicies"
	case X509CertificateTypeOrimaryKeyUsageRestriction:
		ret = "OrimaryKeyUsageRestriction"
	case X509CertificateTypeSubjectDirectoryAttributes:
		ret = "SubjectDirectoryAttributes"
	case X509CertificateTypeSubjectKeyIdentifier:
		ret = "SubjectKeyIdentifier"
	case X509CertificateTypeKeyUsage:
		ret = "KeyUsage"
	case X509CertificateTypePrivateKeyUsagePeriod:
		ret = "PrivateKeyUsagePeriod"
	case X509CertificateTypeSubjectAlternativeName:
		ret = "SubjectAlternativeName"
	case X509CertificateTypeIssuerAlternativeName:
		ret = "IssuerAlternativeName"
	case X509CertificateTypeBasicConstraints:
		ret = "BasicConstraints"
	case X509CertificateTypeCrlNumber:
		ret = "CrlNumber"
	case X509CertificateTypeReasonCode:
		ret = "ReasonCode"
	case X509CertificateTypeHoldInstructionCode:
		ret = "HoldInstructionCode"
	case X509CertificateTypeInvalidityDate:
		ret = "InvalidityDate"
	case X509CertificateTypeDeltaCrlIndicator:
		ret = "DeltaCrlIndicator"
	case X509CertificateTypeIssuingDistributionPoint:
		ret = "IssuingDistributionPoint"
	case X509CertificateTypeCertificateIssuer:
		ret = "CertificateIssuer"
	case X509CertificateTypeNameConstraints:
		ret = "NameConstraints"
	case X509CertificateTypeCrlDistributionPoints:
		ret = "CrlDistributionPoints"
	case X509CertificateTypeCertificatePolicies2:
		ret = "CertificatePolicies2"
	case X509CertificateTypePolicyMappings:
		ret = "PolicyMappings"
	case X509CertificateTypeAuthorityKeyIdentifier:
		ret = "AuthorityKeyIdentifier"
	case X509CertificateTypePolicyConstraints:
		ret = "PolicyConstraints"
	case X509CertificateTypeExtendedKeyUsage:
		ret = "ExtendedKeyUsage"
	case X509CertificateTypeFreshestCrl:
		ret = "FreshestCrl"
	}
	return ret
}

// AllX509CertificateType returns a slice containing all defined X509CertificateType values.
func AllX509CertificateType() []X509CertificateType {
	return []X509CertificateType{
		X509CertificateTypeNone,
		X509CertificateTypeOldAuthorityKeyIdentifier,
		X509CertificateTypeOldPrimaryKeyAttributes,
		X509CertificateTypeCertificatePolicies,
		X509CertificateTypeOrimaryKeyUsageRestriction,
		X509CertificateTypeSubjectDirectoryAttributes,
		X509CertificateTypeSubjectKeyIdentifier,
		X509CertificateTypeKeyUsage,
		X509CertificateTypePrivateKeyUsagePeriod,
		X509CertificateTypeSubjectAlternativeName,
		X509CertificateTypeIssuerAlternativeName,
		X509CertificateTypeBasicConstraints,
		X509CertificateTypeCrlNumber,
		X509CertificateTypeReasonCode,
		X509CertificateTypeHoldInstructionCode,
		X509CertificateTypeInvalidityDate,
		X509CertificateTypeDeltaCrlIndicator,
		X509CertificateTypeIssuingDistributionPoint,
		X509CertificateTypeCertificateIssuer,
		X509CertificateTypeNameConstraints,
		X509CertificateTypeCrlDistributionPoints,
		X509CertificateTypeCertificatePolicies2,
		X509CertificateTypePolicyMappings,
		X509CertificateTypeAuthorityKeyIdentifier,
		X509CertificateTypePolicyConstraints,
		X509CertificateTypeExtendedKeyUsage,
		X509CertificateTypeFreshestCrl,
	}
}
