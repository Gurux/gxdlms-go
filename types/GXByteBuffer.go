package types

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
// ---------------------------------------------------------------------------

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"strings"

	"github.com/Gurux/gxdlms-go/internal/buffer"
)

// Array capacity increase size.
const arrayCapacity = 10

// GXByteBuffer is a mutable byte buffer with a read position and dynamic capacity.
//
// It is primarily used to build or parse byte streams in a sequential manner,
// with utility methods for reading/writing values in a variety of integer formats.
type GXByteBuffer struct {
	// Read position (next byte to read).
	position int

	// Number of valid bytes in the buffer.
	size int

	// Underlying data storage.
	data []byte
}

// NewGXByteBuffer creates a new empty buffer.
func NewGXByteBuffer() *GXByteBuffer {
	return &GXByteBuffer{
		position: 0,
		size:     0,
		data:     nil,
	}
}

// NewGXByteBufferWithCapacity creates a new buffer with the given initial capacity.
func NewGXByteBufferWithCapacity(capacity int) *GXByteBuffer {
	return &GXByteBuffer{
		position: 0,
		size:     0,
		data:     make([]byte, capacity),
	}
}

// NewGXByteBufferWithData creates a new buffer initialized with the given data.
func NewGXByteBufferWithData(data []byte) *GXByteBuffer {
	b := &GXByteBuffer{
		position: 0,
		size:     len(data),
		data:     make([]byte, len(data)),
	}
	copy(b.data, data)
	return b
}

// NewGXByteBufferFromByteBuffer creates a new GXByteBuffer with initial data.
func NewGXByteBufferFromByteBuffer(data *GXByteBuffer) *GXByteBuffer {
	b := &GXByteBuffer{
		position: 0,
		size:     data.Available(),
		data:     make([]byte, data.Available()),
	}
	copy(b.data, data.data[data.position:data.position+data.Available()])
	return b
}

// Capacity returns the buffer capacity.
func (g *GXByteBuffer) Capacity() int {
	return len(g.data)
}

// SetCapacity sets the buffer capacity.
func (g *GXByteBuffer) SetCapacity(value int) error {
	if value == 0 {
		g.data = nil
		g.size = 0
	} else {
		if len(g.data) == 0 {
			g.data = make([]byte, value)
		} else {
			tmp := g.data
			g.data = make([]byte, value)
			copy(g.data, tmp[0:g.size])
		}
	}
	return nil
}

// Position returns the byte buffer read position.
func (g *GXByteBuffer) Position() int {
	return g.position
}

// SetPosition sets the byte buffer read position.
func (g *GXByteBuffer) SetPosition(value int) error {
	if value > g.size || value < 0 {
		return fmt.Errorf("Position")
	}
	g.position = value
	return nil
}

// Size returns the byte buffer data size.
func (g *GXByteBuffer) Size() int {
	return g.size
}

// SetSize sets the byte buffer data size.
func (g *GXByteBuffer) SetSize(value int) error {
	if value > g.Capacity() || value < 0 {
		return fmt.Errorf("Size")
	}
	g.size = value
	if g.position > g.size {
		g.position = g.size
	}
	return nil
}

// Available returns the amount of non read bytes in the buffer.
func (g *GXByteBuffer) Available() int {
	return g.size - g.position
}

// SetUint8 appends a single byte to the buffer at the current end position.
//
// Parameters:
//
//	value: Byte to append.
func (g *GXByteBuffer) SetUint8(value uint8) error {
	err := g.SetUint8At(g.size, value)
	if err == nil {
		g.size++
	}
	return err
}

// Clear resets the buffer's read position and size without releasing allocated memory.
func (g *GXByteBuffer) Clear() {
	g.position = 0
	g.size = 0
}

// Array returns the contents of the buffer as a byte slice.
//
// The returned slice contains only the valid bytes (up to Size()).
func (g *GXByteBuffer) Array() []byte {
	if g.Capacity() == g.size {
		return g.data
	}
	ret, _ := g.SubArray(0, g.size)
	return ret
}

// SubArray returns a copy of a range of bytes from the buffer.
//
// Parameters:
//
//	index: Start index.
//	count: Number of bytes to copy.
func (g *GXByteBuffer) SubArray(index int, count int) ([]byte, error) {
	tmp := make([]byte, count)
	copy(tmp, g.data[index:index+count])
	return tmp, nil
}

// Move shifts a range of bytes within the buffer.
//
// Parameters:
//
//	srcPos: Source index.
//	destPos: Destination index.
//	count: Number of bytes to move.
func (g *GXByteBuffer) Move(srcPos int, destPos int, count int) error {
	if count < 0 {
		return fmt.Errorf("count")
	}
	if count != 0 {
		if destPos+count > g.size {
			g.SetCapacity(destPos + count)
		}
		copy(g.data[destPos:], g.data[srcPos:srcPos+count])
		g.size = destPos + count
		if g.position > g.size {
			g.position = g.size
		}
	}
	return nil
}

// Trim discards all bytes that have already been read and resets the position to 0.
func (g *GXByteBuffer) Trim() {
	if g.size == g.position {
		g.size = 0
	} else {
		g.Move(g.position, 0, g.size-g.position)
	}
	g.position = 0
}

// Uint8 reads a byte at the current position and advances the position by one.
func (g *GXByteBuffer) Uint8() (uint8, error) {
	value, err := g.Uint8At(g.position)
	if err == nil {
		g.position++
	}
	return value, err
}

// SetUint8At writes a byte to the buffer at the specified index.
//
// Parameters:
//
//	index: Zero-based byte index.
//	value: Value to write.
func (g *GXByteBuffer) SetUint8At(index int, value uint8) error {
	if index >= g.Capacity() {
		err := g.SetCapacity(index + arrayCapacity)
		if err != nil {
			return err
		}
	}
	g.data[index] = value
	return nil
}

// SetUint16 appends a uint16 value to the buffer (big-endian) at the current end position.
//
// Parameters:
//
//	value: Value to append.
func (g *GXByteBuffer) SetUint16(value uint16) error {
	err := g.SetUint16At(g.size, value)
	if err == nil {
		g.size += 2
	}
	return err
}

// SetInt16 appends an int16 value to the buffer (big-endian) at the current end position.
//
// Parameters:
//
//	value: Value to append.
func (g *GXByteBuffer) SetInt16(value int16) error {
	err := g.SetInt16At(g.size, value)
	if err == nil {
		g.size += 2
	}
	return err
}

// Uint16 reads a big-endian uint16 from the current position and advances the position by 2.
func (g *GXByteBuffer) Uint16() (uint16, error) {
	value, err := g.Uint16At(g.position)
	if err == nil {
		g.position += 2
	}
	return value, err
}

// SetUint16At writes a big-endian uint16 to the buffer at the specified index.
//
// Parameters:
//
//	index: Zero-based byte index.
//	value: Value to write.
func (g *GXByteBuffer) SetUint16At(index int, value uint16) error {
	if index+2 > g.Capacity() {
		err := g.SetCapacity(index + arrayCapacity)
		if err != nil {
			return err
		}
	}
	g.data[index] = uint8(((value >> 8) & 0xFF))
	g.data[index+1] = uint8((value & 0xFF))
	return nil
}

// SetInt16At writes a big-endian int16 to the buffer at the specified index.
//
// Parameters:
//
//	index: Zero-based byte index.
//	value: Value to write.
func (g *GXByteBuffer) SetInt16At(index int, value int16) error {
	if index+2 > g.Capacity() {
		err := g.SetCapacity((index + arrayCapacity))
		if err != nil {
			return err
		}
	}
	g.data[index] = uint8(((value >> 8) & 0xFF))
	g.data[index+1] = uint8((value & 0xFF))
	return nil
}

// SetUint32 appends a uint32 value (big-endian) to the buffer at the current end position.
//
// Parameters:
//
//	value: Value to append.
func (g *GXByteBuffer) SetUint32(value uint32) error {
	err := g.SetUint32At(g.size, value)
	if err == nil {
		g.size += 4
	}
	return err
}

// SetInt32 appends an int32 value (big-endian) to the buffer at the current end position.
//
// Parameters:
//
//	value: Value to append.
func (g *GXByteBuffer) SetInt32(value int32) error {
	err := g.SetInt32At(g.size, value)
	if err == nil {
		g.size += 4
	}
	return err
}

// Uint32 reads a big-endian uint32 from the current position and advances the position by 4.
func (g *GXByteBuffer) Uint32() (uint32, error) {
	value, err := g.Uint32At(g.position)
	if err == nil {
		g.position += 4
	}
	return value, err
}

// SetUint32At writes a big-endian uint32 value to the buffer at the specified index.
//
// Parameters:
//
//	index: Zero-based byte index.
//	value: Value to write.
func (g *GXByteBuffer) SetUint32At(index int, value uint32) error {
	if index+4 > g.Capacity() {
		err := g.SetCapacity((index + arrayCapacity))
		if err != nil {
			return err
		}
	}
	g.data[index] = uint8(((value >> 24) & 0xFF))
	g.data[index+1] = uint8(((value >> 16) & 0xFF))
	g.data[index+2] = uint8(((value >> 8) & 0xFF))
	g.data[index+3] = uint8((value & 0xFF))
	return nil
}

// SetInt32At writes a big-endian int32 value to the buffer at the specified index.
//
// Parameters:
//
//	index: Zero-based byte index.
//	value: Value to write.
func (g *GXByteBuffer) SetInt32At(index int, value int32) error {
	if index+4 > g.Capacity() {
		err := g.SetCapacity((index + arrayCapacity))
		if err != nil {
			return err
		}
	}
	g.data[index] = uint8(((value >> 24) & 0xFF))
	g.data[index+1] = uint8(((value >> 16) & 0xFF))
	g.data[index+2] = uint8(((value >> 8) & 0xFF))
	g.data[index+3] = uint8((value & 0xFF))
	return nil
}

// SetUint64 appends a uint64 value (big-endian) to the buffer at the current end position.
//
// Parameters:
//
//	value: Value to append.
func (g *GXByteBuffer) SetUint64(value uint64) error {
	err := g.SetUint64At(g.size, value)
	if err == nil {
		g.size += 8
	}
	return err
}

// SetInt64 appends an int64 value (big-endian) to the buffer at the current end position.
//
// Parameters:
//
//	value: Value to append.
func (g *GXByteBuffer) SetInt64(value int64) error {
	err := g.SetInt64At(g.size, value)
	if err == nil {
		g.size += 8
	}
	return err
}

// Uint64 reads a big-endian uint64 from the current position and advances the position by 8.
func (g *GXByteBuffer) Uint64() (uint64, error) {
	value, err := g.Uint64At(g.position)
	if err == nil {
		g.position += 8
	}
	return value, err
}

// Uint64At reads a big-endian uint64 value from the buffer at the given index.
//
// Parameters:
//
//	index: Byte index.
func (g *GXByteBuffer) Uint64At(index int) (uint64, error) {
	if index+8 > g.size {
		return 0, fmt.Errorf("index out of range")
	}
	value := ((uint64(g.data[index]) & 0xFF) << 56) | ((uint64(g.data[index+1]) & 0xFF) << 48) | ((uint64(g.data[index+2]) & 0xFF) << 40) | ((uint64(g.data[index+3]) & 0xFF) << 32) | ((uint64(g.data[index+4]) & 0xFF) << 24) | ((uint64(g.data[index+5]) & 0xFF) << 16) | ((uint64(g.data[index+6]) & 0xFF) << 8) | (uint64(g.data[index+7]) & 0xFF)
	return value, nil
}

// SetUint64At writes a big-endian uint64 value to the buffer at the specified index.
//
// Parameters:
//
//	index: Zero-based byte index.
//	item: Value to write.
func (g *GXByteBuffer) SetUint64At(index int, item uint64) error {
	if index+8 > g.Capacity() {
		err := g.SetCapacity((index + arrayCapacity))
		if err != nil {
			return err
		}
	}
	g.data[g.size] = uint8(((item >> 56) & 0xFF))
	g.data[g.size+1] = uint8(((item >> 48) & 0xFF))
	g.data[g.size+2] = uint8(((item >> 40) & 0xFF))
	g.data[g.size+3] = uint8(((item >> 32) & 0xFF))
	g.data[g.size+4] = uint8(((item >> 24) & 0xFF))
	g.data[g.size+5] = uint8(((item >> 16) & 0xFF))
	g.data[g.size+6] = uint8(((item >> 8) & 0xFF))
	g.data[g.size+7] = uint8((item & 0xFF))
	return nil
}

// SetInt64At writes a big-endian int64 value to the buffer at the specified index.
//
// Parameters:
//
//	index: Zero-based byte index.
//	item: Value to write.
func (g *GXByteBuffer) SetInt64At(index int, item int64) error {
	if index+8 > g.Capacity() {
		err := g.SetCapacity((index + arrayCapacity))
		if err != nil {
			return err
		}
	}
	g.data[g.size] = uint8(((item >> 56) & 0xFF))
	g.data[g.size+1] = uint8(((item >> 48) & 0xFF))
	g.data[g.size+2] = uint8(((item >> 40) & 0xFF))
	g.data[g.size+3] = uint8(((item >> 32) & 0xFF))
	g.data[g.size+4] = uint8(((item >> 24) & 0xFF))
	g.data[g.size+5] = uint8(((item >> 16) & 0xFF))
	g.data[g.size+6] = uint8(((item >> 8) & 0xFF))
	g.data[g.size+7] = uint8((item & 0xFF))
	return nil
}

// Int8 reads a signed 8-bit integer from the current position and advances the position.
func (g *GXByteBuffer) Int8() (int8, error) {
	val, ret := g.Uint8()
	return int8(val), ret
}

// Uint8At reads a byte at the specified index.
//
// Parameters:
//
//	index: Byte index.
func (g *GXByteBuffer) Uint8At(index int) (uint8, error) {
	if index >= g.size {
		return 0, fmt.Errorf("index out of range")
	}
	return g.data[index], nil
}

// Uint16At reads a big-endian uint16 value from the specified index.
//
// Parameters:
//
//	index: Byte index.
func (g *GXByteBuffer) Uint16At(index int) (uint16, error) {
	if index+2 > g.size {
		return 0, fmt.Errorf("index out of range")
	}
	return uint16(((uint16(g.data[index]&0xFF) << 8) | uint16(g.data[index+1]&0xFF))), nil
}

// Int32 reads a big-endian int32 from the current position and advances the position by 4.
func (g *GXByteBuffer) Int32() (int32, error) {
	var value, err = g.Uint32At(g.position)
	if err == nil {
		g.position += 4
	}
	return int32(value), err
}

// Int16 reads a big-endian int16 from the current position and advances the position by 2.
func (g *GXByteBuffer) Int16() (int16, error) {
	var value, err = g.Int16At(g.position)
	if err == nil {
		g.position += 2
	}
	return value, err
}

// Int16At reads a big-endian int16 from the specified index.
//
// Parameters:
//
//	index: Byte index.
func (g *GXByteBuffer) Int16At(index int) (int16, error) {
	if index+2 > g.size {
		return 0, fmt.Errorf("index out of range")
	}
	return int16(((int16(g.data[index]&0xFF) << 8) | (int16(g.data[index+1] & 0xFF)))), nil
}

// Uint32At reads a big-endian uint32 value from the specified index.
//
// Parameters:
//
//	index: Byte index.
func (g *GXByteBuffer) Uint32At(index int) (uint32, error) {
	if index+4 > g.size {
		return 0, fmt.Errorf("index out of range")
	}
	return uint32((uint32(g.data[index]&0xFF)<<24 | uint32(g.data[index+1]&0xFF)<<16 | uint32(g.data[index+2]&0xFF)<<8 | uint32(g.data[index+3]&0xFF))), nil
}

// Uint24At reads a 24-bit unsigned integer from the specified index.
//
// Parameters:
//
//	index: Byte index.
func (g *GXByteBuffer) Uint24At(index int) (int, error) {
	if index+3 > g.size {
		return 0, fmt.Errorf("index out of range")
	}
	return int((int(g.data[index]&0xFF)<<16 | int(g.data[index+1]&0xFF)<<8 | int(g.data[index+2]&0xFF))), nil
}

// Uint24 reads a 24-bit unsigned integer from the current position and advances the position by 3.
func (g *GXByteBuffer) Uint24() (int, error) {
	var value, err = g.Uint24At(g.position)
	if err == nil {
		g.position += 3
	}
	return value, err
}

// Float reads a float32 (big-endian) from the current position and advances the position by 4.
func (g *GXByteBuffer) Float() (float32, error) {
	var value, err = g.Int32()
	return math.Float32frombits(uint32(value)), err
}

// SetFloat appends a float32 value (big-endian) to the buffer.
func (g *GXByteBuffer) SetFloat(value float32) error {
	tmp := make([]byte, 4)
	binary.BigEndian.PutUint32(tmp, math.Float32bits(value))
	return g.Set(tmp)
}

// SetDouble appends a float64 value (big-endian) to the buffer.
func (g *GXByteBuffer) SetDouble(value float64) error {
	tmp := make([]byte, 8)
	binary.BigEndian.PutUint64(tmp, math.Float64bits(value))
	return g.Set(tmp)
}

// Double reads a float64 (big-endian) from the current position and advances the position by 8.
func (g *GXByteBuffer) Double() (float64, error) {
	value, err := g.Int64()
	return math.Float64frombits(uint64(value)), err
}

// Int64 reads a big-endian int64 from the current position and advances the position by 8.
func (g *GXByteBuffer) Int64() (int64, error) {
	if g.position+8 > g.size {
		return 0, fmt.Errorf("index out of range")
	}
	value := ((int64(g.data[g.position]) & 0xFF) << 56) | ((int64(g.data[g.position+1]) & 0xFF) << 48) | ((int64(g.data[g.position+2]) & 0xFF) << 40) | ((int64(g.data[g.position+3]) & 0xFF) << 32) | ((int64(g.data[g.position+4]) & 0xFF) << 24) | ((int64(g.data[g.position+5]) & 0xFF) << 16) | ((int64(g.data[g.position+6]) & 0xFF) << 8) | (int64(g.data[g.position+7]) & 0xFF)
	g.position += 8
	return value, nil
}

// IsAsciiString reports whether the given byte slice contains only ASCII characters.
//
// Tabs, line feeds, carriage returns, and NUL are permitted.
func IsAsciiString(value []byte) bool {
	if len(value) != 0 {
		for _, it := range value {
			if (it < 32 || it > 127) && it != '\r' && it != '\n' && it != '\t' && it != 0 {
				return false
			}
		}
	}
	return true
}

// String returns the buffer contents as a string.
//
// If the contents are ASCII, it returns the ASCII string (trimming at the first NUL).
// Otherwise it returns a hex representation.
func (g *GXByteBuffer) String() string {
	tmp := g.data[:g.size]
	if IsAsciiString(tmp) {
		var str = string(tmp)
		var pos = strings.IndexByte(str, 0)
		if pos != -1 {
			str = str[:pos]
		}
		return str
	}
	return ToHex(tmp, true)
}

// StringWithRange returns a string constructed from a range of bytes in the buffer.
//
// Parameters:
//
//	index: Start byte index.
//	count: Number of bytes.
func (g *GXByteBuffer) StringWithRange(index int, count int) (string, error) {
	if index < 0 {
		return "", fmt.Errorf("index")
	}
	if count < 0 {
		return "", fmt.Errorf("count")
	}
	if index+count > g.size {
		return "", fmt.Errorf("index out of range")
	}
	return string(g.data[index : index+count]), nil
}

// StringUtf8At returns a UTF-8 string from a range of bytes in the buffer.
//
// Parameters:
//
//	index: Start byte index.
//	count: Number of bytes.
func (g *GXByteBuffer) StringUtf8At(index int, count int) (string, error) {
	if index < 0 {
		return "", fmt.Errorf("index")
	}
	if count < 0 {
		return "", fmt.Errorf("count")
	}
	return string(g.data[index : index+count]), nil
}

// Set appends the given byte slice to the buffer at the current end position.
//
// Parameters:
//
//	value: Bytes to append.
func (g *GXByteBuffer) Set(value []byte) error {
	if len(value) != 0 {
		return g.SetAt(value, 0, len(value))
	}
	return nil
}

// SetAt appends a slice of bytes to the buffer.
//
// Parameters:
//
//	value: Source byte slice.
//	index: Offset in the source slice.
//	count: Number of bytes to append (-1 means "rest of slice").
func (g *GXByteBuffer) SetAt(value []byte, index int, count int) error {
	if len(value) != 0 && count != 0 {
		if count == -1 {
			count = len(value) - index
		}
		if g.size+count > g.Capacity() {
			g.SetCapacity((g.size + count + arrayCapacity))
		}
		copy(g.data[g.size:g.size+count], value[index:index+count])
		g.size += count
	}
	return nil
}

// InsertBytes inserts data at the given index and shifts existing data to the right.
func (g *GXByteBuffer) InsertBytes(index int, value []byte) error {
	if index < 0 || index > g.size {
		return fmt.Errorf("index out of range")
	}
	if len(value) == 0 {
		return nil
	}
	needed := g.size + len(value)
	if needed > g.Capacity() {
		if err := g.SetCapacity(needed + arrayCapacity); err != nil {
			return err
		}
	}
	copy(g.data[index+len(value):], g.data[index:g.size])
	copy(g.data[index:index+len(value)], value)
	g.size = needed
	if g.position > g.size {
		g.position = g.size
	}
	return nil
}

func (g *GXByteBuffer) SetByteBuffer(value *GXByteBuffer) error {
	var err error
	if value != nil {
		err := g.Set(value.data[value.position:value.size])
		if err == nil {
			value.position = value.size
		}
	}
	return err
}

// SetByteBufferByCount sets the  new value to byte array.
// value: Byte array to add.
// count: Byte count.
func (g *GXByteBuffer) SetByteBufferByCount(value *GXByteBuffer, count int) error {
	if g.size+count > g.Capacity() {
		g.SetCapacity(g.size + count + arrayCapacity)
	}
	var err error
	if count != 0 {
		copy(g.data[g.size:g.size+count], value.data[value.position:value.position+count])
		g.size += count
		err = value.SetPosition(count)
	}
	return err
}

// add returns the add new object to the byte buffer.
// value: Value to add.
func (g *GXByteBuffer) Add(value any) error {
	var err error
	if value != nil {
		if v1, ok := value.([]byte); ok {
			err = g.Set(v1)
		} else if v2, ok := value.(uint8); ok {
			err = g.SetUint8(v2)
		} else if v3, ok := value.(uint16); ok {
			err = g.SetUint16(v3)
		} else if v4, ok := value.(uint32); ok {
			err = g.SetUint32(v4)
		} else if v5, ok := value.(uint64); ok {
			err = g.SetUint64(v5)
		} else if v6, ok := value.(int8); ok {
			err = g.SetUint8(uint8(v6))
		} else if v7, ok := value.(int16); ok {
			err = g.SetInt16(v7)
		} else if v8, ok := value.(int32); ok {
			err = g.SetInt32(v8)
		} else if v9, ok := value.(int64); ok {
			err = g.SetInt64(v9)
		} else if v10, ok := value.(string); ok {
			err = g.Set([]byte(v10))
		} else {
			err = fmt.Errorf("Invalid object type.")
		}
	}
	return err
}

// get returns the get value from the byte array.
// target: Target array.
func (g *GXByteBuffer) Get(target []byte) error {
	if g.size-g.position < len(target) {
		return fmt.Errorf("index out of range")
	}
	copy(target, g.data[g.position:g.position+len(target)])
	g.position += len(target)
	return nil
}

// compare returns the compares, whether two given arrays are similar starting from current position.
// arr: Array to compare.
// Returns: True, if arrays are similar. False, if the arrays differ.
func (g *GXByteBuffer) Compare(arr []byte) bool {
	if arr == nil || g.size-g.position < len(arr) {
		return false
	}
	tmp := make([]byte, len(arr))
	g.Get(tmp)
	var ret = bytes.Equal(tmp, arr)
	if !ret {
		g.position -= len(arr)
	}
	return ret
}

// IsHexString checks if a string is a valid hex string.
func IsHexString(value string) bool {
	if len(value) == 0 {
		return false
	}
	for _, ch := range value {
		if ch != ' ' && !((ch > 0x40 && ch < 'G') ||
			(ch > 0x60 && ch < 'g') || (ch > '/' && ch < ':')) {
			return false
		}
	}
	return true
}

// HexToBytes converts a hex string to a byte array.
func HexToBytes(value string) []byte {
	return buffer.HexToBytes(value)
}

// ToHex converts a byte array to a hex string.
func ToHex(bytes []byte, addSpace bool) string {
	return buffer.ToHexWithRange(bytes, addSpace, 0, len(bytes))
}

// ToHexWithRange converts a byte array to a hex string with specified range.
func ToHexWithRange(bytes []byte, addSpace bool, index, count int) string {
	return buffer.ToHexWithRange(bytes, addSpace, index, count)
}

// setHexString returns the push the given hex string as byte array into this buffer at the current position, and then increments the position.value: The hex string to be added.
func (g *GXByteBuffer) SetHexString(value string) error {
	return g.Set(buffer.HexToBytes(value))
}

// setHexString returns the push the given hex string as byte array into this buffer at the current position, and then increments the position.index: Byte index.
// value: The hex string to be added.
func (g *GXByteBuffer) SetHexStringByIndex(index int, value string) {
	tmp := buffer.HexToBytes(value)
	g.SetAt(tmp, index, len(tmp)-index)
}

// setHexString returns the push the given hex string as byte array into this buffer at the current position, and then increments the position.value: Byte array to add.
// index: Byte index.
// count: Byte count.
func (g *GXByteBuffer) SetHexStringWithRange(value string, index int, count int) {
	tmp := HexToBytes(value)
	g.SetAt(tmp, index, count)
}

// SetObjectCount writes the length of a constructed object into the buffer using ASN.1 length encoding.
//
// For lengths < 0x80, a single byte is written. For larger lengths, a multi-byte length field is
// written according to ASN.1 DER length encoding rules.
func SetObjectCount(count int, buff *GXByteBuffer) error {
	var err error
	if count < 0x80 {
		return buff.SetUint8(uint8(count))
	}
	if count < 0x100 {
		err = buff.SetUint8(0x81)
		if err != nil {
			return err
		}
		return buff.SetUint8(uint8(count))
	}
	if count < 0x10000 {
		err = buff.SetUint8(0x82)
		if err != nil {
			return err
		}
		return buff.SetUint16(uint16(count))
	}
	err = buff.SetUint8(0x84)
	if err != nil {
		return err
	}
	return buff.SetUint32(uint32(count))
}

// ObjectCountBuffer is the minimal interface needed to insert an encoded object count.
type ObjectCountBuffer interface {
	InsertBytes(index int, value []byte) error
}

// InsertObjectCount inserts a BER-encoded object count at the given index.
func InsertObjectCount(count int, buff ObjectCountBuffer, index int) error {
	if count < 0 {
		return fmt.Errorf("count")
	}
	var tmp [5]byte
	var n int
	switch {
	case count < 0x80:
		tmp[0] = byte(count)
		n = 1
	case count < 0x100:
		tmp[0] = 0x81
		tmp[1] = byte(count)
		n = 2
	case count < 0x10000:
		tmp[0] = 0x82
		binary.BigEndian.PutUint16(tmp[1:], uint16(count))
		n = 3
	default:
		tmp[0] = 0x84
		binary.BigEndian.PutUint32(tmp[1:], uint32(count))
		n = 5
	}
	return buff.InsertBytes(index, tmp[:n])
}

// GetObjectCount reads an ASN.1 length field from the buffer and returns the decoded length.
//
// This complements SetObjectCount and supports the DER length encoding form used when
// reading the length of constructed objects.
func GetObjectCount(data *GXByteBuffer) (int, error) {
	ch, err := data.Uint8()
	if err != nil {
		return 0, err
	}
	if ch < 0x80 {
		return int(ch), nil
	}
	cnt := int(ch & 0x7F)
	if cnt == 1 {
		v, err := data.Uint8()
		return int(v), err
	}
	if cnt == 2 {
		v, err := data.Uint16()
		return int(v), err
	}
	if cnt == 4 {
		v, err := data.Uint32()
		return int(v), err
	}
	return 0, fmt.Errorf("invalid object count header: 0x%X", ch)
}

// remaining returns the get remaining data.
//
//	Returns: Remaining data as byte array.
func (g *GXByteBuffer) Remaining() ([]byte, error) {
	return g.SubArray(g.position, g.size-g.position)
}

// RemainingHexString returns the get remaining data as hex string.
// addSpace: Add space between bytes.
//
//	Returns:
//	    Remaining data as hex string
func (g *GXByteBuffer) RemainingHexString(addSpace bool) string {
	return buffer.ToHexWithRange(g.data, addSpace, g.position, g.size-g.position)
}

// ToHexByIndex returns the get data as hex string.
// addSpace: Add space between bytes.
// index: Byte index.
//
//	Returns:
//	    Data as hex string.
func (g *GXByteBuffer) ToHexByIndex(addSpace bool, index int) string {
	return buffer.ToHexWithRange(g.data, addSpace, index, g.size-index)
}
