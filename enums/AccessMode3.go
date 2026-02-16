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

// AccessMode3 The AccessMode3 enumerates the access modes for Logical Name Association version 3.
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
	switch strings.ToUpper(value) {
	case "NOACCESS":
		ret = AccessMode3NoAccess
	case "READ":
		ret = AccessMode3Read
	case "WRITE":
		ret = AccessMode3Write
	case "AUTHENTICATEDREQUEST":
		ret = AccessMode3AuthenticatedRequest
	case "ENCRYPTEDREQUEST":
		ret = AccessMode3EncryptedRequest
	case "DIGITALLYSIGNEDREQUEST":
		ret = AccessMode3DigitallySignedRequest
	case "AUTHENTICATEDRESPONSE":
		ret = AccessMode3AuthenticatedResponse
	case "ENCRYPTEDRESPONSE":
		ret = AccessMode3EncryptedResponse
	case "DIGITALLYSIGNEDRESPONSE":
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
		ret = "NOACCESS"
	case AccessMode3Read:
		ret = "READ"
	case AccessMode3Write:
		ret = "WRITE"
	case AccessMode3AuthenticatedRequest:
		ret = "AUTHENTICATEDREQUEST"
	case AccessMode3EncryptedRequest:
		ret = "ENCRYPTEDREQUEST"
	case AccessMode3DigitallySignedRequest:
		ret = "DIGITALLYSIGNEDREQUEST"
	case AccessMode3AuthenticatedResponse:
		ret = "AUTHENTICATEDRESPONSE"
	case AccessMode3EncryptedResponse:
		ret = "ENCRYPTEDRESPONSE"
	case AccessMode3DigitallySignedResponse:
		ret = "DIGITALLYSIGNEDRESPONSE"
	}
	return ret
}
