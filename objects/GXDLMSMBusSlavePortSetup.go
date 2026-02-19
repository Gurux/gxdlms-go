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

// Model and configure communication channels.
// Online help:
// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSMBusSlavePortSetup
type GXDLMSMBusSlavePortSetup struct {
	GXDLMSObject
	// Defines the baud rate for the opening sequence.
	DefaultBaud enums.BaudRate

	// Defines the baud rate for the opening sequence.
	AvailableBaud enums.BaudRate

	// Defines whether or not the device has been assigned an address
	// since last power up of the device.
	AddressState enums.AddressState

	// The currently assigned device address.
	BusAddress uint8
}

// base returns the base GXDLMSObject of the object.
func (g *GXDLMSMBusSlavePortSetup) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

// Invoke returns the invokes method.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Invoke parameters.
func (g *GXDLMSMBusSlavePortSetup) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
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
func (g *GXDLMSMBusSlavePortSetup) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	// LN is static and read only once.
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	// DefaultBaud
	if all || !g.IsRead(2) {
		attributes = append(attributes, 2)
	}
	// AvailableBaud
	if all || !g.IsRead(3) {
		attributes = append(attributes, 3)
	}
	// AddressState
	if all || !g.IsRead(4) {
		attributes = append(attributes, 4)
	}
	// BusAddress
	if all || !g.IsRead(5) {
		attributes = append(attributes, 5)
	}
	return attributes
}

// GetNames returns the names of attribute indexes.
func (g *GXDLMSMBusSlavePortSetup) GetNames() []string {
	return []string{"Logical Name", "Default Baud Rate", "Available Baud rate", "Address State", "Bus Address"}
}

// GetMethodNames returns the names of method indexes.
func (g *GXDLMSMBusSlavePortSetup) GetMethodNames() []string {
	return []string{}
}

// GetAttributeCount returns the amount of attributes.
//
// Returns:
//
//	Count of attributes.
func (g *GXDLMSMBusSlavePortSetup) GetAttributeCount() int {
	return 5
}

// GetMethodCount returns the amount of methods.
func (g *GXDLMSMBusSlavePortSetup) GetMethodCount() int {
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
func (g *GXDLMSMBusSlavePortSetup) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	if e.Index == 1 {
		v, err := helpers.LogicalNameToBytes(g.LogicalName())
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return v, err
	}
	if e.Index == 2 {
		return g.DefaultBaud, nil
	}
	if e.Index == 3 {
		return g.AvailableBaud, nil
	}
	if e.Index == 4 {
		return g.AddressState, nil
	}
	if e.Index == 5 {
		return g.BusAddress, nil
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
func (g *GXDLMSMBusSlavePortSetup) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	if e.Index == 1 {
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		err = g.SetLogicalName(ln)
	} else if e.Index == 2 {
		if e.Value == nil {
			g.DefaultBaud = enums.BaudRateBaudrate300
		} else {
			g.DefaultBaud = enums.BaudRate(e.Value.(types.GXEnum).Value)
		}
	} else if e.Index == 3 {
		if e.Value == nil {
			g.AvailableBaud = enums.BaudRateBaudrate300
		} else {
			g.AvailableBaud = enums.BaudRate(e.Value.(types.GXEnum).Value)
		}
	} else if e.Index == 4 {
		if e.Value == nil {
			g.AddressState = enums.AddressStateNone
		} else {
			g.AddressState = enums.AddressState(e.Value.(types.GXEnum).Value)
		}
	} else if e.Index == 5 {
		if e.Value == nil {
			g.BusAddress = 0
		} else {
			g.BusAddress = e.Value.(uint8)
		}
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
func (g *GXDLMSMBusSlavePortSetup) Load(reader *GXXmlReader) error {
	ret, err := reader.ReadElementContentAsInt("DefaultBaud", 0)
	if err != nil {
		return err
	}
	g.DefaultBaud = enums.BaudRate(ret)
	ret, err = reader.ReadElementContentAsInt("AvailableBaud", 0)
	if err != nil {
		return err
	}
	g.AvailableBaud = enums.BaudRate(ret)
	ret, err = reader.ReadElementContentAsInt("AddressState", 0)
	if err != nil {
		return err
	}
	g.AddressState = enums.AddressState(ret)
	ret, err = reader.ReadElementContentAsInt("BusAddress", 0)
	if err != nil {
		return err
	}
	g.BusAddress = uint8(ret)
	return nil
}

// Save returns the save object content to XML.
//
// Parameters:
//
//	writer: XML writer.
func (g *GXDLMSMBusSlavePortSetup) Save(writer *GXXmlWriter) error {
	err := writer.WriteElementString("DefaultBaud", int(g.DefaultBaud))
	if err != nil {
		return err
	}
	err = writer.WriteElementString("AvailableBaud", int(g.AvailableBaud))
	if err != nil {
		return err
	}
	err = writer.WriteElementString("AddressState", int(g.AddressState))
	if err != nil {
		return err
	}
	err = writer.WriteElementString("BusAddress", g.BusAddress)
	return err
}

// PostLoad returns the handle actions after Load.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSMBusSlavePortSetup) PostLoad(reader *GXXmlReader) error {
	return nil
}

// GetValues returns the an array containing the COSEM object's attribute values.
func (g *GXDLMSMBusSlavePortSetup) GetValues() []any {
	return []any{g.LogicalName(), g.DefaultBaud, g.AvailableBaud, g.AddressState, g.BusAddress}
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
func (g *GXDLMSMBusSlavePortSetup) GetDataType(index int) (enums.DataType, error) {
	if index == 1 {
		return enums.DataTypeOctetString, nil
	}
	if index == 2 {
		return enums.DataTypeEnum, nil
	}
	if index == 3 {
		return enums.DataTypeEnum, nil
	}
	if index == 4 {
		return enums.DataTypeEnum, nil
	}
	if index == 5 {
		return enums.DataTypeUint8, nil
	}
	return 0, errors.New("GetDataType failed. Invalid attribute index.")
}

// Constructor.
// ln: Logical Name of the object.
// sn: Short Name of the object.
func NewGXDLMSMBusSlavePortSetup(ln string, sn int16) (*GXDLMSMBusSlavePortSetup, error) {
	err := ValidateLogicalName(ln)
	if err != nil {
		return nil, err
	}
	return &GXDLMSMBusSlavePortSetup{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeMBusSlavePortSetup,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
