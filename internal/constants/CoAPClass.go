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

// CoAP class.
type CoAPClass int

const (
	// CoAPClassMethod defines that the // Request method.
	CoAPClassMethod CoAPClass = iota
	// CoAPClassSuccess defines that the // Success response.
	CoAPClassSuccess CoAPClass = 2
	// CoAPClassClientError defines that the // Client error response.
	CoAPClassClientError CoAPClass = 4
	// CoAPClassServerError defines that the // Server error response.
	CoAPClassServerError CoAPClass = 5
	// CoAPClassSignaling defines that the // Signaling.
	CoAPClassSignaling CoAPClass = 7
)

// CoAPClassParse converts the given string into a CoAPClass value.
//
// It returns the corresponding CoAPClass constant if the string matches
// a known level name, or an error if the input is invalid.
func CoAPClassParse(value string) (CoAPClass, error) {
	var ret CoAPClass
	var err error
	switch strings.ToUpper(value) {
	case "METHOD":
		ret = CoAPClassMethod
	case "SUCCESS":
		ret = CoAPClassSuccess
	case "CLIENTERROR":
		ret = CoAPClassClientError
	case "SERVERERROR":
		ret = CoAPClassServerError
	case "SIGNALING":
		ret = CoAPClassSignaling
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the CoAPClass.
// It satisfies fmt.Stringer.
func (g CoAPClass) String() string {
	var ret string
	switch g {
	case CoAPClassMethod:
		ret = "METHOD"
	case CoAPClassSuccess:
		ret = "SUCCESS"
	case CoAPClassClientError:
		ret = "CLIENTERROR"
	case CoAPClassServerError:
		ret = "SERVERERROR"
	case CoAPClassSignaling:
		ret = "SIGNALING"
	}
	return ret
}
