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
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"fmt"
	"math/big"

	"github.com/Gurux/gxdlms-go/enums"
)

// GXEcdsa provides helpers for ECDSA signing, verification, and key agreement.
type GXEcdsa struct {
	// Public key.
	publicKey *GXPublicKey

	// Private key.
	privateKey *GXPrivateKey

	curve *gxCurve
}

// NewGXEcdsaFromPublicKey creates an ECDSA helper from an existing public key.
//
// The returned object can be used to verify signatures.
func NewGXEcdsaFromPublicKey(key *GXPublicKey) (*GXEcdsa, error) {
	ret := GXEcdsa{}
	var err error
	ret.curve, err = newGxCurve(key.scheme)
	if err != nil {
		return nil, err
	}
	ret.publicKey = key
	return &ret, nil
}

// NewGXEcdsaFromPrivateKey creates an ECDSA helper from a private key.
//
// The returned object can be used to sign and verify data.
func NewGXEcdsaFromPrivateKey(key *GXPrivateKey) (*GXEcdsa, error) {
	ret := GXEcdsa{}
	var err error
	ret.curve, err = newGxCurve(key.scheme)
	if err != nil {
		return nil, err
	}
	ret.privateKey = key
	return &ret, nil
}

// schemeSize returns the key size in bytes for the given ECDSA curve scheme.
func schemeSize(scheme enums.Ecc) int {
	if scheme == enums.EccP256 {
		return 32
	}
	return 48
}

// getRandomNumber returns a cryptographically secure random integer for the given curve scheme.
//
// Parameters:
//
//	scheme: Curve scheme.
//
// Returns:
//
//	A random big integer and any error encountered.
func getRandomNumber(scheme enums.Ecc) (*big.Int, error) {
	size := schemeSize(scheme)
	bytes := make([]byte, size)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}
	return new(big.Int).SetBytes(bytes), nil
}

// Sign computes an ECDSA signature over the provided data using the private key.
//
// The signature is returned as raw bytes in the form (r || s).
//
// Parameters:
//
//	data: Data to sign.
//
// Returns:
//
//	Signature bytes or an error.
func (g *GXEcdsa) Sign(data []byte) ([]byte, error) {
	if g.privateKey == nil {
		return nil, fmt.Errorf("invalid private key")
	}
	var msg *big.Int
	switch g.privateKey.scheme {
	case enums.EccP256:
		hash := sha256.Sum256(data)
		msg = new(big.Int).SetBytes(hash[:])
	case enums.EccP384:
		hash := sha512.Sum384(data)
		msg = new(big.Int).SetBytes(hash[:])
	default:
		return nil, fmt.Errorf("invalid private key scheme")
	}

	// R = k * G, r = R[x]
	k, err := getRandomNumber(g.privateKey.scheme)
	if err != nil {
		return nil, err
	}

	pk := new(big.Int).SetBytes(g.privateKey.rawValue)
	R := gxEccPoint{x: new(big.Int), y: new(big.Int)}
	shamirsPointMulti(g.curve, &R, &g.curve.g, k)

	r := new(big.Int).Set(R.x)
	r.Mod(r, g.curve.n)

	// s = (k^-1 * (e + d * r)) mod n
	s := new(big.Int).Set(pk)
	s.Mul(s, r)
	s.Add(s, msg)

	kinv := new(big.Int).ModInverse(k, g.curve.n)
	if kinv == nil {
		return nil, fmt.Errorf("failed to compute modular inverse")
	}
	s.Mul(s, kinv)
	s.Mod(s, g.curve.n)

	signature := GXByteBuffer{}
	err = signature.Set(r.Bytes())
	if err != nil {
		return nil, err
	}
	err = signature.Set(s.Bytes())
	if err != nil {
		return nil, err
	}
	return signature.Array(), nil
}

// GenerateSecret computes a shared secret using ECDH between the private key and the provided public key.
//
// Parameters:
//
//	publicKey: Public key to use in the key agreement.
//
// Returns:
//
//	Shared secret bytes or an error.
func (g *GXEcdsa) GenerateSecret(publicKey *GXPublicKey) ([]byte, error) {
	if g.privateKey == nil {
		return nil, errors.New("Invalid private key.")
	}
	if g.privateKey.scheme != publicKey.Scheme() {
		return nil, errors.New("Private key scheme is different than public key.")
	}
	pk := new(big.Int).SetBytes(g.privateKey.rawValue)
	bb := GXByteBuffer{}
	err := bb.Set(publicKey.RawValue())
	if err != nil {
		return nil, err
	}
	size := schemeSize(g.privateKey.scheme)
	v, err := bb.SubArray(1, size)
	if err != nil {
		return nil, err
	}
	x := new(big.Int).SetBytes(v)
	v, err = bb.SubArray(1+size, size)
	if err != nil {
		return nil, err
	}
	y := new(big.Int).SetBytes(v)
	p := &gxEccPoint{x, y}
	curve, err := newGxCurve(g.privateKey.scheme)
	if err != nil {
		return nil, err
	}
	ret := &gxEccPoint{}
	shamirsPointMulti(curve, ret, p, pk)
	return ret.x.Bytes(), nil
}

// GXEcdsaGenerateKeyPair generates a new ECDSA public/private key pair for the given curve scheme.
//
// Returns:
//
//	The generated public and private keys (wrapped in a key/value pair) or an error.
func GXEcdsaGenerateKeyPair(scheme enums.Ecc) (*GXKeyValuePair[*GXPublicKey, *GXPrivateKey], error) {
	raw, err := getRandomNumber(scheme)
	if err != nil {
		return nil, err
	}
	pk, err := PrivateKeyFromRawBytes(raw.Bytes())
	if err != nil {
		return nil, err
	}
	pub, err := pk.GetPublicKey()
	if err != nil {
		return nil, err
	}
	return NewGXKeyValuePair(pub, pk), nil
}

// Verify checks whether the given signature is valid for the provided data.
//
// Parameters:
//
//	signature: Signature bytes in (r || s) form.
//	data: Data that was signed.
//
// Returns:
//
//	True if the signature is valid; otherwise false.
func (g *GXEcdsa) Verify(signature []byte, data []byte) (bool, error) {
	var msg *big.Int
	if g.publicKey == nil {
		if g.privateKey == nil {
			return false, fmt.Errorf("invalid private key")
		}
		var err error
		g.publicKey, err = g.privateKey.GetPublicKey()
		if err != nil {
			return false, err
		}
	}
	if g.publicKey.scheme == enums.EccP256 {
		hash := sha256.Sum256(data)
		msg = new(big.Int).SetBytes(hash[:])
	} else {
		hash := sha512.Sum384(data)
		msg = new(big.Int).SetBytes(hash[:])
	}

	bb := GXByteBuffer{}
	err := bb.Set(signature)
	if err != nil {
		return false, err
	}

	size := schemeSize(g.publicKey.scheme)
	sigRBytes, err := bb.SubArray(0, size)
	if err != nil {
		return false, err
	}
	sigR := new(big.Int).SetBytes(sigRBytes)

	sigSBytes, err := bb.SubArray(size, size)
	if err != nil {
		return false, err
	}
	sigS := new(big.Int).SetBytes(sigSBytes)

	w := new(big.Int).Set(sigS)
	w.ModInverse(w, g.curve.n)
	if w == nil {
		return false, fmt.Errorf("failed to compute modular inverse")
	}
	u1 := new(big.Int).Set(msg)
	u1.Mul(u1, w)
	u1.Mod(u1, g.curve.n)
	u2 := new(big.Int).Set(sigR)
	u2.Mul(u2, w)
	u2.Mod(u2, g.curve.n)
	tmp := gxEccPoint{x: new(big.Int), y: new(big.Int)}
	shamirsTrick(g.curve, g.publicKey, &tmp, u1, u2)
	tmp.x.Mod(tmp.x, g.curve.n)
	return tmp.x.Cmp(sigR) == 0, nil
}

// EcdsaValidate checks that the given public key lies on the expected elliptic curve.
//
// This can be used to verify that a public key is valid for the curve implied by its scheme.
func EcdsaValidate(publicKey *GXPublicKey) error {
	if publicKey == nil {
		return errors.New("Invalid public key.")
	}
	bb := GXByteBuffer{}
	err := bb.Set(publicKey.RawValue())
	if err != nil {
		return err
	}
	size := schemeSize(publicKey.Scheme())
	v, err := bb.SubArray(1, size)
	if err != nil {
		return err
	}
	x := new(big.Int).SetBytes(v)
	v, err = bb.SubArray(1+size, size)
	if err != nil {
		return err
	}
	y := new(big.Int).SetBytes(v)
	curve, err := newGxCurve(publicKey.Scheme())
	if err != nil {
		return err
	}
	y.Mul(y, y)
	y.Mod(y, curve.p)
	tmpX := new(big.Int).Set(x)
	tmpX.Mul(x, x)
	tmpX.Mod(tmpX, curve.p)
	tmpX.Add(tmpX, curve.a)
	tmpX.Mul(x, x)
	tmpX.Add(tmpX, curve.b)
	tmpX.Mod(tmpX, curve.p)
	if y.Cmp(tmpX) != 0 {
		return errors.New("Public key validate failed. Public key is not valid ECDSA public key.")
	}
	return nil
}
