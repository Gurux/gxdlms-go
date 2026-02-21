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
	"time"

	"github.com/Gurux/gxdlms-go/dlmserrors"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// Online help:
// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSAccount
type GXDLMSAccount struct {
	GXDLMSObject
	// Payment mode.
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSAccount
	PaymentMode enums.PaymentMode

	// Account status.
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSAccount
	AccountStatus enums.AccountStatus

	// Index into the credit reference list indicating which Credit object is In use
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSAccount
	CurrentCreditInUse uint8

	// Current credit status.
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSAccount
	CurrentCreditStatus enums.AccountCreditStatus

	// The available_credit attribute is the sum of the positive current credit amount values in the instances of the Credit class.
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSAccount
	AvailableCredit int

	// Amount to clear.
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSAccount
	AmountToClear int

	// Conjunction with the amount to clear, and is included in the description of that attribute.
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSAccount
	ClearanceThreshold int

	// Simple sum of total_amount_remaining of all the Charge objects which are listed in the Account object.
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSAccount
	AggregatedDebt int

	// Credit references.
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSAccount
	CreditReferences []string

	// Charge references.
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSAccount
	ChargeReferences []string

	// Credit charge configurations
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSAccount
	CreditChargeConfigurations []GXCreditChargeConfiguration

	// Token gateway configurations.
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSAccount
	TokenGatewayConfigurations []GXTokenGatewayConfiguration

	// Account activation time.
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSAccount
	AccountActivationTime types.GXDateTime

	// Account closure time.
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSAccount
	AccountClosureTime types.GXDateTime

	// Currency settings.
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSAccount
	Currency GXCurrency

	// Low credit threshold.
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSAccount
	LowCreditThreshold int

	// Next credit available threshold.
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSAccount
	NextCreditAvailableThreshold int

	// Max provision.
	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSAccount
	MaxProvision uint16

	// Online help:
	// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSAccount
	MaxProvisionPeriod int
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSAccount) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

//Invoke returns the invokes method.
//
// Parameters:
//   settings: DLMS settings.
//   e: Invoke parameters.
func (g *GXDLMSAccount) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	switch e.Index {
	case 1:
		g.AccountStatus = enums.AccountStatusAccountActive
		g.AccountActivationTime = types.GXDateTime{Value: time.Now()}
	case 2:
		g.AccountStatus = enums.AccountStatusAccountClosed
		g.AccountClosureTime = types.GXDateTime{Value: time.Now()}
	case 3:
		//Meter must handle this.
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
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
func (g *GXDLMSAccount) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	// LN is static and read only once.
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	// PaymentMode, AccountStatus
	if all || g.CanRead(2) {
		attributes = append(attributes, 2)
	}
	// CurrentCreditInUse
	if all || g.CanRead(3) {
		attributes = append(attributes, 3)
	}
	// CurrentCreditStatus
	if all || g.CanRead(4) {
		attributes = append(attributes, 4)
	}
	// AvailableCredit
	if all || g.CanRead(5) {
		attributes = append(attributes, 5)
	}
	// AmountToClear
	if all || g.CanRead(6) {
		attributes = append(attributes, 6)
	}
	// ClearanceThreshold
	if all || g.CanRead(7) {
		attributes = append(attributes, 7)
	}
	// AggregatedDebt
	if all || g.CanRead(8) {
		attributes = append(attributes, 8)
	}
	// CreditReferences
	if all || g.CanRead(9) {
		attributes = append(attributes, 9)
	}
	// ChargeReferences
	if all || g.CanRead(10) {
		attributes = append(attributes, 10)
	}
	// CreditChargeConfigurations
	if all || g.CanRead(11) {
		attributes = append(attributes, 11)
	}
	// TokenGatewayConfigurations
	if all || g.CanRead(12) {
		attributes = append(attributes, 12)
	}
	// AccountActivationTime
	if all || g.CanRead(13) {
		attributes = append(attributes, 13)
	}
	// AccountClosureTime
	if all || g.CanRead(14) {
		attributes = append(attributes, 14)
	}
	// Currency
	if all || g.CanRead(15) {
		attributes = append(attributes, 15)
	}
	// LowCreditThreshold
	if all || g.CanRead(16) {
		attributes = append(attributes, 16)
	}
	// NextCreditAvailableThreshold
	if all || g.CanRead(17) {
		attributes = append(attributes, 17)
	}
	// MaxProvision
	if all || g.CanRead(18) {
		attributes = append(attributes, 18)
	}
	// MaxProvisionPeriod
	if all || g.CanRead(19) {
		attributes = append(attributes, 19)
	}
	return attributes
}

//GetNames returns the names of attribute indexes.
func (g *GXDLMSAccount) GetNames() []string {
	return []string{"Logical Name", "Payment mode",
		"Current credit in use", "Current credit status",
		"Available credit", "Amount to clear", "Clearance threshold",
		"Aggregated debt", "Credit references", "Charge references",
		"Credit charge configurations", "Token gateway configurations",
		"Account activation time", "Account closure time", "Currency",
		"Low credit threshold", "Next credit available threshold",
		"Max provision", "Max provision period"}
}

//GetMethodNames returns the names of method indexes.
func (g *GXDLMSAccount) GetMethodNames() []string {
	return []string{"Activate account", "Close account", "Reset account"}
}

//GetAttributeCount returns the amount of attributes.
//
// Returns:
//   Count of attributes.
func (g *GXDLMSAccount) GetAttributeCount() int {
	return 19
}

//GetMethodCount returns the amount of methods.
func (g *GXDLMSAccount) GetMethodCount() int {
	return 3
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
func (g *GXDLMSAccount) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	switch e.Index {
	case 1:
		return helpers.LogicalNameToBytes(g.LogicalName())
	case 2:
		data := types.NewGXByteBuffer()
		if err := data.SetUint8(uint8(enums.DataTypeStructure)); err != nil {
			return nil, err
		}
		if err := data.SetUint8(2); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, data, enums.DataTypeEnum, g.PaymentMode); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, data, enums.DataTypeEnum, g.AccountStatus); err != nil {
			return nil, err
		}
		return data.Array(), nil
	case 3:
		return g.CurrentCreditInUse, nil
	case 4:
		bs, err := types.NewGXBitStringFromInteger(int(g.CurrentCreditStatus), 8)
		if err != nil {
			return nil, err
		}
		return *bs, nil
	case 5:
		return g.AvailableCredit, nil
	case 6:
		return g.AmountToClear, nil
	case 7:
		return g.ClearanceThreshold, nil
	case 8:
		return g.AggregatedDebt, nil
	case 9:
		data := types.NewGXByteBuffer()
		if err := data.SetUint8(uint8(enums.DataTypeArray)); err != nil {
			return nil, err
		}
		if err := types.SetObjectCount(len(g.CreditReferences), data); err != nil {
			return nil, err
		}
		for _, it := range g.CreditReferences {
			ln, err := helpers.LogicalNameToBytes(it)
			if err != nil {
				return nil, err
			}
			if err = internal.SetData(settings, data, enums.DataTypeOctetString, ln); err != nil {
				return nil, err
			}
		}
		return data.Array(), nil
	case 10:
		data := types.NewGXByteBuffer()
		if err := data.SetUint8(uint8(enums.DataTypeArray)); err != nil {
			return nil, err
		}
		if err := types.SetObjectCount(len(g.ChargeReferences), data); err != nil {
			return nil, err
		}
		for _, it := range g.ChargeReferences {
			ln, err := helpers.LogicalNameToBytes(it)
			if err != nil {
				return nil, err
			}
			if err = internal.SetData(settings, data, enums.DataTypeOctetString, ln); err != nil {
				return nil, err
			}
		}
		return data.Array(), nil
	case 11:
		data := types.NewGXByteBuffer()
		if err := data.SetUint8(uint8(enums.DataTypeArray)); err != nil {
			return nil, err
		}
		if err := types.SetObjectCount(len(g.CreditChargeConfigurations), data); err != nil {
			return nil, err
		}
		for _, it := range g.CreditChargeConfigurations {
			if err := data.SetUint8(uint8(enums.DataTypeStructure)); err != nil {
				return nil, err
			}
			if err := data.SetUint8(3); err != nil {
				return nil, err
			}
			creditLN, err := helpers.LogicalNameToBytes(it.CreditReference)
			if err != nil {
				return nil, err
			}
			if err = internal.SetData(settings, data, enums.DataTypeOctetString, creditLN); err != nil {
				return nil, err
			}
			chargeLN, err := helpers.LogicalNameToBytes(it.ChargeReference)
			if err != nil {
				return nil, err
			}
			if err = internal.SetData(settings, data, enums.DataTypeOctetString, chargeLN); err != nil {
				return nil, err
			}
			bs, err := types.NewGXBitStringFromInteger(int(it.CollectionConfiguration), 3)
			if err != nil {
				return nil, err
			}
			if err = internal.SetData(settings, data, enums.DataTypeBitString, bs); err != nil {
				return nil, err
			}
		}
		return data.Array(), nil
	case 12:
		data := types.NewGXByteBuffer()
		if err := data.SetUint8(uint8(enums.DataTypeArray)); err != nil {
			return nil, err
		}
		if err := types.SetObjectCount(len(g.TokenGatewayConfigurations), data); err != nil {
			return nil, err
		}
		for _, it := range g.TokenGatewayConfigurations {
			if err := data.SetUint8(uint8(enums.DataTypeStructure)); err != nil {
				return nil, err
			}
			if err := data.SetUint8(2); err != nil {
				return nil, err
			}
			creditLN, err := helpers.LogicalNameToBytes(it.CreditReference)
			if err != nil {
				return nil, err
			}
			if err = internal.SetData(settings, data, enums.DataTypeOctetString, creditLN); err != nil {
				return nil, err
			}
			if err = internal.SetData(settings, data, enums.DataTypeUint8, it.TokenProportion); err != nil {
				return nil, err
			}
		}
		return data.Array(), nil
	case 13:
		return g.AccountActivationTime, nil
	case 14:
		return g.AccountClosureTime, nil
	case 15:
		data := types.NewGXByteBuffer()
		if err := data.SetUint8(uint8(enums.DataTypeStructure)); err != nil {
			return nil, err
		}
		if err := data.SetUint8(3); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, data, enums.DataTypeStringUTF8, g.Currency.Name); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, data, enums.DataTypeInt8, g.Currency.Scale); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, data, enums.DataTypeEnum, g.Currency.Unit); err != nil {
			return nil, err
		}
		return data.Array(), nil
	case 16:
		return g.LowCreditThreshold, nil
	case 17:
		return g.NextCreditAvailableThreshold, nil
	case 18:
		return g.MaxProvision, nil
	case 19:
		return g.MaxProvisionPeriod, nil
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
func (g *GXDLMSAccount) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	toSlice := func(value any) ([]any, bool) {
		switch v := value.(type) {
		case types.GXArray:
			return []any(v), true
		case types.GXStructure:
			return []any(v), true
		case []any:
			return v, true
		default:
			return nil, false
		}
	}
	toInt := func(value any) (int, error) {
		switch v := value.(type) {
		case int:
			return v, nil
		case int8:
			return int(v), nil
		case int16:
			return int(v), nil
		case int32:
			return int(v), nil
		case int64:
			return int(v), nil
		case uint8:
			return int(v), nil
		case uint16:
			return int(v), nil
		case uint32:
			return int(v), nil
		case uint64:
			return int(v), nil
		case types.GXEnum:
			return int(v.Value), nil
		case types.GXBitString:
			return v.ToInteger(), nil
		case *types.GXBitString:
			return v.ToInteger(), nil
		default:
			return 0, fmt.Errorf("unsupported numeric value type: %T", value)
		}
	}

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
			g.AccountStatus = enums.AccountStatusNewInactiveAccount
			g.PaymentMode = enums.PaymentModeCredit
			return nil
		}
		arr, ok := toSlice(e.Value)
		if !ok || len(arr) < 2 {
			e.Error = enums.ErrorCodeReadWriteDenied
			return fmt.Errorf("invalid payment mode structure: %T", e.Value)
		}
		paymentMode, err := toInt(arr[0])
		if err != nil {
			return err
		}
		accountStatus, err := toInt(arr[1])
		if err != nil {
			return err
		}
		g.PaymentMode = enums.PaymentMode(paymentMode)
		g.AccountStatus = enums.AccountStatus(accountStatus)
	case 3:
		value, err := toInt(e.Value)
		if err != nil {
			return err
		}
		g.CurrentCreditInUse = uint8(value)
	case 4:
		value, err := toInt(e.Value)
		if err != nil {
			return err
		}
		g.CurrentCreditStatus = enums.AccountCreditStatus(value)
	case 5:
		value, err := toInt(e.Value)
		if err != nil {
			return err
		}
		g.AvailableCredit = value
	case 6:
		value, err := toInt(e.Value)
		if err != nil {
			return err
		}
		g.AmountToClear = value
	case 7:
		value, err := toInt(e.Value)
		if err != nil {
			return err
		}
		g.ClearanceThreshold = value
	case 8:
		value, err := toInt(e.Value)
		if err != nil {
			return err
		}
		g.AggregatedDebt = value
	case 9:
		g.CreditReferences = make([]string, 0)
		if e.Value != nil {
			arr, ok := toSlice(e.Value)
			if !ok {
				return fmt.Errorf("invalid credit references type: %T", e.Value)
			}
			for _, it := range arr {
				ln, err := helpers.ToLogicalName(it)
				if err != nil {
					return err
				}
				g.CreditReferences = append(g.CreditReferences, ln)
			}
		}
	case 10:
		g.ChargeReferences = make([]string, 0)
		if e.Value != nil {
			arr, ok := toSlice(e.Value)
			if !ok {
				return fmt.Errorf("invalid charge references type: %T", e.Value)
			}
			for _, it := range arr {
				ln, err := helpers.ToLogicalName(it)
				if err != nil {
					return err
				}
				g.ChargeReferences = append(g.ChargeReferences, ln)
			}
		}
	case 11:
		g.CreditChargeConfigurations = make([]GXCreditChargeConfiguration, 0)
		if e.Value != nil {
			arr, ok := toSlice(e.Value)
			if !ok {
				return fmt.Errorf("invalid credit-charge configuration type: %T", e.Value)
			}
			for _, tmp := range arr {
				it, ok := toSlice(tmp)
				if !ok || len(it) < 3 {
					return fmt.Errorf("invalid credit-charge item type: %T", tmp)
				}
				creditLN, err := helpers.ToLogicalName(it[0])
				if err != nil {
					return err
				}
				chargeLN, err := helpers.ToLogicalName(it[1])
				if err != nil {
					return err
				}
				configuration, err := toInt(it[2])
				if err != nil {
					return err
				}
				g.CreditChargeConfigurations = append(g.CreditChargeConfigurations, GXCreditChargeConfiguration{
					CreditReference:         creditLN,
					ChargeReference:         chargeLN,
					CollectionConfiguration: enums.CreditCollectionConfiguration(configuration),
				})
			}
		}
	case 12:
		g.TokenGatewayConfigurations = make([]GXTokenGatewayConfiguration, 0)
		if e.Value != nil {
			arr, ok := toSlice(e.Value)
			if !ok {
				return fmt.Errorf("invalid token gateway configuration type: %T", e.Value)
			}
			for _, tmp := range arr {
				it, ok := toSlice(tmp)
				if !ok || len(it) < 2 {
					return fmt.Errorf("invalid token gateway item type: %T", tmp)
				}
				creditLN, err := helpers.ToLogicalName(it[0])
				if err != nil {
					return err
				}
				tokenProportion, err := toInt(it[1])
				if err != nil {
					return err
				}
				g.TokenGatewayConfigurations = append(g.TokenGatewayConfigurations, GXTokenGatewayConfiguration{
					CreditReference: creditLN,
					TokenProportion: uint8(tokenProportion),
				})
			}
		}
	case 13:
		if e.Value == nil {
			g.AccountActivationTime = types.GXDateTime{}
			return nil
		}
		switch v := e.Value.(type) {
		case []byte:
			value, err := internal.ChangeTypeFromByteArray(settings, v, enums.DataTypeDateTime)
			if err != nil {
				return err
			}
			g.AccountActivationTime = value.(types.GXDateTime)
		case string:
			value, err := types.NewGXDateTimeFromString(v, nil)
			if err != nil {
				return err
			}
			g.AccountActivationTime = *value
		case types.GXDateTime:
			g.AccountActivationTime = v
		case *types.GXDateTime:
			g.AccountActivationTime = *v
		default:
			return fmt.Errorf("invalid account activation time type: %T", e.Value)
		}
	case 14:
		if e.Value == nil {
			g.AccountClosureTime = types.GXDateTime{}
			return nil
		}
		switch v := e.Value.(type) {
		case []byte:
			value, err := internal.ChangeTypeFromByteArray(settings, v, enums.DataTypeDateTime)
			if err != nil {
				return err
			}
			g.AccountClosureTime = value.(types.GXDateTime)
		case string:
			value, err := types.NewGXDateTimeFromString(v, nil)
			if err != nil {
				return err
			}
			g.AccountClosureTime = *value
		case types.GXDateTime:
			g.AccountClosureTime = v
		case *types.GXDateTime:
			g.AccountClosureTime = *v
		default:
			return fmt.Errorf("invalid account closure time type: %T", e.Value)
		}
	case 15:
		if e.Value == nil {
			g.Currency = GXCurrency{}
			return nil
		}
		tmp, ok := toSlice(e.Value)
		if !ok || len(tmp) < 3 {
			return fmt.Errorf("invalid currency value type: %T", e.Value)
		}
		switch v := tmp[0].(type) {
		case []byte:
			g.Currency.Name = string(v)
		case string:
			g.Currency.Name = v
		default:
			return fmt.Errorf("invalid currency name type: %T", tmp[0])
		}
		scale, err := toInt(tmp[1])
		if err != nil {
			return err
		}
		unit, err := toInt(tmp[2])
		if err != nil {
			return err
		}
		g.Currency.Scale = int8(scale)
		g.Currency.Unit = enums.Currency(unit)
	case 16:
		value, err := toInt(e.Value)
		if err != nil {
			return err
		}
		g.LowCreditThreshold = value
	case 17:
		value, err := toInt(e.Value)
		if err != nil {
			return err
		}
		g.NextCreditAvailableThreshold = value
	case 18:
		value, err := toInt(e.Value)
		if err != nil {
			return err
		}
		g.MaxProvision = uint16(value)
	case 19:
		value, err := toInt(e.Value)
		if err != nil {
			return err
		}
		g.MaxProvisionPeriod = value
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil
}

func (g *GXDLMSAccount) LoadReferences(reader *GXXmlReader, name string, list []string) error {
	list = list[:0]
	var err error
	if ret, err := reader.IsStartElementNamed(name, true); ret && err == nil {
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
			list = append(list, ret)
		}
		reader.ReadEndElement(name)
	}
	return err
}

func (g *GXDLMSAccount) LoadCreditChargeConfigurations(reader *GXXmlReader, list []GXCreditChargeConfiguration) error {
	var err error
	list = list[:0]
	if ret, err := reader.IsStartElementNamed("CreditChargeConfigurations", true); ret && err == nil {
		for {
			ret, err = reader.IsStartElementNamed("Item", true)
			if err != nil {
				return err
			}
			if !ret {
				break
			}
			it := GXCreditChargeConfiguration{}
			it.CreditReference, err = reader.ReadElementContentAsString("Credit", "")
			if err != nil {
				return err
			}
			it.ChargeReference, err = reader.ReadElementContentAsString("Charge", "")
			if err != nil {
				return err
			}
			ret, err := reader.ReadElementContentAsInt("Configuration", 0)
			if err != nil {
				return err
			}
			it.CollectionConfiguration = enums.CreditCollectionConfiguration(ret)
			list = append(list, it)
		}
		reader.ReadEndElement("CreditChargeConfigurations")
	}
	return err
}

func (g *GXDLMSAccount) LoadTokenGatewayConfigurations(reader *GXXmlReader, list []GXTokenGatewayConfiguration) error {
	var err error
	list = list[:0]
	if ret, err := reader.IsStartElementNamed("TokenGatewayConfigurations", true); ret && err == nil {
		for {
			ret, err = reader.IsStartElementNamed("Item", true)
			if err != nil {
				return err
			}
			if !ret {
				break
			}
			it := GXTokenGatewayConfiguration{}
			it.CreditReference, err = reader.ReadElementContentAsString("Credit", "")
			if err != nil {
				return err
			}
			it.TokenProportion, err = reader.ReadElementContentAsUInt8("Token", 0)
			if err != nil {
				return err
			}
			list = append(list, it)
		}
		reader.ReadEndElement("TokenGatewayConfigurations")
	}
	return err
}

//Load returns the load object content from XML.
//
// Parameters:
//   reader: XML reader.
func (g *GXDLMSAccount) Load(reader *GXXmlReader) error {
	ret, err := reader.ReadElementContentAsInt("PaymentMode", 0)
	if err != nil {
		return err
	}
	g.PaymentMode = enums.PaymentMode(ret)
	ret, err = reader.ReadElementContentAsInt("AccountStatus", 0)
	if err != nil {
		return err
	}
	g.AccountStatus = enums.AccountStatus(ret)
	g.CurrentCreditInUse, err = reader.ReadElementContentAsUInt8("CurrentCreditInUse", 0)
	if err != nil {
		return err
	}
	ret, err = reader.ReadElementContentAsInt("CurrentCreditStatus", 0)
	if err != nil {
		return err
	}
	g.CurrentCreditStatus = enums.AccountCreditStatus(ret)
	g.AvailableCredit, err = reader.ReadElementContentAsInt("AvailableCredit", 0)
	if err != nil {
		return err
	}
	g.AmountToClear, err = reader.ReadElementContentAsInt("AmountToClear", 0)
	if err != nil {
		return err
	}
	g.ClearanceThreshold, err = reader.ReadElementContentAsInt("ClearanceThreshold", 0)
	if err != nil {
		return err
	}
	g.AggregatedDebt, err = reader.ReadElementContentAsInt("AggregatedDebt", 0)
	if err != nil {
		return err
	}
	err = g.LoadReferences(reader, "CreditReferences", g.CreditReferences)
	if err != nil {
		return err
	}
	err = g.LoadReferences(reader, "ChargeReferences", g.ChargeReferences)
	if err != nil {
		return err
	}
	err = g.LoadCreditChargeConfigurations(reader, g.CreditChargeConfigurations)
	if err != nil {
		return err
	}
	err = g.LoadTokenGatewayConfigurations(reader, g.TokenGatewayConfigurations)
	if err != nil {
		return err
	}
	g.AccountActivationTime, err = reader.ReadElementContentAsDateTime("AccountActivationTime", nil)
	if err != nil {
		return err
	}
	g.AccountClosureTime, err = reader.ReadElementContentAsDateTime("AccountClosureTime", nil)
	if err != nil {
		return err
	}
	g.Currency.Name, err = reader.ReadElementContentAsString("CurrencyName", "")
	if err != nil {
		return err
	}
	g.Currency.Scale, err = reader.ReadElementContentAsInt8("CurrencyScale", 0)
	if err != nil {
		return err
	}
	ret, err = reader.ReadElementContentAsInt("CurrencyUnit", 0)
	if err != nil {
		return err
	}
	g.Currency.Unit = enums.Currency(ret)
	g.LowCreditThreshold, err = reader.ReadElementContentAsInt("LowCreditThreshold", 0)
	if err != nil {
		return err
	}
	g.NextCreditAvailableThreshold, err = reader.ReadElementContentAsInt("NextCreditAvailableThreshold", 0)
	if err != nil {
		return err
	}
	g.MaxProvision, err = reader.ReadElementContentAsUInt16("MaxProvision", 0)
	if err != nil {
		return err
	}
	g.MaxProvisionPeriod, err = reader.ReadElementContentAsInt("MaxProvisionPeriod", 0)
	return err
}

func (g *GXDLMSAccount) SaveReferences(writer *GXXmlWriter, list []string, name string) error {
	var err error
	writer.WriteStartElement(name)
	for _, it := range list {
		writer.WriteStartElement("Item")
		err = writer.WriteElementString("Name", it)
		if err != nil {
			return err
		}
		writer.WriteEndElement()
	}
	writer.WriteEndElement()
	return err
}

func (g *GXDLMSAccount) SaveCreditChargeConfigurations(writer *GXXmlWriter, list []GXCreditChargeConfiguration) error {
	var err error
	writer.WriteStartElement("CreditChargeConfigurations")
	for _, it := range list {
		writer.WriteStartElement("Item")
		err = writer.WriteElementString("Credit", it.CreditReference)
		if err != nil {
			return err
		}
		err = writer.WriteElementString("Charge", it.ChargeReference)
		if err != nil {
			return err
		}
		err = writer.WriteElementString("Configuration", int(it.CollectionConfiguration))
		if err != nil {
			return err
		}
		writer.WriteEndElement()
	}
	writer.WriteEndElement()
	return err
}

func (g *GXDLMSAccount) SaveTokenGatewayConfigurations(writer *GXXmlWriter, list []GXTokenGatewayConfiguration) error {
	var err error
	if list != nil {
		writer.WriteStartElement("TokenGatewayConfigurations")
		for _, it := range list {
			writer.WriteStartElement("Item")
			err = writer.WriteElementString("Credit", it.CreditReference)
			if err != nil {
				return err
			}
			err = writer.WriteElementString("Token", it.TokenProportion)
			if err != nil {
				return err
			}
			writer.WriteEndElement()
		}
		writer.WriteEndElement()
	}
	return err
}

//Save returns the save object content to XML.
//
// Parameters:
//   writer: XML writer.
func (g *GXDLMSAccount) Save(writer *GXXmlWriter) error {
	err := writer.WriteElementString("PaymentMode", int(g.PaymentMode))
	if err != nil {
		return err
	}
	err = writer.WriteElementString("AccountStatus", int(g.AccountStatus))
	if err != nil {
		return err
	}
	err = writer.WriteElementString("CurrentCreditInUse", g.CurrentCreditInUse)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("CurrentCreditStatus", uint8(g.CurrentCreditStatus))
	if err != nil {
		return err
	}
	err = writer.WriteElementString("AvailableCredit", g.AvailableCredit)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("AmountToClear", g.AmountToClear)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("ClearanceThreshold", g.ClearanceThreshold)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("AggregatedDebt", g.AggregatedDebt)
	if err != nil {
		return err
	}
	g.SaveReferences(writer, g.CreditReferences, "CreditReferences")
	g.SaveReferences(writer, g.ChargeReferences, "ChargeReferences")
	g.SaveCreditChargeConfigurations(writer, g.CreditChargeConfigurations)
	g.SaveTokenGatewayConfigurations(writer, g.TokenGatewayConfigurations)
	err = writer.WriteElementString("AccountActivationTime", g.AccountActivationTime)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("AccountClosureTime", g.AccountClosureTime)
	if err != nil {
		return err
	}
	writer.WriteElementString("CurrencyName", g.Currency.Name)
	writer.WriteElementString("CurrencyScale", g.Currency.Scale)
	writer.WriteElementString("CurrencyUnit", int(g.Currency.Unit))
	err = writer.WriteElementString("LowCreditThreshold", g.LowCreditThreshold)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("NextCreditAvailableThreshold", g.NextCreditAvailableThreshold)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("MaxProvision", g.MaxProvision)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("MaxProvisionPeriod", g.MaxProvisionPeriod)
	if err != nil {
		return err
	}
	return err
}

//PostLoad returns the handle actions after Load.
//
// Parameters:
//   reader: XML reader.
func (g *GXDLMSAccount) PostLoad(reader *GXXmlReader) error {
	return nil
}

//GetValues returns the an array containing the COSEM object's attribute values.
func (g *GXDLMSAccount) GetValues() []any {
	return []any{g.LogicalName(), []any{g.PaymentMode, g.AccountStatus}, g.CurrentCreditInUse, g.CurrentCreditStatus,
		g.AvailableCredit, g.AmountToClear, g.ClearanceThreshold, g.AggregatedDebt,
		g.CreditReferences, g.ChargeReferences, g.CreditChargeConfigurations,
		g.TokenGatewayConfigurations, g.AccountActivationTime,
		g.AccountClosureTime, g.Currency, g.LowCreditThreshold, g.NextCreditAvailableThreshold,
		g.MaxProvision, g.MaxProvisionPeriod}
}

//Activate returns the activate account.
//
// Parameters:
//   client: DLMS client.
//
// Returns:
//   Action bytes.
func (g *GXDLMSAccount) Activate(client IGXDLMSClient) ([][]uint8, error) {
	return client.Method(g, 1, int8(0), enums.DataTypeInt8)
}

//Close returns the close account.
//
// Parameters:
//   client: DLMS client.
//
// Returns:
//   Action bytes.
func (g *GXDLMSAccount) Close(client IGXDLMSClient) ([][]uint8, error) {
	return client.Method(g, 2, int8(0), enums.DataTypeInt8)
}

//Reset returns the reset account.
//
// Parameters:
//   client: DLMS client.
//
// Returns:
//   Action bytes.
func (g *GXDLMSAccount) Reset(client IGXDLMSClient) ([][]uint8, error) {
	return client.Method(g, 3, int8(0), enums.DataTypeInt8)
}

//GetUIDataType returns the UI data type of selected index.
//
// Parameters:
//   index: Attribute index of the object.
//
// Returns:
//   UI data type of the object.
func (g *GXDLMSAccount) GetUIDataType(index int) enums.DataType {
	// AccountActivationTime or AccountClosureTime
	if index == 13 || index == 14 {
		return enums.DataTypeDateTime
	}
	return g.Base().GetUIDataType(index)
}

//GetDataType returns the device data type of selected attribute index.
//
// Parameters:
//   index: Attribute index of the object.
//
// Returns:
//   Device data type of the object.
func (g *GXDLMSAccount) GetDataType(index int) (enums.DataType, error) {
	switch index {
	case 1:
		return enums.DataTypeOctetString, nil
	case 2:
		return enums.DataTypeStructure, nil
	case 3:
		return enums.DataTypeUint8, nil
	case 4:
		return enums.DataTypeBitString, nil
	case 5:
		return enums.DataTypeInt32, nil
	case 6:
		return enums.DataTypeInt32, nil
	case 7:
		return enums.DataTypeInt32, nil
	case 8:
		return enums.DataTypeInt32, nil
	case 9:
		return enums.DataTypeArray, nil
	case 10:
		return enums.DataTypeArray, nil
	case 11:
		return enums.DataTypeArray, nil
	case 12:
		return enums.DataTypeArray, nil
	case 13:
		return enums.DataTypeOctetString, nil
	case 14:
		return enums.DataTypeOctetString, nil
	case 15:
		return enums.DataTypeStructure, nil
	case 16:
		return enums.DataTypeInt32, nil
	case 17:
		return enums.DataTypeInt32, nil
	case 18:
		return enums.DataTypeUint16, nil
	case 19:
		return enums.DataTypeInt32, nil
	default:
		return 0, dlmserrors.ErrInvalidAttributeIndex
	}
}

// NewGXDLMSAccount creates a new Account object instance.
//
// The function validates `ln` before creating the object.
//`ln` is the Logical Name and `sn` is the Short Name of the object.
func NewGXDLMSAccount(ln string, sn int16) (*GXDLMSAccount, error) {
	err := ValidateLogicalName(ln)
	if err != nil {
		return nil, err
	}
	return &GXDLMSAccount{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeAccount,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
