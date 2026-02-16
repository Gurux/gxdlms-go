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

// AccessMode The AccessMode enumerates the access modes.
type AccessMode int

const (
	// AccessModeNoAccess defines that there is No access.
	AccessModeNoAccess AccessMode = iota
	// AccessModeRead defines that the client is allowed only reading from the server.
	AccessModeRead
	// AccessModeWrite defines that the client is allowed only writing to the server.
	AccessModeWrite
	// AccessModeReadWrite defines that the client is allowed both reading from the server and writing to it.
	AccessModeReadWrite
	// AccessModeAuthenticatedRead defines that the authenticated read is used.
	AccessModeAuthenticatedRead
	// AccessModeAuthenticatedWrite defines that the authenticated write is used.
	AccessModeAuthenticatedWrite
	// AccessModeAuthenticatedReadWrite defines that the authenticated Read Write is used.
	AccessModeAuthenticatedReadWrite
)

// AccessModeParse converts the given string into a AccessMode value.
//
// It returns the corresponding AccessMode constant if the string matches
// a known level name, or an error if the input is invalid.
func AccessModeParse(value string) (AccessMode, error) {
	var ret AccessMode
	var err error
	switch strings.ToUpper(value) {
	case "NOACCESS":
		ret = AccessModeNoAccess
	case "READ":
		ret = AccessModeRead
	case "WRITE":
		ret = AccessModeWrite
	case "READWRITE":
		ret = AccessModeReadWrite
	case "AUTHENTICATEDREAD":
		ret = AccessModeAuthenticatedRead
	case "AUTHENTICATEDWRITE":
		ret = AccessModeAuthenticatedWrite
	case "AUTHENTICATEDREADWRITE":
		ret = AccessModeAuthenticatedReadWrite
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the AccessMode.
// It satisfies fmt.Stringer.
func (g AccessMode) String() string {
	var ret string
	switch g {
	case AccessModeNoAccess:
		ret = "NOACCESS"
	case AccessModeRead:
		ret = "READ"
	case AccessModeWrite:
		ret = "WRITE"
	case AccessModeReadWrite:
		ret = "READWRITE"
	case AccessModeAuthenticatedRead:
		ret = "AUTHENTICATEDREAD"
	case AccessModeAuthenticatedWrite:
		ret = "AUTHENTICATEDWRITE"
	case AccessModeAuthenticatedReadWrite:
		ret = "AUTHENTICATEDREADWRITE"
	}
	return ret
}
