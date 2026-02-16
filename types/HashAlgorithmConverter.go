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

// HashAlgorithmToString converts the given HashAlgorithm value into a string.
func HashAlgorithmToString(value enums.HashAlgorithm) (string, error) {
	switch value {
	case enums.HashAlgorithmSha1Rsa:
		return "1.2.840.113549.1.1.5", nil
	case enums.HashAlgorithmMd5Rsa:
		return "1.2.840.113549.1.1.4", nil
	case enums.HashAlgorithmSha1Dsa:
		return "1.2.840.10040.4.3", nil
	case enums.HashAlgorithmSha1Rsa1:
		return "1.3.14.3.2.29", nil
	case enums.HashAlgorithmShaRsa:
		return "1.3.14.3.2.15", nil
	case enums.HashAlgorithmMd5Rsa1:
		return "1.3.14.3.2.3", nil
	case enums.HashAlgorithmMd2Rsa1:
		return "1.2.840.113549.1.1.2", nil
	case enums.HashAlgorithmMd4Rsa:
		return "1.2.840.113549.1.1.3", nil
	case enums.HashAlgorithmMd4Rsa1:
		return "1.3.14.3.2.2", nil
	case enums.HashAlgorithmMd4Rsa2:
		return "1.3.14.3.2.4", nil
	case enums.HashAlgorithmMd2Rsa:
		return "1.3.14.7.2.3.1", nil
	case enums.HashAlgorithmSha1Dsa1:
		return "1.3.14.3.2.13", nil
	case enums.HashAlgorithmDsaSha1:
		return "1.3.14.3.2.27", nil
	case enums.HashAlgorithmMosaicUpdatedSig:
		return "2.16.840.1.101.2.1.1.19", nil
	case enums.HashAlgorithmSha1NoSign:
		return "1.3.14.3.2.26", nil
	case enums.HashAlgorithmMd5NoSign:
		return "1.2.840.113549.2.5", nil
	case enums.HashAlgorithmSha256NoSign:
		return "2.16.840.1.101.3.4.2.1", nil
	case enums.HashAlgorithmSha384NoSign:
		return "2.16.840.1.101.3.4.2.2", nil
	case enums.HashAlgorithmSha512NoSign:
		return "2.16.840.1.101.3.4.2.3", nil
	case enums.HashAlgorithmSha256Rsa:
		return "1.2.840.113549.1.1.11", nil
	case enums.HashAlgorithmSha384Rsa:
		return "1.2.840.113549.1.1.12", nil
	case enums.HashAlgorithmSha512Rsa:
		return "1.2.840.113549.1.1.13", nil
	case enums.HashAlgorithmRsaSsaPss:
		return "1.2.840.113549.1.1.10", nil
	case enums.HashAlgorithmSha1withecdsa:
		return "1.2.840.10045.4.1", nil
	case enums.HashAlgorithmSha256WithEcdsa:
		return "1.2.840.10045.4.3.2", nil
	case enums.HashAlgorithmSha384WithEcdsa:
		return "1.2.840.10045.4.3.3", nil
	case enums.HashAlgorithmSha512WithEcdsa:
		return "1.2.840.10045.4.3.4", nil
	case enums.HashAlgorithmSpecifiedEcdsa:
		return "1.2.840.10045.4.3", nil
	default:
		return "", fmt.Errorf("Invalid HashAlgorithm. %s", value)
	}
}

// HashAlgorithmFromString converts the given string into a HashAlgorithm value.
func HashAlgorithmFromString(value string) enums.HashAlgorithm {
	switch value {
	case "1.2.840.113549.1.1.5":
		return enums.HashAlgorithmSha1Rsa
	case "1.2.840.113549.1.1.4":
		return enums.HashAlgorithmMd5Rsa
	case "1.2.840.10040.4.3":
		return enums.HashAlgorithmSha1Dsa
	case "1.3.14.3.2.29":
		return enums.HashAlgorithmSha1Rsa1
	case "1.3.14.3.2.15":
		return enums.HashAlgorithmShaRsa
	case "1.3.14.3.2.3":
		return enums.HashAlgorithmMd5Rsa1
	case "1.2.840.113549.1.1.2":
		return enums.HashAlgorithmMd2Rsa1
	case "1.2.840.113549.1.1.3":
		return enums.HashAlgorithmMd4Rsa
	case "1.3.14.3.2.2":
		return enums.HashAlgorithmMd4Rsa1
	case "1.3.14.3.2.4":
		return enums.HashAlgorithmMd4Rsa2
	case "1.3.14.7.2.3.1":
		return enums.HashAlgorithmMd2Rsa
	case "1.3.14.3.2.13":
		return enums.HashAlgorithmSha1Dsa1
	case "1.3.14.3.2.27":
		return enums.HashAlgorithmDsaSha1
	case "2.16.840.1.101.2.1.1.19":
		return enums.HashAlgorithmMosaicUpdatedSig
	case "1.3.14.3.2.26":
		return enums.HashAlgorithmSha1NoSign
	case "1.2.840.113549.2.5":
		return enums.HashAlgorithmMd5NoSign
	case "2.16.840.1.101.3.4.2.1":
		return enums.HashAlgorithmSha256NoSign
	case "2.16.840.1.101.3.4.2.2":
		return enums.HashAlgorithmSha384NoSign
	case "2.16.840.1.101.3.4.2.3":
		return enums.HashAlgorithmSha512NoSign
	case "1.2.840.113549.1.1.11":
		return enums.HashAlgorithmSha256Rsa
	case "1.2.840.113549.1.1.12":
		return enums.HashAlgorithmSha384Rsa
	case "1.2.840.113549.1.1.13":
		return enums.HashAlgorithmSha512Rsa
	case "1.2.840.113549.1.1.10":
		return enums.HashAlgorithmRsaSsaPss
	case "1.2.840.10045.4.1":
		return enums.HashAlgorithmSha1withecdsa
	case "1.2.840.10045.4.3.2":
		return enums.HashAlgorithmSha256WithEcdsa
	case "1.2.840.10045.4.3.3":
		return enums.HashAlgorithmSha384WithEcdsa
	case "1.2.840.10045.4.3.4":
		return enums.HashAlgorithmSha512WithEcdsa
	case "1.2.840.10045.4.3":
		return enums.HashAlgorithmSpecifiedEcdsa
	default:
		return enums.HashAlgorithmNone
	}
}
