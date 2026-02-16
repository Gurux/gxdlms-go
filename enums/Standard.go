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

// Standard Used DLMS standard.
type Standard int

const (
	// StandardDLMS defines that the meter uses default DLMS IEC 62056 standard. https://dlms.com
	StandardDLMS Standard = iota
	// StandardIndia defines that the meter uses India DLMS standard IS 15959-2. https://www.standardsbis.in
	StandardIndia
	// StandardItaly defines that the meter uses Italy DLMS standard UNI/TS 11291-11-2. https://uni.com
	StandardItaly
	// StandardSaudiArabia defines that the meter uses Saudi Arabia DLMS standard.
	StandardSaudiArabia
	// StandardIdis defines that the meter uses IDIS DLMS standard. https://www.idis-association.com/
	StandardIdis
	// StandardSpain defines that the meter uses Spain DLMS standard.
	StandardSpain
)

// StandardParse converts the given string into a Standard value.
//
// It returns the corresponding Standard constant if the string matches
// a known level name, or an error if the input is invalid.
func StandardParse(value string) (Standard, error) {
	var ret Standard
	var err error
	switch strings.ToUpper(value) {
	case "DLMS":
		ret = StandardDLMS
	case "INDIA":
		ret = StandardIndia
	case "ITALY":
		ret = StandardItaly
	case "SAUDIARABIA":
		ret = StandardSaudiArabia
	case "IDIS":
		ret = StandardIdis
	case "SPAIN":
		ret = StandardSpain
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the Standard.
// It satisfies fmt.Stringer.
func (g Standard) String() string {
	var ret string
	switch g {
	case StandardDLMS:
		ret = "DLMS"
	case StandardIndia:
		ret = "INDIA"
	case StandardItaly:
		ret = "ITALY"
	case StandardSaudiArabia:
		ret = "SAUDIARABIA"
	case StandardIdis:
		ret = "IDIS"
	case StandardSpain:
		ret = "SPAIN"
	}
	return ret
}
