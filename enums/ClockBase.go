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

type ClockBase int

const (
	// ClockBaseNone defines that the not defined
	ClockBaseNone ClockBase = iota
	// ClockBaseCrystal defines that the internal Crystal
	ClockBaseCrystal
	// ClockBaseFrequency50 defines that the mains frequency 50 Hz,
	ClockBaseFrequency50
	// ClockBaseFrequency60 defines that the mains frequency 60 Hz,
	ClockBaseFrequency60
	// ClockBaseGPS defines that the global Positioning System.
	ClockBaseGPS
	// ClockBaseRadio defines that the radio controlled.
	ClockBaseRadio
)

// ClockBaseParse converts the given string into a ClockBase value.
//
// It returns the corresponding ClockBase constant if the string matches
// a known level name, or an error if the input is invalid.
func ClockBaseParse(value string) (ClockBase, error) {
	var ret ClockBase
	var err error
	switch strings.ToUpper(value) {
	case "NONE":
		ret = ClockBaseNone
	case "CRYSTAL":
		ret = ClockBaseCrystal
	case "FREQUENCY50":
		ret = ClockBaseFrequency50
	case "FREQUENCY60":
		ret = ClockBaseFrequency60
	case "GPS":
		ret = ClockBaseGPS
	case "RADIO":
		ret = ClockBaseRadio
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the ClockBase.
// It satisfies fmt.Stringer.
func (g ClockBase) String() string {
	var ret string
	switch g {
	case ClockBaseNone:
		ret = "NONE"
	case ClockBaseCrystal:
		ret = "CRYSTAL"
	case ClockBaseFrequency50:
		ret = "FREQUENCY50"
	case ClockBaseFrequency60:
		ret = "FREQUENCY60"
	case ClockBaseGPS:
		ret = "GPS"
	case ClockBaseRadio:
		ret = "RADIO"
	}
	return ret
}
