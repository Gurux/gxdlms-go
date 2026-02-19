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

// ZigBeeStatus Defines the ZigBee status enumeration values.
type ZigBeeStatus int

const (
	// ZigBeeStatusAuthorised defines that authorised on PAN bit is set.
	ZigBeeStatusAuthorised ZigBeeStatus = 0x1
	// ZigBeeStatusReporting defines that actively reporting on PAN bit is set.
	ZigBeeStatusReporting ZigBeeStatus = 0x2
	// ZigBeeStatusUnauthorised defines that unauthorised on PAN but has reported bit is set.
	ZigBeeStatusUnauthorised ZigBeeStatus = 0x4
	// ZigBeeStatusAuthorisedSwapOut defines that authorised after swap-out bit is set.
	ZigBeeStatusAuthorisedSwapOut ZigBeeStatus = 0x8
	// ZigBeeStatusSepTransmitting defines that sEP Transmitting bit is set.
	ZigBeeStatusSepTransmitting ZigBeeStatus = 0x10
)

// ZigBeeStatusParse converts the given string into a ZigBeeStatus value.
//
// It returns the corresponding ZigBeeStatus constant if the string matches
// a known level name, or an error if the input is invalid.
func ZigBeeStatusParse(value string) (ZigBeeStatus, error) {
	var ret ZigBeeStatus
	var err error
	switch {
	case strings.EqualFold(value, "Authorised"):
		ret = ZigBeeStatusAuthorised
	case strings.EqualFold(value, "Reporting"):
		ret = ZigBeeStatusReporting
	case strings.EqualFold(value, "Unauthorised"):
		ret = ZigBeeStatusUnauthorised
	case strings.EqualFold(value, "AuthorisedSwapOut"):
		ret = ZigBeeStatusAuthorisedSwapOut
	case strings.EqualFold(value, "SepTransmitting"):
		ret = ZigBeeStatusSepTransmitting
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the ZigBeeStatus.
// It satisfies fmt.Stringer.
func (g ZigBeeStatus) String() string {
	var ret string
	switch g {
	case ZigBeeStatusAuthorised:
		ret = "Authorised"
	case ZigBeeStatusReporting:
		ret = "Reporting"
	case ZigBeeStatusUnauthorised:
		ret = "Unauthorised"
	case ZigBeeStatusAuthorisedSwapOut:
		ret = "AuthorisedSwapOut"
	case ZigBeeStatusSepTransmitting:
		ret = "SepTransmitting"
	}
	return ret
}

// AllZigBeeStatus returns a slice containing all defined ZigBeeStatus values.
func AllZigBeeStatus() []ZigBeeStatus {
	return []ZigBeeStatus{
		ZigBeeStatusAuthorised,
		ZigBeeStatusReporting,
		ZigBeeStatusUnauthorised,
		ZigBeeStatusAuthorisedSwapOut,
		ZigBeeStatusSepTransmitting,
	}
}
