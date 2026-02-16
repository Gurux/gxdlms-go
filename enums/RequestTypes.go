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

// RequestTypes defines the request types for more data handling.
type RequestTypes int

const (
	// RequestTypesNone defines that there is no more data available.
	RequestTypesNone RequestTypes = 0
	// RequestTypesDataBlock defines that data block request.
	RequestTypesDataBlock RequestTypes = 1
	// RequestTypesFrame defines that frame request.
	RequestTypesFrame RequestTypes = 2
	// RequestTypesGBT defines that general block transfer request.
	RequestTypesGBT RequestTypes = 4
)

// RequestTypesParse converts the given string into a RequestTypes value.
//
// It returns the corresponding RequestTypes constant if the string matches
// a known level name, or an error if the input is invalid.
func RequestTypesParse(value string) (RequestTypes, error) {
	var ret RequestTypes
	var err error
	switch strings.ToUpper(value) {
	case "NONE":
		ret = RequestTypesNone
	case "DATABLOCK":
		ret = RequestTypesDataBlock
	case "FRAME":
		ret = RequestTypesFrame
	case "GBT":
		ret = RequestTypesGBT
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the RequestTypes.
// It satisfies fmt.Stringer.
func (r RequestTypes) String() string {
	var ret string
	switch r {
	case RequestTypesNone:
		ret = "NONE"
	case RequestTypesDataBlock:
		ret = "DATABLOCK"
	case RequestTypesFrame:
		ret = "FRAME"
	case RequestTypesGBT:
		ret = "GBT"
	}
	return ret
}
