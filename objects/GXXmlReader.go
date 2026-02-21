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
	"io"
	"strconv"
	"strings"

	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/types"
	"golang.org/x/text/language"
)

type GXXmlReader struct {
	dec     *xml.Decoder
	tok     xml.Token // current token
	Objects *GXDLMSObjectCollection
}

func NewGXXmlReaderFromStream(stream *bufio.Reader) *GXXmlReader {
	//var tok xml.Token
	r := bufio.NewReader(stream)
	dec := xml.NewDecoder(r)
	tok, _ := dec.Token()
	if _, ok := tok.(xml.ProcInst); ok {
		tok, _ = dec.Token()
	}
	return &GXXmlReader{
		dec: dec, tok: tok}
}

func (r *GXXmlReader) Close() error { return nil }
func (r *GXXmlReader) EOF() bool {
	return r.tok == nil || r.tok == io.EOF
}

func (r *GXXmlReader) IsStartElement() bool {
	r.getNext()
	r.nodeType()
	_, ok := r.tok.(xml.StartElement)
	return ok
}

func (r *GXXmlReader) Name() string {
	switch r.tok.(type) {
	case xml.StartElement:
		return r.tok.(xml.StartElement).Name.Local
	case xml.EndElement:
		return r.tok.(xml.EndElement).Name.Local
	}
	return ""
}

func (r *GXXmlReader) Read() error {
	v, err := r.dec.Token()
	r.tok = v
	if err == io.EOF {
		r.tok = nil
		err = nil
	}
	return err
}

func (r *GXXmlReader) GetAttribute(index int) string { return "" }

func (x *GXXmlReader) nodeType() string {
	switch x.tok.(type) {
	case xml.StartElement:
		return "StartElement"
	case xml.EndElement:
		return "EndElement"
	case xml.CharData:
		return "CharData"
	case xml.Comment:
		return "Comment"
	case xml.Directive:
		return "Directive"
	case xml.ProcInst:
		return "ProcInst"
	default:
		return "None"
	}
}

// getNext skips comments and whitespace-only character data.
func (x *GXXmlReader) getNext() error {
	for {
		switch t := x.tok.(type) {
		case xml.Comment:
			err := x.Read()
			if err != nil {
				return err
			}
		case xml.CharData:
			if strings.TrimSpace(string(t)) == "" {
				err := x.Read()
				if err != nil {
					return err
				}
			} else {
				return nil
			}
		default:
			return nil
		}
	}
}
func (x *GXXmlReader) isStartElementNamed(name string) bool {
	se, ok := x.tok.(xml.StartElement)
	if !ok {
		return false
	}
	return strings.EqualFold(se.Name.Local, name)
}

// ReadEndElement(name)
func (x *GXXmlReader) ReadEndElement(name string) error {
	if err := x.getNext(); err != nil {
		return err
	}
	if ee, ok := x.tok.(xml.EndElement); ok && strings.EqualFold(ee.Name.Local, name) {
		if err := x.Read(); err != nil {
			return err
		}
		return x.getNext()
	}
	return nil
}

func (x *GXXmlReader) IsStartElementNamed(name string, getNext bool) (bool, error) {
	if err := x.getNext(); err != nil {
		return false, err
	}
	ret := x.isStartElementNamed(name)

	if getNext && (ret || x.isEndElementNamed(name)) {
		// C# checks IsEmptyElement; in Go StartElement has no direct IsEmpty.
		// We approximate: if it's StartElement and next token is EndElement immediately, treat as empty.
		var wasStart bool
		if ret {
			wasStart = true
		}

		if err := x.Read(); err != nil {
			return false, err
		}

		if !ret {
			return x.IsStartElementNamed(name, getNext)
		}

		if wasStart && x.isEndElementNamed(name) {
			// empty element <name/>
			// advance past end element
			if err := x.Read(); err != nil {
				return false, err
			}
			if err := x.getNext(); err != nil {
				return false, err
			}
			return false, nil
		}
	}

	if err := x.getNext(); err != nil {
		return false, err
	}
	return ret, nil
}

func (x *GXXmlReader) isEndElementNamed(name string) bool {
	ee, ok := x.tok.(xml.EndElement)
	if !ok {
		return false
	}
	return strings.EqualFold(ee.Name.Local, name)
}

func (x *GXXmlReader) ReadElementContentAsString(name string, def string) (string, error) {
	if err := x.getNext(); err != nil {
		return def, err
	}
	if !strings.EqualFold(x.Name(), name) {
		return def, nil
	}
	s, err := x.readElementText(name)
	if err != nil {
		return def, err
	}
	return s, nil
}

func (x *GXXmlReader) ReadElementContentAsGXDateTime(name string) (types.GXDateTime, error) {
	ret, err := x.ReadElementContentAsString(name, "")
	if err != nil || ret == "" {
		return types.GXDateTime{}, err
	}
	ret2, err := types.NewGXDateTimeFromString(ret, &language.AmericanEnglish)
	return *ret2, err
}

func (x *GXXmlReader) ReadElementContentAsGXDate(name string) (types.GXDate, error) {
	ret, err := x.ReadElementContentAsString(name, "")
	if err != nil || ret == "" {
		return types.GXDate{}, err
	}
	sdt, err := types.NewGXDateFromString(ret, &language.AmericanEnglish)
	return *sdt, err
}

func (x *GXXmlReader) ReadElementContentAsTime(name string) (types.GXTime, error) {
	ret, err := x.ReadElementContentAsString(name, "")
	if err != nil || ret == "" {
		return types.GXTime{}, err
	}
	sdt, err := types.NewGXTimeFromString(ret, &language.AmericanEnglish)
	return *sdt, err
}

func (x *GXXmlReader) ReadElementContentAsInt(name string, def int) (int, error) {
	if err := x.getNext(); err != nil {
		return def, err
	}
	if !strings.EqualFold(x.Name(), name) {
		return def, nil
	}
	s, err := x.readElementText(name)
	if err != nil {
		return def, err
	}
	v, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		return def, err
	}
	return v, nil
}

func (x *GXXmlReader) ReadElementContentAsBool(name string, def bool) (bool, error) {
	ret, err := x.ReadElementContentAsString(name, "")
	if err != nil {
		return false, err
	}
	return strconv.ParseBool(ret)
}

func (x *GXXmlReader) ReadElementContentAsInt8(name string, def int) (int8, error) {
	ret, err := x.ReadElementContentAsInt(name, def)
	if err != nil {
		return 0, err
	}
	return int8(ret), nil
}

func (x *GXXmlReader) ReadElementContentAsInt16(name string, def int) (int16, error) {
	ret, err := x.ReadElementContentAsInt(name, def)
	if err != nil {
		return 0, err
	}
	return int16(ret), nil
}

func (x *GXXmlReader) ReadElementContentAsUInt8(name string, def int) (uint8, error) {
	ret, err := x.ReadElementContentAsInt(name, def)
	if err != nil {
		return 0, err
	}
	return uint8(ret), nil
}

func (x *GXXmlReader) ReadElementContentAsUInt16(name string, def int) (uint16, error) {
	ret, err := x.ReadElementContentAsInt(name, def)
	if err != nil {
		return 0, err
	}
	return uint16(ret), nil
}

func (x *GXXmlReader) ReadElementContentAsUInt32(name string, def int) (uint32, error) {
	ret, err := x.ReadElementContentAsInt(name, def)
	if err != nil {
		return 0, err
	}
	return uint32(ret), nil
}

func (x *GXXmlReader) ReadElementContentAsLong(name string, def int64) (int64, error) {
	if err := x.getNext(); err != nil {
		return def, err
	}
	if !strings.EqualFold(x.Name(), name) {
		return def, nil
	}
	s, err := x.readElementText(name)
	if err != nil {
		return def, err
	}
	v, err := strconv.ParseInt(strings.TrimSpace(s), 10, 64)
	if err != nil {
		return def, err
	}
	return v, nil
}

func (x *GXXmlReader) ReadElementContentAsULong(name string, def uint64) (uint64, error) {
	if err := x.getNext(); err != nil {
		return def, err
	}
	if !strings.EqualFold(x.Name(), name) {
		return def, nil
	}
	s, err := x.readElementText(name)
	if err != nil {
		return def, err
	}
	v, err := strconv.ParseUint(strings.TrimSpace(s), 10, 64)
	if err != nil {
		return def, err
	}
	return v, nil
}

func (x *GXXmlReader) ReadElementContentAsDouble(name string, def float64) (float64, error) {
	if err := x.getNext(); err != nil {
		return def, err
	}
	if !strings.EqualFold(x.Name(), name) {
		return def, nil
	}
	s, err := x.readElementText(name)
	if err != nil {
		return def, err
	}
	v, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
	if err != nil {
		return def, err
	}
	return v, nil
}

func (x *GXXmlReader) ReadElementContentAsDateTime(name string, def *types.GXDateTime) (types.GXDateTime, error) {
	ret, err := x.ReadElementContentAsString(name, "")
	if err != nil {
		if def == nil {
			return types.GXDateTime{}, err
		}
		return *def, err
	}

	ret2, err := types.NewGXDateTimeFromString(ret, &language.AmericanEnglish)
	if err != nil {
		return types.GXDateTime{}, err
	}
	return *ret2, err
}

func (x *GXXmlReader) ReadElementContentAsObject(name string, def any, obj IGXDLMSBase, index int) (any, error) {
	if !strings.EqualFold(x.Name(), name) {
		return def, nil
	}
	se, _ := x.tok.(xml.StartElement)
	attrCount := len(se.Attr)

	if attrCount == 0 {
		// consume element text
		_, err := x.readElementText(name)
		if err != nil {
			return nil, err
		}
		if obj != nil {
			obj.Base().SetDataType(index, enums.DataTypeNone)
		}
		return nil, nil
	}

	dtStr := se.Attr[0].Value
	v, err := strconv.Atoi(strings.TrimSpace(dtStr))
	if err != nil {
		return nil, err
	}
	dt := enums.DataType(v)
	if obj != nil {
		obj.Base().SetDataType(index, dt)
	}

	uiType := dt
	if attrCount > 1 {
		if t, err := enums.DataTypeParse(se.Attr[1].Value); err == nil {
			uiType = t
		}
	}
	if obj != nil && obj.GetUIDataType(index) == enums.DataTypeNone {
		obj.Base().SetUIDataType(index, uiType)
	}

	ret, err := x.readElementText(name)
	if err != nil {
		return def, err
	}
	if dt == enums.DataTypeArray || dt == enums.DataTypeStructure {
		// Move inside element and read items.
		if err := x.Read(); err != nil {
			return def, err
		}
		if err := x.getNext(); err != nil {
			return def, err
		}
		arr, err := x.readArray()
		if err != nil {
			return def, err
		}
		if err := x.ReadEndElement(name); err != nil {
			return def, err
		}
		return arr, nil
	}
	if dt != enums.DataTypeNone {
		return internal.Convert(ret, dt)
	}
	return ret, err
}

func (x *GXXmlReader) readArray() ([]any, error) {
	var list []any
	for {
		ok, err := x.IsStartElementNamed("Item", false)
		if err != nil {
			return nil, err
		}
		if !ok {
			break
		}
		v, err := x.ReadElementContentAsObject("Item", nil, nil, 0)
		if err != nil {
			return nil, err
		}
		list = append(list, v)
	}
	return list, nil
}

// readElementText reads the text content for a start element currently at x.tok,
// consuming until it has read the matching end element.
func (x *GXXmlReader) readElementText(name string) (string, error) {
	// Current token should be StartElement(name)
	// We'll use Decoder.Token() loop until end element.
	var b strings.Builder

	// Read next token after start
	for {
		t, err := x.dec.Token()
		if err == io.EOF {
			x.tok = nil
			return b.String(), nil
		}
		if err != nil {
			return "", err
		}
		x.tok = t

		switch tt := t.(type) {
		case xml.CharData:
			b.WriteString(string(tt))
		case xml.EndElement:
			if strings.EqualFold(tt.Name.Local, name) {
				// advance once more so x.tok points to next thing (like C# does often)
				err := x.Read()
				if err != nil {
					return "", err
				}
				x.getNext()
				return b.String(), nil
			}
		}
	}
}

func (x *GXXmlReader) String() string {
	if x.dec != nil {
		return fmt.Sprintf("%s, Name=%q", x.nodeType(), x.Name())
	}
	return "GXXmlReader(nil)"
}
