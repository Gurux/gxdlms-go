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

// Weekdays defines the weekdays.
type Weekdays int

const (
	// WeekdaysNone defines that no day of week is selected bit is set.
	WeekdaysNone Weekdays = iota
	// WeekdaysMonday defines that indicates Monday bit is set.
	WeekdaysMonday Weekdays = 0x1
	// WeekdaysTuesday defines that indicates Tuesday bit is set.
	WeekdaysTuesday Weekdays = 0x2
	// WeekdaysWednesday defines that indicates Wednesday bit is set.
	WeekdaysWednesday Weekdays = 0x4
	// WeekdaysThursday defines that indicates Thursday bit is set.
	WeekdaysThursday Weekdays = 0x8
	// WeekdaysFriday defines that indicates Friday bit is set.
	WeekdaysFriday Weekdays = 0x10
	// WeekdaysSaturday defines that indicates Saturday bit is set.
	WeekdaysSaturday Weekdays = 0x20
	// WeekdaysSunday defines that indicates Sunday bit is set.
	WeekdaysSunday Weekdays = 0x40
)

// WeekdaysParse converts the given string into a Weekdays value.
//
// It returns the corresponding Weekdays constant if the string matches
// a known level name, or an error if the input is invalid.
func WeekdaysParse(value string) (Weekdays, error) {
	var ret Weekdays
	var err error
	switch {
	case strings.EqualFold(value, "None"):
		ret = WeekdaysNone
	case strings.EqualFold(value, "Monday"):
		ret = WeekdaysMonday
	case strings.EqualFold(value, "Tuesday"):
		ret = WeekdaysTuesday
	case strings.EqualFold(value, "Wednesday"):
		ret = WeekdaysWednesday
	case strings.EqualFold(value, "Thursday"):
		ret = WeekdaysThursday
	case strings.EqualFold(value, "Friday"):
		ret = WeekdaysFriday
	case strings.EqualFold(value, "Saturday"):
		ret = WeekdaysSaturday
	case strings.EqualFold(value, "Sunday"):
		ret = WeekdaysSunday
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the Weekdays.
// It satisfies fmt.Stringer.
func (g Weekdays) String() string {
	var ret string
	switch g {
	case WeekdaysNone:
		ret = "None"
	case WeekdaysMonday:
		ret = "Monday"
	case WeekdaysTuesday:
		ret = "Tuesday"
	case WeekdaysWednesday:
		ret = "Wednesday"
	case WeekdaysThursday:
		ret = "Thursday"
	case WeekdaysFriday:
		ret = "Friday"
	case WeekdaysSaturday:
		ret = "Saturday"
	case WeekdaysSunday:
		ret = "Sunday"
	}
	return ret
}

// AllWeekdays returns a slice containing all defined Weekdays values.
func AllWeekdays() []Weekdays {
	return []Weekdays{
		WeekdaysNone,
		WeekdaysMonday,
		WeekdaysTuesday,
		WeekdaysWednesday,
		WeekdaysThursday,
		WeekdaysFriday,
		WeekdaysSaturday,
		WeekdaysSunday,
	}
}
