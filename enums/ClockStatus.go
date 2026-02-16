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

// ClockStatus Defines : .
type ClockStatus int

const (
	// ClockStatusOk defines that none status bits is set.
	ClockStatusOk ClockStatus = iota
	// ClockStatusInvalidValue defines that the invalid a value is set.
	ClockStatusInvalidValue = 0x1
	// ClockStatusDoubtfulValue defines that the doubtful value bit is set
	ClockStatusDoubtfulValue = 0x2
	// ClockStatusDifferentClockBase defines that the different clock base bit is set.
	ClockStatusDifferentClockBase = 0x4
	// ClockStatusInvalidClockStatus defines that the invalid clock status is set.
	ClockStatusInvalidClockStatus = 0x8
	// ClockStatusDaylightSavingActive defines that the daylight saving active bit it set.
	ClockStatusDaylightSavingActive = 0x80
	// ClockStatusSkip defines that the clock status is skipped.
	ClockStatusSkip = 0xFF
)

// ClockStatusParse converts the given string into a ClockStatus value.
//
// It returns the corresponding ClockStatus constant if the string matches
// a known level name, or an error if the input is invalid.
func ClockStatusParse(value string) (ClockStatus, error) {
	var ret ClockStatus
	var err error
	switch strings.ToUpper(value) {
	case "OK":
		ret = ClockStatusOk
	case "INVALIDVALUE":
		ret = ClockStatusInvalidValue
	case "DOUBTFULVALUE":
		ret = ClockStatusDoubtfulValue
	case "DIFFERENTCLOCKBASE":
		ret = ClockStatusDifferentClockBase
	case "INVALIDCLOCKSTATUS":
		ret = ClockStatusInvalidClockStatus
	case "DAYLIGHTSAVINGACTIVE":
		ret = ClockStatusDaylightSavingActive
	case "SKIP":
		ret = ClockStatusSkip
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the ClockStatus.
// It satisfies fmt.Stringer.
func (g ClockStatus) String() string {
	var ret string
	switch g {
	case ClockStatusOk:
		ret = "OK"
	case ClockStatusInvalidValue:
		ret = "INVALIDVALUE"
	case ClockStatusDoubtfulValue:
		ret = "DOUBTFULVALUE"
	case ClockStatusDifferentClockBase:
		ret = "DIFFERENTCLOCKBASE"
	case ClockStatusInvalidClockStatus:
		ret = "INVALIDCLOCKSTATUS"
	case ClockStatusDaylightSavingActive:
		ret = "DAYLIGHTSAVINGACTIVE"
	case ClockStatusSkip:
		ret = "SKIP"
	}
	return ret
}
