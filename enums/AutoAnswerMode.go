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

type AutoAnswerMode int

const (
	// AutoAnswerModeDevice defines that the line dedicated to the device.
	AutoAnswerModeDevice AutoAnswerMode = iota
	// AutoAnswerModeCall defines that the shared line management with a limited number of calls allowed. Once the number of calls is reached,
	//  the window status becomes inactive until the next start date, whatever the result of the call,
	AutoAnswerModeCall
	// AutoAnswerModeConnected defines that the shared line management with a limited number of successful calls allowed. Once the number of
	//  successful communications is reached, the window status becomes inactive until the next start date,
	AutoAnswerModeConnected
	// AutoAnswerModeNone defines that the currently no modem connected.
	AutoAnswerModeNone
)

// AutoAnswerModeParse converts the given string into a AutoAnswerMode value.
//
// It returns the corresponding AutoAnswerMode constant if the string matches
// a known level name, or an error if the input is invalid.
func AutoAnswerModeParse(value string) (AutoAnswerMode, error) {
	var ret AutoAnswerMode
	var err error
	switch {
	case strings.EqualFold(value, "Device"):
		ret = AutoAnswerModeDevice
	case strings.EqualFold(value, "Call"):
		ret = AutoAnswerModeCall
	case strings.EqualFold(value, "Connected"):
		ret = AutoAnswerModeConnected
	case strings.EqualFold(value, "None"):
		ret = AutoAnswerModeNone
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the AutoAnswerMode.
// It satisfies fmt.Stringer.
func (g AutoAnswerMode) String() string {
	var ret string
	switch g {
	case AutoAnswerModeDevice:
		ret = "Device"
	case AutoAnswerModeCall:
		ret = "Call"
	case AutoAnswerModeConnected:
		ret = "Connected"
	case AutoAnswerModeNone:
		ret = "None"
	}
	return ret
}

// AllAutoAnswerMode returns a slice containing all defined AutoAnswerMode values.
func AllAutoAnswerMode() []AutoAnswerMode {
	return []AutoAnswerMode{
	AutoAnswerModeDevice,
	AutoAnswerModeCall,
	AutoAnswerModeConnected,
	AutoAnswerModeNone,
	}
}
