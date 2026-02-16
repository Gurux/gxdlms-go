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

// ConnectionState Enumerates Action request types.
type ConnectionState int

const (
	// ConnectionStateNone defines that connection is not made for the meter bit is set.
	ConnectionStateNone ConnectionState = iota
	// ConnectionStateHdlc defines that connection is made for HDLC level bit is set.
	ConnectionStateHdlc
	// ConnectionStateDlms defines that connection is made for DLMS level bit is set.
	ConnectionStateDlms
	// ConnectionStateIec defines that connection is made for optical IEC 62056-21 level bit is set.
	ConnectionStateIec ConnectionState = 4
)

// ConnectionStateParse converts the given string into a ConnectionState value.
//
// It returns the corresponding ConnectionState constant if the string matches
// a known level name, or an error if the input is invalid.
func ConnectionStateParse(value string) (ConnectionState, error) {
	var ret ConnectionState
	var err error
	switch strings.ToUpper(value) {
	case "NONE":
		ret = ConnectionStateNone
	case "HDLC":
		ret = ConnectionStateHdlc
	case "DLMS":
		ret = ConnectionStateDlms
	case "IEC":
		ret = ConnectionStateIec
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the ConnectionState.
// It satisfies fmt.Stringer.
func (g ConnectionState) String() string {
	var ret string
	switch g {
	case ConnectionStateNone:
		ret = "NONE"
	case ConnectionStateHdlc:
		ret = "HDLC"
	case ConnectionStateDlms:
		ret = "DLMS"
	case ConnectionStateIec:
		ret = "IEC"
	}
	return ret
}
