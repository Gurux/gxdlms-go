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

// ObjectType ObjectType enumerates the usable types of DLMS objects in GuruxDLMS.
type ObjectType int

const (
	// ObjectTypeNone defines that the default value, no object type is set.
	ObjectTypeNone ObjectType = iota
	// ObjectTypeActionSchedule defines that the when communicating with a meter, the application may demand periodical
	// actions. If these actions are not linked to tariffication  = ActivityCalendar
	// or Schedule, use an object of type ActionSchedule  = 0x16.
	ObjectTypeActionSchedule = 22
	// ObjectTypeActivityCalendar defines that the when handling tariffication structures, you can use an object of type
	// ActivityCalendar. It determines, when to activate specified scripts to
	// perform certain activities in the meter. The activities, simply said,
	// scheduled actions, are operations that are carried out on a specified day,
	// at a specified time.
	// ActivityCalendar can be used together with a more general object type,
	// Schedule, and they can even overlap. If multiple actions are timed to the
	// same moment, the actions determined in the Schedule are executed first,
	// and then the ones determined in the ActivityCalendar. If using object
	// type SpecialDaysTable, with ActivityCalendar, simultaneous actions determined
	// in SpecialDaysTable are executed over the ones determined in ActivityCalendar.
	// <p /><b>Note: </b>To make sure that tariffication is correct after a
	// power failure, only the latest missed action from ActivityCalendar is
	// executed, with a delay. In a case of power failure, if a Schedule object
	// coexists, the latest missed action from ActivityCalendar has to be executed
	// at the correct time, sequentially with actions determined in the Schedule.
	ObjectTypeActivityCalendar = 20
	// ObjectTypeAssociationLogicalName defines that the associationLogicalName object type is used with meters that utilize
	// Logical Name associations within a COSEM.
	ObjectTypeAssociationLogicalName = 15
	// ObjectTypeAssociationShortName defines that the associationShortName object type is used with meters that utilize Short
	// Name associations within a COSEM.
	ObjectTypeAssociationShortName = 12
	// ObjectTypeAutoAnswer defines that the to determine auto answering settings  = for data transfer between device
	// and modem = s to answer incoming calls, use AutoAnswer object.
	ObjectTypeAutoAnswer = 28
	// ObjectTypeAutoConnect defines that the to determine auto connecting settings  = for data transfer from the meter
	// to defined destinations, use AutoConnect  = previously known as AutoDial
	// object.
	ObjectTypeAutoConnect = 29
	// ObjectTypeClock defines that the an object of type Clock is used to handle the information of a date  = day,
	// month and year and/or a time  = hundredths of a second, seconds, minutes
	// and hours.
	ObjectTypeClock = 8
	// ObjectTypeData defines that the an object of type Data typically stores manufacturer specific information
	// of the meter, for example configuration data and logical name.
	ObjectTypeData = 1
	// ObjectTypeDemandRegister defines that the an object of type DemandRegister stores a value, information of the item,
	// which the value belongs to, the status of the item, and the time of the value.
	// DemandRegister object type enables both current, and last average, it
	// supports both block, and sliding demand calculation, and it also provides
	// resetting the value average, and periodic averages.
	ObjectTypeDemandRegister = 5
	// ObjectTypeMacAddressSetup defines that the mAC address of the physical device.
	ObjectTypeMacAddressSetup = 43
	// ObjectTypeExtendedRegister defines that the extendedRegister stores a value, and understands the type of the value.
	// Refer to an object of this type by its logical name, using the OBIS
	// identification code.
	ObjectTypeExtendedRegister = 4
	// ObjectTypeGprsSetup defines that the to determine the GPRS settings, use GprsSetup object.
	ObjectTypeGprsSetup = 45
	// ObjectTypeIecHdlcSetup defines that the to determine the HDLC = High-level Data Link Control settings, use the
	// IecHdlcSetup object.
	ObjectTypeIecHdlcSetup = 23
	// ObjectTypeIecLocalPortSetup defines that the to determine the Local Port settings, use the IecLocalPortSetup object.
	ObjectTypeIecLocalPortSetup = 19
	// ObjectTypeIecTwistedPairSetup defines that the to determine the Twisted Pair settings, use the IecTwistedPairSetup object.
	ObjectTypeIecTwistedPairSetup = 24
	// ObjectTypeIP4Setup defines that the to determine the IP 4 settings, use the IP4Setup object.
	ObjectTypeIP4Setup = 42
	// ObjectTypeGSMDiagnostic defines that the gSM diagnostic settings.
	ObjectTypeGSMDiagnostic = 47
	// ObjectTypeIP6Setup defines that the to determine the IP 6 settings, use the Ip6Setup object.
	ObjectTypeIP6Setup = 48
	// ObjectTypeMBusSlavePortSetup defines that the to determine the M-BUS settings, use the MbusSetup object.
	ObjectTypeMBusSlavePortSetup = 25
	// ObjectTypeModemConfiguration defines that the to determine modem settings, use ModemConfiguration object.
	ObjectTypeModemConfiguration = 27
	// ObjectTypePushSetup defines that to determine the push settings, use Push setup object.
	ObjectTypePushSetup = 40
	// ObjectTypePppSetup defines that the to determine PPP  = Point-to-Point Protocol settings, use the PppSetup object.
	ObjectTypePppSetup = 44
	// ObjectTypeProfileGeneric defines that the profileGeneric determines a general way of gathering values from a profile.
	// The data is retrieved either by a period of time, or by an occuring event.
	// When gathering values from a profile, you need to understand the concept
	// of the profile buffer, in which the profile data is stored. The buffer may
	// be sorted by a register, or by a clock, within the profile, or the data
	// can be just piled in it, in order: last in, first out.
	// You can retrieve a part of the buffer, within a certain range of values,
	// or by a range of entry numbers. You can also determine objects, whose
	// values are to be retained. To determine, what to retrieve, and what to
	// retain, you need to assign the objects to the profile. You can use static
	// assignments, as all entries in a buffer are alike  = same size, same structure.
	// <p /><b>Note: </b>When you modify any assignment, the buffer of the
	// corresponding profile is cleared, and all other profiles, using the
	// modified one, will be cleared too. This is to make sure that their
	// entries stay alike by size and structure.
	ObjectTypeProfileGeneric = 7
	// ObjectTypeRegister defines that the register stores a value, and understands the type of the value. Refer to
	// an object of this type by its logical name, using the OBIS identification
	// code.
	ObjectTypeRegister = 3
	// ObjectTypeRegisterActivation defines that the when handling tariffication structures, you can use RegisterActivation to
	// determine, what objects to enable, when activating a certain activation mask.
	// The objects, assigned to the register, but not determined in the mask,
	// are disabled.
	// <p /><b>Note: </b>If an object is not assigned to any register, it is,
	// by default, enabled.
	ObjectTypeRegisterActivation = 6
	// ObjectTypeRegisterMonitor defines that the registerMonitor allows you to determine scripts to execute, when a register
	// value crosses a specified threshold. To use RegisterMonitor, also ScriptTable
	// needs to be instantiated in the same logical device.
	ObjectTypeRegisterMonitor = 21
	// ObjectTypeIec8802LlcType1Setup defines that the iSO/IEC 8802-2 LLC Type 1 setup.
	ObjectTypeIec8802LlcType1Setup = 57
	// ObjectTypeIec8802LlcType2Setup defines that the iSO/IEC 8802-2 LLC Type 2 setup.
	ObjectTypeIec8802LlcType2Setup = 58
	// ObjectTypeIec8802LlcType3Setup defines that the iSO/IEC 8802-2 LLC Type 3 setup.
	ObjectTypeIec8802LlcType3Setup = 59
	// ObjectTypeDisconnectControl defines that the instances of the Disconnect control IC manage an internal or external disconnect unit
	//  of the meter (e.g. electricity breaker, gas valve) in order to connect or disconnect
	//  – partly or entirely – the premises of the consumer to / from the supply.
	ObjectTypeDisconnectControl = 70
	// ObjectTypeLimiter defines limiter COSEM object.
	ObjectTypeLimiter = 71
	// ObjectTypeMBusClient defines M-Bus client COSEM object.
	ObjectTypeMBusClient = 72
	// ObjectTypeCompactData defines compact data COSEM object.
	ObjectTypeCompactData = 62
	// ObjectTypeParameterMonitor defines parameter monitor COSEM object.
	ObjectTypeParameterMonitor = 65
	// ObjectTypeWirelessModeQchannel defines that the defines the operational parameters for
	//  communication using the mode Q interfaces.
	ObjectTypeWirelessModeQchannel = 73
	// ObjectTypeMBusMasterPortSetup defines that the defines the operational parameters for communication using the
	//  EN 13757-2 interfaces if the device acts as an M-bus master.
	ObjectTypeMBusMasterPortSetup = 74
	// ObjectTypeMBusPortSetup defines that the servers hosted by M-Bus slave devices.
	ObjectTypeMBusPortSetup = 76
	// ObjectTypeMBusDiagnostic defines that the holds information related to the operation of the M-Bus network
	ObjectTypeMBusDiagnostic = 77
	// ObjectTypeLlcSscsSetup defines that the addresses that are provided by the base node during the opening of the
	//  convergence layer.
	ObjectTypeLlcSscsSetup = 80
	// ObjectTypePrimeNbOfdmPlcPhysicalLayerCounters defines that the counters related to the physical layers exchanges.
	ObjectTypePrimeNbOfdmPlcPhysicalLayerCounters = 81
	// ObjectTypePrimeNbOfdmPlcMacSetup defines that the a necessary parameters to set up and manage the PRIME NB OFDM PLC MAC layer.
	ObjectTypePrimeNbOfdmPlcMacSetup = 82
	// ObjectTypePrimeNbOfdmPlcMacFunctionalParameters defines that the functional behaviour of MAC.
	ObjectTypePrimeNbOfdmPlcMacFunctionalParameters = 83
	// ObjectTypePrimeNbOfdmPlcMacCounters defines that the statistical information on the operation of the MAC layer for management purposes.
	ObjectTypePrimeNbOfdmPlcMacCounters = 84
	// ObjectTypePrimeNbOfdmPlcMacNetworkAdministrationData defines that the parameters related to the management of the devices connected to the network.
	ObjectTypePrimeNbOfdmPlcMacNetworkAdministrationData = 85
	// ObjectTypePrimeNbOfdmPlcApplicationsIdentification defines that the identification information related to administration and maintenance of PRIME NB OFDM PLC devices.
	ObjectTypePrimeNbOfdmPlcApplicationsIdentification = 86
	// ObjectTypeRegisterTable defines that the registerTable stores identical attributes of objects, in a selected
	// collection of objects. All the objects in the collection need to be of
	// the same type. Also, the value in value groups A to D and F in their
	// logical name  = OBIS identification code needs to be identical.
	// <p />Clause 5 determines the possible values in value group E, as a table,
	// where header = the common part, and each cell = a possible E value,
	// of the OBIS code.
	ObjectTypeRegisterTable = 61
	// ObjectTypeNtpSetup defines that the nTP Setup is used for time synchronisation.
	ObjectTypeNtpSetup = 100
	// ObjectTypeZigBeeSasStartup defines that the configure a ZigBee PRO device with information necessary
	// to create or join the network.
	ObjectTypeZigBeeSasStartup = 101
	// ObjectTypeZigBeeSasJoin defines that the configure the behaviour of a ZigBee PRO device on
	// joining or loss of connection to the network.
	ObjectTypeZigBeeSasJoin = 102
	// ObjectTypeZigBeeSasApsFragmentation defines that the configure the fragmentation feature of ZigBee PRO transport layer.
	ObjectTypeZigBeeSasApsFragmentation = 103
	// ObjectTypeZigBeeNetworkControl defines ZigBee network control COSEM object.
	ObjectTypeZigBeeNetworkControl = 104
	// ObjectTypeDataProtection defines data protection COSEM object.
	ObjectTypeDataProtection = 30
	// ObjectTypeAccount defines account COSEM object.
	ObjectTypeAccount = 111
	// ObjectTypeCredit defines credit COSEM object.
	ObjectTypeCredit = 112
	// ObjectTypeCharge defines charge COSEM object.
	ObjectTypeCharge = 113
	// ObjectTypeTokenGateway defines token gateway COSEM object.
	ObjectTypeTokenGateway = 115
	// ObjectTypeIEC6205541Attributes defines IEC 6205541 attributes COSEM object.
	ObjectTypeIEC6205541Attributes = 116
	// ObjectTypeArrayManager defines that the allow managing attributes of type array of other interface objects.
	ObjectTypeArrayManager = 123
	// ObjectTypeSapAssignment defines that the sapAssigment stores information of assignment of the logical devices to
	// their Service Access Points.
	ObjectTypeSapAssignment = 17
	// ObjectTypeImageTransfer defines that the instances of the Image transfer IC model the mechanism of
	//  transferring binary files, called firmware Images to COSEM servers.
	ObjectTypeImageTransfer = 18
	// ObjectTypeSchedule defines that the to handle time and date driven actions, use Schedule, with an object of
	// type SpecialDaysTable.
	ObjectTypeSchedule = 10
	// ObjectTypeScriptTable defines that the to trigger a set of actions with an execute method, use object type
	// ScriptTable. Each table entry  = script includes a unique identifier, and
	// a set of action specifications, which either execute a method, or modify
	// the object attributes, within the logical device. The script can be
	// triggered by other objects  = within the same logical device, or from the
	// outside.
	ObjectTypeScriptTable = 9
	// ObjectTypeSMTPSetup defines that the to determine the SMTP protocol settings, use the SMTPSetup object.
	ObjectTypeSMTPSetup = 2
	// ObjectTypeSpecialDaysTable defines that the with SpecialDaysTable you can determine dates to override a preset behaviour,
	// for specific days  = data item day_id. SpecialDaysTable works together with
	// objects of Schedule, or Activity Calendar.
	ObjectTypeSpecialDaysTable = 11
	// ObjectTypeStatusMapping defines that the statusMapping object stores status words with mapping. Each bit in the
	// status word is mapped to position = s in referencing status table.
	ObjectTypeStatusMapping = 63
	// ObjectTypeSecuritySetup defines security setup COSEM object.
	ObjectTypeSecuritySetup = 64
	// ObjectTypeTCPUDPSetup defines that the to determine Internet TCP/UDP protocol settings, use the TCPUDPSetup object.
	ObjectTypeTCPUDPSetup = 41
	// ObjectTypeUtilityTables defines that the in an object of type UtilityTables each "Table"  = ANSI C12.19:1997 table data
	// is represented as an instance, and identified by its logical name.
	ObjectTypeUtilityTables = 26
	// ObjectTypeSFSKPhyMacSetUp defines that the s-FSK Phy MAC Setup
	ObjectTypeSFSKPhyMacSetUp = 50
	// ObjectTypeSFSKActiveInitiator defines that the s-FSK Active initiator.
	ObjectTypeSFSKActiveInitiator = 51
	// ObjectTypeSFSKMacSynchronizationTimeouts defines that the s-FSK MAC synchronization timeouts
	ObjectTypeSFSKMacSynchronizationTimeouts = 52
	// ObjectTypeSFSKMacCounters defines that the s-FSK MAC Counters.
	ObjectTypeSFSKMacCounters = 53
	// ObjectTypeIec61334_4_32LlcSetup defines that the iEC 61334-4-32 LLC setup
	ObjectTypeIec61334_4_32LlcSetup = 55
	// ObjectTypeSFSKReportingSystemList defines that the s-FSK Reporting system list.
	ObjectTypeSFSKReportingSystemList = 56
	// ObjectTypeArbitrator defines that the arbitrator.
	ObjectTypeArbitrator = 68
	// ObjectTypeG3PlcMacLayerCounters defines that the g3-PLC MAC layer counters
	ObjectTypeG3PlcMacLayerCounters = 90
	// ObjectTypeG3PlcMacSetup defines that the g3-PLC MAC setup.
	ObjectTypeG3PlcMacSetup = 91
	// ObjectTypeG3Plc6LoWPan defines that the g3-PLC 6LoWPAN.
	ObjectTypeG3Plc6LoWPan = 92
	// ObjectTypeFunctionControl defines that the function control.
	ObjectTypeFunctionControl = 122
	// ObjectTypeCommunicationPortProtection defines that the communication port protection.
	ObjectTypeCommunicationPortProtection = 124
	// ObjectTypeLteMonitoring defines that the lTE monitoring.
	ObjectTypeLteMonitoring = 151
	// ObjectTypeCoAPSetup defines that the coAP setup.
	ObjectTypeCoAPSetup = 152
	// ObjectTypeCoAPDiagnostic defines that the coAP diagnostic.
	ObjectTypeCoAPDiagnostic = 153
	// ObjectTypeG3PlcHybridRfMacLayerCounters defines that the g3-PLC Hybrid RF MAC layer counters.
	ObjectTypeG3PlcHybridRfMacLayerCounters = 160
	// ObjectTypeG3PlcHybridRfMacSetup defines that the g3-PLC Hybrid RF MAC setup.
	ObjectTypeG3PlcHybridRfMacSetup = 161
	// ObjectTypeG3PlcHybrid6LoWPANAdaptationLayerSetup defines that the g3-PLC Hybrid 6LoWPAN adaptation layer setup.
	ObjectTypeG3PlcHybrid6LoWPANAdaptationLayerSetup = 162
	// ObjectTypeTariffPlan defines that the tariff Plan (Piano Tariffario) is used in Italian standard UNI/TS 11291-11.
	ObjectTypeTariffPlan = 8192
)

// ObjectTypeParse converts the given string into a ObjectType value.
//
// It returns the corresponding ObjectType constant if the string matches
// a known level name, or an error if the input is invalid.
func ObjectTypeParse(value string) (ObjectType, error) {
	var ret ObjectType
	var err error
	switch {
	case strings.EqualFold(value, "None"):
		ret = ObjectTypeNone
	case strings.EqualFold(value, "ActionSchedule"):
		ret = ObjectTypeActionSchedule
	case strings.EqualFold(value, "ActivityCalendar"):
		ret = ObjectTypeActivityCalendar
	case strings.EqualFold(value, "AssociationLogicalName"):
		ret = ObjectTypeAssociationLogicalName
	case strings.EqualFold(value, "AssociationShortName"):
		ret = ObjectTypeAssociationShortName
	case strings.EqualFold(value, "AutoAnswer"):
		ret = ObjectTypeAutoAnswer
	case strings.EqualFold(value, "AutoConnect"):
		ret = ObjectTypeAutoConnect
	case strings.EqualFold(value, "Clock"):
		ret = ObjectTypeClock
	case strings.EqualFold(value, "Data"):
		ret = ObjectTypeData
	case strings.EqualFold(value, "DemandRegister"):
		ret = ObjectTypeDemandRegister
	case strings.EqualFold(value, "MacAddressSetup"):
		ret = ObjectTypeMacAddressSetup
	case strings.EqualFold(value, "ExtendedRegister"):
		ret = ObjectTypeExtendedRegister
	case strings.EqualFold(value, "GprsSetup"):
		ret = ObjectTypeGprsSetup
	case strings.EqualFold(value, "IecHdlcSetup"):
		ret = ObjectTypeIecHdlcSetup
	case strings.EqualFold(value, "IecLocalPortSetup"):
		ret = ObjectTypeIecLocalPortSetup
	case strings.EqualFold(value, "IecTwistedPairSetup"):
		ret = ObjectTypeIecTwistedPairSetup
	case strings.EqualFold(value, "IP4Setup"):
		ret = ObjectTypeIP4Setup
	case strings.EqualFold(value, "GSMDiagnostic"):
		ret = ObjectTypeGSMDiagnostic
	case strings.EqualFold(value, "IP6Setup"):
		ret = ObjectTypeIP6Setup
	case strings.EqualFold(value, "MBusSlavePortSetup"):
		ret = ObjectTypeMBusSlavePortSetup
	case strings.EqualFold(value, "ModemConfiguration"):
		ret = ObjectTypeModemConfiguration
	case strings.EqualFold(value, "PushSetup"):
		ret = ObjectTypePushSetup
	case strings.EqualFold(value, "PppSetup"):
		ret = ObjectTypePppSetup
	case strings.EqualFold(value, "ProfileGeneric"):
		ret = ObjectTypeProfileGeneric
	case strings.EqualFold(value, "Register"):
		ret = ObjectTypeRegister
	case strings.EqualFold(value, "RegisterActivation"):
		ret = ObjectTypeRegisterActivation
	case strings.EqualFold(value, "RegisterMonitor"):
		ret = ObjectTypeRegisterMonitor
	case strings.EqualFold(value, "Iec8802LlcType1Setup"):
		ret = ObjectTypeIec8802LlcType1Setup
	case strings.EqualFold(value, "Iec8802LlcType2Setup"):
		ret = ObjectTypeIec8802LlcType2Setup
	case strings.EqualFold(value, "Iec8802LlcType3Setup"):
		ret = ObjectTypeIec8802LlcType3Setup
	case strings.EqualFold(value, "DisconnectControl"):
		ret = ObjectTypeDisconnectControl
	case strings.EqualFold(value, "Limiter"):
		ret = ObjectTypeLimiter
	case strings.EqualFold(value, "MBusClient"):
		ret = ObjectTypeMBusClient
	case strings.EqualFold(value, "CompactData"):
		ret = ObjectTypeCompactData
	case strings.EqualFold(value, "ParameterMonitor"):
		ret = ObjectTypeParameterMonitor
	case strings.EqualFold(value, "WirelessModeQchannel"):
		ret = ObjectTypeWirelessModeQchannel
	case strings.EqualFold(value, "MBusMasterPortSetup"):
		ret = ObjectTypeMBusMasterPortSetup
	case strings.EqualFold(value, "MBusPortSetup"):
		ret = ObjectTypeMBusPortSetup
	case strings.EqualFold(value, "MBusDiagnostic"):
		ret = ObjectTypeMBusDiagnostic
	case strings.EqualFold(value, "LlcSscsSetup"):
		ret = ObjectTypeLlcSscsSetup
	case strings.EqualFold(value, "PrimeNbOfdmPlcPhysicalLayerCounters"):
		ret = ObjectTypePrimeNbOfdmPlcPhysicalLayerCounters
	case strings.EqualFold(value, "PrimeNbOfdmPlcMacSetup"):
		ret = ObjectTypePrimeNbOfdmPlcMacSetup
	case strings.EqualFold(value, "PrimeNbOfdmPlcMacFunctionalParameters"):
		ret = ObjectTypePrimeNbOfdmPlcMacFunctionalParameters
	case strings.EqualFold(value, "PrimeNbOfdmPlcMacCounters"):
		ret = ObjectTypePrimeNbOfdmPlcMacCounters
	case strings.EqualFold(value, "PrimeNbOfdmPlcMacNetworkAdministrationData"):
		ret = ObjectTypePrimeNbOfdmPlcMacNetworkAdministrationData
	case strings.EqualFold(value, "PrimeNbOfdmPlcApplicationsIdentification"):
		ret = ObjectTypePrimeNbOfdmPlcApplicationsIdentification
	case strings.EqualFold(value, "RegisterTable"):
		ret = ObjectTypeRegisterTable
	case strings.EqualFold(value, "NtpSetup"):
		ret = ObjectTypeNtpSetup
	case strings.EqualFold(value, "ZigBeeSasStartup"):
		ret = ObjectTypeZigBeeSasStartup
	case strings.EqualFold(value, "ZigBeeSasJoin"):
		ret = ObjectTypeZigBeeSasJoin
	case strings.EqualFold(value, "ZigBeeSasApsFragmentation"):
		ret = ObjectTypeZigBeeSasApsFragmentation
	case strings.EqualFold(value, "ZigBeeNetworkControl"):
		ret = ObjectTypeZigBeeNetworkControl
	case strings.EqualFold(value, "DataProtection"):
		ret = ObjectTypeDataProtection
	case strings.EqualFold(value, "Account"):
		ret = ObjectTypeAccount
	case strings.EqualFold(value, "Credit"):
		ret = ObjectTypeCredit
	case strings.EqualFold(value, "Charge"):
		ret = ObjectTypeCharge
	case strings.EqualFold(value, "TokenGateway"):
		ret = ObjectTypeTokenGateway
	case strings.EqualFold(value, "IEC6205541Attributes"):
		ret = ObjectTypeIEC6205541Attributes
	case strings.EqualFold(value, "ArrayManager"):
		ret = ObjectTypeArrayManager
	case strings.EqualFold(value, "SapAssignment"):
		ret = ObjectTypeSapAssignment
	case strings.EqualFold(value, "ImageTransfer"):
		ret = ObjectTypeImageTransfer
	case strings.EqualFold(value, "Schedule"):
		ret = ObjectTypeSchedule
	case strings.EqualFold(value, "ScriptTable"):
		ret = ObjectTypeScriptTable
	case strings.EqualFold(value, "SMTPSetup"):
		ret = ObjectTypeSMTPSetup
	case strings.EqualFold(value, "SpecialDaysTable"):
		ret = ObjectTypeSpecialDaysTable
	case strings.EqualFold(value, "StatusMapping"):
		ret = ObjectTypeStatusMapping
	case strings.EqualFold(value, "SecuritySetup"):
		ret = ObjectTypeSecuritySetup
	case strings.EqualFold(value, "TCPUDPSetup"):
		ret = ObjectTypeTCPUDPSetup
	case strings.EqualFold(value, "UtilityTables"):
		ret = ObjectTypeUtilityTables
	case strings.EqualFold(value, "SFSKPhyMacSetUp"):
		ret = ObjectTypeSFSKPhyMacSetUp
	case strings.EqualFold(value, "SFSKActiveInitiator"):
		ret = ObjectTypeSFSKActiveInitiator
	case strings.EqualFold(value, "SFSKMacSynchronizationTimeouts"):
		ret = ObjectTypeSFSKMacSynchronizationTimeouts
	case strings.EqualFold(value, "SFSKMacCounters"):
		ret = ObjectTypeSFSKMacCounters
	case strings.EqualFold(value, "Iec61334_4_32LlcSetup"):
		ret = ObjectTypeIec61334_4_32LlcSetup
	case strings.EqualFold(value, "SFSKReportingSystemList"):
		ret = ObjectTypeSFSKReportingSystemList
	case strings.EqualFold(value, "Arbitrator"):
		ret = ObjectTypeArbitrator
	case strings.EqualFold(value, "G3PlcMacLayerCounters"):
		ret = ObjectTypeG3PlcMacLayerCounters
	case strings.EqualFold(value, "G3PlcMacSetup"):
		ret = ObjectTypeG3PlcMacSetup
	case strings.EqualFold(value, "G3Plc6LoWPan"):
		ret = ObjectTypeG3Plc6LoWPan
	case strings.EqualFold(value, "FunctionControl"):
		ret = ObjectTypeFunctionControl
	case strings.EqualFold(value, "CommunicationPortProtection"):
		ret = ObjectTypeCommunicationPortProtection
	case strings.EqualFold(value, "LteMonitoring"):
		ret = ObjectTypeLteMonitoring
	case strings.EqualFold(value, "CoAPSetup"):
		ret = ObjectTypeCoAPSetup
	case strings.EqualFold(value, "CoAPDiagnostic"):
		ret = ObjectTypeCoAPDiagnostic
	case strings.EqualFold(value, "G3PlcHybridRfMacLayerCounters"):
		ret = ObjectTypeG3PlcHybridRfMacLayerCounters
	case strings.EqualFold(value, "G3PlcHybridRfMacSetup"):
		ret = ObjectTypeG3PlcHybridRfMacSetup
	case strings.EqualFold(value, "G3PlcHybrid6LoWPANAdaptationLayerSetup"):
		ret = ObjectTypeG3PlcHybrid6LoWPANAdaptationLayerSetup
	case strings.EqualFold(value, "TariffPlan"):
		ret = ObjectTypeTariffPlan
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the ObjectType.
// It satisfies fmt.Stringer.
func (g ObjectType) String() string {
	var ret string
	switch g {
	case ObjectTypeNone:
		ret = "None"
	case ObjectTypeActionSchedule:
		ret = "ActionSchedule"
	case ObjectTypeActivityCalendar:
		ret = "ActivityCalendar"
	case ObjectTypeAssociationLogicalName:
		ret = "AssociationLogicalName"
	case ObjectTypeAssociationShortName:
		ret = "AssociationShortName"
	case ObjectTypeAutoAnswer:
		ret = "AutoAnswer"
	case ObjectTypeAutoConnect:
		ret = "AutoConnect"
	case ObjectTypeClock:
		ret = "Clock"
	case ObjectTypeData:
		ret = "Data"
	case ObjectTypeDemandRegister:
		ret = "DemandRegister"
	case ObjectTypeMacAddressSetup:
		ret = "MacAddressSetup"
	case ObjectTypeExtendedRegister:
		ret = "ExtendedRegister"
	case ObjectTypeGprsSetup:
		ret = "GprsSetup"
	case ObjectTypeIecHdlcSetup:
		ret = "IecHdlcSetup"
	case ObjectTypeIecLocalPortSetup:
		ret = "IecLocalPortSetup"
	case ObjectTypeIecTwistedPairSetup:
		ret = "IecTwistedPairSetup"
	case ObjectTypeIP4Setup:
		ret = "IP4Setup"
	case ObjectTypeGSMDiagnostic:
		ret = "GSMDiagnostic"
	case ObjectTypeIP6Setup:
		ret = "IP6Setup"
	case ObjectTypeMBusSlavePortSetup:
		ret = "MBusSlavePortSetup"
	case ObjectTypeModemConfiguration:
		ret = "ModemConfiguration"
	case ObjectTypePushSetup:
		ret = "PushSetup"
	case ObjectTypePppSetup:
		ret = "PppSetup"
	case ObjectTypeProfileGeneric:
		ret = "ProfileGeneric"
	case ObjectTypeRegister:
		ret = "Register"
	case ObjectTypeRegisterActivation:
		ret = "RegisterActivation"
	case ObjectTypeRegisterMonitor:
		ret = "RegisterMonitor"
	case ObjectTypeIec8802LlcType1Setup:
		ret = "Iec8802LlcType1Setup"
	case ObjectTypeIec8802LlcType2Setup:
		ret = "Iec8802LlcType2Setup"
	case ObjectTypeIec8802LlcType3Setup:
		ret = "Iec8802LlcType3Setup"
	case ObjectTypeDisconnectControl:
		ret = "DisconnectControl"
	case ObjectTypeLimiter:
		ret = "Limiter"
	case ObjectTypeMBusClient:
		ret = "MBusClient"
	case ObjectTypeCompactData:
		ret = "CompactData"
	case ObjectTypeParameterMonitor:
		ret = "ParameterMonitor"
	case ObjectTypeWirelessModeQchannel:
		ret = "WirelessModeQchannel"
	case ObjectTypeMBusMasterPortSetup:
		ret = "MBusMasterPortSetup"
	case ObjectTypeMBusPortSetup:
		ret = "MBusPortSetup"
	case ObjectTypeMBusDiagnostic:
		ret = "MBusDiagnostic"
	case ObjectTypeLlcSscsSetup:
		ret = "LlcSscsSetup"
	case ObjectTypePrimeNbOfdmPlcPhysicalLayerCounters:
		ret = "PrimeNbOfdmPlcPhysicalLayerCounters"
	case ObjectTypePrimeNbOfdmPlcMacSetup:
		ret = "PrimeNbOfdmPlcMacSetup"
	case ObjectTypePrimeNbOfdmPlcMacFunctionalParameters:
		ret = "PrimeNbOfdmPlcMacFunctionalParameters"
	case ObjectTypePrimeNbOfdmPlcMacCounters:
		ret = "PrimeNbOfdmPlcMacCounters"
	case ObjectTypePrimeNbOfdmPlcMacNetworkAdministrationData:
		ret = "PrimeNbOfdmPlcMacNetworkAdministrationData"
	case ObjectTypePrimeNbOfdmPlcApplicationsIdentification:
		ret = "PrimeNbOfdmPlcApplicationsIdentification"
	case ObjectTypeRegisterTable:
		ret = "RegisterTable"
	case ObjectTypeNtpSetup:
		ret = "NtpSetup"
	case ObjectTypeZigBeeSasStartup:
		ret = "ZigBeeSasStartup"
	case ObjectTypeZigBeeSasJoin:
		ret = "ZigBeeSasJoin"
	case ObjectTypeZigBeeSasApsFragmentation:
		ret = "ZigBeeSasApsFragmentation"
	case ObjectTypeZigBeeNetworkControl:
		ret = "ZigBeeNetworkControl"
	case ObjectTypeDataProtection:
		ret = "DataProtection"
	case ObjectTypeAccount:
		ret = "Account"
	case ObjectTypeCredit:
		ret = "Credit"
	case ObjectTypeCharge:
		ret = "Charge"
	case ObjectTypeTokenGateway:
		ret = "TokenGateway"
	case ObjectTypeIEC6205541Attributes:
		ret = "IEC6205541Attributes"
	case ObjectTypeArrayManager:
		ret = "ArrayManager"
	case ObjectTypeSapAssignment:
		ret = "SapAssignment"
	case ObjectTypeImageTransfer:
		ret = "ImageTransfer"
	case ObjectTypeSchedule:
		ret = "Schedule"
	case ObjectTypeScriptTable:
		ret = "ScriptTable"
	case ObjectTypeSMTPSetup:
		ret = "SMTPSetup"
	case ObjectTypeSpecialDaysTable:
		ret = "SpecialDaysTable"
	case ObjectTypeStatusMapping:
		ret = "StatusMapping"
	case ObjectTypeSecuritySetup:
		ret = "SecuritySetup"
	case ObjectTypeTCPUDPSetup:
		ret = "TCPUDPSetup"
	case ObjectTypeUtilityTables:
		ret = "UtilityTables"
	case ObjectTypeSFSKPhyMacSetUp:
		ret = "SFSKPhyMacSetUp"
	case ObjectTypeSFSKActiveInitiator:
		ret = "SFSKActiveInitiator"
	case ObjectTypeSFSKMacSynchronizationTimeouts:
		ret = "SFSKMacSynchronizationTimeouts"
	case ObjectTypeSFSKMacCounters:
		ret = "SFSKMacCounters"
	case ObjectTypeIec61334_4_32LlcSetup:
		ret = "Iec61334_4_32LlcSetup"
	case ObjectTypeSFSKReportingSystemList:
		ret = "SFSKReportingSystemList"
	case ObjectTypeArbitrator:
		ret = "Arbitrator"
	case ObjectTypeG3PlcMacLayerCounters:
		ret = "G3PlcMacLayerCounters"
	case ObjectTypeG3PlcMacSetup:
		ret = "G3PlcMacSetup"
	case ObjectTypeG3Plc6LoWPan:
		ret = "G3Plc6LoWPan"
	case ObjectTypeFunctionControl:
		ret = "FunctionControl"
	case ObjectTypeCommunicationPortProtection:
		ret = "CommunicationPortProtection"
	case ObjectTypeLteMonitoring:
		ret = "LteMonitoring"
	case ObjectTypeCoAPSetup:
		ret = "CoAPSetup"
	case ObjectTypeCoAPDiagnostic:
		ret = "CoAPDiagnostic"
	case ObjectTypeG3PlcHybridRfMacLayerCounters:
		ret = "G3PlcHybridRfMacLayerCounters"
	case ObjectTypeG3PlcHybridRfMacSetup:
		ret = "G3PlcHybridRfMacSetup"
	case ObjectTypeG3PlcHybrid6LoWPANAdaptationLayerSetup:
		ret = "G3PlcHybrid6LoWPANAdaptationLayerSetup"
	case ObjectTypeTariffPlan:
		ret = "TariffPlan"
	}
	return ret
}

// AllObjectType returns a slice containing all defined ObjectType values.
func AllObjectType() []ObjectType {
	return []ObjectType{
		ObjectTypeNone,
		ObjectTypeActionSchedule,
		ObjectTypeActivityCalendar,
		ObjectTypeAssociationLogicalName,
		ObjectTypeAssociationShortName,
		ObjectTypeAutoAnswer,
		ObjectTypeAutoConnect,
		ObjectTypeClock,
		ObjectTypeData,
		ObjectTypeDemandRegister,
		ObjectTypeMacAddressSetup,
		ObjectTypeExtendedRegister,
		ObjectTypeGprsSetup,
		ObjectTypeIecHdlcSetup,
		ObjectTypeIecLocalPortSetup,
		ObjectTypeIecTwistedPairSetup,
		ObjectTypeIP4Setup,
		ObjectTypeGSMDiagnostic,
		ObjectTypeIP6Setup,
		ObjectTypeMBusSlavePortSetup,
		ObjectTypeModemConfiguration,
		ObjectTypePushSetup,
		ObjectTypePppSetup,
		ObjectTypeProfileGeneric,
		ObjectTypeRegister,
		ObjectTypeRegisterActivation,
		ObjectTypeRegisterMonitor,
		ObjectTypeIec8802LlcType1Setup,
		ObjectTypeIec8802LlcType2Setup,
		ObjectTypeIec8802LlcType3Setup,
		ObjectTypeDisconnectControl,
		ObjectTypeLimiter,
		ObjectTypeMBusClient,
		ObjectTypeCompactData,
		ObjectTypeParameterMonitor,
		ObjectTypeWirelessModeQchannel,
		ObjectTypeMBusMasterPortSetup,
		ObjectTypeMBusPortSetup,
		ObjectTypeMBusDiagnostic,
		ObjectTypeLlcSscsSetup,
		ObjectTypePrimeNbOfdmPlcPhysicalLayerCounters,
		ObjectTypePrimeNbOfdmPlcMacSetup,
		ObjectTypePrimeNbOfdmPlcMacFunctionalParameters,
		ObjectTypePrimeNbOfdmPlcMacCounters,
		ObjectTypePrimeNbOfdmPlcMacNetworkAdministrationData,
		ObjectTypePrimeNbOfdmPlcApplicationsIdentification,
		ObjectTypeRegisterTable,
		ObjectTypeNtpSetup,
		ObjectTypeZigBeeSasStartup,
		ObjectTypeZigBeeSasJoin,
		ObjectTypeZigBeeSasApsFragmentation,
		ObjectTypeZigBeeNetworkControl,
		ObjectTypeDataProtection,
		ObjectTypeAccount,
		ObjectTypeCredit,
		ObjectTypeCharge,
		ObjectTypeTokenGateway,
		ObjectTypeIEC6205541Attributes,
		ObjectTypeArrayManager,
		ObjectTypeSapAssignment,
		ObjectTypeImageTransfer,
		ObjectTypeSchedule,
		ObjectTypeScriptTable,
		ObjectTypeSMTPSetup,
		ObjectTypeSpecialDaysTable,
		ObjectTypeStatusMapping,
		ObjectTypeSecuritySetup,
		ObjectTypeTCPUDPSetup,
		ObjectTypeUtilityTables,
		ObjectTypeSFSKPhyMacSetUp,
		ObjectTypeSFSKActiveInitiator,
		ObjectTypeSFSKMacSynchronizationTimeouts,
		ObjectTypeSFSKMacCounters,
		ObjectTypeIec61334_4_32LlcSetup,
		ObjectTypeSFSKReportingSystemList,
		ObjectTypeArbitrator,
		ObjectTypeG3PlcMacLayerCounters,
		ObjectTypeG3PlcMacSetup,
		ObjectTypeG3Plc6LoWPan,
		ObjectTypeFunctionControl,
		ObjectTypeCommunicationPortProtection,
		ObjectTypeLteMonitoring,
		ObjectTypeCoAPSetup,
		ObjectTypeCoAPDiagnostic,
		ObjectTypeG3PlcHybridRfMacLayerCounters,
		ObjectTypeG3PlcHybridRfMacSetup,
		ObjectTypeG3PlcHybrid6LoWPANAdaptationLayerSetup,
		ObjectTypeTariffPlan,
	}
}
