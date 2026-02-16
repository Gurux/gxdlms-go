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

// BaudRate Defines the baudrates.
type BaudRate int

const (
	// BaudRateBaudrate300 defines that the baudrate is 300.
	BaudRateBaudrate300 BaudRate = iota
	// BaudRateBaudrate600 defines that the baudrate is 600.
	BaudRateBaudrate600
	// BaudRateBaudrate1200 defines that the baudrate is 1200.
	BaudRateBaudrate1200
	// BaudRateBaudrate2400 defines that the baudrate is 2400.
	BaudRateBaudrate2400
	// BaudRateBaudrate4800 defines that the baudrate is 4800.
	BaudRateBaudrate4800
	// BaudRateBaudrate9600 defines that the baudrate is 9600.
	BaudRateBaudrate9600
	// BaudRateBaudrate19200 defines that the baudrate is 19200.
	BaudRateBaudrate19200
	// BaudRateBaudrate38400 defines that the baudrate is 38400.
	BaudRateBaudrate38400
	// BaudRateBaudrate57600 defines that the baudrate is 57600.
	BaudRateBaudrate57600
	// BaudRateBaudrate115200 defines that the baudrate is 115200.
	BaudRateBaudrate115200
)

// BaudRateParse converts the given string into a BaudRate value.
//
// It returns the corresponding BaudRate constant if the string matches
// a known level name, or an error if the input is invalid.
func BaudRateParse(value string) (BaudRate, error) {
	var ret BaudRate
	var err error
	switch strings.ToUpper(value) {
	case "BAUDRATE300":
		ret = BaudRateBaudrate300
	case "BAUDRATE600":
		ret = BaudRateBaudrate600
	case "BAUDRATE1200":
		ret = BaudRateBaudrate1200
	case "BAUDRATE2400":
		ret = BaudRateBaudrate2400
	case "BAUDRATE4800":
		ret = BaudRateBaudrate4800
	case "BAUDRATE9600":
		ret = BaudRateBaudrate9600
	case "BAUDRATE19200":
		ret = BaudRateBaudrate19200
	case "BAUDRATE38400":
		ret = BaudRateBaudrate38400
	case "BAUDRATE57600":
		ret = BaudRateBaudrate57600
	case "BAUDRATE115200":
		ret = BaudRateBaudrate115200
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
	case BaudRateBaudrate300:
		ret = "BAUDRATE300"
	case BaudRateBaudrate600:
		ret = "BAUDRATE600"
	case BaudRateBaudrate1200:
		ret = "BAUDRATE1200"
	case BaudRateBaudrate2400:
		ret = "BAUDRATE2400"
	case BaudRateBaudrate4800:
		ret = "BAUDRATE4800"
	case BaudRateBaudrate9600:
		ret = "BAUDRATE9600"
	case BaudRateBaudrate19200:
		ret = "BAUDRATE19200"
	case BaudRateBaudrate38400:
		ret = "BAUDRATE38400"
	case BaudRateBaudrate57600:
		ret = "BAUDRATE57600"
	case BaudRateBaudrate115200:
		ret = "BAUDRATE115200"
	}
	return ret
}
