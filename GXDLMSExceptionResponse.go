package dlms

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

	"github.com/Gurux/gxdlms-go/enums"
)

// DLMS specific exception response.
// https://www.gurux.fi/Gurux.DLMS.ErrorCodes
type GXDLMSExceptionResponse struct {
	exceptionStateError   enums.ExceptionStateError
	exceptionServiceError enums.ExceptionServiceError
	value                 any
	message               string
}

func (g *GXDLMSExceptionResponse) ExceptionStateError() enums.ExceptionStateError {
	return g.exceptionStateError
}
func (g *GXDLMSExceptionResponse) ExceptionServiceError() enums.ExceptionServiceError {
	return g.exceptionServiceError
}
func (g *GXDLMSExceptionResponse) Value() any {
	return g.value
}

func getStateError(stateError enums.ExceptionStateError) string {
	switch stateError {
	case enums.ExceptionStateErrorServiceNotAllowed:
		return "Service not allowed"
	case enums.ExceptionStateErrorServiceUnknown:
		return "Service unknown"
	}
	return ""
}

func getServiceError(serviceError enums.ExceptionServiceError, value *any) string {
	switch serviceError {
	case enums.ExceptionServiceErrorOperationNotPossible:
		return "Operation not possible"
	case enums.ExceptionServiceErrorOtherReason:
		return "Other reason"
	case enums.ExceptionServiceErrorServiceNotSupported:
		return "Service not supported"
	case enums.ExceptionServiceErrorPduTooLong:
		return "PDU is too long"
	case enums.ExceptionServiceErrorDecipheringError:
		return "Deciphering failed"
	case enums.ExceptionServiceErrorInvocationCounterError:
		return "Invocation counter is invalid. Expected value is " + fmt.Sprint(value)
	}
	return ""
}

func NewGXDLMSExceptionResponse(stateError enums.ExceptionStateError, type_ enums.ExceptionServiceError, value any) *GXDLMSExceptionResponse {
	return &GXDLMSExceptionResponse{
		exceptionStateError:   stateError,
		exceptionServiceError: type_,
		value:                 value,
		message:               "https://www.gurux.fi/Gurux.DLMS.ErrorCodes",
	}
}

// Error implements the error interface.
func (e *GXDLMSExceptionResponse) Error() string {
	return fmt.Sprintf(
		"Exception response %s exception. %s",
		getStateError(e.exceptionStateError),
		getServiceError(e.exceptionServiceError, &e.value))
}
