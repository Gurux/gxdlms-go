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

// AlgorithmID enumerates Available algorithms.
type AlgorithmID int

const (
	// AlgorithmIDAesGcm128 defines that the // AES-GCM-128
	AlgorithmIDAesGcm128 AlgorithmID = iota
	// AlgorithmIDAesGcm256 defines that the // AES-GCM-256
	AlgorithmIDAesGcm256
	// AlgorithmIDAesWrap128 defines that the // AES-WRAP-128
	AlgorithmIDAesWrap128
	// AlgorithmIDAesWrap256 defines that the // AES-WRAP-256.
	AlgorithmIDAesWrap256
)

// AlgorithmIDParse converts the given string into a AlgorithmId value.
//
// It returns the corresponding AlgorithmId constant if the string matches
// a known level name, or an error if the input is invalid.
func AlgorithmIDParse(value string) (AlgorithmID, error) {
	var ret AlgorithmID
	var err error
	switch {
	case strings.EqualFold(value, "AesGcm128"):
		ret = AlgorithmIDAesGcm128
	case strings.EqualFold(value, "AesGcm256"):
		ret = AlgorithmIDAesGcm256
	case strings.EqualFold(value, "AesWrap128"):
		ret = AlgorithmIDAesWrap128
	case strings.EqualFold(value, "AesWrap256"):
		ret = AlgorithmIDAesWrap256
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the AlgorithmID.
// It satisfies fmt.Stringer.
func (g AlgorithmID) String() string {
	var ret string
	switch g {
	case AlgorithmIDAesGcm128:
		ret = "AesGcm128"
	case AlgorithmIDAesGcm256:
		ret = "AesGcm256"
	case AlgorithmIDAesWrap128:
		ret = "AesWrap128"
	case AlgorithmIDAesWrap256:
		ret = "AesWrap256"
	}
	return ret
}

// AllAlgorithmID returns a slice containing all defined AlgorithmID values.
func AllAlgorithmID() []AlgorithmID {
	return []AlgorithmID{
		AlgorithmIDAesGcm128,
		AlgorithmIDAesGcm256,
		AlgorithmIDAesWrap128,
		AlgorithmIDAesWrap256,
	}
}
