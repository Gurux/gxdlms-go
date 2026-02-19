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

// VdeState error describes  s.
type VdeStateError int

const (
	// VdeStateErrorOther defines that the Other error.
	VdeStateErrorOther VdeStateError = iota
	// VdeStateErrorNoDlmsContext defines that the no DLMS context.
	VdeStateErrorNoDlmsContext
	// VdeStateErrorLoadingDataSet defines that the loading dataset error.
	VdeStateErrorLoadingDataSet
	// VdeStateErrorStatusNochange defines that the status No change.
	VdeStateErrorStatusNochange
	// VdeStateErrorStatusInoperable defines that the status Inoperable.
	VdeStateErrorStatusInoperable
)

// VdeStateErrorParse converts the given string into a VdeStateError value.
//
// It returns the corresponding VdeStateError constant if the string matches
// a known level name, or an error if the input is invalid.
func VdeStateErrorParse(value string) (VdeStateError, error) {
	var ret VdeStateError
	var err error
	switch {
	case strings.EqualFold(value, "Other"):
		ret = VdeStateErrorOther
	case strings.EqualFold(value, "NoDlmsContext"):
		ret = VdeStateErrorNoDlmsContext
	case strings.EqualFold(value, "LoadingDataSet"):
		ret = VdeStateErrorLoadingDataSet
	case strings.EqualFold(value, "StatusNochange"):
		ret = VdeStateErrorStatusNochange
	case strings.EqualFold(value, "StatusInoperable"):
		ret = VdeStateErrorStatusInoperable
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the VdeStateError.
// It satisfies fmt.Stringer.
func (g VdeStateError) String() string {
	var ret string
	switch g {
	case VdeStateErrorOther:
		ret = "Other"
	case VdeStateErrorNoDlmsContext:
		ret = "NoDlmsContext"
	case VdeStateErrorLoadingDataSet:
		ret = "LoadingDataSet"
	case VdeStateErrorStatusNochange:
		ret = "StatusNochange"
	case VdeStateErrorStatusInoperable:
		ret = "StatusInoperable"
	}
	return ret
}

// AllVdeStateError returns a slice containing all defined VdeStateError values.
func AllVdeStateError() []VdeStateError {
	return []VdeStateError{
		VdeStateErrorOther,
		VdeStateErrorNoDlmsContext,
		VdeStateErrorLoadingDataSet,
		VdeStateErrorStatusNochange,
		VdeStateErrorStatusInoperable,
	}
}
