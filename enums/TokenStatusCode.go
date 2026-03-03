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

// TokenStatusCode enumerates token status codes.
type TokenStatusCode int

const (
	// TokenStatusCodeFormatOk defines that the token is OK.
	TokenStatusCodeFormatOk TokenStatusCode = iota
	// TokenStatusCodeAuthenticationOk defines that the token authentication result is OK.
	TokenStatusCodeAuthenticationOk
	// TokenStatusCodeValidationOk defines that the token validation result is OK.
	TokenStatusCodeValidationOk
	// TokenStatusCodeTokenExecutionOk defines that the token execution result is OK.
	TokenStatusCodeTokenExecutionOk
	// TokenStatusCodeTokenFormatFailure defines that the there is an error in token format.
	TokenStatusCodeTokenFormatFailure
	// TokenStatusCodeAuthenticationFailure defines token authentication failure.
	TokenStatusCodeAuthenticationFailure
	// TokenStatusCodeValidationResultFailure defines defines token validation failure.
	TokenStatusCodeValidationResultFailure
	// TokenStatusCodeTokenExecutionResultFailure defines token execution result failure.
	TokenStatusCodeTokenExecutionResultFailure
	// TokenStatusCodeTokenReceived defines that the the token is received and not yet processed.
	TokenStatusCodeTokenReceived
)

// TokenStatusCodeParse converts the given string into a TokenStatusCode value.
//
// It returns the corresponding TokenStatusCode constant if the string matches
// a known level name, or an error if the input is invalid.
func TokenStatusCodeParse(value string) (TokenStatusCode, error) {
	var ret TokenStatusCode
	var err error
	switch {
	case strings.EqualFold(value, "FormatOk"):
		ret = TokenStatusCodeFormatOk
	case strings.EqualFold(value, "AuthenticationOk"):
		ret = TokenStatusCodeAuthenticationOk
	case strings.EqualFold(value, "ValidationOk"):
		ret = TokenStatusCodeValidationOk
	case strings.EqualFold(value, "TokenExecutionOk"):
		ret = TokenStatusCodeTokenExecutionOk
	case strings.EqualFold(value, "TokenFormatFailure"):
		ret = TokenStatusCodeTokenFormatFailure
	case strings.EqualFold(value, "AuthenticationFailure"):
		ret = TokenStatusCodeAuthenticationFailure
	case strings.EqualFold(value, "ValidationResultFailure"):
		ret = TokenStatusCodeValidationResultFailure
	case strings.EqualFold(value, "TokenExecutionResultFailure"):
		ret = TokenStatusCodeTokenExecutionResultFailure
	case strings.EqualFold(value, "TokenReceived"):
		ret = TokenStatusCodeTokenReceived
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the TokenStatusCode.
// It satisfies fmt.Stringer.
func (g TokenStatusCode) String() string {
	var ret string
	switch g {
	case TokenStatusCodeFormatOk:
		ret = "FormatOk"
	case TokenStatusCodeAuthenticationOk:
		ret = "AuthenticationOk"
	case TokenStatusCodeValidationOk:
		ret = "ValidationOk"
	case TokenStatusCodeTokenExecutionOk:
		ret = "TokenExecutionOk"
	case TokenStatusCodeTokenFormatFailure:
		ret = "TokenFormatFailure"
	case TokenStatusCodeAuthenticationFailure:
		ret = "AuthenticationFailure"
	case TokenStatusCodeValidationResultFailure:
		ret = "ValidationResultFailure"
	case TokenStatusCodeTokenExecutionResultFailure:
		ret = "TokenExecutionResultFailure"
	case TokenStatusCodeTokenReceived:
		ret = "TokenReceived"
	}
	return ret
}

// AllTokenStatusCode returns a slice containing all defined TokenStatusCode values.
func AllTokenStatusCode() []TokenStatusCode {
	return []TokenStatusCode{
		TokenStatusCodeFormatOk,
		TokenStatusCodeAuthenticationOk,
		TokenStatusCodeValidationOk,
		TokenStatusCodeTokenExecutionOk,
		TokenStatusCodeTokenFormatFailure,
		TokenStatusCodeAuthenticationFailure,
		TokenStatusCodeValidationResultFailure,
		TokenStatusCodeTokenExecutionResultFailure,
		TokenStatusCodeTokenReceived,
	}
}
