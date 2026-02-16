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
// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSGprsSetup
type GXDLMSGprsSetup struct {
	GXDLMSObject
	APN string

	PINCode uint16

	DefaultQualityOfService GXDLMSQosElement

	RequestedQualityOfService GXDLMSQosElement
}

// base returns the base GXDLMSObject of the object.
func (g *GXDLMSGprsSetup) Base() *GXDLMSObject {
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
func (g *GXDLMSGprsSetup) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	// LN is static and read only once.
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	// APN
	if all || !g.CanRead(2) {
		attributes = append(attributes, 2)
	}
	// PINCode
	if all || !g.CanRead(3) {
		attributes = append(attributes, 3)
	}
	// DefaultQualityOfService + RequestedQualityOfService
	if all || !g.CanRead(4) {
		attributes = append(attributes, 4)
	}
	return attributes
}

// GetNames returns the names of attribute indexes.
func (g *GXDLMSGprsSetup) GetNames() []string {
	return []string{"Logical Name", "APN", "PIN Code", "Default Quality Of Service and Requested Quality Of Service"}
}

// GetMethodNames returns the names of method indexes.
func (g *GXDLMSGprsSetup) GetMethodNames() []string {
	return []string{}
}

// GetAttributeCount returns the amount of attributes.
//
// Returns:
//
//	Count of attributes.
func (g *GXDLMSGprsSetup) GetAttributeCount() int {
	return 4
}

// GetMethodCount returns the amount of methods.
func (g *GXDLMSGprsSetup) GetMethodCount() int {
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
func (g *GXDLMSGprsSetup) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	var ret any
	var err error
	if e.Index == 1 {
		ret, err = helpers.LogicalNameToBytes(g.LogicalName())
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
	} else if e.Index == 2 {
		if g.APN != "" {
			ret = []byte(g.APN)
		}
	} else if e.Index == 3 {
		ret = g.PINCode
	} else if e.Index == 4 {
		data := types.NewGXByteBuffer()
		err = data.SetUint8(uint8(enums.DataTypeStructure))
		if err != nil {
			return nil, err
		}
		err = data.SetUint8(uint8(2))
		if err != nil {
			return nil, err
		}
		err = data.SetUint8(uint8(enums.DataTypeStructure))
		if err != nil {
			return nil, err
		}
		err = data.SetUint8(uint8(5))
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, data, enums.DataTypeUint8, g.DefaultQualityOfService.Precedence)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, data, enums.DataTypeUint8, g.DefaultQualityOfService.Delay)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, data, enums.DataTypeUint8, g.DefaultQualityOfService.Reliability)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, data, enums.DataTypeUint8, g.DefaultQualityOfService.PeakThroughput)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, data, enums.DataTypeUint8, g.DefaultQualityOfService.MeanThroughput)
		if err != nil {
			return nil, err
		}
		err = data.SetUint8(uint8(enums.DataTypeStructure))
		if err != nil {
			return nil, err
		}
		err = data.SetUint8(uint8(5))
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, data, enums.DataTypeUint8, g.RequestedQualityOfService.Precedence)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, data, enums.DataTypeUint8, g.RequestedQualityOfService.Delay)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, data, enums.DataTypeUint8, g.RequestedQualityOfService.Reliability)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, data, enums.DataTypeUint8, g.RequestedQualityOfService.PeakThroughput)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, data, enums.DataTypeUint8, g.RequestedQualityOfService.MeanThroughput)
		if err != nil {
			return nil, err
		}
		ret = data.Array()
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
func (g *GXDLMSGprsSetup) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	var err error
	switch e.Index {
	case 1:
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		err = g.SetLogicalName(ln)
	case 2:
		if v, ok := e.Value.(string); ok {
			g.APN = v
		} else {
			g.APN = string(e.Value.([]byte))
		}
	case 3:
		g.PINCode = e.Value.(uint16)
	case 4:
		g.DefaultQualityOfService.Precedence = 0
		g.DefaultQualityOfService.Delay = 0
		g.DefaultQualityOfService.Reliability = 0
		g.DefaultQualityOfService.PeakThroughput = 0
		g.DefaultQualityOfService.MeanThroughput = 0
		g.RequestedQualityOfService.Precedence = 0
		g.RequestedQualityOfService.Delay = 0
		g.RequestedQualityOfService.Reliability = 0
		g.RequestedQualityOfService.PeakThroughput = 0
		g.RequestedQualityOfService.MeanThroughput = 0
		if e.Value != nil {
			tmp := e.Value.(types.GXStructure)
			t := tmp[0].(types.GXStructure)
			g.DefaultQualityOfService.Precedence = t[0].(uint8)
			g.DefaultQualityOfService.Delay = t[1].(uint8)
			g.DefaultQualityOfService.Reliability = t[2].(uint8)
			g.DefaultQualityOfService.PeakThroughput = t[3].(uint8)
			g.DefaultQualityOfService.MeanThroughput = t[4].(uint8)
			t = tmp[1].(types.GXStructure)
			g.RequestedQualityOfService.Precedence = t[0].(uint8)
			g.RequestedQualityOfService.Delay = t[1].(uint8)
			g.RequestedQualityOfService.Reliability = t[2].(uint8)
			g.RequestedQualityOfService.PeakThroughput = t[3].(uint8)
			g.RequestedQualityOfService.MeanThroughput = t[4].(uint8)
		}
	default:
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
func (g *GXDLMSGprsSetup) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, nil
}

// Load returns the load object content from XML.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSGprsSetup) Load(reader *GXXmlReader) error {
	var err error
	g.APN, err = reader.ReadElementContentAsString("APN", "")
	if err != nil {
		return err
	}
	ret, err := reader.ReadElementContentAsInt("PINCode", 0)
	if err != nil {
		return err
	}
	g.PINCode = uint16(ret)
	if reader.isStartElementNamed2("DefaultQualityOfService", true) {
		ret, err := reader.ReadElementContentAsInt("Precedence", 0)
		if err != nil {
			return err
		}
		g.DefaultQualityOfService.Precedence = uint8(ret)
		ret, err = reader.ReadElementContentAsInt("Delay", 0)
		if err != nil {
			return err
		}
		g.DefaultQualityOfService.Delay = uint8(ret)
		ret, err = reader.ReadElementContentAsInt("Reliability", 0)
		if err != nil {
			return err
		}
		g.DefaultQualityOfService.Reliability = uint8(ret)
		ret, err = reader.ReadElementContentAsInt("PeakThroughput", 0)
		if err != nil {
			return err
		}
		g.DefaultQualityOfService.PeakThroughput = uint8(ret)
		ret, err = reader.ReadElementContentAsInt("MeanThroughput", 0)
		if err != nil {
			return err
		}
		g.DefaultQualityOfService.MeanThroughput = uint8(ret)
		reader.ReadEndElement("DefaultQualityOfService")
	}
	if reader.isStartElementNamed2("RequestedQualityOfService", true) {
		ret, err = reader.ReadElementContentAsInt("Precedence", 0)
		if err != nil {
			return err
		}
		g.RequestedQualityOfService.Precedence = uint8(ret)
		ret, err = reader.ReadElementContentAsInt("Delay", 0)
		if err != nil {
			return err
		}
		g.RequestedQualityOfService.Delay = uint8(ret)
		ret, err = reader.ReadElementContentAsInt("Reliability", 0)
		if err != nil {
			return err
		}
		g.RequestedQualityOfService.Reliability = uint8(ret)
		ret, err = reader.ReadElementContentAsInt("PeakThroughput", 0)
		if err != nil {
			return err
		}
		g.RequestedQualityOfService.PeakThroughput = uint8(ret)
		ret, err = reader.ReadElementContentAsInt("MeanThroughput", 0)
		if err != nil {
			return err
		}
		g.RequestedQualityOfService.MeanThroughput = uint8(ret)
		reader.ReadEndElement("RequestedQualityOfService")
	}
	return err
}

// Save returns the save object content to XML.
//
// Parameters:
//
//	writer: XML writer.
func (g *GXDLMSGprsSetup) Save(writer *GXXmlWriter) error {
	err := writer.WriteElementString("APN", g.APN)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("PINCode", g.PINCode)
	if err != nil {
		return err
	}
	writer.WriteStartElement("DefaultQualityOfService")
	err = writer.WriteElementString("Precedence", g.DefaultQualityOfService.Precedence)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("Delay", g.DefaultQualityOfService.Delay)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("Reliability", g.DefaultQualityOfService.Reliability)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("PeakThroughput", g.DefaultQualityOfService.PeakThroughput)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("MeanThroughput", g.DefaultQualityOfService.MeanThroughput)
	if err != nil {
		return err
	}
	writer.WriteEndElement()
	writer.WriteStartElement("RequestedQualityOfService")
	err = writer.WriteElementString("Precedence", g.RequestedQualityOfService.Precedence)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("Delay", g.RequestedQualityOfService.Delay)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("Reliability", g.RequestedQualityOfService.Reliability)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("PeakThroughput", g.RequestedQualityOfService.PeakThroughput)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("MeanThroughput", g.RequestedQualityOfService.MeanThroughput)
	if err != nil {
		return err
	}
	writer.WriteEndElement()
	return err
}

// PostLoad returns the handle actions after Load.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSGprsSetup) PostLoad(reader *GXXmlReader) error {
	return nil
}

// GetValues returns the an array containing the COSEM object's attribute values.
func (g *GXDLMSGprsSetup) GetValues() []any {
	return []any{g.LogicalName, g.APN, g.PINCode, []any{g.DefaultQualityOfService, g.RequestedQualityOfService}}
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
func (g *GXDLMSGprsSetup) GetUIDataType(index int) enums.DataType {
	if index == 2 {
		return enums.DataTypeString
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
func (g *GXDLMSGprsSetup) GetDataType(index int) (enums.DataType, error) {
	if index == 1 {
		return enums.DataTypeOctetString, nil
	}
	if index == 2 {
		return enums.DataTypeOctetString, nil
	}
	if index == 3 {
		return enums.DataTypeUint16, nil
	}
	if index == 4 {
		return enums.DataTypeArray, nil
	}
	return enums.DataTypeNone, errors.New("GetDataType failed. Invalid attribute index.")
}

// Constructor.
// ln: Logical Name of the object.
// sn: Short Name of the object.
func NewGXDLMSGprsSetup(ln string, sn int16) (*GXDLMSGprsSetup, error) {
	err := ValidateLogicalName(ln)
	if err != nil {
		return nil, err
	}
	return &GXDLMSGprsSetup{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeGprsSetup,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
