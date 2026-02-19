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

// GlobalKeyType Enumerates global key types.
type GlobalKeyType int

const (
	// GlobalKeyTypeUnicastEncryption defines that the global unicast encryption key.
	GlobalKeyTypeUnicastEncryption GlobalKeyType = iota
	// GlobalKeyTypeBroadcastEncryption defines that the global broadcast encryption key.
	GlobalKeyTypeBroadcastEncryption
	// GlobalKeyTypeAuthentication defines that the authentication key.
	GlobalKeyTypeAuthentication
	// GlobalKeyTypeKek defines that the key Encrypting Key, also known as Master key.
	GlobalKeyTypeKek
)

// GlobalKeyTypeParse converts the given string into a GlobalKeyType value.
//
// It returns the corresponding GlobalKeyType constant if the string matches
// a known level name, or an error if the input is invalid.
func GlobalKeyTypeParse(value string) (GlobalKeyType, error) {
	var ret GlobalKeyType
	var err error
	switch {
	case strings.EqualFold(value, "UnicastEncryption"):
		ret = GlobalKeyTypeUnicastEncryption
	case strings.EqualFold(value, "BroadcastEncryption"):
		ret = GlobalKeyTypeBroadcastEncryption
	case strings.EqualFold(value, "Authentication"):
		ret = GlobalKeyTypeAuthentication
	case strings.EqualFold(value, "Kek"):
		ret = GlobalKeyTypeKek
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the GlobalKeyType.
// It satisfies fmt.Stringer.
func (g GlobalKeyType) String() string {
	var ret string
	switch g {
	case GlobalKeyTypeUnicastEncryption:
		ret = "UnicastEncryption"
	case GlobalKeyTypeBroadcastEncryption:
		ret = "BroadcastEncryption"
	case GlobalKeyTypeAuthentication:
		ret = "Authentication"
	case GlobalKeyTypeKek:
		ret = "Kek"
	}
	return ret
}

// AllGlobalKeyType returns a slice containing all defined GlobalKeyType values.
func AllGlobalKeyType() []GlobalKeyType {
	return []GlobalKeyType{
		GlobalKeyTypeUnicastEncryption,
		GlobalKeyTypeBroadcastEncryption,
		GlobalKeyTypeAuthentication,
		GlobalKeyTypeKek,
	}
}
