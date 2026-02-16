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
type ReleaseResponseReason int

const (
	// ReleaseResponseReasonNormal defines that the // Client closes connection as normal.
	ReleaseResponseReasonNormal ReleaseResponseReason = iota
	// ReleaseResponseReasonNotFinished defines that the // Connection is not finished.
	ReleaseResponseReasonNotFinished
	// ReleaseResponseReasonUserDefined defines that the // Client closes connection user defined reason.
	ReleaseResponseReasonUserDefined ReleaseResponseReason = 30
)

// ReleaseResponseReasonParse converts the given string into a ReleaseResponseReason value.
//
// It returns the corresponding ReleaseResponseReason constant if the string matches
// a known level name, or an error if the input is invalid.
func ReleaseResponseReasonParse(value string) (ReleaseResponseReason, error) {
	var ret ReleaseResponseReason
	var err error
	switch strings.ToUpper(value) {
	case "NORMAL":
		ret = ReleaseResponseReasonNormal
	case "NOTFINISHED":
		ret = ReleaseResponseReasonNotFinished
	case "USERDEFINED":
		ret = ReleaseResponseReasonUserDefined
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the ReleaseResponseReason.
// It satisfies fmt.Stringer.
func (g ReleaseResponseReason) String() string {
	var ret string
	switch g {
	case ReleaseResponseReasonNormal:
		ret = "NORMAL"
	case ReleaseResponseReasonNotFinished:
		ret = "NOTFINISHED"
	case ReleaseResponseReasonUserDefined:
		ret = "USERDEFINED"
	}
	return ret
}
