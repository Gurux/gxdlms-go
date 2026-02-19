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

// Service describes service errors.
type Service int

const (
	// ServiceOther defines that the other error.
	ServiceOther Service = iota
	// ServicePduSize defines that the PDU size is wrong.
	ServicePduSize
	// ServiceUnsupported defines that the service is unsupported.
	ServiceUnsupported
)

// ServiceParse converts the given string into a Service value.
//
// It returns the corresponding Service constant if the string matches
// a known level name, or an error if the input is invalid.
func ServiceParse(value string) (Service, error) {
	var ret Service
	var err error
	switch {
	case strings.EqualFold(value, "Other"):
		ret = ServiceOther
	case strings.EqualFold(value, "PduSize"):
		ret = ServicePduSize
	case strings.EqualFold(value, "Unsupported"):
		ret = ServiceUnsupported
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the Service.
// It satisfies fmt.Stringer.
func (g Service) String() string {
	var ret string
	switch g {
	case ServiceOther:
		ret = "Other"
	case ServicePduSize:
		ret = "PduSize"
	case ServiceUnsupported:
		ret = "Unsupported"
	}
	return ret
}

// AllService returns a slice containing all defined Service values.
func AllService() []Service {
	return []Service{
		ServiceOther,
		ServicePduSize,
		ServiceUnsupported,
	}
}
