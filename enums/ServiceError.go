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

// ServiceError ServiceError enumerates service errors.
type ServiceError int

const (
	// ServiceErrorApplicationReference defines that the application error has occurred.
	ServiceErrorApplicationReference ServiceError = iota
	// ServiceErrorHardwareResource defines that the hardware error.
	ServiceErrorHardwareResource
	// ServiceErrorVdeStateError defines that the VDE state error.
	ServiceErrorVdeStateError
	// ServiceErrorService defines that the service error.
	ServiceErrorService
	// ServiceErrorDefinition defines that the definition error.
	ServiceErrorDefinition
	// ServiceErrorAccess defines that the access error.
	ServiceErrorAccess
	// ServiceErrorInitiate defines that the initiate error.
	ServiceErrorInitiate
	// ServiceErrorLoadDataSet defines that the loadDataSet error.
	ServiceErrorLoadDataSet
	// ServiceErrorTask defines that the task error.
	ServiceErrorTask
	// ServiceErrorOtherError defines that the other error has occurred. The other error describes manufacturer specific error code.
	ServiceErrorOtherError
)

// ServiceErrorParse converts the given string into a ServiceError value.
//
// It returns the corresponding ServiceError constant if the string matches
// a known level name, or an error if the input is invalid.
func ServiceErrorParse(value string) (ServiceError, error) {
	var ret ServiceError
	var err error
	switch strings.ToUpper(value) {
	case "APPLICATIONREFERENCE":
		ret = ServiceErrorApplicationReference
	case "HARDWARERESOURCE":
		ret = ServiceErrorHardwareResource
	case "VDESTATEERROR":
		ret = ServiceErrorVdeStateError
	case "SERVICE":
		ret = ServiceErrorService
	case "DEFINITION":
		ret = ServiceErrorDefinition
	case "ACCESS":
		ret = ServiceErrorAccess
	case "INITIATE":
		ret = ServiceErrorInitiate
	case "LOADDATASET":
		ret = ServiceErrorLoadDataSet
	case "TASK":
		ret = ServiceErrorTask
	case "OTHERERROR":
		ret = ServiceErrorOtherError
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the ServiceError.
// It satisfies fmt.Stringer.
func (g ServiceError) String() string {
	var ret string
	switch g {
	case ServiceErrorApplicationReference:
		ret = "APPLICATIONREFERENCE"
	case ServiceErrorHardwareResource:
		ret = "HARDWARERESOURCE"
	case ServiceErrorVdeStateError:
		ret = "VDESTATEERROR"
	case ServiceErrorService:
		ret = "SERVICE"
	case ServiceErrorDefinition:
		ret = "DEFINITION"
	case ServiceErrorAccess:
		ret = "ACCESS"
	case ServiceErrorInitiate:
		ret = "INITIATE"
	case ServiceErrorLoadDataSet:
		ret = "LOADDATASET"
	case ServiceErrorTask:
		ret = "TASK"
	case ServiceErrorOtherError:
		ret = "OTHERERROR"
	}
	return ret
}
