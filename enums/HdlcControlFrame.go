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

// HDLC control frame types.
type HdlcControlFrame int

const (
	// HdlcControlFrameReceiveReady defines that the receive is ready.
	HdlcControlFrameReceiveReady = iota
	// HdlcControlFrameReceiveNotReady defines that the receive is not ready.
	HdlcControlFrameReceiveNotReady
	// HdlcControlFrameReject defines that the frame is rejected.
	HdlcControlFrameReject
	// HdlcControlFrameSelectiveReject defines that the frame is selective rejected. Not all meters support this.
	HdlcControlFrameSelectiveReject
)

// HdlcControlFrameParse converts the given string into a HdlcControlFrame value.
//
// It returns the corresponding HdlcControlFrame constant if the string matches
// a known level name, or an error if the input is invalid.
func HdlcControlFrameParse(value string) (HdlcControlFrame, error) {
	var ret HdlcControlFrame
	var err error
	switch strings.ToUpper(value) {
	case "RECEIVEREADY":
		ret = HdlcControlFrameReceiveReady
	case "RECEIVENOTREADY":
		ret = HdlcControlFrameReceiveNotReady
	case "REJECT":
		ret = HdlcControlFrameReject
	case "SELECTIVEREJECT":
		ret = HdlcControlFrameSelectiveReject
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the HdlcControlFrame.
// It satisfies fmt.Stringer.
func (g HdlcControlFrame) String() string {
	var ret string
	switch g {
	case HdlcControlFrameReceiveReady:
		ret = "RECEIVEREADY"
	case HdlcControlFrameReceiveNotReady:
		ret = "RECEIVENOTREADY"
	case HdlcControlFrameReject:
		ret = "REJECT"
	case HdlcControlFrameSelectiveReject:
		ret = "SELECTIVEREJECT"
	}
	return ret
}
