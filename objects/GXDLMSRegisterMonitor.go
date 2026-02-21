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
	"github.com/Gurux/gxdlms-go/dlmserrors"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// Online help:
// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSRegisterMonitor
type GXDLMSRegisterMonitor struct {
	GXDLMSObject
	Thresholds []any

	MonitoredValue GXDLMSMonitoredValue

	Actions []GXDLMSActionSet
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSRegisterMonitor) Base() *GXDLMSObject {
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
func (g *GXDLMSRegisterMonitor) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	// LN is static and read only once.
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	// Thresholds
	if all || !g.IsRead(2) {
		attributes = append(attributes, 2)
	}
	// MonitoredValue
	if all || !g.IsRead(3) {
		attributes = append(attributes, 3)
	}
	// Actions
	if all || !g.IsRead(4) {
		attributes = append(attributes, 4)
	}
	return attributes
}

// GetNames returns the names of attribute indexes.
func (g *GXDLMSRegisterMonitor) GetNames() []string {
	return []string{"Logical Name", "Thresholds", "Monitored Value", "Actions"}
}

// GetMethodNames returns the names of method indexes.
func (g *GXDLMSRegisterMonitor) GetMethodNames() []string {
	return []string{}
}

// GetAttributeCount returns the amount of attributes.
//
// Returns:
//
//	Count of attributes.
func (g *GXDLMSRegisterMonitor) GetAttributeCount() int {
	return 4
}

// GetMethodCount returns the amount of methods.
func (g *GXDLMSRegisterMonitor) GetMethodCount() int {
	return 0
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
func (g *GXDLMSRegisterMonitor) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	var err error
	if e.Index == 1 {
		v, err := helpers.LogicalNameToBytes(g.LogicalName())
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return v, err
	}
	if e.Index == 2 {
		return g.Thresholds, nil
	}
	if e.Index == 3 {
		data := types.NewGXByteBuffer()
		err = data.SetUint8(byte(enums.DataTypeStructure))
		if err != nil {
			return nil, err
		}
		err = data.SetUint8(3)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, data, enums.DataTypeUint16, g.MonitoredValue.ObjectType)
		ln, err := helpers.LogicalNameToBytes(g.MonitoredValue.LogicalName)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, data, enums.DataTypeOctetString, ln)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, data, enums.DataTypeInt8, g.MonitoredValue.AttributeIndex)
		if err != nil {
			return nil, err
		}
		return data.Array(), nil
	}
	if e.Index == 4 {
		data := types.NewGXByteBuffer()
		err = data.SetUint8(byte(enums.DataTypeArray))
		if err != nil {
			return nil, err
		}
		if g.Actions == nil {
			err = data.SetUint8(0)
			if err != nil {
				return nil, err
			}
		} else {
			data.SetUint8(byte(len(g.Actions)))
			for _, it := range g.Actions {
				err = data.SetUint8(byte(enums.DataTypeStructure))
				if err != nil {
					return nil, err
				}
				err = data.SetUint8(2)
				if err != nil {
					return nil, err
				}
				err = data.SetUint8(byte(enums.DataTypeStructure))
				if err != nil {
					return nil, err
				}
				err = data.SetUint8(2)
				if err != nil {
					return nil, err
				}
				ln, err := helpers.LogicalNameToBytes(it.ActionUp.LogicalName)
				if err != nil {
					return nil, err
				}
				err = internal.SetData(settings, data, enums.DataTypeOctetString, ln)
				if err != nil {
					return nil, err
				}
				err = internal.SetData(settings, data, enums.DataTypeUint16, it.ActionUp.ScriptSelector)
				if err != nil {
					return nil, err
				}
				err = data.SetUint8(byte(enums.DataTypeStructure))
				if err != nil {
					return nil, err
				}
				err = data.SetUint8(2)
				if err != nil {
					return nil, err
				}
				ln, err = helpers.LogicalNameToBytes(it.ActionDown.LogicalName)
				if err != nil {
					return nil, err
				}
				err = internal.SetData(settings, data, enums.DataTypeOctetString, ln)
				if err != nil {
					return nil, err
				}
				err = internal.SetData(settings, data, enums.DataTypeUint16, it.ActionDown.ScriptSelector)
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

// SetValue returns the set value of given attribute.
// When raw parameter us not used example register multiplies value by scalar.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Set parameters.
func (g *GXDLMSRegisterMonitor) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	var err error
	if e.Index == 1 {
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return g.SetLogicalName(ln)
	} else if e.Index == 2 {
		if e.Value != nil {
			g.Thresholds = e.Value.(types.GXArray)
		} else {
			g.Thresholds = g.Thresholds[:0]
		}
	} else if e.Index == 3 {
		if e.Value != nil {
			arr := e.Value.(types.GXStructure)
			g.MonitoredValue.ObjectType = enums.ObjectType(arr[0].(uint16))
			g.MonitoredValue.LogicalName, err = helpers.ToLogicalName(arr[1])
			if err != nil {
				return err
			}
			g.MonitoredValue.AttributeIndex = arr[2].(int8)
		} else {
			g.MonitoredValue.ObjectType = enums.ObjectTypeNone
			g.MonitoredValue.LogicalName = ""
			g.MonitoredValue.AttributeIndex = 0
		}
	} else if e.Index == 4 {
		g.Actions = g.Actions[:0]
		if e.Value != nil {
			items := []GXDLMSActionSet{}
			for _, tmp := range e.Value.(types.GXArray) {
				action_set := tmp.(types.GXStructure)
				it := action_set[0].(types.GXStructure)
				set := GXDLMSActionSet{}
				set.ActionUp.LogicalName, err = helpers.ToLogicalName(it[0])
				if err != nil {
					return err
				}
				set.ActionUp.ScriptSelector = it[1].(uint16)
				it = action_set[1].(types.GXStructure)
				set.ActionDown.LogicalName, err = (helpers.ToLogicalName(it[0]))
				if err != nil {
					return err
				}
				set.ActionDown.ScriptSelector = it[1].(uint16)
				items = append(items, set)
			}
			g.Actions = items
		}
	} else {
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil
}

// Invoke returns the invokes method.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Invoke parameters.
func (g *GXDLMSRegisterMonitor) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, nil
}

// Load returns the load object content from XML.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSRegisterMonitor) Load(reader *GXXmlReader) error {
	var err error
	g.Thresholds = []any{}
	if ret, err := reader.IsStartElementNamed("Thresholds", true); ret && err == nil {
		for {
			ret, err = reader.IsStartElementNamed("Value", true)
			if err != nil {
				return err
			}
			if !ret {
				break
			}
			//TODO: it := reader.ReadElementContentAsObject("Value", nil, nil, 0)
			//TODO: g.Thresholds = append(g.Thresholds, it)
		}
		reader.ReadEndElement("Thresholds")
	}
	if ret, err := reader.IsStartElementNamed("MonitoredValue", true); ret && err == nil {
		ret, err := reader.ReadElementContentAsInt("ObjectType", 0)
		if err != nil {
			return err
		}
		g.MonitoredValue.ObjectType = enums.ObjectType(ret)
		if err != nil {
			return err
		}
		g.MonitoredValue.LogicalName, err = reader.ReadElementContentAsString("LN", "0.0.0.0.0.0")
		if err != nil {
			return err
		}
		g.MonitoredValue.AttributeIndex, err = reader.ReadElementContentAsInt8("Index", 0)
		if err != nil {
			return err
		}
		reader.ReadEndElement("MonitoredValue")
	}
	g.Actions = g.Actions[:0]
	if ret, err := reader.IsStartElementNamed("Actions", true); ret && err == nil {
		for {
			ret, err = reader.IsStartElementNamed("Item", true)
			if err != nil {
				return err
			}
			if !ret {
				break
			}
			it := GXDLMSActionSet{}
			g.Actions = append(g.Actions, it)
			if ret, err := reader.IsStartElementNamed("Up", true); ret && err == nil {
				it.ActionUp.LogicalName, err = reader.ReadElementContentAsString("LN", "0.0.0.0.0.0")
				if err != nil {
					return err
				}
				if it.ActionUp.LogicalName == "" {
					it.ActionUp.LogicalName = "0.0.0.0.0.0"
				}
				it.ActionUp.ScriptSelector, err = reader.ReadElementContentAsUInt16("Selector", 0)
				if err != nil {
					return err
				}
				reader.ReadEndElement("Up")
			}
			if ret, err := reader.IsStartElementNamed("Down", true); ret && err == nil {
				it.ActionDown.LogicalName, err = reader.ReadElementContentAsString("LN", "0.0.0.0.0.0")
				if err != nil {
					return err
				}
				if it.ActionDown.LogicalName == "" {
					it.ActionDown.LogicalName = "0.0.0.0.0.0"
				}
				it.ActionDown.ScriptSelector, err = reader.ReadElementContentAsUInt16("Selector", 0)
				if err != nil {
					return err
				}
				reader.ReadEndElement("Down")
			}
		}
		reader.ReadEndElement("Actions")
	}
	return err
}

// Save returns the save object content to XML.
//
// Parameters:
//
//	writer: XML writer.
func (g *GXDLMSRegisterMonitor) Save(writer *GXXmlWriter) error {
	var err error
	if g.Thresholds != nil {
		writer.WriteStartElement("Thresholds")
		for _, it := range g.Thresholds {
			err = writer.WriteElementObject("Value", it, enums.DataTypeNone, enums.DataTypeNone)
			if err != nil {
				return err
			}
		}
		writer.WriteEndElement()
	}
	writer.WriteStartElement("MonitoredValue")
	err = writer.WriteElementString("ObjectType", int(g.MonitoredValue.ObjectType))
	if err != nil {
		return err
	}
	err = writer.WriteElementString("LN", g.MonitoredValue.LogicalName)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("Index", g.MonitoredValue.AttributeIndex)
	if err != nil {
		return err
	}
	writer.WriteEndElement()
	if g.Actions != nil {
		writer.WriteStartElement("Actions")
		for _, it := range g.Actions {
			writer.WriteStartElement("Item")
			writer.WriteStartElement("Up")
			err = writer.WriteElementString("LN", it.ActionUp.LogicalName)
			if err != nil {
				return err
			}
			err = writer.WriteElementString("Selector", it.ActionUp.ScriptSelector)
			if err != nil {
				return err
			}
			writer.WriteEndElement()
			writer.WriteStartElement("Down")
			err = writer.WriteElementString("LN", it.ActionDown.LogicalName)
			if err != nil {
				return err
			}
			err = writer.WriteElementString("Selector", it.ActionDown.ScriptSelector)
			if err != nil {
				return err
			}
			writer.WriteEndElement()
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
func (g *GXDLMSRegisterMonitor) PostLoad(reader *GXXmlReader) error {
	return nil
}

// GetValues returns the an array containing the COSEM object's attribute values.
func (g *GXDLMSRegisterMonitor) GetValues() []any {
	return []any{g.LogicalName(), g.Thresholds, g.MonitoredValue, g.Actions}
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
func (g *GXDLMSRegisterMonitor) GetDataType(index int) (enums.DataType, error) {
	if index == 1 {
		return enums.DataTypeOctetString, nil
	}
	if index == 2 {
		return enums.DataTypeArray, nil
	}
	if index == 3 {
		return enums.DataTypeArray, nil
	}
	if index == 4 {
		return enums.DataTypeArray, nil
	}
	return 0, dlmserrors.ErrInvalidAttributeIndex
}

// NewGXDLMSRegisterMonitor creates a new register monitor object instance.
//
// The function validates `ln` before creating the object.
//`ln` is the Logical Name and `sn` is the Short Name of the object.
func NewGXDLMSRegisterMonitor(ln string, sn int16) (*GXDLMSRegisterMonitor, error) {
	err := ValidateLogicalName(ln)
	if err != nil {
		return nil, err
	}
	return &GXDLMSRegisterMonitor{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeRegisterMonitor,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
