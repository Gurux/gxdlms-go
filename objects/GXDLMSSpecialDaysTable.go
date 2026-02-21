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
	"errors"

	"github.com/Gurux/gxdlms-go/dlmserrors"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// Online help:
// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSSpecialDaysTable
type GXDLMSSpecialDaysTable struct {
	GXDLMSObject
	// Value of COSEM Data object.
	Entries []GXDLMSSpecialDay
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSSpecialDaysTable) Base() *GXDLMSObject {
	return &g.GXDLMSObject
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
func (g *GXDLMSSpecialDaysTable) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	// LN is static and read only once.
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	// Entries
	if all || !g.IsRead(2) {
		attributes = append(attributes, 2)
	}
	return attributes
}

// GetNames returns the names of attribute indexes.
func (g *GXDLMSSpecialDaysTable) GetNames() []string {
	return []string{"Logical Name", "Entries"}
}

// GetMethodNames returns the names of method indexes.
func (g *GXDLMSSpecialDaysTable) GetMethodNames() []string {
	return []string{"Insert", "Delete"}
}

// GetAttributeCount returns the amount of attributes.
//
// Returns:
//
//	Count of attributes.
func (g *GXDLMSSpecialDaysTable) GetAttributeCount() int {
	return 2
}

// GetMethodCount returns the amount of methods.
func (g *GXDLMSSpecialDaysTable) GetMethodCount() int {
	return 2
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
func (g *GXDLMSSpecialDaysTable) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	if e.Index == 1 {
		v, err := helpers.LogicalNameToBytes(g.LogicalName())
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return v, err
	}
	if e.Index == 2 {
		data := types.NewGXByteBuffer()
		err := data.SetUint8(uint8(enums.DataTypeArray))
		if err != nil {
			return nil, err
		}
		if g.Entries == nil {
			err = data.SetUint8(0)
			if err != nil {
				return nil, err
			}
		} else {
			cnt := len(g.Entries)
			types.SetObjectCount(cnt, data)
			if cnt != 0 {
				for _, it := range g.Entries {
					err = data.SetUint8(uint8(enums.DataTypeStructure))
					if err != nil {
						return nil, err
					}
					err = data.SetUint8(uint8(3))
					if err != nil {
						return nil, err
					}
					err = internal.SetData(settings, data, enums.DataTypeUint16, it.Index)
					if err != nil {
						return nil, err
					}
					if settings.Standard == enums.StandardSaudiArabia {
						err = internal.SetData(settings, data, enums.DataTypeDate, it.Date)
					} else {
						err = internal.SetData(settings, data, enums.DataTypeOctetString, it.Date)
					}
					if err != nil {
						return nil, err
					}
					err = internal.SetData(settings, data, enums.DataTypeUint8, it.DayId)
					if err != nil {
						return nil, err
					}
				}
			}
		}
		return data, nil
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
func (g *GXDLMSSpecialDaysTable) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	var err error
	if e.Index == 1 {
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return g.SetLogicalName(ln)
	} else if e.Index == 2 {
		g.Entries = nil
		if e.Value != nil {
			items := []GXDLMSSpecialDay{}
			for _, tmp := range e.Value.(types.GXArray) {
				item := tmp.(types.GXStructure)
				it := GXDLMSSpecialDay{}
				it.Index = item[0].(uint16)
				if v, ok := item[1].(types.GXDate); ok {
					it.Date = v
				} else if v, ok := item[1].([]byte); ok {
					ret, err := internal.ChangeTypeFromByteArray(settings, v, enums.DataTypeDate)
					if err != nil {
						return err
					}
					it.Date = ret.(types.GXDate)
				} else {
					return errors.New("Invalid date.")
				}
				it.DayId = item[2].(byte)
				items = append(items, it)
			}
			g.Entries = items
		}
	} else {
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return err
}

// Invoke returns the invokes method.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Invoke parameters.
func (g *GXDLMSSpecialDaysTable) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	if e.Index != 1 && e.Index != 2 {
		e.Error = enums.ErrorCodeReadWriteDenied
	} else {
		items := []GXDLMSSpecialDay{}
		if len(g.Entries) != 0 {
			items = append(items, g.Entries...)
		}
		if e.Index == 1 {
			item := e.Parameters.(types.GXStructure)
			it := GXDLMSSpecialDay{}
			it.Index = item[0].(uint16)
			ret, err := internal.ChangeTypeFromByteArray(settings, item[1].([]byte), enums.DataTypeDate)
			if err != nil {
				return nil, err
			}
			it.Date = ret.(types.GXDate)
			it.DayId = item[2].(byte)
			for _, item2 := range items {
				if item2.Index == it.Index {
					internal.Remove(items, item2)
					break
				}
			}
			items = append(items, it)
		} else if e.Index == 2 {
			index := e.Parameters.(uint16)
			for _, item := range items {
				if item.Index == index {
					internal.Remove(items, item)
					break
				}
			}
		}
		g.Entries = items
	}
	return nil, nil
}

// Load returns the load object content from XML.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSSpecialDaysTable) Load(reader *GXXmlReader) error {
	var err error
	list := []GXDLMSSpecialDay{}
	if ret, err := reader.IsStartElementNamed("Entries", true); ret && err == nil {
		for {
			ret, err = reader.IsStartElementNamed("Entry", true)
			if err != nil {
				return err
			}
			if !ret {
				break
			}
			it := GXDLMSSpecialDay{}
			ret, err := reader.ReadElementContentAsInt("Index", 0)
			if err != nil {
				return err
			}
			it.Index = uint16(ret)
			it.Date, err = reader.ReadElementContentAsGXDate("Date")
			if err != nil {
				return err
			}
			ret, err = reader.ReadElementContentAsInt("DayId", 0)
			it.DayId = uint8(ret)
			if err != nil {
				return err
			}
			list = append(list, it)
		}
		reader.ReadEndElement("Entries")
	}
	g.Entries = list
	return err
}

// Save returns the save object content to XML.
//
// Parameters:
//
//	writer: XML writer.
func (g *GXDLMSSpecialDaysTable) Save(writer *GXXmlWriter) error {
	var err error
	if g.Entries != nil {
		writer.WriteStartElement("Entries")
		for _, it := range g.Entries {
			writer.WriteStartElement("Entry")
			err = writer.WriteElementString("Index", it.Index)
			if err != nil {
				return err
			}
			err = writer.WriteElementString("Date", it.Date)
			if err != nil {
				return err
			}
			err = writer.WriteElementString("DayId", it.DayId)
			if err != nil {
				return err
			}
			writer.WriteEndElement()
		}
		writer.WriteEndElement()
	}
	return err
}

// PostLoad returns the handle actions after Load.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSSpecialDaysTable) PostLoad(reader *GXXmlReader) error {
	return nil
}

// Insert returns the inserts a new entry in the table.
//
// Returns:
//
//	If a special day with the same index or with the same date as an already defined day is inserted,
//	the old entry will be overwritten.
func (g *GXDLMSSpecialDaysTable) Insert(client IGXDLMSClient, entry *GXDLMSSpecialDay) ([][]uint8, error) {
	bb := types.NewGXByteBuffer()
	err := bb.SetUint8(enums.DataTypeStructure)
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(3)
	if err != nil {
		return nil, err
	}
	err = internal.SetData(nil, bb, enums.DataTypeUint16, entry.Index)
	if err != nil {
		return nil, err
	}
	err = internal.SetData(nil, bb, enums.DataTypeOctetString, entry.Date)
	if err != nil {
		return nil, err
	}
	err = internal.SetData(nil, bb, enums.DataTypeUint8, entry.DayId)
	if err != nil {
		return nil, err
	}
	return client.Method(g, 1, bb.Array(), enums.DataTypeArray)
}

// Delete returns the deletes an entry from the table.
func (g *GXDLMSSpecialDaysTable) Delete(client IGXDLMSClient, entry *GXDLMSSpecialDay) ([][]uint8, error) {
	return client.Method(g, 2, uint16(entry.Index), enums.DataTypeUint16)
}

// GetValues returns the an array containing the COSEM object's attribute values.
func (g *GXDLMSSpecialDaysTable) GetValues() []any {
	return []any{g.LogicalName(), g.Entries}
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
func (g *GXDLMSSpecialDaysTable) GetDataType(index int) (enums.DataType, error) {
	if index == 1 {
		return enums.DataTypeOctetString, nil
	}
	if index == 2 {
		return enums.DataTypeArray, nil
	}
	return 0, dlmserrors.ErrInvalidAttributeIndex
}

// NewGXDLMSSpecialDaysTable creates a new special days table object instance.
//
// The function validates `ln` before creating the object.
//`ln` is the Logical Name and `sn` is the Short Name of the object.
func NewGXDLMSSpecialDaysTable(ln string, sn int16) (*GXDLMSSpecialDaysTable, error) {
	err := ValidateLogicalName(ln)
	if err != nil {
		return nil, err
	}
	return &GXDLMSSpecialDaysTable{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeSpecialDaysTable,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
