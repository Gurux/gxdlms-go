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

// HardwareResource describes hardware errors.
type HardwareResource int

const (
	// HardwareResourceOther defines that the other hardware resource error has occurred.
	HardwareResourceOther HardwareResource = iota
	// HardwareResourceMemoryUnavailable defines that the memory is unavailable.
	HardwareResourceMemoryUnavailable
	// HardwareResourceProcessorResourceUnavailable defines that the  processor resource is unavailable.
	HardwareResourceProcessorResourceUnavailable
	// HardwareResourceMassStorageUnavailable defines that the mass storage is unavailable.
	HardwareResourceMassStorageUnavailable
	// HardwareResourceOtherResourceUnavailable defines that the other resource is unavailable.
	HardwareResourceOtherResourceUnavailable
)

// HardwareResourceParse converts the given string into a HardwareResource value.
//
// It returns the corresponding HardwareResource constant if the string matches
// a known level name, or an error if the input is invalid.
func HardwareResourceParse(value string) (HardwareResource, error) {
	var ret HardwareResource
	var err error
	switch {
	case strings.EqualFold(value, "Other"):
		ret = HardwareResourceOther
	case strings.EqualFold(value, "MemoryUnavailable"):
		ret = HardwareResourceMemoryUnavailable
	case strings.EqualFold(value, "ProcessorResourceUnavailable"):
		ret = HardwareResourceProcessorResourceUnavailable
	case strings.EqualFold(value, "MassStorageUnavailable"):
		ret = HardwareResourceMassStorageUnavailable
	case strings.EqualFold(value, "OtherResourceUnavailable"):
		ret = HardwareResourceOtherResourceUnavailable
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the HardwareResource.
// It satisfies fmt.Stringer.
func (g HardwareResource) String() string {
	var ret string
	switch g {
	case HardwareResourceOther:
		ret = "Other"
	case HardwareResourceMemoryUnavailable:
		ret = "MemoryUnavailable"
	case HardwareResourceProcessorResourceUnavailable:
		ret = "ProcessorResourceUnavailable"
	case HardwareResourceMassStorageUnavailable:
		ret = "MassStorageUnavailable"
	case HardwareResourceOtherResourceUnavailable:
		ret = "OtherResourceUnavailable"
	}
	return ret
}

// AllHardwareResource returns a slice containing all defined HardwareResource values.
func AllHardwareResource() []HardwareResource {
	return []HardwareResource{
		HardwareResourceOther,
		HardwareResourceMemoryUnavailable,
		HardwareResourceProcessorResourceUnavailable,
		HardwareResourceMassStorageUnavailable,
		HardwareResourceOtherResourceUnavailable,
	}
}
