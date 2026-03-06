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
	"errors"

	"github.com/Gurux/gxcommon-go"
	"github.com/Gurux/gxdlms-go/internal/buffer"
)

// GXAsn1PublicKey represents an ASN.1 public key encoded as a fixed-length byte array.
//
// This type is used when decoding X.509/PKCS keys embedded within ASN.1 structures.
type GXAsn1PublicKey struct {
	// value holds the raw public key bytes.
	value []byte
}

// Value returns the raw public key bytes.
func (g *GXAsn1PublicKey) Value() []byte {
	return g.value
}

func (g *GXAsn1PublicKey) init(key []byte) error {
	if len(key) != 270 {
		return gxcommon.ErrInvalidArgument
	}
	g.value = make([]byte, len(key))
	copy(g.value, key)
	return nil
}

// String implements fmt.Stringer and returns the hex encoding of the public key.
func (g *GXAsn1PublicKey) String() string {
	if len(g.value) == 0 {
		return ""
	}
	return buffer.ToHex(g.value, false)
}

// NewGXAsn1PublicKey creates a GXAsn1PublicKey from the raw public key bytes.
func NewGXAsn1PublicKey(key []byte) (*GXAsn1PublicKey, error) {
	g := &GXAsn1PublicKey{}
	err := g.init(key)
	if err != nil {
		return nil, err
	}
	return g, nil
}

// NewGXAsn1PublicKeyFromBitString creates a public key from an ASN.1 BIT STRING.
func NewGXAsn1PublicKeyFromBitString(data *GXBitString) (*GXAsn1PublicKey, error) {
	if data == nil {
		return nil, errors.New("data is nil")
	}
	v, err := Asn1FromByteArray(data.Value())
	if err != nil {
		return nil, err
	}
	seq := v.(GXAsn1Sequence)
	tmp, err := Asn1ToByteArray([]any{seq[0], seq[1]})
	if err != nil {
		return nil, err
	}
	g := &GXAsn1PublicKey{}
	err = g.init(tmp)
	if err != nil {
		return nil, err
	}
	return g, nil
}
