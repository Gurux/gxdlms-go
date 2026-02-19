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

type X9ObjectIdentifier int

const (
	X9ObjectIdentifierNone X9ObjectIdentifier = iota
	X9ObjectIdentifierIdFieldType
	X9ObjectIdentifierPrimeField
	X9ObjectIdentifierCharacteristicTwoField
	X9ObjectIdentifierGNBasis
	X9ObjectIdentifierTPBasis
	X9ObjectIdentifierPPBasis
	X9ObjectIdentifierECDsaWithSha1
	X9ObjectIdentifierIdECPublicKey
	X9ObjectIdentifierECDsaWithSha2
	X9ObjectIdentifierECDsaWithSha224
	X9ObjectIdentifierECDsaWithSha256
	X9ObjectIdentifierECDsaWithSha384
	X9ObjectIdentifierECDsaWithSha512
	X9ObjectIdentifierEllipticCurve
	X9ObjectIdentifierCTwoCurve
	X9ObjectIdentifierC2Pnb163v1
	X9ObjectIdentifierC2Pnb163v2
	X9ObjectIdentifierC2Pnb163v3
	X9ObjectIdentifierC2Pnb176w1
	X9ObjectIdentifierC2Tnb191v1
	X9ObjectIdentifierC2Tnb191v2
	X9ObjectIdentifierC2Tnb191v3
	X9ObjectIdentifierC2Onb191v4
	X9ObjectIdentifierC2Onb191v5
	X9ObjectIdentifierC2Pnb208w1
	X9ObjectIdentifierC2Tnb239v1
	X9ObjectIdentifierC2Tnb239v2
	X9ObjectIdentifierC2Tnb239v3
	X9ObjectIdentifierC2Onb239v4
	X9ObjectIdentifierC2Onb239v5
	X9ObjectIdentifierC2Pnb272w1
	X9ObjectIdentifierC2Pnb304w1
	X9ObjectIdentifierC2Tnb359v1
	X9ObjectIdentifierC2Pnb368w1
	X9ObjectIdentifierC2Tnb431r1
	X9ObjectIdentifierPrimeCurve
	X9ObjectIdentifierPrime192v1
	X9ObjectIdentifierPrime192v2
	X9ObjectIdentifierPrime192v3
	X9ObjectIdentifierPrime239v1
	X9ObjectIdentifierPrime239v2
	X9ObjectIdentifierPrime239v3
	X9ObjectIdentifierPrime256v1
	X9ObjectIdentifierIdDsa
	X9ObjectIdentifierIdDsaWithSha1
	X9ObjectIdentifierX9x63Scheme
	X9ObjectIdentifierDHSinglePassStdDHSha1KdfScheme
	X9ObjectIdentifierDHSinglePassCofactorDHSha1KdfScheme
	X9ObjectIdentifierMqvSinglePassSha1KdfScheme
	X9ObjectIdentifieransi_x9_42
	X9ObjectIdentifierDHPublicNumber
	X9ObjectIdentifierX9x42Schemes
	X9ObjectIdentifierDHStatic
	X9ObjectIdentifierDHEphem
	X9ObjectIdentifierDHOneFlow
	X9ObjectIdentifierDHHybrid1
	X9ObjectIdentifierDHHybrid2
	X9ObjectIdentifierDHHybridOneFlow
	X9ObjectIdentifierMqv2
	X9ObjectIdentifierMqv1
	X9ObjectIdentifierSecp384r1
)

// X9ObjectIdentifierParse converts the given string into a X9ObjectIdentifier value.
//
// It returns the corresponding X9ObjectIdentifier constant if the string matches
// a known level name, or an error if the input is invalid.
func X9ObjectIdentifierParse(value string) (X9ObjectIdentifier, error) {
	var ret X9ObjectIdentifier
	var err error
	switch {
	case strings.EqualFold(value, "None"):
		ret = X9ObjectIdentifierNone
	case strings.EqualFold(value, "IdFieldType"):
		ret = X9ObjectIdentifierIdFieldType
	case strings.EqualFold(value, "PrimeField"):
		ret = X9ObjectIdentifierPrimeField
	case strings.EqualFold(value, "CharacteristicTwoField"):
		ret = X9ObjectIdentifierCharacteristicTwoField
	case strings.EqualFold(value, "GNBasis"):
		ret = X9ObjectIdentifierGNBasis
	case strings.EqualFold(value, "TPBasis"):
		ret = X9ObjectIdentifierTPBasis
	case strings.EqualFold(value, "PPBasis"):
		ret = X9ObjectIdentifierPPBasis
	case strings.EqualFold(value, "ECDsaWithSha1"):
		ret = X9ObjectIdentifierECDsaWithSha1
	case strings.EqualFold(value, "IdECPublicKey"):
		ret = X9ObjectIdentifierIdECPublicKey
	case strings.EqualFold(value, "ECDsaWithSha2"):
		ret = X9ObjectIdentifierECDsaWithSha2
	case strings.EqualFold(value, "ECDsaWithSha224"):
		ret = X9ObjectIdentifierECDsaWithSha224
	case strings.EqualFold(value, "ECDsaWithSha256"):
		ret = X9ObjectIdentifierECDsaWithSha256
	case strings.EqualFold(value, "ECDsaWithSha384"):
		ret = X9ObjectIdentifierECDsaWithSha384
	case strings.EqualFold(value, "ECDsaWithSha512"):
		ret = X9ObjectIdentifierECDsaWithSha512
	case strings.EqualFold(value, "EllipticCurve"):
		ret = X9ObjectIdentifierEllipticCurve
	case strings.EqualFold(value, "CTwoCurve"):
		ret = X9ObjectIdentifierCTwoCurve
	case strings.EqualFold(value, "C2Pnb163v1"):
		ret = X9ObjectIdentifierC2Pnb163v1
	case strings.EqualFold(value, "C2Pnb163v2"):
		ret = X9ObjectIdentifierC2Pnb163v2
	case strings.EqualFold(value, "C2Pnb163v3"):
		ret = X9ObjectIdentifierC2Pnb163v3
	case strings.EqualFold(value, "C2Pnb176w1"):
		ret = X9ObjectIdentifierC2Pnb176w1
	case strings.EqualFold(value, "C2Tnb191v1"):
		ret = X9ObjectIdentifierC2Tnb191v1
	case strings.EqualFold(value, "C2Tnb191v2"):
		ret = X9ObjectIdentifierC2Tnb191v2
	case strings.EqualFold(value, "C2Tnb191v3"):
		ret = X9ObjectIdentifierC2Tnb191v3
	case strings.EqualFold(value, "C2Onb191v4"):
		ret = X9ObjectIdentifierC2Onb191v4
	case strings.EqualFold(value, "C2Onb191v5"):
		ret = X9ObjectIdentifierC2Onb191v5
	case strings.EqualFold(value, "C2Pnb208w1"):
		ret = X9ObjectIdentifierC2Pnb208w1
	case strings.EqualFold(value, "C2Tnb239v1"):
		ret = X9ObjectIdentifierC2Tnb239v1
	case strings.EqualFold(value, "C2Tnb239v2"):
		ret = X9ObjectIdentifierC2Tnb239v2
	case strings.EqualFold(value, "C2Tnb239v3"):
		ret = X9ObjectIdentifierC2Tnb239v3
	case strings.EqualFold(value, "C2Onb239v4"):
		ret = X9ObjectIdentifierC2Onb239v4
	case strings.EqualFold(value, "C2Onb239v5"):
		ret = X9ObjectIdentifierC2Onb239v5
	case strings.EqualFold(value, "C2Pnb272w1"):
		ret = X9ObjectIdentifierC2Pnb272w1
	case strings.EqualFold(value, "C2Pnb304w1"):
		ret = X9ObjectIdentifierC2Pnb304w1
	case strings.EqualFold(value, "C2Tnb359v1"):
		ret = X9ObjectIdentifierC2Tnb359v1
	case strings.EqualFold(value, "C2Pnb368w1"):
		ret = X9ObjectIdentifierC2Pnb368w1
	case strings.EqualFold(value, "C2Tnb431r1"):
		ret = X9ObjectIdentifierC2Tnb431r1
	case strings.EqualFold(value, "PrimeCurve"):
		ret = X9ObjectIdentifierPrimeCurve
	case strings.EqualFold(value, "Prime192v1"):
		ret = X9ObjectIdentifierPrime192v1
	case strings.EqualFold(value, "Prime192v2"):
		ret = X9ObjectIdentifierPrime192v2
	case strings.EqualFold(value, "Prime192v3"):
		ret = X9ObjectIdentifierPrime192v3
	case strings.EqualFold(value, "Prime239v1"):
		ret = X9ObjectIdentifierPrime239v1
	case strings.EqualFold(value, "Prime239v2"):
		ret = X9ObjectIdentifierPrime239v2
	case strings.EqualFold(value, "Prime239v3"):
		ret = X9ObjectIdentifierPrime239v3
	case strings.EqualFold(value, "Prime256v1"):
		ret = X9ObjectIdentifierPrime256v1
	case strings.EqualFold(value, "IdDsa"):
		ret = X9ObjectIdentifierIdDsa
	case strings.EqualFold(value, "IdDsaWithSha1"):
		ret = X9ObjectIdentifierIdDsaWithSha1
	case strings.EqualFold(value, "X9x63Scheme"):
		ret = X9ObjectIdentifierX9x63Scheme
	case strings.EqualFold(value, "DHSinglePassStdDHSha1KdfScheme"):
		ret = X9ObjectIdentifierDHSinglePassStdDHSha1KdfScheme
	case strings.EqualFold(value, "DHSinglePassCofactorDHSha1KdfScheme"):
		ret = X9ObjectIdentifierDHSinglePassCofactorDHSha1KdfScheme
	case strings.EqualFold(value, "MqvSinglePassSha1KdfScheme"):
		ret = X9ObjectIdentifierMqvSinglePassSha1KdfScheme
	case strings.EqualFold(value, "ansi_x9_42"):
		ret = X9ObjectIdentifieransi_x9_42
	case strings.EqualFold(value, "DHPublicNumber"):
		ret = X9ObjectIdentifierDHPublicNumber
	case strings.EqualFold(value, "X9x42Schemes"):
		ret = X9ObjectIdentifierX9x42Schemes
	case strings.EqualFold(value, "DHStatic"):
		ret = X9ObjectIdentifierDHStatic
	case strings.EqualFold(value, "DHEphem"):
		ret = X9ObjectIdentifierDHEphem
	case strings.EqualFold(value, "DHOneFlow"):
		ret = X9ObjectIdentifierDHOneFlow
	case strings.EqualFold(value, "DHHybrid1"):
		ret = X9ObjectIdentifierDHHybrid1
	case strings.EqualFold(value, "DHHybrid2"):
		ret = X9ObjectIdentifierDHHybrid2
	case strings.EqualFold(value, "DHHybridOneFlow"):
		ret = X9ObjectIdentifierDHHybridOneFlow
	case strings.EqualFold(value, "Mqv2"):
		ret = X9ObjectIdentifierMqv2
	case strings.EqualFold(value, "Mqv1"):
		ret = X9ObjectIdentifierMqv1
	case strings.EqualFold(value, "Secp384r1"):
		ret = X9ObjectIdentifierSecp384r1
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the X9ObjectIdentifier.
// It satisfies fmt.Stringer.
func (g X9ObjectIdentifier) String() string {
	var ret string
	switch g {
	case X9ObjectIdentifierNone:
		ret = "None"
	case X9ObjectIdentifierIdFieldType:
		ret = "IdFieldType"
	case X9ObjectIdentifierPrimeField:
		ret = "PrimeField"
	case X9ObjectIdentifierCharacteristicTwoField:
		ret = "CharacteristicTwoField"
	case X9ObjectIdentifierGNBasis:
		ret = "GNBasis"
	case X9ObjectIdentifierTPBasis:
		ret = "TPBasis"
	case X9ObjectIdentifierPPBasis:
		ret = "PPBasis"
	case X9ObjectIdentifierECDsaWithSha1:
		ret = "ECDsaWithSha1"
	case X9ObjectIdentifierIdECPublicKey:
		ret = "IdECPublicKey"
	case X9ObjectIdentifierECDsaWithSha2:
		ret = "ECDsaWithSha2"
	case X9ObjectIdentifierECDsaWithSha224:
		ret = "ECDsaWithSha224"
	case X9ObjectIdentifierECDsaWithSha256:
		ret = "ECDsaWithSha256"
	case X9ObjectIdentifierECDsaWithSha384:
		ret = "ECDsaWithSha384"
	case X9ObjectIdentifierECDsaWithSha512:
		ret = "ECDsaWithSha512"
	case X9ObjectIdentifierEllipticCurve:
		ret = "EllipticCurve"
	case X9ObjectIdentifierCTwoCurve:
		ret = "CTwoCurve"
	case X9ObjectIdentifierC2Pnb163v1:
		ret = "C2Pnb163v1"
	case X9ObjectIdentifierC2Pnb163v2:
		ret = "C2Pnb163v2"
	case X9ObjectIdentifierC2Pnb163v3:
		ret = "C2Pnb163v3"
	case X9ObjectIdentifierC2Pnb176w1:
		ret = "C2Pnb176w1"
	case X9ObjectIdentifierC2Tnb191v1:
		ret = "C2Tnb191v1"
	case X9ObjectIdentifierC2Tnb191v2:
		ret = "C2Tnb191v2"
	case X9ObjectIdentifierC2Tnb191v3:
		ret = "C2Tnb191v3"
	case X9ObjectIdentifierC2Onb191v4:
		ret = "C2Onb191v4"
	case X9ObjectIdentifierC2Onb191v5:
		ret = "C2Onb191v5"
	case X9ObjectIdentifierC2Pnb208w1:
		ret = "C2Pnb208w1"
	case X9ObjectIdentifierC2Tnb239v1:
		ret = "C2Tnb239v1"
	case X9ObjectIdentifierC2Tnb239v2:
		ret = "C2Tnb239v2"
	case X9ObjectIdentifierC2Tnb239v3:
		ret = "C2Tnb239v3"
	case X9ObjectIdentifierC2Onb239v4:
		ret = "C2Onb239v4"
	case X9ObjectIdentifierC2Onb239v5:
		ret = "C2Onb239v5"
	case X9ObjectIdentifierC2Pnb272w1:
		ret = "C2Pnb272w1"
	case X9ObjectIdentifierC2Pnb304w1:
		ret = "C2Pnb304w1"
	case X9ObjectIdentifierC2Tnb359v1:
		ret = "C2Tnb359v1"
	case X9ObjectIdentifierC2Pnb368w1:
		ret = "C2Pnb368w1"
	case X9ObjectIdentifierC2Tnb431r1:
		ret = "C2Tnb431r1"
	case X9ObjectIdentifierPrimeCurve:
		ret = "PrimeCurve"
	case X9ObjectIdentifierPrime192v1:
		ret = "Prime192v1"
	case X9ObjectIdentifierPrime192v2:
		ret = "Prime192v2"
	case X9ObjectIdentifierPrime192v3:
		ret = "Prime192v3"
	case X9ObjectIdentifierPrime239v1:
		ret = "Prime239v1"
	case X9ObjectIdentifierPrime239v2:
		ret = "Prime239v2"
	case X9ObjectIdentifierPrime239v3:
		ret = "Prime239v3"
	case X9ObjectIdentifierPrime256v1:
		ret = "Prime256v1"
	case X9ObjectIdentifierIdDsa:
		ret = "IdDsa"
	case X9ObjectIdentifierIdDsaWithSha1:
		ret = "IdDsaWithSha1"
	case X9ObjectIdentifierX9x63Scheme:
		ret = "X9x63Scheme"
	case X9ObjectIdentifierDHSinglePassStdDHSha1KdfScheme:
		ret = "DHSinglePassStdDHSha1KdfScheme"
	case X9ObjectIdentifierDHSinglePassCofactorDHSha1KdfScheme:
		ret = "DHSinglePassCofactorDHSha1KdfScheme"
	case X9ObjectIdentifierMqvSinglePassSha1KdfScheme:
		ret = "MqvSinglePassSha1KdfScheme"
	case X9ObjectIdentifieransi_x9_42:
		ret = "ansi_x9_42"
	case X9ObjectIdentifierDHPublicNumber:
		ret = "DHPublicNumber"
	case X9ObjectIdentifierX9x42Schemes:
		ret = "X9x42Schemes"
	case X9ObjectIdentifierDHStatic:
		ret = "DHStatic"
	case X9ObjectIdentifierDHEphem:
		ret = "DHEphem"
	case X9ObjectIdentifierDHOneFlow:
		ret = "DHOneFlow"
	case X9ObjectIdentifierDHHybrid1:
		ret = "DHHybrid1"
	case X9ObjectIdentifierDHHybrid2:
		ret = "DHHybrid2"
	case X9ObjectIdentifierDHHybridOneFlow:
		ret = "DHHybridOneFlow"
	case X9ObjectIdentifierMqv2:
		ret = "Mqv2"
	case X9ObjectIdentifierMqv1:
		ret = "Mqv1"
	case X9ObjectIdentifierSecp384r1:
		ret = "Secp384r1"
	}
	return ret
}

// AllX9ObjectIdentifier returns a slice containing all defined X9ObjectIdentifier values.
func AllX9ObjectIdentifier() []X9ObjectIdentifier {
	return []X9ObjectIdentifier{
	X9ObjectIdentifierNone,
	X9ObjectIdentifierIdFieldType,
	X9ObjectIdentifierPrimeField,
	X9ObjectIdentifierCharacteristicTwoField,
	X9ObjectIdentifierGNBasis,
	X9ObjectIdentifierTPBasis,
	X9ObjectIdentifierPPBasis,
	X9ObjectIdentifierECDsaWithSha1,
	X9ObjectIdentifierIdECPublicKey,
	X9ObjectIdentifierECDsaWithSha2,
	X9ObjectIdentifierECDsaWithSha224,
	X9ObjectIdentifierECDsaWithSha256,
	X9ObjectIdentifierECDsaWithSha384,
	X9ObjectIdentifierECDsaWithSha512,
	X9ObjectIdentifierEllipticCurve,
	X9ObjectIdentifierCTwoCurve,
	X9ObjectIdentifierC2Pnb163v1,
	X9ObjectIdentifierC2Pnb163v2,
	X9ObjectIdentifierC2Pnb163v3,
	X9ObjectIdentifierC2Pnb176w1,
	X9ObjectIdentifierC2Tnb191v1,
	X9ObjectIdentifierC2Tnb191v2,
	X9ObjectIdentifierC2Tnb191v3,
	X9ObjectIdentifierC2Onb191v4,
	X9ObjectIdentifierC2Onb191v5,
	X9ObjectIdentifierC2Pnb208w1,
	X9ObjectIdentifierC2Tnb239v1,
	X9ObjectIdentifierC2Tnb239v2,
	X9ObjectIdentifierC2Tnb239v3,
	X9ObjectIdentifierC2Onb239v4,
	X9ObjectIdentifierC2Onb239v5,
	X9ObjectIdentifierC2Pnb272w1,
	X9ObjectIdentifierC2Pnb304w1,
	X9ObjectIdentifierC2Tnb359v1,
	X9ObjectIdentifierC2Pnb368w1,
	X9ObjectIdentifierC2Tnb431r1,
	X9ObjectIdentifierPrimeCurve,
	X9ObjectIdentifierPrime192v1,
	X9ObjectIdentifierPrime192v2,
	X9ObjectIdentifierPrime192v3,
	X9ObjectIdentifierPrime239v1,
	X9ObjectIdentifierPrime239v2,
	X9ObjectIdentifierPrime239v3,
	X9ObjectIdentifierPrime256v1,
	X9ObjectIdentifierIdDsa,
	X9ObjectIdentifierIdDsaWithSha1,
	X9ObjectIdentifierX9x63Scheme,
	X9ObjectIdentifierDHSinglePassStdDHSha1KdfScheme,
	X9ObjectIdentifierDHSinglePassCofactorDHSha1KdfScheme,
	X9ObjectIdentifierMqvSinglePassSha1KdfScheme,
	X9ObjectIdentifieransi_x9_42,
	X9ObjectIdentifierDHPublicNumber,
	X9ObjectIdentifierX9x42Schemes,
	X9ObjectIdentifierDHStatic,
	X9ObjectIdentifierDHEphem,
	X9ObjectIdentifierDHOneFlow,
	X9ObjectIdentifierDHHybrid1,
	X9ObjectIdentifierDHHybrid2,
	X9ObjectIdentifierDHHybridOneFlow,
	X9ObjectIdentifierMqv2,
	X9ObjectIdentifierMqv1,
	X9ObjectIdentifierSecp384r1,
	}
}
