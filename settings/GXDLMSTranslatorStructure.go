package settings

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
	"errors"
	"fmt"
	"strings"

	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal/buffer"
)

const DataTypeOffset = int(0xFF0000)

// This class is used internally in GXDLMSTranslator to save generated xml.
type GXDLMSTranslatorStructure struct {
	// Is comment added already. Nested comments are not allowed in a XML.
	comments int

	outputType enums.TranslatorOutputType

	// Name space is omit.
	omitNameSpace bool

	// Amount of spaces.
	offset int

	tags map[int]string

	showNumericsAsHex bool

	sb strings.Builder

	ShowStringAsHex bool

	// Are comments added.
	Comments bool

	// Are spaces ignored.
	IgnoreSpaces bool
}

func (g *GXDLMSTranslatorStructure) OutputType() enums.TranslatorOutputType {
	return g.outputType
}

// Name space is omit.
func (g *GXDLMSTranslatorStructure) OmitNameSpace() bool {
	return g.omitNameSpace
}

// Offset returns the amount of spaces.
func (g *GXDLMSTranslatorStructure) Offset() int {
	return g.offset
}

// SetOffset sets the amount of spaces.
func (g *GXDLMSTranslatorStructure) SetOffset(value int) error {
	if value < 0 {
		return errors.New("offset")
	}
	g.offset = value
	return nil
}

// appendSpaces returns the append spaces.
func (g *GXDLMSTranslatorStructure) appendSpaces() {
	if g.IgnoreSpaces {
		g.sb.WriteString(" ")
	} else if g.offset > 0 {
		g.sb.WriteString(strings.Repeat(" ", 2*g.offset))
	}
}

func (g *GXDLMSTranslatorStructure) GetTag(tag int) string {
	if g.outputType == enums.TranslatorOutputTypeSimpleXML || g.omitNameSpace {
		return g.tags[tag]
	}
	return "x:" + g.tags[tag]
}

func (g *GXDLMSTranslatorStructure) GetDataType(type_ enums.DataType) string {
	return g.GetTag(DataTypeOffset + int(type_))
}

func (g *GXDLMSTranslatorStructure) String() string {
	return g.sb.String()
}

func (g *GXDLMSTranslatorStructure) AppendStringLine(value string) {
	if g.IgnoreSpaces {
		g.sb.WriteString(value)
	} else {
		g.appendSpaces()
		g.sb.WriteString(value)
	}
}

func (g *GXDLMSTranslatorStructure) AppendLineFromTag(tag int, name string, value any) {
	g.AppendLine(g.GetTag(tag), name, value)
}
func (g *GXDLMSTranslatorStructure) AppendLine(tag string, name string, value any) {
	g.appendSpaces()
	g.sb.WriteString("<")
	g.sb.WriteString(tag)
	if g.outputType == enums.TranslatorOutputTypeSimpleXML {
		g.sb.WriteString(" ")
		if name == "" {
			g.sb.WriteString("Value")
		} else {
			g.sb.WriteString(name)
		}
		g.sb.WriteString("=\"")
	} else {
		g.sb.WriteString(">")
	}
	if v1, ok := value.(uint8); ok {
		g.sb.WriteString(g.IntegerToHex(int64(v1), 2, false))
	} else if v2, ok := value.(int8); ok {
		g.sb.WriteString(g.IntegerToHex(int64(v2), 2, false))
	} else if v3, ok := value.(uint16); ok {
		g.sb.WriteString(g.IntegerToHex(int64(v3), 4, false))
	} else if v4, ok := value.(int16); ok {
		g.sb.WriteString(g.IntegerToHex(int64(v4), 4, false))
	} else if v5, ok := value.(uint32); ok {
		g.sb.WriteString(g.IntegerToHex(int64(v5), 8, false))
	} else if v6, ok := value.(int32); ok {
		g.sb.WriteString(g.IntegerToHex(int64(v6), 8, false))
	} else if v7, ok := value.(uint64); ok {
		g.sb.WriteString(g.IntegerToHex(int64(v7), 16, false))
	} else if v8, ok := value.(int64); ok {
		g.sb.WriteString(g.IntegerToHex(v8, 16, false))
	} else if bv, ok := value.([]byte); ok {
		g.sb.WriteString(buffer.ToHex(bv, true))
	} else if _, ok := value.([]int8); ok {
		g.sb.WriteString(buffer.ToHex(value.([]byte), true))
	} else {
		g.sb.WriteString(fmt.Sprint(value))
	}
	if g.outputType == enums.TranslatorOutputTypeSimpleXML {
		g.sb.WriteString("\" />")
	} else {
		g.sb.WriteString("</")
		g.sb.WriteString(tag)
		g.sb.WriteString(">")
	}
	g.sb.WriteString("\r")
	g.sb.WriteString("\n")
}

// StartComment returns the start comment section.
//
// Parameters:
//
//	comment: Comment to add.
func (g *GXDLMSTranslatorStructure) StartComment(comment string) {
	if g.Comments {
		g.appendSpaces()
		if g.comments == 0 {
			g.sb.WriteString("<!-- ")
		} else {
			g.sb.WriteString("# ")
		}
		g.comments++
		g.sb.WriteString(comment)
		g.sb.WriteString("\r")
		g.sb.WriteString("\n")
		g.offset++
	}
}

// EndComment returns the end comment section.
func (g *GXDLMSTranslatorStructure) EndComment() {
	if g.Comments {
		g.comments--
		g.offset--
		if g.comments == 0 {
			g.appendSpaces()
			g.sb.WriteString("-->")
		}
		g.sb.WriteString("\r")
		g.sb.WriteString("\n")
	}
}

// AppendComment returns the append comment.
//
// Parameters:
//
//	comment: Comment to add.
func (g *GXDLMSTranslatorStructure) AppendComment(comment string) {
	if g.Comments {
		g.appendSpaces()
		if g.comments == 0 {
			g.sb.WriteString("<!-- ")
			g.sb.WriteString(comment)
			g.sb.WriteString(" -->")
		} else {
			g.sb.WriteString("# ")
			g.sb.WriteString(comment)
		}
		g.sb.WriteString("\r")
		g.sb.WriteString("\n")
	}
}

func (g *GXDLMSTranslatorStructure) AppendString(value string) {
	g.sb.WriteString(value)
}

func (g *GXDLMSTranslatorStructure) Append(tag int, start bool) {
	if start {
		g.sb.WriteString("<")
	} else {
		g.sb.WriteString("</")
	}
	g.sb.WriteString(g.GetTag(tag))
	g.sb.WriteString(">")
}

func (g *GXDLMSTranslatorStructure) AppendStartTag(tag int, name string, value string, plain bool) {
	g.appendSpaces()
	g.sb.WriteString("<")
	g.sb.WriteString(g.GetTag(tag))
	if g.outputType == enums.TranslatorOutputTypeSimpleXML && name != "" {
		g.sb.WriteString(" ")
		g.sb.WriteString(name)
		g.sb.WriteString("=\"")
		g.sb.WriteString(value)
		g.sb.WriteString("\" >")
	} else {
		g.sb.WriteString(">")
	}
	if !plain {
		g.sb.WriteString("")
	}
	g.offset++
}

func (g *GXDLMSTranslatorStructure) AppendEndTag(tag int, plain bool) {
	g.offset--
	if !plain {
		g.appendSpaces()
	}
	g.sb.WriteString("</")
	g.sb.WriteString(g.GetTag(tag))
	g.sb.WriteString(">")
	g.sb.WriteString("")
}

func (g *GXDLMSTranslatorStructure) AppendEmptyTag(tag int) {
	g.AppendEmptyStringTag(g.tags[tag])
}

func (g *GXDLMSTranslatorStructure) AppendEmptyStringTag(tag string) {
	g.appendSpaces()
	g.sb.WriteString("<")
	g.sb.WriteString(tag)
	g.sb.WriteString("/>")
}

// Trim returns the remove \r\n.
func (g *GXDLMSTranslatorStructure) Trim() {
	tmp := g.sb.String()
	tmp = strings.TrimSuffix(tmp, "\r\n")
	g.sb.Reset()
	g.sb.WriteString(tmp)
}

// GetXmlLength returns the get XML Length.
//
// Returns:
//
//	XML Length.
func (g *GXDLMSTranslatorStructure) GetXmlLength() int {
	return g.sb.Len()
}

// SetXmlLength returns the set XML Length.
//
// Parameters:
//
//	value: XML Length.
func (g *GXDLMSTranslatorStructure) SetXmlLength(value int) {
	tmp := g.sb.String()[:value]
	g.sb.Reset()
	g.sb.WriteString(tmp)
}

// IntegerToHex returns the convert integer to string.
//
// Parameters:
//
//	value: Conveted value.
//	desimals: Desimal count.
//	forceHex: Force value as hex.
//
// Returns:
//
//	Integer value as a string.
func (g *GXDLMSTranslatorStructure) IntegerToHex(value any, desimals int, forceHex bool) string {
	if forceHex || (g.showNumericsAsHex && g.outputType == enums.TranslatorOutputTypeSimpleXML) {
		ret := fmt.Sprintf("%0*X", desimals, value)
		return ret
	}
	return fmt.Sprint(value)
}
