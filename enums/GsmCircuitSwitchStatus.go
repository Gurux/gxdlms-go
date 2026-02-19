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

// GSM circuit switced status.
type GsmCircuitSwitchStatus int

const (
	// GsmCircuitSwitchStatusInactive defines that the // Inactive.
	GsmCircuitSwitchStatusInactive GsmCircuitSwitchStatus = iota
	// GsmCircuitSwitchStatusIncomingCall defines that the // Incoming call.
	GsmCircuitSwitchStatusIncomingCall
	// GsmCircuitSwitchStatusActive defines that the // Active.
	GsmCircuitSwitchStatusActive
)

// GsmCircuitSwitchStatusParse converts the given string into a GsmCircuitSwitchStatus value.
//
// It returns the corresponding GsmCircuitSwitchStatus constant if the string matches
// a known level name, or an error if the input is invalid.
func GsmCircuitSwitchStatusParse(value string) (GsmCircuitSwitchStatus, error) {
	var ret GsmCircuitSwitchStatus
	var err error
	switch {
	case strings.EqualFold(value, "Inactive"):
		ret = GsmCircuitSwitchStatusInactive
	case strings.EqualFold(value, "IncomingCall"):
		ret = GsmCircuitSwitchStatusIncomingCall
	case strings.EqualFold(value, "Active"):
		ret = GsmCircuitSwitchStatusActive
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the GsmCircuitSwitchStatus.
// It satisfies fmt.Stringer.
func (g GsmCircuitSwitchStatus) String() string {
	var ret string
	switch g {
	case GsmCircuitSwitchStatusInactive:
		ret = "Inactive"
	case GsmCircuitSwitchStatusIncomingCall:
		ret = "IncomingCall"
	case GsmCircuitSwitchStatusActive:
		ret = "Active"
	}
	return ret
}

// AllGsmCircuitSwitchStatus returns a slice containing all defined GsmCircuitSwitchStatus values.
func AllGsmCircuitSwitchStatus() []GsmCircuitSwitchStatus {
	return []GsmCircuitSwitchStatus{
	GsmCircuitSwitchStatusInactive,
	GsmCircuitSwitchStatusIncomingCall,
	GsmCircuitSwitchStatusActive,
	}
}
