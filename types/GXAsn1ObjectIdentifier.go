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
	"strconv"
	"strings"

	"github.com/Gurux/gxdlms-go/enums"
)

type GXAsn1ObjectIdentifier struct {
	objectIdentifier string
}

// ObjectIdentifier returns the unique identifier for the associated object.
func (g *GXAsn1ObjectIdentifier) ObjectIdentifier() string {
	return g.objectIdentifier
}

func (g *GXAsn1ObjectIdentifier) Encoded() ([]byte, error) {
	return g.OidStringtoBytes(g.objectIdentifier)
}

// NewGXAsn1ObjectIdentifier creates a new object identifier.
func NewGXAsn1ObjectIdentifier(objectIdentifier string) *GXAsn1ObjectIdentifier {
	return &GXAsn1ObjectIdentifier{objectIdentifier: objectIdentifier}
}

// NewGXAsn1ObjectIdentifier creates a new object identifier.
func NewGXAsn1ObjectIdentifierFromByteBuffer(bb *GXByteBuffer, count int) *GXAsn1ObjectIdentifier {
	objectIdentifier, err := oidStringFromByteArray(bb, count)
	if err != nil {
		return nil
	}
	return &GXAsn1ObjectIdentifier{objectIdentifier: objectIdentifier}
}

// Description returns a human-readable description of the object identifier.
func (g *GXAsn1ObjectIdentifier) Description() (string, error) {
	n := X509NameFromString(g.objectIdentifier)
	if n != enums.X509NameNone {
		return X509NameToString(n)
	}
	ha := HashAlgorithmFromString(g.objectIdentifier)
	if ha != enums.HashAlgorithmNone {
		return HashAlgorithmToString(ha)
	}
	oi := X9ObjectIdentifierFromString(g.objectIdentifier)
	if oi != enums.X9ObjectIdentifierNone {
		return X9ObjectIdentifierToString(oi)
	}
	tmp := PkcsObjectIdentifierFromString(g.objectIdentifier)
	if tmp != enums.PkcsObjectIdentifierNone {
		return PkcsObjectIdentifierToString(tmp)
	}
	ct := X509CertificateTypeFromString(g.objectIdentifier)
	if ct != enums.X509CertificateTypeNone {
		return X509CertificateTypeToString(ct)
	}
	return "", fmt.Errorf("Description not found for OID: %s", g.objectIdentifier)
}

// oidStringFromByteArray  returns the get OID string from bytes.
//
// Parameters:
//
//	bb: converted bytes.
//	len: byte count.
//
// Returns:
//
//	OID string.
func oidStringFromByteArray(bb *GXByteBuffer, len int) (string, error) {
	value := 0
	sb := strings.Builder{}
	if len != 0 {
		// Get first byte.
		tmp, err := bb.Uint8()
		if err != nil {
			return "", err
		}
		sb.WriteString(strconv.Itoa(int(tmp / 40)))
		sb.WriteString(".")
		sb.WriteString(strconv.Itoa(int(tmp % 40)))
		for pos := 1; pos != len; pos++ {
			tmp, err := bb.Uint8()
			if err != nil {
				return "", err
			}
			if (tmp & 0x80) != 0 {
				value += int(tmp & 0x7F)
				value <<= 7
			} else {
				value += int(tmp)
				sb.WriteString(".")
				sb.WriteString(strconv.Itoa(value))
				value = 0
			}
		}
	}
	return sb.String(), nil
}

// OidStringtoBytes returns the convert OID string to bytes.
func (g *GXAsn1ObjectIdentifier) OidStringtoBytes(oid string) ([]byte, error) {
	var value int64
	arr := strings.Split(strings.TrimSpace(oid), ".")
	// Make first byte.
	tmp := GXByteBuffer{}
	v0, err := strconv.Atoi(arr[0])
	if err != nil {
		return nil, err
	}
	v1, err := strconv.Atoi(arr[1])
	if err != nil {
		return nil, err
	}
	value = int64(v0)*40 + int64(v1)
	err = tmp.SetUint8(uint8(value))
	if err != nil {
		return nil, err
	}
	for pos := 2; pos != len(arr); pos++ {
		value, err := strconv.Atoi(arr[pos])
		if err != nil {
			return nil, err
		}
		if value < 0x80 {
			err = tmp.SetUint8(uint8(value))
			if err != nil {
				return nil, err
			}
		} else if value < 0x4000 {
			err = tmp.SetUint8(uint8((0x80 | (value >> 7))))
			if err != nil {
				return nil, err
			}
			err = tmp.SetUint8(uint8((value & 0x7F)))
			if err != nil {
				return nil, err
			}
		} else if value < 0x200000 {
			err = tmp.SetUint8(uint8((0x80 | (value >> 14))))
			if err != nil {
				return nil, err
			}
			err = tmp.SetUint8(uint8((0x80 | (value >> 7))))
			if err != nil {
				return nil, err
			}
			err = tmp.SetUint8(uint8((value & 0x7F)))
			if err != nil {
				return nil, err
			}
		} else if value < 0x10000000 {
			err = tmp.SetUint8(uint8((0x80 | (value >> 21))))
			if err != nil {
				return nil, err
			}
			err = tmp.SetUint8(uint8((0x80 | (value >> 14))))
			if err != nil {
				return nil, err
			}
			err = tmp.SetUint8(uint8((0x80 | (value >> 7))))
			if err != nil {
				return nil, err
			}
			err = tmp.SetUint8(uint8((value & 0x7F)))
			if err != nil {
				return nil, err
			}
		} else if value < 0x800000000 {
			err = tmp.SetUint8(uint8((0x80 | (value >> 49))))
			if err != nil {
				return nil, err
			}
			err = tmp.SetUint8(uint8((0x80 | (value >> 42))))
			if err != nil {
				return nil, err
			}
			err = tmp.SetUint8(uint8((0x80 | (value >> 35))))
			if err != nil {
				return nil, err
			}
			err = tmp.SetUint8(uint8((0x80 | (value >> 28))))
			if err != nil {
				return nil, err
			}
			err = tmp.SetUint8(uint8((0x80 | (value >> 21))))
			if err != nil {
				return nil, err
			}
			err = tmp.SetUint8(uint8((0x80 | (value >> 14))))
			if err != nil {
				return nil, err
			}
			err = tmp.SetUint8(uint8((0x80 | (value >> 7))))
			if err != nil {
				return nil, err
			}
			err = tmp.SetUint8(uint8((value & 0x7F)))
			if err != nil {
				return nil, err
			}
		} else {
			return nil, errors.New("Invalid OID.")
		}
	}
	return tmp.Array(), nil
}

// String returns OID as string.
func (g *GXAsn1ObjectIdentifier) String() string {
	return g.objectIdentifier
}
