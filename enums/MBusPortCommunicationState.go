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

// MBusPortCommunicationState Defines the M-Bus port communication status.
type MBusPortCommunicationState int

const (
	// MBusPortCommunicationStateNoAccess defines that the no access.
	MBusPortCommunicationStateNoAccess MBusPortCommunicationState = iota
	// MBusPortCommunicationStateTemporaryNoAccess defines that the temporary no access
	MBusPortCommunicationStateTemporaryNoAccess
	// MBusPortCommunicationStateLimitedAccess defines that the limited access
	MBusPortCommunicationStateLimitedAccess
	// MBusPortCommunicationStateUnlimitedAccess defines that the unlimited access.
	MBusPortCommunicationStateUnlimitedAccess
	// MBusPortCommunicationStatewMBus defines that the wM-Bus.
	MBusPortCommunicationStatewMBus
)

// MBusPortCommunicationStateParse converts the given string into a MBusPortCommunicationState value.
//
// It returns the corresponding MBusPortCommunicationState constant if the string matches
// a known level name, or an error if the input is invalid.
func MBusPortCommunicationStateParse(value string) (MBusPortCommunicationState, error) {
	var ret MBusPortCommunicationState
	var err error
	switch {
	case strings.EqualFold(value, "NoAccess"):
		ret = MBusPortCommunicationStateNoAccess
	case strings.EqualFold(value, "TemporaryNoAccess"):
		ret = MBusPortCommunicationStateTemporaryNoAccess
	case strings.EqualFold(value, "LimitedAccess"):
		ret = MBusPortCommunicationStateLimitedAccess
	case strings.EqualFold(value, "UnlimitedAccess"):
		ret = MBusPortCommunicationStateUnlimitedAccess
	case strings.EqualFold(value, "wMBus"):
		ret = MBusPortCommunicationStatewMBus
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the MBusPortCommunicationState.
// It satisfies fmt.Stringer.
func (g MBusPortCommunicationState) String() string {
	var ret string
	switch g {
	case MBusPortCommunicationStateNoAccess:
		ret = "NoAccess"
	case MBusPortCommunicationStateTemporaryNoAccess:
		ret = "TemporaryNoAccess"
	case MBusPortCommunicationStateLimitedAccess:
		ret = "LimitedAccess"
	case MBusPortCommunicationStateUnlimitedAccess:
		ret = "UnlimitedAccess"
	case MBusPortCommunicationStatewMBus:
		ret = "wMBus"
	}
	return ret
}

// AllMBusPortCommunicationState returns a slice containing all defined MBusPortCommunicationState values.
func AllMBusPortCommunicationState() []MBusPortCommunicationState {
	return []MBusPortCommunicationState{
		MBusPortCommunicationStateNoAccess,
		MBusPortCommunicationStateTemporaryNoAccess,
		MBusPortCommunicationStateLimitedAccess,
		MBusPortCommunicationStateUnlimitedAccess,
		MBusPortCommunicationStatewMBus,
	}
}
