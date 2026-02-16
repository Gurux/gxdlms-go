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

// ExtendedKeyUsage .
type ExtendedKeyUsage int

const (
	// ExtendedKeyUsageNone defines that extended key usage is not used bit is set.
	ExtendedKeyUsageNone ExtendedKeyUsage = iota
	// ExtendedKeyUsageServerAuth defines that certificate can be used as an TLS server certificate bit is set.
	ExtendedKeyUsageServerAuth
	// ExtendedKeyUsageClientAuth defines that certificate can be used as an TLS client certificate bit is set.
	ExtendedKeyUsageClientAuth
)

// ExtendedKeyUsageParse converts the given string into a ExtendedKeyUsage value.
//
// It returns the corresponding ExtendedKeyUsage constant if the string matches
// a known level name, or an error if the input is invalid.
func ExtendedKeyUsageParse(value string) (ExtendedKeyUsage, error) {
	var ret ExtendedKeyUsage
	var err error
	switch strings.ToUpper(value) {
	case "NONE":
		ret = ExtendedKeyUsageNone
	case "SERVERAUTH":
		ret = ExtendedKeyUsageServerAuth
	case "CLIENTAUTH":
		ret = ExtendedKeyUsageClientAuth
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the ExtendedKeyUsage.
// It satisfies fmt.Stringer.
func (g ExtendedKeyUsage) String() string {
	var ret string
	switch g {
	case ExtendedKeyUsageNone:
		ret = "NONE"
	case ExtendedKeyUsageServerAuth:
		ret = "SERVERAUTH"
	case ExtendedKeyUsageClientAuth:
		ret = "CLIENTAUTH"
	}
	return ret
}
