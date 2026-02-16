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

// X509CertificateTypeToString converts the given X509CertificateType value into a string.
func X509CertificateTypeToString(value enums.X509CertificateType) (string, error) {
	var ret string
	switch value {
	case enums.X509CertificateTypeOldAuthorityKeyIdentifier:
		ret = "2.5.29.1"
	case enums.X509CertificateTypeOldPrimaryKeyAttributes:
		ret = "2.5.29.2"
	case enums.X509CertificateTypeCertificatePolicies:
		ret = "2.5.29.3"
	case enums.X509CertificateTypeOrimaryKeyUsageRestriction:
		ret = "2.5.29.4"
	case enums.X509CertificateTypeSubjectDirectoryAttributes:
		ret = "2.5.29.9"
	case enums.X509CertificateTypeSubjectKeyIdentifier:
		ret = "2.5.29.14"
	case enums.X509CertificateTypeKeyUsage:
		ret = "2.5.29.15"
	case enums.X509CertificateTypePrivateKeyUsagePeriod:
		ret = "2.5.29.16"
	case enums.X509CertificateTypeSubjectAlternativeName:
		ret = "2.5.29.17"
	case enums.X509CertificateTypeIssuerAlternativeName:
		ret = "2.5.29.18"
	case enums.X509CertificateTypeBasicConstraints:
		ret = "2.5.29.19"
	case enums.X509CertificateTypeCrlNumber:
		ret = "2.5.29.20"
	case enums.X509CertificateTypeReasonCode:
		ret = "2.5.29.21"
	case enums.X509CertificateTypeHoldInstructionCode:
		ret = "2.5.29.23"
	case enums.X509CertificateTypeInvalidityDate:
		ret = "2.5.29.24"
	case enums.X509CertificateTypeDeltaCrlIndicator:
		ret = "2.5.29.27"
	case enums.X509CertificateTypeIssuingDistributionPoint:
		ret = "2.5.29.28"
	case enums.X509CertificateTypeCertificateIssuer:
		ret = "2.5.29.29"
	case enums.X509CertificateTypeNameConstraints:
		ret = "2.5.29.30"
	case enums.X509CertificateTypeCrlDistributionPoints:
		ret = "2.5.29.31"
	case enums.X509CertificateTypeCertificatePolicies2:
		ret = "2.5.29.32"
	case enums.X509CertificateTypePolicyMappings:
		ret = "2.5.29.33"
	case enums.X509CertificateTypeAuthorityKeyIdentifier:
		ret = "2.5.29.35"
	case enums.X509CertificateTypePolicyConstraints:
		ret = "2.5.29.36"
	case enums.X509CertificateTypeExtendedKeyUsage:
		ret = "2.5.29.37"
	case enums.X509CertificateTypeFreshestCrl:
		ret = "2.5.29.46"
	default:
		return "", fmt.Errorf("Invalid X509 Certificate Type. %s ", value)
	}
	return ret, nil
}

// X509CertificateTypeFromString converts the given string value into a X509CertificateType enum.
func X509CertificateTypeFromString(value string) enums.X509CertificateType {
	var ret enums.X509CertificateType
	switch value {
	case "2.5.29.1":
		ret = enums.X509CertificateTypeOldAuthorityKeyIdentifier
	case "2.5.29.2":
		ret = enums.X509CertificateTypeOldPrimaryKeyAttributes
	case "2.5.29.3":
		ret = enums.X509CertificateTypeCertificatePolicies
	case "2.5.29.4":
		ret = enums.X509CertificateTypeOrimaryKeyUsageRestriction
	case "2.5.29.9":
		ret = enums.X509CertificateTypeSubjectDirectoryAttributes
	case "2.5.29.14":
		ret = enums.X509CertificateTypeSubjectKeyIdentifier
	case "2.5.29.15":
		ret = enums.X509CertificateTypeKeyUsage
	case "2.5.29.16":
		ret = enums.X509CertificateTypePrivateKeyUsagePeriod
	case "2.5.29.17":
		ret = enums.X509CertificateTypeSubjectAlternativeName
	case "2.5.29.18":
		ret = enums.X509CertificateTypeIssuerAlternativeName
	case "2.5.29.19":
		ret = enums.X509CertificateTypeBasicConstraints
	case "2.5.29.20":
		ret = enums.X509CertificateTypeCrlNumber
	case "2.5.29.21":
		ret = enums.X509CertificateTypeReasonCode
	case "2.5.29.23":
		ret = enums.X509CertificateTypeHoldInstructionCode
	case "2.5.29.24":
		ret = enums.X509CertificateTypeInvalidityDate
	case "2.5.29.27":
		ret = enums.X509CertificateTypeDeltaCrlIndicator
	case "2.5.29.28":
		ret = enums.X509CertificateTypeIssuingDistributionPoint
	case "2.5.29.29":
		ret = enums.X509CertificateTypeCertificateIssuer
	case "2.5.29.30":
		ret = enums.X509CertificateTypeNameConstraints
	case "2.5.29.31":
		ret = enums.X509CertificateTypeCrlDistributionPoints
	case "2.5.29.32":
		ret = enums.X509CertificateTypeCertificatePolicies2
	case "2.5.29.33":
		ret = enums.X509CertificateTypePolicyMappings
	case "2.5.29.35":
		ret = enums.X509CertificateTypeAuthorityKeyIdentifier
	case "2.5.29.36":
		ret = enums.X509CertificateTypePolicyConstraints
	case "2.5.29.37":
		ret = enums.X509CertificateTypeExtendedKeyUsage
	case "2.5.29.46":
		ret = enums.X509CertificateTypeFreshestCrl
	default:
		ret = enums.X509CertificateTypeNone
	}
	return ret
}
