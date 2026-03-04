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

//   enumerator values.
type Repeater int

const (
	// RepeaterNever defines that the // Newer repeater.
	RepeaterNever Repeater = iota
	// RepeaterAlways defines that the // Always repeater.
	RepeaterAlways
	// RepeaterDynamic defines that the // Dynamic repeater.
	RepeaterDynamic
)

// RepeaterParse converts the given string into a Repeater value.
//
// It returns the corresponding Repeater constant if the string matches
// a known level name, or an error if the input is invalid.
func RepeaterParse(value string) (Repeater, error) {
	var ret Repeater
	var err error
	switch {
	case strings.EqualFold(value, "Never"):
		ret = RepeaterNever
	case strings.EqualFold(value, "Always"):
		ret = RepeaterAlways
	case strings.EqualFold(value, "Dynamic"):
		ret = RepeaterDynamic
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the Repeater.
// It satisfies fmt.Stringer.
func (g Repeater) String() string {
	var ret string
	switch g {
	case RepeaterNever:
		ret = "Never"
	case RepeaterAlways:
		ret = "Always"
	case RepeaterDynamic:
		ret = "Dynamic"
	}
	return ret
}

// AllRepeater returns a slice containing all defined Repeater values.
func AllRepeater() []Repeater {
	return []Repeater{
	RepeaterNever,
	RepeaterAlways,
	RepeaterDynamic,
	}
}
