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

// SecuritySuite :  Specifies authentication, encryption and key wrapping algorithm.
type SecuritySuite int

const (
	// SecuritySuite0 defines that the gMAC ciphering is used with security setup version 0.
	SecuritySuite0 SecuritySuite = iota
	// SecuritySuite1 defines that the eCDSA P-256 ciphering is used.
	SecuritySuite1
	// SecuritySuite2 defines that the eCDSA P-384 ciphering is used.
	SecuritySuite2
)

// SecuritySuiteParse converts the given string into a SecuritySuite value.
//
// It returns the corresponding SecuritySuite constant if the string matches
// a known level name, or an error if the input is invalid.
func SecuritySuiteParse(value string) (SecuritySuite, error) {
	var ret SecuritySuite
	var err error
	switch strings.ToUpper(value) {
	case "SUITE0":
		ret = SecuritySuite0
	case "SUITE1":
		ret = SecuritySuite1
	case "SUITE2":
		ret = SecuritySuite2
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the SecuritySuite.
// It satisfies fmt.Stringer.
func (g SecuritySuite) String() string {
	var ret string
	switch g {
	case SecuritySuite0:
		ret = "SUITE0"
	case SecuritySuite1:
		ret = "SUITE1"
	case SecuritySuite2:
		ret = "SUITE2"
	}
	return ret
}
