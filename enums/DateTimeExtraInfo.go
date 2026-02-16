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

// DateTimeExtraInfo : .
type DateTimeExtraInfo int

const (
	// DateTimeExtraInfoNone defines that there is no extra info.
	DateTimeExtraInfoNone DateTimeExtraInfo = 0x0
	// DateTimeExtraInfoDstBegin defines that daylight savings begin bit is set.
	DateTimeExtraInfoDstBegin DateTimeExtraInfo = 0x1
	// DateTimeExtraInfoDstEnd defines that daylight savings end bit is set.
	DateTimeExtraInfoDstEnd DateTimeExtraInfo = 0x2
	// DateTimeExtraInfoLastDay defines that last day of month bit is set.
	DateTimeExtraInfoLastDay DateTimeExtraInfo = 0x4
	// DateTimeExtraInfoLastDay2 defines that 2nd last day of month bit is set.
	DateTimeExtraInfoLastDay2 DateTimeExtraInfo = 0x8
)

// DateTimeExtraInfoParse converts the given string into a DateTimeExtraInfo value.
//
// It returns the corresponding DateTimeExtraInfo constant if the string matches
// a known level name, or an error if the input is invalid.
func DateTimeExtraInfoParse(value string) (DateTimeExtraInfo, error) {
	var ret DateTimeExtraInfo
	var err error
	switch strings.ToUpper(value) {
	case "NONE":
		ret = DateTimeExtraInfoNone
	case "DSTBEGIN":
		ret = DateTimeExtraInfoDstBegin
	case "DSTEND":
		ret = DateTimeExtraInfoDstEnd
	case "LASTDAY":
		ret = DateTimeExtraInfoLastDay
	case "LASTDAY2":
		ret = DateTimeExtraInfoLastDay2
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the DateTimeExtraInfo.
// It satisfies fmt.Stringer.
func (g DateTimeExtraInfo) String() string {
	var ret string
	switch g {
	case DateTimeExtraInfoNone:
		ret = "NONE"
	case DateTimeExtraInfoDstBegin:
		ret = "DSTBEGIN"
	case DateTimeExtraInfoDstEnd:
		ret = "DSTEND"
	case DateTimeExtraInfoLastDay:
		ret = "LASTDAY"
	case DateTimeExtraInfoLastDay2:
		ret = "LASTDAY2"
	}
	return ret
}
