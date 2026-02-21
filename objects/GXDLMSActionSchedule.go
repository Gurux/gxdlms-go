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
	"time"

	"github.com/Gurux/gxcommon-go"
	"github.com/Gurux/gxdlms-go/dlmserrors"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
	"golang.org/x/text/language"
)

// GXDLMSActionSchedule represents the DLMS Action Schedule object.
// Online help: https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSActionSchedule
type GXDLMSActionSchedule struct {
	GXDLMSObject
	// Script to execute.
	Target *GXDLMSScriptTable

	ExecutedScriptLogicalName string

	// Zero based script index to execute.
	ExecutedScriptSelector uint16

	Type enums.SingleActionScheduleType

	ExecutionTime []types.GXDateTime
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSActionSchedule) Base() *GXDLMSObject {
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
func (g *GXDLMSActionSchedule) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	// LN is static and read only once.
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	// ExecutedScriptLogicalName is static and read only once.
	if all || !g.CanRead(2) {
		attributes = append(attributes, 2)
	}
	// Type is static and read only once.
	if all || !g.CanRead(3) {
		attributes = append(attributes, 3)
	}
	// ExecutionTime is static and read only once.
	if all || !g.CanRead(4) {
		attributes = append(attributes, 4)
	}
	return attributes
}

// GetNames returns the names of attribute indexes.
func (g *GXDLMSActionSchedule) GetNames() []string {
	return []string{"Logical Name", "Executed script logical name", "Type", "Execution time"}
}

// GetMethodNames returns the names of method indexes.
func (g *GXDLMSActionSchedule) GetMethodNames() []string {
	return []string{}
}

// GetAttributeCount returns the amount of attributes.
//
// Returns:
//
//	Count of attributes.
func (g *GXDLMSActionSchedule) GetAttributeCount() int {
	return 4
}

// GetMethodCount returns the amount of methods.
func (g *GXDLMSActionSchedule) GetMethodCount() int {
	return 0
}

func (g *GXDLMSActionSchedule) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
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
		err = data.SetUint8(uint8(enums.DataTypeStructure))
		if err != nil {
			return nil, err
		}
		err = data.SetUint8(2)
		if err != nil {
			return nil, err
		}
		// LN
		var ln []byte
		if g.Target != nil {
			ln, err = helpers.LogicalNameToBytes(g.Target.Base().LogicalName())
			if err != nil {
				return nil, err
			}
			err = internal.SetData(settings, data, enums.DataTypeOctetString, ln)
		} else {
			ln, err = helpers.LogicalNameToBytes(g.ExecutedScriptLogicalName)
			if err != nil {
				return nil, err
			}
			err = internal.SetData(settings, data, enums.DataTypeOctetString, ln)
		}
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, data, enums.DataTypeUint16, g.ExecutedScriptSelector)
		if err != nil {
			return nil, err
		}
		return data.Array(), nil
	}
	if e.Index == 3 {
		return int(g.Type), nil
	}
	if e.Index == 4 {
		data := types.NewGXByteBuffer()
		err = data.SetUint8(uint8(enums.DataTypeArray))
		if err != nil {
			return nil, err
		}
		err = types.SetObjectCount(len(g.ExecutionTime), data)
		if err != nil {
			return nil, err
		}
		for _, it := range g.ExecutionTime {
			err = data.SetUint8(uint8(enums.DataTypeStructure))
			if err != nil {
				return nil, err
			}
			err = data.SetUint8(uint8(2))
			if err != nil {
				return nil, err
			}
			if settings != nil && settings.Standard == enums.StandardSaudiArabia {
				err = internal.SetData(settings, data, enums.DataTypeTime, *types.NewGXTimeFromDateTime(&it))
				if err != nil {
					return nil, err
				}
				err = internal.SetData(settings, data, enums.DataTypeDate, *types.NewGXDateFromDateTime(&it))
			} else {
				err = internal.SetData(settings, data, enums.DataTypeOctetString, *types.NewGXTimeFromDateTime(&it))
				if err != nil {
					return nil, err
				}
				err = internal.SetData(settings, data, enums.DataTypeOctetString, *types.NewGXDateFromDateTime(&it))
			}
			if err != nil {
				return nil, err
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
func (g *GXDLMSActionSchedule) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	if e.Index == 1 {
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return err
		}
		return g.SetLogicalName(ln)
	} else if e.Index == 2 {
		if e.Value != nil {
			arr := e.Value.(types.GXStructure)
			if len(arr) != 0 {
				ln, err := helpers.ToLogicalName(arr[0])
				if err != nil {
					return err
				}
				ret := getObjectCollection(settings.Objects).FindByLN(enums.ObjectTypeScriptTable, ln)
				if ret != nil {
					g.Target = ret.(*GXDLMSScriptTable)
					if g.Target == nil {
						g.Target, _ = NewGXDLMSScriptTable(ln, 0)
					}
					g.ExecutedScriptSelector = arr[1].(uint16)
				}
			} else {
				g.Target = nil
				g.ExecutedScriptSelector = 0
			}
		} else {
			g.Target = nil
			g.ExecutedScriptSelector = 0
		}
	} else if e.Index == 3 {
		g.Type = enums.SingleActionScheduleType(e.Value.(types.GXEnum).Value)
	} else if e.Index == 4 {
		g.ExecutionTime = nil
		if e.Value != nil {
			items := []types.GXDateTime{}
			for _, tmp2 := range e.Value.(types.GXArray) {
				it := tmp2.(types.GXStructure)
				var time_ types.GXTime
				if _, ok := it[0].([]byte); ok {
					ret, err := internal.ChangeTypeFromByteArray(settings, it[0].([]byte), enums.DataTypeTime)
					if err != nil {
						return err
					}
					time_ = ret.(types.GXTime)
				} else if v, ok := it[0].(types.GXDateTime); ok {
					time_ = *types.NewGXTimeFromDateTime(&v)
				} else {
					return gxcommon.ErrInvalidArgument
				}
				time_.Skip &^= (enums.DateTimeSkipsYear | enums.DateTimeSkipsMonth | enums.DateTimeSkipsDay | enums.DateTimeSkipsDayOfWeek)
				var date_ types.GXDate
				if _, ok := it[1].([]byte); ok {
					ret, err := internal.ChangeTypeFromByteArray(settings, it[1].([]byte), enums.DataTypeDate)
					if err != nil {
						return err
					}
					date_ = ret.(types.GXDate)
				} else if v, ok := it[1].(types.GXDate); ok {
					date_ = v
				} else {
					return gxcommon.ErrInvalidArgument
				}
				date_.Skip &^= (enums.DateTimeSkipsHour | enums.DateTimeSkipsMinute | enums.DateTimeSkipsSecond | enums.DateTimeSkipsMs)
				tmp := types.GXDateTime{
					Value: time.Date(int(date_.Value.Year()), time.Month(date_.Value.Month()), int(date_.Value.Day()),
						0, 0, 0, 0, time.Local),
				}
				tmp.Value = tmp.Value.Add(time.Duration(time_.Value.Hour()) * time.Hour)
				tmp.Value = tmp.Value.Add(time.Duration(time_.Value.Minute()) * time.Minute)
				tmp.Value = tmp.Value.Add(time.Duration(time_.Value.Second()) * time.Second)
				tmp.Skip = date_.Skip | time_.Skip
				items = append(items, tmp)
			}
			g.ExecutionTime = items
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
func (g *GXDLMSActionSchedule) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, nil
}

// Load returns the load object content from XML.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSActionSchedule) Load(reader *GXXmlReader) error {
	ret, err := reader.ReadElementContentAsInt("ObjectType", 0)
	if err != nil {
		return err
	}
	ot := enums.ObjectType(ret)
	ln, err := reader.ReadElementContentAsString("LN", "")
	if err != nil {
		return err
	}
	if ot != enums.ObjectTypeNone && ln != "" {
		g.Target = reader.Objects.FindByLN(ot, ln).(*GXDLMSScriptTable)
		// if object is not load yet.
		if g.Target == nil {
			g.Target, err = NewGXDLMSScriptTable(ln, 0)
			if err != nil {
				return err
			}
		}
	}
	ret, err = reader.ReadElementContentAsInt("ExecutedScriptSelector", 0)
	if err != nil {
		return err
	}
	g.ExecutedScriptSelector = uint16(ret)
	ret, err = reader.ReadElementContentAsInt("Type", 0)
	if err != nil {
		return err
	}
	g.Type = enums.SingleActionScheduleType(ret)
	list := []types.GXDateTime{}
	if ret, err := reader.IsStartElementNamed("ExecutionTime", true); ret && err == nil {
		for {
			ret, err = reader.IsStartElementNamed("Time", true)
			if err != nil {
				return err
			}
			if !ret {
				break
			}
			it, err := reader.ReadElementContentAsString("Time", "")
			if err != nil {
				return err
			}
			t, err := types.NewGXDateTimeFromString(it, &language.AmericanEnglish)
			if err != nil {
				return err
			}
			list = append(list, *t)
		}
		reader.ReadEndElement("ExecutionTime")
	}
	g.ExecutionTime = list
	return err
}

// Save returns the save object content to XML.
//
// Parameters:
//
//	writer: XML writer.
func (g *GXDLMSActionSchedule) Save(writer *GXXmlWriter) error {
	var err error
	if g.Target != nil {
		err = writer.WriteElementString("ObjectType", int(g.Target.objectType))
		if err != nil {
			return err
		}
		err = writer.WriteElementString("LN", g.Target.LogicalName)
		if err != nil {
			return err
		}
	}
	err = writer.WriteElementString("ExecutedScriptSelector", g.ExecutedScriptSelector)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("Type", int(g.Type))
	if err != nil {
		return err
	}
	if g.ExecutionTime != nil {
		writer.WriteStartElement("ExecutionTime")
		for _, it := range g.ExecutionTime {
			err = writer.WriteElementString("Time", it)
			if err != nil {
				return err
			}
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
func (g *GXDLMSActionSchedule) PostLoad(reader *GXXmlReader) error {
	// Upload target after load.
	if g.Target != nil {
		target := reader.Objects.FindByLN(enums.ObjectTypeScriptTable, g.Target.LogicalName()).(*GXDLMSScriptTable)
		if target != nil && target != g.Target {
			g.Target = target
		}
	}
	return nil
}

// GetValues returns the an array containing the COSEM object's attribute values.
func (g *GXDLMSActionSchedule) GetValues() []any {
	if g.Target != nil {
		return []any{g.LogicalName(), fmt.Sprintf("%s %d", g.Target.LogicalName(), g.ExecutedScriptSelector), g.Type, g.ExecutionTime}
	}
	return []any{g.LogicalName(), fmt.Sprintf("%s %d", g.ExecutedScriptLogicalName, g.ExecutedScriptSelector), g.Type, g.ExecutionTime}
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
func (g *GXDLMSActionSchedule) GetDataType(index int) (enums.DataType, error) {
	if index == 1 {
		return enums.DataTypeOctetString, nil
	}
	if index == 2 {
		return enums.DataTypeArray, nil
	}
	if index == 3 {
		return enums.DataTypeEnum, nil
	}
	if index == 4 {
		return enums.DataTypeArray, nil
	}
	return enums.DataTypeNone, dlmserrors.ErrInvalidAttributeIndex
}

// NewGXDLMSActionSchedule creates a new Action Schedule object instance.
//
// The var attributes []int` before creating the object.
// `ln` is the Logical Name and `sn` is the Short Name of the object.
func NewGXDLMSActionSchedule(ln string, sn int16) (*GXDLMSActionSchedule, error) {
	err := ValidateLogicalName(ln)
	if err != nil {
		return nil, err
	}
	return &GXDLMSActionSchedule{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeActionSchedule,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
