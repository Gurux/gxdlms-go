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
	"fmt"
	"strings"
)

// GXBitString represents an ASN.1 BIT STRING.
//
// The bit string is stored as a byte slice and a padBits count indicates how many
// bits at the end are unused and should be ignored.
type GXBitString struct {
	// Number of extra bits at the end of the string.
	padBits int

	// Bit string bytes.
	value []byte
}

// NewGXBitStringFromString creates a bit string value from a string of '0' and '1' characters.
//
// NOTE: The input is not validated and must only contain '0' and '1' characters.
func NewGXBitStringFromString(value string) (*GXBitString, error) {
	// TODO: Validate that string contains only '0' and '1' characters.
	return &GXBitString{value: []byte(value), padBits: 0}, nil
}

// NewGXBitString creates a GXBitString from raw bytes and a pad bit count.
func NewGXBitString(value []byte, padBits int) (*GXBitString, error) {
	return &GXBitString{value: value, padBits: padBits}, nil
}

// NewGXBitStringFromByteArray creates a GXBitString from an ASN.1 BIT STRING encoded value.
//
// The first byte is treated as the pad bit count, and the remaining bytes are the bit string.
func NewGXBitStringFromByteArray(value []byte) (*GXBitString, error) {
	padBits := int(value[0])
	return NewGXBitString(value[1:], padBits)
}

// NewGXBitStringFromInteger creates a GXBitString from an integer value.
//
// The integer bits are converted into a bit string. padCount specifies the total
// number of bits to include, and any unused bits at the end are represented using padBits.
func NewGXBitStringFromInteger(value int, padCount int) (*GXBitString, error) {
	padBits := padCount % 8
	arr := make([]byte, padCount/8+padBits)
	for pos := 0; pos != len(arr); pos++ {
		arr[pos] = SwapBits(byte(value))
		value >>= 8
	}
	return &GXBitString{value: arr, padBits: padBits}, nil
}

// PadBits returns the number of unused bits at the end of the bit string.
func (g *GXBitString) PadBits() int {
	return g.padBits
}

// SetPadBits sets the number of unused bits at the end of the bit string.
func (g *GXBitString) SetPadBits(value int) error {
	if value < 0 {
		return errors.New("PadBits")
	}
	g.padBits = value
	return nil
}

// Value returns the underlying byte slice of the bit string.
func (g *GXBitString) Value() []byte {
	return g.value
}

// Length returns the number of valid bits in the bit string.
func (g *GXBitString) Length() int {
	if len(g.value) == 0 {
		return 0
	}
	return (8 * len(g.value)) - g.padBits
}

// AppendZeros appends 'count' zero characters to the provided string builder.
func (g *GXBitString) AppendZeros(sb *strings.Builder, count int) {
	for pos := 0; pos != count; pos++ {
		sb.WriteString("0")
	}
}

func (g *GXBitString) String() string {
	if len(g.value) == 0 {
		return ""
	}
	return g.ToString(false)
}

// ToBitString appends a bitstring to a string builder.
func ToBitString(sb *strings.Builder, value byte, count int) {
	if count > 0 {
		if count > 8 {
			count = 8
		}
		for pos := 7; pos > 8-count-1; pos-- {
			if (value & (1 << uint(pos))) != 0 {
				sb.WriteByte('1')
			} else {
				sb.WriteByte('0')
			}
		}
	}
}

// ToString converts the bit string into a textual representation of '0' and '1'.
//
// If showBits is true, the output is prefixed with the bit count (e.g., "8 bit 01010101").
func (g *GXBitString) ToString(showBits bool) string {
	if len(g.value) == 0 {
		return ""
	}
	sb := strings.Builder{}
	sb.Grow(8 * len(g.value))
	for _, it := range g.value {
		ToBitString(&sb, it, 8)
	}
	tmp := sb.String()
	// Remove extra bits.
	tmp = tmp[:len(tmp)-g.padBits]
	if showBits {
		tmp = fmt.Sprintf("%d bit ", (8*len(g.value))-g.padBits) + tmp
	}
	return tmp
}

// ToInteger converts the bit string into an integer value.
//
// The bit string is treated as a little-endian integer (lowest-order byte first).
func (g *GXBitString) ToInteger() int {
	ret := uint32(0)
	if g.value != nil {
		bytePos := 0
		for _, it := range g.value {
			tmp := uint32(SwapBits(it))
			tmp <<= bytePos
			ret |= tmp
			bytePos = bytePos + 8
		}
	}
	return int(ret)
}

// SwapBits reverses the bit order in a byte (least significant bit becomes most significant).
//
// This is used internally when converting between integer values and bit-string representations.
func SwapBits(value byte) byte {
	var ret byte = 0
	for pos := 0; pos != 8; pos++ {
		ret = (ret << 1) | (value & 0x01)
		value = (byte)(value >> 1)
	}
	return ret
}
