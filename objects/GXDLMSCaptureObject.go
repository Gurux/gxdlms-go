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

import "github.com/Gurux/gxdlms-go/types"

type GXDLMSCaptureObject struct {
	// Attribute Index of DLMS object in the profile generic table.
	AttributeIndex int

	// Data index of DLMS object in the profile generic table.
	DataIndex int

	// Restriction element for compact or push data.
	Restriction GXDLMSRestriction

	// Push data columns.
	Columns []types.GXKeyValuePair[IGXDLMSBase, GXDLMSCaptureObject]
}

// NewGXDLMSCaptureObject creates a new instance of GXDLMSCaptureObject.
func NewGXDLMSCaptureObject(attributeIndex int, dataIndex int) *GXDLMSCaptureObject {
	return &GXDLMSCaptureObject{
		AttributeIndex: attributeIndex,
		DataIndex:      dataIndex,
	}
}
