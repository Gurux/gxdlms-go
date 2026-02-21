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
	"fmt"

	"github.com/Gurux/gxdlms-go/dlmserrors"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// Online help:
// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSIECLocalPortSetup
type GXDLMSIECLocalPortSetup struct {
	GXDLMSObject
	// Start communication mode.
	DefaultMode enums.OpticalProtocolMode

	// Default Baudrate.
	DefaultBaudrate enums.BaudRate

	// Proposed Baudrate.
	ProposedBaudrate enums.BaudRate

	// Defines the minimum time between the reception of a request
	// (end of request telegram) and the transmission of the response (begin of response telegram).
	ResponseTime enums.LocalPortResponseTime

	// Device address according to IEC 62056-21.
	DeviceAddress string

	// Password 1 according to IEC 62056-21.
	Password1 string

	// Password 2 according to IEC 62056-21.
	Password2 string

	// Password W5 reserved for national applications.
	Password5 string
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSIECLocalPortSetup) Base() *GXDLMSObject {
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
func (g *GXDLMSIECLocalPortSetup) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	// LN is static and read only once.
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	// DefaultMode
	if all || g.CanRead(2) {
		attributes = append(attributes, 2)
	}
	// DefaultBaudrate
	if all || g.CanRead(3) {
		attributes = append(attributes, 3)
	}
	// ProposedBaudrate
	if all || g.CanRead(4) {
		attributes = append(attributes, 4)
	}
	// ResponseTime
	if all || g.CanRead(5) {
		attributes = append(attributes, 5)
	}
	// DeviceAddress
	if all || g.CanRead(6) {
		attributes = append(attributes, 6)
	}
	// Password1
	if all || g.CanRead(7) {
		attributes = append(attributes, 7)
	}
	// Password2
	if all || g.CanRead(8) {
		attributes = append(attributes, 8)
	}
	// Password5
	if all || g.CanRead(9) {
		attributes = append(attributes, 9)
	}
	return attributes
}

// GetNames returns the names of attribute indexes.
func (g *GXDLMSIECLocalPortSetup) GetNames() []string {
	return []string{"Logical Name", "Default Mode", "Default Baud rate", "Proposed Baud rate", "Response Time", "Device Address", "Password 1", "Password 2", "Password 5"}
}

// GetMethodNames returns the names of method indexes.
func (g *GXDLMSIECLocalPortSetup) GetMethodNames() []string {
	return []string{}
}

// GetAttributeCount returns the amount of attributes.
//
// Returns:
//
//	Count of attributes.
func (g *GXDLMSIECLocalPortSetup) GetAttributeCount() int {
	return 9
}

// GetMethodCount returns the amount of methods.
func (g *GXDLMSIECLocalPortSetup) GetMethodCount() int {
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
func (g *GXDLMSIECLocalPortSetup) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	var ret any
	switch e.Index {
	case 1:
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		err = g.SetLogicalName(ln)
	case 2:
		ret = g.DefaultMode
	case 3:
		ret = g.DefaultBaudrate
	case 4:
		ret = g.ProposedBaudrate
	case 5:
		ret = g.ResponseTime
	case 6:
		if g.DeviceAddress == "" {
			return nil, nil
		}
		ret = []byte(g.DeviceAddress)
	case 7:
		if g.Password1 == "" {
			return nil, nil
		}
		ret = []byte(g.Password1)
	case 8:
		if g.Password2 == "" {
			return nil, nil
		}
		ret = []byte(g.Password2)
	case 9:
		if g.Password5 == "" {
			return nil, nil
		}
		ret = []byte(g.Password5)
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return ret, nil
}

// SetValue returns the set value of given attribute.
// When raw parameter us not used example register multiplies value by scalar.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Set parameters.
func (g *GXDLMSIECLocalPortSetup) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	switch e.Index {
	case 1:
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		err = g.SetLogicalName(ln)
	case 2:
		g.DefaultMode = enums.OpticalProtocolMode(e.Value.(types.GXEnum).Value)
	case 3:
		g.DefaultBaudrate = enums.BaudRate(e.Value.(types.GXEnum).Value)
	case 4:
		g.ProposedBaudrate = enums.BaudRate(e.Value.(types.GXEnum).Value)
	case 5:
		g.ResponseTime = enums.LocalPortResponseTime(e.Value.(types.GXEnum).Value)
	case 6:
		if v, ok := e.Value.([]byte); ok {
			ret, err := internal.ChangeTypeFromByteArray(settings, v, enums.DataTypeString)
			if err != nil {
				return err
			} else {
				g.DeviceAddress = ret.(string)
			}
		} else {
			g.DeviceAddress = fmt.Sprint(e.Value)
		}
	case 7:
		if v, ok := e.Value.([]byte); ok {
			g.Password1 = string(v)
		} else {
			g.Password1 = fmt.Sprint(e.Value)
		}
	case 8:
		if v, ok := e.Value.([]byte); ok {
			g.Password2 = string(v)
		} else {
			g.Password2 = fmt.Sprint(e.Value)
		}
	case 9:
		if v, ok := e.Value.([]byte); ok {
			g.Password5 = string(v)
		} else {
			g.Password5 = fmt.Sprint(e.Value)
		}
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil
}

// Invoke returns the invokes method.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Invoke parameters.
func (g *GXDLMSIECLocalPortSetup) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, nil
}

// Load returns the load object content from XML.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSIECLocalPortSetup) Load(reader *GXXmlReader) error {
	ret, err := reader.ReadElementContentAsInt("DefaultMode", 0)
	if err != nil {
		return err
	}
	g.DefaultMode = enums.OpticalProtocolMode(ret)
	ret, err = reader.ReadElementContentAsInt("DefaultBaudrate", 0)
	if err != nil {
		return err
	}
	g.DefaultBaudrate = enums.BaudRate(ret)
	ret, err = reader.ReadElementContentAsInt("ProposedBaudrate", 0)
	if err != nil {
		return err
	}
	g.ProposedBaudrate = enums.BaudRate(ret)
	ret, err = reader.ReadElementContentAsInt("ResponseTime", 0)
	if err != nil {
		return err
	}
	g.ResponseTime = enums.LocalPortResponseTime(ret)
	g.DeviceAddress, err = reader.ReadElementContentAsString("DeviceAddress", "")
	if err != nil {
		return err
	}
	g.Password1, err = reader.ReadElementContentAsString("Password1", "")
	if err != nil {
		return err
	}
	g.Password2, err = reader.ReadElementContentAsString("Password2", "")
	if err != nil {
		return err
	}
	g.Password5, err = reader.ReadElementContentAsString("Password5", "")
	return err
}

// Save returns the save object content to XML.
//
// Parameters:
//
//	writer: XML writer.
func (g *GXDLMSIECLocalPortSetup) Save(writer *GXXmlWriter) error {
	err := writer.WriteElementString("DefaultMode", int(g.DefaultMode))
	if err != nil {
		return err
	}
	err = writer.WriteElementString("DefaultBaudrate", int(g.DefaultBaudrate))
	if err != nil {
		return err
	}
	err = writer.WriteElementString("ProposedBaudrate", int(g.ProposedBaudrate))
	if err != nil {
		return err
	}
	err = writer.WriteElementString("ResponseTime", int(g.ResponseTime))
	if err != nil {
		return err
	}
	err = writer.WriteElementString("DeviceAddress", g.DeviceAddress)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("Password1", g.Password1)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("Password2", g.Password2)
	if err != nil {
		return err
	}
	return writer.WriteElementString("Password5", g.Password5)
}

// PostLoad returns the handle actions after Load.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSIECLocalPortSetup) PostLoad(reader *GXXmlReader) error {
	return nil
}

// GetValues returns the an array containing the COSEM object's attribute values.
func (g *GXDLMSIECLocalPortSetup) GetValues() []any {
	return []any{g.LogicalName(), g.DefaultMode, g.DefaultBaudrate, g.ProposedBaudrate,
		g.ResponseTime, g.DeviceAddress, g.Password1, g.Password2, g.Password5}
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
func (g *GXDLMSIECLocalPortSetup) GetDataType(index int) (enums.DataType, error) {
	switch index {
	case 1:
		return enums.DataTypeOctetString, nil
	case 2:
		return enums.DataTypeEnum, nil
	case 3:
		return enums.DataTypeEnum, nil
	case 4:
		return enums.DataTypeEnum, nil
	case 5:
		return enums.DataTypeEnum, nil
	case 6:
		return enums.DataTypeOctetString, nil
	case 7:
		return enums.DataTypeOctetString, nil
	case 8:
		return enums.DataTypeOctetString, nil
	case 9:
		return enums.DataTypeOctetString, nil
	}
	return 0, dlmserrors.ErrInvalidAttributeIndex
}

// NewGXDLMSIECLocalPortSetup creates a new IEC local port setup object instance.
//
// The function validates `ln` before creating the object.
//`ln` is the Logical Name and `sn` is the Short Name of the object.
func NewGXDLMSIECLocalPortSetup(ln string, sn int16) (*GXDLMSIECLocalPortSetup, error) {
	err := ValidateLogicalName(ln)
	if err != nil {
		return nil, err
	}
	return &GXDLMSIECLocalPortSetup{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeIecLocalPortSetup,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
