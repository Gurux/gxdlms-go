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

// ExceptionStateError StateError enumerates state errors.
type ExceptionStateError int

const (
	//ExceptionStateErrorServiceNotAllowed defines that the service is not allowed.
	ExceptionStateErrorServiceNotAllowed ExceptionStateError = 1
	// ExceptionStateErrorServiceUnknown defines that the service is unknown.
	ExceptionStateErrorServiceUnknown ExceptionStateError = 2
)

// ExceptionStateErrorParse converts the given string into a ExceptionStateError value.
//
// It returns the corresponding ExceptionStateError constant if the string matches
// a known level name, or an error if the input is invalid.
func ExceptionStateErrorParse(value string) (ExceptionStateError, error) {
	var ret ExceptionStateError
	var err error
	switch {
	case strings.EqualFold(value, "ServiceNotAllowed"):
		ret = ExceptionStateErrorServiceNotAllowed
	case strings.EqualFold(value, "ServiceUnknown"):
		ret = ExceptionStateErrorServiceUnknown
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the ExceptionStateError.
// It satisfies fmt.Stringer.
func (g ExceptionStateError) String() string {
	var ret string
	switch g {
	case ExceptionStateErrorServiceNotAllowed:
		ret = "ServiceNotAllowed"
	case ExceptionStateErrorServiceUnknown:
		ret = "ServiceUnknown"
	}
	return ret
}

// AllExceptionStateError returns a slice containing all defined ExceptionStateError values.
func AllExceptionStateError() []ExceptionStateError {
	return []ExceptionStateError{
		ExceptionStateErrorServiceNotAllowed,
		ExceptionStateErrorServiceUnknown,
	}
}
