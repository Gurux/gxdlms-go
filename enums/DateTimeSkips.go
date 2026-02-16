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

// DateTimeSkips Enumerates skipped fields from date time.
type DateTimeSkips int

const (
	// DateTimeSkipsNone defines that nothing is skipped from date time.
	DateTimeSkipsNone DateTimeSkips = 0x0
	// DateTimeSkipsYear defines that year part of date time is skipped bit is set.
	DateTimeSkipsYear DateTimeSkips = 0x1
	// DateTimeSkipsMonth defines that month part of date time is skipped bit is set.
	DateTimeSkipsMonth DateTimeSkips = 0x2
	// DateTimeSkipsDay defines that day part is skipped bit is set.
	DateTimeSkipsDay DateTimeSkips = 0x4
	// DateTimeSkipsDayOfWeek defines that day of week part of date time is skipped bit is set.
	DateTimeSkipsDayOfWeek DateTimeSkips = 0x8
	// DateTimeSkipsHour defines that hours part of date time is skipped bit is set.
	DateTimeSkipsHour DateTimeSkips = 0x10
	// DateTimeSkipsMinute defines that minute part of date time is skipped bit is set.
	DateTimeSkipsMinute DateTimeSkips = 0x20
	// DateTimeSkipsSecond defines that seconds part of date time is skipped bit is set.
	DateTimeSkipsSecond DateTimeSkips = 0x40
	// DateTimeSkipsMs defines that hundreds of seconds part of date time is skipped bit is set.
	DateTimeSkipsMs DateTimeSkips = 0x80
	// DateTimeSkipsDeviation defines that deviation is skipped on write bit is set.
	DateTimeSkipsDeviation DateTimeSkips = 0x100
	// DateTimeSkipsStatus defines that status is skipped on write bit is set.
	DateTimeSkipsStatus DateTimeSkips = 0x200
)

// DateTimeSkipsParse converts the given string into a DateTimeSkips value.
//
// It returns the corresponding DateTimeSkips constant if the string matches
// a known level name, or an error if the input is invalid.
func DateTimeSkipsParse(value string) (DateTimeSkips, error) {
	var ret DateTimeSkips
	var err error
	switch strings.ToUpper(value) {
	case "NONE":
		ret = DateTimeSkipsNone
	case "YEAR":
		ret = DateTimeSkipsYear
	case "MONTH":
		ret = DateTimeSkipsMonth
	case "DAY":
		ret = DateTimeSkipsDay
	case "DAYOFWEEK":
		ret = DateTimeSkipsDayOfWeek
	case "HOUR":
		ret = DateTimeSkipsHour
	case "MINUTE":
		ret = DateTimeSkipsMinute
	case "SECOND":
		ret = DateTimeSkipsSecond
	case "MS":
		ret = DateTimeSkipsMs
	case "DEVIATION":
		ret = DateTimeSkipsDeviation
	case "STATUS":
		ret = DateTimeSkipsStatus
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the DateTimeSkips.
// It satisfies fmt.Stringer.
func (g DateTimeSkips) String() string {
	var ret string
	switch g {
	case DateTimeSkipsNone:
		ret = "NONE"
	case DateTimeSkipsYear:
		ret = "YEAR"
	case DateTimeSkipsMonth:
		ret = "MONTH"
	case DateTimeSkipsDay:
		ret = "DAY"
	case DateTimeSkipsDayOfWeek:
		ret = "DAYOFWEEK"
	case DateTimeSkipsHour:
		ret = "HOUR"
	case DateTimeSkipsMinute:
		ret = "MINUTE"
	case DateTimeSkipsSecond:
		ret = "SECOND"
	case DateTimeSkipsMs:
		ret = "MS"
	case DateTimeSkipsDeviation:
		ret = "DEVIATION"
	case DateTimeSkipsStatus:
		ret = "STATUS"
	}
	return ret
}
