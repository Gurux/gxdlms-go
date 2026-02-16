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

// GsmStatus GSM status.
type GsmStatus int

const (
	// GsmStatusNone defines that the not registered.
	GsmStatusNone GsmStatus = iota
	// GsmStatusHomeNetwork defines that the registered, home network.
	GsmStatusHomeNetwork
	// GsmStatusSearching defines that the not registered, but MT is currently searching a new operator to register to.
	GsmStatusSearching
	// GsmStatusDenied defines that the registration denied.
	GsmStatusDenied
	// GsmStatusUnknown defines that the unknown.
	GsmStatusUnknown
	// GsmStatusRoaming defines that the registered, roaming.
	GsmStatusRoaming
)

// GsmStatusParse converts the given string into a GsmStatus value.
//
// It returns the corresponding GsmStatus constant if the string matches
// a known level name, or an error if the input is invalid.
func GsmStatusParse(value string) (GsmStatus, error) {
	var ret GsmStatus
	var err error
	switch strings.ToUpper(value) {
	case "NONE":
		ret = GsmStatusNone
	case "HOMENETWORK":
		ret = GsmStatusHomeNetwork
	case "SEARCHING":
		ret = GsmStatusSearching
	case "DENIED":
		ret = GsmStatusDenied
	case "UNKNOWN":
		ret = GsmStatusUnknown
	case "ROAMING":
		ret = GsmStatusRoaming
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the GsmStatus.
// It satisfies fmt.Stringer.
func (g GsmStatus) String() string {
	var ret string
	switch g {
	case GsmStatusNone:
		ret = "NONE"
	case GsmStatusHomeNetwork:
		ret = "HOMENETWORK"
	case GsmStatusSearching:
		ret = "SEARCHING"
	case GsmStatusDenied:
		ret = "DENIED"
	case GsmStatusUnknown:
		ret = "UNKNOWN"
	case GsmStatusRoaming:
		ret = "ROAMING"
	}
	return ret
}
