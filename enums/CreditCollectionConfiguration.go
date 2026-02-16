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

// CreditCollectionConfiguration Defines behaviour under specific conditions
type CreditCollectionConfiguration int

const (
	// CreditCollectionConfigurationNone defines that none bit is set.
	CreditCollectionConfigurationNone CreditCollectionConfiguration = iota
	// CreditCollectionConfigurationDisconnected defines that collect when supply disconnected bit is set.
	CreditCollectionConfigurationDisconnected CreditCollectionConfiguration = 0x1
	// CreditCollectionConfigurationLoadLimiting defines that collect in load limiting periods bit is set.
	CreditCollectionConfigurationLoadLimiting CreditCollectionConfiguration = 0x2
	// CreditCollectionConfigurationFriendlyCredit defines that collect in friendly credit periods bit is set.
	CreditCollectionConfigurationFriendlyCredit CreditCollectionConfiguration = 0x4
)

// CreditCollectionConfigurationParse converts the given string into a CreditCollectionConfiguration value.
//
// It returns the corresponding CreditCollectionConfiguration constant if the string matches
// a known level name, or an error if the input is invalid.
func CreditCollectionConfigurationParse(value string) (CreditCollectionConfiguration, error) {
	var ret CreditCollectionConfiguration
	var err error
	switch strings.ToUpper(value) {
	case "NONE":
		ret = CreditCollectionConfigurationNone
	case "DISCONNECTED":
		ret = CreditCollectionConfigurationDisconnected
	case "LOADLIMITING":
		ret = CreditCollectionConfigurationLoadLimiting
	case "FRIENDLYCREDIT":
		ret = CreditCollectionConfigurationFriendlyCredit
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the CreditCollectionConfiguration.
// It satisfies fmt.Stringer.
func (g CreditCollectionConfiguration) String() string {
	var ret string
	switch g {
	case CreditCollectionConfigurationNone:
		ret = "NONE"
	case CreditCollectionConfigurationDisconnected:
		ret = "DISCONNECTED"
	case CreditCollectionConfigurationLoadLimiting:
		ret = "LOADLIMITING"
	case CreditCollectionConfigurationFriendlyCredit:
		ret = "FRIENDLYCREDIT"
	}
	return ret
}
