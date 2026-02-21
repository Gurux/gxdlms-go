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

// AccessMode3 enumerates allowed access rights.
type AccessMode3 int

const (
	// AccessMode3NoAccess defines that there is no access.
	AccessMode3NoAccess AccessMode3 = iota
	// AccessMode3Read defines that the client is allowed only reading from the server.
	//  This is used in version 1.
	AccessMode3Read = 1
	// AccessMode3Write defines that the client is allowed only writing to the server.
	AccessMode3Write = 2
	// AccessMode3AuthenticatedRequest defines that the request messages are authenticated.
	AccessMode3AuthenticatedRequest = 4
	// AccessMode3EncryptedRequest defines that the request messages are encrypted.
	AccessMode3EncryptedRequest = 8
	// AccessMode3DigitallySignedRequest defines that the request messages are digitally signed.
	AccessMode3DigitallySignedRequest = 16
	// AccessMode3AuthenticatedResponse defines that the response messages are authenticated.
	AccessMode3AuthenticatedResponse = 32
	// AccessMode3EncryptedResponse defines that the response messages are encrypted.
	AccessMode3EncryptedResponse = 64
	// AccessMode3DigitallySignedResponse defines that the response messages are digitally signed.
	AccessMode3DigitallySignedResponse = 128
)

// AccessMode3Parse converts the given string into a AccessMode3 value.
//
// It returns the corresponding AccessMode3 constant if the string matches
// a known level name, or an error if the input is invalid.
func AccessMode3Parse(value string) (AccessMode3, error) {
	var ret AccessMode3
	var err error
	switch {
	case strings.EqualFold(value, "NoAccess"):
		ret = AccessMode3NoAccess
	case strings.EqualFold(value, "Read"):
		ret = AccessMode3Read
	case strings.EqualFold(value, "Write"):
		ret = AccessMode3Write
	case strings.EqualFold(value, "AuthenticatedRequest"):
		ret = AccessMode3AuthenticatedRequest
	case strings.EqualFold(value, "EncryptedRequest"):
		ret = AccessMode3EncryptedRequest
	case strings.EqualFold(value, "DigitallySignedRequest"):
		ret = AccessMode3DigitallySignedRequest
	case strings.EqualFold(value, "AuthenticatedResponse"):
		ret = AccessMode3AuthenticatedResponse
	case strings.EqualFold(value, "EncryptedResponse"):
		ret = AccessMode3EncryptedResponse
	case strings.EqualFold(value, "DigitallySignedResponse"):
		ret = AccessMode3DigitallySignedResponse
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the AccessMode3.
// It satisfies fmt.Stringer.
func (g AccessMode3) String() string {
	var ret string
	switch g {
	case AccessMode3NoAccess:
		ret = "NoAccess"
	case AccessMode3Read:
		ret = "Read"
	case AccessMode3Write:
		ret = "Write"
	case AccessMode3AuthenticatedRequest:
		ret = "AuthenticatedRequest"
	case AccessMode3EncryptedRequest:
		ret = "EncryptedRequest"
	case AccessMode3DigitallySignedRequest:
		ret = "DigitallySignedRequest"
	case AccessMode3AuthenticatedResponse:
		ret = "AuthenticatedResponse"
	case AccessMode3EncryptedResponse:
		ret = "EncryptedResponse"
	case AccessMode3DigitallySignedResponse:
		ret = "DigitallySignedResponse"
	}
	return ret
}

// AllAccessMode3 returns a slice containing all defined AccessMode3 values.
func AllAccessMode3() []AccessMode3 {
	return []AccessMode3{
		AccessMode3NoAccess,
		AccessMode3Read,
		AccessMode3Write,
		AccessMode3AuthenticatedRequest,
		AccessMode3EncryptedRequest,
		AccessMode3DigitallySignedRequest,
		AccessMode3AuthenticatedResponse,
		AccessMode3EncryptedResponse,
		AccessMode3DigitallySignedResponse,
	}
}
