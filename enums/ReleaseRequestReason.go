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

// RequestTypes enumerates the replies of the server to a client's request,
// indicating the request type.
type ReleaseRequestReason int

const (
	// ReleaseRequestReasonNormal defines that the // Client closes connection as normal.
	ReleaseRequestReasonNormal ReleaseRequestReason = iota
	// ReleaseRequestReasonUrgent defines that the // Client closes connection as urgent.
	ReleaseRequestReasonUrgent
	// ReleaseRequestReasonUserDefined defines that the // Client closes connection user defined reason.
	ReleaseRequestReasonUserDefined ReleaseRequestReason = 30
)

// ReleaseRequestReasonParse converts the given string into a ReleaseRequestReason value.
//
// It returns the corresponding ReleaseRequestReason constant if the string matches
// a known level name, or an error if the input is invalid.
func ReleaseRequestReasonParse(value string) (ReleaseRequestReason, error) {
	var ret ReleaseRequestReason
	var err error
	switch strings.ToUpper(value) {
	case "NORMAL":
		ret = ReleaseRequestReasonNormal
	case "URGENT":
		ret = ReleaseRequestReasonUrgent
	case "USERDEFINED":
		ret = ReleaseRequestReasonUserDefined
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the ReleaseRequestReason.
// It satisfies fmt.Stringer.
func (g ReleaseRequestReason) String() string {
	var ret string
	switch g {
	case ReleaseRequestReasonNormal:
		ret = "NORMAL"
	case ReleaseRequestReasonUrgent:
		ret = "URGENT"
	case ReleaseRequestReasonUserDefined:
		ret = "USERDEFINED"
	}
	return ret
}
