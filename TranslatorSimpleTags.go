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
	"fmt"
	"strconv"
	"strings"

	"github.com/Gurux/gxcommon-go"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/constants"
	"github.com/Gurux/gxdlms-go/settings"
)

func addTag(list map[int]string, value int, text string) {
	list[value] = text
}

func simpleGetServiceErrors() map[enums.ServiceError]string {
	list := make(map[enums.ServiceError]string)
	list[enums.ServiceErrorApplicationReference] = "ApplicationReference"
	list[enums.ServiceErrorHardwareResource] = "HardwareResource"
	list[enums.ServiceErrorVdeStateError] = "VdeStateError"
	list[enums.ServiceErrorService] = "Service"
	list[enums.ServiceErrorDefinition] = "Definition"
	list[enums.ServiceErrorAccess] = "Access"
	list[enums.ServiceErrorInitiate] = "Initiate"
	list[enums.ServiceErrorLoadDataSet] = "LoadDataSet"
	list[enums.ServiceErrorTask] = "Task"
	return list
}

func simpleGetApplicationReference() map[enums.ApplicationReference]string {
	list := make(map[enums.ApplicationReference]string)
	list[enums.ApplicationReferenceApplicationContextUnsupported] = "ApplicationContextUnsupported"
	list[enums.ApplicationReferenceApplicationReferenceInvalid] = "ApplicationReferenceInvalid"
	list[enums.ApplicationReferenceApplicationUnreachable] = "ApplicationUnreachable"
	list[enums.ApplicationReferenceDecipheringError] = "DecipheringError"
	list[enums.ApplicationReferenceOther] = "Other"
	list[enums.ApplicationReferenceProviderCommunicationError] = "ProviderCommunicationError"
	list[enums.ApplicationReferenceTimeElapsed] = "TimeElapsed"
	return list
}

func simpleGetHardwareResource() map[enums.HardwareResource]string {
	list := make(map[enums.HardwareResource]string)
	list[enums.HardwareResourceMassStorageUnavailable] = "MassStorageUnavailable"
	list[enums.HardwareResourceMemoryUnavailable] = "MemoryUnavailable"
	list[enums.HardwareResourceOther] = "Other"
	list[enums.HardwareResourceOtherResourceUnavailable] = "OtherResourceUnavailable"
	list[enums.HardwareResourceProcessorResourceUnavailable] = "ProcessorResourceUnavailable"
	return list
}

func simpleGetVdeStateError() map[enums.VdeStateError]string {
	list := make(map[enums.VdeStateError]string)
	list[enums.VdeStateErrorLoadingDataSet] = "LoadingDataSet"
	list[enums.VdeStateErrorNoDlmsContext] = "NoDlmsContext"
	list[enums.VdeStateErrorOther] = "Other"
	list[enums.VdeStateErrorStatusInoperable] = "StatusInoperable"
	list[enums.VdeStateErrorStatusNochange] = "StatusNochange"
	return list
}

func simpleGetService() map[enums.Service]string {
	list := make(map[enums.Service]string)
	list[enums.ServiceOther] = "Other"
	list[enums.ServicePduSize] = "PduSize"
	list[enums.ServiceUnsupported] = "ServiceUnsupported"
	return list
}

func simpleGetDefinition() map[enums.Definition]string {
	list := make(map[enums.Definition]string)
	list[enums.DefinitionObjectAttributeInconsistent] = "ObjectAttributeInconsistent"
	list[enums.DefinitionObjectClassInconsistent] = "ObjectClassInconsistent"
	list[enums.DefinitionObjectUndefined] = "ObjectUndefined"
	list[enums.DefinitionOther] = "Other"
	return list
}

func simpleGetAccess() map[enums.Access]string {
	list := make(map[enums.Access]string)
	list[enums.AccessHardwareFault] = "HardwareFault"
	list[enums.AccessObjectAccessInvalid] = "ObjectAccessInvalid"
	list[enums.AccessObjectUnavailable] = "ObjectUnavailable"
	list[enums.AccessOther] = "Other"
	list[enums.AccessScopeOfAccessViolated] = "ScopeOfAccessViolated"
	return list
}

func simpleGetInitiate() map[enums.Initiate]string {
	list := make(map[enums.Initiate]string)
	list[enums.InitiateDlmsVersionTooLow] = "DlmsVersionTooLow"
	list[enums.InitiateIncompatibleConformance] = "IncompatibleConformance"
	list[enums.InitiateOther] = "Other"
	list[enums.InitiatePduSizeTooShort] = "PduSizeTooShort"
	list[enums.InitiateRefusedByTheVDEHandler] = "RefusedByTheVDEHandler"
	return list
}

func simpleGetLoadDataSet() map[enums.LoadDataSet]string {
	list := make(map[enums.LoadDataSet]string)
	list[enums.LoadDataSetDatasetNotReady] = "DataSetNotReady"
	list[enums.LoadDataSetDatasetSizeTooLarge] = "DatasetSizeTooLarge"
	list[enums.LoadDataSetInterpretationFailure] = "InterpretationFailure"
	list[enums.LoadDataSetNotAwaitedSegment] = "NotAwaitedSegment"
	list[enums.LoadDataSetNotLoadable] = "NotLoadable"
	list[enums.LoadDataSetOther] = "Other"
	list[enums.LoadDataSetPrimitiveOutOfSequence] = "PrimitiveOutOfSequence"
	list[enums.LoadDataSetStorageFailure] = "StorageFailure"
	return list
}

func simpleGetTask() map[enums.Task]string {
	list := make(map[enums.Task]string)
	list[enums.TaskNoRemoteControl] = "NoRemoteControl"
	list[enums.TaskOther] = "Other"
	list[enums.TaskTiRunning] = "tiRunning"
	list[enums.TaskTiStopped] = "tiStopped"
	list[enums.TaskTiUnusable] = "tiUnusable"
	return list
}

func simpleGetApplicationReferenceByName(value string) (int, error) {
	ret := -1
	for k, v := range simpleGetApplicationReference() {
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

func simpleGetHardwareResourceByName(value string) (int, error) {
	ret := -1
	for k, v := range simpleGetHardwareResource() {
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

func simpleGetVdeStateErrorByName(value string) (int, error) {
	ret := -1
	for k, v := range simpleGetVdeStateError() {
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

func simpleGetServiceByName(value string) (int, error) {
	ret := -1
	for k, v := range simpleGetService() {
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

func simpleGetDefinitionByName(value string) (int, error) {
	ret := -1
	for k, v := range simpleGetDefinition() {
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

func simpleGetAccessByName(value string) (int, error) {
	ret := -1
	for k, v := range simpleGetAccess() {
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

func simpleGetLoadDataSetByName(value string) (int, error) {
	ret := -1
	for k, v := range simpleGetLoadDataSet() {
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

func simpleGetTaskByName(value string) (int, error) {
	ret := -1
	for k, v := range simpleGetTask() {
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
func simpleGetGeneralTags(type_ enums.TranslatorOutputType, list map[int]string) {
	addTag(list, enums.CommandSnrm, "Snrm")
	addTag(list, enums.CommandUnacceptableFrame, "UnacceptableFrame")
	addTag(list, enums.CommandDisconnectMode, "DisconnectMode")
	addTag(list, enums.CommandUa, "Ua")
	addTag(list, enums.CommandAarq, "AssociationRequest")
	addTag(list, enums.CommandAare, "AssociationResponse")
	addTag(list, int(constants.TranslatorGeneralTagsApplicationContextName), "ApplicationContextName")
	addTag(list, enums.CommandInitiateResponse, "InitiateResponse")
	addTag(list, enums.CommandInitiateRequest, "InitiateRequest")
	addTag(list, int(constants.TranslatorGeneralTagsNegotiatedQualityOfService), "NegotiatedQualityOfService")
	addTag(list, int(constants.TranslatorGeneralTagsProposedQualityOfService), "ProposedQualityOfService")
	addTag(list, int(constants.TranslatorGeneralTagsProposedDlmsVersionNumber), "ProposedDlmsVersionNumber")
	addTag(list, int(constants.TranslatorGeneralTagsProposedMaxPduSize), "ProposedMaxPduSize")
	addTag(list, int(constants.TranslatorGeneralTagsProposedConformance), "ProposedConformance")
	addTag(list, int(constants.TranslatorGeneralTagsVaaName), "VaaName")
	addTag(list, int(constants.TranslatorGeneralTagsNegotiatedConformance), "NegotiatedConformance")
	addTag(list, int(constants.TranslatorGeneralTagsNegotiatedDlmsVersionNumber), "NegotiatedDlmsVersionNumber")
	addTag(list, int(constants.TranslatorGeneralTagsNegotiatedMaxPduSize), "NegotiatedMaxPduSize")
	addTag(list, int(constants.TranslatorGeneralTagsConformanceBit), "ConformanceBit")
	addTag(list, int(constants.TranslatorGeneralTagsSenderACSERequirements), "SenderACSERequirements")
	addTag(list, int(constants.TranslatorGeneralTagsResponderACSERequirement), "ResponderACSERequirement")
	addTag(list, int(constants.TranslatorGeneralTagsRespondingMechanismName), "MechanismName")
	addTag(list, int(constants.TranslatorGeneralTagsCallingMechanismName), "MechanismName")
	addTag(list, int(constants.TranslatorGeneralTagsCallingAuthentication), "CallingAuthentication")
	addTag(list, int(constants.TranslatorGeneralTagsRespondingAuthentication), "RespondingAuthentication")
	addTag(list, int(enums.CommandReleaseRequest), "ReleaseRequest")
	addTag(list, int(enums.CommandReleaseResponse), "ReleaseResponse")
	addTag(list, int(enums.CommandDisconnectRequest), "DisconnectRequest")
	addTag(list, int(constants.TranslatorGeneralTagsAssociationResult), "AssociationResult")
	addTag(list, int(constants.TranslatorGeneralTagsResultSourceDiagnostic), "ResultSourceDiagnostic")
	addTag(list, int(constants.TranslatorGeneralTagsACSEServiceUser), "ACSEServiceUser")
	addTag(list, int(constants.TranslatorGeneralTagsACSEServiceProvider), "ACSEServiceProvider")
	addTag(list, int(constants.TranslatorGeneralTagsCallingAPTitle), "CallingAPTitle")
	addTag(list, int(constants.TranslatorGeneralTagsRespondingAPTitle), "RespondingAPTitle")
	addTag(list, int(constants.TranslatorGeneralTagsDedicatedKey), "DedicatedKey")
	addTag(list, int(constants.TranslatorGeneralTagsUserInformation), "UserInformation")
	addTag(list, int(enums.CommandConfirmedServiceError), "ConfirmedServiceError")
	addTag(list, int(enums.CommandInformationReport), "InformationReportRequest")
	addTag(list, int(enums.CommandEventNotification), "EventNotificationRequest")
	addTag(list, int(enums.CommandGeneralBlockTransfer), "GeneralBlockTransfer")
	addTag(list, int(constants.TranslatorGeneralTagsCallingAeInvocationId), "CallingAEInvocationId")
	addTag(list, int(constants.TranslatorGeneralTagsCalledAeInvocationId), "CalledAEInvocationId")
	addTag(list, int(constants.TranslatorGeneralTagsRespondingAeInvocationId), "RespondingAeInvocationId")
	addTag(list, int(constants.TranslatorGeneralTagsCallingAeQualifier), "CallingAEQualifier")
	addTag(list, int(constants.TranslatorGeneralTagsPrimeNewDeviceNotification), "PrimeNewDeviceNotification")
	addTag(list, int(constants.TranslatorGeneralTagsPrimeRemoveDeviceNotification), "PrimeRemoveDeviceNotification")
	addTag(list, int(constants.TranslatorGeneralTagsPrimeStartReportingMeters), "PrimeStartReportingMeters")
	addTag(list, int(constants.TranslatorGeneralTagsPrimeDeleteMeters), "PrimeDeleteMeters")
	addTag(list, int(constants.TranslatorGeneralTagsPrimeEnableAutoClose), "PrimeEnableAutoClose")
	addTag(list, int(constants.TranslatorGeneralTagsPrimeDisableAutoClose), "PrimeDisableAutoClose")
	addTag(list, int(enums.CommandExceptionResponse), "ExceptionResponse")
	addTag(list, int(internal.TranslatorTagsStateError), "StateError")
}

// GetSnTags returns the get SN tags.
func simpleGetSnTags(type_ enums.TranslatorOutputType, list map[int]string) {
	list[int(enums.CommandReadRequest)] = "ReadRequest"
	list[int(enums.CommandWriteRequest)] = "WriteRequest"
	list[int(enums.CommandWriteResponse)] = "WriteResponse"
	list[int((enums.CommandReadRequest))<<8|int(constants.VariableAccessSpecificationVariableName)] = "VariableName"
	list[int((enums.CommandReadRequest))<<8|int(constants.VariableAccessSpecificationParameterisedAccess)] = "ParameterisedAccess"
	list[int((enums.CommandReadRequest))<<8|int(constants.VariableAccessSpecificationBlockNumberAccess)] = "BlockNumberAccess"
	list[int(enums.CommandWriteRequest)<<8|int(constants.VariableAccessSpecificationVariableName)] = "VariableName"
	list[int(enums.CommandReadResponse)] = "ReadResponse"
	list[int((enums.CommandReadResponse))<<8|int(constants.SingleReadResponseDataBlockResult)] = "DataBlockResult"
	list[int((enums.CommandReadResponse))<<8|int(constants.SingleReadResponseData)] = "Data"
	list[int(enums.CommandReadResponse)<<8|int(constants.SingleReadResponseDataAccessError)] = "DataAccessError"
}

// GetLnTags returns the get LN tags.
func simpleGetLnTags(type_ enums.TranslatorOutputType, list map[int]string) {
	addTag(list, enums.CommandGetRequest, "GetRequest")
	list[int((enums.CommandGetRequest))<<8|int(constants.GetCommandTypeNormal)] = "GetRequestNormal"
	list[int((enums.CommandGetRequest))<<8|int(constants.GetCommandTypeNextDataBlock)] = "GetRequestForNextDataBlock"
	list[int((enums.CommandGetRequest))<<8|int(constants.GetCommandTypeWithList)] = "GetRequestWithList"
	addTag(list, enums.CommandSetRequest, "SetRequest")
	list[int((enums.CommandSetRequest))<<8|int(constants.SetRequestTypeNormal)] = "SetRequestNormal"
	list[int((enums.CommandSetRequest))<<8|int(constants.SetRequestTypeFirstDataBlock)] = "SetRequestFirstDataBlock"
	list[int((enums.CommandSetRequest))<<8|int(constants.SetRequestTypeWithDataBlock)] = "SetRequestWithDataBlock"
	list[int((enums.CommandSetRequest))<<8|int(constants.SetRequestTypeWithList)] = "SetRequestWithList"
	addTag(list, enums.CommandMethodRequest, "ActionRequest")
	list[int((enums.CommandMethodRequest))<<8|int(constants.ActionRequestTypeNormal)] = "ActionRequestNormal"
	list[int((enums.CommandMethodRequest))<<8|int(constants.ActionRequestTypeNextBlock)] = "ActionRequestForNextPBlock"
	list[int((enums.CommandMethodRequest))<<8|int(constants.ActionRequestTypeWithList)] = "ActionRequestWithList"
	list[int(enums.CommandMethodRequest)<<8|int(constants.ActionRequestTypeWithFirstBlock)] = "ActionRequestWithFirstBlock"
	list[int(enums.CommandMethodRequest)<<8|int(constants.ActionRequestTypeWithListAndFirstBlock)] = "ActionRequestWithListAndFirstBlock"
	list[int(enums.CommandMethodRequest)<<8|int(constants.ActionRequestTypeWithBlock)] = "ActionRequestWithBlock"
	addTag(list, enums.CommandMethodResponse, "ActionResponse")
	list[int((enums.CommandMethodResponse))<<8|int(constants.ActionResponseTypeNormal)] = "ActionResponseNormal"
	list[int((enums.CommandMethodResponse))<<8|int(constants.ActionResponseTypeWithBlock)] = "ActionResponseWithPBlock"
	list[int((enums.CommandMethodResponse))<<8|int(constants.ActionResponseTypeWithList)] = "ActionResponseWithList"
	list[int((enums.CommandMethodResponse))<<8|int(constants.ActionResponseTypeNextBlock)] = "ActionResponseNextBlock"
	list[int(enums.CommandDataNotification)] = "DataNotification"
	list[int(enums.CommandGetResponse)] = "GetResponse"
	list[int(enums.CommandGetResponse)<<8|int(constants.GetCommandTypeNormal)] = "GetResponseNormal"
	list[int(enums.CommandGetResponse)<<8|int(constants.GetCommandTypeNextDataBlock)] = "GetResponsewithDataBlock"
	list[int(enums.CommandGetResponse)<<8|int(constants.GetCommandTypeWithList)] = "GetResponseWithList"
	list[int(enums.CommandSetResponse)] = "SetResponse"
	list[int(enums.CommandSetResponse)<<8|int(constants.SetResponseTypeNormal)] = "SetResponseNormal"
	list[int(enums.CommandSetResponse)<<8|int(constants.SetResponseTypeDataBlock)] = "SetResponseDataBlock"
	list[int(enums.CommandSetResponse)<<8|int(constants.SetResponseTypeLastDataBlock)] = "SetResponseWithLastDataBlock"
	list[int(enums.CommandSetResponse)<<8|int(constants.SetResponseTypeWithList)] = "SetResponseWithList"
	addTag(list, enums.CommandAccessRequest, "AccessRequest")
	list[int((enums.CommandAccessRequest))<<8|int(enums.AccessServiceCommandTypeGet)] = "AccessRequestGet"
	list[int((enums.CommandAccessRequest))<<8|int(enums.AccessServiceCommandTypeSet)] = "AccessRequestSet"
	list[int((enums.CommandAccessRequest))<<8|int(enums.AccessServiceCommandTypeAction)] = "AccessRequestAction"
	addTag(list, enums.CommandAccessResponse, "AccessResponse")
	list[int((enums.CommandAccessResponse))<<8|int(enums.AccessServiceCommandTypeGet)] = "AccessResponseGet"
	list[int((enums.CommandAccessResponse))<<8|int(enums.AccessServiceCommandTypeSet)] = "AccessResponseSet"
	list[int((enums.CommandAccessResponse))<<8|int(enums.AccessServiceCommandTypeAction)] = "AccessResponseAction"
	list[int(internal.TranslatorTagsAccessRequestBody)] = "AccessRequestBody"
	list[int(internal.TranslatorTagsListOfAccessRequestSpecification)] = "AccessRequestSpecification"
	list[int(internal.TranslatorTagsAccessRequestSpecification)] = "_AccessRequestSpecification"
	list[int(internal.TranslatorTagsAccessRequestListOfData)] = "AccessRequestListOfData"
	list[int(internal.TranslatorTagsAccessResponseBody)] = "AccessResponseBody"
	list[int(internal.TranslatorTagsListOfAccessResponseSpecification)] = "AccessResponseSpecification"
	list[int(internal.TranslatorTagsAccessResponseSpecification)] = "_AccessResponseSpecification"
	list[int(internal.TranslatorTagsAccessResponseListOfData)] = "AccessResponseListOfData"
	list[int(internal.TranslatorTagsService)] = "Service"
	list[int(internal.TranslatorTagsServiceError)] = "ServiceError"
	addTag(list, enums.CommandGatewayRequest, "GatewayRequest")
	addTag(list, enums.CommandGatewayResponse, "GatewayResponse")
}

// GetPlcTags returns the get PLC tags.
func simpleGetPlcTags(list map[int]string) {
	addTag(list, enums.CommandDiscoverRequest, "DiscoverRequest")
	addTag(list, enums.CommandDiscoverReport, "DiscoverReport")
	addTag(list, enums.CommandRegisterRequest, "RegisterRequest")
	addTag(list, enums.CommandPingRequest, "PingRequest")
	addTag(list, enums.CommandPingResponse, "PingResponse")
}

// GetGloTags returns the get glo tags.
func simpleGetGloTags(type_ enums.TranslatorOutputType, list map[int]string) {
	addTag(list, enums.CommandGloInitiateRequest, "glo_InitiateRequest")
	addTag(list, enums.CommandGloInitiateResponse, "glo_InitiateResponse")
	addTag(list, enums.CommandGloGetRequest, "glo_GetRequest")
	addTag(list, enums.CommandGloGetResponse, "glo_GetResponse")
	addTag(list, enums.CommandGloSetRequest, "glo_SetRequest")
	addTag(list, enums.CommandGloSetResponse, "glo_SetResponse")
	addTag(list, enums.CommandGloMethodRequest, "glo_ActionRequest")
	addTag(list, enums.CommandGloMethodResponse, "glo_ActionResponse")
	addTag(list, enums.CommandGloReadRequest, "glo_ReadRequest")
	addTag(list, enums.CommandGloReadResponse, "glo_ReadResponse")
	addTag(list, enums.CommandGloWriteRequest, "glo_WriteRequest")
	addTag(list, enums.CommandGloWriteResponse, "glo_WriteResponse")
	addTag(list, enums.CommandGeneralGloCiphering, "GeneralGloCiphering")
	addTag(list, enums.CommandGeneralCiphering, "GeneralCiphering")
	addTag(list, enums.CommandGeneralSigning, "GeneralSigning")
	addTag(list, enums.CommandGloConfirmedServiceError, "glo_GloConfirmedServiceError")
}

// GetDedTags returns the get ded tags.
func simpleGetDedTags(type_ enums.TranslatorOutputType, list map[int]string) {
	addTag(list, int(enums.CommandDedInitiateRequest), "ded_InitiateRequest")
	addTag(list, int(enums.CommandDedInitiateResponse), "ded_InitiateResponse")
	addTag(list, int(enums.CommandDedGetRequest), "ded_GetRequest")
	addTag(list, int(enums.CommandDedGetResponse), "ded_GetResponse")
	addTag(list, int(enums.CommandDedSetRequest), "ded_SetRequest")
	addTag(list, int(enums.CommandDedSetResponse), "ded_SetResponse")
	addTag(list, int(enums.CommandDedMethodRequest), "ded_ActionRequest")
	addTag(list, int(enums.CommandDedMethodResponse), "ded_ActionResponse")
	addTag(list, int(enums.CommandGeneralDedCiphering), "generalDedCiphering")
	addTag(list, int(enums.CommandDedConfirmedServiceError), "ded_GloConfirmedServiceError")
}

// getTranslatorTags returns the get translator tags.
func simpleGetTranslatorTags(type_ enums.TranslatorOutputType, list map[int]string) {
	addTag(list, int(internal.TranslatorTagsWrapper), "Wrapper")
	addTag(list, int(internal.TranslatorTagsHdlc), "Hdlc")
	addTag(list, int(internal.TranslatorTagsPduDlms), "Pdu")
	addTag(list, int(internal.TranslatorTagsTargetAddress), "TargetAddress")
	addTag(list, int(internal.TranslatorTagsSourceAddress), "SourceAddress")
	addTag(list, int(internal.TranslatorTagsListOfVariableAccessSpecification), "ListOfVariableAccessSpecification")
	addTag(list, int(internal.TranslatorTagsListOfData), "ListOfData")
	addTag(list, int(internal.TranslatorTagsSuccess), "Ok")
	addTag(list, int(internal.TranslatorTagsDataAccessError), "DataAccessError")
	addTag(list, int(internal.TranslatorTagsAttributeDescriptor), "AttributeDescriptor")
	addTag(list, int(internal.TranslatorTagsClassId), "ClassId")
	addTag(list, int(internal.TranslatorTagsInstanceId), "InstanceId")
	addTag(list, int(internal.TranslatorTagsAttributeId), "AttributeId")
	addTag(list, int(internal.TranslatorTagsMethodInvocationParameters), "MethodInvocationParameters")
	addTag(list, int(internal.TranslatorTagsSelector), "Selector")
	addTag(list, int(internal.TranslatorTagsParameter), "Parameter")
	addTag(list, int(internal.TranslatorTagsLastBlock), "LastBlock")
	addTag(list, int(internal.TranslatorTagsBlockNumber), "BlockNumber")
	addTag(list, int(internal.TranslatorTagsRawData), "RawData")
	addTag(list, int(internal.TranslatorTagsMethodDescriptor), "MethodDescriptor")
	addTag(list, int(internal.TranslatorTagsMethodId), "MethodId")
	addTag(list, int(internal.TranslatorTagsResult), "Result")
	addTag(list, int(internal.TranslatorTagsPblock), "PBlock")
	addTag(list, int(internal.TranslatorTagsContent), "Content")
	addTag(list, int(internal.TranslatorTagsSignature), "Signature")
	addTag(list, int(internal.TranslatorTagsReturnParameters), "ReturnParameters")
	addTag(list, int(internal.TranslatorTagsAccessSelection), "AccessSelection")
	addTag(list, int(internal.TranslatorTagsValue), "Value")
	addTag(list, int(internal.TranslatorTagsAccessSelector), "AccessSelector")
	addTag(list, int(internal.TranslatorTagsAccessParameters), "AccessParameters")
	addTag(list, int(internal.TranslatorTagsAttributeDescriptorList), "AttributeDescriptorList")
	addTag(list, int(internal.TranslatorTagsAttributeDescriptorWithSelection), "AttributeDescriptorWithSelection")
	addTag(list, int(internal.TranslatorTagsReadDataBlockAccess), "ReadDataBlockAccess")
	addTag(list, int(internal.TranslatorTagsWriteDataBlockAccess), "WriteDataBlockAccess")
	addTag(list, int(internal.TranslatorTagsData), "Data")
	addTag(list, int(internal.TranslatorTagsInvokeId), "InvokeIdAndPriority")
	addTag(list, int(internal.TranslatorTagsLongInvokeId), "LongInvokeIdAndPriority")
	addTag(list, int(internal.TranslatorTagsDateTime), "DateTime")
	addTag(list, int(internal.TranslatorTagsCurrentTime), "CurrentTime")
	addTag(list, int(internal.TranslatorTagsTime), "DateTime")
	addTag(list, int(internal.TranslatorTagsReason), "Reason")
	addTag(list, int(internal.TranslatorTagsNotificationBody), "NotificationBody")
	addTag(list, int(internal.TranslatorTagsDataValue), "DataValue")
	addTag(list, int(internal.TranslatorTagsCipheredService), "CipheredService")
	addTag(list, int(internal.TranslatorTagsSystemTitle), "SystemTitle")
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
	addTag(list, int(internal.TranslatorTagsAttributeValue), "AttributeValue")
	addTag(list, int(internal.TranslatorTagsMaxInfoRX), "MaxInfoRX")
	addTag(list, int(internal.TranslatorTagsMaxInfoTX), "MaxInfoTX")
	addTag(list, int(internal.TranslatorTagsWindowSizeRX), "WindowSizeRX")
	addTag(list, int(internal.TranslatorTagsWindowSizeTX), "WindowSizeTX")
	addTag(list, int(internal.TranslatorTagsValueList), "ValueList")
	addTag(list, int(internal.TranslatorTagsDataAccessResult), "DataAccessResult")
	addTag(list, int(internal.TranslatorTagsBlockControl), "BlockControl")
	addTag(list, int(internal.TranslatorTagsBlockNumberAck), "BlockNumberAck")
	addTag(list, int(internal.TranslatorTagsBlockData), "BlockData")
	addTag(list, int(internal.TranslatorTagsContentsDescription), "ContentsDescription")
	addTag(list, int(internal.TranslatorTagsArrayContents), "ArrayContents")
	addTag(list, int(internal.TranslatorTagsNetworkId), "NetworkId")
	addTag(list, int(internal.TranslatorTagsPhysicalDeviceAddress), "PhysicalDeviceAddress")
	addTag(list, int(internal.TranslatorTagsProtocolVersion), "ProtocolVersion")
	addTag(list, int(internal.TranslatorTagsCalledAPTitle), "CalledAPTitle")
	addTag(list, int(internal.TranslatorTagsCalledAPInvocationId), "CalledAPInvocationId")
	addTag(list, int(internal.TranslatorTagsCalledAEInvocationId), "CalledAEInvocationId")
	addTag(list, int(internal.TranslatorTagsCallingApInvocationId), "CallingAPInvocationId")
	addTag(list, int(internal.TranslatorTagsCalledAEQualifier), "CalledAEQualifier")
}

func simpleGetServiceErrorValue(error enums.ServiceError, value uint8) string {
	switch error {
	case enums.ServiceErrorApplicationReference:
		return simpleGetApplicationReference()[enums.ApplicationReference(value)]
	case enums.ServiceErrorHardwareResource:
		return simpleGetHardwareResource()[enums.HardwareResource(value)]
	case enums.ServiceErrorVdeStateError:
		return simpleGetVdeStateError()[enums.VdeStateError(value)]
	case enums.ServiceErrorService:
		return simpleGetService()[enums.Service(value)]
	case enums.ServiceErrorDefinition:
		return simpleGetDefinition()[enums.Definition(value)]
	case enums.ServiceErrorAccess:
		return simpleGetAccess()[enums.Access(value)]
	case enums.ServiceErrorInitiate:
		return simpleGetInitiate()[enums.Initiate(value)]
	case enums.ServiceErrorLoadDataSet:
		return simpleGetLoadDataSet()[enums.LoadDataSet(value)]
	case enums.ServiceErrorTask:
		return simpleGetTask()[enums.Task(value)]
	case enums.ServiceErrorOtherError:
		return strconv.Itoa(int(value))
	default:
	}
	return ""
}

// ServiceErrorToString returns the parameters:
//
//	error: Service error enumeration value.
//
// Returns:
//
//	Service error simple XML tag.
func simpleServiceErrorToString(error enums.ServiceError) string {
	return simpleGetServiceErrors()[error]
}

// GetServiceErrorByName returns the parameters:
//
//	value: Service error simple XML tag.
//
// Returns:
//
//	Service error enumeration value.
func simpleGetServiceErrorByName(value string) (enums.ServiceError, error) {
	for k, v := range simpleGetServiceErrors() {
		if strings.EqualFold(value, v) {
			return k, nil
		}
	}
	return 0, gxcommon.ErrInvalidArgument
}

func simpleGetError(serviceError enums.ServiceError, value string) (uint8, error) {
	var err error
	ret := 0
	switch serviceError {
	case enums.ServiceErrorApplicationReference:
		ret, err = simpleGetApplicationReferenceByName(value)
	case enums.ServiceErrorHardwareResource:
		ret, err = simpleGetHardwareResourceByName(value)
	case enums.ServiceErrorVdeStateError:
		ret, err = simpleGetVdeStateErrorByName(value)
	case enums.ServiceErrorService:
		ret, err = simpleGetServiceByName(value)
	case enums.ServiceErrorDefinition:
		ret, err = simpleGetDefinitionByName(value)
	case enums.ServiceErrorAccess:
		ret, err = simpleGetAccessByName(value)
	case enums.ServiceErrorInitiate:
		ret, err = simpleGetInitiateByName(value)
	case enums.ServiceErrorLoadDataSet:
		ret, err = simpleGetLoadDataSetByName(value)
	case enums.ServiceErrorTask:
		ret, err = simpleGetTaskByName(value)
	case enums.ServiceErrorOtherError:
		ret, err = strconv.Atoi(value)
		if err != nil {
			return 0, err
		}
	default:
		err = gxcommon.ErrInvalidArgument
	}
	return uint8(ret), err
}

func simpleGetDataTypeTags(list map[int]string) {
	list[settings.DataTypeOffset+int(enums.DataTypeNone)] = "None"
	list[settings.DataTypeOffset+int(enums.DataTypeArray)] = "Array"
	list[settings.DataTypeOffset+int(enums.DataTypeBcd)] = "Bcd"
	list[settings.DataTypeOffset+int(enums.DataTypeBitString)] = "BitString"
	list[settings.DataTypeOffset+int(enums.DataTypeBoolean)] = "Boolean"
	list[settings.DataTypeOffset+int(enums.DataTypeCompactArray)] = "CompactArray"
	list[settings.DataTypeOffset+int(enums.DataTypeDate)] = "Date"
	list[settings.DataTypeOffset+int(enums.DataTypeDateTime)] = "DateTime"
	list[settings.DataTypeOffset+int(enums.DataTypeEnum)] = "Enum"
	list[settings.DataTypeOffset+int(enums.DataTypeFloat32)] = "Float32"
	list[settings.DataTypeOffset+int(enums.DataTypeFloat64)] = "Float64"
	list[settings.DataTypeOffset+int(enums.DataTypeInt16)] = "Int16"
	list[settings.DataTypeOffset+int(enums.DataTypeInt32)] = "Int32"
	list[settings.DataTypeOffset+int(enums.DataTypeInt64)] = "Int64"
	list[settings.DataTypeOffset+int(enums.DataTypeInt8)] = "Int8"
	list[settings.DataTypeOffset+int(enums.DataTypeOctetString)] = "OctetString"
	list[settings.DataTypeOffset+int(enums.DataTypeString)] = "String"
	list[settings.DataTypeOffset+int(enums.DataTypeStringUTF8)] = "StringUTF8"
	list[settings.DataTypeOffset+int(enums.DataTypeStructure)] = "Structure"
	list[settings.DataTypeOffset+int(enums.DataTypeTime)] = "Time"
	list[settings.DataTypeOffset+int(enums.DataTypeUint16)] = "UInt16"
	list[settings.DataTypeOffset+int(enums.DataTypeUint32)] = "UInt32"
	list[settings.DataTypeOffset+int(enums.DataTypeUint64)] = "UInt64"
	list[settings.DataTypeOffset+int(enums.DataTypeUint8)] = "UInt8"
	list[settings.DataTypeOffset+int(enums.DataTypeDeltaInt8)] = "Delta-Int8"
	list[settings.DataTypeOffset+int(enums.DataTypeDeltaInt16)] = "Delta-Int16"
	list[settings.DataTypeOffset+int(enums.DataTypeDeltaInt32)] = "Delta-Int32"
	list[settings.DataTypeOffset+int(enums.DataTypeDeltaUint8)] = "Delta-UInt8"
	list[settings.DataTypeOffset+int(enums.DataTypeDeltaUint16)] = "Delta-UInt16"
	list[settings.DataTypeOffset+int(enums.DataTypeDeltaUint32)] = "Delta-UInt32"
}

func simpleErrorCodeToString(value enums.ErrorCode) (string, error) {
	var str string
	switch value {
	case enums.ErrorCodeAccessViolated:
		str = "AccessViolated"
	case enums.ErrorCodeDataBlockNumberInvalid:
		str = "DataBlockNumberInvalid"
	case enums.ErrorCodeDataBlockUnavailable:
		str = "DataBlockUnavailable"
	case enums.ErrorCodeHardwareFault:
		str = "HardwareFault"
	case enums.ErrorCodeInconsistentClass:
		str = "InconsistentClass"
	case enums.ErrorCodeLongGetOrReadAborted:
		str = "LongGetOrReadAborted"
	case enums.ErrorCodeLongSetOrWriteAborted:
		str = "LongSetOrWriteAborted"
	case enums.ErrorCodeNoLongGetOrReadInProgress:
		str = "NoLongGetOrReadInProgress"
	case enums.ErrorCodeNoLongSetOrWriteInProgress:
		str = "NoLongSetOrWriteInProgress"
	case enums.ErrorCodeOk:
		str = "Success"
	case enums.ErrorCodeOtherReason:
		str = "OtherReason"
	case enums.ErrorCodeReadWriteDenied:
		str = "ReadWriteDenied"
	case enums.ErrorCodeTemporaryFailure:
		str = "TemporaryFailure"
	case enums.ErrorCodeUnavailableObject:
		str = "UnavailableObject"
	case enums.ErrorCodeUndefinedObject:
		str = "UndefinedObject"
	case enums.ErrorCodeUnmatchedType:
		str = "UnmatchedType"
	default:
		return "", fmt.Errorf("Error code: %d", value)
	}
	return str, nil
}

func simpleValueOfErrorCode(value string) (enums.ErrorCode, error) {
	var v enums.ErrorCode
	if strings.EqualFold("AccessViolated", value) {
		v = enums.ErrorCodeAccessViolated
	} else if strings.EqualFold("DataBlockNumberInvalid", value) {
		v = enums.ErrorCodeDataBlockNumberInvalid
	} else if strings.EqualFold("DataBlockUnavailable", value) {
		v = enums.ErrorCodeDataBlockUnavailable
	} else if strings.EqualFold("HardwareFault", value) {
		v = enums.ErrorCodeHardwareFault
	} else if strings.EqualFold("InconsistentClass", value) {
		v = enums.ErrorCodeInconsistentClass
	} else if strings.EqualFold("LongGetOrReadAborted", value) {
		v = enums.ErrorCodeLongGetOrReadAborted
	} else if strings.EqualFold("LongSetOrWriteAborted", value) {
		v = enums.ErrorCodeLongSetOrWriteAborted
	} else if strings.EqualFold("NoLongGetOrReadInProgress", value) {
		v = enums.ErrorCodeNoLongGetOrReadInProgress
	} else if strings.EqualFold("NoLongSetOrWriteInProgress", value) {
		v = enums.ErrorCodeNoLongSetOrWriteInProgress
	} else if strings.EqualFold("Ok", value) {
		v = enums.ErrorCodeOk
	} else if strings.EqualFold("OtherReason", value) {
		v = enums.ErrorCodeOtherReason
	} else if strings.EqualFold("ReadWriteDenied", value) {
		v = enums.ErrorCodeReadWriteDenied
	} else if strings.EqualFold("TemporaryFailure", value) {
		v = enums.ErrorCodeTemporaryFailure
	} else if strings.EqualFold("UnavailableObject", value) {
		v = enums.ErrorCodeUnavailableObject
	} else if strings.EqualFold("UndefinedObject", value) {
		v = enums.ErrorCodeUndefinedObject
	} else if strings.EqualFold("UnmatchedType", value) {
		v = enums.ErrorCodeUnmatchedType
	} else {
		return 0, fmt.Errorf("Error code: %s", value)
	}
	return v, nil
}

func simpleGetInitiateByName(value string) (int, error) {
	ret := -1
	for k, v := range simpleGetInitiate() {
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

func simpleReleaseResponseReasonToString(value constants.ReleaseResponseReason) (string, error) {
	var str string
	switch value {
	case constants.ReleaseResponseReasonNormal:
		str = "Normal"
	case constants.ReleaseResponseReasonNotFinished:
		str = "NotFinished"
	case constants.ReleaseResponseReasonUserDefined:
		str = "UserDefined"
	default:
		return "", gxcommon.ErrInvalidArgument
	}
	return str, nil
}

func simplevalueOfReleaseResponseReason(value string) (constants.ReleaseResponseReason, error) {
	var ret constants.ReleaseResponseReason
	if strings.EqualFold(value, "Normal") {
		ret = constants.ReleaseResponseReasonNormal
	} else if strings.EqualFold(value, "NotFinished") {
		ret = constants.ReleaseResponseReasonNotFinished
	} else if strings.EqualFold(value, "UserDefined") {
		ret = constants.ReleaseResponseReasonUserDefined
	} else {
		return 0, gxcommon.ErrInvalidArgument
	}
	return ret, nil
}

func simpleReleaseRequestReasonToString(value constants.ReleaseRequestReason) (string, error) {
	var str string
	switch value {
	case constants.ReleaseRequestReasonNormal:
		str = "Normal"
	case constants.ReleaseRequestReasonUrgent:
		str = "Urgent"
	case constants.ReleaseRequestReasonUserDefined:
		str = "UserDefined"
	default:
		return "", gxcommon.ErrInvalidArgument
	}
	return str, nil
}

func simplevalueOfReleaseRequestReason(value string) (constants.ReleaseRequestReason, error) {
	var ret constants.ReleaseRequestReason
	if strings.EqualFold(value, "Normal") {
		ret = constants.ReleaseRequestReasonNormal
	} else if strings.EqualFold(value, "Urgent") {
		ret = constants.ReleaseRequestReasonUrgent
	} else if strings.EqualFold(value, "UserDefined") {
		ret = constants.ReleaseRequestReasonUserDefined
	} else {
		return 0, gxcommon.ErrInvalidArgument
	}
	return ret, nil
}

// simpleStateErrorToString returns the gets state error description.
//
// Parameters:
//
//	error: State error enumerator value.
//
// Returns:
//
//	State error as an string.
func simpleStateErrorToString(error enums.ExceptionStateError) (string, error) {
	switch error {
	case enums.ExceptionStateErrorServiceNotAllowed:
		return "ServiceNotAllowed", nil
	case enums.ExceptionStateErrorServiceUnknown:
		return "ServiceUnknown", nil
	default:
		return "", gxcommon.ErrInvalidArgument
	}
}

// simpleExceptionServiceErrorToString returns the gets service error description.
//
// Parameters:
//
//	error: Service error enumerator value.
//
// Returns:
//
//	Service error as an string.
func simpleExceptionServiceErrorToString(error enums.ExceptionServiceError) (string, error) {
	switch error {
	case enums.ExceptionServiceErrorOperationNotPossible:
		return "OperationNotPossible", nil
	case enums.ExceptionServiceErrorServiceNotSupported:
		return "ServiceNotSupported", nil
	case enums.ExceptionServiceErrorOtherReason:
		return "OtherReason", nil
	case enums.ExceptionServiceErrorPduTooLong:
		return "PduTooLong", nil
	case enums.ExceptionServiceErrorDecipheringError:
		return "DecipheringError", nil
	case enums.ExceptionServiceErrorInvocationCounterError:
		return "InvocationCounterError", nil
	default:
		return "", gxcommon.ErrInvalidArgument
	}
}

// simplevalueofStateError returns the parameters:
//
//	value: State error string value.
//
// Returns:
//
//	State error enum value.
func simplevalueofStateError(value string) (enums.ExceptionStateError, error) {
	if "ServiceNotAllowed" == value {
		return enums.ExceptionStateErrorServiceNotAllowed, nil
	}
	if "ServiceUnknown" == value {
		return enums.ExceptionStateErrorServiceUnknown, nil
	}
	return 0, gxcommon.ErrInvalidArgument
}

// simplevalueOfExceptionServiceError returns the parameters:
//
//	value: Service error string value.
//
// Returns:
//
//	Service error enum value.
func simplevalueOfExceptionServiceError(value string) (enums.ExceptionServiceError, error) {
	if value == "OperationNotPossible" {
		return enums.ExceptionServiceErrorOperationNotPossible, nil
	}
	if value == "ServiceNotSupported" {
		return enums.ExceptionServiceErrorServiceNotSupported, nil
	}
	if value == "OtherReason" {
		return enums.ExceptionServiceErrorOtherReason, nil
	}
	if value == "PduTooLong" {
		return enums.ExceptionServiceErrorPduTooLong, nil
	}
	if value == "DecipheringError" {
		return enums.ExceptionServiceErrorDecipheringError, nil
	}
	if value == "InvocationCounterError" {
		return enums.ExceptionServiceErrorInvocationCounterError, nil
	}
	return 0, gxcommon.ErrInvalidArgument
}
