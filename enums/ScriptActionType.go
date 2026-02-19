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

type ScriptActionType int

const (
	// ScriptActionTypeNone defines that the nothing is going to execute.
	ScriptActionTypeNone ScriptActionType = iota
	// ScriptActionTypeWrite defines that the write attribute.
	ScriptActionTypeWrite
	// ScriptActionTypeExecute defines that the execute specific method
	ScriptActionTypeExecute
)

// ScriptActionTypeParse converts the given string into a ScriptActionType value.
//
// It returns the corresponding ScriptActionType constant if the string matches
// a known level name, or an error if the input is invalid.
func ScriptActionTypeParse(value string) (ScriptActionType, error) {
	var ret ScriptActionType
	var err error
	switch {
	case strings.EqualFold(value, "None"):
		ret = ScriptActionTypeNone
	case strings.EqualFold(value, "Write"):
		ret = ScriptActionTypeWrite
	case strings.EqualFold(value, "Execute"):
		ret = ScriptActionTypeExecute
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the ScriptActionType.
// It satisfies fmt.Stringer.
func (g ScriptActionType) String() string {
	var ret string
	switch g {
	case ScriptActionTypeNone:
		ret = "None"
	case ScriptActionTypeWrite:
		ret = "Write"
	case ScriptActionTypeExecute:
		ret = "Execute"
	}
	return ret
}

// AllScriptActionType returns a slice containing all defined ScriptActionType values.
func AllScriptActionType() []ScriptActionType {
	return []ScriptActionType{
		ScriptActionTypeNone,
		ScriptActionTypeWrite,
		ScriptActionTypeExecute,
	}
}
