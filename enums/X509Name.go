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
	switch strings.ToUpper(value) {
	case "NONE":
		ret = X509NameNone
	case "C":
		ret = X509NameC
	case "O":
		ret = X509NameO
	case "OU":
		ret = X509NameOU
	case "T":
		ret = X509NameT
	case "CN":
		ret = X509NameCN
	case "STREET":
		ret = X509NameSTREET
	case "SERIALNUMBER":
		ret = X509NameSerialNumber
	case "L":
		ret = X509NameL
	case "ST":
		ret = X509NameST
	case "SURNAME":
		ret = X509NameSurName
	case "GIVENNAME":
		ret = X509NameGivenName
	case "INITIALS":
		ret = X509NameInitials
	case "GENERATION":
		ret = X509NameGeneration
	case "UNIQUEIDENTIFIER":
		ret = X509NameUniqueIdentifier
	case "BUSINESSCATEGORY":
		ret = X509NameBusinessCategory
	case "POSTALCODE":
		ret = X509NamePostalCode
	case "DNQUALIFIER":
		ret = X509NameDnQualifier
	case "PSEUDONYM":
		ret = X509NamePseudonym
	case "DATEOFBIRTH":
		ret = X509NameDateOfBirth
	case "PLACEOFBIRTH":
		ret = X509NamePlaceOfBirth
	case "GENDER":
		ret = X509NameGender
	case "COUNTRYOFCITIZENSHIP":
		ret = X509NameCountryOfCitizenship
	case "COUNTRYOFRESIDENCE":
		ret = X509NameCountryOfResidence
	case "NAMEATBIRTH":
		ret = X509NameNameAtBirth
	case "POSTALADDRESS":
		ret = X509NamePostalAddress
	case "DMDNAME":
		ret = X509NameDmdName
	case "TELEPHONENUMBER":
		ret = X509NameTelephoneNumber
	case "NAME":
		ret = X509NameName
	case "E":
		ret = X509NameE
	case "DC":
		ret = X509NameDC
	case "UID":
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
		ret = "NONE"
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
		ret = "SERIALNUMBER"
	case X509NameL:
		ret = "L"
	case X509NameST:
		ret = "ST"
	case X509NameSurName:
		ret = "SURNAME"
	case X509NameGivenName:
		ret = "GIVENNAME"
	case X509NameInitials:
		ret = "INITIALS"
	case X509NameGeneration:
		ret = "GENERATION"
	case X509NameUniqueIdentifier:
		ret = "UNIQUEIDENTIFIER"
	case X509NameBusinessCategory:
		ret = "BUSINESSCATEGORY"
	case X509NamePostalCode:
		ret = "POSTALCODE"
	case X509NameDnQualifier:
		ret = "DNQUALIFIER"
	case X509NamePseudonym:
		ret = "PSEUDONYM"
	case X509NameDateOfBirth:
		ret = "DATEOFBIRTH"
	case X509NamePlaceOfBirth:
		ret = "PLACEOFBIRTH"
	case X509NameGender:
		ret = "GENDER"
	case X509NameCountryOfCitizenship:
		ret = "COUNTRYOFCITIZENSHIP"
	case X509NameCountryOfResidence:
		ret = "COUNTRYOFRESIDENCE"
	case X509NameNameAtBirth:
		ret = "NAMEATBIRTH"
	case X509NamePostalAddress:
		ret = "POSTALADDRESS"
	case X509NameDmdName:
		ret = "DMDNAME"
	case X509NameTelephoneNumber:
		ret = "TELEPHONENUMBER"
	case X509NameName:
		ret = "NAME"
	case X509NameE:
		ret = "E"
	case X509NameDC:
		ret = "DC"
	case X509NameUID:
		ret = "UID"
	}
	return ret
}
