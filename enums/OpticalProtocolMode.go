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

// OpticalProtocolMode Defines the protocol used by the meter on the port.
type OpticalProtocolMode int

const (
	// OpticalProtocolModeDefault defines that the protocol according to IEC 62056-21 (modes A…E),
	OpticalProtocolModeDefault OpticalProtocolMode = iota
	// OpticalProtocolModeNet defines that the protocol according to IEC 62056-46.
	//  Using this enumeration value all other attributes of this IC are not applicable.
	OpticalProtocolModeNet
	// OpticalProtocolModeUnknown defines that the protocol not specified. Using this enumeration value,
	//  ProposedBaudrate is used for setting the communication speed on the port.
	//  All other attributes are not applicable.
	OpticalProtocolModeUnknown
)

// OpticalProtocolModeParse converts the given string into a OpticalProtocolMode value.
//
// It returns the corresponding OpticalProtocolMode constant if the string matches
// a known level name, or an error if the input is invalid.
func OpticalProtocolModeParse(value string) (OpticalProtocolMode, error) {
	var ret OpticalProtocolMode
	var err error
	switch strings.ToUpper(value) {
	case "DEFAULT":
		ret = OpticalProtocolModeDefault
	case "NET":
		ret = OpticalProtocolModeNet
	case "UNKNOWN":
		ret = OpticalProtocolModeUnknown
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the OpticalProtocolMode.
// It satisfies fmt.Stringer.
func (g OpticalProtocolMode) String() string {
	var ret string
	switch g {
	case OpticalProtocolModeDefault:
		ret = "DEFAULT"
	case OpticalProtocolModeNet:
		ret = "NET"
	case OpticalProtocolModeUnknown:
		ret = "UNKNOWN"
	}
	return ret
}
