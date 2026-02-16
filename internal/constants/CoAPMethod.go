package constants

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

// CoAP methods.
type CoAPMethod int

const (
	// CoAPMethodNone defines that the // Empty command.
	CoAPMethodNone CoAPMethod = iota
	// CoAPMethodGet defines that the // Get command.
	CoAPMethodGet
	// CoAPMethodPost defines that the // Post command.
	CoAPMethodPost
	// CoAPMethodPut defines that the // Put command.
	CoAPMethodPut
	// CoAPMethodDelete defines that the // Delete command.
	CoAPMethodDelete
	// CoAPMethodFetch defines that the // Fetch command.
	CoAPMethodFetch
	// CoAPMethodPatch defines that the // Patch command.
	CoAPMethodPatch
	// CoAPMethodIPatch defines that the // IPatch command.
	CoAPMethodIPatch
)

// CoAPMethodParse converts the given string into a CoAPMethod value.
//
// It returns the corresponding CoAPMethod constant if the string matches
// a known level name, or an error if the input is invalid.
func CoAPMethodParse(value string) (CoAPMethod, error) {
	var ret CoAPMethod
	var err error
	switch strings.ToUpper(value) {
	case "NONE":
		ret = CoAPMethodNone
	case "GET":
		ret = CoAPMethodGet
	case "POST":
		ret = CoAPMethodPost
	case "PUT":
		ret = CoAPMethodPut
	case "DELETE":
		ret = CoAPMethodDelete
	case "FETCH":
		ret = CoAPMethodFetch
	case "PATCH":
		ret = CoAPMethodPatch
	case "IPATCH":
		ret = CoAPMethodIPatch
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the CoAPMethod.
// It satisfies fmt.Stringer.
func (g CoAPMethod) String() string {
	var ret string
	switch g {
	case CoAPMethodNone:
		ret = "NONE"
	case CoAPMethodGet:
		ret = "GET"
	case CoAPMethodPost:
		ret = "POST"
	case CoAPMethodPut:
		ret = "PUT"
	case CoAPMethodDelete:
		ret = "DELETE"
	case CoAPMethodFetch:
		ret = "FETCH"
	case CoAPMethodPatch:
		ret = "PATCH"
	case CoAPMethodIPatch:
		ret = "IPATCH"
	}
	return ret
}
