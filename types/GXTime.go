package types

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
	"time"

	"github.com/Gurux/gxdlms-go/enums"
	"golang.org/x/text/language"
)

// GXTime represents a COSEM time-only value (date fields are skipped).
//
// It embeds GXDateTime and configures the Skip flags so that year/month/day and
// day-of-week components are omitted when formatting or serializing.
type GXTime struct {
	GXDateTime
}

// NewGXTimeFromTime creates a GXTime from a Go time.Time.
//
// The returned value represents only the time of day; date fields are skipped.
func NewGXTimeFromTime(value time.Time) *GXTime {
	ret := &GXTime{}
	ret.Value = value
	ret.Skip = enums.DateTimeSkipsYear | enums.DateTimeSkipsMonth | enums.DateTimeSkipsDay | enums.DateTimeSkipsDayOfWeek
	return ret
}

// NewGXTimeFromDateTime creates a GXTime from GXDateTime.
func NewGXTimeFromDateTime(value *GXDateTime) *GXTime {
	ret := &GXTime{}
	ret.Value = value.Value
	ret.Skip = value.Skip | enums.DateTimeSkipsYear | enums.DateTimeSkipsMonth | enums.DateTimeSkipsDay | enums.DateTimeSkipsDayOfWeek
	ret.Extra = value.Extra
	return ret
}

// NewGXTime creates a GXTime from hour, minute, second, and millisecond.
func NewGXTime(hour int, minute int, second int, millisecond int) (*GXTime, error) {
	ret := &GXTime{}
	ret.Value = time.Date(0, 1, 1, hour, minute, second, millisecond*1000000, time.UTC)
	ret.Skip |= enums.DateTimeSkipsYear | enums.DateTimeSkipsMonth | enums.DateTimeSkipsDay | enums.DateTimeSkipsDayOfWeek
	return ret, nil
}

// NewGXTimeFromString parses a time string and returns a GXTime instance.
//
// The returned value will omit date-related fields (year/month/day/day-of-week) and
// uses the provided language tag to determine time formatting rules.
func NewGXTimeFromString(value string, language *language.Tag) (*GXTime, error) {
	ret := &GXTime{}
	ret.Skip |= enums.DateTimeSkipsYear | enums.DateTimeSkipsMonth | enums.DateTimeSkipsDay | enums.DateTimeSkipsDayOfWeek
	err := parseInternal(ret, value, language)
	if err != nil {
		return nil, err
	}
	return ret, err
}

// String implements fmt.Stringer and returns a localized string representation of the time.
//
// It formats the value in the current locale and uses the local timezone. Date fields are
// ignored.
func (g *GXTime) String() string {
	return g.ToString(nil, true)
}

// ToString returns a localized string representation of the time.
//
// Date-related fields are skipped (rendered as '*'). The output format is selected based on
// the provided language tag (or the current locale if nil).
//
// If useLocalTime is true, the value is converted to the local timezone before formatting.
func (g *GXTime) ToString(language *language.Tag, useLocalTime bool) string {
	g.Skip |= enums.DateTimeSkipsYear | enums.DateTimeSkipsMonth | enums.DateTimeSkipsDay | enums.DateTimeSkipsDayOfWeek
	return toString(g, language, useLocalTime, false)
}

// ToFormatString returns the time using a fixed, deterministic output format.
//
// Date-related fields are skipped, and the returned format is not localized beyond
// language-specific separators.
//
// If useLocalTime is true, the value is converted to the local timezone before formatting.
func (g *GXTime) ToFormatString(language *language.Tag, useLocalTime bool) string {
	return toString(g, language, useLocalTime, true)
}

// ToFormatMeterString returns the time using meter timezone formatting.
//
// This produces a fixed format string that includes the meter timezone offset. The date
// components are still skipped.
func (g *GXTime) ToFormatMeterString(language *language.Tag) string {
	return toString(g, language, false, true)
}
