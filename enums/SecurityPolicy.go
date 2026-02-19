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
	switch {
	case strings.EqualFold(value, "None"):
		ret = SecurityPolicyNone
	case strings.EqualFold(value, "Authenticated"):
		ret = SecurityPolicyAuthenticated
	case strings.EqualFold(value, "Encrypted"):
		ret = SecurityPolicyEncrypted
	case strings.EqualFold(value, "AuthenticatedEncrypted"):
		ret = SecurityPolicyAuthenticatedEncrypted
	case strings.EqualFold(value, "AuthenticatedRequest"):
		ret = SecurityPolicyAuthenticatedRequest
	case strings.EqualFold(value, "EncryptedRequest"):
		ret = SecurityPolicyEncryptedRequest
	case strings.EqualFold(value, "DigitallySignedRequest"):
		ret = SecurityPolicyDigitallySignedRequest
	case strings.EqualFold(value, "AuthenticatedResponse"):
		ret = SecurityPolicyAuthenticatedResponse
	case strings.EqualFold(value, "EncryptedResponse"):
		ret = SecurityPolicyEncryptedResponse
	case strings.EqualFold(value, "DigitallySignedResponse"):
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
		ret = "None"
	case SecurityPolicyAuthenticated:
		ret = "Authenticated"
	case SecurityPolicyEncrypted:
		ret = "Encrypted"
	case SecurityPolicyAuthenticatedEncrypted:
		ret = "AuthenticatedEncrypted"
	case SecurityPolicyAuthenticatedRequest:
		ret = "AuthenticatedRequest"
	case SecurityPolicyEncryptedRequest:
		ret = "EncryptedRequest"
	case SecurityPolicyDigitallySignedRequest:
		ret = "DigitallySignedRequest"
	case SecurityPolicyAuthenticatedResponse:
		ret = "AuthenticatedResponse"
	case SecurityPolicyEncryptedResponse:
		ret = "EncryptedResponse"
	case SecurityPolicyDigitallySignedResponse:
		ret = "DigitallySignedResponse"
	}
	return ret
}

// AllSecurityPolicy returns a slice containing all defined SecurityPolicy values.
func AllSecurityPolicy() []SecurityPolicy {
	return []SecurityPolicy{
		SecurityPolicyNone,
		SecurityPolicyAuthenticated,
		SecurityPolicyEncrypted,
		SecurityPolicyAuthenticatedEncrypted,
		SecurityPolicyAuthenticatedRequest,
		SecurityPolicyEncryptedRequest,
		SecurityPolicyDigitallySignedRequest,
		SecurityPolicyAuthenticatedResponse,
		SecurityPolicyEncryptedResponse,
		SecurityPolicyDigitallySignedResponse,
	}
}
