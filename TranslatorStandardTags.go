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
	"strconv"
	"strings"

	"github.com/Gurux/gxcommon-go"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/constants"
	"github.com/Gurux/gxdlms-go/settings"
)

func standardGetServiceErrors() map[enums.ServiceError]string {
	list := make(map[enums.ServiceError]string)
	list[enums.ServiceErrorApplicationReference] = "application-reference"
	list[enums.ServiceErrorHardwareResource] = "hardware-resource"
	list[enums.ServiceErrorVdeStateError] = "vde-state-error"
	list[enums.ServiceErrorService] = "service"
	list[enums.ServiceErrorDefinition] = "definition"
	list[enums.ServiceErrorAccess] = "access"
	list[enums.ServiceErrorInitiate] = "initiate"
	list[enums.ServiceErrorLoadDataSet] = "load-data-set"
	list[enums.ServiceErrorTask] = "task"
	return list
}

func standardGetApplicationReference() map[enums.ApplicationReference]string {
	list := make(map[enums.ApplicationReference]string)
	list[enums.ApplicationReferenceApplicationContextUnsupported] = "application-context-unsupported"
	list[enums.ApplicationReferenceApplicationReferenceInvalid] = "application-reference-invalid"
	list[enums.ApplicationReferenceApplicationUnreachable] = "application-unreachable"
	list[enums.ApplicationReferenceDecipheringError] = "deciphering-error"
	list[enums.ApplicationReferenceOther] = "other"
	list[enums.ApplicationReferenceProviderCommunicationError] = "provider-communication-error"
	list[enums.ApplicationReferenceTimeElapsed] = "time-elapsed"
	return list
}

func standardGetHardwareResource() map[enums.HardwareResource]string {
	list := make(map[enums.HardwareResource]string)
	list[enums.HardwareResourceMassStorageUnavailable] = "mass-storage-unavailable"
	list[enums.HardwareResourceMemoryUnavailable] = "memory-unavailable"
	list[enums.HardwareResourceOther] = "other"
	list[enums.HardwareResourceOtherResourceUnavailable] = "other-resource-unavailable"
	list[enums.HardwareResourceProcessorResourceUnavailable] = "processor-resource-unavailable"
	return list
}

func standardGetVdeStateError() map[enums.VdeStateError]string {
	list := make(map[enums.VdeStateError]string)
	list[enums.VdeStateErrorLoadingDataSet] = "loading-data-set"
	list[enums.VdeStateErrorNoDlmsContext] = "no-dlms-context"
	list[enums.VdeStateErrorOther] = "other"
	list[enums.VdeStateErrorStatusInoperable] = "status-inoperable"
	list[enums.VdeStateErrorStatusNochange] = "status-nochange"
	return list
}

func standardGetService() map[enums.Service]string {
	list := make(map[enums.Service]string)
	list[enums.ServiceOther] = "other"
	list[enums.ServicePduSize] = "pdu-size"
	list[enums.ServiceUnsupported] = "service-unsupported"
	return list
}

func standardGetDefinition() map[enums.Definition]string {
	list := make(map[enums.Definition]string)
	list[enums.DefinitionObjectAttributeInconsistent] = "object-attribute-inconsistent"
	list[enums.DefinitionObjectClassInconsistent] = "object-class-inconsistent"
	list[enums.DefinitionObjectUndefined] = "object-undefined"
	list[enums.DefinitionOther] = "other"
	return list
}

func standardGetAccess() map[enums.Access]string {
	list := make(map[enums.Access]string)
	list[enums.AccessHardwareFault] = "hardware-fault"
	list[enums.AccessObjectAccessInvalid] = "object-access-violated"
	list[enums.AccessObjectUnavailable] = "object-unavailable"
	list[enums.AccessOther] = "other"
	list[enums.AccessScopeOfAccessViolated] = "scope-of-access-violated"
	return list
}

func standardGetInitiate() map[enums.Initiate]string {
	list := make(map[enums.Initiate]string)
	list[enums.InitiateDlmsVersionTooLow] = "dlms-version-too-low"
	list[enums.InitiateIncompatibleConformance] = "incompatible-conformance"
	list[enums.InitiateOther] = "other"
	list[enums.InitiatePduSizeTooShort] = "pdu-size-too-short"
	list[enums.InitiateRefusedByTheVDEHandler] = "refused-by-the-VDE-Handler"
	return list
}

func standardGetLoadDataSet() map[enums.LoadDataSet]string {
	list := make(map[enums.LoadDataSet]string)
	list[enums.LoadDataSetDatasetNotReady] = "data-set-not-ready"
	list[enums.LoadDataSetDatasetSizeTooLarge] = "dataset-size-too-large"
	list[enums.LoadDataSetInterpretationFailure] = "interpretation-failure"
	list[enums.LoadDataSetNotAwaitedSegment] = "not-awaited-segment"
	list[enums.LoadDataSetNotLoadable] = "not-loadable"
	list[enums.LoadDataSetOther] = "other"
	list[enums.LoadDataSetPrimitiveOutOfSequence] = "primitive-out-of-sequence"
	list[enums.LoadDataSetStorageFailure] = "storage-failure"
	return list
}

func standardGetTask() map[enums.Task]string {
	list := make(map[enums.Task]string)
	list[enums.TaskNoRemoteControl] = "no-remote-control"
	list[enums.TaskOther] = "other"
	list[enums.TaskTiRunning] = "ti-running"
	list[enums.TaskTiStopped] = "ti-stopped"
	list[enums.TaskTiUnusable] = "ti-unusable"
	return list
}

func standardGetApplicationReferenceByName(value string) (int, error) {
	ret := -1
	for k, v := range standardGetApplicationReference() {
		if strings.EqualFold(value, v) {
			ret = int(k)
			break
		}
	}
	if ret == -1 {
		return 0, gxcommon.ErrInvalidArgument
	}
	return ret, nil
}

func standardGetHardwareResourceByName(value string) (int, error) {
	ret := -1
	for k, v := range standardGetHardwareResource() {
		if strings.EqualFold(value, v) {
			ret = int(k)
			break
		}
	}
	if ret == -1 {
		return 0, gxcommon.ErrInvalidArgument
	}
	return ret, nil
}

func standardGetVdeStateErrorByName(value string) (int, error) {
	ret := -1
	for k, v := range standardGetVdeStateError() {
		if strings.EqualFold(value, v) {
			ret = int(k)
			break
		}
	}
	if ret == -1 {
		return 0, gxcommon.ErrInvalidArgument
	}
	return ret, nil
}

func standardGetServiceByName(value string) (int, error) {
	ret := -1
	for k, v := range standardGetService() {
		if strings.EqualFold(value, v) {
			ret = int(k)
			break
		}
	}
	if ret == -1 {
		return 0, gxcommon.ErrInvalidArgument
	}
	return ret, nil
}

func standardGetDefinitionByName(value string) (int, error) {
	ret := -1
	for k, v := range standardGetDefinition() {
		if strings.EqualFold(value, v) {
			ret = int(k)
			break
		}
	}
	if ret == -1 {
		return 0, gxcommon.ErrInvalidArgument
	}
	return ret, nil
}

func standardGetAccessByName(value string) (int, error) {
	ret := -1
	for k, v := range standardGetAccess() {
		if strings.EqualFold(value, v) {
			ret = int(k)
			break
		}
	}
	if ret == -1 {
		return 0, gxcommon.ErrInvalidArgument
	}
	return ret, nil
}

func standardGetInitiateByName(value string) (int, error) {
	ret := -1
	for k, v := range standardGetInitiate() {
		if strings.EqualFold(value, v) {
			ret = int(k)
			break
		}
	}
	if ret == -1 {
		return 0, gxcommon.ErrInvalidArgument
	}
	return ret, nil
}

func standardGetLoadDataSetByName(value string) (int, error) {
	ret := -1
	for k, v := range standardGetLoadDataSet() {
		if strings.EqualFold(value, v) {
			ret = int(k)
			break
		}
	}
	if ret == -1 {
		return 0, gxcommon.ErrInvalidArgument
	}
	return ret, nil
}

func standardGetTaskByName(value string) (int, error) {
	ret := -1
	for k, v := range standardGetTask() {
		if strings.EqualFold(value, v) {
			ret = int(k)
			break
		}
	}
	if ret == -1 {
		return 0, gxcommon.ErrInvalidArgument
	}
	return ret, nil
}

// GetGeneralTags returns the get general tags.
func standardGetGeneralTags(type_ enums.TranslatorOutputType, list map[int]string) {
	addTag(list, enums.CommandSnrm, "Snrm")
	addTag(list, enums.CommandUnacceptableFrame, "UnacceptableFrame")
	addTag(list, enums.CommandDisconnectMode, "DisconnectMode")
	addTag(list, enums.CommandUa, "Ua")
	addTag(list, enums.CommandAarq, "aarq")
	addTag(list, enums.CommandAare, "aare")
	addTag(list, int(constants.TranslatorGeneralTagsApplicationContextName), "application-context-name")
	addTag(list, enums.CommandInitiateResponse, "InitiateResponse")
	addTag(list, enums.CommandInitiateRequest, "initiateRequest")
	addTag(list, int(constants.TranslatorGeneralTagsNegotiatedQualityOfService), "negotiated-quality-of-service")
	addTag(list, int(constants.TranslatorGeneralTagsProposedQualityOfService), "proposed-quality-of-service")
	addTag(list, int(constants.TranslatorGeneralTagsProposedDlmsVersionNumber), "proposed-dlms-version-number")
	addTag(list, int(constants.TranslatorGeneralTagsProposedMaxPduSize), "client-max-receive-pdu-size")
	addTag(list, int(constants.TranslatorGeneralTagsProposedConformance), "proposed-conformance")
	addTag(list, int(constants.TranslatorGeneralTagsVaaName), "VaaName")
	addTag(list, int(constants.TranslatorGeneralTagsNegotiatedConformance), "NegotiatedConformance")
	addTag(list, int(constants.TranslatorGeneralTagsNegotiatedDlmsVersionNumber), "NegotiatedDlmsVersionNumber")
	addTag(list, int(constants.TranslatorGeneralTagsNegotiatedMaxPduSize), "NegotiatedMaxPduSize")
	addTag(list, int(constants.TranslatorGeneralTagsConformanceBit), "ConformanceBit")
	addTag(list, int(constants.TranslatorGeneralTagsSenderACSERequirements), "sender-acse-requirements")
	addTag(list, int(constants.TranslatorGeneralTagsResponderACSERequirement), "responder-acse-requirements")
	addTag(list, int(constants.TranslatorGeneralTagsRespondingMechanismName), "mechanism-name")
	addTag(list, int(constants.TranslatorGeneralTagsCallingMechanismName), "mechanism-name")
	addTag(list, int(constants.TranslatorGeneralTagsCallingAuthentication), "calling-authentication-value")
	addTag(list, int(constants.TranslatorGeneralTagsRespondingAuthentication), "responding-authentication-value")
	addTag(list, enums.CommandReleaseRequest, "rlrq")
	addTag(list, int(enums.CommandReleaseResponse), "rlre")
	addTag(list, int(enums.CommandDisconnectRequest), "Disc")
	addTag(list, int(constants.TranslatorGeneralTagsAssociationResult), "result")
	addTag(list, int(constants.TranslatorGeneralTagsResultSourceDiagnostic), "result-source-diagnostic")
	addTag(list, int(constants.TranslatorGeneralTagsACSEServiceUser), "acse-service-user")
	addTag(list, int(constants.TranslatorGeneralTagsACSEServiceProvider), "acse-service-provider")
	addTag(list, int(constants.TranslatorGeneralTagsCallingAPTitle), "CallingAPTitle")
	addTag(list, int(constants.TranslatorGeneralTagsRespondingAPTitle), "RespondingAPTitle")
	addTag(list, int(constants.TranslatorGeneralTagsCharString), "charstring")
	addTag(list, int(constants.TranslatorGeneralTagsDedicatedKey), "dedicated-key")
	addTag(list, int(internal.TranslatorTagsResponseAllowed), "response-allowed")
	addTag(list, int(constants.TranslatorGeneralTagsUserInformation), "user-information")
	addTag(list, int(enums.CommandConfirmedServiceError), "confirmedServiceError")
	addTag(list, int(enums.CommandInformationReport), "InformationReport")
	addTag(list, int(enums.CommandEventNotification), "event-notification-request")
	addTag(list, int(constants.TranslatorGeneralTagsCallingAeInvocationId), "calling-AE-invocation-id")
	addTag(list, int(constants.TranslatorGeneralTagsCalledAeInvocationId), "called-AE-invocation-id")
	addTag(list, int(constants.TranslatorGeneralTagsRespondingAeInvocationId), "responding-AE-invocation-id")
	addTag(list, int(constants.TranslatorGeneralTagsCallingAeQualifier), "calling-Ae-qualifier")
	addTag(list, int(enums.CommandExceptionResponse), "exception-response")
	addTag(list, int(internal.TranslatorTagsStateError), "state-error")
}

// GetSnTags returns the get SN tags.
func standardGetSnTags(type_ enums.TranslatorOutputType, list map[int]string) {
	list[int(enums.CommandReadRequest)] = "readRequest"
	list[int(enums.CommandWriteRequest)] = "writeRequest"
	list[int(enums.CommandWriteResponse)] = "writeResponse"
	list[int(enums.CommandWriteRequest)<<8|int(constants.SingleReadResponseData)] = "Data"
	list[int(enums.CommandReadRequest)<<8|int(constants.VariableAccessSpecificationVariableName)] = "variable-name"
	list[int(enums.CommandReadRequest)<<8|int(constants.VariableAccessSpecificationParameterisedAccess)] = "parameterized-access"
	list[int(enums.CommandReadRequest)<<8|int(constants.VariableAccessSpecificationBlockNumberAccess)] = "BlockNumberAccess"
	list[int(enums.CommandWriteRequest)<<8|int(constants.VariableAccessSpecificationVariableName)] = "variable-name"
	list[int(enums.CommandReadResponse)] = "readResponse"
	list[int(enums.CommandReadResponse)<<8|int(constants.SingleReadResponseDataBlockResult)] = "DataBlockResult"
	list[int(enums.CommandReadResponse)<<8|int(constants.SingleReadResponseData)] = "data"
	list[int(enums.CommandWriteResponse)<<8|int(constants.SingleReadResponseData)] = "data"
	list[int(enums.CommandReadResponse)<<8|int(constants.SingleReadResponseDataAccessError)] = "data-access-error"
}

// GetLnTags returns the get LN tags.
func standardGetLnTags(type_ enums.TranslatorOutputType, list map[int]string) {
	addTag(list, enums.CommandGetRequest, "get-request")
	list[int(enums.CommandGetRequest)<<8|int(constants.GetCommandTypeNormal)] = "get-request-normal"
	list[int(enums.CommandGetRequest)<<8|int(constants.GetCommandTypeNextDataBlock)] = "get-request-next"
	list[int(enums.CommandGetRequest)<<8|int(constants.GetCommandTypeWithList)] = "get-request-with-list"
	addTag(list, enums.CommandSetRequest, "set-request")
	list[int(enums.CommandSetRequest)<<8|int(constants.SetRequestTypeNormal)] = "set-request-normal"
	list[int(enums.CommandSetRequest)<<8|int(constants.SetRequestTypeFirstDataBlock)] = "set-request-first-data-block"
	list[int(enums.CommandSetRequest)<<8|int(constants.SetRequestTypeWithDataBlock)] = "set-request-with-data-block"
	list[int(enums.CommandSetRequest)<<8|int(constants.SetRequestTypeWithList)] = "set-request-with-list"
	addTag(list, enums.CommandMethodRequest, "action-request")
	list[int(enums.CommandMethodRequest)<<8|int(constants.ActionRequestTypeNormal)] = "action-request-normal"
	list[int(enums.CommandMethodRequest)<<8|int(constants.ActionRequestTypeNextBlock)] = "action-request-next-pblock"
	list[int(enums.CommandMethodRequest)<<8|int(constants.ActionRequestTypeWithList)] = "action-request-with-list"
	list[int(enums.CommandMethodRequest)<<8|int(constants.ActionRequestTypeWithFirstBlock)] = "action-request-with-first-block"
	list[int(enums.CommandMethodRequest)<<8|int(constants.ActionRequestTypeWithListAndFirstBlock)] = "action-request-with-list-and-first-block"
	list[int(enums.CommandMethodRequest)<<8|int(constants.ActionRequestTypeWithBlock)] = "action-request-with-list-and-block"
	addTag(list, enums.CommandMethodResponse, "action-response")
	list[int(enums.CommandMethodResponse)<<8|int(constants.ActionResponseTypeNormal)] = "action-response-normal"
	list[int(enums.CommandMethodResponse)<<8|int(constants.ActionResponseTypeWithBlock)] = "action-response-with-pblock"
	list[int(enums.CommandMethodResponse)<<8|int(constants.ActionResponseTypeWithList)] = "action-response-with-list"
	list[int(enums.CommandMethodResponse)<<8|int(constants.ActionResponseTypeNextBlock)] = "action-response-next-pblock"
	list[int(internal.TranslatorTagsSingleResponse)] = "single-response"
	list[int(enums.CommandDataNotification)] = "data-notification"
	list[int(enums.CommandGetResponse)] = "get-response"
	list[int(enums.CommandGetResponse)<<8|int(constants.GetCommandTypeNormal)] = "get-response-normal"
	list[int(enums.CommandGetResponse)<<8|int(constants.GetCommandTypeNextDataBlock)] = "get-response-with-data-block"
	list[int(enums.CommandGetResponse)<<8|int(constants.GetCommandTypeWithList)] = "get-response-with-list"
	addTag(list, enums.CommandSetResponse, "set-response")
	list[int(enums.CommandSetResponse)<<8|int(constants.SetResponseTypeNormal)] = "set-response-normal"
	list[int(enums.CommandSetResponse)<<8|int(constants.SetResponseTypeDataBlock)] = "set-response-datablock"
	list[int(enums.CommandSetResponse)<<8|int(constants.SetResponseTypeLastDataBlock)] = "set-response-with-last-data-block"
	list[int(enums.CommandSetResponse)<<8|int(constants.SetResponseTypeWithList)] = "set-response-with-list"
	addTag(list, enums.CommandAccessRequest, "access-request")
	list[int(enums.CommandAccessRequest)<<8|int(enums.AccessServiceCommandTypeGet)] = "access-request-get"
	list[int(enums.CommandAccessRequest)<<8|int(enums.AccessServiceCommandTypeSet)] = "access-request-set"
	list[int(enums.CommandAccessRequest)<<8|int(enums.AccessServiceCommandTypeAction)] = "access-request-action"
	addTag(list, enums.CommandAccessResponse, "access-response")
	list[int(enums.CommandAccessResponse)<<8|int(enums.AccessServiceCommandTypeGet)] = "access-response-get"
	list[int(enums.CommandAccessResponse)<<8|int(enums.AccessServiceCommandTypeSet)] = "access-response-set"
	list[int(enums.CommandAccessResponse)<<8|int(enums.AccessServiceCommandTypeAction)] = "access-response-action"
	list[int(internal.TranslatorTagsAccessRequestBody)] = "access-request-body"
	list[int(internal.TranslatorTagsListOfAccessRequestSpecification)] = "access-request-specification"
	list[int(internal.TranslatorTagsAccessRequestSpecification)] = "Access-Request-Specification"
	list[int(internal.TranslatorTagsAccessRequestListOfData)] = "access-request-list-of-data"
	list[int(internal.TranslatorTagsAccessResponseBody)] = "access-response-body"
	list[int(internal.TranslatorTagsListOfAccessResponseSpecification)] = "access-response-specification"
	list[int(internal.TranslatorTagsAccessResponseSpecification)] = "Access-Response-Specification"
	list[int(internal.TranslatorTagsAccessResponseListOfData)] = "access-response-list-of-data"
	list[int(internal.TranslatorTagsService)] = "service"
	list[int(internal.TranslatorTagsServiceError)] = "service-error"
	addTag(list, enums.CommandGeneralBlockTransfer, "general-block-transfer")
	addTag(list, enums.CommandGatewayRequest, "gateway-request")
	addTag(list, enums.CommandGatewayResponse, "gateway-response")
}

// GetPlcTags returns the get PLC tags.
func standardGetPlcTags(list map[int]string) {
	addTag(list, enums.CommandDiscoverRequest, "discover-request")
	addTag(list, enums.CommandDiscoverReport, "discover-report")
	addTag(list, enums.CommandRegisterRequest, "register-request")
	addTag(list, enums.CommandPingRequest, "ping-request")
	addTag(list, enums.CommandPingResponse, "ping-response")
}

// GetGloTags returns the get glo tags.
func standardGetGloTags(type_ enums.TranslatorOutputType, list map[int]string) {
	addTag(list, enums.CommandGloInitiateRequest, "glo-initiate-request")
	addTag(list, enums.CommandGloInitiateResponse, "glo-initiate-response")
	addTag(list, enums.CommandGloGetRequest, "glo-get-request")
	addTag(list, enums.CommandGloGetResponse, "glo-get-response")
	addTag(list, enums.CommandGloSetRequest, "glo-set-request")
	addTag(list, enums.CommandGloSetResponse, "glo-set-response")
	addTag(list, enums.CommandGloMethodRequest, "glo-action-request")
	addTag(list, enums.CommandGloMethodResponse, "glo-action-response")
	addTag(list, enums.CommandGloReadRequest, "glo-read-request")
	addTag(list, enums.CommandGloReadResponse, "glo-read-response")
	addTag(list, enums.CommandGloWriteRequest, "glo-write-request")
	addTag(list, enums.CommandGloWriteResponse, "glo-write-response")
	addTag(list, enums.CommandGeneralGloCiphering, "general-glo-ciphering")
	addTag(list, enums.CommandGloConfirmedServiceError, "glo-confirmed-service-error")
}

// GetDedTags returns the get ded tags.
func standardGetDedTags(type_ enums.TranslatorOutputType, list map[int]string) {
	addTag(list, enums.CommandDedInitiateRequest, "ded-initiate-request")
	addTag(list, enums.CommandDedInitiateResponse, "ded-initiate-response")
	addTag(list, enums.CommandDedGetRequest, "ded-get-request")
	addTag(list, enums.CommandDedGetResponse, "ded-get-response")
	addTag(list, enums.CommandDedSetRequest, "ded-set-request")
	addTag(list, enums.CommandDedSetResponse, "ded-set-response")
	addTag(list, enums.CommandDedMethodRequest, "ded-action-request")
	addTag(list, enums.CommandDedMethodResponse, "ded-action-response")
	addTag(list, enums.CommandGeneralDedCiphering, "general-ded-ciphering")
	addTag(list, enums.CommandDedConfirmedServiceError, "ded-confirmed-service-error")
	addTag(list, enums.CommandGeneralCiphering, "general-ciphering")
	addTag(list, enums.CommandGeneralSigning, "general-signing")
}

// getTranslatorTags returns the get translator tags.
func standardGetTranslatorTags(type_ enums.TranslatorOutputType, list map[int]string) {
	addTag(list, int(internal.TranslatorTagsWrapper), "Wrapper")
	addTag(list, int(internal.TranslatorTagsHdlc), "Hdlc")
	addTag(list, int(internal.TranslatorTagsPduDlms), "xDLMS-APDU")
	addTag(list, int(internal.TranslatorTagsPduCse), "aCSE-APDU")
	addTag(list, int(internal.TranslatorTagsTargetAddress), "TargetAddress")
	addTag(list, int(internal.TranslatorTagsSourceAddress), "SourceAddress")
	addTag(list, int(internal.TranslatorTagsListOfVariableAccessSpecification), "variable-access-specification")
	addTag(list, int(internal.TranslatorTagsListOfData), "list-of-data")
	addTag(list, int(internal.TranslatorTagsSuccess), "Success")
	addTag(list, int(internal.TranslatorTagsDataAccessError), "data-access-result")
	addTag(list, int(internal.TranslatorTagsAttributeDescriptor), "cosem-attribute-descriptor")
	addTag(list, int(internal.TranslatorTagsClassId), "class-id")
	addTag(list, int(internal.TranslatorTagsInstanceId), "instance-id")
	addTag(list, int(internal.TranslatorTagsAttributeId), "attribute-id")
	addTag(list, int(internal.TranslatorTagsMethodInvocationParameters), "method-invocation-parameters")
	addTag(list, int(internal.TranslatorTagsSelector), "selector")
	addTag(list, int(internal.TranslatorTagsParameter), "parameter")
	addTag(list, int(internal.TranslatorTagsLastBlock), "last-block")
	addTag(list, int(internal.TranslatorTagsBlockNumber), "block-number")
	addTag(list, int(internal.TranslatorTagsRawData), "raw-data")
	addTag(list, int(internal.TranslatorTagsMethodDescriptor), "cosem-method-descriptor")
	addTag(list, int(internal.TranslatorTagsMethodId), "method-id")
	addTag(list, int(internal.TranslatorTagsResult), "result")
	addTag(list, int(internal.TranslatorTagsPblock), "pblock")
	addTag(list, int(internal.TranslatorTagsContent), "content")
	addTag(list, int(internal.TranslatorTagsSignature), "signature")
	addTag(list, int(internal.TranslatorTagsReturnParameters), "return-parameters")
	addTag(list, int(internal.TranslatorTagsAccessSelection), "access-selection")
	addTag(list, int(internal.TranslatorTagsValue), "value")
	addTag(list, int(internal.TranslatorTagsAccessSelector), "access-selector")
	addTag(list, int(internal.TranslatorTagsAccessParameters), "access-parameters")
	addTag(list, int(internal.TranslatorTagsAttributeDescriptorList), "attribute-descriptor-list")
	addTag(list, int(internal.TranslatorTagsAttributeDescriptorWithSelection), "Cosem-Attribute-Descriptor-With-Selection")
	addTag(list, int(internal.TranslatorTagsReadDataBlockAccess), "ReadDataBlockAccess")
	addTag(list, int(internal.TranslatorTagsWriteDataBlockAccess), "WriteDataBlockAccess")
	addTag(list, int(internal.TranslatorTagsData), "data")
	addTag(list, int(internal.TranslatorTagsInvokeId), "invoke-id-and-priority")
	addTag(list, int(internal.TranslatorTagsLongInvokeId), "long-invoke-id-and-priority")
	addTag(list, int(internal.TranslatorTagsDateTime), "date-time")
	addTag(list, int(internal.TranslatorTagsCurrentTime), "current-time")
	addTag(list, int(internal.TranslatorTagsTime), "time")
	addTag(list, int(internal.TranslatorTagsReason), "Reason")
	addTag(list, int(internal.TranslatorTagsVariableAccessSpecification), "Variable-Access-Specification")
	addTag(list, int(internal.TranslatorTagsChoice), "CHOICE")
	addTag(list, int(internal.TranslatorTagsNotificationBody), "notification-body")
	addTag(list, int(internal.TranslatorTagsDataValue), "data-value")
	addTag(list, int(internal.TranslatorTagsInitiateError), "initiateError")
	addTag(list, int(internal.TranslatorTagsCipheredService), "ciphered-content")
	addTag(list, int(internal.TranslatorTagsSystemTitle), "system-title")
	addTag(list, int(internal.TranslatorTagsDataBlock), "DataBlock")
	addTag(list, int(internal.TranslatorTagsTransactionId), "TransactionId")
	addTag(list, int(internal.TranslatorTagsOriginatorSystemTitle), "OriginatorSystemTitle")
	addTag(list, int(internal.TranslatorTagsRecipientSystemTitle), "RecipientSystemTitle")
	addTag(list, int(internal.TranslatorTagsOtherInformation), "OtherInformation")
	addTag(list, int(internal.TranslatorTagsKeyInfo), "KeyInfo")
	addTag(list, int(internal.TranslatorTagsCipheredContent), "CipheredContent")
	addTag(list, int(internal.TranslatorTagsAgreedKey), "AgreedKey")
	addTag(list, int(internal.TranslatorTagsKeyParameters), "KeyParameters")
	addTag(list, int(internal.TranslatorTagsKeyCipheredData), "KeyCipheredData")
	addTag(list, int(internal.TranslatorTagsAttributeValue), "attribute-value")
	addTag(list, int(internal.TranslatorTagsMaxInfoRX), "MaxInfoRX")
	addTag(list, int(internal.TranslatorTagsMaxInfoTX), "MaxInfoTX")
	addTag(list, int(internal.TranslatorTagsWindowSizeRX), "WindowSizeRX")
	addTag(list, int(internal.TranslatorTagsWindowSizeTX), "WindowSizeTX")
	addTag(list, int(internal.TranslatorTagsValueList), "value-list")
	addTag(list, int(internal.TranslatorTagsDataAccessResult), "data-access-result")
	addTag(list, int(internal.TranslatorTagsBlockControl), "block-control")
	addTag(list, int(internal.TranslatorTagsBlockNumberAck), "block-number-ack")
	addTag(list, int(internal.TranslatorTagsBlockData), "block-data")
	addTag(list, int(internal.TranslatorTagsContentsDescription), "contents-description")
	addTag(list, int(internal.TranslatorTagsArrayContents), "array-contents")
	addTag(list, int(internal.TranslatorTagsNetworkId), "network-id")
	addTag(list, int(internal.TranslatorTagsPhysicalDeviceAddress), "physical-device-address")
	addTag(list, int(internal.TranslatorTagsProtocolVersion), "protocol-version")
	addTag(list, int(internal.TranslatorTagsCalledAPTitle), "called-ap-title")
	addTag(list, int(internal.TranslatorTagsCalledAPInvocationId), "called-ap-invocation-id")
	addTag(list, int(internal.TranslatorTagsCalledAEInvocationId), "called-ae-invocation-id")
	addTag(list, int(internal.TranslatorTagsCallingApInvocationId), "calling-ap-invocation-id")
	addTag(list, int(internal.TranslatorTagsCalledAEQualifier), "called-ae-qualifier")
}

func standardGetServiceErrorValue(error enums.ServiceError, value uint8) string {
	switch error {
	case enums.ServiceErrorApplicationReference:
		return standardGetApplicationReference()[enums.ApplicationReference(value)]
	case enums.ServiceErrorHardwareResource:
		return standardGetHardwareResource()[enums.HardwareResource(value)]
	case enums.ServiceErrorVdeStateError:
		return standardGetVdeStateError()[enums.VdeStateError(value)]
	case enums.ServiceErrorService:
		return standardGetService()[enums.Service(value)]
	case enums.ServiceErrorDefinition:
		return standardGetDefinition()[enums.Definition(value)]
	case enums.ServiceErrorAccess:
		return standardGetAccess()[enums.Access(value)]
	case enums.ServiceErrorInitiate:
		return standardGetInitiate()[enums.Initiate(value)]
	case enums.ServiceErrorLoadDataSet:
		return standardGetLoadDataSet()[enums.LoadDataSet(value)]
	case enums.ServiceErrorTask:
		return standardGetTask()[enums.Task(value)]
	case enums.ServiceErrorOtherError:
		return strconv.Itoa(int(value))
	default:
	}
	return ""
}

// standardServiceErrorToString returns the parameters:
//
//	error: Service error enumeration value.
//
// Returns:
//
//	Service error standard XML tag.
func standardServiceErrorToString(error enums.ServiceError) string {
	return standardGetServiceErrors()[error]
}

// GetServiceError returns the parameters:
//
//	value: Service error standard XML tag.
//
// Returns:
//
//	Service error enumeration value.
func standardGetServiceErrorByName(value string) (enums.ServiceError, error) {
	for k, v := range standardGetServiceErrors() {
		if strings.EqualFold(value, v) {
			return k, nil
		}
	}
	return 0, gxcommon.ErrInvalidArgument
}

func standardGetError(serviceError enums.ServiceError, value string) (uint8, error) {
	var err error
	ret := 0
	switch serviceError {
	case enums.ServiceErrorApplicationReference:
		ret, err = standardGetApplicationReferenceByName(value)
	case enums.ServiceErrorHardwareResource:
		ret, err = standardGetHardwareResourceByName(value)
	case enums.ServiceErrorVdeStateError:
		ret, err = standardGetVdeStateErrorByName(value)
	case enums.ServiceErrorService:
		ret, err = standardGetServiceByName(value)
	case enums.ServiceErrorDefinition:
		ret, err = standardGetDefinitionByName(value)
	case enums.ServiceErrorAccess:
		ret, err = standardGetAccessByName(value)
	case enums.ServiceErrorInitiate:
		ret, err = standardGetInitiateByName(value)
	case enums.ServiceErrorLoadDataSet:
		ret, err = standardGetLoadDataSetByName(value)
	case enums.ServiceErrorTask:
		ret, err = standardGetTaskByName(value)
	case enums.ServiceErrorOtherError:
		ret, err = strconv.Atoi(value)
		if err != nil {
			return 0, err
		}
	default:
	}
	return uint8(ret), nil
}

func standardvalueOfConformance(value string) (enums.Conformance, error) {
	var ret enums.Conformance
	if strings.EqualFold("access", value) {
		ret = enums.ConformanceAccess
	} else if strings.EqualFold("action", value) {
		ret = enums.ConformanceAction
	} else if strings.EqualFold("attribute0-supported-with-get", value) {
		ret = enums.ConformanceAttribute0SupportedWithGet
	} else if strings.EqualFold("attribute0-supported-with-set", value) {
		ret = enums.ConformanceAttribute0SupportedWithSet
	} else if strings.EqualFold("block-transfer-with-action", value) {
		ret = enums.ConformanceBlockTransferWithAction
	} else if strings.EqualFold("block-transfer-with-get-or-read", value) {
		ret = enums.ConformanceBlockTransferWithGetOrRead
	} else if strings.EqualFold("block-transfer-with-set-or-write", value) {
		ret = enums.ConformanceBlockTransferWithSetOrWrite
	} else if strings.EqualFold("data-notification", value) {
		ret = enums.ConformanceDataNotification
	} else if strings.EqualFold("event-notification", value) {
		ret = enums.ConformanceEventNotification
	} else if strings.EqualFold("general-block-transfer", value) {
		ret = enums.ConformanceGeneralBlockTransfer
	} else if strings.EqualFold("general-protection", value) {
		ret = enums.ConformanceGeneralProtection
	} else if strings.EqualFold("delta-value-encoding", value) {
		ret = enums.ConformanceDeltaValueEncoding
	} else if strings.EqualFold("get", value) {
		ret = enums.ConformanceGet
	} else if strings.EqualFold("information-report", value) {
		ret = enums.ConformanceInformationReport
	} else if strings.EqualFold("multiple-references", value) {
		ret = enums.ConformanceMultipleReferences
	} else if strings.EqualFold("parameterized-access", value) {
		ret = enums.ConformanceParameterizedAccess
	} else if strings.EqualFold("priority-mgmt-supported", value) {
		ret = enums.ConformancePriorityMgmtSupported
	} else if strings.EqualFold("read", value) {
		ret = enums.ConformanceRead
	} else if strings.EqualFold("reserved-seven", value) {
		ret = enums.ConformanceReservedSeven
	} else if strings.EqualFold("reserved-zero", value) {
		ret = enums.ConformanceReservedZero
	} else if strings.EqualFold("selective-access", value) {
		ret = enums.ConformanceSelectiveAccess
	} else if strings.EqualFold("set", value) {
		ret = enums.ConformanceSet
	} else if strings.EqualFold("unconfirmed-write", value) {
		ret = enums.ConformanceUnconfirmedWrite
	} else if strings.EqualFold("write", value) {
		ret = enums.ConformanceWrite
	} else {
		return 0, gxcommon.ErrInvalidArgument
	}
	return ret, nil
}

func standardGetDataTypeTags(list map[int]string) {
	list[settings.DataTypeOffset+int(enums.DataTypeNone)] = "nil-data"
	list[settings.DataTypeOffset+int(enums.DataTypeArray)] = "array"
	list[settings.DataTypeOffset+int(enums.DataTypeBcd)] = "bcd"
	list[settings.DataTypeOffset+int(enums.DataTypeBitString)] = "bit-string"
	list[settings.DataTypeOffset+int(enums.DataTypeBoolean)] = "boolean"
	list[settings.DataTypeOffset+int(enums.DataTypeCompactArray)] = "compact-array"
	list[settings.DataTypeOffset+int(enums.DataTypeDate)] = "date"
	list[settings.DataTypeOffset+int(enums.DataTypeDateTime)] = "date-time"
	list[settings.DataTypeOffset+int(enums.DataTypeEnum)] = "enum"
	list[settings.DataTypeOffset+int(enums.DataTypeFloat32)] = "float32"
	list[settings.DataTypeOffset+int(enums.DataTypeFloat64)] = "float64,"
	list[settings.DataTypeOffset+int(enums.DataTypeInt16)] = "long"
	list[settings.DataTypeOffset+int(enums.DataTypeInt32)] = "double-long"
	list[settings.DataTypeOffset+int(enums.DataTypeInt64)] = "long64"
	list[settings.DataTypeOffset+int(enums.DataTypeInt8)] = "integer"
	list[settings.DataTypeOffset+int(enums.DataTypeOctetString)] = "octet-string"
	list[settings.DataTypeOffset+int(enums.DataTypeString)] = "visible-string"
	list[settings.DataTypeOffset+int(enums.DataTypeStringUTF8)] = "utf8-string"
	list[settings.DataTypeOffset+int(enums.DataTypeStructure)] = "structure"
	list[settings.DataTypeOffset+int(enums.DataTypeTime)] = "time"
	list[settings.DataTypeOffset+int(enums.DataTypeUint16)] = "long-unsigned"
	list[settings.DataTypeOffset+int(enums.DataTypeUint32)] = "double-long-unsigned"
	list[settings.DataTypeOffset+int(enums.DataTypeUint64)] = "long64-unsigned"
	list[settings.DataTypeOffset+int(enums.DataTypeUint8)] = "unsigned"
	list[settings.DataTypeOffset+int(enums.DataTypeDeltaInt8)] = "integer-delta"
	list[settings.DataTypeOffset+int(enums.DataTypeDeltaInt16)] = "long-integer-delta"
	list[settings.DataTypeOffset+int(enums.DataTypeDeltaInt32)] = "double-long-integer-delta"
	list[settings.DataTypeOffset+int(enums.DataTypeDeltaUint8)] = "unsigned-delta"
	list[settings.DataTypeOffset+int(enums.DataTypeDeltaUint16)] = "long-unsigned-delta"
	list[settings.DataTypeOffset+int(enums.DataTypeDeltaUint32)] = "double-long-unsigned-delta"
}

func standardErrorCodeToString(value enums.ErrorCode) (string, error) {
	var str string
	switch value {
	case enums.ErrorCodeAccessViolated:
		str = "scope-of-access-violated"
	case enums.ErrorCodeDataBlockNumberInvalid:
		str = "data-block-number-invalid"
	case enums.ErrorCodeDataBlockUnavailable:
		str = "data-block-unavailable"
	case enums.ErrorCodeHardwareFault:
		str = "hardware-fault"
	case enums.ErrorCodeInconsistentClass:
		str = "object-class-inconsistent"
	case enums.ErrorCodeLongGetOrReadAborted:
		str = "long-Get-aborted"
	case enums.ErrorCodeLongSetOrWriteAborted:
		str = "long-set-aborted"
	case enums.ErrorCodeNoLongGetOrReadInProgress:
		str = "no-long-Get-in-progress"
	case enums.ErrorCodeNoLongSetOrWriteInProgress:
		str = "no-long-set-in-progress"
	case enums.ErrorCodeOk:
		str = "success"
	case enums.ErrorCodeOtherReason:
		str = "other-reason"
	case enums.ErrorCodeReadWriteDenied:
		str = "read-write-denied"
	case enums.ErrorCodeTemporaryFailure:
		str = "temporary-failure"
	case enums.ErrorCodeUnavailableObject:
		str = "object-unavailable"
	case enums.ErrorCodeUndefinedObject:
		str = "object-undefined"
	case enums.ErrorCodeUnmatchedType:
		str = "type-unmatched"
	default:
		return "", gxcommon.ErrInvalidArgument
	}
	return str, nil
}

func standardvalueOfErrorCode(value string) (enums.ErrorCode, error) {
	var v enums.ErrorCode
	if strings.EqualFold("scope-of-access-violated", value) {
		v = enums.ErrorCodeAccessViolated
	} else if strings.EqualFold("data-block-number-invalid", value) {
		v = enums.ErrorCodeDataBlockNumberInvalid
	} else if strings.EqualFold("data-block-unavailable", value) {
		v = enums.ErrorCodeDataBlockUnavailable
	} else if strings.EqualFold("hardware-fault", value) {
		v = enums.ErrorCodeHardwareFault
	} else if strings.EqualFold("object-class-inconsistent", value) {
		v = enums.ErrorCodeInconsistentClass
	} else if strings.EqualFold("long-Get-aborted", value) {
		v = enums.ErrorCodeLongGetOrReadAborted
	} else if strings.EqualFold("long-set-aborted", value) {
		v = enums.ErrorCodeLongSetOrWriteAborted
	} else if strings.EqualFold("no-long-Get-in-progress", value) {
		v = enums.ErrorCodeNoLongGetOrReadInProgress
	} else if strings.EqualFold("no-long-set-in-progress", value) {
		v = enums.ErrorCodeNoLongSetOrWriteInProgress
	} else if strings.EqualFold("success", value) {
		v = enums.ErrorCodeOk
	} else if strings.EqualFold("other-reason", value) {
		v = enums.ErrorCodeOtherReason
	} else if strings.EqualFold("read-write-denied", value) {
		v = enums.ErrorCodeReadWriteDenied
	} else if strings.EqualFold("temporary-failure", value) {
		v = enums.ErrorCodeTemporaryFailure
	} else if strings.EqualFold("object-unavailable", value) {
		v = enums.ErrorCodeUnavailableObject
	} else if strings.EqualFold("object-undefined", value) {
		v = enums.ErrorCodeUndefinedObject
	} else if strings.EqualFold("type-unmatched", value) {
		v = enums.ErrorCodeUnmatchedType
	} else {
		return 0, errors.New("Error code: ")
	}
	return v, nil
}

func standardReleaseResponseReasonToString(value constants.ReleaseResponseReason) (string, error) {
	var str string
	switch value {
	case constants.ReleaseResponseReasonNormal:
		str = "Normal"
	case constants.ReleaseResponseReasonNotFinished:
		str = "not-finished"
	case constants.ReleaseResponseReasonUserDefined:
		str = "user-defined"
	default:
		return "", gxcommon.ErrInvalidArgument
	}
	return str, nil
}

func standardvalueOfReleaseResponseReason(value string) (constants.ReleaseResponseReason, error) {
	var ret constants.ReleaseResponseReason
	if strings.EqualFold(value, "normal") {
		ret = constants.ReleaseResponseReasonNormal
	} else if strings.EqualFold(value, "not-finished") {
		ret = constants.ReleaseResponseReasonNotFinished
	} else if strings.EqualFold(value, "user-defined") {
		ret = constants.ReleaseResponseReasonUserDefined
	} else {
		return 0, gxcommon.ErrInvalidArgument
	}
	return ret, nil
}

func standardReleaseRequestReasonToString(value constants.ReleaseRequestReason) (string, error) {
	var str string
	switch value {
	case constants.ReleaseRequestReasonNormal:
		str = "normal"
	case constants.ReleaseRequestReasonUrgent:
		str = "urgent"
	case constants.ReleaseRequestReasonUserDefined:
		str = "user-defined"
	default:
		return "", gxcommon.ErrInvalidArgument
	}
	return str, nil
}

func standardvalueOfReleaseRequestReason(value string) (constants.ReleaseRequestReason, error) {
	var ret constants.ReleaseRequestReason
	if strings.EqualFold(value, "normal") {
		ret = constants.ReleaseRequestReasonNormal
	} else if strings.EqualFold(value, "urgent") {
		ret = constants.ReleaseRequestReasonUrgent
	} else if strings.EqualFold(value, "user-defined") {
		ret = constants.ReleaseRequestReasonUserDefined
	} else {
		return 0, gxcommon.ErrInvalidArgument
	}
	return ret, nil
}

// StateErrorToString returns the gets state error description.
//
// Parameters:
//
//	error: State error enumerator value.
//
// Returns:
//
//	State error as an string.
func standardStateErrorToString(error enums.ExceptionStateError) (string, error) {
	switch error {
	case enums.ExceptionStateErrorServiceNotAllowed:
		return "service-not-allowed", nil
	case enums.ExceptionStateErrorServiceUnknown:
		return "service-unknown", nil
	default:
		return "", gxcommon.ErrInvalidArgument
	}
}

// ExceptionServiceErrorToString returns the gets service error description.
//
// Parameters:
//
//	error: Service error enumerator value.
//
// Returns:
//
//	Service error as an string.
func standardExceptionServiceErrorToString(error enums.ExceptionServiceError) (string, error) {
	switch error {
	case enums.ExceptionServiceErrorOperationNotPossible:
		return "operation-not-possible", nil
	case enums.ExceptionServiceErrorServiceNotSupported:
		return "service-not-supported", nil
	case enums.ExceptionServiceErrorOtherReason:
		return "other-reason", nil
	case enums.ExceptionServiceErrorPduTooLong:
		return "pdu-too-long", nil
	case enums.ExceptionServiceErrorDecipheringError:
		return "deciphering-error", nil
	case enums.ExceptionServiceErrorInvocationCounterError:
		return "invocation-counter-error", nil
	default:
		return "", gxcommon.ErrInvalidArgument
	}
}

// ValueofStateError returns the parameters:
//
//	value: State error string value.
//
// Returns:
//
//	State error enum value.
func standardvalueofStateError(value string) (enums.ExceptionStateError, error) {
	if "service-not-allowed" == value {
		return enums.ExceptionStateErrorServiceNotAllowed, nil
	}
	if "service-unknown" == value {
		return enums.ExceptionStateErrorServiceUnknown, nil
	}
	return 0, gxcommon.ErrInvalidArgument
}

// ValueOfExceptionServiceError returns the parameters:
//
//	value: Service error string value.
//
// Returns:
//
//	Service error enum value.
func standardvalueOfExceptionServiceError(value string) (enums.ExceptionServiceError, error) {
	if "operation-not-possible" == value {
		return enums.ExceptionServiceErrorOperationNotPossible, nil
	}
	if "service-not-supported" == value {
		return enums.ExceptionServiceErrorServiceNotSupported, nil
	}
	if "other-reason" == value {
		return enums.ExceptionServiceErrorOtherReason, nil
	}
	if "pdu-too-long" == value {
		return enums.ExceptionServiceErrorPduTooLong, nil
	}
	if "deciphering-error" == value {
		return enums.ExceptionServiceErrorDecipheringError, nil
	}
	if "invocation-counter-error" == value {
		return enums.ExceptionServiceErrorInvocationCounterError, nil
	}
	return 0, gxcommon.ErrInvalidArgument
}
