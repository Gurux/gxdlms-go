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

type PppSetupLcpOptionType int

const (
	PppSetupLcpOptionTypeMaxRecUnit               PppSetupLcpOptionType = 1
	PppSetupLcpOptionTypeAsyncControlCharMap      PppSetupLcpOptionType = 2
	PppSetupLcpOptionTypeAuthProtocol             PppSetupLcpOptionType = 3
	PppSetupLcpOptionTypeMagicNumber              PppSetupLcpOptionType = 5
	PppSetupLcpOptionTypeProtocolFieldCompression PppSetupLcpOptionType = 7
	PppSetupLcpOptionTypeAddressAndCtrCompression PppSetupLcpOptionType = 8
	PppSetupLcpOptionTypeFCSAlternatives          PppSetupLcpOptionType = 9
	PppSetupLcpOptionTypeCallback                 PppSetupLcpOptionType = 13
)

// PppSetupLcpOptionTypeParse converts the given string into a PppSetupLcpOptionType value.
//
// It returns the corresponding PppSetupLcpOptionType constant if the string matches
// a known level name, or an error if the input is invalid.
func PppSetupLcpOptionTypeParse(value string) (PppSetupLcpOptionType, error) {
	var ret PppSetupLcpOptionType
	var err error
	switch strings.ToUpper(value) {
	case "MAXRECUNIT":
		ret = PppSetupLcpOptionTypeMaxRecUnit
	case "ASYNCCONTROLCHARMAP":
		ret = PppSetupLcpOptionTypeAsyncControlCharMap
	case "AUTHPROTOCOL":
		ret = PppSetupLcpOptionTypeAuthProtocol
	case "MAGICNUMBER":
		ret = PppSetupLcpOptionTypeMagicNumber
	case "PROTOCOLFIELDCOMPRESSION":
		ret = PppSetupLcpOptionTypeProtocolFieldCompression
	case "ADDRESSANDCTRCOMPRESSION":
		ret = PppSetupLcpOptionTypeAddressAndCtrCompression
	case "FCSALTERNATIVES":
		ret = PppSetupLcpOptionTypeFCSAlternatives
	case "CALLBACK":
		ret = PppSetupLcpOptionTypeCallback
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the PppSetupLcpOptionType.
// It satisfies fmt.Stringer.
func (g PppSetupLcpOptionType) String() string {
	var ret string
	switch g {
	case PppSetupLcpOptionTypeMaxRecUnit:
		ret = "MAXRECUNIT"
	case PppSetupLcpOptionTypeAsyncControlCharMap:
		ret = "ASYNCCONTROLCHARMAP"
	case PppSetupLcpOptionTypeAuthProtocol:
		ret = "AUTHPROTOCOL"
	case PppSetupLcpOptionTypeMagicNumber:
		ret = "MAGICNUMBER"
	case PppSetupLcpOptionTypeProtocolFieldCompression:
		ret = "PROTOCOLFIELDCOMPRESSION"
	case PppSetupLcpOptionTypeAddressAndCtrCompression:
		ret = "ADDRESSANDCTRCOMPRESSION"
	case PppSetupLcpOptionTypeFCSAlternatives:
		ret = "FCSALTERNATIVES"
	case PppSetupLcpOptionTypeCallback:
		ret = "CALLBACK"
	}
	return ret
}
