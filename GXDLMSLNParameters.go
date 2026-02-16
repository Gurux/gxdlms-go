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
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// LN Parameters
type GXDLMSLNParameters struct {
	// DLMS settings.
	settings *settings.GXDLMSSettings

	// DLMS Command.
	command enums.Command

	// Received Ciphered command.
	cipheredCommand enums.Command

	// Request type.
	requestType uint8

	// Attribute descriptor.
	attributeDescriptor *types.GXByteBuffer

	// Data.
	data *types.GXByteBuffer

	// Send date and time. This is used in Data notification messages.
	time *types.GXDateTime

	// Reply status.
	status uint8

	// Are there more data to send or more data to receive.
	multipleBlocks bool

	// Is this last block in send.
	lastBlock bool

	// Block index.
	blockIndex uint32

	// Block number ack.
	blockNumberAck uint16

	// Received invoke ID.
	invokeId uint32

	// GBT window size.
	gbtWindowSize uint8

	// Is GBT streaming used.
	streaming bool

	// Access mode.
	AccessMode int
}

func NewGXDLMSLNParameters(settings *settings.GXDLMSSettings,
	invokeId uint32,
	command enums.Command,
	commandType byte,
	attributeDescriptor *types.GXByteBuffer,
	data *types.GXByteBuffer,
	status byte,
	cipheredCommand enums.Command) *GXDLMSLNParameters {
	settings.Command = command
	if command == enums.CommandGetRequest && commandType != byte(enums.GetCommandTypeNextDataBlock) {
		settings.CommandType = commandType
	}
	return &GXDLMSLNParameters{
		settings:            settings,
		invokeId:            invokeId,
		blockIndex:          settings.BlockIndex,
		blockNumberAck:      settings.BlockNumberAck,
		command:             command,
		cipheredCommand:     cipheredCommand,
		requestType:         commandType,
		attributeDescriptor: attributeDescriptor,
		data:                data,
		status:              status,
		multipleBlocks:      settings.Count != settings.Index,
		lastBlock:           settings.Count == settings.Index,
		gbtWindowSize:       1,
	}
}
