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

// ApplicationContextName : .
type ApplicationContextName int

const (
	// ApplicationContextNameUnknown defines that the invalid application context name.
	ApplicationContextNameUnknown ApplicationContextName = iota
	// ApplicationContextNameLogicalName defines that the logical name.
	ApplicationContextNameLogicalName
	// ApplicationContextNameShortName defines that the short name.
	ApplicationContextNameShortName
	// ApplicationContextNameLogicalNameWithCiphering defines that the logical name with ciphering.
	ApplicationContextNameLogicalNameWithCiphering
	// ApplicationContextNameShortNameWithCiphering defines that the short name with ciphering.
	ApplicationContextNameShortNameWithCiphering
)

// ApplicationContextNameParse converts the given string into a ApplicationContextName value.
//
// It returns the corresponding ApplicationContextName constant if the string matches
// a known level name, or an error if the input is invalid.
func ApplicationContextNameParse(value string) (ApplicationContextName, error) {
	var ret ApplicationContextName
	var err error
	switch strings.ToUpper(value) {
	case "UNKNOWN":
		ret = ApplicationContextNameUnknown
	case "LOGICALNAME":
		ret = ApplicationContextNameLogicalName
	case "SHORTNAME":
		ret = ApplicationContextNameShortName
	case "LOGICALNAMEWITHCIPHERING":
		ret = ApplicationContextNameLogicalNameWithCiphering
	case "SHORTNAMEWITHCIPHERING":
		ret = ApplicationContextNameShortNameWithCiphering
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the ApplicationContextName.
// It satisfies fmt.Stringer.
func (g ApplicationContextName) String() string {
	var ret string
	switch g {
	case ApplicationContextNameUnknown:
		ret = "UNKNOWN"
	case ApplicationContextNameLogicalName:
		ret = "LOGICALNAME"
	case ApplicationContextNameShortName:
		ret = "SHORTNAME"
	case ApplicationContextNameLogicalNameWithCiphering:
		ret = "LOGICALNAMEWITHCIPHERING"
	case ApplicationContextNameShortNameWithCiphering:
		ret = "SHORTNAMEWITHCIPHERING"
	}
	return ret
}
