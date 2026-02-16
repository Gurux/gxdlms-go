package dlms

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
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/objects"
)

// Access item is used to generate Access Service message.
type GXDLMSAccessItem struct {
	// COSEM target object.
	Target objects.IGXDLMSBase

	// Executed command type.
	Command enums.AccessServiceCommandType

	// Attribute index.
	Index uint8

	// Reply error code.
	Error enums.ErrorCode

	// Reply value.
	Value any
}

func NewGXDLMSAccessItem(command enums.AccessServiceCommandType, target objects.IGXDLMSBase, index uint8) *GXDLMSAccessItem {
	return &GXDLMSAccessItem{
		Command: command,
		Target:  target,
		Index:   index,
	}
}
