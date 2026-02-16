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

// ApplicationReference :  describes application errors.
type ApplicationReference int

const (
	// ApplicationReferenceOther defines that the other error is occurred.
	ApplicationReferenceOther ApplicationReference = iota
	// ApplicationReferenceTimeElapsed defines that the time elapsed has occurred.
	ApplicationReferenceTimeElapsed
	// ApplicationReferenceApplicationUnreachable defines that the application unreachable has occurred.
	ApplicationReferenceApplicationUnreachable
	// ApplicationReferenceApplicationReferenceInvalid defines that the application reference is invalid has occurred.
	ApplicationReferenceApplicationReferenceInvalid
	// ApplicationReferenceApplicationContextUnsupported defines that the application context unsupported has occurred.
	ApplicationReferenceApplicationContextUnsupported
	// ApplicationReferenceProviderCommunicationError defines that the provider communication error has occurred.
	ApplicationReferenceProviderCommunicationError
	// ApplicationReferenceDecipheringError defines that the deciphering error has occurred.
	ApplicationReferenceDecipheringError
)

// ApplicationReferenceParse converts the given string into a ApplicationReference value.
//
// It returns the corresponding ApplicationReference constant if the string matches
// a known level name, or an error if the input is invalid.
func ApplicationReferenceParse(value string) (ApplicationReference, error) {
	var ret ApplicationReference
	var err error
	switch strings.ToUpper(value) {
	case "OTHER":
		ret = ApplicationReferenceOther
	case "TIMEELAPSED":
		ret = ApplicationReferenceTimeElapsed
	case "APPLICATIONUNREACHABLE":
		ret = ApplicationReferenceApplicationUnreachable
	case "APPLICATIONREFERENCEINVALID":
		ret = ApplicationReferenceApplicationReferenceInvalid
	case "APPLICATIONCONTEXTUNSUPPORTED":
		ret = ApplicationReferenceApplicationContextUnsupported
	case "PROVIDERCOMMUNICATIONERROR":
		ret = ApplicationReferenceProviderCommunicationError
	case "DECIPHERINGERROR":
		ret = ApplicationReferenceDecipheringError
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the ApplicationReference.
// It satisfies fmt.Stringer.
func (g ApplicationReference) String() string {
	var ret string
	switch g {
	case ApplicationReferenceOther:
		ret = "OTHER"
	case ApplicationReferenceTimeElapsed:
		ret = "TIMEELAPSED"
	case ApplicationReferenceApplicationUnreachable:
		ret = "APPLICATIONUNREACHABLE"
	case ApplicationReferenceApplicationReferenceInvalid:
		ret = "APPLICATIONREFERENCEINVALID"
	case ApplicationReferenceApplicationContextUnsupported:
		ret = "APPLICATIONCONTEXTUNSUPPORTED"
	case ApplicationReferenceProviderCommunicationError:
		ret = "PROVIDERCOMMUNICATIONERROR"
	case ApplicationReferenceDecipheringError:
		ret = "DECIPHERINGERROR"
	}
	return ret
}
