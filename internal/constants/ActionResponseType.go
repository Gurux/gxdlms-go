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

// ActionResponseType enumerates action response types.
type ActionResponseType int

const (
	// ActionResponseTypeNormal defines that the // Normal action.
	ActionResponseTypeNormal ActionResponseType = 1
	// ActionResponseTypeWithBlock defines that the // Action with block.
	ActionResponseTypeWithBlock ActionResponseType = 2
	// ActionResponseTypeWithList defines that the // Action with list.
	ActionResponseTypeWithList ActionResponseType = 3
	// ActionResponseTypeNextBlock defines that the // Action with next block.
	ActionResponseTypeNextBlock ActionResponseType = 4
)

// ActionResponseTypeParse converts the given string into a ActionResponseType value.
//
// It returns the corresponding ActionResponseType constant if the string matches
// a known level name, or an error if the input is invalid.
func ActionResponseTypeParse(value string) (ActionResponseType, error) {
	var ret ActionResponseType
	var err error
	switch {
	case strings.EqualFold(value, "Normal"):
		ret = ActionResponseTypeNormal
	case strings.EqualFold(value, "WithBlock"):
		ret = ActionResponseTypeWithBlock
	case strings.EqualFold(value, "WithList"):
		ret = ActionResponseTypeWithList
	case strings.EqualFold(value, "NextBlock"):
		ret = ActionResponseTypeNextBlock
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the ActionResponseType.
// It satisfies fmt.Stringer.
func (g ActionResponseType) String() string {
	var ret string
	switch g {
	case ActionResponseTypeNormal:
		ret = "Normal"
	case ActionResponseTypeWithBlock:
		ret = "WithBlock"
	case ActionResponseTypeWithList:
		ret = "WithList"
	case ActionResponseTypeNextBlock:
		ret = "NextBlock"
	}
	return ret
}

// AllActionResponseType returns a slice containing all defined ActionResponseType values.
func AllActionResponseType() []ActionResponseType {
	return []ActionResponseType{
		ActionResponseTypeNormal,
		ActionResponseTypeWithBlock,
		ActionResponseTypeWithList,
		ActionResponseTypeNextBlock,
	}
}
