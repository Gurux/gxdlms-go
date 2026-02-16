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

// ExceptionServiceError ServiceError enumerates exception service errors.
type ExceptionServiceError int

const (
	// ExceptionServiceErrorNone defines that no exception has occurred.
	ExceptionServiceErrorNone ExceptionServiceError = iota
	// ExceptionServiceErrorOperationNotPossible defines that the operation not possible.
	ExceptionServiceErrorOperationNotPossible
	// ExceptionServiceErrorServiceNotSupported defines that the service not supported.
	ExceptionServiceErrorServiceNotSupported
	// ExceptionServiceErrorOtherReason defines that the other reason has occurred.
	ExceptionServiceErrorOtherReason
	// ExceptionServiceErrorPduTooLong defines that the PDU is too long.
	ExceptionServiceErrorPduTooLong
	// ExceptionServiceErrorDecipheringError defines that the ciphering failed.
	ExceptionServiceErrorDecipheringError
	// ExceptionServiceErrorInvocationCounterError defines that the invocation counter is invalid.
	ExceptionServiceErrorInvocationCounterError
)

// ExceptionServiceErrorParse converts the given string into a ExceptionServiceError value.
//
// It returns the corresponding ExceptionServiceError constant if the string matches
// a known level name, or an error if the input is invalid.
func ExceptionServiceErrorParse(value string) (ExceptionServiceError, error) {
	var ret ExceptionServiceError
	var err error
	switch strings.ToUpper(value) {
	case "NONE":
		ret = ExceptionServiceErrorNone
	case "OPERATIONNOTPOSSIBLE":
		ret = ExceptionServiceErrorOperationNotPossible
	case "SERVICENOTSUPPORTED":
		ret = ExceptionServiceErrorServiceNotSupported
	case "OTHERREASON":
		ret = ExceptionServiceErrorOtherReason
	case "PDUTOOLONG":
		ret = ExceptionServiceErrorPduTooLong
	case "DECIPHERINGERROR":
		ret = ExceptionServiceErrorDecipheringError
	case "INVOCATIONCOUNTERERROR":
		ret = ExceptionServiceErrorInvocationCounterError
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the ExceptionServiceError.
// It satisfies fmt.Stringer.
func (g ExceptionServiceError) String() string {
	var ret string
	switch g {
	case ExceptionServiceErrorNone:
		ret = "NONE"
	case ExceptionServiceErrorOperationNotPossible:
		ret = "OPERATIONNOTPOSSIBLE"
	case ExceptionServiceErrorServiceNotSupported:
		ret = "SERVICENOTSUPPORTED"
	case ExceptionServiceErrorOtherReason:
		ret = "OTHERREASON"
	case ExceptionServiceErrorPduTooLong:
		ret = "PDUTOOLONG"
	case ExceptionServiceErrorDecipheringError:
		ret = "DECIPHERINGERROR"
	case ExceptionServiceErrorInvocationCounterError:
		ret = "INVOCATIONCOUNTERERROR"
	}
	return ret
}
