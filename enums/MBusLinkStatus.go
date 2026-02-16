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

// MBusLinkStatus M-Bus link status.
type MBusLinkStatus int

const (
	// MBusLinkStatusNone defines that the data never received.
	MBusLinkStatusNone MBusLinkStatus = iota
	// MBusLinkStatusNormal defines that the normal operation.
	MBusLinkStatusNormal
	// MBusLinkStatusTemporarilyInterrupted defines that the link temporarily interrupted.
	MBusLinkStatusTemporarilyInterrupted
	// MBusLinkStatusPermanentlyInterrupted defines that the link permanently interrupted.
	MBusLinkStatusPermanentlyInterrupted
)

// MBusLinkStatusParse converts the given string into a MBusLinkStatus value.
//
// It returns the corresponding MBusLinkStatus constant if the string matches
// a known level name, or an error if the input is invalid.
func MBusLinkStatusParse(value string) (MBusLinkStatus, error) {
	var ret MBusLinkStatus
	var err error
	switch strings.ToUpper(value) {
	case "NONE":
		ret = MBusLinkStatusNone
	case "NORMAL":
		ret = MBusLinkStatusNormal
	case "TEMPORARILYINTERRUPTED":
		ret = MBusLinkStatusTemporarilyInterrupted
	case "PERMANENTLYINTERRUPTED":
		ret = MBusLinkStatusPermanentlyInterrupted
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the MBusLinkStatus.
// It satisfies fmt.Stringer.
func (g MBusLinkStatus) String() string {
	var ret string
	switch g {
	case MBusLinkStatusNone:
		ret = "NONE"
	case MBusLinkStatusNormal:
		ret = "NORMAL"
	case MBusLinkStatusTemporarilyInterrupted:
		ret = "TEMPORARILYINTERRUPTED"
	case MBusLinkStatusPermanentlyInterrupted:
		ret = "PERMANENTLYINTERRUPTED"
	}
	return ret
}
