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
package dlms

import "strings"

// OBIS Code class is used to find out default descrition to OBIS Code.
type GXStandardObisCode struct {
	// OBIS code.
	OBIS []string

	// OBIS code description.
	Description string

	// Interfaces that are using this OBIS code.
	Interfaces string

	// Standard data types.
	DataType string

	// Standard UI data type.
	UIDataType string
}

// String produces convert to string.
//             Returns:
func (g *GXStandardObisCode) String() string {
	return strings.Join(g.OBIS, ".") + " " + g.Description
}

// NewGXStandardObisCode creates new OBIS code object.
//             Parameters:
//                 obis:      OBIS code.
//				   desc:      OBIS code description.
//                 interfaces:Interfaces that are using this OBIS code.
//                 dataType:  Standard data types.
//             Returns:
//                 New OBIS code object.
func NewGXStandardObisCode(obis []string, desc string, interfaces string, dataType string) *GXStandardObisCode {
	return &GXStandardObisCode{
		OBIS:        obis,
		Description: desc,
		Interfaces:  interfaces,
		DataType:    dataType,
	}
}
