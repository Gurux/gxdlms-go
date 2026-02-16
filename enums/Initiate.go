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

// Initiate :  describes onitiate errors.
type Initiate int

const (
	//InitiateOther defines that other error has occurred.
	InitiateOther Initiate = iota
	// InitiateDlmsVersionTooLow defines that the DLMS version is too low.
	InitiateDlmsVersionTooLow
	// InitiateIncompatibleConformance defines that the incompatible conformance.
	InitiateIncompatibleConformance
	// InitiatePduSizeTooShort defines that the PDU size is too short.
	InitiatePduSizeTooShort
	// InitiateRefusedByTheVDEHandler defines that the refused by the VDE handler.
	InitiateRefusedByTheVDEHandler
)

// InitiateParse converts the given string into a Initiate value.
//
// It returns the corresponding Initiate constant if the string matches
// a known level name, or an error if the input is invalid.
func InitiateParse(value string) (Initiate, error) {
	var ret Initiate
	var err error
	switch strings.ToUpper(value) {
	case "OTHER":
		ret = InitiateOther
	case "DLMSVERSIONTOOLOW":
		ret = InitiateDlmsVersionTooLow
	case "INCOMPATIBLECONFORMANCE":
		ret = InitiateIncompatibleConformance
	case "PDUSIZETOOSHORT":
		ret = InitiatePduSizeTooShort
	case "REFUSEDBYTHEVDEHANDLER":
		ret = InitiateRefusedByTheVDEHandler
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the Initiate.
// It satisfies fmt.Stringer.
func (g Initiate) String() string {
	var ret string
	switch g {
	case InitiateOther:
		ret = "OTHER"
	case InitiateDlmsVersionTooLow:
		ret = "DLMSVERSIONTOOLOW"
	case InitiateIncompatibleConformance:
		ret = "INCOMPATIBLECONFORMANCE"
	case InitiatePduSizeTooShort:
		ret = "PDUSIZETOOSHORT"
	case InitiateRefusedByTheVDEHandler:
		ret = "REFUSEDBYTHEVDEHANDLER"
	}
	return ret
}
