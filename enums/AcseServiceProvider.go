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

// AcseServiceProvider numerates ACSE service providers.
type AcseServiceProvider int

const (
	// AcseServiceProviderNone defines that the there is no error.
	AcseServiceProviderNone AcseServiceProvider = iota
	// AcseServiceProviderNoReasonGiven defines that the reason is not given.
	AcseServiceProviderNoReasonGiven
	// AcseServiceProviderNoCommonAcseVersion defines that the invalid ACSE version.
	AcseServiceProviderNoCommonAcseVersion
)

// AcseServiceProviderParse converts the given string into a AcseServiceProvider value.
//
// It returns the corresponding AcseServiceProvider constant if the string matches
// a known level name, or an error if the input is invalid.
func AcseServiceProviderParse(value string) (AcseServiceProvider, error) {
	var ret AcseServiceProvider
	var err error
	switch {
	case strings.EqualFold(value, "None"):
		ret = AcseServiceProviderNone
	case strings.EqualFold(value, "NoReasonGiven"):
		ret = AcseServiceProviderNoReasonGiven
	case strings.EqualFold(value, "NoCommonAcseVersion"):
		ret = AcseServiceProviderNoCommonAcseVersion
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the AcseServiceProvider.
// It satisfies fmt.Stringer.
func (g AcseServiceProvider) String() string {
	var ret string
	switch g {
	case AcseServiceProviderNone:
		ret = "None"
	case AcseServiceProviderNoReasonGiven:
		ret = "NoReasonGiven"
	case AcseServiceProviderNoCommonAcseVersion:
		ret = "NoCommonAcseVersion"
	}
	return ret
}

// AllAcseServiceProvider returns a slice containing all defined AcseServiceProvider values.
func AllAcseServiceProvider() []AcseServiceProvider {
	return []AcseServiceProvider{
		AcseServiceProviderNone,
		AcseServiceProviderNoReasonGiven,
		AcseServiceProviderNoCommonAcseVersion,
	}
}
