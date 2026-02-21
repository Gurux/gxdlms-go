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

	"github.com/Gurux/gxdlms-go/dlmserrors"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// Online help:
// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSUtilityTables
type GXDLMSUtilityTables struct {
	GXDLMSObject
	// Table Id.
	TableId uint16

	// Contents of the table
	Buffer []byte
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSUtilityTables) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

func (g *GXDLMSUtilityTables) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	return nil, errors.New("Invoke method is not supported in GXDLMSUtilityTables object.")
}

//GetAttributeIndexToRead returns the collection of attributes to read.
// If attribute is static and already read or device is returned HW error it is not returned.
//
// Parameters:
//   all: All items are returned even if they are read already.
//
// Returns:
//   Collection of attributes to read.
func (g *GXDLMSUtilityTables) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	// LN is static and read only once.
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	// TableId
	if all || g.CanRead(2) {
		attributes = append(attributes, 2)
	}
	// Length
	if all || g.CanRead(3) {
		attributes = append(attributes, 3)
	}
	// Buffer
	if all || g.CanRead(4) {
		attributes = append(attributes, 4)
	}
	return attributes
}

//GetNames returns the names of attribute indexes.
func (g *GXDLMSUtilityTables) GetNames() []string {
	return []string{"Logical Name", "Table Id", "Length", "Buffer"}
}

//GetMethodNames returns the names of method indexes.
func (g *GXDLMSUtilityTables) GetMethodNames() []string {
	return []string{}
}

//GetAttributeCount returns the amount of attributes.
//
// Returns:
//   Count of attributes.
func (g *GXDLMSUtilityTables) GetAttributeCount() int {
	return 4
}

//GetMethodCount returns the amount of methods.
func (g *GXDLMSUtilityTables) GetMethodCount() int {
	return 0
}

//GetValue returns the value of given attribute.
// When raw parameter us not used example register multiplies value by scalar.
//
// Parameters:
//   settings: DLMS settings.
//   e: Get parameters.
//
// Returns:
//   Value of the attribute index.
func (g *GXDLMSUtilityTables) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	switch e.Index {
	case 1:
		return helpers.LogicalNameToBytes(g.Base().LogicalName())
	case 2:
		return g.TableId, nil
	case 3:
		return len(g.Buffer), nil
	case 4:
		return g.Buffer, nil
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
		break
	}
	return nil, nil
}

//SetValue returns the set value of given attribute.
// When raw parameter us not used example register multiplies value by scalar.
//
// Parameters:
//   settings: DLMS settings.
//   e: Set parameters.
func (g *GXDLMSUtilityTables) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	switch e.Index {
	case 1:
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return g.SetLogicalName(ln)
	case 2:
		g.TableId = e.Value.(uint16)
	case 3:
	case 4:
		if v, ok := e.Value.(string); ok {
			g.Buffer = types.HexToBytes(v)
		} else {
			g.Buffer = e.Value.([]byte)
		}
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil
}

//Load returns the load object content from XML.
//
// Parameters:
//   reader: XML reader.
func (g *GXDLMSUtilityTables) Load(reader *GXXmlReader) error {
	var err error
	g.TableId, err = reader.ReadElementContentAsUInt16("Id", 0)
	if err != nil {
		return err
	}
	ret, err := reader.ReadElementContentAsString("Buffer", "")
	if err != nil {
		return err
	}
	g.Buffer = types.HexToBytes(ret)
	return err
}

//Save returns the save object content to XML.
//
// Parameters:
//   writer: XML writer.
func (g *GXDLMSUtilityTables) Save(writer *GXXmlWriter) error {
	err := writer.WriteElementString("Id", g.TableId)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("Buffer", types.ToHex(g.Buffer, true))
	if err != nil {
		return err
	}
	return err
}

//PostLoad returns the handle actions after Load.
//
// Parameters:
//   reader: XML reader.
func (g *GXDLMSUtilityTables) PostLoad(reader *GXXmlReader) error {
	return nil
}

//GetValues returns the an array containing the COSEM object's attribute values.
func (g *GXDLMSUtilityTables) GetValues() []any {
	return []any{g.LogicalName(), g.TableId, uint16(len(g.Buffer)), g.Buffer}
}

//GetDataType returns the device data type of selected attribute index.
//
// Parameters:
//   index: Attribute index of the object.
//
// Returns:
//   Device data type of the object.
func (g *GXDLMSUtilityTables) GetDataType(index int) (enums.DataType, error) {
	switch index {
	case 1:
		return enums.DataTypeOctetString, nil
	case 2:
		return enums.DataTypeUint16, nil
	case 3:
		return enums.DataTypeUint32, nil
	case 4:
		return enums.DataTypeOctetString, nil
	default:
		return 0, dlmserrors.ErrInvalidAttributeIndex
	}
}

// NewGXDLMSUtilityTables creates a new Utility Tables object instance.
//
// The var attributes []int` before creating the object.
//`ln` is the Logical Name and `sn` is the Short Name of the object.
func NewGXDLMSUtilityTables(ln string, sn int16) (*GXDLMSUtilityTables, error) {
	err := ValidateLogicalName(ln)
	if err != nil {
		return nil, err
	}
	return &GXDLMSUtilityTables{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeUtilityTables,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
