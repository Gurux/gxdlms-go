package types

// --------------------------------------------------------------------------
//
//	Gurux Ltd
//
// Filename:        $HeadURL$
//
// Version:         $Revision$,
//
//	$Date$
//	$Author$
//
// # Copyright (c) Gurux Ltd
//
// ---------------------------------------------------------------------------
//
//	DESCRIPTION
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
// More information of Guruonthx products: https://www.gurux.org
//
// This code is licensed under the GNU General Public License v2.
// Full text may be retrieved at http://www.gnu.org/licenses/gpl-2.0.txt
// ---------------------------------------------------------------------------

import (
	"time"

	"github.com/Gurux/gxdlms-go/enums"
	"golang.org/x/text/language"
)

// GXDate represents a COSEM date value where time fields are skipped.
type GXDate struct {
	GXDateTime
}

// NewGXDateFromString creates a GXDate by parsing the given string.
func NewGXDateFromString(value string, language *language.Tag) (*GXDate, error) {
	ret := &GXDate{}
	ret.Skip = enums.DateTimeSkipsHour | enums.DateTimeSkipsMinute | enums.DateTimeSkipsSecond | enums.DateTimeSkipsDeviation | enums.DateTimeSkipsMs
	err := parseInternal(ret, value, language)
	if err != nil {
		return nil, err
	}
	return ret, err
}

// NewGXDateFromTime creates a GXDate from time.Time.
func NewGXDateFromTime(value time.Time) (*GXDate, error) {
	ret := &GXDate{}
	ret.Value = value
	ret.Skip = enums.DateTimeSkipsHour | enums.DateTimeSkipsMinute | enums.DateTimeSkipsSecond | enums.DateTimeSkipsDeviation | enums.DateTimeSkipsMs
	return ret, nil
}

// NewGXDateFromDateTime creates a GXDate from GXDateTime.
func NewGXDateFromDateTime(value *GXDateTime) *GXDate {
	ret := &GXDate{}
	ret.Value = value.Value
	ret.Skip = enums.DateTimeSkipsHour | enums.DateTimeSkipsMinute | enums.DateTimeSkipsSecond | enums.DateTimeSkipsDeviation | enums.DateTimeSkipsMs
	ret.Extra = value.Extra
	return ret
}

// NewGXDate creates a GXDate from year, month, and day.
func NewGXDate(year int, month int, day int) (*GXDate, error) {
	ret := &GXDate{}
	ret.Value = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	ret.Skip = enums.DateTimeSkipsHour | enums.DateTimeSkipsMinute | enums.DateTimeSkipsSecond | enums.DateTimeSkipsDeviation | enums.DateTimeSkipsMs
	ret.Skip |= enums.DateTimeSkipsDayOfWeek
	return ret, nil
}

// String returns the date as a localized string using local time.
func (g *GXDate) String() string {
	return g.ToString(nil, true)
}

// ToString returns the date formatted for the given language.
// If useLocalTime is true, the value is formatted using the local timezone.
func (g *GXDate) ToString(language *language.Tag, useLocalTime bool) string {
	g.Skip |= enums.DateTimeSkipsHour | enums.DateTimeSkipsMinute | enums.DateTimeSkipsSecond | enums.DateTimeSkipsDeviation | enums.DateTimeSkipsMs
	return toString(g, language, useLocalTime, false)
}

// ToFormatString returns the date using a full fixed output format.
// If useLocalTime is true, the value is formatted using the local timezone.
func (g *GXDate) ToFormatString(language *language.Tag, useLocalTime bool) string {
	return toString(g, language, useLocalTime, true)
}

// ToFormatMeterString returns the date using meter timezone formatting.
func (g *GXDate) ToFormatMeterString(language *language.Tag) string {
	return toString(g, language, false, true)
}
