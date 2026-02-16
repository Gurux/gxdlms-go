package constants

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

// CoAP client error codes.
type CoAPClientError int

const (
	// CoAPClientErrorBadRequest defines that the // Bad Request.
	CoAPClientErrorBadRequest CoAPClientError = iota
	// CoAPClientErrorUnauthorized defines that the // Unauthorized.
	CoAPClientErrorUnauthorized
	// CoAPClientErrorBadOption defines that the // Bad Option.
	CoAPClientErrorBadOption
	// CoAPClientErrorForbidden defines that the // Forbidden.
	CoAPClientErrorForbidden
	// CoAPClientErrorNotFound defines that the // Not Found.
	CoAPClientErrorNotFound
	// CoAPClientErrorMethodNotAllowed defines that the // Method Not Allowed.
	CoAPClientErrorMethodNotAllowed
	// CoAPClientErrorNotAcceptable defines that the // Not Acceptable.
	CoAPClientErrorNotAcceptable
	// CoAPClientErrorRequestEntityIncomplete defines that the // Request Entity Incomplete.
	CoAPClientErrorRequestEntityIncomplete CoAPClientError = 8
	// CoAPClientErrorConflict defines that the // Conflict.
	CoAPClientErrorConflict CoAPClientError = 9
	// CoAPClientErrorPreconditionFailed defines that the // Precondition Failed.
	CoAPClientErrorPreconditionFailed CoAPClientError = 12
	// CoAPClientErrorRequestEntityTooLarge defines that the // Request Entity Too Large.
	CoAPClientErrorRequestEntityTooLarge CoAPClientError = 13
	// CoAPClientErrorUnsupportedContentFormat defines that the // Unsupported Content-Format.
	CoAPClientErrorUnsupportedContentFormat CoAPClientError = 15
)

// CoAPClientErrorParse converts the given string into a CoAPClientError value.
//
// It returns the corresponding CoAPClientError constant if the string matches
// a known level name, or an error if the input is invalid.
func CoAPClientErrorParse(value string) (CoAPClientError, error) {
	var ret CoAPClientError
	var err error
	switch strings.ToUpper(value) {
	case "BADREQUEST":
		ret = CoAPClientErrorBadRequest
	case "UNAUTHORIZED":
		ret = CoAPClientErrorUnauthorized
	case "BADOPTION":
		ret = CoAPClientErrorBadOption
	case "FORBIDDEN":
		ret = CoAPClientErrorForbidden
	case "NOTFOUND":
		ret = CoAPClientErrorNotFound
	case "METHODNOTALLOWED":
		ret = CoAPClientErrorMethodNotAllowed
	case "NOTACCEPTABLE":
		ret = CoAPClientErrorNotAcceptable
	case "REQUESTENTITYINCOMPLETE":
		ret = CoAPClientErrorRequestEntityIncomplete
	case "CONFLICT":
		ret = CoAPClientErrorConflict
	case "PRECONDITIONFAILED":
		ret = CoAPClientErrorPreconditionFailed
	case "REQUESTENTITYTOOLARGE":
		ret = CoAPClientErrorRequestEntityTooLarge
	case "UNSUPPORTEDCONTENTFORMAT":
		ret = CoAPClientErrorUnsupportedContentFormat
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the CoAPClientError.
// It satisfies fmt.Stringer.
func (g CoAPClientError) String() string {
	var ret string
	switch g {
	case CoAPClientErrorBadRequest:
		ret = "BADREQUEST"
	case CoAPClientErrorUnauthorized:
		ret = "UNAUTHORIZED"
	case CoAPClientErrorBadOption:
		ret = "BADOPTION"
	case CoAPClientErrorForbidden:
		ret = "FORBIDDEN"
	case CoAPClientErrorNotFound:
		ret = "NOTFOUND"
	case CoAPClientErrorMethodNotAllowed:
		ret = "METHODNOTALLOWED"
	case CoAPClientErrorNotAcceptable:
		ret = "NOTACCEPTABLE"
	case CoAPClientErrorRequestEntityIncomplete:
		ret = "REQUESTENTITYINCOMPLETE"
	case CoAPClientErrorConflict:
		ret = "CONFLICT"
	case CoAPClientErrorPreconditionFailed:
		ret = "PRECONDITIONFAILED"
	case CoAPClientErrorRequestEntityTooLarge:
		ret = "REQUESTENTITYTOOLARGE"
	case CoAPClientErrorUnsupportedContentFormat:
		ret = "UNSUPPORTEDCONTENTFORMAT"
	}
	return ret
}
