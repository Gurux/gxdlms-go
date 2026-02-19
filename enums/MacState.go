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

// MacState Present functional state of the node.
type MacState int

const (
	// MacStateDisconnected defines that the disconnected.
	MacStateDisconnected MacState = iota
	// MacStateTerminal defines that the terminal.
	MacStateTerminal
	// MacStateSwitch defines that the switch.
	MacStateSwitch
	// MacStateBase defines that the base.
	MacStateBase
)

// MacStateParse converts the given string into a MacState value.
//
// It returns the corresponding MacState constant if the string matches
// a known level name, or an error if the input is invalid.
func MacStateParse(value string) (MacState, error) {
	var ret MacState
	var err error
	switch {
	case strings.EqualFold(value, "Disconnected"):
		ret = MacStateDisconnected
	case strings.EqualFold(value, "Terminal"):
		ret = MacStateTerminal
	case strings.EqualFold(value, "Switch"):
		ret = MacStateSwitch
	case strings.EqualFold(value, "Base"):
		ret = MacStateBase
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the MacState.
// It satisfies fmt.Stringer.
func (g MacState) String() string {
	var ret string
	switch g {
	case MacStateDisconnected:
		ret = "Disconnected"
	case MacStateTerminal:
		ret = "Terminal"
	case MacStateSwitch:
		ret = "Switch"
	case MacStateBase:
		ret = "Base"
	}
	return ret
}

// AllMacState returns a slice containing all defined MacState values.
func AllMacState() []MacState {
	return []MacState{
		MacStateDisconnected,
		MacStateTerminal,
		MacStateSwitch,
		MacStateBase,
	}
}
