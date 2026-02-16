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

// Encryption modes.
type MBusEncryptionMode int

const (
	// MBusEncryptionModeNone defines that the encryption is not used.
	MBusEncryptionModeNone MBusEncryptionMode = iota
	// MBusEncryptionModeAes128 defines AES with Counter Mode (CTR) noPadding and IV.
	MBusEncryptionModeAes128
	// MBusEncryptionModeDesCbc defines that the  DES with Cipher Block Chaining Mode (CBC).
	MBusEncryptionModeDesCbc
	// MBusEncryptionModeDesCbcIv defines that the // DES with Cipher Block Chaining Mode (CBC) and Initial Vector.
	MBusEncryptionModeDesCbcIv
	// MBusEncryptionModeAesCbcIv defines that the // AES with Cipher Block Chaining Mode (CBC) and Initial Vector.
	MBusEncryptionModeAesCbcIv MBusEncryptionMode = 5
	// MBusEncryptionModeAesCbcIv0 defines that the // AES 128 with Cipher Block Chaining Mode (CBC) and dynamic key and Initial Vector with 0.
	MBusEncryptionModeAesCbcIv0 MBusEncryptionMode = 7
	// MBusEncryptionModeTls defines that the // TLS
	MBusEncryptionModeTls MBusEncryptionMode = 13
)

// MBusEncryptionModeParse converts the given string into a MBusEncryptionMode value.
//
// It returns the corresponding MBusEncryptionMode constant if the string matches
// a known level name, or an error if the input is invalid.
func MBusEncryptionModeParse(value string) (MBusEncryptionMode, error) {
	var ret MBusEncryptionMode
	var err error
	switch strings.ToUpper(value) {
	case "NONE":
		ret = MBusEncryptionModeNone
	case "AES128":
		ret = MBusEncryptionModeAes128
	case "DESCBC":
		ret = MBusEncryptionModeDesCbc
	case "DESCBCIV":
		ret = MBusEncryptionModeDesCbcIv
	case "AESCBCIV":
		ret = MBusEncryptionModeAesCbcIv
	case "AESCBCIV0":
		ret = MBusEncryptionModeAesCbcIv0
	case "TLS":
		ret = MBusEncryptionModeTls
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the MBusEncryptionMode.
// It satisfies fmt.Stringer.
func (g MBusEncryptionMode) String() string {
	var ret string
	switch g {
	case MBusEncryptionModeNone:
		ret = "NONE"
	case MBusEncryptionModeAes128:
		ret = "AES128"
	case MBusEncryptionModeDesCbc:
		ret = "DESCBC"
	case MBusEncryptionModeDesCbcIv:
		ret = "DESCBCIV"
	case MBusEncryptionModeAesCbcIv:
		ret = "AESCBCIV"
	case MBusEncryptionModeAesCbcIv0:
		ret = "AESCBCIV0"
	case MBusEncryptionModeTls:
		ret = "TLS"
	}
	return ret
}
