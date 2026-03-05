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
// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSTokenGateway
type GXDLMSTokenGateway struct {
	GXDLMSObject
	// Token.
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSTokenGateway
	Token []byte

	// Time
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSTokenGateway
	Time types.GXDateTime

	// Descriptions.
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSTokenGateway
	Descriptions []string

	// Token Delivery method.
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSTokenGateway
	DeliveryMethod enums.TokenDelivery

	// Token status code.
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSTokenGateway
	StatusCode enums.TokenStatusCode

	// Token data value.
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSTokenGateway
	DataValue string
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSTokenGateway) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

// Invoke returns the invokes method.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Invoke parameters.
func (g *GXDLMSTokenGateway) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
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
func (g *GXDLMSTokenGateway) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	// LN is static and read only once.
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	// Token
	if all || g.CanRead(2) {
		attributes = append(attributes, 2)
	}
	// Time
	if all || g.CanRead(3) {
		attributes = append(attributes, 3)
	}
	// Description
	if all || g.CanRead(4) {
		attributes = append(attributes, 4)
	}
	// DeliveryMethod
	if all || g.CanRead(5) {
		attributes = append(attributes, 5)
	}
	// Status
	if all || g.CanRead(6) {
		attributes = append(attributes, 6)
	}
	return attributes
}

// GetNames returns the names of attribute indexes.
func (g *GXDLMSTokenGateway) GetNames() []string {
	return []string{"Logical Name", "Token", "Time", "Description", "DeliveryMethod", "Status"}
}

// GetMethodNames returns the names of method indexes.
func (g *GXDLMSTokenGateway) GetMethodNames() []string {
	return []string{"Enter"}
}

// GetAttributeCount returns the amount of attributes.
//
// Returns:
//
//	Count of attributes.
func (g *GXDLMSTokenGateway) GetAttributeCount() int {
	return 6
}

// GetMethodCount returns the amount of methods.
func (g *GXDLMSTokenGateway) GetMethodCount() int {
	return 1
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
func (g *GXDLMSTokenGateway) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	var ret any
	var err error
	switch e.Index {
	case 1:
		return helpers.LogicalNameToBytes(g.LogicalName())
	case 2:
		ret = g.Token
	case 3:
		ret = g.Time
	case 4:
		bb := types.GXByteBuffer{}
		err = bb.SetUint8(enums.DataTypeArray)
		if err != nil {
			return nil, err
		}
		err = bb.SetUint8(byte(len(g.Descriptions)))
		if err != nil {
			return nil, err
		}
		for _, it := range g.Descriptions {
			err = bb.SetUint8(byte(enums.DataTypeOctetString))
			if err != nil {
				return nil, err
			}
			err = bb.SetUint8(byte(len(it)))
			if err != nil {
				return nil, err
			}
			err = bb.Set([]byte(it))
			if err != nil {
				return nil, err
			}
		}
		ret = bb.Array()
	case 5:
		ret = uint8(g.DeliveryMethod)
	case 6:
		bb := types.GXByteBuffer{}
		err = bb.SetUint8(byte(enums.DataTypeStructure))
		if err != nil {
			return nil, err
		}
		err = bb.SetUint8(2)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, &bb, enums.DataTypeEnum, uint8(g.StatusCode))
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, &bb, enums.DataTypeBitString, g.DataValue)
		if err != nil {
			return nil, err
		}
		ret = bb.Array()
	default:
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
func (g *GXDLMSTokenGateway) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	var err error
	switch e.Index {
	case 1:
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return g.SetLogicalName(ln)
	case 2:
		g.Token = e.Value.([]byte)
	case 3:
		if v, ok := e.Value.([]byte); ok {
			ret, err := internal.ChangeTypeFromByteArray(settings, v, enums.DataTypeDateTime)
			if err != nil {
				return err
			}
			g.Time = ret.(types.GXDateTime)
		} else if v, ok := e.Value.(string); ok {
			ret, err := types.NewGXDateTimeFromString(v, nil)
			if err != nil {
				return err
			}
			g.Time = *ret
		} else {
			g.Time = e.Value.(types.GXDateTime)
		}
	case 4:
		g.Descriptions = g.Descriptions[:0]
		if e.Value != nil {
			for _, it := range e.Value.(types.GXArray) {
				g.Descriptions = append(g.Descriptions, string(it.([]byte)))
			}
		}
	case 5:
		g.DeliveryMethod = enums.TokenDelivery(e.Value.(types.GXEnum).Value)
	case 6:
		if e.Value != nil {
			arr := e.Value.(types.GXStructure)
			g.StatusCode = enums.TokenStatusCode(arr[0].(types.GXEnum).Value)
			bs := arr[1].(types.GXBitString)
			g.DataValue = bs.String()
		} else {
			g.StatusCode = enums.TokenStatusCodeFormatOk
			g.DataValue = ""
		}
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return err
}

// Load returns the load object content from XML.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSTokenGateway) Load(reader *GXXmlReader) error {
	tmp, err := reader.ReadElementContentAsString("Token", "")
	if err != nil {
		return err
	}
	g.Token = types.HexToBytes(tmp)
	g.Time, err = reader.ReadElementContentAsGXDateTime("Time")
	if err != nil {
		return err
	}
	g.Descriptions = g.Descriptions[:0]
	if ret, err := reader.IsStartElementNamed("Descriptions", true); ret && err == nil {
		for {
			ret, err = reader.IsStartElementNamed("Item", true)
			if err != nil {
				return err
			}
			if !ret {
				break
			}
			ret, err := reader.ReadElementContentAsString("Name", "")
			if err != nil {
				return err
			}
			g.Descriptions = append(g.Descriptions, ret)
		}
		reader.ReadEndElement("Descriptions")
	}
	ret, err := reader.ReadElementContentAsInt("DeliveryMethod", 0)
	if err != nil {
		return err
	}
	g.DeliveryMethod = enums.TokenDelivery(ret)
	ret, err = reader.ReadElementContentAsInt("Status", 0)
	if err != nil {
		return err
	}
	g.StatusCode = enums.TokenStatusCode(ret)
	g.DataValue, err = reader.ReadElementContentAsString("Data", "")
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
func (g *GXDLMSTokenGateway) Save(writer *GXXmlWriter) error {
	err := writer.WriteElementString("Token", types.ToHex(g.Token, false))
	if err != nil {
		return err
	}
	err = writer.WriteElementString("Time", g.Time)
	if err != nil {
		return err
	}
	writer.WriteStartElement("Descriptions")
	if g.Descriptions != nil {
		for _, it := range g.Descriptions {
			writer.WriteStartElement("Item")
			err = writer.WriteElementString("Name", it)
			if err != nil {
				return err
			}
			writer.WriteEndElement()
		}
	}
	writer.WriteEndElement()
	err = writer.WriteElementString("DeliveryMethod", int(g.DeliveryMethod))
	if err != nil {
		return err
	}
	err = writer.WriteElementString("Status", int(g.StatusCode))
	if err != nil {
		return err
	}
	err = writer.WriteElementString("Data", g.DataValue)
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
func (g *GXDLMSTokenGateway) PostLoad(reader *GXXmlReader) error {
	return nil
}

// GetValues returns an array containing the object's current attribute values.
func (g *GXDLMSTokenGateway) GetValues() []any {
	return []any{g.LogicalName(), g.Token, g.Time, g.Descriptions, g.DeliveryMethod, []any{g.StatusCode, g.DataValue}}
}

// Enter returns the transfer a token to the server.
//
// Parameters:
//
//	client: DLMS client.
//	token: The token to send.
//
// Returns:
//
//	Action bytes.
func (g *GXDLMSTokenGateway) Enter(client IGXDLMSClient, token []byte) ([][]uint8, error) {
	return client.Method(g, 1, token, enums.DataTypeOctetString)
}

// GetUIDataType returns the UI data type of selected index.
//
// Parameters:
//
//	index: Attribute index of the object.
//
// Returns:
//
//	UI data type of the object.
func (g *GXDLMSTokenGateway) GetUIDataType(index int) enums.DataType {
	if index == 3 {
		return enums.DataTypeDateTime
	}
	return g.Base().GetUIDataType(index)
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
func (g *GXDLMSTokenGateway) GetDataType(index int) (enums.DataType, error) {
	var ret enums.DataType
	var err error
	switch index {
	case 1:
		ret = enums.DataTypeOctetString
	case 2:
		ret = enums.DataTypeOctetString
	case 3:
		ret = enums.DataTypeOctetString
	case 4:
		ret = enums.DataTypeArray
	case 5:
		ret = enums.DataTypeEnum
	case 6:
		ret = enums.DataTypeStructure
	default:
		err = errors.New("GetDataType failed. Invalid attribute index.")
	}
	return ret, err
}

// NewGXDLMSGXDLMSTokenGateway creates a new Token Gateway object instance.
//
// The function validates `ln` before creating the object.
// `ln` is the Logical Name and `sn` is the Short Name of the object.
func NewGXDLMSTokenGateway(ln string, sn int16) (*GXDLMSTokenGateway, error) {
	err := ValidateLogicalName(ln)
	if err != nil {
		return nil, err
	}
	return &GXDLMSTokenGateway{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeTokenGateway,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
