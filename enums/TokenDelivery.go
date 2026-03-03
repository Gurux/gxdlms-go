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

// TokenDelivery enumerates token delivery methods.
type TokenDelivery int

const (
	// TokenDeliveryRemote defines that the token is delivered via remote communications.
	TokenDeliveryRemote TokenDelivery = iota
	// TokenDeliveryLocal defines that the token is delivered via local communications.
	TokenDeliveryLocal
	// TokenDeliveryManual defines that the token is delivered via manual entry.
	TokenDeliveryManual
)

// TokenDeliveryParse converts the given string into a TokenDelivery value.
//
// It returns the corresponding TokenDelivery constant if the string matches
// a known level name, or an error if the input is invalid.
func TokenDeliveryParse(value string) (TokenDelivery, error) {
	var ret TokenDelivery
	var err error
	switch {
	case strings.EqualFold(value, "Remote"):
		ret = TokenDeliveryRemote
	case strings.EqualFold(value, "Local"):
		ret = TokenDeliveryLocal
	case strings.EqualFold(value, "Manual"):
		ret = TokenDeliveryManual
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the TokenDelivery.
// It satisfies fmt.Stringer.
func (g TokenDelivery) String() string {
	var ret string
	switch g {
	case TokenDeliveryRemote:
		ret = "Remote"
	case TokenDeliveryLocal:
		ret = "Local"
	case TokenDeliveryManual:
		ret = "Manual"
	}
	return ret
}

// AllTokenDelivery returns a slice containing all defined TokenDelivery values.
func AllTokenDelivery() []TokenDelivery {
	return []TokenDelivery{
		TokenDeliveryRemote,
		TokenDeliveryLocal,
		TokenDeliveryManual,
	}
}
