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

// GainResolution Enumerates gain resolution steps.
type GainResolution int

const (
	// GainResolutiondB6 defines that the step is 6 dB.
	GainResolutiondB6 GainResolution = iota
	// GainResolutiondB3 defines that the step is 3 dB.
	GainResolutiondB3
)

// GainResolutionParse converts the given string into a GainResolution value.
//
// It returns the corresponding GainResolution constant if the string matches
// a known level name, or an error if the input is invalid.
func GainResolutionParse(value string) (GainResolution, error) {
	var ret GainResolution
	var err error
	switch strings.ToUpper(value) {
	case "DB6":
		ret = GainResolutiondB6
	case "DB3":
		ret = GainResolutiondB3
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the GainResolution.
// It satisfies fmt.Stringer.
func (g GainResolution) String() string {
	var ret string
	switch g {
	case GainResolutiondB6:
		ret = "DB6"
	case GainResolutiondB3:
		ret = "DB3"
	}
	return ret
}
