package dlms

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
	"bytes"
	"errors"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/Gurux/gxcommon-go"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/manufacturersettings"
	"github.com/Gurux/gxdlms-go/objects"
	"github.com/Gurux/gxdlms-go/secure"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// GXDLMSClient implements methods to communicate with DLMS/COSEM metering devices.
// This class provides a complete implementation of the DLMS/COSEM protocol client,
// supporting both Logical Name (LN) and Short Name (SN) referencing modes.
// It handles connection establishment, data exchange, security, and disconnection
// with smart meters and other DLMS-compliant devices.
type GXDLMSClient struct {
	// Manufacturer ID (FLAG ID).
	// Three-character manufacturer identifier used for manufacturer-specific functionality.
	// This ID is standardized and assigned by the DLMS User Association.
	manufacturerID string

	// Initial challenge value that is restored after the connection is closed.
	// Used for High Level Security (HLS) authentication. The challenge is a random
	// value used in the authentication process.
	initializeChallenge []byte

	// Initial PDU size that is restored after the connection is closed.
	// Stores the maximum PDU size negotiated during connection establishment.
	initializePduSize uint16

	// Initial maximum HDLC transmission window size that is restored after the connection is closed.
	// Defines the maximum information field size for transmitting HDLC frames.
	initializeMaxInfoTX uint16

	// Initial maximum HDLC receive window size that is restored after the connection is closed.
	// Defines the maximum information field size for receiving HDLC frames.
	initializeMaxInfoRX uint16

	// Initial maximum HDLC window size in transmission that is restored after the connection is closed.
	// Specifies the number of frames that can be transmitted before an acknowledgment is required.
	initializeWindowSizeTX uint8

	// Initial maximum HDLC window size in receive that is restored after the connection is closed.
	// Specifies the number of frames that can be received before an acknowledgment must be sent.
	initializeWindowSizeRX uint8

	// Gets the DLMS settings object containing all communication parameters.
	settings *settings.GXDLMSSettings

	// Is authentication Required.
	isAuthenticationRequired bool

	// Translator for converting DLMS PDUs to/from XML format.
	translator *GXDLMSTranslator

	// Controls whether XML client throws exceptions or serializes them as default.
	// When set to true, exceptions are thrown normally. When false (default),
	// exceptions are serialized to XML format for XML-based clients.
	// Useful for debugging and protocol analysis.
	throwExceptions bool

	// Gets or sets the collection of custom OBIS codes.
	// This collection is used when reading the Association View from the meter
	// to provide descriptions for objects. Custom OBIS codes can define manufacturer-specific
	// or non-standard objects. If not set, descriptions for custom objects will be empty.
	CustomObisCodes manufacturersettings.GXObisCodeCollection

	// If protected release is used release is including a ciphered xDLMS Initiate request.
	// New DLMS Conformance Tests tests expect protected release.
	// It's not optional anymore.
	UseProtectedRelease bool
}

// Gets the DLMS settings object containing all communication parameters.
func (g *GXDLMSClient) Settings() *settings.GXDLMSSettings {
	return g.settings
}

// ManufacturerID returns the gets or sets the manufacturer ID (FLAG ID).
// The manufacturer ID is used for manufacturer-specific functionality and must be
// exactly 3 characters. This identifier is standardized by the DLMS User Association.
func (g *GXDLMSClient) ManufacturerID() string {
	return g.manufacturerID
}

// SetManufacturerID sets the gets or sets the manufacturer ID (FLAG ID).
// The manufacturer ID is used for manufacturer-specific functionality and must be
// exactly 3 characters. This identifier is standardized by the DLMS User Association.
func (g *GXDLMSClient) SetManufacturerID(value string) error {
	if value != "" && len(value) != 3 {
		return errors.New("Manufacturer ID is 3 chars long string")
	}
	g.manufacturerID = value
	return nil
}

// ClientAddress returns the gets or sets the client address.
// The client address identifies this client to the server (meter).
// The valid range and format depend on the addressing mode:
// - One-byte addressing: 0-127
// - Two-byte addressing: 0-16383
// - Four-byte addressing: logical and physical address combination
// Default value is typically 16 (0x10) for public clients.
func (g *GXDLMSClient) ClientAddress() int {
	return g.settings.ClientAddress
}

// SetClientAddress sets the gets or sets the client address.
// The client address identifies this client to the server (meter).
// The valid range and format depend on the addressing mode:
// - One-byte addressing: 0-127
// - Two-byte addressing: 0-16383
// - Four-byte addressing: logical and physical address combination
// Default value is typically 16 (0x10) for public clients.
func (g *GXDLMSClient) SetClientAddress(value int) error {
	g.settings.ClientAddress = value
	return nil
}

// UseUtc2NormalTime returns the gets or sets a value indicating whether to use UTC to normal time conversion.
// The DLMS standard specifies that time zone is from normal time to UTC in minutes.
// However, some meters (e.g., Italy, Saudi Arabia, India) use UTC time zone instead.
// Set this to true for meters that are configured to use UTC time zone.
// This affects how date/time values are interpreted and converted.
func (g *GXDLMSClient) UseUtc2NormalTime() bool {
	return g.settings.UseUtc2NormalTime
}

// SetUseUtc2NormalTime sets the gets or sets a value indicating whether to use UTC to normal time conversion.
// The DLMS standard specifies that time zone is from normal time to UTC in minutes.
// However, some meters (e.g., Italy, Saudi Arabia, India) use UTC time zone instead.
// Set this to true for meters that are configured to use UTC time zone.
// This affects how date/time values are interpreted and converted.
func (g *GXDLMSClient) SetUseUtc2NormalTime(value bool) error {
	g.settings.UseUtc2NormalTime = value
	return nil
}

// ExpectedInvocationCounter returns the gets or sets the expected invocation (frame) counter value.
// When using ciphered communication, the invocation counter ensures replay protection.
// If this value is set to a non-zero value, any received ciphered PDU with an
// invocation counter lower than this value will be rejected. Setting to 0 disables
// this validation. The counter automatically increments with each secured message.
func (g *GXDLMSClient) ExpectedInvocationCounter() uint64 {
	return g.settings.ExpectedInvocationCounter()
}

// SetExpectedInvocationCounter sets the gets or sets the expected invocation (frame) counter value.
// When using ciphered communication, the invocation counter ensures replay protection.
// If this value is set to a non-zero value, any received ciphered PDU with an
// invocation counter lower than this value will be rejected. Setting to 0 disables
// this validation. The counter automatically increments with each secured message.
func (g *GXDLMSClient) SetExpectedInvocationCounter(value uint64) error {
	g.settings.SetExpectedInvocationCounter(value)
	return nil
}

// DateTimeSkips returns the gets or sets the date/time fields to skip during serialization.
// Some meters cannot handle certain date/time fields such as deviation or status.
// Use this property to skip problematic fields during date/time value encoding.
// Common scenarios include meters that don't support daylight saving time deviation
// or clock status flags.
func (g *GXDLMSClient) DateTimeSkips() enums.DateTimeSkips {
	return g.settings.DateTimeSkips
}

// SetDateTimeSkips sets the gets or sets the date/time fields to skip during serialization.
// Some meters cannot handle certain date/time fields such as deviation or status.
// Use this property to skip problematic fields during date/time value encoding.
// Common scenarios include meters that don't support daylight saving time deviation
// or clock status flags.
func (g *GXDLMSClient) SetDateTimeSkips(value enums.DateTimeSkips) error {
	g.settings.DateTimeSkips = value
	return nil
}

// Standard returns the gets or sets the DLMS standard version being used.
// Different countries and regions may have specific variations or extensions
// to the base DLMS/COSEM standard. This property allows the client to adapt
// to these variations, particularly for date/time handling and specific OBIS codes.
func (g *GXDLMSClient) Standard() enums.Standard {
	return g.settings.Standard
}

// SetStandard sets the gets or sets the DLMS standard version being used.
// Different countries and regions may have specific variations or extensions
// to the base DLMS/COSEM standard. This property allows the client to adapt
// to these variations, particularly for date/time handling and specific OBIS codes.
func (g *GXDLMSClient) SetStandard(value enums.Standard) error {
	g.settings.Standard = value
	return nil
}

// QualityOfService returns the gets or sets the quality of service parameter.
// Quality of Service (QoS) is used in some communication protocols to
// prioritize traffic or specify delivery requirements. The meaning and
// usage depend on the specific protocol and meter implementation.
func (g *GXDLMSClient) QualityOfService() uint8 {
	return g.settings.QualityOfService
}

// SetQualityOfService sets the gets or sets the quality of service parameter.
// Quality of Service (QoS) is used in some communication protocols to
// prioritize traffic or specify delivery requirements. The meaning and
// usage depend on the specific protocol and meter implementation.
func (g *GXDLMSClient) SetQualityOfService(value uint8) error {
	g.settings.QualityOfService = value
	return nil
}

// UserId returns the gets or sets the user identifier.
// The user ID is used when the Association LN object has a user list configured.
// It identifies which user (authentication context) is being used for the connection.
// Value -1 indicates that user ID is not used.
func (g *GXDLMSClient) UserId() int {
	return g.settings.UserID
}

// SetUserId sets the gets or sets the user identifier.
// The user ID is used when the Association LN object has a user list configured.
// It identifies which user (authentication context) is being used for the connection.
// Value -1 indicates that user ID is not used.
func (g *GXDLMSClient) SetUserId(value int) error {
	if value < -1 || value > 255 {
		return errors.New("Invalid user Id.")
	}
	g.settings.UserID = value
	return nil
}

// ServerAddress returns the gets or sets the server (meter) address.
// The server address identifies the target device (meter) in communication.
// The format depends on the addressing mode and can include logical and
// physical address components. Default is typically 1 for single-meter scenarios.
func (g *GXDLMSClient) ServerAddress() int {
	return g.settings.ServerAddress
}

// SetServerAddress sets the gets or sets the server (meter) address.
// The server address identifies the target device (meter) in communication.
// The format depends on the addressing mode and can include logical and
// physical address components. Default is typically 1 for single-meter scenarios.
func (g *GXDLMSClient) SetServerAddress(value int) error {
	g.settings.ServerAddress = value
	return nil
}

// ServerAddressSize returns the gets or sets the size of the server address in bytes.
// This property defines how many bytes are used to encode the server address
// in the HDLC frame. Common values:
// - 1 byte: Simple addressing (0-127)
// - 2 bytes: Extended addressing (0-16383)
// - 4 bytes: Full logical and physical addressing
// The value is typically auto-detected based on the addressing mode.
func (g *GXDLMSClient) ServerAddressSize() uint8 {
	return g.settings.ServerAddressSize
}

// SetServerAddressSize sets the gets or sets the size of the server address in bytes.
// This property defines how many bytes are used to encode the server address
// in the HDLC frame. Common values:
// - 1 byte: Simple addressing (0-127)
// - 2 bytes: Extended addressing (0-16383)
// - 4 bytes: Full logical and physical addressing
// The value is typically auto-detected based on the addressing mode.
func (g *GXDLMSClient) SetServerAddressSize(value uint8) error {
	g.settings.ServerAddressSize = value
	return nil
}

// SourceSystemTitle returns the source system title.
// Meter returns system title when ciphered connection is made or GMAC authentication is used.
func (g *GXDLMSClient) SourceSystemTitle() []byte {
	return g.settings.SourceSystemTitle()
}

// DLMSVersion returns the dLMS version number.
// Gurux DLMS component supports DLMS version number 6.
func (g *GXDLMSClient) DLMSVersion() uint8 {
	return g.settings.DLMSVersion
}

// MaxReceivePDUSize returns the retrieves the maximum size of PDU receiver.
// PDU size tells maximum size of PDU packet.
// Value can be from 0 to 0xFFFF. By default the value is 0xFFFF.
func (g *GXDLMSClient) MaxReceivePDUSize() uint16 {
	return g.settings.MaxPduSize()
}

// SetMaxReceivePDUSize sets the retrieves the maximum size of PDU receiver.
// PDU size tells maximum size of PDU packet.
// Value can be from 0 to 0xFFFF. By default the value is 0xFFFF.
func (g *GXDLMSClient) SetMaxReceivePDUSize(value uint16) error {
	return g.settings.SetMaxPduSize(value)
}

// GbtWindowSize returns the maximum GBT window size.
func (g *GXDLMSClient) GbtWindowSize() uint8 {
	return g.settings.GbtWindowSize()
}

// SetGbtWindowSize sets the maximum GBT window size.
func (g *GXDLMSClient) SetGbtWindowSize(value uint8) error {
	g.settings.SetGbtWindowSize(value)
	return nil
}

// UseLogicalNameReferencing returns the determines, whether Logical, or Short name, referencing is used.
// Referencing depends on the device to communicate with.
// Normally, a device supports only either Logical or Short name referencing.
// The referencing is defined by the device manufacurer.
// If the referencing is wrong, the SNMR message will fail.
func (g *GXDLMSClient) UseLogicalNameReferencing() bool {
	return g.settings.UseLogicalNameReferencing()
}

// SetUseLogicalNameReferencing sets the determines, whether Logical, or Short name, referencing is used.
// Referencing depends on the device to communicate with.
// Normally, a device supports only either Logical or Short name referencing.
// The referencing is defined by the device manufacurer.
// If the referencing is wrong, the SNMR message will fail.
func (g *GXDLMSClient) SetUseLogicalNameReferencing(value bool) error {
	g.settings.SetUseLogicalNameReferencing(value)
	return nil
}

// CtoSChallenge returns the client to Server custom challenge.
// This is for debugging purposes. Reset custom challenge settings CtoSChallenge to nil.
func (g *GXDLMSClient) CtoSChallenge() []byte {
	return g.settings.CtoSChallenge()
}

// SetCtoSChallenge sets the client to Server custom challenge.
// This is for debugging purposes. Reset custom challenge settings CtoSChallenge to nil.
func (g *GXDLMSClient) SetCtoSChallenge(value []byte) error {
	g.settings.UseCustomChallenge = value != nil
	g.settings.SetCtoSChallenge(value)
	return nil
}

// Password returns the retrieves the password that is used in communication.
// If authentication is set to none, password is not used.
// For HighSHA1, HighMD5 and HighGMAC password is worked as a shared secret.
func (g *GXDLMSClient) Password() []byte {
	return g.settings.Password
}

// SetPassword sets the retrieves the password that is used in communication.
// If authentication is set to none, password is not used.
// For HighSHA1, HighMD5 and HighGMAC password is worked as a shared secret.
func (g *GXDLMSClient) SetPassword(value []byte) error {
	g.settings.Password = value
	return nil
}

// ProposedConformance returns the when connection is made client tells what kind of services it want's to use.
func (g *GXDLMSClient) ProposedConformance() enums.Conformance {
	return g.settings.ProposedConformance
}

// SetProposedConformance sets the when connection is made client tells what kind of services it want's to use.
func (g *GXDLMSClient) SetProposedConformance(value enums.Conformance) error {
	g.settings.ProposedConformance = value
	return nil
}

// NegotiatedConformance returns the functionality what server can offer.
func (g *GXDLMSClient) NegotiatedConformance() enums.Conformance {
	return g.settings.NegotiatedConformance
}

// SetNegotiatedConformance sets the functionality what server can offer.
func (g *GXDLMSClient) SetNegotiatedConformance(value enums.Conformance) error {
	g.settings.NegotiatedConformance = value
	return nil
}

// ProtocolVersion returns the protocol version.
func (g *GXDLMSClient) ProtocolVersion() string {
	return g.settings.ProtocolVersion
}

// SetProtocolVersion sets the protocol version.
func (g *GXDLMSClient) SetProtocolVersion(value string) error {
	g.settings.ProtocolVersion = value
	return nil
}

// Authentication returns the retrieves the authentication used in communicating with the device.
// By default authentication is not used. If authentication is used,
// set the password with the Password property.
// Note!
// For HLS authentication password (shared secret) is needed from the manufacturer.
func (g *GXDLMSClient) Authentication() enums.Authentication {
	return g.settings.Authentication
}

// SetAuthentication sets the retrieves the authentication used in communicating with the device.
// By default authentication is not used. If authentication is used,
// set the password with the Password property.
// Note!
// For HLS authentication password (shared secret) is needed from the manufacturer.
func (g *GXDLMSClient) SetAuthentication(value enums.Authentication) error {
	g.settings.Authentication = value
	return nil
}

// ChallengeSize returns the challenge Size.
// Random challenge is used if value is zero.
func (g *GXDLMSClient) ChallengeSize() uint8 {
	return g.settings.ChallengeSize()
}

// SetChallengeSize sets the challenge Size.
// Random challenge is used if value is zero.
func (g *GXDLMSClient) SetChallengeSize(value uint8) error {
	g.settings.SetChallengeSize(value)
	return nil
}

// StartingBlockIndex returns the set starting block index in HDLC framing.
// Default is One based, but some meters use Zero based value.
// Usually this is not used.
func (g *GXDLMSClient) StartingBlockIndex() uint32 {
	return g.settings.StartingBlockIndex
}

// SetStartingBlockIndex sets the set starting block index in HDLC framing.
// Default is One based, but some meters use Zero based value.
// Usually this is not used.
func (g *GXDLMSClient) SetStartingBlockIndex(value uint32) error {
	g.settings.StartingBlockIndex = value
	g.settings.ResetBlockIndex()
	return nil
}

// Priority returns the used priority in HDLC framing.
func (g *GXDLMSClient) Priority() enums.Priority {
	return g.settings.Priority
}

// SetPriority sets the used priority in HDLC framing.
func (g *GXDLMSClient) SetPriority(value enums.Priority) error {
	g.settings.Priority = value
	return nil
}

// ServiceClass returns the used service class in HDLC framing.
func (g *GXDLMSClient) ServiceClass() enums.ServiceClass {
	return g.settings.ServiceClass
}

// SetServiceClass sets the used service class in HDLC framing.
func (g *GXDLMSClient) SetServiceClass(value enums.ServiceClass) error {
	g.settings.ServiceClass = value
	return nil
}

// InvokeID returns the invoke ID.
func (g *GXDLMSClient) InvokeID() byte {
	return g.settings.InvokeID()
}

// SetInvokeID sets the invoke ID.
func (g *GXDLMSClient) SetInvokeID(value byte) error {
	g.settings.SetInvokeID(value)
	return nil
}

// AutoIncreaseInvokeID returns the auto increase Invoke ID.
func (g *GXDLMSClient) AutoIncreaseInvokeID() bool {
	return g.settings.AutoIncreaseInvokeID
}

// SetAutoIncreaseInvokeID sets the auto increase Invoke ID.
func (g *GXDLMSClient) SetAutoIncreaseInvokeID(value bool) error {
	g.settings.AutoIncreaseInvokeID = value
	return nil
}

// InterfaceType returns the determines the type of the connection
// All DLMS meters do not support the IEC 62056-47 standard.
// If the device does not support the standard, and the connection is made
// using TCP/IP, set the type to InterfaceType.General.
func (g *GXDLMSClient) InterfaceType() enums.InterfaceType {
	return g.settings.InterfaceType
}

// SetInterfaceType sets the determines the type of the connection
// All DLMS meters do not support the IEC 62056-47 standard.
// If the device does not support the standard, and the connection is made
// using TCP/IP, set the type to InterfaceType.General.
func (g *GXDLMSClient) SetInterfaceType(value enums.InterfaceType) error {
	g.settings.InterfaceType = value
	return nil
}

// PreEstablishedConnection returns the is pre-established connection used.
// AARQ or release messages are not used with pre-established connections.
func (g *GXDLMSClient) PreEstablishedConnection() bool {
	return g.settings.PreEstablishedSystemTitle != nil
}

// HdlcSettings returns the hDLC connection settings.
func (g *GXDLMSClient) HdlcSettings() *settings.GXHdlcSettings {
	return g.settings.Hdlc
}

// Plc returns the pLC settings.
func (g *GXDLMSClient) Plc() *settings.GXPlcSettings {
	return g.settings.Plc
}

// MBus returns the m-Bus settings.
func (g *GXDLMSClient) MBus() *settings.GXMBusSettings {
	return g.settings.MBus
}

// Coap returns the coAP settings.
func (g *GXDLMSClient) Coap() *settings.GXCoAPSettings {
	return g.settings.Coap
}

// Gateway returns the gateway settings.
func (g *GXDLMSClient) Gateway() *settings.GXDLMSGateway {
	return g.settings.Gateway
}

// SetGateway sets the gateway settings.
func (g *GXDLMSClient) SetGateway(value *settings.GXDLMSGateway) error {
	g.settings.Gateway = value
	return nil
}

// Broacast returns the is data send as a broadcast or unicast.
func (g *GXDLMSClient) Broacast() bool {
	return g.settings.Broadcast
}

// SetBroacast sets the is data send as a broadcast or unicast.
func (g *GXDLMSClient) SetBroacast(value bool) error {
	g.settings.Broadcast = value
	return nil
}

// ConnectionState returns the connection state to the meter.
func (g *GXDLMSClient) ConnectionState() enums.ConnectionState {
	return g.settings.Connected
}

// Is authentication Required.
func (g *GXDLMSClient) IsAuthenticationRequired() bool {
	return g.isAuthenticationRequired
}

// Version returns the the version can be used for backward compatibility.
func (g *GXDLMSClient) Version() int {
	return g.settings.Version
}

// SetVersion sets the the version can be used for backward compatibility.
func (g *GXDLMSClient) SetVersion(value int) error {
	g.settings.Version = value
	return nil
}

// Objects returns the available objects.
func (g *GXDLMSClient) Objects() *objects.GXDLMSObjectCollection {
	ret := g.settings.Objects.(objects.GXDLMSObjectCollection)
	return &ret
}

// OverwriteAttributeAccessRights returns the overwrite attribute access rights if association view tells wrong access rights and they are overwritten.
func (g *GXDLMSClient) OverwriteAttributeAccessRights() bool {
	return g.settings.OverwriteAttributeAccessRights
}

// SetOverwriteAttributeAccessRights sets the overwrite attribute access rights if association view tells wrong access rights and they are overwritten.
func (g *GXDLMSClient) SetOverwriteAttributeAccessRights(value bool) error {
	g.settings.OverwriteAttributeAccessRights = value
	return nil
}

func createDLMSObject(settings *settings.GXDLMSSettings, ClassID uint16, Version any, baseName any, LN any, accessRights any, lnVersion uint8) objects.IGXDLMSBase {
	type_ := enums.ObjectType(ClassID)
	obj := objects.CreateObject(type_)
	if obj != nil {
		updateObjectData(obj, type_, Version, baseName, LN, accessRights.(types.GXStructure), lnVersion)
	}
	return obj
}

func updateObjectData(
	obj objects.IGXDLMSBase,
	objectType enums.ObjectType,
	version any,
	baseName any,
	logicalName any,
	accessRights types.GXStructure,
	lnVersion uint8) {
	var tmp int
	if obj == nil {
		return
	}
	// Some meters return only supported access rights.
	// All access rights are set to NoAccess.
	if lnVersion < 3 {
		for pos := 0; pos != obj.GetAttributeCount(); pos++ {
			obj.Base().SetAccess(pos+1, enums.AccessModeNoAccess)
		}
		for pos := 0; pos != obj.GetMethodCount(); pos++ {
			obj.Base().SetMethodAccess(pos+1, enums.MethodAccessModeNoAccess)
		}
	} else {
		for pos := 0; pos != obj.GetAttributeCount(); pos++ {
			obj.Base().SetAccess3(pos+1, enums.AccessMode3NoAccess)
		}
		for pos := 0; pos != obj.GetMethodCount(); pos++ {
			obj.Base().SetMethodAccess3(pos+1, enums.MethodAccessMode3NoAccess)
		}
	}

	// Check access rights...
	if len(accessRights) == 2 {
		// accessRights[0] = attribute access list
		if attrList, ok := accessRights[0].(types.GXArray); ok {
			for _, item := range attrList {
				attributeAccess, ok := item.(types.GXStructure)
				if !ok || len(attributeAccess) < 2 {
					continue
				}

				id := int(attributeAccess[0].(int8))
				tmp := attributeAccess[1].(types.GXEnum).Value
				// With some meters id is negative.
				if id > 0 {
					if lnVersion < 3 {
						obj.Base().SetAccess(id, enums.AccessMode(tmp))
					} else {
						obj.Base().SetAccess3(id, enums.AccessMode3(tmp))
					}
				}

				// Optional selectors at index 2.
				if len(attributeAccess) > 2 && attributeAccess[2] != nil {
					var value byte = 0
					for _, it := range attributeAccess[2].(types.GXArray) {
						shift := it.(int8)
						if shift >= 0 && shift < 8 {
							value |= (1 << byte(shift))
						}
					}
					obj.Base().SetAccessSelector(id, value)
				}
			}
		}
		// Methods
		methods := accessRights[1].(types.GXArray)
		if len(methods) != 0 {
			for _, item := range methods {
				methodAccess := item.(types.GXStructure)
				id := int(methodAccess[0].(int8))
				// If version is 0 (bool), else version is 1 (int).
				if b, ok := methodAccess[1].(bool); ok {
					if b {
						tmp = 1
					} else {
						tmp = 0
					}
				} else {
					tmp = int(methodAccess[1].(types.GXEnum).Value)
				}
				if id > 0 {
					if lnVersion < 3 {
						obj.Base().SetMethodAccess(id, enums.MethodAccessMode(tmp))
					} else {
						obj.Base().SetMethodAccess3(id, enums.MethodAccessMode3(tmp))
					}
				}
			}
		}
	}
	if baseName != nil {
		obj.Base().ShortName = baseName.(int16)
	}
	if version != nil {
		obj.Base().Version = version.(uint8)
	}
	if logicalName != nil {
		ret, _ := helpers.ToLogicalName(logicalName)
		obj.Base().SetLogicalName(ret)
	}
}

// parseSNObjects returns the reserved for internal use.
func (g *GXDLMSClient) parseSNObjects(buff *types.GXByteBuffer,
	ignoreInactiveObjects bool) (objects.GXDLMSObjectCollection, error) {
	// Get array tag.
	size, err := buff.Uint8()
	if err != nil {
		return nil, err
	}
	// Check that data is in the array
	if size != 0x01 {
		return nil, errors.New("Invalid response.")
	}
	items := objects.GXDLMSObjectCollection{}
	cnt, err := types.GetObjectCount(buff)
	if err != nil {
		return nil, err
	}
	info := internal.GXDataInfo{}
	for objPos := 0; objPos != cnt; objPos++ {
		// Some meters give wrong item count.
		if buff.Position() == buff.Size() {
			break
		}
		ret, err := internal.GetData(g.settings, buff, &info)
		if err != nil {
			return nil, err
		}
		objects := ret.(types.GXArray)
		info.Clear()
		if len(objects) != 4 {
			return nil, errors.New("Invalid structure format.")
		}
		baseName := objects[0].(int8)
		ot := objects[1].(uint16)
		comp := createDLMSObject(g.settings, ot, objects[2], baseName, objects[3], nil, 2)
		if comp != nil {
			if !ignoreInactiveObjects || comp.Base().LogicalName() != "0.0.127.0.0.0" {
				items.Add(comp)
			} else {
				log.Printf("Inactive object : %d %d\n", ot, baseName)
			}
		}
	}
	return items, nil
}

// parseLNObjects returns the reserved for internal use.
func (g *GXDLMSClient) parseLNObjects(buff *types.GXByteBuffer,
	ignoreInactiveObjects bool) (objects.GXDLMSObjectCollection, error) {
	size, err := buff.Uint8()
	if err != nil {
		return nil, err
	}
	// Check that data is in the array.
	if size != 0x01 {
		return nil, errors.New("Invalid response.")
	}
	// get object count
	cnt, err := types.GetObjectCount(buff)
	if err != nil {
		return nil, err
	}
	objectCnt := 0
	items := objects.GXDLMSObjectCollection{}
	info := internal.GXDataInfo{}
	lnVersion := uint8(2)
	// Find LN Version because some meters don't add LN Association the first object.
	pos := buff.Position()
	for buff.Position() != buff.Size() && cnt != objectCnt {
		info.Clear()
		d, err := internal.GetData(g.settings, buff, &info)
		if err != nil {
			return nil, err
		}
		objects := d.(types.GXStructure)
		if len(objects) != 4 {
			return nil, errors.New("Invalid structure format.")
		}
		objectCnt++
		ot := objects[0].(uint16)
		// Get LN association version.
		ln, err := ToLogicalName(objects[2])
		if err != nil {
			return nil, err
		}
		if ot == uint16(enums.ObjectTypeAssociationLogicalName) && ln == "0.0.40.0.0.255" {
			lnVersion = objects[1].(uint8)
			break
		}
	}
	objectCnt = 0
	buff.SetPosition(pos)
	for buff.Position() != buff.Size() && cnt != objectCnt {
		info.Clear()
		ret, err := internal.GetData(g.settings, buff, &info)
		if err != nil {
			return nil, err
		}
		objects := ret.(types.GXStructure)
		if len(objects) != 4 {
			return nil, errors.New("Invalid structure format.")
		}
		objectCnt++
		ot := objects[0].(uint16)
		comp := createDLMSObject(g.settings, ot, objects[1], nil, objects[2], objects[3], lnVersion)
		if comp != nil {
			if !ignoreInactiveObjects || comp.Base().LogicalName() != "0.0.127.0.0.0" {
				items.Add(comp)
			} else {
				log.Printf("Inactive object : %d %s\n", ot, comp.Base().LogicalName())
			}
		} else {
			ln, err := ToLogicalName(objects[2])
			if err != nil {
				return nil, err
			}
			log.Printf("Unknown object : %d %s", ot, ln)

		}
	}
	if cnt != objectCnt {
		log.Printf("Association expect size is %d and actual size is %d\n", cnt, objectCnt)
	}
	return items, nil
}

// Method returns the generate Method (Action) request.
//
// Parameters:
//
//	name: Method object short name or Logical Name.
//	objectType: Object type.
//	index: Method index.
//	value: Additional data.
//	type: Additional data type.
func (g *GXDLMSClient) Method2(name any,
	objectType enums.ObjectType,
	index int,
	value any,
	type_ enums.DataType,
	mode int) ([][]byte, error) {
	if name == nil || (index < 1 && g.Standard() != enums.StandardItaly) {
		return nil, gxcommon.ErrInvalidArgument
	}
	var err error
	g.settings.ResetBlockIndex()
	if type_ == enums.DataTypeNone && value != nil {
		type_, err = internal.GetDLMSDataType(reflect.TypeOf(value))
		if err != nil {
			return nil, err
		}
	}
	attributeDescriptor := types.GXByteBuffer{}
	data := types.GXByteBuffer{}
	err = internal.SetData(g.settings, &data, type_, value)
	if err != nil {
		return nil, err
	}
	if g.UseLogicalNameReferencing() {
		err = attributeDescriptor.SetUint16(uint16(objectType))
		if err != nil {
			return nil, err
		}
		ln, err := LogicalNameToBytes(name.(string))
		if err != nil {
			return nil, err
		}
		err = attributeDescriptor.Set(ln)
		if err != nil {
			return nil, err
		}
		err = attributeDescriptor.SetUint8(uint8(index))
		if err != nil {
			return nil, err
		}
		// Method Invocation Parameters is not used.
		if type_ == enums.DataTypeNone {
			err = attributeDescriptor.SetUint8(0)
			if err != nil {
				return nil, err
			}
		} else {
			err = attributeDescriptor.SetUint8(1)
			if err != nil {
				return nil, err
			}
		}
		p := NewGXDLMSLNParameters(g.settings, 0, enums.CommandMethodRequest, byte(enums.ActionRequestTypeNormal), &attributeDescriptor, &data, 0xff, enums.CommandNone)
		p.AccessMode = mode
		// GBT Window size or streaming is not used with method because there is no information available from theGBT block number and client doesn't know when ACK is expected.
		return getLnMessages(p)
	} else {
		var ind int
		var count int
		getActionInfo(objectType, &ind, &count)
		if index > count {
			return nil, gxcommon.ErrInvalidArgument
		}
		sn := name.(uint16)
		index = (ind + (index-1)*0x8)
		sn = sn + uint16(index)
		err = attributeDescriptor.SetUint16(sn)
		if err != nil {
			return nil, err
		}
		// Add selector.
		if type_ != enums.DataTypeNone {
			err = attributeDescriptor.SetUint8(1)
			if err != nil {
				return nil, err
			}
		}
		return getSnMessages(NewGXDLMSSNParameters(g.settings, enums.CommandWriteRequest, 1, byte(enums.VariableAccessSpecificationVariableName), &attributeDescriptor, &data))
	}
	return nil, nil
}

func (g *GXDLMSClient) Write2(name any, value any, type_ enums.DataType, objectType enums.ObjectType, index int, mode int) ([][]byte, error) {
	g.settings.ResetBlockIndex()
	if type_ == enums.DataTypeNone && value != nil {
		type_, err := internal.GetDLMSDataType(reflect.TypeOf(value))
		if err != nil {
			return nil, err
		}
		if type_ == enums.DataTypeNone {
			return nil, errors.New("Invalid parameter. Unknown value type.")
		}
	}
	attributeDescriptor := types.GXByteBuffer{}
	data := types.GXByteBuffer{}
	var reply [][]byte
	err := internal.SetData(g.settings, &data, type_, value)
	if err != nil {
		return nil, err
	}
	if g.UseLogicalNameReferencing() {
		err = attributeDescriptor.SetUint16(uint16(objectType))
		if err != nil {
			return nil, err
		}
		ln, err := LogicalNameToBytes(name.(string))
		err = attributeDescriptor.Set(ln)
		if err != nil {
			return nil, err
		}
		err = attributeDescriptor.SetUint8(uint8(index))
		if err != nil {
			return nil, err
		}
		err = attributeDescriptor.SetUint8(0)
		if err != nil {
			return nil, err
		}
		p := NewGXDLMSLNParameters(g.settings, 0, enums.CommandSetRequest, byte(enums.SetRequestTypeNormal), &attributeDescriptor, &data, 0xff, enums.CommandNone)
		p.AccessMode = mode
		p.blockIndex = g.settings.BlockIndex
		p.blockNumberAck = g.settings.BlockNumberAck
		p.streaming = false
		reply, err = getLnMessages(p)
		if err != nil {
			return nil, err
		}
	} else {
		// Add name.
		sn := name.(uint16)
		sn = sn + uint16(((index - 1) * 8))
		err = attributeDescriptor.SetUint16(sn)
		if err != nil {
			return nil, err
		}
		err = attributeDescriptor.SetUint8(1)
		if err != nil {
			return nil, err
		}
		p := NewGXDLMSSNParameters(g.settings, enums.CommandWriteRequest, 1, byte(enums.VariableAccessSpecificationVariableName), &attributeDescriptor, &data)
		reply, err = getSnMessages(p)
	}
	return reply, err
}

func (g *GXDLMSClient) Read2(name any, objectType enums.ObjectType, attributeOrdinal int, data *types.GXByteBuffer, mode int) ([][]byte, error) {
	if attributeOrdinal < 0 {
		return nil, gxcommon.ErrInvalidArgument
	}
	g.settings.ResetBlockIndex()
	attributeDescriptor := types.GXByteBuffer{}
	var reply [][]byte
	var err error
	if g.UseLogicalNameReferencing() {
		err = attributeDescriptor.SetUint16(uint16(objectType))
		if err != nil {
			return nil, err
		}
		ln, err := LogicalNameToBytes(name.(string))
		if err != nil {
			return nil, err
		}
		err = attributeDescriptor.Set(ln)
		if err != nil {
			return nil, err
		}
		err = attributeDescriptor.SetUint8(uint8(attributeOrdinal))
		if err != nil {
			return nil, err
		}
		if data == nil || data.Size() == 0 {
			err = attributeDescriptor.SetUint8(0)
			if err != nil {
				return nil, err
			}
		} else {
			err = attributeDescriptor.SetUint8(1)
			if err != nil {
				return nil, err
			}
		}
		p := NewGXDLMSLNParameters(g.settings, 0, enums.CommandGetRequest, byte(enums.GetCommandTypeNormal), &attributeDescriptor, data, 0xff, enums.CommandNone)
		p.AccessMode = mode
		reply, err = getLnMessages(p)
	} else {
		var requestType byte
		sn := name.(uint16)
		sn = sn + uint16(((attributeOrdinal - 1) * 8))
		err = attributeDescriptor.SetUint16(sn)
		if err != nil {
			return nil, err
		}
		// parameterized-access
		if data != nil && data.Size() != 0 {
			requestType = uint8(enums.VariableAccessSpecificationParameterisedAccess)
		} else {
			requestType = uint8(enums.VariableAccessSpecificationVariableName)
		}
		p := NewGXDLMSSNParameters(g.settings, enums.CommandReadRequest, 1, requestType, &attributeDescriptor, data)
		reply, err = getSnMessages(p)
	}
	return reply, err
}

// CopyTo returns the copies all client settings to another client instance.
// This method performs a deep copy of all configuration settings including
// authentication, addresses, security parameters, and object collections.
// Useful when creating multiple clients with similar configurations.
//
// Parameters:
//
//	target: The target client to copy settings to.
func (g *GXDLMSClient) CopyTo(target *GXDLMSClient) {
	g.settings.CopyTo(target.Settings())
}

// SNRMRequest returns the generates SNRM request.
// his method is used to generate send SNRMRequest.
// Before the SNRM request can be generated, at least the following
// properties must be set:
// ClientAddress
// ServerAddress
// Note! According to IEC 62056-47: when communicating using
// TCP/IP, the SNRM request is not send.
//
// Returns:
//
//	SNRM request as byte array.
func (g *GXDLMSClient) SNRMRequest() ([]byte, error) {
	return g.SNRMRequestForce(false)
}

// SNRMRequest returns the generates SNRM request.
// his method is used to generate send SNRMRequest.
// Before the SNRM request can be generated, at least the following
// properties must be set:
// ClientAddress
// ServerAddress
// Note! According to IEC 62056-47: when communicating using
// TCP/IP, the SNRM request is not send.
//
// Parameters:
//
//	forceParameters: Are HDLC parameters forced. Some meters require this.
//
// Returns:
//
//	SNRM request as byte array.
func (g *GXDLMSClient) SNRMRequestForce(forceParameters bool) ([]byte, error) {
	g.settings.Closing = false
	g.initializeMaxInfoTX = g.HdlcSettings().MaxInfoTX()
	g.initializeMaxInfoRX = g.HdlcSettings().MaxInfoRX()
	g.initializeWindowSizeTX = g.HdlcSettings().WindowSizeTX()
	g.initializeWindowSizeRX = g.HdlcSettings().WindowSizeRX()
	g.settings.Connected = enums.ConnectionStateNone
	g.isAuthenticationRequired = false
	g.settings.ResetFrameSequence()
	// SNRM request is not used for all communication channels.
	if g.InterfaceType() == enums.InterfaceTypePlcHdlc {
		return getMacHdlcFrame(g.settings, uint8(enums.CommandSnrm), 0, nil)
	}
	if g.InterfaceType() != enums.InterfaceTypeHDLC && g.InterfaceType() != enums.InterfaceTypeHdlcWithModeE {
		return nil, nil
	}
	data := types.NewGXByteBufferWithCapacity(25)
	err := data.SetUint8(0x81)
	if err != nil {
		return nil, err
	}
	err = data.SetUint8(0x80)
	if err != nil {
		return nil, err
	}
	err = data.SetUint8(0)
	if err != nil {
		return nil, err
	}
	maxInfoTX := g.HdlcSettings().MaxInfoTX()
	maxInfoRX := g.HdlcSettings().MaxInfoRX()
	// If custom HDLC parameters are used.
	if g.InterfaceType() != enums.InterfaceTypePlcHdlc && (forceParameters || defaultMaxInfoTX != maxInfoTX || defaultMaxInfoRX != maxInfoRX ||
		defaultWindowSizeTX != g.HdlcSettings().WindowSizeTX() || defaultWindowSizeRX != g.HdlcSettings().WindowSizeRX()) {
		err = data.SetUint8(uint8(internal.HDLCInfoMaxInfoTX))
		if err != nil {
			return nil, err
		}
		appendHdlcParameter(data, uint16(maxInfoTX))
		err = data.SetUint8(uint8(internal.HDLCInfoMaxInfoRX))
		if err != nil {
			return nil, err
		}
		appendHdlcParameter(data, uint16(maxInfoRX))
		err = data.SetUint8(uint8(internal.HDLCInfoWindowSizeTX))
		if err != nil {
			return nil, err
		}
		err = data.SetUint8(4)
		if err != nil {
			return nil, err
		}
		err = data.SetUint32(uint32(g.HdlcSettings().WindowSizeTX()))
		if err != nil {
			return nil, err
		}
		err = data.SetUint8(uint8(internal.HDLCInfoWindowSizeRX))
		if err != nil {
			return nil, err
		}
		err = data.SetUint8(4)
		if err != nil {
			return nil, err
		}
		err = data.SetUint32(uint32(g.HdlcSettings().WindowSizeRX()))
		if err != nil {
			return nil, err
		}
	}
	// If default HDLC parameters are not used.
	if data.Size() != 3 {
		err = data.SetUint8At(2, uint8((data.Size() - 3)))
		if err != nil {
			return nil, err
		}
	} else {
		data = nil
	}
	return getHdlcFrame(g.settings, uint8(enums.CommandSnrm), data, true)
}

// ParseUAResponse returns the parses UAResponse from byte array.
func (g *GXDLMSClient) ParseUAResponse(data *types.GXByteBuffer) error {
	if g.settings.InterfaceType == enums.InterfaceTypeHDLC ||
		g.settings.InterfaceType == enums.InterfaceTypeHdlcWithModeE ||
		g.settings.InterfaceType == enums.InterfaceTypePlcHdlc {
		ret := parseSnrmUaResponse(data, g.settings)
		if ret != nil {
			return ret
		}
		g.settings.Connected = enums.ConnectionStateHdlc
	}
	return nil
}

// AARQRequest returns the generate AARQ request.
//
// Returns:
//
//	AARQ request as byte array.
func (g *GXDLMSClient) AARQRequest() ([][]byte, error) {
	if g.PreEstablishedConnection() {
		// AARQ is not generate for pre-established connections.
		return nil, nil
	}
	if g.ProposedConformance() == enums.ConformanceNone {
		return nil, errors.New("Invalid conformance.")
	}
	g.settings.Closing = false
	g.initializePduSize = g.MaxReceivePDUSize()
	g.initializeChallenge = g.settings.CtoSChallenge()
	g.settings.NegotiatedConformance = enums.ConformanceNone
	g.settings.ResetBlockIndex()
	g.settings.ServerPublicKeyCertificate = nil
	g.settings.Connected &= ^enums.ConnectionStateDlms
	buff := types.NewGXByteBufferWithCapacity(20)
	err := checkInit(g.settings)
	if err != nil {
		return nil, err
	}
	g.settings.SetStoCChallenge(nil)
	if g.AutoIncreaseInvokeID() {
		g.settings.SetInvokeID(0)
	} else {
		g.settings.SetInvokeID(1)
	}
	g.settings.EphemeralBlockCipherKey = nil
	g.settings.EphemeralBroadcastBlockCipherKey = nil
	g.settings.EphemeralAuthenticationKey = nil
	g.settings.EphemeralKek = nil
	// If High authentication is used.
	if g.settings.Authentication > enums.AuthenticationLow {
		if !g.settings.UseCustomChallenge {
			g.settings.SetCtoSChallenge(settings.GenerateChallenge(g.settings.Authentication, g.settings.ChallengeSize()))
		}
	} else {
		g.settings.SetCtoSChallenge(nil)
	}
	err = generateAarq(g.settings, g.settings.Cipher, nil, buff)
	if err != nil {
		return nil, err
	}
	var reply [][]byte
	if g.UseLogicalNameReferencing() {
		p := NewGXDLMSLNParameters(g.settings, 0, enums.CommandAarq, 0, buff, nil, 0xff, enums.CommandNone)
		reply, err = getLnMessages(p)
	} else {
		reply, err = getSnMessages(NewGXDLMSSNParameters(g.settings, enums.CommandAarq, 0, 0, nil, buff))
	}
	return reply, err
}

// ParseAAREResponse returns the parses the AARE response.
// Parse method will update the following data:
// DLMSVersion
// MaxReceivePDUSize
// UseLogicalNameReferencing
// LNSettings or SNSettings
// LNSettings or SNSettings will be updated, depending on the referencing,
// Logical name or Short name.
//
// Returns:
//
//	The AARE response
func (g *GXDLMSClient) ParseAAREResponse(reply *types.GXByteBuffer) error {
	ret, err := parsePDU(g.settings, g.settings.Cipher, reply, nil)
	if err != nil {
		return err
	}
	g.isAuthenticationRequired = ret == enums.SourceDiagnosticAuthenticationRequired
	if g.isAuthenticationRequired {
		log.Println("Authentication is required.")
	} else {
		g.settings.Connected |= enums.ConnectionStateDlms
	}
	if g.DLMSVersion() != 6 {
		return errors.New("Invalid DLMS version number.")
	}
	return nil
}

// GetApplicationAssociationRequest returns the get challenge request if HLS authentication is used.
func (g *GXDLMSClient) GetApplicationAssociationRequest() ([][]byte, error) {
	return g.GetApplicationAssociationRequestWithLogicalName("")
}

// GetApplicationAssociationRequestWithLogicalName returns the get challenge request if HLS authentication is used.
func (g *GXDLMSClient) GetApplicationAssociationRequestWithLogicalName(ln string) ([][]byte, error) {
	if g.settings.Authentication != enums.AuthenticationHighECDSA &&
		g.settings.Authentication != enums.AuthenticationHighGMAC &&
		len(g.settings.Password) == 0 {
		return nil, errors.New("Password is invalid.")
	}
	g.settings.ResetBlockIndex()
	var err error
	var challenge []byte
	// Count challenge for Landis+Gyr. L+G is using custom way to count the challenge.
	if g.manufacturerID == "LGZ" && g.settings.Authentication == enums.AuthenticationHigh {
		challenge = encryptLandisGyrHighLevelAuthentication(g.settings.Password, g.settings.StoCChallenge())
		if g.UseLogicalNameReferencing() {
			if ln == "" {
				ln = "0.0.40.0.0.255"
			}
			return g.Method2(ln, enums.ObjectTypeAssociationLogicalName, 1, challenge, enums.DataTypeOctetString, 0)
		}
		return g.Method2(0xFA00, enums.ObjectTypeAssociationShortName, 8, challenge, enums.DataTypeOctetString, 0)
	}
	var pw []byte
	if g.settings.Authentication == enums.AuthenticationHighGMAC {
		pw = g.settings.Cipher.SystemTitle()
	} else if g.settings.Authentication == enums.AuthenticationHighSHA256 {
		tmp := types.GXByteBuffer{}
		err = tmp.Set(g.settings.Password)
		if err != nil {
			return nil, err
		}
		err = tmp.Set(g.settings.Cipher.SystemTitle())
		if err != nil {
			return nil, err
		}
		err = tmp.Set(g.settings.SourceSystemTitle())
		if err != nil {
			return nil, err
		}
		err = tmp.Set(g.settings.StoCChallenge())
		if err != nil {
			return nil, err
		}
		err = tmp.Set(g.settings.CtoSChallenge())
		if err != nil {
			return nil, err
		}
		pw = tmp.Array()
	} else if g.settings.Authentication == enums.AuthenticationHighECDSA {
		/*TODO:
		if g.settings.Cipher.SigningKeyPair() == nil {
		Settings.Cipher.SigningKeyPair = types.NewGXKeyValuePair[*types.GXPublicKey, *types.GXPrivateKey](
		 (GXPublicKey)Settings.GetKey(DLMS.Objects.Enums.CertificateType.DigitalSignature,
		 Settings.SourceSystemTitle, false),
		 Settings.Cipher.SigningKeyPair.Value)
		 }
		 Settings.Cipher.SigningKeyPair.Key
		// if (Settings.Cipher.SigningKeyPair.Value == nil)
		// {
		// Settings.Cipher.SigningKeyPair = new KeyValuePair<GXPublicKey, GXPrivateKey>(Settings.Cipher.SigningKeyPair.Key,
		// (GXPrivateKey)Settings.GetKey(DLMS.Objects.Enums.CertificateType.DigitalSignature,
		// Settings.Cipher.SystemTitle, true))
		// }
		// Settings.Cipher.SigningKeyPair.Value
		*/
		tmp := types.GXByteBuffer{}
		err = tmp.Set(g.settings.Cipher.SystemTitle())
		if err != nil {
			return nil, err
		}
		err = tmp.Set(g.settings.SourceSystemTitle())
		if err != nil {
			return nil, err
		}
		err = tmp.Set(g.settings.StoCChallenge())
		if err != nil {
			return nil, err
		}
		err = tmp.Set(g.settings.CtoSChallenge())
		if err != nil {
			return nil, err
		}
		pw = tmp.Array()
	} else {
		pw = g.settings.Password
	}
	challenge, err = settings.Secure(g.settings, g.settings.Cipher, g.settings.Cipher.InvocationCounter(),
		g.settings.StoCChallenge(), pw)
	if err != nil {
		return nil, err
	}
	if g.settings.Cipher != nil {
		g.settings.Cipher.SetInvocationCounter(g.settings.Cipher.InvocationCounter() + 1)
	}
	if g.UseLogicalNameReferencing() {
		if ln == "" {
			ln = "0.0.40.0.0.255"
		}
		return g.Method2(ln, enums.ObjectTypeAssociationLogicalName, 1, challenge, enums.DataTypeOctetString, 0)
	}
	return g.Method2(0xFA00, enums.ObjectTypeAssociationShortName, 8, challenge, enums.DataTypeOctetString, 0)
}

// ParseApplicationAssociationResponse returns the parse server's challenge if HLS authentication is used.
func (g *GXDLMSClient) ParseApplicationAssociationResponse(reply *types.GXByteBuffer) error {
	var err error
	// Landis+Gyr is not returning StoC.
	if g.manufacturerID == "LGZ" && g.settings.Authentication == enums.AuthenticationHigh {
		g.settings.Connected |= enums.ConnectionStateDlms
	} else {
		info := internal.GXDataInfo{}
		equals := false
		value, err := internal.GetData(g.settings, reply, &info)
		if value != nil {
			if g.settings.Authentication == enums.AuthenticationHighECDSA {
				if g.settings.Cipher.SigningKeyPair() == nil {
					return errors.New("SigningKeyPair is empty.")
				}

				tmp2 := types.GXByteBuffer{}
				err = tmp2.Set(g.settings.SourceSystemTitle())
				if err != nil {
					return err
				}
				err = tmp2.Set(g.settings.Cipher.SystemTitle())
				if err != nil {
					return err
				}
				err = tmp2.Set(g.settings.CtoSChallenge())
				if err != nil {
					return err
				}
				err = tmp2.Set(g.settings.StoCChallenge())
				if err != nil {
					return err
				}
				sig, err := types.NewGXEcdsaFromPublicKey(g.settings.Cipher.SigningKeyPair().Key)
				if err != nil {
					return err
				}
				equals, err = sig.Verify(value.([]byte), tmp2.Array())
				if err != nil {
					return err
				}
			} else {
				var secret []byte
				var ic uint32
				if g.settings.Authentication == enums.AuthenticationHighGMAC {
					secret = g.settings.SourceSystemTitle()
					bb := types.NewGXByteBufferWithData(value.([]byte))
					bb.Uint8()
					ic, err = bb.Uint32()
					if err != nil {
						return err
					}
				} else if g.settings.Authentication == enums.AuthenticationHighSHA256 {
					tmp2 := types.GXByteBuffer{}
					err = tmp2.Set(g.settings.Password)
					if err != nil {
						return err
					}
					err = tmp2.Set(g.settings.SourceSystemTitle())
					if err != nil {
						return err
					}
					err = tmp2.Set(g.settings.Cipher.SystemTitle())
					if err != nil {
						return err
					}
					err = tmp2.Set(g.settings.CtoSChallenge())
					if err != nil {
						return err
					}
					err = tmp2.Set(g.settings.StoCChallenge())
					if err != nil {
						return err
					}
					secret = tmp2.Array()
				} else {
					secret = g.settings.Password
				}
				tmp, err := settings.Secure(g.settings, g.settings.Cipher, ic, g.settings.CtoSChallenge(), secret)
				if err != nil {
					return err
				}
				equals = bytes.Compare(value.([]byte), tmp) == 0
			}
			g.settings.Connected |= enums.ConnectionStateDlms
		}
		if !equals {
			g.settings.Connected &= ^enums.ConnectionStateDlms
			return errors.New("Invalid password. Server to Client challenge do not match.")
		}
	}
	return err
}

// ReleaseRequest returns the generates a request to release the connection.
//
// Returns:
//
//	Release request, as byte array.
func (g *GXDLMSClient) ReleaseRequest() ([][]byte, error) {
	return g.ReleaseRequest2(false)
}

// ReleaseRequest2 returns the generates a request to release the connection.
//
// Returns:
//
//	Release request, as byte array.
func (g *GXDLMSClient) ReleaseRequest2(force bool) ([][]byte, error) {
	var err error
	if g.PreEstablishedConnection() {
		// Disconnect message is not used for pre-established connections.
		return nil, nil
	}
	// If connection is not established, there is no need to send DisconnectRequest.
	if !force && (g.settings.Connected&enums.ConnectionStateDlms) == 0 {
		return nil, nil
	}
	buff := types.GXByteBuffer{}
	if !g.UseProtectedRelease || g.settings.Cipher.Security() == enums.SecurityNone {
		err = buff.SetUint8(3)
		if err != nil {
			return nil, err
		}
		err = buff.SetUint8(0x80)
		if err != nil {
			return nil, err
		}
		err = buff.SetUint8(1)
		if err != nil {
			return nil, err
		}
		err = buff.SetUint8(0)
		if err != nil {
			return nil, err
		}
	} else {
		err = buff.SetUint8(0)
		if err != nil {
			return nil, err
		}
		err = buff.SetUint8(0x80)
		if err != nil {
			return nil, err
		}
		err = buff.SetUint8(01)
		if err != nil {
			return nil, err
		}
		err = buff.SetUint8(00)
		if err != nil {
			return nil, err
		}
		g.settings.SetMaxPduSize(g.initializePduSize)
		g.settings.NegotiatedConformance = g.settings.ProposedConformance
		err = GenerateUserInformation(g.settings, g.settings.Cipher, nil, &buff)
		if err != nil {
			return nil, err
		}
		// Increase IC.
		if g.settings.IsCiphered(false) {
			g.settings.Cipher.SetInvocationCounter(g.settings.Cipher.InvocationCounter() + 1)
		}
		err = buff.SetUint8At(0, uint8((buff.Size() - 1)))
		if err != nil {
			return nil, err
		}
	}
	var reply [][]byte
	if g.UseLogicalNameReferencing() {
		p := GXDLMSLNParameters{settings: g.settings, requestType: enums.CommandReleaseRequest, attributeDescriptor: &buff, status: 0xff}
		reply, err = getLnMessages(&p)
	} else {
		reply, err = getSnMessages(NewGXDLMSSNParameters(g.settings, enums.CommandReleaseRequest, 0xFF, 0xFF, nil, &buff))
	}
	if err != nil {
		return nil, err
	}
	g.settings.Connected &= ^enums.ConnectionStateDlms
	g.settings.Closing = true
	g.SetMaxReceivePDUSize(g.initializePduSize)
	g.SetCtoSChallenge(g.initializeChallenge)
	return reply, nil
}

// ParseRelease returns the parses the release response from the meter.
//
// Parameters:
//
//	value: Release request from the meter.
func (g *GXDLMSClient) ParseRelease(value *types.GXByteBuffer) error {
	value.SetPosition(0)
	ret, err := value.Uint8()
	if err != nil {
		return err
	}
	if ret != uint8(enums.CommandReleaseResponse) {
		return errors.New("Invalid release response.")
	}
	len_, err := types.GetObjectCount(value)
	if value.Available() < len_ {
		return errors.New("MemoryError")
	}
	// BerType
	ret, err = value.Uint8()
	if err != nil {
		return err
	}
	if ret != 0x80 {
		return errors.New("Invalid release response.")
	}
	value.Uint8()
	if value.Available() != 0 {
		ret, err := value.Uint8()
		if err != nil {
			return err
		}
		if ret != 0xBE {
			return errors.New("Invalid release response.")
		}
		_, err = parsePDU2(g.settings, g.settings.Cipher, value, nil)
		if err != nil {
			return err
		}

		if g.initializePduSize != g.settings.MaxPduSize() {
			return errors.New("Invalid release response.")
		}
		if g.ProposedConformance() != g.NegotiatedConformance() {
			return errors.New("Invalid release response.")
		}
	}
	return err
}

// DisconnectRequest returns the generates a disconnect request.
//
// Returns:
//
//	Disconnected request, as byte array.
func (g *GXDLMSClient) DisconnectRequest() ([]byte, error) {
	return g.DisconnectRequest2(false)
}

// DisconnectRequest returns the generates a disconnect request.
//
// Returns:
//
//	Disconnected request, as byte array.
func (g *GXDLMSClient) DisconnectRequest2(force bool) ([]byte, error) {
	if !force && g.settings.Connected == enums.ConnectionStateNone {
		return nil, nil
	}
	var err error
	var ret []byte
	if useHdlc(g.settings.InterfaceType) {
		if g.settings.InterfaceType == enums.InterfaceTypePlcHdlc {
			ret, err = getMacHdlcFrame(g.settings, uint8(enums.CommandDisconnectRequest), 0, nil)
		} else {
			ret, err = getHdlcFrame(g.settings, uint8(enums.CommandDisconnectRequest), nil, true)
		}
	} else if force || g.settings.Connected == enums.ConnectionStateDlms {
		ret2, err := g.ReleaseRequest2(force)
		if err != nil {
			return nil, err
		}
		ret = ret2[0]
	}
	if useHdlc(g.settings.InterfaceType) {
		g.HdlcSettings().SetMaxInfoTX(g.initializeMaxInfoTX)
		g.HdlcSettings().SetMaxInfoRX(g.initializeMaxInfoRX)
		g.HdlcSettings().SetWindowSizeTX(g.initializeWindowSizeTX)
		g.HdlcSettings().SetWindowSizeRX(g.initializeWindowSizeRX)
	}
	g.SetMaxReceivePDUSize(g.initializePduSize)
	g.settings.Connected = enums.ConnectionStateNone
	g.settings.ResetFrameSequence()
	g.settings.Closing = true
	return ret, err
}

// ParseObjects returns the parses the COSEM objects of the received data.
//
// Parameters:
//
//	data: Received data, from the device, as byte array.
//	onlyKnownObjects: Parse only DLMS standard objects. Manufacture specific objects are ignored.
//	ignoreInactiveObjects: Inactive objects are ignored.
//
// Returns:
//
//	Collection of COSEM objects.
func (g *GXDLMSClient) ParseObjects(data *types.GXByteBuffer,
	ignoreInactiveObjects bool) (objects.GXDLMSObjectCollection, error) {
	if data == nil || data.Size() == 0 {
		return nil, errors.New("ParseObjects failed. Invalid parameter.")
	}
	var err error
	var objects_ objects.GXDLMSObjectCollection
	if g.UseLogicalNameReferencing() {
		objects_, err = g.parseLNObjects(data, ignoreInactiveObjects)
	} else {
		objects_, err = g.parseSNObjects(data, ignoreInactiveObjects)
	}
	if err != nil {
		return nil, err
	}

	if g.CustomObisCodes != nil {
		for _, it := range g.CustomObisCodes {
			if it.Append {
				obj := objects.CreateObject(it.ObjectType)
				obj.Base().Version = it.Version
				obj.Base().SetLogicalName(it.LogicalName)
				obj.Base().Description = it.Description
				objects_ = append(objects_, obj)
			}
		}
	}
	c := NewGXDLMSConverter(g.Standard())
	err = c.UpdateOBISCodesInformation(objects_)
	if err != nil {
		return nil, err
	}
	g.settings.Objects = objects_
	return objects_, nil
}

// ParsePushObjects returns the collection of push objects.
//
// Parameters:
//
//	data: Received data.
//
// Returns:
//
//	Array of objects and called indexes.
func (g *GXDLMSClient) ParsePushObjects(data []any) ([]*types.GXKeyValuePair[objects.IGXDLMSBase, int], error) {
	objects_ := []*types.GXKeyValuePair[objects.IGXDLMSBase, int]{}
	if data != nil {
		c := NewGXDLMSConverter(g.Standard())
		for _, item := range data {
			it := item.([]any)
			classID := it[0].(uint16)
			if classID > 0 {
				ret, err := ToLogicalName(it[1].([]byte))
				if err != nil {
					return nil, err
				}
				comp := g.Objects().FindByLN(enums.ObjectType(classID), ret)
				if comp == nil {
					comp = createDLMSObject(g.settings, classID, 0, 0, it[1], nil, 2)
					c.UpdateOBISCodeInformation(comp)
				}
				if comp != nil {
					objects_ = append(objects_, types.NewGXKeyValuePair(comp, it[2].(int)))
				} else {
					ret, err := helpers.ToLogicalName(it[1].([]byte))
					if err != nil {
						return nil, err
					}
					log.Println("Unknown object : %d %d", classID, ret)
				}
			}
		}
	}
	return objects_, nil
}

// UpdateValue returns the get Value from byte array received from the meter.
func (g *GXDLMSClient) UpdateValue(target objects.IGXDLMSBase,
	attributeIndex int,
	value any,
	columns *[]types.GXKeyValuePair[objects.IGXDLMSBase, *objects.GXDLMSCaptureObject]) (any, error) {
	// Update data type if value is readable.
	ret, err := target.GetDataType(attributeIndex)
	if err != nil {
		return nil, err
	}
	if value != nil && ret == enums.DataTypeNone {

		type_, err := internal.GetDLMSDataType(reflect.TypeOf(value))
		if err == nil {
			target.Base().SetDataType(attributeIndex, type_)
		}
		//It's ok if this fails.
	}
	if v, ok := value.([]byte); ok {
		type_ := target.GetUIDataType(attributeIndex)
		if type_ != enums.DataTypeNone {
			if type_ == enums.DataTypeDate && len(v) == 5 {
				type_ = enums.DataTypeDate
				target.Base().SetUIDataType(attributeIndex, type_)
			}
			value, err = ChangeTypeFromByteArray(v, type_, g.UseUtc2NormalTime())
			if err != nil {
				return nil, err
			}
		}
	}
	e := internal.NewValueEventArgs(g.settings, target, byte(attributeIndex))
	e.Parameters = columns
	e.Value = value
	err = target.SetValue(g.settings, e)
	if err != nil {
		return nil, err
	}
	return target.GetValues()[attributeIndex-1], nil
}

// UpdateValues updates the list values.
//
// Parameters:
//
//	list: COSEM objects to update.
//	values: Received values.
func (g *GXDLMSClient) UpdateValues(list []types.GXKeyValuePair[objects.IGXDLMSBase, int], values []any) error {
	var err error
	pos := 0
	for _, it := range list {
		e := internal.NewValueEventArgs(g.settings, it.Key, byte(it.Value))
		e.Value = values[pos]
		var type_ enums.DataType
		if v, ok := e.Value.([]byte); ok {
			type_ = it.Key.GetUIDataType(it.Value)
			if type_ != enums.DataTypeNone {
				e.Value, err = ChangeTypeFromByteArray(v, type_, g.UseUtc2NormalTime())
				if err != nil {
					return err
				}
			}
		}
		err := it.Key.SetValue(g.settings, e)
		if err != nil {
			return err
		}
		pos++
	}
	return nil
}

// ParseAccessResponse returns the parse access response.
//
// Parameters:
//
//	list: Collection of access items.
//	data: Received data from the meter.
func (g *GXDLMSClient) ParseAccessResponse(list []GXDLMSAccessItem, data *types.GXByteBuffer) error {
	// Get count
	info := internal.GXDataInfo{}
	cnt, err := types.GetObjectCount(data)
	if err != nil {
		return err
	}
	if cnt != len(list) {
		return errors.New("List size and values size do not match.")
	}
	for _, it := range list {
		info.Clear()
		ret, err := internal.GetData(g.settings, data, &info)
		if err != nil {
			return err
		}
		it.Value = ret
	}
	cnt, err = types.GetObjectCount(data)
	if err != nil {
		return err
	}
	if len(list) != cnt {
		return errors.New("List size and values size do not match.")
	}
	for _, it := range list {
		data.Uint8()
		ret, err := data.Uint8()
		if err != nil {
			return err
		}
		it.Error = enums.ErrorCode(ret)
		if it.Command == enums.AccessServiceCommandTypeGet && it.Error == enums.ErrorCodeOk {
			_, err = g.UpdateValue(it.Target, int(it.Index), it.Value, nil)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (g *GXDLMSClient) GetAttributeInfo(item objects.IGXDLMSBase, index int) *manufacturersettings.GXDLMSAttributeSettings {
	return item.Base().Attributes.Find(index)
}

// ChangeTypeFromByteArray returns the changes byte array received from the meter to given type.
//
// Parameters:
//
//	value: Byte array received from the meter.
//	type: Wanted type.
//	useUtc: Standard says that Time zone is from normal time to UTC in minutes.
//	  If meter is configured to use UTC time (UTC to normal time) set this to true.
//
// Returns:
//
//	Value changed by type.
func ChangeTypeFromByteArray(value []byte, type_ enums.DataType, useUtc bool) (any, error) {
	return ChangeType(types.NewGXByteBufferWithData(value), type_, useUtc)
}

// ChangeTypeFromByteBuffer returns the changes byte array received from the meter to given type.
//
// Parameters:
//
//	value: Byte array received from the meter.
//	type: Wanted type.
//
// Returns:
//
//	Value changed by type.
func ChangeTypeFromByteBuffer(value *types.GXByteBuffer, type_ enums.DataType) (any, error) {
	return ChangeType(value, type_, false)
}

// ChangeType returns the changes byte array received from the meter to given type.
//
// Parameters:
//
//	value: Byte array received from the meter.
//	type: Wanted type.
//	useUtc: Standard says that Time zone is from normal time to UTC in minutes.
//	  If meter is configured to use UTC time (UTC to normal time) set this to true.
//
// Returns:
//
//	Value changed by type.
func ChangeType(value *types.GXByteBuffer, type_ enums.DataType, useUtc bool) (any, error) {
	settings := settings.GXDLMSSettings{}
	settings.UseUtc2NormalTime = useUtc
	return internal.ChangeType(&settings, value, type_)
}

// GetObjectsRequest returns the reads the selected object from the device.
// This method is used to get all registers in the device.
//
// Returns:
//
//	Read request, as byte array.
func (g *GXDLMSClient) GetObjectsRequest() ([][]byte, error) {
	return g.GetObjectsRequest2("")
}

// GetObjectsRequest returns the reads the selected object from the device.
// This method is used to get all registers in the device.
//
// Returns:
//
//	Read request, as byte array.
func (g *GXDLMSClient) GetObjectsRequest2(ln string) ([][]byte, error) {
	var name any
	g.settings.ResetBlockIndex()
	if g.UseLogicalNameReferencing() {
		if ln == "" {
			name = "0.0.40.0.0.255"
		} else {
			name = ln
		}
	} else {
		name = int16(-1536)
	}
	return g.Read2(name, enums.ObjectTypeAssociationLogicalName, 2, nil, 0)
}

// Method returns the generate Method (Action) request.
//
// Parameters:
//
//	item: object to write.
//	index: Attribute index where data is write.
func (g *GXDLMSClient) Method(item *objects.GXDLMSObject, index int, data any) ([][]byte, error) {
	return g.Method2(item.Name(), item.ObjectType(), index, data, enums.DataTypeNone, 0)
}

// Write returns the generates a write message.
//
// Parameters:
//
//	item: object to write.
//	index: Attribute index where data is write.
func (g *GXDLMSClient) Write(item objects.IGXDLMSBase, index int) ([][]byte, error) {
	if item == nil || index < 1 {
		return nil, gxcommon.ErrInvalidArgument
	}
	a := internal.NewValueEventArgs(g.settings, item, uint8(index))
	value, err := item.GetValue(g.settings, a)
	if err != nil {
		return nil, err
	}
	type_, err := item.GetDataType(index)
	if err != nil {
		return nil, err
	}
	if type_ == enums.DataTypeNone {
		type_, err = internal.GetDLMSDataType(reflect.TypeOf(value))
		if err != nil {
			return nil, err
		}
	}
	// If values is show as string, but send as byte array.
	if _, ok := value.(string); ok && type_ == enums.DataTypeOctetString {
		tp := item.Base().GetUIDataType(index)
		if tp == enums.DataTypeString {
			value = []byte(value.(string))
		}
	}
	mode := int(item.Base().GetAccess3(index))
	return g.Write2(item.Base().Name(), value, type_, item.Base().ObjectType(), index, mode)
}

// Read returns the generates a read message.
//
// Parameters:
//
//	item: DLMS object to read.
//	attributeOrdinal: Read attribute index.
//
// Returns:
//
//	Read request as byte array.
func (g *GXDLMSClient) Read(item objects.IGXDLMSBase, attributeOrdinal int) ([][]byte, error) {
	mode := int(item.Base().GetAccess3(attributeOrdinal))
	return g.Read2(item.Base().Name(), item.Base().ObjectType(), attributeOrdinal, nil, mode)
}

// ReadList returns the read list of COSEM objects.
//
// Parameters:
//
//	list: List of COSEM object and attribute index to read.
//
// Returns:
//
//	Read List request as byte array.
func (g *GXDLMSClient) ReadList(list []types.GXKeyValuePair[objects.IGXDLMSBase, int]) ([][]byte, error) {
	if (g.NegotiatedConformance() & enums.ConformanceMultipleReferences) == 0 {
		return nil, errors.New("Meter doesn't support multiple objects reading with one request.")
	}
	if len(list) == 0 {
		return nil, gxcommon.ErrInvalidArgument
	}
	var err error
	g.settings.ResetBlockIndex()
	messages := [][]byte{}
	data := types.GXByteBuffer{}
	if g.UseLogicalNameReferencing() {
		// Find highest access mode.
		mode := 0
		for _, it := range list {
			m := int(it.Key.Base().GetAccess3(it.Value))
			if m > mode {
				mode = m
			}
		}
		p := NewGXDLMSLNParameters(g.settings, 0, enums.CommandGetRequest, byte(enums.GetCommandTypeWithList), &data,
			nil, 0xff, enums.CommandNone)

		p.AccessMode = mode
		// Request service primitive shall always fit in a single APDU.
		pos := 0
		count := int((g.settings.MaxPduSize() - 12) / 10)
		if len(list) < count {
			count = len(list)
		}
		// All meters can handle 10 items.
		if count > 10 {
			count = 10
		}
		err = types.SetObjectCount(count, &data)
		if err != nil {
			return nil, err
		}
		for _, it := range list {
			err := data.SetUint16(uint16(it.Key.Base().ObjectType()))

			ln, err := LogicalNameToBytes(it.Key.Base().LogicalName())
			if err != nil {
				return nil, err
			}
			data.Set(ln)
			// Attribute ID.
			err = data.SetUint8(uint8(it.Value))
			if err != nil {
				return nil, err
			}
			// Attribute selector is not used.
			err = data.SetUint8(0)
			if err != nil {
				return nil, err
			}
			pos++
			if pos%count == 0 && len(list) != pos {
				ret, err := getLnMessages(p)
				if err != nil {
					return nil, err
				}
				messages = append(messages, ret...)
				data.Clear()
				if len(list)-pos < count {
					types.SetObjectCount(len(list)-pos, &data)
				} else {
					types.SetObjectCount(count, &data)
				}
			}
		}
		ret, err := getLnMessages(p)
		if err != nil {
			return nil, err
		}
		messages = append(messages, ret...)
	} else {
		if len(list) == 1 {
			return g.Read(list[0].Key, list[0].Value)
		}
		p := GXDLMSSNParameters{Settings: g.settings, Command: enums.CommandReadRequest, Count: len(list), RequestType: 0xFF, AttributeDescriptor: &data}
		for _, it := range list {
			err = data.SetUint8(byte(enums.VariableAccessSpecificationBlockNumberAccess))
			if err != nil {
				return nil, err
			}
			sn := it.Key.Base().ShortName
			sn += int16((it.Value - 1) * 8)
			err = data.SetUint16(uint16(sn))
			if err != nil {
				return nil, err
			}
			if data.Size() >= int(g.settings.MaxPduSize()) {
				ret, err := getSnMessages(&p)
				if err != nil {
					return nil, err
				}
				messages = append(messages, ret...)
				data.Clear()
			}
		}
		ret, err := getSnMessages(&p)
		if err != nil {
			return nil, err
		}
		messages = append(messages, ret...)
	}
	return messages, nil
}

// WriteList returns the write list of COSEM objects.
//
// Parameters:
//
//	list: List of COSEM object and attribute index to read.
//
// Returns:
//
//	Write List request as byte array.
func (g *GXDLMSClient) WriteList(list []types.GXKeyValuePair[objects.IGXDLMSBase, int]) ([][]byte, error) {
	var err error
	if (int(g.NegotiatedConformance()) & int(enums.ConformanceMultipleReferences)) == 0 {
		return nil, errors.New("Meter doesn't support multiple objects writing with one request.")
	}
	if len(list) == 0 {
		return nil, gxcommon.ErrInvalidArgument
	}
	// Find highest access mode.
	mode := 0
	for _, it := range list {
		m := int(it.Key.Base().GetAccess3(it.Value))
		if m > mode {
			mode = m
		}
	}
	g.settings.ResetBlockIndex()
	var messages [][]byte
	data := types.GXByteBuffer{}
	if g.UseLogicalNameReferencing() {
		p := NewGXDLMSLNParameters(g.settings, 0, enums.CommandSetRequest, byte(internal.SetCommandTypeWithList),
			nil, &data, 0xff, enums.CommandNone)
		p.AccessMode = mode
		err = types.SetObjectCount(len(list), &data)
		if err != nil {
			return nil, err
		}
		for _, it := range list {
			err := data.SetUint16(uint16(it.Key.Base().ObjectType()))
			if err != nil {
				return nil, err
			}
			ln, err := LogicalNameToBytes(it.Key.Base().LogicalName())
			if err != nil {
				return nil, err
			}
			data.Set(ln)
			err = data.SetUint8(uint8(it.Value))
			if err != nil {
				return nil, err
			}
			err = data.SetUint8(0)
			if err != nil {
				return nil, err
			}
		}
		err = types.SetObjectCount(len(list), &data)
		if err != nil {
			return nil, err
		}
		for _, it := range list {
			a := internal.NewValueEventArgs(g.settings, it.Key, uint8(it.Value))
			value, err := it.Key.GetValue(g.settings, a)
			if err != nil {
				return nil, err
			}
			type_, err := it.Key.GetDataType(it.Value)
			if err != nil {
				return nil, err
			}
			if type_ == enums.DataTypeNone {
				type_, err = internal.GetDLMSDataType(reflect.TypeOf(value))
				if err != nil {
					return nil, err
				}
			}
			// If values is show as string, but send as byte array.
			if v, ok := value.(string); ok {
				tp := it.Key.GetUIDataType(it.Value)
				if tp == enums.DataTypeString {
					value = []byte(v)
				}
			}
			err = internal.SetData(g.settings, &data, type_, value)
		}
		if err != nil {
			return nil, err
		}
		ret, err := getLnMessages(p)
		if err != nil {
			return nil, err
		}
		messages = append(messages, ret...)
	} else {
		p := NewGXDLMSSNParameters(g.settings, enums.CommandWriteRequest, len(list), 0xFF, nil, &data)
		for _, it := range list {
			err = data.SetUint8(byte(enums.VariableAccessSpecificationBlockNumberAccess))
			if err != nil {
				return nil, err
			}
			sn := it.Key.Base().ShortName
			sn += int16((it.Value - 1) * 8)
			err = data.SetUint16(uint16(sn))
			if err != nil {
				return nil, err
			}
		}
		// Add length.
		err = types.SetObjectCount(len(list), &data)
		if err != nil {
			return nil, err
		}
		p.Count = len(list)
		for _, it := range list {
			value, err := it.Key.GetValue(g.settings, internal.NewValueEventArgs(g.settings, it.Key, uint8(it.Value)))
			if err != nil {
				return nil, err
			}
			type_, err := it.Key.GetDataType(it.Value)
			if err != nil {
				return nil, err
			}
			if type_ == enums.DataTypeNone {
				type_, err = internal.GetDLMSDataType(reflect.TypeOf(value))
				if err != nil {
					return nil, err
				}
			}
			// If values is show as string, but send as byte array.
			if v, ok := value.(string); ok {
				tp := it.Key.GetUIDataType(it.Value)
				if tp == enums.DataTypeString {
					value = []byte(v)
				}
			}
			err = internal.SetData(g.settings, &data, type_, value)
			if err != nil {
				return nil, err
			}
		}
		ret, err := getSnMessages(p)
		if err != nil {
			return nil, err
		}
		messages = append(messages, ret...)
	}
	return messages, nil
}

// GetKeepAlive returns the generates the keep alive message.
// Keepalive message is needed to keep the connection up. Connection is closed if keepalive is not sent in meter's inactivity timeout period.
//
// Returns:
//
//	Returns Keep alive message, as byte array.
func (g *GXDLMSClient) GetKeepAlive() ([]byte, error) {
	g.settings.ResetBlockIndex()
	if g.UseLogicalNameReferencing() {
		ln, err := objects.NewGXDLMSAssociationLogicalName("0.0.40.0.0.255", 0)
		if err != nil {
			return nil, err
		}
		ret, err := g.Read(ln, 1)
		if err != nil {
			return nil, err
		}
		return ret[0], nil
	}
	sn, err := objects.NewGXDLMSAssociationShortName("0.0.40.0.0.255", -1536)
	if err != nil {
		return nil, err
	}
	ret, err := g.Read(sn, 1)
	if err != nil {
		return nil, err
	}
	return ret[0], nil
}

// ReadRowsByEntry returns the read rows by entry.
// Check Conformance because all meters do not support this.
//
// Parameters:
//
//	pg: Profile generic object to read.
//	index: One based start index.
//	count: Rows count to read.
//	columns: Columns to read.
//
// Returns:
//
//	Read message as byte array.
func (g *GXDLMSClient) ReadRowsByEntryWithColumns(pg *objects.GXDLMSProfileGeneric,
	index uint32,
	count uint32,
	columns []types.GXKeyValuePair[objects.IGXDLMSBase, objects.GXDLMSCaptureObject]) ([][]byte, error) {
	columnIndex := 1
	columnEnd := 0
	pos := 0
	// If columns are given find indexes.
	if columns != nil && len(columns) != 0 {
		if len(pg.CaptureObjects) == 0 {
			return nil, errors.New("Read capture objects first.")
		}
		columnIndex = len(pg.CaptureObjects)
		columnEnd = 1
		for _, c := range columns {
			pos = 0
			found := false
			for _, it := range pg.CaptureObjects {
				pos++
				if it.Key.Base().ObjectType() == c.Key.Base().ObjectType() &&
					it.Key.Base().LogicalName() == c.Key.Base().LogicalName() &&
					it.Value.AttributeIndex == c.Value.AttributeIndex &&
					it.Value.DataIndex == c.Value.DataIndex {
					found = true
					if pos < columnIndex {
						columnIndex = pos
					}
					columnEnd = pos
					break
				}
			}
			if !found {
				return nil, fmt.Errorf("Invalid column: %s", c.Key.Base().LogicalName())
			}
		}
	}
	return g.ReadRowsByEntryWithColumns2(pg, index, count, columnIndex, columnEnd)
}

// ReadRowsByEntry returns the read rows by entry.
// Check Conformance because all meters do not support this.
//
// Parameters:
//
//	pg: Profile generic object to read.
//	index: One based start index.
//	count: Rows count to read.
//
// Returns:
//
//	Read message as byte array.
func (g *GXDLMSClient) ReadRowsByEntry(pg *objects.GXDLMSProfileGeneric, index uint32, count uint32) ([][]byte, error) {
	return g.ReadRowsByEntryWithColumns(pg, index, count, nil)
}

// ReadRowsByEntry returns the read rows by entry.
// Check Conformance because all meters do not support this.
//
// Parameters:
//
//	pg: Profile generic object to read.
//	index: One based start index.
//	count: Rows count to read.
//	columnStart: One based column start index.
//	columnEnd: Column end index.
//
// Returns:
//
//	Read message as byte array.
func (g *GXDLMSClient) ReadRowsByEntryWithColumns2(pg *objects.GXDLMSProfileGeneric,
	index uint32,
	count uint32,
	columnStart int,
	columnEnd int) ([][]byte, error) {
	var err error
	if index < 0 {
		return nil, errors.New("index")
	}
	if count < 0 {
		return nil, errors.New("count")
	}
	if columnStart < 1 {
		return nil, errors.New("columnStart")
	}
	if columnEnd < 0 {
		return nil, errors.New("columnEnd")
	}
	if len(pg.CaptureObjects) == 0 {
		return nil, errors.New("Capture objects not read.")
	}
	pg.Buffer = pg.Buffer[:0]
	g.settings.ResetBlockIndex()
	buff := types.NewGXByteBufferWithCapacity(19)
	err = buff.SetUint8(0x02)
	if err != nil {
		return nil, err
	}
	err = buff.SetUint8(uint8(enums.DataTypeStructure))
	if err != nil {
		return nil, err
	}
	err = buff.SetUint8(0x04)
	if err != nil {
		return nil, err
	}
	internal.SetData(g.settings, buff, enums.DataTypeUint32, index)
	// Add Count
	if count == 0 {
		err = internal.SetData(g.settings, buff, enums.DataTypeUint32, count)
	} else {
		err = internal.SetData(g.settings, buff, enums.DataTypeUint32, index+count-1)
	}
	if err != nil {
		return nil, err
	}
	err = internal.SetData(g.settings, buff, enums.DataTypeUint16, columnStart)
	if err != nil {
		return nil, err
	}
	err = internal.SetData(g.settings, buff, enums.DataTypeUint16, columnEnd)
	if err != nil {
		return nil, err
	}
	mode := int(pg.GetAccess3(2))
	return g.Read2(pg.Name(), enums.ObjectTypeProfileGeneric, 2, buff, mode)
}

// ReadRowsByRange returns the read rows by range.
// Use this method to read Profile Generic table between dates.
// Check Conformance because all meters do not support this.
// Some meters return error if there are no data betweens start and end time.
//
// Parameters:
//
//	pg: Profile generic object to read.
//	start: Start time.
//	end: End time.
//
// Returns:
//
//	Read message as byte array.
func (g *GXDLMSClient) ReadRowsByRange2(pg *objects.GXDLMSProfileGeneric, start time.Time, end time.Time) ([][]byte, error) {

	s := types.NewGXDateTimeFromTime(start)
	e := types.NewGXDateTimeFromTime(end)
	return g.ReadRowsByRangeWithColumns(pg, s, e, nil)
}

// ReadRowsByRange returns the read rows by range.
// Use this method to read Profile Generic table between dates.
// Check Conformance because all meters do not support this.
// Some meters return error if there are no data betweens start and end time.
//
// Parameters:
//
//	pg: Profile generic object to read.
//	start: Start time.
//	end: End time.
//
// Returns:
//
//	Read message as byte array.
func (g *GXDLMSClient) ReadRowsByRange(pg *objects.GXDLMSProfileGeneric, start *types.GXDateTime, end *types.GXDateTime) ([][]byte, error) {
	return g.ReadRowsByRangeWithColumns(pg, start, end, nil)
}

// ReadRowsByRangeWithColumns returns the read rows by range.
// Use this method to read Profile Generic table between dates.
// Check Conformance because all meters do not support this.
// Some meters return error if there are no data betweens start and end time.
//
// Parameters:
//
//	pg: Profile generic object to read.
//	start: Start time.
//	end: End time.
//	columns: Columns to read.
//
// Returns:
//
//	Read message as byte array.
func (g *GXDLMSClient) ReadRowsByRangeWithColumns(pg *objects.GXDLMSProfileGeneric,
	start *types.GXDateTime,
	end *types.GXDateTime,
	columns *[]types.GXKeyValuePair[objects.IGXDLMSBase, objects.GXDLMSCaptureObject]) ([][]byte, error) {
	if len(pg.CaptureObjects) == 0 {
		return nil, errors.New("Capture objects not read.")
	}
	var err error
	pg.Buffer = pg.Buffer[:0]
	g.settings.ResetBlockIndex()
	sort := pg.SortObject
	if sort == nil {
		sort = pg.CaptureObjects[0].Key
	}
	ln := "0.0.1.0.0.255"
	type_ := enums.ObjectTypeClock
	clockType := ClockTypeClock
	// If Unix time is used.
	if _, ok := sort.(*objects.GXDLMSData); ok {
		clockType = ClockTypeUnix
		ln = "0.0.1.1.0.255"
		type_ = enums.ObjectTypeData
	} else if _, ok := sort.(*objects.GXDLMSData); ok {
		//If high resolution time is used.
		clockType = ClockTypeHighResolution
		ln = "0.0.1.2.0.255"
		type_ = enums.ObjectTypeData
	}
	buff := types.NewGXByteBuffer()
	err = buff.SetUint8(0x01)
	if err != nil {
		return nil, err
	}
	err = buff.SetUint8(uint8(enums.DataTypeStructure))
	if err != nil {
		return nil, err
	}
	err = buff.SetUint8(0x04)
	if err != nil {
		return nil, err
	}
	err = buff.SetUint8(uint8(enums.DataTypeStructure))
	if err != nil {
		return nil, err
	}
	err = buff.SetUint8(0x04)
	if err != nil {
		return nil, err
	}
	err = internal.SetData(g.settings, buff, enums.DataTypeUint16, type_)
	if err != nil {
		return nil, err
	}
	ln2, err := LogicalNameToBytes(ln)
	if err != nil {
		return nil, err
	}
	err = internal.SetData(g.settings, buff, enums.DataTypeOctetString, ln2)
	if err != nil {
		return nil, err
	}
	err = internal.SetData(g.settings, buff, enums.DataTypeInt8, 2)
	if err != nil {
		return nil, err
	}
	err = internal.SetData(g.settings, buff, enums.DataTypeUint16, 0)
	if clockType == ClockTypeClock {
		err = internal.SetData(g.settings, buff, enums.DataTypeOctetString, start)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(g.settings, buff, enums.DataTypeOctetString, end)
	} else if clockType == ClockTypeUnix {
		err = internal.SetData(g.settings, buff, enums.DataTypeUint32, start.ToUnixTime())
		if err != nil {
			return nil, err
		}
		err = internal.SetData(g.settings, buff, enums.DataTypeUint32, end.ToUnixTime())
	} else if clockType == ClockTypeHighResolution {
		err = internal.SetData(g.settings, buff, enums.DataTypeUint64, start.ToHighResolutionTime())
		if err != nil {
			return nil, err
		}
		err = internal.SetData(g.settings, buff, enums.DataTypeUint64, end.ToHighResolutionTime())
	}
	if err != nil {
		return nil, err
	}
	err = buff.SetUint8(enums.DataTypeArray)
	if err != nil {
		return nil, err
	}
	if columns == nil {
		err = buff.SetUint8(0x00)
		if err != nil {
			return nil, err
		}
	} else {
		err = types.SetObjectCount(len(*columns), buff)
		if err != nil {
			return nil, err
		}
		for _, it := range *columns {
			err = buff.SetUint8(enums.DataTypeStructure)
			if err != nil {
				return nil, err
			}
			err = buff.SetUint8(4)
			if err != nil {
				return nil, err
			}
			ln, err := LogicalNameToBytes(it.Key.Base().LogicalName())
			if err != nil {
				return nil, err
			}
			err = internal.SetData(g.settings, buff, enums.DataTypeUint16, int(it.Key.Base().ObjectType()))
			if err != nil {
				return nil, err
			}
			err = internal.SetData(g.settings, buff, enums.DataTypeOctetString, ln)
			if err != nil {
				return nil, err
			}
			err = internal.SetData(g.settings, buff, enums.DataTypeInt8, it.Value.AttributeIndex)
			if err != nil {
				return nil, err
			}
			err = internal.SetData(g.settings, buff, enums.DataTypeUint16, it.Value.DataIndex)
			if err != nil {
				return nil, err
			}
		}
	}
	mode := int(pg.GetAccess3(2))
	return g.Read2(pg.Name(), enums.ObjectTypeProfileGeneric, 2, buff, mode)
}

// ReceiverReady returns the generates an acknowledgment message, with which the server is informed to
// send next packets.
//
// Parameters:
//
//	reply: Reply data.
//
// Returns:
//
//	Acknowledgment message as byte array.
func (g *GXDLMSClient) ReceiverReady(reply *GXReplyData) ([]byte, error) {
	return receiverReady(g.settings, reply)
}

// GetData returns the removes the frame from the packet, and returns DLMS PDU.
//
// Parameters:
//
//	reply: The received data from the device.
//	data: Information from the received data.
//
// Returns:
//
//	Is frame complete.
func (g *GXDLMSClient) GetDataFromByteArray(reply []byte, data *GXReplyData, notify *GXReplyData) (bool, error) {
	bb := types.NewGXByteBufferWithData(reply)
	return g.GetData(bb, data, notify)
}

// GetData returns the removes the HDLC frame from the packet, and returns COSEM data only.
//
// Parameters:
//
//	reply: The received data from the device.
//	data: The exported reply information.
//	notify: Optional notify data.
//
// Returns:
//
//	Is frame complete.
func (g *GXDLMSClient) GetData(reply *types.GXByteBuffer, data *GXReplyData, notify *GXReplyData) (bool, error) {
	data.xml = nil
	ret, err := getData(g.settings, reply, data, notify)
	if err != nil {
		if g.translator == nil || g.throwExceptions {
			return false, err
		}
		ret = true
	}
	if ret && g.translator != nil && data.moreData == enums.RequestTypesNone {
		if data.Xml() == nil {
			//TODO:			data.xml = settings.GXDLMSTranslatorStructure{OutputType: g.translator.outputType, OmitXmlNameSpace: g.translator.OmitXmlNameSpace,
			//				Hex: g.translator.Hex, ShowStringAsHex: g.translator.ShowStringAsHex, Comments: g.translator.Comments, Tags: g.translator.tags}
		}
		pos := data.Data.Position()
		data2 := *data.Data
		if data.Command() == enums.CommandGetResponse {
			tmp := types.NewGXByteBufferWithCapacity(4 + data.Data.Size())
			err = tmp.SetUint8(byte(data.Command()))
			if err != nil {
				return false, err
			}
			err = tmp.SetUint8(byte(enums.GetCommandTypeNormal))
			if err != nil {
				return false, err
			}
			err = tmp.SetUint8(uint8(data.InvokeId()))
			if err != nil {
				return false, err
			}
			err = tmp.SetUint8(0)
			if err != nil {
				return false, err
			}
			err = tmp.SetByteBuffer(data.Data)
			if err != nil {
				return false, err
			}
			data.Data = tmp
		} else if data.Command() == enums.CommandMethodResponse {
			tmp := types.NewGXByteBufferWithCapacity(6 + data.Data.Size())
			err = tmp.SetUint8(byte(data.Command()))
			if err != nil {
				return false, err
			}
			err = tmp.SetUint8(byte(enums.GetCommandTypeNormal))
			if err != nil {
				return false, err
			}
			err = tmp.SetUint8(uint8(data.InvokeId()))
			if err != nil {
				return false, err
			}
			err = tmp.SetUint8(0)
			if err != nil {
				return false, err
			}
			err = tmp.SetUint8(1)
			if err != nil {
				return false, err
			}
			err = tmp.SetUint8(0)
			if err != nil {
				return false, err
			}
			err = tmp.SetByteBuffer(data.Data)
			if err != nil {
				return false, err
			}
			data.Data = tmp
		} else if data.Command() == enums.CommandReadResponse {
			tmp := types.NewGXByteBufferWithCapacity(3 + data.Data.Size())
			err = tmp.SetUint8(byte(data.Command()))
			if err != nil {
				return false, err
			}
			err = tmp.SetUint8(byte(enums.VariableAccessSpecificationBlockNumberAccess))
			if err != nil {
				return false, err
			}
			err = tmp.SetUint8(uint8(data.InvokeId()))
			if err != nil {
				return false, err
			}
			err = tmp.SetUint8(0)
			if err != nil {
				return false, err
			}
			err = tmp.SetByteBuffer(data.Data)
			if err != nil {
				return false, err
			}
			data.Data = tmp
		}
		data.Data.SetPosition(0)
		if data.Command() == enums.CommandSnrm || data.Command() == enums.CommandUa {
			data.xml.AppendStartTag(int(data.Command()), "", "", true)
			if data.Data.Size() != 0 {
				//TODO: g.translator.PduToXml(data.Xml(), data.Data(), g.translator.OmitXmlDeclaration, g.translator.OmitXmlNameSpace, true, nil)
			}
			data.xml.AppendEndTag(int(data.Command()), true)
		} else {
			if data.Data.Size() != 0 {
				//TODO: g.translator.PduToXml(data.Xml(), data.Data(), g.translator.OmitXmlDeclaration, g.translator.OmitXmlNameSpace, true, nil)
			}
			data.Data = &data2
			data.Data.SetPosition(pos)
		}
	}
	return ret, nil
}

// GetValue returns the get value from DLMS byte stream.
//
// Parameters:
//
//	data: Received data.
//	useUtc: Standard says that Time zone is from normal time to UTC in minutes.
//	  If meter is configured to use UTC time (UTC to normal time) set this to true.
//
// Returns:
//
//	Parsed value.
func (g *GXDLMSClient) GetValue(data *types.GXByteBuffer, useUtc bool) (any, error) {
	return g.GetValueWithType(data, enums.DataTypeNone, useUtc)
}

// GetValue returns the get value from DLMS byte stream.
//
// Parameters:
//
//	data: Received data.
//	type: Conversion type is used if returned data is byte array.
//	useUtc: Standard says that Time zone is from normal time to UTC in minutes.
//	  If meter is configured to use UTC time (UTC to normal time) set this to true.
//
// Returns:
//
//	Parsed value.
func (g *GXDLMSClient) GetValueWithType(data *types.GXByteBuffer, type_ enums.DataType, useUtc bool) (any, error) {
	info := internal.GXDataInfo{}
	value, err := internal.GetData(g.settings, data, &info)
	if err != nil {
		return nil, err
	}
	if v, ok := value.([]byte); ok {
		value, err = ChangeTypeFromByteArray(v, type_, useUtc)
		if err != nil {
			return nil, err
		}
	}
	return value, nil
}

// GetServerAddress returns the convert physical address and logical address to server address.
//
// Parameters:
//
//	logicalAddress: Server logical address.
//	physicalAddress: Server physical address.
//
// Returns:
//
//	Server address.
func GetServerAddress(logicalAddress int, physicalAddress int) (int, error) {
	return GetServerAddressWithAddressSize(logicalAddress, physicalAddress, 0)
}

// GetServerAddress returns the convert physical address and logical address to server address.
//
// Parameters:
//
//	logicalAddress: Server logical address.
//	physicalAddress: Server physical address.
//	addressSize: Address size in bytes.
//
// Returns:
//
//	Server address.
func GetServerAddressWithAddressSize(logicalAddress int, physicalAddress int, addressSize int) (int, error) {
	var value int
	if addressSize < 4 && physicalAddress < 0x80 && logicalAddress < 0x80 {
		value = logicalAddress<<7 | physicalAddress
	} else if physicalAddress < 0x4000 && logicalAddress < 0x4000 {
		value = logicalAddress<<14 | physicalAddress
	} else {
		return 0, errors.New("Invalid logical or physical address.")
	}
	return value, nil
}

// GetServerAddressFromSerialNumber returns the converts meter serial number to server address.
// Default formula is used.
// All meters do not use standard formula or support serial number addressing at all.
//
// Parameters:
//
//	serialNumber: Meter serial number.
//	logicalAddress: Used logical address.
//
// Returns:
//
//	Server address.
func GetServerAddressFromSerialNumber(serialNumber uint64, logicalAddress int) (int, error) {
	return GetServerAddressFromSerialNumberWithFormula(serialNumber, logicalAddress, "")
}

// GetServerAddressFromSerialNumber returns the converts meter serial number to server address.
// All meters do not use standard formula or support serial number addressing at all.
//
// Parameters:
//
//	serialNumber: Meter serial number.
//	logicalAddress: Server logical address.
//	formula: Formula used to convert serial number to server address.
//
// Returns:
//
//	Server address.
func GetServerAddressFromSerialNumberWithFormula(serialNumber uint64, logicalAddress int, formula string) (int, error) {
	// If formula is not given use default formula.This formula is defined in DLMS specification.
	if formula == "" {
		formula = "SN % 10000 + 1000"
	}
	ret, err := internal.SerialnumberCounterCount(serialNumber, formula)
	if err != nil {
		return 0, err
	}
	if logicalAddress == 0 {
		return int(ret), nil
	}
	return logicalAddress<<14 | int(ret), nil
}

// encryptLandisGyrHighLevelAuthentication returns the encrypt Landis+Gyr High level password.
//
// Parameters:
//
//	password: User password.
//	seed: Seed received from the meter.
func encryptLandisGyrHighLevelAuthentication(password []byte, seed []byte) []byte {
	crypted := make([]byte, len(seed))
	copy(crypted, seed)
	for pos := 0; pos < len(password) && pos < len(crypted); pos++ {
		if password[pos] != 0x30 {
			crypted[pos] += byte(password[pos] - 0x30)

			// Convert to upper case letter.
			if crypted[pos] > '9' && crypted[pos] < 'A' {
				crypted[pos] += 7
			}
			if crypted[pos] > 'F' {
				crypted[pos] = byte('0' + crypted[pos] - 'G')
			}
		}
	}
	return crypted
}

// AccessRequest returns the generates a access service message.
//
// Parameters:
//
//	time: Send time. Set to DateTime.MinValue is not used.
//
// Returns:
//
//	Access request as byte array.
func (g *GXDLMSClient) AccessRequest(time_ *time.Time, list []GXDLMSAccessItem) ([][]byte, error) {
	mode := 0
	var err error
	bb := types.NewGXByteBuffer()
	types.SetObjectCount(len(list), bb)
	for _, it := range list {
		err = bb.SetUint8(byte(it.Command))
		if err != nil {
			return nil, err
		}
		err = bb.SetUint16(uint16(it.Target.Base().ObjectType()))
		if err != nil {
			return nil, err
		}
		ret, err := LogicalNameToBytes(it.Target.Base().LogicalName())
		if err != nil {
			return nil, err
		}
		bb.Set(ret)
		err = bb.SetUint8(it.Index)
		if err != nil {
			return nil, err
		}
		m := int(it.Target.Base().GetAccess3(int(it.Index)))
		if m > mode {
			mode = m
		}
	}
	//Data
	err = types.SetObjectCount(len(list), bb)
	if err != nil {
		return nil, err
	}
	dt := time.Now()
	for _, it := range list {
		if it.Command == enums.AccessServiceCommandTypeGet {
			err = bb.SetUint8(0)
			if err != nil {
				return nil, err
			}
		} else if it.Command == enums.AccessServiceCommandTypeSet {
			value, err := it.Target.GetValue(g.settings, internal.NewValueEventArgs(g.settings, it.Target, it.Index))
			if err != nil {
				return nil, err
			}
			type_, err := it.Target.GetDataType(int(it.Index))
			if err != nil {
				return nil, err
			}
			if type_ == enums.DataTypeNone {
				type_, err = internal.GetDLMSDataType(reflect.TypeOf(value))
				if err != nil {
					return nil, err
				}
			}
			err = internal.SetData(g.settings, bb, type_, value)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, errors.New("Invalid command.")
		}
	}
	p := NewGXDLMSLNParameters(g.settings, 0, enums.CommandAccessRequest, 0xFF, nil, bb, 0xff, enums.CommandNone)
	p.AccessMode = mode
	p.time = types.NewGXDateTimeFromTime(dt)
	return getLnMessages(p)
}

// ParseReport returns the parse received Information reports and Event notifications.
//
// Parameters:
//
//	reply: Reply.
//
// Returns:
//
//	Data notification data.
func (g *GXDLMSClient) ParseReport(reply *GXReplyData, list []*types.GXKeyValuePair[objects.IGXDLMSBase, int]) (any, error) {
	var err error
	if reply.Command() == enums.CommandEventNotification {
		err = handleEventNotification(g.settings, reply, list)
		return nil, err
	} else if reply.Command() == enums.CommandInformationReport {
		err = handleInformationReport(g.settings, reply, list)
		return nil, err
	} else if reply.Command() == enums.CommandDataNotification {
		return reply.Value, nil
	}
	return nil, errors.New("Invalid command. %s" + reply.Command().String())
}

// CustomHdlcFrameRequest returns the generates a invalid HDLC frame.
// This method can be used for sending custom HDLC frames example in testing.
//
// Parameters:
//
//	command: HDLC command.
//	data: data
//
// Returns:
//
//	HDLC frame request, as byte array.
func (g *GXDLMSClient) CustomHdlcFrameRequest(command uint8, data *types.GXByteBuffer) ([]byte, error) {
	if !useHdlc(g.settings.InterfaceType) {
		return nil, errors.New("This method can be used only to generate HDLC custom frames")
	}
	return getHdlcFrame(g.settings, command, data, true)
}

// GetFrameSize returns the get size of the frame.
// When WRAPPER is used this method can be used to check how many bytes we need to read.
//
// Parameters:
//
//	data: Received data.
//
// Returns:
//
//	Size of received bytes on the frame.
func (g *GXDLMSClient) GetFrameSize(data *types.GXByteBuffer) int {
	var ret int
	switch g.InterfaceType() {
	case enums.InterfaceTypeHDLC, enums.InterfaceTypeHdlcWithModeE:
		ret = 0
		index := data.Position()
		// If whole frame is not received yet.
		if data.Available() > 8 {
			for pos := data.Position(); pos < data.Size(); pos++ {
				ch, err := data.Uint8()
				if err != nil {
					break
				}
				if ch == internal.HDLCFrameStartEnd {
					break
				}
			}
			frame, err := data.Uint8()
			if err != nil {
				break
			}
			// Check frame length.
			if (frame & 0x7) != 0 {
				ret = int((frame & 0x7) << 8)
			}
			ch, err := data.Uint8()
			if err != nil {
				break
			}
			ret = ret + 1 + int(ch)
		}
		data.SetPosition(index)
	case enums.InterfaceTypeWRAPPER:
		if data.Available() < 8 {
			ret = 8 - data.Available()
		} else {
			v, err := data.Uint16At(data.Position())
			if err != nil {
				break
			}
			if v != 1 {
				ret = 8 - data.Available()
			} else {
				v, err := data.Uint16At(data.Position() + 6)
				if err != nil {
					break
				}
				ret = 8 + int(v) - data.Available()
			}
		}
	case enums.InterfaceTypePlc:
		if data.Available() < 2 {
			ret = 2 - data.Available()
		} else if v, err := data.Uint8At(data.Position()); v != 2 && err == nil {
			ret = 2 - data.Available()
		} else {
			v, err := data.Uint8At(data.Position() + 1)
			if err != nil {
				break
			}
			ret = 2 + int(v) - data.Available()
		}
	case enums.InterfaceTypePlcHdlc:
		ret := getPlcSfskFrameSize(data)
		if ret == 0 {
			ret = 2
		}
	default:
		ret = 1
	}
	if ret < 1 {
		ret = 1
	}
	return ret
}

// GetHdlcAddressInfo returns the get HDLC sender and receiver address information.
//
// Parameters:
//
//	reply: Received data.
//	target: target (primary) address
//	source: Source (secondary) address.
//	type: DLMS frame type.
func GetHdlcAddressInfo(reply *types.GXByteBuffer, target *int, source *int, type_ *uint8) {
	getHdlcAddressInfo(reply, target, source, type_)
}

// CanRead returns the can client read the object attribute index.
// This method is added because Association Logical Name version #3 where access rights are defined with bitmask.
//
// Parameters:
//
//	target: Object to read.
//	index: Attribute index.
//
// Returns:
//
//	True, if read is allowed.
func (g *GXDLMSClient) CanRead(target *objects.GXDLMSObject, index int) bool {
	// Handle access rights for Association LN Version < 3.
	access := target.GetAccess(index)
	if (access&enums.AccessModeRead) == 0 && access != enums.AccessModeAuthenticatedRead && access != enums.AccessModeAuthenticatedReadWrite {
		// If bit mask is used.
		m := target.GetAccess3(index)
		if (m & enums.AccessMode3Read) == 0 {
			return false
		}
		security := enums.SecurityNone
		signing := enums.SigningNone
		if g.settings.Cipher != nil {
			security = g.settings.Cipher.Security()
			signing = g.settings.Cipher.Signing()
		}
		// If authenticatation is expected, but secured connection is not used.
		if (m&(enums.AccessMode3AuthenticatedRequest|enums.AccessMode3AuthenticatedResponse)) != 0 && (security&(enums.SecurityAuthentication)) == 0 {
			return false
		}
		// If encryption is expected, but secured connection is not used.
		if (m&(enums.AccessMode3EncryptedRequest|enums.AccessMode3EncryptedResponse)) != 0 && (security&(enums.SecurityEncryption)) == 0 {
			return false
		}
		// If signing is expected, but it's not used.
		if (m&(enums.AccessMode3DigitallySignedRequest|enums.AccessMode3DigitallySignedResponse)) != 0 && (signing&(enums.SigningGeneralSigning)) == 0 {
			if g.settings.Cipher.SigningKeyPair() == nil {
				return false
			}
		}
	}
	return true
}

// CanWrite returns the can client write the object attribute index.
// This method is added because Association Logical Name version #3 where access rights are defined with bitmask.
//
// Parameters:
//
//	target: Object to write.
//	index: Attribute index.
//
// Returns:
//
//	True, if write is allowed.
func (g *GXDLMSClient) CanWrite(target *objects.GXDLMSObject, index int) bool {
	// Handle access rights for Association LN Version < 3.
	access := target.GetAccess(index)
	if (access&enums.AccessModeWrite) == 0 && access != enums.AccessModeAuthenticatedWrite && access != enums.AccessModeAuthenticatedReadWrite {
		// If bit mask is used.
		m := target.GetAccess3(index)
		if (m & enums.AccessMode3Write) == 0 {
			return false
		}
		security := enums.SecurityNone
		signing := enums.SigningNone
		if g.settings.Cipher != nil {
			security = g.settings.Cipher.Security()
			signing = g.settings.Cipher.Signing()
		}
		// If authentication is expected, but secured connection is not used.
		if (m&(enums.AccessMode3AuthenticatedRequest|enums.AccessMode3AuthenticatedResponse)) != 0 && (security&(enums.SecurityAuthentication)) == 0 {
			return false
		}
		// If encryption is expected, but secured connection is not used.
		if (m&(enums.AccessMode3EncryptedRequest|enums.AccessMode3EncryptedResponse)) != 0 && (security&(enums.SecurityEncryption)) == 0 {
			return false
		}
		// If signing is expected, but it's not used.
		if (m&(enums.AccessMode3DigitallySignedRequest|enums.AccessMode3DigitallySignedResponse)) != 0 && (signing&(enums.SigningGeneralSigning)) == 0 {
			if g.settings.Cipher.SigningKeyPair() == nil {
				return false
			}
		}
	}
	return true
}

// CanInvoke returns the can client invoke server methods.
// This method is added because Association Logical Name version #3 where access rights are defined with bitmask.
//
// Parameters:
//
//	target: Object to invoke.
//	index: Method attribute index.
//
// Returns:
//
//	True, if client can access meter methods.
func (g *GXDLMSClient) CanInvoke(target *objects.GXDLMSObject, index int) bool {
	// Handle access rights for Association LN Version < 3.
	if target.GetMethodAccess(index) == enums.MethodAccessModeNoAccess {
		// If bit mask is used.
		m := target.GetMethodAccess3(index)
		if (m & enums.MethodAccessMode3Access) == 0 {
			return false
		}
		security := enums.SecurityNone
		signing := enums.SigningNone
		if g.settings.Cipher != nil {
			security = g.settings.Cipher.Security()
			signing = g.settings.Cipher.Signing()
		}
		// If authentication is expected, but secured connection is not used.
		if (m&(enums.MethodAccessMode3AuthenticatedRequest|enums.MethodAccessMode3AuthenticatedResponse)) != 0 && (security&(enums.SecurityAuthentication)) == 0 {
			return false
		}
		// If encryption is expected, but secured connection is not used.
		if (m&(enums.MethodAccessMode3EncryptedRequest|enums.MethodAccessMode3EncryptedResponse)) != 0 && (security&(enums.SecurityEncryption)) == 0 {
			return false
		}
		// If signing is expected, but it's not used.
		if (m&(enums.MethodAccessMode3DigitallySignedRequest|enums.MethodAccessMode3DigitallySignedResponse)) != 0 && (signing&(enums.SigningGeneralSigning)) == 0 {
			if g.settings.Cipher.SigningKeyPair() == nil {
				return false
			}
		}
	}
	return true
}

// NewGXDLMSClient returns the new DLMS client.
//
// Parameters:
//
//	useLogicalNameReferencing: Use logical name referencing.
//	serverAddress: Server address.
//	authentication: Authentication type.
//	password: User password.
//	interfaceType: Interface type.
func NewGXDLMSClient(useLogicalNameReferencing bool, clientAddress int, serverAddress int, authentication enums.Authentication,
	password []byte, interfaceType enums.InterfaceType) (*GXDLMSClient, error) {
	ret := &GXDLMSClient{}
	ret.settings = settings.NewGXDLMSSettingsWithParams(false, useLogicalNameReferencing, interfaceType)
	ret.settings.Cipher = &secure.GXCiphering{}
	err := ret.SetClientAddress(clientAddress)
	if err != nil {
		return nil, err
	}
	err = ret.SetServerAddress(serverAddress)
	if err != nil {
		return nil, err
	}
	err = ret.SetAuthentication(authentication)
	if err != nil {
		return nil, err
	}
	err = ret.SetPassword(password)
	if err != nil {
		return nil, err
	}
	err = ret.SetInterfaceType(interfaceType)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
