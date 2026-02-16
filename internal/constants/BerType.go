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

// BerType BER encoding enumeration values.
type BerType byte

const (
	// BerTypeEOC defines that the end of Content.
	BerTypeEOC BerType = 0x00
	// BerTypeBoolean defines that the boolean.
	BerTypeBoolean BerType = 0x1
	// BerTypeInteger defines that the integer.
	BerTypeInteger BerType = 0x2
	// BerTypeBitString defines that the bit String.
	BerTypeBitString BerType = 0x3
	// BerTypeOctetString defines that the octet string.
	BerTypeOctetString BerType = 0x4
	// BerTypeNull defines that the nil value.
	BerTypeNull BerType = 0x5
	// BerTypeObjectIdentifier defines that the object identifier.
	BerTypeObjectIdentifier BerType = 0x6
	// BerTypeObjectDescriptor defines that the object Descriptor.
	BerTypeObjectDescriptor BerType = 7
	// BerTypeExternal defines that the external
	BerTypeExternal BerType = 8
	// BerTypeReal defines that the real (float).
	BerTypeReal BerType = 9
	// BerTypeEnumerated defines that the enumerated.
	BerTypeEnumerated BerType = 10
	// BerTypeUtf8String  defines that the utf8 String.
	BerTypeUtf8String BerType = 12
	// BerTypeSequence defines that the sequence.
	BerTypeSequence BerType = 0x10
	// BerTypeSet defines that the set.
	BerTypeSet BerType = 0x11
	// BerTypeNumericString defines that the numeric string.
	BerTypeNumericString BerType = 18
	// BerTypePrintableString defines that the printable string.
	BerTypePrintableString BerType = 19
	// BerTypeTeletexString defines that the teletex string.
	BerTypeTeletexString BerType = 20
	// BerTypeVideotexString defines that the videotex string.
	BerTypeVideotexString BerType = 21
	// BerTypeIa5String defines that the ia5 string
	BerTypeIa5String BerType = 22
	// BerTypeUtcTime defines that the utc time.
	BerTypeUtcTime BerType = 23
	// BerTypeGeneralizedTime defines that the generalized time.
	BerTypeGeneralizedTime BerType = 24
	// BerTypeGraphicString defines that the graphic string.
	BerTypeGraphicString BerType = 25
	// BerTypeVisibleString defines that the visible string.
	BerTypeVisibleString BerType = 26
	// BerTypeGeneralString defines that the general string.
	BerTypeGeneralString BerType = 27
	// BerTypeUniversalString defines that the universal string.
	BerTypeUniversalString BerType = 28
	// BerTypeBmpString defines that the bmp string.
	BerTypeBmpString BerType = 30
	// BerTypeApplication defines that the application class.
	BerTypeApplication BerType = 0x40
	// BerTypeContext defines that the context class.
	BerTypeContext BerType = 0x80
	// BerTypePrivate defines that the private class.
	BerTypePrivate BerType = 0xc0
	// BerTypeConstructed defines that the constructed.
	BerTypeConstructed BerType = 0x20
)

// BerTypeParse converts the given string into a BerType value.
//
// It returns the corresponding BerType constant if the string matches
// a known level name, or an error if the input is invalid.
func BerTypeParse(value string) (BerType, error) {
	var ret BerType
	var err error
	switch strings.ToUpper(value) {
	case "EOC":
		ret = BerTypeEOC
	case "BOOLEAN":
		ret = BerTypeBoolean
	case "INTEGER":
		ret = BerTypeInteger
	case "BITSTRING":
		ret = BerTypeBitString
	case "OCTETSTRING":
		ret = BerTypeOctetString
	case "NULL":
		ret = BerTypeNull
	case "OBJECTIDENTIFIER":
		ret = BerTypeObjectIdentifier
	case "OBJECTDESCRIPTOR":
		ret = BerTypeObjectDescriptor
	case "EXTERNAL":
		ret = BerTypeExternal
	case "REAL":
		ret = BerTypeReal
	case "ENUMERATED":
		ret = BerTypeEnumerated
	case "UTF8STRINGTAG":
		ret = BerTypeUtf8String
	case "SEQUENCE":
		ret = BerTypeSequence
	case "SET":
		ret = BerTypeSet
	case "NUMERICSTRING":
		ret = BerTypeNumericString
	case "PRINTABLESTRING":
		ret = BerTypePrintableString
	case "TELETEXSTRING":
		ret = BerTypeTeletexString
	case "VIDEOTEXSTRING":
		ret = BerTypeVideotexString
	case "IA5STRING":
		ret = BerTypeIa5String
	case "UTCTIME":
		ret = BerTypeUtcTime
	case "GENERALIZEDTIME":
		ret = BerTypeGeneralizedTime
	case "GRAPHICSTRING":
		ret = BerTypeGraphicString
	case "VISIBLESTRING":
		ret = BerTypeVisibleString
	case "GENERALSTRING":
		ret = BerTypeGeneralString
	case "UNIVERSALSTRING":
		ret = BerTypeUniversalString
	case "BMPSTRING":
		ret = BerTypeBmpString
	case "APPLICATION":
		ret = BerTypeApplication
	case "CONTEXT":
		ret = BerTypeContext
	case "PRIVATE":
		ret = BerTypePrivate
	case "CONSTRUCTED":
		ret = BerTypeConstructed
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the BerType.
// It satisfies fmt.Stringer.
func (g BerType) String() string {
	var ret string
	switch g {
	case BerTypeEOC:
		ret = "EOC"
	case BerTypeBoolean:
		ret = "BOOLEAN"
	case BerTypeInteger:
		ret = "INTEGER"
	case BerTypeBitString:
		ret = "BITSTRING"
	case BerTypeOctetString:
		ret = "OCTETSTRING"
	case BerTypeNull:
		ret = "NULL"
	case BerTypeObjectIdentifier:
		ret = "OBJECTIDENTIFIER"
	case BerTypeObjectDescriptor:
		ret = "OBJECTDESCRIPTOR"
	case BerTypeExternal:
		ret = "EXTERNAL"
	case BerTypeReal:
		ret = "REAL"
	case BerTypeEnumerated:
		ret = "ENUMERATED"
	case BerTypeUtf8String:
		ret = "UTF8STRINGTAG"
	case BerTypeSequence:
		ret = "SEQUENCE"
	case BerTypeSet:
		ret = "SET"
	case BerTypeNumericString:
		ret = "NUMERICSTRING"
	case BerTypePrintableString:
		ret = "PRINTABLESTRING"
	case BerTypeTeletexString:
		ret = "TELETEXSTRING"
	case BerTypeVideotexString:
		ret = "VIDEOTEXSTRING"
	case BerTypeIa5String:
		ret = "IA5STRING"
	case BerTypeUtcTime:
		ret = "UTCTIME"
	case BerTypeGeneralizedTime:
		ret = "GENERALIZEDTIME"
	case BerTypeGraphicString:
		ret = "GRAPHICSTRING"
	case BerTypeVisibleString:
		ret = "VISIBLESTRING"
	case BerTypeGeneralString:
		ret = "GENERALSTRING"
	case BerTypeUniversalString:
		ret = "UNIVERSALSTRING"
	case BerTypeBmpString:
		ret = "BMPSTRING"
	case BerTypeApplication:
		ret = "APPLICATION"
	case BerTypeContext:
		ret = "CONTEXT"
	case BerTypePrivate:
		ret = "PRIVATE"
	case BerTypeConstructed:
		ret = "CONSTRUCTED"
	}
	return ret
}
