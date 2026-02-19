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

// DataProtectionIdentifiedKeyType : s.
type DataProtectionIdentifiedKeyType int

const (
	// DataProtectionIdentifiedKeyTypeUnicastEncryption defines that the global unicast encryption key.
	DataProtectionIdentifiedKeyTypeUnicastEncryption DataProtectionIdentifiedKeyType = iota
	// DataProtectionIdentifiedKeyTypeBroadcastEncryption defines that the global broadcast encryption key.
	DataProtectionIdentifiedKeyTypeBroadcastEncryption
)

// DataProtectionIdentifiedKeyTypeParse converts the given string into a DataProtectionIdentifiedKeyType value.
//
// It returns the corresponding DataProtectionIdentifiedKeyType constant if the string matches
// a known level name, or an error if the input is invalid.
func DataProtectionIdentifiedKeyTypeParse(value string) (DataProtectionIdentifiedKeyType, error) {
	var ret DataProtectionIdentifiedKeyType
	var err error
	switch {
	case strings.EqualFold(value, "UnicastEncryption"):
		ret = DataProtectionIdentifiedKeyTypeUnicastEncryption
	case strings.EqualFold(value, "BroadcastEncryption"):
		ret = DataProtectionIdentifiedKeyTypeBroadcastEncryption
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the DataProtectionIdentifiedKeyType.
// It satisfies fmt.Stringer.
func (g DataProtectionIdentifiedKeyType) String() string {
	var ret string
	switch g {
	case DataProtectionIdentifiedKeyTypeUnicastEncryption:
		ret = "UnicastEncryption"
	case DataProtectionIdentifiedKeyTypeBroadcastEncryption:
		ret = "BroadcastEncryption"
	}
	return ret
}

// AllDataProtectionIdentifiedKeyType returns a slice containing all defined DataProtectionIdentifiedKeyType values.
func AllDataProtectionIdentifiedKeyType() []DataProtectionIdentifiedKeyType {
	return []DataProtectionIdentifiedKeyType{
	DataProtectionIdentifiedKeyTypeUnicastEncryption,
	DataProtectionIdentifiedKeyTypeBroadcastEncryption,
	}
}
