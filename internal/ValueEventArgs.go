package internal

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
	"github.com/Gurux/gxdlms-go/settings"
)

type ValueEventArgs struct {
	// Target COSEM object.
	Target interface{}

	// DLMS server.
	Server IGXDLMSServer

	// Parameterised access selector.
	Selector uint8

	// Optional parameters.
	Parameters any

	// DLMS settings.
	Settings *settings.GXDLMSSettings

	// Is action. This is reserved for internal use.
	Action bool

	// Attribute index of queried object.
	Index uint8

	// object value
	Value any

	// Is request handled.
	Handled bool

	// Occurred error.
	Error enums.ErrorCode

	// Is user updating the value.
	User bool

	// Is value max PDU size skipped when converting data to bytes.
	SkipMaxPduSize bool

	// Is reply handled as byte array or octet string.
	ByteArray bool

	// Row to PDU is used with Profile Generic to tell how many rows are fit to one PDU.
	RowToPdu uint16

	// Rows begin index.
	RowBeginIndex uint32

	// Rows end index.
	RowEndIndex uint32

	// Received invoke ID.
	InvokeId uint32
}

func NewValueEventArgs(settings *settings.GXDLMSSettings, target any, index uint8) *ValueEventArgs {
	return &ValueEventArgs{
		Settings: settings,
		Target:   target,
		Index:    index,
	}
}

func NewValueEventArgs2(server any, target interface{}, index uint8) *ValueEventArgs {
	return &ValueEventArgs{
		Server: server.(IGXDLMSServer),
		Target: target,
		Index:  index,
	}
}

func NewValueEventArgs3(target interface{}, index uint8, selector uint8, parameters any) *ValueEventArgs {
	return &ValueEventArgs{
		Target:     target,
		Index:      index,
		Selector:   selector,
		Parameters: parameters,
	}
}
