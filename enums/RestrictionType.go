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

// RestrictionType Enumerates compact data and push object restriction types.
type RestrictionType int

const (
	// RestrictionTypeNone defines that the none.
	RestrictionTypeNone RestrictionType = iota
	// RestrictionTypeDate defines that the restriction by date.
	RestrictionTypeDate
	// RestrictionTypeEntry defines that the restriction by entry.
	RestrictionTypeEntry
)

// RestrictionTypeParse converts the given string into a RestrictionType value.
//
// It returns the corresponding RestrictionType constant if the string matches
// a known level name, or an error if the input is invalid.
func RestrictionTypeParse(value string) (RestrictionType, error) {
	var ret RestrictionType
	var err error
	switch {
	case strings.EqualFold(value, "None"):
		ret = RestrictionTypeNone
	case strings.EqualFold(value, "Date"):
		ret = RestrictionTypeDate
	case strings.EqualFold(value, "Entry"):
		ret = RestrictionTypeEntry
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the RestrictionType.
// It satisfies fmt.Stringer.
func (g RestrictionType) String() string {
	var ret string
	switch g {
	case RestrictionTypeNone:
		ret = "None"
	case RestrictionTypeDate:
		ret = "Date"
	case RestrictionTypeEntry:
		ret = "Entry"
	}
	return ret
}

// AllRestrictionType returns a slice containing all defined RestrictionType values.
func AllRestrictionType() []RestrictionType {
	return []RestrictionType{
		RestrictionTypeNone,
		RestrictionTypeDate,
		RestrictionTypeEntry,
	}
}
