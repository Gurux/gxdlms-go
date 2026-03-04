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

type DeltaElectricalPhase int

const (
	// DeltaElectricalPhaseNotDefined defines that the // Not defined.
	DeltaElectricalPhaseNotDefined DeltaElectricalPhase = iota
	// DeltaElectricalPhaseSame defines that the // The server system is connected to the same phase as the client system.
	DeltaElectricalPhaseSame
	DeltaElectricalPhaseDegrees60
	DeltaElectricalPhaseDegrees120
	DeltaElectricalPhaseDegrees180
	DeltaElectricalPhaseDegreesMinus120
	DeltaElectricalPhaseDegreesMinus60
)

// DeltaElectricalPhaseParse converts the given string into a DeltaElectricalPhase value.
//
// It returns the corresponding DeltaElectricalPhase constant if the string matches
// a known level name, or an error if the input is invalid.
func DeltaElectricalPhaseParse(value string) (DeltaElectricalPhase, error) {
	var ret DeltaElectricalPhase
	var err error
	switch {
	case strings.EqualFold(value, "NotDefined"):
		ret = DeltaElectricalPhaseNotDefined
	case strings.EqualFold(value, "Same"):
		ret = DeltaElectricalPhaseSame
	case strings.EqualFold(value, "Degrees60"):
		ret = DeltaElectricalPhaseDegrees60
	case strings.EqualFold(value, "Degrees120"):
		ret = DeltaElectricalPhaseDegrees120
	case strings.EqualFold(value, "Degrees180"):
		ret = DeltaElectricalPhaseDegrees180
	case strings.EqualFold(value, "DegreesMinus120"):
		ret = DeltaElectricalPhaseDegreesMinus120
	case strings.EqualFold(value, "DegreesMinus60"):
		ret = DeltaElectricalPhaseDegreesMinus60
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the DeltaElectricalPhase.
// It satisfies fmt.Stringer.
func (g DeltaElectricalPhase) String() string {
	var ret string
	switch g {
	case DeltaElectricalPhaseNotDefined:
		ret = "NotDefined"
	case DeltaElectricalPhaseSame:
		ret = "Same"
	case DeltaElectricalPhaseDegrees60:
		ret = "Degrees60"
	case DeltaElectricalPhaseDegrees120:
		ret = "Degrees120"
	case DeltaElectricalPhaseDegrees180:
		ret = "Degrees180"
	case DeltaElectricalPhaseDegreesMinus120:
		ret = "DegreesMinus120"
	case DeltaElectricalPhaseDegreesMinus60:
		ret = "DegreesMinus60"
	}
	return ret
}

// AllDeltaElectricalPhase returns a slice containing all defined DeltaElectricalPhase values.
func AllDeltaElectricalPhase() []DeltaElectricalPhase {
	return []DeltaElectricalPhase{
	DeltaElectricalPhaseNotDefined,
	DeltaElectricalPhaseSame,
	DeltaElectricalPhaseDegrees60,
	DeltaElectricalPhaseDegrees120,
	DeltaElectricalPhaseDegrees180,
	DeltaElectricalPhaseDegreesMinus120,
	DeltaElectricalPhaseDegreesMinus60,
	}
}
