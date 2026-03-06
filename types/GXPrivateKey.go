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

// GXPrivateKey represents an EC private key used in DLMS/COSEM security.
//
// It includes a referenced public key (when available) and optional system title
// metadata used mainly for debugging.
//
// The private key bytes are stored in rawValue, and scheme indicates whether the key is
// P-256 or P-384.
type GXPrivateKey struct {
	// Used scheme.
	scheme enums.Ecc

	// Private key raw value.
	rawValue []byte

	publicKey *GXPublicKey

	// SystemTitle is extra information that can be used in debugging.
	// It is not serialized.
	SystemTitle []byte
}

// Scheme returns the ECC curve scheme used by the key.
func (g *GXPrivateKey) Scheme() enums.Ecc {
	return g.scheme
}

// RawValue returns the private key bytes.
func (g *GXPrivateKey) RawValue() []byte {
	return g.rawValue
}

// PrivateKeyFromRawBytes creates a GXPrivateKey from raw private key bytes.
//
// The length of key must be either 32 (P-256) or 48 (P-384).
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

// PrivateKeyFromDer parses a DER-encoded (base64) private key and returns a GXPrivateKey.
//
// The DER payload is expected to follow PKCS#8 private key encoding.
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

// PrivateKeyFromPem parses a PEM-encoded private key string and returns a GXPrivateKey.
//
// The input should include the "-----BEGIN PRIVATE KEY-----" and "-----END PRIVATE KEY-----" markers.
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

// PrivateKeyLoad reads a PEM file from disk and returns the decoded GXPrivateKey.
func PrivateKeyLoad(path string) (*GXPrivateKey, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return PrivateKeyFromPem(string(data))
}

// Save writes the private key to a PEM file at the provided path.
func (g *GXPrivateKey) Save(path string) error {
	ret, err := g.ToPem()
	if err != nil {
		return err
	}
	return os.WriteFile(path, []byte(ret), 0644)
}

// ToDer encodes the private key in DER (PKCS#8) format and returns it as a base64 string.
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

// ToPem returns the private key in PEM format.
func (g *GXPrivateKey) ToPem() (string, error) {
	ret, err := g.ToDer()
	if err != nil {
		return "", err
	}
	return "-----BEGIN EC PRIVATE KEY-----\n" + ret + "\n-----END EC PRIVATE KEY-----", nil
}

// GetPublicKey returns the corresponding public key for this private key.
//
// The public key is derived from the private key if it has not already been computed.
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

// ToHex returns the private key bytes as a hexadecimal string.
func (g *GXPrivateKey) ToHex() string {
	return ToHex(g.rawValue, true)
}

// Equals reports whether obj represents the same private key.
func (g *GXPrivateKey) Equals(obj any) bool {
	if o, ok := obj.(GXPrivateKey); ok {
		return bytes.Equal(g.rawValue, o.rawValue)
	}
	return false
}

// String implements fmt.Stringer by returning the key as a hex string.
func (g *GXPrivateKey) String() string {
	return g.ToHex()
}
