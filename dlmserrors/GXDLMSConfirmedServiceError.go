package dlmserrors

// --------------------------------------------------------------------------
//
//	Gurux Ltd
//
// Filename:        $HeadURL$
//
// Version:         $Revision$,
//
//	$Date$
//	$Author$
//
// # Copyright (c) Gurux Ltd
//
// ---------------------------------------------------------------------------
//
//	DESCRIPTION
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
// ---------------------------------------------------------------------------

import (
	"fmt"

	"github.com/Gurux/gxdlms-go/enums"
)

// GXDLMSConfirmedServiceError corresponds to the C# exception class.
type GXDLMSConfirmedServiceError struct {
	ConfirmedServiceError enums.ConfirmedServiceError
	ServiceError          enums.ServiceError
	ServiceErrorValue     uint32
	HelpLink              string
}

// NewGXDLMSConfirmedServiceError creates a new instance of GXDLMSConfirmedServiceError.
func NewGXDLMSConfirmedServiceError(
	service enums.ConfirmedServiceError,
	errType enums.ServiceError,
	value byte,
) error {

	return &GXDLMSConfirmedServiceError{
		ConfirmedServiceError: service,
		ServiceError:          errType,
		ServiceErrorValue:     uint32(value),
		HelpLink:              "https://www.gurux.fi/Gurux.DLMS.ErrorCodes",
	}
}

// Error implements the error interface.
func (e *GXDLMSConfirmedServiceError) Error() string {
	return fmt.Sprintf(
		"ServiceError %s exception. %s %s",
		getConfirmedServiceError(e.ConfirmedServiceError),
		getServiceError(e.ServiceError),
		getServiceErrorValue(e.ServiceError, uint8(e.ServiceErrorValue)),
	)
}

func getConfirmedServiceError(stateError enums.ConfirmedServiceError) string {
	switch stateError {
	case enums.ConfirmedServiceErrorInitiateError:
		return "Initiate Error"
	case enums.ConfirmedServiceErrorRead:
		return "Read"
	case enums.ConfirmedServiceErrorWrite:
		return "Write"
	default:
		return ""
	}
}

func getServiceError(err enums.ServiceError) string {
	switch err {
	case enums.ServiceErrorApplicationReference:
		return "ApplicationReference"
	case enums.ServiceErrorHardwareResource:
		return "HardwareResource"
	case enums.ServiceErrorVdeStateError:
		return "VdeStateError"
	case enums.ServiceErrorService:
		return "Service"
	case enums.ServiceErrorDefinition:
		return "Definition"
	case enums.ServiceErrorAccess:
		return "Access"
	case enums.ServiceErrorInitiate:
		return "Initiate"
	case enums.ServiceErrorLoadDataSet:
		return "Load data set"
	case enums.ServiceErrorTask:
		return "Task"
	case enums.ServiceErrorOtherError:
		return "Other Error"
	default:
		return ""
	}
}

func getServiceErrorValue(err enums.ServiceError, value byte) string {
	switch err {
	case enums.ServiceErrorApplicationReference:
		return enums.ApplicationReference(value).String()
	case enums.ServiceErrorHardwareResource:
		return enums.HardwareResource(value).String()
	case enums.ServiceErrorVdeStateError:
		return enums.VdeStateError(value).String()
	case enums.ServiceErrorService:
		return enums.Service(value).String()
	case enums.ServiceErrorDefinition:
		return enums.Definition(value).String()
	case enums.ServiceErrorAccess:
		return enums.Access(value).String()
	case enums.ServiceErrorInitiate:
		return enums.Initiate(value).String()
	case enums.ServiceErrorLoadDataSet:
		return enums.LoadDataSet(value).String()
	case enums.ServiceErrorTask:
		return enums.Task(value).String()
	case enums.ServiceErrorOtherError:
		return fmt.Sprintf("%d", value)
	default:
		return ""
	}
}
