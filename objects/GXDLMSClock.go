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
	"time"

	"github.com/Gurux/gxdlms-go/dlmserrors"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// Online help:
// https://www.gurux.fi/Gurux.types.Objects.GXDLMSClock
type GXDLMSClock struct {
	GXDLMSObject
	// Time of COSEM Clock object.
	Time types.GXDateTime

	// TimeZone of COSEM Clock object.
	TimeZone int16

	// Status of COSEM Clock object.
	Status enums.ClockStatus

	Begin types.GXDateTime

	End types.GXDateTime

	Deviation int8

	// Is summer time enabled.
	Enabled bool

	// Clock base of COSEM Clock object.
	ClockBase enums.ClockBase
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSClock) Base() *GXDLMSObject {
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
func (g *GXDLMSClock) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	// LN is static and read only once.
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	// Time
	if all || g.CanRead(2) {
		attributes = append(attributes, 2)
	}
	// TimeZone
	if all || !g.CanRead(3) {
		attributes = append(attributes, 3)
	}
	// Status
	if all || g.CanRead(4) {
		attributes = append(attributes, 4)
	}
	// Begin
	if all || !g.CanRead(5) {
		attributes = append(attributes, 5)
	}
	// End
	if all || !g.CanRead(6) {
		attributes = append(attributes, 6)
	}
	// Deviation
	if all || !g.CanRead(7) {
		attributes = append(attributes, 7)
	}
	// Enabled
	if all || !g.CanRead(8) {
		attributes = append(attributes, 8)
	}
	// ClockBase
	if all || !g.CanRead(9) {
		attributes = append(attributes, 9)
	}
	return attributes
}

// GetNames returns the names of attribute indexes.
func (g *GXDLMSClock) GetNames() []string {
	return []string{"Logical Name", "Time", "Time Zone", "Status", "Begin", "End", "Deviation", "Enabled", "Clock Base"}
}

// GetMethodNames returns the names of method indexes.
func (g *GXDLMSClock) GetMethodNames() []string {
	return []string{"Adjust to quarter", "Adjust to measuring period", "Adjust to minute", "Adjust to preset time", "Preset adjusting time", "Shift time"}
}

// GetAttributeCount returns the amount of attributes.
//
// Returns:
//
//	Count of attributes.
func (g *GXDLMSClock) GetAttributeCount() int {
	return 9
}

// GetMethodCount returns the amount of methods.
func (g *GXDLMSClock) GetMethodCount() int {
	return 6
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
func (g *GXDLMSClock) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	var ret any
	var err error
	switch e.Index {
	case 1:
		ret, err = helpers.LogicalNameToBytes(g.LogicalName())
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return nil, err
		}
	case 2:
		ret = g.Time
	case 3:
		ret = g.TimeZone
	case 4:
		ret = g.Status
	case 5:
		ret = g.Begin
	case 6:
		ret = g.End
	case 7:
		ret = g.Deviation
	case 8:
		ret = g.Enabled
	case 9:
		ret = g.ClockBase
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return ret, nil
}

// SetValue returns the set value of given attribute.
// When raw parameter us not used example register multiplies value by scalar.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Set parameters.
func (g *GXDLMSClock) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	var err error

	switch e.Index {
	case 1:
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return g.SetLogicalName(ln)
	case 2:
		if e.Value == nil {
			g.Time = types.GXDateTime{}
		} else {
			if _, ok := e.Value.([]byte); ok {
				e.Value, err = internal.ChangeTypeFromByteArray(settings, e.Value.([]byte), enums.DataTypeDateTime)
				if err != nil {
					e.Error = enums.ErrorCodeReadWriteDenied
				}
			} else if _, ok := e.Value.(string); ok {
				e.Value, err = types.NewGXDateTimeFromString(e.Value.(string), nil)
			}
			if _, ok := e.Value.(types.GXDateTime); ok {
				g.Time = types.GXDateTime(e.Value.(types.GXDateTime))
			} else if _, ok := e.Value.(string); ok {
				e.Value, err = types.NewGXDateTimeFromString(e.Value.(string), nil)
			} else {
				g.Time = types.GXDateTime(e.Value.(types.GXDateTime))
			}
		}
	case 3:
		g.TimeZone = e.Value.(int16)
	case 4:
		g.Status = enums.ClockStatus(e.Value.(uint8))
	case 5:
		if e.Value == nil {
			g.Begin = types.GXDateTime{}
		} else {
			if _, ok := e.Value.([]byte); ok {
				e.Value, err = internal.ChangeTypeFromByteArray(settings, e.Value.([]byte), enums.DataTypeDateTime)
				if err != nil {
					e.Error = enums.ErrorCodeReadWriteDenied
					return err
				}
			} else if _, ok := e.Value.(string); ok {
				e.Value, err = types.NewGXDateTimeFromString(e.Value.(string), nil)
			}
			g.Begin = e.Value.(types.GXDateTime)
		}
	case 6:
		if e.Value == nil {
			g.End = types.GXDateTime{}
		} else {
			if _, ok := e.Value.([]byte); ok {
				e.Value, err = internal.ChangeTypeFromByteArray(settings, e.Value.([]byte), enums.DataTypeDateTime)
				if err != nil {
					e.Error = enums.ErrorCodeReadWriteDenied
					return err
				}
			} else if _, ok := e.Value.(string); ok {
				e.Value, err = types.NewGXDateTimeFromString(e.Value.(string), nil)
			}
			g.End = e.Value.(types.GXDateTime)
		}
	case 7:
		g.Deviation = e.Value.(int8)
	case 8:
		g.Enabled = e.Value.(bool)
		if settings != nil && settings.IsServer() {
			if g.Enabled {
				g.Status |= enums.ClockStatusDaylightSavingActive
			} else {
				g.Status &= ^enums.ClockStatusDaylightSavingActive
			}
		}
	case 9:
		g.ClockBase = enums.ClockBase(e.Value.(types.GXEnum).Value)
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
func (g *GXDLMSClock) Load(reader *GXXmlReader) error {
	var err error
	var val int
	g.Time, err = reader.ReadElementContentAsGXDateTime("Time")
	val, err = reader.ReadElementContentAsInt("TimeZone", 0)
	if err != nil {
		return err
	}
	g.TimeZone = int16(val)
	val, err = reader.ReadElementContentAsInt("Status", 0)
	if err != nil {
		return err
	}
	g.Status = enums.ClockStatus(val)
	g.Begin, err = reader.ReadElementContentAsGXDateTime("Begin")
	if err != nil {
		return err
	}
	g.End, err = reader.ReadElementContentAsGXDateTime("End")
	if err != nil {
		return err
	}
	val, err = reader.ReadElementContentAsInt("Deviation", 0)
	if err != nil {
		return err
	}
	g.Deviation = int8(val)
	g.Enabled, err = reader.ReadElementContentAsBool("Enabled", false)
	if err != nil {
		return err
	}
	val, err = reader.ReadElementContentAsInt("ClockBase", 0)
	if err != nil {
		return err
	}
	g.ClockBase = enums.ClockBase(val)
	return err
}

// Save returns the save object content to XML.
//
// Parameters:
//
//	writer: XML writer.
func (g *GXDLMSClock) Save(writer *GXXmlWriter) error {
	err := writer.WriteElementStringGXDateTime("Time", &g.Time)
	if err != nil {
		return err
	}
	err = writer.WriteElementStringInt("TimeZone", int(g.TimeZone))
	if err != nil {
		return err
	}
	err = writer.WriteElementStringInt("Status", int(g.Status))
	if err != nil {
		return err
	}
	err = writer.WriteElementStringGXDateTime("Begin", &g.Begin)
	if err != nil {
		return err
	}
	err = writer.WriteElementStringGXDateTime("End", &g.End)
	if err != nil {
		return err
	}
	err = writer.WriteElementStringInt("Deviation", int(g.Deviation))
	if err != nil {
		return err
	}
	err = writer.WriteElementStringBool("Enabled", g.Enabled)
	if err != nil {
		return err
	}
	return writer.WriteElementStringInt("ClockBase", int(g.ClockBase))
}

func (g *GXDLMSClock) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, nil
}

func (g *GXDLMSClock) PostLoad(reader *GXXmlReader) error {
	return nil
}

// GetUIDataType returns the UI data type of selected index.
//
// Parameters:
//
//	index: Attribute index of the object.
//
// Returns:
//
//	UI data type of the object.
func (g *GXDLMSClock) GetUIDataType(index int) enums.DataType {
	if index == 2 || index == 5 || index == 6 {
		return enums.DataTypeDateTime
	}
	return enums.DataTypeNone
}

// GetValues returns the an array containing the COSEM object's attribute values.
func (g *GXDLMSClock) GetValues() []any {
	return []any{g.LogicalName(), g.Time, g.TimeZone, g.Status, g.Begin, g.End, g.Deviation, g.Enabled, g.ClockBase}
}

// AdjustToQuarter returns the sets the meter's time to the nearest (+/-) quarter of an hour value (*:00, *:15, *:30, *:45).
func (g *GXDLMSClock) AdjustToQuarter(client IGXDLMSClient) ([][]byte, error) {
	return client.Method(g, 1, 0, enums.DataTypeInt8)
}

// AdjustToMeasuringPeriod returns the sets the meter's time to the nearest (+/-) starting point of a measuring period.
func (g *GXDLMSClock) AdjustToMeasuringPeriod(client IGXDLMSClient) ([][]byte, error) {
	return client.Method(g, 2, 0, enums.DataTypeInt8)
}

// AdjustToMinute returns the sets the meter's time to the nearest minute.
// If second_counter lower 30 s, so second_counter is set to 0.
// If second_counter higher 30 s, so second_counter is set to 0, and
// minute_counter and all depending clock values are incremented if necessary.
func (g *GXDLMSClock) AdjustToMinute(client IGXDLMSClient) ([][]byte, error) {
	return client.Method(g, 3, 0, enums.DataTypeInt8)
}

// AdjustToPresetTime returns the this Method is used in conjunction with the preset_adjusting_time
// Method. If the meter's time lies between validity_interval_start and
// validity_interval_end, then time is set to preset_time.
func (g *GXDLMSClock) AdjustToPresetTime(client IGXDLMSClient) ([][]byte, error) {
	return client.Method(g, 4, 0, enums.DataTypeInt8)
}

// PresetAdjustingTime returns the presets the time to a new value (preset_time) and defines a validity_interval within which the new time can be activated.
func (g *GXDLMSClock) PresetAdjustingTime(client IGXDLMSClient, presetTime *time.Time, validityIntervalStart *time.Time, validityIntervalEnd *time.Time) ([][]byte, error) {
	var err error
	var buff types.GXByteBuffer
	err = buff.Add(uint8(enums.DataTypeStructure))
	if err != nil {
		return nil, err
	}
	err = buff.Add(uint8(3))
	if err != nil {
		return nil, err
	}
	err = internal.SetData(client.Settings(), &buff, enums.DataTypeOctetString, presetTime)
	if err != nil {
		return nil, err
	}
	err = internal.SetData(client.Settings(), &buff, enums.DataTypeOctetString, validityIntervalStart)
	if err != nil {
		return nil, err
	}
	err = internal.SetData(client.Settings(), &buff, enums.DataTypeOctetString, validityIntervalEnd)
	if err != nil {
		return nil, err
	}
	return client.Method(g, 5, buff.Array(), enums.DataTypeArray)
}

// ShiftTime returns the shifts the time by n (-900 &lt;= n &lt;= 900) s.
func (g *GXDLMSClock) ShiftTime(client IGXDLMSClient, time int) ([][]byte, error) {
	if time < -900 || time > 900 {
		return nil, errors.New("Invalid shift time.")
	}
	return client.Method(g, 6, time, enums.DataTypeInt16)
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
func (g *GXDLMSClock) GetDataType(index int) (enums.DataType, error) {
	if index == 1 {
		return enums.DataTypeOctetString, nil
	}
	if index == 2 {
		return enums.DataTypeOctetString, nil
	}
	if index == 3 {
		return enums.DataTypeInt16, nil
	}
	if index == 4 {
		return enums.DataTypeUint8, nil
	}
	if index == 5 {
		return enums.DataTypeOctetString, nil
	}
	if index == 6 {
		return enums.DataTypeOctetString, nil
	}
	if index == 7 {
		return enums.DataTypeInt8, nil
	}
	if index == 8 {
		return enums.DataTypeBoolean, nil
	}
	if index == 9 {
		return enums.DataTypeEnum, nil
	}
	return enums.DataTypeNone, dlmserrors.ErrInvalidAttributeIndex
}

// NewGXDLMSClock creates a new clock object instance.
//
// The function validates `ln` before creating the object.
//`ln` is the Logical Name and `sn` is the Short Name of the object.
func NewGXDLMSClock(ln string, sn int16) (*GXDLMSClock, error) {
	err := ValidateLogicalName(ln)
	if err != nil {
		return nil, err
	}
	return &GXDLMSClock{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeClock,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
