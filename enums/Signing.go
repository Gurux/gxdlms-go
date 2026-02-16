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

// Signing Enumerates :  types.
type Signing int

const (
	// SigningNone defines that the signing is not used.
	SigningNone Signing = iota
	// SigningEphemeralUnifiedModel defines that the Ephemeral Unified Model scheme. Messages are digitally signed and send with general-ciphering messages.
	SigningEphemeralUnifiedModel
	// SigningOnePassDiffieHellman defines that the One-Pass Diffie-Hellman scheme. Messages are digitally signed and send with general-ciphering messages.
	SigningOnePassDiffieHellman
	// SigningStaticUnifiedModel defines that the Static Unified Model scheme. Messages are digitally signed and send with general-ciphering messages.
	SigningStaticUnifiedModel
	// SigningGeneralSigning defines that the general signing is used.
	SigningGeneralSigning
)

// SigningParse converts the given string into a Signing value.
//
// It returns the corresponding Signing constant if the string matches
// a known level name, or an error if the input is invalid.
func SigningParse(value string) (Signing, error) {
	var ret Signing
	var err error
	switch strings.ToUpper(value) {
	case "NONE":
		ret = SigningNone
	case "EPHEMERALUNIFIEDMODEL":
		ret = SigningEphemeralUnifiedModel
	case "ONEPASSDIFFIEHELLMAN":
		ret = SigningOnePassDiffieHellman
	case "STATICUNIFIEDMODEL":
		ret = SigningStaticUnifiedModel
	case "GENERALSIGNING":
		ret = SigningGeneralSigning
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the Signing.
// It satisfies fmt.Stringer.
func (g Signing) String() string {
	var ret string
	switch g {
	case SigningNone:
		ret = "NONE"
	case SigningEphemeralUnifiedModel:
		ret = "EPHEMERALUNIFIEDMODEL"
	case SigningOnePassDiffieHellman:
		ret = "ONEPASSDIFFIEHELLMAN"
	case SigningStaticUnifiedModel:
		ret = "STATICUNIFIEDMODEL"
	case SigningGeneralSigning:
		ret = "GENERALSIGNING"
	}
	return ret
}
