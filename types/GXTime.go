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

// This class is used because in COSEM object model some fields from date time can be ignored.
//
//	Default behavior of DateTime do not allow this.
type GXTime struct {
	GXDateTime
}

// Constructor.
func NewGXTimeFromTime(value time.Time) *GXTime {
	ret := &GXTime{}
	ret.Value = value
	ret.Skip = enums.DateTimeSkipsYear | enums.DateTimeSkipsMonth | enums.DateTimeSkipsDay | enums.DateTimeSkipsDayOfWeek
	return ret
}

// Constructor.
func NewGXTimeFromDateTime(value *GXDateTime) *GXTime {
	ret := &GXTime{}
	ret.Value = value.Value
	ret.Skip = value.Skip | enums.DateTimeSkipsYear | enums.DateTimeSkipsMonth | enums.DateTimeSkipsDay | enums.DateTimeSkipsDayOfWeek
	ret.Extra = value.Extra
	return ret
}

// Constructor.
func NewGXTime(hour int, minute int, second int, millisecond int) (*GXTime, error) {
	ret := &GXTime{}
	ret.Value = time.Date(0, 1, 1, hour, minute, second, millisecond*1000000, time.UTC)
	ret.Skip |= enums.DateTimeSkipsYear | enums.DateTimeSkipsMonth | enums.DateTimeSkipsDay | enums.DateTimeSkipsDayOfWeek
	return ret, nil
}

// Constructor.
func NewGXTimeFromString(value string, language *language.Tag) (*GXTime, error) {
	ret := &GXTime{}
	ret.Skip |= enums.DateTimeSkipsYear | enums.DateTimeSkipsMonth | enums.DateTimeSkipsDay | enums.DateTimeSkipsDayOfWeek
	err := ret.parseInternal(value, language)
	if err != nil {
		return nil, err
	}
	return ret, err
}
