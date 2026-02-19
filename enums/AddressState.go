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

// AddressState Defines whether or not the device has been assigned an address
//
//	since last power up of the device.
type AddressState int

const (
	// AddressStateNone defines that the not assigned an address yet.
	AddressStateNone AddressState = iota
	// AddressStateAssigned defines that the assigned an address either by manual setting, or by automated method.
	AddressStateAssigned
)

// AddressStateParse converts the given string into a AddressState value.
//
// It returns the corresponding AddressState constant if the string matches
// a known level name, or an error if the input is invalid.
func AddressStateParse(value string) (AddressState, error) {
	var ret AddressState
	var err error
	switch {
	case strings.EqualFold(value, "None"):
		ret = AddressStateNone
	case strings.EqualFold(value, "Assigned"):
		ret = AddressStateAssigned
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the AddressState.
// It satisfies fmt.Stringer.
func (g AddressState) String() string {
	var ret string
	switch g {
	case AddressStateNone:
		ret = "None"
	case AddressStateAssigned:
		ret = "Assigned"
	}
	return ret
}

// AllAddressState returns a slice containing all defined AddressState values.
func AllAddressState() []AddressState {
	return []AddressState{
	AddressStateNone,
	AddressStateAssigned,
	}
}
