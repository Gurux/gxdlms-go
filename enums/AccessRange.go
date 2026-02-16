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

type AccessRange int

const (
	// AccessRangeEntry defines that the read entries.
	AccessRangeEntry AccessRange = iota
	// AccessRangeLast defines that the read last N days.
	AccessRangeLast
	// AccessRangeRange defines that the read between days
	AccessRangeRange
	// AccessRangeAll defines that the read All.
	AccessRangeAll
)

// AccessRangeParse converts the given string into a AccessRange value.
//
// It returns the corresponding AccessRange constant if the string matches
// a known level name, or an error if the input is invalid.
func AccessRangeParse(value string) (AccessRange, error) {
	var ret AccessRange
	var err error
	switch strings.ToUpper(value) {
	case "ENTRY":
		ret = AccessRangeEntry
	case "LAST":
		ret = AccessRangeLast
	case "RANGE":
		ret = AccessRangeRange
	case "ALL":
		ret = AccessRangeAll
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the AccessRange.
// It satisfies fmt.Stringer.
func (g AccessRange) String() string {
	var ret string
	switch g {
	case AccessRangeEntry:
		ret = "ENTRY"
	case AccessRangeLast:
		ret = "LAST"
	case AccessRangeRange:
		ret = "RANGE"
	case AccessRangeAll:
		ret = "ALL"
	}
	return ret
}
