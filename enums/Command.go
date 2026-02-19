package enums

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
	"strings"

	"github.com/Gurux/gxcommon-go"
)

// Command DLMS command enumeration.
type Command int

const (
	// CommandNone defines that the no command to execute.
	CommandNone Command = iota
	// CommandInitiateRequest defines that the initiate request.
	CommandInitiateRequest = 0x1
	// CommandInitiateResponse defines that the initiate response.
	CommandInitiateResponse = 0x8
	// CommandReadRequest defines that the read request.
	CommandReadRequest = 0x5
	// CommandReadResponse defines that the read response.
	CommandReadResponse = 0xC
	// CommandWriteRequest defines that the write request.
	CommandWriteRequest = 0x6
	// CommandWriteResponse defines that the write response.
	CommandWriteResponse = 0xD
	// CommandGetRequest defines that the get request.
	CommandGetRequest = 0xC0
	// CommandGetResponse defines that the get response.
	CommandGetResponse = 0xC4
	// CommandSetRequest defines that the set request.
	CommandSetRequest = 0xC1
	// CommandSetResponse defines that the set response.
	CommandSetResponse = 0xC5
	// CommandMethodRequest defines that the action request.
	CommandMethodRequest = 0xC3
	// CommandMethodResponse defines that the action response.
	CommandMethodResponse = 0xC7
	// CommandDisconnectMode defines that the hDLC Disconnect Mode.
	CommandDisconnectMode = 0x1f
	// CommandUnacceptableFrame defines that the hDLC Unacceptable Frame.
	CommandUnacceptableFrame = 0x97
	// CommandSnrm defines that the hDLC SNRM request.
	CommandSnrm = 0x93
	// CommandUa defines that the hDLC UA request.
	CommandUa = 0x73
	// CommandAarq defines that the aARQ request.
	CommandAarq = 0x60
	// CommandAare defines that the aARE request.
	CommandAare = 0x61
	// CommandDisconnectRequest defines that the disconnect request for HDLC framing. (DISC)
	CommandDisconnectRequest = 0x53
	// CommandReleaseRequest defines that the release request.
	CommandReleaseRequest = 0x62
	// CommandReleaseResponse defines that the release response.
	CommandReleaseResponse = 0x63
	// CommandConfirmedServiceError defines that the confirmed Service Error.
	CommandConfirmedServiceError = 0x0E
	// CommandExceptionResponse defines that the exception Response.
	CommandExceptionResponse = 0xD8
	// CommandGeneralBlockTransfer defines that the general Block Transfer.
	CommandGeneralBlockTransfer = 0xE0
	// CommandAccessRequest defines that the access Request.
	CommandAccessRequest = 0xD9
	// CommandAccessResponse defines that the access Response.
	CommandAccessResponse = 0xDA
	// CommandDataNotification defines that the data Notification request.
	CommandDataNotification = 0x0F
	// CommandDataNotificationConfirm defines that the data notification confirm.
	CommandDataNotificationConfirm = 16
	// CommandGloGetRequest defines that the glo get request.
	CommandGloGetRequest = 0xC8
	// CommandGloGetResponse defines that the glo get response.
	CommandGloGetResponse = 0xCC
	// CommandGloSetRequest defines that the glo set request.
	CommandGloSetRequest = 0xC9
	// CommandGloSetResponse defines that the glo set response.
	CommandGloSetResponse = 0xCD
	// CommandGloEventNotification defines that the glo event notification.
	CommandGloEventNotification = 0xCA
	// CommandGloMethodRequest defines that the glo method request.
	CommandGloMethodRequest = 0xCB
	// CommandGloMethodResponse defines that the glo method response.
	CommandGloMethodResponse = 0xCF
	// CommandGloInitiateRequest defines that the glo Initiate request.
	CommandGloInitiateRequest = 0x21
	// CommandGloReadRequest defines that the glo read request.
	CommandGloReadRequest = 37
	// CommandGloWriteRequest defines that the glo write request.
	CommandGloWriteRequest = 38
	// CommandGloInitiateResponse defines that the glo Initiate response.
	CommandGloInitiateResponse = 0x28
	// CommandGloReadResponse defines that the glo read response.
	CommandGloReadResponse = 44
	// CommandGloWriteResponse defines that the glo write response.
	CommandGloWriteResponse = 45
	// CommandGloConfirmedServiceError defines that the glo confirmed service error.
	CommandGloConfirmedServiceError = 46
	// CommandGloInformationReport defines that the glo information report.
	CommandGloInformationReport = 56
	// CommandGeneralGloCiphering defines that the general GLO ciphering.
	CommandGeneralGloCiphering = 0xDB
	// CommandGeneralDedCiphering defines that the general DED ciphering.
	CommandGeneralDedCiphering = 0xDC
	// CommandGeneralCiphering defines that the general ciphering.
	CommandGeneralCiphering = 0xDD
	// CommandGeneralSigning defines that the general signing.
	CommandGeneralSigning = 0xDF
	// CommandInformationReport defines that the information Report request.
	CommandInformationReport = 0x18
	// CommandEventNotification defines that the event Notification request.
	CommandEventNotification = 0xC2
	// CommandDedInitiateRequest defines that the ded initiate request.
	CommandDedInitiateRequest = 65
	// CommandDedReadRequest defines that the ded read request.
	CommandDedReadRequest = 69
	// CommandDedWriteRequest defines that the ded write request.
	CommandDedWriteRequest = 70
	// CommandDedInitiateResponse defines that the ded initiate response.
	CommandDedInitiateResponse = 72
	// CommandDedReadResponse defines that the ded read response.
	CommandDedReadResponse = 76
	// CommandDedWriteResponse defines that the ded write response.
	CommandDedWriteResponse = 77
	// CommandDedConfirmedServiceError defines that the ded confirmed service error.
	CommandDedConfirmedServiceError = 78
	// CommandDedUnconfirmedWriteRequest defines that the ded confirmed write request.
	CommandDedUnconfirmedWriteRequest = 86
	// CommandDedInformationReport defines that the ded information report.
	CommandDedInformationReport = 88
	// CommandDedGetRequest defines that the ded get request.
	CommandDedGetRequest = 0xD0
	// CommandDedGetResponse defines that the ded get response.
	CommandDedGetResponse = 0xD4
	// CommandDedSetRequest defines that the ded set request.
	CommandDedSetRequest = 0xD1
	// CommandDedSetResponse defines that the ded set response.
	CommandDedSetResponse = 0xD5
	// CommandDedEventNotification defines that the ded event notification.
	CommandDedEventNotification = 0xD2
	// CommandDedMethodRequest defines that the ded method request.
	CommandDedMethodRequest = 0xD3
	// CommandDedMethodResponse defines that the ded method response.
	CommandDedMethodResponse = 0xD7
	// CommandGatewayRequest defines that the request message from client to gateway.
	CommandGatewayRequest = 0xE6
	// CommandGatewayResponse defines that the response message from gateway to client.
	CommandGatewayResponse = 0xE7
	// CommandDiscoverRequest defines that the pLC discover request.
	CommandDiscoverRequest = 0x1D
	// CommandDiscoverReport defines that the pLC discover report.
	CommandDiscoverReport = 0x1E
	// CommandRegisterRequest defines that the pLC register request.
	CommandRegisterRequest = 0x1C
	// CommandPingRequest defines that the pLC ping request.
	CommandPingRequest = 0x19
	// CommandPingResponse defines that the pLC ping response.
	CommandPingResponse = 0x1A
)

// CommandParse converts the given string into a Command value.
//
// It returns the corresponding Command constant if the string matches
// a known level name, or an error if the input is invalid.
func CommandParse(value string) (Command, error) {
	var ret Command
	var err error
	switch {
	case strings.EqualFold(value, "None"):
		ret = CommandNone
	case strings.EqualFold(value, "InitiateRequest"):
		ret = CommandInitiateRequest
	case strings.EqualFold(value, "InitiateResponse"):
		ret = CommandInitiateResponse
	case strings.EqualFold(value, "ReadRequest"):
		ret = CommandReadRequest
	case strings.EqualFold(value, "ReadResponse"):
		ret = CommandReadResponse
	case strings.EqualFold(value, "WriteRequest"):
		ret = CommandWriteRequest
	case strings.EqualFold(value, "WriteResponse"):
		ret = CommandWriteResponse
	case strings.EqualFold(value, "GetRequest"):
		ret = CommandGetRequest
	case strings.EqualFold(value, "GetResponse"):
		ret = CommandGetResponse
	case strings.EqualFold(value, "SetRequest"):
		ret = CommandSetRequest
	case strings.EqualFold(value, "SetResponse"):
		ret = CommandSetResponse
	case strings.EqualFold(value, "MethodRequest"):
		ret = CommandMethodRequest
	case strings.EqualFold(value, "MethodResponse"):
		ret = CommandMethodResponse
	case strings.EqualFold(value, "DisconnectMode"):
		ret = CommandDisconnectMode
	case strings.EqualFold(value, "UnacceptableFrame"):
		ret = CommandUnacceptableFrame
	case strings.EqualFold(value, "Snrm"):
		ret = CommandSnrm
	case strings.EqualFold(value, "Ua"):
		ret = CommandUa
	case strings.EqualFold(value, "Aarq"):
		ret = CommandAarq
	case strings.EqualFold(value, "Aare"):
		ret = CommandAare
	case strings.EqualFold(value, "DisconnectRequest"):
		ret = CommandDisconnectRequest
	case strings.EqualFold(value, "ReleaseRequest"):
		ret = CommandReleaseRequest
	case strings.EqualFold(value, "ReleaseResponse"):
		ret = CommandReleaseResponse
	case strings.EqualFold(value, "ConfirmedServiceError"):
		ret = CommandConfirmedServiceError
	case strings.EqualFold(value, "ExceptionResponse"):
		ret = CommandExceptionResponse
	case strings.EqualFold(value, "GeneralBlockTransfer"):
		ret = CommandGeneralBlockTransfer
	case strings.EqualFold(value, "AccessRequest"):
		ret = CommandAccessRequest
	case strings.EqualFold(value, "AccessResponse"):
		ret = CommandAccessResponse
	case strings.EqualFold(value, "DataNotification"):
		ret = CommandDataNotification
	case strings.EqualFold(value, "DataNotificationConfirm"):
		ret = CommandDataNotificationConfirm
	case strings.EqualFold(value, "GloGetRequest"):
		ret = CommandGloGetRequest
	case strings.EqualFold(value, "GloGetResponse"):
		ret = CommandGloGetResponse
	case strings.EqualFold(value, "GloSetRequest"):
		ret = CommandGloSetRequest
	case strings.EqualFold(value, "GloSetResponse"):
		ret = CommandGloSetResponse
	case strings.EqualFold(value, "GloEventNotification"):
		ret = CommandGloEventNotification
	case strings.EqualFold(value, "GloMethodRequest"):
		ret = CommandGloMethodRequest
	case strings.EqualFold(value, "GloMethodResponse"):
		ret = CommandGloMethodResponse
	case strings.EqualFold(value, "GloInitiateRequest"):
		ret = CommandGloInitiateRequest
	case strings.EqualFold(value, "GloReadRequest"):
		ret = CommandGloReadRequest
	case strings.EqualFold(value, "GloWriteRequest"):
		ret = CommandGloWriteRequest
	case strings.EqualFold(value, "GloInitiateResponse"):
		ret = CommandGloInitiateResponse
	case strings.EqualFold(value, "GloReadResponse"):
		ret = CommandGloReadResponse
	case strings.EqualFold(value, "GloWriteResponse"):
		ret = CommandGloWriteResponse
	case strings.EqualFold(value, "GloConfirmedServiceError"):
		ret = CommandGloConfirmedServiceError
	case strings.EqualFold(value, "GloInformationReport"):
		ret = CommandGloInformationReport
	case strings.EqualFold(value, "GeneralGloCiphering"):
		ret = CommandGeneralGloCiphering
	case strings.EqualFold(value, "GeneralDedCiphering"):
		ret = CommandGeneralDedCiphering
	case strings.EqualFold(value, "GeneralCiphering"):
		ret = CommandGeneralCiphering
	case strings.EqualFold(value, "GeneralSigning"):
		ret = CommandGeneralSigning
	case strings.EqualFold(value, "InformationReport"):
		ret = CommandInformationReport
	case strings.EqualFold(value, "EventNotification"):
		ret = CommandEventNotification
	case strings.EqualFold(value, "DedInitiateRequest"):
		ret = CommandDedInitiateRequest
	case strings.EqualFold(value, "DedReadRequest"):
		ret = CommandDedReadRequest
	case strings.EqualFold(value, "DedWriteRequest"):
		ret = CommandDedWriteRequest
	case strings.EqualFold(value, "DedInitiateResponse"):
		ret = CommandDedInitiateResponse
	case strings.EqualFold(value, "DedReadResponse"):
		ret = CommandDedReadResponse
	case strings.EqualFold(value, "DedWriteResponse"):
		ret = CommandDedWriteResponse
	case strings.EqualFold(value, "DedConfirmedServiceError"):
		ret = CommandDedConfirmedServiceError
	case strings.EqualFold(value, "DedUnconfirmedWriteRequest"):
		ret = CommandDedUnconfirmedWriteRequest
	case strings.EqualFold(value, "DedInformationReport"):
		ret = CommandDedInformationReport
	case strings.EqualFold(value, "DedGetRequest"):
		ret = CommandDedGetRequest
	case strings.EqualFold(value, "DedGetResponse"):
		ret = CommandDedGetResponse
	case strings.EqualFold(value, "DedSetRequest"):
		ret = CommandDedSetRequest
	case strings.EqualFold(value, "DedSetResponse"):
		ret = CommandDedSetResponse
	case strings.EqualFold(value, "DedEventNotification"):
		ret = CommandDedEventNotification
	case strings.EqualFold(value, "DedMethodRequest"):
		ret = CommandDedMethodRequest
	case strings.EqualFold(value, "DedMethodResponse"):
		ret = CommandDedMethodResponse
	case strings.EqualFold(value, "GatewayRequest"):
		ret = CommandGatewayRequest
	case strings.EqualFold(value, "GatewayResponse"):
		ret = CommandGatewayResponse
	case strings.EqualFold(value, "DiscoverRequest"):
		ret = CommandDiscoverRequest
	case strings.EqualFold(value, "DiscoverReport"):
		ret = CommandDiscoverReport
	case strings.EqualFold(value, "RegisterRequest"):
		ret = CommandRegisterRequest
	case strings.EqualFold(value, "PingRequest"):
		ret = CommandPingRequest
	case strings.EqualFold(value, "PingResponse"):
		ret = CommandPingResponse
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the Command.
// It satisfies fmt.Stringer.
func (g Command) String() string {
	var ret string
	switch g {
	case CommandNone:
		ret = "None"
	case CommandInitiateRequest:
		ret = "InitiateRequest"
	case CommandInitiateResponse:
		ret = "InitiateResponse"
	case CommandReadRequest:
		ret = "ReadRequest"
	case CommandReadResponse:
		ret = "ReadResponse"
	case CommandWriteRequest:
		ret = "WriteRequest"
	case CommandWriteResponse:
		ret = "WriteResponse"
	case CommandGetRequest:
		ret = "GetRequest"
	case CommandGetResponse:
		ret = "GetResponse"
	case CommandSetRequest:
		ret = "SetRequest"
	case CommandSetResponse:
		ret = "SetResponse"
	case CommandMethodRequest:
		ret = "MethodRequest"
	case CommandMethodResponse:
		ret = "MethodResponse"
	case CommandDisconnectMode:
		ret = "DisconnectMode"
	case CommandUnacceptableFrame:
		ret = "UnacceptableFrame"
	case CommandSnrm:
		ret = "Snrm"
	case CommandUa:
		ret = "Ua"
	case CommandAarq:
		ret = "Aarq"
	case CommandAare:
		ret = "Aare"
	case CommandDisconnectRequest:
		ret = "DisconnectRequest"
	case CommandReleaseRequest:
		ret = "ReleaseRequest"
	case CommandReleaseResponse:
		ret = "ReleaseResponse"
	case CommandConfirmedServiceError:
		ret = "ConfirmedServiceError"
	case CommandExceptionResponse:
		ret = "ExceptionResponse"
	case CommandGeneralBlockTransfer:
		ret = "GeneralBlockTransfer"
	case CommandAccessRequest:
		ret = "AccessRequest"
	case CommandAccessResponse:
		ret = "AccessResponse"
	case CommandDataNotification:
		ret = "DataNotification"
	case CommandDataNotificationConfirm:
		ret = "DataNotificationConfirm"
	case CommandGloGetRequest:
		ret = "GloGetRequest"
	case CommandGloGetResponse:
		ret = "GloGetResponse"
	case CommandGloSetRequest:
		ret = "GloSetRequest"
	case CommandGloSetResponse:
		ret = "GloSetResponse"
	case CommandGloEventNotification:
		ret = "GloEventNotification"
	case CommandGloMethodRequest:
		ret = "GloMethodRequest"
	case CommandGloMethodResponse:
		ret = "GloMethodResponse"
	case CommandGloInitiateRequest:
		ret = "GloInitiateRequest"
	case CommandGloReadRequest:
		ret = "GloReadRequest"
	case CommandGloWriteRequest:
		ret = "GloWriteRequest"
	case CommandGloInitiateResponse:
		ret = "GloInitiateResponse"
	case CommandGloReadResponse:
		ret = "GloReadResponse"
	case CommandGloWriteResponse:
		ret = "GloWriteResponse"
	case CommandGloConfirmedServiceError:
		ret = "GloConfirmedServiceError"
	case CommandGloInformationReport:
		ret = "GloInformationReport"
	case CommandGeneralGloCiphering:
		ret = "GeneralGloCiphering"
	case CommandGeneralDedCiphering:
		ret = "GeneralDedCiphering"
	case CommandGeneralCiphering:
		ret = "GeneralCiphering"
	case CommandGeneralSigning:
		ret = "GeneralSigning"
	case CommandInformationReport:
		ret = "InformationReport"
	case CommandEventNotification:
		ret = "EventNotification"
	case CommandDedInitiateRequest:
		ret = "DedInitiateRequest"
	case CommandDedReadRequest:
		ret = "DedReadRequest"
	case CommandDedWriteRequest:
		ret = "DedWriteRequest"
	case CommandDedInitiateResponse:
		ret = "DedInitiateResponse"
	case CommandDedReadResponse:
		ret = "DedReadResponse"
	case CommandDedWriteResponse:
		ret = "DedWriteResponse"
	case CommandDedConfirmedServiceError:
		ret = "DedConfirmedServiceError"
	case CommandDedUnconfirmedWriteRequest:
		ret = "DedUnconfirmedWriteRequest"
	case CommandDedInformationReport:
		ret = "DedInformationReport"
	case CommandDedGetRequest:
		ret = "DedGetRequest"
	case CommandDedGetResponse:
		ret = "DedGetResponse"
	case CommandDedSetRequest:
		ret = "DedSetRequest"
	case CommandDedSetResponse:
		ret = "DedSetResponse"
	case CommandDedEventNotification:
		ret = "DedEventNotification"
	case CommandDedMethodRequest:
		ret = "DedMethodRequest"
	case CommandDedMethodResponse:
		ret = "DedMethodResponse"
	case CommandGatewayRequest:
		ret = "GatewayRequest"
	case CommandGatewayResponse:
		ret = "GatewayResponse"
	case CommandDiscoverRequest:
		ret = "DiscoverRequest"
	case CommandDiscoverReport:
		ret = "DiscoverReport"
	case CommandRegisterRequest:
		ret = "RegisterRequest"
	case CommandPingRequest:
		ret = "PingRequest"
	case CommandPingResponse:
		ret = "PingResponse"
	}
	return ret
}

// AllCommand returns a slice containing all defined Command values.
func AllCommand() []Command {
	return []Command{
		CommandNone,
		CommandInitiateRequest,
		CommandInitiateResponse,
		CommandReadRequest,
		CommandReadResponse,
		CommandWriteRequest,
		CommandWriteResponse,
		CommandGetRequest,
		CommandGetResponse,
		CommandSetRequest,
		CommandSetResponse,
		CommandMethodRequest,
		CommandMethodResponse,
		CommandDisconnectMode,
		CommandUnacceptableFrame,
		CommandSnrm,
		CommandUa,
		CommandAarq,
		CommandAare,
		CommandDisconnectRequest,
		CommandReleaseRequest,
		CommandReleaseResponse,
		CommandConfirmedServiceError,
		CommandExceptionResponse,
		CommandGeneralBlockTransfer,
		CommandAccessRequest,
		CommandAccessResponse,
		CommandDataNotification,
		CommandDataNotificationConfirm,
		CommandGloGetRequest,
		CommandGloGetResponse,
		CommandGloSetRequest,
		CommandGloSetResponse,
		CommandGloEventNotification,
		CommandGloMethodRequest,
		CommandGloMethodResponse,
		CommandGloInitiateRequest,
		CommandGloReadRequest,
		CommandGloWriteRequest,
		CommandGloInitiateResponse,
		CommandGloReadResponse,
		CommandGloWriteResponse,
		CommandGloConfirmedServiceError,
		CommandGloInformationReport,
		CommandGeneralGloCiphering,
		CommandGeneralDedCiphering,
		CommandGeneralCiphering,
		CommandGeneralSigning,
		CommandInformationReport,
		CommandEventNotification,
		CommandDedInitiateRequest,
		CommandDedReadRequest,
		CommandDedWriteRequest,
		CommandDedInitiateResponse,
		CommandDedReadResponse,
		CommandDedWriteResponse,
		CommandDedConfirmedServiceError,
		CommandDedUnconfirmedWriteRequest,
		CommandDedInformationReport,
		CommandDedGetRequest,
		CommandDedGetResponse,
		CommandDedSetRequest,
		CommandDedSetResponse,
		CommandDedEventNotification,
		CommandDedMethodRequest,
		CommandDedMethodResponse,
		CommandGatewayRequest,
		CommandGatewayResponse,
		CommandDiscoverRequest,
		CommandDiscoverReport,
		CommandRegisterRequest,
		CommandPingRequest,
		CommandPingResponse,
	}
}
