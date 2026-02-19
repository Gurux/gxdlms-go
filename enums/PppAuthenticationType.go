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

// PppAuthenticationType Ppp Authentication Type
type PppAuthenticationType int

const (
	// PppAuthenticationTypeNone defines that the no authentication.
	PppAuthenticationTypeNone PppAuthenticationType = iota
	// PppAuthenticationTypePAP defines that the pAP Login
	PppAuthenticationTypePAP
	// PppAuthenticationTypeCHAP defines that the cHAP-algorithm
	PppAuthenticationTypeCHAP
)

// PppAuthenticationTypeParse converts the given string into a PppAuthenticationType value.
//
// It returns the corresponding PppAuthenticationType constant if the string matches
// a known level name, or an error if the input is invalid.
func PppAuthenticationTypeParse(value string) (PppAuthenticationType, error) {
	var ret PppAuthenticationType
	var err error
	switch {
	case strings.EqualFold(value, "None"):
		ret = PppAuthenticationTypeNone
	case strings.EqualFold(value, "PAP"):
		ret = PppAuthenticationTypePAP
	case strings.EqualFold(value, "CHAP"):
		ret = PppAuthenticationTypeCHAP
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the PppAuthenticationType.
// It satisfies fmt.Stringer.
func (g PppAuthenticationType) String() string {
	var ret string
	switch g {
	case PppAuthenticationTypeNone:
		ret = "None"
	case PppAuthenticationTypePAP:
		ret = "PAP"
	case PppAuthenticationTypeCHAP:
		ret = "CHAP"
	}
	return ret
}

// AllPppAuthenticationType returns a slice containing all defined PppAuthenticationType values.
func AllPppAuthenticationType() []PppAuthenticationType {
	return []PppAuthenticationType{
		PppAuthenticationTypeNone,
		PppAuthenticationTypePAP,
		PppAuthenticationTypeCHAP,
	}
}
