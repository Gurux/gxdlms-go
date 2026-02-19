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

// MBusDataHeaderType Defines the M-Bus data header types.
type MBusDataHeaderType int

const (
	// MBusDataHeaderTypeNone defines that the m-Bus data header is not used.
	MBusDataHeaderTypeNone MBusDataHeaderType = iota
	// MBusDataHeaderTypeShort defines that the short data header is used.
	MBusDataHeaderTypeShort
	// MBusDataHeaderTypeLong defines that the long data header is used.
	MBusDataHeaderTypeLong
)

// MBusDataHeaderTypeParse converts the given string into a MBusDataHeaderType value.
//
// It returns the corresponding MBusDataHeaderType constant if the string matches
// a known level name, or an error if the input is invalid.
func MBusDataHeaderTypeParse(value string) (MBusDataHeaderType, error) {
	var ret MBusDataHeaderType
	var err error
	switch {
	case strings.EqualFold(value, "None"):
		ret = MBusDataHeaderTypeNone
	case strings.EqualFold(value, "Short"):
		ret = MBusDataHeaderTypeShort
	case strings.EqualFold(value, "Long"):
		ret = MBusDataHeaderTypeLong
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the MBusDataHeaderType.
// It satisfies fmt.Stringer.
func (g MBusDataHeaderType) String() string {
	var ret string
	switch g {
	case MBusDataHeaderTypeNone:
		ret = "None"
	case MBusDataHeaderTypeShort:
		ret = "Short"
	case MBusDataHeaderTypeLong:
		ret = "Long"
	}
	return ret
}

// AllMBusDataHeaderType returns a slice containing all defined MBusDataHeaderType values.
func AllMBusDataHeaderType() []MBusDataHeaderType {
	return []MBusDataHeaderType{
		MBusDataHeaderTypeNone,
		MBusDataHeaderTypeShort,
		MBusDataHeaderTypeLong,
	}
}
