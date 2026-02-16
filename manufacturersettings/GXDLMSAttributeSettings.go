package manufacturersettings

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

type GXDLMSAttributeSettings struct {
	// Attribute data type.
	_type enums.DataType

	// Attribute name.
	Name string

	// Attribute Index.
	Index int

	Parent GXAttributeCollection

	// Data type that user ẃant's to see.
	UIType enums.DataType

	Access enums.AccessMode

	Access3 enums.AccessMode3

	MethodAccess enums.MethodAccessMode

	MethodAccess3 enums.MethodAccessMode3

	Static bool

	// Force that data is always sent as blocks.
	ForceToBlocks bool

	// Attribute values.
	Values GXObisValueItemCollection

	// XML Data template.
	Xml string

	// Read order.
	Order int

	// Available Access selector values.
	AccessSelector uint8
}

// Type returns the attribute data type.
func (g *GXDLMSAttributeSettings) Type() enums.DataType {
	if g.Index == 1 && g._type == enums.DataTypeNone {
		return enums.DataTypeOctetString
	}
	return g._type
}

// SetType sets the attribute data type.
func (g *GXDLMSAttributeSettings) SetType(value enums.DataType) {
	g._type = value
}

func (g *GXDLMSAttributeSettings) CopyTo(target GXDLMSAttributeSettings) {
	target.Name = g.Name
	target.Index = g.Index
	target._type = g._type
	target.UIType = g.UIType
	target.Access = g.Access
	target.Static = g.Static
	target.Values = g.Values
	target.Order = g.Order
	target.Xml = g.Xml
}

// String returns string representation of the attribute settings.
func (g *GXDLMSAttributeSettings) String() string {
	if g.Name == "" {
		return "Attribute #" + fmt.Sprint(g.Index)
	}
	return g.Name + fmt.Sprint(g.Index)
}
