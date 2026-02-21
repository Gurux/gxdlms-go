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
func CreateObject(objectType enums.ObjectType, ln string, sn int16) (IGXDLMSBase, error) {
	var ret IGXDLMSBase
	var err error
	switch objectType {
	case enums.ObjectTypeActionSchedule:
		ret, err = NewGXDLMSActionSchedule(ln, sn)
	case enums.ObjectTypeActivityCalendar:
		ret, err = NewGXDLMSActivityCalendar(ln, sn)
	case enums.ObjectTypeAssociationLogicalName:
		ret, err = NewGXDLMSAssociationLogicalName(ln, sn)
	case enums.ObjectTypeAssociationShortName:
		ret, err = NewGXDLMSAssociationShortName(ln, sn)
	case enums.ObjectTypeAutoAnswer:
		ret, err = NewGXDLMSAutoAnswer(ln, sn)
	case enums.ObjectTypeAutoConnect:
		ret, err = NewGXDLMSAutoConnect(ln, sn)
	case enums.ObjectTypeClock:
		ret, err = NewGXDLMSClock(ln, sn)
	case enums.ObjectTypeData:
		ret, err = NewGXDLMSData(ln, sn)
	case enums.ObjectTypeDemandRegister:
		ret, err = NewGXDLMSDemandRegister(ln, sn)
	case enums.ObjectTypeMacAddressSetup:
		ret, err = NewGXDLMSMacAddressSetup(ln, sn)
	case enums.ObjectTypeExtendedRegister:
		ret, err = NewGXDLMSExtendedRegister(ln, sn)
	case enums.ObjectTypeGprsSetup:
		ret, err = NewGXDLMSGprsSetup(ln, sn)
	case enums.ObjectTypeIecHdlcSetup:
		ret, err = NewGXDLMSHdlcSetup(ln, sn)
	case enums.ObjectTypeIecLocalPortSetup:
		ret, err = NewGXDLMSIECLocalPortSetup(ln, sn)
	case enums.ObjectTypeIecTwistedPairSetup:
		ret, err = NewGXDLMSIecTwistedPairSetup(ln, sn)
	case enums.ObjectTypeIP4Setup:
		//TODO: ret, err = NewGXDLMSIp4Setup(ln, sn)
	case enums.ObjectTypeIP6Setup:
		//TODO: ret, err = NewGXDLMSIp6Setup(ln, sn)
	case enums.ObjectTypeMBusSlavePortSetup:
		ret, err = NewGXDLMSMBusSlavePortSetup(ln, sn)
	case enums.ObjectTypeImageTransfer:
		//TODO: ret, err = NewGXDLMSImageTransfer(ln, sn)
	case enums.ObjectTypeSecuritySetup:
		ret, err = NewGXDLMSSecuritySetup(ln, sn)
	case enums.ObjectTypeDisconnectControl:
		ret, err = NewGXDLMSDisconnectControl(ln, sn)
	case enums.ObjectTypeLimiter:
		ret, err = NewGXDLMSLimiter(ln, sn)
	case enums.ObjectTypeMBusClient:
		ret, err = NewGXDLMSMBusClient(ln, sn)
	case enums.ObjectTypeModemConfiguration:
		ret, err = NewGXDLMSModemConfiguration(ln, sn)
	case enums.ObjectTypePppSetup:
		ret, err = NewGXDLMSPppSetup(ln, sn)
	case enums.ObjectTypeProfileGeneric:
		ret, err = NewGXDLMSProfileGeneric(ln, sn)
	case enums.ObjectTypeRegister:
		ret, err = NewGXDLMSRegister(ln, sn)
	case enums.ObjectTypeRegisterActivation:
		ret, err = NewGXDLMSRegisterActivation(ln, sn)
	case enums.ObjectTypeRegisterMonitor:
		ret, err = NewGXDLMSRegisterMonitor(ln, sn)
	case enums.ObjectTypeSapAssignment:
		ret, err = NewGXDLMSSapAssignment(ln, sn)
	case enums.ObjectTypeSchedule:
		ret, err = NewGXDLMSSchedule(ln, sn)
	case enums.ObjectTypeScriptTable:
		ret, err = NewGXDLMSScriptTable(ln, sn)
	case enums.ObjectTypeSpecialDaysTable:
		ret, err = NewGXDLMSSpecialDaysTable(ln, sn)
	case enums.ObjectTypeTCPUDPSetup:
		ret, err = NewGXDLMSTcpUdpSetup(ln, sn)
	case enums.ObjectTypeUtilityTables:
		ret, err = NewGXDLMSUtilityTables(ln, sn)
	case enums.ObjectTypePushSetup:
		ret, err = NewGXDLMSPushSetup(ln, sn)
	case enums.ObjectTypeMBusMasterPortSetup:
		ret, err = NewGXDLMSMBusMasterPortSetup(ln, sn)
	case enums.ObjectTypeGSMDiagnostic:
		ret, err = NewGXDLMSGSMDiagnostic(ln, sn)
	case enums.ObjectTypeAccount:
		ret, err = NewGXDLMSAccount(ln, sn)
		/*
		   case enums.ObjectTypeCredit:
		       ret, err = NewGXDLMSCredit(ln, sn)
		   case enums.ObjectTypeCharge:
		       ret, err = NewGXDLMSCharge(ln, sn)
		   case enums.ObjectTypeTokenGateway:
		       ret, err = NewGXDLMSTokenGateway(ln, sn)
		   case enums.ObjectTypeParameterMonitor:
		       ret, err = NewGXDLMSParameterMonitor(ln, sn)
		   case enums.ObjectTypeCompactData:
		       ret, err = NewGXDLMSCompactData(ln, sn)
		   case enums.ObjectTypeLLCSSCSSetup:
		       ret, err = NewGXDLMSLlcSscsSetup(ln, sn)
		   case enums.ObjectTypePRIMENBOFDMPLCPHYSICALLAYERCOUNTERS:
		       ret, err = NewGXDLMSPrimeNbOfdmPlcPhysicalLayerCounters(ln, sn)
		   case enums.ObjectTypePRIMENBOFDMPLCMACSETUP:
		       ret, err = NewGXDLMSPrimeNbOfdmPlcMacSetup(ln, sn)
		   case enums.ObjectTypePRIMENBOFDMPLCMACFUNCTIONALPARAMETERS:
		       ret, err = NewGXDLMSPrimeNbOfdmPlcMacFunctionalParameters(ln, sn)
		   case enums.ObjectTypePRIMENBOFDMPLCMACCOUNTERS:
		       ret, err = NewGXDLMSPrimeNbOfdmPlcMacCounters(ln, sn)
		   case enums.ObjectTypePRIMENBOFDMPLCMACNETWORKADMINISTRATIONDATA:
		       ret, err = NewGXDLMSPrimeNbOfdmPlcMacNetworkAdministrationData(ln, sn)
		   case enums.ObjectTypePRIMENBOFDMPLCAPPLICATIONSIDENTIFICATION:
		       ret, err = NewGXDLMSPrimeNbOfdmPlcApplicationsIdentification(ln, sn)
		   case enums.ObjectTypeIEC8802LLCTYPE1SETUP:
		       ret, err = NewGXDLMSIec8802LlcType1Setup(ln, sn)
		   case enums.ObjectTypeIEC8802LLCTYPE2SETUP:
		       ret, err = NewGXDLMSIec8802LlcType2Setup(ln, sn)
		   case enums.ObjectTypeIEC8802LLCTYPE3SETUP:
		       ret, err = NewGXDLMSIec8802LlcType3Setup(ln, sn)
		   case enums.ObjectTypeSFSKREPORTINGSYSTEMLIST:
		       ret, err = NewGXDLMSSFSKReportingSystemList(ln, sn)
		   case enums.ObjectTypeArbitrator:
		       ret, err = NewGXDLMSArbitrator(ln, sn)
		   case enums.ObjectTypeSFSKMacCounters:
		       ret, err = NewGXDLMSSFSKMacCounters(ln, sn)
		   case enums.ObjectTypeSFSKMacSynchronizationTimeouts:
		       ret, err = NewGXDLMSSFSKMacSynchronizationTimeouts(ln, sn)
		   case enums.ObjectTypeSFSKActiveInitiator:
		       ret, err = NewGXDLMSSFSKActiveInitiator(ln, sn)
		   case enums.ObjectTypeSFSKPhyMacSetup:
		       ret, err = NewGXDLMSSFSKPhyMacSetup(ln, sn)
		   case enums.ObjectTypeNtpSetup:
		       ret, err = NewGXDLMSNtpSetup(ln, sn)
		*/
	case enums.ObjectTypeCommunicationPortProtection:
		ret, err = NewGXDLMSCommunicationPortProtection(ln, sn)
		/*
		   case enums.ObjectTypeG3PlcMacLayerCounters:
		       ret, err = NewGXDLMSG3PlcMacLayerCounters(ln, sn)
		   case enums.ObjectTypeG3PPlc6LoWPan:
		       ret, err = NewGXDLMSG3Plc6LoWPan(ln, sn)
		   case enums.ObjectTypeG3PlcMacSetup:
		       ret, err = NewGXDLMSG3PlcMacSetup(ln, sn)
		   case enums.ObjectTypeArrayManager:
		       ret, err = NewGXDLMSArrayManager(ln, sn)
		   case enums.ObjectTypeLteMonitoring:
		       ret, err = NewGXDLMSLteMonitoring(ln, sn)
		   case enums.ObjectTypeFunctionControl:
		       ret, err = NewGXDLMSFunctionControl(ln, sn)
		   case enums.ObjectTypeCoapSetup:
		       ret, err = NewGXDLMSCoAPSetup(ln, sn)
		   case enums.ObjectTypeCoapDiagnostic:
		       ret, err = NewGXDLMSCoAPDiagnostic(ln, sn)
		   case enums.ObjectTypeMBusPortSetup:
		       ret, err = NewGXDLMSMBusPortSetup(ln, sn)
		   case enums.ObjectTypeMBusDiagnostic:
		       ret, err = NewGXDLMSMBusDiagnostic(ln, sn)
		   case enums.ObjectTypeIEC6205541Attributes:
		       ret, err = NewGXDLMSIec6205541Attributes(ln, sn)
		*/
	default:
	}
	return ret, err
}
