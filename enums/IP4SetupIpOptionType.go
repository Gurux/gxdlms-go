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

type IP4SetupIpOptionType int

const (
	// IP4SetupIpOptionTypeSecurity defines that the if this option is present, the device shall be allowed to send security,
	//  compartmentation, handling restrictions and TCC (closed user group)
	//  parameters within its IP Datagrams. The value of the IP-Option-
	//  Length Field must be 11, and the IP-Option-Data shall contain the
	//  value of the Security, Compartments, Handling Restrictions and
	//  Transmission Control Code values, as specified in STD0005 / RFC791.
	IP4SetupIpOptionTypeSecurity IP4SetupIpOptionType = 0x82
	// IP4SetupIpOptionTypeLooseSourceAndRecordRoute defines that the if this option is present, the device shall supply routing information to be
	//  used by the gateways in forwarding the datagram to the destination, and to
	//  record the route information.
	//  The IP-Option-length and IP-Option-Data values are specified in STD0005 / RFC 791.
	IP4SetupIpOptionTypeLooseSourceAndRecordRoute IP4SetupIpOptionType = 0x83
	// IP4SetupIpOptionTypeStrictSourceAndRecordRoute defines that the if this option is present, the device shall supply routing information to be
	//  used by the gateways in forwarding the datagram to the destination, and to
	//  record the route information.
	//  The IP-Option-length and IP-Option-Data values are specified in STD0005 / RFC 791.
	IP4SetupIpOptionTypeStrictSourceAndRecordRoute IP4SetupIpOptionType = 0x89
	// IP4SetupIpOptionTypeRecordRoute defines that the if this option is present, the device shall as well:
	//  send originated IP Datagrams with that option, providing means
	//  to record the route of these Datagrams
	//  as a router, send routed IP Datagrams with the route option
	//  adjusted according to this option.
	//  The IP-Option-length and IP-Option-Data values are specified in
	//  STD0005 / RFC 791.
	IP4SetupIpOptionTypeRecordRoute IP4SetupIpOptionType = 0x07
	// IP4SetupIpOptionTypeInternetTimestamp defines that the if this option is present, the device shall as well:
	//  send originated IP Datagrams with that option, providing means
	//  to time-stamp the datagram in the route to its destination
	//  as a router, send routed IP Datagrams with the time-stamp option
	//  adjusted according to this option.
	//  The IP-Option-length and IP-Option-Data values are specified in STD0005 / RFC 791.
	IP4SetupIpOptionTypeInternetTimestamp IP4SetupIpOptionType = 0x44
)

// IP4SetupIpOptionTypeParse converts the given string into a IP4SetupIpOptionType value.
//
// It returns the corresponding IP4SetupIpOptionType constant if the string matches
// a known level name, or an error if the input is invalid.
func IP4SetupIpOptionTypeParse(value string) (IP4SetupIpOptionType, error) {
	var ret IP4SetupIpOptionType
	var err error
	switch {
	case strings.EqualFold(value, "Security"):
		ret = IP4SetupIpOptionTypeSecurity
	case strings.EqualFold(value, "LooseSourceAndRecordRoute"):
		ret = IP4SetupIpOptionTypeLooseSourceAndRecordRoute
	case strings.EqualFold(value, "StrictSourceAndRecordRoute"):
		ret = IP4SetupIpOptionTypeStrictSourceAndRecordRoute
	case strings.EqualFold(value, "RecordRoute"):
		ret = IP4SetupIpOptionTypeRecordRoute
	case strings.EqualFold(value, "InternetTimestamp"):
		ret = IP4SetupIpOptionTypeInternetTimestamp
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the IP4SetupIpOptionType.
// It satisfies fmt.Stringer.
func (g IP4SetupIpOptionType) String() string {
	var ret string
	switch g {
	case IP4SetupIpOptionTypeSecurity:
		ret = "Security"
	case IP4SetupIpOptionTypeLooseSourceAndRecordRoute:
		ret = "LooseSourceAndRecordRoute"
	case IP4SetupIpOptionTypeStrictSourceAndRecordRoute:
		ret = "StrictSourceAndRecordRoute"
	case IP4SetupIpOptionTypeRecordRoute:
		ret = "RecordRoute"
	case IP4SetupIpOptionTypeInternetTimestamp:
		ret = "InternetTimestamp"
	}
	return ret
}

// AllIP4SetupIpOptionType returns a slice containing all defined IP4SetupIpOptionType values.
func AllIP4SetupIpOptionType() []IP4SetupIpOptionType {
	return []IP4SetupIpOptionType{
		IP4SetupIpOptionTypeSecurity,
		IP4SetupIpOptionTypeLooseSourceAndRecordRoute,
		IP4SetupIpOptionTypeStrictSourceAndRecordRoute,
		IP4SetupIpOptionTypeRecordRoute,
		IP4SetupIpOptionTypeInternetTimestamp,
	}
}
