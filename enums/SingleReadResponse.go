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

// Enumerates single read response types.
type SingleReadResponse int

const (
	// SingleReadResponseData defines normal data reply.
	SingleReadResponseData SingleReadResponse = iota
	// SingleReadResponseDataAccessError defines that the an error has occured on read.
	SingleReadResponseDataAccessError
	// SingleReadResponseDataBlockResult defines return data as blocks.
	SingleReadResponseDataBlockResult
	// SingleReadResponseBlockNumber defines return block number.
	SingleReadResponseBlockNumber
)

// SingleReadResponseParse converts the given string into a SingleReadResponse value.
//
// It returns the corresponding SingleReadResponse constant if the string matches
// a known level name, or an error if the input is invalid.
func SingleReadResponseParse(value string) (SingleReadResponse, error) {
	var ret SingleReadResponse
	var err error
	switch strings.ToUpper(value) {
	case "DATA":
		ret = SingleReadResponseData
	case "DATAACCESSERROR":
		ret = SingleReadResponseDataAccessError
	case "DATABLOCKRESULT":
		ret = SingleReadResponseDataBlockResult
	case "BLOCKNUMBER":
		ret = SingleReadResponseBlockNumber
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the SingleReadResponse.
// It satisfies fmt.Stringer.
func (g SingleReadResponse) String() string {
	var ret string
	switch g {
	case SingleReadResponseData:
		ret = "DATA"
	case SingleReadResponseDataAccessError:
		ret = "DATAACCESSERROR"
	case SingleReadResponseDataBlockResult:
		ret = "DATABLOCKRESULT"
	case SingleReadResponseBlockNumber:
		ret = "BLOCKNUMBER"
	}
	return ret
}
