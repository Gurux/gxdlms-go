package internal

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

// PduType APDU types.
type PduType int

const (
	// PduTypeProtocolVersion defines that the iMPLICIT BIT STRING {version1 (0)} DEFAULT {version1}
	PduTypeProtocolVersion PduType = iota
	// PduTypeApplicationContextName defines that the application-context-name
	PduTypeApplicationContextName
	// PduTypeCalledApTitle defines that the aP-title OPTIONAL
	PduTypeCalledApTitle
	// PduTypeCalledAeQualifier defines that the aE-qualifier OPTIONAL.
	PduTypeCalledAeQualifier
	// PduTypeCalledApInvocationId defines that the aP-invocation-identifier OPTIONAL.
	PduTypeCalledApInvocationId
	// PduTypeCalledAeInvocationId defines that the aE-invocation-identifier OPTIONAL
	PduTypeCalledAeInvocationId
	// PduTypeCallingApTitle defines that the aP-title OPTIONAL
	PduTypeCallingApTitle
	// PduTypeCallingAeQualifier defines that the aE-qualifier OPTIONAL
	PduTypeCallingAeQualifier
	// PduTypeCallingApInvocationId defines that the aP-invocation-identifier OPTIONAL
	PduTypeCallingApInvocationId
	// PduTypeCallingAeInvocationId defines that the aE-invocation-identifier OPTIONAL
	PduTypeCallingAeInvocationId
	// PduTypeSenderAcseRequirements defines that the the following field shall not be present if only the kernel is used.
	PduTypeSenderAcseRequirements
	// PduTypeMechanismName defines that the the following field shall only be present if the authentication functional unit is selected.
	PduTypeMechanismName
	// PduTypeCallingAuthenticationValue defines that the the following field shall only be present if the authentication functional unit is selected.
	PduTypeCallingAuthenticationValue
	// PduTypeImplementationInformation defines that the implementation-data.
	PduTypeImplementationInformation PduType = 29
	// PduTypeUserInformation defines that the association-information OPTIONAL
	PduTypeUserInformation PduType = 30
)

// PduTypeParse converts the given string into a PduType value.
//
// It returns the corresponding PduType constant if the string matches
// a known level name, or an error if the input is invalid.
func PduTypeParse(value string) (PduType, error) {
	var ret PduType
	var err error
	switch strings.ToUpper(value) {
	case "PROTOCOLVERSION":
		ret = PduTypeProtocolVersion
	case "APPLICATIONCONTEXTNAME":
		ret = PduTypeApplicationContextName
	case "CALLEDAPTITLE":
		ret = PduTypeCalledApTitle
	case "CALLEDAEQUALIFIER":
		ret = PduTypeCalledAeQualifier
	case "CALLEDAPINVOCATIONID":
		ret = PduTypeCalledApInvocationId
	case "CALLEDAEINVOCATIONID":
		ret = PduTypeCalledAeInvocationId
	case "CALLINGAPTITLE":
		ret = PduTypeCallingApTitle
	case "CALLINGAEQUALIFIER":
		ret = PduTypeCallingAeQualifier
	case "CALLINGAPINVOCATIONID":
		ret = PduTypeCallingApInvocationId
	case "CALLINGAEINVOCATIONID":
		ret = PduTypeCallingAeInvocationId
	case "SENDERACSEREQUIREMENTS":
		ret = PduTypeSenderAcseRequirements
	case "MECHANISMNAME":
		ret = PduTypeMechanismName
	case "CALLINGAUTHENTICATIONVALUE":
		ret = PduTypeCallingAuthenticationValue
	case "IMPLEMENTATIONINFORMATION":
		ret = PduTypeImplementationInformation
	case "USERINFORMATION":
		ret = PduTypeUserInformation
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the PduType.
// It satisfies fmt.Stringer.
func (g PduType) String() string {
	var ret string
	switch g {
	case PduTypeProtocolVersion:
		ret = "PROTOCOLVERSION"
	case PduTypeApplicationContextName:
		ret = "APPLICATIONCONTEXTNAME"
	case PduTypeCalledApTitle:
		ret = "CALLEDAPTITLE"
	case PduTypeCalledAeQualifier:
		ret = "CALLEDAEQUALIFIER"
	case PduTypeCalledApInvocationId:
		ret = "CALLEDAPINVOCATIONID"
	case PduTypeCalledAeInvocationId:
		ret = "CALLEDAEINVOCATIONID"
	case PduTypeCallingApTitle:
		ret = "CALLINGAPTITLE"
	case PduTypeCallingAeQualifier:
		ret = "CALLINGAEQUALIFIER"
	case PduTypeCallingApInvocationId:
		ret = "CALLINGAPINVOCATIONID"
	case PduTypeCallingAeInvocationId:
		ret = "CALLINGAEINVOCATIONID"
	case PduTypeSenderAcseRequirements:
		ret = "SENDERACSEREQUIREMENTS"
	case PduTypeMechanismName:
		ret = "MECHANISMNAME"
	case PduTypeCallingAuthenticationValue:
		ret = "CALLINGAUTHENTICATIONVALUE"
	case PduTypeImplementationInformation:
		ret = "IMPLEMENTATIONINFORMATION"
	case PduTypeUserInformation:
		ret = "USERINFORMATION"
	}
	return ret
}
