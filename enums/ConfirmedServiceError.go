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

// ConfirmedServiceError :  tells when error has occurred.
type ConfirmedServiceError int

const (
	// ConfirmedServiceErrorInitiateError defines that the error has occurred on initialize.
	ConfirmedServiceErrorInitiateError ConfirmedServiceError = 1
	// ConfirmedServiceErrorRead defines that the error has occurred on read.
	ConfirmedServiceErrorRead ConfirmedServiceError = 5
	// ConfirmedServiceErrorWrite defines that the error has occurred on write.
	ConfirmedServiceErrorWrite ConfirmedServiceError = 6
)

// ConfirmedServiceErrorParse converts the given string into a ConfirmedServiceError value.
//
// It returns the corresponding ConfirmedServiceError constant if the string matches
// a known level name, or an error if the input is invalid.
func ConfirmedServiceErrorParse(value string) (ConfirmedServiceError, error) {
	var ret ConfirmedServiceError
	var err error
	switch {
	case strings.EqualFold(value, "InitiateError"):
		ret = ConfirmedServiceErrorInitiateError
	case strings.EqualFold(value, "Read"):
		ret = ConfirmedServiceErrorRead
	case strings.EqualFold(value, "Write"):
		ret = ConfirmedServiceErrorWrite
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the ConfirmedServiceError.
// It satisfies fmt.Stringer.
func (g ConfirmedServiceError) String() string {
	var ret string
	switch g {
	case ConfirmedServiceErrorInitiateError:
		ret = "InitiateError"
	case ConfirmedServiceErrorRead:
		ret = "Read"
	case ConfirmedServiceErrorWrite:
		ret = "Write"
	}
	return ret
}

// AllConfirmedServiceError returns a slice containing all defined ConfirmedServiceError values.
func AllConfirmedServiceError() []ConfirmedServiceError {
	return []ConfirmedServiceError{
	ConfirmedServiceErrorInitiateError,
	ConfirmedServiceErrorRead,
	ConfirmedServiceErrorWrite,
	}
}
