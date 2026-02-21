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

// BaudRate enumerates available baudrates.
type BaudRate int

const (
	// BaudRate300 defines that the baudrate is 300.
	BaudRate300 BaudRate = iota
	// BaudRate600 defines that the baudrate is 600.
	BaudRate600
	// BaudRate1200 defines that the baudrate is 1200.
	BaudRate1200
	// BaudRate2400 defines that the baudrate is 2400.
	BaudRate2400
	// BaudRate4800 defines that the baudrate is 4800.
	BaudRate4800
	// BaudRate9600 defines that the baudrate is 9600.
	BaudRate9600
	// BaudRate19200 defines that the baudrate is 19200.
	BaudRate19200
	// BaudRate38400 defines that the baudrate is 38400.
	BaudRate38400
	// BaudRate57600 defines that the baudrate is 57600.
	BaudRate57600
	// BaudRate115200 defines that the baudrate is 115200.
	BaudRate115200
)

// BaudRateParse converts the given string into a BaudRate value.
//
// It returns the corresponding BaudRate constant if the string matches
// a known level name, or an error if the input is invalid.
func BaudRateParse(value string) (BaudRate, error) {
	var ret BaudRate
	var err error
	switch {
	case strings.EqualFold(value, "Baudrate300"):
		ret = BaudRate300
	case strings.EqualFold(value, "Baudrate600"):
		ret = BaudRate600
	case strings.EqualFold(value, "Baudrate1200"):
		ret = BaudRate1200
	case strings.EqualFold(value, "Baudrate2400"):
		ret = BaudRate2400
	case strings.EqualFold(value, "Baudrate4800"):
		ret = BaudRate4800
	case strings.EqualFold(value, "Baudrate9600"):
		ret = BaudRate9600
	case strings.EqualFold(value, "Baudrate19200"):
		ret = BaudRate19200
	case strings.EqualFold(value, "Baudrate38400"):
		ret = BaudRate38400
	case strings.EqualFold(value, "Baudrate57600"):
		ret = BaudRate57600
	case strings.EqualFold(value, "Baudrate115200"):
		ret = BaudRate115200
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the BaudRate.
// It satisfies fmt.Stringer.
func (g BaudRate) String() string {
	var ret string
	switch g {
	case BaudRate300:
		ret = "Baudrate300"
	case BaudRate600:
		ret = "Baudrate600"
	case BaudRate1200:
		ret = "Baudrate1200"
	case BaudRate2400:
		ret = "Baudrate2400"
	case BaudRate4800:
		ret = "Baudrate4800"
	case BaudRate9600:
		ret = "Baudrate9600"
	case BaudRate19200:
		ret = "Baudrate19200"
	case BaudRate38400:
		ret = "Baudrate38400"
	case BaudRate57600:
		ret = "Baudrate57600"
	case BaudRate115200:
		ret = "Baudrate115200"
	}
	return ret
}

// AllBaudRate returns a slice containing all defined BaudRate values.
func AllBaudRate() []BaudRate {
	return []BaudRate{
		BaudRate300,
		BaudRate600,
		BaudRate1200,
		BaudRate2400,
		BaudRate4800,
		BaudRate9600,
		BaudRate19200,
		BaudRate38400,
		BaudRate57600,
		BaudRate115200,
	}
}
