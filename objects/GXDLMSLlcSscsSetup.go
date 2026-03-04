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

	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
)

// Online help:
// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSLlcSscsSetup
type GXDLMSLlcSscsSetup struct {
	GXDLMSObject
	// Address assigned to the service node during its registration by the base node.
	ServiceNodeAddress uint16

	// Base node address to which the service node is registered.
	BaseNodeAddress uint16
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSLlcSscsSetup) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

// Invoke returns the invokes method.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Invoke parameters.
func (g *GXDLMSLlcSscsSetup) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	if e.Index == 1 {
		g.ServiceNodeAddress = 0xFFE
		g.BaseNodeAddress = 0
	} else {
		e.Error = enums.ErrorCodeReadWriteDenied
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
func (g *GXDLMSLlcSscsSetup) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	// LN is static and read only once.
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	// ServiceNodeAddress
	if all || g.CanRead(2) {
		attributes = append(attributes, 2)
	}
	// BaseNodeAddress
	if all || g.CanRead(3) {
		attributes = append(attributes, 3)
	}
	return attributes
}

// GetNames returns the names of attribute indexes.
func (g *GXDLMSLlcSscsSetup) GetNames() []string {
	return []string{"Logical Name", "ServiceNodeAddress", "BaseNodeAddress"}
}

// GetMethodNames returns the names of method indexes.
func (g *GXDLMSLlcSscsSetup) GetMethodNames() []string {
	return []string{"Reset"}
}

// GetAttributeCount returns the amount of attributes.
//
// Returns:
//
//	Count of attributes.
func (g *GXDLMSLlcSscsSetup) GetAttributeCount() int {
	return 3
}

// GetMethodCount returns the amount of methods.
func (g *GXDLMSLlcSscsSetup) GetMethodCount() int {
	return 1
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
func (g *GXDLMSLlcSscsSetup) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	var ret any
	var err error
	switch e.Index {
	case 1:
		ret, err = helpers.LogicalNameToBytes(g.LogicalName())
	case 2:
		ret = g.ServiceNodeAddress
	case 3:
		ret = g.BaseNodeAddress
	default:
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
func (g *GXDLMSLlcSscsSetup) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	switch e.Index {
	case 1:
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return g.SetLogicalName(ln)
	case 2:
		g.ServiceNodeAddress = e.Value.(uint16)
	case 3:
		g.BaseNodeAddress = e.Value.(uint16)
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil
}

// Load returns the load object content from XML.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSLlcSscsSetup) Load(reader *GXXmlReader) error {
	var err error
	g.ServiceNodeAddress, err = reader.ReadElementContentAsUInt16("ServiceNodeAddress", 0)
	if err != nil {
		return err
	}
	g.BaseNodeAddress, err = reader.ReadElementContentAsUInt16("BaseNodeAddress", 0)
	if err != nil {
		return err
	}
	return err
}

// Save returns the save object content to XML.
//
// Parameters:
//
//	writer: XML writer.
func (g *GXDLMSLlcSscsSetup) Save(writer *GXXmlWriter) error {
	err := writer.WriteElementString("ServiceNodeAddress", g.ServiceNodeAddress)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("BaseNodeAddress", g.BaseNodeAddress)
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
func (g *GXDLMSLlcSscsSetup) PostLoad(reader *GXXmlReader) error {
	return nil
}

// GetValues returns the an array containing the COSEM object's attribute values.
func (g *GXDLMSLlcSscsSetup) GetValues() []any {
	return []any{g.LogicalName(), g.ServiceNodeAddress, g.BaseNodeAddress}
}

// Reset returns the deallocating the service node address.
// The value of the ServiceNodeAddress becomes NEW and the value of the BaseNodeAddress becomes 0.
//
// Parameters:
//
//	client: DLMS client.
//
// Returns:
//
//	Action bytes.
func (g *GXDLMSLlcSscsSetup) Reset(client IGXDLMSClient) ([][]uint8, error) {
	return client.Method(g, 1, int8(0), enums.DataTypeInt8)
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
func (g *GXDLMSLlcSscsSetup) GetDataType(index int) (enums.DataType, error) {
	var ret enums.DataType
	var err error
	switch index {
	case 1:
		ret = enums.DataTypeOctetString
	case 2:
		ret = enums.DataTypeUint16
	case 3:
		ret = enums.DataTypeUint16
	default:
		err = errors.New("GetDataType failed. Invalid attribute index.")
	}
	return ret, err
}

// NewGXDLMSLlcSscsSetup creates a new Llc Sscs Setup object instance.
//
// The function validates `ln` before creating the object.
// `ln` is the Logical Name and `sn` is the Short Name of the object.
func NewGXDLMSLlcSscsSetup(ln string, sn int16) (*GXDLMSLlcSscsSetup, error) {
	err := ValidateLogicalName(ln)
	if err != nil {
		return nil, err
	}
	return &GXDLMSLlcSscsSetup{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeLlcSscsSetup,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
