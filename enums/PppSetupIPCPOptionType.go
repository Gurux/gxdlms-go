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

// PppSetupIPCPOptionType PPP Setup IPCP option types.
type PppSetupIPCPOptionType int

const (
	PppSetupIPCPOptionTypeIPCompressionProtocol PppSetupIPCPOptionType = 2
	PppSetupIPCPOptionTypePrefLocalIP           PppSetupIPCPOptionType = 3
	PppSetupIPCPOptionTypePrefPeerIP            PppSetupIPCPOptionType = 20
	PppSetupIPCPOptionTypeGAO                   PppSetupIPCPOptionType = 21
	PppSetupIPCPOptionTypeUSIP                  PppSetupIPCPOptionType = 22
)

// PppSetupIPCPOptionTypeParse converts the given string into a PppSetupIPCPOptionType value.
//
// It returns the corresponding PppSetupIPCPOptionType constant if the string matches
// a known level name, or an error if the input is invalid.
func PppSetupIPCPOptionTypeParse(value string) (PppSetupIPCPOptionType, error) {
	var ret PppSetupIPCPOptionType
	var err error
	switch {
	case strings.EqualFold(value, "IPCompressionProtocol"):
		ret = PppSetupIPCPOptionTypeIPCompressionProtocol
	case strings.EqualFold(value, "PrefLocalIP"):
		ret = PppSetupIPCPOptionTypePrefLocalIP
	case strings.EqualFold(value, "PrefPeerIP"):
		ret = PppSetupIPCPOptionTypePrefPeerIP
	case strings.EqualFold(value, "GAO"):
		ret = PppSetupIPCPOptionTypeGAO
	case strings.EqualFold(value, "USIP"):
		ret = PppSetupIPCPOptionTypeUSIP
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the PppSetupIPCPOptionType.
// It satisfies fmt.Stringer.
func (g PppSetupIPCPOptionType) String() string {
	var ret string
	switch g {
	case PppSetupIPCPOptionTypeIPCompressionProtocol:
		ret = "IPCompressionProtocol"
	case PppSetupIPCPOptionTypePrefLocalIP:
		ret = "PrefLocalIP"
	case PppSetupIPCPOptionTypePrefPeerIP:
		ret = "PrefPeerIP"
	case PppSetupIPCPOptionTypeGAO:
		ret = "GAO"
	case PppSetupIPCPOptionTypeUSIP:
		ret = "USIP"
	}
	return ret
}

// AllPppSetupIPCPOptionType returns a slice containing all defined PppSetupIPCPOptionType values.
func AllPppSetupIPCPOptionType() []PppSetupIPCPOptionType {
	return []PppSetupIPCPOptionType{
		PppSetupIPCPOptionTypeIPCompressionProtocol,
		PppSetupIPCPOptionTypePrefLocalIP,
		PppSetupIPCPOptionTypePrefPeerIP,
		PppSetupIPCPOptionTypeGAO,
		PppSetupIPCPOptionTypeUSIP,
	}
}
