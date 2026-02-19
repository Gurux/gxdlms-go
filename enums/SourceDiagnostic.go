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

// SourceDiagnostic enumerates the error codes for reasons that can cause the server to reject the client.
type SourceDiagnostic int

const (
	// SourceDiagnosticNone defines that source diagnostic is not used.
	SourceDiagnosticNone SourceDiagnostic = iota
	// SourceDiagnosticNoReasonGiven defines that no reason is given for the error.
	SourceDiagnosticNoReasonGiven
	// SourceDiagnosticApplicationContextNameNotSupported defines that the application context name is not supported.
	SourceDiagnosticApplicationContextNameNotSupported
	// SourceDiagnosticCallingApTitleNotRecognized defines that the calling AP title not recognized.
	SourceDiagnosticCallingApTitleNotRecognized
	// SourceDiagnosticCallingApInvocationIdentifierNotRecognized defines that the calling AP invocation identifier not recognized.
	SourceDiagnosticCallingApInvocationIdentifierNotRecognized
	// SourceDiagnosticCallingAeQualifierNotRecognized defines that the calling AE qualifier not recognized
	SourceDiagnosticCallingAeQualifierNotRecognized
	// SourceDiagnosticCallingAeInvocationIdentifierNotRecognized defines that the calling AE invocation identifier not recognized
	SourceDiagnosticCallingAeInvocationIdentifierNotRecognized
	// SourceDiagnosticCalledApTitleNotRecognized defines that the called AP title not recognized
	SourceDiagnosticCalledApTitleNotRecognized
	// SourceDiagnosticCalledApInvocationIdentifierNotRecognized defines that the called AP invocation identifier not recognized
	SourceDiagnosticCalledApInvocationIdentifierNotRecognized
	// SourceDiagnosticCalledAeQualifierNotRecognized defines that the called AE qualifier not recognized
	SourceDiagnosticCalledAeQualifierNotRecognized
	// SourceDiagnosticCalledAeInvocationIdentifierNotRecognized defines that the called AE invocation identifier not recognized
	SourceDiagnosticCalledAeInvocationIdentifierNotRecognized
	// SourceDiagnosticAuthenticationMechanismNameNotRecognized defines that the authentication mechanism name is not recognized.
	SourceDiagnosticAuthenticationMechanismNameNotRecognized
	// SourceDiagnosticAuthenticationMechanismNameReguired defines that the authentication mechanism name is required.
	SourceDiagnosticAuthenticationMechanismNameReguired
	// SourceDiagnosticAuthenticationFailure defines that the authentication failure.
	SourceDiagnosticAuthenticationFailure
	// SourceDiagnosticAuthenticationRequired defines that the authentication is required.
	SourceDiagnosticAuthenticationRequired
)

// SourceDiagnosticParse converts the given string into a SourceDiagnostic value.
//
// It returns the corresponding SourceDiagnostic constant if the string matches
// a known level name, or an error if the input is invalid.
func SourceDiagnosticParse(value string) (SourceDiagnostic, error) {
	var ret SourceDiagnostic
	var err error
	switch {
	case strings.EqualFold(value, "None"):
		ret = SourceDiagnosticNone
	case strings.EqualFold(value, "NoReasonGiven"):
		ret = SourceDiagnosticNoReasonGiven
	case strings.EqualFold(value, "ApplicationContextNameNotSupported"):
		ret = SourceDiagnosticApplicationContextNameNotSupported
	case strings.EqualFold(value, "CallingApTitleNotRecognized"):
		ret = SourceDiagnosticCallingApTitleNotRecognized
	case strings.EqualFold(value, "CallingApInvocationIdentifierNotRecognized"):
		ret = SourceDiagnosticCallingApInvocationIdentifierNotRecognized
	case strings.EqualFold(value, "CallingAeQualifierNotRecognized"):
		ret = SourceDiagnosticCallingAeQualifierNotRecognized
	case strings.EqualFold(value, "CallingAeInvocationIdentifierNotRecognized"):
		ret = SourceDiagnosticCallingAeInvocationIdentifierNotRecognized
	case strings.EqualFold(value, "CalledApTitleNotRecognized"):
		ret = SourceDiagnosticCalledApTitleNotRecognized
	case strings.EqualFold(value, "CalledApInvocationIdentifierNotRecognized"):
		ret = SourceDiagnosticCalledApInvocationIdentifierNotRecognized
	case strings.EqualFold(value, "CalledAeQualifierNotRecognized"):
		ret = SourceDiagnosticCalledAeQualifierNotRecognized
	case strings.EqualFold(value, "CalledAeInvocationIdentifierNotRecognized"):
		ret = SourceDiagnosticCalledAeInvocationIdentifierNotRecognized
	case strings.EqualFold(value, "AuthenticationMechanismNameNotRecognized"):
		ret = SourceDiagnosticAuthenticationMechanismNameNotRecognized
	case strings.EqualFold(value, "AuthenticationMechanismNameReguired"):
		ret = SourceDiagnosticAuthenticationMechanismNameReguired
	case strings.EqualFold(value, "AuthenticationFailure"):
		ret = SourceDiagnosticAuthenticationFailure
	case strings.EqualFold(value, "AuthenticationRequired"):
		ret = SourceDiagnosticAuthenticationRequired
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the SourceDiagnostic.
// It satisfies fmt.Stringer.
func (g SourceDiagnostic) String() string {
	var ret string
	switch g {
	case SourceDiagnosticNone:
		ret = "None"
	case SourceDiagnosticNoReasonGiven:
		ret = "NoReasonGiven"
	case SourceDiagnosticApplicationContextNameNotSupported:
		ret = "ApplicationContextNameNotSupported"
	case SourceDiagnosticCallingApTitleNotRecognized:
		ret = "CallingApTitleNotRecognized"
	case SourceDiagnosticCallingApInvocationIdentifierNotRecognized:
		ret = "CallingApInvocationIdentifierNotRecognized"
	case SourceDiagnosticCallingAeQualifierNotRecognized:
		ret = "CallingAeQualifierNotRecognized"
	case SourceDiagnosticCallingAeInvocationIdentifierNotRecognized:
		ret = "CallingAeInvocationIdentifierNotRecognized"
	case SourceDiagnosticCalledApTitleNotRecognized:
		ret = "CalledApTitleNotRecognized"
	case SourceDiagnosticCalledApInvocationIdentifierNotRecognized:
		ret = "CalledApInvocationIdentifierNotRecognized"
	case SourceDiagnosticCalledAeQualifierNotRecognized:
		ret = "CalledAeQualifierNotRecognized"
	case SourceDiagnosticCalledAeInvocationIdentifierNotRecognized:
		ret = "CalledAeInvocationIdentifierNotRecognized"
	case SourceDiagnosticAuthenticationMechanismNameNotRecognized:
		ret = "AuthenticationMechanismNameNotRecognized"
	case SourceDiagnosticAuthenticationMechanismNameReguired:
		ret = "AuthenticationMechanismNameReguired"
	case SourceDiagnosticAuthenticationFailure:
		ret = "AuthenticationFailure"
	case SourceDiagnosticAuthenticationRequired:
		ret = "AuthenticationRequired"
	}
	return ret
}

// AllSourceDiagnostic returns a slice containing all defined SourceDiagnostic values.
func AllSourceDiagnostic() []SourceDiagnostic {
	return []SourceDiagnostic{
		SourceDiagnosticNone,
		SourceDiagnosticNoReasonGiven,
		SourceDiagnosticApplicationContextNameNotSupported,
		SourceDiagnosticCallingApTitleNotRecognized,
		SourceDiagnosticCallingApInvocationIdentifierNotRecognized,
		SourceDiagnosticCallingAeQualifierNotRecognized,
		SourceDiagnosticCallingAeInvocationIdentifierNotRecognized,
		SourceDiagnosticCalledApTitleNotRecognized,
		SourceDiagnosticCalledApInvocationIdentifierNotRecognized,
		SourceDiagnosticCalledAeQualifierNotRecognized,
		SourceDiagnosticCalledAeInvocationIdentifierNotRecognized,
		SourceDiagnosticAuthenticationMechanismNameNotRecognized,
		SourceDiagnosticAuthenticationMechanismNameReguired,
		SourceDiagnosticAuthenticationFailure,
		SourceDiagnosticAuthenticationRequired,
	}
}
