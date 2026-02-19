package objects

// --------------------------------------------------------------------------
//
//	Gurux Ltd
//
// Filename:        $HeadURL$
//
// Version:         $Revision$,
//
//	$Date$
//	$Author$
//
// # Copyright (c) Gurux Ltd
//
// ---------------------------------------------------------------------------
//
//	DESCRIPTION
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
// ---------------------------------------------------------------------------

import (
	"github.com/Gurux/gxdlms-go/enums"
)

// CreateObject creates a COSEM object based on the object type.
func CreateObject(objectType enums.ObjectType) IGXDLMSBase {
	var ret IGXDLMSBase
	switch objectType {
	case enums.ObjectTypeActionSchedule:
		ret = &GXDLMSActionSchedule{}
	case enums.ObjectTypeActivityCalendar:
		ret = &GXDLMSActivityCalendar{}
	case enums.ObjectTypeAssociationLogicalName:
		ret = &GXDLMSAssociationLogicalName{}
	case enums.ObjectTypeAssociationShortName:
		ret = &GXDLMSAssociationShortName{}
	case enums.ObjectTypeAutoAnswer:
		ret = &GXDLMSAutoAnswer{}
	case enums.ObjectTypeAutoConnect:
		ret = &GXDLMSAutoConnect{}
	case enums.ObjectTypeClock:
		ret = &GXDLMSClock{}
	case enums.ObjectTypeData:
		ret = &GXDLMSData{}
	case enums.ObjectTypeDemandRegister:
		ret = &GXDLMSDemandRegister{}
	case enums.ObjectTypeMacAddressSetup:
		ret = &GXDLMSMacAddressSetup{}
	case enums.ObjectTypeExtendedRegister:
		ret = &GXDLMSExtendedRegister{}
	case enums.ObjectTypeGprsSetup:
		ret = &GXDLMSGprsSetup{}
	case enums.ObjectTypeIecHdlcSetup:
		ret = &GXDLMSHdlcSetup{}
	case enums.ObjectTypeIecLocalPortSetup:
		ret = &GXDLMSIECLocalPortSetup{}
	case enums.ObjectTypeIecTwistedPairSetup:
		ret = &GXDLMSIecTwistedPairSetup{}
	case enums.ObjectTypeIP4Setup:
		//TODO: ret = &GXDLMSIp4Setup{}
	case enums.ObjectTypeIP6Setup:
		//TODO: ret = &GXDLMSIp6Setup{}
	case enums.ObjectTypeMBusSlavePortSetup:
		ret = &GXDLMSMBusSlavePortSetup{}
	case enums.ObjectTypeImageTransfer:
		//TODO: ret = &GXDLMSImageTransfer{}
	case enums.ObjectTypeSecuritySetup:
		ret = &GXDLMSSecuritySetup{}
	case enums.ObjectTypeDisconnectControl:
		ret = &GXDLMSDisconnectControl{}
	case enums.ObjectTypeLimiter:
		ret = &GXDLMSLimiter{}
	case enums.ObjectTypeMBusClient:
		ret = &GXDLMSMBusClient{}
	case enums.ObjectTypeModemConfiguration:
		ret = &GXDLMSModemConfiguration{}
	case enums.ObjectTypePppSetup:
		ret = &GXDLMSPppSetup{}
	case enums.ObjectTypeProfileGeneric:
		ret = &GXDLMSProfileGeneric{}
	case enums.ObjectTypeRegister:
		ret = &GXDLMSRegister{}
	case enums.ObjectTypeRegisterActivation:
		ret = &GXDLMSRegisterActivation{}
	case enums.ObjectTypeRegisterMonitor:
		ret = &GXDLMSRegisterMonitor{}
	case enums.ObjectTypeSapAssignment:
		ret = &GXDLMSSapAssignment{}
	case enums.ObjectTypeSchedule:
		ret = &GXDLMSSchedule{}
	case enums.ObjectTypeScriptTable:
		ret = &GXDLMSScriptTable{}
	case enums.ObjectTypeSpecialDaysTable:
		ret = &GXDLMSSpecialDaysTable{}
	case enums.ObjectTypeTCPUDPSetup:
		ret = &GXDLMSTcpUdpSetup{}
		/*
		   case enums.ObjectTypeUtilityTables:
		       ret = &GXDLMSUtilityTables{}
		   case enums.ObjectTypePushSetup:
		       ret = &GXDLMSPushSetup{}
		   case enums.ObjectTypeMBusMasterPortSetup:
		       ret = &GXDLMSMBusMasterPortSetup{}
		   case enums.ObjectTypeGsmDiagnostic:
		       ret = &GXDLMSGsmDiagnostic{}
		   case enums.ObjectTypeAccount:
		       ret = &GXDLMSAccount{}
		   case enums.ObjectTypeCredit:
		       ret = &GXDLMSCredit{}
		   case enums.ObjectTypeCharge:
		       ret = &GXDLMSCharge{}
		   case enums.ObjectTypeTokenGateway:
		       ret = &GXDLMSTokenGateway{}
		   case enums.ObjectTypeParameterMonitor:
		       ret = &GXDLMSParameterMonitor{}
		   case enums.ObjectTypeCompactData:
		       ret = &GXDLMSCompactData{}
		   case enums.ObjectTypeLLCSSCSSetup:
		       ret = &GXDLMSLlcSscsSetup{}
		   case enums.ObjectTypePRIMENBOFDMPLCPHYSICALLAYERCOUNTERS:
		       ret = &GXDLMSPrimeNbOfdmPlcPhysicalLayerCounters{}
		   case enums.ObjectTypePRIMENBOFDMPLCMACSETUP:
		       ret = &GXDLMSPrimeNbOfdmPlcMacSetup{}
		   case enums.ObjectTypePRIMENBOFDMPLCMACFUNCTIONALPARAMETERS:
		       ret = &GXDLMSPrimeNbOfdmPlcMacFunctionalParameters{}
		   case enums.ObjectTypePRIMENBOFDMPLCMACCOUNTERS:
		       ret = &GXDLMSPrimeNbOfdmPlcMacCounters{}
		   case enums.ObjectTypePRIMENBOFDMPLCMACNETWORKADMINISTRATIONDATA:
		       ret = &GXDLMSPrimeNbOfdmPlcMacNetworkAdministrationData{}
		   case enums.ObjectTypePRIMENBOFDMPLCAPPLICATIONSIDENTIFICATION:
		       ret = &GXDLMSPrimeNbOfdmPlcApplicationsIdentification{}
		   case enums.ObjectTypeIEC8802LLCTYPE1SETUP:
		       ret = &GXDLMSIec8802LlcType1Setup{}
		   case enums.ObjectTypeIEC8802LLCTYPE2SETUP:
		       ret = &GXDLMSIec8802LlcType2Setup{}
		   case enums.ObjectTypeIEC8802LLCTYPE3SETUP:
		       ret = &GXDLMSIec8802LlcType3Setup{}
		   case enums.ObjectTypeSFSKREPORTINGSYSTEMLIST:
		       ret = &GXDLMSSFSKReportingSystemList{}
		   case enums.ObjectTypeArbitrator:
		       ret = &GXDLMSArbitrator{}
		   case enums.ObjectTypeSFSKMacCounters:
		       ret = &GXDLMSSFSKMacCounters{}
		   case enums.ObjectTypeSFSKMacSynchronizationTimeouts:
		       ret = &GXDLMSSFSKMacSynchronizationTimeouts{}
		   case enums.ObjectTypeSFSKActiveInitiator:
		       ret = &GXDLMSSFSKActiveInitiator{}
		   case enums.ObjectTypeSFSKPhyMacSetup:
		       ret = &GXDLMSSFSKPhyMacSetup{}
		   case enums.ObjectTypeNtpSetup:
		       ret = &GXDLMSNtpSetup{}
		*/
	case enums.ObjectTypeCommunicationPortProtection:
		ret = &GXDLMSCommunicationPortProtection{}
		/*
		   case enums.ObjectTypeG3PlcMacLayerCounters:
		       ret = &GXDLMSG3PlcMacLayerCounters{}
		   case enums.ObjectTypeG3PPlc6LoWPan:
		       ret = &GXDLMSG3Plc6LoWPan{}
		   case enums.ObjectTypeG3PlcMacSetup:
		       ret = &GXDLMSG3PlcMacSetup{}
		   case enums.ObjectTypeArrayManager:
		       ret = &GXDLMSArrayManager{}
		   case enums.ObjectTypeLteMonitoring:
		       ret = &GXDLMSLteMonitoring{}
		   case enums.ObjectTypeFunctionControl:
		       ret = &GXDLMSFunctionControl{}
		   case enums.ObjectTypeCoapSetup:
		       ret = &GXDLMSCoAPSetup{}
		   case enums.ObjectTypeCoapDiagnostic:
		       ret = &GXDLMSCoAPDiagnostic{}
		   case enums.ObjectTypeMBusPortSetup:
		       ret = &GXDLMSMBusPortSetup{}
		   case enums.ObjectTypeMBusDiagnostic:
		       ret = &GXDLMSMBusDiagnostic{}
		   case enums.ObjectTypeIEC6205541Attributes:
		       ret = &GXDLMSIec6205541Attributes{}
		*/
	default:
	}
	if ret != nil {
		ret.Base().objectType = objectType
	}
	return ret
}
