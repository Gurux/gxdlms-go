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

// GetCommandType Enumerates Get request and response types.
type GetCommandType int

const (
	// GetCommandTypeNormal defines that the normal Get.
	GetCommandTypeNormal GetCommandType = 1
	// GetCommandTypeNextDataBlock defines that the next data block.
	GetCommandTypeNextDataBlock GetCommandType = 2
	// GetCommandTypeWithList defines that the get request with list.
	GetCommandTypeWithList GetCommandType = 3
)

// GetCommandTypeParse converts the given string into a GetCommandType value.
//
// It returns the corresponding GetCommandType constant if the string matches
// a known level name, or an error if the input is invalid.
func GetCommandTypeParse(value string) (GetCommandType, error) {
	var ret GetCommandType
	var err error
	switch strings.ToUpper(value) {
	case "NORMAL":
		ret = GetCommandTypeNormal
	case "NEXTDATABLOCK":
		ret = GetCommandTypeNextDataBlock
	case "WITHLIST":
		ret = GetCommandTypeWithList
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the GetCommandType.
// It satisfies fmt.Stringer.
func (g GetCommandType) String() string {
	var ret string
	switch g {
	case GetCommandTypeNormal:
		ret = "NORMAL"
	case GetCommandTypeNextDataBlock:
		ret = "NEXTDATABLOCK"
	case GetCommandTypeWithList:
		ret = "WITHLIST"
	}
	return ret
}
