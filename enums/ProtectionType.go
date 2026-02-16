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

// ProtectionType Enumerated data protection types.
type ProtectionType int

const (
	// ProtectionTypeAuthentication defines that the authentication.
	ProtectionTypeAuthentication ProtectionType = iota
	// ProtectionTypeEncryption defines that the encryption
	ProtectionTypeEncryption
	// ProtectionTypeAuthenticationEncryption defines that the authentication and encryption.
	ProtectionTypeAuthenticationEncryption
	// ProtectionTypeDigitalSignature defines that the digital signature
	ProtectionTypeDigitalSignature
)

// ProtectionTypeParse converts the given string into a ProtectionType value.
//
// It returns the corresponding ProtectionType constant if the string matches
// a known level name, or an error if the input is invalid.
func ProtectionTypeParse(value string) (ProtectionType, error) {
	var ret ProtectionType
	var err error
	switch strings.ToUpper(value) {
	case "AUTHENTICATION":
		ret = ProtectionTypeAuthentication
	case "ENCRYPTION":
		ret = ProtectionTypeEncryption
	case "AUTHENTICATIONENCRYPTION":
		ret = ProtectionTypeAuthenticationEncryption
	case "DIGITALSIGNATURE":
		ret = ProtectionTypeDigitalSignature
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the ProtectionType.
// It satisfies fmt.Stringer.
func (g ProtectionType) String() string {
	var ret string
	switch g {
	case ProtectionTypeAuthentication:
		ret = "AUTHENTICATION"
	case ProtectionTypeEncryption:
		ret = "ENCRYPTION"
	case ProtectionTypeAuthenticationEncryption:
		ret = "AUTHENTICATIONENCRYPTION"
	case ProtectionTypeDigitalSignature:
		ret = "DIGITALSIGNATURE"
	}
	return ret
}
