// --------------------------------------------------------------------------
//
//	Gurux Ltd
//
// Filename:        $HeadURL$
//
// Version:         $Revision$,
//
//	$Date$
//	$Author$
//
// # Copyright (c) Gurux Ltd
//
// ---------------------------------------------------------------------------
//
//	DESCRIPTION
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
// ---------------------------------------------------------------------------
package objects

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"os"
	"strconv"

	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/types"
)

type GXXmlWriter struct {
	enc      *xml.Encoder
	out      *bufio.Writer
	file     *os.File
	settings *GXXmlWriterSettings

	// Simple element stack so we can implement WriteString easily.
	open []string
}

func NewGXXmlWriterStream(stream *bufio.Writer, settings *GXXmlWriterSettings) *GXXmlWriter {
	w := bufio.NewWriter(stream)
	return &GXXmlWriter{
		out:      w,
		enc:      xml.NewEncoder(w),
		settings: settings,
	}
}
func (w *GXXmlWriter) Close() error { return nil }

func (x *GXXmlWriter) ignoreDefaultValues() bool {
	return x.settings != nil && x.settings.IgnoreDefaultValues
}

func (x *GXXmlWriter) useMeterTime() bool {
	return x.settings != nil && x.settings.UseMeterTime
}

// WriteStartDocument ~ writer.WriteStartDocument()
func (x *GXXmlWriter) WriteStartDocument() error {
	// xml.Encoder doesn't have explicit StartDocument; write XML header manually.
	_, err := x.out.WriteString(xml.Header)
	return err
}

func (x *GXXmlWriter) WriteStartElement(name string) error {
	x.open = append(x.open, name)
	return x.enc.EncodeToken(xml.StartElement{Name: xml.Name{Local: name}})
}

func (x *GXXmlWriter) WriteAttributeString(name, value string) error {
	if len(x.open) == 0 {
		return fmt.Errorf("WriteAttributeString called with no open element")
	}
	// Go encoder can't "add attribute to already emitted StartElement" unless we emit the start with attrs.
	// So we implement a stricter pattern: use WriteStartElementWithAttrs.
	return fmt.Errorf("use WriteStartElementWithAttrs in Go version (attributes must be known at start)")
}

// Idiomatic helper: start element with attributes in one go.
func (x *GXXmlWriter) WriteStartElementWithAttrs(name string, attrs map[string]string) error {
	x.open = append(x.open, name)
	var a []xml.Attr
	for k, v := range attrs {
		a = append(a, xml.Attr{Name: xml.Name{Local: k}, Value: v})
	}
	return x.enc.EncodeToken(xml.StartElement{Name: xml.Name{Local: name}, Attr: a})
}

func (x *GXXmlWriter) WriteEndElement() error {
	if len(x.open) == 0 {
		return fmt.Errorf("WriteEndElement: no open elements")
	}
	name := x.open[len(x.open)-1]
	x.open = x.open[:len(x.open)-1]
	return x.enc.EncodeToken(xml.EndElement{Name: xml.Name{Local: name}})
}

func (x *GXXmlWriter) WriteEndDocument() error {
	// Just flush.
	return x.enc.Flush()
}

func (x *GXXmlWriter) Flush() error {
	if err := x.enc.Flush(); err != nil {
		return err
	}
	return x.out.Flush()
}

// WriteString writes character data inside current element.
func (x *GXXmlWriter) WriteString(s string) error {
	return x.enc.EncodeToken(xml.CharData([]byte(s)))
}

// --- Element string writers (matching C# behavior) ---------------------

func (x *GXXmlWriter) WriteElementStringU64(name string, value uint64) error {
	if !x.ignoreDefaultValues() || value != 0 {
		return x.writeSimpleElement(name, strconv.FormatUint(value, 10))
	}
	return nil
}

func (x *GXXmlWriter) WriteElementStringU32(name string, value uint32) error {
	return x.WriteElementStringU64(name, uint64(value))
}

func (x *GXXmlWriter) WriteElementStringDouble(name string, value float64, defaultValue float64) error {
	if !x.ignoreDefaultValues() || value != defaultValue {
		// invariant culture = '.' decimal, Go does that by default.
		return x.writeSimpleElement(name, strconv.FormatFloat(value, 'g', -1, 64))
	}
	return nil
}

func (x *GXXmlWriter) WriteElementStringInt(name string, value int) error {
	if !x.ignoreDefaultValues() || value != 0 {
		return x.writeSimpleElement(name, strconv.Itoa(value))
	}
	return nil
}

func (x *GXXmlWriter) WriteElementString(name string, value any) error {
	if value != nil {
		return x.writeSimpleElement(name, fmt.Sprint(value))
	}
	if !x.ignoreDefaultValues() {
		return x.writeSimpleElement(name, "")
	}
	return nil
}

func (x *GXXmlWriter) WriteElementStringBool(name string, value bool) error {
	if !x.ignoreDefaultValues() || value {
		if value {
			return x.writeSimpleElement(name, "1")
		}
		return x.writeSimpleElement(name, "0")
	}
	return nil
}

func (x *GXXmlWriter) WriteElementStringGXDateTime(name string, value *types.GXDateTime) error {
	if value != nil {
		var s string
		if x.useMeterTime() {
			s = value.ToFormatMeterString(nil)
		} else {
			s = value.ToFormatString(nil, false)
		}
		return x.writeSimpleElement(name, s)
	}
	if !x.ignoreDefaultValues() {
		return x.writeSimpleElement(name, "")
	}
	return nil
}

func (x *GXXmlWriter) writeSimpleElement(name, text string) error {
	start := xml.StartElement{Name: xml.Name{Local: name}}
	if err := x.enc.EncodeToken(start); err != nil {
		return err
	}
	if err := x.enc.EncodeToken(xml.CharData([]byte(text))); err != nil {
		return err
	}
	return x.enc.EncodeToken(xml.EndElement{Name: start.Name})
}

func (x *GXXmlWriter) WriteElementObject(name string, value any, dt enums.DataType, uiType enums.DataType) error {
	if value == nil {
		if x.ignoreDefaultValues() {
			return nil
		}
		// <name Type="0" />
		return x.writeEmptyElementWithAttrs(name, map[string]string{"Type": "0"})
	}

	attrs := map[string]string{
		"Type": strconv.Itoa(int(dt)),
	}

	// UIType writing rules (match C# roughly)
	if uiType != enums.DataTypeNone && dt != uiType && (uiType != enums.DataTypeString || dt == enums.DataTypeOctetString) {
		attrs["UIType"] = strconv.Itoa(int(uiType))
	} else {
		switch value.(type) {
		case float64:
			attrs["UIType"] = strconv.Itoa(int(enums.DataTypeFloat64))
		case float32:
			attrs["UIType"] = strconv.Itoa(int(enums.DataTypeFloat32))
		}
	}

	// Start element with attrs
	start := xml.StartElement{Name: xml.Name{Local: name}}
	for k, v := range attrs {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: k}, Value: v})
	}
	if err := x.enc.EncodeToken(start); err != nil {
		return err
	}

	// Content
	if dt == enums.DataTypeArray || dt == enums.DataTypeStructure {
		if err := x.writeArray(value); err != nil {
			return err
		}
	} else {
		var s string
		switch v := value.(type) {
		case *types.GXDateTime:
			if x.useMeterTime() {
				s = v.ToFormatMeterString(nil)
			} else {
				s = v.ToFormatString(nil, false)
			}
		case []byte:
			s = types.ToHex(v, false)
		default:
			s = fmt.Sprint(v)
		}
		if err := x.enc.EncodeToken(xml.CharData([]byte(s))); err != nil {
			return err
		}
	}

	// End
	return x.enc.EncodeToken(xml.EndElement{Name: start.Name})
}

func (x *GXXmlWriter) writeEmptyElementWithAttrs(name string, attrs map[string]string) error {
	start := xml.StartElement{Name: xml.Name{Local: name}}
	for k, v := range attrs {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: k}, Value: v})
	}
	// <name ...></name> (Encoder doesn't emit <name/> reliably with tokens; this is fine)
	if err := x.enc.EncodeToken(start); err != nil {
		return err
	}
	return x.enc.EncodeToken(xml.EndElement{Name: start.Name})
}

func (x *GXXmlWriter) writeArray(data any) error {
	// C# checks List<object>, GXArray, GXStructure
	var list []any

	switch v := data.(type) {
	case []any:
		list = v
	case types.GXArray:
		list = []any(v)
	case types.GXStructure:
		list = []any(v)
	default:
		return nil
	}

	for _, tmp := range list {
		switch t := tmp.(type) {
		case []byte:
			//TODO:
			if err := x.WriteElementObject("Item", t, enums.DataTypeNone, enums.DataTypeNone); err != nil {
				return err
			}
		case types.GXArray:
			// <Item Type="Array"> ... </Item>
			if err := x.startItemWithType(enums.DataTypeArray); err != nil {
				return err
			}
			if err := x.writeArray(t); err != nil {
				return err
			}
			if err := x.enc.EncodeToken(xml.EndElement{Name: xml.Name{Local: "Item"}}); err != nil {
				return err
			}
		case types.GXStructure:
			if err := x.startItemWithType(enums.DataTypeStructure); err != nil {
				return err
			}
			if err := x.writeArray(t); err != nil {
				return err
			}
			if err := x.enc.EncodeToken(xml.EndElement{Name: xml.Name{Local: "Item"}}); err != nil {
				return err
			}
		default:
			if err := x.WriteElementObject("Item", t, enums.DataTypeNone, enums.DataTypeNone); err != nil {
				return err
			}
		}
	}
	return nil
}

func (x *GXXmlWriter) startItemWithType(dt enums.DataType) error {
	start := xml.StartElement{
		Name: xml.Name{Local: "Item"},
		Attr: []xml.Attr{{Name: xml.Name{Local: "Type"}, Value: strconv.Itoa(int(dt))}},
	}
	return x.enc.EncodeToken(start)
}
