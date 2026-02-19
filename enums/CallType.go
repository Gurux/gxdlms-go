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

// CallType The purpose of the call.
type CallType int

const (
	// CallTypeNormal defines that the normal CSD call.
	CallTypeNormal CallType = iota
	// CallTypeWakeUp defines that the wake-up request.
	CallTypeWakeUp
)

// CallTypeParse converts the given string into a CallType value.
//
// It returns the corresponding CallType constant if the string matches
// a known level name, or an error if the input is invalid.
func CallTypeParse(value string) (CallType, error) {
	var ret CallType
	var err error
	switch {
	case strings.EqualFold(value, "Normal"):
		ret = CallTypeNormal
	case strings.EqualFold(value, "WakeUp"):
		ret = CallTypeWakeUp
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the CallType.
// It satisfies fmt.Stringer.
func (g CallType) String() string {
	var ret string
	switch g {
	case CallTypeNormal:
		ret = "Normal"
	case CallTypeWakeUp:
		ret = "WakeUp"
	}
	return ret
}

// AllCallType returns a slice containing all defined CallType values.
func AllCallType() []CallType {
	return []CallType{
	CallTypeNormal,
	CallTypeWakeUp,
	}
}
