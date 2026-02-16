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
	"strings"

	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// Online help:
// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSMacAddressSetup
type GXDLMSMacAddressSetup struct {
	GXDLMSObject
	// Value of COSEM Data object.
	MacAddress string
}

// base returns the base GXDLMSObject of the object.
func (g *GXDLMSMacAddressSetup) Base() *GXDLMSObject {
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
func (g *GXDLMSMacAddressSetup) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	// LN is static and read only once.
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	// MacAddress
	if all || !g.CanRead(2) {
		attributes = append(attributes, 2)
	}
	return attributes
}

// GetNames returns the names of attribute indexes.
func (g *GXDLMSMacAddressSetup) GetNames() []string {
	return []string{"Logical Name", "MAC Address"}
}

// GetMethodNames returns the names of method indexes.
func (g *GXDLMSMacAddressSetup) GetMethodNames() []string {
	return []string{}
}

// GetAttributeCount returns the amount of attributes.
//
// Returns:
//
//	Count of attributes.
func (g *GXDLMSMacAddressSetup) GetAttributeCount() int {
	return 2
}

// GetMethodCount returns the amount of methods.
func (g *GXDLMSMacAddressSetup) GetMethodCount() int {
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
func (g *GXDLMSMacAddressSetup) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	var ret any
	var err error
	switch e.Index {
	case 1:
		ret, err = helpers.LogicalNameToBytes(g.LogicalName())
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
	case 2:
		ret = types.HexToBytes(g.MacAddress)
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
		err = errors.New("GetValue failed. Invalid attribute index.")
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
func (g *GXDLMSMacAddressSetup) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	var err error
	switch e.Index {
	case 1:
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		err = g.SetLogicalName(ln)
	case 2:
		if _, ok := e.Value.([]byte); ok {
			g.MacAddress = types.ToHex(e.Value.([]byte), true)
			g.MacAddress = strings.ReplaceAll(g.MacAddress, " ", ":")
		} else {
			g.MacAddress = types.ToHex(types.HexToBytes(string(e.Value.([]byte))), true)
			g.MacAddress = strings.ReplaceAll(g.MacAddress, " ", ":")
		}
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return err
}

// Invoke returns the invokes method.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Invoke parameters.
func (g *GXDLMSMacAddressSetup) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	return nil, errors.New("Invoke method is not supported in GXDLMSMacAddressSetup object.")
}

// Load returns the load object content from XML.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSMacAddressSetup) Load(reader *GXXmlReader) error {
	var err error
	g.MacAddress, err = reader.ReadElementContentAsString("MacAddress", "")
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
func (g *GXDLMSMacAddressSetup) Save(writer *GXXmlWriter) error {
	err := writer.WriteElementString("MacAddress", g.MacAddress)
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
func (g *GXDLMSMacAddressSetup) PostLoad(reader *GXXmlReader) error {
	return nil
}

// GetValues returns the an array containing the COSEM object's attribute values.
func (g *GXDLMSMacAddressSetup) GetValues() []any {
	return []any{g.LogicalName, g.MacAddress}
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
func (g *GXDLMSMacAddressSetup) GetDataType(index int) (enums.DataType, error) {
	if index == 1 {
		return enums.DataTypeOctetString, nil
	}
	if index == 2 {
		return enums.DataTypeOctetString, nil
	}
	return enums.DataTypeNone, errors.New("GetDataType failed. Invalid attribute index.")
}

// Constructor.
// ln: Logical Name of the object.
// sn: Short Name of the object.
func NewGXDLMSMacAddressSetup(ln string, sn int16) (*GXDLMSMacAddressSetup, error) {
	err := ValidateLogicalName(ln)
	if err != nil {
		return nil, err
	}
	return &GXDLMSMacAddressSetup{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeMacAddressSetup,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
