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

// ControlMode Configures the behaviour of the disconnect control object for all
//
//	triggers, i.e. the possible state transitions.
type ControlMode int

const (
	// ControlModeNone defines that the the disconnect control object is always in 'connected' state,
	ControlModeNone ControlMode = iota
	// ControlModeMode1 defines that the disconnection: Remote (b, c), manual (f), local (g)
	//  Reconnection: Remote (d), manual (e).
	ControlModeMode1
	// ControlModeMode2 defines that the disconnection: Remote (b, c), manual (f), local (g)
	//  Reconnection: Remote (a), manual (e).
	ControlModeMode2
	// ControlModeMode3 defines that the disconnection: Remote (b, c), manual (-), local (g)
	//  Reconnection: Remote (d), manual (e).
	ControlModeMode3
	// ControlModeMode4 defines that the disconnection: Remote (b, c), manual (-), local (g)
	//  Reconnection: Remote (a), manual (e)
	ControlModeMode4
	// ControlModeMode5 defines that the disconnection: Remote (b, c), manual (f), local (g)
	//  Reconnection: Remote (d), manual (e), local (h),
	ControlModeMode5
	// ControlModeMode6 defines that the disconnection: Remote (b, c), manual (-), local (g)
	// Reconnection: Remote (d), manual (e), local (h)
	ControlModeMode6
	// ControlModeMode7 defines that the disconnection: Remote (b, c), manual (-), local (g)
	// Reconnection: Remote (a, i), manual (e), local (h)
	ControlModeMode7
)

// ControlModeParse converts the given string into a ControlMode value.
//
// It returns the corresponding ControlMode constant if the string matches
// a known level name, or an error if the input is invalid.
func ControlModeParse(value string) (ControlMode, error) {
	var ret ControlMode
	var err error
	switch strings.ToUpper(value) {
	case "NONE":
		ret = ControlModeNone
	case "MODE1":
		ret = ControlModeMode1
	case "MODE2":
		ret = ControlModeMode2
	case "MODE3":
		ret = ControlModeMode3
	case "MODE4":
		ret = ControlModeMode4
	case "MODE5":
		ret = ControlModeMode5
	case "MODE6":
		ret = ControlModeMode6
	case "MODE7":
		ret = ControlModeMode7
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the ControlMode.
// It satisfies fmt.Stringer.
func (g ControlMode) String() string {
	var ret string
	switch g {
	case ControlModeNone:
		ret = "NONE"
	case ControlModeMode1:
		ret = "MODE1"
	case ControlModeMode2:
		ret = "MODE2"
	case ControlModeMode3:
		ret = "MODE3"
	case ControlModeMode4:
		ret = "MODE4"
	case ControlModeMode5:
		ret = "MODE5"
	case ControlModeMode6:
		ret = "MODE6"
	case ControlModeMode7:
		ret = "MODE7"
	}
	return ret
}
