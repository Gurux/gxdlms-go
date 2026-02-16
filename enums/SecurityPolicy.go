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

package enums

import (
	"fmt"
	"strings"

	"github.com/Gurux/gxcommon-go"
)

// SecurityPolicy :  Enforces authentication and/or encryption algorithm provided with security suite version 0.
type SecurityPolicy int

const (
	// SecurityPolicyNone defines that security is not used bit is set.
	SecurityPolicyNone SecurityPolicy = iota
	// SecurityPolicyAuthenticated defines that all messages are authenticated using Security Suite 0 bit is set.
	SecurityPolicyAuthenticated SecurityPolicy = 0x1
	// SecurityPolicyEncrypted defines that all messages are encrypted using Security Suite 0 bit is set.
	SecurityPolicyEncrypted SecurityPolicy = 0x2
	// SecurityPolicyAuthenticatedEncrypted defines that all messages are authenticated and encrypted using Security Suite 0 bit is set.
	SecurityPolicyAuthenticatedEncrypted SecurityPolicy = 0x3
	// SecurityPolicyAuthenticatedRequest defines that request is authenticated bit is set.
	SecurityPolicyAuthenticatedRequest SecurityPolicy = 0x4
	// SecurityPolicyEncryptedRequest defines that request is encrypted bit is set.
	SecurityPolicyEncryptedRequest SecurityPolicy = 0x8
	// SecurityPolicyDigitallySignedRequest defines that request is digitally signed bit is set.
	SecurityPolicyDigitallySignedRequest SecurityPolicy = 0x10
	// SecurityPolicyAuthenticatedResponse defines that response is authenticated bit is set.
	SecurityPolicyAuthenticatedResponse SecurityPolicy = 0x20
	// SecurityPolicyEncryptedResponse defines that response is encrypted bit is set.
	SecurityPolicyEncryptedResponse SecurityPolicy = 0x40
	// SecurityPolicyDigitallySignedResponse defines that response is digitally signed bit is set.
	SecurityPolicyDigitallySignedResponse SecurityPolicy = 0x80
)

// SecurityPolicyParse converts the given string into a SecurityPolicy value.
//
// It returns the corresponding SecurityPolicy constant if the string matches
// a known level name, or an error if the input is invalid.
func SecurityPolicyParse(value string) (SecurityPolicy, error) {
	var ret SecurityPolicy
	var err error
	switch strings.ToUpper(value) {
	case "NONE":
		ret = SecurityPolicyNone
	case "AUTHENTICATED":
		ret = SecurityPolicyAuthenticated
	case "ENCRYPTED":
		ret = SecurityPolicyEncrypted
	case "AUTHENTICATEDENCRYPTED":
		ret = SecurityPolicyAuthenticatedEncrypted
	case "AUTHENTICATEDREQUEST":
		ret = SecurityPolicyAuthenticatedRequest
	case "ENCRYPTEDREQUEST":
		ret = SecurityPolicyEncryptedRequest
	case "DIGITALLYSIGNEDREQUEST":
		ret = SecurityPolicyDigitallySignedRequest
	case "AUTHENTICATEDRESPONSE":
		ret = SecurityPolicyAuthenticatedResponse
	case "ENCRYPTEDRESPONSE":
		ret = SecurityPolicyEncryptedResponse
	case "DIGITALLYSIGNEDRESPONSE":
		ret = SecurityPolicyDigitallySignedResponse
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the SecurityPolicy.
// It satisfies fmt.Stringer.
func (g SecurityPolicy) String() string {
	var ret string
	switch g {
	case SecurityPolicyNone:
		ret = "NONE"
	case SecurityPolicyAuthenticated:
		ret = "AUTHENTICATED"
	case SecurityPolicyEncrypted:
		ret = "ENCRYPTED"
	case SecurityPolicyAuthenticatedEncrypted:
		ret = "AUTHENTICATEDENCRYPTED"
	case SecurityPolicyAuthenticatedRequest:
		ret = "AUTHENTICATEDREQUEST"
	case SecurityPolicyEncryptedRequest:
		ret = "ENCRYPTEDREQUEST"
	case SecurityPolicyDigitallySignedRequest:
		ret = "DIGITALLYSIGNEDREQUEST"
	case SecurityPolicyAuthenticatedResponse:
		ret = "AUTHENTICATEDRESPONSE"
	case SecurityPolicyEncryptedResponse:
		ret = "ENCRYPTEDRESPONSE"
	case SecurityPolicyDigitallySignedResponse:
		ret = "DIGITALLYSIGNEDRESPONSE"
	}
	return ret
}
