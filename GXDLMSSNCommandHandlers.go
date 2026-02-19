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
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/constants"
	"github.com/Gurux/gxdlms-go/objects"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

func handleRead(settings *settings.GXDLMSSettings,
	server *GXDLMSServer,
	type_ uint8,
	data *types.GXByteBuffer,
	list []*internal.ValueEventArgs,
	reads []*internal.ValueEventArgs,
	replyData *types.GXByteBuffer,
	xml *settings.GXDLMSTranslatorStructure,
	cipheredCommand enums.Command) error {
	sn, err := data.Int16()
	if err != nil {
		return err
	}
	if xml != nil {
		if xml.OutputType() == enums.TranslatorOutputTypeStandardXML {
			xml.AppendStartTag(int(internal.TranslatorTagsVariableAccessSpecification), "", "", true)
		}
		if type_ == uint8(constants.VariableAccessSpecificationParameterisedAccess) {
			xml.AppendStartTag(int(enums.CommandReadRequest)<<8|int(constants.VariableAccessSpecificationParameterisedAccess), "", "", true)
			xml.AppendLineFromTag(int(enums.CommandReadRequest)<<8|int(constants.VariableAccessSpecificationVariableName), "Value", xml.IntegerToHex(sn, 4, false))
			ret, err := data.Uint8()
			if err != nil {
				return err
			}
			xml.AppendLineFromTag(int(internal.TranslatorTagsSelector), "Value", xml.IntegerToHex(ret, 2, false))
			di := internal.GXDataInfo{}
			di.Xml = xml
			xml.AppendStartTag(int(internal.TranslatorTagsParameter), "", "", true)
			internal.GetData(settings, data, &di)
			xml.AppendEndTag(int(internal.TranslatorTagsParameter), false)
			xml.AppendEndTag(int(enums.CommandReadRequest)<<8|int(constants.VariableAccessSpecificationParameterisedAccess), true)
		} else {
			xml.AppendLineFromTag(int(enums.CommandReadRequest)<<8|int(constants.VariableAccessSpecificationVariableName), "Value", xml.IntegerToHex(sn, 4, false))
		}
		if xml.OutputType() == enums.TranslatorOutputTypeStandardXML {
			xml.AppendEndTag(int(internal.TranslatorTagsVariableAccessSpecification), false)
		}
		return nil
	}
	info := FindServerSNObject(server, sn)
	e := internal.NewValueEventArgs2(server, info.Item, info.Index)
	e.Action = info.IsAction
	if type_ == uint8(constants.VariableAccessSpecificationParameterisedAccess) {
		e.Selector, err = data.Uint8()
		if err != nil {
			return err
		}
		di := internal.GXDataInfo{}
		e.Parameters, err = internal.GetData(settings, data, &di)
		if err != nil {
			return err
		}
	}
	// Return error if connection is not established.
	if (settings.Connected&enums.ConnectionStateDlms) == 0 && cipheredCommand == enums.CommandNone &&
		(!e.Action || e.Target.(objects.IGXDLMSBase).Base().ShortName != -1536 || e.Index != 8) {
		replyData.Add(GenerateConfirmedServiceError(enums.ConfirmedServiceErrorInitiateError, enums.ServiceErrorService, uint8(enums.ServiceUnsupported)))
		return nil
	}
	if v, ok := e.Target.(*objects.GXDLMSProfileGeneric); ok {
		e.RowToPdu, err = RowsToPdu(settings, v)
		if err != nil {
			return err
		}
	}
	list = append(list, e)
	if !e.Action && server.NotifyGetAttributeAccess(e) == 0 {
		e.Error = enums.ErrorCodeReadWriteDenied
	} else if e.Action && (server.NotifyGetMethodAccess(e)&int(enums.MethodAccessModeAccess)) == 0 {
		e.Error = enums.ErrorCodeReadWriteDenied
	} else {
		reads = append(reads, e)
	}
	return nil
}

// HandleReadBlockNumberAccess returns the handle read Block in blocks.
//
// Parameters:
//
//	data: Received data.
func handleReadBlockNumberAccess(settings *settings.GXDLMSSettings,
	server *GXDLMSServer,
	data *types.GXByteBuffer,
	replyData *types.GXByteBuffer,
	xml *settings.GXDLMSTranslatorStructure) error {
	blockNumber, err := data.Uint16()
	if err != nil {
		return err
	}
	if xml != nil {
		xml.AppendStartTag(int(enums.CommandReadRequest)<<8|int(constants.VariableAccessSpecificationBlockNumberAccess), "", "", true)
		xml.AppendLine("<BlockNumber Value=\""+xml.IntegerToHex(blockNumber, 4, false)+"\" />", "", nil)
		xml.AppendEndTag(int(enums.CommandReadRequest)<<8|int(constants.VariableAccessSpecificationBlockNumberAccess), false)
		return nil
	}
	if uint32(blockNumber) != settings.BlockIndex {
		bb := types.GXByteBuffer{}
		log.Printf("handleReadRequest failed. Invalid block number. %d/%d", settings.BlockIndex, blockNumber)
		if err := bb.SetUint8(uint8(enums.ErrorCodeDataBlockNumberInvalid)); err != nil {
			return err
		}
		err := getSNPdu(NewGXDLMSSNParameters(settings, enums.CommandReadResponse, 1, byte(constants.SingleReadResponseDataAccessError), &bb, nil), replyData)
		settings.ResetBlockIndex()
		return err
	}
	if settings.Index != settings.Count && server.transaction.Data.Size() < int(settings.MaxPduSize()) {
		reads := []*internal.ValueEventArgs{}
		actions := []*internal.ValueEventArgs{}
		for _, it := range server.transaction.Targets {
			if it.Action {
				actions = append(actions, it)
			} else {
				reads = append(reads, it)
			}
		}
		if len(reads) != 0 {
			server.NotifyRead(reads)
		}
		if len(actions) != 0 {
			server.NotifyPreAction(actions)
		}
		if _, err := getReadData(settings, server.transaction.Targets, server.transaction.Data); err != nil {
			return err
		}
		if len(reads) != 0 {
			server.NotifyPostRead(reads)
		}
		if len(actions) != 0 {
			server.NotifyPostAction(actions)
		}
	}
	settings.IncreaseBlockIndex()
	p := NewGXDLMSSNParameters(settings, enums.CommandReadResponse, 1, byte(constants.SingleReadResponseDataBlockResult),
		nil, server.transaction.Data)
	p.MultipleBlocks = true
	if err := getSNPdu(p, replyData); err != nil {
		return err
	}

	if server.transaction.Data.Size() == server.transaction.Data.Position() {
		server.transaction = nil
		settings.ResetBlockIndex()
	} else {
		server.transaction.Data.Trim()
	}
	return nil
}

func handleReadDataBlockAccess(settings *settings.GXDLMSSettings,
	server *GXDLMSServer,
	command enums.Command,
	data *types.GXByteBuffer,
	cnt int,
	replyData *types.GXByteBuffer,
	xml *settings.GXDLMSTranslatorStructure,
	cipheredCommand enums.Command) error {
	bb := types.GXByteBuffer{}
	lastBlock, err := data.Uint8()
	if err != nil {
		return err
	}
	blockNumber, err := data.Uint16()
	if err != nil {
		return err
	}
	if xml != nil {
		if command == enums.CommandWriteResponse {
			xml.AppendStartTag(int(internal.TranslatorTagsWriteDataBlockAccess), "", "", true)
		} else {
			xml.AppendStartTag(int(internal.TranslatorTagsReadDataBlockAccess), "", "", true)
		}
		xml.AppendStringLine("<LastBlock Value=\"" + xml.IntegerToHex(lastBlock, 2, true) + "\" />")
		xml.AppendStringLine("<BlockNumber Value=\"" + xml.IntegerToHex(blockNumber, 4, true) + "\" />")
		if command == enums.CommandWriteResponse {
			xml.AppendEndTag(int(internal.TranslatorTagsWriteDataBlockAccess), true)
		} else {
			xml.AppendEndTag(int(internal.TranslatorTagsReadDataBlockAccess), true)
		}
		return nil
	}
	if blockNumber != uint16(settings.BlockIndex) {
		log.Printf("handleReadRequest failed. Invalid block number. %d/%d", settings.BlockIndex, blockNumber)
		err = bb.SetUint8(uint8(enums.ErrorCodeDataBlockNumberInvalid))
		if err != nil {
			return err
		}
		err := getSNPdu(NewGXDLMSSNParameters(settings, command, 1, byte(constants.SingleReadResponseDataAccessError), &bb, nil), replyData)
		settings.ResetBlockIndex()
		return err
	}
	count := 1
	type_ := uint8(enums.DataTypeOctetString)
	if command == enums.CommandWriteResponse {
		count, err = types.GetObjectCount(data)
		if err != nil {
			return err
		}
		type_, err = data.Uint8()
		if err != nil {
			return err
		}
	}
	size, err := types.GetObjectCount(data)
	if err != nil {
		return err
	}
	realSize := data.Size() - data.Position()
	if count != 1 || type_ != uint8(enums.DataTypeOctetString) || size != realSize {
		log.Println("handleGetRequest failed. Invalid block size.")
		err = bb.SetUint8(uint8(enums.ErrorCodeDataBlockUnavailable))
		if err != nil {
			return err
		}
		err = getSNPdu(NewGXDLMSSNParameters(settings, command, cnt, byte(constants.SingleReadResponseDataAccessError), &bb, nil), replyData)
		settings.ResetBlockIndex()
		return err
	}
	if server.transaction == nil {
		server.transaction = newGXDLMSLongTransaction(nil, command, data)
	} else {
		server.transaction.Data.SetByteBuffer(data)
	}
	if lastBlock == 0 {
		err = bb.SetUint16(blockNumber)
		if err != nil {
			return err
		}
		settings.IncreaseBlockIndex()
		if command == enums.CommandReadResponse {
			type_ = uint8(constants.SingleReadResponseBlockNumber)
		} else {
			type_ = uint8(constants.SingleWriteResponseBlockNumber)
		}
		return getSNPdu(NewGXDLMSSNParameters(settings, command, cnt, type_, nil, &bb), replyData)
	} else {
		if server.transaction != nil {
			data.SetSize(0)
			err = data.SetByteBuffer(server.transaction.Data)
			if err != nil {
				return err
			}
			server.transaction = nil
		}
		if command == enums.CommandReadResponse {
			handleReadRequest(settings, server, data, replyData, xml, cipheredCommand)
		} else {
			handleWriteRequest(settings, server, data, replyData, xml, cipheredCommand)
		}
		settings.ResetBlockIndex()
	}
	return err
}

// FindSNObject returns the find Short Name object.
//
// Parameters:
//
//	sn: Short name to find.
func FindServerSNObject(server *GXDLMSServer, sn int16) gxSNInfo {
	i := FindSNObject(server.Items().(objects.GXDLMSObjectCollection), sn)
	if i.Item == nil {
		i.Item = server.NotifyFindObject(enums.ObjectTypeNone, int(sn), "").(objects.IGXDLMSBase)
	}
	return i
}

func ReturnSNError(settings *settings.GXDLMSSettings,
	server *GXDLMSServer,
	cmd enums.Command,
	error enums.ErrorCode,
	replyData *types.GXByteBuffer) error {
	bb := types.GXByteBuffer{}
	err := bb.SetUint8(uint8(error))
	if err != nil {
		return err
	}
	err = getSNPdu(NewGXDLMSSNParameters(settings, cmd, 1, byte(constants.SingleReadResponseDataAccessError), &bb, nil), replyData)
	settings.ResetBlockIndex()
	return err
}

// GetReadData returns the get data for Read command.
//
// Parameters:
//
//	list: received objects.
//	data: Data as byte array.
//
// Returns:
//
//	Response type_.
func getReadData(settings *settings.GXDLMSSettings,
	list []*internal.ValueEventArgs,
	data *types.GXByteBuffer) (constants.SingleReadResponse, error) {
	var value any
	var err error
	first := true
	type_ := constants.SingleReadResponseData
	for _, e := range list {
		if e.Handled {
			value = e.Value
		} else {
			// If action.
			if e.Action {
				value, err = e.Target.(objects.IGXDLMSBase).Invoke(settings, e)
			} else {
				value, err = e.Target.(objects.IGXDLMSBase).GetValue(settings, e)
			}
			if err != nil {
				return type_, err
			}
		}
		if e.Error == 0 {
			if !first && len(list) != 1 {
				err := data.SetUint8(uint8(constants.SingleReadResponseData))
				if err != nil {
					return type_, err
				}
			}
			// If action.
			if e.Action {
				dt, err := internal.GetDLMSDataType(reflect.TypeOf(value))
				if err != nil {
					return type_, err
				}
				internal.SetData(settings, data, dt, value)
			} else {
				err = appendData(settings, e.Target.(objects.IGXDLMSBase), e.Index, data, value)
				if err != nil {
					return type_, err
				}
			}
		} else {
			if !first && len(list) != 1 {
				data.SetUint8(uint8(constants.SingleReadResponseDataAccessError))
			}
			err = data.SetUint8(uint8(e.Error))
			if err != nil {
				return type_, err
			}
			type_ = constants.SingleReadResponseDataAccessError
		}
		first = false
	}
	return type_, nil
}

// GenerateWriteResponse returns the generate write reply.
func GenerateWriteResponse(settings *settings.GXDLMSSettings, results *types.GXByteBuffer, replyData *types.GXByteBuffer) error {
	bb := types.NewGXByteBufferWithCapacity(2 * results.Size())
	for pos := 0; pos != results.Size(); pos++ {
		ret, err := results.Uint8At(pos)
		if err != nil {
			return err
		}
		// If meter returns error.
		if ret != 0 {
			err = bb.SetUint8(1)
			if err != nil {
				return err
			}
		}
		err = bb.SetUint8(ret)
		if err != nil {
			return err
		}
	}
	p := NewGXDLMSSNParameters(settings, enums.CommandWriteResponse, results.Size(), 0xFF, nil, bb)
	return getSNPdu(p, replyData)
}

func FindSNObject(items objects.GXDLMSObjectCollection, sn int16) gxSNInfo {
	i := gxSNInfo{}
	var offset int
	var count int
	for _, it := range items {
		tmp := int16(it.Base().ShortName)
		if sn >= tmp {
			if sn < tmp+int16(it.GetAttributeCount())*8 {
				i.IsAction = false
				i.Item = it
				i.Index = uint8((sn-int16(i.Item.Base().ShortName))/8) + 1
				break
			} else {
				//If method is accessed.
				getActionInfo(it.Base().ObjectType(), &offset, &count)
				if sn < tmp+int16(offset)+(8*int16(count)) {
					i.Item = it
					i.IsAction = true
					i.Index = uint8((sn-tmp-int16(offset))/8) + 1
					break
				}
			}
		}
	}
	return i
}

// handleReadRequest returns the handle read request.
//
// Parameters:
//
//	settings: DLMS settings.
//	server: DLMS server.
//	data: Received data.
func handleReadRequest(settings *settings.GXDLMSSettings, server *GXDLMSServer,
	data *types.GXByteBuffer,
	replyData *types.GXByteBuffer,
	xml *settings.GXDLMSTranslatorStructure,
	cipheredCommand enums.Command) error {
	// Return error if connection is not established.
	if xml == nil && (settings.Connected&enums.ConnectionStateDlms) == 0 && cipheredCommand == enums.CommandNone {
		return replyData.Add(GenerateConfirmedServiceError(enums.ConfirmedServiceErrorInitiateError, enums.ServiceErrorService, uint8(enums.ServiceUnsupported)))
	}
	bb := types.GXByteBuffer{}
	list := []*internal.ValueEventArgs{}
	reads := []*internal.ValueEventArgs{}
	// If get next frame.
	if xml == nil && data.Size() == 0 {
		if replyData.Available() != 0 {
			// Return existing PDU first.
			return nil
		}
		for _, it := range server.transaction.Targets {
			list = append(list, it)
		}
	} else {
		cnt, err := types.GetObjectCount(data)
		if err != nil {
			return err
		}
		if xml != nil {
			xml.AppendStartTag(int(enums.CommandReadRequest), "Qty", xml.IntegerToHex(cnt, 2, false), true)
		}
		for pos := 0; pos != cnt; pos++ {
			type_, err := data.Uint8()
			if err != nil {
				return err
			}
			if type_ == uint8(constants.VariableAccessSpecificationBlockNumberAccess) || type_ == uint8(constants.VariableAccessSpecificationParameterisedAccess) {
				err = handleRead(settings, server, type_, data, list, reads, replyData, xml, cipheredCommand)
			} else if type_ == uint8(constants.VariableAccessSpecificationBlockNumberAccess) {
				err = handleReadBlockNumberAccess(settings, server, data, replyData, xml)
				if xml != nil {
					xml.AppendEndTag(int(enums.CommandReadRequest), true)
				}
				return err
			} else if type_ == uint8(constants.VariableAccessSpecificationReadDataBlockAccess) {
				err = handleReadDataBlockAccess(settings, server, enums.CommandReadResponse, data, cnt, replyData, xml, cipheredCommand)
				if xml != nil {
					xml.AppendEndTag(int(enums.CommandReadRequest), true)
				}
				return err
			} else {
				if xml != nil {
					xml.AppendEndTag(int(enums.CommandReadRequest), true)
				} else {
					err = ReturnSNError(settings, server, enums.CommandReadResponse, enums.ErrorCodeReadWriteDenied, replyData)
				}
				return err
			}
		}
		if len(reads) != 0 {
			server.NotifyRead(reads)
		}
	}
	if xml != nil {
		xml.AppendEndTag(int(enums.CommandReadRequest), true)
		return nil
	}
	requestType, err := getReadData(settings, list, &bb)
	if err != nil {
		return err
	}
	if len(reads) != 0 {
		server.NotifyPostRead(reads)
	}
	p := NewGXDLMSSNParameters(settings, enums.CommandReadResponse, len(list), byte(requestType), nil, &bb)
	err = getSNPdu(p, replyData)
	if server.transaction == nil && (bb.Available() != 0 || settings.Count != settings.Index) {
		for _, it := range list {
			reads = append(reads, it)
		}
		server.transaction = newGXDLMSLongTransaction(reads, enums.CommandReadRequest, &bb)
	} else if server.transaction != nil {
		err = replyData.SetByteBuffer(&bb)
	}
	return err
}

// handleWriteRequest returns the handle write request.
func handleWriteRequest(conf *settings.GXDLMSSettings,
	server *GXDLMSServer,
	data *types.GXByteBuffer,
	replyData *types.GXByteBuffer,
	xml *settings.GXDLMSTranslatorStructure,
	cipheredCommand enums.Command) error {
	var value any
	// Get object count.
	targets := []gxSNInfo{}
	cnt, err := types.GetObjectCount(data)
	if err != nil {
		return err
	}
	if xml != nil {
		xml.AppendStartTag(int(enums.CommandWriteRequest), "", "", true)
		xml.AppendStartTag(int(internal.TranslatorTagsListOfVariableAccessSpecification), "Qty", xml.IntegerToHex(cnt, 2, false), true)
		if xml.OutputType() == enums.TranslatorOutputTypeStandardXML {
			xml.AppendStartTag(int(internal.TranslatorTagsVariableAccessSpecification), "", "", true)
		}
	}
	var di internal.GXDataInfo
	var info gxSNInfo
	results := types.NewGXByteBufferWithCapacity(cnt)
	for pos := 0; pos != cnt; pos++ {
		type_, err := data.Uint8()
		if err != nil {
			return err
		}
		if type_ == uint8(constants.VariableAccessSpecificationVariableName) {
			sn, err := data.Uint16()
			if err != nil {
				return err
			}
			if xml != nil {
				xml.AppendLineFromTag(int(enums.CommandWriteRequest)<<8|int(type_), "Value", xml.IntegerToHex(sn, 4, false))
			} else {
				info = FindServerSNObject(server, int16(sn))
				targets = append(targets, info)
				// If target is unknown.
				if info.Item == nil {
					err = results.SetUint8(enums.ErrorCodeUndefinedObject)
					if err != nil {
						return err
					}
				} else {
					err = results.SetUint8(enums.ErrorCodeOk)
					if err != nil {
						return err
					}
				}
			}
		} else if type_ == uint8(constants.VariableAccessSpecificationWriteDataBlockAccess) {
			// Return error if connection is not established.
			if xml == nil && (conf.Connected&enums.ConnectionStateDlms) == 0 && cipheredCommand == enums.CommandNone {
				replyData.Add(GenerateConfirmedServiceError(enums.ConfirmedServiceErrorInitiateError, enums.ServiceErrorService, uint8(enums.ServiceUnsupported)))
				return nil
			}
			handleReadDataBlockAccess(conf, server, enums.CommandWriteResponse, data, cnt, replyData, xml, cipheredCommand)
			if xml == nil {
				return nil
			}
		} else {
			err = results.SetUint8(enums.ErrorCodeHardwareFault)
			if err != nil {
				return err
			}
		}
	}
	if xml != nil {
		if xml.OutputType() == enums.TranslatorOutputTypeStandardXML {
			xml.AppendEndTag(int(internal.TranslatorTagsVariableAccessSpecification), true)
		}
		xml.AppendEndTag(int(internal.TranslatorTagsListOfVariableAccessSpecification), true)
	}
	cnt, err = types.GetObjectCount(data)
	if err != nil {
		return err
	}
	di = internal.GXDataInfo{}
	di.Xml = xml
	if xml != nil {
		xml.AppendStartTag(int(internal.TranslatorTagsListOfData), "Qty", xml.IntegerToHex(cnt, 2, false), true)
	}
	for pos := 0; pos != cnt; pos++ {
		di.Clear()
		if xml != nil {
			if xml.OutputType() == enums.TranslatorOutputTypeStandardXML {
				xml.AppendStartTag(int(enums.CommandWriteRequest<<8|int(constants.SingleReadResponseData)), "", "", true)
			}
			value, err = internal.GetData(conf, data, &di)
			if err != nil {
				return err
			}
			if !di.Complete {
				value = types.ToHexWithRange(data.Array(), false, data.Position(), data.Size()-data.Position())
				xml.AppendLineFromTag(settings.DataTypeOffset+int(di.Type), "Value", value.(string))
			}
			if xml.OutputType() == enums.TranslatorOutputTypeStandardXML {
				xml.AppendEndTag(int(enums.CommandWriteRequest<<8|constants.SingleReadResponseData), true)
			}
		} else if ret, err := results.Uint8At(pos); err == nil && ret == 0 {
			access := true
			// If object has found.
			target := targets[pos]
			value, err = internal.GetData(conf, data, &di)
			if err != nil {
				return err
			}
			e := internal.NewValueEventArgs2(server, target.Item, target.Index)
			if target.IsAction {
				am := server.NotifyGetMethodAccess(e)
				// If action is denied.
				if (am & int(enums.MethodAccessModeAccess)) != 0 {
					access = false
				}
			} else {
				if _, ok := value.([]byte); ok {
					dt, err := target.Item.GetDataType(int(target.Index))
					if err != nil {
						return err
					}
					value, err = internal.ChangeTypeFromByteArray(conf, value.([]byte), dt)
					if err != nil {
						return err
					}
				}
			}
			am := server.NotifyGetAttributeAccess(e)
			// If write is denied.
			if (am & int(enums.AccessModeWrite)) == 0 {
				access = false
			}
			if access {
				if target.IsAction {
					e.Parameters = value
					actions := []*internal.ValueEventArgs{e}
					server.NotifyPreAction(actions)
					if !e.Handled {
						reply, err := target.Item.Invoke(conf, e)
						if err != nil {
							return err
						}
						server.NotifyPostAction(actions)
						if _, ok := target.Item.(*objects.GXDLMSAssociationShortName); ok {
							bb := types.NewGXByteBuffer()
							err = bb.SetUint8(uint8(enums.DataTypeOctetString))
							if err != nil {
								return err
							}
							err = bb.SetUint8(uint8(len(reply)))
							if err != nil {
								return err
							}
							err = bb.Set(reply)
							if err != nil {
								return err
							}

							p := NewGXDLMSSNParameters(conf, enums.CommandReadResponse, 1, 0, nil, bb)
							return getSNPdu(p, replyData)
						}
					}
				} else {
					e.Value = value
					server.NotifyWrite([]*internal.ValueEventArgs{e})
					if e.Error != 0 {
						err = results.SetUint8At(pos, uint8(e.Error))
						if err != nil {
							return err
						}
					} else if !e.Handled {
						target.Item.SetValue(conf, e)
						server.NotifyPostWrite([]*internal.ValueEventArgs{e})
					}
				}
			} else {
				err = results.SetUint8At(pos, uint8(enums.ErrorCodeReadWriteDenied))
				if err != nil {
					return err
				}
			}
		}
	}
	if xml != nil {
		xml.AppendEndTag(int(internal.TranslatorTagsListOfData), true)
		xml.AppendEndTag(int(enums.CommandWriteRequest), true)
		return nil
	}
	return GenerateWriteResponse(conf, results, replyData)
}

// HandleInformationReport returns the handle Information Report.
//
// Parameters:
//
//	settings: DLMS settings.
func handleInformationReport(settings *settings.GXDLMSSettings,
	reply *GXReplyData,
	list []*types.GXKeyValuePair[objects.IGXDLMSBase, int]) error {
	reply.Time = time.Time{}
	len, err := reply.Data.Uint8()
	if err != nil {
		return err
	}
	// If date time is given.
	var tmp []byte
	if len != 0 {
		tmp = make([]byte, len)
		err = reply.Data.Get(tmp)
		if err != nil {
			return err
		}
		ret, err := internal.ChangeTypeFromByteArray(settings, tmp, enums.DataTypeDateTime)
		if err != nil {
			return err
		}
		reply.Time = ret.(types.GXDateTime).Value
	}
	var type_ uint8
	count, err := types.GetObjectCount(reply.Data)
	if err != nil {
		return err
	}
	if reply.xml != nil {
		reply.xml.AppendStartTag(int(enums.CommandInformationReport), "", "", true)
		if !reply.Time.IsZero() {
			reply.xml.AppendComment(fmt.Sprint(reply.Time))
			if reply.xml.OutputType() == enums.TranslatorOutputTypeSimpleXML {
				reply.xml.AppendLineFromTag(int(internal.TranslatorTagsCurrentTime), "", types.ToHex(tmp, false))
			} else {
				reply.xml.AppendLineFromTag(int(internal.TranslatorTagsCurrentTime), "", internal.GeneralizedTime(&reply.Time))
			}
		}
		reply.xml.AppendStartTag(int(internal.TranslatorTagsListOfVariableAccessSpecification), "Qty",
			reply.xml.IntegerToHex(count, 2, false), true)
	}
	for pos := 0; pos != count; pos++ {
		type_, err = reply.Data.Uint8()
		if err != nil {
			return err
		}
		if type_ == uint8(constants.VariableAccessSpecificationVariableName) {
			sn, err := reply.Data.Uint16()
			if err != nil {
				return err
			}
			if reply.xml != nil {
				reply.xml.AppendLineFromTag(int(enums.CommandWriteRequest)<<8|int(constants.VariableAccessSpecificationVariableName), "Value",
					reply.xml.IntegerToHex(sn, 4, false))
			} else {
				info := FindSNObject(settings.Objects.(objects.GXDLMSObjectCollection), int16(sn))
				if info.Item != nil {
					list = append(list, types.NewGXKeyValuePair[objects.IGXDLMSBase, int](info.Item, int(info.Index)))
				} else {
					log.Println(fmt.Sprintf("Unknown object : %v.", sn))
				}
			}
		}
	}
	if reply.xml != nil {
		reply.xml.AppendEndTag(int(internal.TranslatorTagsListOfVariableAccessSpecification), true)
		reply.xml.AppendStartTag(int(internal.TranslatorTagsListOfData), "Qty", reply.xml.IntegerToHex(count, 2, false), true)
	}
	count, err = types.GetObjectCount(reply.Data)
	if err != nil {
		return err
	}
	di := internal.GXDataInfo{}
	di.Xml = reply.xml
	for pos := 0; pos != count; pos++ {
		di.Clear()
		if reply.xml != nil {
			_, err = internal.GetData(settings, reply.Data, &di)
			if err != nil {
				return err
			}
		} else {
			v := internal.NewValueEventArgs(settings, list[pos].Key, uint8(list[pos].Value))
			v.Value, err = internal.GetData(settings, reply.Data, &di)
			if err != nil {
				return err
			}
			list[pos].Key.SetValue(settings, v)
		}
	}
	if reply.xml != nil {
		reply.xml.AppendEndTag(int(internal.TranslatorTagsListOfData), true)
		reply.xml.AppendEndTag(int(enums.CommandInformationReport), true)
	}
	return nil
}
