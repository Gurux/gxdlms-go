package internal

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

// Reserved for internal use.
type HDLCInfo int

const (
	HDLCInfoMaxInfoTX    HDLCInfo = 0x5
	HDLCInfoMaxInfoRX    HDLCInfo = 0x6
	HDLCInfoWindowSizeTX HDLCInfo = 0x7
	HDLCInfoWindowSizeRX HDLCInfo = 0x8
)

// HDLCInfoParse converts the given string into a HDLCInfo value.
//
// It returns the corresponding HDLCInfo constant if the string matches
// a known level name, or an error if the input is invalid.
func HDLCInfoParse(value string) (HDLCInfo, error) {
	var ret HDLCInfo
	var err error
	switch strings.ToUpper(value) {
	case "MAXINFOTX":
		ret = HDLCInfoMaxInfoTX
	case "MAXINFORX":
		ret = HDLCInfoMaxInfoRX
	case "WINDOWSIZETX":
		ret = HDLCInfoWindowSizeTX
	case "WINDOWSIZERX":
		ret = HDLCInfoWindowSizeRX
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the HDLCInfo.
// It satisfies fmt.Stringer.
func (g HDLCInfo) String() string {
	var ret string
	switch g {
	case HDLCInfoMaxInfoTX:
		ret = "MAXINFOTX"
	case HDLCInfoMaxInfoRX:
		ret = "MAXINFORX"
	case HDLCInfoWindowSizeTX:
		ret = "WINDOWSIZETX"
	case HDLCInfoWindowSizeRX:
		ret = "WINDOWSIZERX"
	}
	return ret
}
