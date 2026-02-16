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
	switch strings.ToUpper(value) {
	case "NONE":
		ret = X9ObjectIdentifierNone
	case "IDFIELDTYPE":
		ret = X9ObjectIdentifierIdFieldType
	case "PRIMEFIELD":
		ret = X9ObjectIdentifierPrimeField
	case "CHARACTERISTICTWOFIELD":
		ret = X9ObjectIdentifierCharacteristicTwoField
	case "GNBASIS":
		ret = X9ObjectIdentifierGNBasis
	case "TPBASIS":
		ret = X9ObjectIdentifierTPBasis
	case "PPBASIS":
		ret = X9ObjectIdentifierPPBasis
	case "ECDSAWITHSHA1":
		ret = X9ObjectIdentifierECDsaWithSha1
	case "IDECPUBLICKEY":
		ret = X9ObjectIdentifierIdECPublicKey
	case "ECDSAWITHSHA2":
		ret = X9ObjectIdentifierECDsaWithSha2
	case "ECDSAWITHSHA224":
		ret = X9ObjectIdentifierECDsaWithSha224
	case "ECDSAWITHSHA256":
		ret = X9ObjectIdentifierECDsaWithSha256
	case "ECDSAWITHSHA384":
		ret = X9ObjectIdentifierECDsaWithSha384
	case "ECDSAWITHSHA512":
		ret = X9ObjectIdentifierECDsaWithSha512
	case "ELLIPTICCURVE":
		ret = X9ObjectIdentifierEllipticCurve
	case "CTWOCURVE":
		ret = X9ObjectIdentifierCTwoCurve
	case "C2PNB163V1":
		ret = X9ObjectIdentifierC2Pnb163v1
	case "C2PNB163V2":
		ret = X9ObjectIdentifierC2Pnb163v2
	case "C2PNB163V3":
		ret = X9ObjectIdentifierC2Pnb163v3
	case "C2PNB176W1":
		ret = X9ObjectIdentifierC2Pnb176w1
	case "C2TNB191V1":
		ret = X9ObjectIdentifierC2Tnb191v1
	case "C2TNB191V2":
		ret = X9ObjectIdentifierC2Tnb191v2
	case "C2TNB191V3":
		ret = X9ObjectIdentifierC2Tnb191v3
	case "C2ONB191V4":
		ret = X9ObjectIdentifierC2Onb191v4
	case "C2ONB191V5":
		ret = X9ObjectIdentifierC2Onb191v5
	case "C2PNB208W1":
		ret = X9ObjectIdentifierC2Pnb208w1
	case "C2TNB239V1":
		ret = X9ObjectIdentifierC2Tnb239v1
	case "C2TNB239V2":
		ret = X9ObjectIdentifierC2Tnb239v2
	case "C2TNB239V3":
		ret = X9ObjectIdentifierC2Tnb239v3
	case "C2ONB239V4":
		ret = X9ObjectIdentifierC2Onb239v4
	case "C2ONB239V5":
		ret = X9ObjectIdentifierC2Onb239v5
	case "C2PNB272W1":
		ret = X9ObjectIdentifierC2Pnb272w1
	case "C2PNB304W1":
		ret = X9ObjectIdentifierC2Pnb304w1
	case "C2TNB359V1":
		ret = X9ObjectIdentifierC2Tnb359v1
	case "C2PNB368W1":
		ret = X9ObjectIdentifierC2Pnb368w1
	case "C2TNB431R1":
		ret = X9ObjectIdentifierC2Tnb431r1
	case "PRIMECURVE":
		ret = X9ObjectIdentifierPrimeCurve
	case "PRIME192V1":
		ret = X9ObjectIdentifierPrime192v1
	case "PRIME192V2":
		ret = X9ObjectIdentifierPrime192v2
	case "PRIME192V3":
		ret = X9ObjectIdentifierPrime192v3
	case "PRIME239V1":
		ret = X9ObjectIdentifierPrime239v1
	case "PRIME239V2":
		ret = X9ObjectIdentifierPrime239v2
	case "PRIME239V3":
		ret = X9ObjectIdentifierPrime239v3
	case "PRIME256V1":
		ret = X9ObjectIdentifierPrime256v1
	case "IDDSA":
		ret = X9ObjectIdentifierIdDsa
	case "IDDSAWITHSHA1":
		ret = X9ObjectIdentifierIdDsaWithSha1
	case "X9X63SCHEME":
		ret = X9ObjectIdentifierX9x63Scheme
	case "DHSINGLEPASSSTDDHSHA1KDFSCHEME":
		ret = X9ObjectIdentifierDHSinglePassStdDHSha1KdfScheme
	case "DHSINGLEPASSCOFACTORDHSHA1KDFSCHEME":
		ret = X9ObjectIdentifierDHSinglePassCofactorDHSha1KdfScheme
	case "MQVSINGLEPASSSHA1KDFSCHEME":
		ret = X9ObjectIdentifierMqvSinglePassSha1KdfScheme
	case "ANSI_X9_42":
		ret = X9ObjectIdentifieransi_x9_42
	case "DHPUBLICNUMBER":
		ret = X9ObjectIdentifierDHPublicNumber
	case "X9X42SCHEMES":
		ret = X9ObjectIdentifierX9x42Schemes
	case "DHSTATIC":
		ret = X9ObjectIdentifierDHStatic
	case "DHEPHEM":
		ret = X9ObjectIdentifierDHEphem
	case "DHONEFLOW":
		ret = X9ObjectIdentifierDHOneFlow
	case "DHHYBRID1":
		ret = X9ObjectIdentifierDHHybrid1
	case "DHHYBRID2":
		ret = X9ObjectIdentifierDHHybrid2
	case "DHHYBRIDONEFLOW":
		ret = X9ObjectIdentifierDHHybridOneFlow
	case "MQV2":
		ret = X9ObjectIdentifierMqv2
	case "MQV1":
		ret = X9ObjectIdentifierMqv1
	case "SECP384R1":
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
		ret = "NONE"
	case X9ObjectIdentifierIdFieldType:
		ret = "IDFIELDTYPE"
	case X9ObjectIdentifierPrimeField:
		ret = "PRIMEFIELD"
	case X9ObjectIdentifierCharacteristicTwoField:
		ret = "CHARACTERISTICTWOFIELD"
	case X9ObjectIdentifierGNBasis:
		ret = "GNBASIS"
	case X9ObjectIdentifierTPBasis:
		ret = "TPBASIS"
	case X9ObjectIdentifierPPBasis:
		ret = "PPBASIS"
	case X9ObjectIdentifierECDsaWithSha1:
		ret = "ECDSAWITHSHA1"
	case X9ObjectIdentifierIdECPublicKey:
		ret = "IDECPUBLICKEY"
	case X9ObjectIdentifierECDsaWithSha2:
		ret = "ECDSAWITHSHA2"
	case X9ObjectIdentifierECDsaWithSha224:
		ret = "ECDSAWITHSHA224"
	case X9ObjectIdentifierECDsaWithSha256:
		ret = "ECDSAWITHSHA256"
	case X9ObjectIdentifierECDsaWithSha384:
		ret = "ECDSAWITHSHA384"
	case X9ObjectIdentifierECDsaWithSha512:
		ret = "ECDSAWITHSHA512"
	case X9ObjectIdentifierEllipticCurve:
		ret = "ELLIPTICCURVE"
	case X9ObjectIdentifierCTwoCurve:
		ret = "CTWOCURVE"
	case X9ObjectIdentifierC2Pnb163v1:
		ret = "C2PNB163V1"
	case X9ObjectIdentifierC2Pnb163v2:
		ret = "C2PNB163V2"
	case X9ObjectIdentifierC2Pnb163v3:
		ret = "C2PNB163V3"
	case X9ObjectIdentifierC2Pnb176w1:
		ret = "C2PNB176W1"
	case X9ObjectIdentifierC2Tnb191v1:
		ret = "C2TNB191V1"
	case X9ObjectIdentifierC2Tnb191v2:
		ret = "C2TNB191V2"
	case X9ObjectIdentifierC2Tnb191v3:
		ret = "C2TNB191V3"
	case X9ObjectIdentifierC2Onb191v4:
		ret = "C2ONB191V4"
	case X9ObjectIdentifierC2Onb191v5:
		ret = "C2ONB191V5"
	case X9ObjectIdentifierC2Pnb208w1:
		ret = "C2PNB208W1"
	case X9ObjectIdentifierC2Tnb239v1:
		ret = "C2TNB239V1"
	case X9ObjectIdentifierC2Tnb239v2:
		ret = "C2TNB239V2"
	case X9ObjectIdentifierC2Tnb239v3:
		ret = "C2TNB239V3"
	case X9ObjectIdentifierC2Onb239v4:
		ret = "C2ONB239V4"
	case X9ObjectIdentifierC2Onb239v5:
		ret = "C2ONB239V5"
	case X9ObjectIdentifierC2Pnb272w1:
		ret = "C2PNB272W1"
	case X9ObjectIdentifierC2Pnb304w1:
		ret = "C2PNB304W1"
	case X9ObjectIdentifierC2Tnb359v1:
		ret = "C2TNB359V1"
	case X9ObjectIdentifierC2Pnb368w1:
		ret = "C2PNB368W1"
	case X9ObjectIdentifierC2Tnb431r1:
		ret = "C2TNB431R1"
	case X9ObjectIdentifierPrimeCurve:
		ret = "PRIMECURVE"
	case X9ObjectIdentifierPrime192v1:
		ret = "PRIME192V1"
	case X9ObjectIdentifierPrime192v2:
		ret = "PRIME192V2"
	case X9ObjectIdentifierPrime192v3:
		ret = "PRIME192V3"
	case X9ObjectIdentifierPrime239v1:
		ret = "PRIME239V1"
	case X9ObjectIdentifierPrime239v2:
		ret = "PRIME239V2"
	case X9ObjectIdentifierPrime239v3:
		ret = "PRIME239V3"
	case X9ObjectIdentifierPrime256v1:
		ret = "PRIME256V1"
	case X9ObjectIdentifierIdDsa:
		ret = "IDDSA"
	case X9ObjectIdentifierIdDsaWithSha1:
		ret = "IDDSAWITHSHA1"
	case X9ObjectIdentifierX9x63Scheme:
		ret = "X9X63SCHEME"
	case X9ObjectIdentifierDHSinglePassStdDHSha1KdfScheme:
		ret = "DHSINGLEPASSSTDDHSHA1KDFSCHEME"
	case X9ObjectIdentifierDHSinglePassCofactorDHSha1KdfScheme:
		ret = "DHSINGLEPASSCOFACTORDHSHA1KDFSCHEME"
	case X9ObjectIdentifierMqvSinglePassSha1KdfScheme:
		ret = "MQVSINGLEPASSSHA1KDFSCHEME"
	case X9ObjectIdentifieransi_x9_42:
		ret = "ANSI_X9_42"
	case X9ObjectIdentifierDHPublicNumber:
		ret = "DHPUBLICNUMBER"
	case X9ObjectIdentifierX9x42Schemes:
		ret = "X9X42SCHEMES"
	case X9ObjectIdentifierDHStatic:
		ret = "DHSTATIC"
	case X9ObjectIdentifierDHEphem:
		ret = "DHEPHEM"
	case X9ObjectIdentifierDHOneFlow:
		ret = "DHONEFLOW"
	case X9ObjectIdentifierDHHybrid1:
		ret = "DHHYBRID1"
	case X9ObjectIdentifierDHHybrid2:
		ret = "DHHYBRID2"
	case X9ObjectIdentifierDHHybridOneFlow:
		ret = "DHHYBRIDONEFLOW"
	case X9ObjectIdentifierMqv2:
		ret = "MQV2"
	case X9ObjectIdentifierMqv1:
		ret = "MQV1"
	case X9ObjectIdentifierSecp384r1:
		ret = "SECP384R1"
	}
	return ret
}
