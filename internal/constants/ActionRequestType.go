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

// ActionRequestType enumerates action request types.
type ActionRequestType int

const (
	// ActionRequestTypeNormal defines that the // Normal action.
	ActionRequestTypeNormal ActionRequestType = 1
	// ActionRequestTypeNextBlock defines that the // Next block.
	ActionRequestTypeNextBlock ActionRequestType = 2
	// ActionRequestTypeWithList defines that the // Action with list.
	ActionRequestTypeWithList ActionRequestType = 3
	// ActionRequestTypeWithFirstBlock defines that the // Action with first block.
	ActionRequestTypeWithFirstBlock ActionRequestType = 4
	// ActionRequestTypeWithListAndFirstBlock defines that the // Action with list and first block.
	ActionRequestTypeWithListAndFirstBlock ActionRequestType = 5
	// ActionRequestTypeWithBlock defines that the // Action with list and next block.
	ActionRequestTypeWithBlock ActionRequestType = 6
)

// ActionRequestTypeParse converts the given string into a ActionRequestType value.
//
// It returns the corresponding ActionRequestType constant if the string matches
// a known level name, or an error if the input is invalid.
func ActionRequestTypeParse(value string) (ActionRequestType, error) {
	var ret ActionRequestType
	var err error
	switch {
	case strings.EqualFold(value, "Normal"):
		ret = ActionRequestTypeNormal
	case strings.EqualFold(value, "NextBlock"):
		ret = ActionRequestTypeNextBlock
	case strings.EqualFold(value, "WithList"):
		ret = ActionRequestTypeWithList
	case strings.EqualFold(value, "WithFirstBlock"):
		ret = ActionRequestTypeWithFirstBlock
	case strings.EqualFold(value, "WithListAndFirstBlock"):
		ret = ActionRequestTypeWithListAndFirstBlock
	case strings.EqualFold(value, "WithBlock"):
		ret = ActionRequestTypeWithBlock
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the ActionRequestType.
// It satisfies fmt.Stringer.
func (g ActionRequestType) String() string {
	var ret string
	switch g {
	case ActionRequestTypeNormal:
		ret = "Normal"
	case ActionRequestTypeNextBlock:
		ret = "NextBlock"
	case ActionRequestTypeWithList:
		ret = "WithList"
	case ActionRequestTypeWithFirstBlock:
		ret = "WithFirstBlock"
	case ActionRequestTypeWithListAndFirstBlock:
		ret = "WithListAndFirstBlock"
	case ActionRequestTypeWithBlock:
		ret = "WithBlock"
	}
	return ret
}

// AllActionRequestType returns a slice containing all defined ActionRequestType values.
func AllActionRequestType() []ActionRequestType {
	return []ActionRequestType{
		ActionRequestTypeNormal,
		ActionRequestTypeNextBlock,
		ActionRequestTypeWithList,
		ActionRequestTypeWithFirstBlock,
		ActionRequestTypeWithListAndFirstBlock,
		ActionRequestTypeWithBlock,
	}
}
