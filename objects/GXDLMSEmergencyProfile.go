package objects

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

	"github.com/Gurux/gxdlms-go/types"
)

type GXDLMSEmergencyProfile struct {
	ID uint16

	ActivationTime types.GXDateTime

	Duration uint32
}

func (g *GXDLMSEmergencyProfile) String() string {
	ret := g.ActivationTime.ToFormatString(nil, true)
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("%d", g.ID))
	sb.WriteString(" ")
	sb.WriteString(ret)
	sb.WriteString(" ")
	sb.WriteString(fmt.Sprintf("%d", g.Duration))
	return sb.String()
}
