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

// ServiceType Type of service used to push the data.
type ServiceType int

const (
	// ServiceTypeTCP defines that the transport service type is TCP/IP.
	ServiceTypeTCP ServiceType = iota
	// ServiceTypeUDP defines that the transport service type is UDP.
	ServiceTypeUDP
	// ServiceTypeFTP defines that the transport service type is FTP.
	ServiceTypeFTP
	// ServiceTypeSMTP defines that the transport service type is SMTP.
	ServiceTypeSMTP
	// ServiceTypeSMS defines that the transport service type is SMS.
	ServiceTypeSMS
	// ServiceTypeHDLC defines that the transport service type is HDLC.
	ServiceTypeHDLC
	// ServiceTypeMBUS defines that the transport service type is M-Bus.
	ServiceTypeMBUS
	// ServiceTypeZigBee defines that the transport service type is ZigBee.
	ServiceTypeZigBee
	// ServiceTypeDlmsGateway defines that the dLMS Gateway.
	ServiceTypeDlmsGateway
	// ServiceTypeReliableCoAP defines that the reliable CoAP.
	ServiceTypeReliableCoAP
	// ServiceTypeUnreliableCoAP defines that the unreliable CoAP.
	ServiceTypeUnreliableCoAP
)

// ServiceTypeParse converts the given string into a ServiceType value.
//
// It returns the corresponding ServiceType constant if the string matches
// a known level name, or an error if the input is invalid.
func ServiceTypeParse(value string) (ServiceType, error) {
	var ret ServiceType
	var err error
	switch {
	case strings.EqualFold(value, "TCP"):
		ret = ServiceTypeTCP
	case strings.EqualFold(value, "UDP"):
		ret = ServiceTypeUDP
	case strings.EqualFold(value, "Ftp"):
		ret = ServiceTypeFTP
	case strings.EqualFold(value, "SMTP"):
		ret = ServiceTypeSMTP
	case strings.EqualFold(value, "Sms"):
		ret = ServiceTypeSMS
	case strings.EqualFold(value, "Hdlc"):
		ret = ServiceTypeHDLC
	case strings.EqualFold(value, "MBus"):
		ret = ServiceTypeMBUS
	case strings.EqualFold(value, "ZigBee"):
		ret = ServiceTypeZigBee
	case strings.EqualFold(value, "DlmsGateway"):
		ret = ServiceTypeDlmsGateway
	case strings.EqualFold(value, "ReliableCoAP"):
		ret = ServiceTypeReliableCoAP
	case strings.EqualFold(value, "UnreliableCoAP"):
		ret = ServiceTypeUnreliableCoAP
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the ServiceType.
// It satisfies fmt.Stringer.
func (g ServiceType) String() string {
	var ret string
	switch g {
	case ServiceTypeTCP:
		ret = "TCP"
	case ServiceTypeUDP:
		ret = "UDP"
	case ServiceTypeFTP:
		ret = "Ftp"
	case ServiceTypeSMTP:
		ret = "SMTP"
	case ServiceTypeSMS:
		ret = "Sms"
	case ServiceTypeHDLC:
		ret = "Hdlc"
	case ServiceTypeMBUS:
		ret = "MBus"
	case ServiceTypeZigBee:
		ret = "ZigBee"
	case ServiceTypeDlmsGateway:
		ret = "DlmsGateway"
	case ServiceTypeReliableCoAP:
		ret = "ReliableCoAP"
	case ServiceTypeUnreliableCoAP:
		ret = "UnreliableCoAP"
	}
	return ret
}

// AllServiceType returns a slice containing all defined ServiceType values.
func AllServiceType() []ServiceType {
	return []ServiceType{
		ServiceTypeTCP,
		ServiceTypeUDP,
		ServiceTypeFTP,
		ServiceTypeSMTP,
		ServiceTypeSMS,
		ServiceTypeHDLC,
		ServiceTypeMBUS,
		ServiceTypeZigBee,
		ServiceTypeDlmsGateway,
		ServiceTypeReliableCoAP,
		ServiceTypeUnreliableCoAP,
	}
}
