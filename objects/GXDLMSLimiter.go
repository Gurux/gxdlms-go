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
// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSLimiter
type GXDLMSLimiter struct {
	GXDLMSObject
	// Defines an attribute of an object to be monitored.
	MonitoredValue IGXDLMSBase

	// Attribute index of monitored value.
	MonitoredAttributeIndex int8

	// Provides the active threshold value to which the attribute monitored is compared.
	ThresholdActive any

	// Provides the threshold value to which the attribute monitored
	// is compared when in normal operation.
	ThresholdNormal any

	// Provides the threshold value to which the attribute monitored
	// is compared when an emergency profile is active.
	ThresholdEmergency any

	// Defines minimal over threshold duration in seconds required
	// to execute the over threshold action.
	MinOverThresholdDuration uint32

	// Defines minimal under threshold duration in seconds required to
	// execute the under threshold action.
	MinUnderThresholdDuration uint32

	EmergencyProfile GXDLMSEmergencyProfile

	EmergencyProfileGroupIDs []uint16

	// Is Emergency Profile active.
	EmergencyProfileActive bool

	// Defines the scripts to be executed when the monitored value
	// crosses the threshold for minimal duration time.
	ActionOverThreshold GXDLMSActionItem

	// Defines the scripts to be executed when the monitored value
	// crosses the threshold for minimal duration time.
	ActionUnderThreshold GXDLMSActionItem
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSLimiter) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

// Invoke returns the invokes method.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Invoke parameters.
func (g *GXDLMSLimiter) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	e.Error = enums.ErrorCodeReadWriteDenied
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
func (g *GXDLMSLimiter) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	// LN is static and read only once.
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	// MonitoredValue
	if all || g.CanRead(2) {
		attributes = append(attributes, 2)
	}
	// ThresholdActive
	if all || g.CanRead(3) {
		attributes = append(attributes, 3)
	}
	// ThresholdNormal
	if all || g.CanRead(4) {
		attributes = append(attributes, 4)
	}
	// ThresholdEmergency
	if all || g.CanRead(5) {
		attributes = append(attributes, 5)
	}
	// MinOverThresholdDuration
	if all || g.CanRead(6) {
		attributes = append(attributes, 6)
	}
	// MinUnderThresholdDuration
	if all || g.CanRead(7) {
		attributes = append(attributes, 7)
	}
	// EmergencyProfile
	if all || g.CanRead(8) {
		attributes = append(attributes, 8)
	}
	// EmergencyProfileGroup
	if all || g.CanRead(9) {
		attributes = append(attributes, 9)
	}
	// EmergencyProfileActive
	if all || g.CanRead(10) {
		attributes = append(attributes, 10)
	}
	// Actions
	if all || g.CanRead(11) {
		attributes = append(attributes, 11)
	}
	return attributes
}

// GetNames returns the names of attribute indexes.
func (g *GXDLMSLimiter) GetNames() []string {
	return []string{"Logical Name", "Monitored Value", "Active Threshold", "Normal Threshold",
		"Emergency Threshold", "Threshold Duration Min Over",
		"Threshold Duration Min Under", "Emergency Profile",
		"Emergency Profile Group", "Emergency Profile Active", "Actions"}
}

// GetMethodNames returns the names of method indexes.
func (g *GXDLMSLimiter) GetMethodNames() []string {
	return []string{}
}

// GetAttributeCount returns the amount of attributes.
//
// Returns:
//
//	Count of attributes.
func (g *GXDLMSLimiter) GetAttributeCount() int {
	return 11
}

// GetMethodCount returns the amount of methods.
func (g *GXDLMSLimiter) GetMethodCount() int {
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
func (g *GXDLMSLimiter) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	var err error
	var ret any
	if e.Index == 1 {
		ret, err = helpers.LogicalNameToBytes(g.LogicalName())
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
	} else if e.Index == 2 {
		data := types.NewGXByteBuffer()
		err = data.SetUint8(uint8(enums.DataTypeStructure))
		if err != nil {
			return nil, err
		}
		err = data.SetUint8(3)
		if err != nil {
			return nil, err
		}
		if g.MonitoredValue == nil {
			err = internal.SetData(settings, data, enums.DataTypeUint16, 0)
			if err != nil {
				return nil, err
			}
			ln, err := helpers.LogicalNameToBytes("")
			err = internal.SetData(settings, data, enums.DataTypeOctetString, ln)
			if err != nil {
				return nil, err
			}
			err = internal.SetData(settings, data, enums.DataTypeInt8, 0)
			if err != nil {
				return nil, err
			}
		} else {
			err = internal.SetData(settings, data, enums.DataTypeUint16, g.MonitoredValue.Base().ObjectType)
			if err != nil {
				return nil, err
			}
			ln, err := helpers.LogicalNameToBytes(g.MonitoredValue.Base().LogicalName())
			err = internal.SetData(settings, data, enums.DataTypeOctetString, ln)
			if err != nil {
				return nil, err
			}
			err = internal.SetData(settings, data, enums.DataTypeInt8, g.MonitoredAttributeIndex)
			if err != nil {
				return nil, err
			}
		}
		ret = data.Array()
	} else if e.Index == 3 {
		ret = g.ThresholdActive
	} else if e.Index == 4 {
		ret = g.ThresholdNormal
	} else if e.Index == 5 {
		ret = g.ThresholdEmergency
	} else if e.Index == 6 {
		ret = g.MinOverThresholdDuration
	} else if e.Index == 7 {
		ret = g.MinUnderThresholdDuration
	} else if e.Index == 8 {
		data := types.NewGXByteBuffer()
		err = data.SetUint8(uint8(enums.DataTypeStructure))
		if err != nil {
			return nil, err
		}
		err = data.SetUint8(3)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, data, enums.DataTypeUint16, g.EmergencyProfile.ID)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, data, enums.DataTypeOctetString, g.EmergencyProfile.ActivationTime)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, data, enums.DataTypeUint32, g.EmergencyProfile.Duration)
		if err != nil {
			return nil, err
		}
		ret = data.Array()
	} else if e.Index == 9 {
		data := types.NewGXByteBuffer()
		err = data.SetUint8(uint8(enums.DataTypeArray))
		if err != nil {
			return nil, err
		}
		if g.EmergencyProfileGroupIDs == nil {
			err = data.SetUint8(0)
			if err != nil {
				return nil, err
			}
		} else {
			err = data.SetUint8(byte(len(g.EmergencyProfileGroupIDs)))
			if err != nil {
				return nil, err
			}
			for _, it := range g.EmergencyProfileGroupIDs {
				err = internal.SetData(settings, data, enums.DataTypeUint16, it)
				if err != nil {
					return nil, err
				}
			}
		}
		ret = data.Array()
	} else if e.Index == 10 {
		ret = g.EmergencyProfileActive
	} else if e.Index == 11 {
		data := types.NewGXByteBuffer()
		err = data.SetUint8(uint8(enums.DataTypeStructure))
		if err != nil {
			return nil, err
		}
		err = data.SetUint8(2)
		if err != nil {
			return nil, err
		}
		err = data.SetUint8(uint8(enums.DataTypeStructure))
		if err != nil {
			return nil, err
		}
		err = data.SetUint8(2)
		if err != nil {
			return nil, err
		}
		ln, err := helpers.LogicalNameToBytes(g.ActionOverThreshold.LogicalName)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, data, enums.DataTypeOctetString, ln)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, data, enums.DataTypeUint16, g.ActionOverThreshold.ScriptSelector)
		if err != nil {
			return nil, err
		}
		err = data.SetUint8(uint8(enums.DataTypeStructure))
		if err != nil {
			return nil, err
		}
		err = data.SetUint8(2)
		if err != nil {
			return nil, err
		}
		ln, err = helpers.LogicalNameToBytes(g.ActionUnderThreshold.LogicalName)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, data, enums.DataTypeOctetString, ln)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, data, enums.DataTypeUint16, g.ActionUnderThreshold.ScriptSelector)
		if err != nil {
			return nil, err
		}
		ret = data.Array()
		if err != nil {
			return nil, err
		}
		ret = data.Array()
	} else {
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return ret, err
}

// SetValue returns the set value of given attribute.
// When raw parameter us not used example register multiplies value by scalar.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Set parameters.
func (g *GXDLMSLimiter) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	if e.Index == 1 {
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		err = g.SetLogicalName(ln)
	} else if e.Index == 2 {
		tmp := e.Value.(types.GXStructure)
		ot := enums.ObjectType(tmp[0].(uint16))
		ln, err := helpers.ToLogicalName(tmp[1])
		if err != nil {
			return err
		}
		attIndex := tmp[2].(int8)
		if ot != enums.ObjectTypeNone {
			g.MonitoredValue = getObjectCollection(settings.Objects).FindByLN(ot, ln)
			if g.MonitoredValue == nil {
				g.MonitoredValue, err = CreateObject(ot, ln, 0)
				if err != nil {
					return err
				}
			}
		} else {
			g.MonitoredValue = nil
		}
		g.MonitoredAttributeIndex = attIndex
		if g.MonitoredValue != nil && attIndex != 0 {
			dt, err := g.MonitoredValue.GetDataType(int(attIndex))
			if err != nil {
				return err
			}
			g.SetDataType(3, dt)
			g.SetDataType(4, dt)
			g.SetDataType(5, dt)
			g.SetDataType(6, dt)
			g.SetDataType(7, dt)
		}
	} else if e.Index == 3 {
		g.ThresholdActive = e.Value
	} else if e.Index == 4 {
		g.ThresholdNormal = e.Value
	} else if e.Index == 5 {
		g.ThresholdEmergency = e.Value
	} else if e.Index == 6 {
		g.MinOverThresholdDuration = e.Value.(uint32)
	} else if e.Index == 7 {
		g.MinUnderThresholdDuration = e.Value.(uint32)
	} else if e.Index == 8 {
		tmp := e.Value.(types.GXStructure)
		g.EmergencyProfile.ID = tmp[0].(uint16)
		ret, err := internal.ChangeTypeFromByteArray(settings, tmp[1].([]byte), enums.DataTypeDateTime)
		if err != nil {
			return err
		}
		g.EmergencyProfile.ActivationTime = ret.(types.GXDateTime)
		g.EmergencyProfile.Duration = tmp[2].(uint32)
	} else if e.Index == 9 {
		var list []uint16
		if e.Value != nil {
			for _, it := range e.Value.(types.GXArray) {
				list = append(list, it.(uint16))
			}
		}
		g.EmergencyProfileGroupIDs = list
	} else if e.Index == 10 {
		g.EmergencyProfileActive = e.Value.(bool)
	} else if e.Index == 11 {
		tmp := e.Value.(types.GXStructure)
		tmp1 := tmp[0].(types.GXStructure)
		tmp2 := tmp[1].(types.GXStructure)
		ln, err := helpers.ToLogicalName(tmp1[0])
		if err != nil {
			return err
		}
		g.ActionOverThreshold.LogicalName = ln
		g.ActionOverThreshold.ScriptSelector = tmp1[1].(uint16)
		ln, err = helpers.ToLogicalName(tmp2[0])
		if err != nil {
			return err
		}
		g.ActionUnderThreshold.LogicalName = ln
		g.ActionUnderThreshold.ScriptSelector = tmp2[1].(uint16)
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
func (g *GXDLMSLimiter) Load(reader *GXXmlReader) error {
	b, err := reader.IsStartElementNamed("MonitoredValue", true)
	if b {
		ret, err := reader.ReadElementContentAsInt("ObjectType", 0)
		if err != nil {
			return err
		}
		ot := enums.ObjectType(ret)
		ln, err := reader.ReadElementContentAsString("LN", "")
		ret, err = reader.ReadElementContentAsInt("Index", 0)
		if err != nil {
			return err
		}
		g.MonitoredAttributeIndex = int8(ret)
		if ot != enums.ObjectTypeNone && ln != "" {
			g.MonitoredValue = reader.Objects.FindByLN(ot, ln)
			// If item is not serialized yet.
			if g.MonitoredValue == nil {
				g.MonitoredValue, err = CreateObject(ot, ln, 0)
				if err != nil {
					return err
				}
			}
		}
		reader.ReadEndElement("MonitoredValue")
	}
	g.ThresholdActive, err = reader.ReadElementContentAsObject("ThresholdActive", nil, g, 3)
	if err != nil {
		return err
	}
	g.ThresholdNormal, err = reader.ReadElementContentAsObject("ThresholdNormal", nil, g, 4)
	if err != nil {
		return err
	}
	g.ThresholdEmergency, err = reader.ReadElementContentAsObject("ThresholdEmergency", nil, g, 5)
	if err != nil {
		return err
	}
	ret2, err := reader.ReadElementContentAsLong("MinOverThresholdDuration", 0)
	if err != nil {
		return err
	}
	g.MinOverThresholdDuration = uint32(ret2)
	ret2, err = reader.ReadElementContentAsLong("MinUnderThresholdDuration", 0)
	if err != nil {
		return err
	}
	g.MinUnderThresholdDuration = uint32(ret2)
	b, err = reader.IsStartElementNamed("EmergencyProfile", true)
	if err != nil {
		return err
	}
	if b {
		ret, err := reader.ReadElementContentAsInt("ID", 0)
		if err != nil {
			return err
		}
		g.EmergencyProfile.ID = uint16(ret)
		g.EmergencyProfile.ActivationTime, err = reader.ReadElementContentAsDateTime("Time", nil)
		if err != nil {
			return err
		}
		ret, err = reader.ReadElementContentAsInt("Duration", 0)
		if err != nil {
			return err
		}
		g.EmergencyProfile.Duration = uint32(ret)
		reader.ReadEndElement("EmergencyProfile")
	}
	list := []uint16{}
	b, err = reader.IsStartElementNamed("EmergencyProfileGroupIDs", true)
	if err != nil {
		return err
	}
	if b {
		b, err = reader.IsStartElementNamed("Value", false)
		if err != nil {
			return err
		}
		for b {
			ret, err := reader.ReadElementContentAsInt("Value", 0)
			if err != nil {
				return err
			}
			list = append(list, uint16(ret))
		}
		reader.ReadEndElement("EmergencyProfileGroupIDs")
	}
	g.EmergencyProfileGroupIDs = list
	ret, err := reader.ReadElementContentAsInt("Active", 0)
	if err != nil {
		return err
	}
	g.EmergencyProfileActive = ret != 0
	if b, err := reader.IsStartElementNamed("ActionOverThreshold", true); b {
		g.ActionOverThreshold.LogicalName, err = reader.ReadElementContentAsString("LN", "")
		if err != nil {
			return err
		}
		ret, err = reader.ReadElementContentAsInt("ScriptSelector", 0)
		if err != nil {
			return err
		}
		g.ActionOverThreshold.ScriptSelector = uint16(ret)
		reader.ReadEndElement("ActionOverThreshold")
	}
	if b, err := reader.IsStartElementNamed("ActionUnderThreshold", true); b {
		g.ActionUnderThreshold.LogicalName, err = reader.ReadElementContentAsString("LN", "")
		if err != nil {
			return err
		}
		ret, err = reader.ReadElementContentAsInt("ScriptSelector", 0)
		if err != nil {
			return err
		}
		g.ActionUnderThreshold.ScriptSelector = uint16(ret)
		if err != nil {
			return err
		}
		reader.ReadEndElement("ActionUnderThreshold")
	}
	return err
}

// Save returns the save object content to XML.
//
// Parameters:
//
//	writer: XML writer.
func (g *GXDLMSLimiter) Save(writer *GXXmlWriter) error {
	var err error
	writer.WriteStartElement("MonitoredValue")
	if g.MonitoredValue != nil {
		err = writer.WriteElementString("ObjectType", int(g.MonitoredValue.Base().ObjectType()))
		if err != nil {
			return err
		}
		err = writer.WriteElementString("LN", g.MonitoredValue.Base().LogicalName())
		if err != nil {
			return err
		}
		err = writer.WriteElementString("Index", g.MonitoredAttributeIndex)
		if err != nil {
			return err
		}
	}
	writer.WriteEndElement()
	err = writer.WriteElementObject("ThresholdActive", g.ThresholdActive, enums.DataTypeNone, enums.DataTypeNone)
	if err != nil {
		return err
	}
	err = writer.WriteElementObject("ThresholdNormal", g.ThresholdNormal, enums.DataTypeNone, enums.DataTypeNone)
	if err != nil {
		return err
	}
	err = writer.WriteElementObject("ThresholdEmergency", g.ThresholdEmergency, enums.DataTypeNone, enums.DataTypeNone)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("MinOverThresholdDuration", g.MinOverThresholdDuration)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("MinUnderThresholdDuration", g.MinUnderThresholdDuration)
	if err != nil {
		return err
	}
	writer.WriteStartElement("EmergencyProfile")
	err = writer.WriteElementString("ID", g.EmergencyProfile.ID)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("Time", g.EmergencyProfile.ActivationTime)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("Duration", g.EmergencyProfile.Duration)
	if err != nil {
		return err
	}
	writer.WriteEndElement()
	writer.WriteStartElement("EmergencyProfileGroupIDs")
	if g.EmergencyProfileGroupIDs != nil {
		for _, it := range g.EmergencyProfileGroupIDs {
			err = writer.WriteElementString("Value", it)
			if err != nil {
				return err
			}
		}
	}
	writer.WriteEndElement()
	err = writer.WriteElementString("Active", g.EmergencyProfileActive)
	if err != nil {
		return err
	}
	writer.WriteStartElement("ActionOverThreshold")
	err = writer.WriteElementString("LN", g.ActionOverThreshold.LogicalName)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("ScriptSelector", g.ActionOverThreshold.ScriptSelector)
	if err != nil {
		return err
	}
	writer.WriteEndElement()
	writer.WriteStartElement("ActionUnderThreshold")
	err = writer.WriteElementString("LN", g.ActionUnderThreshold.LogicalName)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("ScriptSelector", g.ActionUnderThreshold.ScriptSelector)
	if err != nil {
		return err
	}
	writer.WriteEndElement()
	return err
}

// PostLoad returns the handle actions after Load.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSLimiter) PostLoad(reader *GXXmlReader) error {
	// Upload Monitored Value after load.
	if g.MonitoredValue != nil {
		target := reader.Objects.FindByLN(g.MonitoredValue.Base().ObjectType(), g.MonitoredValue.Base().LogicalName())
		if target != nil && target != g.MonitoredValue {
			g.MonitoredValue = target
		}
	}
	return nil
}

// GetValues returns the an array containing the COSEM object's attribute values.
func (g *GXDLMSLimiter) GetValues() []any {
	return []any{g.LogicalName(), g.MonitoredValue, g.ThresholdActive, g.ThresholdNormal,
		g.ThresholdEmergency, g.MinOverThresholdDuration, g.MinUnderThresholdDuration,
		g.EmergencyProfile, g.EmergencyProfileGroupIDs, g.EmergencyProfileActive, []any{g.ActionOverThreshold, g.ActionUnderThreshold}}
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
func (g *GXDLMSLimiter) GetDataType(index int) (enums.DataType, error) {
	var ret enums.DataType
	if index == 1 {
		ret = enums.DataTypeOctetString
	} else if index == 2 {
		ret = enums.DataTypeStructure
	} else if index == 3 {
		ret, _ = g.Base().GetDataType(index)
	} else if index == 4 {
		ret, _ = g.Base().GetDataType(index)
	} else if index == 5 {
		ret, _ = g.Base().GetDataType(index)
	} else if index == 6 {
		ret = enums.DataTypeUint32
	} else if index == 7 {
		ret = enums.DataTypeUint32
	} else if index == 8 {
		ret = enums.DataTypeStructure
	} else if index == 9 {
		ret = enums.DataTypeArray
	} else if index == 10 {
		ret = enums.DataTypeBoolean
	} else if index == 11 {
		ret = enums.DataTypeStructure
	} else {
		return 0, dlmserrors.ErrInvalidAttributeIndex
	}
	return ret, nil
}

// NewGXDLMSLimiter creates a new limiter object instance.
//
// The function validates `ln` before creating the object.
//`ln` is the Logical Name and `sn` is the Short Name of the object.
func NewGXDLMSLimiter(ln string, sn int16) (*GXDLMSLimiter, error) {
	err := ValidateLogicalName(ln)
	if err != nil {
		return nil, err
	}
	return &GXDLMSLimiter{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeLimiter,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
