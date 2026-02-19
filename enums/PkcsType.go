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

// PkcsType Public-Key Cryptography Standards (PKCS) type.
type PkcsType int

const (
	// PkcsTypeNone defines that the unknown certificate type.
	PkcsTypeNone PkcsType = iota
	// PkcsTypePkcs8 defines that the pKCS 8 is used with private key.
	PkcsTypePkcs8
	// PkcsTypePkcs10 defines that the pKCS 10 is used with Certificate Signing Request.
	PkcsTypePkcs10
	// PkcsTypex509Certificate defines that the x509Certificate is used with public key.
	PkcsTypex509Certificate
)

// PkcsTypeParse converts the given string into a PkcsType value.
//
// It returns the corresponding PkcsType constant if the string matches
// a known level name, or an error if the input is invalid.
func PkcsTypeParse(value string) (PkcsType, error) {
	var ret PkcsType
	var err error
	switch {
	case strings.EqualFold(value, "None"):
		ret = PkcsTypeNone
	case strings.EqualFold(value, "Pkcs8"):
		ret = PkcsTypePkcs8
	case strings.EqualFold(value, "Pkcs10"):
		ret = PkcsTypePkcs10
	case strings.EqualFold(value, "x509Certificate"):
		ret = PkcsTypex509Certificate
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the PkcsType.
// It satisfies fmt.Stringer.
func (g PkcsType) String() string {
	var ret string
	switch g {
	case PkcsTypeNone:
		ret = "None"
	case PkcsTypePkcs8:
		ret = "Pkcs8"
	case PkcsTypePkcs10:
		ret = "Pkcs10"
	case PkcsTypex509Certificate:
		ret = "x509Certificate"
	}
	return ret
}

// AllPkcsType returns a slice containing all defined PkcsType values.
func AllPkcsType() []PkcsType {
	return []PkcsType{
		PkcsTypeNone,
		PkcsTypePkcs8,
		PkcsTypePkcs10,
		PkcsTypex509Certificate,
	}
}
