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
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/Gurux/gxdlms-go/enums"
)

// GXEcdsa provides helpers for ECDSA signing, verification, and key agreement.
type GXEcdsa struct {
	// Public key.
	publicKey *ecdsa.PublicKey

	// Private key.
	privateKey *ecdsa.PrivateKey
}

// NewGXEcdsaFromPublicKey creates an ECDSA helper from an existing public key.
//
// The returned object can be used to verify signatures.
func NewGXEcdsaFromPublicKey(key *ecdsa.PublicKey) (*GXEcdsa, error) {
	ret := GXEcdsa{}
	ret.publicKey = key
	return &ret, nil
}

// NewGXEcdsaFromPrivateKey creates an ECDSA helper from a private key.
//
// The returned object can be used to sign and verify data.
func NewGXEcdsaFromPrivateKey(key *ecdsa.PrivateKey) (*GXEcdsa, error) {
	ret := GXEcdsa{}
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

func PublicKeyFromECDSAPrivate(priv *ecdsa.PrivateKey) (*ecdsa.PublicKey, error) {
	ecdhPriv, err := priv.ECDH()
	if err != nil {
		return nil, err
	}

	raw := ecdhPriv.PublicKey().Bytes()
	return ecdsa.ParseUncompressedPublicKey(priv.Curve, raw)
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
	pub, err := PublicKeyFromECDSAPrivate(g.privateKey)
	if err != nil {
		return nil, err
	}
	if pub == nil {
		return nil, errors.New("invalid public key")
	}
	var digest []byte
	switch g.privateKey.Curve.Params().BitSize {
	case 256:
		sum := sha256.Sum256(data)
		digest = sum[:]
	case 384:
		sum := sha512.Sum384(data)
		digest = sum[:]
	default:
		return nil, errors.New("unsupported ECC scheme")
	}
	sig, err := ecdsa.SignASN1(rand.Reader, g.privateKey, digest)
	if err != nil {
		return nil, err
	}
	return sig, nil
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
func (g *GXEcdsa) GenerateSecret(publicKey *ecdsa.PublicKey) ([]byte, error) {
	if g.privateKey == nil {
		return nil, errors.New("Invalid private key.")
	}
	if publicKey == nil {
		return nil, errors.New("invalid public key")
	}
	if g.privateKey.Curve.Params().BitSize != publicKey.Curve.Params().BitSize {
		return nil, errors.New("Private key scheme is different than public key.")
	}
	ecdhPriv, err := g.privateKey.ECDH()
	if err != nil {
		return nil, err
	}
	ecdhPub, err := publicKey.ECDH()
	if err != nil {
		return nil, err
	}
	secret, err := ecdhPriv.ECDH(ecdhPub)
	if err != nil {
		return nil, err
	}
	return secret, nil
}

// GXEcdsaGenerateKeyPair generates a new ECDSA public/private key pair for the given curve scheme.
//
// Returns:
//
//	The generated public and private keys (wrapped in a key/value pair) or an error.
func GXEcdsaGenerateKeyPair(scheme enums.Ecc) (*GXKeyValuePair[*ecdsa.PublicKey, *ecdsa.PrivateKey], error) {

	var curve elliptic.Curve
	switch scheme {
	case enums.EccP256:
		curve = elliptic.P256()
	case enums.EccP384:
		curve = elliptic.P384()
	}
	pk, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, err
	}
	pub, err := PublicKeyFromECDSAPrivate(pk)
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
	if g.publicKey == nil {
		if g.privateKey == nil {
			return false, fmt.Errorf("invalid private key")
		}
		var err error
		g.publicKey, err = PublicKeyFromECDSAPrivate(g.privateKey)
		if err != nil {
			return false, err
		}
	}
	var digest []byte
	var size int
	sc, err := PublicKeyScheme(g.publicKey)
	if err != nil {
		return false, err
	}
	switch sc {
	case enums.EccP256:
		sum := sha256.Sum256(data)
		//Convert array to slice ([]byte).
		digest = sum[:]
		size = 32
	case enums.EccP384:
		sum := sha512.Sum384(data)
		//Convert array to slice ([]byte).
		digest = sum[:]
		size = 48
	default:
		return false, errors.New("unsupported ECC scheme")
	}
	r := new(big.Int).SetBytes(signature[:size])
	s := new(big.Int).SetBytes(signature[size:])
	ok := ecdsa.Verify(g.publicKey, digest, r, s)
	return ok, nil
}

func PublicKeyScheme(key *ecdsa.PublicKey) (enums.Ecc, error) {
	switch key.Curve.Params().BitSize {
	case 256:
		return enums.EccP256, nil
	case 384:
		return enums.EccP384, nil
	default:
		return enums.EccP256, errors.New("unsupported ECC scheme")
	}
}

func PrivateKeyScheme(key *ecdsa.PrivateKey) (enums.Ecc, error) {
	switch key.Curve.Params().BitSize {
	case 256:
		return enums.EccP256, nil
	case 384:
		return enums.EccP384, nil
	default:
		return enums.EccP256, errors.New("unsupported ECC scheme")
	}
}

func PrivateKeyFromRawBytes(raw []byte) (*ecdsa.PrivateKey, error) {
	var curve elliptic.Curve
	switch 8 * len(raw) {
	case 256:
		curve = elliptic.P256()
	case 384:
		curve = elliptic.P384()
	default:
		return nil, errors.New("unsupported ECC scheme")
	}
	d := new(big.Int).SetBytes(raw)
	priv := &ecdsa.PrivateKey{
		D: d,
	}
	priv.Curve = curve
	ret, err := PublicKeyFromECDSAPrivate(priv)
	if err != nil {
		return nil, err
	}
	priv.PublicKey = *ret
	return priv, nil
}

func PublicKeyFromRawBytes(raw []byte) (*ecdsa.PublicKey, error) {
	var curve elliptic.Curve
	switch len(raw) {
	case 65, 64:
		//Compression tag is not send in DLMS messages and size is 64.
		curve = elliptic.P256()
	case 97, 96:
		//Compression tag is not send in DLMS messages and size is 96.
		curve = elliptic.P384()
	default:
		return nil, errors.New("unsupported ECC scheme")
	}
	pub, err := ecdsa.ParseUncompressedPublicKey(curve, raw)
	if err != nil {
		return nil, errors.New("public key validate failed. public key is not valid ECDSA public key")
	}
	return pub, nil
}

func PrivateKeyToBytes(priv *ecdsa.PrivateKey) []byte {
	size := (priv.Curve.Params().BitSize + 7) / 8
	b := priv.D.Bytes()
	data := make([]byte, size)
	copy(data[size-len(b):], b)
	return data
}

func PublicKeyToBytes(pub *ecdsa.PublicKey) []byte {
	data := pub.X.Bytes()
	size := (pub.Curve.Params().BitSize + 7) / 8
	if len(data) < size {
		newData := make([]byte, size)
		copy(newData[size-len(data):], data)
		data = newData
	}
	return data
}

func PrivateKeyToHex(value *ecdsa.PrivateKey) string {
	if value == nil {
		return ""
	}
	size := (value.Curve.Params().BitSize + 7) / 8
	b := value.D.Bytes()
	data := make([]byte, size)
	copy(data[size-len(b):], b)
	return strings.ToUpper(hex.EncodeToString(data))
}

func PublicKeyToHex(value *ecdsa.PublicKey) string {
	if value == nil {
		return ""
	}
	data := value.X.Bytes()
	size := (value.Curve.Params().BitSize + 7) / 8
	if len(data) < size {
		newData := make([]byte, size)
		copy(newData[size-len(data):], data)
		data = newData
	}
	return strings.ToUpper(hex.EncodeToString(data))
}
