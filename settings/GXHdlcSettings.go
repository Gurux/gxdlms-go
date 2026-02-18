package settings

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
)

// HDLC settings contains commands for retrieving and setting the limits of
//
//	field length and window size, when communicating with the server.
type GXHdlcSettings struct {
	maxInfoTX uint16
	maxInfoRX uint16

	windowSizeTX uint8
	windowSizeRX uint8
}

// MaxInfoTX returns the the maximum information field length in transmit.
func (g *GXHdlcSettings) MaxInfoTX() uint16 {
	return g.maxInfoTX
}

// SetMaxInfoTX sets the the maximum information field length in transmit.
func (g *GXHdlcSettings) SetMaxInfoTX(value uint16) error {
	if value < 32 || value > 2030 {
		return fmt.Errorf("MaxInfoTX")
	}
	g.maxInfoTX = value
	return nil
}

// MaxInfoRX returns the the maximum information field length in receive.
func (g *GXHdlcSettings) MaxInfoRX() uint16 {
	return g.maxInfoRX
}

// SetMaxInfoRX sets the the maximum information field length in receive.
func (g *GXHdlcSettings) SetMaxInfoRX(value uint16) error {
	if value < 32 || value > 2030 {
		return fmt.Errorf("MaxInfoRX")
	}
	g.maxInfoRX = value
	return nil
}

// WindowSizeTX returns the the window size in transmit.
func (g *GXHdlcSettings) WindowSizeTX() uint8 {
	return g.windowSizeTX
}

// SetWindowSizeTX sets the the window size in transmit.
func (g *GXHdlcSettings) SetWindowSizeTX(value uint8) error {
	if value > 7 {
		return fmt.Errorf("WindowSizeTX")
	}
	g.windowSizeTX = value
	return nil
}

// WindowSizeRX returns the the window size in receive.
func (g *GXHdlcSettings) WindowSizeRX() uint8 {
	return g.windowSizeRX
}

// SetWindowSizeRX sets the the window size in receive.
func (g *GXHdlcSettings) SetWindowSizeRX(value uint8) error {
	if value > 7 {
		return fmt.Errorf("WindowSizeRX")
	}
	g.windowSizeRX = value
	return nil
}

func NewGXHdlcSettings() *GXHdlcSettings {
	return &GXHdlcSettings{
		maxInfoTX:    128,
		maxInfoRX:    128,
		windowSizeTX: 1,
		windowSizeRX: 1,
	}
}
