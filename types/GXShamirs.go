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
	"math/big"
)

// GXShamirs implements Shamir's trick for elliptic curve operations.

// isBitSet checks if a bit is set at the given position in a big integer.
func isBitSet(n *big.Int, pos uint) bool {
	return n.Bit(int(pos)) == 1
}

// usedBits returns the number of bits used in a big integer.
func usedBits(n *big.Int) uint16 {
	return uint16(n.BitLen())
}

// Trick implements Shamir's trick for efficient point multiplication.
//
// Parameters:
//
//	curve: Used curve.
//	pub: Public key.
//	ret: Result point.
//	u1: First scalar.
//	u2: Second scalar.
func shamirsTrick(curve *gxCurve, pub *GXPublicKey, ret *gxEccPoint, u1, u2 *big.Int) {
	sum := gxEccPoint{x: new(big.Int), y: new(big.Int)}
	op2 := gxEccPoint{x: new(big.Int).SetBytes(pub.X()), y: new(big.Int).SetBytes(pub.Y())}
	shamirsPointAdd(curve, &sum, &curve.g, &op2)
	bits1 := usedBits(u1)
	bits2 := usedBits(u2)
	var pos uint16
	if bits1 > bits2 {
		pos = bits1
	} else {
		pos = bits2
	}
	pos--

	if isBitSet(u1, uint(pos)) && isBitSet(u2, uint(pos)) {
		ret.x.Set(sum.x)
		ret.y.Set(sum.y)
	} else if isBitSet(u1, uint(pos)) {
		ret.x.Set(curve.g.x)
		ret.y.Set(curve.g.y)
	} else if isBitSet(u2, uint(pos)) {
		ret.x.SetBytes(pub.X())
		ret.y.SetBytes(pub.Y())
	}

	tmp := gxEccPoint{x: new(big.Int), y: new(big.Int)}
	pos--

	for {
		tmp.x.Set(ret.x)
		tmp.y.Set(ret.y)
		shamirsPointDouble(curve, ret, &tmp)
		tmp.x.Set(ret.x)
		tmp.y.Set(ret.y)

		if isBitSet(u1, uint(pos)) && isBitSet(u2, uint(pos)) {
			shamirsPointAdd(curve, ret, &tmp, &sum)
		} else if isBitSet(u1, uint(pos)) {
			shamirsPointAdd(curve, ret, &tmp, &curve.g)
		} else if isBitSet(u2, uint(pos)) {
			shamirsPointAdd(curve, ret, &tmp, &op2)
		}

		if pos == 0 {
			break
		}
		pos--
	}
}

// PointAdd adds two elliptic curve points.
//
// Parameters:
//
//	curve: Used curve.
//	ret: Result point.
//	p1: First point.
//	p2: Second point.
func shamirsPointAdd(curve *gxCurve, ret, p1, p2 *gxEccPoint) {
	// Calculate lambda.
	ydiff := new(big.Int).Set(p2.y)
	ydiff.Sub(ydiff, p1.y)

	xdiff := new(big.Int).Set(p2.x)
	xdiff.Sub(xdiff, p1.x)
	xdiff.ModInverse(xdiff, curve.p)

	lambda := new(big.Int).Set(ydiff)
	lambda.Mul(lambda, xdiff)
	lambda.Mod(lambda, curve.p)
	// Calculate resulting x coord.
	ret.x.Set(lambda)
	ret.x.Mul(ret.x, lambda)
	ret.x.Sub(ret.x, p1.x)
	ret.x.Sub(ret.x, p2.x)
	ret.x.Mod(ret.x, curve.p)
	// Calculate resulting y coord.
	ret.y.Set(p1.x)
	ret.y.Sub(ret.y, ret.x)
	ret.y.Mul(ret.y, lambda)
	ret.y.Sub(ret.y, p1.y)
	ret.y.Mod(ret.y, curve.p)
}

// PointDouble doubles an elliptic curve point.
//
// Parameters:
//
//	curve: Used curve.
//	ret: Result value.
//	p1: Point to double.
func shamirsPointDouble(curve *gxCurve, ret, p1 *gxEccPoint) {
	numer := new(big.Int).Set(p1.x)
	numer.Mul(numer, p1.x)
	numer.Mul(numer, big.NewInt(3))
	numer.Add(numer, curve.a)

	denom := new(big.Int).Set(p1.y)
	denom.Mul(denom, big.NewInt(2))
	denom.ModInverse(denom, curve.p)

	lambda := new(big.Int).Set(numer)
	lambda.Mul(lambda, denom)
	lambda.Mod(lambda, curve.p)
	// Calculate resulting x coord.
	ret.x.Set(lambda)
	ret.x.Mul(ret.x, lambda)
	ret.x.Sub(ret.x, p1.x)
	ret.x.Sub(ret.x, p1.x)
	ret.x.Mod(ret.x, curve.p)
	// Calculate resulting y coord.
	ret.y.Set(p1.x)
	ret.y.Sub(ret.y, ret.x)
	ret.y.Mul(ret.y, lambda)
	ret.y.Sub(ret.y, p1.y)
	ret.y.Mod(ret.y, curve.p)
}

// shamirsPointMulti multiplies an elliptic curve point with a big integer value.
//
// Parameters:
//
//	curve: Used curve.
//	ret: Return value.
//	point: Point to multiply.
//	scalar: Scalar multiplier.
func shamirsPointMulti(curve *gxCurve, ret, point *gxEccPoint, scalar *big.Int) {
	R0 := gxEccPoint{x: new(big.Int).Set(point.x), y: new(big.Int).Set(point.y)}
	R1 := gxEccPoint{x: new(big.Int), y: new(big.Int)}
	tmp := gxEccPoint{x: new(big.Int), y: new(big.Int)}

	shamirsPointDouble(curve, &R1, point)

	dbits := usedBits(scalar)
	dbits -= 2

	for {
		if isBitSet(scalar, uint(dbits)) {
			tmp.x.Set(R0.x)
			tmp.y.Set(R0.y)
			shamirsPointAdd(curve, &R0, &R1, &tmp)
			tmp.x.Set(R1.x)
			tmp.y.Set(R1.y)
			shamirsPointDouble(curve, &R1, &tmp)
		} else {
			tmp.x.Set(R1.x)
			tmp.y.Set(R1.y)
			shamirsPointAdd(curve, &R1, &R0, &tmp)
			tmp.x.Set(R0.x)
			tmp.y.Set(R0.y)
			shamirsPointDouble(curve, &R0, &tmp)
		}

		if dbits == 0 {
			break
		}
		dbits--
	}
	ret.x.Set(R0.x)
	ret.y.Set(R0.y)
}
