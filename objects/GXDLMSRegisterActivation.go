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

import (
	"bytes"
	"errors"
	"strings"

	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// Online help:
// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSRegisterActivation
type GXDLMSRegisterActivation struct {
	GXDLMSObject
	// Assignment list.
	RegisterAssignment []GXDLMSObjectDefinition
	// Mask list.
	MaskList []types.GXKeyValuePair[[]byte, []byte]
	// Active mask.
	ActiveMask []byte
}

// base returns the base GXDLMSObject of the object.
func (g *GXDLMSRegisterActivation) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

func (g *GXDLMSRegisterActivation) getMaskIndex(maskName []byte) int {
	index := 0
	for _, v := range g.MaskList {
		if bytes.Equal(v.Key, maskName) {
			return index
		}
		index++
	}
	return -1
}

// Invoke returns the invokes method.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Invoke parameters.
func (g *GXDLMSRegisterActivation) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	var err error
	if e.Index == 1 {
		items := []GXDLMSObjectDefinition{}
		if g.RegisterAssignment != nil {
			items = append(items, g.RegisterAssignment...)
		}
		item := GXDLMSObjectDefinition{}
		s := e.Parameters.(types.GXStructure)
		item.objectType = enums.ObjectType(s[0].(int))
		item.LogicalName, err = helpers.ToLogicalName(s[1].([]byte))
		if err != nil {
			return nil, err
		}
		items = append(items, item)
		g.RegisterAssignment = items
	} else if e.Index == 2 {
		s := e.Parameters.(types.GXStructure)
		index := g.getMaskIndex(s[0].([]byte))
		var index_list []byte
		for _, b := range s[1].([]any) {
			index_list = append(index_list, b.(byte))
		}
		if index == -1 {
			g.MaskList = append(g.MaskList, *types.NewGXKeyValuePair(s[0].([]byte), index_list))
		} else {
			g.MaskList = append(g.MaskList[:index], g.MaskList[index+1:]...)
			g.MaskList, err = internal.InsertAt(g.MaskList, index, *types.NewGXKeyValuePair(s[0].([]byte), index_list))
			if err != nil {
				return nil, err
			}
		}
	} else if e.Index == 3 {
		index := g.getMaskIndex(e.Parameters.([]byte))
		if index != -1 {
			g.MaskList = append(g.MaskList[:index], g.MaskList[index+1:]...)
		}
	} else {
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil, nil
}

// GetAttributeIndexToRead returns the collection of attributes to read.
// If attribute is static and already read or device is returned HW error it is not returned.
//
// Parameters:
//
//	all: All items are returned even if they are read already.
//
// Returns:
//
//	Collection of attributes to read.
func (g *GXDLMSRegisterActivation) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	// LN is static and read only once.
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	// RegisterAssignment
	if all || !g.IsRead(2) {
		attributes = append(attributes, 2)
	}
	// MaskList
	if all || !g.IsRead(3) {
		attributes = append(attributes, 3)
	}
	// ActiveMask
	if all || !g.IsRead(4) {
		attributes = append(attributes, 4)
	}
	return attributes
}

// GetNames returns the names of attribute indexes.
func (g *GXDLMSRegisterActivation) GetNames() []string {
	return []string{"Logical Name", "Register Assignment", "Mask List", "Active Mask"}
}

// GetMethodNames returns the names of method indexes.
func (g *GXDLMSRegisterActivation) GetMethodNames() []string {
	return []string{"Add register", "Add mask", "Delete mask"}
}

// GetAttributeCount returns the amount of attributes.
//
// Returns:
//
//	Count of attributes.
func (g *GXDLMSRegisterActivation) GetAttributeCount() int {
	return 4
}

// GetMethodCount returns the amount of methods.
func (g *GXDLMSRegisterActivation) GetMethodCount() int {
	return 3
}

// GetValue returns the value of given attribute.
// When raw parameter us not used example register multiplies value by scalar.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Get parameters.
//
// Returns:
//
//	Value of the attribute index.
func (g *GXDLMSRegisterActivation) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	var err error
	if e.Index == 1 {
		v, err := helpers.LogicalNameToBytes(g.LogicalName())
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return v, err
	}
	if e.Index == 2 {
		data := types.NewGXByteBuffer()
		err = data.SetUint8(uint8(enums.DataTypeArray))
		if err != nil {
			return nil, err
		}
		if g.RegisterAssignment == nil {
			err = data.SetUint8(0)
			if err != nil {
				return nil, err
			}
		} else {
			data.SetUint8(byte(len(g.RegisterAssignment)))
			for _, it := range g.RegisterAssignment {
				err = data.SetUint8(uint8(enums.DataTypeStructure))
				if err != nil {
					return nil, err
				}
				err = data.SetUint8(2)
				if err != nil {
					return nil, err
				}
				err = internal.SetData(settings, data, enums.DataTypeUint16, it.ObjectType())
				if err != nil {
					return nil, err
				}
				ln, err := helpers.LogicalNameToBytes(it.LogicalName)
				if err != nil {
					return nil, err
				}
				err = internal.SetData(settings, data, enums.DataTypeOctetString, ln)
				if err != nil {
					return nil, err
				}
			}
		}
		return data.Array(), nil
	}
	if e.Index == 3 {
		data := types.NewGXByteBuffer()
		err = data.SetUint8(uint8(enums.DataTypeArray))
		if err != nil {
			return nil, err
		}
		if g.MaskList == nil {
			err = data.SetUint8(0)
			if err != nil {
				return nil, err
			}
		} else {
			err = data.SetUint8(uint8(len(g.MaskList)))
			if err != nil {
				return nil, err
			}
			for _, it := range g.MaskList {
				err = data.SetUint8(uint8(enums.DataTypeStructure))
				if err != nil {
					return nil, err
				}
				err = data.SetUint8(2)
				if err != nil {
					return nil, err
				}
				err = internal.SetData(settings, data, enums.DataTypeOctetString, it.Key)
				if err != nil {
					return nil, err
				}
				err = data.SetUint8(uint8(enums.DataTypeArray))
				if err != nil {
					return nil, err
				}
				data.SetUint8(byte(len(it.Value)))
				for _, b := range it.Value {
					err = internal.SetData(settings, data, enums.DataTypeUint8, b)
				}
			}
		}
		return data.Array(), nil
	}
	if e.Index == 4 {
		return g.ActiveMask, nil
	}
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, nil
}

// SetValue returns the set value of given attribute.
// When raw parameter us not used example register multiplies value by scalar.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Set parameters.
func (g *GXDLMSRegisterActivation) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	var err error
	if e.Index == 1 {
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return g.SetLogicalName(ln)
	} else if e.Index == 2 {
		items := []GXDLMSObjectDefinition{}
		if e.Value != nil {
			for _, tmp := range e.Value.(types.GXArray) {
				it := tmp.(types.GXStructure)
				item := GXDLMSObjectDefinition{}
				item.objectType = enums.ObjectType(it[0].(uint16))
				item.LogicalName, err = helpers.ToLogicalName(it[1].([]byte))
				if err != nil {
					return err
				}
				items = append(items, item)
			}
		}
		g.RegisterAssignment = items
	} else if e.Index == 3 {
		g.MaskList = g.MaskList[:0]
		if e.Value != nil {
			for _, tmp := range e.Value.(types.GXArray) {
				it := tmp.(types.GXStructure)
				var index_list []byte
				for _, b := range it[1].(types.GXArray) {
					index_list = append(index_list, b.(byte))
				}
				g.MaskList = append(g.MaskList, *types.NewGXKeyValuePair(it[0].([]byte), index_list))
			}
		}
	} else if e.Index == 4 {
		if e.Value == nil {
			g.ActiveMask = nil
		} else {
			g.ActiveMask = e.Value.([]byte)
		}
	} else {
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil
}

// Load returns the load object content from XML.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSRegisterActivation) Load(reader *GXXmlReader) error {
	list := []GXDLMSObjectDefinition{}
	if reader.isStartElementNamed2("RegisterAssignment", true) {
		for reader.isStartElementNamed2("Item", true) {
			it := GXDLMSObjectDefinition{}
			ret, err := reader.ReadElementContentAsInt("ObjectType", 0)
			if err != nil {
				return err
			}
			it.objectType = enums.ObjectType(ret)
			if err != nil {
				return err
			}
			it.LogicalName, err = reader.ReadElementContentAsString("LN", "")
			if err != nil {
				return err
			}
			list = append(list, it)
		}
		reader.ReadEndElement("RegisterAssignment")
	}
	g.RegisterAssignment = list
	g.MaskList = g.MaskList[:0]
	if reader.isStartElementNamed2("MaskList", true) {
		for reader.isStartElementNamed2("Item", true) {
			ret, err := reader.ReadElementContentAsString("Mask", "")
			if err != nil {
				return err
			}
			mask := types.HexToBytes(ret)
			ret, err = reader.ReadElementContentAsString("Index", "")
			i := types.HexToBytes(ret)
			g.MaskList = append(g.MaskList, *types.NewGXKeyValuePair(mask, i))
		}
		reader.ReadEndElement("MaskList")
	}
	ret, err := reader.ReadElementContentAsString("ActiveMask", "")
	if err != nil {
		return err
	}
	g.ActiveMask = types.HexToBytes(ret)
	return err
}

// Save returns the save object content to XML.
//
// Parameters:
//
//	writer: XML writer.
func (g *GXDLMSRegisterActivation) Save(writer *GXXmlWriter) error {
	var err error
	writer.WriteStartElement("RegisterAssignment")
	if g.RegisterAssignment != nil {
		for _, it := range g.RegisterAssignment {
			writer.WriteStartElement("Item")
			err = writer.WriteElementString("ObjectType", int(it.ObjectType()))
			if err != nil {
				return err
			}
			err = writer.WriteElementString("LN", it.LogicalName)
			if err != nil {
				return err
			}
			writer.WriteEndElement()
		}
	}
	writer.WriteEndElement()
	writer.WriteStartElement("MaskList")
	if g.MaskList != nil {
		for _, v := range g.MaskList {
			writer.WriteStartElement("Item")
			err = writer.WriteElementString("Mask", types.ToHex(v.Key, false))
			if err != nil {
				return err
			}
			ret := types.ToHex(v.Value, false)
			err = writer.WriteElementString("Index", strings.ReplaceAll(ret, " ", ";"))
			if err != nil {
				return err
			}
			writer.WriteEndElement()
		}
	}
	writer.WriteEndElement()
	if g.ActiveMask != nil {
		err = writer.WriteElementString("ActiveMask", types.ToHex(g.ActiveMask, false))
		if err != nil {
			return err
		}
	}
	return err
}

// PostLoad returns the handle actions after Load.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSRegisterActivation) PostLoad(reader *GXXmlReader) error {
	return nil
}

// GetValues returns the an array containing the COSEM object's attribute values.
func (g *GXDLMSRegisterActivation) GetValues() []any {
	return []any{g.LogicalName(), g.RegisterAssignment, g.MaskList, g.ActiveMask}
}

// AddRegister returns the add new register.
//
// Parameters:
//
//	client: DLMS client.
//	script: Register to add.
//
// Returns:
//
//	Action bytes.
func (g *GXDLMSRegisterActivation) AddRegister(client IGXDLMSClient, target IGXDLMSBase) ([][]uint8, error) {
	bb := types.NewGXByteBuffer()
	err := bb.SetUint8(byte(enums.DataTypeStructure))
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(2)
	if err != nil {
		return nil, err
	}
	err = internal.SetData(nil, bb, enums.DataTypeUint16, target.Base().ObjectType())
	if err != nil {
		return nil, err
	}
	ln, err := helpers.LogicalNameToBytes(target.Base().LogicalName())
	if err != nil {
		return nil, err
	}

	err = internal.SetData(nil, bb, enums.DataTypeOctetString, ln)
	if err != nil {
		return nil, err
	}
	return client.Method(g, 1, bb.Array(), enums.DataTypeArray)
}

// AddMask returns the add new register activation mask.
//
// Parameters:
//
//	client: DLMS client.
//	name: Register activation mask name.
//	indexes: Register activation indexes.
//
// Returns:
//
//	Action bytes.
func (g *GXDLMSRegisterActivation) AddMask(client IGXDLMSClient, name []byte, indexes []byte) ([][]uint8, error) {
	bb := types.NewGXByteBuffer()
	err := bb.SetUint8(enums.DataTypeStructure)
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(2)
	if err != nil {
		return nil, err
	}
	err = internal.SetData(nil, bb, enums.DataTypeOctetString, name)
	err = bb.SetUint8(enums.DataTypeArray)
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(uint8(len(indexes)))
	if err != nil {
		return nil, err
	}
	for _, it := range indexes {
		err = internal.SetData(nil, bb, enums.DataTypeUint8, it)
	}
	return client.Method(g, 2, bb.Array(), enums.DataTypeArray)
}

// RemoveMask returns the remove register activation mask.
//
// Parameters:
//
//	client: DLMS client.
//	name: Register activation mask name.
//
// Returns:
//
//	Action bytes.
func (g *GXDLMSRegisterActivation) RemoveMask(client IGXDLMSClient, name []byte) ([][]uint8, error) {
	return client.Method(g, 3, name, enums.DataTypeOctetString)
}

// GetDataType returns the device data type of selected attribute index.
//
// Parameters:
//
//	index: Attribute index of the object.
//
// Returns:
//
//	Device data type of the object.
func (g *GXDLMSRegisterActivation) GetDataType(index int) (enums.DataType, error) {
	var ret enums.DataType
	if index == 1 {
		ret = enums.DataTypeOctetString
	} else if index == 2 {
		ret = enums.DataTypeArray
	} else if index == 3 {
		ret = enums.DataTypeArray
	} else if index == 4 {
		ret = enums.DataTypeOctetString
	} else {
		return 0, errors.New("GetDataType failed. Invalid attribute index.")
	}
	return ret, nil
}

// Constructor.
// ln: Logical Name of the object.
// sn: Short Name of the object.
func NewGXDLMSRegisterActivation(ln string, sn int16) (*GXDLMSRegisterActivation, error) {
	err := ValidateLogicalName(ln)
	if err != nil {
		return nil, err
	}
	return &GXDLMSRegisterActivation{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeRegisterActivation,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
