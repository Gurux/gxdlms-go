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

	"github.com/Gurux/gxdlms-go/enums"
)

type GXDLMSMonitoredValue struct {
	ObjectType enums.ObjectType

	LogicalName string

	AttributeIndex int8
}

func (g *GXDLMSMonitoredValue) update(value IGXDLMSBase, attributeIndex int8) {
	g.ObjectType = value.Base().ObjectType()
	g.LogicalName = value.Base().LogicalName()
	g.AttributeIndex = attributeIndex
}

func (g *GXDLMSMonitoredValue) String() string {
	return fmt.Sprintf("%s %s %d", g.ObjectType.String(), g.LogicalName, g.AttributeIndex)
}
