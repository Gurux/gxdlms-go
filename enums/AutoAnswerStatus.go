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

type AutoAnswerStatus int

const (
	// AutoAnswerStatusInactive defines that the the device will manage no new incoming call.
	//  This status is automatically reset to Active when the next listening window starts,
	AutoAnswerStatusInactive AutoAnswerStatus = iota
	// AutoAnswerStatusActive defines that the the device can answer to the next incoming call.
	AutoAnswerStatusActive
	// AutoAnswerStatusLocked defines that the this value can be set automatically by the device or by a specific client when this client has
	//  completed its reading session and wants to give the line back to the customer before the end of the
	//  window duration. This status is automatically reset to Active when the next listening window starts.
	AutoAnswerStatusLocked
)

// AutoAnswerStatusParse converts the given string into a AutoAnswerStatus value.
//
// It returns the corresponding AutoAnswerStatus constant if the string matches
// a known level name, or an error if the input is invalid.
func AutoAnswerStatusParse(value string) (AutoAnswerStatus, error) {
	var ret AutoAnswerStatus
	var err error
	switch strings.ToUpper(value) {
	case "INACTIVE":
		ret = AutoAnswerStatusInactive
	case "ACTIVE":
		ret = AutoAnswerStatusActive
	case "LOCKED":
		ret = AutoAnswerStatusLocked
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the AutoAnswerStatus.
// It satisfies fmt.Stringer.
func (g AutoAnswerStatus) String() string {
	var ret string
	switch g {
	case AutoAnswerStatusInactive:
		ret = "INACTIVE"
	case AutoAnswerStatusActive:
		ret = "ACTIVE"
	case AutoAnswerStatusLocked:
		ret = "LOCKED"
	}
	return ret
}
