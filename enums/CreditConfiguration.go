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

// CreditConfiguration Enumerated :  values.
type CreditConfiguration int

const (
	// CreditConfigurationNone defines that none bit is set.
	CreditConfigurationNone CreditConfiguration = iota
	// CreditConfigurationVisual defines that requires visual indication, bit is set.
	CreditConfigurationVisual CreditConfiguration = 0x1
	// CreditConfigurationConfirmation defines that requires confirmation before it can be selected/invoked bit is set.
	CreditConfigurationConfirmation CreditConfiguration = 0x2
	// CreditConfigurationPaidBack defines that requires the credit amount to be paid back bit is set.
	CreditConfigurationPaidBack CreditConfiguration = 0x4
	// CreditConfigurationResettable defines that resettable bit is set.
	CreditConfigurationResettable CreditConfiguration = 0x8
	// CreditConfigurationTokens defines that able to receive credit amounts from tokens bit is set.
	CreditConfigurationTokens CreditConfiguration = 0x10
)

// CreditConfigurationParse converts the given string into a CreditConfiguration value.
//
// It returns the corresponding CreditConfiguration constant if the string matches
// a known level name, or an error if the input is invalid.
func CreditConfigurationParse(value string) (CreditConfiguration, error) {
	var ret CreditConfiguration
	var err error
	switch {
	case strings.EqualFold(value, "None"):
		ret = CreditConfigurationNone
	case strings.EqualFold(value, "Visual"):
		ret = CreditConfigurationVisual
	case strings.EqualFold(value, "Confirmation"):
		ret = CreditConfigurationConfirmation
	case strings.EqualFold(value, "PaidBack"):
		ret = CreditConfigurationPaidBack
	case strings.EqualFold(value, "Resettable"):
		ret = CreditConfigurationResettable
	case strings.EqualFold(value, "Tokens"):
		ret = CreditConfigurationTokens
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the CreditConfiguration.
// It satisfies fmt.Stringer.
func (g CreditConfiguration) String() string {
	var ret string
	switch g {
	case CreditConfigurationNone:
		ret = "None"
	case CreditConfigurationVisual:
		ret = "Visual"
	case CreditConfigurationConfirmation:
		ret = "Confirmation"
	case CreditConfigurationPaidBack:
		ret = "PaidBack"
	case CreditConfigurationResettable:
		ret = "Resettable"
	case CreditConfigurationTokens:
		ret = "Tokens"
	}
	return ret
}

// AllCreditConfiguration returns a slice containing all defined CreditConfiguration values.
func AllCreditConfiguration() []CreditConfiguration {
	return []CreditConfiguration{
	CreditConfigurationNone,
	CreditConfigurationVisual,
	CreditConfigurationConfirmation,
	CreditConfigurationPaidBack,
	CreditConfigurationResettable,
	CreditConfigurationTokens,
	}
}
