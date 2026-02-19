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

// InterfaceType InterfaceType enumerates the usable types of connection in GuruxDLMS.
type InterfaceType int

const (
	// InterfaceTypeHDLC is a general interface type is used for meters that supports
	//  IEC 62056-46 Data link layer using HDLC protocol.
	InterfaceTypeHDLC InterfaceType = iota
	// InterfaceTypeWRAPPER is a network interface type is used for meters that supports
	//  IEC 62056-47 COSEM transport layers for IPv4 networks.
	InterfaceTypeWRAPPER
	// InterfaceTypePDU is for plain PDU.
	InterfaceTypePDU
	// InterfaceTypeWirelessMBus is EN 13757-4/-5 Wireless M-Bus profile is used.
	InterfaceTypeWirelessMBus
	// InterfaceTypeHdlcWithModeE is IEC 62056-21 E-Mode is used to initialize communication before moving to HDLC protocol.
	InterfaceTypeHdlcWithModeE
	// InterfaceTypePlc is PLC Logical link control (LLC) profile is used with IEC 61334-4-32 connectionless LLC sublayer.
	InterfaceTypePlc
	// InterfaceTypePlcHdlc is PLC Logical link control (LLC) profile is used with HDLC.
	InterfaceTypePlcHdlc
	// InterfaceTypeLPWAN is LowPower Wide Area Networks (LPWAN) profile is used.
	InterfaceTypeLPWAN
	// InterfaceTypeWiSUN is Wi-SUN FAN mesh network is used.
	InterfaceTypeWiSUN
	// InterfaceTypePlcPrime is OFDM PLC PRIME is defined in IEC 62056-8-4.
	InterfaceTypePlcPrime
	// InterfaceTypeWiredMBus is EN 13757-2 wired (twisted pair based) M-Bus scheme is used.
	InterfaceTypeWiredMBus
	// InterfaceTypeSMS is SMS short wrapper scheme is used.
	InterfaceTypeSMS
	// InterfaceTypePrimeDcWrapper is PRIME data concentrator wrapper.
	InterfaceTypePrimeDcWrapper
	// InterfaceTypeCoAP is Constrained Application Protocol (CoAP).
	InterfaceTypeCoAP
)

// InterfaceTypeParse converts the given string into a InterfaceType value.
//
// It returns the corresponding InterfaceType constant if the string matches
// a known level name, or an error if the input is invalid.
func InterfaceTypeParse(value string) (InterfaceType, error) {
	var ret InterfaceType
	var err error
	switch {
	case strings.EqualFold(value, "HDLC"):
		ret = InterfaceTypeHDLC
	case strings.EqualFold(value, "WRAPPER"):
		ret = InterfaceTypeWRAPPER
	case strings.EqualFold(value, "PDU"):
		ret = InterfaceTypePDU
	case strings.EqualFold(value, "WirelessMBus"):
		ret = InterfaceTypeWirelessMBus
	case strings.EqualFold(value, "HdlcWithModeE"):
		ret = InterfaceTypeHdlcWithModeE
	case strings.EqualFold(value, "Plc"):
		ret = InterfaceTypePlc
	case strings.EqualFold(value, "PlcHdlc"):
		ret = InterfaceTypePlcHdlc
	case strings.EqualFold(value, "LPWAN"):
		ret = InterfaceTypeLPWAN
	case strings.EqualFold(value, "WiSUN"):
		ret = InterfaceTypeWiSUN
	case strings.EqualFold(value, "PlcPrime"):
		ret = InterfaceTypePlcPrime
	case strings.EqualFold(value, "WiredMBus"):
		ret = InterfaceTypeWiredMBus
	case strings.EqualFold(value, "SMS"):
		ret = InterfaceTypeSMS
	case strings.EqualFold(value, "PrimeDcWrapper"):
		ret = InterfaceTypePrimeDcWrapper
	case strings.EqualFold(value, "CoAP"):
		ret = InterfaceTypeCoAP
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the InterfaceType.
// It satisfies fmt.Stringer.
func (g InterfaceType) String() string {
	var ret string
	switch g {
	case InterfaceTypeHDLC:
		ret = "HDLC"
	case InterfaceTypeWRAPPER:
		ret = "WRAPPER"
	case InterfaceTypePDU:
		ret = "PDU"
	case InterfaceTypeWirelessMBus:
		ret = "WirelessMBus"
	case InterfaceTypeHdlcWithModeE:
		ret = "HdlcWithModeE"
	case InterfaceTypePlc:
		ret = "Plc"
	case InterfaceTypePlcHdlc:
		ret = "PlcHdlc"
	case InterfaceTypeLPWAN:
		ret = "LPWAN"
	case InterfaceTypeWiSUN:
		ret = "WiSUN"
	case InterfaceTypePlcPrime:
		ret = "PlcPrime"
	case InterfaceTypeWiredMBus:
		ret = "WiredMBus"
	case InterfaceTypeSMS:
		ret = "SMS"
	case InterfaceTypePrimeDcWrapper:
		ret = "PrimeDcWrapper"
	case InterfaceTypeCoAP:
		ret = "CoAP"
	}
	return ret
}

// AllInterfaceType returns a slice containing all defined InterfaceType values.
func AllInterfaceType() []InterfaceType {
	return []InterfaceType{
		InterfaceTypeHDLC,
		InterfaceTypeWRAPPER,
		InterfaceTypePDU,
		InterfaceTypeWirelessMBus,
		InterfaceTypeHdlcWithModeE,
		InterfaceTypePlc,
		InterfaceTypePlcHdlc,
		InterfaceTypeLPWAN,
		InterfaceTypeWiSUN,
		InterfaceTypePlcPrime,
		InterfaceTypeWiredMBus,
		InterfaceTypeSMS,
		InterfaceTypePrimeDcWrapper,
		InterfaceTypeCoAP,
	}
}
