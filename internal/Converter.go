package internal

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
	"reflect"
	"strconv"

	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/types"
)

// Convert converts value to requested DLMS data type.
func Convert(value any, dt enums.DataType) (any, error) {
	if value == nil {
		if dt == enums.DataTypeNone {
			return nil, nil
		}
		return nil, fmt.Errorf("value is nil")
	}

	switch dt {
	case enums.DataTypeNone:
		return value, nil
	case enums.DataTypeOctetString:
		switch v := value.(type) {
		case []byte:
			return v, nil
		case string:
			if v == "" {
				return []byte{}, nil
			}
			ret, err := HexToBytes(v)
			if err != nil {
				return nil, err
			}
			return ret, nil
		default:
			return nil, fmt.Errorf("cannot convert %T to %v", value, dt)
		}
	case enums.DataTypeDateTime:
		switch v := value.(type) {
		case types.GXDateTime:
			return v, nil
		case *types.GXDateTime:
			if v == nil {
				return nil, fmt.Errorf("value is nil")
			}
			return *v, nil
		case string:
			ret, err := types.NewGXDateTimeFromString(v, nil)
			if err != nil {
				return nil, err
			}
			return *ret, nil
		case []byte:
			return ChangeTypeFromByteArray(nil, v, dt)
		default:
			return nil, fmt.Errorf("cannot convert %T to %v", value, dt)
		}
	case enums.DataTypeDate:
		switch v := value.(type) {
		case types.GXDate:
			return v, nil
		case *types.GXDate:
			if v == nil {
				return nil, fmt.Errorf("value is nil")
			}
			return *v, nil
		case string:
			ret, err := types.NewGXDateFromString(v, nil)
			if err != nil {
				return nil, err
			}
			return *ret, nil
		case []byte:
			return ChangeTypeFromByteArray(nil, v, dt)
		default:
			return nil, fmt.Errorf("cannot convert %T to %v", value, dt)
		}
	case enums.DataTypeTime:
		switch v := value.(type) {
		case types.GXTime:
			return v, nil
		case *types.GXTime:
			if v == nil {
				return nil, fmt.Errorf("value is nil")
			}
			return *v, nil
		case string:
			ret, err := types.NewGXTimeFromString(v, nil)
			if err != nil {
				return nil, err
			}
			return *ret, nil
		case []byte:
			return ChangeTypeFromByteArray(nil, v, dt)
		default:
			return nil, fmt.Errorf("cannot convert %T to %v", value, dt)
		}
	case enums.DataTypeEnum:
		if v, ok := value.(types.GXEnum); ok {
			return v, nil
		}
		if v, ok := value.(*types.GXEnum); ok {
			if v == nil {
				return nil, fmt.Errorf("value is nil")
			}
			return *v, nil
		}
		ret, err := Convert(value, enums.DataTypeUint8)
		if err != nil {
			return nil, err
		}
		return types.GXEnum{Value: ret.(uint8)}, nil
	case enums.DataTypeBitString:
		switch v := value.(type) {
		case types.GXBitString:
			return v, nil
		case *types.GXBitString:
			if v == nil {
				return nil, fmt.Errorf("value is nil")
			}
			return *v, nil
		case string:
			ret, err := types.NewGXBitStringFromString(v)
			if err != nil {
				return nil, err
			}
			return *ret, nil
		case []byte:
			ret, err := types.NewGXBitStringFromByteArray(v)
			if err != nil {
				return nil, err
			}
			return *ret, nil
		default:
			return nil, fmt.Errorf("cannot convert %T to %v", value, dt)
		}
	case enums.DataTypeArray:
		switch v := value.(type) {
		case types.GXArray:
			return v, nil
		case []any:
			return types.GXArray(v), nil
		default:
			return nil, fmt.Errorf("cannot convert %T to %v", value, dt)
		}
	case enums.DataTypeStructure:
		switch v := value.(type) {
		case types.GXStructure:
			return v, nil
		case []any:
			return types.GXStructure(v), nil
		default:
			return nil, fmt.Errorf("cannot convert %T to %v", value, dt)
		}
	default:
		dst, err := GetDataType(dt)
		if err != nil {
			return nil, err
		}

		v := reflect.ValueOf(value)
		src := v.Type()
		if src.AssignableTo(dst) {
			return v.Interface(), nil
		}
		if src.ConvertibleTo(dst) {
			return v.Convert(dst).Interface(), nil
		}

		switch s := value.(type) {
		case string:
			return convertFromString(s, dst)
		case []byte:
			return convertFromString(string(s), dst)
		}
		return nil, fmt.Errorf("cannot convert %v to %v", src, dst)
	}
}

func convertFromString(s string, dst reflect.Type) (any, error) {
	switch dst.Kind() {
	case reflect.String:
		return s, nil
	case reflect.Bool:
		ret, err := strconv.ParseBool(s)
		if err != nil {
			return nil, err
		}
		return ret, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		ret, err := strconv.ParseInt(s, 10, dst.Bits())
		if err != nil {
			return nil, err
		}
		out := reflect.New(dst).Elem()
		out.SetInt(ret)
		return out.Interface(), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		ret, err := strconv.ParseUint(s, 10, dst.Bits())
		if err != nil {
			return nil, err
		}
		out := reflect.New(dst).Elem()
		out.SetUint(ret)
		return out.Interface(), nil
	case reflect.Float32, reflect.Float64:
		ret, err := strconv.ParseFloat(s, dst.Bits())
		if err != nil {
			return nil, err
		}
		out := reflect.New(dst).Elem()
		out.SetFloat(ret)
		return out.Interface(), nil
	default:
		return nil, fmt.Errorf("cannot convert string to %v", dst)
	}
}
