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

// X509Name X509 names.
type X509Name int

const (
	X509NameNone X509Name = iota
	// X509NameC defines that the country code - StringType(SIZE(2))
	X509NameC
	// X509NameO defines that the organization - StringType(SIZE(1..64))
	X509NameO
	// X509NameOU defines that the organizational unit name - StringType(SIZE(1..64))
	X509NameOU
	// X509NameT defines that the title
	X509NameT
	// X509NameCN defines that the common name - StringType(SIZE(1..64))
	X509NameCN
	// X509NameSTREET defines that the street - StringType(SIZE(1..64))
	X509NameSTREET
	// X509NameSerialNumber defines that the device serial number name - StringType(SIZE(1..64))
	X509NameSerialNumber
	// X509NameL defines that the locality name - StringType(SIZE(1..64))
	X509NameL
	// X509NameST defines that the state, or province name - StringType(SIZE(1..64))
	X509NameST
	// X509NameSurName defines that the naming attributes of type X520name
	X509NameSurName
	// X509NameGivenName defines that the given name.
	X509NameGivenName
	// X509NameInitials defines that the initials.
	X509NameInitials
	// X509NameGeneration defines that the generation.
	X509NameGeneration
	// X509NameUniqueIdentifier defines that the unique identifier.
	X509NameUniqueIdentifier
	// X509NameBusinessCategory defines that the businessCategory - DirectoryString(SIZE(1..128))
	X509NameBusinessCategory
	// X509NamePostalCode defines that the postalCode - DirectoryString(SIZE(1..40))
	X509NamePostalCode
	// X509NameDnQualifier defines that the dnQualifier - DirectoryString(SIZE(1..64))
	X509NameDnQualifier
	// X509NamePseudonym defines that the rFC 3039 Pseudonym - DirectoryString(SIZE(1..64))
	X509NamePseudonym
	// X509NameDateOfBirth defines that the rFC 3039 DateOfBirth - GeneralizedTime - YYYYMMDD000000Z
	X509NameDateOfBirth
	// X509NamePlaceOfBirth defines that the rFC 3039 PlaceOfBirth - DirectoryString(SIZE(1..128))
	X509NamePlaceOfBirth
	// X509NameGender defines that the rFC 3039 DateOfBirth - PrintableString (SIZE(1 -- "M", "F", "m" or "f")
	X509NameGender
	// X509NameCountryOfCitizenship defines that the rFC 3039 CountryOfCitizenship - PrintableString (SIZE (2 -- ISO 3166)) codes only
	X509NameCountryOfCitizenship
	// X509NameCountryOfResidence defines that the rFC 3039 CountryOfCitizenship - PrintableString (SIZE (2 -- ISO 3166)) codes only
	X509NameCountryOfResidence
	// X509NameNameAtBirth defines that the iSIS-MTT NameAtBirth - DirectoryString(SIZE(1..64))
	X509NameNameAtBirth
	// X509NamePostalAddress defines that the rFC 3039 PostalAddress - SEQUENCE SIZE (1..6 OF DirectoryString(SIZE(1..30)))
	X509NamePostalAddress
	// X509NameDmdName defines that the rFC 2256 dmdName
	X509NameDmdName
	// X509NameTelephoneNumber defines that the id-at-telephoneNumber
	X509NameTelephoneNumber
	// X509NameName defines that the id-at-name
	X509NameName
	// X509NameE defines that the email address in Verisign certificates
	X509NameE
	// X509NameDC defines that the domain component
	X509NameDC
	// X509NameUID defines that the lDAP User id.
	X509NameUID
)

// X509NameParse converts the given string into a X509Name value.
//
// It returns the corresponding X509Name constant if the string matches
// a known level name, or an error if the input is invalid.
func X509NameParse(value string) (X509Name, error) {
	var ret X509Name
	var err error
	switch {
	case strings.EqualFold(value, "None"):
		ret = X509NameNone
	case strings.EqualFold(value, "C"):
		ret = X509NameC
	case strings.EqualFold(value, "O"):
		ret = X509NameO
	case strings.EqualFold(value, "OU"):
		ret = X509NameOU
	case strings.EqualFold(value, "T"):
		ret = X509NameT
	case strings.EqualFold(value, "CN"):
		ret = X509NameCN
	case strings.EqualFold(value, "STREET"):
		ret = X509NameSTREET
	case strings.EqualFold(value, "SerialNumber"):
		ret = X509NameSerialNumber
	case strings.EqualFold(value, "L"):
		ret = X509NameL
	case strings.EqualFold(value, "ST"):
		ret = X509NameST
	case strings.EqualFold(value, "SurName"):
		ret = X509NameSurName
	case strings.EqualFold(value, "GivenName"):
		ret = X509NameGivenName
	case strings.EqualFold(value, "Initials"):
		ret = X509NameInitials
	case strings.EqualFold(value, "Generation"):
		ret = X509NameGeneration
	case strings.EqualFold(value, "UniqueIdentifier"):
		ret = X509NameUniqueIdentifier
	case strings.EqualFold(value, "BusinessCategory"):
		ret = X509NameBusinessCategory
	case strings.EqualFold(value, "PostalCode"):
		ret = X509NamePostalCode
	case strings.EqualFold(value, "DnQualifier"):
		ret = X509NameDnQualifier
	case strings.EqualFold(value, "Pseudonym"):
		ret = X509NamePseudonym
	case strings.EqualFold(value, "DateOfBirth"):
		ret = X509NameDateOfBirth
	case strings.EqualFold(value, "PlaceOfBirth"):
		ret = X509NamePlaceOfBirth
	case strings.EqualFold(value, "Gender"):
		ret = X509NameGender
	case strings.EqualFold(value, "CountryOfCitizenship"):
		ret = X509NameCountryOfCitizenship
	case strings.EqualFold(value, "CountryOfResidence"):
		ret = X509NameCountryOfResidence
	case strings.EqualFold(value, "NameAtBirth"):
		ret = X509NameNameAtBirth
	case strings.EqualFold(value, "PostalAddress"):
		ret = X509NamePostalAddress
	case strings.EqualFold(value, "DmdName"):
		ret = X509NameDmdName
	case strings.EqualFold(value, "TelephoneNumber"):
		ret = X509NameTelephoneNumber
	case strings.EqualFold(value, "Name"):
		ret = X509NameName
	case strings.EqualFold(value, "E"):
		ret = X509NameE
	case strings.EqualFold(value, "DC"):
		ret = X509NameDC
	case strings.EqualFold(value, "UID"):
		ret = X509NameUID
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the X509Name.
// It satisfies fmt.Stringer.
func (g X509Name) String() string {
	var ret string
	switch g {
	case X509NameNone:
		ret = "None"
	case X509NameC:
		ret = "C"
	case X509NameO:
		ret = "O"
	case X509NameOU:
		ret = "OU"
	case X509NameT:
		ret = "T"
	case X509NameCN:
		ret = "CN"
	case X509NameSTREET:
		ret = "STREET"
	case X509NameSerialNumber:
		ret = "SerialNumber"
	case X509NameL:
		ret = "L"
	case X509NameST:
		ret = "ST"
	case X509NameSurName:
		ret = "SurName"
	case X509NameGivenName:
		ret = "GivenName"
	case X509NameInitials:
		ret = "Initials"
	case X509NameGeneration:
		ret = "Generation"
	case X509NameUniqueIdentifier:
		ret = "UniqueIdentifier"
	case X509NameBusinessCategory:
		ret = "BusinessCategory"
	case X509NamePostalCode:
		ret = "PostalCode"
	case X509NameDnQualifier:
		ret = "DnQualifier"
	case X509NamePseudonym:
		ret = "Pseudonym"
	case X509NameDateOfBirth:
		ret = "DateOfBirth"
	case X509NamePlaceOfBirth:
		ret = "PlaceOfBirth"
	case X509NameGender:
		ret = "Gender"
	case X509NameCountryOfCitizenship:
		ret = "CountryOfCitizenship"
	case X509NameCountryOfResidence:
		ret = "CountryOfResidence"
	case X509NameNameAtBirth:
		ret = "NameAtBirth"
	case X509NamePostalAddress:
		ret = "PostalAddress"
	case X509NameDmdName:
		ret = "DmdName"
	case X509NameTelephoneNumber:
		ret = "TelephoneNumber"
	case X509NameName:
		ret = "Name"
	case X509NameE:
		ret = "E"
	case X509NameDC:
		ret = "DC"
	case X509NameUID:
		ret = "UID"
	}
	return ret
}

// AllX509Name returns a slice containing all defined X509Name values.
func AllX509Name() []X509Name {
	return []X509Name{
		X509NameNone,
		X509NameC,
		X509NameO,
		X509NameOU,
		X509NameT,
		X509NameCN,
		X509NameSTREET,
		X509NameSerialNumber,
		X509NameL,
		X509NameST,
		X509NameSurName,
		X509NameGivenName,
		X509NameInitials,
		X509NameGeneration,
		X509NameUniqueIdentifier,
		X509NameBusinessCategory,
		X509NamePostalCode,
		X509NameDnQualifier,
		X509NamePseudonym,
		X509NameDateOfBirth,
		X509NamePlaceOfBirth,
		X509NameGender,
		X509NameCountryOfCitizenship,
		X509NameCountryOfResidence,
		X509NameNameAtBirth,
		X509NamePostalAddress,
		X509NameDmdName,
		X509NameTelephoneNumber,
		X509NameName,
		X509NameE,
		X509NameDC,
		X509NameUID,
	}
}
