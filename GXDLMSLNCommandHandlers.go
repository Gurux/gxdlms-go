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
	"github.com/Gurux/gxdlms-go/internal/constants"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/objects"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

func appendAttributeDescriptor(xml *settings.GXDLMSTranslatorStructure, ci int, ln []byte, attributeIndex uint8) error {
	xml.AppendStartTag(int(internal.TranslatorTagsAttributeDescriptor), "", "", true)
	if xml.Comments {
		xml.AppendComment(enums.ObjectType(ci).String())
	}
	xml.AppendLine(internal.TranslatorTagsClassId.String(), "Value", xml.IntegerToHex(int(ci), 4, false))
	ret, err := helpers.ToLogicalName(ln)
	if err != nil {
		return fmt.Errorf("Invalid logical name. %s", err.Error())
	}
	xml.AppendComment(ret)
	xml.AppendLine(internal.TranslatorTagsInstanceId.String(), "Value", types.ToHex(ln, false))
	obj, err := objects.CreateObject(enums.ObjectType(ci), ret, 0)
	if err != nil {
		return err
	}
	if obj != nil {
		xml.AppendComment(obj.GetNames()[attributeIndex-1])
	}
	xml.AppendLine(internal.TranslatorTagsAttributeId.String(), "Value", xml.IntegerToHex(attributeIndex, 2, false))
	xml.AppendEndTag(int(internal.TranslatorTagsAttributeDescriptor), true)
	return nil
}

func AppendMethodDescriptor(xml *settings.GXDLMSTranslatorStructure, ci int, ln []byte, attributeIndex uint8) error {
	xml.AppendStartTag(int(internal.TranslatorTagsMethodDescriptor), "", "", true)
	if xml.Comments {
		xml.AppendComment(enums.ObjectType(ci).String())
	}
	xml.AppendLine(internal.TranslatorTagsClassId.String(), "Value", xml.IntegerToHex(int(ci), 4, false))
	ret, err := helpers.ToLogicalName(ln)
	if err != nil {
		return fmt.Errorf("Invalid logical name. %s", err.Error())
	}
	xml.AppendComment(ret)
	xml.AppendLine(internal.TranslatorTagsInstanceId.String(), "Value", types.ToHex(ln, false))
	obj, err := objects.CreateObject(enums.ObjectType(ci), ret, 0)
	if err != nil {
		return err
	}
	if obj != nil {
		xml.AppendComment(obj.GetMethodNames()[attributeIndex-1])
	}
	xml.AppendLine(internal.TranslatorTagsMethodId.String(), "Value", xml.IntegerToHex(attributeIndex, 2, false))
	xml.AppendEndTag(int(internal.TranslatorTagsMethodDescriptor), true)
	return nil
}

// getRequestNormal returns the handle get request normal enums.Command
//
// Parameters:
//
//	data: Received data.
func getRequestNormal(settings *settings.GXDLMSSettings,
	invokeID uint8,
	server *GXDLMSServer,
	data *types.GXByteBuffer,
	replyData *types.GXByteBuffer,
	xml *settings.GXDLMSTranslatorStructure,
	cipheredCommand enums.Command) error {
	var mode int
	bb := types.NewGXByteBuffer()
	// Get type.
	status := enums.ErrorCodeOk
	settings.Count = 0
	settings.Index = 0
	settings.ResetBlockIndex()
	ret, err := data.Uint16()
	// CI
	ci := enums.ObjectType(ret)
	if data.Available() < 6 {
		if xml != nil {
			xml.AppendComment("Logical name is missing.")
			xml.AppendComment("Attribute Id is missing.")
			xml.AppendComment("Access Selection is missing.")
			return nil
		}
		return errors.New("Get request is not complete.")
	}
	ln := make([]byte, 6)
	data.Get(ln)
	// Attribute Id
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
	info := internal.GXDataInfo{}
	if selection != 0 {
		selector, err = data.Uint8()
		if err != nil {
			return err
		}
	}
	if xml != nil {
		appendAttributeDescriptor(xml, int(ci), ln, attributeIndex)
		if selection != 0 {
			info.Xml = xml
			xml.AppendStartTag(int(internal.TranslatorTagsAccessSelection), "", "", true)
			xml.AppendLine(internal.TranslatorTagsAccessSelector.String(), "Value", xml.IntegerToHex(selector, 2, false))
			xml.AppendStartTag(int(internal.TranslatorTagsAccessParameters), "", "", true)
			internal.GetData(settings, data, &info)
			xml.AppendEndTag(int(internal.TranslatorTagsAccessParameters), false)
			xml.AppendEndTag(int(internal.TranslatorTagsAccessSelection), false)
		}
		return nil
	}
	if selection != 0 {
		parameters, err = internal.GetData(settings, data, &info)
	}
	ln2, err := helpers.ToLogicalName(ln)
	if err != nil {
		return err
	}
	var obj *objects.GXDLMSAssociationLogicalName
	if ci == enums.ObjectTypeAssociationLogicalName && ln2 == "0.0.40.0.0.255" {
		obj = settings.AssignedAssociation().(*objects.GXDLMSAssociationLogicalName)
	}
	if obj == nil {
		obj = getObjectCollection(settings.Objects).FindByLN(ci, ln2).(*objects.GXDLMSAssociationLogicalName)
	}
	if obj == nil {
		obj = server.NotifyFindObject(ci, 0, ln2).(*objects.GXDLMSAssociationLogicalName)
	}
	e := internal.NewValueEventArgs2(server, obj, attributeIndex)
	e.Selector = selector
	e.Parameters = parameters
	e.InvokeId = uint32(invokeID)
	if obj == nil {
		status = enums.ErrorCodeUndefinedObject
	} else {
		access := server.NotifyGetAttributeAccess(e)
		if (access&int(enums.AccessModeRead)) == 0 && (access&int(enums.AccessModeAuthenticatedRead)) == 0 {
			status = enums.ErrorCodeReadWriteDenied
		} else {
			if settings.AssignedAssociation() != nil {
				mode = int(settings.AssignedAssociation().(*objects.GXDLMSAssociationLogicalName).GetObjectAccess3(obj, int(attributeIndex)))
			}
			if (obj.Base().ObjectType() == enums.ObjectTypeAssociationLogicalName || obj.Base().ObjectType() == enums.ObjectTypeAssociationShortName) && attributeIndex == 1 {
				val := []byte{0, 0, 40, 0, 0, 255}
				err = appendData(settings, obj, attributeIndex, bb, val)
				if err != nil {
					return err
				}
			} else {
				if (obj.Base().ObjectType() == enums.ObjectTypeProfileGeneric) && attributeIndex == 2 {
					e.RowToPdu, err = RowsToPdu(settings, e.Target.(*objects.GXDLMSProfileGeneric))
					if err != nil {
						return err
					}
				}
				var value any
				server.NotifyRead([]*internal.ValueEventArgs{e})
				if e.Handled {
					value = e.Value
				} else {
					settings.Count = e.RowEndIndex - e.RowBeginIndex
					value, err = obj.GetValue(settings, e)
					if err != nil {
						return err
					}
				}
				if e.ByteArray {
					err = bb.Set(value.([]byte))
					if err != nil {
						return err
					}
				} else {
					err = appendData(settings, obj, attributeIndex, bb, value)
					if err != nil {
						return err
					}
				}
				server.NotifyPostRead([]*internal.ValueEventArgs{e})
				status = int(e.Error)
			}
		}
	}
	p := NewGXDLMSLNParameters(settings, e.InvokeId, enums.CommandGetResponse, 1, nil, bb, byte(status), cipheredCommand)
	p.AccessMode = int(mode)
	err = getLNPdu(p, replyData)
	if settings.Count != settings.Index || bb.Size() != bb.Position() {
		server.SetTransaction(newGXDLMSLongTransaction([]*internal.ValueEventArgs{e}, enums.CommandGetRequest, bb))
	}
	return err
}

// GetRequestWithList returns the handle get request with list enums.Command
//
// Parameters:
//
//	data: Received data.
func GetRequestWithList(settings *settings.GXDLMSSettings,
	invokeID uint8,
	server *GXDLMSServer,
	data *types.GXByteBuffer,
	replyData *types.GXByteBuffer,
	xml *settings.GXDLMSTranslatorStructure,
	cipheredCommand enums.Command) error {
	bb := types.NewGXByteBuffer()
	var pos int
	cnt, err := types.GetObjectCount(data)
	if err != nil {
		return err
	}
	types.SetObjectCount(cnt, bb)
	list := []*internal.ValueEventArgs{}
	if xml != nil {
		xml.AppendStartTag(int(internal.TranslatorTagsAttributeDescriptorList), "Qty", xml.IntegerToHex(cnt, 2, false), true)
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
		var selector uint8
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
			xml.AppendStartTag(int(internal.TranslatorTagsAttributeDescriptorWithSelection), "", "", true)
			xml.AppendStartTag(int(internal.TranslatorTagsAttributeDescriptor), "", "", true)
			xml.AppendComment(ci.String())
			xml.AppendLineFromTag(int(internal.TranslatorTagsClassId), "Value", xml.IntegerToHex(int(ci), 4, false))
			ln2, err := helpers.ToLogicalName(ln)
			if err != nil {
				return err
			}
			xml.AppendComment(ln2)
			xml.AppendLineFromTag(int(internal.TranslatorTagsInstanceId), "Value", types.ToHex(ln, false))
			xml.AppendLineFromTag(int(internal.TranslatorTagsAttributeId), "Value", xml.IntegerToHex(attributeIndex, 2, false))
			xml.AppendEndTag(int(internal.TranslatorTagsAttributeDescriptor), false)
			xml.AppendEndTag(int(internal.TranslatorTagsAttributeDescriptorWithSelection), false)
		} else {
			var ln2 string
			var obj objects.IGXDLMSBase
			ln2, err = helpers.ToLogicalName(ln)
			if err != nil {
				return err
			}
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
				e := internal.NewValueEventArgs2(server, obj, attributeIndex)
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
		xml.AppendEndTag(int(internal.TranslatorTagsAttributeDescriptorList), true)
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
		err := bb.SetUint8(uint8(it.Error))
		if err != nil {
			return err
		}
		if it.ByteArray {
			err = bb.Set(value.([]byte))
			if err != nil {
				return err
			}
		} else {
			err = appendData(settings, it.Target.(objects.IGXDLMSBase), it.Index, bb, value)
			if err != nil {
				return err
			}
		}
		invokeID = uint8(it.InvokeId)
		pos++
	}
	server.NotifyPostRead(list)
	p := NewGXDLMSLNParameters(settings, uint32(invokeID), enums.CommandGetResponse, 3, nil, bb, 0xFF, cipheredCommand)
	err = getLNPdu(p, replyData)
	if settings.Index != settings.Count || bb.Available() != 0 {
		server.SetTransaction(newGXDLMSLongTransaction(list, enums.CommandGetRequest, bb))
	}
	return err
}

func handleSetRequestNormal(settings *settings.GXDLMSSettings,
	server *GXDLMSServer,
	data *types.GXByteBuffer,
	type_ uint8,
	p *GXDLMSLNParameters,
	replyData *types.GXByteBuffer,
	xml *settings.GXDLMSTranslatorStructure) error {
	var value any
	reply := internal.GXDataInfo{}
	// CI
	ret, err := data.Uint16()
	if err != nil {
		return err
	}
	ci := enums.ObjectType(ret)
	if data.Available() < 8 {
		if xml != nil {
			xml.AppendComment("Logical name is missin")
			xml.AppendComment("Attribute Id is missin")
			xml.AppendComment("Access Selection is missin")
			return nil
		}
		return errors.New("Set request is not complete.")
	}
	ln := make([]byte, 6)
	data.Get(ln)
	// Attribute index.
	index, err := data.Uint8()
	if err != nil {
		return err
	}
	data.Uint8()
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
			log.Printf("HandleSetRequest failed. Invalid block number. %d/%d", settings.BlockIndex, blockNumber)
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
			log.Println("HandleSetRequest failed. Invalid block size.")
			p.status = uint8(enums.ErrorCodeDataBlockUnavailable)
			return nil
		}
		if xml != nil {
			appendAttributeDescriptor(xml, int(ci), ln, index)
			xml.AppendStartTag(int(internal.TranslatorTagsDataBlock), "", "", true)
			xml.AppendLineFromTag(int(internal.TranslatorTagsLastBlock), "Value", xml.IntegerToHex(lastBlock, 2, false))
			xml.AppendLineFromTag(int(internal.TranslatorTagsBlockNumber), "Value", xml.IntegerToHex(blockNumber, 8, false))
			xml.AppendLineFromTag(int(internal.TranslatorTagsRawData), "Value", data.RemainingHexString(false))
			xml.AppendEndTag(int(internal.TranslatorTagsDataBlock), true)
		}
	}
	if xml != nil {
		appendAttributeDescriptor(xml, int(ci), ln, index)
		xml.AppendStartTag(int(internal.TranslatorTagsValue), "", "", true)
		di := internal.GXDataInfo{}
		di.Xml = xml
		value, err = internal.GetData(settings, data, &di)
		if err != nil {
			return err
		}
		if !di.Complete {
			types.ToHexWithRange(data.Array(), false, data.Position(), data.Size()-data.Position())
		} else if _, ok := value.([]byte); ok {
			types.ToHex(value.([]byte), false)
		}
		xml.AppendEndTag(int(internal.TranslatorTagsValue), true)
		return nil
	}
	if !p.multipleBlocks {
		settings.ResetBlockIndex()
		value, err = internal.GetData(settings, data, &reply)
		if err != nil {
			return err
		}

	}
	var obj objects.IGXDLMSBase
	ln2, err := helpers.ToLogicalName(ln)
	if err != nil {
		return err
	}
	if ci == enums.ObjectTypeAssociationLogicalName && ln2 == "0.0.40.0.0.255" {
		obj = settings.AssignedAssociation().(objects.IGXDLMSBase)
	}
	if obj == nil {
		ret, err := helpers.ToLogicalName(ln)
		if err != nil {
			return err
		}
		obj = getObjectCollection(settings.Objects).FindByLN(ci, ret)
	}
	if obj == nil {
		ret, err := helpers.ToLogicalName(ln)
		if err != nil {
			return err
		}
		obj = server.NotifyFindObject(ci, 0, ret).(objects.IGXDLMSBase)
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
				if err != nil {
					return err
				}
				if _, ok := value.([]byte); ok {
					value, err = internal.ChangeTypeFromByteArray(settings, value.([]byte), dt)
					if err != nil {
						return err
					}
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

func hanleSetRequestWithDataBlock(settings *settings.GXDLMSSettings,
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
		ret = uint8(enums.ErrorCodeDataBlockNumberInvalid)
	} else {
		settings.IncreaseBlockIndex()
		size, err := types.GetObjectCount(data)
		if err != nil {
			return err
		}
		realSize := data.Size() - data.Position()
		if size != realSize {
			log.Println("HanleSetRequestWithDataBlock failed. Invalid block size.")
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
		if !p.multipleBlocks {
			value, err := internal.GetData(settings, server.transaction.Data, &reply)
			if err != nil {
				return err
			}
			if _, ok := value.([]byte); ok {
				dt, err := server.transaction.Targets[0].Target.(objects.IGXDLMSBase).GetDataType(int(server.transaction.Targets[0].Index))
				if err != nil {
					return err
				}
				if _, ok := value.([]byte); ok {
					value, err = internal.ChangeTypeFromByteArray(settings, value.([]byte), dt)
					if err != nil {
						return err
					}
				}
			}
			server.transaction.Targets[0].Value = value
			server.NotifyWrite(server.transaction.Targets)
			if !server.transaction.Targets[0].Handled && !p.multipleBlocks {
				server.transaction.Targets[0].Target.(objects.IGXDLMSBase).SetValue(settings,
					server.transaction.Targets[0])
				server.NotifyPostWrite(server.transaction.Targets)
				ret = byte(server.transaction.Targets[0].Error)
			}
			settings.ResetBlockIndex()
		}
	}
	if ret != 0 {
		p.attributeDescriptor = types.NewGXByteBuffer()
		err = p.attributeDescriptor.SetUint8(ret)
	}
	p.multipleBlocks = true
	return err
}

func hanleSetRequestWithList(settings *settings.GXDLMSSettings,
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
		xml.AppendStartTag(int(internal.TranslatorTagsAttributeDescriptorList), "Qty", xml.IntegerToHex(cnt, 2, false), true)
	}
	for pos := 0; pos != cnt; pos++ {
		status[pos] = 0
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
		if xml != nil {
			xml.AppendStartTag(int(internal.TranslatorTagsAttributeDescriptorWithSelection), "", "", true)
			xml.AppendStartTag(int(internal.TranslatorTagsAttributeDescriptor), "", "", true)
			xml.AppendComment(ci.String())
			xml.AppendLineFromTag(int(internal.TranslatorTagsClassId), "Value", xml.IntegerToHex(int(ci), 4, false))
			xml.AppendComment(ln2)
			xml.AppendLineFromTag(int(internal.TranslatorTagsInstanceId), "Value", types.ToHex(ln, false))
			xml.AppendLineFromTag(int(internal.TranslatorTagsAttributeId), "Value", xml.IntegerToHex(attributeIndex, 2, false))
			xml.AppendEndTag(int(internal.TranslatorTagsAttributeDescriptor), false)
			xml.AppendEndTag(int(internal.TranslatorTagsAttributeDescriptorWithSelection), false)
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
	}
	cnt, err = types.GetObjectCount(data)
	if err != nil {
		return err
	}
	if xml != nil {
		xml.AppendEndTag(int(internal.TranslatorTagsAttributeDescriptorList), false)
		xml.AppendStartTag(int(internal.TranslatorTagsValueList), "Qty", xml.IntegerToHex(cnt, 2, false), true)
	}
	for pos := 0; pos != cnt; pos++ {
		if xml != nil || status[pos] == 0 {
			di := internal.GXDataInfo{}
			di.Xml = xml
			if xml != nil && xml.OutputType() == enums.TranslatorOutputTypeStandardXML {
				xml.AppendStartTag(int(enums.CommandWriteRequest<<8|constants.SingleReadResponseData), "", "", true)
			}
			value, err := internal.GetData(settings, data, &di)
			if err != nil {
				return err
			}
			if !di.Complete {
				value = types.ToHexWithRange(data.Array(), false, data.Position(), data.Size()-data.Position())
			} else if _, ok := value.([]byte); ok {
				value = types.ToHex(value.([]byte), false)
			}
			if xml != nil && xml.OutputType() == enums.TranslatorOutputTypeStandardXML {
				xml.AppendEndTag(int(enums.CommandWriteRequest<<8|constants.SingleReadResponseData), false)
			}
		}
	}
	if xml != nil {
		xml.AppendEndTag(int(internal.TranslatorTagsValueList), false)
	}
	p.status = 0xFF
	p.attributeDescriptor = types.NewGXByteBuffer()
	err = types.SetObjectCount(len(status), p.attributeDescriptor)
	if err != nil {
		return err
	}
	for _, it := range status {
		p.attributeDescriptor.SetUint8(it)
	}
	p.requestType = uint8(constants.SetResponseTypeWithList)
	return nil
}

// methodRequestNextDataBlock returns the handle method request next data block command.
//
// Parameters:
//
//	data: Received data.
func methodRequestNextDataBlock(settings *settings.GXDLMSSettings,
	server *GXDLMSServer,
	data *types.GXByteBuffer,
	invokeID uint8,
	replyData *types.GXByteBuffer,
	xml *settings.GXDLMSTranslatorStructure,
	streaming bool,
	cipheredCommand enums.Command) error {
	var err error
	bb := types.NewGXByteBuffer()
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
			log.Printf("handleGetRequest failed. Invalid block number %d/%d.", settings.BlockIndex, index)
			return getLNPdu(NewGXDLMSLNParameters(settings, 0, enums.CommandMethodResponse, 1, nil, bb, byte(enums.ErrorCodeDataBlockNumberInvalid), cipheredCommand), replyData)
		}
	}
	settings.IncreaseBlockIndex()
	cmd := enums.Command(enums.CommandMethodResponse)
	if streaming {
		cmd = enums.CommandGeneralBlockTransfer
	}
	p := NewGXDLMSLNParameters(settings, uint32(invokeID), cmd, byte(constants.ActionResponseTypeWithBlock), nil, bb, byte(enums.ErrorCodeOk), cipheredCommand)
	p.streaming = streaming
	p.gbtWindowSize = settings.GbtWindowSize()
	// If transaction is not in progress.
	if server.transaction == nil {
		p.status = uint8(enums.ErrorCodeNoLongGetOrReadInProgress)
		p.requestType = 1
		err = getLNPdu(p, replyData)
		if err != nil {
			return err
		}
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
						ret, err := internal.GetDLMSDataType(reflect.TypeOf(value))
						if err != nil {
							return err
						}
						err = internal.SetData(settings, bb, ret, value)
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
		getLNPdu(p, replyData)
		if moreData || bb.Size() != bb.Position() {
			server.transaction.Data = bb
		} else {
			server.SetTransaction(nil)
			settings.ResetBlockIndex()
		}
	}
	return err
}

// GetRequestNextDataBlock returns the handle get request next data block enums.Command
//
// Parameters:
//
//	data: Received data.
func GetRequestNextDataBlock(settings *settings.GXDLMSSettings,
	invokeID uint8,
	server *GXDLMSServer,
	data *types.GXByteBuffer,
	replyData *types.GXByteBuffer,
	xml *settings.GXDLMSTranslatorStructure,
	streaming bool,
	cipheredCommand enums.Command) error {
	var err error
	bb := types.NewGXByteBuffer()
	if !streaming {
		index, err := data.Uint32()
		if err != nil {
			return err
		}
		if xml != nil {
			xml.AppendLineFromTag(int(internal.TranslatorTagsBlockNumber), "", xml.IntegerToHex(index, 8, false))
			return nil
		}
		if index != settings.BlockIndex {
			log.Println("handleGetRequest failed. Invalid block number. " + strconv.Itoa(int(settings.BlockIndex)) + "/" + strconv.Itoa(int(index)))
			getLNPdu(NewGXDLMSLNParameters(settings, 0, enums.CommandGetResponse, 2, nil, bb, byte(enums.ErrorCodeDataBlockNumberInvalid), cipheredCommand), replyData)
			return nil
		}
	}
	settings.IncreaseBlockIndex()
	cmd := enums.Command(enums.CommandGetResponse)
	if streaming {
		cmd = enums.CommandGeneralBlockTransfer
	}
	p := NewGXDLMSLNParameters(settings, uint32(invokeID), cmd, 2, nil, bb, byte(enums.ErrorCodeOk), cipheredCommand)
	p.streaming = settings.GbtWindowSize() != 1
	p.gbtWindowSize = settings.GbtWindowSize()
	// If transaction is not in progress.
	if server.Transaction() == nil {
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
						err = appendData(settings, arg.Target.(objects.IGXDLMSBase), arg.Index, bb, value)
					}
				}
				moreData = settings.Index != settings.Count
			}
		}
		p.multipleBlocks = true
		getLNPdu(p, replyData)
		if moreData || bb.Size()-bb.Position() != 0 {
			server.transaction.Data = bb
		} else {
			server.SetTransaction(nil)
			settings.ResetBlockIndex()
		}
	}
	return err
}

// HandleEventNotification returns the handle Event Notification.
func handleEventNotification(settings *settings.GXDLMSSettings,
	reply *GXReplyData,
	list []*types.GXKeyValuePair[objects.IGXDLMSBase, int]) error {
	reply.Time = time.Time{}
	// Check is there date-time.
	len, err := reply.Data.Uint8()
	if err != nil {
		return err
	}
	// If date time is given.
	var tmp []byte
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
		reply.Time = ret.(types.GXDateTime).Value
	}
	if reply.xml != nil {
		reply.xml.AppendStartTag(int(enums.CommandEventNotification), "", "", true)
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
	reply.Data.Get(ln)
	index, err := reply.Data.Uint8()
	if err != nil {
		return err
	}
	if reply.xml != nil {
		appendAttributeDescriptor(reply.xml, int(ci), ln, index)
		reply.xml.AppendStartTag(int(internal.TranslatorTagsAttributeValue), "", "", true)
	}
	di := internal.GXDataInfo{}
	di.Xml = reply.xml
	value, err := internal.GetData(settings, reply.Data, &di)
	if err != nil {
		return err
	}
	if reply.xml != nil {
		reply.xml.AppendEndTag(int(internal.TranslatorTagsAttributeValue), true)
		reply.xml.AppendEndTag(int(enums.CommandEventNotification), true)
	} else {
		var obj objects.IGXDLMSBase
		ln2, err := helpers.ToLogicalName(ln)
		if err != nil {
			return err
		}
		if enums.ObjectType(ci) == enums.ObjectTypeAssociationLogicalName && ln2 == "0.0.40.0.0.255" {
			obj = settings.AssignedAssociation().(objects.IGXDLMSBase)
		}
		if obj == nil {
			obj = getObjectCollection(settings.Objects).FindByLN(enums.ObjectType(ci), ln2)
		}
		if obj != nil {
			v := internal.NewValueEventArgs3(obj, index, 0, nil)
			v.Value = value
			obj.SetValue(settings, v)
			list = append(list, types.NewGXKeyValuePair[objects.IGXDLMSBase, int](obj, int(index)))
		}
	}
	return err
}

func handleGetRequest(settings *settings.GXDLMSSettings,
	server *GXDLMSServer,
	data *types.GXByteBuffer,
	replyData *types.GXByteBuffer,
	xml *settings.GXDLMSTranslatorStructure,
	cipheredCommand enums.Command) error {
	var err error
	// Return error if connection is not established.
	if xml == nil && (settings.Connected&enums.ConnectionStateDlms) == 0 && cipheredCommand == enums.CommandNone {
		return replyData.Set(GenerateConfirmedServiceError(enums.ConfirmedServiceErrorInitiateError, enums.ServiceErrorService, uint8(enums.ServiceUnsupported)))
	}
	invokeID := 0
	type_ := constants.GetCommandTypeNormal
	// If GBT is used data is empty.
	if data.Size() != 0 {
		ret, err := data.Uint8()
		if err != nil {
			return err
		}
		type_ = constants.GetCommandType(ret)
		invokeID, err := data.Uint8()
		if err != nil {
			return err
		}
		settings.UpdateInvokeID(invokeID)
		if xml != nil {
			if type_ <= constants.GetCommandTypeWithList {
				addInvokeId(xml, enums.CommandGetRequest, int(type_), uint32(invokeID))
			} else {
				xml.AppendStartTag(int(enums.CommandGetRequest), "", "", true)
				xml.AppendComment(fmt.Sprintf("Unknown tag: %d.", type_))
				xml.AppendLineFromTag(int(internal.TranslatorTagsInvokeId), "Value", xml.IntegerToHex(invokeID, 2, false))
			}
		}
	}
	// GetRequest normal
	if type_ == constants.GetCommandTypeNormal {
		err = getRequestNormal(settings, uint8(invokeID), server, data, replyData, xml, cipheredCommand)
	} else if type_ == constants.GetCommandTypeNextDataBlock {
		err = GetRequestNextDataBlock(settings, uint8(invokeID), server, data, replyData, xml, false, cipheredCommand)
	} else if type_ == constants.GetCommandTypeWithList {
		err = GetRequestWithList(settings, uint8(invokeID), server, data, replyData, xml, cipheredCommand)
	} else if xml == nil {
		log.Println("HandleGetRequest failed. Invalid command type.")
		settings.ResetBlockIndex()
		type_ = constants.GetCommandTypeNormal
		data.Clear()
		err = getLNPdu(NewGXDLMSLNParameters(settings, uint32(invokeID), enums.CommandGetResponse, byte(type_), nil, nil,
			byte(enums.ErrorCodeReadWriteDenied), cipheredCommand), replyData)
	}
	if xml != nil {
		if type_ <= constants.GetCommandTypeWithList {
			xml.AppendEndTag(int(enums.CommandGetRequest<<8|type_), false)
		}
		xml.AppendEndTag(int(enums.CommandGetRequest), false)
	}
	return err
}

// HandleSetRequest returns the handle set request.
//
// Returns:
//
//	Reply to the client.
func handleSetRequest(settings *settings.GXDLMSSettings,
	server *GXDLMSServer,
	data *types.GXByteBuffer,
	replyData *types.GXByteBuffer,
	xml *settings.GXDLMSTranslatorStructure,
	cipheredCommand enums.Command) error {
	// Return error if connection is not established.
	if xml == nil && (settings.Connected&enums.ConnectionStateDlms) == 0 && cipheredCommand == enums.CommandNone {
		return replyData.Set(GenerateConfirmedServiceError(enums.ConfirmedServiceErrorInitiateError, enums.ServiceErrorService, uint8(enums.ServiceUnsupported)))
	}
	// Get type.
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
			xml.AppendComment(fmt.Sprintf("Unknown tag: %d.", type_))
			xml.AppendLineFromTag(int(internal.TranslatorTagsInvokeId), "Value", xml.IntegerToHex(invoke, 2, false))
		}
	}
	switch type_ {
	case uint8(constants.SetRequestTypeNormal):
	case uint8(constants.SetRequestTypeFirstDataBlock):
		if type_ == uint8(constants.SetRequestTypeNormal) {
			p.status = 0
		}
		err = handleSetRequestNormal(settings, server, data, byte(type_), p, replyData, xml)
	case byte(constants.SetRequestTypeWithDataBlock):
		err = hanleSetRequestWithDataBlock(settings, server, data, p, replyData, xml)
	case byte(constants.SetRequestTypeWithList):
		err = hanleSetRequestWithList(settings, invoke, server, data, p, replyData, xml)
	default:
		log.Println("HandleSetRequest failed. Unknown command.")
		data.Clear()
		settings.ResetBlockIndex()
		p.status = byte(enums.ErrorCodeReadWriteDenied)
	}
	if xml != nil {
		if type_ < 6 {
			xml.AppendEndTag(int(enums.CommandSetRequest)<<8|int(type_), false)
		}
		xml.AppendEndTag(int(enums.CommandSetRequest), false)
		return nil
	}
	return getLNPdu(p, replyData)
}

func methodRequest(settings *settings.GXDLMSSettings,
	type_ constants.ActionRequestType,
	invokeId uint8,
	server *GXDLMSServer,
	data *types.GXByteBuffer,
	connectionInfo *GXDLMSConnectionEventArgs,
	replyData *types.GXByteBuffer,
	xml *settings.GXDLMSTranslatorStructure,
	cipheredCommand enums.Command) error {
	error_ := enums.ErrorCode(enums.ErrorCodeOk)
	bb := types.NewGXByteBuffer()
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
	p := NewGXDLMSLNParameters(settings, uint32(invokeId), enums.CommandMethodResponse,
		byte(constants.ActionResponseTypeNormal), nil, bb, 0, cipheredCommand)
	if type_ == constants.ActionRequestTypeNormal {
		// Get parameters.
		selection, err := data.Uint8()
		if err != nil {
			return err
		}
		if xml != nil {
			AppendMethodDescriptor(xml, int(ci), ln, id)
			if selection != 0 {
				xml.AppendStartTag(int(internal.TranslatorTagsMethodInvocationParameters), "", "", true)
				di := internal.GXDataInfo{}
				di.Xml = xml
				_, err = internal.GetData(settings, data, &di)
				xml.AppendEndTag(int(internal.TranslatorTagsMethodInvocationParameters), false)
			}
			return err
		}
		if selection != 0 {
			info := internal.GXDataInfo{}
			parameters, err = internal.GetData(settings, data, &info)
		}
	} else if type_ == constants.ActionRequestTypeWithFirstBlock {
		p.requestType = uint8(constants.ActionResponseTypeNextBlock)
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
			log.Printf("MethodRequest failed. Invalid block number %d/%d. ", settings.BlockIndex, blockNumber)
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
			AppendMethodDescriptor(xml, int(ci), ln, id)
			xml.AppendStartTag(int(internal.TranslatorTagsDataBlock), "", "", true)
			xml.AppendLineFromTag(int(internal.TranslatorTagsLastBlock), "Value", xml.IntegerToHex(lastBlock, 2, false))
			xml.AppendLineFromTag(int(internal.TranslatorTagsBlockNumber), "Value", xml.IntegerToHex(blockNumber, 8, false))
			xml.AppendLineFromTag(int(internal.TranslatorTagsRawData), "Value", data.RemainingHexString(false))
			xml.AppendEndTag(int(internal.TranslatorTagsDataBlock), false)
			return nil
		}
	}
	var obj objects.IGXDLMSBase
	ln2, err := helpers.ToLogicalName(ln)
	if err != nil {
		return err
	}
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
	var e *internal.ValueEventArgs
	if obj == nil {
		error_ = enums.ErrorCodeUndefinedObject
	} else {
		if settings.AssignedAssociation != nil {
			p.AccessMode = int(settings.AssignedAssociation().(*objects.GXDLMSAssociationLogicalName).GetObjectMethodAccess3(obj, int(id)))
		}
		e = internal.NewValueEventArgs2(server, obj, id)
		e.Parameters = parameters
		e.InvokeId = p.invokeId
		if (server.NotifyGetMethodAccess(e) & int(enums.MethodAccessModeAccess)) == 0 {
			error_ = enums.ErrorCodeReadWriteDenied
		} else {
			if p.multipleBlocks {
				server.transaction = newGXDLMSLongTransaction([]*internal.ValueEventArgs{e}, enums.CommandMethodRequest, data)
			} else if server.Transaction() == nil {
				//Check transaction so invoke is not called multiple times.This might happen when all data can't fit to one PDU.
				p.requestType = uint8(constants.ActionResponseTypeNormal)
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
					if e.ByteArray {
						err = bb.Set(actionReply)
						if err != nil {
							return err
						}
					} else {
						ret, err := internal.GetDLMSDataType(reflect.TypeOf(actionReply))
						if err != nil {
							return err
						}
						err = internal.SetData(settings, bb, ret, actionReply)
						if err != nil {
							return err
						}
					}
				} else {
					error_ = e.Error
					err = bb.SetUint8(0)
					if err != nil {
						return err
					}
				}
			}
		}
		invokeId = uint8(e.InvokeId)
	}
	if error_ != 0 {
		p.status = uint8(error_)
	}
	err = getLNPdu(p, replyData)
	if err != nil {
		return err
	}
	// If all reply data doesn't fit to one PDU.
	if server.Transaction() == nil && p.data.Available() != 0 {
		server.transaction = newGXDLMSLongTransaction([]*internal.ValueEventArgs{e}, enums.CommandMethodResponse, data)
	}
	// If High level authentication fails.
	if a, ok := obj.(*objects.GXDLMSAssociationLogicalName); ok {
		if a.AssociationStatus == enums.AssociationStatusAssociated {
			server.NotifyConnected(connectionInfo)
			settings.Connected |= enums.ConnectionStateDlms
		} else {
			server.NotifyInvalidConnection(connectionInfo)
			settings.Connected &= ^enums.ConnectionStateDlms
		}
	}
	// Start to use new keys.
	if error_ == 0 {
		if a, ok := obj.(*objects.GXDLMSSecuritySetup); ok {
			a.ApplyKeys(settings, e)
		}
	}
	return err
}

func methodRequestNextBlock(settings *settings.GXDLMSSettings,
	server *GXDLMSServer,
	data *types.GXByteBuffer,
	connectionInfo *GXDLMSConnectionEventArgs,
	replyData *types.GXByteBuffer,
	xml *settings.GXDLMSTranslatorStructure,
	streaming bool,
	cipheredCommand enums.Command) error {
	var e *internal.ValueEventArgs
	bb := types.NewGXByteBuffer()
	lastBlock, err := data.Uint8()
	if err != nil {
		return err
	}
	if !streaming {
		var blockNumber uint32
		blockNumber, err := data.Uint32()
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
			xml.AppendLineFromTag(int(internal.TranslatorTagsLastBlock), "", xml.IntegerToHex(lastBlock, 2, false))
			xml.AppendLineFromTag(int(internal.TranslatorTagsBlockNumber), "", xml.IntegerToHex(blockNumber, 8, false))
			xml.AppendLineFromTag(int(internal.TranslatorTagsRawData), "", data.RemainingHexString(false))
			xml.AppendEndTag(int(internal.TranslatorTagsDataBlock), false)
			return nil
		}
		if blockNumber != settings.BlockIndex {
			server.SetTransaction(nil)
			log.Printf("MethodRequestNextBlock failed. Invalid block number %d/%d. ", settings.BlockIndex, blockNumber)
			settings.ResetBlockIndex()
			err := getLNPdu(NewGXDLMSLNParameters(settings, 0, enums.CommandMethodResponse,
				byte(constants.ActionResponseTypeNormal), nil, bb, byte(enums.ErrorCodeDataBlockNumberInvalid), cipheredCommand), replyData)
			return err
		}
		if size < data.Available() {
			server.SetTransaction(nil)
			settings.ResetBlockIndex()
			log.Printf("MethodRequestNextBlock failed. Not enough data. Actual: %d. Expected: %d", data.Available(), size)
			err := getLNPdu(NewGXDLMSLNParameters(settings, 0, enums.CommandMethodResponse,
				byte(constants.ActionResponseTypeNormal), nil, bb, byte(enums.ErrorCodeDataBlockNumberInvalid), cipheredCommand), replyData)
			return err
		}
	}
	cmd := enums.Command(enums.CommandMethodResponse)
	if streaming {
		cmd = enums.CommandGeneralBlockTransfer
	}
	p := NewGXDLMSLNParameters(settings, 0, cmd,
		byte(constants.ActionResponseTypeNormal), nil, bb, byte(enums.ErrorCodeOk), cipheredCommand)
	p.multipleBlocks = lastBlock == 0
	p.streaming = streaming
	p.gbtWindowSize = settings.GbtWindowSize()
	// If transaction is not in progress.
	if server.transaction == nil {
		p.status = uint8(enums.ErrorCodeNoLongGetOrReadInProgress)
	} else {
		err = server.transaction.Data.SetByteBuffer(data)
		if err != nil {
			return err
		}
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
					if s, ok := arg.Target.(*objects.GXDLMSSecuritySetup); ok {
						s.ApplyKeys(settings, e)
					}
					err = bb.SetUint8(1)
					if err != nil {
						return err
					}
					err = bb.SetUint8(0)
					if err != nil {
						return err
					}
					ret, err := internal.GetDLMSDataType(reflect.TypeOf(value))
					if err != nil {
						return err
					}
					err = internal.SetData(settings, bb, ret, value)
				} else {
					p.status = uint8(arg.Error)
					err = bb.SetUint8(0)
					if err != nil {
						return err
					}
				}
			}
			server.SetTransaction(nil)
			settings.ResetBlockIndex()
			p.blockIndex = 1
		} else {
			p.requestType = uint8(constants.ActionResponseTypeNextBlock)
			p.status = 0xFF
		}
	}

	err = getLNPdu(p, replyData)
	if settings.Count != settings.Index || bb.Size() != bb.Position() {
		server.SetTransaction(newGXDLMSLongTransaction([]*internal.ValueEventArgs{e}, enums.CommandMethodRequest, bb))
	}
	if lastBlock == 0 {
		settings.IncreaseBlockIndex()
	}
	return err
}

// HandleMethodRequest returns the handle action request.
func handleMethodRequest(settings *settings.GXDLMSSettings,
	server *GXDLMSServer,
	data *types.GXByteBuffer,
	connectionInfo *GXDLMSConnectionEventArgs,
	replyData *types.GXByteBuffer,
	xml *settings.GXDLMSTranslatorStructure,
	cipheredCommand enums.Command) error {
	// Get type_.
	var invokeID uint8
	ret, err := data.Uint8()
	if err != nil {
		return err
	}
	type_ := constants.ActionRequestType(ret)
	invokeID, err = data.Uint8()
	if err != nil {
		return err
	}
	settings.UpdateInvokeID(invokeID)
	if xml != nil {
		if type_ > 0 && type_ <= constants.ActionRequestTypeWithBlock {
			addInvokeId(xml, enums.CommandMethodRequest, int(type_), uint32(invokeID))
		} else {
			xml.AppendStartTag(int(enums.CommandMethodRequest), "", "", true)
			xml.AppendComment(fmt.Sprintf("Unknown tag: %d.", type_))
			xml.AppendLineFromTag(int(internal.TranslatorTagsInvokeId), "Value", xml.IntegerToHex(invokeID, 2, false))
		}
	}
	switch type_ {
	case constants.ActionRequestTypeNormal:
	case constants.ActionRequestTypeWithFirstBlock:
		err = methodRequest(settings, type_, invokeID, server, data, connectionInfo, replyData, xml, cipheredCommand)
	case constants.ActionRequestTypeNextBlock:
		err = methodRequestNextDataBlock(settings, server, data, invokeID, replyData, xml, false, cipheredCommand)
	case constants.ActionRequestTypeWithBlock:
		err = methodRequestNextBlock(settings, server, data, connectionInfo, replyData, xml, false, cipheredCommand)
	default:
		if xml == nil {
			log.Println("HandleMethodRequest failed. Invalid command type_.")
			settings.ResetBlockIndex()
			type_ = constants.ActionRequestTypeNormal
			data.Clear()
			err = getLNPdu(NewGXDLMSLNParameters(settings, uint32(invokeID), enums.CommandMethodResponse, byte(type_), nil, nil, byte(enums.ErrorCodeReadWriteDenied), cipheredCommand), replyData)
		}
	}
	if xml != nil {
		if type_ > 0 && type_ <= constants.ActionRequestTypeWithBlock {
			xml.AppendEndTag(int(enums.CommandMethodRequest<<8|type_), false)
		} else {
			xml.AppendEndTag(int(enums.CommandMethodRequest), false)
		}
		xml.AppendEndTag(int(enums.CommandMethodRequest), false)
	}
	return err
}

// handleAccessRequest returns the handle Access request.
func handleAccessRequest(settings *settings.GXDLMSSettings,
	server *GXDLMSServer,
	data *types.GXByteBuffer,
	reply *types.GXByteBuffer,
	xml *settings.GXDLMSTranslatorStructure,
	cipheredCommand enums.Command) error {
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
			dt := enums.DataType(enums.DataTypeDateTime)
			if len == 4 {
				dt = enums.DataTypeTime
			} else if len == 5 {
				dt = enums.DataTypeDate
			}
			info := internal.GXDataInfo{}
			info.Type = dt
			_, err := internal.GetData(settings, types.NewGXByteBufferWithData(tmp), &info)
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
		xml.AppendStartTag(int(enums.CommandAccessRequest), "", "", true)
		xml.AppendLineFromTag(int(internal.TranslatorTagsLongInvokeId), "Value", xml.IntegerToHex(invokeId, 2, false))
		xml.AppendLineFromTag(int(internal.TranslatorTagsDateTime), "Value", types.ToHex(tmp, false))
		xml.AppendStartTag(int(internal.TranslatorTagsAccessRequestBody), "", "", true)
		xml.AppendStartTag(int(internal.TranslatorTagsListOfAccessRequestSpecification), "Qty", xml.IntegerToHex(cnt, 2, false), true)
	}
	list := []*GXDLMSAccessItem{}
	var type_ enums.AccessServiceCommandType
	for pos := 0; pos != cnt; pos++ {
		ret, err := data.Uint8()
		if err != nil {
			return err
		}
		type_ = enums.AccessServiceCommandType(ret)
		if !(type_ == enums.AccessServiceCommandTypeGet || type_ == enums.AccessServiceCommandTypeSet ||
			type_ == enums.AccessServiceCommandTypeAction) {
			return errors.New("Invalid access service command type.")
		}
		// CI
		ret2, err := data.Uint16()
		if err != nil {
			return err
		}
		ci := enums.ObjectType(ret2)
		ln := make([]byte, 6)
		err = data.Get(ln)
		if err != nil {
			return err
		}
		// Attribute Id
		attributeIndex, err := data.Uint8()
		if err != nil {
			return err
		}
		if xml != nil {
			xml.AppendStartTag(int(internal.TranslatorTagsAccessRequestSpecification), "", "", true)
			xml.AppendStartTag(int(enums.CommandAccessRequest<<8|int(type_)), "", "", true)
			appendAttributeDescriptor(xml, int(ci), ln, attributeIndex)
			xml.AppendEndTag(int(enums.CommandAccessRequest<<8|int(type_)), false)
			xml.AppendEndTag(int(internal.TranslatorTagsAccessRequestSpecification), false)
		} else {
			var obj objects.IGXDLMSBase
			ln2, err := helpers.ToLogicalName(ln)
			if err != nil {
				return err
			}
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
		xml.AppendEndTag(int(internal.TranslatorTagsListOfAccessRequestSpecification), false)
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
	err = types.SetObjectCount(cnt, bb)
	if err != nil {
		return err
	}
	results := types.NewGXByteBuffer()
	err = types.SetObjectCount(cnt, results)
	if err != nil {
		return err
	}
	for pos := 0; pos != cnt; pos++ {
		di := internal.GXDataInfo{}
		di.Xml = xml
		if xml != nil && xml.OutputType() == enums.TranslatorOutputTypeStandardXML {
			xml.AppendStartTag(int(enums.CommandWriteRequest<<8|constants.SingleReadResponseData), "", "", true)
		}
		value, err := internal.GetData(settings, data, &di)
		if err != nil {
			return err
		}
		if !di.Complete {
			value = types.ToHexWithRange(data.Array(), false, data.Position(), data.Size()-data.Position())
		} else if _, ok := value.([]byte); ok {
			value = types.ToHex(value.([]byte), false)
		}
		if xml == nil {
			it := list[pos]
			err = results.SetUint8(uint8(it.Command))
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
				e := internal.NewValueEventArgs(settings, it.Target, it.Index)
				e.Parameters = value
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
							value, err = it.Target.GetValue(settings, e)
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
								if err != nil {
									return err
								}
							} else {
								appendData(settings, it.Target, it.Index, bb, value)
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
			xml.AppendEndTag(int(enums.CommandWriteRequest<<8|constants.SingleReadResponseData), false)
		}
	}
	if xml != nil {
		xml.AppendEndTag(int(internal.TranslatorTagsAccessRequestListOfData), false)
		xml.AppendEndTag(int(internal.TranslatorTagsAccessRequestBody), false)
		xml.AppendEndTag(int(enums.CommandAccessRequest), false)
	} else {
		err = bb.SetByteBuffer(results)
		if err != nil {
			return err
		}
		err = getLNPdu(NewGXDLMSLNParameters(settings, invokeId, enums.CommandAccessResponse, 0xff, nil, bb, 0xFF, cipheredCommand), reply)
	}
	return err
}
