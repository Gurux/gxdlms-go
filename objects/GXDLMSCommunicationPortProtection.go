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
	"github.com/Gurux/gxdlms-go/dlmserrors"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// Represents a DLMS/COSEM Communication Port Protection object that manages access control and lockout mechanisms
// for a communication port based on failed authentication attempts.
// This class provides configuration and status information for communication port protection,
// including lockout timing, allowed failed attempts, and the current protection status. It is typically used in
// metering or device management scenarios to prevent unauthorized access by disabling the port after a
// configurable number of failed attempts. The protection behavior can be customized using properties such as
// ProtectionMode, AllowedFailedAttempts, and SteepnessFactor. The associated port is referenced via the Port
// property.
// Online help:
// https://www.gurux.fi/Gurux.internal.Objects.GXDLMSCommunicationPortProtection
type GXDLMSCommunicationPortProtection struct {
	GXDLMSObject
	// Controls the protection mode.
	ProtectionMode enums.ProtectionMode

	// Number of allowed failed communication attempts before port is disabled.
	AllowedFailedAttempts uint16

	// The lockout time.
	InitialLockoutTime uint32

	// Holds a factor that controls how the lockout time is increased with
	// each failed attempt.
	SteepnessFactor uint8

	// The lockout time.
	MaxLockoutTime uint32

	// The communication port being protected
	Port IGXDLMSBase

	// Current protection status.
	ProtectionStatus enums.ProtectionStatus

	// Failed attempts.
	FailedAttempts uint32

	// Total failed attempts.
	CumulativeFailedAttempts uint32
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSCommunicationPortProtection) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

// Invoke returns the invokes method.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Invoke parameters.
func (g *GXDLMSCommunicationPortProtection) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	if e.Index == 1 {
		g.FailedAttempts = 0
		if g.ProtectionMode == enums.ProtectionModeLockedOnFailedAttempts {
			g.ProtectionStatus = enums.ProtectionStatusUnlocked
		}
	} else {
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
func (g *GXDLMSCommunicationPortProtection) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	// LN is static and read only once.
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	// ProtectionMode
	if all || g.CanRead(2) {
		attributes = append(attributes, 2)
	}
	// AllowedFailedAttempts
	if all || g.CanRead(3) {
		attributes = append(attributes, 3)
	}
	// InitialLockoutTime
	if all || g.CanRead(4) {
		attributes = append(attributes, 4)
	}
	// SteepnessFactor
	if all || g.CanRead(5) {
		attributes = append(attributes, 5)
	}
	// MaxLockoutTime
	if all || g.CanRead(6) {
		attributes = append(attributes, 6)
	}
	// Port
	if all || g.CanRead(7) {
		attributes = append(attributes, 7)
	}
	// ProtectionStatus
	if all || g.CanRead(8) {
		attributes = append(attributes, 8)
	}
	// FailedAttempts
	if all || g.CanRead(9) {
		attributes = append(attributes, 9)
	}
	// CumulativeFailedAttempts
	if all || g.CanRead(10) {
		attributes = append(attributes, 10)
	}
	return attributes
}

// GetNames returns the names of attribute indexes.
func (g *GXDLMSCommunicationPortProtection) GetNames() []string {
	return []string{"Logical Name", "Protection mode", "Allowed failed attempts", "Initial lockout time", "Steepness factor", "Max lockout time", "Port", "Protection status", "Failed attempts", "Cumulative failed attempts"}
}

// GetMethodNames returns the names of method indexes.
func (g *GXDLMSCommunicationPortProtection) GetMethodNames() []string {
	return []string{"Reset"}
}

// GetAttributeCount returns the amount of attributes.
//
// Returns:
//
//	Count of attributes.
func (g *GXDLMSCommunicationPortProtection) GetAttributeCount() int {
	return 10
}

// GetMethodCount returns the amount of methods.
func (g *GXDLMSCommunicationPortProtection) GetMethodCount() int {
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
func (g *GXDLMSCommunicationPortProtection) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	var ret any
	var err error
	switch e.Index {
	case 1:
		ret, err = helpers.LogicalNameToBytes(g.LogicalName())
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return ret, err
	case 2:
		ret = g.ProtectionMode
	case 3:
		ret = g.AllowedFailedAttempts
	case 4:
		ret = g.InitialLockoutTime
	case 5:
		ret = g.SteepnessFactor
	case 6:
		ret = g.MaxLockoutTime
	case 7:
		if g.Port == nil {
			ret, err = helpers.LogicalNameToBytes("")
			if err != nil {
				e.Error = enums.ErrorCodeReadWriteDenied
			}
		} else {
			ret, err = helpers.LogicalNameToBytes(g.Port.Base().LogicalName())
			if err != nil {
				e.Error = enums.ErrorCodeReadWriteDenied
			}
		}
		return ret, err
	case 8:
		ret = g.ProtectionStatus
	case 9:
		ret = g.FailedAttempts
	case 10:
		ret = g.CumulativeFailedAttempts
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
		ret = nil
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
func (g *GXDLMSCommunicationPortProtection) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	var err error
	switch e.Index {
	case 1:
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return g.SetLogicalName(ln)
	case 2:
		g.ProtectionMode = enums.ProtectionMode(e.Value.(types.GXEnum).Value)
	case 3:
		g.AllowedFailedAttempts = e.Value.(uint16)
	case 4:
		g.InitialLockoutTime = e.Value.(uint32)
	case 5:
		g.SteepnessFactor = e.Value.(uint8)
	case 6:
		g.MaxLockoutTime = e.Value.(uint32)
	case 7:
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			break
		}
		g.Port = getObjectCollection(settings.Objects).FindByLN(enums.ObjectTypeNone, ln)
	case 8:
		g.ProtectionStatus = enums.ProtectionStatus(e.Value.(types.GXEnum).Value)
	case 9:
		g.FailedAttempts = e.Value.(uint32)
	case 10:
		g.CumulativeFailedAttempts = e.Value.(uint32)
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
func (g *GXDLMSCommunicationPortProtection) Load(reader *GXXmlReader) error {
	var err error
	var i int
	var l int64
	i, err = reader.ReadElementContentAsInt("ProtectionMode", 0)
	if err != nil {
		return err
	}
	g.ProtectionMode = enums.ProtectionMode(i)
	i, err = reader.ReadElementContentAsInt("AllowedFailedAttempts", 0)
	if err != nil {
		return err
	}
	g.AllowedFailedAttempts = uint16(i)
	l, err = reader.ReadElementContentAsLong("InitialLockoutTime", 0)
	if err != nil {
		return err
	}
	g.InitialLockoutTime = uint32(l)
	i, err = reader.ReadElementContentAsInt("SteepnessFactor", 0)
	if err != nil {
		return err
	}
	g.SteepnessFactor = uint8(i)
	l, err = reader.ReadElementContentAsLong("MaxLockoutTime", 0)
	if err != nil {
		return err
	}
	g.MaxLockoutTime = uint32(l)
	port, err := reader.ReadElementContentAsString("Port", "")
	if err != nil {
		return err
	}
	if port == "" {
		g.Port = nil
	} else {
		g.Port = reader.Objects.FindByLN(enums.ObjectTypeNone, port)
		// Save port object for data object if it's not loaded yet.
		if g.Port == nil {
			g.Port, err = CreateObject(enums.ObjectTypeData, port, 0)
			if err != nil {
				return err
			}
			g.Port.Base().Version = 0
		}
	}
	i, err = reader.ReadElementContentAsInt("ProtectionStatus", 0)
	g.ProtectionStatus = enums.ProtectionStatus(i)
	l, err = reader.ReadElementContentAsLong("FailedAttempts", 0)
	if err != nil {
		return err
	}
	g.FailedAttempts = uint32(l)
	l, err = reader.ReadElementContentAsLong("CumulativeFailedAttempts", 0)
	if err != nil {
		return err
	}
	g.CumulativeFailedAttempts = uint32(l)
	return nil
}

// Save returns the save object content to XML.
//
// Parameters:
//
//	writer: XML writer.
func (g *GXDLMSCommunicationPortProtection) Save(writer *GXXmlWriter) error {
	err := writer.WriteElementStringInt("ProtectionMode", int(g.ProtectionMode))
	if err != nil {
		return err
	}
	err = writer.WriteElementStringInt("AllowedFailedAttempts", int(g.AllowedFailedAttempts))
	if err != nil {
		return err
	}
	err = writer.WriteElementStringU32("InitialLockoutTime", g.InitialLockoutTime)
	if err != nil {
		return err
	}
	err = writer.WriteElementStringInt("SteepnessFactor", int(g.SteepnessFactor))
	if err != nil {
		return err
	}
	err = writer.WriteElementStringU32("MaxLockoutTime", g.MaxLockoutTime)
	if err != nil {
		return err
	}
	if g.Port == nil {
		writer.WriteElementString("Port", "")
	} else {
		err = writer.WriteElementString("Port", g.Port.Base().LogicalName())
		if err != nil {
			return err
		}
	}
	err = writer.WriteElementStringInt("ProtectionStatus", int(g.ProtectionStatus))
	if err != nil {
		return err
	}
	err = writer.WriteElementStringU32("FailedAttempts", g.FailedAttempts)
	if err != nil {
		return err
	}
	err = writer.WriteElementStringU32("CumulativeFailedAttempts", g.CumulativeFailedAttempts)
	if err != nil {
		return err
	}
	return nil
}

// PostLoad returns the handle actions after Load.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSCommunicationPortProtection) PostLoad(reader *GXXmlReader) error {
	if g.Port != nil {
		g.Port = reader.Objects.FindByLN(enums.ObjectTypeNone, g.Port.Base().LogicalName())
	}
	return nil
}

// Reset returns the resets failed attempts and current lockout time to zero.
// Protection status is set to unlocked.
//
// Parameters:
//
//	client: DLMS client.
//
// Returns:
//
//	Action bytes.
func (g *GXDLMSCommunicationPortProtection) Reset(client IGXDLMSClient) ([][]byte, error) {
	return client.Method(g, 1, int8(0), enums.DataTypeInt8)
}

// GetValues returns the an array containing the COSEM object's attribute values.
func (g *GXDLMSCommunicationPortProtection) GetValues() []any {
	return []any{g.LogicalName(), g.ProtectionMode, g.AllowedFailedAttempts, g.InitialLockoutTime, g.SteepnessFactor,
		g.MaxLockoutTime, g.Port, g.ProtectionStatus, g.FailedAttempts, g.CumulativeFailedAttempts}
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
func (g *GXDLMSCommunicationPortProtection) GetDataType(index int) (enums.DataType, error) {
	var ret enums.DataType
	switch index {
	case 1:
		ret = enums.DataTypeOctetString
	case 2:
		ret = enums.DataTypeEnum
	case 3:
		ret = enums.DataTypeUint16
	case 4:
		ret = enums.DataTypeUint32
	case 5:
		ret = enums.DataTypeUint8
	case 6:
		ret = enums.DataTypeUint32
	case 7:
		ret = enums.DataTypeOctetString
	case 8:
		ret = enums.DataTypeEnum
	case 9:
		ret = enums.DataTypeUint32
	case 10:
		ret = enums.DataTypeUint32
	default:
		return 0, dlmserrors.ErrInvalidAttributeIndex
	}
	return ret, nil
}

// NewGXDLMSCommunicationPortProtection creates a new communication port protection object instance.
//
// The function validates `ln` before creating the object.
//`ln` is the Logical Name and `sn` is the Short Name of the object.
func NewGXDLMSCommunicationPortProtection(ln string, sn int16) (*GXDLMSCommunicationPortProtection, error) {
	err := ValidateLogicalName(ln)
	if err != nil {
		return nil, err
	}
	return &GXDLMSCommunicationPortProtection{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeCommunicationPortProtection,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
