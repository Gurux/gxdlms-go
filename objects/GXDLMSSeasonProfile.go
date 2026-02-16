package objects

import (
	"strings"

	"github.com/Gurux/gxdlms-go/types"
)

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

type GXDLMSSeasonProfile struct {
	// Name of season profile.
	// Some manufacturers are using non ASCII names.
	Name []byte

	// Season Profile start time.
	Start types.GXDateTime

	// Week name of season profile.
	// Some manufacturers are using non ASCII names.
	WeekName []byte
}

func (g *GXDLMSSeasonProfile) String() string {
	sb := strings.Builder{}
	if types.IsAsciiString(g.Name) {
		sb.WriteString(string(g.Name))
	} else {
		sb.WriteString(types.ToHex(g.Name, false))
	}
	sb.WriteString(" ")
	sb.WriteString(g.Start.ToFormatString(nil, true))
	sb.WriteString(" ")
	if types.IsAsciiString(g.WeekName) {
		sb.WriteString(string(g.WeekName))
	} else {
		sb.WriteString(types.ToHex(g.WeekName, false))
	}
	return sb.String()
}
