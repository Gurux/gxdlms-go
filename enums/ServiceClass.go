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

// ServiceClass Is server responding for the request.
type ServiceClass int

const (
	// ServiceClassUnConfirmed defines that the server doesn't respond or send data.
	ServiceClassUnConfirmed ServiceClass = iota
	// ServiceClassConfirmed defines that the the server sends data or an acknowledge message.
	ServiceClassConfirmed
)

// ServiceClassParse converts the given string into a ServiceClass value.
//
// It returns the corresponding ServiceClass constant if the string matches
// a known level name, or an error if the input is invalid.
func ServiceClassParse(value string) (ServiceClass, error) {
	var ret ServiceClass
	var err error
	switch {
	case strings.EqualFold(value, "UnConfirmed"):
		ret = ServiceClassUnConfirmed
	case strings.EqualFold(value, "Confirmed"):
		ret = ServiceClassConfirmed
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the ServiceClass.
// It satisfies fmt.Stringer.
func (g ServiceClass) String() string {
	var ret string
	switch g {
	case ServiceClassUnConfirmed:
		ret = "UnConfirmed"
	case ServiceClassConfirmed:
		ret = "Confirmed"
	}
	return ret
}

// AllServiceClass returns a slice containing all defined ServiceClass values.
func AllServiceClass() []ServiceClass {
	return []ServiceClass{
		ServiceClassUnConfirmed,
		ServiceClassConfirmed,
	}
}
