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

// This class is used because in COSEM object model some fields from date time can be ignored.
//
//	Default behavior of DateTime do not allow this.
type GXDate struct {
	GXDateTime
}

// Constructor.
func NewGXDateFromString(value string, language *language.Tag) (*GXDate, error) {
	ret := &GXDate{}
	ret.Skip = enums.DateTimeSkipsHour | enums.DateTimeSkipsMinute | enums.DateTimeSkipsSecond | enums.DateTimeSkipsDeviation | enums.DateTimeSkipsMs
	err := ret.parseInternal(value, language)
	if err != nil {
		return nil, err
	}
	return ret, err
}

// Constructor.
func NewGXDateFromTime(value time.Time) (*GXDate, error) {
	ret := &GXDate{}
	ret.Value = value
	ret.Skip = enums.DateTimeSkipsHour | enums.DateTimeSkipsMinute | enums.DateTimeSkipsSecond | enums.DateTimeSkipsDeviation | enums.DateTimeSkipsMs
	return ret, nil
}

// Constructor.
func NewGXDateFromDateTime(value *GXDateTime) *GXDate {
	ret := &GXDate{}
	ret.Value = value.Value
	ret.Skip = enums.DateTimeSkipsHour | enums.DateTimeSkipsMinute | enums.DateTimeSkipsSecond | enums.DateTimeSkipsDeviation | enums.DateTimeSkipsMs
	ret.Extra = value.Extra
	return ret
}

// Constructor.
func NewGXDate(year int, month int, day int) (*GXDate, error) {
	ret := &GXDate{}
	ret.Value = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	ret.Skip = enums.DateTimeSkipsHour | enums.DateTimeSkipsMinute | enums.DateTimeSkipsSecond | enums.DateTimeSkipsDeviation | enums.DateTimeSkipsMs
	ret.Skip |= enums.DateTimeSkipsDayOfWeek
	return ret, nil
}

func (g *GXDate) String() string {
	return g.ToString(nil, true)
}

func (g *GXDate) ToString(language *language.Tag, useLocalTime bool) string {
	g.GXDateTime.Skip |= enums.DateTimeSkipsHour | enums.DateTimeSkipsMinute | enums.DateTimeSkipsSecond | enums.DateTimeSkipsDeviation | enums.DateTimeSkipsMs
	return g.GXDateTime.ToString(language, useLocalTime)
}
