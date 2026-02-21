package settings

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
	"log"

	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal/constants"
	"github.com/Gurux/gxdlms-go/types"
)

const (
	// defaultMaxReceivePDUSize is the default Max received PDU size.
	defaultMaxReceivePDUSize uint16 = 0xFFFF
)

// GXDLMSSettings includes DLMS communication
type GXDLMSSettings struct {
	// gbtWindowSize is the General Block Transfer window size.
	gbtWindowSize uint8

	// assignedAssociation is the assigned association for the server.
	assignedAssociation any

	useLogicalNameReferencing bool

	// invokeID is the Invoke ID.
	invokeID uint8

	// LongInvokeID is the long Invoke ID.
	LongInvokeID uint32

	// ctoSChallenge is the Client to server challenge.
	ctoSChallenge []byte

	challengeSize uint8

	// stoCChallenge is the Server to Client challenge.
	stoCChallenge []byte

	// SenderFrame is the HDLC sender frame sequence number.
	SenderFrame uint8

	// ReceiverFrame is the HDLC receiver frame sequence number.
	ReceiverFrame uint8

	sourceSystemTitle []byte

	// ClientPublicKeyCertificate is the optional ECDSA public key certificate that is sent in part of AARQ.
	ClientPublicKeyCertificate *types.GXx509Certificate

	// ServerPublicKeyCertificate is the optional ECDSA public key certificate that is sent in part of AARE.
	ServerPublicKeyCertificate *types.GXx509Certificate

	// PreEstablishedSystemTitle is the pre-established system title.
	PreEstablishedSystemTitle []byte

	// Kek is the Key Encrypting Key, also known as Master key.
	Kek []byte

	// Count is the Long data count.
	Count uint32

	// Index is the Long data index.
	Index uint32

	// TargetEphemeralKey is the target ephemeral public key.
	TargetEphemeralKey interface{} // TODO: Replace with actual GXPublicKey type when implemented

	maxReceivePDUSize uint16

	// ProtocolVersion is the protocol version.
	ProtocolVersion string

	// ProposedConformance is what client tells it wants to use when connection is made.
	ProposedConformance enums.Conformance

	// NegotiatedConformance is what server tells is available and client will know it.
	NegotiatedConformance enums.Conformance

	// Cipher is the cipher interface.
	// GXDLMSAssociationShortName and GXDLMSAssociationLogicalName use this if GMAC authentication is used.
	Cipher GXICipher

	// UserID is the identifier of the user.
	// This value is used if user list on Association LN is used.
	UserID int

	// QualityOfService is the quality of service.
	QualityOfService uint8

	// UseUtc2NormalTime indicates if meter is configured to use UTC time (UTC to normal time).
	// Standard says that Time zone is from normal time to UTC in minutes.
	// If meter is configured to use UTC time (UTC to normal time) set this to true.
	UseUtc2NormalTime bool

	// DateTimeSkips are the skipped date time fields.
	// This value can be used if meter can't handle deviation or status.
	DateTimeSkips enums.DateTimeSkips

	// Standard is the used standard.
	Standard enums.Standard

	// Command is the last executed command.
	Command enums.Command

	// CommandType is the last executed command type.
	CommandType uint8

	expectedInvocationCounter uint64

	// InvocationCounter is the invocation counter object.
	InvocationCounter interface{}

	// EphemeralKek is the ephemeral KEK.
	EphemeralKek []byte

	// EphemeralBlockCipherKey is the ephemeral Block cipher key.
	EphemeralBlockCipherKey []byte

	// EphemeralBroadcastBlockCipherKey is the ephemeral broadcast block cipherKey.
	EphemeralBroadcastBlockCipherKey []byte

	// EphemeralAuthenticationKey is the ephemeral authentication key.
	EphemeralAuthenticationKey []byte

	// Keys is the list of certificates to decrypt the data (for XML).
	Keys []*types.GXKeyValuePair[*types.GXPkcs8, *types.GXx509Certificate]

	// UseCustomChallenge indicates if custom challenges are used.
	// If custom challenge is used new challenge is not generated if it is set.
	// This is for debugging purposes.
	UseCustomChallenge bool

	// StartingBlockIndex is the starting block index.
	// Default is One based, but some meters use Zero based value.
	// Usually this is not used.
	StartingBlockIndex uint32

	// DLMSVersion is the DLMS version number.
	DLMSVersion uint8

	// BlockIndex is the current block index.
	BlockIndex uint32

	// isServer indicates if this is server or client
	isServer bool

	// Priority is the used priority.
	Priority enums.Priority

	// ServiceClass is the used service class.
	ServiceClass enums.ServiceClass

	// Objects is the collection of the objects.
	Objects interface{}

	maxServerPDUSize uint16

	// Authentication is the used authentication.
	Authentication enums.Authentication

	// Password is the client password.
	Password []byte

	// Closing indicates if connection is closing.
	// Some meters might return invalid data when disconnect is called and cause infinity loop.
	// This property is used to ignore it.
	Closing bool

	// Connected is the connection state.
	Connected enums.ConnectionState

	// BlockNumberAck is the block number acknowledged in GBT.
	BlockNumberAck uint16

	// OverwriteAttributeAccessRights indicates if association view tells wrong access rights and they are overwritten.
	OverwriteAttributeAccessRights bool

	// Hdlc contains HDLC framing
	Hdlc *GXHdlcSettings

	// Gateway contains gateway
	Gateway *GXDLMSGateway

	// Plc contains PLC
	Plc *GXPlcSettings

	// MBus contains M-Bus
	MBus *GXMBusSettings

	// Pdu contains PDU
	Pdu *GXPduSettings

	// Coap contains CoAP
	Coap *GXCoAPSettings

	// Broadcast indicates if data is sent as a broadcast or unicast.
	Broadcast bool

	// Compression indicates if V.44 Compression is used.
	Compression bool

	// InterfaceType is the used interface.
	InterfaceType enums.InterfaceType

	// ClientAddress is the client address.
	ClientAddress int

	// PushClientAddress is the address server uses when sending push messages.
	// Client address is used if PushAddress is zero.
	PushClientAddress int

	// GbtCount is the General Block Transfer count in server.
	GbtCount uint8

	// ServerAddress is the server address.
	ServerAddress int

	// ServerAddressSize is the size of Server address.
	ServerAddressSize uint8

	// Version can be used for backward compatibility.
	Version int

	// AutoIncreaseInvokeID indicates if Invoke ID is auto increased.
	AutoIncreaseInvokeID bool

	// CryptoNotifier for external HSM operations.
	CryptoNotifier interface{} // TODO: Replace with GXCryptoNotifier when implemented

	// customObject is the event invoked when custom manufacturer object is created.
	customObject interface{} // TODO: Replace with ObjectCreateEventHandler when implemented

	// customPdu is the event invoked when custom PDU is handled.
	customPdu interface{} // TODO: Replace with CustomPduEventHandler when implemented
}

// NewGXDLMSSettings creates a new DLMS settings instance with default values.
func NewGXDLMSSettings(objects any) *GXDLMSSettings {
	return NewGXDLMSSettingsWithParams(false, true, enums.InterfaceTypeHDLC, objects)
}

// getInitialConformance returns the get initial Conformance
//
// Parameters:
//
//	useLogicalNameReferencing: Is logical name referencing used.
//
// Returns:
//
//	Initial Conformance.
func getInitialConformance(useLogicalNameReferencing bool) enums.Conformance {
	if useLogicalNameReferencing {
		return enums.ConformanceGeneralBlockTransfer | enums.ConformanceBlockTransferWithAction | enums.ConformanceBlockTransferWithSetOrWrite | enums.ConformanceBlockTransferWithGetOrRead | enums.ConformanceSet | enums.ConformanceSelectiveAccess | enums.ConformanceAction | enums.ConformanceMultipleReferences | enums.ConformanceGet | enums.ConformanceAccess | enums.ConformanceGeneralProtection | enums.ConformanceDeltaValueEncoding
	}
	return enums.ConformanceInformationReport | enums.ConformanceRead | enums.ConformanceUnconfirmedWrite | enums.ConformanceWrite | enums.ConformanceParameterizedAccess | enums.ConformanceMultipleReferences
}

// NewGXDLMSSettingsWithParams creates a new DLMS settings instance with specified parameters.
func NewGXDLMSSettingsWithParams(server bool, ln bool, interfaceType enums.InterfaceType, objects any) *GXDLMSSettings {
	s := &GXDLMSSettings{
		InterfaceType:      interfaceType,
		UseCustomChallenge: false,
		StartingBlockIndex: 1,
		BlockIndex:         1,
		DLMSVersion:        6,
		invokeID:           0x1,
		LongInvokeID:       0x1,
		Priority:           enums.PriorityHigh,
		ServiceClass:       enums.ServiceClassConfirmed,
		maxServerPDUSize:   defaultMaxReceivePDUSize,
		maxReceivePDUSize:  defaultMaxReceivePDUSize,
		isServer:           server,
		gbtWindowSize:      1,
		UserID:             -1,
		Standard:           enums.StandardDLMS,
		challengeSize:      16,
	}
	switch interfaceType {
	case enums.InterfaceTypeHDLC, enums.InterfaceTypeHdlcWithModeE:
		s.Hdlc = NewGXHdlcSettings()
	case enums.InterfaceTypePlc:
		s.Plc = &GXPlcSettings{}
	case enums.InterfaceTypeWirelessMBus:
		s.MBus = &GXMBusSettings{}
	case enums.InterfaceTypePDU:
		s.Pdu = &GXPduSettings{}
	case enums.InterfaceTypeCoAP:
		s.Coap = &GXCoAPSettings{}
	default:
	}
	s.Objects = objects
	s.useLogicalNameReferencing = ln
	s.ProposedConformance = getInitialConformance(ln)
	if server {
		s.ProposedConformance |= enums.ConformanceGeneralProtection
	}
	s.ResetFrameSequence()
	return s
}

// IsCiphered checks if data is ciphered.
func (s *GXDLMSSettings) IsCiphered(checkGeneralSigning bool) bool {
	if s.Cipher == nil {
		return false
	}
	return s.Cipher.Security() != enums.SecurityNone ||
		(checkGeneralSigning && s.Cipher.Signing() == enums.SigningGeneralSigning)
}

// SourceSystemTitle returns the source system title.
func (s *GXDLMSSettings) SourceSystemTitle() []byte {
	if s.Cipher != nil {
		return s.Cipher.RecipientSystemTitle()
	}
	return s.sourceSystemTitle
}

// SetSourceSystemTitle sets the source system title.
func (s *GXDLMSSettings) SetSourceSystemTitle(value []byte) {
	if s.Cipher != nil {
		s.Cipher.SetRecipientSystemTitle(value)
	}
	s.sourceSystemTitle = value
}

// ChallengeSize returns the challenge size.
func (s *GXDLMSSettings) ChallengeSize() uint8 {
	return s.challengeSize
}

// SetChallengeSize sets the challenge size.
// Random challenge is used if value is zero.
func (s *GXDLMSSettings) SetChallengeSize(value uint8) error {
	if s.Authentication == enums.AuthenticationHighECDSA && value < 32 {
		return errors.New("invalid challenge size. ECDSA Challenge must be between 32 to 64 uint8s")
	}
	if value < 8 || value > 64 {
		return errors.New("invalid challenge size. Challenge must be between 8 to 64 uint8s")
	}
	s.challengeSize = value
	return nil
}

// CtoSChallenge returns the Client to Server challenge.
func (s *GXDLMSSettings) CtoSChallenge() []byte {
	return s.ctoSChallenge
}

// SetCtoSChallenge sets the Client to Server challenge.
func (s *GXDLMSSettings) SetCtoSChallenge(value []byte) {
	s.ctoSChallenge = value
}

// StoCChallenge returns the Server to Client challenge.
func (s *GXDLMSSettings) StoCChallenge() []byte {
	return s.stoCChallenge
}

// SetStoCChallenge sets the Server to Client challenge.
func (s *GXDLMSSettings) SetStoCChallenge(value []byte) {
	s.stoCChallenge = value
}

// IncreaseReceiverSequence increases the receiver sequence.
func IncreaseReceiverSequence(value uint8) uint8 {
	return value + 0x20 | 0x10 | value&0xE
}

// IncreaseSendSequence increases the sender sequence.
func IncreaseSendSequence(value uint8) uint8 {
	return value&0xF0 | (value+0x2)&0xE
}

// ResetFrameSequence resets frame sequence.
func (s *GXDLMSSettings) ResetFrameSequence() {
	if s.isServer {
		s.SenderFrame = 0x1E
		s.ReceiverFrame = 0xEE
	} else {
		s.SenderFrame = 0xFE
		s.ReceiverFrame = 0xE
	}
}

// CheckFrame checks if the frame is valid.
func (s *GXDLMSSettings) CheckFrame(frame uint8, xml interface{}) bool {
	// If notify
	if frame == 0x13 {
		return true
	}
	// If U frame.
	if (frame & byte(constants.HdlcFrameTypeUframe)) == byte(constants.HdlcFrameTypeUframe) {
		if frame == 0x93 {
			isEcho := !s.isServer && frame == 0x93 &&
				(s.SenderFrame == 0x10 || s.SenderFrame == 0xfe) && s.ReceiverFrame == 0xE
			s.ResetFrameSequence()
			return !isEcho
		}
		if frame == 0x73 && !s.isServer {
			return s.SenderFrame == 0xFE && s.ReceiverFrame == 0xE
		}
		return true
	}
	// If S-frame.
	if (frame & byte(constants.HdlcFrameTypeSframe)) == byte(constants.HdlcFrameTypeSframe) {
		// If echo.
		if frame == (s.SenderFrame & 0xF1) {
			return false
		}
		s.ReceiverFrame = IncreaseReceiverSequence(s.ReceiverFrame)
		return true
	}
	// Handle I-frame.
	var expected uint8
	if (s.SenderFrame & 0x1) == 0 {
		expected = IncreaseReceiverSequence(IncreaseSendSequence(s.ReceiverFrame))
		if frame == expected {
			s.ReceiverFrame = frame
			return true
		}
		// If the final bit is not set.
		if frame == (expected & ^uint8(0x10)) && s.Hdlc.WindowSizeRX() != 1 {
			s.ReceiverFrame = frame
			return true
		}
		// If Final bit is not set for the previous message.
		if (s.ReceiverFrame&0x10) == 0 && s.Hdlc.WindowSizeRX() != 1 {
			expected = 0x10 | IncreaseSendSequence(s.ReceiverFrame)
			if frame == expected {
				s.ReceiverFrame = frame
				return true
			}
			// If the final bit is not set.
			if frame == (expected & ^uint8(0x10)) {
				s.ReceiverFrame = frame
				return true
			}
		}
	} else {
		// If answer for RR.
		expected = IncreaseSendSequence(s.ReceiverFrame)
		if frame == expected {
			s.ReceiverFrame = frame
			return true
		}
		if frame == (expected & ^uint8(0x10)) {
			s.ReceiverFrame = frame
			return true
		}
		if s.Hdlc.WindowSizeRX() != 1 {
			// If HDLC window size is bigger than one.
			if frame == (expected | 0x10) {
				s.ReceiverFrame = frame
				return true
			}
		}
	}
	// If try to find data from bytestream and not real communicating.
	if xml != nil && ((!s.isServer && s.ReceiverFrame == 0xE) ||
		(s.isServer && s.ReceiverFrame == 0xEE)) {
		s.ReceiverFrame = frame
		return true
	}
	log.Printf("Invalid HDLC Frame: %X. Expected: %X\n", frame, expected)
	return false
}

// NextSend generates I-frame.
func (s *GXDLMSSettings) NextSend(first bool) uint8 {
	if first {
		s.SenderFrame = IncreaseReceiverSequence(IncreaseSendSequence(s.SenderFrame))
	} else {
		s.SenderFrame = IncreaseSendSequence(s.SenderFrame)
	}
	return s.SenderFrame
}

// ReceiverReady generates Receiver Ready S-frame.
func (s *GXDLMSSettings) ReceiverReady() uint8 {
	s.SenderFrame = IncreaseReceiverSequence(s.SenderFrame) | 1
	return s.SenderFrame & 0xF1
}

// KeepAlive generates Keep Alive S-frame.
func (s *GXDLMSSettings) KeepAlive() uint8 {
	s.SenderFrame = s.SenderFrame | 1
	return s.SenderFrame & 0xF1
}

// ResetBlockIndex resets block index to default value.
func (s *GXDLMSSettings) ResetBlockIndex() {
	s.BlockIndex = s.StartingBlockIndex
	s.BlockNumberAck = 0
}

// IncreaseBlockIndex increases block index.
func (s *GXDLMSSettings) IncreaseBlockIndex() {
	s.BlockIndex++
}

// GbtWindowSize returns the General Block Transfer window size.
func (s *GXDLMSSettings) GbtWindowSize() uint8 {
	return s.gbtWindowSize
}

// SetGbtWindowSize sets the General Block Transfer window size.
func (s *GXDLMSSettings) SetGbtWindowSize(value uint8) error {
	if value > 63 {
		return errors.New("maximum size for GBT Window is 63 messages")
	}
	s.gbtWindowSize = value
	return nil
}

// MaxPduSize returns the maximum PDU size.
func (s *GXDLMSSettings) MaxPduSize() uint16 {
	return s.maxReceivePDUSize
}

// SetMaxPduSize sets the maximum PDU size.
func (s *GXDLMSSettings) SetMaxPduSize(value uint16) error {
	if value < 64 && value != 0 {
		return errors.New("MaxReceivePDUSize must be at least 64 or 0")
	}
	s.maxReceivePDUSize = value
	return nil
}

// GetMaxServerPDUSize returns the server maximum PDU size.
func (s *GXDLMSSettings) GetMaxServerPDUSize() uint16 {
	return s.maxServerPDUSize
}

// SetMaxServerPDUSize sets the server maximum PDU size.
func (s *GXDLMSSettings) SetMaxServerPDUSize(value uint16) {
	if s.InterfaceType == enums.InterfaceTypePlc {
		value = 134
	}
	s.maxServerPDUSize = value
}

// UseLogicalNameReferencing returns if Logical Name Referencing is used.
func (s *GXDLMSSettings) UseLogicalNameReferencing() bool {
	return s.useLogicalNameReferencing
}

// SetUseLogicalNameReferencing sets if Logical Name Referencing is used.
func (s *GXDLMSSettings) SetUseLogicalNameReferencing(value bool) {
	if s.useLogicalNameReferencing != value {
		s.useLogicalNameReferencing = value
		s.ProposedConformance = getInitialConformance(value)
		if s.isServer {
			s.ProposedConformance |= enums.ConformanceGeneralProtection
		}
	}
}

// AssignedAssociation returns the assigned association for the server.
func (s *GXDLMSSettings) AssignedAssociation() any {
	return s.assignedAssociation
}

// SetAssignedAssociation sets the assigned association for the server.
func (s *GXDLMSSettings) SetAssignedAssociation(value any) {
	/*
		if ln, ok := s.assignedAssociation.(*objects.GXDLMSAssociationLogicalName); ok {
			ln.AssociationStatus = enums.AssociationStatusNonAssociated
			ln.XDLMSContextInfo.CypheringInfo = nil
			s.InvocationCounter = nil
			s.Cipher.SetSecurityPolicy(enums.SecurityPolicyNone)
			s.EphemeralBlockCipherKey = nil
			s.EphemeralBroadcastBlockCipherKey = nil
			s.EphemeralAuthenticationKey = nil
			s.Cipher.SetSecuritySuite(enums.SecuritySuite0)
			s.Cipher.SetSigning(enums.SigningNone)
		}
		s.assignedAssociation = value
		if ln2, ok := s.assignedAssociation.(*objects.GXDLMSAssociationLogicalName); ok {
			s.ProposedConformance = ln2.XDLMSContextInfo.Conformance
			s.maxServerPDUSize = ln2.XDLMSContextInfo.MaxReceivePduSize
			s.Authentication = ln2.AuthenticationMechanismName.MechanismId
			s.UpdateSecuritySettings(nil)
		}
	*/
}

// InvokeID returns the Invoke ID.
func (s *GXDLMSSettings) InvokeID() uint8 {
	return s.invokeID
}

// SetInvokeID sets the Invoke ID.
func (s *GXDLMSSettings) SetInvokeID(value uint8) error {
	if value > 0xF {
		return errors.New("invalid InvokeID")
	}
	s.invokeID = value
	return nil
}

// UpdateInvokeID updates invoke ID and priority.
func (s *GXDLMSSettings) UpdateInvokeID(value uint8) {
	if (value & 0x80) != 0 {
		s.Priority = enums.PriorityHigh
	} else {
		s.Priority = enums.PriorityNormal
	}
	if (value & 0x40) != 0 {
		s.ServiceClass = enums.ServiceClassConfirmed
	} else {
		s.ServiceClass = enums.ServiceClassUnConfirmed
	}
	s.invokeID = value & 0xF
}

// ExpectedInvocationCounter returns the expected Invocation (Frame) counter value.
// Expected Invocation counter is not checked if value is zero.
func (s *GXDLMSSettings) ExpectedInvocationCounter() uint64 {
	return s.expectedInvocationCounter
}

// SetExpectedInvocationCounter sets the expected Invocation (Frame) counter value.
func (s *GXDLMSSettings) SetExpectedInvocationCounter(value uint64) {
	s.expectedInvocationCounter = value
}

// CopyTo copies all settings to target.
func (s *GXDLMSSettings) CopyTo(target *GXDLMSSettings) {
	target.UseCustomChallenge = s.UseCustomChallenge
	target.StartingBlockIndex = s.StartingBlockIndex
	target.DLMSVersion = s.DLMSVersion
	target.BlockIndex = s.BlockIndex
	target.isServer = s.isServer
	target.useLogicalNameReferencing = s.useLogicalNameReferencing
	target.ClientAddress = s.ClientAddress
	target.ServerAddress = s.ServerAddress
	target.PushClientAddress = s.PushClientAddress
	target.ServerAddressSize = s.ServerAddressSize
	target.Authentication = s.Authentication
	target.Password = s.Password
	target.ProposedConformance = s.ProposedConformance
	target.invokeID = s.invokeID
	target.LongInvokeID = s.LongInvokeID
	target.Priority = s.Priority
	target.ServiceClass = s.ServiceClass
	target.ctoSChallenge = s.ctoSChallenge
	target.stoCChallenge = s.stoCChallenge
	target.SenderFrame = s.SenderFrame
	target.ReceiverFrame = s.ReceiverFrame
	target.SetSourceSystemTitle(s.SourceSystemTitle())
	target.Kek = s.Kek
	target.Count = s.Count
	target.Index = s.Index
	target.maxReceivePDUSize = s.maxReceivePDUSize
	target.maxServerPDUSize = s.maxServerPDUSize
	target.ProposedConformance = s.ProposedConformance
	target.NegotiatedConformance = s.NegotiatedConformance
	// TODO: Copy Cipher when implemented
	// if s.Cipher != nil && target.Cipher != nil {
	//     s.Cipher.(*GXCiphering).CopyTo(target.Cipher.(*GXCiphering))
	// }
	target.UserID = s.UserID
	target.UseUtc2NormalTime = s.UseUtc2NormalTime
	target.gbtWindowSize = s.gbtWindowSize
	// TODO: Copy Objects when implemented
	// target.Objects.Clear()
	// target.Objects.AddRange(s.Objects)
	target.Hdlc.SetMaxInfoRX(s.Hdlc.MaxInfoRX())
	target.Hdlc.SetMaxInfoTX(s.Hdlc.MaxInfoTX())
	target.Hdlc.SetWindowSizeRX(s.Hdlc.WindowSizeRX())
	target.Hdlc.SetWindowSizeTX(s.Hdlc.WindowSizeTX())
	// TODO: Copy Gateway when implemented
	// if s.Gateway != nil {
	//     target.Gateway = &GXDLMSGateway{}
	//     target.Gateway.NetworkId = s.Gateway.NetworkId
	//     target.Gateway.PhysicalDeviceAddress = s.Gateway.PhysicalDeviceAddress
	// }
}

// UpdateSecurity updates security settings from security setup object.
func (s *GXDLMSSettings) UpdateSecurity(systemTitle []byte, ss interface{}) {
	// TODO: Implement when GXDLMSSecuritySetup is available
	// This method updates cipher settings, certificates, and invocation counter
	// from the security setup object
}

// UpdateSecuritySettings updates security settings from assigned association.
func (s *GXDLMSSettings) UpdateSecuritySettings(systemTitle []byte) {
	/*
		ln, ok := s.assignedAssociation.(*objects.GXDLMSAssociationLogicalName)
		if ok {
			// Update security settings.
			if ln.SecuritySetupReference != "" &&
				(ln.ApplicationContextName.ContextID == enums.ApplicationContextNameLogicalNameWithCiphering ||
					ln.AuthenticationMechanismName.MechanismId == enums.AuthenticationHighGMAC ||
					ln.AuthenticationMechanismName.MechanismId == enums.AuthenticationHighECDSA) {
				ss := ln.ObjectList.FindByLN(enums.ObjectTypeSecuritySetup, ln.SecuritySetupReference)
				s.UpdateSecurity(systemTitle, ss.(*objects.GXDLMSSecuritySetup))
			} else {
				ss := ln.ObjectList.FindByLN(enums.ObjectTypeSecuritySetup, ln.SecuritySetupReference)
				s.UpdateSecurity(systemTitle, ss.(*objects.GXDLMSSecuritySetup))
			}
		}
	*/
}

// Crypt encrypts or decrypts data using external Hardware Security Module.
func (s *GXDLMSSettings) Crypt(certificateType enums.CertificateType, data []byte, encrypt bool, keyType interface{}) []byte {
	// TODO: Implement when CryptoNotifier is available
	// if s.CryptoNotifier != nil {
	//     // Create args and call crypto notifier
	// }
	return nil
}

// GetKey gets the encryption/signing key.
func (s *GXDLMSSettings) GetKey(certificateType interface{}, systemTitle []byte, encrypt bool) any {
	// TODO: Implement when certificate types and security setup are available
	// This method retrieves signing or key agreement keys from cipher or security setup
	return nil
}

// IsServer indicates if this is server or client
func (s *GXDLMSSettings) IsServer() bool {
	return s.isServer
}
