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

// Used ECC scheme.
type Ecc int

const (
	// EccP256 defines that the // ECC-P256 domain parameters are used.
	EccP256 Ecc = iota
	// EccP384 defines that the // ECC-384 domain parameters are used.
	EccP384
)

// EccParse converts the given string into a Ecc value.
//
// It returns the corresponding Ecc constant if the string matches
// a known level name, or an error if the input is invalid.
func EccParse(value string) (Ecc, error) {
	var ret Ecc
	var err error
	switch {
	case strings.EqualFold(value, "P256"):
		ret = EccP256
	case strings.EqualFold(value, "P384"):
		ret = EccP384
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the Ecc.
// It satisfies fmt.Stringer.
func (g Ecc) String() string {
	var ret string
	switch g {
	case EccP256:
		ret = "P256"
	case EccP384:
		ret = "P384"
	}
	return ret
}

// AllEcc returns a slice containing all defined Ecc values.
func AllEcc() []Ecc {
	return []Ecc{
	EccP256,
	EccP384,
	}
}
