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

	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// Online help:
// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSExtendedRegister
type GXDLMSExtendedRegister struct {
	GXDLMSRegister
	// Status
	Status any

	// Capture time.
	CaptureTime time.Time
}

func (g *GXDLMSExtendedRegister) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	// Resets the value to the default value. The default value is an instance specific constant.
	if e.Index == 1 {
		g.Value = nil
		g.CaptureTime = time.Now()
	} else {
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil, nil
}

func (g *GXDLMSExtendedRegister) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	// LN is static and read only once.
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	// ScalerUnit
	if all || !g.IsRead(3) {
		attributes = append(attributes, 3)
	}
	// Value
	if all || g.CanRead(2) {
		attributes = append(attributes, 2)
	}
	// Status
	if all || g.CanRead(4) {
		attributes = append(attributes, 4)
	}
	// CaptureTime
	if all || g.CanRead(5) {
		attributes = append(attributes, 5)
	}
	return attributes
}

func (g *GXDLMSExtendedRegister) GetNames() []string {
	return []string{"Logical Name", "Value", "Scaler and Unit", "Status", "CaptureTime"}
}

func (g *GXDLMSExtendedRegister) GetAttributeCount() int {
	return 5
}

func (g *GXDLMSExtendedRegister) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	var ret any
	var err error
	switch e.Index {
	case 1:
		ret, err = helpers.LogicalNameToBytes(g.LogicalName())
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
	case 2:
		// If client set new value.
		if !settings.IsServer() && g.scaler != 1 && g.Value != nil {
			dt, err := g.Base().GetDataType(2)
			if err != nil {
				return nil, err
			}
			if dt == enums.DataTypeNone && g.Value != nil {
				dt, err := internal.GetDLMSDataType(reflect.TypeOf(g.Value))
				if err != nil {
					return nil, err
				}
				// If user has set initial value.
				if dt == enums.DataTypeString {
					dt = enums.DataTypeNone
				}
			}
			tmp := internal.AnyToDouble(e.Value) / g.Scaler()
			if dt != enums.DataTypeNone {
				//TODO: tmp = Convert.ChangeType(tmp, internal.GetDataType(dt))
			}
			return tmp, nil
		}
		return g.Value, nil
	case 3:
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
		err = internal.SetData(settings, data, enums.DataTypeEnum, byte(g.Unit))
		if err != nil {
			return nil, err
		}
		return data.Array(), nil
	case 4:
		return g.Status, nil
	case 5:
		return g.CaptureTime, nil
	}
	e.Error = enums.ErrorCodeReadWriteDenied
	return ret, err
}

func (g *GXDLMSExtendedRegister) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
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
			dt, err := internal.GetDLMSDataType(reflect.TypeOf(e.Value))
			if err != nil {
				return err
			}
			g.SetDataType(2, dt)
			if settings.IsServer() {
				g.Value = e.Value
			} else {
				g.Value = internal.AnyToDouble(e.Value) * g.Scaler()
			}
		} else {
			g.Value = e.Value
		}
	case 3:
		if e.Value == nil {
			g.scaler = 1
			g.Unit = enums.UnitNone
		} else {
			arr := e.Value.(types.GXStructure)
			if len(arr) != 2 {
				return errors.New("setValue failed. Invalid scaler unit value.")
			}
			g.scaler = arr[0].(int8)
			g.Unit = enums.Unit(arr[1].(types.GXEnum).Value)
		}
	case 4:
		g.Status = e.Value
	case 5:
		if _, ok := e.Value.([]byte); ok {
			ret, err := internal.ChangeTypeFromByteArray(settings, e.Value.([]byte), enums.DataTypeDateTime)
			if err != nil {
				e.Error = enums.ErrorCodeReadWriteDenied
			}
			g.CaptureTime = ret.(types.GXDateTime).Value
		} else if _, ok := e.Value.(string); ok {
			e.Value, err = types.NewGXDateTimeFromString(e.Value.(string), nil)
			if err != nil {
				e.Error = enums.ErrorCodeReadWriteDenied
			}
		}
		// Actaris meters might return nil.
		if e.Value == nil {
			g.CaptureTime = time.Time{}
		} else {
			g.CaptureTime = e.Value.(types.GXDateTime).Value.Local()
		}
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return err
}

func (g *GXDLMSExtendedRegister) Load(reader *GXXmlReader) error {
	var err error
	ret, err := reader.ReadElementContentAsInt("Unit", 0)
	if err != nil {
		return err
	}
	g.Unit = enums.Unit(ret)
	ret, err = reader.ReadElementContentAsInt("Scaler", 1)
	if err != nil {
		return err
	}
	g.scaler = int8(ret)
	g.Value, err = reader.ReadElementContentAsObject("Value", nil, g, 2)
	if err != nil {
		return err
	}
	g.Status, err = reader.ReadElementContentAsObject("Status", nil, g, 4)
	if err != nil {
		return err
	}
	dt, err := reader.ReadElementContentAsDateTime("CaptureTime", nil)
	if err != nil {
		return err
	}
	g.CaptureTime = dt.Value
	return err
}

func (g *GXDLMSExtendedRegister) Save(writer *GXXmlWriter) error {
	err := writer.WriteElementString("Unit", int(g.Unit))
	if err != nil {
		return err
	}
	err = writer.WriteElementString("Scaler", g.scaler)
	if err != nil {
		return err
	}
	dt, err := g.GetDataType(2)
	if err != nil {
		return err
	}
	udt := g.GetUIDataType(2)
	err = writer.WriteElementObject("Value", g.Value, dt, udt)
	if err != nil {
		return err
	}
	dt, err = g.GetDataType(2)
	if err != nil {
		return err
	}
	udt = g.GetUIDataType(2)
	err = writer.WriteElementObject("Status", g.Status, dt, udt)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("CaptureTime", g.CaptureTime)
	if err != nil {
		return err
	}
	return err
}

func (g *GXDLMSExtendedRegister) PostLoad(reader *GXXmlReader) error {
	return nil
}

func (g *GXDLMSExtendedRegister) GetUIDataType(index int) enums.DataType {
	if index == 5 {
		return enums.DataTypeDateTime
	}
	return g.Base().GetUIDataType(index)
}

func (g *GXDLMSExtendedRegister) GetValues() []any {
	return []any{g.LogicalName, g.Value, []any{g.Scaler, g.Unit}, g.Status, g.CaptureTime}
}

func (g *GXDLMSExtendedRegister) IsRead(index int) bool {
	if index == 3 {
		return g.Unit != 0
	}
	return g.Base().IsRead(index)
}

func (g *GXDLMSExtendedRegister) GetDataType(index int) (enums.DataType, error) {
	if index == 1 {
		return enums.DataTypeOctetString, nil
	}
	if index == 2 {
		return enums.DataTypeNone, nil
	}
	if index == 3 {
		return enums.DataTypeArray, nil
	}
	if index == 4 {
		return g.Base().GetDataType(index)
	}
	if index == 5 {
		return enums.DataTypeOctetString, nil
	}
	return enums.DataTypeNone, errors.New("GetDataType failed. Invalid attribute index.")
}
