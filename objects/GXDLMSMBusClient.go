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
	"reflect"

	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// Use this class to setup M-Bus slave devices and to exchange data with them.
// Online help:
// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSMBusClient
type GXDLMSMBusClient struct {
	GXDLMSObject
	// Provides reference to an M-Bus master port setup object, used to configure
	// an M-Bus port, each interface allowing to exchange data with one or more
	// M-Bus slave devices
	MBusPortReference string

	CaptureDefinition []types.GXKeyValuePair[string, string]

	CapturePeriod uint32

	PrimaryAddress uint8

	IdentificationNumber uint32

	ManufacturerID uint16

	// Carries the Version element of the data header as specified in
	// EN 13757-3 sub-clause 5.6.
	DataHeaderVersion uint8

	DeviceType enums.MBusDeviceType

	AccessNumber uint8

	Status uint8

	Alarm uint8

	Configuration uint16

	EncryptionKeyStatus enums.MBusEncryptionKeyStatus
}

// base returns the base GXDLMSObject of the object.
func (g *GXDLMSMBusClient) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

// Invoke returns the invokes method.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Invoke parameters.
func (g *GXDLMSMBusClient) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
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
func (g *GXDLMSMBusClient) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	// LN is static and read only once.
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	// MBusPortReference
	if all || g.CanRead(2) {
		attributes = append(attributes, 2)
	}
	// CaptureDefinition
	if all || g.CanRead(3) {
		attributes = append(attributes, 3)
	}
	// CapturePeriod
	if all || g.CanRead(4) {
		attributes = append(attributes, 4)
	}
	// PrimaryAddress
	if all || g.CanRead(5) {
		attributes = append(attributes, 5)
	}
	// IdentificationNumber
	if all || g.CanRead(6) {
		attributes = append(attributes, 6)
	}
	// ManufacturerID
	if all || g.CanRead(7) {
		attributes = append(attributes, 7)
	}
	// Version
	if all || g.CanRead(8) {
		attributes = append(attributes, 8)
	}
	// DeviceType
	if all || g.CanRead(9) {
		attributes = append(attributes, 9)
	}
	// AccessNumber
	if all || g.CanRead(10) {
		attributes = append(attributes, 10)
	}
	// Status
	if all || g.CanRead(11) {
		attributes = append(attributes, 11)
	}
	// Alarm
	if all || g.CanRead(12) {
		attributes = append(attributes, 12)
	}
	if g.Version > 0 {
		// Configuration
		if all || g.CanRead(13) {
			attributes = append(attributes, 13)
		}
		// EncryptionKeyStatus
		if all || g.CanRead(14) {
			attributes = append(attributes, 14)
		}
	}
	return attributes
}

// GetNames returns the names of attribute indexes.
func (g *GXDLMSMBusClient) GetNames() []string {
	if g.Version == 0 {
		return []string{"Logical Name", "MBus Port Reference", "Capture Definition", "Capture Period", "Primary Address", "Identification Number", "Manufacturer ID", "Version", "Device Type", "Access Number", "Status", "Alarm"}
	}
	return []string{"Logical Name", "MBus Port Reference", "Capture Definition", "Capture Period", "Primary Address", "Identification Number", "Manufacturer ID", "Version", "Device Type", "Access Number", "Status", "Alarm", "Configuration", "Encryption Key Status"}
}

// GetMethodNames returns the names of method indexes.
func (g *GXDLMSMBusClient) GetMethodNames() []string {
	return []string{"Slave install", "Slave deinstall", "Capture", "Reset alarm", "Synchronize clock", "Data send", "Set encryption key", "Transfer key"}
}

// GetAttributeCount returns the amount of attributes.
//
// Returns:
//
//	Count of attributes.
func (g *GXDLMSMBusClient) GetAttributeCount() int {
	if g.Version == 0 {
		return 12
	}
	return 14
}

// GetMethodCount returns the amount of methods.
func (g *GXDLMSMBusClient) GetMethodCount() int {
	return 8
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
func (g *GXDLMSMBusClient) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	var err error
	if e.Index == 1 {
		v, err := helpers.LogicalNameToBytes(g.LogicalName())
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return v, err
	}
	if e.Index == 2 {
		return helpers.LogicalNameToBytes(g.MBusPortReference)
	}
	if e.Index == 3 {
		buff := types.NewGXByteBuffer()
		err = buff.SetUint8(uint8(enums.DataTypeArray))
		if err != nil {
			return nil, err
		}
		err = types.SetObjectCount(len(g.CaptureDefinition), buff)
		if err != nil {
			return nil, err
		}
		for _, it := range g.CaptureDefinition {
			err = buff.SetUint8(uint8(enums.DataTypeStructure))
			if err != nil {
				return nil, err
			}
			err = buff.SetUint8(2)
			if err != nil {
				return nil, err
			}
			err = internal.SetData(settings, buff, enums.DataTypeUint8, it.Key)
			if err != nil {
				return nil, err
			}
			err = internal.SetData(settings, buff, enums.DataTypeOctetString, []byte(it.Value))
			if err != nil {
				return nil, err
			}
		}
		return buff.Array(), nil
	}
	if e.Index == 4 {
		return g.CapturePeriod, nil
	}
	if e.Index == 5 {
		return g.PrimaryAddress, nil
	}
	if e.Index == 6 {
		return g.IdentificationNumber, nil
	}
	if e.Index == 7 {
		return g.ManufacturerID, nil
	}
	if e.Index == 8 {
		return g.DataHeaderVersion, nil
	}
	if e.Index == 9 {
		return g.DeviceType, nil
	}
	if e.Index == 10 {
		return g.AccessNumber, nil
	}
	if e.Index == 11 {
		return g.Status, nil
	}
	if e.Index == 12 {
		return g.Alarm, nil
	}
	if g.Version > 0 {
		if e.Index == 13 {
			return g.Configuration, nil
		}
		if e.Index == 14 {
			return g.EncryptionKeyStatus, nil
		}
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
func (g *GXDLMSMBusClient) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	var err error
	if e.Index == 1 {
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		err = g.SetLogicalName(ln)
	} else if e.Index == 2 {
		g.MBusPortReference, err = helpers.ToLogicalName(e.Value)
	} else if e.Index == 3 {
		g.CaptureDefinition = g.CaptureDefinition[:0]
		if e.Value != nil {
			for _, tmp := range e.Value.(types.GXArray) {
				it := tmp.(types.GXStructure)
				ret, err := internal.ChangeTypeFromByteArray(settings, it[0].([]byte), enums.DataTypeOctetString)
				if err != nil {
					return err
				}
				g.CaptureDefinition = append(g.CaptureDefinition, *types.NewGXKeyValuePair[string, string](string(ret.([]byte)), ret.(string)))
			}
		}
	} else if e.Index == 4 {
		g.CapturePeriod = e.Value.(uint32)
	} else if e.Index == 5 {
		g.PrimaryAddress = e.Value.(uint8)
	} else if e.Index == 6 {
		g.IdentificationNumber = e.Value.(uint32)
	} else if e.Index == 7 {
		g.ManufacturerID = e.Value.(uint16)
	} else if e.Index == 8 {
		g.DataHeaderVersion = e.Value.(uint8)
	} else if e.Index == 9 {
		g.DeviceType = enums.MBusDeviceType(e.Value.(uint8))
	} else if e.Index == 10 {
		g.AccessNumber = e.Value.(uint8)
	} else if e.Index == 11 {
		g.Status = e.Value.(uint8)
	} else if e.Index == 12 {
		g.Alarm = e.Value.(uint8)
	} else if g.Version > 0 {
		if e.Index == 13 {
			g.Configuration = uint16(e.Value.(int))
		} else if e.Index == 14 {
			g.EncryptionKeyStatus = enums.MBusEncryptionKeyStatus(e.Value.(int))
		} else {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
	} else {
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return err
}

// Load returns the load object content from XML.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSMBusClient) Load(reader *GXXmlReader) error {
	var err error
	var ret int
	g.MBusPortReference, err = reader.ReadElementContentAsString("MBusPortReference", "")
	if err != nil {
		return err
	}
	g.CaptureDefinition = g.CaptureDefinition[:0]
	b, err := reader.IsStartElementNamed("CaptureDefinition", true)
	if err != nil {
		return err
	}
	if b {
		b, err = reader.IsStartElementNamed("Item", true)
		if err != nil {
			return err
		}
		for b {
			d, err := reader.ReadElementContentAsString("Data", "")
			if err != nil {
				return err
			}
			v, err := reader.ReadElementContentAsString("Value", "")
			if err != nil {
				return err
			}
			g.CaptureDefinition = append(g.CaptureDefinition, *types.NewGXKeyValuePair[string, string](d, v))
			b, err = reader.IsStartElementNamed("Item", true)
			if err != nil {
				return err
			}
		}
		reader.ReadEndElement("CaptureDefinition")
	}
	ret, err = reader.ReadElementContentAsInt("CapturePeriod", 0)
	if err != nil {
		return err
	}
	g.CapturePeriod = uint32(ret)
	ret, err = reader.ReadElementContentAsInt("PrimaryAddress", 0)
	if err != nil {
		return err
	}
	g.PrimaryAddress = uint8(ret)
	ret, err = reader.ReadElementContentAsInt("IdentificationNumber", 0)
	if err != nil {
		return err
	}
	g.IdentificationNumber = uint32(ret)
	ret, err = reader.ReadElementContentAsInt("ManufacturerID", 0)
	if err != nil {
		return err
	}
	g.ManufacturerID = uint16(ret)
	ret, err = reader.ReadElementContentAsInt("DataHeaderVersion", 0)
	if err != nil {
		return err
	}
	g.DataHeaderVersion = uint8(ret)
	ret, err = reader.ReadElementContentAsInt("DeviceType", 0)
	if err != nil {
		return err
	}
	g.DeviceType = enums.MBusDeviceType(ret)
	ret, err = reader.ReadElementContentAsInt("AccessNumber", 0)
	if err != nil {
		return err
	}
	g.AccessNumber = uint8(ret)
	ret, err = reader.ReadElementContentAsInt("Status", 0)
	if err != nil {
		return err
	}
	g.Status = uint8(ret)
	ret, err = reader.ReadElementContentAsInt("Alarm", 0)
	if err != nil {
		return err
	}
	g.Alarm = uint8(ret)
	if g.Version > 0 {
		ret, err = reader.ReadElementContentAsInt("Configuration", 0)
		if err != nil {
			return err
		}
		g.Configuration = uint16(ret)
		ret, err = reader.ReadElementContentAsInt("EncryptionKeyStatus", 0)
		if err != nil {
			return err
		}
		g.EncryptionKeyStatus = enums.MBusEncryptionKeyStatus(ret)
	}
	return err
}

// Save returns the save object content to XML.
//
// Parameters:
//
//	writer: XML writer.
func (g *GXDLMSMBusClient) Save(writer *GXXmlWriter) error {
	err := writer.WriteElementString("MBusPortReference", g.MBusPortReference)
	if err != nil {
		return err
	}
	writer.WriteStartElement("CaptureDefinition")
	if g.CaptureDefinition != nil {
		for k, v := range g.CaptureDefinition {
			writer.WriteStartElement("Item")
			err = writer.WriteElementString("Data", k)
			if err != nil {
				return err
			}
			err = writer.WriteElementString("Value", v)
			if err != nil {
				return err
			}
			writer.WriteEndElement()
		}
	}
	writer.WriteEndElement()
	err = writer.WriteElementString("CapturePeriod", g.CapturePeriod)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("PrimaryAddress", g.PrimaryAddress)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("IdentificationNumber", g.IdentificationNumber)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("ManufacturerID", g.ManufacturerID)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("DataHeaderVersion", g.DataHeaderVersion)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("DeviceType", int(g.DeviceType))
	if err != nil {
		return err
	}
	err = writer.WriteElementString("AccessNumber", g.AccessNumber)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("Status", g.Status)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("Alarm", g.Alarm)
	if err != nil {
		return err
	}
	if g.Version > 0 {
		err = writer.WriteElementString("Configuration", g.Configuration)
		if err != nil {
			return err
		}
		err = writer.WriteElementString("EncryptionKeyStatus", int(g.EncryptionKeyStatus))
		if err != nil {
			return err
		}
	}
	return err
}

// PostLoad returns the handle actions after Load.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSMBusClient) PostLoad(reader *GXXmlReader) error {
	return nil
}

// GetManufacturerName returns the EN 61107 manufacturer from the manufacturer ID
func (g *GXDLMSMBusClient) GetManufacturerName(value uint16) string {
	return internal.DecryptManufacturer(value)
}

// SlaveInstall returns the installs a slave device.
//
// Parameters:
//
//	client: DLMS client settings.
//	primaryAddress: Primary address.
//
// Returns:
//
//	Generated DLMS data.
func (g *GXDLMSMBusClient) SlaveInstall(client IGXDLMSClient, primaryAddress uint8) ([][]uint8, error) {
	return client.Method(g, 1, primaryAddress, enums.DataTypeInt8)
}

// SlaveDeInstall returns the de-installs a slave device.
//
// Parameters:
//
//	client: DLMS client settings.
//
// Returns:
//
//	Generated DLMS data.
func (g *GXDLMSMBusClient) SlaveDeInstall(client IGXDLMSClient) ([][]uint8, error) {
	return client.Method(g, 2, 0, enums.DataTypeInt8)
}

// Capture returns the captures values.
//
// Parameters:
//
//	client: DLMS client settings.
//
// Returns:
//
//	Generated DLMS data.
func (g *GXDLMSMBusClient) Capture(client IGXDLMSClient) ([][]uint8, error) {
	return client.Method(g, 3, 0, enums.DataTypeInt8)
}

// ResetAlarm returns the resets alarm state of the M-Bus slave device.
//
// Parameters:
//
//	client: DLMS client settings.
//
// Returns:
//
//	Generated DLMS data.
func (g *GXDLMSMBusClient) ResetAlarm(client IGXDLMSClient) ([][]uint8, error) {
	return client.Method(g, 4, 0, enums.DataTypeInt8)
}

// SynchronizeClock returns the synchronize the clock.
//
// Parameters:
//
//	client: DLMS client settings.
//
// Returns:
//
//	Generated DLMS data.
func (g *GXDLMSMBusClient) SynchronizeClock(client IGXDLMSClient) ([][]uint8, error) {
	return client.Method(g, 5, 0, enums.DataTypeInt8)
}

// SendData returns the sends data to the M-Bus slave device.
//
// Parameters:
//
//	client: DLMS client settings.
//	data: data to send
//
// Returns:
//
//	Generated DLMS data.
func (g *GXDLMSMBusClient) SendData(client IGXDLMSClient, data []GXMBusClientData) ([][]uint8, error) {
	bb := types.NewGXByteBuffer()
	err := bb.SetUint8(enums.DataTypeArray)
	if err != nil {
		return nil, err
	}
	types.SetObjectCount(len(data), bb)
	for _, it := range data {
		err = bb.SetUint8(enums.DataTypeStructure)
		if err != nil {
			return nil, err
		}
		err = bb.SetUint8(3)
		if err != nil {
			return nil, err
		}
		err = bb.SetUint8(enums.DataTypeOctetString)
		if err != nil {
			return nil, err
		}
		types.SetObjectCount(len(it.DataInformation), bb)
		err = bb.Set(it.DataInformation)
		if err != nil {
			return nil, err
		}
		err = bb.SetUint8(enums.DataTypeOctetString)
		if err != nil {
			return nil, err
		}
		types.SetObjectCount(len(it.ValueInformation), bb)
		err = bb.Set(it.ValueInformation)
		if err != nil {
			return nil, err
		}
		ret, err := internal.GetDLMSDataType(reflect.TypeOf(it.Data))
		if err != nil {
			return nil, err
		}
		err = internal.SetData(client.Settings(), bb, ret, it.Data)
		if err != nil {
			return nil, err
		}
	}
	return client.Method(g, 6, bb.Array(), enums.DataTypeArray)
}

// SetEncryptionKey returns the sets the encryption key in the M-Bus client and enables encrypted communication
// with the M-Bus slave device.
//
// Parameters:
//
//	client: DLMS client settings.
//	encryptionKey: encryption key
//
// Returns:
//
//	Generated DLMS data.
func (g *GXDLMSMBusClient) SetEncryptionKey(client IGXDLMSClient, encryptionKey []byte) ([][]uint8, error) {
	bb := types.NewGXByteBuffer()
	err := bb.SetUint8(enums.DataTypeOctetString)
	if err != nil {
		return nil, err
	}
	if encryptionKey == nil {
		err = bb.SetUint8(0)
		if err != nil {
			return nil, err
		}
	} else {
		types.SetObjectCount(len(encryptionKey), bb)
		err = bb.Set(encryptionKey)
		if err != nil {
			return nil, err
		}
	}
	return client.Method(g, 7, bb.Array(), enums.DataTypeArray)
}

// TransferKey returns the transfers an encryption key to the M-Bus slave device.
//
// Parameters:
//
//	client: DLMS client settings.
//	encryptionKey: encryption key
//
// Returns:
//
//	Generated DLMS data.
func (g *GXDLMSMBusClient) TransferKey(client IGXDLMSClient, encryptionKey []byte) ([][]uint8, error) {
	bb := types.NewGXByteBuffer()
	err := bb.SetUint8(enums.DataTypeOctetString)
	if err != nil {
		return nil, err
	}
	if encryptionKey == nil {
		err = bb.SetUint8(0)
		if err != nil {
			return nil, err
		}
	} else {
		types.SetObjectCount(len(encryptionKey), bb)
		err = bb.Set(encryptionKey)
		if err != nil {
			return nil, err
		}
	}
	return client.Method(g, 8, bb.Array(), enums.DataTypeArray)

}

// GetValues returns the an array containing the COSEM object's attribute values.
func (g *GXDLMSMBusClient) GetValues() []any {
	if g.Version == 0 {
		return []any{g.LogicalName(), g.MBusPortReference, g.CaptureDefinition,
			g.CapturePeriod, g.PrimaryAddress, g.IdentificationNumber, g.ManufacturerID,
			g.DataHeaderVersion, g.DeviceType, g.AccessNumber, g.Status, g.Alarm}
	}
	return []any{g.LogicalName(), g.MBusPortReference, g.CaptureDefinition, g.CapturePeriod,
		g.PrimaryAddress, g.IdentificationNumber, g.ManufacturerID, g.DataHeaderVersion,
		g.DeviceType, g.AccessNumber, g.Status, g.Alarm, g.Configuration, g.EncryptionKeyStatus}
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
func (g *GXDLMSMBusClient) GetDataType(index int) (enums.DataType, error) {
	if index == 1 {
		return enums.DataTypeOctetString, nil
	}
	if index == 2 {
		return enums.DataTypeOctetString, nil
	}
	if index == 3 {
		return enums.DataTypeArray, nil
	}
	if index == 4 {
		return enums.DataTypeUint32, nil
	}
	if index == 5 {
		return enums.DataTypeUint8, nil
	}
	if index == 6 {
		return enums.DataTypeUint32, nil
	}
	if index == 7 {
		return enums.DataTypeUint16, nil
	}
	if index == 8 {
		return enums.DataTypeUint8, nil
	}
	if index == 9 {
		return enums.DataTypeUint8, nil
	}
	if index == 10 {
		return enums.DataTypeUint8, nil
	}
	if index == 11 {
		return enums.DataTypeUint8, nil
	}
	if index == 12 {
		return enums.DataTypeUint8, nil
	}
	if g.Version > 0 {
		if index == 13 {
			return enums.DataTypeUint16, nil
		}
		if index == 14 {
			return enums.DataTypeEnum, nil
		}
	}
	return 0, errors.New("GetDataType failed. Invalid attribute index.")
}

// Constructor.
// ln: Logical Name of the object.
// sn: Short Name of the object.
func NewGXDLMSMBusClient(ln string, sn int16) (*GXDLMSMBusClient, error) {
	err := ValidateLogicalName(ln)
	if err != nil {
		return nil, err
	}
	return &GXDLMSMBusClient{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeMBusClient,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
