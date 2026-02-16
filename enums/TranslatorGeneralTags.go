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

type TranslatorGeneralTags int

const (
	TranslatorGeneralTagsApplicationContextName        TranslatorGeneralTags = 0xA1
	TranslatorGeneralTagsNegotiatedQualityOfService    TranslatorGeneralTags = 0xBE00
	TranslatorGeneralTagsProposedDlmsVersionNumber     TranslatorGeneralTags = 0xBE01
	TranslatorGeneralTagsProposedMaxPduSize            TranslatorGeneralTags = 0xBE02
	TranslatorGeneralTagsProposedConformance           TranslatorGeneralTags = 0xBE03
	TranslatorGeneralTagsVaaName                       TranslatorGeneralTags = 0xBE04
	TranslatorGeneralTagsNegotiatedConformance         TranslatorGeneralTags = 0xBE05
	TranslatorGeneralTagsNegotiatedDlmsVersionNumber   TranslatorGeneralTags = 0xBE06
	TranslatorGeneralTagsNegotiatedMaxPduSize          TranslatorGeneralTags = 0xBE07
	TranslatorGeneralTagsConformanceBit                TranslatorGeneralTags = 0xBE08
	TranslatorGeneralTagsProposedQualityOfService      TranslatorGeneralTags = 0xBE09
	TranslatorGeneralTagsSenderACSERequirements        TranslatorGeneralTags = 0x8A
	TranslatorGeneralTagsResponderACSERequirement      TranslatorGeneralTags = 0x88
	TranslatorGeneralTagsRespondingMechanismName       TranslatorGeneralTags = 0x89
	TranslatorGeneralTagsCallingMechanismName          TranslatorGeneralTags = 0x8B
	TranslatorGeneralTagsCallingAuthentication         TranslatorGeneralTags = 0xAC
	TranslatorGeneralTagsRespondingAuthentication      TranslatorGeneralTags = 0x80
	TranslatorGeneralTagsAssociationResult             TranslatorGeneralTags = 0xA2
	TranslatorGeneralTagsResultSourceDiagnostic        TranslatorGeneralTags = 0xA3
	TranslatorGeneralTagsACSEServiceUser               TranslatorGeneralTags = 0xA301
	TranslatorGeneralTagsACSEServiceProvider           TranslatorGeneralTags = 0xA302
	TranslatorGeneralTagsCallingAPTitle                TranslatorGeneralTags = 0xA6
	TranslatorGeneralTagsRespondingAPTitle             TranslatorGeneralTags = 0xA4
	TranslatorGeneralTagsDedicatedKey                  TranslatorGeneralTags = 0xA8
	TranslatorGeneralTagsCallingAeInvocationId         TranslatorGeneralTags = 0xA9
	TranslatorGeneralTagsCalledAeInvocationId          TranslatorGeneralTags = 0xA5
	TranslatorGeneralTagsCallingAeQualifier            TranslatorGeneralTags = 0xA7
	TranslatorGeneralTagsCharString                    TranslatorGeneralTags = 0xAA
	TranslatorGeneralTagsUserInformation               TranslatorGeneralTags = 0xAB
	TranslatorGeneralTagsRespondingAeInvocationId      TranslatorGeneralTags = 0xAD
	TranslatorGeneralTagsPrimeNewDeviceNotification    TranslatorGeneralTags = 0xAE
	TranslatorGeneralTagsPrimeRemoveDeviceNotification TranslatorGeneralTags = 0xAF
	TranslatorGeneralTagsPrimeStartReportingMeters     TranslatorGeneralTags = 0xB0
	TranslatorGeneralTagsPrimeDeleteMeters             TranslatorGeneralTags = 0xB1
	TranslatorGeneralTagsPrimeEnableAutoClose          TranslatorGeneralTags = 0xB2
	TranslatorGeneralTagsPrimeDisableAutoClose         TranslatorGeneralTags = 0xB3
)

// TranslatorGeneralTagsParse converts the given string into a TranslatorGeneralTags value.
//
// It returns the corresponding TranslatorGeneralTags constant if the string matches
// a known level name, or an error if the input is invalid.
func TranslatorGeneralTagsParse(value string) (TranslatorGeneralTags, error) {
	var ret TranslatorGeneralTags
	var err error
	switch strings.ToUpper(value) {
	case "APPLICATIONCONTEXTNAME":
		ret = TranslatorGeneralTagsApplicationContextName
	case "NEGOTIATEDQUALITYOFSERVICE":
		ret = TranslatorGeneralTagsNegotiatedQualityOfService
	case "PROPOSEDDLMSVERSIONNUMBER":
		ret = TranslatorGeneralTagsProposedDlmsVersionNumber
	case "PROPOSEDMAXPDUSIZE":
		ret = TranslatorGeneralTagsProposedMaxPduSize
	case "PROPOSEDCONFORMANCE":
		ret = TranslatorGeneralTagsProposedConformance
	case "VAANAME":
		ret = TranslatorGeneralTagsVaaName
	case "NEGOTIATEDCONFORMANCE":
		ret = TranslatorGeneralTagsNegotiatedConformance
	case "NEGOTIATEDDLMSVERSIONNUMBER":
		ret = TranslatorGeneralTagsNegotiatedDlmsVersionNumber
	case "NEGOTIATEDMAXPDUSIZE":
		ret = TranslatorGeneralTagsNegotiatedMaxPduSize
	case "CONFORMANCEBIT":
		ret = TranslatorGeneralTagsConformanceBit
	case "PROPOSEDQUALITYOFSERVICE":
		ret = TranslatorGeneralTagsProposedQualityOfService
	case "SENDERACSEREQUIREMENTS":
		ret = TranslatorGeneralTagsSenderACSERequirements
	case "RESPONDERACSEREQUIREMENT":
		ret = TranslatorGeneralTagsResponderACSERequirement
	case "RESPONDINGMECHANISMNAME":
		ret = TranslatorGeneralTagsRespondingMechanismName
	case "CALLINGMECHANISMNAME":
		ret = TranslatorGeneralTagsCallingMechanismName
	case "CALLINGAUTHENTICATION":
		ret = TranslatorGeneralTagsCallingAuthentication
	case "RESPONDINGAUTHENTICATION":
		ret = TranslatorGeneralTagsRespondingAuthentication
	case "ASSOCIATIONRESULT":
		ret = TranslatorGeneralTagsAssociationResult
	case "RESULTSOURCEDIAGNOSTIC":
		ret = TranslatorGeneralTagsResultSourceDiagnostic
	case "ACSESERVICEUSER":
		ret = TranslatorGeneralTagsACSEServiceUser
	case "ACSESERVICEPROVIDER":
		ret = TranslatorGeneralTagsACSEServiceProvider
	case "CALLINGAPTITLE":
		ret = TranslatorGeneralTagsCallingAPTitle
	case "RESPONDINGAPTITLE":
		ret = TranslatorGeneralTagsRespondingAPTitle
	case "DEDICATEDKEY":
		ret = TranslatorGeneralTagsDedicatedKey
	case "CALLINGAEINVOCATIONID":
		ret = TranslatorGeneralTagsCallingAeInvocationId
	case "CALLEDAEINVOCATIONID":
		ret = TranslatorGeneralTagsCalledAeInvocationId
	case "CALLINGAEQUALIFIER":
		ret = TranslatorGeneralTagsCallingAeQualifier
	case "CHARSTRING":
		ret = TranslatorGeneralTagsCharString
	case "USERINFORMATION":
		ret = TranslatorGeneralTagsUserInformation
	case "RESPONDINGAEINVOCATIONID":
		ret = TranslatorGeneralTagsRespondingAeInvocationId
	case "PRIMENEWDEVICENOTIFICATION":
		ret = TranslatorGeneralTagsPrimeNewDeviceNotification
	case "PRIMEREMOVEDEVICENOTIFICATION":
		ret = TranslatorGeneralTagsPrimeRemoveDeviceNotification
	case "PRIMESTARTREPORTINGMETERS":
		ret = TranslatorGeneralTagsPrimeStartReportingMeters
	case "PRIMEDELETEMETERS":
		ret = TranslatorGeneralTagsPrimeDeleteMeters
	case "PRIMEENABLEAUTOCLOSE":
		ret = TranslatorGeneralTagsPrimeEnableAutoClose
	case "PRIMEDISABLEAUTOCLOSE":
		ret = TranslatorGeneralTagsPrimeDisableAutoClose
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the TranslatorGeneralTags.
// It satisfies fmt.Stringer.
func (g TranslatorGeneralTags) String() string {
	var ret string
	switch g {
	case TranslatorGeneralTagsApplicationContextName:
		ret = "APPLICATIONCONTEXTNAME"
	case TranslatorGeneralTagsNegotiatedQualityOfService:
		ret = "NEGOTIATEDQUALITYOFSERVICE"
	case TranslatorGeneralTagsProposedDlmsVersionNumber:
		ret = "PROPOSEDDLMSVERSIONNUMBER"
	case TranslatorGeneralTagsProposedMaxPduSize:
		ret = "PROPOSEDMAXPDUSIZE"
	case TranslatorGeneralTagsProposedConformance:
		ret = "PROPOSEDCONFORMANCE"
	case TranslatorGeneralTagsVaaName:
		ret = "VAANAME"
	case TranslatorGeneralTagsNegotiatedConformance:
		ret = "NEGOTIATEDCONFORMANCE"
	case TranslatorGeneralTagsNegotiatedDlmsVersionNumber:
		ret = "NEGOTIATEDDLMSVERSIONNUMBER"
	case TranslatorGeneralTagsNegotiatedMaxPduSize:
		ret = "NEGOTIATEDMAXPDUSIZE"
	case TranslatorGeneralTagsConformanceBit:
		ret = "CONFORMANCEBIT"
	case TranslatorGeneralTagsProposedQualityOfService:
		ret = "PROPOSEDQUALITYOFSERVICE"
	case TranslatorGeneralTagsSenderACSERequirements:
		ret = "SENDERACSEREQUIREMENTS"
	case TranslatorGeneralTagsResponderACSERequirement:
		ret = "RESPONDERACSEREQUIREMENT"
	case TranslatorGeneralTagsRespondingMechanismName:
		ret = "RESPONDINGMECHANISMNAME"
	case TranslatorGeneralTagsCallingMechanismName:
		ret = "CALLINGMECHANISMNAME"
	case TranslatorGeneralTagsCallingAuthentication:
		ret = "CALLINGAUTHENTICATION"
	case TranslatorGeneralTagsRespondingAuthentication:
		ret = "RESPONDINGAUTHENTICATION"
	case TranslatorGeneralTagsAssociationResult:
		ret = "ASSOCIATIONRESULT"
	case TranslatorGeneralTagsResultSourceDiagnostic:
		ret = "RESULTSOURCEDIAGNOSTIC"
	case TranslatorGeneralTagsACSEServiceUser:
		ret = "ACSESERVICEUSER"
	case TranslatorGeneralTagsACSEServiceProvider:
		ret = "ACSESERVICEPROVIDER"
	case TranslatorGeneralTagsCallingAPTitle:
		ret = "CALLINGAPTITLE"
	case TranslatorGeneralTagsRespondingAPTitle:
		ret = "RESPONDINGAPTITLE"
	case TranslatorGeneralTagsDedicatedKey:
		ret = "DEDICATEDKEY"
	case TranslatorGeneralTagsCallingAeInvocationId:
		ret = "CALLINGAEINVOCATIONID"
	case TranslatorGeneralTagsCalledAeInvocationId:
		ret = "CALLEDAEINVOCATIONID"
	case TranslatorGeneralTagsCallingAeQualifier:
		ret = "CALLINGAEQUALIFIER"
	case TranslatorGeneralTagsCharString:
		ret = "CHARSTRING"
	case TranslatorGeneralTagsUserInformation:
		ret = "USERINFORMATION"
	case TranslatorGeneralTagsRespondingAeInvocationId:
		ret = "RESPONDINGAEINVOCATIONID"
	case TranslatorGeneralTagsPrimeNewDeviceNotification:
		ret = "PRIMENEWDEVICENOTIFICATION"
	case TranslatorGeneralTagsPrimeRemoveDeviceNotification:
		ret = "PRIMEREMOVEDEVICENOTIFICATION"
	case TranslatorGeneralTagsPrimeStartReportingMeters:
		ret = "PRIMESTARTREPORTINGMETERS"
	case TranslatorGeneralTagsPrimeDeleteMeters:
		ret = "PRIMEDELETEMETERS"
	case TranslatorGeneralTagsPrimeEnableAutoClose:
		ret = "PRIMEENABLEAUTOCLOSE"
	case TranslatorGeneralTagsPrimeDisableAutoClose:
		ret = "PRIMEDISABLEAUTOCLOSE"
	}
	return ret
}
