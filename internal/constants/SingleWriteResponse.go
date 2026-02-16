package constants

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

// SingleWriteResponse Enumerates single write response types.
type SingleWriteResponse int

const (
	// SingleWriteResponseSuccess defines that the write succeeded.
	SingleWriteResponseSuccess SingleWriteResponse = iota
	// SingleWriteResponseDataAccessError defines that the write error has occurred.
	SingleWriteResponseDataAccessError
	// SingleWriteResponseBlockNumber defines that the get next block.
	SingleWriteResponseBlockNumber
)

// SingleWriteResponseParse converts the given string into a SingleWriteResponse value.
//
// It returns the corresponding SingleWriteResponse constant if the string matches
// a known level name, or an error if the input is invalid.
func SingleWriteResponseParse(value string) (SingleWriteResponse, error) {
	var ret SingleWriteResponse
	var err error
	switch strings.ToUpper(value) {
	case "SUCCESS":
		ret = SingleWriteResponseSuccess
	case "DATAACCESSERROR":
		ret = SingleWriteResponseDataAccessError
	case "BLOCKNUMBER":
		ret = SingleWriteResponseBlockNumber
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the SingleWriteResponse.
// It satisfies fmt.Stringer.
func (g SingleWriteResponse) String() string {
	var ret string
	switch g {
	case SingleWriteResponseSuccess:
		ret = "SUCCESS"
	case SingleWriteResponseDataAccessError:
		ret = "DATAACCESSERROR"
	case SingleWriteResponseBlockNumber:
		ret = "BLOCKNUMBER"
	}
	return ret
}
