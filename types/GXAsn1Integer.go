package types

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
	"math/big"

	"github.com/Gurux/gxdlms-go/internal/buffer"
)

// ASN1 integer value.
type GXAsn1Integer struct {
	// Bit string.
	value []byte
}

// Constructor from byte array.
func NewGXAsn1Integer(value []byte) *GXAsn1Integer {
	tmp := make([]byte, len(value))
	copy(tmp, value)
	return &GXAsn1Integer{value: tmp}
}

// Constructor from big integer.
func NewGXAsn1IntegerFromBigInteger(value big.Int) *GXAsn1Integer {
	g := &GXAsn1Integer{}
	g.value = value.Bytes()
	buffer.Reverse(g.value)
	return g
}

// Bit string.
func (g *GXAsn1Integer) Value() []byte {
	return g.value
}

// ToBigInteger returns integer value as big integer.
func (g *GXAsn1Integer) ToBigInteger() *big.Int {
	tmp := make([]byte, len(g.value))
	copy(tmp, g.value)
	buffer.Reverse(tmp)
	return new(big.Int).SetBytes(tmp)
}

// ToByte returns integer value as byte.
func (g *GXAsn1Integer) ToByte() (int8, error) {
	bb := GXByteBuffer{}
	err := bb.Set(g.value)
	if err != nil {
		return 0, err
	}
	return bb.Int8()
}

// ToShort returns integer value as short.
func (g *GXAsn1Integer) ToShort() (int16, error) {
	bb := GXByteBuffer{}
	err := bb.Set(g.value)
	if err != nil {
		return 0, err
	}
	return bb.Int16()
}

// ToInt returns integer value as int.
func (g *GXAsn1Integer) ToInt() (int32, error) {
	bb := GXByteBuffer{}
	err := bb.Set(g.value)
	if err != nil {
		return 0, err
	}
	return bb.Int32()
}

// ToLong returns integer value as long.
func (g *GXAsn1Integer) ToLong() (int64, error) {
	bb := GXByteBuffer{}
	err := bb.Set(g.value)
	if err != nil {
		return 0, err
	}
	return bb.Int64()
}

// String returns integer value as string.
func (g *GXAsn1Integer) String() string {
	var str string
	switch len(g.value) {
	case 1:
		ret, err := g.ToByte()
		if err != nil {
			return ""
		}
		str = fmt.Sprint(ret)
	case 2:
		ret, err := g.ToShort()
		if err != nil {
			return ""
		}
		str = fmt.Sprint(ret)
	case 4:
		ret, err := g.ToInt()
		if err != nil {
			return ""
		}
		str = fmt.Sprint(ret)
	case 8:
		ret, err := g.ToLong()
		if err != nil {
			return ""
		}
		str = fmt.Sprint(ret)
	default:
		n := new(big.Int).SetBytes(g.value)
		str = n.String()
	}
	return str
}
