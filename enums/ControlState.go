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

// ControlState The internal states of the disconnect control object.
type ControlState int

const (
	// ControlStateDisconnected defines that the the output_state is set to false and the consumer is disconnected.
	ControlStateDisconnected ControlState = iota
	// ControlStateConnected defines that the the output_state is set to true and the consumer is connected.
	ControlStateConnected
	// ControlStateReadyForReconnection defines that the the output_state is set to false and the consumer is disconnected.
	ControlStateReadyForReconnection
)

// ControlStateParse converts the given string into a ControlState value.
//
// It returns the corresponding ControlState constant if the string matches
// a known level name, or an error if the input is invalid.
func ControlStateParse(value string) (ControlState, error) {
	var ret ControlState
	var err error
	switch strings.ToUpper(value) {
	case "DISCONNECTED":
		ret = ControlStateDisconnected
	case "CONNECTED":
		ret = ControlStateConnected
	case "READYFORRECONNECTION":
		ret = ControlStateReadyForReconnection
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the ControlState.
// It satisfies fmt.Stringer.
func (g ControlState) String() string {
	var ret string
	switch g {
	case ControlStateDisconnected:
		ret = "DISCONNECTED"
	case ControlStateConnected:
		ret = "CONNECTED"
	case ControlStateReadyForReconnection:
		ret = "READYFORRECONNECTION"
	}
	return ret
}
