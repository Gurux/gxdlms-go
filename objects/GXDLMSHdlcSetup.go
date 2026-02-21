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
// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSHdlcSetup
type GXDLMSHdlcSetup struct {
	GXDLMSObject
	windowSizeTransmit        uint8
	windowSizeReceive         uint8
	maximumInfoLengthTransmit uint16
	maximumInfoLengthReceive  uint16
	CommunicationSpeed        enums.BaudRate
	InterCharachterTimeout    uint16
	InactivityTimeout         uint16
	DeviceAddress             uint16
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSHdlcSetup) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

func (g *GXDLMSHdlcSetup) WindowSizeTransmit() uint8 {
	return g.windowSizeTransmit
}

func (g *GXDLMSHdlcSetup) SetWindowSizeTransmit(value uint8) error {
	if value > 7 {
		return errors.New("WindowSizeTransmit")
	}
	g.windowSizeTransmit = value
	return nil
}

func (g *GXDLMSHdlcSetup) WindowSizeReceive() uint8 {
	return g.windowSizeReceive
}

func (g *GXDLMSHdlcSetup) SetWindowSizeReceive(value uint8) error {
	if value > 7 {
		return errors.New("WindowSizeReceive")
	}
	g.windowSizeReceive = value
	return nil
}

func (g *GXDLMSHdlcSetup) MaximumInfoLengthTransmit() uint16 {
	return g.maximumInfoLengthTransmit
}

func (g *GXDLMSHdlcSetup) SetMaximumInfoLengthTransmit(value uint16) error {
	if value > 2030 {
		return errors.New("MaximumInfoLengthTransmit")
	}
	g.maximumInfoLengthTransmit = value
	return nil
}

func (g *GXDLMSHdlcSetup) MaximumInfoLengthReceive() uint16 {
	return g.maximumInfoLengthReceive
}

func (g *GXDLMSHdlcSetup) SetMaximumInfoLengthReceive(value uint16) error {
	if value > 2030 {
		return errors.New("MaximumInfoLengthReceive")
	}
	g.maximumInfoLengthReceive = value
	return nil
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
func (g *GXDLMSHdlcSetup) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	// LN is static and read only once.
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	// CommunicationSpeed
	if all || !g.CanRead(2) {
		attributes = append(attributes, 2)
	}
	// WindowSizeTransmit
	if all || !g.CanRead(3) {
		attributes = append(attributes, 3)
	}
	// WindowSizeReceive
	if all || !g.CanRead(4) {
		attributes = append(attributes, 4)
	}
	// MaximumInfoLengthTransmit
	if all || !g.CanRead(5) {
		attributes = append(attributes, 5)
	}
	// MaximumInfoLengthReceive
	if all || !g.CanRead(6) {
		attributes = append(attributes, 6)
	}
	// InterCharachterTimeout
	if all || !g.CanRead(7) {
		attributes = append(attributes, 7)
	}
	// InactivityTimeout
	if all || !g.CanRead(8) {
		attributes = append(attributes, 8)
	}
	// DeviceAddress
	if all || !g.CanRead(9) {
		attributes = append(attributes, 9)
	}
	return attributes
}

// GetNames returns the names of attribute indexes.
func (g *GXDLMSHdlcSetup) GetNames() []string {
	return []string{"Logical Name", "Communication Speed", "Window Size Transmit", "Window Size Receive", "Maximum Info Length Transmit", "Maximum Info Length Receive", "InterCharachter Timeout", "Inactivity Timeout", "Device Address"}
}

// GetMethodNames returns the names of method indexes.
func (g *GXDLMSHdlcSetup) GetMethodNames() []string {
	return []string{}
}

// GetAttributeCount returns the amount of attributes.
//
// Returns:
//
//	Count of attributes.
func (g *GXDLMSHdlcSetup) GetAttributeCount() int {
	return 9
}

// GetMethodCount returns the amount of methods.
func (g *GXDLMSHdlcSetup) GetMethodCount() int {
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
func (g *GXDLMSHdlcSetup) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	var ret any
	var err error
	if e.Index == 1 {
		ret, err = helpers.LogicalNameToBytes(g.LogicalName())
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
	}
	if e.Index == 2 {
		ret = g.CommunicationSpeed
	} else if e.Index == 3 {
		ret = g.windowSizeTransmit
	} else if e.Index == 4 {
		ret = g.windowSizeReceive
	} else if e.Index == 5 {
		ret = g.maximumInfoLengthTransmit
	} else if e.Index == 6 {
		ret = g.maximumInfoLengthReceive
	} else if e.Index == 7 {
		ret = g.InterCharachterTimeout
	} else if e.Index == 8 {
		ret = g.InactivityTimeout
	} else if e.Index == 9 {
		ret = g.DeviceAddress
	} else {
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
func (g *GXDLMSHdlcSetup) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	var err error
	if e.Index == 1 {
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		err = g.SetLogicalName(ln)
	} else if e.Index == 2 {
		g.CommunicationSpeed = enums.BaudRate(e.Value.(types.GXEnum).Value)
	} else if e.Index == 3 {
		g.windowSizeTransmit = e.Value.(byte)
	} else if e.Index == 4 {
		g.windowSizeReceive = e.Value.(byte)
	} else if e.Index == 5 {
		g.maximumInfoLengthTransmit = e.Value.(uint16)
	} else if e.Index == 6 {
		g.maximumInfoLengthReceive = e.Value.(uint16)
	} else if e.Index == 7 {
		g.InterCharachterTimeout = e.Value.(uint16)
	} else if e.Index == 8 {
		g.InactivityTimeout = e.Value.(uint16)
	} else if e.Index == 9 {
		g.DeviceAddress = e.Value.(uint16)
	} else {
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
func (g *GXDLMSHdlcSetup) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	return nil, errors.New("Invoke method is not supported in GXDLMSGprsSetup object.")
}

// Load returns the load object content from XML.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSHdlcSetup) Load(reader *GXXmlReader) error {
	ret, err := reader.ReadElementContentAsInt("Speed", 0)
	if err != nil {
		return err
	}
	g.CommunicationSpeed = enums.BaudRate(ret)
	ret, err = reader.ReadElementContentAsInt("WindowSizeTx", 0)
	if err != nil {
		return err
	}
	g.windowSizeTransmit = uint8(ret)
	ret, err = reader.ReadElementContentAsInt("WindowSizeRx", 0)
	if err != nil {
		return err
	}
	g.windowSizeReceive = uint8(ret)
	ret, err = reader.ReadElementContentAsInt("MaximumInfoLengthTx", 0)
	if err != nil {
		return err
	}
	g.maximumInfoLengthTransmit = uint16(ret)
	ret, err = reader.ReadElementContentAsInt("MaximumInfoLengthRx", 0)
	if err != nil {
		return err
	}
	g.maximumInfoLengthReceive = uint16(ret)
	ret, err = reader.ReadElementContentAsInt("InterCharachterTimeout", 0)
	if err != nil {
		return err
	}
	g.InterCharachterTimeout = uint16(ret)
	ret, err = reader.ReadElementContentAsInt("InactivityTimeout", 0)
	if err != nil {
		return err
	}
	g.InactivityTimeout = uint16(ret)
	ret, err = reader.ReadElementContentAsInt("DeviceAddress", 0)
	if err != nil {
		return err
	}
	g.DeviceAddress = uint16(ret)
	return err
}

// Save returns the save object content to XML.
//
// Parameters:
//
//	writer: XML writer.
func (g *GXDLMSHdlcSetup) Save(writer *GXXmlWriter) error {
	err := writer.WriteElementString("Speed", int(g.CommunicationSpeed))
	if err != nil {
		return err
	}
	err = writer.WriteElementString("WindowSizeTx", g.windowSizeTransmit)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("WindowSizeRx", g.windowSizeReceive)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("MaximumInfoLengthTx", g.maximumInfoLengthTransmit)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("MaximumInfoLengthRx", g.maximumInfoLengthReceive)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("InterCharachterTimeout", g.InterCharachterTimeout)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("InactivityTimeout", g.InactivityTimeout)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("DeviceAddress", g.DeviceAddress)
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
func (g *GXDLMSHdlcSetup) PostLoad(reader *GXXmlReader) error {
	return nil
}

// GetValues returns the an array containing the COSEM object's attribute values.
func (g *GXDLMSHdlcSetup) GetValues() []any {
	return []any{g.LogicalName(), g.CommunicationSpeed, g.windowSizeTransmit, g.windowSizeReceive,
		g.maximumInfoLengthTransmit, g.maximumInfoLengthReceive,
		g.InterCharachterTimeout, g.InactivityTimeout, g.DeviceAddress}
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
func (g *GXDLMSHdlcSetup) GetDataType(index int) (enums.DataType, error) {
	var dt enums.DataType
	if index == 1 {
		dt = enums.DataTypeOctetString
	} else if index == 2 {
		dt = enums.DataTypeEnum
	} else if index == 3 {
		dt = enums.DataTypeUint8
	} else if index == 4 {
		dt = enums.DataTypeUint8
	} else if index == 5 {
		dt = enums.DataTypeUint16
	} else if index == 6 {
		dt = enums.DataTypeUint16
	} else if index == 7 {
		dt = enums.DataTypeUint16
	} else if index == 8 {
		dt = enums.DataTypeUint16
	} else if index == 9 {
		dt = enums.DataTypeUint16
	} else {
		return enums.DataTypeNone, dlmserrors.ErrInvalidAttributeIndex
	}
	return dt, nil
}

// NewGXDLMSHdlcSetup creates a new HDLC setup object instance.
//
// The function validates `ln` before creating the object.
//`ln` is the Logical Name and `sn` is the Short Name of the object.
func NewGXDLMSHdlcSetup(ln string, sn int16) (*GXDLMSHdlcSetup, error) {
	err := ValidateLogicalName(ln)
	if err != nil {
		return nil, err
	}
	return &GXDLMSHdlcSetup{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeIecHdlcSetup,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
