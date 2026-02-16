package objects

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
)

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

type GXDLMSObjectCollection []IGXDLMSBase

// Clear removes all elements from the collection.
func (g *GXDLMSObjectCollection) Clear() {
	*g = (*g)[:0]
}

// Clear removes all elements from the collection.
func (g *GXDLMSObjectCollection) Add(item IGXDLMSBase) {
	*g = append(*g, item)
}

func (g *GXDLMSObjectCollection) GetObjects(objectType enums.ObjectType) GXDLMSObjectCollection {
	items := GXDLMSObjectCollection{}
	for _, it := range *g {
		if it.Base().objectType == objectType {
			items = append(items, it)
		}
	}
	return items
}

func (g *GXDLMSObjectCollection) GetObjects2(objectTypes []enums.ObjectType) GXDLMSObjectCollection {
	items := GXDLMSObjectCollection{}
	for _, it := range *g {
		if internal.Contains(objectTypes, it.Base().objectType) {
			items = append(items, it)
		}
	}
	return items
}

// FindByLN searches the collection for an object with the given logical name (string).
// If objectType is enums.ObjectTypeNone the type check is skipped.
func (g *GXDLMSObjectCollection) FindByLN(objectType enums.ObjectType, ln string) IGXDLMSBase {
	for i := range *g {
		it := (*g)[i]
		if (objectType == enums.ObjectTypeNone || it.Base().objectType == objectType) && strings.TrimSpace(it.Base().LogicalName()) == ln {
			return it
		}
	}
	return nil
}

// FindByLNBytes searches the collection for an object with the logical name represented by the byte slice.
// The byte slice is converted to the dotted string representation (e.g. 1.0.1.8.0.255).
func (g *GXDLMSObjectCollection) FindByLNBytes(objectType enums.ObjectType, ln []byte) (IGXDLMSBase, error) {
	name, err := helpers.ToLogicalName(ln)
	if err == nil {
		for i := range *g {
			it := (*g)[i]
			if (objectType == enums.ObjectTypeNone || it.Base().objectType == objectType) && strings.TrimSpace(it.Base().LogicalName()) == name {
				return it, nil
			}
		}
	}
	return nil, err
}

// FindBySN searches the collection for an object with the given short name.
func (g *GXDLMSObjectCollection) FindBySN(sn uint16) IGXDLMSBase {
	for i := range *g {
		it := (*g)[i]
		// Use conversion to uint16 in case ShortName is an int-like type.
		if uint16(it.Base().ShortName) == sn {
			return it
		}
	}
	return nil
}

// String returns the string representation of the collection.
func (g *GXDLMSObjectCollection) String() string {
	var sb strings.Builder
	sb.WriteByte('[')
	for i, it := range *g {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprint(it.Base().Name()))
	}
	sb.WriteByte(']')
	return sb.String()
}

// LoadFromFile returns collection of serialized COSEM objects.
func (g *GXDLMSObjectCollection) LoadFromFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	reader := bufio.NewReader(f)
	defer f.Close()
	return g.LoadFromStream(reader)
}

// LoadFromStream returns new collection.
func (g *GXDLMSObjectCollection) LoadFromStream(stream *bufio.Reader) error {
	var obj IGXDLMSBase
	reader := NewGXXmlReaderFromStream(stream)
	defer reader.Close()
	reader.Objects = *g
	for !reader.EOF() {
		if reader.IsStartElement() {
			target := reader.Name()

			if strings.EqualFold("Objects", target) {
				// Skip.
				if err := reader.Read(); err != nil {
					return err
				}
			} else if strings.HasPrefix(target, "GXDLMS") {
				str := target[6:]
				if err := reader.Read(); err != nil {
					return err
				}
				t, err := enums.ObjectTypeParse(str)
				if err != nil {
					return err
				}
				obj = CreateObject(t)
				obj.Base().Version = 0
			} else if strings.EqualFold("SN", target) {
				v, err := reader.ReadElementContentAsInt("SN", 0)
				if err != nil {
					return err
				}
				if obj != nil {
					obj.Base().ShortName = int16(v)
				}
			} else if strings.EqualFold("LN", target) {
				s, err := reader.ReadElementContentAsString("LN", "")
				if err != nil {
					return err
				}
				if obj != nil && s != "" {
					err := obj.Base().SetLogicalName(s)
					if err != nil {
						return err
					}
					tmp := reader.Objects.FindByLN(obj.Base().ObjectType(), obj.Base().LogicalName())
					if tmp == nil {
						*g = append(*g, obj)
					} else {
						// Version must be updated because component might be added to association view.
						tmp.Base().Version = obj.Base().Version
						obj = tmp
					}
				}
			} else if strings.EqualFold("Description", target) {
				s, err := reader.ReadElementContentAsString("Description", "")
				if err != nil {
					return err
				}
				if obj != nil && s != "" {
					obj.Base().Description = s
				}
			} else if strings.EqualFold("Version", target) {
				v, err := reader.ReadElementContentAsInt("Version", 0)
				if err != nil {
					return err
				}
				if obj != nil {
					obj.Base().Version = uint8(v)
				}
			} else if strings.EqualFold("Access", target) {
				s, err := reader.ReadElementContentAsString("Access", "")
				if err != nil {
					return err
				}
				if obj != nil && s != "" {
					pos := 0
					for i := 0; i < len(s); i++ {
						ch := s[i]
						pos++
						obj.Base().SetAccess(pos, enums.AccessMode(ch-0x30))
					}
				}
			} else if strings.EqualFold("Access3", target) {
				s, err := reader.ReadElementContentAsString("Access3", "")
				if err != nil {
					return err
				}
				if obj != nil && s != "" {
					tmp := s
					for pos := 0; pos != len(tmp)/4; pos++ {
						part := tmp[4*pos : 4*pos+4]
						v, err := strconv.ParseInt(part, 16, 32)
						if err != nil {
							return err
						}
						obj.Base().SetAccess3(1+pos, enums.AccessMode3(int(v)&^0x8000))
					}
				}
			} else if strings.EqualFold("MethodAccess", target) {
				s, err := reader.ReadElementContentAsString("MethodAccess", "")
				if err != nil {
					return err
				}
				if obj != nil && s != "" {
					pos := 0
					for i := 0; i < len(s); i++ {
						ch := s[i]
						pos++
						obj.Base().SetMethodAccess(pos, enums.MethodAccessMode(ch-0x30))
					}
				}
			} else if strings.EqualFold("MethodAccess3", target) {
				s, err := reader.ReadElementContentAsString("MethodAccess3", "")
				if err != nil {
					return err
				}
				if obj != nil && s != "" {
					tmp := s
					for pos := 0; pos != len(tmp)/4; pos++ {
						part := tmp[4*pos : 4*pos+4]
						v, err := strconv.ParseInt(part, 16, 32)
						if err != nil {
							return err
						}
						obj.Base().SetMethodAccess3(1+pos, enums.MethodAccessMode3(int(v)&^0x8000))
					}
				}
			} else if obj != nil {
				if base, ok := any(obj).(IGXDLMSBase); ok {
					if err := base.Load(reader); err != nil {
						return err
					}
				}
				obj = nil
			}
		} else {
			if err := reader.Read(); err != nil {
				return err
			}
		}
	}
	// PostLoad for all objects.
	for _, it := range reader.Objects {
		if err := it.PostLoad(reader); err != nil {
			return err
		}
	}
	return nil
}

// ----------------- Save methods on collection -----------------

func (c GXDLMSObjectCollection) SaveToFile(filename string, settings *GXXmlWriterSettings) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	writer := bufio.NewWriter(f)
	return c.SaveToStream(writer, settings)
}

func MissingMethods(v any, ifacePtr any) []string {
	tIface := reflect.TypeOf(ifacePtr).Elem() // ifacePtr pitää olla *SomeInterface
	tVal := reflect.TypeOf(v)
	if tVal == nil {
		return []string{"<nil value>"}
	}

	// Jos v ei ole pointer, mutta metodit on pointer receiverilla, tämä voi olla syy:
	// tValPtr := tVal; if tVal.Kind() != reflect.Pointer { tValPtr = reflect.PointerTo(tVal) }

	missing := []string{}
	for i := 0; i < tIface.NumMethod(); i++ {
		m := tIface.Method(i)

		// MethodByName löytää vain exportatut metodit toisesta paketista.
		// Jos interface sisältää unexportattuja metodeja toisesta paketista, reflect ei auta.
		if _, ok := tVal.MethodByName(m.Name); !ok {
			missing = append(missing, m.Name)
		}
	}
	return missing
}

func (c GXDLMSObjectCollection) SaveToStream(stream *bufio.Writer, settings *GXXmlWriterSettings) error {
	ignoreDescription := settings != nil && settings.IgnoreDescription
	omitXmlDeclaration := settings != nil && settings.OmitXmlDeclaration
	index := 0
	if settings != nil {
		index = settings.Index
	}

	lnVersion := 2
	for _, it := range c {
		if _, ok := any(it).(*GXDLMSAssociationLogicalName); ok {
			lnVersion = int(it.Base().Version)
			break
		}
	}
	writer := NewGXXmlWriterStream(stream, settings)
	defer writer.Close()

	if !omitXmlDeclaration && !ignoreDescription {
		if err := writer.WriteStartDocument(); err != nil {
			return err
		}
	}

	if err := writer.WriteStartElement("Objects"); err != nil {
		return err
	}

	for _, it := range c {
		base, ok := any(it).(IGXDLMSBase)
		if !ok {
			fmt.Printf("Type %T does NOT implement IGXDLMSBase\n", it)
			fmt.Println("Missing:", MissingMethods(it, (*IGXDLMSBase)(nil)))
			continue
		}
		if index < 2 {
			if settings == nil {
				// <GXDLMS{ObjectType}>
				if err := writer.WriteStartElement("GXDLMS" + fmt.Sprint(it.Base().ObjectType())); err != nil {
					return err
				}
			} else {
				// <Object Type="...">
				if err := writer.WriteStartElementWithAttrs("Object", map[string]string{
					"Type": strconv.Itoa(int(it.Base().ObjectType())),
				}); err != nil {
					return err
				}
			}
		}

		if index == 0 {
			// SN
			if it.Base().ShortName != 0 {
				if err := writer.WriteElementStringInt("SN", int(it.Base().ShortName)); err != nil {
					return err
				}
			}
			// LN
			if err := writer.WriteElementString("LN", it.Base().LogicalName()); err != nil {
				return err
			}
			// Version
			if it.Base().Version != 0 {
				if err := writer.WriteElementStringInt("Version", int(it.Base().Version)); err != nil {
					return err
				}
				if _, ok := any(it).(GXDLMSAssociationLogicalName); ok {
					lnVersion = int(it.Base().Version)
				}
			}
			// Description
			if !ignoreDescription && it.Base().Description != "" {
				if err := writer.WriteElementString("Description", it.Base().Description); err != nil {
					return err
				}
			}

			// Access rights
			if lnVersion < 3 {
				var sb strings.Builder
				for pos := 1; pos != base.GetAttributeCount()+1; pos++ {
					sb.WriteString(strconv.Itoa(int(it.Base().GetAccess(pos))))
				}
				if err := writer.WriteElementString("Access", sb.String()); err != nil {
					return err
				}

				sb.Reset()
				for pos := 1; pos != base.GetMethodCount()+1; pos++ {
					sb.WriteString(strconv.Itoa(int(it.Base().GetMethodAccess(pos))))
				}
				if err := writer.WriteElementString("MethodAccess", sb.String()); err != nil {
					return err
				}
			} else {
				var sb strings.Builder
				for pos := 1; pos != base.GetAttributeCount()+1; pos++ {
					// set highest bit; 4 hex chars
					value := 0x8000 | int(it.Base().GetAccess3(pos))
					sb.WriteString(fmt.Sprintf("%04X", value))
				}
				if err := writer.WriteElementString("Access3", sb.String()); err != nil {
					return err
				}

				sb.Reset()
				for pos := 1; pos != base.GetMethodCount()+1; pos++ {
					value := 0x8000 | int(it.Base().GetMethodAccess3(pos))
					sb.WriteString(fmt.Sprintf("%04X", value))
				}
				if err := writer.WriteElementString("MethodAccess3", sb.String()); err != nil {
					return err
				}
			}
		}

		if settings == nil || settings.Values {
			if err := base.Save(writer); err != nil {
				return err
			}
		}

		if index < 2 {
			if err := writer.WriteEndElement(); err != nil {
				return err
			}
		}
	}

	if err := writer.WriteEndElement(); err != nil { // </Objects>
		return err
	}

	if !omitXmlDeclaration && !ignoreDescription {
		if err := writer.WriteEndDocument(); err != nil {
			return err
		}
	}
	writer.Flush()
	return nil
}
