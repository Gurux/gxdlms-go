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

// Sequence number of MAC sub frame.
type PlcMacSubframes int

const (
	PlcMacSubframesOne PlcMacSubframes = 0x6C6C
	PlcMacSubframesTwo PlcMacSubframes = 0x3A3A
	PlcMacSubframesThree PlcMacSubframes = 0x5656
	PlcMacSubframesFour PlcMacSubframes = 0x7171
	PlcMacSubframesFive PlcMacSubframes = 0x1D1D
	PlcMacSubframesSix PlcMacSubframes = 0x4B4B
	PlcMacSubframesSeven PlcMacSubframes = 0x2727
)

// PlcMacSubframesParse converts the given string into a PlcMacSubframes value.
//
// It returns the corresponding PlcMacSubframes constant if the string matches
// a known level name, or an error if the input is invalid.
func PlcMacSubframesParse(value string) (PlcMacSubframes, error) {
	var ret PlcMacSubframes
	var err error
	switch {
	case strings.EqualFold(value, "One"):
		ret = PlcMacSubframesOne
	case strings.EqualFold(value, "Two"):
		ret = PlcMacSubframesTwo
	case strings.EqualFold(value, "Three"):
		ret = PlcMacSubframesThree
	case strings.EqualFold(value, "Four"):
		ret = PlcMacSubframesFour
	case strings.EqualFold(value, "Five"):
		ret = PlcMacSubframesFive
	case strings.EqualFold(value, "Six"):
		ret = PlcMacSubframesSix
	case strings.EqualFold(value, "Seven"):
		ret = PlcMacSubframesSeven
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the PlcMacSubframes.
// It satisfies fmt.Stringer.
func (g PlcMacSubframes) String() string {
	var ret string
	switch g {
	case PlcMacSubframesOne:
		ret = "One"
	case PlcMacSubframesTwo:
		ret = "Two"
	case PlcMacSubframesThree:
		ret = "Three"
	case PlcMacSubframesFour:
		ret = "Four"
	case PlcMacSubframesFive:
		ret = "Five"
	case PlcMacSubframesSix:
		ret = "Six"
	case PlcMacSubframesSeven:
		ret = "Seven"
	}
	return ret
}

// AllPlcMacSubframes returns a slice containing all defined PlcMacSubframes values.
func AllPlcMacSubframes() []PlcMacSubframes {
	return []PlcMacSubframes{
	PlcMacSubframesOne,
	PlcMacSubframesTwo,
	PlcMacSubframesThree,
	PlcMacSubframesFour,
	PlcMacSubframesFive,
	PlcMacSubframesSix,
	PlcMacSubframesSeven,
	}
}
