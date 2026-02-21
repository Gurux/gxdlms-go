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

	"github.com/Gurux/gxdlms-go/dlmserrors"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// Online help:
// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSPushSetup
type GXDLMSPushSetup struct {
	GXDLMSObject
	Service enums.ServiceType

	Destination string

	Message enums.MessageType

	// Defines the list of attributes or objects to be pushed.
	// Upon a call of the push (data) method the selected attributes are sent to the destination
	// defined in send_destination_and_method.
	PushObjectList []types.GXKeyValuePair[IGXDLMSBase, GXDLMSCaptureObject]

	// Contains the start and end date/time
	// stamp when the communication window(s) for the push become active
	// (for the start instant), or inac-tive (for the end instant).
	CommunicationWindow []types.GXKeyValuePair[types.GXDateTime, types.GXDateTime]

	// To avoid simultaneous network connections of a lot of devices at ex-actly
	// the same point in time, a randomisation interval in seconds can be defined.
	// This means that the push operation is not started imme-diately at the
	// beginning of the first communication window but started randomly delayed.
	RandomisationStartInterval uint16

	// The maximum number of retrials in case of unsuccessful push at-tempts. After a successful push no further push attempts are made until the push setup is triggered again.
	// A value of 0 means no repetitions, i.e. only the initial connection at-tempt is made.
	NumberOfRetries uint8

	// Repetition delay.
	// Version 2 is using RepetitionDelay2.
	RepetitionDelay uint16

	// Repetition delay for Version2.
	// Version 0 and 1 use RepetitionDelay.
	RepetitionDelay2 GXRepetitionDelay

	// The logical name of a communication port setup object.
	PortReference IGXDLMSBase

	// Push client SAP.
	PushClientSAP int8

	// Push protection parameters.
	PushProtectionParameters []GXPushProtectionParameters

	// Push operation method.
	PushOperationMethod enums.PushOperationMethod

	// Push confirmation parameter.
	ConfirmationParameters GXPushConfirmationParameter

	// Last confirmation date time.
	LastConfirmationDateTime types.GXDateTime
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSPushSetup) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

// Invoke returns the invokes method.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Invoke parameters.
func (g *GXDLMSPushSetup) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	if e.Index == 1 {
		// Only TCP/IP push is allowed at the moment.
		if g.Service != enums.ServiceTypeTCP || g.Message != enums.MessageTypeCosemApdu || len(g.PushObjectList) == 0 {
			e.Error = enums.ErrorCodeHardwareFault
			return nil, nil
		}
		return nil, nil
	}
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
func (g *GXDLMSPushSetup) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	// LN is static and read only once.
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	// PushObjectList
	if all || g.CanRead(2) {
		attributes = append(attributes, 2)
	}
	// SendDestinationAndMethod
	if all || g.CanRead(3) {
		attributes = append(attributes, 3)
	}
	// CommunicationWindow
	if all || g.CanRead(4) {
		attributes = append(attributes, 4)
	}
	// RandomisationStartInterval
	if all || g.CanRead(5) {
		attributes = append(attributes, 5)
	}
	// NumberOfRetries
	if all || g.CanRead(6) {
		attributes = append(attributes, 6)
	}
	// RepetitionDelay
	if all || g.CanRead(7) {
		attributes = append(attributes, 7)
	}
	if g.Version > 0 {
		// PortReference
		if all || g.CanRead(8) {
			attributes = append(attributes, 8)
		}
		// PushClientSAP
		if all || g.CanRead(9) {
			attributes = append(attributes, 9)
		}
		// PushProtectionParameters
		if all || g.CanRead(10) {
			attributes = append(attributes, 10)
		}
		if g.Version > 1 {
			// PushOperationMethod
			if all || g.CanRead(11) {
				attributes = append(attributes, 11)
			}
			// ConfirmationParameters
			if all || g.CanRead(12) {
				attributes = append(attributes, 12)
			}
			// LastConfirmationDateTime
			if all || g.CanRead(13) {
				attributes = append(attributes, 13)
			}
		}
	}
	return attributes
}

// GetNames returns the names of attribute indexes.
func (g *GXDLMSPushSetup) GetNames() []string {
	return []string{"Logical Name", "Object List", "Send Destination And Method", "Communication Window", "Randomisation Start Interval", "Number Of Retries", "Repetition Delay", "Port reference", "Push client SAP", "Push protection parameters", "Push operation method", "Confirmation parameters", "Last confirmation date time"}
}

// GetMethodNames returns the names of method indexes.
func (g *GXDLMSPushSetup) GetMethodNames() []string {
	if g.Version < 2 {
		return []string{"Push"}
	}
	return []string{"Push", "Reset"}
}

// GetAttributeCount returns the amount of attributes.
//
// Returns:
//
//	Count of attributes.
func (g *GXDLMSPushSetup) GetAttributeCount() int {
	if g.Version == 0 {
		return 7
	}
	if g.Version == 1 {
		return 10
	}
	return 13
}

// GetMethodCount returns the amount of methods.
func (g *GXDLMSPushSetup) GetMethodCount() int {
	if g.Version < 2 {
		return 1
	}
	return 2
}

func (g *GXDLMSPushSetup) GetPushObjectList(settings *settings.GXDLMSSettings) (any, error) {
	buff := types.NewGXByteBuffer()
	if err := buff.SetUint8(uint8(enums.DataTypeArray)); err != nil {
		return nil, err
	}
	if err := types.SetObjectCount(len(g.PushObjectList), buff); err != nil {
		return nil, err
	}
	for _, it := range g.PushObjectList {
		if err := buff.SetUint8(uint8(enums.DataTypeStructure)); err != nil {
			return nil, err
		}
		if g.Version < 1 {
			if err := buff.SetUint8(4); err != nil {
				return nil, err
			}
		} else {
			cnt := uint8(5)
			if g.Version > 1 {
				cnt = 6
			}
			if err := buff.SetUint8(cnt); err != nil {
				return nil, err
			}
		}
		if err := internal.SetData(settings, buff, enums.DataTypeUint16, int(it.Key.Base().ObjectType())); err != nil {
			return nil, err
		}
		ln, err := helpers.LogicalNameToBytes(it.Key.Base().LogicalName())
		if err != nil {
			return nil, err
		}
		if err = internal.SetData(settings, buff, enums.DataTypeOctetString, ln); err != nil {
			return nil, err
		}
		if err = internal.SetData(settings, buff, enums.DataTypeInt8, it.Value.AttributeIndex); err != nil {
			return nil, err
		}
		if err = internal.SetData(settings, buff, enums.DataTypeUint16, it.Value.DataIndex); err != nil {
			return nil, err
		}
		if g.Version > 0 {
			if err = buff.SetUint8(uint8(enums.DataTypeStructure)); err != nil {
				return nil, err
			}
			if err = buff.SetUint8(2); err != nil {
				return nil, err
			}
			if err = internal.SetData(settings, buff, enums.DataTypeEnum, it.Value.Restriction.Type); err != nil {
				return nil, err
			}
			switch it.Value.Restriction.Type {
			case enums.RestrictionTypeNone:
				if err = internal.SetData(settings, buff, enums.DataTypeNone, nil); err != nil {
					return nil, err
				}
			case enums.RestrictionTypeDate:
				if err = buff.SetUint8(uint8(enums.DataTypeStructure)); err != nil {
					return nil, err
				}
				if err = buff.SetUint8(2); err != nil {
					return nil, err
				}
				if err = internal.SetData(settings, buff, enums.DataTypeOctetString, it.Value.Restriction.From); err != nil {
					return nil, err
				}
				if err = internal.SetData(settings, buff, enums.DataTypeOctetString, it.Value.Restriction.To); err != nil {
					return nil, err
				}
			case enums.RestrictionTypeEntry:
				if err = buff.SetUint8(uint8(enums.DataTypeStructure)); err != nil {
					return nil, err
				}
				if err = buff.SetUint8(2); err != nil {
					return nil, err
				}
				if err = internal.SetData(settings, buff, enums.DataTypeUint16, it.Value.Restriction.From); err != nil {
					return nil, err
				}
				if err = internal.SetData(settings, buff, enums.DataTypeUint16, it.Value.Restriction.To); err != nil {
					return nil, err
				}
			default:
				return nil, fmt.Errorf("invalid restriction type: %v", it.Value.Restriction.Type)
			}
			if g.Version > 1 {
				if err = buff.SetUint8(uint8(enums.DataTypeArray)); err != nil {
					return nil, err
				}
				if err = types.SetObjectCount(len(it.Value.Columns), buff); err != nil {
					return nil, err
				}
				for _, it2 := range it.Value.Columns {
					if err = buff.SetUint8(uint8(enums.DataTypeStructure)); err != nil {
						return nil, err
					}
					if err = buff.SetUint8(4); err != nil {
						return nil, err
					}
					if err = internal.SetData(settings, buff, enums.DataTypeUint16, int(it2.Key.Base().ObjectType())); err != nil {
						return nil, err
					}
					ln2, err := helpers.LogicalNameToBytes(it2.Key.Base().LogicalName())
					if err != nil {
						return nil, err
					}
					if err = internal.SetData(settings, buff, enums.DataTypeOctetString, ln2); err != nil {
						return nil, err
					}
					if err = internal.SetData(settings, buff, enums.DataTypeInt8, it2.Value.AttributeIndex); err != nil {
						return nil, err
					}
					if err = internal.SetData(settings, buff, enums.DataTypeUint16, it2.Value.DataIndex); err != nil {
						return nil, err
					}
				}
			}
		}
	}
	return buff.Array(), nil
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
func (g *GXDLMSPushSetup) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	switch e.Index {
	case 1:
		return helpers.LogicalNameToBytes(g.LogicalName())
	case 2:
		return g.GetPushObjectList(settings)
	case 3:
		buff := types.NewGXByteBuffer()
		if err := buff.SetUint8(uint8(enums.DataTypeStructure)); err != nil {
			return nil, err
		}
		if err := buff.SetUint8(3); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, buff, enums.DataTypeEnum, g.Service); err != nil {
			return nil, err
		}
		if g.Destination != "" {
			if g.Service == enums.ServiceTypeHDLC {
				if ln, err := helpers.LogicalNameToBytes(g.Destination); err == nil && len(ln) == 6 && ln[5] == 0xFF {
					if err = internal.SetData(settings, buff, enums.DataTypeOctetString, ln); err != nil {
						return nil, err
					}
				} else if err = internal.SetData(settings, buff, enums.DataTypeOctetString, []byte(g.Destination)); err != nil {
					return nil, err
				}
			} else if err := internal.SetData(settings, buff, enums.DataTypeOctetString, []byte(g.Destination)); err != nil {
				return nil, err
			}
		} else if err := internal.SetData(settings, buff, enums.DataTypeOctetString, nil); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, buff, enums.DataTypeEnum, g.Message); err != nil {
			return nil, err
		}
		return buff.Array(), nil
	case 4:
		buff := types.NewGXByteBuffer()
		if err := buff.SetUint8(uint8(enums.DataTypeArray)); err != nil {
			return nil, err
		}
		if err := types.SetObjectCount(len(g.CommunicationWindow), buff); err != nil {
			return nil, err
		}
		for _, it := range g.CommunicationWindow {
			if err := buff.SetUint8(uint8(enums.DataTypeStructure)); err != nil {
				return nil, err
			}
			if err := buff.SetUint8(2); err != nil {
				return nil, err
			}
			if err := internal.SetData(settings, buff, enums.DataTypeOctetString, it.Key); err != nil {
				return nil, err
			}
			if err := internal.SetData(settings, buff, enums.DataTypeOctetString, it.Value); err != nil {
				return nil, err
			}
		}
		return buff.Array(), nil
	case 5:
		return g.RandomisationStartInterval, nil
	case 6:
		return g.NumberOfRetries, nil
	case 7:
		if g.Version < 2 {
			return g.RepetitionDelay, nil
		}
		buff := types.NewGXByteBuffer()
		if err := buff.SetUint8(uint8(enums.DataTypeStructure)); err != nil {
			return nil, err
		}
		if err := types.SetObjectCount(3, buff); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, buff, enums.DataTypeUint16, g.RepetitionDelay2.Min); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, buff, enums.DataTypeUint16, g.RepetitionDelay2.Exponent); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, buff, enums.DataTypeUint16, g.RepetitionDelay2.Max); err != nil {
			return nil, err
		}
		return buff.Array(), nil
	case 8:
		if g.PortReference.Base().LogicalName() == "" {
			return nil, nil
		}
		return helpers.LogicalNameToBytes(g.PortReference.Base().LogicalName())
	case 9:
		return g.PushClientSAP, nil
	case 10:
		buff := types.NewGXByteBuffer()
		if err := buff.SetUint8(uint8(enums.DataTypeArray)); err != nil {
			return nil, err
		}
		if err := types.SetObjectCount(len(g.PushProtectionParameters), buff); err != nil {
			return nil, err
		}
		for _, it := range g.PushProtectionParameters {
			if err := buff.SetUint8(uint8(enums.DataTypeStructure)); err != nil {
				return nil, err
			}
			if err := buff.SetUint8(2); err != nil {
				return nil, err
			}
			if err := internal.SetData(settings, buff, enums.DataTypeEnum, it.ProtectionType); err != nil {
				return nil, err
			}
			if err := buff.SetUint8(uint8(enums.DataTypeStructure)); err != nil {
				return nil, err
			}
			if err := buff.SetUint8(5); err != nil {
				return nil, err
			}
			if err := internal.SetData(settings, buff, enums.DataTypeOctetString, it.TransactionId); err != nil {
				return nil, err
			}
			if err := internal.SetData(settings, buff, enums.DataTypeOctetString, it.OriginatorSystemTitle); err != nil {
				return nil, err
			}
			if err := internal.SetData(settings, buff, enums.DataTypeOctetString, it.RecipientSystemTitle); err != nil {
				return nil, err
			}
			if err := internal.SetData(settings, buff, enums.DataTypeOctetString, it.OtherInformation); err != nil {
				return nil, err
			}
			if err := buff.SetUint8(uint8(enums.DataTypeStructure)); err != nil {
				return nil, err
			}
			if err := buff.SetUint8(2); err != nil {
				return nil, err
			}
			if err := internal.SetData(settings, buff, enums.DataTypeEnum, it.KeyInfo.DataProtectionKeyType); err != nil {
				return nil, err
			}
			if err := buff.SetUint8(uint8(enums.DataTypeStructure)); err != nil {
				return nil, err
			}
			switch it.KeyInfo.DataProtectionKeyType {
			case enums.DataProtectionKeyTypeIdentified:
				if err := buff.SetUint8(1); err != nil {
					return nil, err
				}
				if err := internal.SetData(settings, buff, enums.DataTypeEnum, it.KeyInfo.IdentifiedKey.KeyType); err != nil {
					return nil, err
				}
			case enums.DataProtectionKeyTypeWrapped:
				if err := buff.SetUint8(2); err != nil {
					return nil, err
				}
				if err := internal.SetData(settings, buff, enums.DataTypeEnum, it.KeyInfo.WrappedKey.KeyType); err != nil {
					return nil, err
				}
				if err := internal.SetData(settings, buff, enums.DataTypeOctetString, it.KeyInfo.WrappedKey.Key); err != nil {
					return nil, err
				}
			case enums.DataProtectionKeyTypeAgreed:
				if err := buff.SetUint8(2); err != nil {
					return nil, err
				}
				if err := internal.SetData(settings, buff, enums.DataTypeOctetString, it.KeyInfo.AgreedKey.Parameters); err != nil {
					return nil, err
				}
				if err := internal.SetData(settings, buff, enums.DataTypeOctetString, it.KeyInfo.AgreedKey.Data); err != nil {
					return nil, err
				}
			default:
				if err := buff.SetUint8(0); err != nil {
					return nil, err
				}
			}
		}
		return buff.Array(), nil
	case 11:
		return g.PushOperationMethod, nil
	case 12:
		buff := types.NewGXByteBuffer()
		if err := buff.SetUint8(uint8(enums.DataTypeStructure)); err != nil {
			return nil, err
		}
		if err := types.SetObjectCount(2, buff); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, buff, enums.DataTypeOctetString, g.ConfirmationParameters.StartDate); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, buff, enums.DataTypeUint32, g.ConfirmationParameters.Interval); err != nil {
			return nil, err
		}
		return buff.Array(), nil
	case 13:
		return g.LastConfirmationDateTime, nil
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
		return nil, nil
	}
}

func (g *GXDLMSPushSetup) SetPushObject(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	toSlice := func(v any) ([]any, bool) {
		switch x := v.(type) {
		case types.GXArray:
			return []any(x), true
		case types.GXStructure:
			return []any(x), true
		case []any:
			return x, true
		default:
			return nil, false
		}
	}
	toInt := func(v any) (int, error) {
		switch x := v.(type) {
		case int:
			return x, nil
		case int8:
			return int(x), nil
		case int16:
			return int(x), nil
		case int32:
			return int(x), nil
		case uint8:
			return int(x), nil
		case uint16:
			return int(x), nil
		case uint32:
			return int(x), nil
		case types.GXEnum:
			return int(x.Value), nil
		default:
			return 0, fmt.Errorf("invalid integer value type: %T", v)
		}
	}

	g.PushObjectList = make([]types.GXKeyValuePair[IGXDLMSBase, GXDLMSCaptureObject], 0)
	if e.Value == nil {
		return nil
	}
	list, ok := toSlice(e.Value)
	if !ok {
		return fmt.Errorf("invalid push object list type: %T", e.Value)
	}
	objects := getObjectCollection(settings.Objects)
	for _, item := range list {
		it, ok := toSlice(item)
		if !ok || len(it) < 4 {
			return fmt.Errorf("invalid push object item type: %T", item)
		}
		ot, err := toInt(it[0])
		if err != nil {
			return err
		}
		ln, err := helpers.ToLogicalName(it[1])
		if err != nil {
			return err
		}
		obj := objects.FindByLN(enums.ObjectType(ot), ln)
		if obj == nil {
			obj, err = CreateObject(enums.ObjectType(ot), ln, 0)
			if err != nil {
				return err
			}
		}
		ai, err := toInt(it[2])
		if err != nil {
			return err
		}
		di, err := toInt(it[3])
		if err != nil {
			return err
		}
		co := GXDLMSCaptureObject{
			AttributeIndex: ai,
			DataIndex:      di,
		}
		if g.Version > 1 && len(it) > 4 {
			restriction, ok := toSlice(it[4])
			if ok && len(restriction) > 0 {
				rt, err := toInt(restriction[0])
				if err != nil {
					return err
				}
				co.Restriction.Type = enums.RestrictionType(rt)
				switch co.Restriction.Type {
				case enums.RestrictionTypeNone:
				case enums.RestrictionTypeDate, enums.RestrictionTypeEntry:
					if len(restriction) > 1 {
						co.Restriction.From = restriction[1]
					}
					if len(restriction) > 2 {
						co.Restriction.To = restriction[2]
					}
				default:
					return fmt.Errorf("invalid restriction type: %v", co.Restriction.Type)
				}
			}
			if len(it) > 5 {
				columns, ok := toSlice(it[5])
				if ok {
					for _, c := range columns {
						column, ok := toSlice(c)
						if !ok || len(column) < 4 {
							return fmt.Errorf("invalid push column type: %T", c)
						}
						columnObjectType, err := toInt(column[0])
						if err != nil {
							return err
						}
						columnLN, err := helpers.ToLogicalName(column[1])
						if err != nil {
							return err
						}
						columnObj := objects.FindByLN(enums.ObjectType(columnObjectType), columnLN)
						if columnObj == nil {
							columnObj, err = CreateObject(enums.ObjectType(columnObjectType), columnLN, 0)
							if err != nil {
								return err
							}
						}
						columnAI, err := toInt(column[2])
						if err != nil {
							return err
						}
						columnDI, err := toInt(column[3])
						if err != nil {
							return err
						}
						co.Columns = append(co.Columns, types.GXKeyValuePair[IGXDLMSBase, GXDLMSCaptureObject]{
							Key: columnObj,
							Value: GXDLMSCaptureObject{
								AttributeIndex: columnAI,
								DataIndex:      columnDI,
							},
						})
					}
				}
			}
		}
		g.PushObjectList = append(g.PushObjectList, types.GXKeyValuePair[IGXDLMSBase, GXDLMSCaptureObject]{
			Key:   obj,
			Value: co,
		})
	}
	return nil
}

func (g *GXDLMSPushSetup) GetPushProtectionParameters(e *internal.ValueEventArgs) error {
	toSlice := func(v any) ([]any, bool) {
		switch x := v.(type) {
		case types.GXArray:
			return []any(x), true
		case types.GXStructure:
			return []any(x), true
		case []any:
			return x, true
		default:
			return nil, false
		}
	}
	toInt := func(v any) (int, error) {
		switch x := v.(type) {
		case int:
			return x, nil
		case int8:
			return int(x), nil
		case int16:
			return int(x), nil
		case int32:
			return int(x), nil
		case uint8:
			return int(x), nil
		case uint16:
			return int(x), nil
		case uint32:
			return int(x), nil
		case types.GXEnum:
			return int(x.Value), nil
		default:
			return 0, fmt.Errorf("invalid integer value type: %T", v)
		}
	}
	toBytes := func(v any) ([]byte, error) {
		switch x := v.(type) {
		case []byte:
			return x, nil
		case string:
			return []byte(x), nil
		default:
			return nil, fmt.Errorf("invalid byte string type: %T", v)
		}
	}

	list := make([]GXPushProtectionParameters, 0)
	if e.Value == nil {
		g.PushProtectionParameters = list
		return nil
	}
	items, ok := toSlice(e.Value)
	if !ok {
		return fmt.Errorf("invalid push protection parameter type: %T", e.Value)
	}
	for _, item := range items {
		it, ok := toSlice(item)
		if !ok || len(it) < 2 {
			return fmt.Errorf("invalid push protection item type: %T", item)
		}
		protectionType, err := toInt(it[0])
		if err != nil {
			return err
		}
		options, ok := toSlice(it[1])
		if !ok || len(options) < 5 {
			return fmt.Errorf("invalid push protection options type: %T", it[1])
		}
		p := GXPushProtectionParameters{
			ProtectionType: enums.ProtectionType(protectionType),
		}
		if p.TransactionId, err = toBytes(options[0]); err != nil {
			return err
		}
		if p.OriginatorSystemTitle, err = toBytes(options[1]); err != nil {
			return err
		}
		if p.RecipientSystemTitle, err = toBytes(options[2]); err != nil {
			return err
		}
		if p.OtherInformation, err = toBytes(options[3]); err != nil {
			return err
		}
		keyInfo, ok := toSlice(options[4])
		if !ok || len(keyInfo) < 2 {
			return fmt.Errorf("invalid data protection key info type: %T", options[4])
		}
		keyType, err := toInt(keyInfo[0])
		if err != nil {
			return err
		}
		p.KeyInfo.DataProtectionKeyType = enums.DataProtectionKeyType(keyType)
		data, ok := toSlice(keyInfo[1])
		if !ok {
			return fmt.Errorf("invalid data protection key data type: %T", keyInfo[1])
		}
		switch p.KeyInfo.DataProtectionKeyType {
		case enums.DataProtectionKeyTypeIdentified:
			if len(data) > 0 {
				keyValue, err := toInt(data[0])
				if err != nil {
					return err
				}
				p.KeyInfo.IdentifiedKey.KeyType = enums.DataProtectionIdentifiedKeyType(keyValue)
			}
		case enums.DataProtectionKeyTypeWrapped:
			if len(data) > 0 {
				keyValue, err := toInt(data[0])
				if err != nil {
					return err
				}
				p.KeyInfo.WrappedKey.KeyType = enums.DataProtectionWrappedKeyType(keyValue)
			}
			if len(data) > 1 {
				p.KeyInfo.WrappedKey.Key, err = toBytes(data[1])
				if err != nil {
					return err
				}
			}
		case enums.DataProtectionKeyTypeAgreed:
			if len(data) > 0 {
				p.KeyInfo.AgreedKey.Parameters, err = toBytes(data[0])
				if err != nil {
					return err
				}
			}
			if len(data) > 1 {
				p.KeyInfo.AgreedKey.Data, err = toBytes(data[1])
				if err != nil {
					return err
				}
			}
		default:
			return fmt.Errorf("invalid data protection key type: %v", p.KeyInfo.DataProtectionKeyType)
		}
		list = append(list, p)
	}
	g.PushProtectionParameters = list
	return nil
}

// SetValue returns the set value of given attribute.
// When raw parameter us not used example register multiplies value by scalar.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Set parameters.
func (g *GXDLMSPushSetup) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	toSlice := func(v any) ([]any, bool) {
		switch x := v.(type) {
		case types.GXArray:
			return []any(x), true
		case types.GXStructure:
			return []any(x), true
		case []any:
			return x, true
		default:
			return nil, false
		}
	}
	toInt := func(v any) (int, error) {
		switch x := v.(type) {
		case int:
			return x, nil
		case int8:
			return int(x), nil
		case int16:
			return int(x), nil
		case int32:
			return int(x), nil
		case int64:
			return int(x), nil
		case uint8:
			return int(x), nil
		case uint16:
			return int(x), nil
		case uint32:
			return int(x), nil
		case uint64:
			return int(x), nil
		case types.GXEnum:
			return int(x.Value), nil
		default:
			return 0, fmt.Errorf("invalid integer value type: %T", v)
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
		return g.SetPushObject(settings, e)
	case 3:
		tmp, ok := toSlice(e.Value)
		if !ok || len(tmp) < 3 {
			return fmt.Errorf("invalid destination and method structure: %T", e.Value)
		}
		service, err := toInt(tmp[0])
		if err != nil {
			return err
		}
		g.Service = enums.ServiceType(service)
		if destinationBytes, ok := tmp[1].([]byte); ok {
			if g.Service == enums.ServiceTypeHDLC && len(destinationBytes) == 6 && destinationBytes[5] == 0xFF {
				g.Destination, err = helpers.ToLogicalName(destinationBytes)
				if err != nil {
					return err
				}
			} else {
				str, err := internal.ChangeTypeFromByteArray(settings, destinationBytes, enums.DataTypeString)
				if err != nil {
					g.Destination = types.ToHex(destinationBytes, true)
				} else if v, ok := str.(string); ok {
					g.Destination = v
				} else {
					g.Destination = types.ToHex(destinationBytes, true)
				}
			}
		} else if destinationText, ok := tmp[1].(string); ok {
			g.Destination = destinationText
		} else {
			return fmt.Errorf("invalid destination type: %T", tmp[1])
		}
		message, err := toInt(tmp[2])
		if err != nil {
			return err
		}
		g.Message = enums.MessageType(message)
	case 4:
		g.CommunicationWindow = make([]types.GXKeyValuePair[types.GXDateTime, types.GXDateTime], 0)
		if e.Value == nil {
			return nil
		}
		items, ok := toSlice(e.Value)
		if !ok {
			return fmt.Errorf("invalid communication window type: %T", e.Value)
		}
		for _, item := range items {
			it, ok := toSlice(item)
			if !ok || len(it) < 2 {
				return fmt.Errorf("invalid communication window item type: %T", item)
			}
			var start types.GXDateTime
			var end types.GXDateTime
			switch v := it[0].(type) {
			case []byte:
				tmp, err := internal.ChangeTypeFromByteArray(settings, v, enums.DataTypeDateTime)
				if err != nil {
					return err
				}
				start = tmp.(types.GXDateTime)
			case types.GXDateTime:
				start = v
			default:
				return fmt.Errorf("invalid communication window start type: %T", it[0])
			}
			switch v := it[1].(type) {
			case []byte:
				tmp, err := internal.ChangeTypeFromByteArray(settings, v, enums.DataTypeDateTime)
				if err != nil {
					return err
				}
				end = tmp.(types.GXDateTime)
			case types.GXDateTime:
				end = v
			default:
				return fmt.Errorf("invalid communication window end type: %T", it[1])
			}
			g.CommunicationWindow = append(g.CommunicationWindow, types.GXKeyValuePair[types.GXDateTime, types.GXDateTime]{Key: start, Value: end})
		}
	case 5:
		value, err := toInt(e.Value)
		if err != nil {
			return err
		}
		g.RandomisationStartInterval = uint16(value)
	case 6:
		value, err := toInt(e.Value)
		if err != nil {
			return err
		}
		g.NumberOfRetries = uint8(value)
	case 7:
		if g.Version < 2 {
			value, err := toInt(e.Value)
			if err != nil {
				return err
			}
			g.RepetitionDelay = uint16(value)
		} else {
			s, ok := toSlice(e.Value)
			if !ok || len(s) < 3 {
				return fmt.Errorf("invalid repetition delay structure: %T", e.Value)
			}
			min, err := toInt(s[0])
			if err != nil {
				return err
			}
			exponent, err := toInt(s[1])
			if err != nil {
				return err
			}
			max, err := toInt(s[2])
			if err != nil {
				return err
			}
			g.RepetitionDelay2.Min = uint16(min)
			g.RepetitionDelay2.Exponent = uint16(exponent)
			g.RepetitionDelay2.Max = uint16(max)
		}
	default:
		if g.Version > 0 && e.Index == 8 {
			g.PortReference = nil
			if bv, ok := e.Value.([]byte); ok {
				ln, err := helpers.ToLogicalName(bv)
				if err != nil {
					return err
				}
				g.PortReference = getObjectCollection(settings.Objects).FindByLN(enums.ObjectTypeNone, ln)
			}
			return nil
		}
		if g.Version > 0 && e.Index == 9 {
			value, err := toInt(e.Value)
			if err != nil {
				return err
			}
			g.PushClientSAP = int8(value)
			return nil
		}
		if g.Version > 0 && e.Index == 10 {
			return g.GetPushProtectionParameters(e)
		}
		if g.Version > 1 && e.Index == 11 {
			value, err := toInt(e.Value)
			if err != nil {
				return err
			}
			g.PushOperationMethod = enums.PushOperationMethod(value)
			return nil
		}
		if g.Version > 1 && e.Index == 12 {
			s, ok := toSlice(e.Value)
			if !ok || len(s) < 2 {
				e.Error = enums.ErrorCodeReadWriteDenied
				return nil
			}
			g.ConfirmationParameters.StartDate = s[0].(types.GXDateTime)
			interval, err := toInt(s[1])
			if err != nil {
				return err
			}
			g.ConfirmationParameters.Interval = uint32(interval)
			return nil
		}
		if g.Version > 1 && e.Index == 13 {
			if dt, ok := e.Value.(types.GXDateTime); ok {
				g.LastConfirmationDateTime = dt
			} else {
				e.Error = enums.ErrorCodeReadWriteDenied
			}
			return nil
		}
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil
}

// Load returns the load object content from XML.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSPushSetup) Load(reader *GXXmlReader) error {
	var err error
	g.PushObjectList = g.PushObjectList[:0]
	// Unknown method data type List.Clear
	if ret, err := reader.IsStartElementNamed("ObjectList", true); ret && err == nil {
		for {
			ret, err = reader.IsStartElementNamed("Item", true)
			if err != nil {
				return err
			}
			if !ret {
				break
			}
			ret, err := reader.ReadElementContentAsInt("ObjectType", 0)
			if err != nil {
				return err
			}
			ot := enums.ObjectType(ret)
			ln, err := reader.ReadElementContentAsString("LN", "")
			if err != nil {
				return err
			}
			ai, err := reader.ReadElementContentAsInt("AI", 0)
			if err != nil {
				return err
			}
			di, err := reader.ReadElementContentAsInt("DI", 0)
			if err != nil {
				return err
			}
			reader.ReadEndElement("Item")
			co := *NewGXDLMSCaptureObject(ai, di)
			obj := reader.Objects.FindByLN(ot, ln)
			if obj == nil {
				obj, err = CreateObject(ot, ln, 0)
				if err != nil {
					return err
				}
			}
			g.PushObjectList = append(g.PushObjectList, *types.NewGXKeyValuePair(obj, co))
		}
		reader.ReadEndElement("ObjectList")
	}
	ret, err := reader.ReadElementContentAsInt("Service", 0)
	if err != nil {
		return err
	}
	g.Service = enums.ServiceType(ret)
	g.Destination, err = reader.ReadElementContentAsString("Destination", "")
	if err != nil {
		return err
	}
	ret, err = reader.ReadElementContentAsInt("Message", 0)
	if err != nil {
		return err
	}
	g.Message = enums.MessageType(ret)
	g.CommunicationWindow = g.CommunicationWindow[:0]
	if ret, err := reader.IsStartElementNamed("CommunicationWindow", true); ret && err == nil {
		for {
			ret, err = reader.IsStartElementNamed("Item", true)
			if err != nil {
				return err
			}
			if !ret {
				break
			}
			start, err := reader.ReadElementContentAsGXDateTime("Start")
			if err != nil {
				return err
			}
			end, err := reader.ReadElementContentAsGXDateTime("End")
			if err != nil {
				return err
			}
			g.CommunicationWindow = append(g.CommunicationWindow, *types.NewGXKeyValuePair(start, end))
		}
		reader.ReadEndElement("CommunicationWindow")
	}
	g.RandomisationStartInterval, err = reader.ReadElementContentAsUInt16("RandomisationStartInterval", 0)
	if err != nil {
		return err
	}
	g.NumberOfRetries, err = reader.ReadElementContentAsUInt8("NumberOfRetries", 0)
	if err != nil {
		return err
	}
	if g.Version < 2 {
		g.RepetitionDelay, err = reader.ReadElementContentAsUInt16("RepetitionDelay", 0)
		if err != nil {
			return err
		}
	} else {
		if ret, err := reader.IsStartElementNamed("RepetitionDelay", true); ret && err == nil {
			g.RepetitionDelay2.Min, err = reader.ReadElementContentAsUInt16("Min", 0)
			if err != nil {
				return err
			}
			g.RepetitionDelay2.Exponent, err = reader.ReadElementContentAsUInt16("Exponent", 0)
			if err != nil {
				return err
			}
			g.RepetitionDelay2.Max, err = reader.ReadElementContentAsUInt16("Max", 0)
			if err != nil {
				return err
			}
		}
		reader.ReadEndElement("RepetitionDelay")
	}
	if g.Version > 0 {
		g.PortReference = nil
		ln, err := reader.ReadElementContentAsString("PortReference", "")
		if err != nil {
			return err
		}
		g.PortReference = reader.Objects.FindByLN(enums.ObjectTypeNone, ln)
		if g.PortReference == nil {
			g.PortReference, err = CreateObject(enums.ObjectTypeIecHdlcSetup, ln, 0)
			if err != nil {
				return err
			}
		}
		g.PushClientSAP, err = reader.ReadElementContentAsInt8("PushClientSAP", 0)
		if err != nil {
			return err
		}
		if ret, err := reader.IsStartElementNamed("PushProtectionParameters", true); ret && err == nil {
			var list []GXPushProtectionParameters
			for {
				ret, err = reader.IsStartElementNamed("Item", true)
				if err != nil {
					return err
				}
				if !ret {
					break
				}
				it := GXPushProtectionParameters{}
				ret, err := reader.ReadElementContentAsInt("ProtectionType", 0)
				if err != nil {
					return err
				}
				it.ProtectionType = enums.ProtectionType(ret)
				ret2, err := reader.ReadElementContentAsString("TransactionId", "")
				if err != nil {
					return err
				}
				it.TransactionId = types.HexToBytes(ret2)
				ret2, err = reader.ReadElementContentAsString("OriginatorSystemTitle", "")
				if err != nil {
					return err
				}
				it.OriginatorSystemTitle = types.HexToBytes(ret2)
				ret2, err = reader.ReadElementContentAsString("RecipientSystemTitle", "")
				if err != nil {
					return err
				}
				it.RecipientSystemTitle = types.HexToBytes(ret2)
				ret2, err = reader.ReadElementContentAsString("OtherInformation", "")
				if err != nil {
					return err
				}
				it.OtherInformation = types.HexToBytes(ret2)
				ret, err = reader.ReadElementContentAsInt("DataProtectionKeyType", 0)
				if err != nil {
					return err
				}
				it.KeyInfo.DataProtectionKeyType = enums.DataProtectionKeyType(ret)
				ret, err = reader.ReadElementContentAsInt("IdentifiedKey", 0)
				if err != nil {
					return err
				}
				it.KeyInfo.IdentifiedKey.KeyType = enums.DataProtectionIdentifiedKeyType(ret)
				ret, err = reader.ReadElementContentAsInt("WrappedKeyType", 0)
				if err != nil {
					return err
				}
				it.KeyInfo.WrappedKey.KeyType = enums.DataProtectionWrappedKeyType(ret)
				ret2, err = reader.ReadElementContentAsString("WrappedKey", "")
				if err != nil {
					return err
				}
				it.KeyInfo.WrappedKey.Key = types.HexToBytes(ret2)
				ret2, err = reader.ReadElementContentAsString("WrappedKeyParameters", "")
				if err != nil {
					return err
				}
				it.KeyInfo.AgreedKey.Parameters = types.HexToBytes(ret2)
				ret2, err = reader.ReadElementContentAsString("AgreedKeyData", "")
				if err != nil {
					return err
				}
				it.KeyInfo.AgreedKey.Data = types.HexToBytes(ret2)
				list = append(list, it)
			}
			reader.ReadEndElement("PushProtectionParameters")
			g.PushProtectionParameters = list
		}
		if g.Version > 1 {
			ret, err = reader.ReadElementContentAsInt("PushOperationMethod", 0)
			if err != nil {
				return err
			}
			g.PushOperationMethod = enums.PushOperationMethod(ret)
			g.ConfirmationParameters.StartDate, err = reader.ReadElementContentAsDateTime("ConfirmationParametersStartDate", nil)
			if err != nil {
				return err
			}
			g.ConfirmationParameters.Interval, err = reader.ReadElementContentAsUInt32("ConfirmationParametersInterval", 0)
			if err != nil {
				return err
			}
			g.LastConfirmationDateTime, err = reader.ReadElementContentAsDateTime("LastConfirmationDateTime", nil)
			if err != nil {
				return err
			}
		}
	}
	return err
}

// Save returns the save object content to XML.
//
// Parameters:
//
//	writer: XML writer.
func (g *GXDLMSPushSetup) Save(writer *GXXmlWriter) error {
	var err error
	if g.PushObjectList != nil {
		writer.WriteStartElement("ObjectList")
		for _, it := range g.PushObjectList {
			writer.WriteStartElement("Item")
			writer.WriteElementString("ObjectType", int(it.Key.Base().ObjectType()))
			writer.WriteElementString("LN", it.Key.Base().LogicalName())
			writer.WriteElementString("AI", it.Value.AttributeIndex)
			writer.WriteElementString("DI", it.Value.DataIndex)
			writer.WriteEndElement()
		}
		writer.WriteEndElement()
	}
	err = writer.WriteElementString("Service", int(g.Service))
	if err != nil {
		return err
	}
	err = writer.WriteElementString("Destination", g.Destination)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("Message", int(g.Message))
	if err != nil {
		return err
	}
	if g.CommunicationWindow != nil {
		writer.WriteStartElement("CommunicationWindow")
		for k, v := range g.CommunicationWindow {
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
	err = writer.WriteElementString("RandomisationStartInterval", g.RandomisationStartInterval)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("NumberOfRetries", g.NumberOfRetries)
	if err != nil {
		return err
	}
	if g.Version < 2 {
		err = writer.WriteElementString("RepetitionDelay", g.RepetitionDelay)
		if err != nil {
			return err
		}
	} else {
		writer.WriteStartElement("RepetitionDelay")
		err = writer.WriteElementString("Min", g.RepetitionDelay2.Min)
		if err != nil {
			return err
		}
		err = writer.WriteElementString("Exponent", g.RepetitionDelay2.Exponent)
		if err != nil {
			return err
		}
		err = writer.WriteElementString("Max", g.RepetitionDelay2.Max)
		if err != nil {
			return err
		}
		writer.WriteEndElement()
	}
	if g.Version > 0 {
		if g.PortReference != nil {
			err = writer.WriteElementString("PortReference", g.PortReference.Base().LogicalName())
			if err != nil {
				return err
			}
		}
		err = writer.WriteElementString("PushClientSAP", g.PushClientSAP)
		if err != nil {
			return err
		}
		if g.PushProtectionParameters != nil {
			writer.WriteStartElement("PushProtectionParameters")
			for _, it := range g.PushProtectionParameters {
				writer.WriteStartElement("Item")
				err = writer.WriteElementString("ProtectionType", uint8(it.ProtectionType))
				if err != nil {
					return err
				}
				err = writer.WriteElementString("TransactionId", types.ToHex(it.TransactionId, false))
				if err != nil {
					return err
				}
				err = writer.WriteElementString("OriginatorSystemTitle", types.ToHex(it.OriginatorSystemTitle, false))
				if err != nil {
					return err
				}
				err = writer.WriteElementString("RecipientSystemTitle", types.ToHex(it.RecipientSystemTitle, false))
				if err != nil {
					return err
				}
				err = writer.WriteElementString("OtherInformation", types.ToHex(it.OtherInformation, false))
				if err != nil {
					return err
				}
				err = writer.WriteElementString("DataProtectionKeyType", int(it.KeyInfo.DataProtectionKeyType))
				if err != nil {
					return err
				}
				writer.WriteElementString("IdentifiedKey", int(it.KeyInfo.IdentifiedKey.KeyType))
				writer.WriteElementString("WrappedKeyType", int(it.KeyInfo.WrappedKey.KeyType))
				writer.WriteElementString("WrappedKey", types.ToHex(it.KeyInfo.WrappedKey.Key, false))
				writer.WriteElementString("WrappedKeyParameters", types.ToHex(it.KeyInfo.AgreedKey.Parameters, false))
				writer.WriteElementString("AgreedKeyData", types.ToHex(it.KeyInfo.AgreedKey.Data, false))
				writer.WriteEndElement()
			}
			writer.WriteEndElement()
		}
		if g.Version > 1 {
			err = writer.WriteElementString("PushOperationMethod", int(g.PushOperationMethod))
			if err != nil {
				return err
			}
			err = writer.WriteElementString("ConfirmationParametersStartDate", g.ConfirmationParameters.StartDate)
			if err != nil {
				return err
			}
			err = writer.WriteElementString("ConfirmationParametersInterval", g.ConfirmationParameters.Interval)
			if err != nil {
				return err
			}
			err = writer.WriteElementString("LastConfirmationDateTime", g.LastConfirmationDateTime)
			if err != nil {
				return err
			}
		}
	}
	return err
}

// PostLoad returns the handle actions after Load.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSPushSetup) PostLoad(reader *GXXmlReader) error {
	// Update port reference.
	if g.PortReference != nil {
		target := reader.Objects.FindByLN(enums.ObjectTypeNone, g.PortReference.Base().LogicalName())
		if target != nil && target != g.PortReference {
			g.PortReference = target
		}
		// Upload object list after load.
		if g.PushObjectList != nil && len(g.PushObjectList) != 0 {
			g.PushObjectList = g.PushObjectList[:0]
			for _, it := range g.PushObjectList {
				var obj = it.Key
				target = reader.Objects.FindByLN(obj.Base().ObjectType(), obj.Base().LogicalName())
				if target != nil && target != obj {
					obj = target
				}
				g.PushObjectList = append(g.PushObjectList, *types.NewGXKeyValuePair(obj, it.Value))
			}
		}
	}
	return nil
}

// GetValues returns the an array containing the COSEM object's attribute values.
func (g *GXDLMSPushSetup) GetValues() []any {
	if g.Version == 0 {
		return []any{g.LogicalName(), g.PushObjectList, []any{g.Service, g.Destination, g.Message},
			g.CommunicationWindow, g.RandomisationStartInterval, g.NumberOfRetries, g.RepetitionDelay}
	}
	if g.Version == 1 {
		return []any{g.LogicalName(), g.PushObjectList, []any{g.Service, g.Destination, g.Message},
			g.CommunicationWindow, g.RandomisationStartInterval, g.NumberOfRetries, g.RepetitionDelay, g.PortReference,
			g.PushClientSAP, g.PushProtectionParameters, g.PushOperationMethod, g.ConfirmationParameters, g.LastConfirmationDateTime}
	}
	return []any{g.LogicalName(), g.PushObjectList, []any{g.Service, g.Destination, g.Message},
		g.CommunicationWindow, g.RandomisationStartInterval, g.NumberOfRetries, g.RepetitionDelay2, g.PortReference,
		g.PushClientSAP, g.PushProtectionParameters, g.PushOperationMethod, g.ConfirmationParameters, g.LastConfirmationDateTime}
}

// GetPushValues returns the get received objects from push message.
//
// Parameters:
//
//	values: Received values.
//
// Returns:
//
//	Clone of captured COSEM objects.
func (g *GXDLMSPushSetup) GetPushValues(client IGXDLMSClient, values []any) error {
	if len(values) != len(g.PushObjectList) {
		return errors.New("size of the push object list is different than values")
	}
	type pushValueUpdater interface {
		UpdateValue(target IGXDLMSBase, attributeIndex int, value any,
			columns *[]types.GXKeyValuePair[IGXDLMSBase, *GXDLMSCaptureObject]) (any, error)
	}
	updater, ok := any(client).(pushValueUpdater)
	if !ok {
		return errors.New("client does not support UpdateValue")
	}
	toSlice := func(v any) ([]any, bool) {
		switch x := v.(type) {
		case []any:
			return x, true
		case types.GXArray:
			return []any(x), true
		case types.GXStructure:
			return []any(x), true
		default:
			return nil, false
		}
	}
	for pos, it := range g.PushObjectList {
		if it.Value.AttributeIndex == 0 {
			tmp, ok := toSlice(values[pos])
			if !ok {
				return fmt.Errorf("invalid push value type at index %d: %T", pos, values[pos])
			}
			attrCount := it.Key.GetAttributeCount()
			if len(tmp) < attrCount {
				return fmt.Errorf("push value count is too small at index %d: expected %d, got %d", pos, attrCount, len(tmp))
			}
			for index := 1; index <= attrCount; index++ {
				if _, err := updater.UpdateValue(it.Key, index, tmp[index-1], nil); err != nil {
					return err
				}
			}
		} else {
			if _, err := updater.UpdateValue(it.Key, it.Value.AttributeIndex, values[pos], nil); err != nil {
				return err
			}
		}
	}
	return nil
}

// Push returns the activates the push process.
//
// Parameters:
//
//	client: DLMS client.
//
// Returns:
//
//	Action bytes.
func (g *GXDLMSPushSetup) Push(client IGXDLMSClient) ([][]uint8, error) {
	return client.Method(g, 1, int8(0), enums.DataTypeInt8)
}

// Reset returns the reset the push process.
//
// Parameters:
//
//	client: DLMS client.
//
// Returns:
//
//	Action bytes.
func (g *GXDLMSPushSetup) Reset(client IGXDLMSClient) ([][]uint8, error) {
	return client.Method(g, 2, int8(0), enums.DataTypeInt8)
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
func (g *GXDLMSPushSetup) GetDataType(index int) (enums.DataType, error) {
	if index == 1 {
		return enums.DataTypeOctetString, nil
	}
	if index == 2 {
		return enums.DataTypeArray, nil
	}
	if index == 3 {
		return enums.DataTypeStructure, nil
	}
	if index == 4 {
		return enums.DataTypeArray, nil
	}
	if index == 5 {
		return enums.DataTypeUint16, nil
	}
	if index == 6 {
		return enums.DataTypeUint8, nil
	}
	if index == 7 {
		if g.Version < 2 {
			return enums.DataTypeUint16, nil
		}
		return enums.DataTypeStructure, nil
	}
	if g.Version > 0 {
		// PortReference
		if index == 8 {
			return enums.DataTypeOctetString, nil
		}
		// PushClientSAP
		if index == 9 {
			return enums.DataTypeInt8, nil
		}
		// PushProtectionParameters
		if index == 10 {
			return enums.DataTypeArray, nil
		}
		if g.Version > 1 {
			// PushOperationMethod
			if index == 11 {
				return enums.DataTypeEnum, nil
			}
			// ConfirmationParameters
			if index == 12 {
				return enums.DataTypeStructure, nil
			}
			// LastConfirmationDateTime
			if index == 13 {
				return enums.DataTypeDateTime, nil
			}
		}
	}
	return 0, dlmserrors.ErrInvalidAttributeIndex
}

// NewGXDLMPushSetup creates a new Push Setup object instance.
//
// The function validates `ln` before creating the object.
// `ln` is the Logical Name and `sn` is the Short Name of the object.
func NewGXDLMSPushSetup(ln string, sn int16) (*GXDLMSPushSetup, error) {
	err := ValidateLogicalName(ln)
	if err != nil {
		return nil, err
	}
	return &GXDLMSPushSetup{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypePushSetup,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
