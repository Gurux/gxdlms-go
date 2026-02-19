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
	switch {
	case strings.EqualFold(value, "Other"):
		ret = MBusDeviceTypeOther
	case strings.EqualFold(value, "Oil"):
		ret = MBusDeviceTypeOil
	case strings.EqualFold(value, "Electricity"):
		ret = MBusDeviceTypeElectricity
	case strings.EqualFold(value, "Gas"):
		ret = MBusDeviceTypeGas
	case strings.EqualFold(value, "Heat"):
		ret = MBusDeviceTypeHeat
	case strings.EqualFold(value, "Steam"):
		ret = MBusDeviceTypeSteam
	case strings.EqualFold(value, "HotWater"):
		ret = MBusDeviceTypeHotWater
	case strings.EqualFold(value, "Water"):
		ret = MBusDeviceTypeWater
	case strings.EqualFold(value, "HeatCostAllocator"):
		ret = MBusDeviceTypeHeatCostAllocator
	case strings.EqualFold(value, "Reserved"):
		ret = MBusDeviceTypeReserved
	case strings.EqualFold(value, "GasMode2"):
		ret = MBusDeviceTypeGasMode2
	case strings.EqualFold(value, "HeatMode2"):
		ret = MBusDeviceTypeHeatMode2
	case strings.EqualFold(value, "HotWaterMode2"):
		ret = MBusDeviceTypeHotWaterMode2
	case strings.EqualFold(value, "WaterMode2"):
		ret = MBusDeviceTypeWaterMode2
	case strings.EqualFold(value, "HeatCostAllocatorMode2"):
		ret = MBusDeviceTypeHeatCostAllocatorMode2
	case strings.EqualFold(value, "Reserved2"):
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
		ret = "Other"
	case MBusDeviceTypeOil:
		ret = "Oil"
	case MBusDeviceTypeElectricity:
		ret = "Electricity"
	case MBusDeviceTypeGas:
		ret = "Gas"
	case MBusDeviceTypeHeat:
		ret = "Heat"
	case MBusDeviceTypeSteam:
		ret = "Steam"
	case MBusDeviceTypeHotWater:
		ret = "HotWater"
	case MBusDeviceTypeWater:
		ret = "Water"
	case MBusDeviceTypeHeatCostAllocator:
		ret = "HeatCostAllocator"
	case MBusDeviceTypeReserved:
		ret = "Reserved"
	case MBusDeviceTypeGasMode2:
		ret = "GasMode2"
	case MBusDeviceTypeHeatMode2:
		ret = "HeatMode2"
	case MBusDeviceTypeHotWaterMode2:
		ret = "HotWaterMode2"
	case MBusDeviceTypeWaterMode2:
		ret = "WaterMode2"
	case MBusDeviceTypeHeatCostAllocatorMode2:
		ret = "HeatCostAllocatorMode2"
	case MBusDeviceTypeReserved2:
		ret = "Reserved2"
	}
	return ret
}

// AllMBusDeviceType returns a slice containing all defined MBusDeviceType values.
func AllMBusDeviceType() []MBusDeviceType {
	return []MBusDeviceType{
		MBusDeviceTypeOther,
		MBusDeviceTypeOil,
		MBusDeviceTypeElectricity,
		MBusDeviceTypeGas,
		MBusDeviceTypeHeat,
		MBusDeviceTypeSteam,
		MBusDeviceTypeHotWater,
		MBusDeviceTypeWater,
		MBusDeviceTypeHeatCostAllocator,
		MBusDeviceTypeReserved,
		MBusDeviceTypeGasMode2,
		MBusDeviceTypeHeatMode2,
		MBusDeviceTypeHotWaterMode2,
		MBusDeviceTypeWaterMode2,
		MBusDeviceTypeHeatCostAllocatorMode2,
		MBusDeviceTypeReserved2,
	}
}
