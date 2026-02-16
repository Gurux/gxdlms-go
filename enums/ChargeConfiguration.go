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

// ChargeConfiguration :  enumeration types.
type ChargeConfiguration int

const (
	// ChargeConfigurationNone defines that none bit is set.
	ChargeConfigurationNone ChargeConfiguration = iota
	// ChargeConfigurationPercentageBasedCollection defines that percentage based collection bit is set.
	ChargeConfigurationPercentageBasedCollection ChargeConfiguration = 0x1
	// ChargeConfigurationContinuousCollection defines that continuous collection bit is set.
	ChargeConfigurationContinuousCollection ChargeConfiguration = 0x2
)

// ChargeConfigurationParse converts the given string into a ChargeConfiguration value.
//
// It returns the corresponding ChargeConfiguration constant if the string matches
// a known level name, or an error if the input is invalid.
func ChargeConfigurationParse(value string) (ChargeConfiguration, error) {
	var ret ChargeConfiguration
	var err error
	switch strings.ToUpper(value) {
	case "NONE":
		ret = ChargeConfigurationNone
	case "PERCENTAGEBASEDCOLLECTION":
		ret = ChargeConfigurationPercentageBasedCollection
	case "CONTINUOUSCOLLECTION":
		ret = ChargeConfigurationContinuousCollection
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the ChargeConfiguration.
// It satisfies fmt.Stringer.
func (g ChargeConfiguration) String() string {
	var ret string
	switch g {
	case ChargeConfigurationNone:
		ret = "NONE"
	case ChargeConfigurationPercentageBasedCollection:
		ret = "PERCENTAGEBASEDCOLLECTION"
	case ChargeConfigurationContinuousCollection:
		ret = "CONTINUOUSCOLLECTION"
	}
	return ret
}
