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

// HashAlgorithm enumerates hash algorithm.
type HashAlgorithm int

const (
	HashAlgorithmNone HashAlgorithm = iota
	HashAlgorithmSha1Rsa
	HashAlgorithmMd5Rsa
	HashAlgorithmSha1Dsa
	HashAlgorithmSha1Rsa1
	HashAlgorithmShaRsa
	HashAlgorithmMd5Rsa1
	HashAlgorithmMd2Rsa1
	HashAlgorithmMd4Rsa
	HashAlgorithmMd4Rsa1
	HashAlgorithmMd4Rsa2
	HashAlgorithmMd2Rsa
	HashAlgorithmSha1Dsa1
	HashAlgorithmDsaSha1
	HashAlgorithmMosaicUpdatedSig
	HashAlgorithmSha1NoSign
	HashAlgorithmMd5NoSign
	HashAlgorithmSha256NoSign
	HashAlgorithmSha384NoSign
	HashAlgorithmSha512NoSign
	HashAlgorithmSha256Rsa
	HashAlgorithmSha384Rsa
	HashAlgorithmSha512Rsa
	HashAlgorithmRsaSsaPss
	HashAlgorithmSha1withecdsa
	HashAlgorithmSha256WithEcdsa
	HashAlgorithmSha384WithEcdsa
	HashAlgorithmSha512WithEcdsa
	HashAlgorithmSpecifiedEcdsa
)

// HashAlgorithmParse converts the given string into a HashAlgorithm value.
//
// It returns the corresponding HashAlgorithm constant if the string matches
// a known level name, or an error if the input is invalid.
func HashAlgorithmParse(value string) (HashAlgorithm, error) {
	var ret HashAlgorithm
	var err error
	switch {
	case strings.EqualFold(value, "None"):
		ret = HashAlgorithmNone
	case strings.EqualFold(value, "Sha1Rsa"):
		ret = HashAlgorithmSha1Rsa
	case strings.EqualFold(value, "Md5Rsa"):
		ret = HashAlgorithmMd5Rsa
	case strings.EqualFold(value, "Sha1Dsa"):
		ret = HashAlgorithmSha1Dsa
	case strings.EqualFold(value, "Sha1Rsa1"):
		ret = HashAlgorithmSha1Rsa1
	case strings.EqualFold(value, "ShaRsa"):
		ret = HashAlgorithmShaRsa
	case strings.EqualFold(value, "Md5Rsa1"):
		ret = HashAlgorithmMd5Rsa1
	case strings.EqualFold(value, "Md2Rsa1"):
		ret = HashAlgorithmMd2Rsa1
	case strings.EqualFold(value, "Md4Rsa"):
		ret = HashAlgorithmMd4Rsa
	case strings.EqualFold(value, "Md4Rsa1"):
		ret = HashAlgorithmMd4Rsa1
	case strings.EqualFold(value, "Md4Rsa2"):
		ret = HashAlgorithmMd4Rsa2
	case strings.EqualFold(value, "Md2Rsa"):
		ret = HashAlgorithmMd2Rsa
	case strings.EqualFold(value, "Sha1Dsa1"):
		ret = HashAlgorithmSha1Dsa1
	case strings.EqualFold(value, "DsaSha1"):
		ret = HashAlgorithmDsaSha1
	case strings.EqualFold(value, "MosaicUpdatedSig"):
		ret = HashAlgorithmMosaicUpdatedSig
	case strings.EqualFold(value, "Sha1NoSign"):
		ret = HashAlgorithmSha1NoSign
	case strings.EqualFold(value, "Md5NoSign"):
		ret = HashAlgorithmMd5NoSign
	case strings.EqualFold(value, "Sha256NoSign"):
		ret = HashAlgorithmSha256NoSign
	case strings.EqualFold(value, "Sha384NoSign"):
		ret = HashAlgorithmSha384NoSign
	case strings.EqualFold(value, "Sha512NoSign"):
		ret = HashAlgorithmSha512NoSign
	case strings.EqualFold(value, "Sha256Rsa"):
		ret = HashAlgorithmSha256Rsa
	case strings.EqualFold(value, "Sha384Rsa"):
		ret = HashAlgorithmSha384Rsa
	case strings.EqualFold(value, "Sha512Rsa"):
		ret = HashAlgorithmSha512Rsa
	case strings.EqualFold(value, "RsaSsaPss"):
		ret = HashAlgorithmRsaSsaPss
	case strings.EqualFold(value, "Sha1withecdsa"):
		ret = HashAlgorithmSha1withecdsa
	case strings.EqualFold(value, "Sha256WithEcdsa"):
		ret = HashAlgorithmSha256WithEcdsa
	case strings.EqualFold(value, "Sha384WithEcdsa"):
		ret = HashAlgorithmSha384WithEcdsa
	case strings.EqualFold(value, "Sha512WithEcdsa"):
		ret = HashAlgorithmSha512WithEcdsa
	case strings.EqualFold(value, "SpecifiedEcdsa"):
		ret = HashAlgorithmSpecifiedEcdsa
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the HashAlgorithm.
// It satisfies fmt.Stringer.
func (g HashAlgorithm) String() string {
	var ret string
	switch g {
	case HashAlgorithmNone:
		ret = "None"
	case HashAlgorithmSha1Rsa:
		ret = "Sha1Rsa"
	case HashAlgorithmMd5Rsa:
		ret = "Md5Rsa"
	case HashAlgorithmSha1Dsa:
		ret = "Sha1Dsa"
	case HashAlgorithmSha1Rsa1:
		ret = "Sha1Rsa1"
	case HashAlgorithmShaRsa:
		ret = "ShaRsa"
	case HashAlgorithmMd5Rsa1:
		ret = "Md5Rsa1"
	case HashAlgorithmMd2Rsa1:
		ret = "Md2Rsa1"
	case HashAlgorithmMd4Rsa:
		ret = "Md4Rsa"
	case HashAlgorithmMd4Rsa1:
		ret = "Md4Rsa1"
	case HashAlgorithmMd4Rsa2:
		ret = "Md4Rsa2"
	case HashAlgorithmMd2Rsa:
		ret = "Md2Rsa"
	case HashAlgorithmSha1Dsa1:
		ret = "Sha1Dsa1"
	case HashAlgorithmDsaSha1:
		ret = "DsaSha1"
	case HashAlgorithmMosaicUpdatedSig:
		ret = "MosaicUpdatedSig"
	case HashAlgorithmSha1NoSign:
		ret = "Sha1NoSign"
	case HashAlgorithmMd5NoSign:
		ret = "Md5NoSign"
	case HashAlgorithmSha256NoSign:
		ret = "Sha256NoSign"
	case HashAlgorithmSha384NoSign:
		ret = "Sha384NoSign"
	case HashAlgorithmSha512NoSign:
		ret = "Sha512NoSign"
	case HashAlgorithmSha256Rsa:
		ret = "Sha256Rsa"
	case HashAlgorithmSha384Rsa:
		ret = "Sha384Rsa"
	case HashAlgorithmSha512Rsa:
		ret = "Sha512Rsa"
	case HashAlgorithmRsaSsaPss:
		ret = "RsaSsaPss"
	case HashAlgorithmSha1withecdsa:
		ret = "Sha1withecdsa"
	case HashAlgorithmSha256WithEcdsa:
		ret = "Sha256WithEcdsa"
	case HashAlgorithmSha384WithEcdsa:
		ret = "Sha384WithEcdsa"
	case HashAlgorithmSha512WithEcdsa:
		ret = "Sha512WithEcdsa"
	case HashAlgorithmSpecifiedEcdsa:
		ret = "SpecifiedEcdsa"
	}
	return ret
}

// AllHashAlgorithm returns a slice containing all defined HashAlgorithm values.
func AllHashAlgorithm() []HashAlgorithm {
	return []HashAlgorithm{
		HashAlgorithmNone,
		HashAlgorithmSha1Rsa,
		HashAlgorithmMd5Rsa,
		HashAlgorithmSha1Dsa,
		HashAlgorithmSha1Rsa1,
		HashAlgorithmShaRsa,
		HashAlgorithmMd5Rsa1,
		HashAlgorithmMd2Rsa1,
		HashAlgorithmMd4Rsa,
		HashAlgorithmMd4Rsa1,
		HashAlgorithmMd4Rsa2,
		HashAlgorithmMd2Rsa,
		HashAlgorithmSha1Dsa1,
		HashAlgorithmDsaSha1,
		HashAlgorithmMosaicUpdatedSig,
		HashAlgorithmSha1NoSign,
		HashAlgorithmMd5NoSign,
		HashAlgorithmSha256NoSign,
		HashAlgorithmSha384NoSign,
		HashAlgorithmSha512NoSign,
		HashAlgorithmSha256Rsa,
		HashAlgorithmSha384Rsa,
		HashAlgorithmSha512Rsa,
		HashAlgorithmRsaSsaPss,
		HashAlgorithmSha1withecdsa,
		HashAlgorithmSha256WithEcdsa,
		HashAlgorithmSha384WithEcdsa,
		HashAlgorithmSha512WithEcdsa,
		HashAlgorithmSpecifiedEcdsa,
	}
}
