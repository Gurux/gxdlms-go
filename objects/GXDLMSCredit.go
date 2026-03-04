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
// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSCredit
type GXDLMSCredit struct {
	GXDLMSObject
	// Current credit amount.
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSCredit
	CurrentCreditAmount int32

	// Type.
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSCredit
	Type enums.CreditType

	// Priority.
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSCredit
	Priority uint8

	// Warning threshold.
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSCredit
	WarningThreshold int32

	// Limit.
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSCredit
	Limit int32

	// Credit configuration.
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSCredit
	CreditConfiguration enums.CreditConfiguration

	// Status.
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSCredit
	Status enums.CreditStatus

	// Preset credit amount.
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSCredit
	PresetCreditAmount int32

	// Credit available threshold.
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSCredit
	CreditAvailableThreshold int32

	// Period.
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSCredit
	Period types.GXDateTime
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSCredit) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

// Invoke returns the invokes method.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Invoke parameters.
func (g *GXDLMSCredit) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	switch e.Index {
	case 1:
		g.CurrentCreditAmount = g.CurrentCreditAmount + e.Value.(int32)
	case 2:
		g.CurrentCreditAmount = e.Value.(int32)
	case 3:
		if (g.CreditConfiguration&enums.CreditConfigurationConfirmation) != 0 && g.Status == enums.CreditStatusSelectable {
			g.Status = enums.CreditStatusInvoked
		}
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
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
func (g *GXDLMSCredit) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	// LN is static and read only once.
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	// CurrentCreditAmount
	if all || g.CanRead(2) {
		attributes = append(attributes, 2)
	}
	// Type
	if all || g.CanRead(3) {
		attributes = append(attributes, 3)
	}
	// Priority
	if all || g.CanRead(4) {
		attributes = append(attributes, 4)
	}
	// WarningThreshold
	if all || g.CanRead(5) {
		attributes = append(attributes, 5)
	}
	// Limit
	if all || g.CanRead(6) {
		attributes = append(attributes, 6)
	}
	// creditConfiguration
	if all || g.CanRead(7) {
		attributes = append(attributes, 7)
	}
	// Status
	if all || g.CanRead(8) {
		attributes = append(attributes, 8)
	}
	// PresetCreditAmount
	if all || g.CanRead(9) {
		attributes = append(attributes, 9)
	}
	// CreditAvailableThreshold
	if all || g.CanRead(10) {
		attributes = append(attributes, 10)
	}
	// Period
	if all || g.CanRead(11) {
		attributes = append(attributes, 11)
	}
	return attributes
}

// GetNames returns the names of attribute indexes.
func (g *GXDLMSCredit) GetNames() []string {
	return []string{"Logical Name", "CurrentCreditAmount", "Type", "Priority", "WarningThreshold", "Limit", "CreditConfiguration", "Status", "PresetCreditAmount", "CreditAvailableThreshold", "Period"}
}

// GetMethodNames returns the names of method indexes.
func (g *GXDLMSCredit) GetMethodNames() []string {
	return []string{"Update amount", "Set amount to value", "Invoke credit"}
}

// GetAttributeCount returns the amount of attributes.
//
// Returns:
//
//	Count of attributes.
func (g *GXDLMSCredit) GetAttributeCount() int {
	return 11
}

// GetMethodCount returns the amount of methods.
func (g *GXDLMSCredit) GetMethodCount() int {
	return 3
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
func (g *GXDLMSCredit) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	var err error
	var ret any
	switch e.Index {
	case 1:
		return helpers.LogicalNameToBytes(g.LogicalName())
	case 2:
		ret = g.CurrentCreditAmount
	case 3:
		ret = g.Type
	case 4:
		ret = g.Priority
	case 5:
		ret = g.WarningThreshold
	case 6:
		ret = g.Limit
	case 7:
		ret, err = types.NewGXBitStringFromInteger(int(g.CreditConfiguration), 5)
	case 8:
		ret = g.Status
	case 9:
		ret = g.PresetCreditAmount
	case 10:
		ret = g.CreditAvailableThreshold
	case 11:
		ret = g.Period
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
func (g *GXDLMSCredit) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	var err error
	switch e.Index {
	case 1:
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return g.SetLogicalName(ln)
	case 2:
		g.CurrentCreditAmount = e.Value.(int32)
	case 3:
		g.Type = enums.CreditType(e.Value.(types.GXEnum).Value)
	case 4:
		g.Priority = e.Value.(uint8)
	case 5:
		g.WarningThreshold = e.Value.(int32)
	case 6:
		g.Limit = e.Value.(int32)
	case 7:
		bs := e.Value.(types.GXBitString)
		g.CreditConfiguration = enums.CreditConfiguration(bs.ToInteger())
	case 8:
		g.Status = enums.CreditStatus(e.Value.(types.GXEnum).Value)
	case 9:
		g.PresetCreditAmount = e.Value.(int32)
	case 10:
		g.CreditAvailableThreshold = e.Value.(int32)
	case 11:
		if e.Value == nil {
			g.Period = types.GXDateTime{}
		} else {
			if v, ok := e.Value.([]byte); ok {
				ret, err := internal.ChangeTypeFromByteArray(settings, v, enums.DataTypeDateTime)
				if err != nil {
					return err
				}
				e.Value = ret.(types.GXDateTime)
			} else if v, ok := e.Value.(string); ok {
				e.Value, err = types.NewGXDateTimeFromString(v, nil)
			}
			if v, ok := e.Value.(types.GXDateTime); ok {
				g.Period = v
			}
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
func (g *GXDLMSCredit) Load(reader *GXXmlReader) error {
	var err error
	g.CurrentCreditAmount, err = reader.ReadElementContentAsInt32("CurrentCreditAmount", 0)
	if err != nil {
		return err
	}
	ret, err := reader.ReadElementContentAsInt("Type", 0)
	if err != nil {
		return err
	}
	g.Type = enums.CreditType(ret)
	g.Priority, err = reader.ReadElementContentAsUInt8("Priority", 0)
	if err != nil {
		return err
	}
	g.WarningThreshold, err = reader.ReadElementContentAsInt32("WarningThreshold", 0)
	if err != nil {
		return err
	}
	g.Limit, err = reader.ReadElementContentAsInt32("Limit", 0)
	if err != nil {
		return err
	}
	ret, err = reader.ReadElementContentAsInt("CreditConfiguration", 0)
	if err != nil {
		return err
	}
	g.CreditConfiguration = enums.CreditConfiguration(ret)
	ret, err = reader.ReadElementContentAsInt("Status", 0)
	if err != nil {
		return err
	}
	g.Status = enums.CreditStatus(ret)
	g.PresetCreditAmount, err = reader.ReadElementContentAsInt32("PresetCreditAmount", 0)
	if err != nil {
		return err
	}
	g.CreditAvailableThreshold, err = reader.ReadElementContentAsInt32("CreditAvailableThreshold", 0)
	if err != nil {
		return err
	}
	g.Period, err = reader.ReadElementContentAsGXDateTime("Period")
	return err
}

// Save returns the save object content to XML.
//
// Parameters:
//
//	writer: XML writer.
func (g *GXDLMSCredit) Save(writer *GXXmlWriter) error {
	err := writer.WriteElementString("CurrentCreditAmount", g.CurrentCreditAmount)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("Type", uint8(g.Type))
	if err != nil {
		return err
	}
	err = writer.WriteElementString("Priority", g.Priority)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("WarningThreshold", g.WarningThreshold)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("Limit", g.Limit)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("CreditConfiguration", uint8(g.CreditConfiguration))
	if err != nil {
		return err
	}
	err = writer.WriteElementString("Status", uint8(g.Status))
	if err != nil {
		return err
	}
	err = writer.WriteElementString("PresetCreditAmount", g.PresetCreditAmount)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("CreditAvailableThreshold", g.CreditAvailableThreshold)
	if err != nil {
		return err
	}
	return writer.WriteElementString("Period", g.Period)
}

// PostLoad returns the handle actions after Load.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSCredit) PostLoad(reader *GXXmlReader) error {
	return nil
}

// GetValues returns the an array containing the COSEM object's attribute values.
func (g *GXDLMSCredit) GetValues() []any {
	return []any{g.LogicalName(), g.CurrentCreditAmount, g.Type, g.Priority,
		g.WarningThreshold, g.Limit, g.CreditConfiguration, g.Status,
		g.PresetCreditAmount, g.CreditAvailableThreshold, g.Period}
}

// UpdateAmount returns the adjusts the value of the current credit amount attribute.
//
// Parameters:
//
//	client: DLMS client.
//	value: Current credit amount
//
// Returns:
//
//	Action bytes.
func (g *GXDLMSCredit) UpdateAmount(client IGXDLMSClient, value int32) ([][]uint8, error) {
	return client.Method(g, 1, value, enums.DataTypeInt32)
}

// SetAmountToValue returns the sets the value of the current credit amount attribute.
//
// Parameters:
//
//	client: DLMS client.
//	value: Current credit amount
//
// Returns:
//
//	Action bytes.
func (g *GXDLMSCredit) SetAmountToValue(client IGXDLMSClient, value int32) ([][]uint8, error) {
	return client.Method(g, 2, value, enums.DataTypeInt32)
}

// InvokeCredit returns the sets the value of the current credit amount attribute.
//
// Parameters:
//
//	client: DLMS client.
//	value: Current credit amount
//
// Returns:
//
//	Action bytes.
func (g *GXDLMSCredit) InvokeCredit(client IGXDLMSClient, value enums.CreditStatus) ([][]uint8, error) {
	return client.Method(g, 3, uint8(value), enums.DataTypeUint8)
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
func (g *GXDLMSCredit) GetUIDataType(index int) enums.DataType {
	// Period
	if index == 11 {
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
func (g *GXDLMSCredit) GetDataType(index int) (enums.DataType, error) {
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
		ret = enums.DataTypeInt32
	case 6:
		ret = enums.DataTypeInt32
	case 7:
		ret = enums.DataTypeBitString
	case 8:
		ret = enums.DataTypeEnum
	case 9:
		ret = enums.DataTypeInt32
	case 10:
		ret = enums.DataTypeInt32
	case 11:
		ret = enums.DataTypeOctetString
	default:
		err = errors.New("GetDataType failed. Invalid attribute index.")
	}
	return ret, err
}

// NewGXDLMSGXDLMSCredit creates a new Credit object instance.
//
// The function validates `ln` before creating the object.
// `ln` is the Logical Name and `sn` is the Short Name of the object.
func NewGXDLMSCredit(ln string, sn int16) (*GXDLMSCredit, error) {
	err := ValidateLogicalName(ln)
	if err != nil {
		return nil, err
	}
	return &GXDLMSCredit{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeCredit,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
