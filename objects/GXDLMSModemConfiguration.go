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
	"strings"

	"github.com/Gurux/gxdlms-go/dlmserrors"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// Online help:
// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSModemConfiguration
type GXDLMSModemConfiguration struct {
	GXDLMSObject
	CommunicationSpeed enums.BaudRate

	InitialisationStrings []GXDLMSModemInitialisation

	ModemProfile []string
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSModemConfiguration) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

func (g *GXDLMSModemConfiguration) DefaultProfiles() []string {
	return []string{"OK", "CONNECT", "RING", "NO CARRIER", "ERROR", "CONNECT 1200", "NO DIAL TONE", "BUSY", "NO ANSWER", "CONNECT 600", "CONNECT 2400", "CONNECT 4800", "CONNECT 9600", "CONNECT 14 400", "CONNECT 28 800", "CONNECT 33 600", "CONNECT 56 000"}
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
func (g *GXDLMSModemConfiguration) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	// LN is static and read only once.
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	// CommunicationSpeed
	if all || !g.IsRead(2) {
		attributes = append(attributes, 2)
	}
	// InitialisationStrings
	if all || !g.IsRead(3) {
		attributes = append(attributes, 3)
	}
	// ModemProfile
	if all || !g.IsRead(4) {
		attributes = append(attributes, 4)
	}
	return attributes
}

// GetNames returns the names of attribute indexes.
func (g *GXDLMSModemConfiguration) GetNames() []string {
	return []string{"Logical Name", "Communication Speed", "Initialisation Strings", "Modem Profile"}
}

// GetMethodNames returns the names of method indexes.
func (g *GXDLMSModemConfiguration) GetMethodNames() []string {
	return []string{}
}

// GetAttributeCount returns the amount of attributes.
//
// Returns:
//
//	Count of attributes.
func (g *GXDLMSModemConfiguration) GetAttributeCount() int {
	return 4
}

// GetMethodCount returns the amount of methods.
func (g *GXDLMSModemConfiguration) GetMethodCount() int {
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
func (g *GXDLMSModemConfiguration) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	var err error
	if e.Index == 1 {
		v, err := helpers.LogicalNameToBytes(g.LogicalName())
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return v, err
	}
	if e.Index == 2 {
		return g.CommunicationSpeed, nil
	}
	if e.Index == 3 {
		data := types.NewGXByteBuffer()
		err = data.SetUint8(uint8(enums.DataTypeArray))
		if err != nil {
			return nil, err
		}
		// Add count
		cnt := 0
		if g.InitialisationStrings != nil {
			cnt = len(g.InitialisationStrings)
		}
		types.SetObjectCount(cnt, data)
		if cnt != 0 {
			for _, it := range g.InitialisationStrings {
				err = data.SetUint8(uint8(enums.DataTypeStructure))
				if err != nil {
					return nil, err
				}
				err = data.SetUint8(uint8(3))
				if err != nil {
					return nil, err
				}
				err = internal.SetData(settings, data, enums.DataTypeOctetString, []byte(it.Request))
				if err != nil {
					return nil, err
				}
				err = internal.SetData(settings, data, enums.DataTypeOctetString, []byte(it.Response))
				if err != nil {
					return nil, err
				}
				err = internal.SetData(settings, data, enums.DataTypeUint16, it.Delay)
				if err != nil {
					return nil, err
				}
			}
		}
		return data.Array(), nil
	}
	if e.Index == 4 {
		data := types.NewGXByteBuffer()
		err = data.SetUint8(uint8(enums.DataTypeArray))
		if err != nil {
			return nil, err
		}
		// Add count
		cnt := 0
		if g.ModemProfile != nil {
			cnt = len(g.ModemProfile)
		}
		types.SetObjectCount(cnt, data)
		if cnt != 0 {
			for _, it := range g.ModemProfile {
				err = internal.SetData(settings, data, enums.DataTypeOctetString, []byte(it))
				if err != nil {
					return nil, err
				}
			}
		}
		return data.Array(), nil
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
func (g *GXDLMSModemConfiguration) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	if e.Index == 1 {
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		err = g.SetLogicalName(ln)
	} else if e.Index == 2 {
		g.CommunicationSpeed = enums.BaudRate(e.Value.(types.GXEnum).Value)
	} else if e.Index == 3 {
		g.InitialisationStrings = nil
		if e.Value != nil {
			items := []GXDLMSModemInitialisation{}
			for _, tmp := range e.Value.(types.GXArray) {
				it := tmp.(types.GXStructure)
				item := GXDLMSModemInitialisation{}
				item.Request = string(it[0].([]byte))
				item.Response = string(it[1].([]byte))
				if len(it) > 2 {
					item.Delay = uint16(it[2].(int))
				}
				items = append(items, item)
			}
			g.InitialisationStrings = items
		}
	} else if e.Index == 4 {
		g.ModemProfile = nil
		if e.Value != nil {
			items := []string{}
			for _, it := range e.Value.(types.GXArray) {
				ret, err := internal.ChangeTypeFromByteArray(settings, it.([]byte), enums.DataTypeString)
				if err != nil {
					return err
				}
				items = append(items, strings.TrimSpace(ret.(string)))
			}
			g.ModemProfile = items
		}
	} else {
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
func (g *GXDLMSModemConfiguration) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, nil
}

// Load returns the load object content from XML.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSModemConfiguration) Load(reader *GXXmlReader) error {
	ret, err := reader.ReadElementContentAsInt("CommunicationSpeed", 0)
	if err != nil {
		return err
	}
	g.CommunicationSpeed = enums.BaudRate(ret)
	b, err := reader.IsStartElementNamed("InitialisationStrings", true)
	if err != nil {
		return err
	}
	if b {
		b, err = reader.IsStartElementNamed("Initialisation", true)
		if err != nil {
			return err
		}
		for b {
			it := GXDLMSModemInitialisation{}
			it.Request, err = reader.ReadElementContentAsString("Request", "")
			if err != nil {
				return err
			}
			it.Response, err = reader.ReadElementContentAsString("Response", "")
			if err != nil {
				return err
			}
			ret, err = reader.ReadElementContentAsInt("Delay", 0)
			if err != nil {
				return err
			}
			it.Delay = uint16(ret)
			// Add the initialisation string to the list
			g.InitialisationStrings = append(g.InitialisationStrings, it)
		}
		reader.ReadEndElement("InitialisationStrings")
	}
	ret2, err := reader.ReadElementContentAsString("ModemProfile", "")
	if err != nil {
		return err
	}
	g.ModemProfile = strings.Split(ret2, ";")
	return err
}

// Save returns the save object content to XML.
//
// Parameters:
//
//	writer: XML writer.
func (g *GXDLMSModemConfiguration) Save(writer *GXXmlWriter) error {
	err := writer.WriteElementString("CommunicationSpeed", int(g.CommunicationSpeed))
	if err != nil {
		return err
	}
	writer.WriteStartElement("InitialisationStrings")
	if g.InitialisationStrings != nil {
		for _, it := range g.InitialisationStrings {
			writer.WriteStartElement("Initialisation")
			err = writer.WriteElementString("Request", it.Request)
			if err != nil {
				return err
			}
			err = writer.WriteElementString("Response", it.Response)
			if err != nil {
				return err
			}
			err = writer.WriteElementString("Delay", it.Delay)
			if err != nil {
				return err
			}
			writer.WriteEndElement()
		}
	}
	writer.WriteEndElement()
	str := ""
	if g.ModemProfile != nil {
		str = strings.Join(g.ModemProfile, ";")
	}
	err = writer.WriteElementString("ModemProfile", str)
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
func (g *GXDLMSModemConfiguration) PostLoad(reader *GXXmlReader) error {
	return nil
}

// GetValues returns the an array containing the COSEM object's attribute values.
func (g *GXDLMSModemConfiguration) GetValues() []any {
	return []any{g.LogicalName(), g.CommunicationSpeed, g.InitialisationStrings, g.ModemProfile}
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
func (g *GXDLMSModemConfiguration) GetDataType(index int) (enums.DataType, error) {
	if index == 1 {
		return enums.DataTypeOctetString, nil
	}
	if index == 2 {
		return enums.DataTypeEnum, nil
	}
	if index == 3 {
		return enums.DataTypeArray, nil
	}
	if index == 4 {
		return enums.DataTypeArray, nil
	}
	return 0, dlmserrors.ErrInvalidAttributeIndex
}

// NewGXDLMSModemConfiguration creates a new modem configuration object instance.
//
// The function validates `ln` before creating the object.
//`ln` is the Logical Name and `sn` is the Short Name of the object.
func NewGXDLMSModemConfiguration(ln string, sn int16) (*GXDLMSModemConfiguration, error) {
	err := ValidateLogicalName(ln)
	if err != nil {
		return nil, err
	}
	return &GXDLMSModemConfiguration{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeModemConfiguration,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
