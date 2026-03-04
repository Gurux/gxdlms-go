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

// MIB variable initiator electrical phase.
type InitiatorElectricalPhase int

const (
	// InitiatorElectricalPhaseNotDefined defines that the // Not defined.
	InitiatorElectricalPhaseNotDefined InitiatorElectricalPhase = iota
	// InitiatorElectricalPhasePhase1 defines that the // Phase 1.
	InitiatorElectricalPhasePhase1
	// InitiatorElectricalPhasePhase2 defines that the // Phase 2.
	InitiatorElectricalPhasePhase2
	// InitiatorElectricalPhasePhase3 defines that the // Phase 3.
	InitiatorElectricalPhasePhase3
)

// InitiatorElectricalPhaseParse converts the given string into a InitiatorElectricalPhase value.
//
// It returns the corresponding InitiatorElectricalPhase constant if the string matches
// a known level name, or an error if the input is invalid.
func InitiatorElectricalPhaseParse(value string) (InitiatorElectricalPhase, error) {
	var ret InitiatorElectricalPhase
	var err error
	switch {
	case strings.EqualFold(value, "NotDefined"):
		ret = InitiatorElectricalPhaseNotDefined
	case strings.EqualFold(value, "Phase1"):
		ret = InitiatorElectricalPhasePhase1
	case strings.EqualFold(value, "Phase2"):
		ret = InitiatorElectricalPhasePhase2
	case strings.EqualFold(value, "Phase3"):
		ret = InitiatorElectricalPhasePhase3
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the InitiatorElectricalPhase.
// It satisfies fmt.Stringer.
func (g InitiatorElectricalPhase) String() string {
	var ret string
	switch g {
	case InitiatorElectricalPhaseNotDefined:
		ret = "NotDefined"
	case InitiatorElectricalPhasePhase1:
		ret = "Phase1"
	case InitiatorElectricalPhasePhase2:
		ret = "Phase2"
	case InitiatorElectricalPhasePhase3:
		ret = "Phase3"
	}
	return ret
}

// AllInitiatorElectricalPhase returns a slice containing all defined InitiatorElectricalPhase values.
func AllInitiatorElectricalPhase() []InitiatorElectricalPhase {
	return []InitiatorElectricalPhase{
	InitiatorElectricalPhaseNotDefined,
	InitiatorElectricalPhasePhase1,
	InitiatorElectricalPhasePhase2,
	InitiatorElectricalPhasePhase3,
	}
}
