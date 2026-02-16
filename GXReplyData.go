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
	"fmt"
	"time"

	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// GXReplyData contains information from received reply data.
type GXReplyData struct {
	// Xml settings. This is used only on xml parser.
	xml *settings.GXDLMSTranslatorStructure

	// Broadcast indicates if data is sent as a broadcast or unicast.
	Broadcast bool

	// DataType is the type of the received data.
	DataType enums.DataType

	// Value is the read value.
	Value any

	// ReadPosition is the last read position. This is used in peek to solve how far data is read.
	ReadPosition int

	// PacketLength is the length of the packet.
	PacketLength int

	// RawPdu indicates PDU is not parsed and it's returned as it is.
	RawPdu bool

	// command is the received command.
	command enums.Command

	// commandType is the received command type.
	commandType byte

	// cipheredCommand is the received ciphered command.
	cipheredCommand enums.Command

	// Data is the received data.
	Data *types.GXByteBuffer

	// isComplete indicates if the frame is complete.
	isComplete bool

	// frameId is the HDLC frame ID.
	frameId byte

	// invokeId is the received invoke ID.
	invokeId uint32

	// systemTitle is the system title of the received PDU.
	// System title is set when ciphered notify packet is received.
	systemTitle []byte

	// Error is the received error code.
	Error int

	// EmptyResponses is true if there are empty frames or blocks.
	EmptyResponses enums.RequestTypes

	// TotalCount is the expected count of elements in the array.
	TotalCount int

	// CipherIndex is the position where data is decrypted or GBT is read.
	CipherIndex int

	// Time is the data notification date time.
	Time time.Time

	// BlockNumber is the GBT block number.
	BlockNumber uint16

	// BlockNumberAck is the GBT block number ACK.
	BlockNumberAck uint16

	// streaming indicates if GBT streaming is in use.
	streaming bool

	// gbtWindowSize is the GBT Window size.
	gbtWindowSize byte

	// hdlcStreaming indicates if HDLC streaming is in progress.
	hdlcStreaming bool

	// TargetAddress is the client address of the notification message.
	// Notification message sets this. This is also used with XML parser.
	TargetAddress int

	// SourceAddress is the server address of the notification message.
	// Notification message sets this. This is also used with XML parser.
	SourceAddress int

	// Gateway information.
	Gateway *settings.GXDLMSGateway

	// PrimeDc is PRIME data concentrator notification information.
	PrimeDc *GXDLMSPrimeDataConcentrator

	// moreData indicates if more data is available. Returns None if more data is not available or Frame or Block type.
	moreData enums.RequestTypes

	// Peek indicates if value is try to peek.
	Peek bool
}

// NewGXReplyData creates a new GXReplyData with default values.
func NewGXReplyData() *GXReplyData {
	r := &GXReplyData{
		Data: types.NewGXByteBuffer(),
	}
	r.Clear()
	return r
}

// newGXReplyDataFull creates a new GXReplyData with specified values.
func newGXReplyDataFull(more enums.RequestTypes, cmd enums.Command, buff *types.GXByteBuffer, complete bool, errorCode byte) *GXReplyData {
	r := &GXReplyData{
		Data: types.NewGXByteBuffer(),
	}
	r.Clear()
	r.moreData = more
	r.command = cmd
	r.Data = buff
	r.isComplete = complete
	r.Error = int(errorCode)
	return r
}

// Xml returns the XML settings. This is used only on xml parser.
func (r *GXReplyData) Xml() *settings.GXDLMSTranslatorStructure {
	return r.xml
}

// SetXml sets the XML settings. This is used only on xml parser.
func (r *GXReplyData) SetXml(value *settings.GXDLMSTranslatorStructure) {
	r.xml = value
}

// Command returns the received command.
func (r *GXReplyData) Command() enums.Command {
	return r.command
}

// SetCommand sets the received command.
func (r *GXReplyData) SetCommand(value enums.Command) {
	r.command = value
}

// CommandType returns the received command type.
func (r *GXReplyData) CommandType() byte {
	return r.commandType
}

// SetCommandType sets the received command type.
func (r *GXReplyData) SetCommandType(value byte) {
	r.commandType = value
}

// CipheredCommand returns the received ciphered command.
func (r *GXReplyData) CipheredCommand() enums.Command {
	return r.cipheredCommand
}

// SetCipheredCommand sets the received ciphered command.
func (r *GXReplyData) SetCipheredCommand(value enums.Command) {
	r.cipheredCommand = value
}

// IsComplete returns true if the frame is complete.
func (r *GXReplyData) IsComplete() bool {
	return r.isComplete
}

// SetIsComplete sets whether the frame is complete.
func (r *GXReplyData) SetIsComplete(value bool) {
	r.isComplete = value
}

// FrameId returns the HDLC frame ID.
func (r *GXReplyData) FrameId() byte {
	return r.frameId
}

// SetFrameId sets the HDLC frame ID.
func (r *GXReplyData) SetFrameId(value byte) {
	r.frameId = value
}

// GetInvokeId returns the received invoke ID.
func (r *GXReplyData) InvokeId() uint32 {
	return r.invokeId
}

// SetInvokeId sets the received invoke ID.
func (r *GXReplyData) SetInvokeId(value uint32) {
	r.invokeId = value
}

// SystemTitle returns the system title of the received PDU.
func (r *GXReplyData) SystemTitle() []byte {
	return r.systemTitle
}

// SetSystemTitle sets the system title of the received PDU.
func (r *GXReplyData) SetSystemTitle(value []byte) {
	r.systemTitle = value
}

// Streaming returns true if GBT streaming is in use.
func (r *GXReplyData) Streaming() bool {
	return r.streaming
}

// SetStreaming sets whether GBT streaming is in use.
func (r *GXReplyData) SetStreaming(value bool) {
	r.streaming = value
}

// GbtWindowSize returns the GBT Window size.
func (r *GXReplyData) GbtWindowSize() byte {
	return r.gbtWindowSize
}

// SetGbtWindowSize sets the GBT Window size.
func (r *GXReplyData) SetGbtWindowSize(value byte) {
	r.gbtWindowSize = value
}

// HdlcStreaming returns true if HDLC streaming is in progress.
func (r *GXReplyData) HdlcStreaming() bool {
	return r.hdlcStreaming
}

// SetHdlcStreaming sets whether HDLC streaming is in progress.
func (r *GXReplyData) SetHdlcStreaming(value bool) {
	r.hdlcStreaming = value
}

// IsStreaming returns true if GBT or HDLC streaming is used.
func (r *GXReplyData) IsStreaming() bool {
	return ((r.moreData&enums.RequestTypesFrame) == 0 &&
		r.streaming && (r.BlockNumberAck*uint16(r.gbtWindowSize))+1 > r.BlockNumber) ||
		((r.moreData&enums.RequestTypesFrame) == enums.RequestTypesFrame && r.hdlcStreaming)
}

// Clear resets data values to default.
func (r *GXReplyData) Clear() {
	r.moreData = enums.RequestTypesNone
	r.cipheredCommand = enums.CommandNone
	r.command = enums.CommandNone
	r.commandType = 0
	r.Data.SetCapacity(0)
	r.isComplete = false
	r.Error = 0
	r.TotalCount = 0
	r.Value = nil
	r.ReadPosition = 0
	r.PacketLength = 0
	r.DataType = enums.DataTypeNone
	r.CipherIndex = 0
	r.Time = time.Time{}
	r.gbtWindowSize = 0
	if r.xml != nil {
		r.xml.SetXmlLength(0)
	}
	r.invokeId = 0
}

// IsMoreData returns true if more data is available.
func (r *GXReplyData) IsMoreData() bool {
	return r.moreData != enums.RequestTypesNone && r.Error == 0
}

// IsNotify returns true if this is a notify message.
func (r *GXReplyData) IsNotify() bool {
	return r.frameId == 0x13 ||
		r.command == enums.CommandEventNotification ||
		r.command == enums.CommandDataNotification ||
		r.command == enums.CommandInformationReport
}

// GetMoreData returns the more data request type.
// Returns None if more data is not available or Frame or Block type.
func (r *GXReplyData) GetMoreData() enums.RequestTypes {
	return r.moreData
}

// setMoreData sets the more data request type.
func (r *GXReplyData) setMoreData(value enums.RequestTypes) {
	r.moreData = value
}

// GetErrorMessage returns the error message description.
func (r *GXReplyData) GetErrorMessage() string {
	return getErrorDescription(enums.ErrorCode(r.Error))
}

// Count returns the count of read elements.
// If this method is used, Peek must be set true.
func (r *GXReplyData) Count() int {
	if list, ok := r.Value.([]interface{}); ok {
		return len(list)
	}
	return 0
}

// String returns the content of reply data as a string.
func (r *GXReplyData) String() string {
	if r.xml != nil {
		return r.xml.String()
	}
	if r.Data == nil {
		return ""
	}
	return types.ToHexWithRange(r.Data.Array(), true, 0, r.Data.Size())
}

func getErrorDescription(errorCode enums.ErrorCode) string {
	switch errorCode {
	case enums.ErrorCodeOk:
		return ""
	case enums.ErrorCodeRejected:
		return "Rejected"
	case enums.ErrorCodeUnacceptableFrame:
		return "Unacceptable Frame"
	case enums.ErrorCodeDisconnectMode:
		return "Disconnect Mode"
	case enums.ErrorCodeHardwareFault:
		return "Hardware Fault"
	case enums.ErrorCodeTemporaryFailure:
		return "Temporary Failure"
	case enums.ErrorCodeReadWriteDenied:
		return "Read Write Denied"
	case enums.ErrorCodeUndefinedObject:
		return "Undefined Object"
	case enums.ErrorCodeInconsistentClass:
		return "Inconsistent Class"
	case enums.ErrorCodeUnavailableObject:
		return "Unavailable Object"
	case enums.ErrorCodeUnmatchedType:
		return "Unmatched Type"
	case enums.ErrorCodeAccessViolated:
		return "Access Violated"
	case enums.ErrorCodeDataBlockUnavailable:
		return "Data Block Unavailable"
	case enums.ErrorCodeLongGetOrReadAborted:
		return "Long Get Or Read Aborted"
	case enums.ErrorCodeNoLongGetOrReadInProgress:
		return "No Long Get Or Read In Progress"
	case enums.ErrorCodeLongSetOrWriteAborted:
		return "Long Set Or Write Aborted"
	case enums.ErrorCodeNoLongSetOrWriteInProgress:
		return "No Long Set Or Write In Progress"
	case enums.ErrorCodeDataBlockNumberInvalid:
		return "Data Block Number Invalid"
	case enums.ErrorCodeOtherReason:
		return "Other Reason"
	default:
		return fmt.Sprintf("Unknown Error (%d)", errorCode)
	}
}
