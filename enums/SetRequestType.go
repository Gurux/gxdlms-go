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

// Enumerates  s.
type SetRequestType int

const (
	// SetRequestTypeNormal defines that the // Normal Set.
	SetRequestTypeNormal SetRequestType = 1
	// SetRequestTypeFirstDataBlock defines that the // Set with first data block.
	SetRequestTypeFirstDataBlock SetRequestType = 2
	// SetRequestTypeWithDataBlock defines that the // Set with data block.
	SetRequestTypeWithDataBlock SetRequestType = 3
	// SetRequestTypeWithList defines that the // Set with list.
	SetRequestTypeWithList SetRequestType = 4
	// SetRequestTypeWithListAndWithFirstDatablock defines that the // Set with list and first data block.
	SetRequestTypeWithListAndWithFirstDatablock SetRequestType = 5
)

// SetRequestTypeParse converts the given string into a SetRequestType value.
//
// It returns the corresponding SetRequestType constant if the string matches
// a known level name, or an error if the input is invalid.
func SetRequestTypeParse(value string) (SetRequestType, error) {
	var ret SetRequestType
	var err error
	switch strings.ToUpper(value) {
	case "NORMAL":
		ret = SetRequestTypeNormal
	case "FIRSTDATABLOCK":
		ret = SetRequestTypeFirstDataBlock
	case "WITHDATABLOCK":
		ret = SetRequestTypeWithDataBlock
	case "WITHLIST":
		ret = SetRequestTypeWithList
	case "WITHLISTANDWITHFIRSTDATABLOCK":
		ret = SetRequestTypeWithListAndWithFirstDatablock
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the SetRequestType.
// It satisfies fmt.Stringer.
func (g SetRequestType) String() string {
	var ret string
	switch g {
	case SetRequestTypeNormal:
		ret = "NORMAL"
	case SetRequestTypeFirstDataBlock:
		ret = "FIRSTDATABLOCK"
	case SetRequestTypeWithDataBlock:
		ret = "WITHDATABLOCK"
	case SetRequestTypeWithList:
		ret = "WITHLIST"
	case SetRequestTypeWithListAndWithFirstDatablock:
		ret = "WITHLISTANDWITHFIRSTDATABLOCK"
	}
	return ret
}
