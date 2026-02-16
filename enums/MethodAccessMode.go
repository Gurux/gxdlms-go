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

// MethodAccessMode The MethodAccessMode enumerates the method access modes.
type MethodAccessMode int

const (
	// MethodAccessModeNoAccess defines that the client can't use method.
	MethodAccessModeNoAccess MethodAccessMode = 0x0
	// MethodAccessModeAccess defines that the method is allowed to use.
	MethodAccessModeAccess MethodAccessMode = 0x1
	// MethodAccessModeAuthenticatedAccess defines that the authenticated access is allowed.
	MethodAccessModeAuthenticatedAccess MethodAccessMode = 0x2
)

// MethodAccessModeParse converts the given string into a MethodAccessMode value.
//
// It returns the corresponding MethodAccessMode constant if the string matches
// a known level name, or an error if the input is invalid.
func MethodAccessModeParse(value string) (MethodAccessMode, error) {
	var ret MethodAccessMode
	var err error
	switch strings.ToUpper(value) {
	case "NOACCESS":
		ret = MethodAccessModeNoAccess
	case "ACCESS":
		ret = MethodAccessModeAccess
	case "AUTHENTICATEDACCESS":
		ret = MethodAccessModeAuthenticatedAccess
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the MethodAccessMode.
// It satisfies fmt.Stringer.
func (g MethodAccessMode) String() string {
	var ret string
	switch g {
	case MethodAccessModeNoAccess:
		ret = "NOACCESS"
	case MethodAccessModeAccess:
		ret = "ACCESS"
	case MethodAccessModeAuthenticatedAccess:
		ret = "AUTHENTICATEDACCESS"
	}
	return ret
}
