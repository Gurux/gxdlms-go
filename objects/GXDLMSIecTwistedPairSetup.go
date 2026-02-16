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
// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSIecTwistedPairSetup
type GXDLMSIecTwistedPairSetup struct {
	GXDLMSObject
	// Working mode.
	Mode enums.IecTwistedPairSetupMode

	// Communication speed.
	Speed enums.BaudRate

	// List of Primary Station Addresses.
	PrimaryAddresses []byte

	// List of the TAB(i) for which the real equipment has been programmed
	// in the case of forgotten station call.
	Tabis []int8
}

// base returns the base GXDLMSObject of the object.
func (g *GXDLMSIecTwistedPairSetup) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

// Invoke returns the invokes method.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Invoke parameters.
func (g *GXDLMSIecTwistedPairSetup) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
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
func (g *GXDLMSIecTwistedPairSetup) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	// LN is static and read only once.
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	// Mode
	if all || g.CanRead(2) {
		attributes = append(attributes, 2)
	}
	// Speed
	if all || g.CanRead(3) {
		attributes = append(attributes, 3)
	}
	// PrimaryAddresses
	if all || g.CanRead(4) {
		attributes = append(attributes, 4)
	}
	// Tabis
	if all || g.CanRead(5) {
		attributes = append(attributes, 5)
	}
	return attributes
}

// GetNames returns the names of attribute indexes.
func (g *GXDLMSIecTwistedPairSetup) GetNames() []string {
	return []string{"Logical Name", "Mode", "Speed", "PrimaryAddresses", "Tabis"}
}

// GetMethodNames returns the names of method indexes.
func (g *GXDLMSIecTwistedPairSetup) GetMethodNames() []string {
	return []string{}
}

// GetAttributeCount returns the amount of attributes.
//
// Returns:
//
//	Count of attributes.
func (g *GXDLMSIecTwistedPairSetup) GetAttributeCount() int {
	return 5
}

// GetMethodCount returns the amount of methods.
func (g *GXDLMSIecTwistedPairSetup) GetMethodCount() int {
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
func (g *GXDLMSIecTwistedPairSetup) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	var ret any
	var err error
	switch e.Index {
	case 1:
		v, err := helpers.LogicalNameToBytes(g.LogicalName())
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return v, err
	case 2:
		ret = uint8(g.Mode)
	case 3:
		ret = uint8(g.Speed)
	case 4:
		data := types.NewGXByteBuffer()
		err = data.SetUint8(uint8(enums.DataTypeArray))
		if err != nil {
			return nil, err
		}
		if g.PrimaryAddresses == nil {
			err = data.SetUint8(0)
			if err != nil {
				return nil, err
			}
		} else {
			err = data.SetUint8(uint8(len(g.PrimaryAddresses)))
			if err != nil {
				return nil, err
			}
			for _, it := range g.PrimaryAddresses {
				err = data.SetUint8(uint8(enums.DataTypeUint8))
				if err != nil {
					return nil, err
				}
				err = data.SetUint8(it)
				if err != nil {
					return nil, err
				}
			}
		}
		ret = data.Array()
	case 5:
		data := types.NewGXByteBuffer()
		err = data.SetUint8(uint8(enums.DataTypeArray))
		if err != nil {
			return nil, err
		}
		if g.Tabis == nil {
			err = data.SetUint8(0)
			if err != nil {
				return nil, err
			}
		} else {
			err = data.SetUint8(uint8(len(g.Tabis)))
			if err != nil {
				return nil, err
			}
			for _, it := range g.Tabis {
				err = data.SetUint8(uint8(enums.DataTypeInt8))
				if err != nil {
					return nil, err
				}
				err = data.SetUint8(uint8(it))
				if err != nil {
					return nil, err
				}
			}
		}
		ret = data.Array()
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
		ret = nil
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
func (g *GXDLMSIecTwistedPairSetup) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	switch e.Index {
	case 1:
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		err = g.SetLogicalName(ln)
	case 2:
		g.Mode = enums.IecTwistedPairSetupMode(e.Value.(types.GXEnum).Value)
	case 3:
		g.Speed = enums.BaudRate(e.Value.(types.GXEnum).Value)
	case 4:
		list := types.NewGXByteBuffer()
		if e.Value != nil {
			for _, it := range e.Value.(types.GXArray) {
				list.Add(it.(byte))
			}
		}
		g.PrimaryAddresses = list.Array()
	case 5:
		list := []int8{}
		if e.Value != nil {
			for _, it := range e.Value.(types.GXArray) {
				list = append(list, it.(int8))
			}
		}
		g.Tabis = list
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
func (g *GXDLMSIecTwistedPairSetup) Load(reader *GXXmlReader) error {
	var err error
	ret, err := reader.ReadElementContentAsInt("Mode", 0)
	if err != nil {
		return err
	}
	g.Mode = enums.IecTwistedPairSetupMode(ret)
	ret, err = reader.ReadElementContentAsInt("Speed", 0)
	if err != nil {
		return err
	}
	g.Speed = enums.BaudRate(ret)
	ret2, err := reader.ReadElementContentAsString("PrimaryAddresses", "")
	if err != nil {
		return err
	}
	g.PrimaryAddresses = types.HexToBytes(ret2)
	ret2, err = reader.ReadElementContentAsString("Tabis", "")
	if err != nil {
		return err
	}
	tmp := types.HexToBytes(ret2)
	g.Tabis = make([]int8, len(tmp))
	if len(tmp) != 0 {
		for i, v := range tmp {
			g.Tabis[i] = int8(v)
		}
	}
	return err
}

// Save returns the save object content to XML.
//
// Parameters:
//
//	writer: XML writer.
func (g *GXDLMSIecTwistedPairSetup) Save(writer *GXXmlWriter) error {
	err := writer.WriteElementString("Mode", int(g.Mode))
	if err != nil {
		return err
	}
	err = writer.WriteElementString("Speed", int(g.Speed))
	if err != nil {
		return err
	}
	err = writer.WriteElementString("PrimaryAddresses", types.ToHex(g.PrimaryAddresses, true))
	if err != nil {
		return err
	}
	if g.Tabis != nil {
		tmp := make([]byte, len(g.Tabis))
		for i, v := range g.Tabis {
			tmp[i] = byte(v)
		}
		err = writer.WriteElementString("Tabis", types.ToHex(tmp, true))
	}
	return err
}

// PostLoad returns the handle actions after Load.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSIecTwistedPairSetup) PostLoad(reader *GXXmlReader) error {
	return nil
}

// GetValues returns the an array containing the COSEM object's attribute values.
func (g *GXDLMSIecTwistedPairSetup) GetValues() []any {
	return []any{g.LogicalName, g.Mode, g.Speed, g.PrimaryAddresses, g.Tabis}
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
func (g *GXDLMSIecTwistedPairSetup) GetDataType(index int) (enums.DataType, error) {
	var ret enums.DataType
	switch index {
	case 1:
		ret = enums.DataTypeOctetString
	case 2:
	case 3:
		ret = enums.DataTypeEnum
	case 4:
	case 5:
		ret = enums.DataTypeArray
	default:
		return 0, errors.New("GetDataType failed. Invalid attribute index.")
	}
	return ret, nil
}

// Constructor.
// ln: Logical Name of the object.
// sn: Short Name of the object.
func NewGXDLMSIecTwistedPairSetup(ln string, sn int16) (*GXDLMSIecTwistedPairSetup, error) {
	err := ValidateLogicalName(ln)
	if err != nil {
		return nil, err
	}
	return &GXDLMSIecTwistedPairSetup{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeIecTwistedPairSetup,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
