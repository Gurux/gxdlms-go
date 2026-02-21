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
	"math"
	"reflect"

	"github.com/Gurux/gxdlms-go/dlmserrors"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
	"golang.org/x/text/language"
)

// Online help:
// https://www.gurux.fi/Gurux.types.Objects.GXDLMSDemandRegister
type GXDLMSDemandRegister struct {
	GXDLMSObject
	scaler int8

	// Current average value of COSEM Data object.
	CurrentAverageValue any

	// Last average value of COSEM Data object.
	LastAverageValue any

	// Unit of COSEM Register object.
	Unit enums.Unit

	// Provides Demand register specific status information.
	Status any

	// Capture time of COSEM Register object.
	CaptureTime *types.GXDateTime

	// Current start time of COSEM Register object.
	StartTimeCurrent *types.GXDateTime

	// Period is the interval between two successive updates of the Last Average Value.
	// (NumberOfPeriods * Period is the denominator for the calculation of the demand).
	Period uint32

	// The number of periods used to calculate the LastAverageValue.
	// NumberOfPeriods >= 1 NumberOfPeriods > 1 indicates that the LastAverageValue represents �sliding demand�.
	// NumberOfPeriods = 1 indicates that the LastAverageValue represents "block demand".
	// The behaviour of the meter after writing a new value to this attribute shall be
	// specified by the manufacturer.
	NumberOfPeriods uint16
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSDemandRegister) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

// Scaler returns the scaler of COSEM Register object.
func (g *GXDLMSDemandRegister) Scaler() float64 {
	return math.Pow(10, float64(g.scaler))
}

// SetScaler sets the scaler of COSEM Register object.
func (g *GXDLMSDemandRegister) SetScaler(value float64) {
	g.scaler = int8(math.Log10(value))
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
func (g *GXDLMSDemandRegister) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	// LN is static and read only once.
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	// Scaler and Unit
	if all || !g.IsRead(4) {
		attributes = append(attributes, 4)
	}
	// CurrentAverageValue
	if all || g.CanRead(2) {
		attributes = append(attributes, 2)
	}
	// LastAverageValue
	if all || g.CanRead(3) {
		attributes = append(attributes, 3)
	}
	// Status
	if all || g.CanRead(5) {
		attributes = append(attributes, 5)
	}
	// CaptureTime
	if all || g.CanRead(6) {
		attributes = append(attributes, 6)
	}
	// StartTimeCurrent
	if all || g.CanRead(7) {
		attributes = append(attributes, 7)
	}
	// Period
	if all || g.CanRead(8) {
		attributes = append(attributes, 8)
	}
	// NumberOfPeriods
	if all || g.CanRead(9) {
		attributes = append(attributes, 9)
	}
	return attributes
}

// GetNames returns the names of attribute indexes.
func (g *GXDLMSDemandRegister) GetNames() []string {
	return []string{"Logical Name", "Current Average Value", "Last Average Value", "Scaler and Unit", "Status", "Capture Time", "Start Time Current", "Period", "Number Of Periods"}
}

// GetMethodNames returns the names of method indexes.
func (g *GXDLMSDemandRegister) GetMethodNames() []string {
	return []string{"Reset", "Next period"}
}

// GetAttributeCount returns the amount of attributes.
//
// Returns:
//
//	Count of attributes.
func (g *GXDLMSDemandRegister) GetAttributeCount() int {
	return 9
}

// GetMethodCount returns the amount of methods.
func (g *GXDLMSDemandRegister) GetMethodCount() int {
	return 2
}

// GetValue returns the returns value of given attribute.
// settings: DLMS settings.
// e: Get parameters.
//
//	Returns:
//	    Value of the attribute index.
func (g *GXDLMSDemandRegister) GetValue(settings *settings.GXDLMSSettings,
	e *internal.ValueEventArgs) (any, error) {
	if e.Index == 1 {
		v, err := helpers.LogicalNameToBytes(g.LogicalName())
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return v, err
	}
	if e.Index == 2 {
		// If client set new value.
		if !settings.IsServer() && g.scaler != 1 && g.CurrentAverageValue != nil {
			dt, _ := g.Base().GetDataType(2)
			if dt == enums.DataTypeNone {
				dt, _ = internal.GetDLMSDataType(reflect.TypeOf(g.CurrentAverageValue))
				// If user has set initial value.
				if dt == enums.DataTypeString {
					dt = enums.DataTypeNone
				}
			}
			tmp := internal.AnyToDouble(e.Value) / g.Scaler()
			if dt != enums.DataTypeNone {
				//TODO: tmp = Convert.ChangeType(tmp, gXCommon.GetDataType(dt))
			}
			return tmp, nil
		}
		return g.CurrentAverageValue, nil
	}
	if e.Index == 3 {
		var err error
		e.ByteArray = true
		data := types.NewGXByteBuffer()
		err = data.SetUint8(uint8(enums.DataTypeStructure))
		if err != nil {
			return nil, err
		}
		err = data.SetUint8(2)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, data, enums.DataTypeInt8, g.scaler)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, data, enums.DataTypeEnum, g.Unit)
		if err != nil {
			return nil, err
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
func (g *GXDLMSDemandRegister) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	var err error
	switch e.Index {
	case 1:
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return g.SetLogicalName(ln)
	case 2:
		if g.scaler != 1 && e.Value != nil && !e.User {
			if settings.IsServer() {
				g.CurrentAverageValue = e.Value
			} else {
				g.CurrentAverageValue = internal.AnyToDouble(e.Value) * g.Scaler()
			}
		} else {
			g.CurrentAverageValue = e.Value
		}
	case 3:
		if g.scaler != 1 && e.Value != nil && !e.User {
			ret, err := internal.GetDLMSDataType(reflect.TypeOf(e.Value))
			if err != nil {
				return err
			}
			g.Base().SetDataType(int(e.Index), ret)
			if settings.IsServer() {
				g.LastAverageValue = e.Value
			} else {
				g.LastAverageValue = internal.AnyToDouble(e.Value) * g.Scaler()
			}
		} else {
			g.LastAverageValue = e.Value
		}
	case 4:
		if e.Value == nil {
			g.scaler = 1
			g.Unit = enums.UnitNone
		} else {
			arr := e.Value.(types.GXStructure)
			if len(arr) == 2 {
				g.scaler = arr[0].(int8)
				g.Unit = enums.Unit((arr[1].(types.GXEnum).Value))
			} else {
				return errors.New("setValue failed. Invalid scaler unit value.")
			}
		}
	case 5:
		g.Status = e.Value
	case 6:
		if e.Value == nil {
			g.CaptureTime = nil
		} else {
			if v, ok := e.Value.([]byte); ok {
				e.Value, err = internal.ChangeTypeFromByteArray(settings, v, enums.DataTypeDateTime)
			} else if _, ok := e.Value.(string); ok {
				e.Value, err = types.NewGXDateTimeFromString(e.Value.(string), nil)
				if err != nil {
					return err
				}
			}
			ret := e.Value.(types.GXDateTime)
			g.CaptureTime = &ret
		}
	case 7:
		if e.Value == nil {
			g.StartTimeCurrent = nil
		} else {
			if _, ok := e.Value.([]byte); ok {
				e.Value, err = internal.ChangeTypeFromByteArray(settings, e.Value.([]byte), enums.DataTypeDateTime)
				if err != nil {
					return err
				}
			} else if _, ok := e.Value.(string); ok {
				e.Value, err = types.NewGXDateTimeFromString(e.Value.(string), nil)
				if err != nil {
					return err
				}
			}
			ret := e.Value.(types.GXDateTime)
			g.StartTimeCurrent = &ret
		}
	case 8:
		g.Period = e.Value.(uint32)
	case 9:
		g.NumberOfPeriods = e.Value.(uint16)
	default:
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
func (g *GXDLMSDemandRegister) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	// Resets the value to the default value. The default value is an instance specific constant.
	switch e.Index {
	case 1:
		g.CurrentAverageValue = nil
		g.LastAverageValue = nil
		g.CaptureTime = nil
		g.StartTimeCurrent = nil

	case 2:
		g.LastAverageValue = g.CurrentAverageValue
		g.CurrentAverageValue = nil
		g.CaptureTime = nil
		g.StartTimeCurrent = nil

	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil, nil
}

// Load returns the load object content from XML.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSDemandRegister) Load(reader *GXXmlReader) error {
	var err error
	g.CurrentAverageValue, err = reader.ReadElementContentAsObject("CurrentAverageValue", nil, g, 2)
	if err != nil {
		return err
	}
	g.LastAverageValue, err = reader.ReadElementContentAsObject("LastAverageValue", nil, g, 3)
	if err != nil {
		return err
	}
	ret, err := reader.ReadElementContentAsInt("Scaler", 1)
	if err != nil {
		return err
	}
	g.scaler = int8(ret)
	ret, err = reader.ReadElementContentAsInt("Unit", 0)
	if err != nil {
		return err
	}
	g.Unit = enums.Unit(ret)
	g.Status, err = reader.ReadElementContentAsObject("Status", nil, g, 5)
	if err != nil {
		return err
	}
	str, err := reader.ReadElementContentAsString("CaptureTime", "")
	if err != nil {
		return err
	}
	if str == "" {
		g.CaptureTime = nil
	} else {
		g.CaptureTime, err = types.NewGXDateTimeFromString(str, &language.AmericanEnglish)
		if err != nil {
			return err
		}
	}
	str, err = reader.ReadElementContentAsString("StartTimeCurrent", "")
	if err != nil {
		return err
	}
	if str == "" {
		g.StartTimeCurrent = nil
	} else {
		g.StartTimeCurrent, err = types.NewGXDateTimeFromString(str, &language.AmericanEnglish)
		if err != nil {
			return err
		}
	}
	ret, err = reader.ReadElementContentAsInt("Period", 0)
	if err != nil {
		return err
	}
	g.Period = uint32(ret)
	ret, err = reader.ReadElementContentAsInt("NumberOfPeriods", 0)
	if err != nil {
		return err
	}
	g.NumberOfPeriods = uint16(ret)
	return err
}

// Save returns the save object content to XML.
//
// Parameters:
//
//	writer: XML writer.
func (g *GXDLMSDemandRegister) Save(writer *GXXmlWriter) error {
	dt, err := g.GetDataType(2)
	if err != nil {
		return err
	}
	err = writer.WriteElementObject("CurrentAverageValue", g.CurrentAverageValue, dt, enums.DataTypeNone)
	if err != nil {
		return err
	}
	dt, err = internal.GetDLMSDataType(reflect.TypeOf(g.LastAverageValue))
	if err != nil {
		return err
	}
	err = writer.WriteElementObject("LastAverageValue", g.LastAverageValue, dt, enums.DataTypeNone)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("Scaler", g.scaler)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("Unit", int(g.Unit))
	if err != nil {
		return err
	}
	dt, err = internal.GetDLMSDataType(reflect.TypeOf(g.Status))
	if err != nil {
		return err
	}
	err = writer.WriteElementObject("Status", g.Status, dt, enums.DataTypeNone)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("CaptureTime", g.CaptureTime)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("StartTimeCurrent", g.StartTimeCurrent)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("Period", g.Period)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("NumberOfPeriods", g.NumberOfPeriods)
	if err != nil {
		return err
	}
	return err
}

// PostLoad returns the handle actions after Load.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSDemandRegister) PostLoad(reader *GXXmlReader) error {
	return nil
}

// IsRead returns the is attribute read.
//
// Parameters:
//
//	index: Attribute index to read.
//
// Returns:
//
//	Returns true if attribute is read.
func (g *GXDLMSDemandRegister) IsRead(index int) bool {
	if index == 4 {
		return g.Unit != enums.UnitNone
	}
	return g.Base().IsRead(index)
}

// GetValues returns the an array containing the COSEM object's attribute values.
func (g *GXDLMSDemandRegister) GetValues() []any {
	return []any{g.LogicalName(), g.CurrentAverageValue, g.LastAverageValue, types.GXStructure{g.Scaler, g.Unit},
		g.Status, g.CaptureTime, g.StartTimeCurrent, g.Period, g.NumberOfPeriods,
	}
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
func (g *GXDLMSDemandRegister) GetUIDataType(index int) enums.DataType {
	if index == 6 || index == 7 {
		return enums.DataTypeDateTime
	}
	return g.Base().GetUIDataType(index)
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
func (g *GXDLMSDemandRegister) GetDataType(index int) (enums.DataType, error) {
	var err error
	var dt enums.DataType
	switch index {
	case 1:
		dt = enums.DataTypeOctetString
	case 2:
		dt, err = g.Base().GetDataType(index)
		if err != nil {
			return enums.DataTypeNone, err
		}
		if dt == enums.DataTypeNone && g.CurrentAverageValue != nil {
			dt, err = internal.GetDLMSDataType(reflect.TypeOf(g.CurrentAverageValue))
			if err != nil {
				return enums.DataTypeNone, err
			}
		}
	case 3:
		dt, err = g.Base().GetDataType(index)
		if err != nil {
			return enums.DataTypeNone, err
		}
		if dt == enums.DataTypeNone && g.LastAverageValue != nil {
			dt, err = internal.GetDLMSDataType(reflect.TypeOf(g.LastAverageValue))
			if err != nil {
				return enums.DataTypeNone, err
			}
		}
	case 4:
		dt = enums.DataTypeArray
	case 5:
		dt, err = g.Base().GetDataType(index)
		if err != nil {
			return enums.DataTypeNone, err
		}
		if dt == enums.DataTypeNone && g.Status != nil {
			dt, err = internal.GetDLMSDataType(reflect.TypeOf(g.Status))
			if err != nil {
				return enums.DataTypeNone, err
			}
		}
	case 6:
		dt = enums.DataTypeOctetString
	case 7:
		dt = enums.DataTypeOctetString
	case 8:
		dt = enums.DataTypeUint32
	case 9:
		dt = enums.DataTypeUint16
	default:
		return 0, dlmserrors.ErrInvalidAttributeIndex
	}
	return dt, nil
}

// Reset returns the reset value.
//
// Parameters:
//
//	client: DLMS client.
//
// Returns:
//
//	Action bytes.
func (g *GXDLMSDemandRegister) Reset(client IGXDLMSClient) ([][]uint8, error) {
	return client.Method(g, 1, int8(0), enums.DataTypeInt8)
}

// NextPeriod returns the closes the current period and starts a new one.
//
// Parameters:
//
//	client: DLMS client.
//
// Returns:
//
//	Action bytes.
func (g *GXDLMSDemandRegister) NextPeriod(client IGXDLMSClient) ([][]uint8, error) {
	return client.Method(g, 2, int8(0), enums.DataTypeInt8)
}

// NewGXDLMSDemandRegister creates a new demand register object instance.
//
// The function validates `ln` before creating the object.
//`ln` is the Logical Name and `sn` is the Short Name of the object.
func NewGXDLMSDemandRegister(ln string, sn int16) (*GXDLMSDemandRegister, error) {
	err := ValidateLogicalName(ln)
	if err != nil {
		return nil, err
	}
	return &GXDLMSDemandRegister{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeDemandRegister,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
