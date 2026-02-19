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

// TranslatorOutputType enumerates Translator output types.
type TranslatorOutputType int

const (
	//TranslatorOutputTypeSimpleXML defines that output is in simple XML format.
	TranslatorOutputTypeSimpleXML TranslatorOutputType = iota
	//TranslatorOutputTypeStandardXML defines that output is in Standard XML format.
	TranslatorOutputTypeStandardXML
)

// TranslatorOutputTypeParse converts the given string into a TranslatorOutputType value.
//
// It returns the corresponding TranslatorOutputType constant if the string matches
// a known level name, or an error if the input is invalid.
func TranslatorOutputTypeParse(value string) (TranslatorOutputType, error) {
	var ret TranslatorOutputType
	var err error
	switch {
	case strings.EqualFold(value, "SimpleXml"):
		ret = TranslatorOutputTypeSimpleXML
	case strings.EqualFold(value, "StandardXml"):
		ret = TranslatorOutputTypeStandardXML
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the TranslatorOutputType.
// It satisfies fmt.Stringer.
func (g TranslatorOutputType) String() string {
	var ret string
	switch g {
	case TranslatorOutputTypeSimpleXML:
		ret = "SimpleXml"
	case TranslatorOutputTypeStandardXML:
		ret = "StandardXml"
	}
	return ret
}

// AllTranslatorOutputType returns a slice containing all defined TranslatorOutputType values.
func AllTranslatorOutputType() []TranslatorOutputType {
	return []TranslatorOutputType{
		TranslatorOutputTypeSimpleXML,
		TranslatorOutputTypeStandardXML,
	}
}
