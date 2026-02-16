package constants

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
// More information of Gurux products: https://www.gurux.org
//
// This code is licensed under the GNU General Public License v2.
// Full text may be retrieved at http://www.gnu.org/licenses/gpl-2.0.txt
// ---------------------------------------------------------------------------

import (
	"fmt"
	"strings"

	"github.com/Gurux/gxcommon-go"
)

// MBusMeterType M-Bus meter type.
type MBusMeterType int

const (
	// MBusMeterTypeOil defines that the oil meter.
	MBusMeterTypeOil MBusMeterType = 1
	// MBusMeterTypeEnergy defines that the energy meter.
	MBusMeterTypeEnergy MBusMeterType = 2
	// MBusMeterTypeGas defines that the gas meter.
	MBusMeterTypeGas MBusMeterType = 3
	// MBusMeterTypeWater defines that the water meter.
	MBusMeterTypeWater MBusMeterType = 7
	// MBusMeterTypeUnknown defines that the unknown meter type.
	MBusMeterTypeUnknown MBusMeterType = 0x0F
)

// MBusMeterTypeParse converts the given string into a MBusMeterType value.
//
// It returns the corresponding MBusMeterType constant if the string matches
// a known level name, or an error if the input is invalid.
func MBusMeterTypeParse(value string) (MBusMeterType, error) {
	var ret MBusMeterType
	var err error
	switch strings.ToUpper(value) {
	case "OIL":
		ret = MBusMeterTypeOil
	case "ENERGY":
		ret = MBusMeterTypeEnergy
	case "GAS":
		ret = MBusMeterTypeGas
	case "WATER":
		ret = MBusMeterTypeWater
	case "UNKNOWN":
		ret = MBusMeterTypeUnknown
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the MBusMeterType.
// It satisfies fmt.Stringer.
func (g MBusMeterType) String() string {
	var ret string
	switch g {
	case MBusMeterTypeOil:
		ret = "OIL"
	case MBusMeterTypeEnergy:
		ret = "ENERGY"
	case MBusMeterTypeGas:
		ret = "GAS"
	case MBusMeterTypeWater:
		ret = "WATER"
	case MBusMeterTypeUnknown:
		ret = "UNKNOWN"
	}
	return ret
}
