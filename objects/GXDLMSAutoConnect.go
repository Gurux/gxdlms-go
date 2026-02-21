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
	"strings"

	"github.com/Gurux/gxdlms-go/dlmserrors"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// Auto Connect implements data transfer from the device to one or several destinations.
// Online help:
// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSAutoConnect
type GXDLMSAutoConnect struct {
	GXDLMSObject
	destinationsAsString bool

	// Defines the mode controlling the auto dial functionality concerning the
	// timing, the message type to be sent and the infrastructure to be used.
	Mode enums.AutoConnectMode

	// The maximum number of trials in the case of unsuccessful dialling attempts.
	Repetitions uint8

	// The time delay, expressed in seconds until an unsuccessful dial attempt can be repeated.
	RepetitionDelay uint16

	// Contains the start and end date/time stamp when the window becomes active.
	CallingWindow []types.GXKeyValuePair[types.GXDateTime, types.GXDateTime]

	// Contains the list of destinations (for example phone numbers, email
	// addresses or their combinations) where the message(s) have to be sent
	// under certain conditions. The conditions and their link to the elements of
	// the array are not defined here.
	Destinations []string
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSAutoConnect) Base() *GXDLMSObject {
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
func (g *GXDLMSAutoConnect) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	// LN is static and read only once.
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	// Mode
	if all || g.CanRead(2) {
		attributes = append(attributes, 2)
	}
	// Repetitions
	if all || g.CanRead(3) {
		attributes = append(attributes, 3)
	}
	// RepetitionDelay
	if all || g.CanRead(4) {
		attributes = append(attributes, 4)
	}
	// CallingWindow
	if all || g.CanRead(5) {
		attributes = append(attributes, 5)
	}
	// Destinations
	if all || g.CanRead(6) {
		attributes = append(attributes, 6)
	}
	return attributes
}

// GetNames returns the names of attribute indexes.
func (g *GXDLMSAutoConnect) GetNames() []string {
	return []string{"Logical Name", "Mode", "Repetitions", "Repetition Delay", "Calling Window", "Destinations"}
}

// GetMethodNames returns the names of method indexes.
func (g *GXDLMSAutoConnect) GetMethodNames() []string {
	if g.Version == 0 {
		return []string{}
	}
	return []string{"Connect"}
}

// GetAttributeCount returns the amount of attributes.
//
// Returns:
//
//	Count of attributes.
func (g *GXDLMSAutoConnect) GetAttributeCount() int {
	return 6
}

// GetMethodCount returns the amount of methods.
func (g *GXDLMSAutoConnect) GetMethodCount() int {
	if g.Version == 0 {
		return 0
	}
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
func (g *GXDLMSAutoConnect) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	var err error
	var ret any
	if e.Index == 1 {
		v, err := helpers.LogicalNameToBytes(g.LogicalName())
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return v, err
	} else if e.Index == 2 {
		ret = uint8(g.Mode)
	} else if e.Index == 3 {
		ret = g.Repetitions
	} else if e.Index == 4 {
		ret = g.RepetitionDelay
	} else if e.Index == 5 {
		cnt := len(g.CallingWindow)
		data := types.NewGXByteBuffer()
		err = data.SetUint8(uint8(enums.DataTypeArray))
		if err != nil {
			return nil, err
		}
		err = types.SetObjectCount(cnt, data)
		if cnt != 0 {
			for _, it := range g.CallingWindow {
				err = data.SetUint8(uint8(enums.DataTypeStructure))
				if err != nil {
					return nil, err
				}
				err = data.SetUint8(uint8(2))
				if err != nil {
					return nil, err
				}
				err = internal.SetData(settings, data, enums.DataTypeOctetString, it.Key)
				if err != nil {
					return nil, err
				}
				err = internal.SetData(settings, data, enums.DataTypeOctetString, it.Value)
				if err != nil {
					return nil, err
				}
			}
		}
		ret = data.Array()
	} else if e.Index == 6 {
		data := types.NewGXByteBuffer()
		err = data.SetUint8(uint8(enums.DataTypeArray))
		if err != nil {
			return nil, err
		}
		if g.Destinations == nil {
			err = types.SetObjectCount(0, data)
		} else {
			cnt := len(g.Destinations)
			err = types.SetObjectCount(cnt, data)
			for _, it := range g.Destinations {
				if g.destinationsAsString {
					err = internal.SetData(settings, data, enums.DataTypeString, it)
				} else {
					err = internal.SetData(settings, data, enums.DataTypeOctetString, []byte(it))
				}
			}
		}
		ret = data.Array()
	} else {
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
func (g *GXDLMSAutoConnect) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	var err error
	if e.Index == 1 {
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return g.SetLogicalName(ln)
	} else if e.Index == 2 {
		g.Mode = enums.AutoConnectMode(e.Value.(types.GXEnum).Value)
	} else if e.Index == 3 {
		g.Repetitions = e.Value.(uint8)
	} else if e.Index == 4 {
		g.RepetitionDelay = e.Value.(uint16)
	} else if e.Index == 5 {
		g.CallingWindow = g.CallingWindow[:0]
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
				g.CallingWindow = append(g.CallingWindow, *types.NewGXKeyValuePair(start.(types.GXDateTime), end.(types.GXDateTime)))
			}
		}
	} else if e.Index == 6 {
		g.Destinations = nil
		if e.Value != nil {
			arr := e.Value.(types.GXArray)
			items := []string{}
			for _, item := range arr {
				var it string
				if _, ok := item.([]byte); ok {
					ret, err := internal.ChangeTypeFromByteArray(settings, item.([]byte), enums.DataTypeString)
					if err != nil {
						return err
					}
					it = ret.(string)
				} else {
					it = fmt.Sprint(item)
					g.destinationsAsString = true
				}
				items = append(items, it)
			}
			g.Destinations = items
		}
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
func (g *GXDLMSAutoConnect) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	if e.Index != 1 {
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil, nil
}

// Load returns the load object content from XML.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSAutoConnect) Load(reader *GXXmlReader) error {
	ret, err := reader.ReadElementContentAsInt("Mode", 1)
	if err != nil {
		return err
	}
	g.Mode = enums.AutoConnectMode(ret)
	ret, err = reader.ReadElementContentAsInt("Repetitions", 1)
	if err != nil {
		return err
	}
	g.Repetitions = uint8(ret)
	ret, err = reader.ReadElementContentAsInt("RepetitionDelay", 1)
	if err != nil {
		return err
	}
	g.RepetitionDelay = uint16(ret)
	g.CallingWindow = g.CallingWindow[:0]
	if ret, err := reader.IsStartElementNamed("CallingWindow", true); ret && err == nil {
		for {
			ret, err = reader.IsStartElementNamed("Item", true)
			if err != nil {
				return err
			}
			if !ret {
				break
			}
			start, err := reader.ReadElementContentAsDateTime("Start", nil)
			if err != nil {
				return err
			}
			end, err := reader.ReadElementContentAsDateTime("End", nil)
			if err != nil {
				return err
			}
			g.CallingWindow = append(g.CallingWindow, *types.NewGXKeyValuePair(start, end))
		}
		reader.ReadEndElement("CallingWindow")
	}
	ret2, err := reader.ReadElementContentAsString("Destinations", "")
	if err != nil {
		return err
	}
	g.Destinations = strings.Split(ret2, ";")
	return err
}

// Save returns the save object content to XML.
//
// Parameters:
//
//	writer: XML writer.
func (g *GXDLMSAutoConnect) Save(writer *GXXmlWriter) error {
	err := writer.WriteElementString("Mode", int(g.Mode))
	if err != nil {
		return err
	}
	err = writer.WriteElementString("Repetitions", g.Repetitions)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("RepetitionDelay", g.RepetitionDelay)
	if err != nil {
		return err
	}
	writer.WriteStartElement("CallingWindow")
	if g.CallingWindow != nil {
		for _, it := range g.CallingWindow {
			writer.WriteStartElement("Item")
			writer.WriteElementString("Start", it.Key)
			writer.WriteElementString("End", it.Value)
			writer.WriteEndElement()
		}
	}
	writer.WriteEndElement()
	if g.Destinations != nil {
		writer.WriteElementString("Destinations", strings.Join(g.Destinations, ";"))
	}
	return err
}

// PostLoad returns the handle actions after Load.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSAutoConnect) PostLoad(reader *GXXmlReader) error {
	return nil
}

// GetValues returns the an array containing the COSEM object's attribute values.
func (g *GXDLMSAutoConnect) GetValues() []any {
	return []any{g.LogicalName(), g.Mode, g.Repetitions, g.RepetitionDelay, g.CallingWindow, g.Destinations}
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
func (g *GXDLMSAutoConnect) GetDataType(index int) (enums.DataType, error) {
	ret := enums.DataTypeNone
	if index == 1 {
		ret = enums.DataTypeOctetString
	} else if index == 2 {
		ret = enums.DataTypeEnum
	} else if index == 3 {
		ret = enums.DataTypeUint8
	} else if index == 4 {
		ret = enums.DataTypeUint16
	} else if index == 5 {
		ret = enums.DataTypeArray
	} else if index == 6 {
		ret = enums.DataTypeArray
	} else {
		return enums.DataTypeNone, dlmserrors.ErrInvalidAttributeIndex
	}
	return ret, nil
}

// Connect returns the initiates the connection process.
//
// Parameters:
//
//	client: DLMS client.
//
// Returns:
//
//	Action bytes.
func (g *GXDLMSAutoConnect) Connect(client IGXDLMSClient) ([][]byte, error) {
	return client.Method(g, 1, int8(0), enums.DataTypeInt8)
}

// NewGXDLMSAutoConnect creates a new auto connect object instance.
//
// The function validates `ln` before creating the object.
//`ln` is the Logical Name and `sn` is the Short Name of the object.
func NewGXDLMSAutoConnect(ln string, sn int16) (*GXDLMSAutoConnect, error) {
	err := ValidateLogicalName(ln)
	if err != nil {
		return nil, err
	}
	return &GXDLMSAutoConnect{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeAutoConnect,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
