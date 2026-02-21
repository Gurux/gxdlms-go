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
	"github.com/Gurux/gxdlms-go/dlmserrors"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// Online help:
// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSMBusMasterPortSetup
type GXDLMSMBusMasterPortSetup struct {
	GXDLMSObject
	// The communication speed supported by the port.
	CommSpeed enums.BaudRate
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSMBusMasterPortSetup) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

//Invoke returns the invokes method.
//
// Parameters:
//   settings: DLMS settings.
//   e: Invoke parameters.
func (g *GXDLMSMBusMasterPortSetup) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, nil
}

//GetAttributeIndexToRead returns the collection of attributes to read.
// If attribute is static and already read or device is returned HW error it is not returned.
//
// Parameters:
//   all: All items are returned even if they are read already.
//
// Returns:
//   Collection of attributes to read.
func (g *GXDLMSMBusMasterPortSetup) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	// LN is static and read only once.
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	// CommSpeed
	if all || g.CanRead(2) {
		attributes = append(attributes, 2)
	}
	return attributes
}

//GetNames returns the names of attribute indexes.
func (g *GXDLMSMBusMasterPortSetup) GetNames() []string {
	return []string{"Logical Name", "Comm Speed"}
}

//GetMethodNames returns the names of method indexes.
func (g *GXDLMSMBusMasterPortSetup) GetMethodNames() []string {
	return []string{}
}

//GetAttributeCount returns the amount of attributes.
//
// Returns:
//   Count of attributes.
func (g *GXDLMSMBusMasterPortSetup) GetAttributeCount() int {
	return 2
}

//GetMethodCount returns the amount of methods.
func (g *GXDLMSMBusMasterPortSetup) GetMethodCount() int {
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
func (g *GXDLMSMBusMasterPortSetup) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	if e.Index == 1 {
		return helpers.LogicalNameToBytes(g.LogicalName())
	}
	if e.Index == 2 {
		return g.CommSpeed, nil
	}
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, nil
}

//SetValue returns the set value of given attribute.
// When raw parameter us not used example register multiplies value by scalar.
//
// Parameters:
//   settings: DLMS settings.
//   e: Set parameters.
func (g *GXDLMSMBusMasterPortSetup) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	if e.Index == 1 {
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return g.SetLogicalName(ln)
	} else if e.Index == 2 {
		g.CommSpeed = enums.BaudRate(e.Value.(types.GXEnum).Value)
	} else {
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil
}

//Load returns the load object content from XML.
//
// Parameters:
//   reader: XML reader.
func (g *GXDLMSMBusMasterPortSetup) Load(reader *GXXmlReader) error {
	ret, err := reader.ReadElementContentAsInt("CommSpeed", 0)
	if err != nil {
		return err
	}
	g.CommSpeed = enums.BaudRate(ret)
	return err
}

//Save returns the save object content to XML.
//
// Parameters:
//   writer: XML writer.
func (g *GXDLMSMBusMasterPortSetup) Save(writer *GXXmlWriter) error {
	err := writer.WriteElementString("CommSpeed", int(g.CommSpeed))
	if err != nil {
		return err
	}
	return err
}

//PostLoad returns the handle actions after Load.
//
// Parameters:
//   reader: XML reader.
func (g *GXDLMSMBusMasterPortSetup) PostLoad(reader *GXXmlReader) error {
	return nil
}

//GetValues returns the an array containing the COSEM object's attribute values.
func (g *GXDLMSMBusMasterPortSetup) GetValues() []any {
	return []any{g.LogicalName(), g.CommSpeed}
}

//GetDataType returns the device data type of selected attribute index.
//
// Parameters:
//   index: Attribute index of the object.
//
// Returns:
//   Device data type of the object.
func (g *GXDLMSMBusMasterPortSetup) GetDataType(index int) (enums.DataType, error) {
	if index == 1 {
		return enums.DataTypeOctetString, nil
	}
	// CommSpeed
	if index == 2 {
		return enums.DataTypeEnum, nil
	}
	return 0, dlmserrors.ErrInvalidAttributeIndex
}

// NewGXDLMMBusMasterPortSetup creates a new M Bus Master Port Setup object instance.
//
// The function validates `ln` before creating the object.
//`ln` is the Logical Name and `sn` is the Short Name of the object.
func NewGXDLMSMBusMasterPortSetup(ln string, sn int16) (*GXDLMSMBusMasterPortSetup, error) {
	err := ValidateLogicalName(ln)
	if err != nil {
		return nil, err
	}
	return &GXDLMSMBusMasterPortSetup{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeMBusMasterPortSetup,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
