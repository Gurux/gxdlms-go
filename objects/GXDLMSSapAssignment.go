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
	"fmt"

	"github.com/Gurux/gxdlms-go/dlmserrors"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// Online help:
// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSSapAssignment
type GXDLMSSapAssignment struct {
	GXDLMSObject
	SapAssignmentList []types.GXKeyValuePair[uint16, string]
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSSapAssignment) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

//GetAttributeIndexToRead returns the collection of attributes to read.
// If attribute is static and already read or device is returned HW error it is not returned.
//
// Parameters:
//   all: All items are returned even if they are read already.
//
// Returns:
//   Collection of attributes to read.
func (g *GXDLMSSapAssignment) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	// LN is static and read only once.
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	// SapAssignmentList
	if all || !g.IsRead(2) {
		attributes = append(attributes, 2)
	}
	return attributes
}

//GetNames returns the names of attribute indexes.
func (g *GXDLMSSapAssignment) GetNames() []string {
	return []string{"Logical Name", "Sap Assignment List"}
}

//GetMethodNames returns the names of method indexes.
func (g *GXDLMSSapAssignment) GetMethodNames() []string {
	return []string{"Connect logical device"}
}

//GetAttributeCount returns the amount of attributes.
//
// Returns:
//   Count of attributes.
func (g *GXDLMSSapAssignment) GetAttributeCount() int {
	return 2
}

//GetMethodCount returns the amount of methods.
func (g *GXDLMSSapAssignment) GetMethodCount() int {
	return 1
}

//GetValue returns the value of given attribute.
// When raw parameter us not used example register multiplies value by scalar.
//
// Parameters:
//   settings: DLMS settings.
//   e: Get parameters.
//
// Returns:
//   Value of the attribute index.
func (g *GXDLMSSapAssignment) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	if e.Index == 1 {
		v, err := helpers.LogicalNameToBytes(g.LogicalName())
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return v, err
	}
	if e.Index == 2 {
		cnt := 0
		if g.SapAssignmentList != nil {
			cnt = len(g.SapAssignmentList)
		}
		data := types.NewGXByteBuffer()
		err := data.SetUint8(uint8(enums.DataTypeArray))
		if err != nil {
			return nil, err
		}
		types.SetObjectCount(cnt, data)
		if cnt != 0 {
			for _, it := range g.SapAssignmentList {
				err = data.SetUint8(uint8(enums.DataTypeStructure))
				if err != nil {
					return nil, err
				}
				err = data.SetUint8(uint8(2))
				if err != nil {
					return nil, err
				}
				err = internal.SetData(settings, data, enums.DataTypeUint16, it.Key)
				if err != nil {
					return nil, err
				}
				err = internal.SetData(settings, data, enums.DataTypeOctetString, []byte(it.Value))
				if err != nil {
					return nil, err
				}
			}
		}
		return data.Array(), nil
	}
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, nil
}

//SetValue returns the set value of given attribute.
// When raw parameter us not used example register multiplies value by scalar.
//
// Parameters:
//   settings: DLMS settings.
//   e: Set parameters.
func (g *GXDLMSSapAssignment) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	if e.Index == 1 {
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return g.SetLogicalName(ln)
	} else if e.Index == 2 {
		g.SapAssignmentList = g.SapAssignmentList[:0]
		if e.Value != nil {
			for _, tmp := range e.Value.(types.GXArray) {
				item := tmp.(types.GXStructure)
				var str string
				if _, ok := item[1].([]byte); ok {
					v, err := internal.ChangeTypeFromByteArray(settings, item[1].([]byte), enums.DataTypeString)
					if err != nil {
						return nil
					}
					str = fmt.Sprint(v)
				} else {
					str = fmt.Sprint(item[1])
				}
				g.SapAssignmentList = append(g.SapAssignmentList, *types.NewGXKeyValuePair(item[0].(uint16), str))
			}
		}
	} else {
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil
}

//Invoke returns the invokes method.
//
// Parameters:
//   settings: DLMS settings.
//   e: Invoke parameters.
func (g *GXDLMSSapAssignment) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	if e.Index == 1 {
		tmp := e.Parameters.(types.GXStructure)
		id := tmp[0].(uint16)
		var str string
		if v, ok := tmp[1].([]byte); ok {
			str = string(v)
		} else {
			str = fmt.Sprint(tmp[0])
		}
		if id == 0 {
			for _, it := range g.SapAssignmentList {
				if it.Value == str {
					internal.Remove(g.SapAssignmentList, it)
					break
				}
			}
		} else {
			g.SapAssignmentList = append(g.SapAssignmentList, *types.NewGXKeyValuePair(id, str))
		}
	} else {
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil, nil
}

//Load returns the load object content from XML.
//
// Parameters:
//   reader: XML reader.
func (g *GXDLMSSapAssignment) Load(reader *GXXmlReader) error {
	var sap uint16
	var err error
	var ldn string
	g.SapAssignmentList = g.SapAssignmentList[:0]
	if ret, err := reader.IsStartElementNamed("SapAssignmentList", true); ret && err == nil {
		for {
			ret, err = reader.IsStartElementNamed("Item", true)
			if err != nil {
				return err
			}
			if !ret {
				break
			}
			sap, err = reader.ReadElementContentAsUInt16("SAP", 0)
			if err != nil {
				return err
			}
			ldn, err = reader.ReadElementContentAsString("LDN", "")
			if err != nil {
				return err
			}
			g.SapAssignmentList = append(g.SapAssignmentList, *types.NewGXKeyValuePair(sap, ldn))
		}
		reader.ReadEndElement("SapAssignmentList")
	}
	return err
}

//Save returns the save object content to XML.
//
// Parameters:
//   writer: XML writer.
func (g *GXDLMSSapAssignment) Save(writer *GXXmlWriter) error {
	var err error
	if g.SapAssignmentList != nil {
		writer.WriteStartElement("SapAssignmentList")
		for k, v := range g.SapAssignmentList {
			writer.WriteStartElement("Item")
			err = writer.WriteElementString("SAP", k)
			if err != nil {
				return err
			}
			err = writer.WriteElementString("LDN", v)
			if err != nil {
				return err
			}
			writer.WriteEndElement()
		}
		writer.WriteEndElement()
	}
	return err
}

//PostLoad returns the handle actions after Load.
//
// Parameters:
//   reader: XML reader.
func (g *GXDLMSSapAssignment) PostLoad(reader *GXXmlReader) error {
	return nil
}

//GetValues returns the an array containing the COSEM object's attribute values.
func (g *GXDLMSSapAssignment) GetValues() []any {
	return []any{g.LogicalName(), g.SapAssignmentList}
}

//AddSap returns the add new SAP item.
//
// Parameters:
//   client: DLMS client.
//   id: SAP ID.
//   name: Logical device name.
//
// Returns:
//   Generated bytes sent to the meter.
func (g *GXDLMSSapAssignment) AddSap(client IGXDLMSClient, id uint16, name string) ([][]uint8, error) {
	data := types.NewGXByteBuffer()
	err := data.SetUint8(uint8(enums.DataTypeStructure))
	if err != nil {
		return nil, err
	}
	err = data.SetUint8(2)
	if err != nil {
		return nil, err
	}
	err = internal.SetData(nil, data, enums.DataTypeUint16, id)
	if err != nil {
		return nil, err
	}
	err = internal.SetData(nil, data, enums.DataTypeOctetString, []byte(name))
	if err != nil {
		return nil, err
	}
	return client.Method(g, 1, data.Array(), enums.DataTypeStructure)
}

//RemoveSap returns the remove SAP item.
//
// Parameters:
//   client: DLMS client.
//   name: Logical device name.
//
// Returns:
//   Generated bytes sent to the meter.
func (g *GXDLMSSapAssignment) RemoveSap(client IGXDLMSClient, name string) ([][]uint8, error) {
	data := types.NewGXByteBuffer()
	err := data.SetUint8(uint8(enums.DataTypeStructure))
	if err != nil {
		return nil, err
	}
	err = data.SetUint8(2)
	if err != nil {
		return nil, err
	}
	err = internal.SetData(nil, data, enums.DataTypeUint16, 0)
	if err != nil {
		return nil, err
	}
	err = internal.SetData(nil, data, enums.DataTypeOctetString, []byte(name))
	if err != nil {
		return nil, err
	}
	return client.Method(g, 1, data.Array(), enums.DataTypeStructure)
}

//GetDataType returns the device data type of selected attribute index.
//
// Parameters:
//   index: Attribute index of the object.
//
// Returns:
//   Device data type of the object.
func (g *GXDLMSSapAssignment) GetDataType(index int) (enums.DataType, error) {
	if index == 1 {
		return enums.DataTypeOctetString, nil
	}
	if index == 2 {
		return enums.DataTypeArray, nil
	}
	return 0, dlmserrors.ErrInvalidAttributeIndex
}

// NewGXDLMSSapAssignment creates a new Sap assignment object instance.
//
// The function validates `ln` before creating the object.
//`ln` is the Logical Name and `sn` is the Short Name of the object.
func NewGXDLMSSapAssignment(ln string, sn int16) (*GXDLMSSapAssignment, error) {
	err := ValidateLogicalName(ln)
	if err != nil {
		return nil, err
	}
	return &GXDLMSSapAssignment{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeSapAssignment,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
