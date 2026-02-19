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

// NtpAuthenticationMethod Defines NTP authentication methods.
type NtpAuthenticationMethod int

const (
	// NtpAuthenticationMethodNoSecurity defines that the no security is used.
	NtpAuthenticationMethodNoSecurity NtpAuthenticationMethod = iota
	// NtpAuthenticationMethodSharedSecrets defines that the shared secrets are used.
	NtpAuthenticationMethodSharedSecrets
	// NtpAuthenticationMethodAutoKeyIff defines that the iFF auto key is used.
	NtpAuthenticationMethodAutoKeyIff
)

// NtpAuthenticationMethodParse converts the given string into a NtpAuthenticationMethod value.
//
// It returns the corresponding NtpAuthenticationMethod constant if the string matches
// a known level name, or an error if the input is invalid.
func NtpAuthenticationMethodParse(value string) (NtpAuthenticationMethod, error) {
	var ret NtpAuthenticationMethod
	var err error
	switch {
	case strings.EqualFold(value, "NoSecurity"):
		ret = NtpAuthenticationMethodNoSecurity
	case strings.EqualFold(value, "SharedSecrets"):
		ret = NtpAuthenticationMethodSharedSecrets
	case strings.EqualFold(value, "AutoKeyIff"):
		ret = NtpAuthenticationMethodAutoKeyIff
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the NtpAuthenticationMethod.
// It satisfies fmt.Stringer.
func (g NtpAuthenticationMethod) String() string {
	var ret string
	switch g {
	case NtpAuthenticationMethodNoSecurity:
		ret = "NoSecurity"
	case NtpAuthenticationMethodSharedSecrets:
		ret = "SharedSecrets"
	case NtpAuthenticationMethodAutoKeyIff:
		ret = "AutoKeyIff"
	}
	return ret
}

// AllNtpAuthenticationMethod returns a slice containing all defined NtpAuthenticationMethod values.
func AllNtpAuthenticationMethod() []NtpAuthenticationMethod {
	return []NtpAuthenticationMethod{
		NtpAuthenticationMethodNoSecurity,
		NtpAuthenticationMethodSharedSecrets,
		NtpAuthenticationMethodAutoKeyIff,
	}
}
