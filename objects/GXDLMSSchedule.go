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
// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSSchedule
type GXDLMSSchedule struct {
	GXDLMSObject
	// Specifies the scripts to be executed at given times.
	Entries []GXScheduleEntry
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSSchedule) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

func (g *GXDLMSSchedule) removeEntry(index uint16) {
	for _, it := range g.Entries {
		if it.Index == index {
			internal.Remove(g.Entries, it)
			break
		}
	}
}

// Invoke returns the invokes method.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Invoke parameters.
func (g *GXDLMSSchedule) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	switch e.Index {
	case 1:

		tmp := e.Parameters.(types.GXArray)
		for index := tmp[0].(uint16); index <= tmp[1].(uint16); index++ {
			if index != 0 {
				for _, it := range g.Entries {
					if it.Index == index {
						it.Enable = false
					}
				}
			}
		}
		for index := tmp[2].(uint16); index <= tmp[3].(uint16); index++ {
			if index != 0 {
				for _, it := range g.Entries {
					if it.Index == index {
						it.Enable = false
					}
				}
			}
		}

	case 2:
		entry, err := g.createEntry(settings, e.Parameters.(types.GXStructure))
		if err != nil {
			return nil, err
		}
		g.removeEntry(entry.Index)
		g.Entries = append(g.Entries, *entry)
	case 3:
		tmp := e.Parameters.(types.GXArray)
		for index := tmp[0].(uint16); index <= tmp[1].(uint16); index++ {
			g.removeEntry(index)
		}
	default:
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
func (g *GXDLMSSchedule) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	// LN is static and read only once.
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	// Entries
	if all || g.CanRead(2) {
		attributes = append(attributes, 2)
	}
	return attributes
}

// GetNames returns the names of attribute indexes.
func (g *GXDLMSSchedule) GetNames() []string {
	return []string{"Logical Name", "Entries"}
}

// GetMethodNames returns the names of method indexes.
func (g *GXDLMSSchedule) GetMethodNames() []string {
	return []string{"Enable/disable", "Insert", "Delete"}
}

// GetAttributeCount returns the amount of attributes.
//
// Returns:
//
//	Count of attributes.
func (g *GXDLMSSchedule) GetAttributeCount() int {
	return 2
}

// GetMethodCount returns the amount of methods.
func (g *GXDLMSSchedule) GetMethodCount() int {
	return 3
}

func (g *GXDLMSSchedule) addEntry(settings *settings.GXDLMSSettings,
	it *GXScheduleEntry,
	data *types.GXByteBuffer) error {
	err := data.SetUint8(uint8(enums.DataTypeStructure))
	if err != nil {
		return err
	}
	err = data.SetUint8(10)
	if err != nil {
		return err
	}
	err = data.SetUint8(uint8(enums.DataTypeUint16))
	if err != nil {
		return err
	}
	err = data.SetUint16(it.Index)
	if err != nil {
		return err
	}
	err = data.SetUint8(uint8(enums.DataTypeBoolean))
	if err != nil {
		return err
	}
	v := uint8(0)
	if it.Enable {
		v = 1
	}
	err = data.SetUint8(v)
	if err != nil {
		return err
	}
	err = data.SetUint8(uint8(enums.DataTypeOctetString))
	if err != nil {
		return err
	}
	err = data.SetUint8(6)
	if err != nil {
		return err
	}
	if it.Script == nil {
		err = data.Set([]byte{0, 0, 0, 0, 0, 0})
		if err != nil {
			return err
		}
	} else {
		ret, err := helpers.LogicalNameToBytes(it.Script.Base().LogicalName())
		if err != nil {
			return err
		}
		err = data.Set(ret)
		if err != nil {
			return err
		}
	}
	err = data.SetUint8(uint8(enums.DataTypeUint16))
	if err != nil {
		return err
	}
	err = data.SetUint16(it.ScriptSelector)
	if err != nil {
		return err
	}
	err = internal.SetData(settings, data, enums.DataTypeOctetString, it.SwitchTime)
	err = data.SetUint8(uint8(enums.DataTypeUint16))
	if err != nil {
		return err
	}
	err = data.SetUint16(it.ValidityWindow)
	if err != nil {
		return err
	}
	ret, err := types.NewGXBitStringFromInteger(int(it.ExecWeekdays), 7)
	if err != nil {
		return err
	}
	err = internal.SetData(settings, data, enums.DataTypeBitString, ret)
	if err != nil {
		return err
	}
	err = internal.SetData(settings, data, enums.DataTypeBitString, it.ExecSpecDays)
	if err != nil {
		return err
	}
	err = internal.SetData(settings, data, enums.DataTypeOctetString, it.BeginDate)
	if err != nil {
		return err
	}
	err = internal.SetData(settings, data, enums.DataTypeOctetString, it.EndDate)
	return err
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
func (g *GXDLMSSchedule) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
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
		types.SetObjectCount(len(g.Entries), data)
		for _, it := range g.Entries {
			g.addEntry(settings, &it, data)
		}
		return data.Array(), nil
	}
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, nil
}

// createEntry returns the create a new entry.
func (g *GXDLMSSchedule) createEntry(settings *settings.GXDLMSSettings, it types.GXStructure) (*GXScheduleEntry, error) {
	item := GXScheduleEntry{}
	item.Index = it[0].(uint16)
	item.Enable = it[1].(bool)
	ln, err := helpers.ToLogicalName(it[2])
	if err != nil {
		return nil, err
	}
	if settings != nil && ln != "0.0.0.0.0.0" {
		obj := getObjectCollection(settings.Objects).FindByLN(enums.ObjectTypeScriptTable, ln)
		item.Script = obj.(*GXDLMSScriptTable)
	}
	if item.Script == nil {
		item.Script, err = NewGXDLMSScriptTable(ln, 0)
		if err != nil {
			return nil, err
		}
	}
	item.ScriptSelector = it[3].(uint16)
	ret, err := internal.ChangeTypeFromByteArray(settings, it[4].([]byte), enums.DataTypeTime)
	if err != nil {
		return nil, err
	}
	item.SwitchTime = ret.(types.GXTime)
	item.ValidityWindow = it[5].(uint16)
	bs := it[6].(types.GXBitString)
	item.ExecWeekdays = enums.Weekdays(bs.ToInteger())
	item.ExecSpecDays = (fmt.Sprint(it[7]))
	ret, err = internal.ChangeTypeFromByteArray(settings, it[8].([]byte), enums.DataTypeDate)
	if err != nil {
		return nil, err
	}
	item.BeginDate = ret.(types.GXDate)
	ret, err = internal.ChangeTypeFromByteArray(settings, it[9].([]byte), enums.DataTypeDate)
	if err != nil {
		return nil, err
	}
	item.EndDate = ret.(types.GXDate)
	return &item, nil
}

// SetValue returns the set value of given attribute.
// When raw parameter us not used example register multiplies value by scalar.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Set parameters.
func (g *GXDLMSSchedule) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	if e.Index == 1 {
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return g.SetLogicalName(ln)
	} else if e.Index == 2 {
		g.Entries = g.Entries[:0]
		arr := e.Value.(types.GXArray)
		if arr != nil {
			for _, it := range arr {
				item, err := g.createEntry(settings, it.(types.GXStructure))
				if err != nil {
					return err
				}
				g.Entries = append(g.Entries, *item)
			}
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
func (g *GXDLMSSchedule) Load(reader *GXXmlReader) error {
	var err error
	g.Entries = g.Entries[:0]
	if ret, err := reader.IsStartElementNamed("Entries", true); ret && err == nil {
		for {
			ret, err = reader.IsStartElementNamed("Item", true)
			if err != nil {
				return err
			}
			if !ret {
				break
			}
			it := GXScheduleEntry{}
			it.Index, err = reader.ReadElementContentAsUInt16("Index", 0)
			if err != nil {
				return err
			}
			it.Enable, err = reader.ReadElementContentAsBool("Enable", true)
			if err != nil {
				return err
			}
			ln, err := reader.ReadElementContentAsString("LogicalName", "")
			if err != nil {
				return err
			}
			if ln != "" {
				it.Script, err = NewGXDLMSScriptTable(ln, 0)
				if err != nil {
					return err
				}
			}
			it.ScriptSelector, err = reader.ReadElementContentAsUInt16("ScriptSelector", 0)
			if err != nil {
				return err
			}
			it.SwitchTime, err = reader.ReadElementContentAsTime("SwitchTime")
			it.ValidityWindow, err = reader.ReadElementContentAsUInt16("ValidityWindow", 0)
			if err != nil {
				return err
			}
			ret, err := reader.ReadElementContentAsInt("ExecWeekdays", 0)
			if err != nil {
				return err
			}
			it.ExecWeekdays = enums.Weekdays(ret)
			it.ExecSpecDays, err = reader.ReadElementContentAsString("ExecSpecDays", "")
			if err != nil {
				return err
			}
			it.BeginDate, err = reader.ReadElementContentAsGXDate("BeginDate")
			if err != nil {
				return err
			}
			it.EndDate, err = reader.ReadElementContentAsGXDate("EndDate")
			if err != nil {
				return err
			}
			g.Entries = append(g.Entries, it)
		}
		reader.ReadEndElement("Entries")
	}
	return err
}

// Save returns the save object content to XML.
//
// Parameters:
//
//	writer: XML writer.
func (g *GXDLMSSchedule) Save(writer *GXXmlWriter) error {
	var err error
	if g.Entries != nil {
		writer.WriteStartElement("Entries")
		for _, it := range g.Entries {
			writer.WriteStartElement("Item")
			err = writer.WriteElementString("Index", it.Index)
			if err != nil {
				return err
			}
			err = writer.WriteElementString("Enable", it.Enable)
			if err != nil {
				return err
			}
			if it.Script != nil {
				err = writer.WriteElementString("LogicalName", it.Script.Base().LogicalName())
				if err != nil {
					return err
				}
			}
			err = writer.WriteElementString("ScriptSelector", it.ScriptSelector)
			if err != nil {
				return err
			}
			err = writer.WriteElementString("SwitchTime", it.SwitchTime)
			if err != nil {
				return err
			}
			err = writer.WriteElementString("ValidityWindow", it.ValidityWindow)
			if err != nil {
				return err
			}
			err = writer.WriteElementString("ExecWeekdays", int(it.ExecWeekdays))
			if err != nil {
				return err
			}
			err = writer.WriteElementString("ExecSpecDays", it.ExecSpecDays)
			if err != nil {
				return err
			}
			err = writer.WriteElementString("BeginDate", it.BeginDate)
			if err != nil {
				return err
			}
			err = writer.WriteElementString("EndDate", it.EndDate)
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
func (g *GXDLMSSchedule) PostLoad(reader *GXXmlReader) error {
	// Upload entries Value after load.
	if g.Entries != nil {
		for _, it := range g.Entries {
			target := reader.Objects.FindByLN(enums.ObjectTypeScriptTable, it.Script.Base().LogicalName())
			if target != nil && target != it.Script {
				it.Script = target.(*GXDLMSScriptTable)
			}
		}
	}
	return nil
}

// GetValues returns the an array containing the COSEM object's attribute values.
func (g *GXDLMSSchedule) GetValues() []any {
	return []any{g.LogicalName(), g.Entries}
}

// Insert returns the add entry to entries list.
//
// Parameters:
//
//	client: DLMS client.
//	entry: Schedule entry.
//
// Returns:
//
//	Action bytes.
func (g *GXDLMSSchedule) Insert(client IGXDLMSClient, entry *GXScheduleEntry) ([][]uint8, error) {
	data := types.NewGXByteBuffer()
	g.addEntry(client.Settings().(*settings.GXDLMSSettings), entry, data)
	return client.Method(g, 2, data.Array(), enums.DataTypeStructure)
}

// Delete returns the remove entry from entries list.
//
// Parameters:
//
//	client: DLMS client.
//	entry: Schedule entry.
//
// Returns:
//
//	Action bytes.
func (g *GXDLMSSchedule) Delete(client IGXDLMSClient, entry *GXScheduleEntry) ([][]uint8, error) {
	data := types.NewGXByteBuffer()
	err := data.SetUint8(uint8(enums.DataTypeStructure))
	if err != nil {
		return nil, err
	}
	err = data.SetUint8(2)
	if err != nil {
		return nil, err
	}
	//firstIndex
	err = internal.SetData(nil, data, enums.DataTypeUint16, entry.Index)
	//lastIndex
	err = internal.SetData(nil, data, enums.DataTypeUint16, entry.Index)
	return client.Method(g, 3, data.Array(), enums.DataTypeStructure)
}

// Enable returns the enable entry from entries list.
//
// Parameters:
//
//	client: DLMS client.
//	entry: Schedule entries.
//
// Returns:
//
//	Action bytes.
func (g *GXDLMSSchedule) Enable(client IGXDLMSClient, entry *GXScheduleEntry) ([][]uint8, error) {
	data := types.NewGXByteBuffer()
	err := data.SetUint8(uint8(enums.DataTypeStructure))
	if err != nil {
		return nil, err
	}
	err = data.SetUint8(4)
	if err != nil {
		return nil, err
	}
	//firstIndex
	err = internal.SetData(nil, data, enums.DataTypeUint16, 0)
	if err != nil {
		return nil, err
	}
	err = internal.SetData(nil, data, enums.DataTypeUint16, 0)
	if err != nil {
		return nil, err
	}
	err = internal.SetData(nil, data, enums.DataTypeUint16, entry.Index)
	if err != nil {
		return nil, err
	}
	//lastIndex
	err = internal.SetData(nil, data, enums.DataTypeUint16, entry.Index)
	if err != nil {
		return nil, err
	}
	return client.Method(g, 1, data.Array(), enums.DataTypeStructure)
}

// Disable returns the disable entry from entries list.
//
// Parameters:
//
//	client: DLMS client.
//	entry: Schedule entries.
//
// Returns:
//
//	Action bytes.
func (g *GXDLMSSchedule) Disable(client IGXDLMSClient, entry *GXScheduleEntry) ([][]uint8, error) {
	data := types.NewGXByteBuffer()
	err := data.SetUint8(uint8(enums.DataTypeStructure))
	if err != nil {
		return nil, err
	}
	err = data.SetUint8(4)
	if err != nil {
		return nil, err
	}
	err = internal.SetData(nil, data, enums.DataTypeUint16, entry.Index)
	if err != nil {
		return nil, err
	}
	err = internal.SetData(nil, data, enums.DataTypeUint16, entry.Index)
	if err != nil {
		return nil, err
	}
	err = internal.SetData(nil, data, enums.DataTypeUint16, 0)
	if err != nil {
		return nil, err
	}
	err = internal.SetData(nil, data, enums.DataTypeUint16, 0)
	if err != nil {
		return nil, err
	}
	return client.Method(g, 1, data.Array(), enums.DataTypeStructure)
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
func (g *GXDLMSSchedule) GetDataType(index int) (enums.DataType, error) {
	if index == 1 {
		return enums.DataTypeOctetString, nil
	}
	if index == 2 {
		return enums.DataTypeArray, nil
	}
	return 0, dlmserrors.ErrInvalidAttributeIndex
}

// NewGXDLMSSchedule creates a new schedule object instance.
//
// The function validates `ln` before creating the object.
//`ln` is the Logical Name and `sn` is the Short Name of the object.
func NewGXDLMSSchedule(ln string, sn int16) (*GXDLMSSchedule, error) {
	err := ValidateLogicalName(ln)
	if err != nil {
		return nil, err
	}
	return &GXDLMSSchedule{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeSchedule,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
