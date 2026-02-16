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

// BitString class is used with Bit strings.
type GXBitString struct {
	// Number of extra bits at the end of the string.
	padBits int

	// Bit string.
	value []byte
}

// NewGXBitString creates new GXBitString struct and initializes it with given parameters.
//
// Parameters:
//
//	value: Bit string.
//	padBits: Number of extra bits at the end of the string.
func NewGXBitStringFromString(value string) (*GXBitString, error) {
	//TODO: Check that string contains only '0' and '1' characters.
	return &GXBitString{value: []byte(value), padBits: 0}, nil
}

// NewGXBitString creates new GXBitString struct and initializes it with given parameters.
//
// Parameters:
//
//	value: Bit string.
//	padBits: Number of extra bits at the end of the string.
func NewGXBitString(value []byte, padBits int) (*GXBitString, error) {
	return &GXBitString{value: value, padBits: padBits}, nil
}

// NewGXBitStringFromByteArray creates new GXBitString struct and initializes it with given parameters.
//
// Parameters:
//
//	value: Bit string.
//	padBits: Number of extra bits at the end of the string.
func NewGXBitStringFromByteArray(value []byte) (*GXBitString, error) {
	padBits := int(value[0])
	return NewGXBitString(value[1:], padBits)
}

// NewGXBitStringFromInteger creates new GXBitString struct and initializes it with given parameters.
//
// Parameters:
//
//	value: Integer value that is converted to bit string.
//	padCount: Number of bits in the bit string.
func NewGXBitStringFromInteger(value int, padCount int) (*GXBitString, error) {
	padBits := padCount % 8
	arr := make([]byte, padCount/8+padBits)
	for pos := 0; pos != len(arr); pos++ {
		arr[pos] = SwapBits(byte(value))
		value >>= 8
	}
	return &GXBitString{value: arr, padBits: padBits}, nil
}

// PadBits returns the number of extra bits at the end of the string.
func (g *GXBitString) PadBits() int {
	return g.padBits
}

// SetPadBits sets the number of extra bits at the end of the string.
func (g *GXBitString) SetPadBits(value int) error {
	if value < 0 {
		return errors.New("PadBits")
	}
	g.padBits = value
	return nil
}

// Bit string.
func (g *GXBitString) Value() []byte {
	return g.value
}

// Length returns the number of extra bits at the end of the string.
func (g *GXBitString) Length() int {
	if len(g.value) == 0 {
		return 0
	}
	return (8 * len(g.value)) - g.padBits
}

// AppendZeros returns the append zeroes to the buffer.
//
// Parameters:
//
//	sb: Buffer where zeros are added.
//	count: Amount of zeroes.
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

// String produces convert bit string to string.
//
// Parameters:
//
//	showBits: Is the number of the bits shown.
//
// Returns:
//
//	Bit string as an string.
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

// ToInteger returns the converts ASN1 bit-string to integer value.
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

// Reserved for internal use.
func SwapBits(value byte) byte {
	var ret byte = 0
	for pos := 0; pos != 8; pos++ {
		ret = (ret << 1) | (value & 0x01)
		value = (byte)(value >> 1)
	}
	return ret
}
