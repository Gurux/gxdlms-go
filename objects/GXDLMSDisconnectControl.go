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
	"github.com/Gurux/gxdlms-go/types"
)

// Online help:
// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSDisconnectControl
type GXDLMSDisconnectControl struct {
	GXDLMSObject
	// Output state of COSEM Disconnect Control object.
	OutputState bool

	// Output state of COSEM Disconnect Control object.
	ControlState enums.ControlState

	// Control mode of COSEM Disconnect Control object.
	ControlMode enums.ControlMode
}

// base returns the base GXDLMSObject of the object.
func (g *GXDLMSDisconnectControl) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

// Invoke returns the invokes method.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Invoke parameters.
func (g *GXDLMSDisconnectControl) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	e.Error = enums.ErrorCodeReadWriteDenied
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
func (g *GXDLMSDisconnectControl) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	// LN is static and read only once.
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	// OutputState
	if all || g.CanRead(2) {
		attributes = append(attributes, 2)
	}
	// ControlState
	if all || g.CanRead(3) {
		attributes = append(attributes, 3)
	}
	// ControlMode
	if all || g.CanRead(4) {
		attributes = append(attributes, 4)
	}
	return attributes
}

// GetNames returns the names of attribute indexes.
func (g *GXDLMSDisconnectControl) GetNames() []string {
	return []string{"Logical Name", "Output State", "Control State", "Control Mode"}
}

// GetMethodNames returns the names of method indexes.
func (g *GXDLMSDisconnectControl) GetMethodNames() []string {
	return []string{"Remote disconnect", "Remote reconnect"}
}

// GetAttributeCount returns the amount of attributes.
//
// Returns:
//
//	Count of attributes.
func (g *GXDLMSDisconnectControl) GetAttributeCount() int {
	return 4
}

// GetMethodCount returns the amount of methods.
func (g *GXDLMSDisconnectControl) GetMethodCount() int {
	return 2
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
func (g *GXDLMSDisconnectControl) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	if e.Index == 1 {
		v, err := helpers.LogicalNameToBytes(g.LogicalName())
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return v, err
	}
	if e.Index == 2 {
		return g.OutputState, nil
	}
	if e.Index == 3 {
		return g.ControlState, nil
	}
	if e.Index == 4 {
		return g.ControlMode, nil
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
func (g *GXDLMSDisconnectControl) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	if e.Index == 1 {
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		err = g.SetLogicalName(ln)
	} else if e.Index == 2 {
		g.OutputState = e.Value.(bool)
	} else if e.Index == 3 {
		g.ControlState = enums.ControlState(e.Value.(types.GXEnum).Value)
	} else if e.Index == 4 {
		g.ControlMode = enums.ControlMode(e.Value.(types.GXEnum).Value)
	} else {
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil
}

// Load returns the load object content from XML.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSDisconnectControl) Load(reader *GXXmlReader) error {
	var err error
	ret, err := reader.ReadElementContentAsInt("OutputState", 0)
	if err != nil {
		return err
	}
	g.OutputState = ret != 0
	ret, err = reader.ReadElementContentAsInt("ControlState", 0)
	if err != nil {
		return err
	}
	g.ControlState = enums.ControlState(ret)
	ret, err = reader.ReadElementContentAsInt("ControlMode", 0)
	if err != nil {
		return err
	}
	g.ControlMode = enums.ControlMode(ret)
	return err
}

// Save returns the save object content to XML.
//
// Parameters:
//
//	writer: XML writer.
func (g *GXDLMSDisconnectControl) Save(writer *GXXmlWriter) error {
	err := writer.WriteElementString("OutputState", g.OutputState)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("ControlState", int(g.ControlState))
	if err != nil {
		return err
	}
	return writer.WriteElementString("ControlMode", int(g.ControlMode))
}

// PostLoad returns the handle actions after Load.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSDisconnectControl) PostLoad(reader *GXXmlReader) error {
	return nil
}

// RemoteDisconnect returns the forces the disconnect control object into 'disconnected' state
// if remote disconnection is enabled(control mode > 0).
//
// Parameters:
//
//	client: DLMS client.
//
// Returns:
//
//	Action bytes.
func (g *GXDLMSDisconnectControl) RemoteDisconnect(client IGXDLMSClient) ([][]uint8, error) {
	return client.Method(g, 1, int8(0), enums.DataTypeInt8)
}

// RemoteReconnect returns the forces the disconnect control object into the 'ready_for_reconnection'
// state if a direct remote reconnection is disabled(control_mode = 1, 3, 5, 6).
// Forces the disconnect control object into the 'connected' state if
// a direct remote reconnection is enabled(control_mode = 2, 4).
//
// Parameters:
//
//	client: DLMS client.
//
// Returns:
//
//	Action bytes.
func (g *GXDLMSDisconnectControl) RemoteReconnect(client IGXDLMSClient) ([][]uint8, error) {
	return client.Method(g, 2, int8(0), enums.DataTypeInt8)
}

// GetValues returns the an array containing the COSEM object's attribute values.
func (g *GXDLMSDisconnectControl) GetValues() []any {
	return []any{g.LogicalName(), g.OutputState, g.ControlState, g.ControlMode}
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
func (g *GXDLMSDisconnectControl) GetDataType(index int) (enums.DataType, error) {
	if index == 1 {
		return enums.DataTypeOctetString, nil
	}
	if index == 2 {
		return enums.DataTypeBoolean, nil
	}
	if index == 3 {
		return enums.DataTypeEnum, nil
	}
	if index == 4 {
		return enums.DataTypeEnum, nil
	}
	return 0, errors.New("GetDataType failed. Invalid attribute index.")
}

// Constructor.
// ln: Logical Name of the object.
// sn: Short Name of the object.
func NewGXDLMSDisconnectControl(ln string, sn int16) (*GXDLMSDisconnectControl, error) {
	err := ValidateLogicalName(ln)
	if err != nil {
		return nil, err
	}
	return &GXDLMSDisconnectControl{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeDisconnectControl,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
