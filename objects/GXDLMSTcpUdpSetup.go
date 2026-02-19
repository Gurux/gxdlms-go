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
// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSTcpUdpSetup
type GXDLMSTcpUdpSetup struct {
	GXDLMSObject
	// TCP/UDP port number on which the physical device is
	// listening for the DLMS/COSEM application.
	Port uint16

	// References an IP setup object by its logical name. The referenced object
	// contains information about the IP Address settings of the IP layer
	// supporting the TCP-UDP layer.
	IPReference string

	// TCP can indicate the maximum receive segment size to its partner.
	MaximumSegmentSize uint16

	// The maximum number of simultaneous connections the COSEM
	// TCP/UDP based transport layer is able to support.
	MaximumSimultaneousConnections uint8

	// Defines the time, expressed in seconds over which, if no frame is
	// received from the COSEM client, the inactive TCP connection shall be aborted.
	// When this value is set to 0, this means that the inactivity_time_out is
	// not operational. In other words, a TCP connection, once established,
	// in normal conditions no power failure, etc. will never be aborted by the COSEM server.
	InactivityTimeout uint16
}

// base returns the base GXDLMSObject of the object.
func (g *GXDLMSTcpUdpSetup) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

// Invoke returns the invokes method.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Invoke parameters.
func (g *GXDLMSTcpUdpSetup) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
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
func (g *GXDLMSTcpUdpSetup) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	// LN is static and read only once.
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	// Port
	if all || !g.IsRead(2) {
		attributes = append(attributes, 2)
	}
	// IPReference
	if all || !g.IsRead(3) {
		attributes = append(attributes, 3)
	}
	// MaximumSegmentSize
	if all || !g.IsRead(4) {
		attributes = append(attributes, 4)
	}
	// MaximumSimultaneousConnections
	if all || !g.IsRead(5) {
		attributes = append(attributes, 5)
	}
	// InactivityTimeout
	if all || !g.IsRead(6) {
		attributes = append(attributes, 6)
	}
	return attributes
}

// GetNames returns the names of attribute indexes.
func (g *GXDLMSTcpUdpSetup) GetNames() []string {
	return []string{"Logical Name", "Port",
		"IP Reference", "Maximum Segment Size",
		"Maximum Simultaneous Connections",
		"Inactivity Timeout"}
}

// GetMethodNames returns the names of method indexes.
func (g *GXDLMSTcpUdpSetup) GetMethodNames() []string {
	return []string{}
}

// GetAttributeCount returns the amount of attributes.
//
// Returns:
//
//	Count of attributes.
func (g *GXDLMSTcpUdpSetup) GetAttributeCount() int {
	return 6
}

// GetMethodCount returns the amount of methods.
func (g *GXDLMSTcpUdpSetup) GetMethodCount() int {
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
func (g *GXDLMSTcpUdpSetup) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	if e.Index == 1 {
		v, err := helpers.LogicalNameToBytes(g.LogicalName())
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return v, err
	} else if e.Index == 2 {
		return g.Port, nil
	} else if e.Index == 3 {
		return helpers.LogicalNameToBytes(g.IPReference)
	} else if e.Index == 4 {
		return g.MaximumSegmentSize, nil
	} else if e.Index == 5 {
		return g.MaximumSimultaneousConnections, nil
	} else if e.Index == 6 {
		return g.InactivityTimeout, nil
	} else {
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil, nil
}

// SetValue returns the set value of given attribute.
// When raw parameter us not used example register multiplies value by scalar.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Set parameters.
func (g *GXDLMSTcpUdpSetup) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	var err error
	if e.Index == 1 {
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return g.SetLogicalName(ln)
	} else if e.Index == 2 {
		g.Port = e.Value.(uint16)
	} else if e.Index == 3 {
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			return err
		}
		g.IPReference = ln
	} else if e.Index == 4 {
		g.MaximumSegmentSize = e.Value.(uint16)
	} else if e.Index == 5 {
		g.MaximumSimultaneousConnections = e.Value.(uint8)
	} else if e.Index == 6 {
		g.InactivityTimeout = e.Value.(uint16)
	} else {
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return err
}

// Load returns the load object content from XML.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSTcpUdpSetup) Load(reader *GXXmlReader) error {
	var err error
	g.Port, err = reader.ReadElementContentAsUInt16("Port", 0)
	if err != nil {
		return err
	}
	g.IPReference, err = reader.ReadElementContentAsString("IPReference", "")
	if err != nil {
		return err
	}
	g.MaximumSegmentSize, err = reader.ReadElementContentAsUInt16("MaximumSegmentSize", 0)
	if err != nil {
		return err
	}
	g.MaximumSimultaneousConnections, err = reader.ReadElementContentAsUInt8("MaximumSimultaneousConnections", 0)
	if err != nil {
		return err
	}
	g.InactivityTimeout, err = reader.ReadElementContentAsUInt16("InactivityTimeout", 0)
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
func (g *GXDLMSTcpUdpSetup) Save(writer *GXXmlWriter) error {
	err := writer.WriteElementString("Port", g.Port)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("IPReference", g.IPReference)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("MaximumSegmentSize", g.MaximumSegmentSize)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("MaximumSimultaneousConnections", g.MaximumSimultaneousConnections)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("InactivityTimeout", g.InactivityTimeout)
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
func (g *GXDLMSTcpUdpSetup) PostLoad(reader *GXXmlReader) error {
	return nil
}

// GetValues returns the an array containing the COSEM object's attribute values.
func (g *GXDLMSTcpUdpSetup) GetValues() []any {
	return []any{g.LogicalName(), g.Port, g.IPReference, g.MaximumSegmentSize, g.MaximumSimultaneousConnections, g.InactivityTimeout}
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
func (g *GXDLMSTcpUdpSetup) GetDataType(index int) (enums.DataType, error) {
	var ret enums.DataType
	if index == 1 {
		ret = enums.DataTypeOctetString
	} else if index == 2 {
		ret = enums.DataTypeUint16
	} else if index == 3 {
		ret = enums.DataTypeOctetString
	} else if index == 4 {
		ret = enums.DataTypeUint16
	} else if index == 5 {
		ret = enums.DataTypeUint8
	} else if index == 6 {
		ret = enums.DataTypeUint16
	} else {
		return 0, errors.New("GetDataType failed. Invalid attribute index.")
	}
	return ret, nil
}

// Constructor.
// ln: Logical Name of the object.
// sn: Short Name of the object.
func NewGXDLMSTcpUdpSetup(ln string, sn int16) (*GXDLMSTcpUdpSetup, error) {
	err := ValidateLogicalName(ln)
	if err != nil {
		return nil, err
	}
	return &GXDLMSTcpUdpSetup{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeTCPUDPSetup,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
