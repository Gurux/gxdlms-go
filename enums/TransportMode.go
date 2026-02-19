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

// TransportMode CoAP transport modes.
type TransportMode int

const (
	// TransportModeReliable defines that the reliable operation supported only.
	TransportModeReliable TransportMode = 1
	// TransportModeUnreliable defines that the unreliable operation supported only,
	TransportModeUnreliable TransportMode = 2
	// TransportModeReliableUnreliable defines that the reliable and Unreliable operation supported.
	TransportModeReliableUnreliable TransportMode = 3
)

// TransportModeParse converts the given string into a TransportMode value.
//
// It returns the corresponding TransportMode constant if the string matches
// a known level name, or an error if the input is invalid.
func TransportModeParse(value string) (TransportMode, error) {
	var ret TransportMode
	var err error
	switch {
	case strings.EqualFold(value, "Reliable"):
		ret = TransportModeReliable
	case strings.EqualFold(value, "Unreliable"):
		ret = TransportModeUnreliable
	case strings.EqualFold(value, "ReliableUnreliable"):
		ret = TransportModeReliableUnreliable
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the TransportMode.
// It satisfies fmt.Stringer.
func (g TransportMode) String() string {
	var ret string
	switch g {
	case TransportModeReliable:
		ret = "Reliable"
	case TransportModeUnreliable:
		ret = "Unreliable"
	case TransportModeReliableUnreliable:
		ret = "ReliableUnreliable"
	}
	return ret
}

// AllTransportMode returns a slice containing all defined TransportMode values.
func AllTransportMode() []TransportMode {
	return []TransportMode{
		TransportModeReliable,
		TransportModeUnreliable,
		TransportModeReliableUnreliable,
	}
}
