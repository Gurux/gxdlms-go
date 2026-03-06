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

import "fmt"

// GXAsn1Context represents a BER/DER context-specific tagged value.
//
// It is used when decoding ASN.1 data to store the items contained within a
// context tag. The Index is the context tag number (0-15), and Constructed
// indicates whether the tag is constructed (contains nested elements) or primitive.
type GXAsn1Context struct {
	// Items contains decoded child elements within the context.
	Items []any
	// Index is the numeric context tag identifier.
	Index int
	// Constructed indicates whether the context tag is constructed (true) or primitive.
	Constructed bool
}

// String implements fmt.Stringer and provides a brief debug representation.
func (g *GXAsn1Context) String() string {
	if g.Constructed {
		return fmt.Sprintf("[%d] (Constructed) (%d) elem", g.Index, len(g.Items))
	}
	return fmt.Sprintf("[%d] (%d) elem", g.Index, len(g.Items))
}

// NewGXAsn1Context creates a new constructed ASN.1 context tag container.
func NewGXAsn1Context() *GXAsn1Context {
	return &GXAsn1Context{
		Items:       make([]any, 0),
		Constructed: true,
	}
}
