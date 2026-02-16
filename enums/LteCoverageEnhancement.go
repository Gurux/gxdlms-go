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

// LteCoverageEnhancement .
type LteCoverageEnhancement int

const (
	// LteCoverageEnhancementLevel0 defines that the cE Mode A in LTE Cat M1 and CE Level 0 in NB-Iot.
	LteCoverageEnhancementLevel0 LteCoverageEnhancement = iota
	// LteCoverageEnhancementLevel1 defines that the cE Mode B in LTE Cat M1 and CE Level 1 in NB-Iot.
	LteCoverageEnhancementLevel1
	// LteCoverageEnhancementLevel2 defines that the cE Level 2 in NB-Iot.
	LteCoverageEnhancementLevel2
)

// LteCoverageEnhancementParse converts the given string into a LteCoverageEnhancement value.
//
// It returns the corresponding LteCoverageEnhancement constant if the string matches
// a known level name, or an error if the input is invalid.
func LteCoverageEnhancementParse(value string) (LteCoverageEnhancement, error) {
	var ret LteCoverageEnhancement
	var err error
	switch strings.ToUpper(value) {
	case "LEVEL0":
		ret = LteCoverageEnhancementLevel0
	case "LEVEL1":
		ret = LteCoverageEnhancementLevel1
	case "LEVEL2":
		ret = LteCoverageEnhancementLevel2
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the LteCoverageEnhancement.
// It satisfies fmt.Stringer.
func (g LteCoverageEnhancement) String() string {
	var ret string
	switch g {
	case LteCoverageEnhancementLevel0:
		ret = "LEVEL0"
	case LteCoverageEnhancementLevel1:
		ret = "LEVEL1"
	case LteCoverageEnhancementLevel2:
		ret = "LEVEL2"
	}
	return ret
}
