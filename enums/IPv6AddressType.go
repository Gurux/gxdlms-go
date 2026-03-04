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

// IPv6AddressType enumerates possible address type to add.
type IPv6AddressType int

const (
	// IPv6AddressTypeUnicast defines that the unicast address is used.
	IPv6AddressTypeUnicast IPv6AddressType = iota
	// IPv6AddressTypeMulticast defines that the multicast address is used.
	IPv6AddressTypeMulticast
	// IPv6AddressTypeGateway defines that the Gateway address is used.
	IPv6AddressTypeGateway
)

// IPv6AddressTypeParse converts the given string into a IPv6AddressType value.
//
// It returns the corresponding IPv6AddressType constant if the string matches
// a known level name, or an error if the input is invalid.
func IPv6AddressTypeParse(value string) (IPv6AddressType, error) {
	var ret IPv6AddressType
	var err error
	switch {
	case strings.EqualFold(value, "Unicast"):
		ret = IPv6AddressTypeUnicast
	case strings.EqualFold(value, "Multicast"):
		ret = IPv6AddressTypeMulticast
	case strings.EqualFold(value, "Gateway"):
		ret = IPv6AddressTypeGateway
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the IPv6AddressType.
// It satisfies fmt.Stringer.
func (g IPv6AddressType) String() string {
	var ret string
	switch g {
	case IPv6AddressTypeUnicast:
		ret = "Unicast"
	case IPv6AddressTypeMulticast:
		ret = "Multicast"
	case IPv6AddressTypeGateway:
		ret = "Gateway"
	}
	return ret
}

// AllIPv6AddressType returns a slice containing all defined IPv6AddressType values.
func AllIPv6AddressType() []IPv6AddressType {
	return []IPv6AddressType{
		IPv6AddressTypeUnicast,
		IPv6AddressTypeMulticast,
		IPv6AddressTypeGateway,
	}
}
