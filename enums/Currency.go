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

// Currency Used currency.
type Currency int

const (
	// CurrencyTime defines that the time.
	CurrencyTime Currency = iota
	// CurrencyConsumption defines that the consumption.
	CurrencyConsumption
	// CurrencyMonetary defines that the monetary.
	CurrencyMonetary
)

// CurrencyParse converts the given string into a Currency value.
//
// It returns the corresponding Currency constant if the string matches
// a known level name, or an error if the input is invalid.
func CurrencyParse(value string) (Currency, error) {
	var ret Currency
	var err error
	switch {
	case strings.EqualFold(value, "Time"):
		ret = CurrencyTime
	case strings.EqualFold(value, "Consumption"):
		ret = CurrencyConsumption
	case strings.EqualFold(value, "Monetary"):
		ret = CurrencyMonetary
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the Currency.
// It satisfies fmt.Stringer.
func (g Currency) String() string {
	var ret string
	switch g {
	case CurrencyTime:
		ret = "Time"
	case CurrencyConsumption:
		ret = "Consumption"
	case CurrencyMonetary:
		ret = "Monetary"
	}
	return ret
}

// AllCurrency returns a slice containing all defined Currency values.
func AllCurrency() []Currency {
	return []Currency{
	CurrencyTime,
	CurrencyConsumption,
	CurrencyMonetary,
	}
}
