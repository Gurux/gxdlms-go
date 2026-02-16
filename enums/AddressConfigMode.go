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

// AddressConfigMode Enumerated : s.
type AddressConfigMode int

const (
	// AddressConfigModeAuto defines that the auto Configuration.
	AddressConfigModeAuto AddressConfigMode = iota
	// AddressConfigModeDHCPv6 defines that the dHCP v6.
	AddressConfigModeDHCPv6
	// AddressConfigModeManual defines that the manual
	AddressConfigModeManual
	// AddressConfigModeNeighbourDiscovery defines that the neighbour Discovery.
	AddressConfigModeNeighbourDiscovery
)

// AddressConfigModeParse converts the given string into a AddressConfigMode value.
//
// It returns the corresponding AddressConfigMode constant if the string matches
// a known level name, or an error if the input is invalid.
func AddressConfigModeParse(value string) (AddressConfigMode, error) {
	var ret AddressConfigMode
	var err error
	switch strings.ToUpper(value) {
	case "AUTO":
		ret = AddressConfigModeAuto
	case "DHCPV6":
		ret = AddressConfigModeDHCPv6
	case "MANUAL":
		ret = AddressConfigModeManual
	case "NEIGHBOURDISCOVERY":
		ret = AddressConfigModeNeighbourDiscovery
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the AddressConfigMode.
// It satisfies fmt.Stringer.
func (g AddressConfigMode) String() string {
	var ret string
	switch g {
	case AddressConfigModeAuto:
		ret = "AUTO"
	case AddressConfigModeDHCPv6:
		ret = "DHCPV6"
	case AddressConfigModeManual:
		ret = "MANUAL"
	case AddressConfigModeNeighbourDiscovery:
		ret = "NEIGHBOURDISCOVERY"
	}
	return ret
}
