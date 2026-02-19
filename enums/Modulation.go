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

// Modulation Enumerates modulation types.
type Modulation int

const (
	// ModulationRobustMode defines that the robust Mode.
	ModulationRobustMode Modulation = iota
	// ModulationDBPsk defines that the DBPSK modulation is used.
	ModulationDBPsk
	// ModulationDqPsk defines that the dQPSK modulation is used.
	ModulationDqPsk
	// ModulationD8Psk defines that the d8PSK modulation is used.
	ModulationD8Psk
	// ModulationQam16 defines that the 16-QAM modulation is used.
	ModulationQam16
)

// ModulationParse converts the given string into a Modulation value.
//
// It returns the corresponding Modulation constant if the string matches
// a known level name, or an error if the input is invalid.
func ModulationParse(value string) (Modulation, error) {
	var ret Modulation
	var err error
	switch {
	case strings.EqualFold(value, "RobustMode"):
		ret = ModulationRobustMode
	case strings.EqualFold(value, "DbPsk"):
		ret = ModulationDBPsk
	case strings.EqualFold(value, "DqPsk"):
		ret = ModulationDqPsk
	case strings.EqualFold(value, "D8Psk"):
		ret = ModulationD8Psk
	case strings.EqualFold(value, "Qam16"):
		ret = ModulationQam16
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the Modulation.
// It satisfies fmt.Stringer.
func (g Modulation) String() string {
	var ret string
	switch g {
	case ModulationRobustMode:
		ret = "RobustMode"
	case ModulationDBPsk:
		ret = "DbPsk"
	case ModulationDqPsk:
		ret = "DqPsk"
	case ModulationD8Psk:
		ret = "D8Psk"
	case ModulationQam16:
		ret = "Qam16"
	}
	return ret
}

// AllModulation returns a slice containing all defined Modulation values.
func AllModulation() []Modulation {
	return []Modulation{
		ModulationRobustMode,
		ModulationDBPsk,
		ModulationDqPsk,
		ModulationD8Psk,
		ModulationQam16,
	}
}
