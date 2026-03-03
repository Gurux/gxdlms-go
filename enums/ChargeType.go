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

// ChargeType enumerates the charge types.
type ChargeType int

const (
	// ChargeTypeConsumptionBasedCollection defines that the charge is based for consumption based collection.
	ChargeTypeConsumptionBasedCollection ChargeType = iota
	// ChargeTypeTimeBasedCollection defines that the the charge is based for time based collection.
	ChargeTypeTimeBasedCollection
	// ChargeTypePaymentEventBasedCollection defines that the charge is based for payment based collection.
	ChargeTypePaymentEventBasedCollection
)

// ChargeTypeParse converts the given string into a ChargeType value.
//
// It returns the corresponding ChargeType constant if the string matches
// a known level name, or an error if the input is invalid.
func ChargeTypeParse(value string) (ChargeType, error) {
	var ret ChargeType
	var err error
	switch {
	case strings.EqualFold(value, "ConsumptionBasedCollection"):
		ret = ChargeTypeConsumptionBasedCollection
	case strings.EqualFold(value, "TimeBasedCollection"):
		ret = ChargeTypeTimeBasedCollection
	case strings.EqualFold(value, "PaymentEventBasedCollection"):
		ret = ChargeTypePaymentEventBasedCollection
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the ChargeType.
// It satisfies fmt.Stringer.
func (g ChargeType) String() string {
	var ret string
	switch g {
	case ChargeTypeConsumptionBasedCollection:
		ret = "ConsumptionBasedCollection"
	case ChargeTypeTimeBasedCollection:
		ret = "TimeBasedCollection"
	case ChargeTypePaymentEventBasedCollection:
		ret = "PaymentEventBasedCollection"
	}
	return ret
}

// AllChargeType returns a slice containing all defined ChargeType values.
func AllChargeType() []ChargeType {
	return []ChargeType{
		ChargeTypeConsumptionBasedCollection,
		ChargeTypeTimeBasedCollection,
		ChargeTypePaymentEventBasedCollection,
	}
}
