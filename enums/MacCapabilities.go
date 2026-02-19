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

// MacCapabilities Present functional state of the node.
type MacCapabilities int

const (
	// MacCapabilitiesSwitchCapable defines that switch capable bit is set.
	MacCapabilitiesSwitchCapable MacCapabilities = 1
	// MacCapabilitiesPacketAggregation defines that packet aggregation bit is set.
	MacCapabilitiesPacketAggregation MacCapabilities = 2
	// MacCapabilitiesContentionFreePeriod defines that contention free period bit is set.
	MacCapabilitiesContentionFreePeriod MacCapabilities = 4
	// MacCapabilitiesDirectConnection defines that direct connection bit is set.
	MacCapabilitiesDirectConnection MacCapabilities = 8
	// MacCapabilitiesMulticast defines that multicast bit is set.
	MacCapabilitiesMulticast MacCapabilities = 0x10
	// MacCapabilitiesPhyRobustnessManagement defines that pHY Robustness Management bit is set.
	MacCapabilitiesPhyRobustnessManagement MacCapabilities = 0x20
	// MacCapabilitiesArq defines that aRQ bit is set.
	MacCapabilitiesArq MacCapabilities = 0x40
	// MacCapabilitiesReservedForFutureUse defines that reserved for future use bit is set.
	MacCapabilitiesReservedForFutureUse MacCapabilities = 0x80
	// MacCapabilitiesDirectConnectionSwitching defines that direct Connection Switching bit is set.
	MacCapabilitiesDirectConnectionSwitching MacCapabilities = 0x100
	// MacCapabilitiesMulticastSwitchingCapability defines that multicast Switching Capability bit is set.
	MacCapabilitiesMulticastSwitchingCapability MacCapabilities = 0x200
	// MacCapabilitiesPhyRobustnessManagementSwitchingCapability defines that pHY Robustness Management Switching Capability bit is set.
	MacCapabilitiesPhyRobustnessManagementSwitchingCapability MacCapabilities = 0x400
	// MacCapabilitiesArqBufferingSwitchingCapability defines that aRQ Buffering Switching Capability bit is set.
	MacCapabilitiesArqBufferingSwitchingCapability MacCapabilities = 0x800
)

// MacCapabilitiesParse converts the given string into a MacCapabilities value.
//
// It returns the corresponding MacCapabilities constant if the string matches
// a known level name, or an error if the input is invalid.
func MacCapabilitiesParse(value string) (MacCapabilities, error) {
	var ret MacCapabilities
	var err error
	switch {
	case strings.EqualFold(value, "SwitchCapable"):
		ret = MacCapabilitiesSwitchCapable
	case strings.EqualFold(value, "PacketAggregation"):
		ret = MacCapabilitiesPacketAggregation
	case strings.EqualFold(value, "ContentionFreePeriod"):
		ret = MacCapabilitiesContentionFreePeriod
	case strings.EqualFold(value, "DirectConnection"):
		ret = MacCapabilitiesDirectConnection
	case strings.EqualFold(value, "Multicast"):
		ret = MacCapabilitiesMulticast
	case strings.EqualFold(value, "PhyRobustnessManagement"):
		ret = MacCapabilitiesPhyRobustnessManagement
	case strings.EqualFold(value, "Arq"):
		ret = MacCapabilitiesArq
	case strings.EqualFold(value, "ReservedForFutureUse"):
		ret = MacCapabilitiesReservedForFutureUse
	case strings.EqualFold(value, "DirectConnectionSwitching"):
		ret = MacCapabilitiesDirectConnectionSwitching
	case strings.EqualFold(value, "MulticastSwitchingCapability"):
		ret = MacCapabilitiesMulticastSwitchingCapability
	case strings.EqualFold(value, "PhyRobustnessManagementSwitchingCapability"):
		ret = MacCapabilitiesPhyRobustnessManagementSwitchingCapability
	case strings.EqualFold(value, "ArqBufferingSwitchingCapability"):
		ret = MacCapabilitiesArqBufferingSwitchingCapability
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the MacCapabilities.
// It satisfies fmt.Stringer.
func (g MacCapabilities) String() string {
	var ret string
	switch g {
	case MacCapabilitiesSwitchCapable:
		ret = "SwitchCapable"
	case MacCapabilitiesPacketAggregation:
		ret = "PacketAggregation"
	case MacCapabilitiesContentionFreePeriod:
		ret = "ContentionFreePeriod"
	case MacCapabilitiesDirectConnection:
		ret = "DirectConnection"
	case MacCapabilitiesMulticast:
		ret = "Multicast"
	case MacCapabilitiesPhyRobustnessManagement:
		ret = "PhyRobustnessManagement"
	case MacCapabilitiesArq:
		ret = "Arq"
	case MacCapabilitiesReservedForFutureUse:
		ret = "ReservedForFutureUse"
	case MacCapabilitiesDirectConnectionSwitching:
		ret = "DirectConnectionSwitching"
	case MacCapabilitiesMulticastSwitchingCapability:
		ret = "MulticastSwitchingCapability"
	case MacCapabilitiesPhyRobustnessManagementSwitchingCapability:
		ret = "PhyRobustnessManagementSwitchingCapability"
	case MacCapabilitiesArqBufferingSwitchingCapability:
		ret = "ArqBufferingSwitchingCapability"
	}
	return ret
}

// AllMacCapabilities returns a slice containing all defined MacCapabilities values.
func AllMacCapabilities() []MacCapabilities {
	return []MacCapabilities{
		MacCapabilitiesSwitchCapable,
		MacCapabilitiesPacketAggregation,
		MacCapabilitiesContentionFreePeriod,
		MacCapabilitiesDirectConnection,
		MacCapabilitiesMulticast,
		MacCapabilitiesPhyRobustnessManagement,
		MacCapabilitiesArq,
		MacCapabilitiesReservedForFutureUse,
		MacCapabilitiesDirectConnectionSwitching,
		MacCapabilitiesMulticastSwitchingCapability,
		MacCapabilitiesPhyRobustnessManagementSwitchingCapability,
		MacCapabilitiesArqBufferingSwitchingCapability,
	}
}
