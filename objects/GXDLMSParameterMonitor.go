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
	"reflect"
	"time"

	"github.com/Gurux/gxcommon-go"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// Online help:
// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSParameterMonitor
type GXDLMSParameterMonitor struct {
	GXDLMSObject
	// Changed parameter.
	ChangedParameter *GXDLMSTarget

	// Capture time.
	CaptureTime time.Time

	// Changed Parameter
	Parameters []GXDLMSTarget
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSParameterMonitor) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

// Invoke returns the invokes method.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Invoke parameters.
func (g *GXDLMSParameterMonitor) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	if e.Index != 1 && e.Index != 2 {
		e.Error = enums.ErrorCodeReadWriteDenied
	} else {
		if e.Index == 1 {
			tmp := e.Parameters.(types.GXStructure)
			ot := enums.ObjectType(tmp[0].(uint16))
			ln, err := helpers.ToLogicalName(tmp[1].([]byte))
			if err != nil {
				return nil, err
			}
			index := uint8(tmp[2].(int8))
			for _, item := range g.Parameters {
				if item.Target.Base().ObjectType() == ot && item.Target.Base().LogicalName() == ln && item.AttributeIndex == index {
					internal.Remove(g.Parameters, item)
					break
				}
			}
			it := GXDLMSTarget{}
			it.Target = getObjectCollection(settings.Objects).FindByLN(ot, ln)
			if it.Target == nil {
				it.Target, err = CreateObject(ot, ln, 0)
				if err != nil {
					return nil, err
				}
			}
			it.AttributeIndex = index
			g.Parameters = append(g.Parameters, it)
		} else if e.Index == 2 {
			tmp := e.Parameters.(types.GXStructure)
			ot := enums.ObjectType(tmp[0].(uint16))
			ln, err := helpers.ToLogicalName(tmp[1].([]byte))
			if err != nil {
				return nil, err
			}
			index := tmp[2].(byte)
			for _, item := range g.Parameters {
				if item.Target.Base().ObjectType() == ot && item.Target.Base().LogicalName() == ln && item.AttributeIndex == index {
					internal.Remove(g.Parameters, item)
					break
				}
			}
		}
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
func (g *GXDLMSParameterMonitor) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	// LN is static and read only once.
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	// ChangedParameter
	if all || g.CanRead(2) {
		attributes = append(attributes, 2)
	}
	// CaptureTime
	if all || g.CanRead(3) {
		attributes = append(attributes, 3)
	}
	// Parameters
	if all || g.CanRead(4) {
		attributes = append(attributes, 4)
	}
	return attributes
}

// GetNames returns the names of attribute indexes.
func (g *GXDLMSParameterMonitor) GetNames() []string {
	return []string{"Logical Name", "ChangedParameter", "CaptureTime", "Parameters"}
}

// GetMethodNames returns the names of method indexes.
func (g *GXDLMSParameterMonitor) GetMethodNames() []string {
	return []string{"Add parameter", "Delete parameter"}
}

// GetAttributeCount returns the amount of attributes.
//
// Returns:
//
//	Count of attributes.
func (g *GXDLMSParameterMonitor) GetAttributeCount() int {
	return 4
}

// GetMethodCount returns the amount of methods.
func (g *GXDLMSParameterMonitor) GetMethodCount() int {
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
func (g *GXDLMSParameterMonitor) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	var err error
	switch e.Index {
	case 1:
		return helpers.LogicalNameToBytes(g.LogicalName())
	case 2:
		{
			data := types.NewGXByteBuffer()
			err = data.SetUint8(uint8(enums.DataTypeStructure))
			if err != nil {
				return nil, err
			}
			err = data.SetUint8(4)
			if err != nil {
				return nil, err
			}
			if g.ChangedParameter == nil || g.ChangedParameter.Target == nil {
				err = internal.SetData(settings, data, enums.DataTypeUint16, 0)
				if err != nil {
					return nil, err
				}
				ln := []byte{0, 0, 0, 0, 0, 0}
				err = internal.SetData(settings, data, enums.DataTypeOctetString, ln)
				if err != nil {
					return nil, err
				}
				err = internal.SetData(settings, data, enums.DataTypeInt8, 1)
				if err != nil {
					return nil, err
				}
				err = internal.SetData(settings, data, enums.DataTypeNone, nil)
				if err != nil {
					return nil, err
				}
			} else {
				err = internal.SetData(settings, data, enums.DataTypeUint16, g.ChangedParameter.Target.Base().ObjectType())
				if err != nil {
					return nil, err
				}
				ln, err := helpers.LogicalNameToBytes(g.ChangedParameter.Target.Base().LogicalName())
				if err != nil {
					return nil, err
				}

				err = internal.SetData(settings, data, enums.DataTypeOctetString, ln)
				if err != nil {
					return nil, err
				}
				err = internal.SetData(settings, data, enums.DataTypeInt8, g.ChangedParameter.AttributeIndex)
				if err != nil {
					return nil, err
				}
				dt, err := internal.GetDLMSDataType(reflect.TypeOf(g.ChangedParameter.Value))
				if err != nil {
					return nil, err
				}
				err = internal.SetData(settings, data, dt, g.ChangedParameter.Value)
				if err != nil {
					return nil, err
				}
			}
			return data.Array(), nil
		}
	case 3:
		return g.CaptureTime, nil
	case 4:
		{
			data := types.NewGXByteBuffer()
			err = data.SetUint8(uint8(enums.DataTypeArray))
			if err != nil {
				return nil, err
			}
			if g.Parameters == nil {
				err = data.SetUint8(0)
				if err != nil {
					return nil, err
				}
			} else {
				err = data.SetUint8(uint8(len(g.Parameters)))
				if err != nil {
					return nil, err
				}
				for _, it := range g.Parameters {
					err = data.SetUint8(uint8(enums.DataTypeStructure))
					if err != nil {
						return nil, err
					}
					err = data.SetUint8(uint8(3))
					if err != nil {
						return nil, err
					}
					err = internal.SetData(settings, data, enums.DataTypeUint16, it.Target.Base().ObjectType())
					if err != nil {
						return nil, err
					}
					ln, err := helpers.LogicalNameToBytes(it.Target.Base().LogicalName())
					if err != nil {
						return nil, err
					}
					err = internal.SetData(settings, data, enums.DataTypeOctetString, ln)
					if err != nil {
						return nil, err
					}
					err = internal.SetData(settings, data, enums.DataTypeInt8, it.AttributeIndex)
					if err != nil {
						return nil, err
					}
				}
			}
			return data.Array(), nil
		}
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil, nil
}

// SetValue returns the set value of given attribute.
// When raw parameter us not used example register multiplies value by scalar.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Set parameters.
func (g *GXDLMSParameterMonitor) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	var err error
	switch e.Index {
	case 1:
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return g.SetLogicalName(ln)
	case 2:
		g.ChangedParameter = &GXDLMSTarget{}
		if tmp, ok := e.Value.(types.GXStructure); ok {
			if len(tmp) != 4 {
				return errors.New("Invalid structure format.")
			}
			ot := enums.ObjectType(tmp[0].(uint16))
			ln, err := helpers.ToLogicalName(tmp[1].([]byte))
			if err != nil {
				return err
			}
			g.ChangedParameter.Target = getObjectCollection(settings.Objects).FindByLN(ot, ln)
			if g.ChangedParameter.Target == nil {
				g.ChangedParameter.Target, err = CreateObject(ot, ln, 0)
				if err != nil {
					return err
				}
			}
			g.ChangedParameter.AttributeIndex = uint8(tmp[2].(int8))
			g.ChangedParameter.Value = tmp[3]
		}
	case 3:
		if e.Value == nil {
			g.CaptureTime = time.Time{}
		} else {
			if v, ok := e.Value.([]byte); ok {
				e.Value, err = internal.ChangeTypeFromByteArray(settings, v, enums.DataTypeDateTime)
				if err != nil {
					return err
				}
			} else if v, ok := e.Value.(string); ok {
				e.Value, err = types.NewGXDateTimeFromString(v, nil)
				if err != nil {
					return err
				}
			}
			if v, ok := e.Value.(types.GXDateTime); ok {
				g.CaptureTime = v.Value
			} else {
				return gxcommon.ErrInvalidArgument
			}
		}
	case 4:
		g.Parameters = g.Parameters[:0]
		if e.Value != nil {
			for _, it := range e.Value.(types.GXArray) {
				tmp := it.(types.GXStructure)
				if len(tmp) != 3 {
					return errors.New("Invalid structure format.")
				}
				obj := GXDLMSTarget{}
				ot := enums.ObjectType(tmp[0].(uint16))
				ln, err := helpers.ToLogicalName(tmp[1].([]byte))
				if err != nil {
					return err
				}
				obj.Target = getObjectCollection(settings.Objects).FindByLN(ot, ln)
				if obj.Target == nil {
					obj.Target, err = CreateObject(ot, ln, 0)
					if err != nil {
						return err
					}
				}
				obj.AttributeIndex = uint8(tmp[2].(int8))
				g.Parameters = append(g.Parameters, obj)
			}
		}
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return err
}

// Load returns the load object content from XML.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSParameterMonitor) Load(reader *GXXmlReader) error {
	var err error
	g.ChangedParameter = &GXDLMSTarget{}
	if ret, err := reader.IsStartElementNamed("ChangedParameter", true); ret && err == nil {
		ret, err := reader.ReadElementContentAsInt("Type", 0)
		if err != nil {
			return err
		}
		ot := enums.ObjectType(ret)
		ln, err := reader.ReadElementContentAsString("LN", "")
		if err != nil {
			return err
		}
		g.ChangedParameter.Target = getObjectCollection(reader.Objects).FindByLN(ot, ln)
		if g.ChangedParameter.Target == nil {
			g.ChangedParameter.Target, err = CreateObject(ot, ln, 0)
			if err != nil {
				return err
			}
		}
		g.ChangedParameter.AttributeIndex, err = reader.ReadElementContentAsUInt8("Index", 0)
		if err != nil {
			return err
		}
		g.ChangedParameter.Value, err = reader.ReadElementContentAsObject("Value", nil, nil, 0)
		if err != nil {
			return err
		}
		reader.ReadEndElement("ChangedParameter")
	}
	ret, err := reader.ReadElementContentAsGXDateTime("Time")
	if err != nil {
		return err
	}
	g.CaptureTime = ret.Value
	g.Parameters = g.Parameters[:0]
	if ret, err := reader.IsStartElementNamed("Parameters", true); ret && err == nil {
		for {
			ret, err = reader.IsStartElementNamed("Item", true)
			if err != nil {
				return err
			}
			if !ret {
				break
			}
			obj := GXDLMSTarget{}
			ret, err := reader.ReadElementContentAsInt("Type", 0)
			if err != nil {
				return err
			}
			ot := enums.ObjectType(ret)
			ln, err := reader.ReadElementContentAsString("LN", "")
			if err != nil {
				return err
			}
			obj.Target = reader.Objects.FindByLN(ot, ln)
			if obj.Target == nil {
				obj.Target, err = CreateObject(ot, ln, 0)
				if err != nil {
					return err
				}
			}
			obj.AttributeIndex, err = reader.ReadElementContentAsUInt8("Index", 0)
			if err != nil {
				return err
			}
			g.Parameters = append(g.Parameters, obj)
		}
		reader.ReadEndElement("Parameters")
	}
	return err
}

// Save returns the save object content to XML.
//
// Parameters:
//
//	writer: XML writer.
func (g *GXDLMSParameterMonitor) Save(writer *GXXmlWriter) error {
	writer.WriteStartElement("ChangedParameter")
	if g.ChangedParameter != nil && g.ChangedParameter.Target != nil {
		err := writer.WriteElementString("Type", int(g.ChangedParameter.Target.Base().ObjectType()))
		if err != nil {
			return err
		}
		err = writer.WriteElementString("LN", g.ChangedParameter.Target.Base().LogicalName())
		if err != nil {
			return err
		}
		err = writer.WriteElementString("Index", g.ChangedParameter.AttributeIndex)
		if err != nil {
			return err
		}
		err = writer.WriteElementObject("Value", g.ChangedParameter.Value, enums.DataTypeNone, enums.DataTypeNone)
		if err != nil {
			return err
		}
	}
	writer.WriteEndElement()
	err := writer.WriteElementString("Time", g.CaptureTime)
	if err != nil {
		return err
	}
	writer.WriteStartElement("Parameters")
	if g.Parameters != nil && len(g.Parameters) != 0 {
		for _, it := range g.Parameters {
			writer.WriteStartElement("Item")
			err = writer.WriteElementString("Type", int(it.Target.Base().ObjectType()))
			if err != nil {
				return err
			}
			err = writer.WriteElementString("LN", it.Target.Base().LogicalName())
			if err != nil {
				return err
			}
			err = writer.WriteElementString("Index", it.AttributeIndex)
			if err != nil {
				return err
			}
			writer.WriteEndElement()
		}
	}
	writer.WriteEndElement()
	return err
}

// PostLoad returns the handle actions after Load.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSParameterMonitor) PostLoad(reader *GXXmlReader) error {
	return nil
}

// GetValues returns the an array containing the COSEM object's attribute values.
func (g *GXDLMSParameterMonitor) GetValues() []any {
	return []any{g.LogicalName(), g.ChangedParameter, g.CaptureTime, g.Parameters}
}

// Insert returns the inserts a new entry in the table.
func (g *GXDLMSParameterMonitor) Insert(client IGXDLMSClient, entry *GXDLMSTarget) ([][]uint8, error) {
	bb := types.NewGXByteBuffer()
	err := bb.SetUint8(enums.DataTypeStructure)
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(3)
	if err != nil {
		return nil, err
	}
	err = internal.SetData(nil, bb, enums.DataTypeUint16, entry.Target.Base().ObjectType())
	if err != nil {
		return nil, err
	}
	ln, err := helpers.LogicalNameToBytes(entry.Target.Base().LogicalName())
	if err != nil {
		return nil, err
	}
	err = internal.SetData(nil, bb, enums.DataTypeOctetString, ln)
	if err != nil {
		return nil, err
	}
	err = internal.SetData(nil, bb, enums.DataTypeInt8, entry.AttributeIndex)
	if err != nil {
		return nil, err
	}
	return client.Method(g, 1, bb.Array(), enums.DataTypeArray)
}

// Delete returns the deletes an entry from the table.
func (g *GXDLMSParameterMonitor) Delete(client IGXDLMSClient, entry *GXDLMSTarget) ([][]uint8, error) {
	bb := types.NewGXByteBuffer()
	err := bb.SetUint8(enums.DataTypeStructure)
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(3)
	if err != nil {
		return nil, err
	}
	err = internal.SetData(nil, bb, enums.DataTypeUint16, entry.Target.Base().ObjectType())
	if err != nil {
		return nil, err
	}
	ln, err := helpers.LogicalNameToBytes(entry.Target.Base().LogicalName())
	if err != nil {
		return nil, err
	}
	err = internal.SetData(nil, bb, enums.DataTypeOctetString, ln)
	if err != nil {
		return nil, err
	}
	err = internal.SetData(nil, bb, enums.DataTypeInt8, entry.AttributeIndex)
	if err != nil {
		return nil, err
	}
	return client.Method(g, 2, bb.Array(), enums.DataTypeArray)
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
func (g *GXDLMSParameterMonitor) GetDataType(index int) (enums.DataType, error) {
	var ret enums.DataType
	var err error
	switch index {
	case 1:
		ret = enums.DataTypeOctetString
	case 2:
		ret = enums.DataTypeStructure
	case 3:
		ret = enums.DataTypeOctetString
	case 4:
		ret = enums.DataTypeArray
	default:
		err = errors.New("GetDataType failed. Invalid attribute index.")
	}
	return ret, err
}

// NewGXDLMSParameterMonitor creates a new Parameter Monitor object instance.
//
// The function validates `ln` before creating the object.
// `ln` is the Logical Name and `sn` is the Short Name of the object.
func NewGXDLMSParameterMonitor(ln string, sn int16) (*GXDLMSParameterMonitor, error) {
	err := ValidateLogicalName(ln)
	if err != nil {
		return nil, err
	}
	return &GXDLMSParameterMonitor{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeParameterMonitor,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
