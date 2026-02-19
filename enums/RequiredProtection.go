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

// RequiredProtection enumerates the access modes for data protection.
type RequiredProtection int

const (
	// RequiredProtectionAuthenticatedRequest defines that request messages are authenticated bit is set.
	RequiredProtectionAuthenticatedRequest RequiredProtection = 4
	// RequiredProtectionEncryptedRequest defines that request messages are encrypted bit is set.
	RequiredProtectionEncryptedRequest RequiredProtection = 8
	// RequiredProtectionDigitallySignedRequest defines that request messages are digitally signed bit is set.
	RequiredProtectionDigitallySignedRequest RequiredProtection = 16
	// RequiredProtectionAuthenticatedResponse defines that response messages are authenticated bit is set.
	RequiredProtectionAuthenticatedResponse RequiredProtection = 32
	// RequiredProtectionEncryptedResponse defines that response messages are encrypted bit is set.
	RequiredProtectionEncryptedResponse RequiredProtection = 64
	// RequiredProtectionDigitallySignedResponse defines that response messages are digitally signed bit is set.
	RequiredProtectionDigitallySignedResponse RequiredProtection = 128
)

// RequiredProtectionParse converts the given string into a RequiredProtection value.
//
// It returns the corresponding RequiredProtection constant if the string matches
// a known level name, or an error if the input is invalid.
func RequiredProtectionParse(value string) (RequiredProtection, error) {
	var ret RequiredProtection
	var err error
	switch {
	case strings.EqualFold(value, "AuthenticatedRequest"):
		ret = RequiredProtectionAuthenticatedRequest
	case strings.EqualFold(value, "EncryptedRequest"):
		ret = RequiredProtectionEncryptedRequest
	case strings.EqualFold(value, "DigitallySignedRequest"):
		ret = RequiredProtectionDigitallySignedRequest
	case strings.EqualFold(value, "AuthenticatedResponse"):
		ret = RequiredProtectionAuthenticatedResponse
	case strings.EqualFold(value, "EncryptedResponse"):
		ret = RequiredProtectionEncryptedResponse
	case strings.EqualFold(value, "DigitallySignedResponse"):
		ret = RequiredProtectionDigitallySignedResponse
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the RequiredProtection.
// It satisfies fmt.Stringer.
func (g RequiredProtection) String() string {
	var ret string
	switch g {
	case RequiredProtectionAuthenticatedRequest:
		ret = "AuthenticatedRequest"
	case RequiredProtectionEncryptedRequest:
		ret = "EncryptedRequest"
	case RequiredProtectionDigitallySignedRequest:
		ret = "DigitallySignedRequest"
	case RequiredProtectionAuthenticatedResponse:
		ret = "AuthenticatedResponse"
	case RequiredProtectionEncryptedResponse:
		ret = "EncryptedResponse"
	case RequiredProtectionDigitallySignedResponse:
		ret = "DigitallySignedResponse"
	}
	return ret
}

// AllRequiredProtection returns a slice containing all defined RequiredProtection values.
func AllRequiredProtection() []RequiredProtection {
	return []RequiredProtection{
		RequiredProtectionAuthenticatedRequest,
		RequiredProtectionEncryptedRequest,
		RequiredProtectionDigitallySignedRequest,
		RequiredProtectionAuthenticatedResponse,
		RequiredProtectionEncryptedResponse,
		RequiredProtectionDigitallySignedResponse,
	}
}
