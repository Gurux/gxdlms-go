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

// GsmPacketSwitchStatus Packet switched status of the modem.
type GsmPacketSwitchStatus int

const (
	// GsmPacketSwitchStatusInactive defines that the inactive
	GsmPacketSwitchStatusInactive GsmPacketSwitchStatus = iota
	// GsmPacketSwitchStatusGPRS defines that the gPRS
	GsmPacketSwitchStatusGPRS
	// GsmPacketSwitchStatusEDGE defines that the eDGE
	GsmPacketSwitchStatusEDGE
	// GsmPacketSwitchStatusUMTS defines that the uMTS
	GsmPacketSwitchStatusUMTS
	// GsmPacketSwitchStatusHSDPA defines that the hSDPA
	GsmPacketSwitchStatusHSDPA
	// GsmPacketSwitchStatusLTE defines that the lTE
	GsmPacketSwitchStatusLTE
	// GsmPacketSwitchStatusCDMA defines that the cDMA
	GsmPacketSwitchStatusCDMA
	// GsmPacketSwitchStatusLteCatM1 defines that the lTE Cat M1.
	GsmPacketSwitchStatusLteCatM1
	// GsmPacketSwitchStatusLteCatNb1 defines that the lTE Cat NB1.
	GsmPacketSwitchStatusLteCatNb1
	// GsmPacketSwitchStatusLteCatNb2 defines that the lTE Cat NB2.
	GsmPacketSwitchStatusLteCatNb2
)

// GsmPacketSwitchStatusParse converts the given string into a GsmPacketSwitchStatus value.
//
// It returns the corresponding GsmPacketSwitchStatus constant if the string matches
// a known level name, or an error if the input is invalid.
func GsmPacketSwitchStatusParse(value string) (GsmPacketSwitchStatus, error) {
	var ret GsmPacketSwitchStatus
	var err error
	switch {
	case strings.EqualFold(value, "Inactive"):
		ret = GsmPacketSwitchStatusInactive
	case strings.EqualFold(value, "GPRS"):
		ret = GsmPacketSwitchStatusGPRS
	case strings.EqualFold(value, "EDGE"):
		ret = GsmPacketSwitchStatusEDGE
	case strings.EqualFold(value, "UMTS"):
		ret = GsmPacketSwitchStatusUMTS
	case strings.EqualFold(value, "HSDPA"):
		ret = GsmPacketSwitchStatusHSDPA
	case strings.EqualFold(value, "LTE"):
		ret = GsmPacketSwitchStatusLTE
	case strings.EqualFold(value, "CDMA"):
		ret = GsmPacketSwitchStatusCDMA
	case strings.EqualFold(value, "LteCatM1"):
		ret = GsmPacketSwitchStatusLteCatM1
	case strings.EqualFold(value, "LteCatNb1"):
		ret = GsmPacketSwitchStatusLteCatNb1
	case strings.EqualFold(value, "LteCatNb2"):
		ret = GsmPacketSwitchStatusLteCatNb2
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the GsmPacketSwitchStatus.
// It satisfies fmt.Stringer.
func (g GsmPacketSwitchStatus) String() string {
	var ret string
	switch g {
	case GsmPacketSwitchStatusInactive:
		ret = "Inactive"
	case GsmPacketSwitchStatusGPRS:
		ret = "GPRS"
	case GsmPacketSwitchStatusEDGE:
		ret = "EDGE"
	case GsmPacketSwitchStatusUMTS:
		ret = "UMTS"
	case GsmPacketSwitchStatusHSDPA:
		ret = "HSDPA"
	case GsmPacketSwitchStatusLTE:
		ret = "LTE"
	case GsmPacketSwitchStatusCDMA:
		ret = "CDMA"
	case GsmPacketSwitchStatusLteCatM1:
		ret = "LteCatM1"
	case GsmPacketSwitchStatusLteCatNb1:
		ret = "LteCatNb1"
	case GsmPacketSwitchStatusLteCatNb2:
		ret = "LteCatNb2"
	}
	return ret
}

// AllGsmPacketSwitchStatus returns a slice containing all defined GsmPacketSwitchStatus values.
func AllGsmPacketSwitchStatus() []GsmPacketSwitchStatus {
	return []GsmPacketSwitchStatus{
		GsmPacketSwitchStatusInactive,
		GsmPacketSwitchStatusGPRS,
		GsmPacketSwitchStatusEDGE,
		GsmPacketSwitchStatusUMTS,
		GsmPacketSwitchStatusHSDPA,
		GsmPacketSwitchStatusLTE,
		GsmPacketSwitchStatusCDMA,
		GsmPacketSwitchStatusLteCatM1,
		GsmPacketSwitchStatusLteCatNb1,
		GsmPacketSwitchStatusLteCatNb2,
	}
}
