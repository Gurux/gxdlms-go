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

// HashAlgorithm s.
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
	switch strings.ToUpper(value) {
	case "NONE":
		ret = HashAlgorithmNone
	case "SHA1RSA":
		ret = HashAlgorithmSha1Rsa
	case "MD5RSA":
		ret = HashAlgorithmMd5Rsa
	case "SHA1DSA":
		ret = HashAlgorithmSha1Dsa
	case "SHA1RSA1":
		ret = HashAlgorithmSha1Rsa1
	case "SHARSA":
		ret = HashAlgorithmShaRsa
	case "MD5RSA1":
		ret = HashAlgorithmMd5Rsa1
	case "MD2RSA1":
		ret = HashAlgorithmMd2Rsa1
	case "MD4RSA":
		ret = HashAlgorithmMd4Rsa
	case "MD4RSA1":
		ret = HashAlgorithmMd4Rsa1
	case "MD4RSA2":
		ret = HashAlgorithmMd4Rsa2
	case "MD2RSA":
		ret = HashAlgorithmMd2Rsa
	case "SHA1DSA1":
		ret = HashAlgorithmSha1Dsa1
	case "DSASHA1":
		ret = HashAlgorithmDsaSha1
	case "MOSAICUPDATEDSIG":
		ret = HashAlgorithmMosaicUpdatedSig
	case "SHA1NOSIGN":
		ret = HashAlgorithmSha1NoSign
	case "MD5NOSIGN":
		ret = HashAlgorithmMd5NoSign
	case "SHA256NOSIGN":
		ret = HashAlgorithmSha256NoSign
	case "SHA384NOSIGN":
		ret = HashAlgorithmSha384NoSign
	case "SHA512NOSIGN":
		ret = HashAlgorithmSha512NoSign
	case "SHA256RSA":
		ret = HashAlgorithmSha256Rsa
	case "SHA384RSA":
		ret = HashAlgorithmSha384Rsa
	case "SHA512RSA":
		ret = HashAlgorithmSha512Rsa
	case "RSASSAPSS":
		ret = HashAlgorithmRsaSsaPss
	case "SHA1WITHECDSA":
		ret = HashAlgorithmSha1withecdsa
	case "SHA256WITHECDSA":
		ret = HashAlgorithmSha256WithEcdsa
	case "SHA384WITHECDSA":
		ret = HashAlgorithmSha384WithEcdsa
	case "SHA512WITHECDSA":
		ret = HashAlgorithmSha512WithEcdsa
	case "SPECIFIEDECDSA":
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
		ret = "NONE"
	case HashAlgorithmSha1Rsa:
		ret = "SHA1RSA"
	case HashAlgorithmMd5Rsa:
		ret = "MD5RSA"
	case HashAlgorithmSha1Dsa:
		ret = "SHA1DSA"
	case HashAlgorithmSha1Rsa1:
		ret = "SHA1RSA1"
	case HashAlgorithmShaRsa:
		ret = "SHARSA"
	case HashAlgorithmMd5Rsa1:
		ret = "MD5RSA1"
	case HashAlgorithmMd2Rsa1:
		ret = "MD2RSA1"
	case HashAlgorithmMd4Rsa:
		ret = "MD4RSA"
	case HashAlgorithmMd4Rsa1:
		ret = "MD4RSA1"
	case HashAlgorithmMd4Rsa2:
		ret = "MD4RSA2"
	case HashAlgorithmMd2Rsa:
		ret = "MD2RSA"
	case HashAlgorithmSha1Dsa1:
		ret = "SHA1DSA1"
	case HashAlgorithmDsaSha1:
		ret = "DSASHA1"
	case HashAlgorithmMosaicUpdatedSig:
		ret = "MOSAICUPDATEDSIG"
	case HashAlgorithmSha1NoSign:
		ret = "SHA1NOSIGN"
	case HashAlgorithmMd5NoSign:
		ret = "MD5NOSIGN"
	case HashAlgorithmSha256NoSign:
		ret = "SHA256NOSIGN"
	case HashAlgorithmSha384NoSign:
		ret = "SHA384NOSIGN"
	case HashAlgorithmSha512NoSign:
		ret = "SHA512NOSIGN"
	case HashAlgorithmSha256Rsa:
		ret = "SHA256RSA"
	case HashAlgorithmSha384Rsa:
		ret = "SHA384RSA"
	case HashAlgorithmSha512Rsa:
		ret = "SHA512RSA"
	case HashAlgorithmRsaSsaPss:
		ret = "RSASSAPSS"
	case HashAlgorithmSha1withecdsa:
		ret = "SHA1WITHECDSA"
	case HashAlgorithmSha256WithEcdsa:
		ret = "SHA256WITHECDSA"
	case HashAlgorithmSha384WithEcdsa:
		ret = "SHA384WITHECDSA"
	case HashAlgorithmSha512WithEcdsa:
		ret = "SHA512WITHECDSA"
	case HashAlgorithmSpecifiedEcdsa:
		ret = "SPECIFIEDECDSA"
	}
	return ret
}
