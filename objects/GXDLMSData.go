package objects

// --------------------------------------------------------------------------
//
//	Gurux Ltd
//
// Filename:        $HeadURL$
//
// Version:         $Revision$,
//
//	$Date$
//	$Author$
//
// # Copyright (c) Gurux Ltd
//
// ---------------------------------------------------------------------------
//
//	DESCRIPTION
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
// ---------------------------------------------------------------------------

import (
	"errors"
	"reflect"

	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// Online help:
//
//	https://www.gurux.fi/Gurux.internal.Objects.GXDLMSData
type GXDLMSData struct {
	GXDLMSObject
	// Value of COSEM Data object.
	Value any
}

// base returns the base GXDLMSObject of the object.
func (g *GXDLMSData) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

// GetAttributeIndexToRead returns the returns collection of attributes to read.all: All items are returned even if they are read already.
//
//	Returns:
//	    Collection of attributes to read.
func (g *GXDLMSData) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	// LN is static and read only once.
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	// Value
	if all || g.CanRead(2) {
		attributes = append(attributes, 2)
	}
	return attributes
}

// GetNames returns the returns names of attribute indexes.
//
//	Returns:
func (g *GXDLMSData) GetNames() []string {
	return []string{"Logical Name", "Value"}
}

// getMethodNames returns the returns names of method indexes.
func (g *GXDLMSData) GetMethodNames() []string {
	return []string{}
}

// GetAttributeCount returns the returns amount of attributes.
//
//	Returns:
//	    Count of attributes.
func (g *GXDLMSData) GetAttributeCount() int {
	return 2
}

// GetMethodCount returns the returns amount of methods.
//
//	Returns:
func (g *GXDLMSData) GetMethodCount() int {
	return 0
}

// GetValue returns the returns value of given attribute.
// settings: DLMS settings.
// e: Get parameters.
//
//	Returns:
//	    Value of the attribute index.
func (g *GXDLMSData) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	var ret any
	var err error
	switch e.Index {
	case 1:
		ret, err = helpers.LogicalNameToBytes(g.LogicalName())
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
	case 2:
		ret = g.Value
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return ret, err
}

// setValue returns the set value of given attribute.settings: DLMS settings.
// e: Set parameters.
func (g *GXDLMSData) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	var err error
	switch e.Index {
	case 1:
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return g.SetLogicalName(ln)
	case 2:
		dt, err := g.GetDataType(2)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			break
		}
		if !e.User && e.Value != nil && (dt == enums.DataTypeNone || dt == enums.DataTypeDateTime || dt == enums.DataTypeString) {
			dt2, err := internal.GetDLMSDataType(reflect.TypeOf(e.Value))
			if err != nil {
				e.Error = enums.ErrorCodeReadWriteDenied
				break
			}
			if dt != dt2 {
				g.SetDataType(2, dt2)
			}
		}
		dt = g.GetUIDataType(2)
		if dt == enums.DataTypeDateTime {
			switch e.Value.(type) {
			case uint32, uint64, int32, int64:
				e.Value = types.GXDateTimeFromUnixTime(e.Value.(int64)).Value.Unix()
			default:
			}
		}
		g.Value = e.Value
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return err
}

// GetValues returns the returns attributes as an array.
//
//	Returns:
//	    Collection of COSEM object values.
func (g *GXDLMSData) GetValues() []any {
	return []any{g.LogicalName(), g.Value}
}

// GetDataType returns the returns device data type of selected attribute index.index: Attribute index of the object.
//
//	Returns:
//	    Device data type of the object.
func (g *GXDLMSData) GetDataType(index int) (enums.DataType, error) {
	switch index {
	case 1:
		// Logical name.
		return enums.DataTypeOctetString, nil
	case 2:
		dt, err := g.GXDLMSObject.GetDataType(index)
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
	default:
		return enums.DataTypeNone, errors.New("GetDataType failed. Invalid attribute index.")
	}
}

// GetUIDataType returns UI data type of selected index.
// index: Attribute index of the object.
//
//	Returns:
//	    UseUI data type of the object.
func (g *GXDLMSData) GetUIDataType(index int) enums.DataType {
	return g.Base().GetUIDataType(index)
}

// Constructor.
// ln: Logical Name of the object.
// sn: Short Name of the object.
func NewGXDLMSData(ln string, sn int16) (*GXDLMSData, error) {
	err := ValidateLogicalName(ln)
	if err != nil {
		return nil, err
	}
	return &GXDLMSData{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeData,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}

// Load returns the load object content from XML.
// reader: XML reader.
func (g *GXDLMSData) Load(reader *GXXmlReader) error {
	var err error
	g.Value, err = reader.ReadElementContentAsObject("Value", nil, g, 2)
	return err
}

// Save returns the save object content to XML.writer: XML writer.
func (g *GXDLMSData) Save(writer *GXXmlWriter) error {
	dt, err := g.GetDataType(2)
	if err != nil {
		return err
	}
	return writer.WriteElementObject("Value", g.Value, dt, g.GetUIDataType(2))
}

func (g *GXDLMSData) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	return nil, errors.New("Invoke method is not supported in GXDLMSData object.")
}

func (g *GXDLMSData) PostLoad(reader *GXXmlReader) error {
	return nil
}
