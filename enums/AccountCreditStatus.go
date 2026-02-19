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

// AccountCreditStatus enumerates the account redit status.
type AccountCreditStatus int

const (
	// AccountCreditStatusNone defines that none bit is set.
	AccountCreditStatusNone AccountCreditStatus = iota
	// AccountCreditStatusInCredit defines that in credit bit is set.
	AccountCreditStatusInCredit AccountCreditStatus = 0x1
	// AccountCreditStatusLowCredit defines that low credit bit is set.
	AccountCreditStatusLowCredit AccountCreditStatus = 0x2
	// AccountCreditStatusNextCreditEnabled defines that next credit enabled bit is set.
	AccountCreditStatusNextCreditEnabled AccountCreditStatus = 0x4
	// AccountCreditStatusNextCreditSelectable defines that next credit selectable bit is set.
	AccountCreditStatusNextCreditSelectable AccountCreditStatus = 0x8
	// AccountCreditStatusCreditReferenceList defines that credit reference list bit is set.
	AccountCreditStatusCreditReferenceList AccountCreditStatus = 0x10
	// AccountCreditStatusSelectableCreditInUse defines that selectable credit in use bit is set.
	AccountCreditStatusSelectableCreditInUse AccountCreditStatus = 0x20
	// AccountCreditStatusOutOfCredit defines that out of credit bit is set.
	AccountCreditStatusOutOfCredit AccountCreditStatus = 0x40
	// AccountCreditStatusReserved defines that reserved bit is set.
	AccountCreditStatusReserved AccountCreditStatus = 0x80
)

// AccountCreditStatusParse converts the given string into a AccountCreditStatus value.
//
// It returns the corresponding AccountCreditStatus constant if the string matches
// a known level name, or an error if the input is invalid.
func AccountCreditStatusParse(value string) (AccountCreditStatus, error) {
	var ret AccountCreditStatus
	var err error
	switch {
	case strings.EqualFold(value, "None"):
		ret = AccountCreditStatusNone
	case strings.EqualFold(value, "InCredit"):
		ret = AccountCreditStatusInCredit
	case strings.EqualFold(value, "LowCredit"):
		ret = AccountCreditStatusLowCredit
	case strings.EqualFold(value, "NextCreditEnabled"):
		ret = AccountCreditStatusNextCreditEnabled
	case strings.EqualFold(value, "NextCreditSelectable"):
		ret = AccountCreditStatusNextCreditSelectable
	case strings.EqualFold(value, "CreditReferenceList"):
		ret = AccountCreditStatusCreditReferenceList
	case strings.EqualFold(value, "SelectableCreditInUse"):
		ret = AccountCreditStatusSelectableCreditInUse
	case strings.EqualFold(value, "OutOfCredit"):
		ret = AccountCreditStatusOutOfCredit
	case strings.EqualFold(value, "Reserved"):
		ret = AccountCreditStatusReserved
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the AccountCreditStatus.
// It satisfies fmt.Stringer.
func (g AccountCreditStatus) String() string {
	var ret string
	switch g {
	case AccountCreditStatusNone:
		ret = "None"
	case AccountCreditStatusInCredit:
		ret = "InCredit"
	case AccountCreditStatusLowCredit:
		ret = "LowCredit"
	case AccountCreditStatusNextCreditEnabled:
		ret = "NextCreditEnabled"
	case AccountCreditStatusNextCreditSelectable:
		ret = "NextCreditSelectable"
	case AccountCreditStatusCreditReferenceList:
		ret = "CreditReferenceList"
	case AccountCreditStatusSelectableCreditInUse:
		ret = "SelectableCreditInUse"
	case AccountCreditStatusOutOfCredit:
		ret = "OutOfCredit"
	case AccountCreditStatusReserved:
		ret = "Reserved"
	}
	return ret
}

// AllAccountCreditStatus returns a slice containing all defined AccountCreditStatus values.
func AllAccountCreditStatus() []AccountCreditStatus {
	return []AccountCreditStatus{
		AccountCreditStatusNone,
		AccountCreditStatusInCredit,
		AccountCreditStatusLowCredit,
		AccountCreditStatusNextCreditEnabled,
		AccountCreditStatusNextCreditSelectable,
		AccountCreditStatusCreditReferenceList,
		AccountCreditStatusSelectableCreditInUse,
		AccountCreditStatusOutOfCredit,
		AccountCreditStatusReserved,
	}
}
