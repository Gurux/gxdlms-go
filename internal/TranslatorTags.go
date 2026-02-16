package internal

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

type TranslatorTags int

const (
	TranslatorTagsWrapper                           TranslatorTags = 0xFF01
	TranslatorTagsHdlc                              TranslatorTags = 65282
	TranslatorTagsPduDlms                           TranslatorTags = 65283
	TranslatorTagsPduCse                            TranslatorTags = 65284
	TranslatorTagsTargetAddress                     TranslatorTags = 65285
	TranslatorTagsSourceAddress                     TranslatorTags = 65286
	TranslatorTagsListOfVariableAccessSpecification TranslatorTags = 65287
	TranslatorTagsListOfData                        TranslatorTags = 65288
	TranslatorTagsSuccess                           TranslatorTags = 65289
	TranslatorTagsDataAccessError                   TranslatorTags = 65290
	TranslatorTagsAttributeDescriptor               TranslatorTags = 65291
	TranslatorTagsClassId                           TranslatorTags = 65292
	TranslatorTagsInstanceId                        TranslatorTags = 65293
	TranslatorTagsAttributeId                       TranslatorTags = 65294
	TranslatorTagsMethodInvocationParameters        TranslatorTags = 65295
	TranslatorTagsSelector                          TranslatorTags = 65296
	TranslatorTagsParameter                         TranslatorTags = 65297
	TranslatorTagsLastBlock                         TranslatorTags = 65298
	TranslatorTagsBlockNumber                       TranslatorTags = 65299
	TranslatorTagsRawData                           TranslatorTags = 65300
	TranslatorTagsMethodDescriptor                  TranslatorTags = 65301
	TranslatorTagsMethodId                          TranslatorTags = 65302
	TranslatorTagsResult                            TranslatorTags = 65303
	TranslatorTagsReturnParameters                  TranslatorTags = 65304
	TranslatorTagsAccessSelection                   TranslatorTags = 65305
	TranslatorTagsValue                             TranslatorTags = 65306
	TranslatorTagsAccessSelector                    TranslatorTags = 65307
	TranslatorTagsAccessParameters                  TranslatorTags = 65308
	TranslatorTagsAttributeDescriptorList           TranslatorTags = 65309
	TranslatorTagsAttributeDescriptorWithSelection  TranslatorTags = 65310
	TranslatorTagsReadDataBlockAccess               TranslatorTags = 65311
	TranslatorTagsWriteDataBlockAccess              TranslatorTags = 65312
	TranslatorTagsData                              TranslatorTags = 65313
	TranslatorTagsInvokeId                          TranslatorTags = 65314
	TranslatorTagsLongInvokeId                      TranslatorTags = 65315
	TranslatorTagsDateTime                          TranslatorTags = 65316
	TranslatorTagsReason                            TranslatorTags = 65317
	TranslatorTagsVariableAccessSpecification       TranslatorTags = 65318
	TranslatorTagsChoice                            TranslatorTags = 65319
	TranslatorTagsNotificationBody                  TranslatorTags = 65320
	TranslatorTagsDataValue                         TranslatorTags = 65321
	TranslatorTagsAccessRequestBody                 TranslatorTags = 65322
	TranslatorTagsListOfAccessRequestSpecification  TranslatorTags = 65323
	TranslatorTagsAccessRequestSpecification        TranslatorTags = 65324
	TranslatorTagsAccessRequestListOfData           TranslatorTags = 65325
	TranslatorTagsAccessResponseBody                TranslatorTags = 65326
	TranslatorTagsListOfAccessResponseSpecification TranslatorTags = 65327
	TranslatorTagsAccessResponseSpecification       TranslatorTags = 65328
	TranslatorTagsAccessResponseListOfData          TranslatorTags = 65329
	TranslatorTagsSingleResponse                    TranslatorTags = 65330
	TranslatorTagsService                           TranslatorTags = 65331
	TranslatorTagsServiceError                      TranslatorTags = 65332
	TranslatorTagsInitiateError                     TranslatorTags = 65333
	TranslatorTagsCipheredService                   TranslatorTags = 65334
	TranslatorTagsSystemTitle                       TranslatorTags = 65335
	TranslatorTagsDataBlock                         TranslatorTags = 65336
	TranslatorTagsTransactionId                     TranslatorTags = 65337
	TranslatorTagsOriginatorSystemTitle             TranslatorTags = 65338
	TranslatorTagsRecipientSystemTitle              TranslatorTags = 65339
	TranslatorTagsOtherInformation                  TranslatorTags = 65340
	TranslatorTagsKeyInfo                           TranslatorTags = 65341
	TranslatorTagsAgreedKey                         TranslatorTags = 65342
	TranslatorTagsKeyParameters                     TranslatorTags = 65343
	TranslatorTagsKeyCipheredData                   TranslatorTags = 65344
	TranslatorTagsCipheredContent                   TranslatorTags = 65345
	TranslatorTagsAttributeValue                    TranslatorTags = 65346
	TranslatorTagsCurrentTime                       TranslatorTags = 65347
	TranslatorTagsTime                              TranslatorTags = 65348
	TranslatorTagsMaxInfoRX                         TranslatorTags = 65349
	TranslatorTagsMaxInfoTX                         TranslatorTags = 65350
	TranslatorTagsWindowSizeRX                      TranslatorTags = 65351
	TranslatorTagsWindowSizeTX                      TranslatorTags = 65352
	TranslatorTagsValueList                         TranslatorTags = 65353
	TranslatorTagsDataAccessResult                  TranslatorTags = 65354
	TranslatorTagsFrameType                         TranslatorTags = 65355
	TranslatorTagsBlockControl                      TranslatorTags = 65356
	TranslatorTagsBlockNumberAck                    TranslatorTags = 65357
	TranslatorTagsBlockData                         TranslatorTags = 65358
	TranslatorTagsContentsDescription               TranslatorTags = 65359
	TranslatorTagsArrayContents                     TranslatorTags = 65360
	TranslatorTagsNetworkId                         TranslatorTags = 65361
	TranslatorTagsPhysicalDeviceAddress             TranslatorTags = 65362
	TranslatorTagsProtocolVersion                   TranslatorTags = 65363
	TranslatorTagsCalledAPTitle                     TranslatorTags = 65364
	TranslatorTagsCalledAPInvocationId              TranslatorTags = 65365
	TranslatorTagsCalledAEInvocationId              TranslatorTags = 65366
	TranslatorTagsCallingApInvocationId             TranslatorTags = 65367
	TranslatorTagsCalledAEQualifier                 TranslatorTags = 65368
	TranslatorTagsResponseAllowed                   TranslatorTags = 65369
	TranslatorTagsExceptionResponse                 TranslatorTags = 65370
	TranslatorTagsStateError                        TranslatorTags = 65371
	TranslatorTagsPblock                            TranslatorTags = 65372
	TranslatorTagsContent                           TranslatorTags = 65373
	TranslatorTagsSignature                         TranslatorTags = 65374
)

// TranslatorTagsParse converts the given string into a TranslatorTags value.
//
// It returns the corresponding TranslatorTags constant if the string matches
// a known level name, or an error if the input is invalid.
func TranslatorTagsParse(value string) (TranslatorTags, error) {
	var ret TranslatorTags
	var err error
	switch strings.ToUpper(value) {
	case "WRAPPER":
		ret = TranslatorTagsWrapper
	case "HDLC":
		ret = TranslatorTagsHdlc
	case "PDUDLMS":
		ret = TranslatorTagsPduDlms
	case "PDUCSE":
		ret = TranslatorTagsPduCse
	case "TARGETADDRESS":
		ret = TranslatorTagsTargetAddress
	case "SOURCEADDRESS":
		ret = TranslatorTagsSourceAddress
	case "LISTOFVARIABLEACCESSSPECIFICATION":
		ret = TranslatorTagsListOfVariableAccessSpecification
	case "LISTOFDATA":
		ret = TranslatorTagsListOfData
	case "SUCCESS":
		ret = TranslatorTagsSuccess
	case "DATAACCESSERROR":
		ret = TranslatorTagsDataAccessError
	case "ATTRIBUTEDESCRIPTOR":
		ret = TranslatorTagsAttributeDescriptor
	case "CLASSID":
		ret = TranslatorTagsClassId
	case "INSTANCEID":
		ret = TranslatorTagsInstanceId
	case "ATTRIBUTEID":
		ret = TranslatorTagsAttributeId
	case "METHODINVOCATIONPARAMETERS":
		ret = TranslatorTagsMethodInvocationParameters
	case "SELECTOR":
		ret = TranslatorTagsSelector
	case "PARAMETER":
		ret = TranslatorTagsParameter
	case "LASTBLOCK":
		ret = TranslatorTagsLastBlock
	case "BLOCKNUMBER":
		ret = TranslatorTagsBlockNumber
	case "RAWDATA":
		ret = TranslatorTagsRawData
	case "METHODDESCRIPTOR":
		ret = TranslatorTagsMethodDescriptor
	case "METHODID":
		ret = TranslatorTagsMethodId
	case "RESULT":
		ret = TranslatorTagsResult
	case "RETURNPARAMETERS":
		ret = TranslatorTagsReturnParameters
	case "ACCESSSELECTION":
		ret = TranslatorTagsAccessSelection
	case "VALUE":
		ret = TranslatorTagsValue
	case "ACCESSSELECTOR":
		ret = TranslatorTagsAccessSelector
	case "ACCESSPARAMETERS":
		ret = TranslatorTagsAccessParameters
	case "ATTRIBUTEDESCRIPTORLIST":
		ret = TranslatorTagsAttributeDescriptorList
	case "ATTRIBUTEDESCRIPTORWITHSELECTION":
		ret = TranslatorTagsAttributeDescriptorWithSelection
	case "READDATABLOCKACCESS":
		ret = TranslatorTagsReadDataBlockAccess
	case "WRITEDATABLOCKACCESS":
		ret = TranslatorTagsWriteDataBlockAccess
	case "DATA":
		ret = TranslatorTagsData
	case "INVOKEID":
		ret = TranslatorTagsInvokeId
	case "LONGINVOKEID":
		ret = TranslatorTagsLongInvokeId
	case "DATETIME":
		ret = TranslatorTagsDateTime
	case "REASON":
		ret = TranslatorTagsReason
	case "VARIABLEACCESSSPECIFICATION":
		ret = TranslatorTagsVariableAccessSpecification
	case "CHOICE":
		ret = TranslatorTagsChoice
	case "NOTIFICATIONBODY":
		ret = TranslatorTagsNotificationBody
	case "DATAVALUE":
		ret = TranslatorTagsDataValue
	case "ACCESSREQUESTBODY":
		ret = TranslatorTagsAccessRequestBody
	case "LISTOFACCESSREQUESTSPECIFICATION":
		ret = TranslatorTagsListOfAccessRequestSpecification
	case "ACCESSREQUESTSPECIFICATION":
		ret = TranslatorTagsAccessRequestSpecification
	case "ACCESSREQUESTLISTOFDATA":
		ret = TranslatorTagsAccessRequestListOfData
	case "ACCESSRESPONSEBODY":
		ret = TranslatorTagsAccessResponseBody
	case "LISTOFACCESSRESPONSESPECIFICATION":
		ret = TranslatorTagsListOfAccessResponseSpecification
	case "ACCESSRESPONSESPECIFICATION":
		ret = TranslatorTagsAccessResponseSpecification
	case "ACCESSRESPONSELISTOFDATA":
		ret = TranslatorTagsAccessResponseListOfData
	case "SINGLERESPONSE":
		ret = TranslatorTagsSingleResponse
	case "SERVICE":
		ret = TranslatorTagsService
	case "SERVICEERROR":
		ret = TranslatorTagsServiceError
	case "INITIATEERROR":
		ret = TranslatorTagsInitiateError
	case "CIPHEREDSERVICE":
		ret = TranslatorTagsCipheredService
	case "SYSTEMTITLE":
		ret = TranslatorTagsSystemTitle
	case "DATABLOCK":
		ret = TranslatorTagsDataBlock
	case "TRANSACTIONID":
		ret = TranslatorTagsTransactionId
	case "ORIGINATORSYSTEMTITLE":
		ret = TranslatorTagsOriginatorSystemTitle
	case "RECIPIENTSYSTEMTITLE":
		ret = TranslatorTagsRecipientSystemTitle
	case "OTHERINFORMATION":
		ret = TranslatorTagsOtherInformation
	case "KEYINFO":
		ret = TranslatorTagsKeyInfo
	case "AGREEDKEY":
		ret = TranslatorTagsAgreedKey
	case "KEYPARAMETERS":
		ret = TranslatorTagsKeyParameters
	case "KEYCIPHEREDDATA":
		ret = TranslatorTagsKeyCipheredData
	case "CIPHEREDCONTENT":
		ret = TranslatorTagsCipheredContent
	case "ATTRIBUTEVALUE":
		ret = TranslatorTagsAttributeValue
	case "CURRENTTIME":
		ret = TranslatorTagsCurrentTime
	case "TIME":
		ret = TranslatorTagsTime
	case "MAXINFORX":
		ret = TranslatorTagsMaxInfoRX
	case "MAXINFOTX":
		ret = TranslatorTagsMaxInfoTX
	case "WINDOWSIZERX":
		ret = TranslatorTagsWindowSizeRX
	case "WINDOWSIZETX":
		ret = TranslatorTagsWindowSizeTX
	case "VALUELIST":
		ret = TranslatorTagsValueList
	case "DATAACCESSRESULT":
		ret = TranslatorTagsDataAccessResult
	case "FRAMETYPE":
		ret = TranslatorTagsFrameType
	case "BLOCKCONTROL":
		ret = TranslatorTagsBlockControl
	case "BLOCKNUMBERACK":
		ret = TranslatorTagsBlockNumberAck
	case "BLOCKDATA":
		ret = TranslatorTagsBlockData
	case "CONTENTSDESCRIPTION":
		ret = TranslatorTagsContentsDescription
	case "ARRAYCONTENTS":
		ret = TranslatorTagsArrayContents
	case "NETWORKID":
		ret = TranslatorTagsNetworkId
	case "PHYSICALDEVICEADDRESS":
		ret = TranslatorTagsPhysicalDeviceAddress
	case "PROTOCOLVERSION":
		ret = TranslatorTagsProtocolVersion
	case "CALLEDAPTITLE":
		ret = TranslatorTagsCalledAPTitle
	case "CALLEDAPINVOCATIONID":
		ret = TranslatorTagsCalledAPInvocationId
	case "CALLEDAEINVOCATIONID":
		ret = TranslatorTagsCalledAEInvocationId
	case "CALLINGAPINVOCATIONID":
		ret = TranslatorTagsCallingApInvocationId
	case "CALLEDAEQUALIFIER":
		ret = TranslatorTagsCalledAEQualifier
	case "RESPONSEALLOWED":
		ret = TranslatorTagsResponseAllowed
	case "EXCEPTIONRESPONSE":
		ret = TranslatorTagsExceptionResponse
	case "STATEERROR":
		ret = TranslatorTagsStateError
	case "PBLOCK":
		ret = TranslatorTagsPblock
	case "CONTENT":
		ret = TranslatorTagsContent
	case "SIGNATURE":
		ret = TranslatorTagsSignature
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the TranslatorTags.
// It satisfies fmt.Stringer.
func (g TranslatorTags) String() string {
	var ret string
	switch g {
	case TranslatorTagsWrapper:
		ret = "WRAPPER"
	case TranslatorTagsHdlc:
		ret = "HDLC"
	case TranslatorTagsPduDlms:
		ret = "PDUDLMS"
	case TranslatorTagsPduCse:
		ret = "PDUCSE"
	case TranslatorTagsTargetAddress:
		ret = "TARGETADDRESS"
	case TranslatorTagsSourceAddress:
		ret = "SOURCEADDRESS"
	case TranslatorTagsListOfVariableAccessSpecification:
		ret = "LISTOFVARIABLEACCESSSPECIFICATION"
	case TranslatorTagsListOfData:
		ret = "LISTOFDATA"
	case TranslatorTagsSuccess:
		ret = "SUCCESS"
	case TranslatorTagsDataAccessError:
		ret = "DATAACCESSERROR"
	case TranslatorTagsAttributeDescriptor:
		ret = "ATTRIBUTEDESCRIPTOR"
	case TranslatorTagsClassId:
		ret = "CLASSID"
	case TranslatorTagsInstanceId:
		ret = "INSTANCEID"
	case TranslatorTagsAttributeId:
		ret = "ATTRIBUTEID"
	case TranslatorTagsMethodInvocationParameters:
		ret = "METHODINVOCATIONPARAMETERS"
	case TranslatorTagsSelector:
		ret = "SELECTOR"
	case TranslatorTagsParameter:
		ret = "PARAMETER"
	case TranslatorTagsLastBlock:
		ret = "LASTBLOCK"
	case TranslatorTagsBlockNumber:
		ret = "BLOCKNUMBER"
	case TranslatorTagsRawData:
		ret = "RAWDATA"
	case TranslatorTagsMethodDescriptor:
		ret = "METHODDESCRIPTOR"
	case TranslatorTagsMethodId:
		ret = "METHODID"
	case TranslatorTagsResult:
		ret = "RESULT"
	case TranslatorTagsReturnParameters:
		ret = "RETURNPARAMETERS"
	case TranslatorTagsAccessSelection:
		ret = "ACCESSSELECTION"
	case TranslatorTagsValue:
		ret = "VALUE"
	case TranslatorTagsAccessSelector:
		ret = "ACCESSSELECTOR"
	case TranslatorTagsAccessParameters:
		ret = "ACCESSPARAMETERS"
	case TranslatorTagsAttributeDescriptorList:
		ret = "ATTRIBUTEDESCRIPTORLIST"
	case TranslatorTagsAttributeDescriptorWithSelection:
		ret = "ATTRIBUTEDESCRIPTORWITHSELECTION"
	case TranslatorTagsReadDataBlockAccess:
		ret = "READDATABLOCKACCESS"
	case TranslatorTagsWriteDataBlockAccess:
		ret = "WRITEDATABLOCKACCESS"
	case TranslatorTagsData:
		ret = "DATA"
	case TranslatorTagsInvokeId:
		ret = "INVOKEID"
	case TranslatorTagsLongInvokeId:
		ret = "LONGINVOKEID"
	case TranslatorTagsDateTime:
		ret = "DATETIME"
	case TranslatorTagsReason:
		ret = "REASON"
	case TranslatorTagsVariableAccessSpecification:
		ret = "VARIABLEACCESSSPECIFICATION"
	case TranslatorTagsChoice:
		ret = "CHOICE"
	case TranslatorTagsNotificationBody:
		ret = "NOTIFICATIONBODY"
	case TranslatorTagsDataValue:
		ret = "DATAVALUE"
	case TranslatorTagsAccessRequestBody:
		ret = "ACCESSREQUESTBODY"
	case TranslatorTagsListOfAccessRequestSpecification:
		ret = "LISTOFACCESSREQUESTSPECIFICATION"
	case TranslatorTagsAccessRequestSpecification:
		ret = "ACCESSREQUESTSPECIFICATION"
	case TranslatorTagsAccessRequestListOfData:
		ret = "ACCESSREQUESTLISTOFDATA"
	case TranslatorTagsAccessResponseBody:
		ret = "ACCESSRESPONSEBODY"
	case TranslatorTagsListOfAccessResponseSpecification:
		ret = "LISTOFACCESSRESPONSESPECIFICATION"
	case TranslatorTagsAccessResponseSpecification:
		ret = "ACCESSRESPONSESPECIFICATION"
	case TranslatorTagsAccessResponseListOfData:
		ret = "ACCESSRESPONSELISTOFDATA"
	case TranslatorTagsSingleResponse:
		ret = "SINGLERESPONSE"
	case TranslatorTagsService:
		ret = "SERVICE"
	case TranslatorTagsServiceError:
		ret = "SERVICEERROR"
	case TranslatorTagsInitiateError:
		ret = "INITIATEERROR"
	case TranslatorTagsCipheredService:
		ret = "CIPHEREDSERVICE"
	case TranslatorTagsSystemTitle:
		ret = "SYSTEMTITLE"
	case TranslatorTagsDataBlock:
		ret = "DATABLOCK"
	case TranslatorTagsTransactionId:
		ret = "TRANSACTIONID"
	case TranslatorTagsOriginatorSystemTitle:
		ret = "ORIGINATORSYSTEMTITLE"
	case TranslatorTagsRecipientSystemTitle:
		ret = "RECIPIENTSYSTEMTITLE"
	case TranslatorTagsOtherInformation:
		ret = "OTHERINFORMATION"
	case TranslatorTagsKeyInfo:
		ret = "KEYINFO"
	case TranslatorTagsAgreedKey:
		ret = "AGREEDKEY"
	case TranslatorTagsKeyParameters:
		ret = "KEYPARAMETERS"
	case TranslatorTagsKeyCipheredData:
		ret = "KEYCIPHEREDDATA"
	case TranslatorTagsCipheredContent:
		ret = "CIPHEREDCONTENT"
	case TranslatorTagsAttributeValue:
		ret = "ATTRIBUTEVALUE"
	case TranslatorTagsCurrentTime:
		ret = "CURRENTTIME"
	case TranslatorTagsTime:
		ret = "TIME"
	case TranslatorTagsMaxInfoRX:
		ret = "MAXINFORX"
	case TranslatorTagsMaxInfoTX:
		ret = "MAXINFOTX"
	case TranslatorTagsWindowSizeRX:
		ret = "WINDOWSIZERX"
	case TranslatorTagsWindowSizeTX:
		ret = "WINDOWSIZETX"
	case TranslatorTagsValueList:
		ret = "VALUELIST"
	case TranslatorTagsDataAccessResult:
		ret = "DATAACCESSRESULT"
	case TranslatorTagsFrameType:
		ret = "FRAMETYPE"
	case TranslatorTagsBlockControl:
		ret = "BLOCKCONTROL"
	case TranslatorTagsBlockNumberAck:
		ret = "BLOCKNUMBERACK"
	case TranslatorTagsBlockData:
		ret = "BLOCKDATA"
	case TranslatorTagsContentsDescription:
		ret = "CONTENTSDESCRIPTION"
	case TranslatorTagsArrayContents:
		ret = "ARRAYCONTENTS"
	case TranslatorTagsNetworkId:
		ret = "NETWORKID"
	case TranslatorTagsPhysicalDeviceAddress:
		ret = "PHYSICALDEVICEADDRESS"
	case TranslatorTagsProtocolVersion:
		ret = "PROTOCOLVERSION"
	case TranslatorTagsCalledAPTitle:
		ret = "CALLEDAPTITLE"
	case TranslatorTagsCalledAPInvocationId:
		ret = "CALLEDAPINVOCATIONID"
	case TranslatorTagsCalledAEInvocationId:
		ret = "CALLEDAEINVOCATIONID"
	case TranslatorTagsCallingApInvocationId:
		ret = "CALLINGAPINVOCATIONID"
	case TranslatorTagsCalledAEQualifier:
		ret = "CALLEDAEQUALIFIER"
	case TranslatorTagsResponseAllowed:
		ret = "RESPONSEALLOWED"
	case TranslatorTagsExceptionResponse:
		ret = "EXCEPTIONRESPONSE"
	case TranslatorTagsStateError:
		ret = "STATEERROR"
	case TranslatorTagsPblock:
		ret = "PBLOCK"
	case TranslatorTagsContent:
		ret = "CONTENT"
	case TranslatorTagsSignature:
		ret = "SIGNATURE"
	}
	return ret
}
