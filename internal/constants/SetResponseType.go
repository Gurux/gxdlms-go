package constants

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

// Enumerates  s.
type SetResponseType int

const (
	// SetResponseTypeNormal defines that the // Normal set response.
	SetResponseTypeNormal SetResponseType = 1
	// SetResponseTypeDataBlock defines that the // Set response in data blocks.
	SetResponseTypeDataBlock SetResponseType = 2
	// SetResponseTypeLastDataBlock defines that the // Set response in last data block.
	SetResponseTypeLastDataBlock SetResponseType = 3
	// SetResponseTypeLastDataBlockWithList defines that the // Set response with list in last data block.
	SetResponseTypeLastDataBlockWithList SetResponseType = 4
	// SetResponseTypeWithList defines that the // Set with list response.
	SetResponseTypeWithList SetResponseType = 5
)

// SetResponseTypeParse converts the given string into a SetResponseType value.
//
// It returns the corresponding SetResponseType constant if the string matches
// a known level name, or an error if the input is invalid.
func SetResponseTypeParse(value string) (SetResponseType, error) {
	var ret SetResponseType
	var err error
	switch {
	case strings.EqualFold(value, "Normal"):
		ret = SetResponseTypeNormal
	case strings.EqualFold(value, "DataBlock"):
		ret = SetResponseTypeDataBlock
	case strings.EqualFold(value, "LastDataBlock"):
		ret = SetResponseTypeLastDataBlock
	case strings.EqualFold(value, "LastDataBlockWithList"):
		ret = SetResponseTypeLastDataBlockWithList
	case strings.EqualFold(value, "WithList"):
		ret = SetResponseTypeWithList
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the SetResponseType.
// It satisfies fmt.Stringer.
func (g SetResponseType) String() string {
	var ret string
	switch g {
	case SetResponseTypeNormal:
		ret = "Normal"
	case SetResponseTypeDataBlock:
		ret = "DataBlock"
	case SetResponseTypeLastDataBlock:
		ret = "LastDataBlock"
	case SetResponseTypeLastDataBlockWithList:
		ret = "LastDataBlockWithList"
	case SetResponseTypeWithList:
		ret = "WithList"
	}
	return ret
}

// AllSetResponseType returns a slice containing all defined SetResponseType values.
func AllSetResponseType() []SetResponseType {
	return []SetResponseType{
		SetResponseTypeNormal,
		SetResponseTypeDataBlock,
		SetResponseTypeLastDataBlock,
		SetResponseTypeLastDataBlockWithList,
		SetResponseTypeWithList,
	}
}
