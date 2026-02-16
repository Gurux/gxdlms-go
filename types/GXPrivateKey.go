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

// Private key.
type GXPrivateKey struct {
	// Used scheme.
	scheme enums.Ecc

	// Private key raw value.
	rawValue []byte

	publicKey *GXPublicKey

	// SystemTitle is an extra information that can be used in debugging.
	// SystemTitle is not serialized.
	SystemTitle []byte
}

// Used scheme.
func (g *GXPrivateKey) Scheme() enums.Ecc {
	return g.scheme
}

// Private key raw value.
func (g *GXPrivateKey) RawValue() []byte {
	return g.rawValue
}

// FromRawBytes returns the create the private key from raw bytes.
//
// Parameters:
//
//	key: Raw data
//
// Returns:
//
//	Private key.
func PrivateKeyFromRawBytes(key []byte) (*GXPrivateKey, error) {
	value := GXPrivateKey{}
	// If private key is given
	if len(key) == 32 {
		value.scheme = enums.EccP256
		value.rawValue = key
	} else if len(key) == 48 {
		value.scheme = enums.EccP384
		value.rawValue = key
	} else {
		return nil, errors.New("Invalid private key.")
	}
	return &value, nil
}

// FromDer returns the create the private key from DER.
//
// Parameters:
//
//	key: DER Base64 coded string.
//
// Returns:
//
//	Private key.
func PrivateKeyFromDer(der string) (*GXPrivateKey, error) {
	der = strings.ReplaceAll(der, "\r\n", "")
	der = strings.ReplaceAll(der, "\n", "")
	key, err := base64.StdEncoding.DecodeString(der)
	if err != nil {
		return nil, err
	}
	ret, err := Asn1FromByteArray(key)
	if err != nil {
		return nil, err
	}
	seq := *ret.(*GXAsn1Sequence)
	if seq[0].(int8) > 3 {
		return nil, errors.New("Invalid private key version.")
	}
	tmp := seq[2].([]any)
	value := GXPrivateKey{}
	id := X9ObjectIdentifierFromString(tmp[0].(string))
	switch id {
	case enums.X9ObjectIdentifierPrime256v1:
		value.scheme = enums.EccP256
	case enums.X9ObjectIdentifierSecp384r1:
		value.scheme = enums.EccP384
	default:
		if id == enums.X9ObjectIdentifierNone {
			return nil, fmt.Errorf("Invalid private key %s.", tmp[0].(string))
		} else {
			return nil, fmt.Errorf("Invalid private key %s %s.", id.String(), tmp[0].(string))
		}
	}
	value.rawValue = seq[1].([]byte)
	if v1, ok := seq[3].([]byte); ok {
		value.publicKey, err = PublicKeyFromRawBytes(v1)
		if err != nil {
			return nil, err
		}
	} else if v2, ok := seq[3].(*GXBitString); ok {
		value.publicKey, err = PublicKeyFromRawBytes(v2.Value())
		if err != nil {
			return nil, err
		}
	} else {
		v := seq[3].([]any)
		bs := (v[0]).(*GXBitString)
		value.publicKey, err = PublicKeyFromRawBytes(bs.Value())
		if err != nil {
			return nil, err
		}
	}
	return &value, nil
}

// FromPem returns the create the private key from PEM.
//
// Parameters:
//
//	pem: PEM in Base64 coded string.
//
// Returns:
//
//	Private key.
func PrivateKeyFromPem(pem string) (*GXPrivateKey, error) {
	pem = strings.ReplaceAll(pem, "\r\n", "\n")
	const START = "-----BEGIN PRIVATE KEY-----"
	const END = "-----END PRIVATE KEY-----"
	index := strings.Index(pem, START)
	if index == -1 {
		return nil, errors.New("Invalid PEM file.")
	}
	pem = pem[index+len(START):]
	index = strings.Index(pem, END)
	if index == -1 {
		return nil, errors.New("Invalid PEM file.")
	}
	return PrivateKeyFromDer(pem[0:index])
}

// PrivateKeyLoad returns the create the private key from PEM file.
//
// Parameters:
//
//	path: Path to the PEM file.
//
// Returns:
//
//	Private key.
func PrivateKeyLoad(path string) (*GXPrivateKey, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return PrivateKeyFromPem(string(data))
}

// Save returns the save private key to PEM file.
//
// Parameters:
//
//	path: File path.
func (g *GXPrivateKey) Save(path string) error {
	ret, err := g.ToPem()
	if err != nil {
		return err
	}
	return os.WriteFile(path, []byte(ret), 0644)
}

// ToDer returns the get private key as DER format.
func (g *GXPrivateKey) ToDer() (string, error) {
	d := GXAsn1Sequence{}
	d = append(d, int8(enums.CertificateVersionVersion2))
	d = append(d, g.rawValue)
	d1 := GXAsn1Sequence{}
	switch g.scheme {
	case enums.EccP256:
		d1 = append(d1, NewGXAsn1ObjectIdentifier("1.2.840.10045.3.1.7"))
	case enums.EccP384:
		d1 = append(d1, NewGXAsn1ObjectIdentifier("1.3.132.0.34"))
	default:
		return "", errors.New("Invalid ECC scheme.")
	}
	d = append(d, d1)
	d2 := GXAsn1Context{Index: 1}
	pk, err := g.GetPublicKey()
	if err != nil {
		return "", err
	}
	bs, err := NewGXBitString(pk.RawValue(), 0)
	if err != nil {
		return "", err
	}
	d2.Items = append(d2.Items, bs)
	d = append(d, d2)
	v, err := Asn1ToByteArray(d)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(v), nil
}

// ToPem returns the get private key as PEM format.
func (g *GXPrivateKey) ToPem() (string, error) {
	ret, err := g.ToDer()
	if err != nil {
		return "", err
	}
	return "-----BEGIN EC PRIVATE KEY-----\n" + ret + "\n-----END EC PRIVATE KEY-----", nil
}

// GetPublicKey returns the get public key from private key.
//
// Returns:
//
//	Public key.
func (g *GXPrivateKey) GetPublicKey() (*GXPublicKey, error) {
	if g.publicKey == nil {
		// Public key = private key multiple by curve.G.
		pk := new(big.Int).SetBytes(g.rawValue)
		curve, err := newGxCurve(g.scheme)
		if err != nil {
			return nil, err
		}
		ret := &gxEccPoint{}
		shamirsPointMulti(curve, ret, &curve.g, pk)
		size := 0
		if g.scheme == enums.EccP256 {
			size = 32
		} else {
			size = 48
		}
		key := NewGXByteBufferWithCapacity(65)
		tmp := ret.x.Bytes()
		err = key.SetAt(tmp, len(tmp)%size, size)
		if err != nil {
			return nil, err
		}
		//Public key is un-compressed format.
		key.SetUint8(4)
		tmp = ret.y.Bytes()
		err = key.SetAt(tmp, len(tmp)%size, size)
		g.publicKey, err = PublicKeyFromRawBytes(key.Array())
		if err != nil {
			return nil, err
		}
	}
	return g.publicKey, nil
}

// ToHex returns the the private key as a hex string.
//
// Returns:
//
//	Private key as hex string.
func (g *GXPrivateKey) ToHex() string {
	return ToHex(g.rawValue, true)
}

// Determines whether the specified object is equal to the current GXPrivateKey instance.
func (g *GXPrivateKey) Equals(obj any) bool {
	if o, ok := obj.(GXPrivateKey); ok {
		return bytes.Equal(g.rawValue, o.rawValue)
	}
	return false
}

// Returns a string that represents the current elliptic curve private key.
func (g *GXPrivateKey) String() string {
	return g.ToHex()
}
