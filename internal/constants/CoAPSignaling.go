package constants

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

// CoAP Signaling .
type CoAPSignaling int

const (
	// CoAPSignalingUnassigned defines that the // Unassigned.
	CoAPSignalingUnassigned CoAPSignaling = iota
	// CoAPSignalingCSM defines that the // CSM.
	CoAPSignalingCSM
	// CoAPSignalingPing defines that the // Ping.
	CoAPSignalingPing
	// CoAPSignalingPong defines that the // Pong.
	CoAPSignalingPong
	// CoAPSignalingRelease defines that the // Release.
	CoAPSignalingRelease
	// CoAPSignalingAbort defines that the // Forbidden.
	CoAPSignalingAbort
)

// CoAPSignalingParse converts the given string into a CoAPSignaling value.
//
// It returns the corresponding CoAPSignaling constant if the string matches
// a known level name, or an error if the input is invalid.
func CoAPSignalingParse(value string) (CoAPSignaling, error) {
	var ret CoAPSignaling
	var err error
	switch strings.ToUpper(value) {
	case "UNASSIGNED":
		ret = CoAPSignalingUnassigned
	case "CSM":
		ret = CoAPSignalingCSM
	case "PING":
		ret = CoAPSignalingPing
	case "PONG":
		ret = CoAPSignalingPong
	case "RELEASE":
		ret = CoAPSignalingRelease
	case "ABORT":
		ret = CoAPSignalingAbort
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the CoAPSignaling.
// It satisfies fmt.Stringer.
func (g CoAPSignaling) String() string {
	var ret string
	switch g {
	case CoAPSignalingUnassigned:
		ret = "UNASSIGNED"
	case CoAPSignalingCSM:
		ret = "CSM"
	case CoAPSignalingPing:
		ret = "PING"
	case CoAPSignalingPong:
		ret = "PONG"
	case CoAPSignalingRelease:
		ret = "RELEASE"
	case CoAPSignalingAbort:
		ret = "ABORT"
	}
	return ret
}
