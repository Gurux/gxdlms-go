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

// Credit Type.
type CreditType int

const (
	// CreditTypeToken defines that the // Token credit.
	CreditTypeToken CreditType = iota
	// CreditTypeReserved defines that the // Reserved credit.
	CreditTypeReserved
	// CreditTypeEmergency defines that the // Emergency credit.
	CreditTypeEmergency
	// CreditTypeTimeBased defines that the // TimeBased credit.
	CreditTypeTimeBased
	// CreditTypeConsumptionBased defines that the // Consumption based credit.
	CreditTypeConsumptionBased
)

// CreditTypeParse converts the given string into a CreditType value.
//
// It returns the corresponding CreditType constant if the string matches
// a known level name, or an error if the input is invalid.
func CreditTypeParse(value string) (CreditType, error) {
	var ret CreditType
	var err error
	switch {
	case strings.EqualFold(value, "Token"):
		ret = CreditTypeToken
	case strings.EqualFold(value, "Reserved"):
		ret = CreditTypeReserved
	case strings.EqualFold(value, "Emergency"):
		ret = CreditTypeEmergency
	case strings.EqualFold(value, "TimeBased"):
		ret = CreditTypeTimeBased
	case strings.EqualFold(value, "ConsumptionBased"):
		ret = CreditTypeConsumptionBased
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the CreditType.
// It satisfies fmt.Stringer.
func (g CreditType) String() string {
	var ret string
	switch g {
	case CreditTypeToken:
		ret = "Token"
	case CreditTypeReserved:
		ret = "Reserved"
	case CreditTypeEmergency:
		ret = "Emergency"
	case CreditTypeTimeBased:
		ret = "TimeBased"
	case CreditTypeConsumptionBased:
		ret = "ConsumptionBased"
	}
	return ret
}

// AllCreditType returns a slice containing all defined CreditType values.
func AllCreditType() []CreditType {
	return []CreditType{
		CreditTypeToken,
		CreditTypeReserved,
		CreditTypeEmergency,
		CreditTypeTimeBased,
		CreditTypeConsumptionBased,
	}
}
