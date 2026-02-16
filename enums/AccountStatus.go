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

type AccountStatus int

const (
	// AccountStatusNewInactiveAccount defines that the new (inactive) account.
	AccountStatusNewInactiveAccount AccountStatus = 1
	// AccountStatusAccountActive defines that the account active.
	AccountStatusAccountActive AccountStatus = 2
	// AccountStatusAccountClosed defines that the account closed.
	AccountStatusAccountClosed AccountStatus = 3
)

// AccountStatusParse converts the given string into a AccountStatus value.
//
// It returns the corresponding AccountStatus constant if the string matches
// a known level name, or an error if the input is invalid.
func AccountStatusParse(value string) (AccountStatus, error) {
	var ret AccountStatus
	var err error
	switch strings.ToUpper(value) {
	case "NEWINACTIVEACCOUNT":
		ret = AccountStatusNewInactiveAccount
	case "ACCOUNTACTIVE":
		ret = AccountStatusAccountActive
	case "ACCOUNTCLOSED":
		ret = AccountStatusAccountClosed
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the AccountStatus.
// It satisfies fmt.Stringer.
func (g AccountStatus) String() string {
	var ret string
	switch g {
	case AccountStatusNewInactiveAccount:
		ret = "NEWINACTIVEACCOUNT"
	case AccountStatusAccountActive:
		ret = "ACCOUNTACTIVE"
	case AccountStatusAccountClosed:
		ret = "ACCOUNTCLOSED"
	}
	return ret
}
