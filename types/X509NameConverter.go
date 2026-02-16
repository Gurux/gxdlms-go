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

// GetString converts the given X509Name value into its string representation.
func X509NameToString(value enums.X509Name) (string, error) {
	var ret string
	switch value {
	case enums.X509NameC:
		ret = "2.5.4.6"
	case enums.X509NameO:
		ret = "2.5.4.10"
	case enums.X509NameOU:
		ret = "2.5.4.11"
	case enums.X509NameT:
		ret = "2.5.4.12"
	case enums.X509NameCN:
		ret = "2.5.4.3"
	case enums.X509NameSTREET:
		ret = "2.5.4.9"
	case enums.X509NameSerialNumber:
		ret = "2.5.4.5"
	case enums.X509NameL:
		ret = "2.5.4.7"
	case enums.X509NameST:
		ret = "2.5.4.8"
	case enums.X509NameSurName:
		ret = "2.5.4.4"
	case enums.X509NameGivenName:
		ret = "2.5.4.42"
	case enums.X509NameInitials:
		ret = "2.5.4.43"
	case enums.X509NameGeneration:
		ret = "2.5.4.44"
	case enums.X509NameUniqueIdentifier:
		ret = "2.5.4.45"
	case enums.X509NameBusinessCategory:
		ret = "2.5.4.15"
	case enums.X509NamePostalCode:
		ret = "2.5.4.17"
	case enums.X509NameDnQualifier:
		ret = "2.5.4.46"
	case enums.X509NamePseudonym:
		ret = "2.5.4.65"
	case enums.X509NameDateOfBirth:
		ret = "1.3.6.1.5.5.7.9.1"
	case enums.X509NamePlaceOfBirth:
		ret = "1.3.6.1.5.5.7.9.2"
	case enums.X509NameGender:
		ret = "1.3.6.1.5.5.7.9.3"
	case enums.X509NameCountryOfCitizenship:
		ret = "1.3.6.1.5.5.7.9.4"
	case enums.X509NameCountryOfResidence:
		ret = "1.3.6.1.5.5.7.9.5"
	case enums.X509NameNameAtBirth:
		ret = "1.3.36.8.3.14"
	case enums.X509NamePostalAddress:
		ret = "2.5.4.16"
	case enums.X509NameDmdName:
		ret = "2.5.4.54"
	case enums.X509NameTelephoneNumber:
		ret = "2.5.4.20"
	case enums.X509NameName:
		ret = "2.5.4.41"
	case enums.X509NameE:
		ret = "1.2.840.113549.1.9.1"
	case enums.X509NameDC:
		ret = "0.9.2342.19200300.100.1.25"
	case enums.X509NameUID:
		ret = "0.9.2342.19200300.100.1.1"
	default:
		return "", fmt.Errorf("Invalid X509Name. %s", value)
	}
	return ret, nil
}

// FromString converts the given string into an X509Name value.
func X509NameFromString(value string) enums.X509Name {
	var ret enums.X509Name
	switch value {
	case "2.5.4.6":
		ret = enums.X509NameC
	case "2.5.4.10":
		ret = enums.X509NameO
	case "2.5.4.11":
		ret = enums.X509NameOU
	case "2.5.4.12":
		ret = enums.X509NameT
	case "2.5.4.3":
		ret = enums.X509NameCN
	case "2.5.4.9":
		ret = enums.X509NameSTREET
	case "2.5.4.5":
		ret = enums.X509NameSerialNumber
	case "2.5.4.7":
		ret = enums.X509NameL
	case "2.5.4.8":
		ret = enums.X509NameST
	case "2.5.4.4":
		ret = enums.X509NameSurName
	case "2.5.4.42":
		ret = enums.X509NameGivenName
	case "2.5.4.43":
		ret = enums.X509NameInitials
	case "2.5.4.44":
		ret = enums.X509NameGeneration
	case "2.5.4.45":
		ret = enums.X509NameUniqueIdentifier
	case "2.5.4.15":
		ret = enums.X509NameBusinessCategory
	case "2.5.4.17":
		ret = enums.X509NamePostalCode
	case "2.5.4.46":
		ret = enums.X509NameDnQualifier
	case "2.5.4.65":
		ret = enums.X509NamePseudonym
	case "1.3.6.1.5.5.7.9.1":
		ret = enums.X509NameDateOfBirth
	case "1.3.6.1.5.5.7.9.2":
		ret = enums.X509NamePlaceOfBirth
	case "1.3.6.1.5.5.7.9.3":
		ret = enums.X509NameGender
	case "1.3.6.1.5.5.7.9.4":
		ret = enums.X509NameCountryOfCitizenship
	case "1.3.6.1.5.5.7.9.5":
		ret = enums.X509NameCountryOfResidence
	case "1.3.36.8.3.14":
		ret = enums.X509NameNameAtBirth
	case "2.5.4.16":
		ret = enums.X509NamePostalAddress
	case "2.5.4.54":
		ret = enums.X509NameDmdName
	case "2.5.4.20":
		ret = enums.X509NameTelephoneNumber
	case "2.5.4.41":
		ret = enums.X509NameName
	case "1.2.840.113549.1.9.1":
		ret = enums.X509NameE
	case "0.9.2342.19200300.100.1.25":
		ret = enums.X509NameDC
	case "0.9.2342.19200300.100.1.1":
		ret = enums.X509NameUID
	default:
		ret = enums.X509NameNone
	}
	return ret
}
