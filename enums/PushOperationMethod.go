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

// PushOperationMethod :  defines what service class is used with push messages.
type PushOperationMethod int

const (
	// PushOperationMethodUnconfirmedFailure defines that the unconfirmed, retry on supporting protocol layer failure.
	PushOperationMethodUnconfirmedFailure PushOperationMethod = iota
	// PushOperationMethodUnconfirmedMissing defines that the unconfirmed, retry on missing supporting protocol layer confirmation.
	PushOperationMethodUnconfirmedMissing
	// PushOperationMethodConfirmed defines that the confirmed, retry on missing confirmation.
	PushOperationMethodConfirmed
)

// PushOperationMethodParse converts the given string into a PushOperationMethod value.
//
// It returns the corresponding PushOperationMethod constant if the string matches
// a known level name, or an error if the input is invalid.
func PushOperationMethodParse(value string) (PushOperationMethod, error) {
	var ret PushOperationMethod
	var err error
	switch {
	case strings.EqualFold(value, "UnconfirmedFailure"):
		ret = PushOperationMethodUnconfirmedFailure
	case strings.EqualFold(value, "UnconfirmedMissing"):
		ret = PushOperationMethodUnconfirmedMissing
	case strings.EqualFold(value, "Confirmed"):
		ret = PushOperationMethodConfirmed
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the PushOperationMethod.
// It satisfies fmt.Stringer.
func (g PushOperationMethod) String() string {
	var ret string
	switch g {
	case PushOperationMethodUnconfirmedFailure:
		ret = "UnconfirmedFailure"
	case PushOperationMethodUnconfirmedMissing:
		ret = "UnconfirmedMissing"
	case PushOperationMethodConfirmed:
		ret = "Confirmed"
	}
	return ret
}

// AllPushOperationMethod returns a slice containing all defined PushOperationMethod values.
func AllPushOperationMethod() []PushOperationMethod {
	return []PushOperationMethod{
		PushOperationMethodUnconfirmedFailure,
		PushOperationMethodUnconfirmedMissing,
		PushOperationMethodConfirmed,
	}
}
