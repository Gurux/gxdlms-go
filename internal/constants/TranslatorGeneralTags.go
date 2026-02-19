package constants

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
	switch {
	case strings.EqualFold(value, "ApplicationContextName"):
		ret = TranslatorGeneralTagsApplicationContextName
	case strings.EqualFold(value, "NegotiatedQualityOfService"):
		ret = TranslatorGeneralTagsNegotiatedQualityOfService
	case strings.EqualFold(value, "ProposedDlmsVersionNumber"):
		ret = TranslatorGeneralTagsProposedDlmsVersionNumber
	case strings.EqualFold(value, "ProposedMaxPduSize"):
		ret = TranslatorGeneralTagsProposedMaxPduSize
	case strings.EqualFold(value, "ProposedConformance"):
		ret = TranslatorGeneralTagsProposedConformance
	case strings.EqualFold(value, "VaaName"):
		ret = TranslatorGeneralTagsVaaName
	case strings.EqualFold(value, "NegotiatedConformance"):
		ret = TranslatorGeneralTagsNegotiatedConformance
	case strings.EqualFold(value, "NegotiatedDlmsVersionNumber"):
		ret = TranslatorGeneralTagsNegotiatedDlmsVersionNumber
	case strings.EqualFold(value, "NegotiatedMaxPduSize"):
		ret = TranslatorGeneralTagsNegotiatedMaxPduSize
	case strings.EqualFold(value, "ConformanceBit"):
		ret = TranslatorGeneralTagsConformanceBit
	case strings.EqualFold(value, "ProposedQualityOfService"):
		ret = TranslatorGeneralTagsProposedQualityOfService
	case strings.EqualFold(value, "SenderACSERequirements"):
		ret = TranslatorGeneralTagsSenderACSERequirements
	case strings.EqualFold(value, "ResponderACSERequirement"):
		ret = TranslatorGeneralTagsResponderACSERequirement
	case strings.EqualFold(value, "RespondingMechanismName"):
		ret = TranslatorGeneralTagsRespondingMechanismName
	case strings.EqualFold(value, "CallingMechanismName"):
		ret = TranslatorGeneralTagsCallingMechanismName
	case strings.EqualFold(value, "CallingAuthentication"):
		ret = TranslatorGeneralTagsCallingAuthentication
	case strings.EqualFold(value, "RespondingAuthentication"):
		ret = TranslatorGeneralTagsRespondingAuthentication
	case strings.EqualFold(value, "AssociationResult"):
		ret = TranslatorGeneralTagsAssociationResult
	case strings.EqualFold(value, "ResultSourceDiagnostic"):
		ret = TranslatorGeneralTagsResultSourceDiagnostic
	case strings.EqualFold(value, "ACSEServiceUser"):
		ret = TranslatorGeneralTagsACSEServiceUser
	case strings.EqualFold(value, "ACSEServiceProvider"):
		ret = TranslatorGeneralTagsACSEServiceProvider
	case strings.EqualFold(value, "CallingAPTitle"):
		ret = TranslatorGeneralTagsCallingAPTitle
	case strings.EqualFold(value, "RespondingAPTitle"):
		ret = TranslatorGeneralTagsRespondingAPTitle
	case strings.EqualFold(value, "DedicatedKey"):
		ret = TranslatorGeneralTagsDedicatedKey
	case strings.EqualFold(value, "CallingAeInvocationId"):
		ret = TranslatorGeneralTagsCallingAeInvocationId
	case strings.EqualFold(value, "CalledAeInvocationId"):
		ret = TranslatorGeneralTagsCalledAeInvocationId
	case strings.EqualFold(value, "CallingAeQualifier"):
		ret = TranslatorGeneralTagsCallingAeQualifier
	case strings.EqualFold(value, "CharString"):
		ret = TranslatorGeneralTagsCharString
	case strings.EqualFold(value, "UserInformation"):
		ret = TranslatorGeneralTagsUserInformation
	case strings.EqualFold(value, "RespondingAeInvocationId"):
		ret = TranslatorGeneralTagsRespondingAeInvocationId
	case strings.EqualFold(value, "PrimeNewDeviceNotification"):
		ret = TranslatorGeneralTagsPrimeNewDeviceNotification
	case strings.EqualFold(value, "PrimeRemoveDeviceNotification"):
		ret = TranslatorGeneralTagsPrimeRemoveDeviceNotification
	case strings.EqualFold(value, "PrimeStartReportingMeters"):
		ret = TranslatorGeneralTagsPrimeStartReportingMeters
	case strings.EqualFold(value, "PrimeDeleteMeters"):
		ret = TranslatorGeneralTagsPrimeDeleteMeters
	case strings.EqualFold(value, "PrimeEnableAutoClose"):
		ret = TranslatorGeneralTagsPrimeEnableAutoClose
	case strings.EqualFold(value, "PrimeDisableAutoClose"):
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
		ret = "ApplicationContextName"
	case TranslatorGeneralTagsNegotiatedQualityOfService:
		ret = "NegotiatedQualityOfService"
	case TranslatorGeneralTagsProposedDlmsVersionNumber:
		ret = "ProposedDlmsVersionNumber"
	case TranslatorGeneralTagsProposedMaxPduSize:
		ret = "ProposedMaxPduSize"
	case TranslatorGeneralTagsProposedConformance:
		ret = "ProposedConformance"
	case TranslatorGeneralTagsVaaName:
		ret = "VaaName"
	case TranslatorGeneralTagsNegotiatedConformance:
		ret = "NegotiatedConformance"
	case TranslatorGeneralTagsNegotiatedDlmsVersionNumber:
		ret = "NegotiatedDlmsVersionNumber"
	case TranslatorGeneralTagsNegotiatedMaxPduSize:
		ret = "NegotiatedMaxPduSize"
	case TranslatorGeneralTagsConformanceBit:
		ret = "ConformanceBit"
	case TranslatorGeneralTagsProposedQualityOfService:
		ret = "ProposedQualityOfService"
	case TranslatorGeneralTagsSenderACSERequirements:
		ret = "SenderACSERequirements"
	case TranslatorGeneralTagsResponderACSERequirement:
		ret = "ResponderACSERequirement"
	case TranslatorGeneralTagsRespondingMechanismName:
		ret = "RespondingMechanismName"
	case TranslatorGeneralTagsCallingMechanismName:
		ret = "CallingMechanismName"
	case TranslatorGeneralTagsCallingAuthentication:
		ret = "CallingAuthentication"
	case TranslatorGeneralTagsRespondingAuthentication:
		ret = "RespondingAuthentication"
	case TranslatorGeneralTagsAssociationResult:
		ret = "AssociationResult"
	case TranslatorGeneralTagsResultSourceDiagnostic:
		ret = "ResultSourceDiagnostic"
	case TranslatorGeneralTagsACSEServiceUser:
		ret = "ACSEServiceUser"
	case TranslatorGeneralTagsACSEServiceProvider:
		ret = "ACSEServiceProvider"
	case TranslatorGeneralTagsCallingAPTitle:
		ret = "CallingAPTitle"
	case TranslatorGeneralTagsRespondingAPTitle:
		ret = "RespondingAPTitle"
	case TranslatorGeneralTagsDedicatedKey:
		ret = "DedicatedKey"
	case TranslatorGeneralTagsCallingAeInvocationId:
		ret = "CallingAeInvocationId"
	case TranslatorGeneralTagsCalledAeInvocationId:
		ret = "CalledAeInvocationId"
	case TranslatorGeneralTagsCallingAeQualifier:
		ret = "CallingAeQualifier"
	case TranslatorGeneralTagsCharString:
		ret = "CharString"
	case TranslatorGeneralTagsUserInformation:
		ret = "UserInformation"
	case TranslatorGeneralTagsRespondingAeInvocationId:
		ret = "RespondingAeInvocationId"
	case TranslatorGeneralTagsPrimeNewDeviceNotification:
		ret = "PrimeNewDeviceNotification"
	case TranslatorGeneralTagsPrimeRemoveDeviceNotification:
		ret = "PrimeRemoveDeviceNotification"
	case TranslatorGeneralTagsPrimeStartReportingMeters:
		ret = "PrimeStartReportingMeters"
	case TranslatorGeneralTagsPrimeDeleteMeters:
		ret = "PrimeDeleteMeters"
	case TranslatorGeneralTagsPrimeEnableAutoClose:
		ret = "PrimeEnableAutoClose"
	case TranslatorGeneralTagsPrimeDisableAutoClose:
		ret = "PrimeDisableAutoClose"
	}
	return ret
}

// AllTranslatorGeneralTags returns a slice containing all defined TranslatorGeneralTags values.
func AllTranslatorGeneralTags() []TranslatorGeneralTags {
	return []TranslatorGeneralTags{
		TranslatorGeneralTagsApplicationContextName,
		TranslatorGeneralTagsNegotiatedQualityOfService,
		TranslatorGeneralTagsProposedDlmsVersionNumber,
		TranslatorGeneralTagsProposedMaxPduSize,
		TranslatorGeneralTagsProposedConformance,
		TranslatorGeneralTagsVaaName,
		TranslatorGeneralTagsNegotiatedConformance,
		TranslatorGeneralTagsNegotiatedDlmsVersionNumber,
		TranslatorGeneralTagsNegotiatedMaxPduSize,
		TranslatorGeneralTagsConformanceBit,
		TranslatorGeneralTagsProposedQualityOfService,
		TranslatorGeneralTagsSenderACSERequirements,
		TranslatorGeneralTagsResponderACSERequirement,
		TranslatorGeneralTagsRespondingMechanismName,
		TranslatorGeneralTagsCallingMechanismName,
		TranslatorGeneralTagsCallingAuthentication,
		TranslatorGeneralTagsRespondingAuthentication,
		TranslatorGeneralTagsAssociationResult,
		TranslatorGeneralTagsResultSourceDiagnostic,
		TranslatorGeneralTagsACSEServiceUser,
		TranslatorGeneralTagsACSEServiceProvider,
		TranslatorGeneralTagsCallingAPTitle,
		TranslatorGeneralTagsRespondingAPTitle,
		TranslatorGeneralTagsDedicatedKey,
		TranslatorGeneralTagsCallingAeInvocationId,
		TranslatorGeneralTagsCalledAeInvocationId,
		TranslatorGeneralTagsCallingAeQualifier,
		TranslatorGeneralTagsCharString,
		TranslatorGeneralTagsUserInformation,
		TranslatorGeneralTagsRespondingAeInvocationId,
		TranslatorGeneralTagsPrimeNewDeviceNotification,
		TranslatorGeneralTagsPrimeRemoveDeviceNotification,
		TranslatorGeneralTagsPrimeStartReportingMeters,
		TranslatorGeneralTagsPrimeDeleteMeters,
		TranslatorGeneralTagsPrimeEnableAutoClose,
		TranslatorGeneralTagsPrimeDisableAutoClose,
	}
}
