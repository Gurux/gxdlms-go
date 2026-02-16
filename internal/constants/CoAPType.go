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

// CoAP type.
type CoAPType int

const (
	// CoAPTypeConfirmable defines that the // Request confirmable.
	CoAPTypeConfirmable CoAPType = iota
	// CoAPTypeNonConfirmable defines that the // Request non-confirmable.
	CoAPTypeNonConfirmable
	// CoAPTypeAcknowledgement defines that the // Response acknowledgement.
	CoAPTypeAcknowledgement
	// CoAPTypeReset defines that the // Meter receives a message, but can't process it.
	CoAPTypeReset
)

// CoAPTypeParse converts the given string into a CoAPType value.
//
// It returns the corresponding CoAPType constant if the string matches
// a known level name, or an error if the input is invalid.
func CoAPTypeParse(value string) (CoAPType, error) {
	var ret CoAPType
	var err error
	switch strings.ToUpper(value) {
	case "CONFIRMABLE":
		ret = CoAPTypeConfirmable
	case "NONCONFIRMABLE":
		ret = CoAPTypeNonConfirmable
	case "ACKNOWLEDGEMENT":
		ret = CoAPTypeAcknowledgement
	case "RESET":
		ret = CoAPTypeReset
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the CoAPType.
// It satisfies fmt.Stringer.
func (g CoAPType) String() string {
	var ret string
	switch g {
	case CoAPTypeConfirmable:
		ret = "CONFIRMABLE"
	case CoAPTypeNonConfirmable:
		ret = "NONCONFIRMABLE"
	case CoAPTypeAcknowledgement:
		ret = "ACKNOWLEDGEMENT"
	case CoAPTypeReset:
		ret = "RESET"
	}
	return ret
}
