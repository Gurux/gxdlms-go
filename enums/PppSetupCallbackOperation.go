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

// PppSetupCallbackOperation Point to point setup callback operations.
type PppSetupCallbackOperation int

const (
	// PppSetupCallbackOperationUser defines that the location is determined by user authentication.
	PppSetupCallbackOperationUser PppSetupCallbackOperation = iota
	// PppSetupCallbackOperationDialling defines that the dialling string.
	PppSetupCallbackOperationDialling
	// PppSetupCallbackOperationLocation defines that the location identifier.
	PppSetupCallbackOperationLocation
	// PppSetupCallbackOperationE164 defines that the e.164 number.
	PppSetupCallbackOperationE164
	// PppSetupCallbackOperationX500 defines that the x500 distinguished name.
	PppSetupCallbackOperationX500
	// PppSetupCallbackOperationCBCP defines that the location is determined during CBCP negotiation.
	PppSetupCallbackOperationCBCP PppSetupCallbackOperation = 6
)

// PppSetupCallbackOperationParse converts the given string into a PppSetupCallbackOperation value.
//
// It returns the corresponding PppSetupCallbackOperation constant if the string matches
// a known level name, or an error if the input is invalid.
func PppSetupCallbackOperationParse(value string) (PppSetupCallbackOperation, error) {
	var ret PppSetupCallbackOperation
	var err error
	switch strings.ToUpper(value) {
	case "USER":
		ret = PppSetupCallbackOperationUser
	case "DIALLING":
		ret = PppSetupCallbackOperationDialling
	case "LOCATION":
		ret = PppSetupCallbackOperationLocation
	case "E_164":
		ret = PppSetupCallbackOperationE164
	case "X500":
		ret = PppSetupCallbackOperationX500
	case "CBCP":
		ret = PppSetupCallbackOperationCBCP
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the PppSetupCallbackOperation.
// It satisfies fmt.Stringer.
func (g PppSetupCallbackOperation) String() string {
	var ret string
	switch g {
	case PppSetupCallbackOperationUser:
		ret = "USER"
	case PppSetupCallbackOperationDialling:
		ret = "DIALLING"
	case PppSetupCallbackOperationLocation:
		ret = "LOCATION"
	case PppSetupCallbackOperationE164:
		ret = "E_164"
	case PppSetupCallbackOperationX500:
		ret = "X500"
	case PppSetupCallbackOperationCBCP:
		ret = "CBCP"
	}
	return ret
}
