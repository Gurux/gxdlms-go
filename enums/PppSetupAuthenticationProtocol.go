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

// PppSetupAuthenticationProtocol The value of the Auth-Prot (Authentication Protocol) element indicates
//
//	the authentication protocol used on the given PPP link.
type PppSetupAuthenticationProtocol int

const (
	// PppSetupAuthenticationProtocolNone defines that the no authentication protocol is used.
	PppSetupAuthenticationProtocolNone PppSetupAuthenticationProtocol = iota
	// PppSetupAuthenticationProtocolPAP defines that the the PAP protocol is used.
	PppSetupAuthenticationProtocolPAP PppSetupAuthenticationProtocol = 0xc023
	// PppSetupAuthenticationProtocolCHAP defines that the the CHAP protocol is used.
	PppSetupAuthenticationProtocolCHAP PppSetupAuthenticationProtocol = 0xc223
	// PppSetupAuthenticationProtocolEAP defines that the the EAP protocol is used.
	PppSetupAuthenticationProtocolEAP PppSetupAuthenticationProtocol = 0xc227
)

// PppSetupAuthenticationProtocolParse converts the given string into a PppSetupAuthenticationProtocol value.
//
// It returns the corresponding PppSetupAuthenticationProtocol constant if the string matches
// a known level name, or an error if the input is invalid.
func PppSetupAuthenticationProtocolParse(value string) (PppSetupAuthenticationProtocol, error) {
	var ret PppSetupAuthenticationProtocol
	var err error
	switch strings.ToUpper(value) {
	case "NONE":
		ret = PppSetupAuthenticationProtocolNone
	case "PAP":
		ret = PppSetupAuthenticationProtocolPAP
	case "CHAP":
		ret = PppSetupAuthenticationProtocolCHAP
	case "EAP":
		ret = PppSetupAuthenticationProtocolEAP
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the PppSetupAuthenticationProtocol.
// It satisfies fmt.Stringer.
func (g PppSetupAuthenticationProtocol) String() string {
	var ret string
	switch g {
	case PppSetupAuthenticationProtocolNone:
		ret = "NONE"
	case PppSetupAuthenticationProtocolPAP:
		ret = "PAP"
	case PppSetupAuthenticationProtocolCHAP:
		ret = "CHAP"
	case PppSetupAuthenticationProtocolEAP:
		ret = "EAP"
	}
	return ret
}
