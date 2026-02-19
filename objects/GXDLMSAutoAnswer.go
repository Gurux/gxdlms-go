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
	"fmt"

	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// Online help:
// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSAutoAnswer
type GXDLMSAutoAnswer struct {
	GXDLMSObject
	Mode enums.AutoAnswerMode

	ListeningWindow []types.GXKeyValuePair[types.GXDateTime, types.GXDateTime]

	Status enums.AutoAnswerStatus

	NumberOfCalls uint8

	// Number of rings within the window defined by ListeningWindow.
	NumberOfRingsInListeningWindow uint8

	// Number of rings outside the window defined by ListeningWindow.
	NumberOfRingsOutListeningWindow uint8

	// Number of rings outside the window defined by ListeningWindow.
	AllowedCallers []types.GXKeyValuePair[string, enums.CallType]
}

// base returns the base GXDLMSObject of the object.
func (g *GXDLMSAutoAnswer) Base() *GXDLMSObject {
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
func (g *GXDLMSAutoAnswer) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	// LN is static and read only once.
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	// Mode is static and read only once.
	if all || !g.CanRead(2) {
		attributes = append(attributes, 2)
	}
	// ListeningWindow is static and read only once.
	if all || !g.CanRead(3) {
		attributes = append(attributes, 3)
	}
	// Status is not static.
	if all || g.CanRead(4) {
		attributes = append(attributes, 4)
	}
	// NumberOfCalls is static and read only once.
	if all || !g.CanRead(5) {
		attributes = append(attributes, 5)
	}
	// NumberOfRingsInListeningWindow is static and read only once.
	if all || !g.CanRead(6) {
		attributes = append(attributes, 6)
	}
	if g.Version > 1 {
		// Allowed callers.
		if all || !g.CanRead(7) {
			attributes = append(attributes, 7)
		}
	}
	return attributes
}

// GetNames returns the names of attribute indexes.
func (g *GXDLMSAutoAnswer) GetNames() []string {
	if g.Version > 1 {
		return []string{"Logical Name", "Mode", "Listening Window", "Status", "Number Of Calls", "Number Of Rings In Listening Window", "Allowed callers"}
	}
	return []string{"Logical Name", "Mode", "Listening Window", "Status", "Number Of Calls", "Number Of Rings In Listening Window"}
}

// GetMethodNames returns the names of method indexes.
func (g *GXDLMSAutoAnswer) GetMethodNames() []string {
	return []string{}
}

// GetAttributeCount returns the amount of attributes.
//
// Returns:
//
//	Count of attributes.
func (g *GXDLMSAutoAnswer) GetAttributeCount() int {
	if g.Version > 1 {
		return 7
	}
	return 6
}

// GetMethodCount returns the amount of methods.
func (g *GXDLMSAutoAnswer) GetMethodCount() int {
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
func (g *GXDLMSAutoAnswer) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	var err error
	if e.Index == 1 {
		v, err := helpers.LogicalNameToBytes(g.LogicalName())
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return v, err
	}
	if e.Index == 2 {
		return g.Mode, nil
	}
	if e.Index == 3 {
		cnt := len(g.ListeningWindow)
		data := types.NewGXByteBuffer()
		err = data.SetUint8(uint8(enums.DataTypeArray))
		if err != nil {
			return nil, err
		}
		err = types.SetObjectCount(cnt, data)
		if cnt != 0 {
			for _, it := range g.ListeningWindow {
				err = data.SetUint8(uint8(enums.DataTypeStructure))
				if err != nil {
					return nil, err
				}
				err = data.SetUint8(uint8(2))
				if err != nil {
					return nil, err
				}
				err = internal.SetData(settings, data, enums.DataTypeOctetString, it.Key)
				err = internal.SetData(settings, data, enums.DataTypeOctetString, it.Value)
			}
		}
		return data.Array(), nil
	}
	if e.Index == 4 {
		return g.Status, nil
	}
	if e.Index == 5 {
		return g.NumberOfCalls, nil
	}
	if e.Index == 6 {
		data := types.NewGXByteBuffer()
		err = data.SetUint8(uint8(enums.DataTypeStructure))
		if err != nil {
			return nil, err
		}
		err = types.SetObjectCount(2, data)
		err = internal.SetData(settings, data, enums.DataTypeUint8, g.NumberOfRingsInListeningWindow)
		err = internal.SetData(settings, data, enums.DataTypeUint8, g.NumberOfRingsOutListeningWindow)
		return data.Array(), nil
	}
	if e.Index == 7 {
		cnt := len(g.AllowedCallers)
		data := types.NewGXByteBuffer()
		err = data.SetUint8(uint8(enums.DataTypeArray))
		if err != nil {
			return nil, err
		}
		err = types.SetObjectCount(cnt, data)
		if cnt != 0 {
			for _, it := range g.AllowedCallers {
				err = data.SetUint8(uint8(enums.DataTypeStructure))
				if err != nil {
					return nil, err
				}
				err = data.SetUint8(uint8(2))
				if err != nil {
					return nil, err
				}
				err = internal.SetData(settings, data, enums.DataTypeOctetString, []byte(it.Key))
				if err != nil {
					return nil, err
				}
				err = internal.SetData(settings, data, enums.DataTypeEnum, byte(it.Value))
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
func (g *GXDLMSAutoAnswer) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	if e.Index == 1 {
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		err = g.SetLogicalName(ln)
	} else if e.Index == 2 {
		g.Mode = enums.AutoAnswerMode(e.Value.(types.GXEnum).Value)
	} else if e.Index == 3 {
		g.ListeningWindow = g.ListeningWindow[:0]
		if e.Value != nil {
			for _, tmp := range e.Value.(types.GXArray) {
				item := tmp.(types.GXStructure)
				start, err := internal.ChangeTypeFromByteArray(settings, item[0].([]byte), enums.DataTypeDateTime)
				if err != nil {
					return err
				}
				end, err := internal.ChangeTypeFromByteArray(settings, item[1].([]byte), enums.DataTypeDateTime)
				if err != nil {
					return err
				}
				g.ListeningWindow = append(g.ListeningWindow, *types.NewGXKeyValuePair(start.(types.GXDateTime), end.(types.GXDateTime)))
			}
		}
	} else if e.Index == 4 {
		g.Status = enums.AutoAnswerStatus(e.Value.(types.GXEnum).Value)
	} else if e.Index == 5 {
		g.NumberOfCalls = e.Value.(uint8)
	} else if e.Index == 6 {
		g.NumberOfRingsInListeningWindow = 0
		g.NumberOfRingsOutListeningWindow = 0

		if e.Value != nil {
			arr := e.Value.(types.GXStructure)
			g.NumberOfRingsInListeningWindow = arr[0].(uint8)
			g.NumberOfRingsOutListeningWindow = arr[1].(uint8)
		}
	} else if e.Index == 7 {
		g.AllowedCallers = g.AllowedCallers[:0]
		if e.Value != nil {
			for _, tmp := range e.Value.(types.GXArray) {
				item := tmp.(types.GXStructure)
				callerId := string(item[0].([]uint8))
				callType := enums.CallType(item[1].(types.GXEnum).Value)
				g.AllowedCallers = append(g.AllowedCallers, *types.NewGXKeyValuePair(callerId, callType))
			}
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
func (g *GXDLMSAutoAnswer) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, nil
}

// Load returns the load object content from XML.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSAutoAnswer) Load(reader *GXXmlReader) error {
	ret, err := reader.ReadElementContentAsInt("Mode", 0)
	if err != nil {
		return err
	}
	g.Mode = enums.AutoAnswerMode(ret)
	g.ListeningWindow = g.ListeningWindow[:0]
	if reader.isStartElementNamed2("ListeningWindow", true) {
		for reader.isStartElementNamed2("Item", true) {
			start, err := reader.ReadElementContentAsDateTime("Start", nil)
			if err != nil {
				return err
			}
			end, err := reader.ReadElementContentAsDateTime("End", nil)
			if err != nil {
				return err
			}
			g.ListeningWindow = append(g.ListeningWindow, *types.NewGXKeyValuePair(start, end))
		}
		reader.ReadEndElement("ListeningWindow")
	}
	ret, err = reader.ReadElementContentAsInt("Status", 0)
	if err != nil {
		return err
	}
	g.Status = enums.AutoAnswerStatus(ret)
	ret, err = reader.ReadElementContentAsInt("NumberOfCalls", 0)
	if err != nil {
		return err
	}
	g.NumberOfCalls = uint8(ret)
	ret, err = reader.ReadElementContentAsInt("NumberOfRingsInListeningWindow", 0)
	if err != nil {
		return err
	}
	g.NumberOfRingsInListeningWindow = uint8(ret)
	ret, err = reader.ReadElementContentAsInt("NumberOfRingsOutListeningWindow", 0)
	if err != nil {
		return err
	}
	g.NumberOfRingsOutListeningWindow = uint8(ret)
	if reader.isStartElementNamed2("AllowedCallers", true) {
		for reader.isStartElementNamed2("Item", true) {
			callerId, err := reader.ReadElementContentAsString("Id", "")
			if err != nil {
				return err
			}
			ret, err = reader.ReadElementContentAsInt("Type", 0)
			if err != nil {
				return err
			}
			callType := enums.CallType(ret)
			g.AllowedCallers = append(g.AllowedCallers, *types.NewGXKeyValuePair(callerId, callType))
		}
		reader.ReadEndElement("AllowedCallers")
	}
	return err
}

// Save returns the save object content to XML.
//
// Parameters:
//
//	writer: XML writer.
func (g *GXDLMSAutoAnswer) Save(writer *GXXmlWriter) error {
	err := writer.WriteElementString("Mode", int(g.Mode))
	if err != nil {
		return err
	}
	if g.ListeningWindow != nil {
		writer.WriteStartElement("ListeningWindow")
		for k, v := range g.ListeningWindow {
			writer.WriteStartElement("Item")
			err = writer.WriteElementString("Start", k)
			if err != nil {
				return err
			}
			err = writer.WriteElementString("End", v)
			if err != nil {
				return err
			}
			writer.WriteEndElement()
		}
		writer.WriteEndElement()
	}
	err = writer.WriteElementString("Status", int(g.Status))
	if err != nil {
		return err
	}
	err = writer.WriteElementString("NumberOfCalls", g.NumberOfCalls)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("NumberOfRingsInListeningWindow", g.NumberOfRingsInListeningWindow)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("NumberOfRingsOutListeningWindow", g.NumberOfRingsOutListeningWindow)
	if err != nil {
		return err
	}
	if g.AllowedCallers != nil {
		writer.WriteStartElement("AllowedCallers")
		for _, it := range g.AllowedCallers {
			writer.WriteStartElement("Item")
			err = writer.WriteElementString("Id", it.Key)
			if err != nil {
				return err
			}
			err = writer.WriteElementString("Type", int(it.Value))
			if err != nil {
				return err
			}
			writer.WriteEndElement()
		}
		writer.WriteEndElement()
	}
	return err
}

// PostLoad returns the handle actions after Load.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSAutoAnswer) PostLoad(reader *GXXmlReader) error {
	return nil
}

// GetValues returns the an array containing the COSEM object's attribute values.
func (g *GXDLMSAutoAnswer) GetValues() []any {
	return []any{g.LogicalName(), g.Mode, g.ListeningWindow, g.Status, g.NumberOfCalls,
		fmt.Sprintf("%d/%d", g.NumberOfRingsInListeningWindow, g.NumberOfRingsOutListeningWindow), g.AllowedCallers}
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
func (g *GXDLMSAutoAnswer) GetDataType(index int) (enums.DataType, error) {
	ret := enums.DataTypeNone
	if index == 1 {
		ret = enums.DataTypeOctetString
	} else if index == 2 {
		ret = enums.DataTypeEnum
	} else if index == 3 {
		ret = enums.DataTypeArray
	} else if index == 4 {
		ret = enums.DataTypeEnum
	} else if index == 5 {
		ret = enums.DataTypeUint8
	} else if index == 6 {
		ret = enums.DataTypeArray
	} else if index == 7 {
		if g.Version > 1 {
			ret = enums.DataTypeArray
		}
	}
	if ret == enums.DataTypeNone {
		return enums.DataTypeNone, errors.New("GetDataType failed. Invalid attribute index.")
	}
	return ret, nil
}

// Constructor.
// ln: Logical Name of the object.
// sn: Short Name of the object.
func NewGXDLMSAutoAnswer(ln string, sn int16) (*GXDLMSAutoAnswer, error) {
	err := ValidateLogicalName(ln)
	if err != nil {
		return nil, err
	}
	return &GXDLMSAutoAnswer{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeAutoAnswer,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
