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
	"time"

	"github.com/Gurux/gxdlms-go/dlmserrors"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// Online help:
// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSGSMDiagnostic
type GXDLMSGSMDiagnostic struct {
	GXDLMSObject
	// Name of network operator.
	Operator string

	// Registration status of the modem.
	Status enums.GsmStatus

	// Registration status of the modem.
	CircuitSwitchStatus enums.GsmCircuitSwitchStatus

	// Registration status of the modem.
	PacketSwitchStatus enums.GsmPacketSwitchStatus

	// Registration status of the modem.
	CellInfo GXDLMSGSMCellInfo

	// Registration status of the modem.
	AdjacentCells []AdjacentCell

	// Date and time when the data have been last captured.
	CaptureTime types.GXDateTime
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSGSMDiagnostic) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

//Invoke returns the invokes method.
//
// Parameters:
//   settings: DLMS settings.
//   e: Invoke parameters.
func (g *GXDLMSGSMDiagnostic) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, nil
}

//GetAttributeIndexToRead returns the collection of attributes to read.
// If attribute is static and already read or device is returned HW error it is not returned.
//
// Parameters:
//   all: All items are returned even if they are read already.
//
// Returns:
//   Collection of attributes to read.
func (g *GXDLMSGSMDiagnostic) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	// LN is static and read only once.
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	// Operator
	if all || g.CanRead(2) {
		attributes = append(attributes, 2)
	}
	// Status
	if all || g.CanRead(3) {
		attributes = append(attributes, 3)
	}
	// CircuitSwitchStatus
	if all || g.CanRead(4) {
		attributes = append(attributes, 4)
	}
	// PacketSwitchStatus
	if all || g.CanRead(5) {
		attributes = append(attributes, 5)
	}
	// CellInfo
	if all || g.CanRead(6) {
		attributes = append(attributes, 6)
	}
	// AdjacentCells
	if all || g.CanRead(7) {
		attributes = append(attributes, 7)
	}
	// CaptureTime
	if all || g.CanRead(8) {
		attributes = append(attributes, 8)
	}
	return attributes
}

//GetNames returns the names of attribute indexes.
func (g *GXDLMSGSMDiagnostic) GetNames() []string {
	return []string{"Logical Name", "Operator", "Status", "CircuitSwitchStatus", "PacketSwitchStatus", "CellInfo", "AdjacentCells", "CaptureTime"}
}

//GetMethodNames returns the names of method indexes.
func (g *GXDLMSGSMDiagnostic) GetMethodNames() []string {
	return []string{}
}

//GetAttributeCount returns the amount of attributes.
//
// Returns:
//   Count of attributes.
func (g *GXDLMSGSMDiagnostic) GetAttributeCount() int {
	return 8
}

//GetMethodCount returns the amount of methods.
func (g *GXDLMSGSMDiagnostic) GetMethodCount() int {
	return 0
}

func anyToUInt32(v any) uint32 {
	switch n := v.(type) {
	case uint8:
		return uint32(n)
	case uint16:
		return uint32(n)
	case uint32:
		return n
	case uint64:
		return uint32(n)
	case int8:
		return uint32(n)
	case int16:
		return uint32(n)
	case int32:
		return uint32(n)
	case int:
		return uint32(n)
	case types.GXEnum:
		return uint32(n.Value)
	default:
		return 0
	}
}

//GetValue returns the value of given attribute.
// When raw parameter us not used example register multiplies value by scalar.
//
// Parameters:
//   settings: DLMS settings.
//   e: Get parameters.
//
// Returns:
//   Value of the attribute index.
func (g *GXDLMSGSMDiagnostic) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	switch e.Index {
	case 1:
		ret, err := helpers.LogicalNameToBytes(g.LogicalName())
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return nil, err
		}
		return ret, nil
	case 2:
		if g.Operator == "" {
			return nil, nil
		}
		return []byte(g.Operator), nil
	case 3:
		return g.Status, nil
	case 4:
		return g.CircuitSwitchStatus, nil
	case 5:
		return g.PacketSwitchStatus, nil
	case 6:
		bb := types.NewGXByteBuffer()
		if err := bb.SetUint8(uint8(enums.DataTypeStructure)); err != nil {
			return nil, err
		}
		if g.Version == 0 {
			if err := bb.SetUint8(4); err != nil {
				return nil, err
			}
			if err := internal.SetData(settings, bb, enums.DataTypeUint16, uint16(g.CellInfo.CellId)); err != nil {
				return nil, err
			}
		} else {
			if err := bb.SetUint8(7); err != nil {
				return nil, err
			}
			if err := internal.SetData(settings, bb, enums.DataTypeUint32, g.CellInfo.CellId); err != nil {
				return nil, err
			}
		}
		if err := internal.SetData(settings, bb, enums.DataTypeUint16, g.CellInfo.LocationId); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeUint8, g.CellInfo.SignalQuality); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeUint8, g.CellInfo.Ber); err != nil {
			return nil, err
		}
		if g.Version > 0 {
			if err := internal.SetData(settings, bb, enums.DataTypeUint16, g.CellInfo.MobileCountryCode); err != nil {
				return nil, err
			}
			if err := internal.SetData(settings, bb, enums.DataTypeUint16, g.CellInfo.MobileNetworkCode); err != nil {
				return nil, err
			}
			if err := internal.SetData(settings, bb, enums.DataTypeUint32, g.CellInfo.ChannelNumber); err != nil {
				return nil, err
			}
		}
		return bb.Array(), nil
	case 7:
		bb := types.NewGXByteBuffer()
		if err := bb.SetUint8(uint8(enums.DataTypeArray)); err != nil {
			return nil, err
		}
		if err := types.SetObjectCount(len(g.AdjacentCells), bb); err != nil {
			return nil, err
		}
		cellIDType := enums.DataType(enums.DataTypeUint32)
		for _, it := range g.AdjacentCells {
			if err := bb.SetUint8(uint8(enums.DataTypeStructure)); err != nil {
				return nil, err
			}
			if err := bb.SetUint8(2); err != nil {
				return nil, err
			}
			if g.Version == 0 {
				cellIDType = enums.DataTypeUint16
				if err := internal.SetData(settings, bb, cellIDType, uint16(it.CellID)); err != nil {
					return nil, err
				}
			} else {
				if err := internal.SetData(settings, bb, cellIDType, it.CellID); err != nil {
					return nil, err
				}
			}
			if err := internal.SetData(settings, bb, enums.DataTypeUint8, it.SignalQuality); err != nil {
				return nil, err
			}
		}
		return bb.Array(), nil
	case 8:
		return g.CaptureTime, nil
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
		return nil, nil
	}
}

//SetValue returns the set value of given attribute.
// When raw parameter us not used example register multiplies value by scalar.
//
// Parameters:
//   settings: DLMS settings.
//   e: Set parameters.
func (g *GXDLMSGSMDiagnostic) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	switch e.Index {
	case 1:
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return err
		}
		return g.SetLogicalName(ln)
	case 2:
		if e.Value == nil {
			g.Operator = ""
		} else if v, ok := e.Value.([]byte); ok {
			g.Operator = string(v)
		} else {
			g.Operator = e.Value.(string)
		}
	case 3:
		if v, ok := e.Value.(types.GXEnum); ok {
			g.Status = enums.GsmStatus(v.Value)
		} else {
			g.Status = enums.GsmStatus(e.Value.(uint8))
		}
	case 4:
		if v, ok := e.Value.(types.GXEnum); ok {
			g.CircuitSwitchStatus = enums.GsmCircuitSwitchStatus(v.Value)
		} else {
			g.CircuitSwitchStatus = enums.GsmCircuitSwitchStatus(e.Value.(uint8))
		}
	case 5:
		if v, ok := e.Value.(types.GXEnum); ok {
			g.PacketSwitchStatus = enums.GsmPacketSwitchStatus(v.Value)
		} else {
			g.PacketSwitchStatus = enums.GsmPacketSwitchStatus(e.Value.(uint8))
		}
	case 6:
		if e.Value != nil {
			var tmp types.GXStructure
			switch v := e.Value.(type) {
			case types.GXStructure:
				tmp = v
			case []any:
				tmp = types.GXStructure(v)
			default:
				e.Error = enums.ErrorCodeReadWriteDenied
				return errors.New("setValue failed. Invalid CellInfo value.")
			}
			g.CellInfo.CellId = anyToUInt32(tmp[0])
			g.CellInfo.LocationId = uint16(anyToUInt32(tmp[1]))
			g.CellInfo.SignalQuality = uint8(anyToUInt32(tmp[2]))
			g.CellInfo.Ber = uint8(anyToUInt32(tmp[3]))
			if g.Version > 0 {
				g.CellInfo.MobileCountryCode = uint16(anyToUInt32(tmp[4]))
				g.CellInfo.MobileNetworkCode = uint16(anyToUInt32(tmp[5]))
				g.CellInfo.ChannelNumber = anyToUInt32(tmp[6])
			}
		}
	case 7:
		g.AdjacentCells = g.AdjacentCells[:0]
		if e.Value != nil {
			var list types.GXArray
			switch v := e.Value.(type) {
			case types.GXArray:
				list = v
			case []any:
				list = types.GXArray(v)
			default:
				e.Error = enums.ErrorCodeReadWriteDenied
				return errors.New("setValue failed. Invalid AdjacentCells value.")
			}
			for _, tmp := range list {
				var it types.GXStructure
				switch v := tmp.(type) {
				case types.GXStructure:
					it = v
				case []any:
					it = types.GXStructure(v)
				default:
					e.Error = enums.ErrorCodeReadWriteDenied
					return errors.New("setValue failed. Invalid AdjacentCell item.")
				}
				g.AdjacentCells = append(g.AdjacentCells, AdjacentCell{
					CellID:        anyToUInt32(it[0]),
					SignalQuality: uint8(anyToUInt32(it[1])),
				})
			}
		}
	case 8:
		if e.Value == nil {
			g.CaptureTime = types.GXDateTime{}
			return nil
		}
		if v, ok := e.Value.([]byte); ok {
			ret, err := internal.ChangeTypeFromByteArray(settings, v, enums.DataTypeDateTime)
			if err != nil {
				e.Error = enums.ErrorCodeReadWriteDenied
				return err
			}
			g.CaptureTime = ret.(types.GXDateTime)
		} else if v, ok := e.Value.(string); ok {
			dt, err := types.NewGXDateTimeFromString(v, nil)
			if err != nil {
				e.Error = enums.ErrorCodeReadWriteDenied
				return err
			}
			g.CaptureTime = *dt
		} else if v, ok := e.Value.(time.Time); ok {
			g.CaptureTime = *types.NewGXDateTimeFromTime(v)
		} else {
			g.CaptureTime = e.Value.(types.GXDateTime)
		}
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil
}

//Load returns the load object content from XML.
//
// Parameters:
//   reader: XML reader.
func (g *GXDLMSGSMDiagnostic) Load(reader *GXXmlReader) error {
	var err error
	g.Operator, err = reader.ReadElementContentAsString("Operator", "")
	if err != nil {
		return err
	}
	ret, err := reader.ReadElementContentAsInt("Status", 0)
	if err != nil {
		return err
	}
	g.Status = enums.GsmStatus(ret)
	ret, err = reader.ReadElementContentAsInt("CircuitSwitchStatus", 0)
	if err != nil {
		return err
	}
	g.CircuitSwitchStatus = enums.GsmCircuitSwitchStatus(ret)
	ret, err = reader.ReadElementContentAsInt("PacketSwitchStatus", 0)
	if err != nil {
		return err
	}
	g.PacketSwitchStatus = enums.GsmPacketSwitchStatus(ret)
	if ret, err := reader.IsStartElementNamed("CellInfo", true); ret && err == nil {
		g.CellInfo.CellId, err = reader.ReadElementContentAsUInt32("CellId", 0)
		if err != nil {
			return err
		}
		g.CellInfo.LocationId, err = reader.ReadElementContentAsUInt16("LocationId", 0)
		if err != nil {
			return err
		}
		g.CellInfo.SignalQuality, err = reader.ReadElementContentAsUInt8("SignalQuality", 0)
		if err != nil {
			return err
		}
		g.CellInfo.Ber, err = reader.ReadElementContentAsUInt8("Ber", 0)
		if err != nil {
			return err
		}
		reader.ReadEndElement("CellInfo")
	}
	g.AdjacentCells = g.AdjacentCells[:0]
	if ret, err := reader.IsStartElementNamed("AdjacentCells", true); ret && err == nil {
		for {
			ret, err = reader.IsStartElementNamed("Item", true)
			if err != nil {
				return err
			}
			if !ret {
				break
			}
			it := AdjacentCell{}
			it.CellID, err = reader.ReadElementContentAsUInt32("CellId", 0)
			if err != nil {
				return err
			}
			it.SignalQuality, err = reader.ReadElementContentAsUInt8("SignalQuality", 0)
			if err != nil {
				return err
			}
			g.AdjacentCells = append(g.AdjacentCells, it)
		}
		reader.ReadEndElement("AdjacentCells")
	}
	g.CaptureTime, err = reader.ReadElementContentAsDateTime("CaptureTime", nil)
	if err != nil {
		return err
	}
	return err
}

//Save returns the save object content to XML.
//
// Parameters:
//   writer: XML writer.
func (g *GXDLMSGSMDiagnostic) Save(writer *GXXmlWriter) error {
	err := writer.WriteElementString("Operator", g.Operator)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("Status", int(g.Status))
	if err != nil {
		return err
	}
	err = writer.WriteElementString("CircuitSwitchStatus", int(g.CircuitSwitchStatus))
	if err != nil {
		return err
	}
	err = writer.WriteElementString("PacketSwitchStatus", int(g.PacketSwitchStatus))
	if err != nil {
		return err
	}
	writer.WriteStartElement("CellInfo")
	err = writer.WriteElementString("CellId", g.CellInfo.CellId)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("LocationId", g.CellInfo.LocationId)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("SignalQuality", g.CellInfo.SignalQuality)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("Ber", g.CellInfo.Ber)
	if err != nil {
		return err
	}
	writer.WriteEndElement()
	writer.WriteStartElement("AdjacentCells")
	if g.AdjacentCells != nil {
		for _, it := range g.AdjacentCells {
			writer.WriteStartElement("Item")
			err = writer.WriteElementString("CellId", it.CellID)
			if err != nil {
				return err
			}
			err = writer.WriteElementString("SignalQuality", it.SignalQuality)
			if err != nil {
				return err
			}
			writer.WriteEndElement()
		}
	}
	writer.WriteEndElement()
	err = writer.WriteElementString("CaptureTime", g.CaptureTime)
	if err != nil {
		return err
	}
	return err
}

//PostLoad returns the handle actions after Load.
//
// Parameters:
//   reader: XML reader.
func (g *GXDLMSGSMDiagnostic) PostLoad(reader *GXXmlReader) error {
	return nil
}

//GetValues returns the an array containing the COSEM object's attribute values.
func (g *GXDLMSGSMDiagnostic) GetValues() []any {
	return []any{g.LogicalName(), g.Operator, g.Status,
		g.CircuitSwitchStatus, g.PacketSwitchStatus,
		g.CellInfo, g.AdjacentCells, g.CaptureTime}
}

//GetDataType returns the device data type of selected attribute index.
//
// Parameters:
//   index: Attribute index of the object.
//
// Returns:
//   Device data type of the object.
func (g *GXDLMSGSMDiagnostic) GetDataType(index int) (enums.DataType, error) {
	switch index {
	case 1:
		return enums.DataTypeOctetString, nil
	case 2:
		return enums.DataTypeString, nil
	case 3:
		return enums.DataTypeEnum, nil
	case 4:
		return enums.DataTypeEnum, nil
	case 5:
		return enums.DataTypeEnum, nil
	case 6:
		return enums.DataTypeStructure, nil
	case 7:
		return enums.DataTypeArray, nil
	case 8:
		return enums.DataTypeDateTime, nil
	}
	return 0, dlmserrors.ErrInvalidAttributeIndex
}

// NewGXDLMGSMDiagnostic creates a new G S M Diagnostic object instance.
//
// The function validates `ln` before creating the object.
//`ln` is the Logical Name and `sn` is the Short Name of the object.
func NewGXDLMSGSMDiagnostic(ln string, sn int16) (*GXDLMSGSMDiagnostic, error) {
	err := ValidateLogicalName(ln)
	if err != nil {
		return nil, err
	}
	return &GXDLMSGSMDiagnostic{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeGSMDiagnostic,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
