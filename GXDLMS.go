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
	"errors"
	"fmt"
	"log"
	"math"
	"reflect"
	"strconv"
	"time"

	"github.com/Gurux/gxdlms-go/dlmserrors"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/buffer"
	"github.com/Gurux/gxdlms-go/internal/constants"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/objects"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

const defaultMaxInfoRX = 128
const defaultMaxInfoTX = 128
const defaultWindowSizeRX = 1
const defaultWindowSizeTX = 1

func getObjectCollection(value interface{}) *objects.GXDLMSObjectCollection {
	return value.(*objects.GXDLMSObjectCollection)
}

// getInvokeIDPriority returns the generates Invoke ID and priority.
//
// Parameters:
//
//	settings: DLMS settings.
//	increase: Is invoke ID increased.
//
// Returns:
//
//	Invoke ID and priority.
func getInvokeIDPriority(settings *settings.GXDLMSSettings, increase bool) uint8 {
	value := uint8(0)
	if settings.Priority == enums.PriorityHigh {
		value = 0x80
	}
	if settings.ServiceClass == enums.ServiceClassConfirmed {
		value |= 0x40
	}
	if increase {
		settings.SetInvokeID(uint8(((settings.InvokeID() + 1) & 0xF)))
	}
	value |= uint8((settings.InvokeID() & 0xF))
	return value
}

// getLongInvokeIDPriority returns the generates Invoke ID and priority.
//
// Parameters:
//
//	settings: DLMS settings.
//
// Returns:
//
//	Invoke ID and priority.
func getLongInvokeIDPriority(settings *settings.GXDLMSSettings) uint32 {
	value := uint32(0)
	if settings.Priority == enums.PriorityHigh {
		value = 0x80000000
	}
	if settings.ServiceClass == enums.ServiceClassConfirmed {
		value |= 0x40000000
	}
	value |= uint32((settings.LongInvokeID & 0xFFFFFF))
	settings.LongInvokeID++
	return value
}

// getGloMessage returns the get used glo message.
//
// Parameters:
//
//	cmd: Executed enums.Command
//
// Returns:
//
//	Integer value of glo message.
func getGloMessage(cmd enums.Command) (uint8, error) {
	switch cmd {
	case enums.CommandReadRequest:
		cmd = enums.CommandGloReadRequest
	case enums.CommandGetRequest:
		cmd = enums.CommandGloGetRequest
	case enums.CommandWriteRequest:
		cmd = enums.CommandGloWriteRequest
	case enums.CommandSetRequest:
		cmd = enums.CommandGloSetRequest
	case enums.CommandMethodRequest:
		cmd = enums.CommandGloMethodRequest
	case enums.CommandReadResponse:
		cmd = enums.CommandGloReadResponse
	case enums.CommandGetResponse:
		cmd = enums.CommandGloGetResponse
	case enums.CommandWriteResponse:
		cmd = enums.CommandGloWriteResponse
	case enums.CommandSetResponse:
		cmd = enums.CommandGloSetResponse
	case enums.CommandMethodResponse:
		cmd = enums.CommandGloMethodResponse
	case enums.CommandDataNotification:
		cmd = enums.CommandGeneralGloCiphering
	case enums.CommandReleaseRequest:
		cmd = enums.CommandReleaseRequest
	case enums.CommandReleaseResponse:
		cmd = enums.CommandReleaseResponse
	case enums.CommandAccessRequest:
	case enums.CommandAccessResponse:
		cmd = enums.CommandGeneralCiphering
	default:
		return 0, dlmserrors.ErrInvalidGloCommand
	}
	return uint8(cmd), nil
}

// getDedMessage returns the get used ded message.
//
// Parameters:
//
//	cmd: Executed enums.Command
//
// Returns:
//
//	Integer value of ded message.
func getDedMessage(cmd enums.Command) (uint8, error) {
	switch cmd {
	case enums.CommandGetRequest:
		cmd = enums.CommandDedGetRequest
	case enums.CommandSetRequest:
		cmd = enums.CommandDedSetRequest
	case enums.CommandMethodRequest:
		cmd = enums.CommandDedMethodRequest
	case enums.CommandGetResponse:
		cmd = enums.CommandDedGetResponse
	case enums.CommandSetResponse:
		cmd = enums.CommandDedSetResponse
	case enums.CommandMethodResponse:
		cmd = enums.CommandDedMethodResponse
	case enums.CommandDataNotification:
		cmd = enums.CommandGeneralDedCiphering
	case enums.CommandReleaseRequest:
		cmd = enums.CommandReleaseRequest
	case enums.CommandReleaseResponse:
		cmd = enums.CommandReleaseResponse
	case enums.CommandAccessRequest:
	case enums.CommandAccessResponse:
		cmd = enums.CommandGeneralDedCiphering
	default:
		return 0, errors.New("Invalid DED enums.Command")
	}
	return uint8(cmd), nil
}

// multipleBlocks returns the check is all data fit to one data block.
//
// Parameters:
//
//	p: LN parameters.
//	reply: Generated reply.
func multipleBlocks(p *GXDLMSLNParameters, reply *types.GXByteBuffer, ciphering bool) {
	// Check is all data fit to one message if data is given.
	len := p.data.Available()
	if p.attributeDescriptor != nil {
		len = len + p.attributeDescriptor.Size()
	}
	if ciphering {
		len = len + internal.CipheringHeaderSize
	}
	// If system title is sent.
	if (p.settings.NegotiatedConformance & enums.ConformanceGeneralProtection) != 0 {
		// System title is not send with Italy protocol.
		if p.settings.Standard != enums.StandardItaly {
			len = len + 9
		} else {
			len = len + 1
		}
	}
	len = len + getSigningSize(p)
	if !p.multipleBlocks {
		p.multipleBlocks = 2+reply.Size()+len > int(p.settings.MaxPduSize())
	}
	if p.multipleBlocks {
		p.lastBlock = !(8+reply.Available()+len > int(p.settings.MaxPduSize()))
	}
	if p.lastBlock {
		p.lastBlock = !(8+reply.Available()+len > int(p.settings.MaxPduSize()))
	}
}

func isGloMessage(cmd enums.Command) bool {
	return cmd == enums.CommandGloGetRequest || cmd == enums.CommandGloSetRequest || cmd == enums.CommandGloMethodRequest
}

func getBlockCipherKey(settings *settings.GXDLMSSettings) ([]byte, error) {
	if settings.Broadcast {
		if len(settings.Cipher.BroadcastBlockCipherKey()) == 0 {
			return nil, errors.New("Invalid Broadcast block cipher key.")
		}
		return settings.Cipher.BroadcastBlockCipherKey(), nil
	}
	if len(settings.EphemeralBlockCipherKey) != 0 {
		return settings.EphemeralBlockCipherKey, nil
	}
	return settings.Cipher.BlockCipherKey(), nil
}

func getAuthenticationKey(settings *settings.GXDLMSSettings) ([]byte, error) {
	if len(settings.EphemeralAuthenticationKey) != 0 {
		return settings.EphemeralAuthenticationKey, nil
	}
	return settings.Cipher.AuthenticationKey(), nil
}

// shoudSign returns the should the message sign.
func shoudSign(p *GXDLMSLNParameters) bool {
	signing := p.cipheredCommand == enums.CommandGeneralSigning ||
		(p.settings.Cipher != nil && p.settings.Cipher.Signing() != enums.SigningNone)
	if !signing {
		// Association LN V3 and signing is not needed.
		if p.settings.IsServer() {
			signing = (p.AccessMode & int((enums.AccessMode3DigitallySignedResponse))) != 0
		} else {
			signing = (p.AccessMode & int((enums.AccessMode3DigitallySignedRequest))) != 0
		}
	}
	return signing
}

// GetSigningSize returns the return amount of the bytes that signing requires.
func getSigningSize(p *GXDLMSLNParameters) int {
	size := 0
	if p.settings.Cipher != nil && p.settings.Cipher.Signing() == enums.SigningGeneralSigning {
		if p.settings.Cipher.SecuritySuite() == enums.SecuritySuite1 {
			size = 65
		} else if p.settings.Cipher.SecuritySuite() == enums.SecuritySuite2 {
			size = 99
		}
	}
	return size
}

func appendMultipleSNBlocks(p *GXDLMSSNParameters, reply *types.GXByteBuffer) (int, error) {
	var err error
	ciphering := p.Settings.IsCiphered(false)
	hSize := uint16(reply.Size() + 3)
	// Add LLC bytes.
	if p.Command == enums.CommandWriteRequest || p.Command == enums.CommandReadRequest {
		hSize = hSize + 1 + uint16(helpers.GetObjectCountSizeInBytes(p.Count))
	}
	maxSize := int(p.Settings.MaxPduSize() - hSize)
	if ciphering {
		maxSize = maxSize - internal.CipheringHeaderSize
		if useHdlc(p.Settings.InterfaceType) {
			maxSize = maxSize - 3
		}
	}
	maxSize = maxSize - int(helpers.GetObjectCountSizeInBytes(maxSize))
	if p.Data.Size()-p.Data.Position() > maxSize {
		err := reply.SetUint8(0)
		if err != nil {
			return 0, err
		}
	} else {
		err = reply.SetUint8(1)
		if err != nil {
			return 0, err
		}
		maxSize = p.Data.Size() - p.Data.Position()
	}
	err = reply.SetUint16(p.BlockIndex)
	if err != nil {
		return 0, err
	}
	if p.Command == enums.CommandWriteRequest {
		p.BlockIndex++
		types.SetObjectCount(p.Count, reply)
		err = reply.SetUint8(enums.DataTypeOctetString)
		if err != nil {
			return 0, err
		}
	}
	if p.Command == enums.CommandReadRequest {
		p.BlockIndex++
	}
	types.SetObjectCount(maxSize, reply)
	return maxSize, nil
}

// getPlcFrame returns the get MAC LLC frame for data.
//
// Parameters:
//
//	settings: DLMS settings.
//	data: Data to add.
//
// Returns:
//
//	MAC frame.
func getPlcFrame(settings *settings.GXDLMSSettings, creditFields uint8, data *types.GXByteBuffer) ([]byte, error) {
	var err error
	frameSize := data.Available()
	// Max frame size is 124 bytes.
	if frameSize > 134 {
		frameSize = 134
	}
	// PAD Length.
	padLen := (36 - ((11 + frameSize) % 36)) % 36
	bb := types.GXByteBuffer{}
	bb.SetCapacity(15 + frameSize + padLen)
	err = bb.SetUint8(2)
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(uint8((11 + frameSize)))
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(0x50)
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(creditFields)
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(uint8((settings.Plc.MacSourceAddress >> 4)))
	if err != nil {
		return nil, err
	}
	val := settings.Plc.MacSourceAddress << 12
	val |= settings.Plc.MacDestinationAddress & 0xFFF
	err = bb.SetUint16(uint16(val))
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(uint8(padLen))
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(uint8(enums.PlcDataLinkDataRequest))
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(uint8(settings.ServerAddress))
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(uint8(settings.ClientAddress))
	if err != nil {
		return nil, err
	}
	err = bb.SetAt(data.Array(), bb.Position(), frameSize)
	if err != nil {
		return nil, err
	}
	for padLen != 0 {
		err = bb.SetUint8(0)
		if err != nil {
			return nil, err
		}
		padLen--
	}
	// Checksum.
	crc := countFCS16(bb.Array(), 0, bb.Size())
	err = bb.SetUint16(crc)
	if err != nil {
		return nil, err
	}
	// Remove sent data in server side.
	if settings.IsServer() {
		if data.Size() == data.Position() {
			data.Clear()
		} else {
			data.Move(data.Position(), 0, data.Size()-data.Position())
			data.SetPosition(0)
		}
	}
	return bb.Array(), nil
}

// getLLCBytes returns the check LLC bytes.
//
// Parameters:
//
//	server: Is server.
//	data: Received data.
func getLLCBytes(server bool, data *types.GXByteBuffer) bool {
	if server {
		return data.Compare(internal.LLCSendBytes)
	}
	return data.Compare(internal.LLCReplyBytes)
}

func getHdlcData(server bool, settings *settings.GXDLMSSettings, reply *types.GXByteBuffer, data *GXReplyData, notify *GXReplyData) (uint8, error) {
	ch := uint8(0)
	var err error
	packetStartID := reply.Position()
	frameLen := 0
	var crc uint16
	var crcRead uint16
	first := (data.moreData&enums.RequestTypesFrame) == 0 || (notify != nil && (notify.moreData&enums.RequestTypesFrame) != 0)
	// If whole frame is not received yet.
	if reply.Size()-reply.Position() < 9 {
		data.isComplete = false
		if notify != nil {
			notify.isComplete = false
		}
		return 0, nil
	}
	data.isComplete = true
	if notify != nil {
		notify.isComplete = true
	}
	isNotify := false
	for pos := reply.Position(); pos < reply.Size(); pos++ {
		ch, err = reply.Uint8()
		if ch == internal.HDLCFrameStartEnd {
			packetStartID = pos
			break
		}
	}
	// Not a HDLC frame. Sometimes meters can send some strange data between DLMS frames.
	if reply.Position() == reply.Size() {
		data.isComplete = false
		if notify != nil {
			notify.isComplete = false
		}
		// Not enough data to parse
		return 0, nil
	}
	frame, err := reply.Uint8()
	if (frame & 0xF0) != 0xA0 {
		reply.SetPosition(reply.Position() - 1)
		return getHdlcData(server, settings, reply, data, notify)
	}
	// Check frame length.
	if (frame & 0x7) != 0 {
		ret := int(frame & 0x7)
		ret = ret << 8
		frameLen = ret
	}
	ch, err = reply.Uint8()
	frameLen = frameLen + int(ch)
	if reply.Size()-reply.Position()+1 < frameLen {
		data.isComplete = false
		reply.SetPosition(packetStartID)
		// Not enough data to parse
		return 0, nil
	}
	eopPos := frameLen + packetStartID + 1
	ch, err = reply.Uint8At(eopPos)
	if ch != internal.HDLCFrameStartEnd {
		reply.SetPosition(2)
		return getHdlcData(server, settings, reply, data, notify)
	}
	// Check addresses.
	var source int
	var target int
	var ret bool
	ret, err = checkHdlcAddress(server, settings, reply, eopPos, &source, &target)
	if err != nil {
		ret = false
		source = 0
		target = 0
	}

	if !ret {
		// If not notify.
		ch1, _ := reply.Uint8At(reply.Position())
		ch2, _ := reply.Uint8At(reply.Position() + 1)
		if !(reply.Position() < reply.Size() && (ch1 == 0x13 || ch2 == 0x3)) {
			reply.SetPosition(1 + eopPos)
			return getHdlcData(server, settings, reply, data, notify)
		}
		if notify != nil {
			isNotify = true
			notify.TargetAddress = target
			notify.SourceAddress = source
		}
	}
	// Is there more data available.
	moreData := (frame & 0x8) != 0
	frame, _ = reply.Uint8()
	if data.xml == nil && !settings.CheckFrame(frame, data.xml) {
		reply.SetPosition(eopPos + 1)
		return getHdlcData(server, settings, reply, data, notify)
	}
	// If server is using same client and server address for notifications.
	if (frame == 0x13 || frame == 0x3) && !isNotify && notify != nil {
		isNotify = true
		notify.TargetAddress = target
		notify.SourceAddress = source
	}
	if moreData {
		if isNotify {
			notify.moreData = enums.RequestTypes((notify.moreData | enums.RequestTypesFrame))
			notify.hdlcStreaming = (frame & 0x10) == 0
		} else {
			data.moreData = enums.RequestTypes((data.moreData | enums.RequestTypesFrame))
			data.hdlcStreaming = (frame & 0x10) == 0
		}
	} else if (frame&0x10) != 0 || settings.Hdlc.WindowSizeRX() == 1 {
		//If the final bit is set. This is used when Window size > 1.
		if isNotify {
			notify.moreData = enums.RequestTypes((notify.moreData & ^enums.RequestTypesFrame))
			notify.hdlcStreaming = false
		} else {
			data.moreData = enums.RequestTypes((data.moreData & ^enums.RequestTypesFrame))
			data.hdlcStreaming = false
		}
	}
	crc = countFCS16(reply.Array(), int(packetStartID+1), int(reply.Position()-packetStartID-1))
	crcRead, _ = reply.Uint16()
	if crc != crcRead {
		if reply.Size()-reply.Position() > 8 {
			return getHdlcData(server, settings, reply, data, notify)
		}
		if data.xml == nil {
			return 0, errors.New("Invalid header checksum.")
		}
		data.xml.AppendComment("Invalid header checksum.")
	}
	// Check that packet CRC match only if there is a data part.
	if reply.Position() != packetStartID+frameLen+1 {
		crc = countFCS16(reply.Array(), int(packetStartID+1), int(frameLen-2))
		crcRead, _ = reply.Uint16At(int(packetStartID + frameLen - 1))
		if crc != crcRead {
			if data.xml == nil {
				return 0, errors.New("Invalid data checksum.")
			}
			data.xml.AppendComment("Invalid data checksum.")
		}
		// Remove CRC and EOP from packet length.
		if isNotify {
			notify.PacketLength = eopPos - 2
		} else {
			data.PacketLength = eopPos - 2
		}
	} else {
		if isNotify {
			notify.PacketLength = reply.Position() + 1
		} else {
			data.PacketLength = reply.Position() + 1
		}
	}
	// If client want to know used server and client address.
	if data.TargetAddress == 0 && data.SourceAddress == 0 {
		data.TargetAddress = target
		data.SourceAddress = source
	}
	if frame != 0x13 && frame != 0x3 && (frame&uint8(constants.HdlcFrameTypeUframe)) == uint8(constants.HdlcFrameTypeUframe) {
		// Get Eop if there is no data.
		if reply.Position() == packetStartID+frameLen+1 {
			reply.Uint8()
		}
		switch frame {
		case 0x97:
			data.Error = int(enums.ErrorCodeUnacceptableFrame)
		case 0x1F:
			data.Error = int(enums.ErrorCodeDisconnectMode)
		case 0x17:
			data.Error = int(enums.ErrorCodeDisconnectMode)
		}
		data.command = enums.Command(frame)
		if data.command == enums.CommandSnrm {
			settings.Connected &= ^enums.ConnectionStateIec
		}
	} else if frame != 0x13 && frame != 0x3 && (frame&uint8(constants.HdlcFrameTypeSframe)) == uint8(constants.HdlcFrameTypeSframe) {
		// If frame is rejected.
		tmp := (frame >> 2) & 0x3
		if tmp == uint8(constants.HdlcControlFrameReject) {
			data.Error = int(enums.ErrorCodeRejected)
		} else if tmp == uint8(constants.HdlcControlFrameReceiveNotReady) {
			data.Error = int(enums.ErrorCodeReceiveNotReady)
		} else if tmp == uint8(constants.HdlcControlFrameReceiveReady) {
			log.Println("ReceiveReady.")
		}
		// Get Eop if there is no data.
		if reply.Position() == packetStartID+frameLen+1 {
			reply.Uint8()
		}
	} else {
		// Get Eop if there is no data.
		if reply.Position() == packetStartID+frameLen+1 {
			reply.Uint8()
			if (frame & 0x1) == 0x1 {
				data.moreData = enums.RequestTypes((data.moreData | enums.RequestTypesFrame))
			}
		} else {
			if first {
				llc := getLLCBytes(server, reply)
				if data.xml == nil {
					if !llc && (frame == 0x13 || frame == 0x3) {
						llc = getLLCBytes(!server, reply)
					}
					if !llc {
						return 0, errors.New("LLC bytes are missing from the message.")
					}
				} else if !llc {
					getLLCBytes(!server, reply)
				}
			}
		}
	}
	return frame, nil
}

func getServerAddress(address int, logical int, physical int) {
	if address < 0x4000 {
		logical = address >> 7
		physical = address & 0x7F
	} else {
		logical = address >> 14
		physical = address & 0x3FFF
	}
}

// checkHdlcAddress returns the check that client and server address match.
//
// Parameters:
//
//	server: Is server.
//	settings: DLMS settings.
//	reply: Received data.
//	index: Position.
//
// Returns:
//
//	True, if client and server address match.
func checkHdlcAddress(server bool, settings *settings.GXDLMSSettings, reply *types.GXByteBuffer, index int, source *int, target *int) (bool, error) {
	var err error
	*target, err = GetHDLCAddressFromByteBuffer(reply)
	if err != nil {
		return false, err
	}
	*source, err = GetHDLCAddressFromByteBuffer(reply)
	if err != nil {
		return false, err
	}
	if server {
		// Check that server addresses match.
		if settings.ServerAddress != 0 && settings.ServerAddress != *target {
			value, err := reply.Uint8At(reply.Position())
			if err != nil {
				return false, err
			}
			if value == uint8(enums.CommandSnrm) {
				settings.ServerAddress = *target
			} else {
				return false, errors.New("Destination addresses do not match. It is " + fmt.Sprint(target) + ". It should be " + fmt.Sprint() + ".")
			}
		} else {
			settings.ServerAddress = *target
		}
		// Check that client addresses match.
		if settings.ClientAddress != 0 && settings.ClientAddress != *source {
			value, err := reply.Uint8At(reply.Position())
			if err != nil {
				return false, err
			}
			if value == uint8(enums.CommandSnrm) {
				settings.ClientAddress = *source
			} else {
				return false, errors.New("Source addresses do not match. It is " + fmt.Sprint(source) + ". It should be " + fmt.Sprint() + ".")
			}
		} else {
			settings.ClientAddress = *source
		}
	} else {
		// Check that client addresses match.
		if settings.ClientAddress != *target {
			// If echo.
			if settings.ClientAddress == *source && settings.ServerAddress == *target {
				reply.SetPosition(index + 1)
			} else if settings.ClientAddress == 0 && (settings.ServerAddress == 0x7F || settings.ServerAddress == 0) {
				//If client wants to know used client and server address.
				return true, nil
			}
			return false, nil
		}
		// Check that server addresses match.
		if settings.ServerAddress != *source && (settings.ServerAddress&0x7F) != 0x7F && (settings.ServerAddress&0x3FFF) != 0x3FFF {
			// Check logical and physical address separately.This is done because some meters might send four byteswhen only two bytes is needed.
			var readLogical int
			var readPhysical int
			var logical int
			var physical int
			getServerAddress(*source, readLogical, readPhysical)
			getServerAddress(settings.ServerAddress, logical, physical)
			if readLogical != logical || readPhysical != physical {
				return false, nil
			}
		}
	}
	return true, nil
}

// getTcpData returns the get data from TCP/IP frame.
//
// Parameters:
//
//	settings: DLMS settings.
//	buff: Received data.
//	data: Reply information.
//	notify: Notify information.
func getTcpData(settings *settings.GXDLMSSettings, buff *types.GXByteBuffer, data *GXReplyData, notify *GXReplyData) bool {
	// If whole frame is not received yet.
	if buff.Available() < 8 {
		data.isComplete = false
		return true
	}
	var value uint16
	isData := true
	pos := buff.Position()
	data.isComplete = false
	if notify != nil {
		notify.isComplete = false
	}
	for buff.Available() > 2 {
		value, _ = buff.Uint16()
		if value == 1 {
			if buff.Available() < 6 {
				isData = false
				break
			}
			// Check TCP/IP addresses.
			ret, _ := checkWrapperAddress(settings, buff, data, notify)
			if !ret {
				data = notify
				isData = false
			}
			value, _ := buff.Uint16()
			data.isComplete = !(uint16(buff.Size()-buff.Position()) < value)
			if !data.isComplete {
				buff.SetPosition(pos)
			} else {
				data.PacketLength = int(buff.Position()) + int(value)
			}
			break
		} else {
			buff.SetPosition(buff.Position() - 1)
		}
	}
	return isData
}

// getSmsData returns the get data from SMS frame.
//
// Parameters:
//
//	settings: DLMS settings.
//	buff: Received data.
//	data: Reply information.
//	notify: Notify information.
func getSmsData(settings *settings.GXDLMSSettings, buff *types.GXByteBuffer, data *GXReplyData, notify *GXReplyData) bool {
	// If whole frame is not received yet.
	if buff.Size()-buff.Position() < 3 {
		data.isComplete = false
		return true
	}
	isData := true
	pos := buff.Position()
	data.isComplete = false
	if notify != nil {
		notify.isComplete = false
	}
	// Check SMS addresses.
	ret, err := checkSMSAddress(settings, buff, data, notify)
	if err != nil {
		return false
	}
	if !ret {
		data = notify
		isData = false
	}
	data.isComplete = buff.Available() != 0
	if !data.isComplete {
		buff.SetPosition(pos)
	} else {
		data.PacketLength = buff.Size()
	}
	return isData
}

func getCoAPValueAsInteger(buff *types.GXByteBuffer, len int) (uint64, error) {
	var token uint64
	switch len {
	case 1:
		v, err := buff.Uint8()
		if err != nil {
			return 0, err
		}
		token = uint64(v)
	case 2:
		v, err := buff.Uint16()
		if err != nil {
			return 0, err
		}
		token = uint64(v)
	case 4:
		v, err := buff.Uint32()
		if err != nil {
			return 0, err
		}
		token = uint64(v)
	case 8:
		v, err := buff.Uint64()
		if err != nil {
			return 0, err
		}
		token = uint64(v)
	default:
		return 0, errors.New("Invalid Coap data.")
	}
	return token, nil
}

// validateCheckSum returns the validate M-Bus checksum
func validateCheckSum(bb *types.GXByteBuffer, count int) bool {
	value := 0
	for pos := 0; pos != count; pos++ {
		v, err := bb.Uint8At(bb.Position() + pos)
		if err != nil {
			return false
		}
		value = value + int(v)
	}
	v, _ := bb.Uint8At(bb.Position() + count)
	return value == int(v)
}

// getWiredMBusData returns the get data from wired M-Bus frame.
//
// Parameters:
//
//	settings: DLMS settings.
//	buff: Received data.
//	data: Reply information.
func getWiredMBusData(settings *settings.GXDLMSSettings, buff *types.GXByteBuffer, data *GXReplyData) {
	packetStartID := buff.Position()
	ch, err := buff.Uint8()
	if err != nil {
		data.isComplete = false
		return
	}
	if ch != 0x68 || buff.Available() < 5 {
		data.isComplete = false
		buff.SetPosition(buff.Position() - 1)
	} else {
		// L-field.
		len, err := buff.Uint8()
		if err != nil {
			data.isComplete = false
			return
		}
		// L-field.
		ch, err := buff.Uint8()
		if err != nil {
			data.isComplete = false
			return
		}
		ch2, err := buff.Uint8()
		if err != nil {
			data.isComplete = false
			return
		}
		if ch != len || buff.Available() < 3+int(len) || ch2 != 0x68 {
			data.isComplete = false
			buff.SetPosition(packetStartID)
		} else {
			crc := validateCheckSum(buff, int(len))
			if !crc && data.xml == nil {
				data.isComplete = false
				buff.SetPosition(packetStartID)
			} else {
				if !crc {
					data.xml.AppendComment("Invalid checksum.")
				}
				// Check EOP.
				ch, err := buff.Uint8At(buff.Position() + int(len) + 1)
				if err != nil {
					data.isComplete = false
					return
				}
				if ch != 0x16 {
					data.isComplete = false
					buff.SetPosition(packetStartID)
					return
				}
				data.PacketLength = buff.Position() + int(len)
				data.isComplete = true
				// Control field (C-Field)
				tmp, err := buff.Uint8()
				cmd := constants.MBusCommand(tmp & 0xF)
				// Address (A-field)
				id, err := buff.Uint8()
				// The Control Information Field (CI-field)
				ci, err := buff.Uint8()
				if ci == 0x0 {
					data.moreData = enums.RequestTypesFrame
				} else if (ci >> 4) == (ci & 0xf) {
					data.moreData &= ^enums.RequestTypesFrame
				}
				// If M-Bus data header is present
				if ci != 0 {
				}
				if (tmp & 0x40) != 0 {

					v, err := buff.Uint8()
					if err != nil {
						data.isComplete = false
						return
					}
					settings.ClientAddress = int(v)
					v, err = buff.Uint8()
					if err != nil {
						data.isComplete = false
						return
					}
					settings.ServerAddress = int(v)
				} else {
					v, err := buff.Uint8()
					if err != nil {
						data.isComplete = false
						return
					}
					settings.ServerAddress = int(v)
					v, err = buff.Uint8()
					if err != nil {
						data.isComplete = false
						return
					}
					settings.ClientAddress = int(v)
				}
				if data.xml != nil && data.xml.Comments {
					data.xml.AppendComment("Command: " + cmd.String())
					data.xml.AppendComment("A-Field: " + strconv.Itoa(int(id)))
					data.xml.AppendComment("CI-Field: " + strconv.Itoa(int(ci)))
					if (tmp & 0x40) != 0 {
						data.xml.AppendComment("Primary station: " + strconv.Itoa(settings.ServerAddress))
						data.xml.AppendComment("Secondary station: " + strconv.Itoa(settings.ClientAddress))
					} else {
						data.xml.AppendComment("Primary station: " + strconv.Itoa(settings.ClientAddress))
						data.xml.AppendComment("Secondary station: " + strconv.Itoa(settings.ServerAddress))
					}
				}
			}
		}
	}
}

// getWirelessMBusData returns the get data from Wireless M-Bus frame.
//
// Parameters:
//
//	settings: DLMS settings.
//	buff: Received data.
//	data: Reply information.
func getWirelessMBusData(settings *settings.GXDLMSSettings, buff *types.GXByteBuffer, data *GXReplyData) error {
	// L-field.
	len_, err := buff.Uint8()
	if err != nil {
		return err
	}
	// Some meters are counting length to frame size.
	if buff.Size() < int(len_)-1 {
		data.isComplete = false
		buff.SetPosition(buff.Position() - 1)
	} else {
		// Some meters are counting length to frame size.
		if buff.Size() < int(len_) {
			len_--
		}
		data.PacketLength = int(len_)
		data.isComplete = true
		// C-field.
		ch, err := buff.Uint8()
		if err != nil {
			return err
		}
		cmd := constants.MBusCommand(ch & 0x4)
		// M-Field.
		manufacturerID, err := buff.Uint16()
		if err != nil {
			return err
		}
		man := internal.DecryptManufacturer(manufacturerID)
		// A-Field.
		buff.Uint32()

		meterVersion, err := buff.Uint8()
		if err != nil {
			return err
		}
		type_, err := buff.Uint8()
		if err != nil {
			return err
		}
		// CI-Field
		ci, err := buff.Uint8()
		if err != nil {
			return err
		}
		// Access number.
		_, err = buff.Uint8()
		if err != nil {
			return err
		}
		// State of the meter
		_, err = buff.Uint8()
		if err != nil {
			return err
		}
		// Configuration word.
		configurationWord, err := buff.Uint16()
		if err != nil {
			return err
		}
		encryption := constants.MBusEncryptionMode(configurationWord & 7)
		v, err := buff.Uint8()
		if err != nil {
			return err
		}
		settings.ClientAddress = int(v)
		v, err = buff.Uint8()
		if err != nil {
			return err
		}
		settings.ServerAddress = int(v)
		if data.xml != nil && data.xml.Comments {
			data.xml.AppendComment("Command: " + strconv.Itoa(int(cmd)))
			data.xml.AppendComment("Manufacturer: " + man)
			data.xml.AppendComment("Meter Version: " + strconv.Itoa(int(meterVersion)))
			data.xml.AppendComment("Meter Type: " + strconv.Itoa(int(type_)))
			data.xml.AppendComment("Control Info: " + strconv.Itoa(int(ci)))
			data.xml.AppendComment("Encryption: " + strconv.Itoa(int(encryption)))
		} else if settings.MBus != nil {
			settings.MBus.ManufacturerId = internal.DecryptManufacturer(manufacturerID)
			settings.MBus.Version = meterVersion
			settings.MBus.MeterType = constants.MBusMeterType(type_)
		}
	}
	return nil
}

// getPlcData returns the get data from S-FSK PLC frame.
//
// Parameters:
//
//	settings: DLMS settings.
//	buff: Received data.
//	data: Reply information.
func getPlcData(settings *settings.GXDLMSSettings, buff *types.GXByteBuffer, data *GXReplyData) error {
	if buff.Available() < 9 {
		data.isComplete = false
		return nil
	}
	packetStartID := buff.Position()
	// Find STX.
	for pos := buff.Position(); pos < buff.Size(); pos++ {
		stx, err := buff.Uint8()
		if err != nil {
			return err
		}
		if stx == 2 {
			packetStartID = pos
			break
		}
	}
	// Not a PLC frame.
	if buff.Position() == buff.Size() {
		data.isComplete = false
		buff.SetPosition(packetStartID)
		return nil
	}
	len, err := buff.Uint8()
	if err != nil {
		return err
	}
	index := buff.Position()
	if buff.Available() < int(len) {
		data.isComplete = false
		buff.SetPosition(buff.Position() - 2)
	} else {
		buff.Uint8()
		// Credit fields.  IC, CC, DC
		buff.Uint8()
		// MAC Addresses.
		mac, err := buff.Uint24()
		if err != nil {
			return err
		}
		// SA.
		macSa := (uint16)(mac >> 12)
		// DA.
		macDa := (uint16)(mac & 0xFFF)
		// PAD length.
		padLen, err := buff.Uint8()
		if err != nil {
			return err
		}
		if buff.Size() < int(len)+int(padLen)+2 {
			data.isComplete = false
			buff.SetPosition(index + 6)
		} else {
			// DL.Data.request
			ch, err := buff.Uint8()
			if err != nil {
				return err
			}
			if int(ch) != int(enums.PlcDataLinkDataRequest) {
				return errors.New("Parsing MAC LLC data failed. Invalid DataLink data request.")
			}
			//da
			_, err = buff.Uint8()
			if err != nil {
				return err
			}
			//sa
			_, err = buff.Uint8()
			if err != nil {
				return err
			}
			if settings.IsServer() {
				data.isComplete = (data.xml != nil || ((macDa == uint16(enums.PlcDestinationAddressAllPhysical) || macDa == settings.Plc.MacSourceAddress) && (macSa == uint16(enums.PlcSourceAddressInitiator) || macSa == settings.Plc.MacDestinationAddress)))
				data.SourceAddress = int(macDa)
				data.TargetAddress = int(macSa)
			} else {
				data.isComplete = (data.xml != nil || (macDa == uint16(enums.PlcDestinationAddressAllPhysical) || macDa == uint16(enums.PlcSourceAddressInitiator) || macDa == settings.Plc.MacDestinationAddress))
				data.TargetAddress = int(macDa)
				data.SourceAddress = int(macSa)
			}
			// Skip padding.
			if data.isComplete {
				crcCount := countFCS16(buff.Array(), 0, int(len+padLen))
				crc, err := buff.Uint16At(int(len + padLen))
				if err != nil {
					return err
				}
				// Check CRC.
				if crc != crcCount {
					if data.xml == nil {
						return errors.New("Invalid data checksum.")
					}
					data.xml.AppendComment("Invalid data checksum.")
				}
				data.PacketLength = int(len)
			} else {
				buff.SetPosition(packetStartID)
			}
		}
	}
	return nil
}

// getPlcHdlcData returns the get data from S-FSK PLC Hdlc frame.
//
// Parameters:
//
//	settings: DLMS settings.
//	buff: Received data.
//	data: Reply information.
func getPlcHdlcData(settings *settings.GXDLMSSettings, buff *types.GXByteBuffer, data *GXReplyData) (uint8, error) {
	if buff.Available() < 2 {
		data.isComplete = false
		return 0, nil
	}
	var frame uint8
	// SN field.
	frameLen := getPlcSfskFrameSize(buff)
	if frameLen == 0 {
		return 0, errors.New("Invalid PLC frame size.")
	}
	if buff.Available() < int(frameLen) {
		data.isComplete = false
	} else {
		buff.SetPosition(2)
		index := buff.Position()
		// Credit fields.  IC, CC, DC
		_, err := buff.Uint8()
		if err != nil {
			return 0, err
		}
		// MAC Addresses.
		mac, err := buff.Uint24()
		// SA.
		sa := (mac >> 12)
		// DA.
		da := (mac & 0xFFF)
		if settings.IsServer() {
			data.isComplete = data.xml != nil || (da == int(enums.PlcDestinationAddressAllPhysical) || da == int(settings.Plc.MacSourceAddress) && (sa == int(enums.PlcHdlcSourceAddressInitiator)) || sa == int(settings.Plc.MacDestinationAddress))
			data.SourceAddress = da
			data.TargetAddress = sa
		} else {
			data.isComplete = data.xml != nil || (da == int(enums.PlcHdlcSourceAddressInitiator) || da == int(settings.Plc.MacDestinationAddress))
			data.TargetAddress = da
			data.SourceAddress = sa
		}
		if data.isComplete {
			// PAD length.
			padLen, err := buff.Uint8()
			if err != nil {
				return 0, err
			}
			frame, err = getHdlcData(settings.IsServer(), settings, buff, data, nil)
			if err != nil {
				return 0, err
			}
			getDataFromFrame(buff, data, enums.InterfaceTypeHDLC)
			buff.SetPosition(int(padLen))
			crcCount := countFCS24(buff.Array(), index, buff.Position()-index)
			crc, err := buff.Uint24At(buff.Position())
			if err != nil {
				return 0, err
			}
			// Check CRC.
			if crc != int(crcCount) {
				if data.xml == nil {
					return 0, errors.New("Invalid data checksum.")
				}
				data.xml.AppendComment("Invalid data checksum.")
			}
			data.PacketLength = 2 + buff.Position() - index + 3
		} else {
			buff.SetPosition(int(frameLen) - index - 4)
		}
	}
	return frame, nil
}

// readResponseDataBlockResult returns the handle read response data block result.
//
// Parameters:
//
//	settings: DLMS settings.
//	reply: Received reply.
//	index: Starting index.
func readResponseDataBlockResult(settings *settings.GXDLMSSettings, reply *GXReplyData, index int) (bool, error) {
	reply.Error = 0
	lastBlock, err := reply.Data.Uint8()
	if err != nil {
		return false, err
	}
	number, err := reply.Data.Uint16()
	if err != nil {
		return false, err
	}
	blockLength, err := types.GetObjectCount(reply.Data)
	if err != nil {
		return false, err
	}
	// Is not Last block.
	if lastBlock == 0 {
		reply.moreData |= enums.RequestTypesDataBlock
	} else {
		reply.moreData &= ^enums.RequestTypesDataBlock
	}
	// If meter's block index is zero based.
	if number != 1 && settings.BlockIndex == 1 {
		settings.BlockIndex = uint32(number)
	}
	expectedIndex := settings.BlockIndex
	if uint32(number) != expectedIndex {
		return false, fmt.Errorf("Invalid Block number. It is %d and it should be %d.", number, expectedIndex)
	}
	// If whole block is not read.
	if (reply.moreData & enums.RequestTypesFrame) != 0 {
		getDataFromBlock(reply.Data, index)
		return false, nil
	}
	// Check block length when all data is received.
	if blockLength != reply.Data.Size()-reply.Data.Position() {
		return false, errors.New("Invalid block length.")
	}
	reply.command = enums.CommandNone
	if reply.xml != nil {
		reply.Data.Trim()
		reply.xml.AppendStartTag(enums.CommandReadResponse<<8|int(constants.SingleReadResponseDataBlockResult), "", "", false)
		reply.xml.AppendLine(internal.TranslatorTagsLastBlock.String(), "Value", reply.xml.IntegerToHex(lastBlock, 2, false))
		reply.xml.AppendLine(internal.TranslatorTagsBlockNumber.String(), "Value", reply.xml.IntegerToHex(number, 4, false))
		reply.xml.AppendLine(internal.TranslatorTagsRawData.String(), "Value", buffer.ToHexWithRange(reply.Data.Array(), false, 0, reply.Data.Size()))
		reply.xml.AppendEndTag(enums.CommandReadResponse<<8|int(constants.SingleReadResponseDataBlockResult), false)
		return false, nil
	}
	getDataFromBlock(reply.Data, index)
	reply.TotalCount = 0
	// If last packet and data is not try to peek.
	if reply.moreData == enums.RequestTypesNone {
		settings.ResetBlockIndex()
	}
	return true, nil
}

// handleReadResponse returns the handle read response and get data from block and/or update error status.
//
// Parameters:
//
//	reply: Received data from the client.
func handleReadResponse(settings *settings.GXDLMSSettings, reply *GXReplyData, index int) (bool, error) {
	cnt := reply.TotalCount
	// If we are reading value first time or block is handed.
	first := cnt == 0 || reply.commandType == uint8(constants.SingleReadResponseDataBlockResult)
	if first {
		cnt, err := types.GetObjectCount(reply.Data)
		if err != nil {
			return false, err
		}
		reply.TotalCount = cnt
	}
	var type_ constants.SingleReadResponse
	var values []any
	if cnt != 1 {
		// Parse data after all data is received when readlist is used.
		if reply.IsMoreData() {
			getDataFromBlock(reply.Data, 0)
			return false, nil
		}
		if !first {
			reply.Data.SetPosition(0)
		}
		values = []any{}
		if _, ok := reply.Value.([]any); ok {
			values = reply.Value.([]any)
		}
		reply.Value = nil
	}
	if reply.xml != nil {
		reply.xml.AppendStartTag(enums.CommandReadResponse, "Qty", reply.xml.IntegerToHex(cnt, 2, false), false)
	}
	standardXml := reply.xml != nil && reply.xml.OutputType() == enums.TranslatorOutputTypeStandardXML
	for pos := 0; pos != cnt; pos++ {
		// Get response type code.
		if first {
			ch, err := reply.Data.Uint8()
			if err != nil {
				return false, err
			}
			type_ = constants.SingleReadResponse(ch)
			reply.commandType = uint8(type_)
		} else {
			type_ = constants.SingleReadResponse(reply.commandType)
		}
		switch type_ {
		case constants.SingleReadResponseData:
			reply.Error = 0
			if reply.xml != nil {
				if standardXml {
					reply.xml.AppendStartTag(int(internal.TranslatorTagsChoice), "", "", false)
				}
				reply.xml.AppendStartTag(int(enums.CommandReadResponse)<<8|int(constants.SingleReadResponseData), "", "", false)
				di := internal.GXDataInfo{}
				di.Xml = reply.xml
				_, err := internal.GetData(settings, reply.Data, &di)
				if err != nil {
					return false, err
				}
				reply.xml.AppendEndTag(int(enums.CommandReadResponse)<<8|int(constants.SingleReadResponseData), false)
				if standardXml {
					reply.xml.AppendEndTag(int(internal.TranslatorTagsChoice), false)
				}
			} else if cnt == 1 {
				getDataFromBlock(reply.Data, 0)
			} else {
				reply.ReadPosition = reply.Data.Position()
				getValueFromData(settings, reply)
				reply.Data.SetPosition(reply.ReadPosition)
				values = append(values, reply.Value)
				reply.Value = nil
			}
		case constants.SingleReadResponseDataAccessError:
			// Get error code.
			v, err := reply.Data.Uint8()
			if err != nil {
				return false, err
			}
			reply.Error = int(v)
			if reply.xml != nil {
				if standardXml {
					reply.xml.AppendStartTag(int(internal.TranslatorTagsChoice), "", "", false)
				}
				ret, err := ErrorCodeToString(reply.xml.OutputType(), enums.ErrorCode(reply.Error))
				if err != nil {
					return false, err
				}
				reply.xml.AppendLine(reply.xml.GetTag(int(enums.CommandReadResponse)<<8|int(constants.SingleReadResponseDataAccessError)),
					"",
					ret)
				if standardXml {
					reply.xml.AppendEndTag(int(internal.TranslatorTagsChoice), false)
				}
			}
		case constants.SingleReadResponseDataBlockResult:
			ret, err := readResponseDataBlockResult(settings, reply, index)
			if err != nil {
				return false, err
			}
			if !ret {
				// If xml only received bytes are shown. Data is not try to parse.
				if reply.xml != nil {
					reply.xml.AppendEndTag(int(enums.CommandReadResponse), false)
				}
				return false, nil
			}
		case constants.SingleReadResponseBlockNumber:
			// Get Block number.
			number, err := reply.Data.Uint16()
			if err != nil {
				return false, err
			}
			if uint32(number) != settings.BlockIndex {
				return false, fmt.Errorf("Invalid Block number. It is %d and it should be %d.", number, settings.BlockIndex)
			}
			settings.IncreaseBlockIndex()
			reply.moreData = (enums.RequestTypes)(reply.moreData | enums.RequestTypesDataBlock)
		default:
			return false, errors.New("handleReadResponse failed. Invalid tag.")
		}
	}
	if reply.xml != nil {
		reply.xml.AppendEndTag(int(enums.CommandReadResponse), false)
		return true, nil
	}
	if values != nil {
		reply.Value = values
	}
	return cnt == 1, nil
}

func handleActionResponseNormal(settings *settings.GXDLMSSettings, data *GXReplyData) (bool, error) {
	ret, err := data.Data.Uint8()
	if err != nil {
		return false, err
	}
	if ret != 0 {
		data.Error = int(ret)
	}
	if data.xml != nil {
		if data.xml.OutputType() == enums.TranslatorOutputTypeStandardXML {
			data.xml.AppendStartTag(int(internal.TranslatorTagsSingleResponse), "", "", false)
		}
		ret, err := ErrorCodeToString(data.xml.OutputType(), enums.ErrorCode(data.Error))
		if err != nil {
			return false, err
		}
		data.xml.AppendLine(data.xml.GetTag(int(internal.TranslatorTagsResult)), "", ret)
	}
	settings.ResetBlockIndex()
	// Response normal. Get data if exists. Some meters do not return here anything.
	if data.Error == 0 && data.Data.Position() < data.Data.Size() {
		ret, err := data.Data.Uint8()
		if err != nil {
			return false, err
		}
		// If data.
		if ret == 0 {
			getDataFromBlock(data.Data, 0)
		} else if ret == 1 {
			ret, err = data.Data.Uint8()
			if err != nil {
				return false, err
			}
			if ret != 0 {
				ret, err = data.Data.Uint8()
				if err != nil {
					return false, err
				}
				data.Error = int(ret)
				// Handle Texas Instrument missing byte here.
				if ret == 9 && data.Error == 16 {
					data.Data.SetPosition(2)
					getDataFromBlock(data.Data, 0)
					data.Error = 0
					ret = 0
				}
			} else {
				getDataFromBlock(data.Data, 0)
			}
		} else {
			return false, errors.New("handleActionResponseNormal failed. Invalid tag.")
		}
		if data.xml != nil && (ret != 0 || data.Data.Position() < data.Data.Size()) {
			data.xml.AppendStartTag(int(internal.TranslatorTagsReturnParameters), "", "", false)
			if ret != 0 {
				ret, err := ErrorCodeToString(data.xml.OutputType(), enums.ErrorCode(data.Error))
				if err != nil {
					return false, err
				}
				data.xml.AppendLine(data.xml.GetTag(int(internal.TranslatorTagsDataAccessError)), "", ret)
			} else {
				data.xml.AppendStartTag(int(internal.TranslatorTagsData), "", "", false)
				di := internal.GXDataInfo{}
				di.Xml = data.xml
				internal.GetData(settings, data.Data, &di)
				data.xml.AppendEndTag(int(internal.TranslatorTagsData), false)
			}
			data.xml.AppendEndTag(int(internal.TranslatorTagsReturnParameters), false)
			if data.xml.OutputType() == enums.TranslatorOutputTypeStandardXML {
				data.xml.AppendEndTag(int(internal.TranslatorTagsSingleResponse), false)
			}
		}
	}
	return true, nil
}

func handleActionResponseWithBlock(settings *settings.GXDLMSSettings, reply *GXReplyData, index int) (bool, error) {
	ret := true
	ch, err := reply.Data.Uint8()
	if err != nil {
		return false, err
	}
	if reply.xml != nil {
		reply.xml.AppendStartTag(int(internal.TranslatorTagsPblock), "", "", false)
		reply.xml.AppendLine(reply.xml.GetTag(int(internal.TranslatorTagsLastBlock)), "Value", reply.xml.IntegerToHex(ch, 2, false))
	}
	if ch == 0 {
		reply.moreData |= enums.RequestTypesDataBlock
	} else {
		reply.moreData &= ^enums.RequestTypesDataBlock
	}
	number, err := reply.Data.Uint32()
	if err != nil {
		return false, err
	}
	if reply.xml != nil {
		reply.xml.AppendLine(reply.xml.GetTag(int(internal.TranslatorTagsBlockNumber)), "Value", reply.xml.IntegerToHex(number, 8, false))
	} else {
		// Update  initial block index. This is critical if message is send and received in multiple blocks.
		if number == 1 {
			settings.ResetBlockIndex()
		}
		if number != settings.BlockIndex {
			return false, fmt.Errorf("Invalid Block number. It is %d and it should be %d.", number, settings.BlockIndex)
		}
	}
	// Note! There is no status!!
	if reply.xml != nil {
		if reply.Data.Available() != 0 {
			// Get data size.
			blockLength, err := types.GetObjectCount(reply.Data)
			if err != nil {
				return false, err
			}
			// if whole block is read.
			if (reply.moreData & enums.RequestTypesFrame) == 0 {
				// Check Block length.
				if blockLength > reply.Data.Size()-reply.Data.Position() {
					reply.xml.AppendComment(fmt.Sprintf("Block is not complete %d/%d.", reply.Data.Size()-reply.Data.Position(), blockLength))
				}
			}
			reply.xml.AppendLine(reply.xml.GetTag(int(internal.TranslatorTagsRawData)), "Value", buffer.ToHexWithRange(reply.Data.Array(), false, reply.Data.Position(), reply.Data.Available()))
		}
		reply.xml.AppendEndTag(int(internal.TranslatorTagsPblock), false)
	} else if reply.Data.Available() != 0 {
		// Get data size.
		blockLength, err := types.GetObjectCount(reply.Data)
		if err != nil {
			return false, err
		}
		// if whole block is read.
		if (reply.moreData & enums.RequestTypesFrame) == 0 {
			// Check Block length.
			if blockLength > reply.Data.Available() {
				return false, errors.New("out of memory")
			}
			// Keep command if this is last block for XML Client.
			if (reply.moreData & enums.RequestTypesDataBlock) != 0 {
				reply.command = enums.CommandNone
			}
		}
		if blockLength == 0 {
			reply.Data.SetSize(index)
		} else {
			getDataFromBlock(reply.Data, index)
		}
		// If last packet and data is not try to peek.
		if reply.moreData == enums.RequestTypesNone {
			if !reply.Peek {
				reply.Data.SetPosition(0)
			}
			settings.ResetBlockIndex()
		}
	} else if reply.Data.Position() == reply.Data.Size() {
		reply.EmptyResponses = enums.RequestTypesDataBlock
	}
	if reply.moreData == enums.RequestTypesNone && settings != nil && settings.Command == enums.CommandMethodRequest && settings.CommandType == uint8(constants.ActionResponseTypeWithList) {
		return false, errors.New("not implemented")
	}
	return ret, nil
}

// handleMethodResponse returns the handle method response and get data from block and/or update error status.
//
// Parameters:
//
//	settings: DLMS settings.
//	data: Received data from the client.
func handleMethodResponse(settings *settings.GXDLMSSettings, data *GXReplyData, index int) error {
	// Get type.
	ret, err := data.Data.Uint8()
	if err != nil {
		return err
	}
	type_ := constants.ActionResponseType(ret)
	ret, err = data.Data.Uint8()
	if err != nil {
		return err
	}
	data.invokeId = uint32(ret)
	verifyInvokeId(settings, data)
	addInvokeId(data.xml, enums.CommandMethodResponse, int(type_), data.invokeId)
	switch type_ {
	case constants.ActionResponseTypeNormal:
		handleActionResponseNormal(settings, data)
	case constants.ActionResponseTypeWithBlock:
		handleActionResponseWithBlock(settings, data, index)
	case constants.ActionResponseTypeWithList:
		return errors.New("Invalid enums.Command")
	case constants.ActionResponseTypeNextBlock:
		number, err := data.Data.Uint32()
		if err != nil {
			return err
		}
		if data.xml != nil {
			data.xml.AppendLine(data.xml.GetTag(int(internal.TranslatorTagsBlockNumber)), "Value", data.xml.IntegerToHex(number, 8, false))
		} else if number != settings.BlockIndex {
			return fmt.Errorf("Invalid Block number. It is %d and it should be %d.", number, settings.BlockIndex)
		}
		settings.IncreaseBlockIndex()
	default:
		return errors.New("Invalid enums.Command")
	}
	if data.xml != nil {
		data.xml.AppendEndTag(int(enums.CommandMethodResponse)<<8|int(type_), false)
		data.xml.AppendEndTag(int(enums.CommandMethodResponse), false)
	}
	return err
}

func handleSetResponse(settings *settings.GXDLMSSettings, data *GXReplyData) error {
	ret, err := data.Data.Uint8()
	if err != nil {
		return err
	}
	type_ := constants.SetResponseType(ret)
	ch, err := data.Data.Uint8()
	if err != nil {
		return err
	}
	data.invokeId = uint32(ch)
	verifyInvokeId(settings, data)
	addInvokeId(data.xml, enums.CommandSetResponse, int(type_), data.invokeId)
	// SetResponseNormal
	if type_ == constants.SetResponseTypeNormal {
		ch, err = data.Data.Uint8()
		if err != nil {
			return err
		}
		data.Error = int(ch)
		if data.xml != nil {
			ret, err := ErrorCodeToString(data.xml.OutputType(), enums.ErrorCode(data.Error))
			if err != nil {
				return err
			}
			data.xml.AppendLine(data.xml.GetTag(int(internal.TranslatorTagsResult)), "Value", ret)
		}
	} else if type_ == constants.SetResponseTypeDataBlock {
		number, err := data.Data.Uint32()
		if err != nil {
			return err
		}
		if data.xml != nil {
			data.xml.AppendLine(data.xml.GetTag(int(internal.TranslatorTagsBlockNumber)), "Value", data.xml.IntegerToHex(number, 8, false))
		}
	} else if type_ == constants.SetResponseTypeLastDataBlock {
		ch, err := data.Data.Uint8()
		if err != nil {
			return err
		}
		data.Error = int(ch)
		number, err := data.Data.Uint32()
		if err != nil {
			return err
		}
		if data.xml != nil {
			ret, err := ErrorCodeToString(data.xml.OutputType(), enums.ErrorCode(data.Error))
			if err != nil {
				return err
			}
			data.xml.AppendLine(data.xml.GetTag(int(internal.TranslatorTagsResult)), "Value", ret)
			data.xml.AppendLine(data.xml.GetTag(int(internal.TranslatorTagsBlockNumber)), "Value", data.xml.IntegerToHex(number, 8, false))
		}
	} else if type_ == constants.SetResponseTypeWithList {
		cnt, err := types.GetObjectCount(data.Data)
		if err != nil {
			return err
		}
		if data.xml != nil {
			data.xml.AppendStartTag(int(internal.TranslatorTagsResult), "Qty", fmt.Sprint(cnt), false)
			for pos := 0; pos != cnt; pos++ {
				ret, err := data.Data.Uint8()
				if err != nil {
					return err
				}
				ret2, err := ErrorCodeToString(data.xml.OutputType(), enums.ErrorCode(int(ret)))
				if err != nil {
					return err
				}
				data.xml.AppendLine(data.xml.GetTag(int(internal.TranslatorTagsDataAccessResult)), "Value", ret2)
			}
			data.xml.AppendEndTag(int(internal.TranslatorTagsResult), false)
		} else {
			for pos := 0; pos != cnt; pos++ {
				ret, err := data.Data.Uint8()
				if err != nil {
					return err
				}
				if data.Error == 0 && ret != 0 {
					data.Error = int(ret)
				}
			}
		}
	} else {
		return errors.New("Invalid data type.")
	}
	if data.xml != nil {
		data.xml.AppendEndTag(int(enums.CommandSetResponse)<<8|int(type_), false)
		data.xml.AppendEndTag(int(enums.CommandSetResponse), false)
	}
	return nil
}

// handleWriteResponse returns the handle write response and get data from block.
//
// Parameters:
//
//	reply: Received data from the client.
func handleWriteResponse(data *GXReplyData) error {
	cnt, err := types.GetObjectCount(data.Data)
	if err != nil {
		return err
	}
	var ret byte
	if data.xml != nil {
		data.xml.AppendStartTag(int(enums.CommandWriteResponse), "Qty", data.xml.IntegerToHex(cnt, 2, false), false)
	}
	for pos := 0; pos != cnt; pos++ {
		ret, err = data.Data.Uint8()
		if err != nil {
			return err
		}
		if ret != 0 {
			ch, err := data.Data.Uint8()
			if err != nil {
				return err
			}
			data.Error = int(ch)
		}
		if data.xml != nil {
			if ret == 0 {
				data.xml.AppendLine("<"+enums.ErrorCode(ret).String()+" />", "", "")
			} else {
				str_ := enums.ErrorCode(data.Error).String()
				data.xml.AppendLineFromTag(int(internal.TranslatorTagsDataAccessError), "Value", str_)
			}
		}
	}
	if data.xml != nil {
		data.xml.AppendEndTag(int(enums.CommandWriteResponse), false)
	}
	return nil
}

// handleGetResponseWithList returns the handle get response with list.
//
// Parameters:
//
//	settings: DLMS settings.
//	reply: Received data from the client.
func handleGetResponseWithList(settings *settings.GXDLMSSettings, reply *GXReplyData) error {
	// Get object count.
	cnt, err := types.GetObjectCount(reply.Data)
	if err != nil {
		return err
	}
	values := make([]any, 0, cnt)
	if reply.xml != nil {
		reply.xml.AppendStartTag(int(internal.TranslatorTagsResult), "Qty", reply.xml.IntegerToHex(cnt, 2, false), false)
	}
	for pos := 0; pos != cnt; pos++ {
		ch, err := reply.Data.Uint8()
		if err != nil {
			return err
		}
		if ch != 0 {
			ch, err = reply.Data.Uint8()
			if err != nil {
				return err
			}
			reply.Error = int(ch)
		} else {
			reply.ReadPosition = reply.Data.Position()
			if reply.xml != nil {
				di := internal.GXDataInfo{}
				di.Xml = reply.xml
				di.Xml.AppendStartTag(int(enums.CommandReadResponse)<<8|int(constants.SingleReadResponseData), "", "", false)
				internal.GetData(settings, reply.Data, &di)
				di.Xml.AppendEndTag(int(enums.CommandReadResponse)<<8|int(constants.SingleReadResponseData), false)
				reply.ReadPosition = reply.Data.Position()
			} else {
				getValueFromData(settings, reply)
				values = append(values, reply.Value)
				reply.Value = nil
			}
			reply.Data.SetPosition(reply.ReadPosition)
		}
	}
	reply.Value = values
	return nil
}

func verifyInvokeId(settings *settings.GXDLMSSettings, reply *GXReplyData) error {
	if reply.xml == nil && settings.AutoIncreaseInvokeID && uint8(reply.invokeId) != getInvokeIDPriority(settings, false) {
		return errors.New("Invalid invoke ID. Expected: " + fmt.Sprintf("%X", getInvokeIDPriority(settings, false)) + fmt.Sprintf(" Actual: %X", reply.invokeId))
	}
	return nil
}

// handleGetResponse returns the handle get response and get data from block and/or update error status.
//
// Parameters:
//
//	settings: DLMS settings.
//	reply: Received data from the client.
//	index: Block index number.
func handleGetResponse(settings *settings.GXDLMSSettings, reply *GXReplyData, index int) (bool, error) {
	ret := true
	data := reply.Data
	empty := false
	// Get type.
	ch, err := data.Uint8()
	if err != nil {
		return false, err
	}
	type_ := constants.GetCommandType(ch)
	reply.commandType = ch
	ch, err = data.Uint8()
	if err != nil {
		return false, err
	}
	reply.invokeId = uint32(ch)
	verifyInvokeId(settings, reply)
	addInvokeId(reply.xml, enums.CommandGetResponse, int(type_), reply.invokeId)
	switch type_ {
	case constants.GetCommandTypeNormal:
		empty, err = handleGetResponseNormal(settings, reply, data)
		if err != nil {
			return false, err
		}
	case constants.GetCommandTypeNextDataBlock:
		ret, err = handleGetResponseNextDataBlock(settings, reply, index, data)
		if err != nil {
			return false, err
		}
	case constants.GetCommandTypeWithList:
		if !reply.IsMoreData() {
			handleGetResponseWithList(settings, reply)
		}
		ret = false
	default:
		return false, errors.New("Invalid Get response.")
	}
	if reply.xml != nil {
		if !empty {
			reply.xml.AppendEndTag(int(internal.TranslatorTagsResult), false)
		}
		reply.xml.AppendEndTag(int(enums.CommandGetResponse)<<8|int(type_), false)
		reply.xml.AppendEndTag(int(enums.CommandGetResponse), false)
	}
	return ret, nil
}

func handleGetResponseNextDataBlock(settings *settings.GXDLMSSettings, reply *GXReplyData, index int, data *types.GXByteBuffer) (bool, error) {
	ret := true
	ch, err := data.Uint8()
	if err != nil {
		return false, err
	}
	if reply.xml != nil {
		reply.xml.AppendStartTag(int(internal.TranslatorTagsResult), "", "", false)
		reply.xml.AppendLineFromTag(int(internal.TranslatorTagsLastBlock), "Value", reply.xml.IntegerToHex(ch, 2, false))
	}
	if ch == 0 {
		reply.moreData |= enums.RequestTypesDataBlock
	} else {
		reply.moreData &^= enums.RequestTypesDataBlock
	}
	number, err := data.Uint32()
	if err != nil {
		return false, err
	}
	if reply.xml != nil {
		reply.xml.AppendLineFromTag(int(internal.TranslatorTagsBlockNumber), "Value", reply.xml.IntegerToHex(number, 8, false))
	} else {
		// If meter's block index is zero based.
		if number != 1 && settings.BlockIndex == 1 {
			settings.BlockIndex = number
		}
		if number != settings.BlockIndex {
			return false, fmt.Errorf("Invalid Block number. It is %d and it should be %d.", number, settings.BlockIndex)
		}
	}
	ch, err = data.Uint8()
	if err != nil {
		return false, err
	}
	if ch != 0 {
		ch, err = data.Uint8()
		if err != nil {
			return false, err
		}
		reply.Error = int(ch)
	}
	if reply.xml != nil {
		reply.xml.AppendStartTag(int(internal.TranslatorTagsResult), "", "", false)
		if reply.Error != 0 {
			ret, err := ErrorCodeToString(reply.xml.OutputType(), enums.ErrorCode(reply.Error))
			if err != nil {
				return false, err
			}
			reply.xml.AppendLineFromTag(int(internal.TranslatorTagsDataAccessResult), "Value", ret)
		} else if reply.Data.Available() != 0 {
			// Get data size.
			blockLength, err := types.GetObjectCount(data)
			if err != nil {
				return false, err
			}
			// if whole block is read.
			if (reply.moreData & enums.RequestTypesFrame) == 0 {
				// Check Block length.
				if blockLength > data.Available() {

					reply.xml.AppendComment(fmt.Sprintf("Block is not complete %d/%d.", data.Available(), blockLength))
				}
			}
			reply.xml.AppendLineFromTag(int(internal.TranslatorTagsRawData), "Value", buffer.ToHexWithRange(reply.Data.Array(), false, data.Position(), reply.Data.Available()))
		}
		reply.xml.AppendEndTag(int(internal.TranslatorTagsResult), false)
	} else if reply.Data.Available() != 0 {
		// Get data size.
		blockLength, err := types.GetObjectCount(data)
		if err != nil {
			return false, err
		}
		// if whole block is read.
		if (reply.moreData & enums.RequestTypesFrame) == 0 {
			// Check Block length.
			if blockLength > data.Available() {
				return false, errors.New("Block length is greater than available data.")
			}
			// Keep command if this is last block for XML Client.
			if (reply.moreData & enums.RequestTypesDataBlock) != 0 {
				reply.command = enums.CommandNone
			}
		}
		if blockLength == 0 {
			data.SetSize(index)
		} else {
			getDataFromBlock(data, index)
		}
		// If last packet and data is not try to peek.
		if reply.moreData == enums.RequestTypesNone {
			if !reply.Peek {
				data.SetPosition(0)
			}
			settings.ResetBlockIndex()
		}
	} else if data.Position() == data.Size() {
		reply.EmptyResponses = enums.RequestTypesDataBlock
	}
	if reply.moreData == enums.RequestTypesNone && settings != nil && settings.Command == enums.CommandGetRequest && settings.CommandType == uint8(constants.GetCommandTypeWithList) {
		handleGetResponseWithList(settings, reply)
		ret = false
	}
	return ret, nil
}

func handleGetResponseNormal(settings *settings.GXDLMSSettings, reply *GXReplyData, data *types.GXByteBuffer) (bool, error) {
	empty := false
	if data.Available() == 0 {
		empty = true
		getDataFromBlock(data, 0)
	} else {
		// Result
		ch, err := data.Uint8()
		if err != nil {
			return false, err
		}
		if ch != 0 {
			ch, err := data.Uint8()
			if err != nil {
				return false, err
			}
			reply.Error = int(ch)
		}
		if reply.xml != nil {
			reply.xml.AppendStartTag(int(internal.TranslatorTagsResult), "", "", false)
			if reply.Error != 0 {
				ret, err := ErrorCodeToString(reply.xml.OutputType(), enums.ErrorCode(reply.Error))
				if err != nil {
					return false, err
				}
				reply.xml.AppendLineFromTag(int(internal.TranslatorTagsDataAccessError), "Value", ret)
			} else {
				reply.xml.AppendStartTag(int(internal.TranslatorTagsData), "", "", false)
				di := internal.GXDataInfo{}
				di.Xml = reply.xml
				internal.GetData(settings, reply.Data, &di)
				reply.xml.AppendEndTag(int(internal.TranslatorTagsData), false)
			}
		} else {
			getDataFromBlock(data, 0)
		}
	}
	return empty, nil
}

func handleGloDedRequest(settings *settings.GXDLMSSettings, data *GXReplyData) error {
	if data.xml != nil && !data.xml.Comments {
		data.Data.SetPosition(data.Data.Position() - 1)
	} else {
		if settings.Cipher == nil {
			return errors.New("Secure connection is not supported.")
		}
		// If all frames are read.
		if (data.moreData & enums.RequestTypesFrame) == 0 {
			data.Data.SetPosition(data.Data.Position() - 1)
			/*
				TODO:
					pos := data.Data.Position()
					encrypted := data.Data.Array()
					ret := settings.Crypt(enums.CertificateTypeDigitalSignature, encrypted, false, enums.CryptoKeyTypeEcdsa)
					if ret != nil {
						encrypted = data.Data.Array()
						data.Data.SetSize(0)
						data.Data.Set(ret)
					} else {
						p := AesGcmParameter{}
						cipher = settings
						if data.Command() == enums.CommandGeneralDedCiphering {
							p = settings.AesGcmParameter(settings, settings.SourceSystemTitle(), cipher.dedicatedKey(), getAuthenticationKey(settings))
						} else if data.Command() == enums.CommandGeneralGloCiphering {
							p = settings.AesGcmParameter(settings, settings.SourceSystemTitle(), getBlockCipherKey(settings), getAuthenticationKey(settings))
						} else if cipher.dedicatedKey() == nil || isGloMessage(data.Command()) {
							p = settings.AesGcmParameter(settings, settings.SourceSystemTitle(), getBlockCipherKey(settings), getAuthenticationKey(settings))
						} else {
							p = settings.AesGcmParameter(settings, settings.SourceSystemTitle(), cipher.dedicatedKey(), getAuthenticationKey(settings))
						}
						tmp := secure.AesEncrypt(p, data.Data())
						cipher.securitySuite = p.SecuritySuite
						data.Data.Clear()
						data.Data.Set(tmp)
					}
					data.cipheredCommand = data.command
					ch, err := data.Data.Uint8()
					if err != nil {
						return err
					}
					data.command = enums.Command(ch)
					if data.command == enums.CommandDataNotification || data.command == enums.CommandInformationReport || data.command == enums.CommandGatewayRequest || data.command == enums.CommandGatewayResponse {
						data.command = enums.CommandNone
						data.Data.SetPosition(data.Data.Position() - 1)
						GetPdu(settings, data)
					}*/
		} else {
			data.Data.SetPosition(data.Data.Position() - 1)
		}
	}
	return nil
}

func handleGloDedResponse(settings *settings.GXDLMSSettings, data *GXReplyData, index int) error {
	/*
		if data.xml != nil && !data.xml.Comments {
			data.Data.Position()--
		} else {
			if settings.Cipher() == nil {
				return errors.New("Secure connection is not supported.")
			}
			// If all frames are read.
			if (data.moreData & enums.RequestTypesFrame) == 0 {
				data.Data.Position()--
				bb := types.GXByteBuffer{}
				bb.Set(data.Data.data[data.Data.Position():data.Data.Size()])
				data.Data.SetPosition(data.Data.Position() - index)
				// If external Hardware Security Module is used.
				ret := settings.Crypt(enums.CertificateTypeDigitalSignature, bb.Array(), false, enums.CryptoKeyTypeEcdsa)
				if ret != nil {
					data.Data.Set(ret)
				} else {
					var p settings.AesGcmParameter
					cipher = settings
					if cipher.dedicatedKey() != nil && (settings.connected&enums.ConnectionStateDlms) != 0 {
						p = settings.AesGcmParameter(settings, settings.SourceSystemTitle(), cipher.dedicatedKey(), getAuthenticationKey(settings))
					} else {
						if settings.preEstablishedSystemTitle != nil && (settings.connected&enums.ConnectionStateDlms) == 0 {
							p = settings.AesGcmParameter(settings, settings.preEstablishedSystemTitle, getBlockCipherKey(settings), getAuthenticationKey(settings))
						} else {
							if settings.SourceSystemTitle() == nil && (settings.connected&enums.ConnectionStateDlms) != 0 {
								if settings.IsServer() {
									return 0, errors.New("Ciphered failed. Client system title is unknown.")
								} else {
									return "", errors.New("Ciphered failed. Server system title is unknown.")
								}
							}
							p = settings.AesGcmParameter(settings, settings.SourceSystemTitle(), getBlockCipherKey(settings), getAuthenticationKey(settings))
						}
					}
					tmp := GXCiphering.Decrypt(p, bb)
					data.SetSystemTitle(p.SystemTitle())
					data.Data().Set(tmp)
					// If target is sending data ciphered using different security policy.
					if !settings.Cipher().SecurityChangeCheck() && (settings.connected&enums.ConnectionStateDlms) != 0 && settings.Cipher().Security() != enums.SecurityNone && settings.Cipher().Signing() != enums.SigningGeneralSigning && settings.Cipher().Security() != p.Security() {
						return errors.New(string.Format("Data is ciphered using different security level. Actual: {0}. Expected: {1}", p.Security, settings.Cipher.Security))
					}
					if settings.ExpectedInvocationCounter() != 0 {
						if p.InvocationCounter() < settings.ExpectedInvocationCounter() {
							return errors.New(string.Format("Data is ciphered using invalid invocation counter value. Actual: {0}. Expected: {1}", p.InvocationCounter, settings.ExpectedInvocationCounter()))
						}
						settings.SetExpectedInvocationCounter(p.InvocationCounter())
					}
				}
				data.CipheredCommand(data.Command())
				data.Command(enums.CommandNone)
				GetPdu(settings, data)
				data.SetCipherIndex(data.Data().Size())
			}
		}
	*/
	return nil
}

func handleAccessResponse(settings *settings.GXDLMSSettings, reply *GXReplyData) error {
	invokeId, err := reply.Data.Uint32()
	if err != nil {
		return err
	}
	reply.Time = time.Time{}
	len, err := reply.Data.Uint8()
	if err != nil {
		return err
	}
	var tmp []byte
	// If date time is given.
	if len != 0 {
		tmp = make([]byte, len)
		reply.Data.Get(tmp)
		ret, err := internal.ChangeTypeFromByteArray(settings, tmp, enums.DataTypeDateTime)
		if err != nil {
			return err
		}
		reply.Time = ret.(*types.GXDateTime).Value
	}
	if reply.xml != nil {
		reply.xml.AppendStartTag(int(enums.CommandAccessResponse), "", "", false)
		reply.xml.AppendLineFromTag(int(internal.TranslatorTagsLongInvokeId), "Value", reply.xml.IntegerToHex(invokeId, 8, false))
		if !reply.Time.IsZero() {
			reply.xml.AppendComment(fmt.Sprint(reply.Time))
		}
		reply.xml.AppendLineFromTag(int(internal.TranslatorTagsDateTime), "Value", buffer.ToHex(tmp, false))
		reply.Data.Uint8()
		len, err := types.GetObjectCount(reply.Data)
		if err != nil {
			return err
		}
		reply.xml.AppendStartTag(int(internal.TranslatorTagsAccessResponseBody), "", "", false)
		reply.xml.AppendStartTag(int(internal.TranslatorTagsAccessResponseListOfData), "Qty", reply.xml.IntegerToHex(len, 2, false), false)
		for pos := 0; pos != len; pos++ {
			if reply.xml.OutputType() == enums.TranslatorOutputTypeStandardXML {
				reply.xml.AppendStartTag(int(enums.CommandWriteRequest)<<8|int(constants.SingleReadResponseData), "", "", false)
			}
			di := internal.GXDataInfo{}
			di.Xml = reply.xml
			internal.GetData(settings, reply.Data, &di)
			if reply.xml.OutputType() == enums.TranslatorOutputTypeStandardXML {
				reply.xml.AppendEndTag(int(enums.CommandWriteRequest)<<8|int(constants.SingleReadResponseData), false)
			}
		}
		reply.xml.AppendEndTag(int(internal.TranslatorTagsAccessResponseListOfData), false)
		// access-response-specification
		len, err = types.GetObjectCount(reply.Data)
		if err != nil {
			return err
		}
		reply.xml.AppendStartTag(int(internal.TranslatorTagsListOfAccessResponseSpecification), "Qty", reply.xml.IntegerToHex(len, 2, false), false)
		for pos := 0; pos != len; pos++ {
			ch, err := reply.Data.Uint8()
			if err != nil {
				return err
			}
			type_ := enums.AccessServiceCommandType(ch)
			ret, err := reply.Data.Uint8()
			if err != nil {
				return err
			}
			if ret != 0 {
				ret, err = reply.Data.Uint8()
				if err != nil {
					return err
				}
			}
			reply.xml.AppendStartTag(int(internal.TranslatorTagsAccessResponseSpecification), "", "", false)
			reply.xml.AppendStartTag(int(enums.CommandAccessResponse)<<8|int(type_), "", "", false)
			ret2, err := ErrorCodeToString(reply.xml.OutputType(), enums.ErrorCode(int(ret)))
			if err != nil {
				return err
			}
			reply.xml.AppendLineFromTag(int(internal.TranslatorTagsResult), "", ret2)
			reply.xml.AppendEndTag(int(enums.CommandAccessResponse)<<8|int(type_), false)
			reply.xml.AppendEndTag(int(internal.TranslatorTagsAccessResponseSpecification), false)
		}
		reply.xml.AppendEndTag(int(internal.TranslatorTagsListOfAccessResponseSpecification), false)
		reply.xml.AppendEndTag(int(internal.TranslatorTagsAccessResponseBody), false)
		reply.xml.AppendEndTag(int(enums.CommandAccessResponse), false)
	} else {
		_, err = reply.Data.Uint8()
	}
	return err
}

func handleDataNotification(settings *settings.GXDLMSSettings, reply *GXReplyData) error {
	start := reply.Data.Position() - 1
	invokeId, err := reply.Data.Uint32()
	if err != nil {
		return err
	}
	reply.Time = time.Time{}
	len, err := reply.Data.Uint8()
	if err != nil {
		return err
	}
	var tmp []byte
	// If date time is given.
	if len != 0 {
		tmp = make([]byte, len)
		reply.Data.Get(tmp)
		ret, err := internal.ChangeTypeFromByteArray(settings, tmp, enums.DataTypeDateTime)
		if err != nil {
			return err
		}
		reply.Time = ret.(*types.GXDateTime).Value
	}
	if reply.xml != nil {
		reply.xml.AppendStartTag(int(enums.CommandDataNotification), "", "", false)
		if (invokeId & 0x80000000) != 0 {
			reply.xml.AppendComment("High priority.")
		}
		if (invokeId & 0x40000000) != 0 {
			reply.xml.AppendComment("Confirmed service.")
		}
		reply.xml.AppendComment(fmt.Sprintf("Invoke ID: %d", (invokeId & 0x3FFFFFFF)))
		reply.xml.AppendLineFromTag(int(internal.TranslatorTagsLongInvokeId), "Value", reply.xml.IntegerToHex(invokeId, 8, false))
		if !reply.Time.IsZero() {
			reply.xml.AppendComment(fmt.Sprint(reply.Time))
		}
		reply.xml.AppendLineFromTag(int(internal.TranslatorTagsDateTime), "Value", buffer.ToHex(tmp, false))
		reply.xml.AppendStartTag(int(internal.TranslatorTagsNotificationBody), "", "", false)
		reply.xml.AppendStartTag(int(internal.TranslatorTagsDataValue), "", "", false)
		di := internal.GXDataInfo{}
		di.Xml = reply.xml
		internal.GetData(settings, reply.Data, &di)
		reply.xml.AppendEndTag(int(internal.TranslatorTagsDataValue), false)
		reply.xml.AppendEndTag(int(internal.TranslatorTagsNotificationBody), false)
		reply.xml.AppendEndTag(int(enums.CommandDataNotification), false)
	} else {
		getDataFromBlock(reply.Data, start)
		getValueFromData(settings, reply)
	}
	return nil
}

func handleGeneralCiphering(settings *settings.GXDLMSSettings, data *GXReplyData) error {
	if settings.Cipher == nil {
		return errors.New("Secure connection is not supported.")
	}
	/*TODO:
	// If all frames are read.
	if (data.moreData & enums.RequestTypesFrame) == 0 {
		origPos := 0
		if data.xml != nil {
			origPos = data.xml.GetXmlLength()
		}
		data.Data.Position()--
		p := settings.AesGcmParameter{settings: settings, systemTitle: settings.SourceSystemTitle(),
			blockCipherKey: getBlockCipherKey(settings), authenticationKey: getAuthenticationKey(settings)}
		p.xml = data.xml
		tmp, err := GXCiphering.Decrypt(p, data.Data())
		if err != nil {
			if data.xml == nil {
				return err
			}
			data.xml.SetXmlLength(origPos)
			if data.xml != nil && p != nil && p.Xml.g.Comments {
				data.xml.AppendStartTag(enums.CommandGeneralCiphering)
				data.xml.AppendLine(internal.TranslatorTagsTransactionId, nil, IntegerToHex(p.TransactionId(), 16, true))
				data.xml.AppendLine(internal.TranslatorTagsOriginatorSystemTitle, nil, ToHexWithRange(p.SystemTitle(), false))
				data.xml.AppendLine(internal.TranslatorTagsRecipientSystemTitle, nil, ToHexWithRange(p.RecipientSystemTitle(), false))
				data.xml.AppendLine(internal.TranslatorTagsDateTime, nil, ToHexWithRange(p.DateTime(), false))
				data.xml.AppendLine(internal.TranslatorTagsOtherInformation, nil, ToHexWithRange(p.OtherInformation(), false))
				data.xml.AppendStartTag(internal.TranslatorTagsKeyInfo)
				data.xml.AppendStartTag(internal.TranslatorTagsAgreedKey)
				data.xml.AppendLine(internal.TranslatorTagsKeyParameters, nil, IntegerToHex(p.KeyParameters(), 2, true))
				data.xml.AppendLine(internal.TranslatorTagsKeyCipheredData, nil, ToHexWithRange(p.KeyCipheredData(), false))
				data.xml.AppendEndTag(internal.TranslatorTagsAgreedKey)
				data.xml.AppendEndTag(internal.TranslatorTagsKeyInfo)
				data.xml.AppendLine(internal.TranslatorTagsCipheredContent, nil, ToHexWithRange(p.CipheredContent(), false))
				data.xml.AppendEndTag(enums.CommandGeneralCiphering)
			}
		}

		data.Data().Clear()
		data.Data().Set(tmp)
		data.cipheredCommand(enums.CommandGeneralCiphering)
		data.command(enums.CommandNone)
		if p.Security() != uint8(enums.SecurityNone) {
			if data.xml != nil && p != nil && p.Xml.g.Comments {
				data.xml.AppendStartTag(enums.CommandGeneralCiphering)
				data.xml.AppendLine(internal.TranslatorTagsTransactionId, nil, IntegerToHex(p.TransactionId(), 16, true))
				data.xml.AppendLine(internal.TranslatorTagsOriginatorSystemTitle, nil, ToHexWithRange(p.SystemTitle(), false))
				data.xml.AppendLine(internal.TranslatorTagsRecipientSystemTitle, nil, ToHexWithRange(p.RecipientSystemTitle(), false))
				data.xml.AppendLine(internal.TranslatorTagsDateTime, nil, ToHexWithRange(p.DateTime(), false))
				data.xml.AppendLine(internal.TranslatorTagsOtherInformation, nil, ToHexWithRange(p.OtherInformation(), false))
				data.xml.AppendStartTag(internal.TranslatorTagsKeyInfo)
				data.xml.AppendStartTag(internal.TranslatorTagsAgreedKey)
				data.xml.AppendLine(internal.TranslatorTagsKeyParameters, nil, IntegerToHex(p.KeyParameters(), 2, true))
				data.xml.AppendLine(internal.TranslatorTagsKeyCipheredData, nil, ToHexWithRange(p.KeyCipheredData(), false))
				data.xml.AppendEndTag(internal.TranslatorTagsAgreedKey)
				data.xml.AppendEndTag(internal.TranslatorTagsKeyInfo)
				data.xml.StartComment("")
			}
			GetPdu(settings, data)
			if data.xml != nil && p != nil && p.Xml.g.Comments {
				data.xml.EndComment()
				data.xml.AppendLine(internal.TranslatorTagsCipheredContent, nil, ToHexWithRange(p.CipheredContent(), false))
				data.xml.AppendEndTag(enums.CommandGeneralCiphering)
			}
		}
	}
	*/
	return nil
}

// getDataFromFrame returns the get data from HDLC or wrapper frame.
//
// Parameters:
//
//	reply: Received data that includes HDLC frame.
//	info: Reply data.
func getDataFromFrame(reply *types.GXByteBuffer, info *GXReplyData, interfaceType enums.InterfaceType) error {
	offset := info.Data.Size()
	cnt := info.PacketLength - reply.Position()
	if cnt != 0 {
		info.Data.SetCapacity((offset + cnt))
		err := info.Data.SetAt(reply.Array(), reply.Position(), cnt)
		if err != nil {
			return err
		}
		reply.SetPosition(reply.Position() + cnt)
		// Remove CRC and EOP.
		if useHdlc(interfaceType) {
			reply.SetPosition(reply.Position() + 3)
		} else if interfaceType == enums.InterfaceTypeWiredMBus {
			reply.SetPosition(reply.Position() + 2)
		}
	}
	info.Data.SetPosition(offset)
	return nil
}

// getDataFromBlock returns the get data from Block.
//
// Parameters:
//
//	data: Stored data block.
//	index: Position where data starts.
//
// Returns:
//
//	Amount of removed bytes.
func getDataFromBlock(data *types.GXByteBuffer, index int) int {
	if data.Size() == data.Position() {
		data.Clear()
		return 0
	}
	pos := data.Position()
	len_ := pos - index
	data.SetPosition(pos - len_)
	err := data.Move(pos, pos-len_, data.Size()-pos)
	if err != nil {
		panic("getDataFromBlock failed")
	}
	return len_
}

func useHdlc(type_ enums.InterfaceType) bool {
	return type_ == enums.InterfaceTypeHDLC || type_ == enums.InterfaceTypeHdlcWithModeE || type_ == enums.InterfaceTypePlcHdlc
}

// receiverReady returns the generates an acknowledgment message, with which the server is informed to send next packets.
//
// Parameters:
//
//	reply: Reply data.
//
// Returns:
//
//	Acknowledgment message as byte array.
func receiverReady(settings *settings.GXDLMSSettings, reply *GXReplyData) ([]byte, error) {
	if reply.moreData == enums.RequestTypesNone {
		// Generate RR.
		id := settings.KeepAlive()
		if settings.InterfaceType == enums.InterfaceTypePlcHdlc {
			return getMacHdlcFrame(settings, id, 0, nil)
		}
		return getHdlcFrame(settings, id, nil, true)
	}
	// Get next frame.
	if (reply.moreData & enums.RequestTypesFrame) != 0 {
		id := settings.ReceiverReady()
		if settings.InterfaceType == enums.InterfaceTypePlcHdlc {
			return getMacHdlcFrame(settings, id, 0, nil)
		}
		return getHdlcFrame(settings, id, nil, true)
	}
	cmd := settings.Command
	// Get next block.
	var data [][]byte
	var err error
	if reply.moreData == enums.RequestTypesGBT {
		p := NewGXDLMSLNParameters(settings, 0, enums.CommandGeneralBlockTransfer, 0, nil, nil, 0xff, enums.CommandNone)
		p.gbtWindowSize = reply.gbtWindowSize
		p.blockNumberAck = reply.BlockNumber
		p.blockIndex = settings.BlockIndex
		data, err = getLnMessages(p)
		if err != nil {
			return nil, err
		}
	} else {
		// Get next block.
		bb := types.NewGXByteBufferWithCapacity(4)
		if settings.UseLogicalNameReferencing() {
			err = bb.SetUint32(settings.BlockIndex)
			if err != nil {
				return nil, err
			}
		} else {
			err = bb.SetUint16(uint16(settings.BlockIndex))
			if err != nil {
				return nil, err
			}
		}
		settings.IncreaseBlockIndex()
		if settings.UseLogicalNameReferencing() {
			p := NewGXDLMSLNParameters(settings, 0, enums.Command(cmd), byte(constants.GetCommandTypeNextDataBlock), bb, nil, 0xff, reply.cipheredCommand)
			data, err = getLnMessages(p)
		} else {
			p := NewGXDLMSSNParameters(settings, enums.Command(cmd), 1, uint8(constants.VariableAccessSpecificationBlockNumberAccess), bb, nil)
			data, err = getSnMessages(p)
			if err != nil {
				return nil, err
			}
		}
	}
	return data[0], nil
}

// GetDescription returns the get error description.
//
// Parameters:
//
//	error: Error number.
//
// Returns:
//
//	Error as plain text.
func GetDescription(error enums.ErrorCode) string {
	var str string
	switch error {
	case enums.ErrorCodeOk:
		str = ""
	case enums.ErrorCodeRejected:
		str = "HDLC message rejected."
	case enums.ErrorCodeUnacceptableFrame:
		str = "Unacceptable frame."
	case enums.ErrorCodeDisconnectMode:
		str = "Disconnect mode."
	case enums.ErrorCodeHardwareFault:
		str = "Access Error : Device reports a hardware fault."
	case enums.ErrorCodeTemporaryFailure:
		str = "Access Error : Device reports a temporary failure."
	case enums.ErrorCodeReadWriteDenied:
		str = "Access Error : Device reports Read-Write denied."
	case enums.ErrorCodeUndefinedObject:
		str = "Access Error : Device reports a undefined object."
	case enums.ErrorCodeInconsistentClass:
		str = "Access Error : Device reports a inconsistent Class or object."
	case enums.ErrorCodeUnavailableObject:
		str = "Access Error : Device reports a unavailable object."
	case enums.ErrorCodeUnmatchedType:
		str = "Access Error : Device reports a unmatched type."
	case enums.ErrorCodeAccessViolated:
		str = "Access Error : Device reports scope of access violated."
	case enums.ErrorCodeDataBlockUnavailable:
		str = "Access Error : Data Block Unavailable."
	case enums.ErrorCodeLongGetOrReadAborted:
		str = "Access Error : Long Get Or Read Aborted."
	case enums.ErrorCodeNoLongGetOrReadInProgress:
		str = "Access Error : No Long Get Or Read In Progress."
	case enums.ErrorCodeLongSetOrWriteAborted:
		str = "Access Error : Long Set Or Write Aborted."
	case enums.ErrorCodeNoLongSetOrWriteInProgress:
		str = "Access Error : No Long Set Or Write In Progress."
	case enums.ErrorCodeDataBlockNumberInvalid:
		str = "Access Error : Data Block Number Invalid."
	case enums.ErrorCodeOtherReason:
		str = "Access Error : Other Reason."
	default:
		str = "Unknown Error."
	}
	return str
}

// checkInit returns the reserved for internal use.
//
// Parameters:
//
//	settings: DLMS settings.
func checkInit(settings *settings.GXDLMSSettings) error {
	if settings.InterfaceType != enums.InterfaceTypePDU {
		if settings.ClientAddress == 0 {
			return errors.New("Invalid Client Address")
		}
		if settings.ServerAddress == 0 {
			return errors.New("Invalid Server Address")
		}
	}
	return nil
}

func appendData(settings *settings.GXDLMSSettings, obj objects.IGXDLMSBase, index uint8, bb *types.GXByteBuffer, value any) error {
	tp, err := obj.GetDataType(int(index))
	if err != nil {
		return err
	}
	if tp == enums.DataTypeArray {
		if a, ok := value.([]byte); ok {
			return bb.Set(a)
		} else if bb, ok := value.(types.GXByteBuffer); ok {
			return bb.Set(bb.Array())
		}
	} else {
		if tp == enums.DataTypeNone {
			tp, err = internal.GetDLMSDataType(reflect.TypeOf(value))
			if err != nil {
				return err
			}
		} else if _, ok := value.(string); ok {
			ui := obj.Base().GetUIDataType(int(index))
			if _, ok := value.(string); ok {
				value = []byte(value.(string))
			} else if ui == enums.DataTypeOctetString {
				value = buffer.HexToBytes(value.(string))
			}
		}
	}
	internal.SetData(settings, bb, tp, value)
	return err
}

// addLLCBytes returns the add LLC bytes to generated message.
//
// Parameters:
//
//	settings: DLMS settings.
//	data: Data where bytes are added.
func addLLCBytes(settings *settings.GXDLMSSettings, data *types.GXByteBuffer) error {
	var err error
	if settings.IsServer() {
		err = data.InsertBytes(0, internal.LLCReplyBytes)
	} else {
		err = data.InsertBytes(0, internal.LLCSendBytes)
	}
	return err
}

func getCipheringParameters(p *GXDLMSLNParameters) (*settings.AesGcmParameter, error) {
	var err error
	var cmd byte
	var key []byte
	cipher := p.settings.Cipher
	// If client.
	if p.cipheredCommand == enums.CommandNone {
		if ((p.settings.Connected&enums.ConnectionStateDlms) == 0 ||
			(p.settings.NegotiatedConformance&enums.ConformanceGeneralProtection) == 0) &&
			(len(p.settings.PreEstablishedSystemTitle) == 0 || (p.settings.ProposedConformance&enums.ConformanceGeneralProtection) == 0) {
			if cipher.DedicatedKey() != nil && (p.settings.Connected&enums.ConnectionStateDlms) != 0 {
				cmd, err = getDedMessage(p.command)
				if err != nil {
					return nil, err
				}
				key = cipher.DedicatedKey()
			} else {
				cmd, err = getGloMessage(p.command)
				if err != nil {
					return nil, err
				}
				key, err = getBlockCipherKey(p.settings)
			}
		} else {
			if p.settings.Cipher.DedicatedKey() != nil {
				cmd = byte(enums.CommandGeneralDedCiphering)
				key = cipher.DedicatedKey()
			} else {
				cmd = byte(enums.CommandGeneralGloCiphering)
				key, err = getBlockCipherKey(p.settings)
			}
		}
	} else {
		if p.cipheredCommand == enums.CommandGeneralDedCiphering {
			cmd = byte(enums.CommandGeneralDedCiphering)
			key = cipher.DedicatedKey()
		} else if p.cipheredCommand == enums.CommandGeneralGloCiphering {
			cmd = byte(enums.CommandGeneralGloCiphering)
			key, err = getBlockCipherKey(p.settings)
		} else if p.settings.Cipher.DedicatedKey() == nil || isGloMessage(p.cipheredCommand) {
			cmd, err = getGloMessage(p.command)
			key, err = getBlockCipherKey(p.settings)
		} else {
			cmd, err = getDedMessage(p.command)
			key = cipher.DedicatedKey()
		}
	}
	ak, err := getAuthenticationKey(p.settings)
	if err != nil {
		return nil, err
	}
	s := settings.NewAesGcmParameter(cmd, p.settings, cipher.Security(), cipher.SecuritySuite(), uint64(cipher.InvocationCounter()), cipher.SystemTitle(), key, ak)
	s.IgnoreSystemTitle = p.settings.Standard == enums.StandardItaly
	s.SetRecipientSystemTitle(p.settings.SourceSystemTitle())
	return s, nil
}

func Cipher0(p *GXDLMSLNParameters, data []byte) ([]byte, error) {
	par, err := getCipheringParameters(p)
	if err != nil {
		return nil, err
	}
	ret, err := settings.EncryptAesGcm(par, data)
	if err == nil {
		p.settings.Cipher.SetInvocationCounter(p.settings.Cipher.InvocationCounter() + 1)
	}
	return ret, err
}

// getLNPdu returns the get next logical name PDU.
//
// Parameters:
//
//	p: LN parameters.
//	reply: Generated message.
func getLNPdu(p *GXDLMSLNParameters, reply *types.GXByteBuffer) error {
	ciphering := p.command != enums.CommandAarq && p.command != enums.CommandAare &&
		(p.settings.IsCiphered(true) || p.cipheredCommand != enums.CommandNone ||
			(p.settings.Cipher != nil && p.settings.Cipher.Signing() == enums.SigningGeneralSigning))
	len_ := 0
	var err error
	if p.command == enums.CommandAarq {
		if p.settings.Gateway != nil && p.settings.Gateway.PhysicalDeviceAddress != nil {
			reply.SetUint8(enums.CommandGatewayRequest)
			reply.SetUint8(p.settings.Gateway.NetworkID)
			reply.SetUint8(byte(len(p.settings.Gateway.PhysicalDeviceAddress)))
			reply.Set(p.settings.Gateway.PhysicalDeviceAddress)
		}
		reply.SetByteBuffer(p.attributeDescriptor)
	} else {
		// Add enums.Command
		if p.command != enums.CommandGeneralBlockTransfer {
			err = reply.SetUint8(uint8(p.command))
			if err != nil {
				return err
			}
		}
		if p.command == enums.CommandEventNotification || p.command == enums.CommandDataNotification || p.command == enums.CommandAccessRequest || p.command == enums.CommandAccessResponse {
			// Add Long-Invoke-Id-And-Priority
			if p.command != enums.CommandEventNotification {
				if p.invokeId != 0 {
					err = reply.SetUint32(p.invokeId)
					if err != nil {
						return err
					}
				} else {
					err = reply.SetUint32(getLongInvokeIDPriority(p.settings))
					if err != nil {
						return err
					}
				}
			}
			if p.time == nil {
				err := reply.SetUint8(uint8(enums.DataTypeNone))
				if err != nil {
					return err
				}
			} else {
				// // Data is send in octet string. Remove data type except from Event Notification.
				pos := reply.Size()
				err := internal.SetData(p.settings, reply, enums.DataTypeOctetString, p.time)
				if err != nil {
					return err
				}
				if p.command != enums.CommandEventNotification {
					reply.Move(pos+1, pos, reply.Size()-pos-1)
				}
			}
			multipleBlocks(p, reply, ciphering)
		} else if p.command != enums.CommandReleaseRequest && p.command != enums.CommandExceptionResponse {
			// Get request size can be bigger than PDU size.
			if p.command != enums.CommandGetRequest && p.data != nil && p.data.Size() != 0 {
				multipleBlocks(p, reply, ciphering)
			}
			// Change Request type if Set request and multiple blocks is needed.
			if p.command == enums.CommandSetRequest {
				if p.multipleBlocks && (p.settings.NegotiatedConformance&enums.ConformanceGeneralBlockTransfer) == 0 {
					if p.requestType == uint8(constants.SetRequestTypeNormal) {
						p.requestType = uint8(constants.SetRequestTypeFirstDataBlock)
					} else if p.requestType == uint8(constants.SetRequestTypeFirstDataBlock) {
						p.requestType = uint8(constants.SetRequestTypeWithDataBlock)
					}
				}
			} else if p.command == enums.CommandMethodRequest {
				if p.multipleBlocks && (p.settings.NegotiatedConformance&enums.ConformanceGeneralBlockTransfer) == 0 {
					if p.requestType == uint8(constants.ActionRequestTypeNormal) {
						p.attributeDescriptor.SetSize(p.attributeDescriptor.Size() - 1)
						p.requestType = uint8(constants.ActionRequestTypeWithFirstBlock)
					} else if p.requestType == uint8(constants.ActionRequestTypeWithFirstBlock) {
						p.requestType = uint8(constants.ActionRequestTypeWithBlock)
					}
				}
			} else if p.command == enums.CommandMethodResponse {
				if p.multipleBlocks && (p.settings.NegotiatedConformance&enums.ConformanceGeneralBlockTransfer) == 0 {
					p.status = 0xFF
					if p.requestType == uint8(constants.ActionResponseTypeNormal) {
						//Remove Method Invocation Parameters tag.
						p.data.SetPosition(p.data.Position() + 2)
						p.requestType = uint8(constants.ActionResponseTypeWithBlock)
					} else if p.requestType == uint8(constants.ActionResponseTypeWithBlock) && p.data.Available() == 0 {
						p.requestType = uint8(constants.ActionResponseTypeNextBlock)
					}
				}
			} else if p.command == enums.CommandGetResponse {
				if p.multipleBlocks && (p.settings.NegotiatedConformance&enums.ConformanceGeneralBlockTransfer) == 0 {
					if p.requestType == 1 {
						p.requestType = 2
					}
				}
			}
			if p.command != enums.CommandGeneralBlockTransfer {
				err = reply.SetUint8(p.requestType)
				if err != nil {
					return err
				}
				// Add Invoke Id And Priority.
				if p.invokeId != 0 {
					err = reply.SetUint8(uint8(p.invokeId))
					if err != nil {
						return err
					}
				} else {
					err = reply.SetUint8(getInvokeIDPriority(p.settings, p.settings.AutoIncreaseInvokeID))
					if err != nil {
						return err
					}
				}
			}
		}
		err = reply.SetByteBuffer(p.attributeDescriptor)
		if err != nil {
			return err
		}
		// If multiple blocks.
		if p.multipleBlocks && (p.settings.NegotiatedConformance&enums.ConformanceGeneralBlockTransfer) == 0 {
			if p.command != enums.CommandSetResponse && (p.command != enums.CommandMethodResponse || p.data.Size() != 0) {
				// Is last block.
				if p.lastBlock {
					err = reply.SetUint8(1)
					if err != nil {
						return err
					}
					p.settings.Count = 0
					p.settings.Index = 0

				} else {
					err = reply.SetUint8(0)
					if err != nil {
						return err
					}
				}
			}
			err = reply.SetUint32(p.blockIndex)
			if err != nil {
				return err
			}
			p.blockIndex++
			// Add status if reply.
			if p.status != 0xFF {
				if p.status != 0 && p.command == enums.CommandGetResponse {
					err = reply.SetUint8(1)
					if err != nil {
						return err
					}
				}
				err = reply.SetUint8(p.status)
				if err != nil {
					return err
				}
			}
			// Block size.
			if p.data != nil {
				len_ = p.data.Available()
			} else {
				len_ = 0
			}
			totalLength := len_ + reply.Available()
			if ciphering {
				totalLength = totalLength + internal.CipheringHeaderSize
			}
			totalLength = totalLength + getSigningSize(p)
			if totalLength > int(p.settings.MaxPduSize()) {
				len_ = int(p.settings.MaxPduSize()) - reply.Available()
				if len_ < 0 {
					len_ = int(p.settings.MaxPduSize())
				}
				if ciphering {
					len_ = len_ - internal.CipheringHeaderSize
				}
				len_ = len_ - getSigningSize(p)
				len_ = len_ - int(helpers.GetObjectCountSizeInBytes(len_))
			}
			// If server is not asking the next block.
			if !(len_ == 0 && p.command == enums.CommandMethodResponse && p.requestType == uint8(constants.ActionResponseTypeNextBlock)) {
				types.SetObjectCount(len_, reply)
				err = reply.SetByteBufferByCount(p.data, len_)
				if err != nil {
					return err
				}
			}
		}
		// Add data that fits to one block.
		if len_ == 0 {
			// Add status if reply.
			if p.status != 0xFF && p.command != enums.CommandGeneralBlockTransfer {
				if p.status != 0 && p.command == enums.CommandGetResponse {
					err = reply.SetUint8(1)
					if err != nil {
						return err
					}
				}
				err = reply.SetUint8(p.status)
				if err != nil {
					return err
				}
			}
			if p.data != nil && p.data.Size() != 0 {
				len_ = p.data.Size() - p.data.Position()
				// Get request size can be bigger than PDU size.
				if (p.settings.NegotiatedConformance & enums.ConformanceGeneralBlockTransfer) != 0 {
					if 7+len_+reply.Size() > int(p.settings.MaxPduSize()) {
						len_ = int(p.settings.MaxPduSize()) - reply.Size() - 7
					}
					// Cipher data only once.
					if ciphering && p.command != enums.CommandGeneralBlockTransfer {
						err = reply.SetByteBuffer(p.data)
						if err != nil {
							return err
						}
						sign := shoudSign(p)
						var tmp []byte
						if (p.settings.Connected&enums.ConnectionStateDlms) == 0 || !sign {
							tmp, err = Cipher0(p, reply.Array())
						}
						if err != nil {
							return err
						}
						p.data.SetSize(0)
						err = p.data.Set(tmp)
						if err != nil {
							return err
						}
						reply.SetSize(0)
						len_ = p.data.Size()
						if 7+len_ > int(p.settings.MaxPduSize()) {
							len_ = int(p.settings.MaxPduSize()) - 7
						}
						if len_+getSigningSize(p) > int(p.settings.MaxPduSize()) {
							len_ = len_ - getSigningSize(p)
						}
						ciphering = false
					}
				} else if p.command != enums.CommandGetRequest && len_+reply.Size() > int(p.settings.MaxPduSize()) {
					len_ = int(p.settings.MaxPduSize()) - reply.Size()
					len_ = len_ - getSigningSize(p)
				}
				if p.settings.Gateway != nil && p.settings.Gateway.PhysicalDeviceAddress != nil {
					if 3+len_+len(p.settings.Gateway.PhysicalDeviceAddress) > int(p.settings.MaxPduSize()) {
						len_ -= (3 + len(p.settings.Gateway.PhysicalDeviceAddress))
					}
				}
				err = reply.SetByteBufferByCount(p.data, len_)
				if err != nil {
					return err
				}
			}
		}
		// TODO:
		// if (reply.Size != 0 && p.command != enums.CommandGeneralBlockTransfer && p.settings.CryptoNotifier != nil && p.settings.CryptoNotifier.pdu != nil)
		// {
		// p.settings.CryptoNotifier.pdu(p.settings.CryptoNotifier, reply.Array())
		// }
		if ciphering && reply.Size() != 0 && p.command != enums.CommandReleaseRequest && (!p.multipleBlocks || (p.settings.NegotiatedConformance&enums.ConformanceGeneralBlockTransfer) == 0) {
			// GBT ciphering is done for all the data, not just block.
			var tmp []byte
			sign := shoudSign(p)
			if (p.settings.Connected&enums.ConnectionStateDlms) == 0 || !sign {
				tmp, err = Cipher0(p, reply.Array())
			}
			if err != nil {
				return err
			}
			reply.SetSize(0)
			err = reply.Set(tmp)
			if err != nil {
				return err
			}
		}
		if p.command == enums.CommandGeneralBlockTransfer || (p.multipleBlocks && (p.settings.NegotiatedConformance&enums.ConformanceGeneralBlockTransfer) != 0) {
			bb := types.GXByteBuffer{}
			err = bb.SetByteBuffer(reply)
			if err != nil {
				return err
			}
			reply.Clear()
			err = reply.SetUint8(uint8(enums.CommandGeneralBlockTransfer))
			if err != nil {
				return err
			}
			value := uint8(0)
			// Is last block
			if p.lastBlock {
				value = 0x80
			}
			if p.streaming {
				value |= 0x40
			}
			value |= p.gbtWindowSize
			err = reply.SetUint8(value)
			if err != nil {
				return err
			}
			err = reply.SetUint16(uint16(p.blockIndex))
			if err != nil {
				return err
			}
			p.blockIndex++
			// Set block number acknowledged
			if p.command != enums.CommandDataNotification && p.blockNumberAck != 0 {
				err = reply.SetUint16(p.blockNumberAck)
				if err != nil {
					return err
				}
				p.blockNumberAck++
			} else {
				p.blockNumberAck = math.MaxUint16
				err = reply.SetUint16(0)
				if err != nil {
					return err
				}
			}
			types.SetObjectCount(bb.Size(), reply)
			err = reply.SetByteBuffer(&bb)
			if err != nil {
				return err
			}
			p.blockNumberAck++
			if p.command != enums.CommandGeneralBlockTransfer {
				p.command = enums.CommandGeneralBlockTransfer
				p.blockNumberAck = uint16((p.settings.BlockNumberAck + 1))
			}
		}
		if p.settings.Gateway != nil && p.settings.Gateway.PhysicalDeviceAddress != nil && p.command != enums.CommandGatewayRequest {
			tmp := types.NewGXByteBufferFromByteBuffer(reply)
			reply.SetSize(0)
			reply.SetUint8(enums.CommandGatewayRequest)
			reply.SetUint8(byte(p.settings.Gateway.NetworkID))
			reply.SetUint8(byte(len(p.settings.Gateway.PhysicalDeviceAddress)))
			reply.Set(p.settings.Gateway.PhysicalDeviceAddress)
			reply.SetByteBuffer(tmp)
			p.command = enums.CommandGatewayRequest
		}
	}
	if useHdlc(p.settings.InterfaceType) {
		addLLCBytes(p.settings, reply)
	}
	return err
}

// getLnMessages returns the get all Logical name messages. Client uses this to generate messages.
//
// Parameters:
//
//	p: LN settings.
//
// Returns:
//
//	Generated messages.
func getLnMessages(p *GXDLMSLNParameters) ([][]byte, error) {
	reply := types.GXByteBuffer{}
	messages := make([][]byte, 0)
	var err error
	var frame byte = 0
	if p.command == enums.CommandDataNotification || p.command == enums.CommandEventNotification {
		frame = 0x13
	}
	for {
		err = getLNPdu(p, &reply)
		if err != nil {
			return nil, err
		}
		p.lastBlock = true
		if p.attributeDescriptor == nil {
			p.settings.BlockIndex++
		}

		if p.command != enums.CommandAarq && p.command != enums.CommandAare {
			if int(p.settings.MaxPduSize()) < reply.Size() {
				panic("assert failed: MaxPduSize() < reply.Size")
			}
		}
		for reply.Position() != reply.Size() {
			switch p.settings.InterfaceType {
			case enums.InterfaceTypeWRAPPER, enums.InterfaceTypePrimeDcWrapper:
				tmp, err := getWrapperFrame(p.settings, p.command, &reply)
				if err != nil {
					return nil, err
				}
				messages = append(messages, tmp)
			case enums.InterfaceTypeHDLC, enums.InterfaceTypeHdlcWithModeE:
				tmp, err := getHdlcFrame(p.settings, frame, &reply, true)
				if err != nil {
					return nil, err
				}
				messages = append(messages, tmp)
				if reply.Position() != reply.Size() {
					frame = p.settings.NextSend(false)
				}
			case enums.InterfaceTypePDU:
				messages = append(messages, reply.Array())
				reply.SetPosition(reply.Size())
			case enums.InterfaceTypePlc:
				tmp, err := getPlcFrame(p.settings, 0x90, &reply)
				if err != nil {
					return nil, err
				}
				messages = append(messages, tmp)
			case enums.InterfaceTypePlcHdlc:
				tmp, err := getMacHdlcFrame(p.settings, frame, 0, &reply)
				if err != nil {
					return nil, err
				}
				messages = append(messages, tmp)
			case enums.InterfaceTypeSMS:
				tmp, err := getSMSFrame(p.settings, p.command, &reply)
				if err != nil {
					return nil, err
				}
				messages = append(messages, tmp)
			default:
				panic("InterfaceType out of range")
			}
		}
		reply.Clear()
		frame = 0
		if p.data == nil || p.data.Position() == p.data.Size() {
			break
		}
	}
	return messages, err
}

// getSnMessages returns the get all Short Name messages. Client uses this to generate messages.
//
// Parameters:
//
//	p: DLMS SN Parameters.
//
// Returns:
//
//	Generated messages.
func getSnMessages(p *GXDLMSSNParameters) ([][]byte, error) {
	reply := types.GXByteBuffer{}
	messages := make([][]byte, 0)
	var frame byte = 0x00
	if p.Command == enums.CommandInformationReport || p.Command == enums.CommandDataNotification {
		if (p.Settings.Connected & enums.ConnectionStateDlms) != 0 {
			// If connection is established.
			frame = 0x13
		} else {
			frame = 0x03
		}
	}

	for {
		err := getSNPdu(p, &reply)
		if err != nil {
			return nil, err
		}
		if p.Command != enums.CommandAarq && p.Command != enums.CommandAare {
			if int(p.Settings.MaxPduSize()) < reply.Size() {
				panic("assert failed: MaxPduSize() < reply.Size")
			}
		}

		// Command is not add to next PDUs.
		for reply.Position() != reply.Size() {
			if p.Settings.InterfaceType == enums.InterfaceTypeWRAPPER {
				tmp, err := getWrapperFrame(p.Settings, p.Command, &reply)
				if err != nil {
					return nil, err
				}
				messages = append(messages, tmp)
			} else if p.Settings.InterfaceType == enums.InterfaceTypeHDLC ||
				p.Settings.InterfaceType == enums.InterfaceTypeHdlcWithModeE {
				tmp, err := getHdlcFrame(p.Settings, frame, &reply, true)
				if err != nil {
					return nil, err
				}
				messages = append(messages, tmp)
				if reply.Position() != reply.Size() {
					frame = p.Settings.NextSend(false)
				}

			} else if p.Settings.InterfaceType == enums.InterfaceTypePDU {
				messages = append(messages, reply.Array())
				break

			} else if p.Settings.InterfaceType == enums.InterfaceTypePlc {
				val := 0
				if p.Command == enums.CommandAarq {
					val = 0x90
				} else {
					val = int(p.Settings.Plc.InitialCredit) << 5
					val |= int(p.Settings.Plc.CurrentCredit) << 2
					val |= int(p.Settings.Plc.DeltaCredit) & 0x03
				}
				tmp, err := getPlcFrame(p.Settings, byte(val), &reply)
				if err != nil {
					return nil, err
				}
				messages = append(messages, tmp)
				break

			} else {
				panic("InterfaceType out of range")
			}
		}

		reply.Clear()
		frame = 0

		// do..while (p.data != nil && p.data.Position != p.data.Size)
		if p.Data == nil || p.Data.Position() == p.Data.Size() {
			break
		}
	}
	return messages, nil
}

func getSNPdu(p *GXDLMSSNParameters, reply *types.GXByteBuffer) error {
	var err error
	ciphering := p.Command != enums.CommandAarq && p.Command != enums.CommandAare && p.Settings.IsCiphered(false)
	if !ciphering && useHdlc(p.Settings.InterfaceType) {
		if p.Settings.IsServer() {
			err := reply.Set(internal.LLCReplyBytes)
			if err != nil {
				return err
			}
		} else if reply.Size() == 0 {
			err := reply.Set(internal.LLCSendBytes)
			if err != nil {
				return err
			}
		}
	}
	cnt := 0
	cipherSize := 0
	if ciphering {
		cipherSize = internal.CipheringHeaderSize
	}
	if p.Data != nil {
		cnt = p.Data.Size() - p.Data.Position()
	}
	// Add command
	if p.Command == enums.CommandInformationReport {
		err := reply.SetUint8(uint8(p.Command))
		if err != nil {
			return err
		}
		if p.Time == nil {
			reply.SetUint8(byte(enums.DataTypeNone))
		} else {
			// Data is send in octet string. Remove data type.
			pos := reply.Size()
			err = internal.SetData(p.Settings, reply, enums.DataTypeOctetString, p.Time)
			if err != nil {
				return err
			}
			err = reply.Move(pos+1, pos, reply.Size()-pos-1)
			if err != nil {
				return err
			}
		}
		types.SetObjectCount(p.Count, reply)
		err = reply.SetByteBuffer(p.AttributeDescriptor)
		if err != nil {
			return err
		}
	} else if p.Command != enums.CommandAarq && p.Command != enums.CommandAare {
		err = reply.SetUint8(uint8(p.Command))
		if err != nil {
			return err
		}
		if p.Count != 0xFF {
			types.SetObjectCount(p.Count, reply)
		}
		if p.RequestType != 0xFF {
			err = reply.SetUint8(p.RequestType)
			if err != nil {
				return err
			}
		}
		err = reply.SetByteBuffer(p.AttributeDescriptor)
		if err != nil {
			return err
		}
		if !p.MultipleBlocks {
			p.MultipleBlocks = reply.Size()+cipherSize+cnt > int(p.Settings.MaxPduSize())
			// If reply data is not fit to one PDU.
			if p.MultipleBlocks {
				reply.SetSize(0)
				if !ciphering && useHdlc(p.Settings.InterfaceType) {
					if p.Settings.IsServer() {
						err = reply.Set(internal.LLCReplyBytes)
						if err != nil {
							return err
						}
					} else if reply.Size() == 0 {
						err = reply.Set(internal.LLCSendBytes)
						if err != nil {
							return err
						}
					}
				}
				if p.Command == enums.CommandWriteRequest {
					p.RequestType = uint8(constants.VariableAccessSpecificationWriteDataBlockAccess)
				} else if p.Command == enums.CommandReadRequest {
					p.RequestType = uint8(constants.VariableAccessSpecificationReadDataBlockAccess)
				} else if p.Command == enums.CommandReadResponse {
					p.RequestType = uint8(constants.SingleReadResponseDataBlockResult)
				} else {
					return errors.New("Invalid enums.Command")
				}
				err = reply.SetUint8(uint8(p.Command))
				if err != nil {
					return err
				}
				err = reply.SetUint8(1)
				if err != nil {
					return err
				}
				if p.RequestType != 0xFF {
					err = reply.SetUint8(p.RequestType)
					if err != nil {
						return err
					}
				}
				cnt, err = appendMultipleSNBlocks(p, reply)
				if err != nil {
					return err
				}
			}
		} else {
			cnt, err = appendMultipleSNBlocks(p, reply)
			if err != nil {
				return err
			}
		}
	}
	err = reply.SetByteBufferByCount(p.Data, cnt)
	if err != nil {
		return err
	}
	// Af all data is transfered.
	if p.Data != nil && p.Data.Position() == p.Data.Size() {
		p.Settings.Index = 0
		p.Settings.Count = 0
	}
	// If Ciphering is used.
	if ciphering && p.Command != enums.CommandAarq && p.Command != enums.CommandAare {

		cipher := p.Settings.Cipher
		tag, err := getGloMessage(p.Command)
		bk, err := getBlockCipherKey(p.Settings)
		if err != nil {
			return err
		}
		ak, err := getAuthenticationKey(p.Settings)
		if err != nil {
			return err
		}
		s := settings.NewAesGcmParameter(tag, p.Settings, cipher.Security(), cipher.SecuritySuite(), uint64(cipher.InvocationCounter()), cipher.SystemTitle(), bk, ak)
		cipher.SetInvocationCounter(cipher.InvocationCounter() + 1)
		tmp, err := settings.EncryptAesGcm(s, reply.Array())
		if err != nil {
			return err
		}
		reply.SetSize(0)
		if useHdlc(p.Settings.InterfaceType) {
			if p.Settings.IsServer() {
				err = reply.Set(internal.LLCReplyBytes)
				if err != nil {
					return err
				}
			} else if reply.Size() == 0 {
				err = reply.Set(internal.LLCSendBytes)
				if err != nil {
					return err
				}
			}
		}
		err = reply.Set(tmp)
		if err != nil {
			return err
		}
	}
	return err
}

// getHdlcAddress returns the get HDLC address.
//
// Parameters:
//
//	value: HDLC address.
//	size: HDLC address size. This is optional.
//
// Returns:
//
//	HDLC address.
func getHdlcAddress(value int, size uint8) (any, error) {
	if size < 2 && value < 0x80 {
		return uint8((value<<1 | 1)), nil
	}
	if size < 4 && value < 0x4000 {
		return uint16(((value&0x3F80)<<2 | (value&0x7F)<<1 | 1)), nil
	}
	if value < 0x10000000 {
		return uint32(((value&0xFE00000)<<4 | (value&0x1FC000)<<3 | (value&0x3F80)<<2 | (value&0x7F)<<1 | 1)), nil
	}
	return 0, errors.New("Invalid address.")
}

// getHdlcAddressBytes returns the convert HDLC address to bytes.
//
// Parameters:
//
//	size: Address size in bytes.
func getHdlcAddressBytes(value int, size uint8) ([]byte, error) {
	tmp, err := getHdlcAddress(value, size)
	if err != nil {
		return nil, err
	}
	bb := types.GXByteBuffer{}
	if v, ok := tmp.(uint8); ok {
		err = bb.SetUint8(v)
		if err != nil {
			return nil, err
		}
	} else if v, ok := tmp.(uint16); ok {
		err = bb.SetUint16(v)
		if err != nil {
			return nil, err
		}
	} else if v, ok := tmp.(uint32); ok {
		err = bb.SetUint32(v)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("Invalid address type.")
	}
	return bb.Array(), nil
}

// GetWrapperFrame returns the split DLMS PDU to wrapper frames.
//
// Parameters:
//
//	settings: DLMS settings.
//	command: DLMS enums.Command
//	data: Wrapped data.
//
// Returns:
//
//	Wrapper frames
func getWrapperFrame(settings *settings.GXDLMSSettings, command enums.Command, data *types.GXByteBuffer) ([]byte, error) {
	bb := types.GXByteBuffer{}
	// Add version.
	err := bb.SetUint16(1)
	if err != nil {
		return nil, err
	}
	if settings.IsServer() {
		err = bb.SetUint16(uint16(settings.ServerAddress))
		if err != nil {
			return nil, err
		}
		if settings.PushClientAddress != 0 && (command == enums.CommandDataNotification || command == enums.CommandEventNotification) {
			err = bb.SetUint16(uint16(settings.PushClientAddress))
			if err != nil {
				return nil, err
			}
		} else {
			err = bb.SetUint16(uint16(settings.ClientAddress))
			if err != nil {
				return nil, err
			}
		}
	} else {
		err = bb.SetUint16(uint16(settings.ClientAddress))
		if err != nil {
			return nil, err
		}
		err = bb.SetUint16(uint16(settings.ServerAddress))
		if err != nil {
			return nil, err
		}
	}
	if data == nil {
		err = bb.SetUint16(0)
		if err != nil {
			return nil, err
		}
	} else {
		err = bb.SetUint16(uint16(data.Size()))
		if err != nil {
			return nil, err
		}
		err = bb.SetByteBuffer(data)
		if err != nil {
			return nil, err
		}
	}
	// Remove sent data in server side.
	if settings.IsServer() {
		if data.Size() == data.Position() {
			data.Clear()
		} else {
			data.Move(data.Position(), 0, data.Size()-data.Position())
			data.SetPosition(0)
		}
	}
	return bb.Array(), nil
}

// GetSMSFrame returns the split DLMS PDU to SMS frame.
//
// Parameters:
//
//	settings: DLMS settings.
//	command: DLMS enums.Command
//	data: Wrapped data.
//
// Returns:
//
//	SMS frame
func getSMSFrame(settings *settings.GXDLMSSettings, command enums.Command, data *types.GXByteBuffer) ([]byte, error) {
	var err error
	bb := types.GXByteBuffer{}
	if settings.IsServer() {
		err := bb.SetUint8(uint8(settings.ServerAddress))
		if err != nil {
			return nil, err
		}
		if settings.PushClientAddress != 0 && (command == enums.CommandDataNotification || command == enums.CommandEventNotification) {
			err = bb.SetUint8(uint8(settings.PushClientAddress))
			if err != nil {
				return nil, err
			}
		} else {
			err = bb.SetUint8(uint8(settings.ClientAddress))
			if err != nil {
				return nil, err
			}
		}
	} else {
		err = bb.SetUint8(uint8(settings.ClientAddress))
		if err != nil {
			return nil, err
		}
		err = bb.SetUint8(uint8(settings.ServerAddress))
		if err != nil {
			return nil, err
		}
	}
	err = bb.SetByteBuffer(data)
	if err != nil {
		return nil, err
	}
	// Remove sent data in server side.
	if settings.IsServer() {
		data.Clear()
	}
	return bb.Array(), nil
}

// getHdlcFrame returns the get HDLC frame for data.
//
// Parameters:
//
//	settings: DLMS settings.
//	frame: Frame ID. If zero new is generated.
//	data: Data to add.
//
// Returns:
//
//	HDLC frames.
func getHdlcFrame(settings *settings.GXDLMSSettings, frame uint8, data *types.GXByteBuffer, final bool) ([]byte, error) {
	bb := types.GXByteBuffer{}
	var err error
	var frameSize int
	var len_ int
	var primaryAddress []byte
	var secondaryAddress []byte
	if settings.IsServer() {
		if (frame == 0x13 || frame == 0x3) && settings.PushClientAddress != 0 {
			primaryAddress, err = getHdlcAddressBytes(settings.PushClientAddress, 0)
		} else {
			primaryAddress, err = getHdlcAddressBytes(settings.ClientAddress, 0)
		}
		if err != nil {
			return nil, err
		}
		secondaryAddress, err = getHdlcAddressBytes(settings.ServerAddress, settings.ServerAddressSize)
		len_ = len(secondaryAddress)
	} else {
		primaryAddress, err = getHdlcAddressBytes(settings.ServerAddress, settings.ServerAddressSize)
		if err != nil {
			return nil, err
		}
		secondaryAddress, err = getHdlcAddressBytes(settings.ClientAddress, 0)
		if err != nil {
			return nil, err
		}
		len_ = len(primaryAddress)
	}
	err = bb.SetUint8(internal.HDLCFrameStartEnd)
	if err != nil {
		return nil, err
	}
	frameSize = int(settings.Hdlc.MaxInfoTX())
	// Remove BOP, type, len, primaryAddress, secondaryAddress, frame, header CRC, data CRC and EOP from data length.
	if data != nil && data.Size() == 0 {
		frameSize = frameSize - 3
	}
	// If no data
	if data == nil || data.Size() == 0 {
		len_ = 0
		err = bb.SetUint8(0xA0)
		if err != nil {
			return nil, err
		}
	} else if data.Size()-data.Position() <= frameSize {
		len_ = data.Available()
		err = bb.SetUint8(uint8((0xA0 | ((7+len(primaryAddress)+len(secondaryAddress)+len_)>>8)&0x7)))
		if err != nil {
			return nil, err
		}
	} else {
		len_ = frameSize
		err = bb.SetUint8(uint8((0xA8 | ((7+len(primaryAddress)+len(secondaryAddress)+len_)>>8)&0x7)))
		if err != nil {
			return nil, err
		}
	}
	// Frame len.
	if len_ == 0 {
		err = bb.SetUint8(uint8((5 + len(primaryAddress) + len(secondaryAddress) + len_)))
		if err != nil {
			return nil, err
		}
	} else {
		err = bb.SetUint8(uint8((7 + len(primaryAddress) + len(secondaryAddress) + len_)))
		if err != nil {
			return nil, err
		}
	}
	err = bb.Set(primaryAddress)
	if err != nil {
		return nil, err
	}
	err = bb.Set(secondaryAddress)
	if err != nil {
		return nil, err
	}
	// Add frame ID.
	if frame == 0 {
		frame = settings.NextSend(true)
	}
	if !final {
		frame = uint8(int(frame) & ^0x10)
	}
	err = bb.SetUint8(frame)
	if err != nil {
		return nil, err
	}
	// Add header CRC.
	crc := countFCS16(bb.Array(), 1, bb.Size()-1)
	err = bb.SetUint16(crc)
	if err != nil {
		return nil, err
	}
	if len_ != 0 {
		err = bb.SetByteBufferByCount(data, len_)
		if err != nil {
			return nil, err
		}
		crc = countFCS16(bb.Array(), 1, bb.Size()-1)
		err = bb.SetUint16(crc)
		if err != nil {
			return nil, err
		}
	}
	// Add EOP
	err = bb.SetUint8(internal.HDLCFrameStartEnd)
	if err != nil {
		return nil, err
	}
	if data != nil {
		// Remove sent data in server side.
		if settings.IsServer() {
			if data.Size() == data.Position() {
				data.Clear()
			} else {
				data.Move(data.Position(), 0, data.Size()-data.Position())
				data.SetPosition(0)
			}
		}
	}
	return bb.Array(), nil
}

// GetMacFrame returns the get MAC LLC frame for data.
//
// Parameters:
//
//	settings: DLMS settings.
//	frame: HDLC frame sequence number.
//	creditFields: Credit fields.
//	data: Data to add.
//
// Returns:
//
//	MAC frame.
func GetMacFrame(settings *settings.GXDLMSSettings, frame uint8, creditFields uint8, data *types.GXByteBuffer) ([]byte, error) {
	if settings.InterfaceType == enums.InterfaceTypePlc {
		return getPlcFrame(settings, creditFields, data)
	}
	return getMacHdlcFrame(settings, frame, creditFields, data)
}

// getMacHdlcFrame returns the get MAC HDLC frame for data.
//
// Parameters:
//
//	settings: DLMS settings.
//	frame: HDLC frame.
//	creditFields: Credit fields.
//	data: Data to add.
//
// Returns:
//
//	MAC frame.
func getMacHdlcFrame(settings *settings.GXDLMSSettings, frame uint8, creditFields uint8, data *types.GXByteBuffer) ([]byte, error) {
	var val int
	if settings.Hdlc.MaxInfoTX() > 126 {
		settings.Hdlc.SetMaxInfoTX(86)
	}
	bb := types.GXByteBuffer{}
	err := bb.SetUint16(0)
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(creditFields)
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(uint8((settings.Plc.MacSourceAddress >> 4)))
	if err != nil {
		return nil, err
	}
	val = int(settings.Plc.MacSourceAddress << 12)
	val |= int(settings.Plc.MacDestinationAddress & 0xFFF)
	err = bb.SetUint16(uint16(val))
	if err != nil {
		return nil, err
	}
	tmp, err := getHdlcFrame(settings, frame, data, true)
	if err != nil {
		return nil, err
	}
	padLen := (36 - ((10 + len(tmp)) % 36)) % 36
	err = bb.SetUint8(uint8(padLen))
	if err != nil {
		return nil, err
	}
	err = bb.Set(tmp)
	if err != nil {
		return nil, err
	}
	for padLen != 0 {
		err = bb.SetUint8(0)
		if err != nil {
			return nil, err
		}
		padLen--
	}
	// Checksum.
	crc := countFCS24(bb.Array(), 2, bb.Size()-2-padLen)
	err = bb.SetUint8(uint8((crc >> 16)))
	if err != nil {
		return nil, err
	}
	err = bb.SetUint16(uint16(crc))
	if err != nil {
		return nil, err
	}
	val = bb.Size() / 36
	if bb.Size()%36 != 0 {
		val++
	}
	if val == 1 {
		val = int(enums.PlcMacSubframesOne)
	} else if val == 2 {
		val = int(enums.PlcMacSubframesTwo)
	} else if val == 3 {
		val = int(enums.PlcMacSubframesThree)
	} else if val == 4 {
		val = int(enums.PlcMacSubframesFour)
	} else if val == 5 {
		val = int(enums.PlcMacSubframesFive)
	} else if val == 6 {
		val = int(enums.PlcMacSubframesSix)
	} else if val == 7 {
		val = int(enums.PlcMacSubframesSeven)
	} else {
		return nil, errors.New("Data length is too high.")
	}
	err = bb.SetUint16At(0, uint16(val))
	if err != nil {
		return nil, err
	}
	return bb.Array(), nil
}

// getHdlcAddressInfo returns the get HDLC sender and receiver address information.
//
// Parameters:
//
//	reply: Received data.
//	target: target (primary) address
//	source: Source (secondary) address.
//	type: DLMS frame type.
func getHdlcAddressInfo(reply *types.GXByteBuffer, target *int, source *int, type_ *uint8) error {
	position := reply.Position()
	*target = 0
	*source = 0
	*type_ = 0
	packetStartID := reply.Position()
	frameLen := 0
	// If whole frame is not received yet.
	if reply.Size()-reply.Position() < 9 {
		return nil
	}
	for pos := reply.Position(); pos < reply.Size(); pos++ {
		ch, err := reply.Uint8()
		if err != nil {
			return err
		}
		if ch == internal.HDLCFrameStartEnd {
			packetStartID = pos
			break
		}
	}
	// Not a HDLC frame. Sometimes meters can send some strange data between DLMS frames.
	if reply.Position() == reply.Size() {
		// Not enough data to parse
		return nil
	}
	frame, err := reply.Uint8()
	if err != nil {
		return err
	}
	// Check frame length.
	if (frame & 0x7) != 0 {
		frameLen = (int(frame&0x7) << 8)
	}
	ch, err := reply.Uint8()
	frameLen = frameLen + int(ch)
	if reply.Size()-reply.Position()+1 < frameLen {
		reply.SetPosition(packetStartID)
		// Not enough data to parse
		return nil
	}
	eopPos := frameLen + packetStartID + 1
	ch, err = reply.Uint8At(eopPos)
	if err != nil {
		return err
	}
	if ch != internal.HDLCFrameStartEnd {
		return errors.New("Invalid data format.")
	}
	*target, err = GetHDLCAddressFromByteBuffer(reply)
	if err != nil {
		return err
	}
	*source, err = GetHDLCAddressFromByteBuffer(reply)
	if err != nil {
		return err
	}
	*type_, err = reply.Uint8()
	if err != nil {
		return err
	}
	reply.SetPosition(position)
	return nil
}

// isWirelessMBusData returns the check is this wireless M-Bus message.
//
// Parameters:
//
//	buff: Received data.
//
// Returns:
//
//	True, if this is wireless M-Bus message.
func isWirelessMBusData(buff *types.GXByteBuffer) bool {
	if buff.Size()-buff.Position() < 2 {
		return false
	}
	cmd, err := buff.Uint8At(buff.Position() + 1)
	if err != nil {
		return false
	}
	return (cmd&uint8(constants.MBusCommandSndNr)) != 0 || (cmd&uint8(constants.MBusCommandSndUd)) != 0 || (cmd&uint8(constants.MBusCommandRspUd)) != 0
}

// IsWiredMBusData returns the check is this wired M-Bus message.
//
// Parameters:
//
//	buff: Received data.
//
// Returns:
//
//	True, if this is wired M-Bus message.
func IsWiredMBusData(buff *types.GXByteBuffer) bool {
	if buff.Size()-buff.Position() < 1 {
		return false
	}
	cmd, err := buff.Uint8At(buff.Position())
	if err != nil {
		return false
	}
	return cmd == 0x68
}

// getPlcHdlcData returns the check is this PLC S-FSK message.
//
// Parameters:
//
//	buff: Received data.
//
// Returns:
//
//	S-FSK frame size in bytes.
func getPlcSfskFrameSize(buff *types.GXByteBuffer) uint8 {
	var ret byte
	if buff.Size()-buff.Position() < 2 {
		ret = 0
	} else {
		len, err := buff.Uint16At(buff.Position())
		if err != nil {
			return 0
		}
		switch len {
		case uint16(enums.PlcMacSubframesOne):
			ret = 36
		case uint16(enums.PlcMacSubframesTwo):
			ret = 2 * 36
		case uint16(enums.PlcMacSubframesThree):
			ret = 3 * 36
		case uint16(enums.PlcMacSubframesFour):
			ret = 4 * 36
		case uint16(enums.PlcMacSubframesFive):
			ret = 5 * 36
		case uint16(enums.PlcMacSubframesSix):
			ret = 6 * 36
		case uint16(enums.PlcMacSubframesSeven):
			ret = 7 * 36
		default:
			ret = 0
		}
	}
	return ret
}

func checkWrapperAddress(settings *settings.GXDLMSSettings, buff *types.GXByteBuffer, data *GXReplyData, notify *GXReplyData) (bool, error) {
	ret := true
	if settings.IsServer() {
		value, err := buff.Uint16()
		if err != nil {
			return false, err
		}
		data.SourceAddress = int(value)
		// Check that client addresses match.
		if data.xml == nil && settings.ClientAddress != 0 && settings.ClientAddress != int(value) {
			return false, errors.New("Source addresses do not match. It is " + fmt.Sprint(value) + ". It should be " + fmt.Sprint())
		}
		settings.ClientAddress = int(value)
		value, err = buff.Uint16()
		if err != nil {
			return false, err
		}
		data.TargetAddress = int(value)
		// Check that server addresses match.
		if data.xml == nil && settings.ServerAddress != 0 && settings.ServerAddress != int(value) {
			return false, errors.New("Destination addresses do not match. It is " + fmt.Sprint(value) + ". It should be " + fmt.Sprint() + ".")
		}
		settings.ServerAddress = int(value)
	} else {
		value, err := buff.Uint16()
		if err != nil {
			return false, err
		}
		data.TargetAddress = int(value)
		// Check that server addresses match.
		if data.xml == nil && settings.ServerAddress != 0 && settings.ServerAddress != int(value) {
			if notify == nil {
				return false, errors.New("Source addresses do not match. It is " + fmt.Sprint(value) + ". It should be " + fmt.Sprint() + ".")
			}
			notify.SourceAddress = int(value)
			ret = false
		} else {
			settings.ServerAddress = int(value)
		}
		value, err = buff.Uint16()
		if err != nil {
			return false, err
		}
		data.SourceAddress = int(value)
		// Check that client addresses match.
		if data.xml == nil && settings.ClientAddress != 0 && settings.ClientAddress != int(value) {
			if notify == nil {
				return false, errors.New("Destination addresses do not match. It is " + fmt.Sprint(value) + ". It should be " + fmt.Sprint() + ".")
			}
			ret = false
			notify.TargetAddress = int(value)
		} else {
			settings.ClientAddress = int(value)
		}
	}
	return ret, nil
}

func checkSMSAddress(settings *settings.GXDLMSSettings, buff *types.GXByteBuffer, data *GXReplyData, notify *GXReplyData) (bool, error) {
	ret := true
	if settings.IsServer() {
		value, err := buff.Uint8()
		if err != nil {
			return false, err
		}
		data.SourceAddress = int(value)
		// Check that client addresses match.
		if data.xml == nil && settings.ClientAddress != 0 && settings.ClientAddress != int(value) {
			return false, errors.New("Source addresses do not match. It is " + fmt.Sprint(value) + ". It should be " + fmt.Sprint())
		}
		settings.ClientAddress = int(value)
		value, err = buff.Uint8()
		if err != nil {
			return false, err
		}
		data.TargetAddress = int(value)
		// Check that server addresses match.
		if data.xml == nil && settings.ServerAddress != 0 && settings.ServerAddress != int(value) {
			return false, errors.New("Destination addresses do not match. It is " + fmt.Sprint(value) + ". It should be " + fmt.Sprint() + ".")
		}
		settings.ServerAddress = int(value)
	} else {
		value, err := buff.Uint8()
		if err != nil {
			return false, err
		}
		data.TargetAddress = int(value)
		// Check that server addresses match.
		if data.xml == nil && settings.ServerAddress != 0 && settings.ServerAddress != int(value) {
			if notify == nil {
				return false, errors.New("Source addresses do not match. It is " + fmt.Sprint(value) + ". It should be " + fmt.Sprint() + ".")
			}
			notify.SourceAddress = int(value)
			ret = false
		} else {
			settings.ServerAddress = int(value)
		}
		value, err = buff.Uint8()
		if err != nil {
			return false, err
		}
		data.SourceAddress = int(value)
		// Check that client addresses match.
		if data.xml == nil && settings.ClientAddress != 0 && settings.ClientAddress != int(value) {
			if notify == nil {
				return false, errors.New("Destination addresses do not match. It is " + fmt.Sprint(value) + ". It should be " + fmt.Sprint() + ".")
			}
			ret = false
			notify.TargetAddress = int(value)
		} else {
			settings.ClientAddress = int(value)
		}
	}
	return ret, nil
}

func addInvokeId(xml *settings.GXDLMSTranslatorStructure, command enums.Command, type_ int, invokeId uint32) {
	if xml != nil {
		xml.AppendStartTag(int(command), "", "", true)
		xml.AppendStartTag(int(command)<<8|int(type_), "", "", true)
		// InvokeIdAndPriority
		if xml.Comments {
			sb := ""
			if (invokeId & 0x80) != 0 {
				sb += "Priority: High, "
			} else {
				sb += "Priority: Normal, "
			}
			if (invokeId & 0x40) != 0 {
				sb += "ServiceClass: Confirmed, "
			} else {
				sb += "ServiceClass: UnConfirmed, "
			}
			sb += "Invoke ID: " + fmt.Sprint((invokeId & 0xF))
			xml.AppendComment(sb)
		}
		xml.AppendLineFromTag(int(internal.TranslatorTagsInvokeId), "", xml.IntegerToHex(invokeId, 2, false))
	}
}

// handleGbt returns the handle General block transfer message.
//
// Parameters:
//
//	settings: DLMS settings.
func handleGbt(settings *settings.GXDLMSSettings, data *GXReplyData) error {
	index := data.Data.Position() - 1
	data.gbtWindowSize = settings.GbtWindowSize()
	bc, err := data.Data.Uint8()
	if err != nil {
		return err
	}
	data.streaming = (bc & 0x40) != 0
	// GBT Window size.
	windowSize := (uint8)(bc & 0x3F)
	bn, err := data.Data.Uint16()
	if err != nil {
		return err
	}
	bna, err := data.Data.Uint16()
	if data.xml == nil {
		// Remove existing data when first block is received.
		if bn == 1 {
			index = 0
		} else if bna != uint16(settings.BlockIndex-1) {
			data.Data.SetSize(index)
			data.command = enums.CommandNone
			return nil
		}
	}
	data.BlockNumber = bn
	data.BlockNumberAck = bna
	settings.BlockNumberAck = data.BlockNumber
	data.command = enums.CommandNone
	len, err := types.GetObjectCount(data.Data)
	if err != nil {
		return err
	}
	if len > data.Data.Size()-data.Data.Position() {
		data.isComplete = false
		return nil
	}
	if data.xml != nil {
		if (data.Data.Size() - data.Data.Position()) != len {
			data.xml.AppendComment("Data length is " + strconv.Itoa(len) + " and there are " + strconv.Itoa(data.Data.Size()-data.Data.Position()) + " bytes.")
		}
		data.xml.AppendStartTag(int(enums.CommandGeneralBlockTransfer), "", "", true)
		if data.xml.Comments {
			data.xml.AppendComment("Last block: " + strconv.FormatBool((bc&0x80) != 0))
			data.xml.AppendComment("Streaming: " + strconv.FormatBool(data.streaming))
			data.xml.AppendComment("Window size: " + strconv.Itoa(int(windowSize)))
		}
		data.xml.AppendLine(internal.TranslatorTagsBlockControl.String(), "", data.xml.IntegerToHex(bc, 2, false))
		data.xml.AppendLine(internal.TranslatorTagsBlockNumber.String(), "", data.xml.IntegerToHex(data.BlockNumber, 4, false))
		data.xml.AppendLine(internal.TranslatorTagsBlockNumberAck.String(), "", data.xml.IntegerToHex(data.BlockNumberAck, 4, false))
		// If last block and not streaming and comments.
		if (bc&0x80) != 0 && !data.streaming && data.xml.Comments && data.Data.Available() != 0 {
			pos := data.Data.Position()
			len2 := data.xml.GetXmlLength()
			reply := GXReplyData{}
			reply.Data = data.Data
			reply.xml = data.xml
			reply.xml.StartComment("")
			err = GetPdu(settings, &reply)
			reply.xml.EndComment()
			if err != nil {
				data.xml.SetXmlLength(len2)
				//It's ok if this fails.
				return err
			}
			data.Data.SetPosition(pos)
		}
		data.xml.AppendLine(internal.TranslatorTagsBlockData.String(), "", data.Data.RemainingHexString(true))
		data.xml.AppendEndTag(int(enums.CommandGeneralBlockTransfer), true)
		return nil
	}
	getDataFromBlock(data.Data, index)
	// Is Last block,
	if (bc & 0x80) == 0 {
		data.moreData = enums.RequestTypesGBT
	} else {
		data.moreData &= ^enums.RequestTypesGBT
		if data.Data.Size() != 0 {
			data.Data.SetPosition(0)
			GetPdu(settings, data)
		}
		// Get data if all data is read or we want to peek data.
		if data.Data.Position() != data.Data.Size() && (data.command == enums.CommandReadResponse || data.command == enums.CommandGetResponse) &&
			(data.moreData == enums.RequestTypesNone || data.Peek) {
			data.Data.SetPosition(0)
			getValueFromData(settings, data)
		}
	}
	return nil
}

func HandleConfirmedServiceError(data *GXReplyData) error {
	if data.xml != nil {
		data.xml.AppendStartTag(int(enums.CommandConfirmedServiceError), "", "", true)
		if data.xml.OutputType() == enums.TranslatorOutputTypeStandardXML {
			data.Data.Uint8()
			data.xml.AppendStartTag(int(internal.TranslatorTagsInitiateError), "", "", true)
			ret, err := data.Data.Uint8()
			if err != nil {
				return err
			}
			type_ := enums.ServiceError(ret)
			tag := standardServiceErrorToString(type_)
			ret, err = data.Data.Uint8()
			if err != nil {
				return err
			}
			value := standardGetServiceErrorValue(type_, ret)
			data.xml.AppendLine("x:"+tag, "", value)
			data.xml.AppendEndTag(int(internal.TranslatorTagsInitiateError), false)
		} else {
			ret, err := data.Data.Uint8()
			if err != nil {
				return err
			}
			data.xml.AppendLineFromTag(int(internal.TranslatorTagsService), "Value", data.xml.IntegerToHex(ret, 2, false))
			ret, err = data.Data.Uint8()
			if err != nil {
				return err
			}
			type_ := enums.ServiceError(ret)
			data.xml.AppendStartTag(int(internal.TranslatorTagsServiceError), "", "", true)
			ret, err = data.Data.Uint8()
			if err != nil {
				return err
			}
			data.xml.AppendLine(simpleServiceErrorToString(type_), "Value", simpleGetServiceErrorValue(type_, ret))
			data.xml.AppendEndTag(int(internal.TranslatorTagsServiceError), false)
		}
		data.xml.AppendEndTag(int(enums.CommandConfirmedServiceError), false)
	} else {
		ret, err := data.Data.Uint8()
		if err != nil {
			return err
		}
		service := enums.ConfirmedServiceError(ret)
		ret, err = data.Data.Uint8()
		if err != nil {
			return err
		}
		type_ := enums.ServiceError(ret)
		ret, err = data.Data.Uint8()
		if err != nil {
			return err
		}
		return dlmserrors.NewGXDLMSConfirmedServiceError(service, type_, ret)
	}
	return nil
}

func HandleExceptionResponse(data *GXReplyData) error {
	ret, err := data.Data.Uint8()
	if err != nil {
		return err
	}
	state := enums.ExceptionStateError(ret)
	ret, err = data.Data.Uint8()
	if err != nil {
		return err
	}
	error_ := enums.ExceptionServiceError(ret)
	var value uint32
	if error_ == enums.ExceptionServiceErrorInvocationCounterError && data.Data.Available() > 3 {
		value, err = data.Data.Uint32()
		if err != nil {
			return err
		}
	}
	if data.xml != nil {
		data.xml.AppendStartTag(int(enums.CommandExceptionResponse), "", "", true)
		if data.xml.OutputType() == enums.TranslatorOutputTypeStandardXML {
			ret, err := standardStateErrorToString(state)
			if err != nil {
				return err
			}
			data.xml.AppendLine(internal.TranslatorTagsStateError.String(), "", ret)
			ret, err = standardExceptionServiceErrorToString(error_)
			if err != nil {
				return err
			}
			data.xml.AppendLine(internal.TranslatorTagsServiceError.String(), "", ret)
		} else {
			ret, err := simpleStateErrorToString(state)
			if err != nil {
				return err
			}
			data.xml.AppendLine(internal.TranslatorTagsStateError.String(), "", ret)
			ret, err = simpleExceptionServiceErrorToString(error_)
			if err != nil {
				return err
			}
			data.xml.AppendLine(internal.TranslatorTagsServiceError.String(), "", ret)
		}
		data.xml.AppendEndTag(int(enums.CommandExceptionResponse), false)
	} else {
		return NewGXDLMSExceptionResponse(state, error_, value)
	}
	return nil
}

// getValueFromData returns the get value from the data.
//
// Parameters:
//
//	settings: DLMS settings.
//	reply: Received data.
func getValueFromData(settings *settings.GXDLMSSettings, reply *GXReplyData) error {
	data := reply.Data
	info := internal.GXDataInfo{}
	if _, ok := reply.Value.([]any); ok {
		info.Type = enums.DataTypeArray
		info.Count = reply.TotalCount
		info.Index = reply.Count()
	}
	index := data.Position()
	data.SetPosition(reply.ReadPosition)
	value, err := internal.GetData(settings, data, &info)
	if err != nil {
		return err
	}
	if value != nil {
		// If new data.
		if a, ok := value.(types.GXArray); ok {
			if reply.Value == nil {
				reply.Value = value
			} else if len(a) != 0 {
				// Add items to collection.
				reply.Value = append(reply.Value.(types.GXArray), a...)
			}
		} else {
			reply.DataType = info.Type
			reply.Value = value
			reply.TotalCount = 0
			reply.ReadPosition = data.Position()
			reply.ReadPosition = data.Position()
			// Element count.
			reply.TotalCount = info.Count
		}
	} else if info.Complete && reply.command == enums.CommandDataNotification {
		reply.ReadPosition = data.Position()
	}
	data.SetPosition(index)
	// If last data frame of the data block is read.
	if reply.command != enums.CommandDataNotification && info.Complete && reply.moreData == enums.RequestTypesNone {
		settings.ResetBlockIndex()
		data.SetPosition(0)
	}
	return nil
}

// GetActionInfo returns the action method information.
//
// Parameters:
//
//	objectType: object type.
//	value: Starting address of action methods.
//	count: Count of action methods
func getActionInfo(objectType enums.ObjectType, value *int, count *int) {
	*count = 0
	*value = 0
	switch objectType {
	case enums.ObjectTypeImageTransfer:
		*value = 0x40
		*count = 4
	case enums.ObjectTypeActivityCalendar:
		*value = 0x50
		*count = 1
	case enums.ObjectTypeAssociationLogicalName:
		*value = 0x60
		*count = 4
	case enums.ObjectTypeAssociationShortName:
		*value = 0x20
		*count = 8
	case enums.ObjectTypeClock:
		*value = 0x60
		*count = 6
	case enums.ObjectTypeDemandRegister:
		*value = 0x48
		*count = 2
	case enums.ObjectTypeExtendedRegister:
		*value = 0x38
		*count = 1
	case enums.ObjectTypeIP4Setup:
		*value = 0x60
		*count = 3
	case enums.ObjectTypeMBusSlavePortSetup:
		*value = 0x60
		*count = 8
	case enums.ObjectTypeProfileGeneric:
		*value = 0x58
		*count = 4
	case enums.ObjectTypeRegister:
		*value = 0x28
		*count = 1
	case enums.ObjectTypeRegisterActivation:
		*value = 0x30
		*count = 3
	case enums.ObjectTypeRegisterTable:
		*value = 0x28
		*count = 2
	case enums.ObjectTypeScriptTable:
		*value = 0x20
		*count = 1
	case enums.ObjectTypeSpecialDaysTable:
		*value = 0x10
		*count = 2
	case enums.ObjectTypeDisconnectControl:
		*value = 0x20
		*count = 2
	case enums.ObjectTypeSecuritySetup:
		*value = 0x30
		*count = 8
	case enums.ObjectTypePushSetup:
		*value = 0x38
		*count = 1
	default:
		*value = 0
		*count = 0
	}
}

func getAttributeSize(obj objects.IGXDLMSBase, attributeIndex int) (int, error) {
	rowsize := 0
	if attributeIndex == 0 {
		for pos := 1; pos < obj.GetAttributeCount(); pos++ {
			size, err := getAttributeSize(obj, pos)
			if err != nil {
				return 0, err
			}
			rowsize += size
		}
	} else {
		dt, err := obj.GetDataType(attributeIndex)
		if err != nil {
			return 0, err
		}
		if dt == enums.DataTypeOctetString {
			udt := obj.GetUIDataType(attributeIndex)
			if err != nil {
				return 0, err
			}
			if udt == enums.DataTypeDateTime || udt == enums.DataTypeDate || udt == enums.DataTypeTime {
				rowsize = internal.GetDataTypeSize(udt)
			}
			if udt == enums.DataTypeDateTime || udt == enums.DataTypeDate || udt == enums.DataTypeTime {
				rowsize = internal.GetDataTypeSize(udt)
			}
		} else if dt == enums.DataTypeNone {
			rowsize = 2
		} else {
			rowsize = internal.GetDataTypeSize(dt)
		}
	}
	return rowsize, nil
}

func RowsToPdu(settings *settings.GXDLMSSettings, pg *objects.GXDLMSProfileGeneric) (uint16, error) {
	// Count how many rows we can fit to one PDU.
	rowsize := 0
	for _, it := range pg.CaptureObjects {
		ret, err := getAttributeSize(it.Key, it.Value.AttributeIndex)
		if err != nil {
			return 0, err
		}
		rowsize += ret
	}
	if rowsize != 0 {
		return uint16((int(settings.MaxPduSize()) / rowsize)), nil
	}
	return 0, nil
}

// ParseSnrmUaResponse returns the parses SNRM or UA Response from byte array and update settings.
//
// Parameters:
//
//	data: Received data
func parseSnrmUaResponse(data *types.GXByteBuffer, settings *settings.GXDLMSSettings) error {
	// If default settings are used.
	if data.Size() == 0 {
		settings.Hdlc.SetMaxInfoRX(defaultMaxInfoRX)
		settings.Hdlc.SetMaxInfoTX(defaultMaxInfoTX)
		settings.Hdlc.SetWindowSizeRX(defaultWindowSizeRX)
		settings.Hdlc.SetWindowSizeTX(defaultWindowSizeTX)
		return nil
	}
	data.Uint8() // Skip FromatID
	data.Uint8() // Skip Group ID.
	data.Uint8() // Skip Group length.
	var val any
	for data.Position() < data.Size() {
		ret, err := data.Uint8()
		if err != nil {
			return err
		}
		id := internal.HDLCInfo(ret)
		len, err := data.Uint8()
		if err != nil {
			return err
		}
		switch len {
		case 1:
			val, err = data.Uint8()
			if err != nil {
				return err
			}
		case 2:
			val, err = data.Uint16()
			if err != nil {
				return err
			}
		case 4:
			val, err = data.Uint32()
			if err != nil {
				return err
			}
		default:
			return errors.New("Invalid Exception.")
		}
		switch id {
		case internal.HDLCInfoMaxInfoTX:
			if v, ok := val.(byte); ok {
				settings.Hdlc.SetMaxInfoRX(uint16(v))
			} else {
				settings.Hdlc.SetMaxInfoRX(val.(uint16))
			}
		case internal.HDLCInfoMaxInfoRX:
			if v, ok := val.(byte); ok {
				settings.Hdlc.SetMaxInfoTX(uint16(v))
			} else {
				settings.Hdlc.SetMaxInfoTX(val.(uint16))
			}
		case internal.HDLCInfoWindowSizeTX:
			if v, ok := val.(uint32); ok {
				settings.Hdlc.SetWindowSizeRX(byte(v))
			} else {
				settings.Hdlc.SetWindowSizeRX(val.(byte))
			}
		case internal.HDLCInfoWindowSizeRX:
			if v, ok := val.(uint32); ok {
				settings.Hdlc.SetWindowSizeTX(byte(v))
			} else {
				settings.Hdlc.SetWindowSizeTX(val.(byte))
			}
		default:
			return errors.New("Invalid UA response.")
		}
	}
	return nil
}

// appendHdlcParameter returns the add HDLC parameter.
func appendHdlcParameter(data *types.GXByteBuffer, value uint16) error {
	var err error
	if value < 0x100 {
		err = data.SetUint8(1)
		if err != nil {
			return err
		}
		err = data.SetUint8(uint8(value))
	} else {
		err = data.SetUint8(2)
		if err != nil {
			return err
		}
		err = data.SetUint16(value)
	}
	return err
}

func IsCiphered(cmd uint8) bool {
	switch enums.Command(cmd) {
	case enums.CommandGloReadRequest:
	case enums.CommandGloWriteRequest:
	case enums.CommandGloGetRequest:
	case enums.CommandGloSetRequest:
	case enums.CommandGloReadResponse:
	case enums.CommandGloWriteResponse:
	case enums.CommandGloGetResponse:
	case enums.CommandGloSetResponse:
	case enums.CommandGloMethodRequest:
	case enums.CommandGloMethodResponse:
	case enums.CommandDedGetRequest:
	case enums.CommandDedSetRequest:
	case enums.CommandDedReadResponse:
	case enums.CommandDedGetResponse:
	case enums.CommandDedSetResponse:
	case enums.CommandDedMethodRequest:
	case enums.CommandDedMethodResponse:
	case enums.CommandGeneralGloCiphering:
	case enums.CommandGeneralDedCiphering:
	case enums.CommandAare:
	case enums.CommandAarq:
	case enums.CommandGloConfirmedServiceError:
	case enums.CommandDedConfirmedServiceError:
	case enums.CommandGeneralCiphering:
	case enums.CommandReleaseRequest:
	case enums.CommandGeneralSigning:
		return true
	}
	return false
}

// GetPdu returns the get PDU from the packet.
//
// Parameters:
//
//	settings: DLMS settings.
//	data: received data.
func GetPdu(conf *settings.GXDLMSSettings, data *GXReplyData) error {
	var err error
	cmd := data.command
	// If header is not read yet or GBT message.
	if data.command == enums.CommandNone {
		// If PDU is missing.
		if data.Data.Size()-data.Data.Position() == 0 {
			return errors.New("Invalid PDU.")
		}
		if conf.InterfaceType == enums.InterfaceTypePrimeDcWrapper {
			ret, err := primeDcHandleNotification(data.Data, data, nil)
			if err != nil {
				return err
			}
			if ret {
				return nil
			}
		}
		index := data.Data.Position()
		ch, err := data.Data.Uint8()
		if err != nil {
			return err
		}
		cmd = enums.Command(ch)
		data.command = cmd
		if conf.Closing {
			conf.Closing = false
			if data.xml == nil && cmd != enums.CommandReleaseResponse && cmd != enums.CommandDisconnectMode && cmd != enums.CommandConfirmedServiceError && cmd != enums.CommandExceptionResponse && cmd != enums.CommandAare {
				return errors.New("Invalid reply. The client has closed the connection.")
			}
		}
		switch cmd {
		case enums.CommandReadResponse:
			if !conf.UseLogicalNameReferencing() {
				ret, err := handleReadResponse(conf, data, index)
				if err != nil {
					return err
				}
				if !ret {
					return nil
				}
			} else {
				data.command = enums.CommandNone
				return errors.New("Invalid command.")
			}
		case enums.CommandGetResponse:
			ret, err := handleGetResponse(conf, data, index)
			if err != nil {
				return err
			}
			if !ret {
				return nil
			}
		case enums.CommandSetResponse:
			err = handleSetResponse(conf, data)
		case enums.CommandWriteResponse:
			err = handleWriteResponse(data)
		case enums.CommandMethodResponse:
			err = handleMethodResponse(conf, data, index)
		case enums.CommandAccessRequest:
			if data.xml != nil || (!conf.IsServer() && (data.moreData&enums.RequestTypesFrame) == 0) {
				err = handleAccessRequest(conf, nil, data.Data, nil, data.xml, enums.CommandNone)
			}
		case enums.CommandAccessResponse:
			if data.xml != nil || (!conf.IsServer() && (data.moreData&enums.RequestTypesFrame) == 0) {
				err = handleAccessResponse(conf, data)
			}
		case enums.CommandGeneralBlockTransfer:
			if data.xml != nil || (!conf.IsServer() && (data.moreData&enums.RequestTypesFrame) == 0) {
				err = handleGbt(conf, data)
			}
		case enums.CommandAarq, enums.CommandAare:
			data.Data.SetPosition(data.Data.Position() - 1)
		case enums.CommandReleaseResponse:
			break
		case enums.CommandConfirmedServiceError:
			err = HandleConfirmedServiceError(data)
		case enums.CommandExceptionResponse:
			err = HandleExceptionResponse(data)
		case enums.CommandGetRequest:
			if data.xml != nil || (!conf.IsServer() && (data.moreData&enums.RequestTypesFrame) == 0) {
				err = handleGetRequest(conf, nil, data.Data, nil, data.xml, enums.CommandNone)
			}
		case enums.CommandReadRequest:
			if data.xml != nil || (!conf.IsServer() && (data.moreData&enums.RequestTypesFrame) == 0) {
				err = handleReadRequest(conf, nil, data.Data, nil, data.xml, enums.CommandNone)
			}
		case enums.CommandWriteRequest:
			if data.xml != nil || (!conf.IsServer() && (data.moreData&enums.RequestTypesFrame) == 0) {
				err = handleWriteRequest(conf, nil, data.Data, nil, data.xml, enums.CommandNone)
			}
		case enums.CommandSetRequest:
			if data.xml != nil || (!conf.IsServer() && (data.moreData&enums.RequestTypesFrame) == 0) {
				err = handleSetRequest(conf, nil, data.Data, nil, data.xml, enums.CommandNone)
			}
		case enums.CommandMethodRequest:
			if data.xml != nil || (!conf.IsServer() && (data.moreData&enums.RequestTypesFrame) == 0) {
				err = handleMethodRequest(conf, nil, data.Data, nil, nil, data.xml, enums.CommandNone)
			}
		case enums.CommandReleaseRequest:
			// Server handles this.
			if (data.moreData & enums.RequestTypesFrame) != 0 {
				break
			}
		case enums.CommandGloReadRequest,
			enums.CommandGloWriteRequest,
			enums.CommandGloGetRequest,
			enums.CommandGloSetRequest,
			enums.CommandGloMethodRequest,
			enums.CommandDedGetRequest,
			enums.CommandDedSetRequest,
			enums.CommandDedMethodRequest:
			err = handleGloDedRequest(conf, data)
		case enums.CommandGloReadResponse, enums.CommandGloWriteResponse,
			enums.CommandGloGetResponse,
			enums.CommandGloSetResponse,
			enums.CommandGloMethodResponse,
			enums.CommandGloEventNotification,
			enums.CommandDedGetResponse,
			enums.CommandDedSetResponse,
			enums.CommandDedMethodResponse,
			enums.CommandDedEventNotification,
			enums.CommandGloConfirmedServiceError,
			enums.CommandDedConfirmedServiceError:
			err = handleGloDedResponse(conf, data, index)
		case enums.CommandGeneralGloCiphering,
			enums.CommandGeneralDedCiphering:
			if conf.IsServer() {
				err = handleGloDedRequest(conf, data)
			} else {
				err = handleGloDedResponse(conf, data, index)
			}
		case enums.CommandGeneralSigning:
			if conf.IsServer() {
				err = handleGloDedRequest(conf, data)
			} else {
				err = handleGloDedResponse(conf, data, index)
			}
		case enums.CommandDataNotification:
			handleDataNotification(conf, data)
		case enums.CommandEventNotification:
			break
		case enums.CommandInformationReport:
			break
		case enums.CommandGeneralCiphering:
			handleGeneralCiphering(conf, data)
		case enums.CommandGatewayRequest,
			enums.CommandGatewayResponse:
			data.Gateway = &settings.GXDLMSGateway{}
			data.Gateway.NetworkID, err = data.Data.Uint8()
			if err != nil {
				return err
			}
			len, err := types.GetObjectCount(data.Data)
			if err != nil {
				return err
			}
			data.Gateway.PhysicalDeviceAddress = make([]byte, len)
			data.Data.Get(data.Gateway.PhysicalDeviceAddress)
			getDataFromBlock(data.Data, index)
			data.command = enums.CommandNone
			GetPdu(conf, data)
		case enums.CommandPingResponse,
			enums.CommandDiscoverReport,
			enums.CommandDiscoverRequest,
			enums.CommandRegisterRequest:
			break
		default:
			/*TODO:
			if settings.customPdu != nil {
				data.Data.SetPosition(data.Data.Position() - 1)
				// GXCustomPduArgs e = new GXCustomPduArgs()
				// {
				// Data = data.Data.Remaining()
				// }
				// Tyypin Gurux.LanguageConverter.UnknownDataTypeException poikkeus.
				// settings.customPdu(e)
				// data.Value = e.Data
				return nil
			} else
			*/
			{
				data.command = enums.CommandNone
				return errors.New("Invalid enums Command")
			}
		}
	} else if (data.moreData & enums.RequestTypesFrame) == 0 {
		// Is whole block is read and if last packet and data is not try to peek.
		if !data.Peek && data.moreData == enums.RequestTypesNone {
			data.Data.SetPosition(0)
			if data.command == enums.CommandAare || data.command == enums.CommandAarq {
				data.Data.SetPosition(0)
			} else {
				data.Data.SetPosition(1)
			}
		}
		if cmd == enums.CommandGeneralBlockTransfer {
			if data.xml != nil || !conf.IsServer() {
				data.Data.SetPosition(data.CipherIndex + 1)
				handleGbt(conf, data)
				data.CipherIndex = data.Data.Size()
				data.command = enums.CommandNone
			}
		} else if conf.IsServer() {
			switch cmd {
			case enums.CommandGloReadRequest,
				enums.CommandGloWriteRequest,
				enums.CommandGloGetRequest,
				enums.CommandGloSetRequest,
				enums.CommandGloMethodRequest,
				enums.CommandGeneralSigning:
				data.command = enums.CommandNone
				data.Data.SetPosition(data.CipherIndex)
				GetPdu(conf, data)
			default:
				break
			}
		} else {
			// Client do not need a command any more.
			if data.IsMoreData() {
				data.command = enums.CommandNone
			}
			switch cmd {
			case enums.CommandGloReadResponse,
				enums.CommandGloWriteResponse,
				enums.CommandGloGetResponse,
				enums.CommandGloSetResponse,
				enums.CommandGloMethodResponse,
				enums.CommandDedReadResponse,
				enums.CommandDedWriteResponse,
				enums.CommandDedGetResponse,
				enums.CommandDedSetResponse,
				enums.CommandDedMethodResponse,
				enums.CommandGeneralGloCiphering,
				enums.CommandGeneralDedCiphering,
				enums.CommandGeneralCiphering,
				enums.CommandAccessResponse,
				enums.CommandGeneralSigning:
				data.command = enums.CommandNone
				data.Data.SetPosition(data.CipherIndex)
				GetPdu(conf, data)
			default:
				break
			}
			if cmd == enums.CommandReadResponse && data.TotalCount > 1 {
				_, err := handleReadResponse(conf, data, 0)
				return err
			}
		}
	}
	// Get data only blocks if SN is used. This is faster.
	if cmd == enums.CommandReadResponse && data.CommandType() == byte(constants.SingleReadResponseDataBlockResult) && (data.moreData&enums.RequestTypesFrame) != 0 {
		return nil
	}
	// Get data if all data is read or we want to peek data.
	if data.Error == 0 && data.xml == nil && data.Data.Position() != data.Data.Size() && (cmd == enums.CommandReadResponse || cmd == enums.CommandGetResponse || cmd == enums.CommandMethodResponse) && (data.moreData == enums.RequestTypesNone || data.Peek) {
		err = getValueFromData(conf, data)
	}
	return err
}

func getData(settings *settings.GXDLMSSettings, reply *types.GXByteBuffer, data *GXReplyData, notify *GXReplyData) (bool, error) {
	var err error
	frame := uint8(0)
	isLast := true
	isNotify := false
	switch settings.InterfaceType {
	case enums.InterfaceTypeHDLC, enums.InterfaceTypeHdlcWithModeE:
		frame, err = getHdlcData(settings.IsServer(), settings, reply, data, notify)
		if err != nil {
			return false, err
		}
		isLast = (frame & 0x10) != 0
		if notify != nil && (frame == 0x13 || frame == 0x3) {
			data = notify
			isNotify = true
		}
		data.SetFrameId(frame)
	case enums.InterfaceTypeWRAPPER, enums.InterfaceTypePrimeDcWrapper:
		if !getTcpData(settings, reply, data, notify) {
			if notify != nil {
				data = notify
			}
			isNotify = true
		}
	case enums.InterfaceTypeWirelessMBus:
		getWirelessMBusData(settings, reply, data)
	case enums.InterfaceTypePDU:
		data.PacketLength = reply.Size()
		data.SetIsComplete(reply.Size() != 0)
	case enums.InterfaceTypePlc:
		getPlcData(settings, reply, data)
	case enums.InterfaceTypePlcHdlc:
		_, err := getPlcHdlcData(settings, reply, data)
		if err != nil {
			return false, err
		}
	case enums.InterfaceTypeWiredMBus:
		getWiredMBusData(settings, reply, data)
	case enums.InterfaceTypeSMS:
		if !getSmsData(settings, reply, data, notify) {
			if notify != nil {
				data = notify
			}
			isNotify = true
		}
	default:
		return false, errors.New("Invalid Interface type.")
	}
	// If all data is not read yet.
	if !data.IsComplete() {
		return false, nil
	}
	if settings.InterfaceType != enums.InterfaceTypePlcHdlc {
		getDataFromFrame(reply, data, settings.InterfaceType)
	}
	// If keepalive or get next frame request.
	if data.xml != nil || (((frame != 0x13 && frame != 0x3) || data.IsMoreData()) && (frame&0x1) != 0) {
		if (settings.InterfaceType == enums.InterfaceTypeHDLC || settings.InterfaceType == enums.InterfaceTypeHdlcWithModeE) && (data.Error == int(enums.ErrorCodeRejected) || data.Data.Size() != 0) {
		}
		if frame == 0x3 && data.IsMoreData() {
			tmp, err := getData(settings, reply, data, notify)
			if err != nil {
				return false, err
			}
			data.Data.SetPosition(0)
			return tmp, nil
		}
		return true, nil
	}
	if data.RawPdu {
		data.Data.SetPosition(0)
		return !isNotify, nil
	}
	if frame == 0x13 && !data.IsMoreData() {
		data.Data.SetPosition(0)
	}
	if settings.InterfaceType == enums.InterfaceTypeCoAP {
		if (data.moreData & enums.RequestTypesFrame) == 0 {
			data.Data.SetPosition(0)
			err = GetPdu(settings, data)
		}
	} else {
		err = GetPdu(settings, data)
	}
	if err != nil {
		return false, err
	}
	if notify != nil && !isNotify {
		switch data.command {
		case enums.CommandDataNotification,
			enums.CommandGloEventNotification,
			enums.CommandInformationReport,
			enums.CommandEventNotification,
			enums.CommandDedInformationReport,
			enums.CommandDedEventNotification:
			isNotify = true
			notify.SetIsComplete(data.IsComplete())
			notify.moreData = data.moreData
			notify.command = data.command
			data.command = enums.CommandNone
			notify.Time = data.Time
			data.Time = time.Time{}
			notify.Data.SetByteBuffer(data.Data)
			notify.Data.Trim()
			notify.Value = data.Value
			data.Value = nil
		default:
			break
		}
	}
	if !isLast || (data.moreData == enums.RequestTypesGBT && reply.Available() != 0) {
		// Clear received notify message.
		if data.command == enums.CommandDataNotification && data.moreData == enums.RequestTypesNone {
			return !isNotify, nil
		}
		return getData(settings, reply, data, notify)
	}
	return !isNotify, nil
}

// GetHDLCAddressFromByteBuffer extracts the HDLC address from a byte buffer.
func GetHDLCAddressFromByteBuffer(buff *types.GXByteBuffer) (int, error) {
	size := 0
	for pos := buff.Position(); pos < buff.Size(); pos++ {
		size++
		val, _ := buff.Uint8At(pos)
		if (val & 0x1) == 1 {
			break
		}
	}
	if size == 1 {
		val, _ := buff.Uint8()
		return int((val & 0xFE) >> 1), nil
	} else if size == 2 {
		val, _ := buff.Uint16()
		result := ((val & 0xFE) >> 1) | ((val & 0xFE00) >> 2)
		return int(result), nil
	} else if size == 4 {
		val, _ := buff.Uint32()
		result := ((val & 0xFE) >> 1) | ((val & 0xFE00) >> 2) |
			((val & 0xFE0000) >> 3) | ((val & 0xFE000000) >> 4)
		return int(result), nil
	}
	return 0, errors.New("wrong size")
}

// AddString adds a string to a byte buffer.
func AddString(value string, bb *types.GXByteBuffer) {
	bb.SetUint8(byte(enums.DataTypeOctetString))
	if value == "" {
		types.SetObjectCount(0, bb)
	} else {
		types.SetObjectCount(len(value), bb)
		bb.Set([]byte(value))
	}
}
