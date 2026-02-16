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

// Enumerates all comformance bits.
// More information:
// https://www.gurux.fi/Gurux.DLMS.
type Conformance int

const (
	// ConformanceNone defines that conformance is not used.
	ConformanceNone Conformance = iota
	// ConformanceReservedZero defines that reserved zero conformance bit bit is set.
	ConformanceReservedZero Conformance = 0x1
	// ConformanceGeneralProtection defines that general protection conformance bit bit is set.
	ConformanceGeneralProtection Conformance = 0x2
	// ConformanceGeneralBlockTransfer defines that general block transfer conformance bit bit is set.
	ConformanceGeneralBlockTransfer Conformance = 0x4
	// ConformanceRead defines that read conformance bit bit is set.
	ConformanceRead Conformance = 0x8
	// ConformanceWrite defines that write conformance bit bit is set.
	ConformanceWrite Conformance = 0x10
	// ConformanceUnconfirmedWrite defines that un confirmed write conformance bit bit is set.
	ConformanceUnconfirmedWrite Conformance = 0x20
	// ConformanceDeltaValueEncoding defines that delta value encoding bit is set.
	ConformanceDeltaValueEncoding Conformance = 0x40
	// ConformanceReservedSeven defines that reserved seven conformance bit bit is set.
	ConformanceReservedSeven Conformance = 0x80
	// ConformanceAttribute0SupportedWithSet defines that attribute 0 supported with set conformance bit bit is set.
	ConformanceAttribute0SupportedWithSet Conformance = 0x100
	// ConformancePriorityMgmtSupported defines that priority mgmt supported conformance bit bit is set.
	ConformancePriorityMgmtSupported Conformance = 0x200
	// ConformanceAttribute0SupportedWithGet defines that attribute 0 supported with get conformance bit bit is set.
	ConformanceAttribute0SupportedWithGet Conformance = 0x400
	// ConformanceBlockTransferWithGetOrRead defines that block transfer with get or read conformance bit bit is set.
	ConformanceBlockTransferWithGetOrRead Conformance = 0x800
	// ConformanceBlockTransferWithSetOrWrite defines that block transfer with set or write conformance bit bit is set.
	ConformanceBlockTransferWithSetOrWrite Conformance = 0x1000
	// ConformanceBlockTransferWithAction defines that block transfer with action conformance bit bit is set.
	ConformanceBlockTransferWithAction Conformance = 0x2000
	// ConformanceMultipleReferences defines that multiple references conformance bit bit is set.
	ConformanceMultipleReferences Conformance = 0x4000
	// ConformanceInformationReport defines that information report conformance bit bit is set.
	ConformanceInformationReport Conformance = 0x8000
	// ConformanceDataNotification defines that data notification conformance bit bit is set.
	ConformanceDataNotification Conformance = 0x10000
	// ConformanceAccess defines that access conformance bit bit is set.
	ConformanceAccess Conformance = 0x20000
	// ConformanceParameterizedAccess defines that parameterized access conformance bit bit is set.
	ConformanceParameterizedAccess Conformance = 0x40000
	// ConformanceGet defines that get conformance bit bit is set.
	ConformanceGet Conformance = 0x80000
	// ConformanceSet defines that set conformance bit bit is set.
	ConformanceSet Conformance = 0x100000
	// ConformanceSelectiveAccess defines that selective access conformance bit bit is set.
	ConformanceSelectiveAccess Conformance = 0x200000
	// ConformanceEventNotification defines that event notification conformance bit bit is set.
	ConformanceEventNotification Conformance = 0x400000
	// ConformanceAction defines that action conformance bit bit is set.
	ConformanceAction Conformance = 0x800000
)

// ConformanceParse converts the given string into a Conformance value.
//
// It returns the corresponding Conformance constant if the string matches
// a known level name, or an error if the input is invalid.
func ConformanceParse(value string) (Conformance, error) {
	var ret Conformance
	arr := strings.Split(value, ",")
	for _, item := range arr {
		switch {
		case strings.EqualFold(item, "None"):
			ret = ConformanceNone
		case strings.EqualFold(item, "ReservedZero"):
			ret = ConformanceReservedZero
		case strings.EqualFold(item, "GeneralProtection"):
			ret = ConformanceGeneralProtection
		case strings.EqualFold(item, "GeneralBlockTransfer"):
			ret = ConformanceGeneralBlockTransfer
		case strings.EqualFold(item, "Read"):
			ret = ConformanceRead
		case strings.EqualFold(item, "Write"):
			ret = ConformanceWrite
		case strings.EqualFold(item, "UnconfirmedWrite"):
			ret = ConformanceUnconfirmedWrite
		case strings.EqualFold(item, "DeltaValueEncoding"):
			ret = ConformanceDeltaValueEncoding
		case strings.EqualFold(item, "ReservedSeven"):
			ret = ConformanceReservedSeven
		case strings.EqualFold(item, "Attribute0SupportedWithSet"):
			ret = ConformanceAttribute0SupportedWithSet
		case strings.EqualFold(item, "PriorityMgmtSupported"):
			ret = ConformancePriorityMgmtSupported
		case strings.EqualFold(item, "Attribute0SupportedWithGet"):
			ret = ConformanceAttribute0SupportedWithGet
		case strings.EqualFold(item, "BlockTransferWithGetOrRead"):
			ret = ConformanceBlockTransferWithGetOrRead
		case strings.EqualFold(item, "BlockTransferWithSetOrWrite"):
			ret = ConformanceBlockTransferWithSetOrWrite
		case strings.EqualFold(item, "BlockTransferWithAction"):
			ret = ConformanceBlockTransferWithAction
		case strings.EqualFold(item, "MultipleReferences"):
			ret = ConformanceMultipleReferences
		case strings.EqualFold(item, "InformationReport"):
			ret = ConformanceInformationReport
		case strings.EqualFold(item, "DataNotification"):
			ret = ConformanceDataNotification
		case strings.EqualFold(item, "Access"):
			ret = ConformanceAccess
		case strings.EqualFold(item, "ParameterizedAccess"):
			ret = ConformanceParameterizedAccess
		case strings.EqualFold(item, "Get"):
			ret = ConformanceGet
		case strings.EqualFold(item, "Set"):
			ret = ConformanceSet
		case strings.EqualFold(item, "SelectiveAccess"):
			ret = ConformanceSelectiveAccess
		case strings.EqualFold(item, "EventNotification"):
			ret = ConformanceEventNotification
		case strings.EqualFold(item, "Action"):
			ret = ConformanceAction
		default:
			return ConformanceNone, fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
		}
	}
	return ret, nil
}

// String returns the canonical name of the Conformance.
// It satisfies fmt.Stringer.
func (g Conformance) String() string {
	var ret string
	if g == ConformanceNone {
		ret = "None"
	} else {
		list := []string{}
		if g&ConformanceReservedZero != 0 {
			list = append(list, "ReservedZero")
		}
		if g&ConformanceGeneralProtection != 0 {
			list = append(list, "GeneralProtection")
		}
		if g&ConformanceGeneralBlockTransfer != 0 {
			list = append(list, "GeneralBlockTransfer")
		}
		if g&ConformanceRead != 0 {
			list = append(list, "Read")
		}
		if g&ConformanceWrite != 0 {
			list = append(list, "Write")
		}
		if g&ConformanceUnconfirmedWrite != 0 {
			list = append(list, "UnconfirmedWrite")
		}
		if g&ConformanceDeltaValueEncoding != 0 {
			list = append(list, "DeltaValueEncoding")
		}
		if g&ConformanceReservedSeven != 0 {
			list = append(list, "ReservedSeven")
		}
		if g&ConformanceAttribute0SupportedWithSet != 0 {
			list = append(list, "Attribute0SupportedWithSet")
		}
		if g&ConformancePriorityMgmtSupported != 0 {
			list = append(list, "PriorityMgmtSupported")
		}
		if g&ConformanceAttribute0SupportedWithGet != 0 {
			list = append(list, "Attribute0SupportedWithGet")
		}
		if g&ConformanceBlockTransferWithGetOrRead != 0 {
			list = append(list, "BlockTransferWithGetOrRead")
		}
		if g&ConformanceBlockTransferWithSetOrWrite != 0 {
			list = append(list, "BlockTransferWithSetOrWrite")
		}
		if g&ConformanceBlockTransferWithAction != 0 {
			list = append(list, "BlockTransferWithAction")
		}
		if g&ConformanceMultipleReferences != 0 {
			list = append(list, "MultipleReferences")
		}
		if g&ConformanceInformationReport != 0 {
			list = append(list, "InformationReport")
		}
		if g&ConformanceDataNotification != 0 {
			list = append(list, "DataNotification")
		}
		if g&ConformanceAccess != 0 {
			list = append(list, "Access")
		}
		if g&ConformanceParameterizedAccess != 0 {
			list = append(list, "ParameterizedAccess")
		}
		if g&ConformanceGet != 0 {
			list = append(list, "Get")
		}
		if g&ConformanceSet != 0 {
			list = append(list, "Set")
		}
		if g&ConformanceSelectiveAccess != 0 {
			list = append(list, "SelectiveAccess")
		}
		if g&ConformanceEventNotification != 0 {
			list = append(list, "EventNotification")
		}
		if g&ConformanceAction != 0 {
			list = append(list, "Action")
		}
		ret = strings.Join(list, ",")
	}
	return ret
}

// AllConformance returns a slice containing all defined conformance values.
func AllConformance() []Conformance {
	return []Conformance{
		ConformanceNone,
		ConformanceReservedZero,
		ConformanceGeneralProtection,
		ConformanceGeneralBlockTransfer,
		ConformanceRead,
		ConformanceWrite,
		ConformanceUnconfirmedWrite,
		ConformanceDeltaValueEncoding,
		ConformanceReservedSeven,
		ConformanceAttribute0SupportedWithSet,
		ConformancePriorityMgmtSupported,
		ConformanceAttribute0SupportedWithGet,
		ConformanceBlockTransferWithGetOrRead,
		ConformanceBlockTransferWithSetOrWrite,
		ConformanceBlockTransferWithAction,
		ConformanceMultipleReferences,
		ConformanceInformationReport,
		ConformanceDataNotification,
		ConformanceAccess,
		ConformanceParameterizedAccess,
		ConformanceGet,
		ConformanceSet,
		ConformanceSelectiveAccess,
		ConformanceEventNotification,
		ConformanceAction,
	}
}
