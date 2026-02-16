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
// Gurux Device Framework is Open Source software you can redistribute it
// and/or modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation version 2 of the License.
// Gurux Device Framework is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY without even the implied warranty of
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
	"github.com/Gurux/gxdlms-go/types"
)

// SN Parameters
type GXDLMSSNParameters struct {
	// DLMS settings.
	Settings *settings.GXDLMSSettings

	// DLMS Command.
	Command enums.Command

	// Request type.
	RequestType uint8

	// Attribute descriptor.
	AttributeDescriptor *types.GXByteBuffer

	// Data.
	Data *types.GXByteBuffer

	// Send date and time. This is used in Data notification messages.
	Time *types.GXDateTime

	// Item Count.
	Count int

	// Are there more data to send or more data to receive.
	MultipleBlocks bool

	// Block index.
	BlockIndex uint16
}

func NewGXDLMSSNParameters(settings *settings.GXDLMSSettings, command enums.Command, count int, commandType byte, attributeDescriptor *types.GXByteBuffer, data *types.GXByteBuffer) *GXDLMSSNParameters {
	ret := &GXDLMSSNParameters{
		Settings:            settings,
		BlockIndex:          uint16(settings.BlockIndex),
		Command:             command,
		Count:               count,
		RequestType:         commandType,
		AttributeDescriptor: attributeDescriptor,
		Data:                data,
	}
	if settings != nil {
		ret.Settings.Command = command
	}
	return ret

}
