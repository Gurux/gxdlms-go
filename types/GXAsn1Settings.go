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
	"strings"
)

type gxAsn1Settings struct {
	count int

	sb strings.Builder

	tags map[int16]string

	tagbyName map[string]int16

	// Are comments used.
	Comments bool
}

func (g *gxAsn1Settings) XmlLength() int {
	return g.sb.Len()
}

func (g *gxAsn1Settings) AddTag(key int16, value string) {
	g.tags[key] = value
	g.tagbyName[strings.ToLower(value)] = key
}

func (g *gxAsn1Settings) GetTag(value int16) string {
	return g.tags[value]
}

func (g *gxAsn1Settings) GetTagByName(value string) int16 {
	return g.tagbyName[strings.ToLower(value)]
}

// AppendComment returns the add comment.
//
// Parameters:
//
//	offset: Offset.
//	value: Comment value.
func (g *gxAsn1Settings) AppendComment(offset int, value string) {
	if g.Comments {
		empty := g.sb.Len() == 0
		var tmp strings.Builder
		if empty {
			tmp = g.sb
		} else {
			tmp = strings.Builder{}
		}
		for pos := 0; pos < g.count-1; pos++ {
			tmp.WriteString(" ")
		}
		tmp.WriteString("<!--")
		tmp.WriteString(value)
		tmp.WriteString("-->\r\n")
		if !empty {
			s := g.sb.String()
			s = s[:offset] + tmp.String() + s[offset:]
			g.sb.Reset()
			g.sb.WriteString(s)
		}
	}
}

// AppendSpaces returns the append spaces to the buffer.
func (g *gxAsn1Settings) AppendSpaces() {
	for pos := 0; pos != g.count; pos++ {
		g.sb.WriteString(" ")
	}
}

func (g *gxAsn1Settings) Append(value string) {
	g.sb.WriteString(value)
}

// Increments the indentation level and appends a new line to the current content.
func (g *gxAsn1Settings) Increase() {
	g.count++
	g.Append("\r\n")
}

// Decrease decreases indentation level.
func (g *gxAsn1Settings) Decrease() {
	g.count--
	g.AppendSpaces()
}

// String returns settings as string.
func (g *gxAsn1Settings) String() string {
	return g.sb.String()
}
