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

// Used Algorithm ID.
type AlgorithmId int

const (
	// AlgorithmIdAesGcm128 defines that the // AES-GCM-128
	AlgorithmIdAesGcm128 AlgorithmId = iota
	// AlgorithmIdAesGcm256 defines that the // AES-GCM-256
	AlgorithmIdAesGcm256
	// AlgorithmIdAesWrap128 defines that the // AES-WRAP-128
	AlgorithmIdAesWrap128
	// AlgorithmIdAesWrap256 defines that the // AES-WRAP-256.
	AlgorithmIdAesWrap256
)

// AlgorithmIdParse converts the given string into a AlgorithmId value.
//
// It returns the corresponding AlgorithmId constant if the string matches
// a known level name, or an error if the input is invalid.
func AlgorithmIdParse(value string) (AlgorithmId, error) {
	var ret AlgorithmId
	var err error
	switch strings.ToUpper(value) {
	case "AESGCM128":
		ret = AlgorithmIdAesGcm128
	case "AESGCM256":
		ret = AlgorithmIdAesGcm256
	case "AESWRAP128":
		ret = AlgorithmIdAesWrap128
	case "AESWRAP256":
		ret = AlgorithmIdAesWrap256
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the AlgorithmId.
// It satisfies fmt.Stringer.
func (g AlgorithmId) String() string {
	var ret string
	switch g {
	case AlgorithmIdAesGcm128:
		ret = "AESGCM128"
	case AlgorithmIdAesGcm256:
		ret = "AESGCM256"
	case AlgorithmIdAesWrap128:
		ret = "AESWRAP128"
	case AlgorithmIdAesWrap256:
		ret = "AESWRAP256"
	}
	return ret
}
