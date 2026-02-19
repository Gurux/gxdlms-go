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

// Enumerates single read response types.
type SingleReadResponse int

const (
	// SingleReadResponseData defines that the // Normal data.
	SingleReadResponseData SingleReadResponse = iota
	// SingleReadResponseDataAccessError defines that the // Error has occured on read.
	SingleReadResponseDataAccessError
	// SingleReadResponseDataBlockResult defines that the // Return data as blocks.
	SingleReadResponseDataBlockResult
	// SingleReadResponseBlockNumber defines that the // Return block number.
	SingleReadResponseBlockNumber
)

// SingleReadResponseParse converts the given string into a SingleReadResponse value.
//
// It returns the corresponding SingleReadResponse constant if the string matches
// a known level name, or an error if the input is invalid.
func SingleReadResponseParse(value string) (SingleReadResponse, error) {
	var ret SingleReadResponse
	var err error
	switch {
	case strings.EqualFold(value, "Data"):
		ret = SingleReadResponseData
	case strings.EqualFold(value, "DataAccessError"):
		ret = SingleReadResponseDataAccessError
	case strings.EqualFold(value, "DataBlockResult"):
		ret = SingleReadResponseDataBlockResult
	case strings.EqualFold(value, "BlockNumber"):
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
		ret = "Data"
	case SingleReadResponseDataAccessError:
		ret = "DataAccessError"
	case SingleReadResponseDataBlockResult:
		ret = "DataBlockResult"
	case SingleReadResponseBlockNumber:
		ret = "BlockNumber"
	}
	return ret
}

// AllSingleReadResponse returns a slice containing all defined SingleReadResponse values.
func AllSingleReadResponse() []SingleReadResponse {
	return []SingleReadResponse{
		SingleReadResponseData,
		SingleReadResponseDataAccessError,
		SingleReadResponseDataBlockResult,
		SingleReadResponseBlockNumber,
	}
}
