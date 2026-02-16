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

// MBusDeviceType M-Bus device type enumerations.
type MBusDeviceType int

const (
	// MBusDeviceTypeOther defines that the other.
	MBusDeviceTypeOther MBusDeviceType = iota
	// MBusDeviceTypeOil defines that the oil meter.
	MBusDeviceTypeOil
	// MBusDeviceTypeElectricity defines that the electricity meter.
	MBusDeviceTypeElectricity
	// MBusDeviceTypeGas defines that the gas meter.
	MBusDeviceTypeGas
	// MBusDeviceTypeHeat defines that the heat meter.
	MBusDeviceTypeHeat
	// MBusDeviceTypeSteam defines that the steam meter.
	MBusDeviceTypeSteam
	// MBusDeviceTypeHotWater defines that the hot water meter.
	MBusDeviceTypeHotWater
	// MBusDeviceTypeWater defines that the water meter.
	MBusDeviceTypeWater
	// MBusDeviceTypeHeatCostAllocator defines that the heat cost allocator meter.
	MBusDeviceTypeHeatCostAllocator
	// MBusDeviceTypeReserved defines that the reserved.
	MBusDeviceTypeReserved
	// MBusDeviceTypeGasMode2 defines that the gas mode 2 meter.
	MBusDeviceTypeGasMode2
	// MBusDeviceTypeHeatMode2 defines that the heat mode 2 meter.
	MBusDeviceTypeHeatMode2
	// MBusDeviceTypeHotWaterMode2 defines that the hot water mode 2 meter.
	MBusDeviceTypeHotWaterMode2
	// MBusDeviceTypeWaterMode2 defines that the water mode 2 meter.
	MBusDeviceTypeWaterMode2
	// MBusDeviceTypeHeatCostAllocatorMode2 defines that the heat cost allocator mode 2 meter.
	MBusDeviceTypeHeatCostAllocatorMode2
	// MBusDeviceTypeReserved2 defines that the reserver.
	MBusDeviceTypeReserved2
)

// MBusDeviceTypeParse converts the given string into a MBusDeviceType value.
//
// It returns the corresponding MBusDeviceType constant if the string matches
// a known level name, or an error if the input is invalid.
func MBusDeviceTypeParse(value string) (MBusDeviceType, error) {
	var ret MBusDeviceType
	var err error
	switch strings.ToUpper(value) {
	case "OTHER":
		ret = MBusDeviceTypeOther
	case "OIL":
		ret = MBusDeviceTypeOil
	case "ELECTRICITY":
		ret = MBusDeviceTypeElectricity
	case "GAS":
		ret = MBusDeviceTypeGas
	case "HEAT":
		ret = MBusDeviceTypeHeat
	case "STEAM":
		ret = MBusDeviceTypeSteam
	case "HOTWATER":
		ret = MBusDeviceTypeHotWater
	case "WATER":
		ret = MBusDeviceTypeWater
	case "HEATCOSTALLOCATOR":
		ret = MBusDeviceTypeHeatCostAllocator
	case "RESERVED":
		ret = MBusDeviceTypeReserved
	case "GASMODE2":
		ret = MBusDeviceTypeGasMode2
	case "HEATMODE2":
		ret = MBusDeviceTypeHeatMode2
	case "HOTWATERMODE2":
		ret = MBusDeviceTypeHotWaterMode2
	case "WATERMODE2":
		ret = MBusDeviceTypeWaterMode2
	case "HEATCOSTALLOCATORMODE2":
		ret = MBusDeviceTypeHeatCostAllocatorMode2
	case "RESERVED2":
		ret = MBusDeviceTypeReserved2
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the MBusDeviceType.
// It satisfies fmt.Stringer.
func (g MBusDeviceType) String() string {
	var ret string
	switch g {
	case MBusDeviceTypeOther:
		ret = "OTHER"
	case MBusDeviceTypeOil:
		ret = "OIL"
	case MBusDeviceTypeElectricity:
		ret = "ELECTRICITY"
	case MBusDeviceTypeGas:
		ret = "GAS"
	case MBusDeviceTypeHeat:
		ret = "HEAT"
	case MBusDeviceTypeSteam:
		ret = "STEAM"
	case MBusDeviceTypeHotWater:
		ret = "HOTWATER"
	case MBusDeviceTypeWater:
		ret = "WATER"
	case MBusDeviceTypeHeatCostAllocator:
		ret = "HEATCOSTALLOCATOR"
	case MBusDeviceTypeReserved:
		ret = "RESERVED"
	case MBusDeviceTypeGasMode2:
		ret = "GASMODE2"
	case MBusDeviceTypeHeatMode2:
		ret = "HEATMODE2"
	case MBusDeviceTypeHotWaterMode2:
		ret = "HOTWATERMODE2"
	case MBusDeviceTypeWaterMode2:
		ret = "WATERMODE2"
	case MBusDeviceTypeHeatCostAllocatorMode2:
		ret = "HEATCOSTALLOCATORMODE2"
	case MBusDeviceTypeReserved2:
		ret = "RESERVED2"
	}
	return ret
}
