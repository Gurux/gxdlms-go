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

// Access :  describes access errors.
type Access int

const (
	// AccessOther defines that the other error has occurred.
	AccessOther Access = iota
	// AccessScopeOfAccessViolated defines that the scope of access violated.
	AccessScopeOfAccessViolated
	// AccessObjectAccessInvalid defines that the object access is invalid.
	AccessObjectAccessInvalid
	// AccessHardwareFault defines that the hardware fault has occurred.
	AccessHardwareFault
	// AccessObjectUnavailable defines that the object is unavailable.
	AccessObjectUnavailable
)

// AccessParse converts the given string into a Access value.
//
// It returns the corresponding Access constant if the string matches
// a known level name, or an error if the input is invalid.
func AccessParse(value string) (Access, error) {
	var ret Access
	var err error
	switch strings.ToUpper(value) {
	case "OTHER":
		ret = AccessOther
	case "SCOPEOFACCESSVIOLATED":
		ret = AccessScopeOfAccessViolated
	case "OBJECTACCESSINVALID":
		ret = AccessObjectAccessInvalid
	case "HARDWAREFAULT":
		ret = AccessHardwareFault
	case "OBJECTUNAVAILABLE":
		ret = AccessObjectUnavailable
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the Access.
// It satisfies fmt.Stringer.
func (g Access) String() string {
	var ret string
	switch g {
	case AccessOther:
		ret = "OTHER"
	case AccessScopeOfAccessViolated:
		ret = "SCOPEOFACCESSVIOLATED"
	case AccessObjectAccessInvalid:
		ret = "OBJECTACCESSINVALID"
	case AccessHardwareFault:
		ret = "HARDWAREFAULT"
	case AccessObjectUnavailable:
		ret = "OBJECTUNAVAILABLE"
	}
	return ret
}
