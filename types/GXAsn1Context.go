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

// GXAsn1Context ASN.1 context class.
type GXAsn1Context struct {
	//Context items.
	Items []any
	//Context index.
	Index int
	//Is constructed type.
	Constructed bool
}

// String implements the fmt.Stringer interface.
func (g *GXAsn1Context) String() string {
	if g.Constructed {
		return fmt.Sprintf("[%d] (Constructed) (%d) elem", g.Index, len(g.Items))
	}
	return fmt.Sprintf("[%d] (%d) elem", g.Index, len(g.Items))
}

// Constructor.
func NewGXAsn1Context() *GXAsn1Context {
	return &GXAsn1Context{
		Items:       make([]any, 0),
		Constructed: true,
	}
}
