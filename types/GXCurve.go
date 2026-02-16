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
	"fmt"
	"math/big"

	"github.com/Gurux/gxdlms-go/enums"
)

// ECC x and y points in the curve.
type gxCurve struct {
	// ECC curve a value.
	a *big.Int
	// ECC curve p value.
	p *big.Int
	// ECC curve b parameter.
	b *big.Int
	// x and y-coordinate of base point G
	g gxEccPoint
	// Order of point G in ECC curve.
	n *big.Int
}

// newGxCurve creates a new ECC curve with the specified scheme.
//
// Parameters:
//
//	scheme: ECC scheme (P256 or P384).
//
// Returns:
//
//	Initialized gxCurve or error if scheme is invalid.
func newGxCurve(scheme enums.Ecc) (*gxCurve, error) {
	curve := &gxCurve{}
	switch scheme {
	case enums.EccP256:
		// Table A.1 – ECC_P256_Domain_Parameters
		curve.a.SetString("FFFFFFFF00000001000000000000000000000000FFFFFFFFFFFFFFFFFFFFFFFC", 16)
		gx := new(big.Int)
		gx.SetString("6B17D1F2E12C4247F8BCE6E563A440F277037D812DEB33A0F4A13945D898C296", 16)
		gy := new(big.Int)
		gy.SetString("4FE342E2FE1A7F9B8EE7EB4A7C0F9E162BCE33576B315ECECBB6406837BF51F5", 16)
		curve.g = gxEccPoint{x: gx, y: gy}
		curve.n.SetString("FFFFFFFF00000000FFFFFFFFFFFFFFFFBCE6FAADA7179E84F3B9CAC2FC632551", 16)
		curve.p.SetString("FFFFFFFF00000001000000000000000000000000FFFFFFFFFFFFFFFFFFFFFFFF", 16)
		curve.b.SetString("5AC635D8AA3A93E7B3EBBD55769886BC651D06B0CC53B0F63BCE3C3E27D2604B", 16)
	case enums.EccP384:
		// Table A.2 – ECC_P384_Domain_Parameters
		curve.a.SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFFFF0000000000000000FFFFFFFC", 16)
		gx := new(big.Int)
		gx.SetString("AA87CA22BE8B05378EB1C71EF320AD746E1D3B628BA79B9859F741E082542A385502F25DBF55296C3A545E3872760AB7", 16)
		gy := new(big.Int)
		gy.SetString("3617DE4A96262C6F5D9E98BF9292DC29F8F41DBD289A147CE9DA3113B5F0B8C00A60B1CE1D7E819D7A431D7C90EA0E5F", 16)
		curve.g = gxEccPoint{x: gx, y: gy}
		curve.n.SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFC7634D81F4372DDF581A0DB248B0A77AECEC196ACCC52973", 16)
		curve.p.SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFFFF0000000000000000FFFFFFFF", 16)
		curve.b.SetString("B3312FA7E23EE7E4988E056BE3F82D19181D9C6EFE8141120314088F5013875AC656398D8A2ED19D2A85C8EDD3EC2AEF", 16)
	default:
		return nil, fmt.Errorf("invalid ECC scheme: %v", scheme)
	}
	return curve, nil
}
