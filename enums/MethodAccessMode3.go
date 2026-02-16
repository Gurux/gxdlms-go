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

// MethodAccessMode3 The MethodAccessMode enumerates the method access modes for logical name association version 3.
type MethodAccessMode3 int

const (
	// MethodAccessMode3NoAccess defines that client can't use method.
	MethodAccessMode3NoAccess MethodAccessMode3 = 0x0
	// MethodAccessMode3Access defines that access is allowed bit is set.
	MethodAccessMode3Access MethodAccessMode3 = 0x1
	// MethodAccessMode3AuthenticatedRequest defines that authenticated request bit is set.
	MethodAccessMode3AuthenticatedRequest MethodAccessMode3 = 0x4
	// MethodAccessMode3EncryptedRequest defines that encrypted request bit is set.
	MethodAccessMode3EncryptedRequest MethodAccessMode3 = 0x8
	// MethodAccessMode3DigitallySignedRequest defines that digitally signed request bit is set.
	MethodAccessMode3DigitallySignedRequest MethodAccessMode3 = 0x10
	// MethodAccessMode3AuthenticatedResponse defines that authenticated response bit is set.
	MethodAccessMode3AuthenticatedResponse MethodAccessMode3 = 0x20
	// MethodAccessMode3EncryptedResponse defines that encrypted response bit is set.
	MethodAccessMode3EncryptedResponse MethodAccessMode3 = 0x40
	// MethodAccessMode3DigitallySignedResponse defines that digitally signed response bit is set.
	MethodAccessMode3DigitallySignedResponse MethodAccessMode3 = 0x80
)

// MethodAccessMode3Parse converts the given string into a MethodAccessMode3 value.
//
// It returns the corresponding MethodAccessMode3 constant if the string matches
// a known level name, or an error if the input is invalid.
func MethodAccessMode3Parse(value string) (MethodAccessMode3, error) {
	var ret MethodAccessMode3
	var err error
	switch strings.ToUpper(value) {
	case "NOACCESS":
		ret = MethodAccessMode3NoAccess
	case "ACCESS":
		ret = MethodAccessMode3Access
	case "AUTHENTICATEDREQUEST":
		ret = MethodAccessMode3AuthenticatedRequest
	case "ENCRYPTEDREQUEST":
		ret = MethodAccessMode3EncryptedRequest
	case "DIGITALLYSIGNEDREQUEST":
		ret = MethodAccessMode3DigitallySignedRequest
	case "AUTHENTICATEDRESPONSE":
		ret = MethodAccessMode3AuthenticatedResponse
	case "ENCRYPTEDRESPONSE":
		ret = MethodAccessMode3EncryptedResponse
	case "DIGITALLYSIGNEDRESPONSE":
		ret = MethodAccessMode3DigitallySignedResponse
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the MethodAccessMode3.
// It satisfies fmt.Stringer.
func (g MethodAccessMode3) String() string {
	var ret string
	switch g {
	case MethodAccessMode3NoAccess:
		ret = "NOACCESS"
	case MethodAccessMode3Access:
		ret = "ACCESS"
	case MethodAccessMode3AuthenticatedRequest:
		ret = "AUTHENTICATEDREQUEST"
	case MethodAccessMode3EncryptedRequest:
		ret = "ENCRYPTEDREQUEST"
	case MethodAccessMode3DigitallySignedRequest:
		ret = "DIGITALLYSIGNEDREQUEST"
	case MethodAccessMode3AuthenticatedResponse:
		ret = "AUTHENTICATEDRESPONSE"
	case MethodAccessMode3EncryptedResponse:
		ret = "ENCRYPTEDRESPONSE"
	case MethodAccessMode3DigitallySignedResponse:
		ret = "DIGITALLYSIGNEDRESPONSE"
	}
	return ret
}
