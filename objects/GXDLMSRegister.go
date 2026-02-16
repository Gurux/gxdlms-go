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

package objects

import (
	"errors"
	"math"
	"reflect"

	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// Online help:
//
//	https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSRegister
type GXDLMSRegister struct {
	GXDLMSObject
	scaler int8

	// Unit of COSEM Register object.
	Unit enums.Unit

	// Value of COSEM Register object.
	Value any
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSRegister) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

// Scaler returns the scaler of COSEM Register object.
func (g *GXDLMSRegister) Scaler() float64 {
	return math.Pow(10, float64(g.scaler))
}

// SetScaler sets the scaler of COSEM Register object.
func (g *GXDLMSRegister) SetScaler(value float64) {
	g.scaler = int8(math.Log10(value))
}

// Invoke returns the invokes method.settings: DLMS settings.
// e: Invoke parameters.
func (g *GXDLMSRegister) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	// Resets the value to the default value. The default value is an instance specific constant.
	if e.Index == 1 {
		g.Value = nil
	} else {
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil, nil
}

// GetAttributeIndexToRead returns the returns collection of attributes to read.
// all: All items are returned even if they are read already.
//
//	Returns:
//	    Collection of attributes to read.
func (g *GXDLMSRegister) GetAttributeIndexToRead(all bool) []int {
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
	return attributes
}

// GetNames returns the returns names of attribute indexes.
func (g *GXDLMSRegister) GetNames() []string {
	return []string{"Logical Name", "Value", "Scaler and Unit"}
}

// GetMethodNames returns the returns names of method indexes.
func (g *GXDLMSRegister) GetMethodNames() []string {
	return []string{"Reset"}
}

// GetAttributeCount returns the returns amount of attributes.
func (g *GXDLMSRegister) GetAttributeCount() int {
	return 3
}

// GetMethodCount returns the returns amount of methods.
func (g *GXDLMSRegister) GetMethodCount() int {
	return 1
}

// GetValue returns the returns value of given attribute.
// settings: DLMS settings.
// e: Get parameters.
//
//	Returns:
//	    Value of the attribute index.
func (g *GXDLMSRegister) GetValue(settings *settings.GXDLMSSettings,
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
		if !settings.IsServer() && g.scaler != 1 && g.Value != nil {
			dt, _ := g.Base().GetDataType(2)
			if dt == enums.DataTypeNone && g.Value != nil {
				dt, _ = internal.GetDLMSDataType(reflect.TypeOf(g.Value))
				// If user has set initial value.
				if dt == enums.DataTypeString {
					dt = enums.DataTypeNone
				}
			}
			tmp := internal.AnyToDouble(e.Value) / g.Scaler()
			if dt != enums.DataTypeNone {
				//TODO: ret, err := internal.ChangeType(tmp, dt)
				// return ret, err
			}
			return tmp, nil
		}
		return g.Value, nil
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
// settings: DLMS settings.
// e: Set parameters.
func (g *GXDLMSRegister) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
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
			if settings != nil && settings.IsServer() {
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
			if len(arr) == 2 {
				g.scaler = arr[0].(int8)
				g.Unit = enums.Unit((arr[1].(types.GXEnum).Value))
			}
		}
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return err
}

// Constructor.
// ln: Logical Name of the object.
// sn: Short Name of the object.
func NewGXDLMSRegister(ln string, sn int16) (*GXDLMSRegister, error) {
	err := ValidateLogicalName(ln)
	if err != nil {
		return nil, err
	}
	return &GXDLMSRegister{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeRegister,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}

// Load returns the load object content from XML.reader: XML reader.
func (g *GXDLMSRegister) Load(reader *GXXmlReader) error {
	v, err := reader.ReadElementContentAsInt("Unit", 0)
	if err != nil {
		return err
	}
	g.Unit = enums.Unit(v)
	ret, err := reader.ReadElementContentAsInt("Scaler", 1)
	if err != nil {
		return err
	}
	g.scaler = int8(ret)
	g.Value, err = reader.ReadElementContentAsObject("Value", nil, g, 2)
	return err
}

// Save returns the save object content to XML.writer: XML writer.
func (g *GXDLMSRegister) Save(writer *GXXmlWriter) error {
	writer.WriteElementStringInt("Unit", int(g.Unit))
	writer.WriteElementStringInt("Scaler", int(g.scaler))
	dt, err := g.GetDataType(2)
	if err != nil {
		return err
	}
	writer.WriteElementObject("Value", g.Value, dt, g.GetUIDataType(2))
	return nil
}

// PostLoad returns the handle actions after Load.
// reader: XML reader.
func (g *GXDLMSRegister) PostLoad(reader *GXXmlReader) error {
	return nil
}

// Reset returns the reset value.
// client: DLMS client.
//
//	Returns:
//	    Action bytes.
func (g *GXDLMSRegister) Reset(client IGXDLMSClient) ([][]byte, error) {
	return client.Method(g, 1, 0, enums.DataTypeInt8)
}

// GetValues returns the returns attributes as an array.
//
//	Returns:
//	    Collection of COSEM object values.
func (g *GXDLMSRegister) GetValues() []any {
	return []any{g.LogicalName(), g.Value, types.GXStructure{g.Scaler, g.Unit}}
}

// IsRead returns the is attribute read.
// index: Attribute index to read.
//
//	Returns:
//	    Returns true if attribute is read.
func (g *GXDLMSRegister) IsRead(index int) bool {
	if index == 3 {
		return g.Unit != enums.UnitNone
	}
	return g.Base().IsRead(index)
}

// GetDataType returns the returns device data type of selected attribute index.index: Attribute index of the object.
//
//	Returns:
//	    Device data type of the object.
func (g *GXDLMSRegister) GetDataType(index int) (enums.DataType, error) {
	if index == 1 {
		return enums.DataTypeOctetString, nil
	}
	if index == 2 {
		dt, err := g.Base().GetDataType(index)
		if err == nil && dt == enums.DataTypeNone && g.Value != nil {
			dt, err = internal.GetDLMSDataType(reflect.TypeOf(g.Value))
			if err != nil {
				return enums.DataTypeNone, err
			}
			if dt == enums.DataTypeString {
				dt = enums.DataTypeNone
			}
		}
		return dt, err
	}
	if index == 3 {
		return enums.DataTypeStructure, nil
	}
	return enums.DataTypeNone, errors.New("GetDataType failed. Invalid attribute index.")
}

// GetUIDataType returns UI data type of selected index.
// index: Attribute index of the object.
//
//	Returns:
//	    UseUI data type of the object.
func (g *GXDLMSRegister) GetUIDataType(index int) enums.DataType {
	return g.Base().GetUIDataType(index)
}
