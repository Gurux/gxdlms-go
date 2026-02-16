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

// PrimeDcMsgType Prime Data Concentrator message type.
type PrimeDcMsgType int

const (
	// PrimeDcMsgTypeNewDeviceNotification defines that the new device notification.
	PrimeDcMsgTypeNewDeviceNotification PrimeDcMsgType = 1
	// PrimeDcMsgTypeRemoveDeviceNotification defines that the remove device notification.
	PrimeDcMsgTypeRemoveDeviceNotification PrimeDcMsgType = 2
	// PrimeDcMsgTypeStartReportingMeters defines that the start reporting meters.
	PrimeDcMsgTypeStartReportingMeters PrimeDcMsgType = 3
	// PrimeDcMsgTypeDeleteMeters defines that the delete meters.
	PrimeDcMsgTypeDeleteMeters PrimeDcMsgType = 4
	// PrimeDcMsgTypeEnableAutoClose defines that the enable auto close.
	PrimeDcMsgTypeEnableAutoClose PrimeDcMsgType = 5
	// PrimeDcMsgTypeDisableAutoClose defines that the disable auto close.
	PrimeDcMsgTypeDisableAutoClose PrimeDcMsgType = 6
)

// PrimeDcMsgTypeParse converts the given string into a PrimeDcMsgType value.
//
// It returns the corresponding PrimeDcMsgType constant if the string matches
// a known level name, or an error if the input is invalid.
func PrimeDcMsgTypeParse(value string) (PrimeDcMsgType, error) {
	var ret PrimeDcMsgType
	var err error
	switch strings.ToUpper(value) {
	case "NEWDEVICENOTIFICATION":
		ret = PrimeDcMsgTypeNewDeviceNotification
	case "REMOVEDEVICENOTIFICATION":
		ret = PrimeDcMsgTypeRemoveDeviceNotification
	case "STARTREPORTINGMETERS":
		ret = PrimeDcMsgTypeStartReportingMeters
	case "DELETEMETERS":
		ret = PrimeDcMsgTypeDeleteMeters
	case "ENABLEAUTOCLOSE":
		ret = PrimeDcMsgTypeEnableAutoClose
	case "DISABLEAUTOCLOSE":
		ret = PrimeDcMsgTypeDisableAutoClose
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the PrimeDcMsgType.
// It satisfies fmt.Stringer.
func (g PrimeDcMsgType) String() string {
	var ret string
	switch g {
	case PrimeDcMsgTypeNewDeviceNotification:
		ret = "NEWDEVICENOTIFICATION"
	case PrimeDcMsgTypeRemoveDeviceNotification:
		ret = "REMOVEDEVICENOTIFICATION"
	case PrimeDcMsgTypeStartReportingMeters:
		ret = "STARTREPORTINGMETERS"
	case PrimeDcMsgTypeDeleteMeters:
		ret = "DELETEMETERS"
	case PrimeDcMsgTypeEnableAutoClose:
		ret = "ENABLEAUTOCLOSE"
	case PrimeDcMsgTypeDisableAutoClose:
		ret = "DISABLEAUTOCLOSE"
	}
	return ret
}
