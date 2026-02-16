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

// HardwareResource :  describes hardware errors.
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
	switch strings.ToUpper(value) {
	case "OTHER":
		ret = HardwareResourceOther
	case "MEMORYUNAVAILABLE":
		ret = HardwareResourceMemoryUnavailable
	case "PROCESSORRESOURCEUNAVAILABLE":
		ret = HardwareResourceProcessorResourceUnavailable
	case "MASSSTORAGEUNAVAILABLE":
		ret = HardwareResourceMassStorageUnavailable
	case "OTHERRESOURCEUNAVAILABLE":
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
		ret = "OTHER"
	case HardwareResourceMemoryUnavailable:
		ret = "MEMORYUNAVAILABLE"
	case HardwareResourceProcessorResourceUnavailable:
		ret = "PROCESSORRESOURCEUNAVAILABLE"
	case HardwareResourceMassStorageUnavailable:
		ret = "MASSSTORAGEUNAVAILABLE"
	case HardwareResourceOtherResourceUnavailable:
		ret = "OTHERRESOURCEUNAVAILABLE"
	}
	return ret
}
