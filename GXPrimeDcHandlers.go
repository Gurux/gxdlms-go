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
	"reflect"
	"strconv"
	"time"

	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/objects"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// HandleNewDeviceNotification returns the handle new device notification message.
//
// Parameters:
//
//	data: Received data
//	replyData: Reply data.
//	xml: XML settings.
func primeDcHandleNewDeviceNotification(data *types.GXByteBuffer, replyData *GXReplyData, xml *settings.GXDLMSTranslatorStructure) error {
	deviceId, err := data.Uint16()
	if err != nil {
		return err
	}
	capabilities, err := data.Uint16()
	if err != nil {
		return err
	}
	len, err := data.Uint8()
	if err != nil {
		return err
	}
	id := make([]byte, len)
	err = data.Get(id)
	if err != nil {
		return err
	}
	eui48 := make([]byte, 6)
	err = data.Get(eui48)
	if err != nil {
		return err
	}
	if replyData != nil {
		replyData.PrimeDc = new(GXDLMSPrimeDataConcentrator)
		replyData.PrimeDc.Type = enums.PrimeDcMsgTypeNewDeviceNotification
		replyData.PrimeDc.DeviceID = deviceId
		replyData.PrimeDc.Capabilities = capabilities
		replyData.PrimeDc.DlmsId = id
		replyData.PrimeDc.Eui48 = eui48
	}
	if xml != nil {
		xml.AppendLine("<DeviceId Value=\""+strconv.Itoa(int(deviceId))+"\" />", "", nil)
		xml.AppendLine("<Capabilities Value=\""+strconv.Itoa(int(capabilities))+"\" />", "", nil)
		if xml.Comments {
			xml.AppendComment("DLMS ID " + string(id))
		}
		xml.AppendLine("<DlmsId Value=\""+types.ToHex(id, false)+"\" />", "", nil)
		xml.AppendLine("<Eui48 Value=\""+types.ToHex(eui48, false)+"\" />", "", nil)
	}
	return nil
}

// HandleRemoveDeviceNotification returns the handle remove device notification message.
//
// Parameters:
//
//	data: Received data
//	replyData: Reply data.
//	xml: XML settings.
func primeDcHandleRemoveDeviceNotification(data *types.GXByteBuffer, replyData *GXReplyData, xml *settings.GXDLMSTranslatorStructure) error {
	deviceId, err := data.Uint16()
	if err != nil {
		return err
	}
	if replyData != nil {
		replyData.PrimeDc = new(GXDLMSPrimeDataConcentrator)
		replyData.PrimeDc.Type = enums.PrimeDcMsgTypeRemoveDeviceNotification
		replyData.PrimeDc.DeviceID = deviceId
	}
	if xml != nil {
		xml.AppendLine("<DeviceId Value=\""+strconv.Itoa(int(deviceId))+"\" />", "", nil)
	}
	return nil
}

// HandleStartReportingMeters returns the handle start reporting meters notification message.
//
// Parameters:
//
//	data: Received data
//	replyData: Reply data.
//	xml: XML settings.
func primeDcHandleStartReportingMeters(data *types.GXByteBuffer, replyData *GXReplyData, xml *settings.GXDLMSTranslatorStructure) error {
	if replyData != nil {
		replyData.PrimeDc = new(GXDLMSPrimeDataConcentrator)
		replyData.PrimeDc.Type = enums.PrimeDcMsgTypeStartReportingMeters
	}
	return nil
}

// HandleDeleteMeters returns the handle delete meters notification message.
//
// Parameters:
//
//	data: Received data
//	replyData: Reply data.
//	xml: XML settings.
func primeDcHandleDeleteMeters(data *types.GXByteBuffer, replyData *GXReplyData, xml *settings.GXDLMSTranslatorStructure) error {
	deviceId, err := data.Uint16()
	if err != nil {
		return err
	}
	if replyData != nil {
		replyData.PrimeDc = new(GXDLMSPrimeDataConcentrator)
		replyData.PrimeDc.Type = enums.PrimeDcMsgTypeDeleteMeters
		replyData.PrimeDc.DeviceID = deviceId
	}
	if xml != nil {
		xml.AppendLine("<DeviceId Value=\""+strconv.Itoa(int(deviceId))+"\" />", "", nil)
	}
	return nil
}

// HandleEnableAutoClose returns the handle enable auto close notification message.
//
// Parameters:
//
//	data: Received data
//	replyData: Reply data.
//	xml: XML settings.
func primeDcHandleEnableAutoClose(data *types.GXByteBuffer, replyData *GXReplyData, xml *settings.GXDLMSTranslatorStructure) error {
	deviceId, err := data.Uint16()
	if err != nil {
		return err
	}
	if replyData != nil {
		replyData.PrimeDc = new(GXDLMSPrimeDataConcentrator)
		replyData.PrimeDc.Type = enums.PrimeDcMsgTypeEnableAutoClose
		replyData.PrimeDc.DeviceID = deviceId
	}
	if xml != nil {
		xml.AppendLine("<DeviceId Value=\""+strconv.Itoa(int(deviceId))+"\" />", "", nil)
	}
	return nil
}

// HandleDisableAutoClose returns the handle disable auto close notification message.
//
// Parameters:
//
//	data: Received data
//	replyData: Reply data.
//	xml: XML settings.
func primeDcHandleDisableAutoClose(data *types.GXByteBuffer, replyData *GXReplyData, xml *settings.GXDLMSTranslatorStructure) error {
	deviceId, err := data.Uint16()
	if err != nil {
		return err
	}
	if replyData != nil {
		replyData.PrimeDc = new(GXDLMSPrimeDataConcentrator)
		replyData.PrimeDc.Type = enums.PrimeDcMsgTypeDisableAutoClose
		replyData.PrimeDc.DeviceID = deviceId
	}
	if xml != nil {
		xml.AppendLine("<DeviceId Value=\""+strconv.Itoa(int(deviceId))+"\" />", "", nil)
	}
	return nil
}

func primeDcAppendAttributeDescriptor(xml *settings.GXDLMSTranslatorStructure, ci int, ln []byte, attributeIndex uint8) error {
	xml.AppendStartTag(int(internal.TranslatorTagsAttributeDescriptor), "", "", false)
	if xml.Comments {
		xml.AppendComment(enums.ObjectType(ci).String())
	}
	xml.AppendLineFromTag(int(internal.TranslatorTagsClassId), "Value", xml.IntegerToHex(int(ci), 4, false))
	ret, err := ToLogicalName(ln)
	if err != nil {
		return err
	}
	xml.AppendComment(ret)
	xml.AppendLineFromTag(int(internal.TranslatorTagsInstanceId), "Value", types.ToHex(ln, false))
	obj := objects.CreateObject(enums.ObjectType(ci))
	if obj != nil {
		xml.AppendComment(obj.GetNames()[attributeIndex-1])
	}

	xml.AppendLineFromTag(int(internal.TranslatorTagsAttributeId), "Value", xml.IntegerToHex(attributeIndex, 2, false))
	xml.AppendEndTag(int(internal.TranslatorTagsAttributeDescriptor), false)
	return nil
}

func primeDcappendMethodDescriptor(xml *settings.GXDLMSTranslatorStructure, ci int, ln []byte, attributeIndex uint8) error {
	xml.AppendStartTag(int(internal.TranslatorTagsMethodDescriptor), "", "", false)
	if xml.Comments {
		xml.AppendComment(enums.ObjectType(ci).String())
	}
	xml.AppendLineFromTag(int(internal.TranslatorTagsClassId), "Value", xml.IntegerToHex(int(ci), 4, false))
	ret, err := ToLogicalName(ln)
	if err != nil {
		return err
	}
	xml.AppendComment(ret)
	xml.AppendLineFromTag(int(internal.TranslatorTagsInstanceId), "Value", types.ToHex(ln, false))
	obj := objects.CreateObject(enums.ObjectType(ci))
	if obj != nil {
		xml.AppendComment(obj.GetMethodNames()[attributeIndex-1])
	}
	xml.AppendLineFromTag(int(internal.TranslatorTagsMethodId), "Value", xml.IntegerToHex(attributeIndex, 2, false))
	xml.AppendEndTag(int(internal.TranslatorTagsMethodDescriptor), false)
	return nil
}

// primeDcGetRequestWithList returns the handle get request with list command.
//
// Parameters:
//
//	data: Received data.
func primeDcGetRequestWithList(settings *settings.GXDLMSSettings,
	invokeID uint8,
	server *GXDLMSServer,
	data *types.GXByteBuffer,
	replyData *types.GXByteBuffer,
	xml *settings.GXDLMSTranslatorStructure,
	cipheredCommand enums.Command) error {
	bb := types.GXByteBuffer{}
	var e *internal.ValueEventArgs
	var pos int
	cnt, err := types.GetObjectCount(data)
	if err != nil {
		return err
	}
	types.SetObjectCount(cnt, &bb)
	list := []*internal.ValueEventArgs{}
	if xml != nil {
		xml.AppendStartTag(int(internal.TranslatorTagsAttributeDescriptorList), "Qty", xml.IntegerToHex(cnt, 2, false), false)
	}
	for pos := 0; pos != cnt; pos++ {
		ret, err := data.Uint16()
		if err != nil {
			return err
		}
		ci := enums.ObjectType(ret)
		ln := make([]byte, 6)
		data.Get(ln)
		attributeIndex, err := data.Uint8()
		if err != nil {
			return err
		}
		// AccessSelection
		selection, err := data.Uint8()
		if err != nil {
			return err
		}
		selector := uint8(0)
		var parameters any
		if selection != 0 {
			selector, err = data.Uint8()
			if err != nil {
				return err
			}
			info := internal.GXDataInfo{}
			parameters, err = internal.GetData(settings, data, &info)
			if err != nil {
				return err
			}
		}
		if xml != nil {
			xml.AppendStartTag(int(internal.TranslatorTagsAttributeDescriptorWithSelection), "", "", false)
			xml.AppendStartTag(int(internal.TranslatorTagsAttributeDescriptor), "", "", false)
			xml.AppendComment(ci.String())
			xml.AppendLineFromTag(int(internal.TranslatorTagsClassId), "Value", xml.IntegerToHex(int(ci), 4, false))
			ret, err := helpers.ToLogicalName(ln)
			if err != nil {
				return err
			}
			xml.AppendComment(ret)
			xml.AppendLineFromTag(int(internal.TranslatorTagsInstanceId), "Value", types.ToHex(ln, false))
			xml.AppendLineFromTag(int(internal.TranslatorTagsAttributeId), "Value", xml.IntegerToHex(attributeIndex, 2, false))
			xml.AppendEndTag(int(internal.TranslatorTagsAttributeDescriptor), false)
			xml.AppendEndTag(int(internal.TranslatorTagsAttributeDescriptorWithSelection), false)
		} else {
			ret, err := helpers.ToLogicalName(ln)
			if err != nil {
				return err
			}
			var obj objects.IGXDLMSBase
			if ci == enums.ObjectTypeAssociationLogicalName && ret == "0.0.40.0.0.255" {
				obj = settings.AssignedAssociation().(objects.IGXDLMSBase)
			}
			if obj == nil {
				obj = getObjectCollection(settings.Objects).FindByLN(ci, ret)
			}
			if obj == nil {
				obj = server.NotifyFindObject(ci, 0, ret).(objects.IGXDLMSBase)
			}
			if obj == nil {
				e = internal.NewValueEventArgs2(server, obj, attributeIndex)
				e.Error = enums.ErrorCodeUndefinedObject
				list = append(list, e)
			} else {
				arg := internal.NewValueEventArgs2(server, obj, attributeIndex)
				arg.Selector = selector
				arg.Parameters = parameters
				arg.InvokeId = uint32(invokeID)
				access := server.NotifyGetAttributeAccess(arg)
				if (access&int(enums.AccessModeRead)) == 0 && (access&int(enums.AccessModeAuthenticatedRead)) == 0 {
					arg.Error = enums.ErrorCodeReadWriteDenied
					list = append(list, arg)
				} else {
					list = append(list, arg)
				}
			}
		}
	}
	if xml != nil {
		xml.AppendEndTag(int(internal.TranslatorTagsAttributeDescriptorList), false)
		return nil
	}
	server.NotifyRead(list)
	var value any
	pos = 0
	for _, it := range list {
		if it.Handled {
			value = it.Value
		} else {
			value, err = it.Target.(objects.IGXDLMSBase).GetValue(settings, it)
			if err != nil {
				return err
			}
		}
		err = bb.SetUint8(uint8(it.Error))
		if err != nil {
			return err
		}
		if it.ByteArray {
			err = bb.Set(value.([]byte))
			if err != nil {
				return err
			}
		} else {
			err = appendData(settings, it.Target.(objects.IGXDLMSBase), it.Index, &bb, value)
			if err != nil {
				return err
			}
		}
		invokeID = uint8(it.InvokeId)
		pos++
	}
	server.NotifyPostRead(list)
	p := NewGXDLMSLNParameters(settings, uint32(invokeID), enums.CommandGetResponse, 3, nil, &bb, 0xFF, cipheredCommand)
	err = getLNPdu(p, replyData)
	if settings.Index != settings.Count || bb.Available() != 0 {
		server.SetTransaction(newGXDLMSLongTransaction(list, enums.CommandGetRequest, &bb))
	}
	return err
}

func primeDcHandleSetRequestNormal(settings *settings.GXDLMSSettings,
	server *GXDLMSServer,
	data *types.GXByteBuffer,
	type_ uint8,
	p *GXDLMSLNParameters,
	replyData *types.GXByteBuffer,
	xml *settings.GXDLMSTranslatorStructure) error {
	reply := internal.GXDataInfo{}
	// CI
	ret, err := data.Uint16()
	if err != nil {
		return err
	}
	ci := enums.ObjectType(ret)
	if data.Available() < 8 {
		if xml != nil {
			xml.AppendComment("Logical name is missing.")
			xml.AppendComment("Attribute Id is missing.")
			xml.AppendComment("Access Selection is missing.")
			return nil
		}
		return errors.New("Set request is not complete.")
	}
	ln := make([]byte, 6)
	err = data.Get(ln)
	if err != nil {
		return err
	}
	// Attribute index.
	index, err := data.Uint8()
	if err != nil {
		return err
	}
	_, err = data.Uint8()
	if err != nil {
		return err
	}
	if type_ == 2 {
		lastBlock, err := data.Uint8()
		if err != nil {
			return err
		}
		p.multipleBlocks = lastBlock == 0
		blockNumber, err := data.Uint32()
		if err != nil {
			return err
		}
		if blockNumber != settings.BlockIndex {
			log.Println("HandleSetRequest failed. Invalid block number %d/%d.", settings.BlockIndex, blockNumber)
			p.status = uint8(enums.ErrorCodeDataBlockNumberInvalid)
			return nil
		}
		settings.IncreaseBlockIndex()
		size, err := types.GetObjectCount(data)
		if err != nil {
			return err
		}
		realSize := data.Size() - data.Position()
		if size != realSize {
			log.Println("HandleSetRequest failed. Invalid block size. %d/%d.", size, realSize)
			p.status = uint8(enums.ErrorCodeDataBlockUnavailable)
			return nil
		}
		if xml != nil {
			primeDcAppendAttributeDescriptor(xml, int(ci), ln, index)
			xml.AppendStartTag(int(internal.TranslatorTagsDataBlock), "", "", true)
			xml.AppendLineFromTag(int(internal.TranslatorTagsLastBlock), "Value", xml.IntegerToHex(lastBlock, 2, false))
			xml.AppendLineFromTag(int(internal.TranslatorTagsBlockNumber), "Value", xml.IntegerToHex(blockNumber, 8, false))
			xml.AppendLineFromTag(int(internal.TranslatorTagsRawData), "Value", data.RemainingHexString(false))
			xml.AppendEndTag(int(internal.TranslatorTagsDataBlock), false)
		}
	}
	if xml != nil {
		primeDcAppendAttributeDescriptor(xml, int(ci), ln, index)
		xml.AppendStartTag(int(internal.TranslatorTagsValue), "", "", true)
		di := internal.GXDataInfo{}
		di.Xml = xml
		value, err := internal.GetData(settings, data, &di)
		if err != nil {
			return err
		}
		if !di.Complete {
			types.ToHexWithRange(data.Array(), false, data.Position(), data.Size()-data.Position())
		} else if v, ok := value.([]byte); ok {
			types.ToHex(v, false)
		}
		xml.AppendEndTag(int(internal.TranslatorTagsValue), true)
		return nil
	}
	var value any
	if !p.multipleBlocks {
		settings.ResetBlockIndex()
		value, err = internal.GetData(settings, data, &reply)
		if err != nil {
			return err
		}
	}
	ln2, err := helpers.ToLogicalName(ln)
	if err != nil {
		return err
	}
	var obj objects.IGXDLMSBase
	if ci == enums.ObjectTypeAssociationLogicalName && ln2 == "0.0.40.0.0.255" {
		obj = settings.AssignedAssociation().(objects.IGXDLMSBase)
	}
	if obj == nil {
		obj = getObjectCollection(settings.Objects).FindByLN(ci, ln2)
	}
	if obj == nil {
		obj = server.NotifyFindObject(ci, 0, ln2).(objects.IGXDLMSBase)
	}
	// If target is unknown.
	if obj == nil {
		p.status = uint8(enums.ErrorCodeUndefinedObject)
	} else {
		e := internal.NewValueEventArgs2(server, obj, index)
		e.InvokeId = p.invokeId
		am := server.NotifyGetAttributeAccess(e)
		// If write is denied.
		if (am & int(enums.AccessModeWrite)) == 0 {
			p.status = uint8(enums.ErrorCodeReadWriteDenied)
		} else {
			if _, ok := value.([]byte); ok {
				dt, err := obj.GetDataType(int(index))
				if dt != enums.DataTypeNone && dt != enums.DataTypeOctetString && dt != enums.DataTypeStructure {
					value, err = internal.ChangeTypeFromByteArray(settings, value.([]byte), dt)
				}
				if err != nil {
					return err
				}
			}
			e.Value = value
			list := []*internal.ValueEventArgs{e}
			if p.multipleBlocks {
				server.SetTransaction(newGXDLMSLongTransaction(list, enums.CommandGetRequest, data))
			}
			server.NotifyWrite(list)
			if e.Error != 0 {
				p.status = uint8(e.Error)
			} else if !e.Handled && !p.multipleBlocks {
				obj.SetValue(settings, e)
				server.NotifyPostWrite(list)
				if e.Error != 0 {
					p.status = uint8(e.Error)
				}
			}
			p.invokeId = e.InvokeId
		}
	}
	return nil
}

func primeDcHandleSetRequestWithDataBlock(settings *settings.GXDLMSSettings,
	server *GXDLMSServer,
	data *types.GXByteBuffer,
	p *GXDLMSLNParameters,
	replyData *types.GXByteBuffer,
	xml *settings.GXDLMSTranslatorStructure) error {
	ret := uint8(0)
	reply := internal.GXDataInfo{}
	reply.Xml = xml
	lastBlock, err := data.Uint8()
	if err != nil {
		return err
	}
	p.multipleBlocks = lastBlock == 0
	blockNumber, err := data.Uint32()
	if err != nil {
		return err
	}
	if xml == nil && blockNumber != settings.BlockIndex {
		log.Println("HanleSetRequestWithDataBlock failed. Invalid block number. %d/%d", settings.BlockIndex, blockNumber)
		ret = uint8(enums.ErrorCodeDataBlockNumberInvalid)
	} else {
		settings.IncreaseBlockIndex()
		size, err := types.GetObjectCount(data)
		if err != nil {
			return err
		}
		realSize := data.Size() - data.Position()
		if size != realSize {
			log.Println("HanleSetRequestWithDataBlock failed. Invalid block size. %d/%d.", size, realSize)
			ret = uint8(enums.ErrorCodeDataBlockUnavailable)
		}
		if xml != nil {
			xml.AppendStartTag(int(internal.TranslatorTagsDataBlock), "", "", true)
			xml.AppendLineFromTag(int(internal.TranslatorTagsLastBlock), "Value", xml.IntegerToHex(lastBlock, 2, false))
			xml.AppendLineFromTag(int(internal.TranslatorTagsBlockNumber), "Value", xml.IntegerToHex(blockNumber, 8, false))
			xml.AppendLineFromTag(int(internal.TranslatorTagsRawData), "Value", data.RemainingHexString(false))
			xml.AppendEndTag(int(internal.TranslatorTagsDataBlock), true)
			return nil
		}
		err = server.transaction.Data.SetByteBuffer(data)
		if err != nil {
			return err
		}
		// If all data is received.
		if !p.multipleBlocks {
			value, err := internal.GetData(settings, server.transaction.Data, &reply)
			if err != nil {
				return err
			}
			if v, ok := value.([]byte); ok {
				t := server.transaction.Targets[0]
				dt, err := t.Target.(objects.IGXDLMSBase).GetDataType(int(t.Index))
				if err != nil {
					return err
				}
				if dt != enums.DataTypeNone && dt != enums.DataTypeOctetString {
					value, err = internal.ChangeTypeFromByteArray(settings, v, dt)
					if err != nil {
						return err
					}
				}
			}
			server.transaction.Targets[0].Value = value
			server.NotifyWrite(server.transaction.Targets)
			if !server.transaction.Targets[0].Handled && !p.multipleBlocks {
				server.transaction.Targets[0].Target.(objects.IGXDLMSBase).SetValue(settings, server.transaction.Targets[0])
				server.NotifyPostWrite(server.transaction.Targets)
				ret = byte(server.transaction.Targets[0].Error)
			}
		}
		settings.ResetBlockIndex()
	}
	if ret != 0 {
		p.attributeDescriptor = types.NewGXByteBuffer()
		err = p.attributeDescriptor.SetUint8(ret)
	}
	p.multipleBlocks = true
	return nil
}

func primeDcHanleSetRequestWithList(settings *settings.GXDLMSSettings,
	invokeID uint8,
	server *GXDLMSServer,
	data *types.GXByteBuffer,
	p *GXDLMSLNParameters,
	replyData *types.GXByteBuffer,
	xml *settings.GXDLMSTranslatorStructure) error {
	cnt, err := types.GetObjectCount(data)
	if err != nil {
		return err
	}
	status := make(map[int]uint8)
	if xml != nil {
		xml.AppendStartTag(int(internal.TranslatorTagsAttributeDescriptorList), "Qty", xml.IntegerToHex(cnt, 2, false), false)
	}
	for pos := 0; pos != cnt; pos++ {
		status[pos] = 0
		ret, er := data.Uint16()
		if er != nil {
			return er
		}
		ci := enums.ObjectType(ret)
		ln := make([]byte, 6)
		data.Get(ln)
		attributeIndex, err := data.Uint8()
		if err != nil {
			return err
		}
		// AccessSelection
		selection, err := data.Uint8()
		if err != nil {
			return err
		}
		selector := uint8(0)
		var parameters any
		if selection != 0 {
			selector, err = data.Uint8()
			if err != nil {
				return err
			}
			info := internal.GXDataInfo{}
			parameters, err = internal.GetData(settings, data, &info)
			if err != nil {
				return err
			}
		}
		ln2, err := helpers.ToLogicalName(ln)
		if err != nil {
			return err
		}
		xml.AppendComment(ln2)
		if xml != nil {
			xml.AppendStartTag(int(internal.TranslatorTagsAttributeDescriptorWithSelection), "", "", false)
			xml.AppendStartTag(int(internal.TranslatorTagsAttributeDescriptor), "", "", false)
			xml.AppendComment(ci.String())
			xml.AppendLineFromTag(int(internal.TranslatorTagsClassId), "Value", xml.IntegerToHex(int(ci), 4, false))
			xml.AppendComment(ln2)
			xml.AppendLineFromTag(int(internal.TranslatorTagsInstanceId), "Value", types.ToHex(ln, false))
			xml.AppendLineFromTag(int(internal.TranslatorTagsAttributeId), "Value", xml.IntegerToHex(attributeIndex, 2, false))
			xml.AppendEndTag(int(internal.TranslatorTagsAttributeDescriptor), true)
			xml.AppendEndTag(int(internal.TranslatorTagsAttributeDescriptorWithSelection), true)
		} else {
			var obj objects.IGXDLMSBase
			if ci == enums.ObjectTypeAssociationLogicalName && ln2 == "0.0.40.0.0.255" {
				obj = settings.AssignedAssociation().(objects.IGXDLMSBase)
			}
			if obj == nil {
				obj = getObjectCollection(settings.Objects).FindByLN(ci, ln2)
			}
			if obj == nil {
				obj = server.NotifyFindObject(ci, 0, ln2).(objects.IGXDLMSBase)
			}
			if obj == nil {
				status[pos] = uint8(enums.ErrorCodeUndefinedObject)
			} else {
				arg := internal.NewValueEventArgs2(server, obj, attributeIndex)
				arg.Selector = selector
				arg.Parameters = parameters
				arg.InvokeId = uint32(invokeID)
				if (server.NotifyGetAttributeAccess(arg) & int(enums.AccessModeWrite)) == 0 {
					status[pos] = uint8(enums.ErrorCodeReadWriteDenied)
				}
			}
		}
		cnt, err = types.GetObjectCount(data)
		if err != nil {
			return err
		}
		if xml != nil {
			xml.AppendEndTag(int(internal.TranslatorTagsAttributeDescriptorList), true)
			xml.AppendStartTag(int(internal.TranslatorTagsValueList), "Qty", xml.IntegerToHex(cnt, 2, false), false)
		}
		for pos := 0; pos != cnt; pos++ {
			if xml != nil || status[pos] == 0 {
				di := internal.GXDataInfo{}
				di.Xml = xml
				if xml != nil && xml.OutputType() == enums.TranslatorOutputTypeStandardXML {
					xml.AppendStartTag(int(enums.CommandWriteRequest)<<8|int(enums.SingleReadResponseData), "", "", false)
				}
				value, err := internal.GetData(settings, data, &di)
				if err != nil {
					return err
				}
				if !di.Complete {
					value = types.ToHexWithRange(data.Array(), false, data.Position(), data.Size()-data.Position())
				} else if v, ok := value.([]byte); ok {
					value = types.ToHex(v, false)
				}
				if xml != nil && xml.OutputType() == enums.TranslatorOutputTypeStandardXML {
					xml.AppendEndTag(int(enums.CommandWriteRequest)<<8|int(enums.SingleReadResponseData), false)
				}
			}
		}
		if xml != nil {
			xml.AppendEndTag(int(internal.TranslatorTagsValueList), true)
		}
	}
	p.status = 0xFF
	p.attributeDescriptor = types.NewGXByteBuffer()
	types.SetObjectCount(len(status), p.attributeDescriptor)
	for _, it := range status {
		err := p.attributeDescriptor.SetUint8(it)
		if err != nil {
			return err
		}
	}
	p.requestType = uint8(enums.SetResponseTypeWithList)
	return nil
}

// primeDcMethodRequestNextDataBlock returns the handle method request next data block command.
//
// Parameters:
//
// primeDcMethodRequestNextDataBlock returns the handle method request next data block command.
//
// Parameters:
//
//	data: Received data.
func primeDcMethodRequestNextDataBlock2(settings *settings.GXDLMSSettings,
	server *GXDLMSServer,
	data *types.GXByteBuffer,
	invokeID uint8,
	replyData *types.GXByteBuffer,
	xml *settings.GXDLMSTranslatorStructure,
	streaming bool,
	cipheredCommand enums.Command) error {
	bb := types.GXByteBuffer{}
	var err error
	if !streaming {
		var index uint32
		index, err := data.Uint32()
		if err != nil {
			return err
		}
		if xml != nil {
			xml.AppendLineFromTag(int(internal.TranslatorTagsBlockNumber), "", xml.IntegerToHex(index, 8, false))
			return nil
		}
		if index != settings.BlockIndex {
			log.Printf("handleGetRequest failed. Invalid block number %d/%d.\n", settings.BlockIndex, index)
			return getLNPdu(NewGXDLMSLNParameters(settings, 0, enums.CommandMethodResponse, 1, nil, &bb, byte(enums.ErrorCodeDataBlockNumberInvalid), cipheredCommand), replyData)
		}
	}
	settings.IncreaseBlockIndex()
	cmd := enums.Command(enums.CommandMethodResponse)
	if streaming {
		cmd = enums.CommandGeneralBlockTransfer
	}
	p := NewGXDLMSLNParameters(settings, uint32(invokeID), cmd, byte(enums.ActionResponseTypeWithBlock), nil, &bb, byte(enums.ErrorCodeOk), cipheredCommand)
	p.streaming = streaming
	p.gbtWindowSize = settings.GbtWindowSize()
	// If transaction is not in progress.
	if server.Transaction() == nil {
		p.status = uint8(enums.ErrorCodeNoLongGetOrReadInProgress)
		p.requestType = 1
		err = getLNPdu(p, replyData)
	} else {
		err = bb.SetByteBuffer(server.transaction.Data)
		if err != nil {
			return err
		}
		moreData := settings.Index != settings.Count
		if moreData {
			// If there is multiple blocks on the buffer.This might happen when Max PDU size is very small.
			if bb.Size() < int(settings.MaxPduSize()) {
				for _, arg := range server.transaction.Targets {
					var value any
					server.NotifyPreAction([]*internal.ValueEventArgs{arg})
					if arg.Handled {
						value = arg.Value
					} else {
						value, err = arg.Target.(objects.IGXDLMSBase).Invoke(settings, arg)
						if err != nil {
							return err
						}
					}
					// Set default action reply if not given.
					if value != nil && arg.Error == 0 {
						err = bb.SetUint8(1)
						if err != nil {
							return err
						}
						err = bb.SetUint8(0)
						if err != nil {
							return err
						}
						dt, err := internal.GetDLMSDataType(reflect.TypeOf(value))
						if err != nil {
							return err
						}
						err = internal.SetData(settings, &bb, dt, value)
					} else {
						p.requestType = 1
						p.status = uint8(arg.Error)
						err = bb.SetUint8(0)
						if err != nil {
							return err
						}
					}
				}
				moreData = settings.Index != settings.Count
			}
		}
		p.multipleBlocks = true
		err = getLNPdu(p, replyData)
		if moreData || bb.Size() != bb.Position() {
			server.transaction.Data = &bb
		} else {
			server.transaction = nil
			settings.ResetBlockIndex()
		}
	}
	return err
}

// GetRequestNextDataBlock returns the handle get request next data block command.
//
// Parameters:
//
//	data: Received data.
func primeDcGetRequestNextDataBlock(settings *settings.GXDLMSSettings,
	invokeID uint8,
	server *GXDLMSServer,
	data *types.GXByteBuffer,
	replyData *types.GXByteBuffer,
	xml *settings.GXDLMSTranslatorStructure,
	streaming bool,
	cipheredCommand enums.Command) error {
	var err error
	bb := types.GXByteBuffer{}
	if !streaming {
		var index uint32
		index, err = data.Uint32()
		if xml != nil {
			xml.AppendLineFromTag(int(internal.TranslatorTagsBlockNumber), "", xml.IntegerToHex(index, 8, false))
			return nil
		}
		if index != settings.BlockIndex {
			log.Printf("handleGetRequest failed. Invalid block number %d/%d.", settings.BlockIndex, index)
			return getLNPdu(NewGXDLMSLNParameters(settings, 0, enums.CommandGetResponse, 2, nil, &bb, byte(enums.ErrorCodeDataBlockNumberInvalid), cipheredCommand), replyData)
		}
	}
	settings.IncreaseBlockIndex()
	cmd := enums.Command(enums.CommandGetResponse)
	if streaming {
		cmd = enums.CommandGeneralBlockTransfer
	}
	p := NewGXDLMSLNParameters(settings, uint32(invokeID), cmd, 2, nil, &bb, byte(enums.ErrorCodeOk), cipheredCommand)
	p.streaming = settings.GbtWindowSize() != 1
	p.gbtWindowSize = settings.GbtWindowSize()
	if server.transaction == nil {
		p.status = uint8(enums.ErrorCodeNoLongGetOrReadInProgress)
	} else {
		err = bb.SetByteBuffer(server.transaction.Data)
		if err != nil {
			return err
		}
		moreData := settings.Index != settings.Count
		if moreData {
			// If there is multiple blocks on the buffer.This might happen when Max PDU size is very small.
			if bb.Size() < int(settings.MaxPduSize()) {
				for _, arg := range server.transaction.Targets {
					var value any
					server.NotifyRead([]*internal.ValueEventArgs{arg})
					if arg.Handled {
						value = arg.Value
					} else {
						value, err = arg.Target.(objects.IGXDLMSBase).GetValue(settings, arg)
						if err != nil {
							return err
						}
					}
					// Add data.
					if arg.ByteArray {
						err = bb.Set(value.([]byte))
						if err != nil {
							return err
						}
					} else {
						err = appendData(settings, arg.Target.(objects.IGXDLMSBase), arg.Index, &bb, value)
						if err != nil {
							return err
						}
					}
				}
				moreData = settings.Index != settings.Count
			}
		}
		p.multipleBlocks = true
		err = getLNPdu(p, replyData)
		if moreData || bb.Size()-bb.Position() != 0 {
			server.transaction.Data = &bb
		} else {
			server.transaction = nil
			settings.ResetBlockIndex()
		}
	}
	return err
}

// HandleEventNotification returns the handle Event Notification.
func primeDcHandleEventNotification(settings *settings.GXDLMSSettings,
	reply *GXReplyData,
	list []*types.GXKeyValuePair[objects.IGXDLMSBase, int]) error {
	reply.Time = time.Time{}
	// Check is there date-time.
	len, err := reply.Data.Uint8()
	if err != nil {
		return err
	}
	var tmp []byte
	// If date time is given.
	if len != 0 {
		len, err = reply.Data.Uint8()
		if err != nil {
			return err
		}
		tmp = make([]byte, len)
		err = reply.Data.Get(tmp)
		if err != nil {
			return err
		}
		ret, err := internal.ChangeTypeFromByteArray(settings, tmp, enums.DataTypeDateTime)
		if err != nil {
			return err
		}
		reply.Time = ret.(*types.GXDateTime).Value
	}
	if reply.xml != nil {
		reply.xml.AppendStartTag(int(enums.CommandEventNotification), "", "", false)
		if !reply.Time.IsZero() {
			reply.xml.AppendComment(fmt.Sprint(reply.Time))
			reply.xml.AppendLineFromTag(int(internal.TranslatorTagsTime), "", types.ToHex(tmp, false))
		}
	}
	ci, err := reply.Data.Uint16()
	if err != nil {
		return err
	}
	ln := make([]byte, 6)
	err = reply.Data.Get(ln)
	if err != nil {
		return err
	}
	index, err := reply.Data.Uint8()
	if err != nil {
		return err
	}
	if reply.xml != nil {
		primeDcAppendAttributeDescriptor(reply.xml, int(ci), ln, index)
		reply.xml.AppendStartTag(int(internal.TranslatorTagsAttributeValue), "", "", true)
	}
	di := internal.GXDataInfo{}
	di.Xml = reply.xml
	value, err := internal.GetData(settings, reply.Data, &di)
	if err != nil {
		return err
	}
	ln2, err := helpers.ToLogicalName(ln)
	if err != nil {
		return err
	}
	if reply.xml != nil {
		reply.xml.AppendEndTag(int(internal.TranslatorTagsAttributeValue), true)
		reply.xml.AppendEndTag(int(enums.CommandEventNotification), true)
	} else {
		var obj objects.IGXDLMSBase
		if enums.ObjectType(ci) == enums.ObjectTypeAssociationLogicalName && ln2 == "0.0.40.0.0.255" {
			obj = settings.AssignedAssociation().(objects.IGXDLMSBase)
		}
		if obj == nil {
			obj = getObjectCollection(settings.Objects).FindByLN(enums.ObjectType(ci), ln2)
		}
		if obj != nil {
			v := internal.NewValueEventArgs3(obj, index, 0, nil)
			v.Value = value
			err = obj.SetValue(settings, v)
			list = append(list, types.NewGXKeyValuePair[objects.IGXDLMSBase, int](obj, int(index)))
		}
	}
	return nil
}

func primeDcHandleNotification(data *types.GXByteBuffer, replyData *GXReplyData, xml *settings.GXDLMSTranslatorStructure) (bool, error) {
	var err error
	ret, err := data.Uint8()
	if err != nil {
		return false, err
	}
	type_ := enums.PrimeDcMsgType(ret)
	switch type_ {
	case enums.PrimeDcMsgTypeNewDeviceNotification:
		if xml != nil {
			xml.AppendStartTag(int(enums.TranslatorGeneralTagsPrimeNewDeviceNotification), "", "", false)
		}
		err = primeDcHandleNewDeviceNotification(data, replyData, xml)
		if xml != nil {
			xml.AppendEndTag(int(enums.TranslatorGeneralTagsPrimeNewDeviceNotification), false)
		}
	case enums.PrimeDcMsgTypeRemoveDeviceNotification:
		if xml != nil {
			xml.AppendStartTag(int(enums.TranslatorGeneralTagsPrimeRemoveDeviceNotification), "", "", false)
		}
		err = primeDcHandleRemoveDeviceNotification(data, replyData, xml)
		if xml != nil {
			xml.AppendEndTag(int(enums.TranslatorGeneralTagsPrimeRemoveDeviceNotification), false)
		}
	case enums.PrimeDcMsgTypeStartReportingMeters:
		if xml != nil {
			xml.AppendStartTag(int(enums.TranslatorGeneralTagsPrimeStartReportingMeters), "", "", false)
		}
		err = primeDcHandleStartReportingMeters(data, replyData, xml)
		if xml != nil {
			xml.AppendEndTag(int(enums.TranslatorGeneralTagsPrimeStartReportingMeters), false)
		}
	case enums.PrimeDcMsgTypeDeleteMeters:
		if xml != nil {
			xml.AppendStartTag(int(enums.TranslatorGeneralTagsPrimeDeleteMeters), "", "", false)
		}
		err = primeDcHandleDeleteMeters(data, replyData, xml)
		if xml != nil {
			xml.AppendEndTag(int(enums.TranslatorGeneralTagsPrimeDeleteMeters), false)
		}
	case enums.PrimeDcMsgTypeEnableAutoClose:
		if xml != nil {
			xml.AppendStartTag(int(enums.TranslatorGeneralTagsPrimeEnableAutoClose), "", "", false)
		}
		err = primeDcHandleEnableAutoClose(data, replyData, xml)
		if xml != nil {
			xml.AppendEndTag(int(enums.TranslatorGeneralTagsPrimeEnableAutoClose), false)
		}
	case enums.PrimeDcMsgTypeDisableAutoClose:
		if xml != nil {
			xml.AppendStartTag(int(enums.TranslatorGeneralTagsPrimeDisableAutoClose), "", "", false)
		}
		err = primeDcHandleDisableAutoClose(data, replyData, xml)
		if xml != nil {
			xml.AppendEndTag(int(enums.TranslatorGeneralTagsPrimeDisableAutoClose), false)
		}
	default:
		data.SetPosition(data.Position() - 1)
		return false, nil
	}
	return true, err
}

// HandleSetRequest returns the handle set request.
//
// Returns:
//
//	Reply to the client.
func primeDcHandleSetRequest(settings *settings.GXDLMSSettings,
	server *GXDLMSServer,
	data *types.GXByteBuffer,
	replyData *types.GXByteBuffer,
	xml *settings.GXDLMSTranslatorStructure,
	cipheredCommand enums.Command) error {
	// Return error if connection is not established.
	if xml == nil && (settings.Connected&enums.ConnectionStateDlms) == 0 && cipheredCommand == enums.CommandNone {
		return replyData.Set(GenerateConfirmedServiceError(enums.ConfirmedServiceErrorInitiateError, enums.ServiceErrorService, uint8(enums.ServiceUnsupported)))
	}
	// Get type_.
	type_, err := data.Uint8()
	if err != nil {
		return err
	}
	// Get invoke ID and priority.
	invoke, err := data.Uint8()
	if err != nil {
		return err
	}
	settings.UpdateInvokeID(invoke)
	p := NewGXDLMSLNParameters(settings, uint32(invoke), enums.CommandSetResponse, uint8(type_), nil, nil, 0xFF, cipheredCommand)
	// SetRequest normal or Set Request With First Data Block
	if xml != nil {
		if type_ < 6 {
			addInvokeId(xml, enums.CommandSetRequest, int(type_), uint32(invoke))
		} else {
			xml.AppendStartTag(int(enums.CommandSetRequest), "", "", true)
			xml.AppendComment(fmt.Sprintf("Unknown tag: %d", type_))
			xml.AppendLineFromTag(int(internal.TranslatorTagsInvokeId), "Value", xml.IntegerToHex(invoke, 2, false))
		}
	}
	switch type_ {
	case byte(enums.SetRequestTypeNormal):
	case byte(enums.SetRequestTypeFirstDataBlock):
		if type_ == byte(enums.SetRequestTypeNormal) {
			p.status = 0
		}
		err = primeDcHandleSetRequestNormal(settings, server, data, byte(type_), p, replyData, xml)
	case byte(enums.SetRequestTypeWithDataBlock):
		err = primeDcHandleSetRequestWithDataBlock(settings, server, data, p, replyData, xml)
	case byte(enums.SetRequestTypeWithList):
		err = primeDcHanleSetRequestWithList(settings, invoke, server, data, p, replyData, xml)
	default:
		log.Println("HandleSetRequest failed. Unknown command.")
		data.Clear()
		settings.ResetBlockIndex()
		p.status = byte(enums.ErrorCodeReadWriteDenied)
	}

	if xml != nil {
		if type_ < 6 {
			xml.AppendEndTag(int(enums.CommandSetRequest)<<8|int(type_), true)
		}
		xml.AppendEndTag(int(enums.CommandSetRequest), true)
		return nil
	}
	return getLNPdu(p, replyData)
}

func primeDcMethodRequest(settings *settings.GXDLMSSettings,
	type_ enums.ActionRequestType,
	invokeId uint8,
	server *GXDLMSServer,
	data *types.GXByteBuffer,
	connectionInfo *GXDLMSConnectionEventArgs,
	replyData *types.GXByteBuffer,
	xml *settings.GXDLMSTranslatorStructure,
	cipheredCommand enums.Command) error {
	error_ := enums.ErrorCodeOk
	bb := types.GXByteBuffer{}
	// CI
	ret, err := data.Uint16()
	if err != nil {
		return err
	}
	ci := enums.ObjectType(ret)
	ln := make([]byte, 6)
	err = data.Get(ln)
	if err != nil {
		return err
	}
	// Attribute Id
	id, err := data.Uint8()
	if err != nil {
		return err
	}
	var parameters any
	p := NewGXDLMSLNParameters(settings, uint32(invokeId), enums.CommandMethodResponse, byte(enums.ActionResponseTypeNormal), nil, &bb, 0, cipheredCommand)
	switch type_ {
	case enums.ActionRequestTypeNormal:
		// Get parameters.
		selection, err := data.Uint8()
		if err != nil {
			return err
		}
		if xml != nil {
			primeDcappendMethodDescriptor(xml, int(ci), ln, id)
			if selection != 0 {
				xml.AppendStartTag(int(internal.TranslatorTagsMethodInvocationParameters), "", "", false)
				di := internal.GXDataInfo{}
				di.Xml = xml
				internal.GetData(settings, data, &di)
				xml.AppendEndTag(int(internal.TranslatorTagsMethodInvocationParameters), true)
			}
			return nil
		}
		if selection != 0 {
			info := internal.GXDataInfo{}
			parameters, err = internal.GetData(settings, data, &info)
			if err != nil {
				return err
			}
		}
	case enums.ActionRequestTypeWithFirstBlock:
		p.requestType = uint8(enums.ActionResponseTypeNextBlock)
		p.status = 0xFF
		lastBlock, err := data.Uint8()
		if err != nil {
			return err
		}
		p.multipleBlocks = lastBlock == 0
		blockNumber, err := data.Uint32()
		if err != nil {
			return err
		}
		if xml == nil && blockNumber != settings.BlockIndex {
			log.Printf("MethodRequest failed. Invalid block number. %d/%d\n", settings.BlockIndex, blockNumber)
			p.status = uint8(enums.ErrorCodeDataBlockNumberInvalid)
			return nil
		}
		settings.IncreaseBlockIndex()
		size, err := types.GetObjectCount(data)
		if err != nil {
			return err
		}
		realSize := data.Size() - data.Position()
		if size != realSize {
			if xml == nil {
				log.Println("MethodRequest failed. Invalid block size.")
				p.status = uint8(enums.ErrorCodeDataBlockUnavailable)
				return nil
			}
			xml.AppendComment("Invalid block size.")
		}
		if xml != nil {
			primeDcappendMethodDescriptor(xml, int(ci), ln, id)
			xml.AppendStartTag(int(internal.TranslatorTagsDataBlock), "", "", true)
			xml.AppendLineFromTag(int(internal.TranslatorTagsLastBlock), "Value", xml.IntegerToHex(lastBlock, 2, false))
			xml.AppendLineFromTag(int(internal.TranslatorTagsBlockNumber), "Value", xml.IntegerToHex(blockNumber, 8, false))
			xml.AppendLineFromTag(int(internal.TranslatorTagsRawData), "Value", data.RemainingHexString(false))
			xml.AppendEndTag(int(internal.TranslatorTagsDataBlock), true)
			return nil
		}
	}
	ln2, err := helpers.ToLogicalName(ln)
	if err != nil {
		return err
	}
	var obj objects.IGXDLMSBase
	var e *internal.ValueEventArgs
	if ci == enums.ObjectTypeAssociationLogicalName && ln2 == "0.0.40.0.0.255" {
		obj = settings.AssignedAssociation().(objects.IGXDLMSBase)
	}
	if obj == nil {
		obj = getObjectCollection(settings.Objects).FindByLN(ci, ln2)
	}
	if (settings.Connected&enums.ConnectionStateDlms) == 0 && cipheredCommand == enums.CommandNone && (ci != enums.ObjectTypeAssociationLogicalName || id != 1) {
		return replyData.Set(GenerateConfirmedServiceError(enums.ConfirmedServiceErrorInitiateError, enums.ServiceErrorService, uint8(enums.ServiceUnsupported)))
	}
	if obj == nil {
		obj = server.NotifyFindObject(ci, 0, ln2).(objects.IGXDLMSBase)
	}
	if obj == nil {
		error_ = enums.ErrorCodeUndefinedObject
	} else {
		if settings.AssignedAssociation() != nil {
			p.AccessMode = int(settings.AssignedAssociation().(*objects.GXDLMSAssociationLogicalName).GetObjectMethodAccess3(obj, int(id)))
		}
		e = internal.NewValueEventArgs2(server, obj, id)
		e.Parameters = parameters
		e.InvokeId = uint32(invokeId)
		if (server.NotifyGetMethodAccess(e) & int(enums.MethodAccessModeAccess)) == 0 {
			error_ = enums.ErrorCodeReadWriteDenied
		} else {
			if p.multipleBlocks {
				server.transaction = newGXDLMSLongTransaction([]*internal.ValueEventArgs{e}, enums.CommandMethodRequest, data)
			} else if server.transaction == nil {
				//Check transaction so invoke is not called multiple times.This might happen when all data can't fit to one PDU.
				p.requestType = uint8(enums.ActionResponseTypeNormal)
				server.NotifyPreAction([]*internal.ValueEventArgs{e})
				var actionReply []byte
				if e.Handled {
					actionReply = e.Value.([]byte)
				} else {
					actionReply, err = obj.Invoke(settings, e)
					if err != nil {
						return err
					}
					server.NotifyPostAction([]*internal.ValueEventArgs{e})
				}
				// Set default action reply if not given.
				if actionReply != nil && e.Error == 0 {
					err = bb.SetUint8(1)
					if err != nil {
						return err
					}
					err = bb.SetUint8(0)
					if err != nil {
						return err
					}
					dt, err := internal.GetDLMSDataType(reflect.TypeOf(actionReply))
					if err != nil {
						return err
					}
					err = internal.SetData(settings, &bb, dt, actionReply)
				} else {
					error_ = int(e.Error)
					err = bb.SetUint8(0)
					if err != nil {
						return err
					}
				}
			}
		}
		p.invokeId = uint32(e.InvokeId)
	}
	if error_ != 0 {
		p.status = uint8(error_)
	}
	err = getLNPdu(p, replyData)
	// If all reply data doesn't fit to one PDU.
	if server.transaction == nil && p.data.Available() != 0 {
		server.transaction = newGXDLMSLongTransaction([]*internal.ValueEventArgs{e}, enums.CommandMethodResponse, p.data)
	}
	// If High level authentication fails.
	if error_ == 0 && (settings.AssignedAssociation() != nil) && ci == enums.ObjectTypeAssociationLogicalName && id == 1 {
		if _, ok := obj.(*objects.GXDLMSAssociationLogicalName); ok {
			server.NotifyConnected(connectionInfo)
			settings.Connected |= enums.ConnectionStateDlms
		} else {
			server.NotifyInvalidConnection(connectionInfo)
			settings.Connected &= ^enums.ConnectionStateDlms
		}
	}
	// Start to use new keys.
	if ss, ok := obj.(*objects.GXDLMSSecuritySetup); ok {
		err = ss.ApplyKeys(settings, e)
	}
	return err
}

func primeDcMethodRequestNextBlock(settings *settings.GXDLMSSettings,
	server *GXDLMSServer,
	data *types.GXByteBuffer,
	connectionInfo *GXDLMSConnectionEventArgs,
	replyData *types.GXByteBuffer,
	xml *settings.GXDLMSTranslatorStructure,
	streaming bool,
	cipheredCommand enums.Command) error {
	var e *internal.ValueEventArgs
	bb := types.GXByteBuffer{}
	lastBlock, err := data.Uint8()
	if err != nil {
		return err
	}
	if !streaming {
		var blockNumber uint32
		blockNumber, err = data.Uint32()
		if err != nil {
			return err
		}
		// Get data size.
		size, err := types.GetObjectCount(data)
		if err != nil {
			return err
		}
		if xml != nil {
			xml.AppendStartTag(int(internal.TranslatorTagsDataBlock), "", "", true)
			xml.AppendLineFromTag(int(internal.TranslatorTagsLastBlock), "Value", xml.IntegerToHex(lastBlock, 2, false))
			xml.AppendLineFromTag(int(internal.TranslatorTagsBlockNumber), "Value", xml.IntegerToHex(blockNumber, 8, false))
			xml.AppendLineFromTag(int(internal.TranslatorTagsRawData), "Value", data.RemainingHexString(false))
			xml.AppendEndTag(int(internal.TranslatorTagsDataBlock), true)
			return nil
		}
		if blockNumber != settings.BlockIndex {
			server.transaction = nil
			log.Println("MethodRequestNextBlock failed. Invalid block number. %d/%d", settings.BlockIndex, blockNumber)
			settings.ResetBlockIndex()
			return getLNPdu(NewGXDLMSLNParameters(settings, 0, enums.CommandMethodResponse, byte(enums.ActionResponseTypeNormal), nil,
				&bb, byte(enums.ErrorCodeDataBlockNumberInvalid), cipheredCommand), replyData)
		}
		if size < data.Available() {
			server.transaction = nil
			settings.ResetBlockIndex()
			log.Println("MethodRequestNextBlock failed. Not enought data. Actual: %d . Expected %d", data.Available, +size)
			return getLNPdu(NewGXDLMSLNParameters(settings, 0, enums.CommandMethodResponse, byte(enums.ActionResponseTypeNormal), nil,
				&bb, byte(enums.ErrorCodeDataBlockNumberInvalid), cipheredCommand), replyData)
		}
	}
	cmd := enums.Command(enums.CommandMethodResponse)
	if streaming {
		cmd = enums.CommandGeneralBlockTransfer
	}
	p := NewGXDLMSLNParameters(settings, 0, cmd, byte(enums.ActionResponseTypeNormal), nil, &bb, byte(enums.ErrorCodeOk), cipheredCommand)
	p.multipleBlocks = lastBlock == 0
	p.streaming = streaming
	p.gbtWindowSize = settings.GbtWindowSize()
	// If transaction is not in progress.
	if server.transaction == nil {
		p.status = uint8(enums.ErrorCodeNoLongGetOrReadInProgress)
	} else {
		server.transaction.Data.SetByteBuffer(data)
		if lastBlock == 1 {
			info := internal.GXDataInfo{}
			parameters, err := internal.GetData(settings, server.transaction.Data, &info)
			if err != nil {
				return err
			}
			for _, arg := range server.transaction.Targets {
				var value any
				arg.Parameters = parameters
				args := []*internal.ValueEventArgs{arg}
				server.NotifyPreAction(args)
				if arg.Handled {
					value = arg.Value
				} else {
					value, err = arg.Target.(objects.IGXDLMSBase).Invoke(settings, arg)
					if err != nil {
						return err
					}
				}
				server.NotifyPostAction(args)
				// Set default action reply if not given.
				if value != nil && arg.Error == 0 {
					// If High level authentication fails.
					if a, ok := arg.Target.(*objects.GXDLMSAssociationLogicalName); ok {
						if a.AssociationStatus == enums.AssociationStatusAssociated {
							server.NotifyConnected(connectionInfo)
							settings.Connected |= enums.ConnectionStateDlms
						} else {
							server.NotifyInvalidConnection(connectionInfo)
							settings.Connected &= ^enums.ConnectionStateDlms
						}
					}
					// Start to use new keys.
					if ss, ok := arg.Target.(objects.GXDLMSSecuritySetup); ok {
						ss.ApplyKeys(settings, e)
					}
					err = bb.SetUint8(1)
					if err != nil {
						return err
					}
					err = bb.SetUint8(0)
					if err != nil {
						return err
					}
					dt, err := internal.GetDLMSDataType(reflect.TypeOf(value))
					if err != nil {
						return err
					}
					err = internal.SetData(settings, &bb, dt, value)
					if err != nil {
						return err
					}
				} else {
					p.status = uint8(arg.Error)
					err = bb.SetUint8(0)
					if err != nil {
						return err
					}
				}
				server.transaction = nil
				settings.ResetBlockIndex()
				p.blockIndex = 1
			}
		} else {
			//Ask next block.
			p.requestType = uint8(enums.ActionResponseTypeNextBlock)
			p.status = 0xFF
		}
	}
	err = getLNPdu(p, replyData)
	if settings.Count != settings.Index || bb.Size() != bb.Position() {
		server.transaction = newGXDLMSLongTransaction([]*internal.ValueEventArgs{e}, enums.CommandMethodRequest, &bb)
	}
	if lastBlock == 0 {
		settings.IncreaseBlockIndex()
	}
	return err
}

// HandleMethodRequest returns the handle action request.
func primeDcHandleMethodRequest(settings *settings.GXDLMSSettings,
	server *GXDLMSServer,
	data *types.GXByteBuffer,
	connectionInfo *GXDLMSConnectionEventArgs,
	replyData *types.GXByteBuffer,
	xml *settings.GXDLMSTranslatorStructure,
	cipheredCommand enums.Command) error {
	// Get type.
	var invokeID uint8
	ret, err := data.Uint8()
	if err != nil {
		return err
	}
	type_ := enums.ActionRequestType(ret)
	invokeID, err = data.Uint8()
	if err != nil {
		return err
	}
	settings.UpdateInvokeID(invokeID)
	if xml != nil {
		if type_ > 0 && type_ <= enums.ActionRequestTypeWithBlock {
			addInvokeId(xml, enums.CommandMethodRequest, int(type_), uint32(invokeID))
		} else {
			xml.AppendStartTag(int(enums.CommandMethodRequest), "", "", true)
			xml.AppendComment(fmt.Sprintf("Unknown tag: %d.", type_))
			xml.AppendLineFromTag(int(internal.TranslatorTagsInvokeId), "Value", xml.IntegerToHex(invokeID, 2, false))
		}
	}
	switch type_ {
	case enums.ActionRequestTypeNormal:
	case enums.ActionRequestTypeWithFirstBlock:
		err = primeDcMethodRequest(settings, type_, invokeID, server, data, connectionInfo, replyData, xml, cipheredCommand)
	case enums.ActionRequestTypeNextBlock:
		err = primeDcMethodRequestNextDataBlock2(settings, server, data, invokeID, replyData, xml, false, cipheredCommand)
	case enums.ActionRequestTypeWithBlock:
		err = primeDcMethodRequestNextBlock(settings, server, data, connectionInfo, replyData, xml, false, cipheredCommand)
	default:
		if xml == nil {
			log.Println("HandleMethodRequest failed. Invalid command type_.")
			settings.ResetBlockIndex()
			type_ = enums.ActionRequestTypeNormal
			data.Clear()
			err = getLNPdu(NewGXDLMSLNParameters(settings, uint32(invokeID), enums.CommandMethodResponse, byte(type_), nil, nil, byte(enums.ErrorCodeReadWriteDenied), cipheredCommand), replyData)
		}
	}
	if xml != nil {
		if type_ > 0 && type_ <= enums.ActionRequestTypeWithBlock {
			xml.AppendEndTag(int(enums.CommandMethodRequest)<<8|int(type_), true)
		} else {
			xml.AppendEndTag(int(enums.CommandMethodRequest), true)
		}
		xml.AppendEndTag(int(enums.CommandMethodRequest), true)
	}
	return err
}

// HandleAccessRequest returns the handle Access request.
func primeDcHandleAccessRequest(settings *settings.GXDLMSSettings,
	server *GXDLMSServer,
	data *types.GXByteBuffer,
	reply *types.GXByteBuffer,
	xml *settings.GXDLMSTranslatorStructure,
	cipheredCommand enums.Command) error {
	var err error
	// Return error if connection is not established.
	if xml == nil && (settings.Connected&enums.ConnectionStateDlms) == 0 && cipheredCommand == enums.CommandNone {
		return reply.Set(GenerateConfirmedServiceError(enums.ConfirmedServiceErrorInitiateError, enums.ServiceErrorService, uint8(enums.ServiceUnsupported)))
	}
	// Get long invoke id and priority.
	invokeId, err := data.Uint32()
	if err != nil {
		return err
	}
	settings.LongInvokeID = invokeId
	len, err := types.GetObjectCount(data)
	if err != nil {
		return err
	}
	var tmp []byte
	// If date time is given.
	if len != 0 {
		tmp = make([]byte, len)
		err = data.Get(tmp)
		if err != nil {
			return err
		}
		if xml == nil {
			var dt enums.DataType
			dt = enums.DataTypeDateTime
			if len == 4 {
				dt = enums.DataTypeTime
			} else if len == 5 {
				dt = enums.DataTypeDate
			}
			info := internal.GXDataInfo{}
			info.Type = dt
			ret := types.NewGXByteBufferWithData(tmp)
			_, err = internal.GetData(settings, ret, &info)
			if err != nil {
				return err
			}
		}
	}
	// Get object count.
	cnt, err := types.GetObjectCount(data)
	if err != nil {
		return err
	}
	if xml != nil {
		xml.AppendStartTag(int(enums.CommandAccessRequest), "", "", false)
		xml.AppendLineFromTag(int(internal.TranslatorTagsLongInvokeId), "Value", xml.IntegerToHex(invokeId, 2, false))
		xml.AppendLineFromTag(int(internal.TranslatorTagsDateTime), "Value", types.ToHex(tmp, false))
		xml.AppendStartTag(int(internal.TranslatorTagsAccessRequestBody), "", "", false)
		xml.AppendStartTag(int(internal.TranslatorTagsListOfAccessRequestSpecification), "Qty", xml.IntegerToHex(cnt, 2, false), false)
	}
	list := []*GXDLMSAccessItem{}
	for pos := 0; pos != cnt; pos++ {
		ch, err := data.Uint8()
		if err != nil {
			return err
		}
		type_ := enums.AccessServiceCommandType(ch)
		if !(type_ == enums.AccessServiceCommandTypeGet || type_ == enums.AccessServiceCommandTypeSet || type_ == enums.AccessServiceCommandTypeAction) {
			return errors.New("Invalid access service command type_.")
		}
		// CI
		ret, err := data.Uint16()
		if err != nil {
			return err
		}
		ci := enums.ObjectType(ret)
		ln := make([]byte, 6)
		data.Get(ln)
		ln2, err := helpers.ToLogicalName(ln)
		if err != nil {
			return err
		}
		xml.AppendComment(ln2)
		// Attribute Id
		attributeIndex, err := data.Uint8()
		if err != nil {
			return err
		}
		if xml != nil {
			xml.AppendStartTag(int(internal.TranslatorTagsAccessRequestSpecification), "", "", false)
			xml.AppendStartTag(int(enums.CommandAccessRequest)<<8|int(type_), "", "", true)
			primeDcAppendAttributeDescriptor(xml, int(ci), ln, attributeIndex)
			xml.AppendEndTag(int(enums.CommandAccessRequest)<<8|int(type_), true)
			xml.AppendEndTag(int(internal.TranslatorTagsAccessRequestSpecification), true)
		} else {
			var obj objects.IGXDLMSBase
			if ci == enums.ObjectTypeAssociationLogicalName && ln2 == "0.0.40.0.0.255" {
				obj = settings.AssignedAssociation().(objects.IGXDLMSBase)
			}
			if obj == nil {
				obj = getObjectCollection(settings.Objects).FindByLN(ci, ln2)
			}
			list = append(list, NewGXDLMSAccessItem(type_, obj, attributeIndex))
		}
	}
	if xml != nil {
		xml.AppendEndTag(int(internal.TranslatorTagsListOfAccessRequestSpecification), true)
		xml.AppendStartTag(int(internal.TranslatorTagsAccessRequestListOfData), "Qty", xml.IntegerToHex(cnt, 2, false), true)
	}
	cnt, err = types.GetObjectCount(data)
	if err != nil {
		return err
	}
	bb := types.NewGXByteBuffer()
	err = bb.SetUint8(0)
	if err != nil {
		return err
	}
	types.SetObjectCount(cnt, bb)
	results := types.NewGXByteBuffer()
	types.SetObjectCount(cnt, results)
	for pos := 0; pos != cnt; pos++ {
		di := internal.GXDataInfo{}
		di.Xml = xml
		if xml != nil && xml.OutputType() == enums.TranslatorOutputTypeStandardXML {
			xml.AppendStartTag(int(enums.CommandWriteRequest)<<8|int(enums.SingleReadResponseData), "", "", true)
		}
		value, err := internal.GetData(settings, data, &di)
		if err != nil {
			return err
		}
		if !di.Complete {
			value = types.ToHexWithRange(data.Array(), false, data.Position(), data.Size()-data.Position())
		} else if v, ok := value.([]byte); ok {
			value = types.ToHex(v, false)
		}
		var e *internal.ValueEventArgs
		if xml == nil {
			it := list[pos]
			err = results.SetUint8(byte(it.Command))
			if err != nil {
				return err
			}
			if it.Target == nil {
				err = bb.SetUint8(0)
				if err != nil {
					return err
				}
				err = results.SetUint8(enums.ErrorCodeUnavailableObject)
				if err != nil {
					return err
				}
			} else {
				e = internal.NewValueEventArgs(settings, it.Target, it.Index)
				if it.Command == enums.AccessServiceCommandTypeGet {
					access := server.NotifyGetAttributeAccess(e)
					if (access&int(enums.AccessModeRead)) == 0 && (access&int(enums.AccessModeAuthenticatedRead)) == 0 {
						err = bb.SetUint8(0)
						if err != nil {
							return err
						}
						err = results.SetUint8(enums.ErrorCodeReadWriteDenied)
						if err != nil {
							return err
						}
					} else {
						server.NotifyRead([]*internal.ValueEventArgs{e})
						if e.Handled {
							value = e.Value
						} else {
							value, err = it.Target.(objects.IGXDLMSBase).GetValue(settings, e)
							if err != nil {
								return err
							}
						}
						// If all data is not fit to PDU and GBT is not used.
						if settings.Index != settings.Count {
							settings.Count = 0
							settings.Index = 0

							err = bb.SetUint8(0)
							if err != nil {
								return err
							}
							err = results.SetUint8(enums.ErrorCodeReadWriteDenied)
							if err != nil {
								return err
							}
						} else {
							if e.ByteArray {
								err = bb.Set(value.([]byte))
							} else {
								err = appendData(settings, it.Target, it.Index, bb, value)
							}
							if err != nil {
								return err
							}
							server.NotifyPostRead([]*internal.ValueEventArgs{e})
							err = results.SetUint8(enums.ErrorCodeOk)
							if err != nil {
								return err
							}
						}
					}
				} else if it.Command == enums.AccessServiceCommandTypeSet {
					err = results.SetUint8(enums.ErrorCodeOk)
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
		}
		if xml != nil && xml.OutputType() == enums.TranslatorOutputTypeStandardXML {
			xml.AppendEndTag(int(enums.CommandWriteRequest)<<8|int(enums.SingleReadResponseData), true)
		}
	}
	if xml != nil {
		xml.AppendEndTag(int(internal.TranslatorTagsAccessRequestListOfData), true)
		xml.AppendEndTag(int(internal.TranslatorTagsAccessRequestBody), true)
		xml.AppendEndTag(int(enums.CommandAccessRequest), true)
	} else {
		err := bb.SetByteBuffer(results)
		if err != nil {
			return err
		}
		err = getLNPdu(NewGXDLMSLNParameters(settings, invokeId, enums.CommandAccessResponse, 0xff, nil, bb, 0xFF, cipheredCommand), reply)
	}
	return err
}
