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

// LocalPortResponseTime Defines the minimum time between the reception of a request
//
//	(end of request telegram) and the transmission of the response (begin of response telegram).
type LocalPortResponseTime int

const (
	// LocalPortResponseTimems20 defines that the minimium time is 20 ms.
	LocalPortResponseTimems20 LocalPortResponseTime = iota
	// LocalPortResponseTimems200 defines that the minimium time is 200 ms.
	LocalPortResponseTimems200
)

// LocalPortResponseTimeParse converts the given string into a LocalPortResponseTime value.
//
// It returns the corresponding LocalPortResponseTime constant if the string matches
// a known level name, or an error if the input is invalid.
func LocalPortResponseTimeParse(value string) (LocalPortResponseTime, error) {
	var ret LocalPortResponseTime
	var err error
	switch {
	case strings.EqualFold(value, "ms20"):
		ret = LocalPortResponseTimems20
	case strings.EqualFold(value, "ms200"):
		ret = LocalPortResponseTimems200
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the LocalPortResponseTime.
// It satisfies fmt.Stringer.
func (g LocalPortResponseTime) String() string {
	var ret string
	switch g {
	case LocalPortResponseTimems20:
		ret = "ms20"
	case LocalPortResponseTimems200:
		ret = "ms200"
	}
	return ret
}

// AllLocalPortResponseTime returns a slice containing all defined LocalPortResponseTime values.
func AllLocalPortResponseTime() []LocalPortResponseTime {
	return []LocalPortResponseTime{
		LocalPortResponseTimems20,
		LocalPortResponseTimems200,
	}
}
