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

type MBusEncryptionKeyStatus int

const (
	MBusEncryptionKeyStatusNoEncryptionKey MBusEncryptionKeyStatus = iota
	MBusEncryptionKeyStatusEncryptionKeySet
	MBusEncryptionKeyStatusEncryptionKeyTransferred
	MBusEncryptionKeyStatusEncryptionKeySetAndTransferred
	MBusEncryptionKeyStatusEncryptionKeyInUse
)

// MBusEncryptionKeyStatusParse converts the given string into a MBusEncryptionKeyStatus value.
//
// It returns the corresponding MBusEncryptionKeyStatus constant if the string matches
// a known level name, or an error if the input is invalid.
func MBusEncryptionKeyStatusParse(value string) (MBusEncryptionKeyStatus, error) {
	var ret MBusEncryptionKeyStatus
	var err error
	switch {
	case strings.EqualFold(value, "NoEncryptionKey"):
		ret = MBusEncryptionKeyStatusNoEncryptionKey
	case strings.EqualFold(value, "EncryptionKeySet"):
		ret = MBusEncryptionKeyStatusEncryptionKeySet
	case strings.EqualFold(value, "EncryptionKeyTransferred"):
		ret = MBusEncryptionKeyStatusEncryptionKeyTransferred
	case strings.EqualFold(value, "EncryptionKeySetAndTransferred"):
		ret = MBusEncryptionKeyStatusEncryptionKeySetAndTransferred
	case strings.EqualFold(value, "EncryptionKeyInUse"):
		ret = MBusEncryptionKeyStatusEncryptionKeyInUse
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the MBusEncryptionKeyStatus.
// It satisfies fmt.Stringer.
func (g MBusEncryptionKeyStatus) String() string {
	var ret string
	switch g {
	case MBusEncryptionKeyStatusNoEncryptionKey:
		ret = "NoEncryptionKey"
	case MBusEncryptionKeyStatusEncryptionKeySet:
		ret = "EncryptionKeySet"
	case MBusEncryptionKeyStatusEncryptionKeyTransferred:
		ret = "EncryptionKeyTransferred"
	case MBusEncryptionKeyStatusEncryptionKeySetAndTransferred:
		ret = "EncryptionKeySetAndTransferred"
	case MBusEncryptionKeyStatusEncryptionKeyInUse:
		ret = "EncryptionKeyInUse"
	}
	return ret
}

// AllMBusEncryptionKeyStatus returns a slice containing all defined MBusEncryptionKeyStatus values.
func AllMBusEncryptionKeyStatus() []MBusEncryptionKeyStatus {
	return []MBusEncryptionKeyStatus{
	MBusEncryptionKeyStatusNoEncryptionKey,
	MBusEncryptionKeyStatusEncryptionKeySet,
	MBusEncryptionKeyStatusEncryptionKeyTransferred,
	MBusEncryptionKeyStatusEncryptionKeySetAndTransferred,
	MBusEncryptionKeyStatusEncryptionKeyInUse,
	}
}
