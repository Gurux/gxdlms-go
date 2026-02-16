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
	"reflect"
	"strings"

	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// Online help:
// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSPppSetup
type GXDLMSPppSetup struct {
	GXDLMSObject
	PHYReference string

	LCPOptions []GXDLMSPppSetupLcpOption

	IPCPOptions []GXDLMSPppSetupIPCPOption

	// PPP authentication procedure type.
	Authentication enums.PppAuthenticationType

	// PPP authentication procedure user name.
	UserName []byte

	// PPP authentication procedure password.
	Password []byte
}

// base returns the base GXDLMSObject of the object.
func (g *GXDLMSPppSetup) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

// Invoke returns the invokes method.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Invoke parameters.
func (g *GXDLMSPppSetup) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
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
func (g *GXDLMSPppSetup) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	// LN is static and read only once.
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	// PHYReference
	if all || !g.IsRead(2) {
		attributes = append(attributes, 2)
	}
	// LCPOptions
	if all || !g.IsRead(3) {
		attributes = append(attributes, 3)
	}
	// IPCPOptions
	if all || !g.IsRead(4) {
		attributes = append(attributes, 4)
	}
	// PPPAuthentication
	if all || !g.IsRead(5) {
		attributes = append(attributes, 5)
	}
	return attributes
}

// GetNames returns the names of attribute indexes.
func (g *GXDLMSPppSetup) GetNames() []string {
	return []string{"Logical Name", "PHY Reference", "LCP Options", "IPCP Options", "PPP Authentication"}
}

// GetMethodNames returns the names of method indexes.
func (g *GXDLMSPppSetup) GetMethodNames() []string {
	return []string{}
}

// GetAttributeCount returns the amount of attributes.
//
// Returns:
//
//	Count of attributes.
func (g *GXDLMSPppSetup) GetAttributeCount() int {
	return 5
}

// GetMethodCount returns the amount of methods.
func (g *GXDLMSPppSetup) GetMethodCount() int {
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
func (g *GXDLMSPppSetup) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	var err error
	if e.Index == 1 {
		v, err := helpers.LogicalNameToBytes(g.LogicalName())
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return v, err
	}
	if e.Index == 2 {
		return helpers.LogicalNameToBytes(g.PHYReference)
	}
	if e.Index == 3 {
		data := types.NewGXByteBuffer()
		err = data.SetUint8(uint8(enums.DataTypeArray))
		if err != nil {
			return nil, err
		}
		if g.LCPOptions == nil {
			err = data.SetUint8(0)
			if err != nil {
				return nil, err
			}
		} else {
			err = data.SetUint8(uint8(len(g.LCPOptions)))
			if err != nil {
				return nil, err
			}
			for _, it := range g.LCPOptions {
				err = data.SetUint8(uint8(enums.DataTypeStructure))
				if err != nil {
					return nil, err
				}
				err = data.SetUint8(uint8(3))
				if err != nil {
					return nil, err
				}
				err = internal.SetData(settings, data, enums.DataTypeUint8, it.Type)
				if err != nil {
					return nil, err
				}
				err = internal.SetData(settings, data, enums.DataTypeUint8, it.Length)
				if err != nil {
					return nil, err
				}
				dt, err := internal.GetDLMSDataType(reflect.TypeOf(it.Data))
				if err != nil {
					return nil, err
				}
				err = internal.SetData(settings, data, dt, it.Data)
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
		if g.IPCPOptions == nil {
			err = data.SetUint8(0)
			if err != nil {
				return nil, err
			}
		} else {
			err = data.SetUint8(uint8(len(g.IPCPOptions)))
			if err != nil {
				return nil, err
			}
			for _, it := range g.IPCPOptions {
				err = data.SetUint8(uint8(enums.DataTypeStructure))
				if err != nil {
					return nil, err
				}
				err = data.SetUint8(uint8(3))
				if err != nil {
					return nil, err
				}
				err = internal.SetData(settings, data, enums.DataTypeUint8, it.Type)
				if err != nil {
					return nil, err
				}
				err = internal.SetData(settings, data, enums.DataTypeUint8, it.Length)
				if err != nil {
					return nil, err
				}
				dt, err := internal.GetDLMSDataType(reflect.TypeOf(it.Data))
				if err != nil {
					return nil, err
				}
				err = internal.SetData(settings, data, dt, it.Data)
				if err != nil {
					return nil, err
				}
			}
		}
		return data.Array(), nil
	} else if e.Index == 5 {
		if len(g.UserName) == 0 {
			return nil, nil
		}
		data := types.NewGXByteBuffer()
		err = data.SetUint8(uint8(enums.DataTypeStructure))
		if err != nil {
			return nil, err
		}
		err = data.SetUint8(2)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, data, enums.DataTypeOctetString, g.UserName)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, data, enums.DataTypeOctetString, g.Password)
		if err != nil {
			return nil, err
		}
		return data.Array(), nil
	}
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, err
}

// SetValue returns the set value of given attribute.
// When raw parameter us not used example register multiplies value by scalar.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Set parameters.
func (g *GXDLMSPppSetup) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	var err error
	if e.Index == 1 {
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		err = g.SetLogicalName(ln)
	} else if e.Index == 2 {
		g.PHYReference, err = helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
	} else if e.Index == 3 {
		items := []GXDLMSPppSetupLcpOption{}
		for _, it := range e.Value.(types.GXArray) {
			item := it.(types.GXStructure)
			it := GXDLMSPppSetupLcpOption{}
			it.Type = enums.PppSetupLcpOptionType(item[0].(types.GXEnum).Value)
			it.Length = item[1].(uint8)
			it.Data = item[2]
			items = append(items, it)
		}
		g.LCPOptions = items
	} else if e.Index == 4 {
		items := []GXDLMSPppSetupIPCPOption{}
		for _, it := range e.Value.(types.GXArray) {
			item := it.(types.GXStructure)
			it := GXDLMSPppSetupIPCPOption{}
			it.Type = enums.PppSetupIPCPOptionType(item[0].(types.GXEnum).Value)
			it.Length = item[1].(uint8)
			it.Data = item[2]
			items = append(items, it)
		}
		g.IPCPOptions = items
	} else if e.Index == 5 {
		if e.Value != nil {
			arr := e.Value.(types.GXStructure)
			if len(arr) == 2 {
				g.UserName = arr[0].([]byte)
				g.Password = arr[1].([]byte)
			} else {
				g.UserName = nil
				g.Password = nil
			}
		}
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
func (g *GXDLMSPppSetup) Load(reader *GXXmlReader) error {
	var err error
	var ret int
	g.PHYReference, err = reader.ReadElementContentAsString("PHYReference", "")
	if err != nil {
		return err
	}
	options := []GXDLMSPppSetupLcpOption{}
	b, err := reader.IsStartElementNamed("LCPOptions", true)
	if err != nil {
		return err
	}
	if b {
		for b, err = reader.IsStartElementNamed("Item", true); b && err == nil; {
			it := GXDLMSPppSetupLcpOption{}
			ret, err = reader.ReadElementContentAsInt("Type", 0)
			if err != nil {
				return err
			}
			it.Type = enums.PppSetupLcpOptionType(ret)
			ret, err = reader.ReadElementContentAsInt("Length", 0)
			if err != nil {
				return err
			}
			it.Length = uint8(ret)
			//TODO: ret, err = reader.ReadElementContentAsObject("Data", nil, nil, 0)
			if err != nil {
				return err
			}
			it.Data = ret
		}
		reader.ReadEndElement("LCPOptions")
	}
	g.LCPOptions = options
	list := []GXDLMSPppSetupIPCPOption{}
	b, err = reader.IsStartElementNamed("IPCPOptions", true)
	if err != nil {
		return err
	}
	if b {
		for b, err = reader.IsStartElementNamed("Item", true); b && err == nil; {
			it := GXDLMSPppSetupIPCPOption{}
			ret, err = reader.ReadElementContentAsInt("Type", 0)
			if err != nil {
				return err
			}
			it.Type = enums.PppSetupIPCPOptionType(ret)
			ret, err = reader.ReadElementContentAsInt("Length", 0)
			if err != nil {
				return err
			}
			it.Length = uint8(ret)
			//TODO: ret, err = reader.ReadElementContentAsObject("Data", nil, nil, 0)
			if err != nil {
				return err
			}
			it.Data = ret
		}
		reader.ReadEndElement("IPCPOptions")
	}
	g.IPCPOptions = list
	ret2, err := reader.ReadElementContentAsString("UserName", "")
	if err != nil {

	}
	g.UserName = types.HexToBytes(ret2)
	if err != nil {
		return err
	}
	ret2, err = reader.ReadElementContentAsString("Password", "")
	if err != nil {
		return err
	}
	g.Password = types.HexToBytes(ret2)
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
func (g *GXDLMSPppSetup) Save(writer *GXXmlWriter) error {
	err := writer.WriteElementString("PHYReference", g.PHYReference)
	if err != nil {
		return err
	}
	writer.WriteStartElement("LCPOptions")
	if g.LCPOptions != nil {
		for _, it := range g.LCPOptions {
			writer.WriteStartElement("Item")
			err = writer.WriteElementString("Type", int(it.Type))
			if err != nil {
				return err
			}
			err = writer.WriteElementString("Length", it.Length)
			if err != nil {
				return err
			}
			dt, err := internal.GetDLMSDataType(reflect.TypeOf(it.Data))
			if err != nil {
				return err
			}
			err = writer.WriteElementObject("Data", it.Data, dt, dt)
			if err != nil {
				return err
			}
			writer.WriteEndElement()
		}
	}
	writer.WriteEndElement()
	writer.WriteStartElement("IPCPOptions")
	if g.IPCPOptions != nil {
		for _, it := range g.IPCPOptions {
			writer.WriteStartElement("Item")
			err = writer.WriteElementString("Type", int(it.Type))
			if err != nil {
				return err
			}
			err = writer.WriteElementString("Length", it.Length)
			if err != nil {
				return err
			}
			dt, err := internal.GetDLMSDataType(reflect.TypeOf(it.Data))
			if err != nil {
				return err
			}
			err = writer.WriteElementObject("Data", it.Data, dt, dt)
			if err != nil {
				return err
			}
			writer.WriteEndElement()
		}
	}
	writer.WriteEndElement()
	err = writer.WriteElementString("UserName", types.ToHex(g.UserName, false))
	if err != nil {
		return err
	}
	err = writer.WriteElementString("Password", types.ToHex(g.Password, false))
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
func (g *GXDLMSPppSetup) PostLoad(reader *GXXmlReader) error {
	return nil
}

// GetValues returns the an array containing the COSEM object's attribute values.
func (g *GXDLMSPppSetup) GetValues() []any {
	sb := strings.Builder{}
	if g.UserName != nil {
		sb.WriteString(string(g.UserName))
	}
	if g.Password != nil {
		if sb.Len() != 0 {
			sb.WriteString(" ")
		}
		sb.WriteString(string(g.Password))
	}
	return []any{g.LogicalName, g.PHYReference, g.LCPOptions, g.IPCPOptions, sb.String()}
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
func (g *GXDLMSPppSetup) GetDataType(index int) (enums.DataType, error) {
	if index == 1 {
		return enums.DataTypeOctetString, nil
	}
	if index == 2 {
		return enums.DataTypeOctetString, nil
	}
	if index == 3 {
		return enums.DataTypeArray, nil
	}
	if index == 4 {
		return enums.DataTypeArray, nil
	}
	if index == 5 {
		if g.UserName == nil || len(g.UserName) == 0 {
			return enums.DataTypeNone, nil
		}
		return enums.DataTypeStructure, nil
	}
	return 0, errors.New("GetDataType failed. Invalid attribute index.")
}

// Constructor.
// ln: Logical Name of the object.
// sn: Short Name of the object.
func NewGXDLMSPppSetup(ln string, sn int16) (*GXDLMSPppSetup, error) {
	err := ValidateLogicalName(ln)
	if err != nil {
		return nil, err
	}
	return &GXDLMSPppSetup{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypePppSetup,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
