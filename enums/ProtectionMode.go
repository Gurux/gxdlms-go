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

// ProtectionMode Enumerates communication port protection mode values.
type ProtectionMode int

const (
	// ProtectionModeLocked defines that the port is locked. Communication is not possible.
	ProtectionModeLocked ProtectionMode = iota
	// ProtectionModeLockedOnFailedAttempts defines that the the port becomes temporarily locked when failed connections exceeds an allowed.
	ProtectionModeLockedOnFailedAttempts
	// ProtectionModeUnlocked defines that the port is unlocked.
	ProtectionModeUnlocked
)

// ProtectionModeParse converts the given string into a ProtectionMode value.
//
// It returns the corresponding ProtectionMode constant if the string matches
// a known level name, or an error if the input is invalid.
func ProtectionModeParse(value string) (ProtectionMode, error) {
	var ret ProtectionMode
	var err error
	switch {
	case strings.EqualFold(value, "Locked"):
		ret = ProtectionModeLocked
	case strings.EqualFold(value, "LockedOnFailedAttempts"):
		ret = ProtectionModeLockedOnFailedAttempts
	case strings.EqualFold(value, "Unlocked"):
		ret = ProtectionModeUnlocked
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the ProtectionMode.
// It satisfies fmt.Stringer.
func (g ProtectionMode) String() string {
	var ret string
	switch g {
	case ProtectionModeLocked:
		ret = "Locked"
	case ProtectionModeLockedOnFailedAttempts:
		ret = "LockedOnFailedAttempts"
	case ProtectionModeUnlocked:
		ret = "Unlocked"
	}
	return ret
}

// AllProtectionMode returns a slice containing all defined ProtectionMode values.
func AllProtectionMode() []ProtectionMode {
	return []ProtectionMode{
		ProtectionModeLocked,
		ProtectionModeLockedOnFailedAttempts,
		ProtectionModeUnlocked,
	}
}
