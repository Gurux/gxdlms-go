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

// KeyUsage Key Usage.
type KeyUsage int

const (
	// KeyUsageNone defines that key is not used bit is set.
	KeyUsageNone KeyUsage = iota
	// KeyUsageDigitalSignature defines that digital signature bit is set.
	KeyUsageDigitalSignature KeyUsage = 0x1
	// KeyUsageNonRepudiation defines that non Repudiation bit is set.
	KeyUsageNonRepudiation KeyUsage = 0x2
	// KeyUsageKeyEncipherment defines that key encipherment bit is set.
	KeyUsageKeyEncipherment KeyUsage = 0x4
	// KeyUsageDataEncipherment defines that data encipherment bit is set.
	KeyUsageDataEncipherment KeyUsage = 0x8
	// KeyUsageKeyAgreement defines that key agreement bit is set.
	KeyUsageKeyAgreement KeyUsage = 0x10
	// KeyUsageKeyCertSign defines that used with CA certificates when the subject public key is used to verify a signature on certificates bit is set.
	KeyUsageKeyCertSign KeyUsage = 0x20
	// KeyUsageCrlSign defines that used when the subject public key is to verify a signature bit is set.
	KeyUsageCrlSign KeyUsage = 0x40
	// KeyUsageEncipherOnly defines that encipher only bit is set.
	KeyUsageEncipherOnly KeyUsage = 0x80
	// KeyUsageDecipherOnly defines that decipher only bit is set.
	KeyUsageDecipherOnly KeyUsage = 0x100
)

// KeyUsageParse converts the given string into a KeyUsage value.
//
// It returns the corresponding KeyUsage constant if the string matches
// a known level name, or an error if the input is invalid.
func KeyUsageParse(value string) (KeyUsage, error) {
	var ret KeyUsage
	var err error
	switch strings.ToUpper(value) {
	case "NONE":
		ret = KeyUsageNone
	case "DIGITALSIGNATURE":
		ret = KeyUsageDigitalSignature
	case "NONREPUDIATION":
		ret = KeyUsageNonRepudiation
	case "KEYENCIPHERMENT":
		ret = KeyUsageKeyEncipherment
	case "DATAENCIPHERMENT":
		ret = KeyUsageDataEncipherment
	case "KEYAGREEMENT":
		ret = KeyUsageKeyAgreement
	case "KEYCERTSIGN":
		ret = KeyUsageKeyCertSign
	case "CRLSIGN":
		ret = KeyUsageCrlSign
	case "ENCIPHERONLY":
		ret = KeyUsageEncipherOnly
	case "DECIPHERONLY":
		ret = KeyUsageDecipherOnly
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the KeyUsage.
// It satisfies fmt.Stringer.
func (g KeyUsage) String() string {
	var ret string
	switch g {
	case KeyUsageNone:
		ret = "NONE"
	case KeyUsageDigitalSignature:
		ret = "DIGITALSIGNATURE"
	case KeyUsageNonRepudiation:
		ret = "NONREPUDIATION"
	case KeyUsageKeyEncipherment:
		ret = "KEYENCIPHERMENT"
	case KeyUsageDataEncipherment:
		ret = "DATAENCIPHERMENT"
	case KeyUsageKeyAgreement:
		ret = "KEYAGREEMENT"
	case KeyUsageKeyCertSign:
		ret = "KEYCERTSIGN"
	case KeyUsageCrlSign:
		ret = "CRLSIGN"
	case KeyUsageEncipherOnly:
		ret = "ENCIPHERONLY"
	case KeyUsageDecipherOnly:
		ret = "DECIPHERONLY"
	}
	return ret
}
