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

// AutoConnectMode Defines the mode controlling the auto dial functionality concerning the timing.
type AutoConnectMode int

const (
	// AutoConnectModeNoAutoConnect defines that the no auto dialling,
	AutoConnectModeNoAutoConnect AutoConnectMode = iota
	// AutoConnectModeAutoDiallingAllowedAnytime defines that the auto dialling allowed anytime,
	AutoConnectModeAutoDiallingAllowedAnytime
	// AutoConnectModeAutoDiallingAllowedCallingWindow defines that the auto dialling allowed within the validity time of the calling window.
	AutoConnectModeAutoDiallingAllowedCallingWindow
	// AutoConnectModeRegularAutoDiallingAllowedCallingWindow defines that the “regular” auto dialling allowed within the validity time
	//  of the calling window; “alarm” initiated auto dialling allowed anytime,
	AutoConnectModeRegularAutoDiallingAllowedCallingWindow
	// AutoConnectModeSmsSendingPlmn defines that the sMS sending via Public Land Mobile Network (PLMN),
	AutoConnectModeSmsSendingPlmn
	// AutoConnectModeSmsSendingPstn defines that the sMS sending via PSTN.
	AutoConnectModeSmsSendingPstn
	// AutoConnectModeEmailSending defines that the email sending.
	AutoConnectModeEmailSending
	// AutoConnectModePermanentlyConnect defines that the the device is permanently connected to the communication network.
	AutoConnectModePermanentlyConnect AutoConnectMode = 101
	// AutoConnectModeConnectWithCallingWindow defines that the the device is permanently connected to the communication network.
	//   No connection possible  outside the calling window.
	AutoConnectModeConnectWithCallingWindow AutoConnectMode = 102
	// AutoConnectModeConnectInvoked defines that the the device is permanently connected to the communication network.
	//  Connection is possible as soon as the connect method is invoked.
	AutoConnectModeConnectInvoked AutoConnectMode = 103
	// AutoConnectModeDisconnectConnectInvoked defines that the the device is usually disconnected.
	//  It connects to the  communication network as soon as the connect method is invoked
	AutoConnectModeDisconnectConnectInvoked AutoConnectMode = 104
)

// AutoConnectModeParse converts the given string into a AutoConnectMode value.
//
// It returns the corresponding AutoConnectMode constant if the string matches
// a known level name, or an error if the input is invalid.
func AutoConnectModeParse(value string) (AutoConnectMode, error) {
	var ret AutoConnectMode
	var err error
	switch strings.ToUpper(value) {
	case "NOAUTOCONNECT":
		ret = AutoConnectModeNoAutoConnect
	case "AUTODIALLINGALLOWEDANYTIME":
		ret = AutoConnectModeAutoDiallingAllowedAnytime
	case "AUTODIALLINGALLOWEDCALLINGWINDOW":
		ret = AutoConnectModeAutoDiallingAllowedCallingWindow
	case "REGULARAUTODIALLINGALLOWEDCALLINGWINDOW":
		ret = AutoConnectModeRegularAutoDiallingAllowedCallingWindow
	case "SMSSENDINGPLMN":
		ret = AutoConnectModeSmsSendingPlmn
	case "SMSSENDINGPSTN":
		ret = AutoConnectModeSmsSendingPstn
	case "EMAILSENDING":
		ret = AutoConnectModeEmailSending
	case "PERMANENTLYCONNECT":
		ret = AutoConnectModePermanentlyConnect
	case "CONNECTWITHCALLINGWINDOW":
		ret = AutoConnectModeConnectWithCallingWindow
	case "CONNECTINVOKED":
		ret = AutoConnectModeConnectInvoked
	case "DISCONNECTCONNECTINVOKED":
		ret = AutoConnectModeDisconnectConnectInvoked
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the AutoConnectMode.
// It satisfies fmt.Stringer.
func (g AutoConnectMode) String() string {
	var ret string
	switch g {
	case AutoConnectModeNoAutoConnect:
		ret = "NOAUTOCONNECT"
	case AutoConnectModeAutoDiallingAllowedAnytime:
		ret = "AUTODIALLINGALLOWEDANYTIME"
	case AutoConnectModeAutoDiallingAllowedCallingWindow:
		ret = "AUTODIALLINGALLOWEDCALLINGWINDOW"
	case AutoConnectModeRegularAutoDiallingAllowedCallingWindow:
		ret = "REGULARAUTODIALLINGALLOWEDCALLINGWINDOW"
	case AutoConnectModeSmsSendingPlmn:
		ret = "SMSSENDINGPLMN"
	case AutoConnectModeSmsSendingPstn:
		ret = "SMSSENDINGPSTN"
	case AutoConnectModeEmailSending:
		ret = "EMAILSENDING"
	case AutoConnectModePermanentlyConnect:
		ret = "PERMANENTLYCONNECT"
	case AutoConnectModeConnectWithCallingWindow:
		ret = "CONNECTWITHCALLINGWINDOW"
	case AutoConnectModeConnectInvoked:
		ret = "CONNECTINVOKED"
	case AutoConnectModeDisconnectConnectInvoked:
		ret = "DISCONNECTCONNECTINVOKED"
	}
	return ret
}
