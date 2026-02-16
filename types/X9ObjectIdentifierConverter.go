package types

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

	"github.com/Gurux/gxdlms-go/enums"
)

// X9ObjectIdentifierToString converts the given X9ObjectIdentifier value into a string.
func X9ObjectIdentifierToString(value enums.X9ObjectIdentifier) (string, error) {
	switch value {
	case enums.X9ObjectIdentifierIdFieldType:
		return "1.2.840.10045.1", nil
	case enums.X9ObjectIdentifierPrimeField:
		return "1.2.840.10045.1", nil
	case enums.X9ObjectIdentifierCharacteristicTwoField:
		return "1.2.840.10045.1.2", nil
	case enums.X9ObjectIdentifierGNBasis:
		return "1.2.840.10045.1.2.3.1", nil
	case enums.X9ObjectIdentifierTPBasis:
		return "1.2.840.10045.1.2.3.2", nil
	case enums.X9ObjectIdentifierPPBasis:
		return "1.2.840.10045.1.2.3.3", nil
	case enums.X9ObjectIdentifierECDsaWithSha1:
		return "1.2.840.10045.4.1", nil
	case enums.X9ObjectIdentifierIdECPublicKey:
		return "1.2.840.10045.2.1", nil
	case enums.X9ObjectIdentifierECDsaWithSha2:
		return "1.2.840.10045.4.3", nil
	case enums.X9ObjectIdentifierECDsaWithSha224:
		return "1.2.840.10045.4.31", nil
	case enums.X9ObjectIdentifierECDsaWithSha256:
		return "1.2.840.10045.4.32", nil
	case enums.X9ObjectIdentifierECDsaWithSha384:
		return "1.2.840.10045.4.33", nil
	case enums.X9ObjectIdentifierECDsaWithSha512:
		return "1.2.840.10045.4.34", nil
	case enums.X9ObjectIdentifierEllipticCurve:
		return "1.2.840.10045.3", nil
	case enums.X9ObjectIdentifierCTwoCurve:
		return "1.2.840.10045.3.0", nil
	case enums.X9ObjectIdentifierC2Pnb163v1:
		return "1.2.840.10045.3.0.1", nil
	case enums.X9ObjectIdentifierC2Pnb163v2:
		return "1.2.840.10045.3.0.2", nil
	case enums.X9ObjectIdentifierC2Pnb163v3:
		return "1.2.840.10045.3.0.3", nil
	case enums.X9ObjectIdentifierC2Pnb176w1:
		return "1.2.840.10045.3.0.4", nil
	case enums.X9ObjectIdentifierC2Tnb191v1:
		return "1.2.840.10045.3.0.5", nil
	case enums.X9ObjectIdentifierC2Tnb191v2:
		return "1.2.840.10045.3.0.6", nil
	case enums.X9ObjectIdentifierC2Tnb191v3:
		return "1.2.840.10045.3.0.7", nil
	case enums.X9ObjectIdentifierC2Onb191v4:
		return "1.2.840.10045.3.0.8", nil
	case enums.X9ObjectIdentifierC2Onb191v5:
		return "1.2.840.10045.3.0.9", nil
	case enums.X9ObjectIdentifierC2Pnb208w1:
		return "1.2.840.10045.3.0.10", nil
	case enums.X9ObjectIdentifierC2Tnb239v1:
		return "1.2.840.10045.3.0.11", nil
	case enums.X9ObjectIdentifierC2Tnb239v2:
		return "1.2.840.10045.3.0.12", nil
	case enums.X9ObjectIdentifierC2Tnb239v3:
		return "1.2.840.10045.3.0.13", nil
	case enums.X9ObjectIdentifierC2Onb239v4:
		return "1.2.840.10045.3.0.14", nil
	case enums.X9ObjectIdentifierC2Onb239v5:
		return "1.2.840.10045.3.0.15", nil
	case enums.X9ObjectIdentifierC2Pnb272w1:
		return "1.2.840.10045.3.0.16", nil
	case enums.X9ObjectIdentifierC2Pnb304w1:
		return "1.2.840.10045.3.0.17", nil
	case enums.X9ObjectIdentifierC2Tnb359v1:
		return "1.2.840.10045.3.0.18", nil
	case enums.X9ObjectIdentifierC2Pnb368w1:
		return "1.2.840.10045.3.0.19", nil
	case enums.X9ObjectIdentifierC2Tnb431r1:
		return "1.2.840.10045.3.0.20", nil
	case enums.X9ObjectIdentifierPrimeCurve:
		return "1.2.840.10045.3.1", nil
	case enums.X9ObjectIdentifierPrime192v1:
		return "1.2.840.10045.3.1.1", nil
	case enums.X9ObjectIdentifierPrime192v2:
		return "1.2.840.10045.3.1.2", nil
	case enums.X9ObjectIdentifierPrime192v3:
		return "1.2.840.10045.3.1.3", nil
	case enums.X9ObjectIdentifierPrime239v1:
		return "1.2.840.10045.3.1.4", nil
	case enums.X9ObjectIdentifierPrime239v2:
		return "1.2.840.10045.3.1.5", nil
	case enums.X9ObjectIdentifierPrime239v3:
		return "1.2.840.10045.3.1.6", nil
	case enums.X9ObjectIdentifierPrime256v1:
		return "1.2.840.10045.3.1.7", nil
	case enums.X9ObjectIdentifierIdDsa:
		return "1.2.840.10040.4.1", nil
	case enums.X9ObjectIdentifierIdDsaWithSha1:
		return "1.2.840.10040.4.3", nil
	case enums.X9ObjectIdentifierX9x63Scheme:
		return "1.3.133.16.840.63.0", nil
	case enums.X9ObjectIdentifierDHSinglePassStdDHSha1KdfScheme:
		return "1.3.133.16.840.63.0.2", nil
	case enums.X9ObjectIdentifierDHSinglePassCofactorDHSha1KdfScheme:
		return "1.3.133.16.840.63.0.3", nil
	case enums.X9ObjectIdentifierMqvSinglePassSha1KdfScheme:
		return "1.3.133.16.840.63.0.16", nil
	case enums.X9ObjectIdentifieransi_x9_42:
		return "1.2.840.10046", nil
	case enums.X9ObjectIdentifierDHPublicNumber:
		return "1.2.840.10046.2.1", nil
	case enums.X9ObjectIdentifierX9x42Schemes:
		return "1.2.840.10046.2.3", nil
	case enums.X9ObjectIdentifierDHStatic:
		return "1.2.840.10046.2.3.1", nil
	case enums.X9ObjectIdentifierDHEphem:
		return "1.2.840.10046.2.3.2", nil
	case enums.X9ObjectIdentifierDHOneFlow:
		return "1.2.840.10046.2.3.3", nil
	case enums.X9ObjectIdentifierDHHybrid1:
		return "1.2.840.10046.2.3.4", nil
	case enums.X9ObjectIdentifierDHHybrid2:
		return "1.2.840.10046.2.3.5", nil
	case enums.X9ObjectIdentifierDHHybridOneFlow:
		return "1.2.840.10046.2.3.6", nil
	case enums.X9ObjectIdentifierMqv2:
		return "1.2.840.10046.2.3.7", nil
	case enums.X9ObjectIdentifierMqv1:
		return "1.2.840.10046.2.3.8", nil
	case enums.X9ObjectIdentifierSecp384r1:
		return "1.3.132.0.34", nil
	default:
		return "", fmt.Errorf("Invalid X509Name. %s ", value)
	}
}

// X9ObjectIdentifierFromString converts the given string value into an X9ObjectIdentifier enum.
func X9ObjectIdentifierFromString(value string) enums.X9ObjectIdentifier {
	if value == "1.2.840.10045.1" {
		return enums.X9ObjectIdentifierIdFieldType
	}
	if value == "1.2.840.10045.1" {
		return enums.X9ObjectIdentifierPrimeField
	}
	if value == "1.2.840.10045.1.2" {
		return enums.X9ObjectIdentifierCharacteristicTwoField
	}
	if value == "1.2.840.10045.1.2.3.1" {
		return enums.X9ObjectIdentifierGNBasis
	}
	if value == "1.2.840.10045.1.2.3.2" {
		return enums.X9ObjectIdentifierTPBasis
	}
	if value == "1.2.840.10045.1.2.3.3" {
		return enums.X9ObjectIdentifierPPBasis
	}
	if value == "1.2.840.10045.4.1" {
		return enums.X9ObjectIdentifierECDsaWithSha1
	}
	if value == "1.2.840.10045.2.1" {
		return enums.X9ObjectIdentifierIdECPublicKey
	}
	if value == "1.2.840.10045.4.3" {
		return enums.X9ObjectIdentifierECDsaWithSha2
	}
	if value == "1.2.840.10045.4.31" {
		return enums.X9ObjectIdentifierECDsaWithSha224
	}
	if value == "1.2.840.10045.4.32" {
		return enums.X9ObjectIdentifierECDsaWithSha256
	}
	if value == "1.2.840.10045.4.33" {
		return enums.X9ObjectIdentifierECDsaWithSha384
	}
	if value == "1.2.840.10045.4.34" {
		return enums.X9ObjectIdentifierECDsaWithSha512
	}
	if value == "1.2.840.10045.3" {
		return enums.X9ObjectIdentifierEllipticCurve
	}
	if value == "1.2.840.10045.3.0" {
		return enums.X9ObjectIdentifierCTwoCurve
	}
	if value == "1.2.840.10045.3.0.1" {
		return enums.X9ObjectIdentifierC2Pnb163v1
	}
	if value == "1.2.840.10045.3.0.2" {
		return enums.X9ObjectIdentifierC2Pnb163v2
	}
	if value == "1.2.840.10045.3.0.3" {
		return enums.X9ObjectIdentifierC2Pnb163v3
	}
	if value == "1.2.840.10045.3.0.4" {
		return enums.X9ObjectIdentifierC2Pnb176w1
	}
	if value == "1.2.840.10045.3.0.5" {
		return enums.X9ObjectIdentifierC2Tnb191v1
	}
	if value == "1.2.840.10045.3.0.6" {
		return enums.X9ObjectIdentifierC2Tnb191v2
	}
	if value == "1.2.840.10045.3.0.7" {
		return enums.X9ObjectIdentifierC2Tnb191v3
	}
	if value == "1.2.840.10045.3.0.8" {
		return enums.X9ObjectIdentifierC2Onb191v4
	}
	if value == "1.2.840.10045.3.0.9" {
		return enums.X9ObjectIdentifierC2Onb191v5
	}
	if value == "1.2.840.10045.3.0.10" {
		return enums.X9ObjectIdentifierC2Pnb208w1
	}
	if value == "1.2.840.10045.3.0.11" {
		return enums.X9ObjectIdentifierC2Tnb239v1
	}
	if value == "1.2.840.10045.3.0.12" {
		return enums.X9ObjectIdentifierC2Tnb239v2
	}
	if value == "1.2.840.10045.3.0.13" {
		return enums.X9ObjectIdentifierC2Tnb239v3
	}
	if value == "1.2.840.10045.3.0.14" {
		return enums.X9ObjectIdentifierC2Onb239v4
	}
	if value == "1.2.840.10045.3.0.15" {
		return enums.X9ObjectIdentifierC2Onb239v5
	}
	if value == "1.2.840.10045.3.0.16" {
		return enums.X9ObjectIdentifierC2Pnb272w1
	}
	if value == "1.2.840.10045.3.0.17" {
		return enums.X9ObjectIdentifierC2Pnb304w1
	}
	if value == "1.2.840.10045.3.0.18" {
		return enums.X9ObjectIdentifierC2Tnb359v1
	}
	if value == "1.2.840.10045.3.0.19" {
		return enums.X9ObjectIdentifierC2Pnb368w1
	}
	if value == "1.2.840.10045.3.0.20" {
		return enums.X9ObjectIdentifierC2Tnb431r1
	}
	if value == "1.2.840.10045.3.1" {
		return enums.X9ObjectIdentifierPrimeCurve
	}
	if value == "1.2.840.10045.3.1.1" {
		return enums.X9ObjectIdentifierPrime192v1
	}
	if value == "1.2.840.10045.3.1.2" {
		return enums.X9ObjectIdentifierPrime192v2
	}
	if value == "1.2.840.10045.3.1.3" {
		return enums.X9ObjectIdentifierPrime192v3
	}
	if value == "1.2.840.10045.3.1.4" {
		return enums.X9ObjectIdentifierPrime239v1
	}
	if value == "1.2.840.10045.3.1.5" {
		return enums.X9ObjectIdentifierPrime239v2
	}
	if value == "1.2.840.10045.3.1.6" {
		return enums.X9ObjectIdentifierPrime239v3
	}
	if value == "1.2.840.10045.3.1.7" {
		return enums.X9ObjectIdentifierPrime256v1
	}
	if value == "1.2.840.10040.4.1" {
		return enums.X9ObjectIdentifierIdDsa
	}
	if value == "1.2.840.10040.4.3" {
		return enums.X9ObjectIdentifierIdDsaWithSha1
	}
	if value == "1.3.133.16.840.63.0" {
		return enums.X9ObjectIdentifierX9x63Scheme
	}
	if value == "1.3.133.16.840.63.0.2" {
		return enums.X9ObjectIdentifierDHSinglePassStdDHSha1KdfScheme
	}
	if value == "1.3.133.16.840.63.0.3" {
		return enums.X9ObjectIdentifierDHSinglePassCofactorDHSha1KdfScheme
	}
	if value == "1.3.133.16.840.63.0.16" {
		return enums.X9ObjectIdentifierMqvSinglePassSha1KdfScheme
	}
	if value == "1.2.840.10046" {
		return enums.X9ObjectIdentifieransi_x9_42
	}
	if value == "1.2.840.10046.2.1" {
		return enums.X9ObjectIdentifierDHPublicNumber
	}
	if value == "1.2.840.10046.2.3" {
		return enums.X9ObjectIdentifierX9x42Schemes
	}
	if value == "1.2.840.10046.2.3.1" {
		return enums.X9ObjectIdentifierDHStatic
	}
	if value == "1.2.840.10046.2.3.2" {
		return enums.X9ObjectIdentifierDHEphem
	}
	if value == "1.2.840.10046.2.3.3" {
		return enums.X9ObjectIdentifierDHOneFlow
	}
	if value == "1.2.840.10046.2.3.4" {
		return enums.X9ObjectIdentifierDHHybrid1
	}
	if value == "1.2.840.10046.2.3.5" {
		return enums.X9ObjectIdentifierDHHybrid2
	}
	if value == "1.2.840.10046.2.3.6" {
		return enums.X9ObjectIdentifierDHHybridOneFlow
	}
	if value == "1.2.840.10046.2.3.7" {
		return enums.X9ObjectIdentifierMqv2
	}
	if value == "1.2.840.10046.2.3.8" {
		return enums.X9ObjectIdentifierMqv1
	}
	if value == "1.3.132.0.34" {
		return enums.X9ObjectIdentifierSecp384r1
	}
	return enums.X9ObjectIdentifierNone
}
