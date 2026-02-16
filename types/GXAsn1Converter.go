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
	"encoding/base64"
	"errors"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal/constants"
)

// ASN1 converter. This class is used to convert public and private keys to byte array and vice verse.
type GXAsn1Converter struct {
}

// getValue parses ASN.1 encoded data from a byte buffer and adds parsed objects to the objects list.
// It handles various BER types including sequences, context tags, strings, integers, and time values.
//
// Parameters:
//
//	bb: Byte buffer containing ASN.1 encoded data.
//	objects: List to store parsed objects.
//	s: Optional XML settings for debug output.
//	getNext: If true, stops after parsing one object.
//
// Returns:
//
//	Error if parsing fails due to invalid data or insufficient buffer space.
func getValue(bb *GXByteBuffer, objects *[]any, s *gxAsn1Settings, getNext bool) error {
	var len_ int
	var tmp []any
	var tmp2 []byte

	type_, err := bb.Uint8()
	if err != nil {
		return err
	}
	len_, err = GetObjectCount(bb)
	if err != nil {
		return err
	}
	if len_ > bb.Available() {
		return errors.New("GXAsn1Converter.GetValue: insufficient buffer space")
	}

	connectPos := 0
	if s != nil {
		connectPos = s.XmlLength()
	}
	start := bb.Position()
	var tagString string
	if s != nil {
		s.AppendSpaces()
		if type_ == uint8(constants.BerTypeInteger) {
			if len_ == 1 || len_ == 2 || len_ == 4 || len_ == 8 {
				tagString = s.GetTag(int16(constants.BerTypeInteger))
			} else {
				tagString = s.GetTag(int16(constants.BerTypeInteger))
			}
		} else {
			tagString = s.GetTag(int16(type_))
		}
		s.Append("<" + tagString + ">")
	}

	// Process based on BER type
	switch type_ {
	case uint8(constants.BerTypeConstructed | constants.BerTypeContext),
		uint8(constants.BerTypeConstructed|constants.BerTypeContext) | 1,
		uint8(constants.BerTypeConstructed|constants.BerTypeContext) | 2,
		uint8(constants.BerTypeConstructed|constants.BerTypeContext) | 3,
		uint8(constants.BerTypeConstructed|constants.BerTypeContext) | 4,
		uint8(constants.BerTypeConstructed|constants.BerTypeContext) | 5:
		if s != nil {
			s.Increase()
		}
		ctx := &GXAsn1Context{Index: int(type_ & 0xF), Constructed: true}
		*objects = append(*objects, ctx)
		for bb.Position() < start+len_ {
			if err := getValue(bb, &ctx.Items, s, false); err != nil {
				return err
			}
		}
		if s != nil {
			s.Decrease()
		}

	case uint8(constants.BerTypeConstructed | constants.BerTypeSequence):
		if s != nil {
			s.Increase()
		}
		seq := &GXAsn1Sequence{}
		*objects = append(*objects, seq)
		cnt := 0
		for bb.Position() < start+len_ {
			cnt++
			if err := getValue(bb, (*[]any)(seq), s, false); err != nil {
				return err
			}
			if getNext {
				break
			}
		}
		if s != nil {
			s.AppendComment(connectPos, fmt.Sprintf("%d elements.", cnt))
			s.Decrease()
		}

	case uint8(constants.BerTypeConstructed | constants.BerTypeSet):
		if s != nil {
			s.Increase()
		}
		tmp = make([]any, 0)
		if err := getValue(bb, &tmp, s, false); err != nil {
			return err
		}
		if len(tmp) > 0 {
			if seq, ok := tmp[0].(*GXAsn1Sequence); ok {
				var val any
				if len(*seq) > 1 {
					val = (*seq)[1]
				}
				*objects = append(*objects, NewGXKeyValuePair((*seq)[0], val))
			} else {
				*objects = append(*objects, NewGXKeyValuePair[any, any](tmp, nil))
			}
		}
		if s != nil {
			s.Decrease()
		}

	case uint8(constants.BerTypeObjectIdentifier):
		oi := NewGXAsn1ObjectIdentifierFromByteBuffer(bb, len_)
		*objects = append(*objects, oi)
		if s != nil {
			desc, err := oi.Description()
			if err == nil {
				s.AppendComment(connectPos, desc)
			}
			s.Append(oi.String())
		}

	case uint8(constants.BerTypePrintableString):
		str, err := bb.StringWithRange(bb.Position(), len_)
		if err != nil {
			return err
		}
		*objects = append(*objects, str)
		if s != nil {
			s.Append(str)
		}

	case uint8(constants.BerTypeBmpString):
		str, err := bb.StringWithRange(bb.Position(), len_)
		if err != nil {
			return err
		}
		*objects = append(*objects, str)
		if s != nil {
			s.Append(str)
		}

	case uint8(constants.BerTypeUtf8String):
		str, err := bb.StringWithRange(bb.Position(), len_)
		if err != nil {
			return err
		}
		utf8Str := NewGXAsn1Utf8String(str)
		*objects = append(*objects, utf8Str)
		if s != nil {
			s.Append(str)
		}

	case uint8(constants.BerTypeIa5String):
		str, err := bb.StringWithRange(bb.Position(), len_)
		if err != nil {
			return err
		}
		ia5Str := &GXAsn1Ia5String{Value: str}
		*objects = append(*objects, ia5Str)
		if s != nil {
			s.Append(str)
		}

	case uint8(constants.BerTypeInteger):
		switch len_ {
		case 1:
			v, err := bb.Int8()
			if err != nil {
				return err
			}
			*objects = append(*objects, v)
		case 2:
			v, err := bb.Int16()
			if err != nil {
				return err
			}
			*objects = append(*objects, v)
		case 4:
			v, err := bb.Int32()
			if err != nil {
				return err
			}
			*objects = append(*objects, v)
		default:
			tmp2 = make([]byte, len_)
			bb.Get(tmp2)
			*objects = append(*objects, NewGXAsn1Integer(tmp2))
		}
		if s != nil && len(*objects) > 0 {
			s.Append(fmt.Sprintf("%v", (*objects)[len(*objects)-1]))
		}

	case uint8(constants.BerTypeNull):
		*objects = append(*objects, nil)

	case uint8(constants.BerTypeBitString):
		bitData, err := bb.SubArray(bb.Position(), len_)
		if err != nil {
			return err
		}
		bb.SetPosition(bb.Position() + len_)
		bitStr, err := NewGXBitString(bitData, 0)
		if err != nil {
			return err
		}
		*objects = append(*objects, bitStr)
		if s != nil {
			s.AppendComment(connectPos, fmt.Sprintf("%d bit.", len_*8))
			s.Append(bitStr.String())
		}

	case uint8(constants.BerTypeUtcTime):
		tmp2 = make([]byte, len_)
		bb.Get(tmp2)
		dateStr := string(tmp2)
		t, err := getUtcTime(dateStr)
		if err != nil {
			return err
		}
		*objects = append(*objects, t)
		if s != nil {
			s.Append(t.Format("2006-01-02 15:04"))
		}

	case uint8(constants.BerTypeGeneralizedTime):
		tmp2 = make([]byte, len_)
		bb.Get(tmp2)
		dateStr := string(tmp2)
		t, _, err := getGeneralizedTime(dateStr)
		if err != nil {
			return err
		}
		*objects = append(*objects, t)
		if s != nil {
			s.Append(t.Format("2006-01-02 15:04:05Z07:00"))
		}

	case uint8(constants.BerTypeContext),
		uint8(constants.BerTypeContext) | 1,
		uint8(constants.BerTypeContext) | 2,
		uint8(constants.BerTypeContext) | 3,
		uint8(constants.BerTypeContext) | 4,
		uint8(constants.BerTypeContext) | 5,
		uint8(constants.BerTypeContext) | 6:
		ctx := &GXAsn1Context{Index: int(type_ & 0xF), Constructed: false}
		tmp2 = make([]byte, len_)
		bb.Get(tmp2)
		ctx.Items = append(ctx.Items, tmp2)
		*objects = append(*objects, ctx)
		if s != nil {
			s.Append(ToHex(tmp2, false))
		}

	case uint8(constants.BerTypeOctetString):
		if bb.Available() > 0 {
			t, err := bb.Uint8At(bb.Position())
			if err != nil {
				return err
			}
			switch t {
			case uint8(constants.BerTypeConstructed | constants.BerTypeSequence),
				uint8(constants.BerTypeBitString):
				if s != nil {
					s.Increase()
				}
				if err := getValue(bb, objects, s, false); err != nil {
					return err
				}
				if s != nil {
					s.Decrease()
				}
			default:
				tmp2 = make([]byte, len_)
				bb.Get(tmp2)
				*objects = append(*objects, tmp2)
				if s != nil {
					s.Append(ToHex(tmp2, false))
				}
			}
		}

	case uint8(constants.BerTypeBoolean):
		b, err := bb.Uint8()
		if err != nil {
			return err
		}
		*objects = append(*objects, b != 0)
		if s != nil {
			s.Append(fmt.Sprintf("%v", b != 0))
		}
	default:
		return fmt.Errorf("invalid BER type: %d", type_)
	}

	if s != nil {
		s.Append("</" + tagString + ">\r\n")
	}

	return nil
}

func getUtcTime(dateString string) (time.Time, error) {
	v, err := strconv.Atoi(dateString[0:2])
	if err != nil {
		return time.Time{}, err
	}
	year := 2000 + v
	v, err = strconv.Atoi(dateString[2:4])
	if err != nil {
		return time.Time{}, err
	}
	month := v
	v, err = strconv.Atoi(dateString[4:6])
	if err != nil {
		return time.Time{}, err
	}
	day := v
	v, err = strconv.Atoi(dateString[6:8])
	if err != nil {
		return time.Time{}, err
	}
	hour := v
	v, err = strconv.Atoi(dateString[8:10])
	if err != nil {
		return time.Time{}, err
	}

	second := 0
	minute := v
	if strings.HasSuffix(dateString, "Z") {
		if len(dateString) > 11 {
			v, err = strconv.Atoi(dateString[10:12])
			if err != nil {
				return time.Time{}, err
			}
			second = v
		}
		return time.Date(year, time.Month(month), day, hour, minute, second, 0, time.UTC), nil
	}
	if len(dateString) > 15 {
		second, err = strconv.Atoi(dateString[10:12])
		if err != nil {
			return time.Time{}, err
		}
	}
	//TODO: tmp := dateString[len(dateString)-6 : len(dateString)-1-len(dateString)-6]
	return time.Date(year, time.Month(month), day, hour, minute, second, 0, time.UTC), nil
}

func getGeneralizedTime(dateString string) (time.Time, int, error) {
	if len(dateString) < 12 {
		return time.Time{}, 0, fmt.Errorf("dateString too short: %q", dateString)
	}
	year, err := strconv.Atoi(dateString[0:4])
	if err != nil {
		return time.Time{}, 0, err
	}
	month, err := strconv.Atoi(dateString[4:6])
	if err != nil {
		return time.Time{}, 0, err
	}
	day, err := strconv.Atoi(dateString[6:8])
	if err != nil {
		return time.Time{}, 0, err
	}
	hour, err := strconv.Atoi(dateString[8:10])
	if err != nil {
		return time.Time{}, 0, err
	}
	minute, err := strconv.Atoi(dateString[10:12])
	if err != nil {
		return time.Time{}, 0, err
	}
	second := 0
	if strings.HasSuffix(dateString, "Z") {
		if len(dateString) > 13 {
			second, err = strconv.Atoi(dateString[12:14])
			if err != nil {
				return time.Time{}, 0, err
			}
		}
		return time.Date(year, time.Month(month), day, hour, minute, second, 0, time.UTC), 0, nil
	}
	if len(dateString) > 17 {
		second, err = strconv.Atoi(dateString[12:14])
		if err != nil {
			return time.Time{}, 0, err
		}
	}
	sign := 1
	last5 := dateString[len(dateString)-5:]
	if last5[0] == '-' {
		sign = -1
	} else if last5[0] != '+' {
		last5 = "+" + dateString[len(dateString)-4:]
	}
	hh, err := strconv.Atoi(last5[1:3])
	if err != nil {
		return time.Time{}, 0, err
	}
	mm, err := strconv.Atoi(last5[3:5])
	if err != nil {
		return time.Time{}, 0, err
	}
	offsetMinutes := sign * (60*hh + mm)
	loc := time.FixedZone("offset", offsetMinutes*60)
	t := time.Date(year, time.Month(month), day, hour, minute, second, 0, loc)
	return t, offsetMinutes, nil
}

func dateToString(date time.Time) string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("%02d", date.Year()-2000))
	sb.WriteString(fmt.Sprintf("%02d", date.Month()))
	sb.WriteString(fmt.Sprintf("%02d", date.Day()))
	sb.WriteString(fmt.Sprintf("%02d", date.Hour()))
	sb.WriteString(fmt.Sprintf("%02d", date.Minute()))
	sb.WriteString(fmt.Sprintf("%02d", date.Second()))
	sb.WriteString("Z")
	return sb.String()
}

// getBytes returns the add ASN1 object to byte buffer.
//
// Parameters:
//
//	bb: Byte buffer where ANS1 object is serialized.
//	target: ANS1 object
//
// Returns:
//
//	Size of object.
func getBytes(bb *GXByteBuffer, target any) (int, error) {
	cnt := 0
	tmp := GXByteBuffer{}
	start := bb.Size()
	switch v := target.(type) {
	case *GXAsn1Context:
		tmp = GXByteBuffer{}
		for _, it := range v.Items {
			n, err := getBytes(&tmp, it)
			if err != nil {
				return 0, err
			}
			cnt += n
		}

		start2 := bb.Size()
		if v.Constructed {
			bb.SetUint8(byte(constants.BerTypeConstructed) | byte(constants.BerTypeContext) | byte(v.Index))
			SetObjectCount(cnt, bb)
		} else {
			tmp.SetUint8At(0, byte(constants.BerTypeContext)|byte(v.Index))
		}
		cnt += bb.Size() - start2
		bb.SetByteBuffer(&tmp)
		return cnt, nil
	case []any:
		tmp := GXByteBuffer{}
		for _, it := range v {
			n, err := getBytes(&tmp, it)
			if err != nil {
				return 0, err
			}
			cnt += n
		}
		start2 := bb.Size()
		bb.SetUint8(byte(constants.BerTypeConstructed | constants.BerTypeSequence))
		SetObjectCount(cnt, bb)
		cnt += bb.Size() - start2
		bb.SetByteBuffer(&tmp)
		return cnt, nil
	case string:
		bb.SetUint8(byte(constants.BerTypePrintableString))
		SetObjectCount(len(v), bb)
		bb.Add(v)
	case int8:
		bb.SetUint8(byte(constants.BerTypeInteger))
		SetObjectCount(1, bb)
		bb.Add(v)
	case int16:
		bb.SetUint8(byte(constants.BerTypeInteger))
		SetObjectCount(2, bb)
		bb.Add(v)
	case int32:
		bb.SetUint8(byte(constants.BerTypeInteger))
		SetObjectCount(4, bb)
		bb.Add(v)
	case int:
		bb.SetUint8(byte(constants.BerTypeInteger))
		SetObjectCount(4, bb)
		bb.Add(int32(v))
	case *GXAsn1Integer:
		bb.SetUint8(byte(constants.BerTypeInteger))
		SetObjectCount(len(v.Value()), bb)
		bb.Add(v.Value())
	case int64:
		bb.SetUint8(byte(constants.BerTypeInteger))
		SetObjectCount(8, bb)
		bb.Add(v)
	case []byte:
		bb.SetUint8(byte(constants.BerTypeOctetString))
		SetObjectCount(len(v), bb)
		bb.Add(v)
	case nil:
		bb.SetUint8(byte(constants.BerTypeNull))
		SetObjectCount(0, bb)
	case bool:
		bb.SetUint8(byte(constants.BerTypeBoolean))
		bb.SetUint8(1)
		if v {
			bb.SetUint8(255)
		} else {
			bb.SetUint8(0)
		}
	case *GXAsn1ObjectIdentifier:
		ret, err := v.Encoded()
		if err != nil {
			return 0, err
		}
		bb.SetUint8(byte(constants.BerTypeObjectIdentifier))
		SetObjectCount(len(ret), bb)
		bb.Add(ret)
	case *GXKeyValuePair[any, any]:
		tmp2 := GXByteBuffer{}
		if v.Value != nil {
			cnt = 0
			n, err := getBytes(&tmp2, v.Key)
			if err != nil {
				return 0, err
			}
			cnt += n
			n, err = getBytes(&tmp2, v.Value)
			if err != nil {
				return 0, err
			}
			cnt += n

			tmp := GXByteBuffer{}
			tmp.SetUint8(byte(constants.BerTypeConstructed | constants.BerTypeSequence))
			SetObjectCount(cnt, &tmp)
			tmp.SetByteBuffer(&tmp2)
		} else {
			list, ok := v.Key.([]any)
			if !ok || len(list) == 0 {
				return 0, fmt.Errorf("Pair.Key must be []any with at least one item when Value is nil")
			}
			if _, err := getBytes(&tmp2, list[0]); err != nil {
				return 0, err
			}
			tmp = tmp2
		}

		// Update len.
		before := bb.Size()
		bb.SetUint8(byte(constants.BerTypeConstructed | constants.BerTypeSet))
		SetObjectCount(tmp.Size(), bb)
		bb.SetByteBuffer(&tmp)
		return bb.Size() - before, nil
	case *GXAsn1Utf8String:
		bb.SetUint8(byte(constants.BerTypeUtf8String))
		s := v.String()
		SetObjectCount(len(s), bb)
		bb.Add(s)
	case *GXAsn1Ia5String:
		bb.SetUint8(byte(constants.BerTypeIa5String))
		s := v.String()
		SetObjectCount(len(s), bb)
		bb.Add(s)

	case *GXBitString:
		bb.SetUint8(byte(constants.BerTypeBitString))
		SetObjectCount(1+len(v.Value()), bb)
		bb.SetUint8(uint8(v.PadBits()))
		bb.Add(v.Value())
	case *GXAsn1PublicKey:
		bb.SetUint8(byte(constants.BerTypeBitString))
		// Size = 64 + padding + uncompressed point indicator.
		SetObjectCount(66, bb)
		bb.SetUint8(0) // padding
		bb.SetUint8(4) // uncompressed point indicator 0x04
		bb.Add(v.Value())
		// Count = type + size + 64 + padding + 0x04 => 68
		return 68, nil

	case time.Time:
		// Save date time in UTC.
		bb.SetUint8(byte(constants.BerTypeUtcTime))
		s := dateToString(v.UTC())
		bb.SetUint8(byte(len(s)))
		bb.Add(s)
	case *GXAsn1Sequence:
		tmp := &GXByteBuffer{}
		for _, it := range *v {
			n, err := getBytes(tmp, it)
			if err != nil {
				return 0, err
			}
			cnt += n
		}
		start2 := bb.Size()
		bb.SetUint8(byte(constants.BerTypeConstructed | constants.BerTypeSequence))
		SetObjectCount(cnt, bb)
		cnt += bb.Size() - start2
		bb.SetByteBuffer(tmp)
		return cnt, nil
	default:
		return 0, fmt.Errorf("invalid type: %T", target)
	}
	return bb.Size() - start, nil
}

func Asn1GetSubject(values *GXAsn1Sequence) string {
	sb := strings.Builder{}
	for _, v := range *values {
		it := v.(GXKeyValuePair[any, any])
		sb.WriteString(X509NameFromString(it.Key.(string)).String())
		sb.WriteString("=")
		sb.WriteString(fmt.Sprintf("%v", it.Value))
		sb.WriteString(", ")
	}
	// Remove last comma.
	if sb.Len() != 0 {
		tmp := sb.String()
		sb.Reset()
		sb.WriteString(tmp[0 : len(tmp)-2])
	}
	return sb.String()
}

// Asn1GetCertificateType returns the get certificate type from byte array.
//
// Parameters:
//
//	data: Byte array.
//	seq: Byte array.
//
// Returns:
//
//	Certificate type.
func Asn1GetCertificateType(data []byte, seq GXAsn1Sequence) (enums.PkcsType, error) {
	if seq == nil {
		ret, err := Asn1FromByteArray(data)
		if err != nil {
			return enums.PkcsTypeNone, err
		}
		if ps, ok := ret.(GXAsn1Sequence); ok {
			seq = ps
		}
	}
	if _, ok := seq[0].(GXAsn1Sequence); ok {
		_, err := NewGXx509Certificate(data)
		if err == nil {
			return enums.PkcsTypex509Certificate, nil
		}
	}
	if _, ok := seq[0].(GXAsn1Sequence); ok {
		_, err := NewGXPkcs10(data)
		if err == nil {
			return enums.PkcsTypePkcs10, nil
		}
	}

	if _, ok := seq[0].(int8); ok {
		_, err := NewGXPkcs8(data)
		if err == nil {
			return enums.PkcsTypePkcs8, nil
		}
	}
	return enums.PkcsTypeNone, nil
}

// GetFilePath returns the default file path.
//
// Parameters:
//
//	scheme: Used scheme.
//	certificateType: Certificate type.
//	systemTitle: System title.
//
// Returns:
//
//	File path.
func (g *GXAsn1Converter) GetFilePath(scheme enums.Ecc, certificateType enums.CertificateType, systemTitle []byte) (string, error) {
	var path string
	switch certificateType {
	case enums.CertificateTypeDigitalSignature:
		path = "D"
	case enums.CertificateTypeKeyAgreement:
		path = "A"
	case enums.CertificateTypeTLS:
		path = "T"
	default:
		return "", errors.New("Unknown certificate type.")
	}
	path = path + ToHex(systemTitle, false) + ".pem"
	if scheme == enums.EccP256 {
		path = filepath.Join("Keys", path)
	} else {
		path = filepath.Join("Keys384", path)
	}
	return path, nil
}

// EncodeSubject converts a subject string into a list of key-value pairs with
// object identifiers and values formatted for X.509 certificate use.
//
// Parameters:
//
//	value: Subject string with format "KEY1=VALUE1,KEY2=VALUE2,..." where
//	       KEY values are X509Name enum names.
//
// Returns:
//
//	List of GXKeyValuePair with GXAsn1ObjectIdentifier keys and string/ASN1 values.
//	Returns error if subject format is invalid or X509Name parsing fails.
func Asn1EncodeSubject(value string) ([]*GXKeyValuePair[*GXAsn1ObjectIdentifier, any], error) {
	list := make([]*GXKeyValuePair[*GXAsn1ObjectIdentifier, any], 0)

	// Split by comma and remove empty entries
	parts := strings.Split(value, ",")
	for _, part := range parts {
		tmp := strings.TrimSpace(part)
		if tmp == "" {
			continue
		}

		// Split by equals sign
		it := strings.Split(tmp, "=")
		if len(it) != 2 {
			return nil, errors.New("invalid subject format: expected KEY=VALUE pairs separated by commas")
		}
		nameStr := strings.TrimSpace(it[0])
		valueStr := strings.TrimSpace(it[1])

		// Parse X509Name enum
		name, err := enums.X509NameParse(nameStr)
		if err != nil {
			return nil, fmt.Errorf("invalid X509Name: %w", err)
		}

		// Determine value type based on X509Name
		var val any
		switch name {
		case enums.X509NameC:
			// Country code is printable string
			val = valueStr
		case enums.X509NameE:
			// Email address in Verisign certificates
			val = &GXAsn1Ia5String{Value: valueStr}
		default:
			// All other fields use UTF8String
			val = NewGXAsn1Utf8String(valueStr)
		}

		// Get OID string from X509Name
		oid, err := X509NameToString(name)
		if err != nil {
			return nil, fmt.Errorf("failed to convert X509Name to OID string: %w", err)
		}

		// Create key-value pair and add to list
		pair := NewGXKeyValuePair(NewGXAsn1ObjectIdentifier(oid), val)
		list = append(list, pair)
	}
	return list, nil
}

// Asn1FromByteArray returns the convert byte array to ASN1 objects.
//
// Parameters:
//
//	data: ASN-1 bytes.
//
// Returns:
//
//	Parsed objects.
func Asn1FromByteArray(data []byte) (any, error) {
	bb := GXByteBuffer{}
	bb.Set(data)
	objects := []any{}
	for bb.Position() != bb.Size() {
		err := getValue(&bb, &objects, nil, false)
		if err != nil {
			return nil, err
		}
	}
	if len(objects) == 0 {
		return nil, nil
	}
	return objects[0], nil
}

// Asn1GetNext returns the get next ASN1 value from the byte buffer.
func Asn1GetNext(data *GXByteBuffer) (any, error) {
	objects := []any{}
	err := getValue(data, &objects, nil, true)
	if err != nil {
		return nil, err
	}
	return objects[0], nil
}

// Asn1ToByteArray returns the convert ASN1 objects to byte array.
//
// Parameters:
//
//	objects: ASN.1 objects.
//
// Returns:
//
//	ASN.1 objects as byte array.
func Asn1ToByteArray(objects any) ([]byte, error) {
	bb := GXByteBuffer{}
	_, err := getBytes(&bb, objects)
	if err != nil {
		return nil, err
	}
	return bb.Array(), nil
}

// Asn1SystemTitleToSubject returns the convert system title to subject.
//
// Parameters:
//
//	systemTitle: System title.
//
// Returns:
//
//	Subject.
func Asn1SystemTitleToSubject(systemTitle []byte) string {
	return "CN=" + ToHex(systemTitle, false)
}

// SystemTitleFromSubject returns the get system title from the subject.
//
// Parameters:
//
//	subject: Subject.
//
// Returns:
//
//	System title.
func SystemTitleFromSubject(subject string) ([]byte, error) {
	hex, err := HexSystemTitleFromSubject(subject)
	if err != nil {
		return nil, err
	}
	return HexToBytes(hex), nil
}

// HexSystemTitleFromSubject returns the get system title in hex string from the subject.
//
// Parameters:
//
//	subject: Subject.
//
// Returns:
//
//	System title.
func HexSystemTitleFromSubject(subject string) (string, error) {
	index := strings.Index(subject, "CN=")
	if index == -1 {
		return "", errors.New("System title not found from the subject.")
	}
	return subject[index+3 : index+19], nil
}

// CertificateTypeToKeyUsage returns the convert ASN1 certificate type to DLMS key usage.
//
// Parameters:
//
//	type: Certificate type.
//
// Returns:
//
//	Key usage.
func (g *GXAsn1Converter) CertificateTypeToKeyUsage(type_ enums.CertificateType) enums.KeyUsage {
	var k enums.KeyUsage
	switch type_ {
	case enums.CertificateTypeDigitalSignature:
		k = enums.KeyUsageDigitalSignature
	case enums.CertificateTypeKeyAgreement:
		k = enums.KeyUsageKeyAgreement
	case enums.CertificateTypeTLS:
		k = enums.KeyUsageKeyCertSign
	case enums.CertificateTypeOther:
		k = enums.KeyUsageCrlSign
	default:
		k = enums.KeyUsageNone
	}
	return k
}

// getCertificateTypeInternal returns certificate type from byte array.
// If seq is nil, it will be parsed from data.
func getCertificateTypeInternal(data []byte, seq *GXAsn1Sequence) enums.PkcsType {
	if seq == nil {
		root, err := Asn1FromByteArray(data)
		if err == nil {
			if s, ok := root.(GXAsn1Sequence); ok {
				seq = &s
			} else if ps, ok := root.(*GXAsn1Sequence); ok && ps != nil {
				seq = ps
			}
		}
	}

	if len(*seq) == 0 {
		return enums.PkcsTypeNone
	}
	if _, ok := (*seq)[0].(GXAsn1Sequence); ok {
		if _, err := NewGXx509Certificate(data); err == nil {
			return enums.PkcsTypex509Certificate
		}
		// ok if fails
	}
	if _, ok := (*seq)[0].(GXAsn1Sequence); ok {
		if _, err := NewGXPkcs10(data); err == nil {
			return enums.PkcsTypePkcs10
		}
		// ok if fails
	}
	if _, ok := (*seq)[0].(int8); ok {
		if _, err := NewGXPkcs8(data); err == nil {
			return enums.PkcsTypePkcs8
		}
		// ok if fails
	}
	return enums.PkcsTypeNone
}

// Asn1GetCertificateTypeFromDer returns the get certificate type from DER string.
//
// Parameters:
//
//	der: DER string
func Asn1GetCertificateTypeFromDer(der string) (enums.PkcsType, error) {
	ret, err := base64.StdEncoding.DecodeString(der)
	if err != nil {
		return enums.PkcsTypeNone, err
	}
	return Asn1GetCertificateType(ret, nil)
}
