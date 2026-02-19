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

// KeyAgreementScheme enumerates key agreement schemes.
type KeyAgreementScheme int

const (
	// KeyAgreementSchemeEphemeralUnifiedModel defines that the the Ephemeral Unified Model scheme is used.
	KeyAgreementSchemeEphemeralUnifiedModel KeyAgreementScheme = iota
	// KeyAgreementSchemeOnePassDiffieHellman defines that the the One-Pass Diffie-Hellman scheme is used.
	KeyAgreementSchemeOnePassDiffieHellman
	// KeyAgreementSchemeStaticUnifiedModel defines that the the Static Unified Model scheme is used.
	KeyAgreementSchemeStaticUnifiedModel
)

// KeyAgreementSchemeParse converts the given string into a KeyAgreementScheme value.
//
// It returns the corresponding KeyAgreementScheme constant if the string matches
// a known level name, or an error if the input is invalid.
func KeyAgreementSchemeParse(value string) (KeyAgreementScheme, error) {
	var ret KeyAgreementScheme
	var err error
	switch {
	case strings.EqualFold(value, "EphemeralUnifiedModel"):
		ret = KeyAgreementSchemeEphemeralUnifiedModel
	case strings.EqualFold(value, "OnePassDiffieHellman"):
		ret = KeyAgreementSchemeOnePassDiffieHellman
	case strings.EqualFold(value, "StaticUnifiedModel"):
		ret = KeyAgreementSchemeStaticUnifiedModel
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the KeyAgreementScheme.
// It satisfies fmt.Stringer.
func (g KeyAgreementScheme) String() string {
	var ret string
	switch g {
	case KeyAgreementSchemeEphemeralUnifiedModel:
		ret = "EphemeralUnifiedModel"
	case KeyAgreementSchemeOnePassDiffieHellman:
		ret = "OnePassDiffieHellman"
	case KeyAgreementSchemeStaticUnifiedModel:
		ret = "StaticUnifiedModel"
	}
	return ret
}

// AllKeyAgreementScheme returns a slice containing all defined KeyAgreementScheme values.
func AllKeyAgreementScheme() []KeyAgreementScheme {
	return []KeyAgreementScheme{
		KeyAgreementSchemeEphemeralUnifiedModel,
		KeyAgreementSchemeOnePassDiffieHellman,
		KeyAgreementSchemeStaticUnifiedModel,
	}
}
