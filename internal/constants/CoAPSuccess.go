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

// CoAP success enumerates.
type CoAPSuccess int

const (
	// CoAPSuccessNone defines that the // Empty success code.
	CoAPSuccessNone CoAPSuccess = iota
	// CoAPSuccessCreated defines that the // Created.
	CoAPSuccessCreated
	// CoAPSuccessDeleted defines that the // Deleted.
	CoAPSuccessDeleted
	// CoAPSuccessValid defines that the // Valid.
	CoAPSuccessValid
	// CoAPSuccessChanged defines that the // Changed.
	CoAPSuccessChanged
	// CoAPSuccessContent defines that the // Content.
	CoAPSuccessContent
	// CoAPSuccessContinue defines that the // Continue.
	CoAPSuccessContinue CoAPSuccess = 31
)

// CoAPSuccessParse converts the given string into a CoAPSuccess value.
//
// It returns the corresponding CoAPSuccess constant if the string matches
// a known level name, or an error if the input is invalid.
func CoAPSuccessParse(value string) (CoAPSuccess, error) {
	var ret CoAPSuccess
	var err error
	switch strings.ToUpper(value) {
	case "NONE":
		ret = CoAPSuccessNone
	case "CREATED":
		ret = CoAPSuccessCreated
	case "DELETED":
		ret = CoAPSuccessDeleted
	case "VALID":
		ret = CoAPSuccessValid
	case "CHANGED":
		ret = CoAPSuccessChanged
	case "CONTENT":
		ret = CoAPSuccessContent
	case "CONTINUE":
		ret = CoAPSuccessContinue
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the CoAPSuccess.
// It satisfies fmt.Stringer.
func (g CoAPSuccess) String() string {
	var ret string
	switch g {
	case CoAPSuccessNone:
		ret = "NONE"
	case CoAPSuccessCreated:
		ret = "CREATED"
	case CoAPSuccessDeleted:
		ret = "DELETED"
	case CoAPSuccessValid:
		ret = "VALID"
	case CoAPSuccessChanged:
		ret = "CHANGED"
	case CoAPSuccessContent:
		ret = "CONTENT"
	case CoAPSuccessContinue:
		ret = "CONTINUE"
	}
	return ret
}
