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

// Definition :  describes definition errors.
type Definition int

const (
	// DefinitionOther defines that the other error has occurred.
	DefinitionOther Definition = iota
	// DefinitionObjectUndefined defines that the object is undefined error has occurred.
	DefinitionObjectUndefined
	// DefinitionObjectClassInconsistent defines that the object class inconsistent error has occurred.
	DefinitionObjectClassInconsistent
	// DefinitionObjectAttributeInconsistent defines that the object attribute inconsistent error has occurred.
	DefinitionObjectAttributeInconsistent
)

// DefinitionParse converts the given string into a Definition value.
//
// It returns the corresponding Definition constant if the string matches
// a known level name, or an error if the input is invalid.
func DefinitionParse(value string) (Definition, error) {
	var ret Definition
	var err error
	switch strings.ToUpper(value) {
	case "OTHER":
		ret = DefinitionOther
	case "OBJECTUNDEFINED":
		ret = DefinitionObjectUndefined
	case "OBJECTCLASSINCONSISTENT":
		ret = DefinitionObjectClassInconsistent
	case "OBJECTATTRIBUTEINCONSISTENT":
		ret = DefinitionObjectAttributeInconsistent
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the Definition.
// It satisfies fmt.Stringer.
func (g Definition) String() string {
	var ret string
	switch g {
	case DefinitionOther:
		ret = "OTHER"
	case DefinitionObjectUndefined:
		ret = "OBJECTUNDEFINED"
	case DefinitionObjectClassInconsistent:
		ret = "OBJECTCLASSINCONSISTENT"
	case DefinitionObjectAttributeInconsistent:
		ret = "OBJECTATTRIBUTEINCONSISTENT"
	}
	return ret
}
