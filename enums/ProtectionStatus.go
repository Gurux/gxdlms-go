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

// ProtectionStatus Enumerates communication port protection status values.
type ProtectionStatus int

const (
	// ProtectionStatusUnlocked defines that the port is unlocked.
	ProtectionStatusUnlocked ProtectionStatus = iota
	// ProtectionStatusTemporarilyLocked defines that the the port is temporarily locked. Communication is not possible.
	ProtectionStatusTemporarilyLocked
	// ProtectionStatusLocked defines that the port is locked. Communication is not possible.
	ProtectionStatusLocked
)

// ProtectionStatusParse converts the given string into a ProtectionStatus value.
//
// It returns the corresponding ProtectionStatus constant if the string matches
// a known level name, or an error if the input is invalid.
func ProtectionStatusParse(value string) (ProtectionStatus, error) {
	var ret ProtectionStatus
	var err error
	switch {
	case strings.EqualFold(value, "Unlocked"):
		ret = ProtectionStatusUnlocked
	case strings.EqualFold(value, "TemporarilyLocked"):
		ret = ProtectionStatusTemporarilyLocked
	case strings.EqualFold(value, "Locked"):
		ret = ProtectionStatusLocked
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the ProtectionStatus.
// It satisfies fmt.Stringer.
func (g ProtectionStatus) String() string {
	var ret string
	switch g {
	case ProtectionStatusUnlocked:
		ret = "Unlocked"
	case ProtectionStatusTemporarilyLocked:
		ret = "TemporarilyLocked"
	case ProtectionStatusLocked:
		ret = "Locked"
	}
	return ret
}

// AllProtectionStatus returns a slice containing all defined ProtectionStatus values.
func AllProtectionStatus() []ProtectionStatus {
	return []ProtectionStatus{
		ProtectionStatusUnlocked,
		ProtectionStatusTemporarilyLocked,
		ProtectionStatusLocked,
	}
}
