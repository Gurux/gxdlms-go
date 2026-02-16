package helpers

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
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Compare(arr1 []byte, arr2 []byte) bool {
	return bytes.Equal(arr1, arr2)
}

func GetObjectCountSizeInBytes(count int) int {
	switch {
	case count < 0x80:
		return 1
	case count < 0x100:
		return 2
	case count < 0x10000:
		return 3
	default:
		return 5
	}
}

func ToLogicalName(value any) (string, error) {
	switch v := value.(type) {
	case string:
		return v, nil
	case []byte:
		if len(v) != 6 {
			return "", fmt.Errorf("logical name length is %d, expected 6", len(v))
		}
		return fmt.Sprintf("%d.%d.%d.%d.%d.%d", v[0], v[1], v[2], v[3], v[4], v[5]), nil
	case [6]byte:
		return fmt.Sprintf("%d.%d.%d.%d.%d.%d", v[0], v[1], v[2], v[3], v[4], v[5]), nil
	default:
		return "", fmt.Errorf("unsupported logical name value: %T", value)
	}
}

func LogicalNameToBytes(value string) ([]byte, error) {
	if value == "" {
		return make([]byte, 6), nil
	}
	parts := strings.Split(value, ".")
	if len(parts) != 6 {
		return nil, errors.New("invalid logical name")
	}
	ret := make([]byte, 6)
	for i, p := range parts {
		n, err := strconv.Atoi(strings.TrimSpace(p))
		if err != nil || n < 0 || n > 255 {
			return nil, fmt.Errorf("invalid logical name part %q", p)
		}
		ret[i] = byte(n)
	}
	return ret, nil
}
