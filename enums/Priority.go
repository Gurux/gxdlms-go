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

// Priority defines used priority.
type Priority int

const (
	// PriorityNormal defines that the normal priority is used.
	PriorityNormal Priority = iota
	// PriorityHigh defines that the high priority is used.
	PriorityHigh
)

// PriorityParse converts the given string into a Priority value.
//
// It returns the corresponding Priority constant if the string matches
// a known level name, or an error if the input is invalid.
func PriorityParse(value string) (Priority, error) {
	var ret Priority
	var err error
	switch {
	case strings.EqualFold(value, "Normal"):
		ret = PriorityNormal
	case strings.EqualFold(value, "High"):
		ret = PriorityHigh
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the Priority.
// It satisfies fmt.Stringer.
func (g Priority) String() string {
	var ret string
	switch g {
	case PriorityNormal:
		ret = "Normal"
	case PriorityHigh:
		ret = "High"
	}
	return ret
}

// AllPriority returns a slice containing all defined Priority values.
func AllPriority() []Priority {
	return []Priority{
		PriorityNormal,
		PriorityHigh,
	}
}
