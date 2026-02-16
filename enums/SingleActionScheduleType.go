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

type SingleActionScheduleType int

const (
	// SingleActionScheduleTypeSingleActionScheduleType1 defines that the size of execution_time = 1. Wildcard in date allowed.
	SingleActionScheduleTypeSingleActionScheduleType1 SingleActionScheduleType = 1
	// SingleActionScheduleTypeSingleActionScheduleType2 defines that the size of execution_time = n.
	//  All time values are the same, wildcards in date not allowed.
	SingleActionScheduleTypeSingleActionScheduleType2 SingleActionScheduleType = 2
	// SingleActionScheduleTypeSingleActionScheduleType3 defines that the size of execution_time = n.
	//  All time values are the same, wildcards in date are allowed,
	SingleActionScheduleTypeSingleActionScheduleType3 SingleActionScheduleType = 3
	// SingleActionScheduleTypeSingleActionScheduleType4 defines that the size of execution_time = n.
	//  Time values may be different, wildcards in date not allowed,
	SingleActionScheduleTypeSingleActionScheduleType4 SingleActionScheduleType = 4
	// SingleActionScheduleTypeSingleActionScheduleType5 defines that the size of execution_time = n.
	//  Time values may be different, wildcards in date are allowed
	SingleActionScheduleTypeSingleActionScheduleType5 SingleActionScheduleType = 5
)

// SingleActionScheduleTypeParse converts the given string into a SingleActionScheduleType value.
//
// It returns the corresponding SingleActionScheduleType constant if the string matches
// a known level name, or an error if the input is invalid.
func SingleActionScheduleTypeParse(value string) (SingleActionScheduleType, error) {
	var ret SingleActionScheduleType
	var err error
	switch strings.ToUpper(value) {
	case "SINGLEACTIONSCHEDULETYPE1":
		ret = SingleActionScheduleTypeSingleActionScheduleType1
	case "SINGLEACTIONSCHEDULETYPE2":
		ret = SingleActionScheduleTypeSingleActionScheduleType2
	case "SINGLEACTIONSCHEDULETYPE3":
		ret = SingleActionScheduleTypeSingleActionScheduleType3
	case "SINGLEACTIONSCHEDULETYPE4":
		ret = SingleActionScheduleTypeSingleActionScheduleType4
	case "SINGLEACTIONSCHEDULETYPE5":
		ret = SingleActionScheduleTypeSingleActionScheduleType5
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the SingleActionScheduleType.
// It satisfies fmt.Stringer.
func (g SingleActionScheduleType) String() string {
	var ret string
	switch g {
	case SingleActionScheduleTypeSingleActionScheduleType1:
		ret = "SINGLEACTIONSCHEDULETYPE1"
	case SingleActionScheduleTypeSingleActionScheduleType2:
		ret = "SINGLEACTIONSCHEDULETYPE2"
	case SingleActionScheduleTypeSingleActionScheduleType3:
		ret = "SINGLEACTIONSCHEDULETYPE3"
	case SingleActionScheduleTypeSingleActionScheduleType4:
		ret = "SINGLEACTIONSCHEDULETYPE4"
	case SingleActionScheduleTypeSingleActionScheduleType5:
		ret = "SINGLEACTIONSCHEDULETYPE5"
	}
	return ret
}
