package manufacturersettings

// --------------------------------------------------------------------------
//
//	Gurux Ltd
//
// Filename:        $HeadURL$
//
// Version:         $Revision$,
//
//	$Date$
//	$Author$
//
// # Copyright (c) Gurux Ltd
//
// ---------------------------------------------------------------------------
//
//	DESCRIPTION
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
// ---------------------------------------------------------------------------

import "github.com/Gurux/gxdlms-go/enums"

type GXObisCode struct {
	// Attribute index.
	AttributeIndex int

	// Logical name of the OBIS item.
	LogicalName string

	// Description of the OBIS item.
	Description string

	// All meters are not supporting Association view.
	//          If OBIS code is wanted to added by default set append to true.
	Append bool

	// UI data type. This is obsolete. Use ObjectType instead.
	UIDataType string

	// Object type.
	ObjectType enums.ObjectType

	// Object version.
	Version uint8

	// object attribute collection.
	Attributes GXAttributeCollection
}

// Constructor.
func NewGXObisCode(ln string, objectType enums.ObjectType, description string) *GXObisCode {
	return &GXObisCode{
		LogicalName: ln,
		ObjectType:  objectType,
		Description: description,
	}
}
