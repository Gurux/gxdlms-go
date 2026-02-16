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

package enums

import (
	"fmt"
	"strings"

	"github.com/Gurux/gxcommon-go"
)

// Security Enumerates used security.
type Security int

const (
	// SecurityNone defines that the transport security is not used.
	SecurityNone Security = iota
	// SecurityAuthentication defines that the authentication security is used.
	SecurityAuthentication Security = 0x10
	// SecurityEncryption defines that the encryption security is used.
	SecurityEncryption Security = 0x20
	// SecurityAuthenticationEncryption defines that the authentication and Encryption security are used.
	SecurityAuthenticationEncryption Security = 0x30
)

// SecurityParse converts the given string into a Security value.
//
// It returns the corresponding Security constant if the string matches
// a known level name, or an error if the input is invalid.
func SecurityParse(value string) (Security, error) {
	var ret Security
	var err error
	switch strings.ToUpper(value) {
	case "NONE":
		ret = SecurityNone
	case "AUTHENTICATION":
		ret = SecurityAuthentication
	case "ENCRYPTION":
		ret = SecurityEncryption
	case "AUTHENTICATIONENCRYPTION":
		ret = SecurityAuthenticationEncryption
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the Security.
// It satisfies fmt.Stringer.
func (g Security) String() string {
	var ret string
	switch g {
	case SecurityNone:
		ret = "NONE"
	case SecurityAuthentication:
		ret = "AUTHENTICATION"
	case SecurityEncryption:
		ret = "ENCRYPTION"
	case SecurityAuthenticationEncryption:
		ret = "AUTHENTICATIONENCRYPTION"
	}
	return ret
}
