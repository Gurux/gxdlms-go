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

// CreditStatus Credit Status.
type CreditStatus int

const (
	// CreditStatusEnabled defines that the enabled state.
	CreditStatusEnabled CreditStatus = iota
	// CreditStatusSelectable defines that the selectable state.
	CreditStatusSelectable
	// CreditStatusInvoked defines that the selected/Invoked state.
	CreditStatusInvoked
	// CreditStatusInUse defines that the in use state.
	CreditStatusInUse
	// CreditStatusConsumed defines that the consumed state.
	CreditStatusConsumed
)

// CreditStatusParse converts the given string into a CreditStatus value.
//
// It returns the corresponding CreditStatus constant if the string matches
// a known level name, or an error if the input is invalid.
func CreditStatusParse(value string) (CreditStatus, error) {
	var ret CreditStatus
	var err error
	switch {
	case strings.EqualFold(value, "Enabled"):
		ret = CreditStatusEnabled
	case strings.EqualFold(value, "Selectable"):
		ret = CreditStatusSelectable
	case strings.EqualFold(value, "Invoked"):
		ret = CreditStatusInvoked
	case strings.EqualFold(value, "InUse"):
		ret = CreditStatusInUse
	case strings.EqualFold(value, "Consumed"):
		ret = CreditStatusConsumed
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the CreditStatus.
// It satisfies fmt.Stringer.
func (g CreditStatus) String() string {
	var ret string
	switch g {
	case CreditStatusEnabled:
		ret = "Enabled"
	case CreditStatusSelectable:
		ret = "Selectable"
	case CreditStatusInvoked:
		ret = "Invoked"
	case CreditStatusInUse:
		ret = "InUse"
	case CreditStatusConsumed:
		ret = "Consumed"
	}
	return ret
}

// AllCreditStatus returns a slice containing all defined CreditStatus values.
func AllCreditStatus() []CreditStatus {
	return []CreditStatus{
	CreditStatusEnabled,
	CreditStatusSelectable,
	CreditStatusInvoked,
	CreditStatusInUse,
	CreditStatusConsumed,
	}
}
