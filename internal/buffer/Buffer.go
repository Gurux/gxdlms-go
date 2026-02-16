package buffer

// --------------------------------------------------------------------------
//
//	Gurux Ltd
//
// Filename:        $HeadURL$
//
// Version:         $Revision$,
//
//	$Date$
//	$Author$
//
// # Copyright (c) Gurux Ltd
//
// ---------------------------------------------------------------------------
//
//	DESCRIPTION
//
// This file is a part of Gurux Device Framework.
//
// Gurux Device Framework is Open Source software you can redistribute it
// and/or modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation version 2 of the License.
// Gurux Device Framework is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
// See the GNU General Public License for more details.
//
// More information of Gurux products: https://www.gurux.org
//
// This code is licensed under the GNU General Public License v2.
// Full text may be retrieved at http://www.gnu.org/licenses/gpl-2.0.txt
// ---------------------------------------------------------------------------

// ToHex converts a byte array to a hex string.
func ToHex(bytes []byte, addSpace bool) string {
	return ToHexWithRange(bytes, addSpace, 0, len(bytes))
}

// ToHexWithRange converts a byte array to a hex string with specified range.
func ToHexWithRange(bytes []byte, addSpace bool, index, count int) string {
	if len(bytes) == 0 || count == 0 {
		return ""
	}
	var str []byte
	if addSpace {
		str = make([]byte, count*3)
	} else {
		str = make([]byte, count*2)
	}
	length := 0
	for pos := 0; pos < count; pos++ {
		tmp := bytes[index+pos] >> 4
		if tmp > 9 {
			str[length] = tmp + 0x37
		} else {
			str[length] = tmp + 0x30
		}
		length++
		tmp = bytes[index+pos] & 0x0F
		if tmp > 9 {
			str[length] = tmp + 0x37
		} else {
			str[length] = tmp + 0x30
		}
		length++
		if addSpace {
			str[length] = ' '
			length++
		}
	}
	if addSpace {
		length--
	}
	return string(str[:length])
}

// getByteValue converts a hex character to its byte value.
// Returns 0xFF if the character is not a valid hex character.
func getByteValue(c byte) byte {
	// If number
	if c > 0x2F && c < 0x3A {
		return c - '0'
	}
	// If uppercase
	if c > 0x40 && c < 'G' {
		return c - 'A' + 10
	}
	// If lowercase
	if c > 0x60 && c < 'g' {
		return c - 'a' + 10
	}
	return 0xFF
}

// isHex checks if a byte is a valid hex character.
func isHex(c byte) bool {
	return getByteValue(c) != 0xFF
}

// HexToBytes converts a hex string to a byte array.
func HexToBytes(value string) []byte {
	if len(value) == 0 {
		return []byte{}
	}
	length := len(value) / 2
	if len(value)%2 != 0 {
		length++
	}
	buffer := make([]byte, length)
	lastValue := -1
	index := 0
	for i := 0; i < len(value); i++ {
		ch := value[i]
		if isHex(ch) {
			if lastValue == -1 {
				lastValue = int(getByteValue(ch))
			} else {
				buffer[index] = byte(lastValue<<4 | int(getByteValue(ch)))
				lastValue = -1
				index++
			}
		} else if lastValue != -1 && ch == ' ' {
			buffer[index] = byte(lastValue)
			lastValue = -1
			index++
		} else {
			lastValue = -1
		}
	}
	if lastValue != -1 {
		buffer[index] = byte(lastValue)
		index++
	}

	// If there are no spaces in the hex string
	if len(buffer) == index {
		return buffer
	}
	return buffer[:index]
}

// Reverse reverses the order of bytes in the array.
func Reverse(value []byte) {
	for i, j := 0, len(value)-1; i < j; i, j = i+1, j-1 {
		value[i], value[j] = value[j], value[i]
	}
}
