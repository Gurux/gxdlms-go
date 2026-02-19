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

// DeviceType Defines the type of the device connected to the modem.
type DeviceType int

const (
	// DeviceTypePanDevice defines that the pAN device.
	DeviceTypePanDevice DeviceType = iota
	// DeviceTypePanCoordinator defines that the pAN coordinator.
	DeviceTypePanCoordinator
	// DeviceTypeNotDefined defines that the not Defined.
	DeviceTypeNotDefined
)

// DeviceTypeParse converts the given string into a DeviceType value.
//
// It returns the corresponding DeviceType constant if the string matches
// a known level name, or an error if the input is invalid.
func DeviceTypeParse(value string) (DeviceType, error) {
	var ret DeviceType
	var err error
	switch {
	case strings.EqualFold(value, "PanDevice"):
		ret = DeviceTypePanDevice
	case strings.EqualFold(value, "PanCoordinator"):
		ret = DeviceTypePanCoordinator
	case strings.EqualFold(value, "NotDefined"):
		ret = DeviceTypeNotDefined
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the DeviceType.
// It satisfies fmt.Stringer.
func (g DeviceType) String() string {
	var ret string
	switch g {
	case DeviceTypePanDevice:
		ret = "PanDevice"
	case DeviceTypePanCoordinator:
		ret = "PanCoordinator"
	case DeviceTypeNotDefined:
		ret = "NotDefined"
	}
	return ret
}

// AllDeviceType returns a slice containing all defined DeviceType values.
func AllDeviceType() []DeviceType {
	return []DeviceType{
	DeviceTypePanDevice,
	DeviceTypePanCoordinator,
	DeviceTypeNotDefined,
	}
}
