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
	switch strings.ToUpper(value) {
	case "NONE":
		ret = CommandNone
	case "INITIATEREQUEST":
		ret = CommandInitiateRequest
	case "INITIATERESPONSE":
		ret = CommandInitiateResponse
	case "READREQUEST":
		ret = CommandReadRequest
	case "READRESPONSE":
		ret = CommandReadResponse
	case "WRITEREQUEST":
		ret = CommandWriteRequest
	case "WRITERESPONSE":
		ret = CommandWriteResponse
	case "GETREQUEST":
		ret = CommandGetRequest
	case "GETRESPONSE":
		ret = CommandGetResponse
	case "SETREQUEST":
		ret = CommandSetRequest
	case "SETRESPONSE":
		ret = CommandSetResponse
	case "METHODREQUEST":
		ret = CommandMethodRequest
	case "METHODRESPONSE":
		ret = CommandMethodResponse
	case "DISCONNECTMODE":
		ret = CommandDisconnectMode
	case "UNACCEPTABLEFRAME":
		ret = CommandUnacceptableFrame
	case "SNRM":
		ret = CommandSnrm
	case "UA":
		ret = CommandUa
	case "AARQ":
		ret = CommandAarq
	case "AARE":
		ret = CommandAare
	case "DISCONNECTREQUEST":
		ret = CommandDisconnectRequest
	case "RELEASEREQUEST":
		ret = CommandReleaseRequest
	case "RELEASERESPONSE":
		ret = CommandReleaseResponse
	case "CONFIRMEDSERVICEERROR":
		ret = CommandConfirmedServiceError
	case "EXCEPTIONRESPONSE":
		ret = CommandExceptionResponse
	case "GENERALBLOCKTRANSFER":
		ret = CommandGeneralBlockTransfer
	case "ACCESSREQUEST":
		ret = CommandAccessRequest
	case "ACCESSRESPONSE":
		ret = CommandAccessResponse
	case "DATANOTIFICATION":
		ret = CommandDataNotification
	case "DATANOTIFICATIONCONFIRM":
		ret = CommandDataNotificationConfirm
	case "GLOGETREQUEST":
		ret = CommandGloGetRequest
	case "GLOGETRESPONSE":
		ret = CommandGloGetResponse
	case "GLOSETREQUEST":
		ret = CommandGloSetRequest
	case "GLOSETRESPONSE":
		ret = CommandGloSetResponse
	case "GLOEVENTNOTIFICATION":
		ret = CommandGloEventNotification
	case "GLOMETHODREQUEST":
		ret = CommandGloMethodRequest
	case "GLOMETHODRESPONSE":
		ret = CommandGloMethodResponse
	case "GLOINITIATEREQUEST":
		ret = CommandGloInitiateRequest
	case "GLOREADREQUEST":
		ret = CommandGloReadRequest
	case "GLOWRITEREQUEST":
		ret = CommandGloWriteRequest
	case "GLOINITIATERESPONSE":
		ret = CommandGloInitiateResponse
	case "GLOREADRESPONSE":
		ret = CommandGloReadResponse
	case "GLOWRITERESPONSE":
		ret = CommandGloWriteResponse
	case "GLOCONFIRMEDSERVICEERROR":
		ret = CommandGloConfirmedServiceError
	case "GLOINFORMATIONREPORT":
		ret = CommandGloInformationReport
	case "GENERALGLOCIPHERING":
		ret = CommandGeneralGloCiphering
	case "GENERALDEDCIPHERING":
		ret = CommandGeneralDedCiphering
	case "GENERALCIPHERING":
		ret = CommandGeneralCiphering
	case "GENERALSIGNING":
		ret = CommandGeneralSigning
	case "INFORMATIONREPORT":
		ret = CommandInformationReport
	case "EVENTNOTIFICATION":
		ret = CommandEventNotification
	case "DEDINITIATEREQUEST":
		ret = CommandDedInitiateRequest
	case "DEDREADREQUEST":
		ret = CommandDedReadRequest
	case "DEDWRITEREQUEST":
		ret = CommandDedWriteRequest
	case "DEDINITIATERESPONSE":
		ret = CommandDedInitiateResponse
	case "DEDREADRESPONSE":
		ret = CommandDedReadResponse
	case "DEDWRITERESPONSE":
		ret = CommandDedWriteResponse
	case "DEDCONFIRMEDSERVICEERROR":
		ret = CommandDedConfirmedServiceError
	case "DEDUNCONFIRMEDWRITEREQUEST":
		ret = CommandDedUnconfirmedWriteRequest
	case "DEDINFORMATIONREPORT":
		ret = CommandDedInformationReport
	case "DEDGETREQUEST":
		ret = CommandDedGetRequest
	case "DEDGETRESPONSE":
		ret = CommandDedGetResponse
	case "DEDSETREQUEST":
		ret = CommandDedSetRequest
	case "DEDSETRESPONSE":
		ret = CommandDedSetResponse
	case "DEDEVENTNOTIFICATION":
		ret = CommandDedEventNotification
	case "DEDMETHODREQUEST":
		ret = CommandDedMethodRequest
	case "DEDMETHODRESPONSE":
		ret = CommandDedMethodResponse
	case "GATEWAYREQUEST":
		ret = CommandGatewayRequest
	case "GATEWAYRESPONSE":
		ret = CommandGatewayResponse
	case "DISCOVERREQUEST":
		ret = CommandDiscoverRequest
	case "DISCOVERREPORT":
		ret = CommandDiscoverReport
	case "REGISTERREQUEST":
		ret = CommandRegisterRequest
	case "PINGREQUEST":
		ret = CommandPingRequest
	case "PINGRESPONSE":
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
		ret = "NONE"
	case CommandInitiateRequest:
		ret = "INITIATEREQUEST"
	case CommandInitiateResponse:
		ret = "INITIATERESPONSE"
	case CommandReadRequest:
		ret = "READREQUEST"
	case CommandReadResponse:
		ret = "READRESPONSE"
	case CommandWriteRequest:
		ret = "WRITEREQUEST"
	case CommandWriteResponse:
		ret = "WRITERESPONSE"
	case CommandGetRequest:
		ret = "GETREQUEST"
	case CommandGetResponse:
		ret = "GETRESPONSE"
	case CommandSetRequest:
		ret = "SETREQUEST"
	case CommandSetResponse:
		ret = "SETRESPONSE"
	case CommandMethodRequest:
		ret = "METHODREQUEST"
	case CommandMethodResponse:
		ret = "METHODRESPONSE"
	case CommandDisconnectMode:
		ret = "DISCONNECTMODE"
	case CommandUnacceptableFrame:
		ret = "UNACCEPTABLEFRAME"
	case CommandSnrm:
		ret = "SNRM"
	case CommandUa:
		ret = "UA"
	case CommandAarq:
		ret = "AARQ"
	case CommandAare:
		ret = "AARE"
	case CommandDisconnectRequest:
		ret = "DISCONNECTREQUEST"
	case CommandReleaseRequest:
		ret = "RELEASEREQUEST"
	case CommandReleaseResponse:
		ret = "RELEASERESPONSE"
	case CommandConfirmedServiceError:
		ret = "CONFIRMEDSERVICEERROR"
	case CommandExceptionResponse:
		ret = "EXCEPTIONRESPONSE"
	case CommandGeneralBlockTransfer:
		ret = "GENERALBLOCKTRANSFER"
	case CommandAccessRequest:
		ret = "ACCESSREQUEST"
	case CommandAccessResponse:
		ret = "ACCESSRESPONSE"
	case CommandDataNotification:
		ret = "DATANOTIFICATION"
	case CommandDataNotificationConfirm:
		ret = "DATANOTIFICATIONCONFIRM"
	case CommandGloGetRequest:
		ret = "GLOGETREQUEST"
	case CommandGloGetResponse:
		ret = "GLOGETRESPONSE"
	case CommandGloSetRequest:
		ret = "GLOSETREQUEST"
	case CommandGloSetResponse:
		ret = "GLOSETRESPONSE"
	case CommandGloEventNotification:
		ret = "GLOEVENTNOTIFICATION"
	case CommandGloMethodRequest:
		ret = "GLOMETHODREQUEST"
	case CommandGloMethodResponse:
		ret = "GLOMETHODRESPONSE"
	case CommandGloInitiateRequest:
		ret = "GLOINITIATEREQUEST"
	case CommandGloReadRequest:
		ret = "GLOREADREQUEST"
	case CommandGloWriteRequest:
		ret = "GLOWRITEREQUEST"
	case CommandGloInitiateResponse:
		ret = "GLOINITIATERESPONSE"
	case CommandGloReadResponse:
		ret = "GLOREADRESPONSE"
	case CommandGloWriteResponse:
		ret = "GLOWRITERESPONSE"
	case CommandGloConfirmedServiceError:
		ret = "GLOCONFIRMEDSERVICEERROR"
	case CommandGloInformationReport:
		ret = "GLOINFORMATIONREPORT"
	case CommandGeneralGloCiphering:
		ret = "GENERALGLOCIPHERING"
	case CommandGeneralDedCiphering:
		ret = "GENERALDEDCIPHERING"
	case CommandGeneralCiphering:
		ret = "GENERALCIPHERING"
	case CommandGeneralSigning:
		ret = "GENERALSIGNING"
	case CommandInformationReport:
		ret = "INFORMATIONREPORT"
	case CommandEventNotification:
		ret = "EVENTNOTIFICATION"
	case CommandDedInitiateRequest:
		ret = "DEDINITIATEREQUEST"
	case CommandDedReadRequest:
		ret = "DEDREADREQUEST"
	case CommandDedWriteRequest:
		ret = "DEDWRITEREQUEST"
	case CommandDedInitiateResponse:
		ret = "DEDINITIATERESPONSE"
	case CommandDedReadResponse:
		ret = "DEDREADRESPONSE"
	case CommandDedWriteResponse:
		ret = "DEDWRITERESPONSE"
	case CommandDedConfirmedServiceError:
		ret = "DEDCONFIRMEDSERVICEERROR"
	case CommandDedUnconfirmedWriteRequest:
		ret = "DEDUNCONFIRMEDWRITEREQUEST"
	case CommandDedInformationReport:
		ret = "DEDINFORMATIONREPORT"
	case CommandDedGetRequest:
		ret = "DEDGETREQUEST"
	case CommandDedGetResponse:
		ret = "DEDGETRESPONSE"
	case CommandDedSetRequest:
		ret = "DEDSETREQUEST"
	case CommandDedSetResponse:
		ret = "DEDSETRESPONSE"
	case CommandDedEventNotification:
		ret = "DEDEVENTNOTIFICATION"
	case CommandDedMethodRequest:
		ret = "DEDMETHODREQUEST"
	case CommandDedMethodResponse:
		ret = "DEDMETHODRESPONSE"
	case CommandGatewayRequest:
		ret = "GATEWAYREQUEST"
	case CommandGatewayResponse:
		ret = "GATEWAYRESPONSE"
	case CommandDiscoverRequest:
		ret = "DISCOVERREQUEST"
	case CommandDiscoverReport:
		ret = "DISCOVERREPORT"
	case CommandRegisterRequest:
		ret = "REGISTERREQUEST"
	case CommandPingRequest:
		ret = "PINGREQUEST"
	case CommandPingResponse:
		ret = "PINGRESPONSE"
	}
	return ret
}
