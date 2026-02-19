package objects

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
// Gurux Device Framework is Open Source software you can redistribute it
// and/or modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation version 2 of the License.
// Gurux Device Framework is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY without even the implied warranty of
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
	"reflect"
	"strings"
	"time"

	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// Online help:
// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSProfileGeneric
type GXDLMSProfileGeneric struct {
	GXDLMSObject
	// Client uses this to save how values are access.
	AccessSelector enums.AccessRange

	// Client uses this to save from which date values are retrieved.
	From any

	// Client uses this to save to which date values are retrieved.
	To any

	// Data of profile generic.
	Buffer [][]any

	CaptureObjects []types.GXKeyValuePair[IGXDLMSBase, *GXDLMSCaptureObject]

	// How often values are captured.
	CapturePeriod uint32

	// How columns are sorted.
	SortMethod enums.SortMethod

	// Column that is used for sorting.
	SortObject IGXDLMSBase

	// Sort object's attribute index.
	SortAttributeIndex int

	// Sort object's data index.
	SortDataIndex int

	// Entries (rows) in Use.
	EntriesInUse uint32

	// Maximum Entries (rows) count.
	ProfileEntries uint32
}

func getObjects(settings *settings.GXDLMSSettings) *GXDLMSObjectCollection {
	ret := settings.Objects.(GXDLMSObjectCollection)
	return &ret
}

// base returns the base GXDLMSObject of the object.
func (g *GXDLMSProfileGeneric) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

// capture returns the copies the values of the objects to capture
// into the buffer by reading capture objects.
func (g *GXDLMSProfileGeneric) capture(server *internal.IGXDLMSServer, isInvoke bool) {
}

// Invoke returns the invokes method.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Invoke parameters.
func (g *GXDLMSProfileGeneric) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	switch e.Index {
	case 1:
		//TODO: g.reset()
	case 2:
		//TODO: g.capture(e.Server, true)
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil, nil
}

// GetAttributeIndexToRead returns the collection of attributes to read.
// If attribute is static and already read or device is returned HW error it is not returned.
//
// Parameters:
//
//	all: All items are returned even if they are read already.
//
// Returns:
//
//	Collection of attributes to read.
func (g *GXDLMSProfileGeneric) GetAttributeIndexToRead(all bool) []int {
	attributes := []int{}
	// LN is static and read only once.
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	// CaptureObjects
	if all || (len(g.CaptureObjects) == 0 && !g.CanRead(3)) {
		attributes = append(attributes, 3)
	}
	// CapturePeriod
	if all || !g.CanRead(4) {
		attributes = append(attributes, 4)
	}
	// SortMethod
	if all || !g.CanRead(5) {
		attributes = append(attributes, 5)
	}
	// SortObject
	if all || !g.CanRead(6) {
		attributes = append(attributes, 6)
	}
	attributes = append(attributes, 2)
	attributes = append(attributes, 7)
	// ProfileEntries
	if all || !g.CanRead(8) {
		attributes = append(attributes, 8)
	}
	return attributes
}

// GetNames returns the names of attribute indexes.
func (g *GXDLMSProfileGeneric) GetNames() []string {
	return []string{"Logical Name", "Buffer", "CaptureObjects", "Capture Period", "Sort Method", "Sort Object", "Entries In Use", "Profile Entries"}
}

// GetMethodNames returns the names of method indexes.
func (g *GXDLMSProfileGeneric) GetMethodNames() []string {
	return []string{"Reset", "Capture"}
}

// GetAttributeCount returns the amount of attributes.
//
// Returns:
//
//	Count of attributes.
func (g *GXDLMSProfileGeneric) GetAttributeCount() int {
	return 8
}

// GetMethodCount returns the amount of methods.
func (g *GXDLMSProfileGeneric) GetMethodCount() int {
	return 2
}

// GetData returns the buffer data.
//
// Parameters:
//
//	settings: DLMS settings.
//	columns: Columns to get. nil if not used.
func (g *GXDLMSProfileGeneric) GetData(settings *settings.GXDLMSSettings,
	e *internal.ValueEventArgs,
	table [][]any,
	columns []types.GXKeyValuePair[IGXDLMSBase, *GXDLMSCaptureObject]) ([]byte, error) {
	var err error
	pos := 0
	cols := columns
	if columns == nil {
		cols = g.CaptureObjects
	}
	types_ := make([]enums.DataType, len(cols))
	for _, it := range cols {
		if it.Value.AttributeIndex == 0 {
			types_[pos] = enums.DataTypeStructure
		} else {
			types_[pos], err = it.Key.GetDataType(it.Value.AttributeIndex)
			if err != nil {
				return nil, err
			}
		}
		pos++
	}
	columnStart := uint16(1)
	columnEnd := uint16(0)
	if e.Selector == 2 {
		arr := e.Parameters.([]any)
		columnStart = arr[2].(uint16)
		columnEnd = arr[3].(uint16)
	}
	if columnStart > 1 || columnEnd != 0 {
		pos := uint16(1)
		cols := []types.GXKeyValuePair[IGXDLMSBase, *GXDLMSCaptureObject]{}
		for _, it := range g.CaptureObjects {
			if !(pos < columnStart || pos > columnEnd) {
				cols = append(cols, it)
			}
			pos++
		}
	}
	data := types.GXByteBuffer{}
	if settings.Index == 0 {
		err = data.SetUint8(uint8(enums.DataTypeArray))
		if err != nil {
			return nil, err
		}
		if e.RowEndIndex != 0 {
			err = types.SetObjectCount(int((e.RowEndIndex - e.RowBeginIndex)), &data)
		} else {
			err = types.SetObjectCount(len(table), &data)
		}
		if err != nil {
			return nil, err
		}
	}
	for _, items := range table {
		err = data.SetUint8(uint8(enums.DataTypeStructure))
		if err != nil {
			return nil, err
		}
		err = types.SetObjectCount(len(cols), &data)
		if err != nil {
			return nil, err
		}
		pos = 0
		var tp enums.DataType
		for _, value := range items {
			if cols == nil || internal.Contains(cols, g.CaptureObjects[pos]) {
				tp = types_[pos]
				if tp == enums.DataTypeNone {
					tp, err = internal.GetDLMSDataType(reflect.TypeOf(value))
					if err != nil {
						return nil, err
					}
					types_[pos] = tp
				}
				if value == nil {
					tp = enums.DataTypeNone
				}
				if v, ok := value.(IGXDLMSBase); ok {
					err = internal.SetData(settings, &data, tp, v.GetValues())
				} else {
					err = internal.SetData(settings, &data, tp, value)
				}
			}
			pos++
		}
		settings.Index++
	}
	if e.RowEndIndex != 0 {
		e.RowBeginIndex = uint32(len(table))
	} else {
		settings.Index = 0
	}
	return data.Array(), nil
}

// GetColumns returns the get selected (filtered) columns.
//
// Parameters:
//
//	cols: Selected columns.
//
// Returns:
//
//	Selected columns.
func (g *GXDLMSProfileGeneric) GetColumns(cols []any) ([]types.GXKeyValuePair[IGXDLMSBase, *GXDLMSCaptureObject], error) {
	var columns []types.GXKeyValuePair[IGXDLMSBase, *GXDLMSCaptureObject]
	if len(cols) != 0 {
		for _, tmp := range cols {
			it := tmp.([]any)
			ot := enums.ObjectType(it[0].(int16))
			ln, err := helpers.ToLogicalName(it[1].([]byte))
			if err != nil {
				return nil, err
			}
			attributeIndex := it[2].(int)
			dataIndex := it[3].(int)
			for _, c := range g.CaptureObjects {
				if c.Key.Base().ObjectType() == ot && c.Value.AttributeIndex == attributeIndex && c.Value.DataIndex == dataIndex && strings.Compare(c.Key.Base().LogicalName(), ln) == 0 {
					columns = append(columns, c)
					break
				}

			}
		}
	}
	return columns, nil
}

func (g *GXDLMSProfileGeneric) getProfileGenericData(settings *settings.GXDLMSSettings,
	e *internal.ValueEventArgs) ([]byte, error) {
	var columns []types.GXKeyValuePair[IGXDLMSBase, *GXDLMSCaptureObject]
	// If all data is read.
	if e.Selector == 0 || e.Parameters == nil || e.RowEndIndex != 0 {
		return g.GetData(settings, e, g.Buffer, columns)
	}
	table := [][]any{}
	return g.GetData(settings, e, table, columns)
}

// GetColumns returns the captured objects.
//
// Parameters:
//
//	settings: DLMS settings.
func (g *GXDLMSProfileGeneric) getColumns(settings *settings.GXDLMSSettings) ([]byte, error) {
	data := types.NewGXByteBuffer()
	err := data.SetUint8(uint8(enums.DataTypeArray))
	if err != nil {
		return nil, err
	}
	err = types.SetObjectCount(len(g.CaptureObjects), data)
	if err != nil {
		return nil, err
	}
	for _, it := range g.CaptureObjects {
		err = data.SetUint8(uint8(enums.DataTypeStructure))
		if err != nil {
			return nil, err
		}
		err = data.SetUint8(4)
		if err != nil {
			return nil, err
		}
		ln, err := helpers.LogicalNameToBytes(it.Key.Base().LogicalName())
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, data, enums.DataTypeUint16, it.Key.Base().ObjectType())
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, data, enums.DataTypeOctetString, ln)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, data, enums.DataTypeInt8, it.Value.AttributeIndex)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, data, enums.DataTypeUint16, it.Value.DataIndex)
		if err != nil {
			return nil, err
		}
	}
	return data.Array(), nil
}

// GetValue returns the value of given attribute.
// When raw parameter us not used example register multiplies value by scalar.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Get parameters.
//
// Returns:
//
//	Value of the attribute index.
func (g *GXDLMSProfileGeneric) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	var err error
	if e.Index == 1 {
		if e.Index == 1 {
			v, err := helpers.LogicalNameToBytes(g.LogicalName())
			if err != nil {
				e.Error = enums.ErrorCodeReadWriteDenied
			}
			return v, err
		}
	}
	if e.Index == 2 {
		return g.getProfileGenericData(settings, e)
	}
	if e.Index == 3 {
		return g.getColumns(settings)
	}
	if e.Index == 4 {
		return g.CapturePeriod, nil
	}
	if e.Index == 5 {
		return g.SortMethod, nil
	}
	if e.Index == 6 {
		data := types.NewGXByteBuffer()
		err = data.SetUint8(uint8(enums.DataTypeStructure))
		if err != nil {
			return nil, err
		}
		err = data.SetUint8(uint8(4))
		if err != nil {
			return nil, err
		}
		if g.SortObject == nil {
			err = internal.SetData(settings, data, enums.DataTypeUint16, 0)
			if err != nil {
				return nil, err
			}
			err = internal.SetData(settings, data, enums.DataTypeOctetString, make([]byte, 6))
			if err != nil {
				return nil, err
			}
			err = internal.SetData(settings, data, enums.DataTypeInt8, 0)
			if err != nil {
				return nil, err
			}
			err = internal.SetData(settings, data, enums.DataTypeUint16, 0)
		} else {
			err = internal.SetData(settings, data, enums.DataTypeUint16, g.SortObject.Base().ObjectType())
			if err != nil {
				return nil, err
			}
			ln, err := helpers.LogicalNameToBytes(g.SortObject.Base().LogicalName())
			if err != nil {
				return nil, err
			}
			err = internal.SetData(settings, data, enums.DataTypeOctetString, ln)
			if err != nil {
				return nil, err
			}
			err = internal.SetData(settings, data, enums.DataTypeInt8, g.SortAttributeIndex)
			if err != nil {
				return nil, err
			}
			err = internal.SetData(settings, data, enums.DataTypeUint16, g.SortDataIndex)
		}
		if err != nil {
			return nil, err
		}
		return data.Array(), nil
	}
	if e.Index == 7 {
		return g.EntriesInUse, nil
	}
	if e.Index == 8 {
		return g.ProfileEntries, nil
	}
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, nil
}

func (g *GXDLMSProfileGeneric) setBuffer(settings *settings.GXDLMSSettings,
	e *internal.ValueEventArgs) error {
	var err error
	var cols []types.GXKeyValuePair[IGXDLMSBase, *GXDLMSCaptureObject]
	if _, ok := e.Parameters.(types.GXKeyValuePair[IGXDLMSBase, *GXDLMSCaptureObject]); ok {
		//TODO:  cols = v
	}
	if cols == nil {
		cols = g.CaptureObjects
	}
	if e.Value != nil {
		var index2 int
		var lastDate time.Time
		for _, tmp := range e.Value.([]any) {
			row := []any{}
			row = append(row, tmp.([]any)...)
			if len(cols) != 0 {
				if len(row) != len(cols) {
					return errors.New("The number of columns does not match.")
				}
				for pos := 0; pos != len(row); pos++ {
					if cols == nil {
						index2 = 0
					} else {
						index2 = cols[pos].Value.AttributeIndex
					}
					var type_ enums.DataType
					// Actaris SL 7000 and ACE 6000 returns 0.
					if index2 > 0 {
						type_ = cols[pos].Key.GetUIDataType(index2)
					} else {
						type_ = enums.DataTypeNone
					}
					if v, ok := row[pos].([]byte); ok {
						if type_ != enums.DataTypeNone {
							row[pos], err = internal.ChangeTypeFromByteArray(settings, v, type_)
							if err != nil {
								return err
							}
							if dt, ok := row[pos].(*types.GXDateTime); ok {
								lastDate = dt.Value
							}
						}
					} else if type_ == enums.DataTypeDateTime && row[pos] == nil && g.CapturePeriod != 0 {
						if !lastDate.IsZero() && len(g.Buffer) != 0 {
							lastDate = g.Buffer[len(g.Buffer)-1][pos].(*types.GXDateTime).Value
						}
						if !lastDate.IsZero() {
							if g.SortMethod == enums.SortMethodFiFo || g.SortMethod == enums.SortMethodSmallest {
								lastDate = lastDate.Add(time.Duration(g.CapturePeriod) * time.Second)
							} else {
								lastDate = lastDate.Add(-time.Duration(g.CapturePeriod) * time.Second)
							}
							row[pos] = types.GXDateTime{Value: lastDate}
						}
					} else if type_ == enums.DataTypeDateTime {
						if _, ok := row[pos].(uint32); ok {
							row[pos] = types.GXDateTimeFromUnixTime(row[pos].(int64))
						} else if _, ok := row[pos].(uint64); ok {
							row[pos] = types.GXDateTimeFromHighResolutionTime(row[pos].(int64))
						}
					}
					if r, ok := cols[pos].Key.(*GXDLMSRegister); ok && index2 == 2 {
						scaler := r.Scaler()
						if scaler != 1 {
							row[pos] = row[pos].(float64) * scaler
						}
					} else if dr, ok := cols[pos].Key.(*GXDLMSDemandRegister); ok && (index2 == 2 || index2 == 3) {
						scaler := dr.Scaler()
						if scaler != 1 {
							row[pos] = row[pos].(float64) * scaler
						}
					} else if r, ok := cols[pos].Key.(*GXDLMSRegister); ok && index2 == 3 {
						v := internal.NewValueEventArgs3(r, 3, 0, nil)
						v.Value = row[pos]
						r.SetValue(nil, v)
						row[pos] = []any{r.Scaler, r.Unit}
					}
				}
				g.Buffer = append(g.Buffer, row)
			}
			if settings.IsServer() {
				g.EntriesInUse = uint32(len(g.Buffer))
			}
		}
	}
	return nil
}

func (g *GXDLMSProfileGeneric) setCaptureObjects(parent any,
	settings *settings.GXDLMSSettings,
	list []types.GXKeyValuePair[IGXDLMSBase, *GXDLMSCaptureObject],
	array types.GXArray) error {
	for _, it := range array {
		tmp := it.(types.GXStructure)
		if len(tmp) != 4 {
			return errors.New("Invalid structure format.")
		}
		v := tmp[0].(uint16)
		if !internal.Contains(enums.AllObjectType(), enums.ObjectType(v)) {
			return errors.New("Invalid object type.")
		}
		type_ := enums.ObjectType(v)
		ln, err := helpers.ToLogicalName(tmp[1].([]byte))
		if err != nil {
			return err
		}
		attributeIndex := tmp[2].(int8)
		dataIndex := tmp[3].(uint16)
		var obj IGXDLMSBase
		if settings != nil && settings.Objects != nil {
			obj = getObjects(settings).FindByLN(type_, ln)
		}
		// Create a new instance to avoid circular references.
		if obj == nil || obj == parent.(IGXDLMSBase) {
			obj = CreateObject(type_)
			/*TODO:
			if c == nil {
				c = GXDLMSConverter(enums.StandardDLMS if settings == nil else settings.Standard())
			}
			c.UpdateOBISCodeInformation(obj)
			*/
		}
		list = append(list, *types.NewGXKeyValuePair(obj, NewGXDLMSCaptureObject(int(attributeIndex), int(dataIndex))))
	}
	return nil
}

// SetValue returns the set value of given attribute.
// When raw parameter us not used example register multiplies value by scalar.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Set parameters.
func (g *GXDLMSProfileGeneric) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	if e.Index == 1 {
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return g.SetLogicalName(ln)
	} else if e.Index == 2 {
		g.setBuffer(settings, e)
	} else if e.Index == 3 {
		if settings != nil && settings.IsServer() {
			g.reset()
		}
		// Clear file
		if e.Server != nil {
			list := []*internal.ValueEventArgs{internal.NewValueEventArgs3(g, 1, 0, nil)}
			e.Server.NotifyPreAction(list)
			e.Server.NotifyPostAction(list)
		}
		g.CaptureObjects = g.CaptureObjects[:0]
		if e.Value != nil {
			g.setCaptureObjects(g, settings, g.CaptureObjects, e.Value.(types.GXArray))
		}
	} else if e.Index == 4 {
		// Any write access to one of the attributes will automatically call a resetand this call will propagate to all other profiles capturing this profile.
		if settings != nil && settings.IsServer() {
			g.reset()
		}
		g.CapturePeriod = e.Value.(uint32)
	} else if e.Index == 5 {
		// Any write access to one of the attributes will automatically call a resetand this call will propagate to all other profiles capturing this profile.
		if settings != nil && settings.IsServer() {
			g.reset()
		}
		g.SortMethod = enums.SortMethod(e.Value.(int))
	} else if e.Index == 6 {
		// Any write access to one of the attributes will automatically call a resetand this call will propagate to all other profiles capturing this profile.
		if settings != nil && settings.IsServer() {
			g.reset()
		}
		tmp := e.Value.([]any)
		if tmp != nil {
			if len(tmp) != 4 {
				return errors.New("Invalid structure format.")
			}
			type_ := enums.ObjectType(tmp[0].(int16))
			if type_ != enums.ObjectTypeNone {
				ln, err := helpers.ToLogicalName(tmp[1].([]byte))
				if err != nil {
					return err
				}
				g.SortAttributeIndex = tmp[2].(int)
				g.SortDataIndex = tmp[3].(int)
				g.SortObject = nil
				for _, it := range g.CaptureObjects {
					if it.Key.Base().ObjectType() == type_ && it.Key.Base().LogicalName() == ln {
						g.SortObject = it.Key
						break
					}
				}
				if g.SortObject == nil {
					g.SortObject = CreateObject(type_)
					g.SortObject.Base().SetLogicalName(ln)
				}
			} else {
				g.SortObject = nil
				g.SortAttributeIndex = 0
				g.SortDataIndex = 0
			}
		} else {
			g.SortObject = nil
		}
	} else if e.Index == 7 {
		g.EntriesInUse = e.Value.(uint32)
	} else if e.Index == 8 {
		// Any write access to one of the attributes will automatically call a resetand this call will propagate to all other profiles capturing this profile.
		if settings != nil && settings.IsServer() {
			g.reset()
		}
		g.ProfileEntries = e.Value.(uint32)
	} else {
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil
}

// Load returns the load object content from XML.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSProfileGeneric) Load(reader *GXXmlReader) error {
	var err error
	g.Buffer = g.Buffer[:0]
	if reader.isStartElementNamed2("Buffer", true) {
		for reader.isStartElementNamed2("Row", true) {
			row := [][]any{}
			for reader.isStartElementNamed2("Cell", false) {
				ret, err := reader.ReadElementContentAsObject("Cell", nil, nil, 0)
				row = append(row, ret.([]any))
				if err != nil {
					return err
				}
			}
			g.Buffer = append(g.Buffer, row...)
		}
		reader.ReadEndElement("Buffer")
	}
	g.CaptureObjects = g.CaptureObjects[:0]
	if reader.isStartElementNamed2("CaptureObjects", true) {
		for reader.isStartElementNamed2("Item", true) {
			ret, err := reader.ReadElementContentAsInt("ObjectType", 0)
			if err != nil {
				return err
			}
			ot := enums.ObjectType(ret)
			ln, err := reader.ReadElementContentAsString("LN", "")
			if err != nil {
				return err
			}
			ai, err := reader.ReadElementContentAsInt("Attribute", 0)
			if err != nil {
				return err
			}
			di, err := reader.ReadElementContentAsInt("Data", 0)
			if err != nil {
				return err
			}
			co := NewGXDLMSCaptureObject(ai, di)
			obj := reader.Objects.FindByLN(ot, ln)
			if obj == nil {
				obj = CreateObject(ot)
				obj.Base().SetLogicalName(ln)
			}
			g.CaptureObjects = append(g.CaptureObjects, *types.NewGXKeyValuePair(obj, co))
		}
		reader.ReadEndElement("CaptureObjects")
	}
	ret, err := reader.ReadElementContentAsInt("CapturePeriod", 0)
	if err != nil {
		return err
	}
	g.CapturePeriod = uint32(ret)
	ret, err = reader.ReadElementContentAsInt("SortMethod", 0)
	if err != nil {
		return err
	}
	g.SortMethod = enums.SortMethod(ret)
	if reader.isStartElementNamed2("SortObject", true) {
		ret, err := reader.ReadElementContentAsInt("ObjectType", 0)
		if err != nil {
			return err
		}
		ot := enums.ObjectType(ret)
		ln, err := reader.ReadElementContentAsString("LN", "")
		g.SortObject = reader.Objects.FindByLN(ot, ln)
		reader.ReadEndElement("SortObject")
	}
	ret, err = reader.ReadElementContentAsInt("EntriesInUse", 0)
	if err != nil {
		return err
	}
	g.EntriesInUse = uint32(ret)
	ret, err = reader.ReadElementContentAsInt("ProfileEntries", 0)
	if err != nil {
		return err
	}
	g.ProfileEntries = uint32(ret)
	return err
}

// Save returns the save object content to XML.
//
// Parameters:
//
//	writer: XML writer.
func (g *GXDLMSProfileGeneric) Save(writer *GXXmlWriter) error {
	var err error
	writer.WriteStartElement("Buffer")
	if g.Buffer != nil {
		var lastdt *time.Time
		add := g.CapturePeriod
		// Some meters are returning 0 if capture period is one hour.
		if add == 0 {
			add = 60
		}
		// Get data types.
		list := []enums.DataType{}
		if len(g.CaptureObjects) != 0 {
			for _, it := range g.CaptureObjects {
				if it.Value.AttributeIndex == 0 {
					ret, err := it.Key.GetDataType(it.Value.AttributeIndex)
					if err != nil {
						return err
					}
					list = append(list, ret)
				} else {
					list = append(list, enums.DataTypeNone)
				}
			}
		}
		for _, row := range g.Buffer {
			writer.WriteStartElement("Row")
			pos := 0
			for _, it := range row {
				//If capture objects is not read.
				if len(g.CaptureObjects) > pos {
					c := g.CaptureObjects[pos]
					pos++
					if v, ok := c.Key.(*GXDLMSClock); ok && c.Value.AttributeIndex == 2 {
						if it != nil {
							lastdt = &v.Time.Value
						} else if lastdt != nil {
							//TODO: lastdt = types.NewGXDateTime(lastdt.Value.AddMinutes(add))
							//TODO: writer.WriteElementObject("Cell", lastdt)
							continue
						} else {
							writer.WriteElementObject("Cell", nil, enums.DataTypeDateTime, enums.DataTypeDateTime)
						}
					}
				}
				writer.WriteElementObject("Cell", it, enums.DataTypeNone, enums.DataTypeNone)
			}
			writer.WriteEndElement()
		}
	}
	writer.WriteEndElement()
	writer.WriteStartElement("CaptureObjects")
	if g.CaptureObjects != nil {
		for _, it := range g.CaptureObjects {
			writer.WriteStartElement("Item")
			writer.WriteElementString("ObjectType", int(it.Key.Base().ObjectType()))
			writer.WriteElementString("LN", it.Key.Base().LogicalName())
			writer.WriteElementString("Attribute", it.Value.AttributeIndex)
			writer.WriteElementString("Data", it.Value.DataIndex)
			writer.WriteEndElement()
		}
	}
	writer.WriteEndElement()
	err = writer.WriteElementString("CapturePeriod", g.CapturePeriod)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("SortMethod", int(g.SortMethod))
	if err != nil {
		return err
	}
	writer.WriteStartElement("SortObject")
	if g.SortObject != nil {
		err = writer.WriteElementString("ObjectType", int(g.SortObject.Base().ObjectType()))
		if err != nil {
			return err
		}
		err = writer.WriteElementString("LN", g.SortObject.Base().LogicalName)
		if err != nil {
			return err
		}
	}
	writer.WriteEndElement()
	err = writer.WriteElementString("EntriesInUse", g.EntriesInUse)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("ProfileEntries", g.ProfileEntries)
	if err != nil {
		return err
	}
	return err
}

// PostLoad returns the handle actions after Load.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSProfileGeneric) PostLoad(reader *GXXmlReader) error {
	//Upload capture objects after load.
	if len(g.CaptureObjects) != 0 {
		columns := []types.GXKeyValuePair[IGXDLMSBase, *GXDLMSCaptureObject]{}
		for _, it := range g.CaptureObjects {
			var obj = it.Key
			target := reader.Objects.FindByLN(obj.Base().ObjectType(), obj.Base().LogicalName())
			if target != nil && target != obj {
				obj = target
			}
			columns = append(columns, types.GXKeyValuePair[IGXDLMSBase, *GXDLMSCaptureObject]{Key: obj, Value: it.Value})
		}
		g.CaptureObjects = g.CaptureObjects[:0]
		g.CaptureObjects = append(g.CaptureObjects, columns...)
	}
	return nil
}

// GetValues returns the an array containing the COSEM object's attribute values.
func (g *GXDLMSProfileGeneric) GetValues() []any {
	return []any{g.LogicalName(), g.Buffer, g.CaptureObjects,
		g.CapturePeriod, g.SortMethod,
		g.SortObject, g.EntriesInUse, g.ProfileEntries,
	}
}

// GetCaptureObject returns the get captured objects.
func (g *GXDLMSProfileGeneric) GetCaptureObject() []IGXDLMSBase {
	list := []IGXDLMSBase{}
	for _, it := range g.CaptureObjects {
		list = append(list, it.Key)
	}
	return list
}

// Reset returns the clears the buffer.
//
// Parameters:
//
//	client: DLMS client.
//
// Returns:
//
//	Action bytes.
func (g *GXDLMSProfileGeneric) Reset(client IGXDLMSClient) ([][]byte, error) {
	return client.Method(g, 1, int8(0), enums.DataTypeInt8)
}

// Capture returns the copies the values of the objects to capture into the buffer by reading each capture object.
//
// Parameters:
//
//	client: DLMS client.
//
// Returns:
//
//	Action bytes.
func (g *GXDLMSProfileGeneric) Capture(client IGXDLMSClient) ([][]byte, error) {
	return client.Method(g, 2, int8(0), enums.DataTypeInt8)
}

// Reset returns the clears the buffer.
func (g *GXDLMSProfileGeneric) reset() {
}

// GetSelectedColumns returns the get selected columns from parameters.
//
// Parameters:
//
//	selector: Is read by entry or range.
//	parameters: Received parameters where columns information is found.
//
// Returns:
//
//	Selected columns.
func (g *GXDLMSProfileGeneric) GetSelectedColumns(selector int,
	parameters any) ([]types.GXKeyValuePair[IGXDLMSBase, *GXDLMSCaptureObject], error) {
	columns := []types.GXKeyValuePair[IGXDLMSBase, *GXDLMSCaptureObject]{}
	switch selector {
	case 0:
		columns = append(columns, g.CaptureObjects...)
		return columns, nil
	case 1:
		arr := parameters.([]any)
		return g.GetColumns(arr)
	case 2:
		arr := parameters.([]any)
		colStart := 1
		colCount := 0
		if len(arr) > 2 {
			colStart = arr[2].(int)
		}
		if len(arr) > 3 {
			colCount = arr[3].(int)
		} else if colStart != 1 {
			colCount = len(g.CaptureObjects)
		}
		if colStart != 1 || colCount != 0 {
			for pos := 0; pos != colCount; pos++ {
				columns = append(columns, g.CaptureObjects[colStart+pos-1])
			}
		} else {
			columns = append(columns, g.CaptureObjects...)
		}
		return columns, nil
	default:
		return nil, errors.New("Invalid selector.")
	}
}

// GetDataType returns the device data type_ of selected attribute index.
//
// Parameters:
//
//	index: Attribute index of the object.
//
// Returns:
//
//	Device data type_ of the object.
func (g *GXDLMSProfileGeneric) GetDataType(index int) (enums.DataType, error) {
	var ret enums.DataType
	switch index {
	case 1:
		ret = enums.DataTypeOctetString
	case 2:
		ret = enums.DataTypeArray
	case 3:
		ret = enums.DataTypeArray
	case 4:
		ret = enums.DataTypeUint32
	case 5:
		ret = enums.DataTypeEnum
	case 6:
		ret = enums.DataTypeStructure
	case 7:
		ret = enums.DataTypeUint32
	case 8:
		ret = enums.DataTypeUint32
	default:
		return 0, errors.New("GetDataType failed. Invalid attribute index.")
	}
	return ret, nil
}

// GetCaptureObjects returns the get capture objects.
//
// Parameters:
//
//	array: Received data.
func (g *GXDLMSProfileGeneric) GetCaptureObjects(array []any) []types.GXKeyValuePair[IGXDLMSBase, *GXDLMSCaptureObject] {
	list := []types.GXKeyValuePair[IGXDLMSBase, *GXDLMSCaptureObject]{}
	g.setCaptureObjects(nil, nil, list, array)
	return list
}

// Constructor.
// ln: Logical Name of the object.
// sn: Short Name of the object.
func NewGXDLMSProfileGeneric(ln string, sn int16) (*GXDLMSProfileGeneric, error) {
	err := ValidateLogicalName(ln)
	if err != nil {
		return nil, err
	}
	return &GXDLMSProfileGeneric{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeProfileGeneric,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
