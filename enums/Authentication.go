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

// Authentication enumerates the authentication levels.
type Authentication int

const (
	// AuthenticationNone defines that the no authentication is used.
	//  India DLMS standard IS 15959 uses name "Public client".
	AuthenticationNone Authentication = iota
	// AuthenticationLow defines that the low authentication is used.
	//  India DLMS standard IS 15959 uses name "Meter reading".
	AuthenticationLow
	// AuthenticationHigh defines that the high authentication is used.
	//  Because DLMS/COSEM specification does not
	//  specify details of the HLS mechanism Indian standard is implemented.
	AuthenticationHigh
	// AuthenticationHighMD5 defines that the high authentication is used. Password is hashed with MD5.
	AuthenticationHighMD5
	// AuthenticationHighSHA1 defines that the high authentication is used. Password is hashed with SHA1.
	AuthenticationHighSHA1
	// AuthenticationHighGMAC defines that the high authentication is used. Password is hashed with GMAC.
	AuthenticationHighGMAC
	// AuthenticationHighSHA256 defines that the high authentication is used. Password is hashed with SHA-256.
	AuthenticationHighSHA256
	// AuthenticationHighECDSA defines that the high authentication is used. Password is hashed with ECDSA.
	AuthenticationHighECDSA
)

// AuthenticationParse converts the given string into a Authentication value.
//
// It returns the corresponding Authentication constant if the string matches
// a known level name, or an error if the input is invalid.
func AuthenticationParse(value string) (Authentication, error) {
	var ret Authentication
	var err error
	switch {
	case strings.EqualFold(value, "None"):
		ret = AuthenticationNone
	case strings.EqualFold(value, "Low"):
		ret = AuthenticationLow
	case strings.EqualFold(value, "High"):
		ret = AuthenticationHigh
	case strings.EqualFold(value, "HighMD5"):
		ret = AuthenticationHighMD5
	case strings.EqualFold(value, "HighSHA1"):
		ret = AuthenticationHighSHA1
	case strings.EqualFold(value, "HighGMAC"):
		ret = AuthenticationHighGMAC
	case strings.EqualFold(value, "HighSHA256"):
		ret = AuthenticationHighSHA256
	case strings.EqualFold(value, "HighECDSA"):
		ret = AuthenticationHighECDSA
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the Authentication.
// It satisfies fmt.Stringer.
func (g Authentication) String() string {
	var ret string
	switch g {
	case AuthenticationNone:
		ret = "None"
	case AuthenticationLow:
		ret = "Low"
	case AuthenticationHigh:
		ret = "High"
	case AuthenticationHighMD5:
		ret = "HighMD5"
	case AuthenticationHighSHA1:
		ret = "HighSHA1"
	case AuthenticationHighGMAC:
		ret = "HighGMAC"
	case AuthenticationHighSHA256:
		ret = "HighSHA256"
	case AuthenticationHighECDSA:
		ret = "HighECDSA"
	}
	return ret
}

// AllAuthentication returns a slice containing all defined Authentication values.
func AllAuthentication() []Authentication {
	return []Authentication{
		AuthenticationNone,
		AuthenticationLow,
		AuthenticationHigh,
		AuthenticationHighMD5,
		AuthenticationHighSHA1,
		AuthenticationHighGMAC,
		AuthenticationHighSHA256,
		AuthenticationHighECDSA,
	}
}
