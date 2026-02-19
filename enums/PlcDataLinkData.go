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

// PLC data link data commands.
type PlcDataLinkData int

const (
	PlcDataLinkDataRequest PlcDataLinkData = 0x90
)

// PlcDataLinkDataParse converts the given string into a PlcDataLinkData value.
//
// It returns the corresponding PlcDataLinkData constant if the string matches
// a known level name, or an error if the input is invalid.
func PlcDataLinkDataParse(value string) (PlcDataLinkData, error) {
	var ret PlcDataLinkData
	var err error
	switch {
	case strings.EqualFold(value, "Request"):
		ret = PlcDataLinkDataRequest
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the PlcDataLinkData.
// It satisfies fmt.Stringer.
func (g PlcDataLinkData) String() string {
	var ret string
	switch g {
	case PlcDataLinkDataRequest:
		ret = "Request"
	}
	return ret
}

// AllPlcDataLinkData returns a slice containing all defined PlcDataLinkData values.
func AllPlcDataLinkData() []PlcDataLinkData {
	return []PlcDataLinkData{
	PlcDataLinkDataRequest,
	}
}
