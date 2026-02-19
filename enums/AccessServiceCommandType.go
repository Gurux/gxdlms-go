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

// Enumerates Access Service types.
type AccessServiceCommandType int

const (
	// AccessServiceCommandTypeGet defines that the // Get request or response.
	AccessServiceCommandTypeGet AccessServiceCommandType = 1
	// AccessServiceCommandTypeSet defines that the // Set request or response.
	AccessServiceCommandTypeSet AccessServiceCommandType = 2
	// AccessServiceCommandTypeAction defines that the // Action request or response.
	AccessServiceCommandTypeAction AccessServiceCommandType = 3
)

// AccessServiceCommandTypeParse converts the given string into a AccessServiceCommandType value.
//
// It returns the corresponding AccessServiceCommandType constant if the string matches
// a known level name, or an error if the input is invalid.
func AccessServiceCommandTypeParse(value string) (AccessServiceCommandType, error) {
	var ret AccessServiceCommandType
	var err error
	switch {
	case strings.EqualFold(value, "Get"):
		ret = AccessServiceCommandTypeGet
	case strings.EqualFold(value, "Set"):
		ret = AccessServiceCommandTypeSet
	case strings.EqualFold(value, "Action"):
		ret = AccessServiceCommandTypeAction
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the AccessServiceCommandType.
// It satisfies fmt.Stringer.
func (g AccessServiceCommandType) String() string {
	var ret string
	switch g {
	case AccessServiceCommandTypeGet:
		ret = "Get"
	case AccessServiceCommandTypeSet:
		ret = "Set"
	case AccessServiceCommandTypeAction:
		ret = "Action"
	}
	return ret
}

// AllAccessServiceCommandType returns a slice containing all defined AccessServiceCommandType values.
func AllAccessServiceCommandType() []AccessServiceCommandType {
	return []AccessServiceCommandType{
	AccessServiceCommandTypeGet,
	AccessServiceCommandTypeSet,
	AccessServiceCommandTypeAction,
	}
}
