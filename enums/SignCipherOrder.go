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

// SignCipherOrder Enumerates Signing and ciphering order.
type SignCipherOrder int

const (
	// SignCipherOrderCipheredFirst defines that the message is ciphered first and then signed.
	SignCipherOrderCipheredFirst SignCipherOrder = iota
	// SignCipherOrderSignedFirst defines that the message is signed first and then ciphered.
	SignCipherOrderSignedFirst
)

// SignCipherOrderParse converts the given string into a SignCipherOrder value.
//
// It returns the corresponding SignCipherOrder constant if the string matches
// a known level name, or an error if the input is invalid.
func SignCipherOrderParse(value string) (SignCipherOrder, error) {
	var ret SignCipherOrder
	var err error
	switch strings.ToUpper(value) {
	case "CIPHEREDFIRST":
		ret = SignCipherOrderCipheredFirst
	case "SIGNEDFIRST":
		ret = SignCipherOrderSignedFirst
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the SignCipherOrder.
// It satisfies fmt.Stringer.
func (g SignCipherOrder) String() string {
	var ret string
	switch g {
	case SignCipherOrderCipheredFirst:
		ret = "CIPHEREDFIRST"
	case SignCipherOrderSignedFirst:
		ret = "SIGNEDFIRST"
	}
	return ret
}
