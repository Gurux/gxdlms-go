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

// PaymentMode Enumerates payment Modes.
type PaymentMode int

const (
	// PaymentModeCredit defines that the credit mode.
	PaymentModeCredit PaymentMode = 1
	// PaymentModePrepayment defines that the prepayment mode.
	PaymentModePrepayment PaymentMode = 2
)

// PaymentModeParse converts the given string into a PaymentMode value.
//
// It returns the corresponding PaymentMode constant if the string matches
// a known level name, or an error if the input is invalid.
func PaymentModeParse(value string) (PaymentMode, error) {
	var ret PaymentMode
	var err error
	switch strings.ToUpper(value) {
	case "CREDIT":
		ret = PaymentModeCredit
	case "PREPAYMENT":
		ret = PaymentModePrepayment
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the PaymentMode.
// It satisfies fmt.Stringer.
func (g PaymentMode) String() string {
	var ret string
	switch g {
	case PaymentModeCredit:
		ret = "CREDIT"
	case PaymentModePrepayment:
		ret = "PREPAYMENT"
	}
	return ret
}
