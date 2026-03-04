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

	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// Online help:
// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSCharge
type GXDLMSCharge struct {
	GXDLMSObject
	// Total amount paid
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSCharge
	TotalAmountPaid int32

	// Charge type.
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSCharge
	ChargeType enums.ChargeType

	// Priority.
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSCharge
	Priority uint8

	// Unit charge active.
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSCharge
	UnitChargeActive GXUnitCharge

	// Unit charge passive.
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSCharge
	UnitChargePassive GXUnitCharge

	// Unit charge activation time.
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSCharge
	UnitChargeActivationTime types.GXDateTime

	// Period.
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSCharge
	Period uint32

	// Charge configuration,
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSCharge
	ChargeConfiguration enums.ChargeConfiguration

	// Last collection time.
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSCharge
	LastCollectionTime types.GXDateTime

	// Last collection amount.
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSCharge
	LastCollectionAmount int32

	// Total amount remaining.
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSCharge
	TotalAmountRemaining int32

	// Proportion.
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSCharge
	Proportion uint16
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSCharge) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

// Invoke returns the invokes method.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Invoke parameters.
func (g *GXDLMSCharge) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
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
func (g *GXDLMSCharge) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	// LN is static and read only once.
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	// TotalAmountPaid
	if all || g.CanRead(2) {
		attributes = append(attributes, 2)
	}
	// ChargeType
	if all || g.CanRead(3) {
		attributes = append(attributes, 3)
	}
	// Priority
	if all || g.CanRead(4) {
		attributes = append(attributes, 4)
	}
	// UnitChargeActive
	if all || g.CanRead(5) {
		attributes = append(attributes, 5)
	}
	// UnitChargePassive
	if all || g.CanRead(6) {
		attributes = append(attributes, 6)
	}
	// UnitChargeActivationTime
	if all || g.CanRead(7) {
		attributes = append(attributes, 7)
	}
	// Period
	if all || g.CanRead(8) {
		attributes = append(attributes, 8)
	}
	// ChargeConfiguration
	if all || g.CanRead(9) {
		attributes = append(attributes, 9)
	}
	// LastCollectionTime
	if all || g.CanRead(10) {
		attributes = append(attributes, 10)
	}
	// LastCollectionAmount
	if all || g.CanRead(11) {
		attributes = append(attributes, 11)
	}
	// TotalAmountRemaining
	if all || g.CanRead(12) {
		attributes = append(attributes, 12)
	}
	// Proportion
	if all || g.CanRead(13) {
		attributes = append(attributes, 13)
	}
	return attributes
}

// GetNames returns the names of attribute indexes.
func (g *GXDLMSCharge) GetNames() []string {
	return []string{"Logical Name", "TotalAmountPaid", "ChargeType", "Priority", "UnitChargeActive", "UnitChargePassive", "UnitChargeActivationTime", "Period", "ChargeConfiguration", "LastCollectionTime", "LastCollectionAmount", "TotalAmountRemaining", "Proportion"}
}

// GetMethodNames returns the names of method indexes.
func (g *GXDLMSCharge) GetMethodNames() []string {
	return []string{"Update unit charge", "Activate passive unit charge", "Collect", "Update total amount remaining", "Set total amount remaining"}
}

// GetAttributeCount returns the amount of attributes.
//
// Returns:
//
//	Count of attributes.
func (g *GXDLMSCharge) GetAttributeCount() int {
	return 13
}

// GetMethodCount returns the amount of methods.
func (g *GXDLMSCharge) GetMethodCount() int {
	return 5
}

func (g *GXDLMSCharge) GetUnitCharge(settings *settings.GXDLMSSettings, charge *GXUnitCharge) ([]byte, error) {
	bb := types.NewGXByteBuffer()
	err := bb.SetUint8(byte(enums.DataTypeStructure))
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(3)
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(byte(enums.DataTypeStructure))
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(2)
	if err != nil {
		return nil, err
	}
	err = internal.SetData(settings, bb, enums.DataTypeInt8, charge.ChargePerUnitScaling.CommodityScale)
	if err != nil {
		return nil, err
	}
	err = internal.SetData(settings, bb, enums.DataTypeInt8, charge.ChargePerUnitScaling.PriceScale)
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(byte(enums.DataTypeStructure))
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(3)
	if err != nil {
		return nil, err
	}
	if charge.Commodity.Target == nil {
		err = internal.SetData(settings, bb, enums.DataTypeUint16, uint16(0))
		if err != nil {
			return nil, err
		}
		err = bb.SetUint8(enums.DataTypeOctetString)
		if err != nil {
			return nil, err
		}
		err = bb.SetUint8(6)
		if err != nil {
			return nil, err
		}
		err = bb.SetUint8(0)
		if err != nil {
			return nil, err
		}
		err = bb.SetUint8(0)
		if err != nil {
			return nil, err
		}
		err = bb.SetUint8(0)
		if err != nil {
			return nil, err
		}
		err = bb.SetUint8(0)
		if err != nil {
			return nil, err
		}
		err = bb.SetUint8(0)
		if err != nil {
			return nil, err
		}
		err = bb.SetUint8(0)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, bb, enums.DataTypeInt8, int8(0))
		if err != nil {
			return nil, err
		}
	} else {
		err = internal.SetData(settings, bb, enums.DataTypeUint16, charge.Commodity.Target.Base().ObjectType())
		if err != nil {
			return nil, err
		}
		ln, err := helpers.LogicalNameToBytes(charge.Commodity.Target.Base().LogicalName())
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, bb, enums.DataTypeOctetString, ln)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, bb, enums.DataTypeInt8, charge.Commodity.Index)
		if err != nil {
			return nil, err
		}
	}
	err = bb.SetUint8(enums.DataTypeArray)
	if err != nil {
		return nil, err
	}
	if charge.ChargeTables == nil {
		err = bb.SetUint8(0)
		if err != nil {
			return nil, err
		}
	} else {
		types.SetObjectCount(len(charge.ChargeTables), bb)
		for _, it := range charge.ChargeTables {
			err = bb.SetUint8(enums.DataTypeStructure)
			if err != nil {
				return nil, err
			}
			err = bb.SetUint8(2)
			if err != nil {
				return nil, err
			}
			err = internal.SetData(settings, bb, enums.DataTypeOctetString, []byte(it.Index))
			if err != nil {
				return nil, err
			}
			err = internal.SetData(settings, bb, enums.DataTypeInt16, it.ChargePerUnit)
			if err != nil {
				return nil, err
			}
		}
	}
	return bb.Array(), nil
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
func (g *GXDLMSCharge) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	var ret any
	var err error
	switch e.Index {
	case 1:
		return helpers.LogicalNameToBytes(g.LogicalName())
	case 2:
		ret = g.TotalAmountPaid
	case 3:
		ret = uint8(g.ChargeType)
	case 4:
		ret = g.Priority
	case 5:
		ret, err = g.GetUnitCharge(settings, &g.UnitChargeActive)
	case 6:
		ret, err = g.GetUnitCharge(settings, &g.UnitChargePassive)
	case 7:
		ret = g.UnitChargeActivationTime
	case 8:
		ret = g.Period
	case 9:
		ret, err = types.NewGXBitStringFromInteger(int(g.ChargeConfiguration), 2)
	case 10:
		ret = g.LastCollectionTime
	case 11:
		ret = g.LastCollectionAmount
	case 12:
		ret = g.TotalAmountRemaining
	case 13:
		ret = g.Proportion
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return ret, err
}

func (g *GXDLMSCharge) SetUnitCharge(charge *GXUnitCharge, value any, objects *GXDLMSObjectCollection) error {
	if value != nil {
		tmp := value.(types.GXStructure)
		tmp2 := tmp[0].(types.GXStructure)
		charge.ChargePerUnitScaling.CommodityScale = tmp2[0].(int8)
		charge.ChargePerUnitScaling.PriceScale = tmp2[1].(int8)
		tmp2 = tmp[1].(types.GXStructure)
		ot := enums.ObjectType(tmp2[0].(uint16))
		ln, err := helpers.ToLogicalName(tmp2[1])
		if err != nil {
			return err
		}
		if ot != enums.ObjectTypeNone {
			if g.parent != nil {
				charge.Commodity.Target = objects.FindByLN(ot, ln)
				// If object is not found from the association view.
				if charge.Commodity.Target == nil {
					charge.Commodity.Target, err = CreateObject(ot, ln, 0)
					if err != nil {
						return err
					}
				}
			}
		} else {
			charge.Commodity.Target = nil
		}
		charge.Commodity.Index = tmp2[2].(int8)
		charge.ChargeTables = charge.ChargeTables[:0]
		for _, tmp3 := range tmp[2].(types.GXArray) {
			it := tmp3.(types.GXStructure)
			item := GXChargeTable{}
			item.Index = string(it[0].([]uint8))
			item.ChargePerUnit = it[1].(int16)
			charge.ChargeTables = append(charge.ChargeTables, item)
		}
	} else {
		charge.ChargePerUnitScaling.CommodityScale = 0
		charge.ChargePerUnitScaling.PriceScale = 0
		charge.Commodity.Target = nil
		charge.Commodity.Index = 0
	}
	return nil
}

// SetValue returns the set value of given attribute.
// When raw parameter us not used example register multiplies value by scalar.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Set parameters.
func (g *GXDLMSCharge) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	var err error
	switch e.Index {
	case 1:
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return g.SetLogicalName(ln)
	case 2:
		g.TotalAmountPaid = e.Value.(int32)
	case 3:
		g.ChargeType = enums.ChargeType(e.Value.(types.GXEnum).Value)
	case 4:
		g.Priority = e.Value.(uint8)
	case 5:
		err = g.SetUnitCharge(&g.UnitChargeActive, e.Value, getObjectCollection(settings.Objects))
	case 6:
		err = g.SetUnitCharge(&g.UnitChargePassive, e.Value, getObjectCollection(settings.Objects))
	case 7:
		if v, ok := e.Value.(types.GXDateTime); ok {
			g.UnitChargeActivationTime = v
		} else if v, ok := e.Value.([]byte); ok {
			val, err := internal.ChangeTypeFromByteArray(settings, v, enums.DataTypeDateTime)
			if err != nil {
				return err
			}
			g.UnitChargeActivationTime = val.(types.GXDateTime)
		} else if v, ok := e.Value.(string); ok {
			ret, err := types.NewGXDateTimeFromString(v, nil)
			if err != nil {
				return err
			}
			g.UnitChargeActivationTime = *ret
		} else {
			g.UnitChargeActivationTime = types.GXDateTime{}
		}
	case 8:
		g.Period = e.Value.(uint32)
	case 9:
		bs := e.Value.(types.GXBitString)
		g.ChargeConfiguration = enums.ChargeConfiguration(bs.ToInteger())
	case 10:
		if v, ok := e.Value.(types.GXDateTime); ok {
			g.LastCollectionTime = v
		} else if v, ok := e.Value.(time.Time); ok {
			g.LastCollectionTime = *types.NewGXDateTimeFromTime(v)
		} else if v, ok := e.Value.([]byte); ok {
			val, err := internal.ChangeTypeFromByteArray(settings, v, enums.DataTypeDateTime)
			if err != nil {
				return err
			}
			g.LastCollectionTime = val.(types.GXDateTime)
		} else if v, ok := e.Value.(string); ok {
			ret, err := types.NewGXDateTimeFromString(v, nil)
			if err != nil {
				return err
			}
			g.LastCollectionTime = *ret
		} else {
			g.LastCollectionTime = types.GXDateTime{}
		}
	case 11:
		g.LastCollectionAmount = e.Value.(int32)
	case 12:
		g.TotalAmountRemaining = e.Value.(int32)
	case 13:
		g.Proportion = e.Value.(uint16)
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return err
}

func (g *GXDLMSCharge) LoadUnitChargeActive(reader *GXXmlReader,
	name string,
	charge *GXUnitCharge) error {
	var err error
	if ret, err := reader.IsStartElementNamed(name, true); ret && err == nil {
		charge.ChargePerUnitScaling.CommodityScale, err = reader.ReadElementContentAsInt8("Scale", 0)
		if err != nil {
			return err
		}
		charge.ChargePerUnitScaling.PriceScale, err = reader.ReadElementContentAsInt8("PriceScale", 0)
		if err != nil {
			return err
		}
		ret, err := reader.ReadElementContentAsInt("Type", 0)
		if err != nil {
			return err
		}
		ot := enums.ObjectType(ret)
		ln, err := reader.ReadElementContentAsString("Ln", "")
		if err != nil {
			return err
		}
		charge.Commodity.Target = reader.Objects.FindByLN(ot, ln)
		charge.Commodity.Index, err = reader.ReadElementContentAsInt8("Index", 0)
		if err != nil {
			return err
		}
		charge.ChargeTables = charge.ChargeTables[:0]
		if ret, err := reader.IsStartElementNamed("ChargeTables", true); ret && err == nil {
			for {
				ret, err = reader.IsStartElementNamed("Time", true)
				if err != nil {
					return err
				}
				if !ret {
					break
				}
				it := GXChargeTable{}
				it.Index, err = reader.ReadElementContentAsString("Index", "")
				if err != nil {
					return err
				}
				it.ChargePerUnit, err = reader.ReadElementContentAsInt16("ChargePerUnit", 0)
				if err != nil {
					return err
				}
				charge.ChargeTables = append(charge.ChargeTables, it)
			}
			reader.ReadEndElement("ChargeTables")
		}
		reader.ReadEndElement(name)
	}
	return err
}

// Load returns the load object content from XML.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSCharge) Load(reader *GXXmlReader) error {
	var err error
	g.TotalAmountPaid, err = reader.ReadElementContentAsInt32("TotalAmountPaid", 0)
	if err != nil {
		return err
	}
	ret, err := reader.ReadElementContentAsInt("ChargeType", 0)
	if err != nil {
		return err
	}
	g.ChargeType = enums.ChargeType(ret)
	g.Priority, err = reader.ReadElementContentAsUInt8("Priority", 0)
	if err != nil {
		return err
	}
	err = g.LoadUnitChargeActive(reader, "UnitChargeActive", &g.UnitChargeActive)
	if err != nil {
		return err
	}
	err = g.LoadUnitChargeActive(reader, "UnitChargePassive", &g.UnitChargePassive)
	if err != nil {
		return err
	}
	g.UnitChargeActivationTime, err = reader.ReadElementContentAsGXDateTime("UnitChargeActivationTime")
	if err != nil {
		return err
	}
	g.Period, err = reader.ReadElementContentAsUInt32("Period", 0)
	if err != nil {
		return err
	}
	ret, err = reader.ReadElementContentAsInt("ChargeConfiguration", 0)
	if err != nil {
		return err
	}
	g.ChargeConfiguration = enums.ChargeConfiguration(ret)
	g.LastCollectionTime, err = reader.ReadElementContentAsGXDateTime("LastCollectionTime")
	if err != nil {
		return err
	}
	g.LastCollectionAmount, err = reader.ReadElementContentAsInt32("LastCollectionAmount", 0)
	if err != nil {
		return err
	}
	g.TotalAmountRemaining, err = reader.ReadElementContentAsInt32("TotalAmountRemaining", 0)
	if err != nil {
		return err
	}
	g.Proportion, err = reader.ReadElementContentAsUInt16("Proportion", 0)
	return err
}

func (g *GXDLMSCharge) SaveUnitCharge(writer *GXXmlWriter, name string, charge *GXUnitCharge) error {
	writer.WriteStartElement(name)
	err := writer.WriteElementString("Scale", charge.ChargePerUnitScaling.CommodityScale)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("PriceScale", charge.ChargePerUnitScaling.PriceScale)
	if err != nil {
		return err
	}
	if charge.Commodity.Target != nil {
		writer.WriteElementString("Type", int(charge.Commodity.Target.Base().ObjectType()))
		writer.WriteElementString("Ln", charge.Commodity.Target.Base().LogicalName())
	} else {
		err = writer.WriteElementString("Type", 0)
		if err != nil {
			return err
		}
		err = writer.WriteElementString("Ln", "0.0.0.0.0.0")
		if err != nil {
			return err
		}
	}
	err = writer.WriteElementString("Index", charge.Commodity.Index)
	if err != nil {
		return err
	}
	writer.WriteStartElement("ChargeTables")
	if charge.ChargeTables != nil {
		for _, it := range charge.ChargeTables {
			writer.WriteStartElement("Item")
			err = writer.WriteElementString("Index", it.Index)
			if err != nil {
				return err
			}
			err = writer.WriteElementString("ChargePerUnit", it.ChargePerUnit)
			if err != nil {
				return err
			}
			writer.WriteEndElement()
		}
	}
	writer.WriteEndElement()
	writer.WriteEndElement()
	return err
}

// Save returns the save object content to XML.
//
// Parameters:
//
//	writer: XML writer.
func (g *GXDLMSCharge) Save(writer *GXXmlWriter) error {
	err := writer.WriteElementString("TotalAmountPaid", g.TotalAmountPaid)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("ChargeType", int(g.ChargeType))
	if err != nil {
		return err
	}
	err = writer.WriteElementString("Priority", g.Priority)
	if err != nil {
		return err
	}
	g.SaveUnitCharge(writer, "UnitChargeActive", &g.UnitChargeActive)
	g.SaveUnitCharge(writer, "UnitChargePassive", &g.UnitChargePassive)
	err = writer.WriteElementString("UnitChargeActivationTime", g.UnitChargeActivationTime)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("Period", g.Period)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("ChargeConfiguration", int(g.ChargeConfiguration))
	if err != nil {
		return err
	}
	err = writer.WriteElementString("LastCollectionTime", g.LastCollectionTime)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("LastCollectionAmount", g.LastCollectionAmount)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("TotalAmountRemaining", g.TotalAmountRemaining)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("Proportion", g.Proportion)
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
func (g *GXDLMSCharge) PostLoad(reader *GXXmlReader) error {
	return nil
}

// GetValues returns the an array containing the COSEM object's attribute values.
func (g *GXDLMSCharge) GetValues() []any {
	return []any{g.LogicalName(), g.TotalAmountPaid, g.ChargeType, g.Priority,
		g.UnitChargeActive, g.UnitChargePassive, g.UnitChargeActivationTime, g.Period,
		g.ChargeConfiguration, g.LastCollectionTime, g.LastCollectionAmount, g.TotalAmountRemaining, g.Proportion}
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
func (g *GXDLMSCharge) GetDataType(index int) (enums.DataType, error) {
	var ret enums.DataType
	var err error
	switch index {
	case 1:
		ret = enums.DataTypeOctetString
	case 2:
		ret = enums.DataTypeInt32
	case 3:
		ret = enums.DataTypeEnum
	case 4:
		ret = enums.DataTypeUint8
	case 5:
		ret = enums.DataTypeStructure
	case 6:
		ret = enums.DataTypeStructure
	case 7:
		ret = enums.DataTypeOctetString
	case 8:
		ret = enums.DataTypeUint32
	case 9:
		ret = enums.DataTypeBitString
	case 10:
		ret = enums.DataTypeDateTime
	case 11:
		ret = enums.DataTypeInt32
	case 12:
		ret = enums.DataTypeInt32
	case 13:
		ret = enums.DataTypeUint16
	default:
		err = errors.New("GetDataType failed. Invalid attribute index.")
	}
	return ret, err
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
func (g *GXDLMSCharge) GetUIDataType(index int) enums.DataType {
	if index == 7 || index == 10 {
		return enums.DataTypeDateTime
	}
	return g.Base().GetUIDataType(index)
}

// NewGXDLMSCharge creates a new Charge object instance.
//
// The function validates `ln` before creating the object.
// `ln` is the Logical Name and `sn` is the Short Name of the object.
func NewGXDLMSCharge(ln string, sn int16) (*GXDLMSCharge, error) {
	err := ValidateLogicalName(ln)
	if err != nil {
		return nil, err
	}
	return &GXDLMSCharge{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeCharge,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
