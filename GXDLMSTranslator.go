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
	"errors"
	"strings"

	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// This class is used to translate DLMS frame or PDU to xml.
type GXDLMSTranslator struct {
	systemTitle []byte

	sAck uint16
	rAck uint16

	sending bool

	sSendSequence uint8

	sReceiveSequence uint8

	rSendSequence uint8

	rReceiveSequence uint8

	// Sending data in multiple frames.
	multipleFrames bool

	// If only PDUs are shown and PDU is received on parts.
	pduFrames types.GXByteBuffer

	outputType enums.TranslatorOutputType

	tags map[int]string

	tagsByName map[string]int

	// Are comments added.
	Comments bool

	// Is only PDU shown when data is parsed with MessageToXml
	PduOnly bool

	// Is only complete PDU parsed and shown.
	CompletePdu bool

	// Are numeric values shown as hex.
	Hex bool

	// Is string serialized as hex.
	ShowStringAsHex bool

	// Is XML declaration skipped.
	OmitXmlDeclaration bool

	// Is XML name space skipped.
	OmitXmlNameSpace bool

	// Used security.
	Security enums.Security

	// Used security suite.
	SecuritySuite enums.SecuritySuite

	// Invocation Counter.
	InvocationCounter uint32

	// Is General Protection used.
	UseGeneralProtection bool

	// Used standard.
	Standard enums.Standard

	// Public/private key signing key pair.
	// Private key is for the initializer and Public key is for the target.
	SigningKeyPair *types.GXKeyValuePair[*types.GXPublicKey, *types.GXPrivateKey]

	// Public/private key TLS key pair.
	// Private key is for the initializer and Public key is for the target.
	TlsKeyPair *types.GXKeyValuePair[*types.GXPublicKey, *types.GXPrivateKey]
	// Ephemeral key pair.
	EphemeralKeyPair *types.GXKeyValuePair[*types.GXPublicKey, *types.GXPrivateKey]

	// Public/private key key agreement key pair.
	KeyAgreementKeyPair *types.GXKeyValuePair[*types.GXPublicKey, *types.GXPrivateKey]

	serverSystemTitle []byte
	dedicatedKey      []byte
	blockCipherKey    []byte
	authenticationKey []byte
}

func ErrorCodeToString(type_ enums.TranslatorOutputType, value enums.ErrorCode) (string, error) {
	if type_ == enums.TranslatorOutputTypeStandardXML {
		return standardErrorCodeToString(value)
	} else {
		return simpleErrorCodeToString(value)
	}
}

func (g *GXDLMSTranslator) OutputType() enums.TranslatorOutputType {
	return g.outputType
}

// SystemTitle returns the system title.
func (g *GXDLMSTranslator) SystemTitle() []byte {
	return g.systemTitle
}

// SetSystemTitle sets the system title.
func (g *GXDLMSTranslator) SetSystemTitle(value []byte) error {
	if value != nil {
		if len(value) == 0 {
			value = nil
		} else if len(value) != 8 {
			return errors.New("Invalid system title. System title size is 8 bytes.")
		}
	}
	g.systemTitle = value
	return nil
}

// ServerSystemTitle returns the server system title.
func (g *GXDLMSTranslator) ServerSystemTitle() []byte {
	return g.serverSystemTitle
}

// SetServerSystemTitle sets the server system title.
func (g *GXDLMSTranslator) SetServerSystemTitle(value []byte) error {
	if value != nil {
		if len(value) == 0 {
			value = nil
		} else if len(value) != 8 {
			return errors.New("Invalid server system title. Server system title size is 8 bytes.")
		}
	}
	g.serverSystemTitle = value
	return nil
}

// DedicatedKey returns the dedicated key.
func (g *GXDLMSTranslator) DedicatedKey() []byte {
	return g.dedicatedKey
}

// SetDedicatedKey sets the dedicated key.
func (g *GXDLMSTranslator) SetDedicatedKey(value []byte) error {
	if value != nil {
		if len(value) == 0 {
			value = nil
		} else if !((g.SecuritySuite != enums.SecuritySuite2 && len(value) == 16) || (g.SecuritySuite == enums.SecuritySuite2 && len(value) == 32)) {
			return errors.New("Invalid dedicated key. Dedicated key size is 16 bytes.")
		}
	}
	g.dedicatedKey = value
	return nil
}

// BlockCipherKey returns the block cipher key.
func (g *GXDLMSTranslator) BlockCipherKey() []byte {
	return g.blockCipherKey
}

// SetBlockCipherKey sets the block cipher key.
func (g *GXDLMSTranslator) SetBlockCipherKey(value []byte) error {
	if value != nil {
		if len(value) == 0 {
			value = nil
		} else if !((g.SecuritySuite != enums.SecuritySuite2 && len(value) == 16) || (g.SecuritySuite == enums.SecuritySuite2 && len(value) == 32)) {
			return errors.New("Invalid block cipher key. Block cipher key size is 16 bytes.")
		}
	}
	g.blockCipherKey = value
	return nil
}

// AuthenticationKey returns the authentication key.
func (g *GXDLMSTranslator) AuthenticationKey() []byte {
	return g.authenticationKey
}

// SetAuthenticationKey sets the authentication key.
func (g *GXDLMSTranslator) SetAuthenticationKey(value []byte) error {
	if value != nil {
		if len(value) == 0 {
			value = nil
		} else if !((g.SecuritySuite != enums.SecuritySuite2 && len(value) == 16) || (g.SecuritySuite == enums.SecuritySuite2 && len(value) == 32)) {
			return errors.New("Invalid authentication key. Authentication key size is 16 bytes.")
		}
	}
	g.authenticationKey = value
	return nil
}

// GetTags returns the get all tags.
//
// Parameters:
//
//	type: Output type.
//	list: List of tags by ID.
//	tagsByName: List of tags by name.
func (g *GXDLMSTranslator) GetTags(type_ enums.TranslatorOutputType, list map[int]string, tagsByName map[string]int) {
	if type_ == enums.TranslatorOutputTypeSimpleXML {
		simpleGetGeneralTags(type_, list)
		simpleGetSnTags(type_, list)
		simpleGetLnTags(type_, list)
		simpleGetGloTags(type_, list)
		simpleGetDedTags(type_, list)
		simpleGetTranslatorTags(type_, list)
		simpleGetDataTypeTags(list)
		simpleGetPlcTags(list)
	} else {
		standardGetGeneralTags(type_, list)
		standardGetSnTags(type_, list)
		standardGetLnTags(type_, list)
		standardGetGloTags(type_, list)
		standardGetDedTags(type_, list)
		standardGetTranslatorTags(type_, list)
		standardGetDataTypeTags(list)
		standardGetPlcTags(list)
	}
	// Simple is not case sensitive.
	lowercase := type_ == enums.TranslatorOutputTypeSimpleXML
	for k, v := range list {
		str := v
		if lowercase {
			str = strings.ToLower(str)
		}
		_, ok := g.tagsByName[str]
		if !ok {
			g.tagsByName[str] = k
		}
	}
}

func (g *GXDLMSTranslator) UpdateAddress(settings *settings.GXDLMSSettings, msg *GXDLMSTranslatorMessage) {
	reply := false
	switch msg.Command {
	case enums.CommandReadRequest:
	case enums.CommandWriteRequest:
	case enums.CommandGetRequest:
	case enums.CommandSetRequest:
	case enums.CommandMethodRequest:
	case enums.CommandSnrm:
	case enums.CommandAarq:
	case enums.CommandDisconnectRequest:
	case enums.CommandReleaseRequest:
	case enums.CommandAccessRequest:
	case enums.CommandGloGetRequest:
	case enums.CommandGloSetRequest:
	case enums.CommandGloMethodRequest:
	case enums.CommandGloInitiateRequest:
	case enums.CommandGloReadRequest:
	case enums.CommandGloWriteRequest:
	case enums.CommandDedInitiateRequest:
	case enums.CommandDedReadRequest:
	case enums.CommandDedWriteRequest:
	case enums.CommandDedGetRequest:
	case enums.CommandDedSetRequest:
	case enums.CommandDedMethodRequest:
	case enums.CommandGatewayRequest:
	case enums.CommandDiscoverRequest:
	case enums.CommandRegisterRequest:
	case enums.CommandPingRequest:
		reply = false
		break
	case enums.CommandReadResponse:
	case enums.CommandWriteResponse:
	case enums.CommandGetResponse:
	case enums.CommandSetResponse:
	case enums.CommandMethodResponse:
	case enums.CommandDisconnectMode:
	case enums.CommandUnacceptableFrame:
	case enums.CommandUa:
	case enums.CommandAare:
	case enums.CommandReleaseResponse:
	case enums.CommandConfirmedServiceError:
	case enums.CommandExceptionResponse:
	case enums.CommandAccessResponse:
	case enums.CommandDataNotification:
	case enums.CommandGloGetResponse:
	case enums.CommandGloSetResponse:
	case enums.CommandGloEventNotification:
	case enums.CommandGloMethodResponse:
	case enums.CommandGloInitiateResponse:
	case enums.CommandGloReadResponse:
	case enums.CommandGloWriteResponse:
	case enums.CommandGloConfirmedServiceError:
	case enums.CommandGloInformationReport:
	case enums.CommandInformationReport:
	case enums.CommandEventNotification:
	case enums.CommandDedInitiateResponse:
	case enums.CommandDedReadResponse:
	case enums.CommandDedWriteResponse:
	case enums.CommandDedConfirmedServiceError:
	case enums.CommandDedUnconfirmedWriteRequest:
	case enums.CommandDedInformationReport:
	case enums.CommandDedGetResponse:
	case enums.CommandDedSetResponse:
	case enums.CommandDedEventNotification:
	case enums.CommandDedMethodResponse:
	case enums.CommandGatewayResponse:
	case enums.CommandDiscoverReport:
	case enums.CommandPingResponse:
		reply = true
		break
	}
	if reply {
		msg.TargetAddress = settings.ClientAddress
		msg.SourceAddress = settings.ServerAddress
	} else {
		msg.SourceAddress = settings.ClientAddress
		msg.TargetAddress = settings.ServerAddress
	}
}
