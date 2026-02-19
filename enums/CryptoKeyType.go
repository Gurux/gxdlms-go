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

// Defines asked crypto key.
type CryptoKeyType int

const (
	// CryptoKeyTypeEcdsa defines that // ECDSA public or private key is asked bit is set.
	CryptoKeyTypeEcdsa CryptoKeyType = 0x0
	// CryptoKeyTypeBlockCipher defines that // Global unicast encryption key (GUEK) is asked bit is set.
	CryptoKeyTypeBlockCipher CryptoKeyType = 0x1
	// CryptoKeyTypeAuthentication defines that // Authentication key is asked bit is set.
	CryptoKeyTypeAuthentication CryptoKeyType = 0x2
	// CryptoKeyTypeBroadcast defines that // Global broadcast encryption key (GBEK) is asked bit is set.
	CryptoKeyTypeBroadcast CryptoKeyType = 0x4
)

// CryptoKeyTypeParse converts the given string into a CryptoKeyType value.
//
// It returns the corresponding CryptoKeyType constant if the string matches
// a known level name, or an error if the input is invalid.
func CryptoKeyTypeParse(value string) (CryptoKeyType, error) {
	var ret CryptoKeyType
	var err error
	switch {
	case strings.EqualFold(value, "Ecdsa"):
		ret = CryptoKeyTypeEcdsa
	case strings.EqualFold(value, "BlockCipher"):
		ret = CryptoKeyTypeBlockCipher
	case strings.EqualFold(value, "Authentication"):
		ret = CryptoKeyTypeAuthentication
	case strings.EqualFold(value, "Broadcast"):
		ret = CryptoKeyTypeBroadcast
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the CryptoKeyType.
// It satisfies fmt.Stringer.
func (g CryptoKeyType) String() string {
	var ret string
	switch g {
	case CryptoKeyTypeEcdsa:
		ret = "Ecdsa"
	case CryptoKeyTypeBlockCipher:
		ret = "BlockCipher"
	case CryptoKeyTypeAuthentication:
		ret = "Authentication"
	case CryptoKeyTypeBroadcast:
		ret = "Broadcast"
	}
	return ret
}

// AllCryptoKeyType returns a slice containing all defined CryptoKeyType values.
func AllCryptoKeyType() []CryptoKeyType {
	return []CryptoKeyType{
	CryptoKeyTypeEcdsa,
	CryptoKeyTypeBlockCipher,
	CryptoKeyTypeAuthentication,
	CryptoKeyTypeBroadcast,
	}
}
