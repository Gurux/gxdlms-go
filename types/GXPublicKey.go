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
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/Gurux/gxdlms-go/enums"
)

// Public key.
type GXPublicKey struct {
	// Used scheme.
	scheme enums.Ecc

	// Private key raw value.
	rawValue []byte

	// SystemTitle is an extra information that can be used in debugging.
	// SystemTitle is not serialized.
	SystemTitle []byte
}

// Used scheme.
func (g *GXPublicKey) Scheme() enums.Ecc {
	return g.scheme
}

// Private key raw value.
func (g *GXPublicKey) RawValue() []byte {
	return g.rawValue
}

// X returns the x Coordinate.
func (g *GXPublicKey) X() []byte {
	pk := GXByteBuffer{}
	pk.Set(g.rawValue)
	size := pk.Size() / 2
	val, err := pk.SubArray(1, size)
	if err != nil {
		return nil
	}
	return val
}

// Y returns the y Coordinate.
func (g *GXPublicKey) Y() []byte {
	pk := GXByteBuffer{}
	pk.Set(g.rawValue)
	size := pk.Size() / 2
	val, err := pk.SubArray(1+size, size)
	if err != nil {
		return nil
	}
	return val
}

// FromRawBytes returns the create the public key from raw bytes.
//
// Parameters:
//
//	key: Raw data
//
// Returns:
//
//	Public key.
func PublicKeyFromRawBytes(key []byte) (*GXPublicKey, error) {
	value := GXPublicKey{}
	if len(key) == 65 {
		value.scheme = enums.EccP256
		value.rawValue = key
	} else if len(key) == 97 {
		value.scheme = enums.EccP384
		value.rawValue = key
	} else if len(key) == 64 {
		value.scheme = enums.EccP256
		value.rawValue = make([]byte, 65)
		value.rawValue[0] = 4
		copy(value.rawValue, key[0:64])
	} else if len(key) == 96 {
		value.scheme = enums.EccP384
		value.rawValue = make([]byte, 96)
		value.rawValue[0] = 4
		copy(value.rawValue, key[0:95])
	} else {
		return nil, errors.New("Invalid public key.")
	}
	return &value, nil
}

// FromDer returns the create the public key from DER.
//
// Parameters:
//
//	der: DER Base64 coded string.
//
// Returns:
//
//	Public key.
func PublicKeyFromDer(der string) (*GXPublicKey, error) {
	der = strings.ReplaceAll(der, "\r\n", "")
	der = strings.ReplaceAll(der, "\n", "")
	value := GXPublicKey{}
	key, err := base64.URLEncoding.DecodeString(der)
	if err != nil {
		return nil, err
	}
	var tmp []any
	var id enums.X9ObjectIdentifier
	ret, err := Asn1FromByteArray(key)
	if err != nil {
		return nil, err
	}
	seq := *ret.(*GXAsn1Sequence)
	tmp = seq[0].([]any)
	id = X9ObjectIdentifierFromString(tmp[1].(string))
	switch id {
	case enums.X9ObjectIdentifierPrime256v1:
		value.scheme = enums.EccP256
	case enums.X9ObjectIdentifierSecp384r1:
		value.scheme = enums.EccP384
	default:
		if id == enums.X9ObjectIdentifierNone {
			return nil, fmt.Errorf("Invalid public key %d.", tmp[0])
		} else {
			return nil, fmt.Errorf("Invalid public key %s.", id.String())
		}
	}
	if _, ok := seq[1].([]byte); ok {
		value.rawValue = []byte(seq[1].([]byte))
	} else {
		bs := seq[1].(GXBitString)
		value.rawValue = bs.Value()
	}
	return &value, nil
}

// FromPem returns the create the public key from PEM.
//
// Parameters:
//
//	pem: PEM Base64 coded string.
//
// Returns:
//
//	Public key.
func PublicKeyFromPem(pem string) (*GXPublicKey, error) {
	pem = strings.ReplaceAll(pem, "\r\n", "\n")
	const START = "-----BEGIN PUBLIC KEY-----"
	const END = "-----END PUBLIC KEY-----"
	index := strings.Index(pem, START)
	if index == -1 {
		return nil, errors.New("Invalid PEM file.")
	}
	pem = pem[index+len(START):]
	index = strings.Index(pem, END)
	if index == -1 {
		return nil, errors.New("Invalid PEM file.")
	}
	tmp := pem[0:index]
	return PublicKeyFromDer(tmp)
}

// Load returns the create the public key from PEM file.
//
// Parameters:
//
//	path: Path to the PEM file.
//
// Returns:
//
//	Public key.
func PublicKeyLoad(path string) (*GXPublicKey, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return PublicKeyFromPem(string(data))
}

// Save returns the save public key to PEM file.
//
// Parameters:
//
//	path: File path.
func (g *GXPublicKey) Save(path string) error {
	ret, err := g.ToPem()
	if err != nil {
		return err
	}
	return os.WriteFile(path, []byte(ret), 0644)
}

// ToHex returns the the public key as a hex string.
func (g *GXPublicKey) ToHex() string {
	return ToHex(g.rawValue, true)
}

// ToDer returns the get public key as DER format.
func (g *GXPublicKey) ToDer() (string, error) {
	ret, err := g.ToEncoded()
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(ret), nil
}

// ToEncoded returns the get public key as encoded format.
func (g *GXPublicKey) ToEncoded() ([]byte, error) {
	// Subject Public Key Info.
	d := GXAsn1Sequence{}
	d1 := GXAsn1Sequence{}
	d1 = append(d1, NewGXAsn1ObjectIdentifier("1.2.840.10045.2.1"))
	switch g.scheme {
	case enums.EccP256:
		d1 = append(d1, NewGXAsn1ObjectIdentifier("1.2.840.10045.3.1.7"))
	case enums.EccP384:
		d1 = append(d1, NewGXAsn1ObjectIdentifier("1.3.132.0.34"))
	default:
		return nil, errors.New("Invalid ECC scheme.")
	}
	d = append(d, d1)
	bs, err := NewGXBitString(g.rawValue, 0)
	if err != nil {
		return nil, err
	}
	d = append(d, bs)
	return Asn1ToByteArray(d)
}

// ToPem returns the get public key as PEM format.
func (g *GXPublicKey) ToPem() (string, error) {
	der, err := g.ToDer()
	if err != nil {
		return "", err
	}
	return "-----BEGIN PUBLIC KEY-----\n" + der + "\n-----END PUBLIC KEY-----\n", nil
}

// Returns a string that represents the current elliptic curve public key, including the scheme and public coordinates.
func (g *GXPublicKey) String() string {
	sb := strings.Builder{}
	sb.WriteString("ECC Public Key:\n")
	switch g.scheme {
	case enums.EccP256:
		sb.WriteString("NIST P-256\n")
	case enums.EccP384:
		sb.WriteString("NIST P-384\n")
	default:
		return "Invalid scheme."
	}
	pk := GXByteBuffer{}
	pk.Set(g.rawValue)
	size := pk.Size() / 2
	sb.WriteString(" public x coord: ")
	tmp, _ := pk.SubArray(1, size)
	sb.WriteString(new(big.Int).SetBytes(tmp).String())
	sb.WriteString(" public y coord: ")
	tmp, _ = pk.SubArray(1+size, size)
	sb.WriteString(new(big.Int).SetBytes(tmp).String())
	return sb.String()
}

// Determines whether the specified object is equal to the current GXPublicKey instance.
func (g *GXPublicKey) Equals(obj any) bool {
	if o, ok := obj.(GXPublicKey); ok {
		return bytes.Equal(g.rawValue, o.rawValue)
	}
	return false
}
