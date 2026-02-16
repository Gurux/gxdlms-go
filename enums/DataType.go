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

// DataType DataType enumerates usable types of data in GuruxDLMS.
type DataType int

const (
	// DataTypeNone defines that no defines that the data type is set.
	DataTypeNone DataType = iota
	// DataTypeArray defines that the defines that the data type is Array.
	DataTypeArray = 1
	// DataTypeBcd defines that the data type is Binary coded decimal.
	DataTypeBcd = 13
	// DataTypeBitString defines that the data type is Bit string.
	DataTypeBitString = 4
	// DataTypeBoolean defines that the data type is Boolean.
	DataTypeBoolean = 3
	// DataTypeCompactArray defines that the data type is Compact array.
	DataTypeCompactArray = 0x13
	// DataTypeDate defines that the data type is Date.
	DataTypeDate = 0x1a
	// DataTypeDateTime defines that the data type is DateTime.
	DataTypeDateTime = 0x19
	// DataTypeEnum defines that the data type is Enum.
	DataTypeEnum = 0x16
	// DataTypeFloat32 defines that the data type is Float32.
	DataTypeFloat32 = 0x17
	// DataTypeFloat64 defines that the data type is Float64.
	DataTypeFloat64 = 0x18
	// DataTypeInt16 defines that the data type is Int16.
	DataTypeInt16 = 0x10
	// DataTypeInt32 defines that the data type is Int32.
	DataTypeInt32 = 5
	// DataTypeInt64 defines that the data type is Int64.
	DataTypeInt64 = 20
	// DataTypeInt8 defines that the data type is Int8.
	DataTypeInt8 = 15
	// DataTypeOctetString defines that the data type is Octet string.
	DataTypeOctetString = 9
	// DataTypeString defines that the data type is String.
	DataTypeString = 10
	// DataTypeStringUTF8 defines that the data type is UTF8 String.
	DataTypeStringUTF8 = 12
	// DataTypeStructure defines that the data type is Structure.
	DataTypeStructure = 2
	// DataTypeTime defines that the data type is Time.
	DataTypeTime = 0x1b
	// DataTypeDeltaInt8 defines that the data type is delta integer.
	DataTypeDeltaInt8 = 28
	// DataTypeDeltaInt16 defines that the data type is delta long.
	DataTypeDeltaInt16 = 29
	// DataTypeDeltaInt32 defines that the data type is delta double long.
	DataTypeDeltaInt32 = 30
	// DataTypeDeltaUint8 defines that the data type is delta unsigned.
	DataTypeDeltaUint8 = 31
	// DataTypeDeltaUint16 defines that the data type is delta long.
	DataTypeDeltaUint16 = 32
	// DataTypeDeltaUint32 defines that the data type is delta double long.
	DataTypeDeltaUint32 = 33
	// DataTypeUint16 defines that the data type is UInt16.
	DataTypeUint16 = 0x12
	// DataTypeUint32 defines that the data type is UInt32.
	DataTypeUint32 = 6
	// DataTypeUint64 defines that the data type is UInt64.
	DataTypeUint64 = 0x15
	// DataTypeUint8 defines that the data type is UInt8.
	DataTypeUint8 = 0x11
)

// DataTypeParse converts the given string into a DataType value.
//
// It returns the corresponding DataType constant if the string matches
// a known level name, or an error if the input is invalid.
func DataTypeParse(value string) (DataType, error) {
	var ret DataType
	var err error
	switch strings.ToUpper(value) {
	case "Array":
		ret = DataTypeArray
	case "Bcd":
		ret = DataTypeBcd
	case "BitString":
		ret = DataTypeBitString
	case "Boolean":
		ret = DataTypeBoolean
	case "CompactArray":
		ret = DataTypeCompactArray
	case "Date":
		ret = DataTypeDate
	case "DateTime":
		ret = DataTypeDateTime
	case "Enum":
		ret = DataTypeEnum
	case "Float32":
		ret = DataTypeFloat32
	case "Float64":
		ret = DataTypeFloat64
	case "Int16":
		ret = DataTypeInt16
	case "Int32":
		ret = DataTypeInt32
	case "Int64":
		ret = DataTypeInt64
	case "Int8":
		ret = DataTypeInt8
	case "None":
		ret = DataTypeNone
	case "OctetString":
		ret = DataTypeOctetString
	case "String":
		ret = DataTypeString
	case "StringUTF8":
		ret = DataTypeStringUTF8
	case "Structure":
		ret = DataTypeStructure
	case "Time":
		ret = DataTypeTime
	case "DeltaInt8":
		ret = DataTypeDeltaInt8
	case "DeltaInt16":
		ret = DataTypeDeltaInt16
	case "DeltaInt32":
		ret = DataTypeDeltaInt32
	case "DeltaUint8":
		ret = DataTypeDeltaUint8
	case "DeltaUint16":
		ret = DataTypeDeltaUint16
	case "DeltaUint32":
		ret = DataTypeDeltaUint32
	case "Uint16":
		ret = DataTypeUint16
	case "Uint32":
		ret = DataTypeUint32
	case "Uint64":
		ret = DataTypeUint64
	case "Uint8":
		ret = DataTypeUint8
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the enums.DataType
// It satisfies fmt.Stringer.
func (g DataType) String() string {
	var ret string
	switch g {
	case DataTypeArray:
		ret = "Array"
	case DataTypeBcd:
		ret = "Bcd"
	case DataTypeBitString:
		ret = "BitString"
	case DataTypeBoolean:
		ret = "Boolean"
	case DataTypeCompactArray:
		ret = "CompactArray"
	case DataTypeDate:
		ret = "Date"
	case DataTypeDateTime:
		ret = "DateTime"
	case DataTypeEnum:
		ret = "Enum"
	case DataTypeFloat32:
		ret = "Float32"
	case DataTypeFloat64:
		ret = "Float64"
	case DataTypeInt16:
		ret = "Int16"
	case DataTypeInt32:
		ret = "Int32"
	case DataTypeInt64:
		ret = "Int64"
	case DataTypeInt8:
		ret = "Int8"
	case DataTypeNone:
		ret = "None"
	case DataTypeOctetString:
		ret = "OctetString"
	case DataTypeString:
		ret = "String"
	case DataTypeStringUTF8:
		ret = "StringUTF8"
	case DataTypeStructure:
		ret = "Structure"
	case DataTypeTime:
		ret = "Time"
	case DataTypeDeltaInt8:
		ret = "DeltaInt8"
	case DataTypeDeltaInt16:
		ret = "DeltaInt16"
	case DataTypeDeltaInt32:
		ret = "DeltaInt32"
	case DataTypeDeltaUint8:
		ret = "DeltaUint8"
	case DataTypeDeltaUint16:
		ret = "DeltaUint16"
	case DataTypeDeltaUint32:
		ret = "DeltaUint32"
	case DataTypeUint16:
		ret = "Uint16"
	case DataTypeUint32:
		ret = "Uint32"
	case DataTypeUint64:
		ret = "Uint64"
	case DataTypeUint8:
		ret = "Uint8"
	}
	return ret
}

// AllDataType returns a slice containing all defined DataType values.
func AllDataType() []DataType {
	return []DataType{
		DataTypeArray,
		DataTypeBcd,
		DataTypeBitString,
		DataTypeBoolean,
		DataTypeCompactArray,
		DataTypeDate,
		DataTypeDateTime,
		DataTypeEnum,
		DataTypeFloat32,
		DataTypeFloat64,
		DataTypeInt16,
		DataTypeInt32,
		DataTypeInt64,
		DataTypeInt8,
		DataTypeNone,
		DataTypeOctetString,
		DataTypeString,
		DataTypeStringUTF8,
		DataTypeStructure,
		DataTypeTime,
		DataTypeDeltaInt8,
		DataTypeDeltaInt16,
		DataTypeDeltaInt32,
		DataTypeDeltaUint8,
		DataTypeDeltaUint16,
		DataTypeDeltaUint32,
		DataTypeUint16,
		DataTypeUint32,
		DataTypeUint64,
		DataTypeUint8,
	}
}
