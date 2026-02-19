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

// SortMethod enumerates possible sort modes.
type SortMethod int

const (
	// SortMethodFiFo defines that the first in first out
	SortMethodFiFo SortMethod = 1
	// SortMethodLiFo defines that the last in first out.
	SortMethodLiFo SortMethod = 2
	// SortMethodLargest defines that the largest is first.
	SortMethodLargest SortMethod = 3
	// SortMethodSmallest defines that the smallest is first.
	SortMethodSmallest SortMethod = 4
	// SortMethodNearestToZero defines that the nearest to zero is first.
	SortMethodNearestToZero SortMethod = 5
	// SortMethodFarestFromZero defines that the farest from zero is first.
	SortMethodFarestFromZero SortMethod = 6
)

// SortMethodParse converts the given string into a SortMethod value.
//
// It returns the corresponding SortMethod constant if the string matches
// a known level name, or an error if the input is invalid.
func SortMethodParse(value string) (SortMethod, error) {
	var ret SortMethod
	var err error
	switch {
	case strings.EqualFold(value, "FiFo"):
		ret = SortMethodFiFo
	case strings.EqualFold(value, "LiFo"):
		ret = SortMethodLiFo
	case strings.EqualFold(value, "Largest"):
		ret = SortMethodLargest
	case strings.EqualFold(value, "Smallest"):
		ret = SortMethodSmallest
	case strings.EqualFold(value, "NearestToZero"):
		ret = SortMethodNearestToZero
	case strings.EqualFold(value, "FarestFromZero"):
		ret = SortMethodFarestFromZero
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the SortMethod.
// It satisfies fmt.Stringer.
func (g SortMethod) String() string {
	var ret string
	switch g {
	case SortMethodFiFo:
		ret = "FiFo"
	case SortMethodLiFo:
		ret = "LiFo"
	case SortMethodLargest:
		ret = "Largest"
	case SortMethodSmallest:
		ret = "Smallest"
	case SortMethodNearestToZero:
		ret = "NearestToZero"
	case SortMethodFarestFromZero:
		ret = "FarestFromZero"
	}
	return ret
}

// AllSortMethod returns a slice containing all defined SortMethod values.
func AllSortMethod() []SortMethod {
	return []SortMethod{
		SortMethodFiFo,
		SortMethodLiFo,
		SortMethodLargest,
		SortMethodSmallest,
		SortMethodNearestToZero,
		SortMethodFarestFromZero,
	}
}
