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

// HDLC frame types.
type HdlcFrameType int

const (
	//HdlcFrameTypeIframe defines that the frame is I-Frame. Information frames are used to transport user data.
	HdlcFrameTypeIframe HdlcFrameType = 0x0
	//HdlcFrameTypeSframe defines that the frame is S-frame. Supervisory Frames are used for flow and error control. Rejected, RNR and RR.
	HdlcFrameTypeSframe HdlcFrameType = 0x1
	//HdlcFrameTypeUframe defines that the frame is U-frame. Unnumbered frames are used for link management. Example SNRM and UA.
	HdlcFrameTypeUframe HdlcFrameType = 0x3
)

// HdlcFrameTypeParse converts the given string into a HdlcFrameType value.
//
// It returns the corresponding HdlcFrameType constant if the string matches
// a known level name, or an error if the input is invalid.
func HdlcFrameTypeParse(value string) (HdlcFrameType, error) {
	var ret HdlcFrameType
	var err error
	switch strings.ToUpper(value) {
	case "IFRAME":
		ret = HdlcFrameTypeIframe
	case "SFRAME":
		ret = HdlcFrameTypeSframe
	case "UFRAME":
		ret = HdlcFrameTypeUframe
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the HdlcFrameType.
// It satisfies fmt.Stringer.
func (g HdlcFrameType) String() string {
	var ret string
	switch g {
	case HdlcFrameTypeIframe:
		ret = "IFRAME"
	case HdlcFrameTypeSframe:
		ret = "SFRAME"
	case HdlcFrameTypeUframe:
		ret = "UFRAME"
	}
	return ret
}
