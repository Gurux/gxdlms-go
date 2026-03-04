package objects

import (
	"fmt"
	"reflect"

	"github.com/Gurux/gxdlms-go/dlmserrors"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// GXDLMSTargetObject identifies a target object and attribute.
type GXDLMSTargetObject struct {
	Target         IGXDLMSBase
	AttributeIndex int8
}

// GXDLMSArrayManagerItem maps an ID to an array target.
type GXDLMSArrayManagerItem struct {
	ID      uint8
	Element *GXDLMSTargetObject
}

// GXDLMSArrayManager controls managed array objects.
type GXDLMSArrayManager struct {
	GXDLMSObject
	Elements []GXDLMSArrayManagerItem
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSArrayManager) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

func (g *GXDLMSArrayManager) GetValues() []any { return []any{g.LogicalName(), g.Elements} }

func (g *GXDLMSArrayManager) GetAttributeIndexToRead(all bool) []int {
	var a []int
	if all || g.LogicalName() == "" {
		a = append(a, 1)
	}
	if all || g.CanRead(2) {
		a = append(a, 2)
	}
	return a
}

func (g *GXDLMSArrayManager) GetNames() []string { return []string{"Logical Name", "Objects"} }

func (g *GXDLMSArrayManager) GetMethodNames() []string {
	return []string{"Amount", "Retrieve", "Insert", "Update", "Remove"}
}

func (g *GXDLMSArrayManager) GetAttributeCount() int { return 2 }

func (g *GXDLMSArrayManager) GetMethodCount() int { return 5 }

func (g *GXDLMSArrayManager) GetDataType(index int) (enums.DataType, error) {
	switch index {
	case 1:
		return enums.DataTypeOctetString, nil
	case 2:
		return enums.DataTypeArray, nil
	default:
		return enums.DataTypeNone, dlmserrors.ErrInvalidAttributeIndex
	}
}

// NumberOfEntries returns the number of entries in the target array.
func (g *GXDLMSArrayManager) NumberOfEntries(client IGXDLMSClient, id uint8) ([][]byte, error) {
	return client.Method(g, 1, id, enums.DataTypeUint8)
}

// ParseNumberOfEntries parses entry count from meter reply.
func (g *GXDLMSArrayManager) ParseNumberOfEntries(reply []byte) (int, error) {
	return g.ParseNumberOfEntriesFromBuffer(types.NewGXByteBufferWithData(reply))
}

// ParseNumberOfEntriesFromBuffer parses entry count from meter reply buffer.
func (g *GXDLMSArrayManager) ParseNumberOfEntriesFromBuffer(reply *types.GXByteBuffer) (int, error) {
	info := &internal.GXDataInfo{}
	v, err := internal.GetData(nil, reply, info)
	if err != nil {
		return 0, err
	}
	u, err := toUint32(v)
	if err != nil {
		return 0, err
	}
	return int(u), nil
}

// RetrieveEntries returns entries from the given range.
func (g *GXDLMSArrayManager) RetrieveEntries(client IGXDLMSClient, id uint8, from uint16, to uint16) ([][]byte, error) {
	conf, _ := client.Settings().(*settings.GXDLMSSettings)
	if conf == nil {
		conf = &settings.GXDLMSSettings{}
	}
	bb := types.NewGXByteBuffer()
	_ = bb.SetUint8(uint8(enums.DataTypeStructure))
	_ = bb.SetUint8(2)
	_ = internal.SetData(conf, bb, enums.DataTypeUint8, id)
	_ = bb.SetUint8(uint8(enums.DataTypeStructure))
	_ = bb.SetUint8(2)
	_ = internal.SetData(conf, bb, enums.DataTypeUint16, from)
	_ = internal.SetData(conf, bb, enums.DataTypeUint16, to)
	return client.Method(g, 2, bb.Array(), enums.DataTypeStructure)
}

// ParseEntries parses entries from meter reply.
func (g *GXDLMSArrayManager) ParseEntries(reply *types.GXByteBuffer) (types.GXArray, error) {
	info := &internal.GXDataInfo{}
	v, err := internal.GetData(nil, reply, info)
	if err != nil {
		return nil, err
	}
	if arr, ok := v.(types.GXArray); ok {
		return arr, nil
	}
	return nil, fmt.Errorf("invalid entry array type: %T", v)
}

// InsertEntry inserts a new entry.
func (g *GXDLMSArrayManager) InsertEntry(client IGXDLMSClient, id uint8, index uint16, entry any) ([][]byte, error) {
	conf, _ := client.Settings().(*settings.GXDLMSSettings)
	if conf == nil {
		conf = &settings.GXDLMSSettings{}
	}
	bb := types.NewGXByteBuffer()
	_ = bb.SetUint8(uint8(enums.DataTypeStructure))
	_ = bb.SetUint8(2)
	_ = internal.SetData(conf, bb, enums.DataTypeUint8, id)
	_ = bb.SetUint8(uint8(enums.DataTypeStructure))
	_ = bb.SetUint8(2)
	_ = internal.SetData(conf, bb, enums.DataTypeUint16, index)
	dt, err := internal.GetDLMSDataType(reflectTypeOf(entry))
	if err != nil {
		return nil, err
	}
	_ = internal.SetData(conf, bb, dt, entry)
	return client.Method(g, 3, bb.Array(), enums.DataTypeStructure)
}

// UpdateEntry updates an entry.
func (g *GXDLMSArrayManager) UpdateEntry(client IGXDLMSClient, id uint8, index uint16, entry any) ([][]byte, error) {
	conf, _ := client.Settings().(*settings.GXDLMSSettings)
	if conf == nil {
		conf = &settings.GXDLMSSettings{}
	}
	bb := types.NewGXByteBuffer()
	_ = bb.SetUint8(uint8(enums.DataTypeStructure))
	_ = bb.SetUint8(2)
	_ = internal.SetData(conf, bb, enums.DataTypeUint8, id)
	_ = bb.SetUint8(uint8(enums.DataTypeStructure))
	_ = bb.SetUint8(2)
	_ = internal.SetData(conf, bb, enums.DataTypeUint16, index)
	dt, err := internal.GetDLMSDataType(reflectTypeOf(entry))
	if err != nil {
		return nil, err
	}
	_ = internal.SetData(conf, bb, dt, entry)
	return client.Method(g, 4, bb.Array(), enums.DataTypeStructure)
}

// RemoveEntries removes entries from the given range.
func (g *GXDLMSArrayManager) RemoveEntries(client IGXDLMSClient, id uint8, from uint16, to uint16) ([][]byte, error) {
	conf, _ := client.Settings().(*settings.GXDLMSSettings)
	if conf == nil {
		conf = &settings.GXDLMSSettings{}
	}
	bb := types.NewGXByteBuffer()
	_ = bb.SetUint8(uint8(enums.DataTypeStructure))
	_ = bb.SetUint8(2)
	_ = internal.SetData(conf, bb, enums.DataTypeUint8, id)
	_ = bb.SetUint8(uint8(enums.DataTypeStructure))
	_ = bb.SetUint8(2)
	_ = internal.SetData(conf, bb, enums.DataTypeUint16, from)
	_ = internal.SetData(conf, bb, enums.DataTypeUint16, to)
	return client.Method(g, 5, bb.Array(), enums.DataTypeStructure)
}

func (g *GXDLMSArrayManager) getTargetArrayData(settings *settings.GXDLMSSettings, target IGXDLMSBase, attributeIndex int8, selector uint8, parameters any) (*types.GXByteBuffer, error) {
	dt, err := target.GetDataType(int(attributeIndex))
	if err != nil {
		return nil, err
	}
	if dt != enums.DataTypeArray {
		return nil, fmt.Errorf("target data type is not array")
	}
	arg := internal.NewValueEventArgs3(target, uint8(attributeIndex), selector, parameters)
	tmp, err := target.GetValue(settings, arg)
	if err != nil {
		return nil, err
	}
	data, ok := tmp.([]byte)
	if !ok {
		return nil, fmt.Errorf("invalid target value type: %T", tmp)
	}
	bb := types.NewGXByteBufferWithData(data)
	t, err := bb.Uint8()
	if err != nil {
		return nil, err
	}
	if enums.DataType(t) != enums.DataTypeArray {
		return nil, fmt.Errorf("target value is not array")
	}
	return bb, nil
}

func (g *GXDLMSArrayManager) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	e.ByteArray = true
	reply := types.NewGXByteBuffer()
	switch e.Index {
	case 1:
		id := e.Parameters.(uint8)
		g.numberOfEntries(settings, e, id, reply)
	case 2:
		args, ok := e.Parameters.(types.GXStructure)
		if !ok {
			e.Error = enums.ErrorCodeReadWriteDenied
			return nil, nil
		}
		g.retrieveEntries(settings, e, args, reply)
	case 3:
		args, ok := e.Parameters.(types.GXStructure)
		if !ok {
			e.Error = enums.ErrorCodeReadWriteDenied
			return nil, nil
		}
		g.insertEntry(settings, e, args)
	case 4:
		args, ok := e.Parameters.(types.GXStructure)
		if !ok {
			e.Error = enums.ErrorCodeReadWriteDenied
			return nil, nil
		}
		g.updateEntry(settings, e, args)
	case 5:
		args, ok := e.Parameters.(types.GXStructure)
		if !ok {
			e.Error = enums.ErrorCodeReadWriteDenied
			return nil, nil
		}
		g.removeEntries(settings, e, args)
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	if reply.Size() == 0 {
		return nil, nil
	}
	return reply.Array(), nil
}

func (g *GXDLMSArrayManager) numberOfEntries(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs, id uint8, reply *types.GXByteBuffer) {
	for _, it := range g.Elements {
		if it.ID != id || it.Element == nil || it.Element.Target == nil {
			continue
		}
		if pg, ok := it.Element.Target.(*GXDLMSProfileGeneric); ok && it.Element.AttributeIndex == 2 {
			_ = internal.SetData(settings, reply, enums.DataTypeUint32, pg.EntriesInUse)
			return
		}
		bb, err := g.getTargetArrayData(settings, it.Element.Target, it.Element.AttributeIndex, 0, nil)
		if err != nil {
			break
		}
		count, err := types.GetObjectCount(bb)
		if err != nil {
			break
		}
		var dt enums.DataType
		switch {
		case count <= 0xFF:
			dt = enums.DataTypeUint8
		case count <= 0xFFFF:
			dt = enums.DataTypeUint16
		default:
			dt = enums.DataTypeUint32
		}
		_ = internal.SetData(settings, reply, dt, count)
		return
	}
	e.Error = enums.ErrorCodeReadWriteDenied
}

func (g *GXDLMSArrayManager) retrieveEntries(settings *settings.GXDLMSSettings,
	e *internal.ValueEventArgs,
	args types.GXStructure,
	reply *types.GXByteBuffer) {
	if len(args) < 2 {
		e.Error = enums.ErrorCodeReadWriteDenied
		return
	}
	id := args[0].(uint8)
	rng := args[1].(types.GXStructure)
	if len(rng) < 2 {
		e.Error = enums.ErrorCodeReadWriteDenied
		return
	}
	from32 := rng[0].(uint32)
	to32, err := toUint32(rng[1])
	if err != nil {
		e.Error = enums.ErrorCodeReadWriteDenied
		return
	}
	if from32 == 0 || from32 > to32 {
		e.Error = enums.ErrorCodeReadWriteDenied
		return
	}
	for _, it := range g.Elements {
		if it.ID != id || it.Element == nil || it.Element.Target == nil {
			continue
		}
		var parameters any
		if _, ok := it.Element.Target.(*GXDLMSProfileGeneric); ok {
			parameters = []any{nil, nil, from32, to32}
		}
		bb, err := g.getTargetArrayData(settings, it.Element.Target, it.Element.AttributeIndex, 2, parameters)
		if err != nil {
			break
		}
		info := &internal.GXDataInfo{}
		data, err := internal.GetData(settings, bb, info)
		if err != nil {
			break
		}
		arr, ok := data.(types.GXArray)
		if !ok {
			break
		}
		from := int(from32 - 1)
		to := int(to32)
		if to < len(arr) {
			arr = arr[:to]
		}
		if from > 0 {
			if from >= len(arr) {
				arr = types.GXArray{}
			} else {
				arr = arr[from:]
			}
		}
		_ = internal.SetData(settings, reply, enums.DataTypeArray, arr)
		return
	}
	e.Error = enums.ErrorCodeReadWriteDenied
}

func (g *GXDLMSArrayManager) insertEntry(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs, args []any) {
	if len(args) < 2 {
		e.Error = enums.ErrorCodeReadWriteDenied
		return
	}
	id := args[0].(uint8)
	item := args[1].(types.GXStructure)
	if len(item) < 2 {
		e.Error = enums.ErrorCodeReadWriteDenied
		return
	}
	index32, err := toUint32(item[0])
	if err != nil {
		e.Error = enums.ErrorCodeReadWriteDenied
		return
	}
	data := item[1]
	for _, it := range g.Elements {
		if it.ID != id || it.Element == nil || it.Element.Target == nil {
			continue
		}
		bb, err := g.getTargetArrayData(settings, it.Element.Target, it.Element.AttributeIndex, 0, nil)
		if err != nil {
			break
		}
		info := &internal.GXDataInfo{}
		v, err := internal.GetData(settings, bb, info)
		if err != nil {
			break
		}
		arr, ok := v.(types.GXArray)
		if !ok {
			break
		}
		index := int(index32)
		if index == 0 {
			arr = append(types.GXArray{data}, arr...)
		} else if index > len(arr) {
			arr = append(arr, data)
		} else {
			pos := index
			arr = append(arr[:pos], append(types.GXArray{data}, arr[pos:]...)...)
		}
		arg := &internal.ValueEventArgs{Target: it.Element.Target, Index: uint8(it.Element.AttributeIndex), Value: arr}
		_ = it.Element.Target.SetValue(settings, arg)
		return
	}
	e.Error = enums.ErrorCodeReadWriteDenied
}

func (g *GXDLMSArrayManager) updateEntry(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs, args []any) {
	if len(args) < 2 {
		e.Error = enums.ErrorCodeReadWriteDenied
		return
	}
	id := args[0].(uint8)
	item := args[1].(types.GXStructure)
	if len(item) < 2 {
		e.Error = enums.ErrorCodeReadWriteDenied
		return
	}
	index32, err := toUint32(item[0])
	if err != nil || index32 == 0 {
		e.Error = enums.ErrorCodeReadWriteDenied
		return
	}
	data := item[1]
	for _, it := range g.Elements {
		if it.ID != id || it.Element == nil || it.Element.Target == nil {
			continue
		}
		bb, err := g.getTargetArrayData(settings, it.Element.Target, it.Element.AttributeIndex, 0, nil)
		if err != nil {
			break
		}
		info := &internal.GXDataInfo{}
		v, err := internal.GetData(settings, bb, info)
		if err != nil {
			break
		}
		arr, ok := v.(types.GXArray)
		if !ok {
			break
		}
		pos := int(index32 - 1)
		if pos < len(arr) {
			arr[pos] = data
			arg := &internal.ValueEventArgs{Target: it.Element.Target, Index: uint8(it.Element.AttributeIndex), Value: arr}
			_ = it.Element.Target.SetValue(settings, arg)
			return
		}
		break
	}
	e.Error = enums.ErrorCodeReadWriteDenied
}

func (g *GXDLMSArrayManager) removeEntries(settings *settings.GXDLMSSettings,
	e *internal.ValueEventArgs,
	args types.GXStructure) {
	if len(args) < 2 {
		e.Error = enums.ErrorCodeReadWriteDenied
		return
	}
	id := args[0].(uint8)
	rng, ok := args[1].(types.GXStructure)
	if !ok || len(rng) < 2 {
		e.Error = enums.ErrorCodeReadWriteDenied
		return
	}
	from32, err := toUint32(rng[0])
	if err != nil || from32 == 0 {
		e.Error = enums.ErrorCodeReadWriteDenied
		return
	}
	to32, err := toUint32(rng[1])
	if err != nil || from32 > to32 {
		e.Error = enums.ErrorCodeReadWriteDenied
		return
	}
	for _, it := range g.Elements {
		if it.ID != id || it.Element == nil || it.Element.Target == nil {
			continue
		}
		bb, err := g.getTargetArrayData(settings, it.Element.Target, it.Element.AttributeIndex, 0, nil)
		if err != nil {
			break
		}
		info := &internal.GXDataInfo{}
		v, err := internal.GetData(settings, bb, info)
		if err != nil {
			break
		}
		arr, ok := v.(types.GXArray)
		if !ok {
			break
		}
		from := int(from32 - 1)
		to := int(to32) // exclusive upper bound for one-based inclusive range.
		if from >= len(arr) {
			break
		}
		if to > len(arr) {
			to = len(arr)
		}
		arr = append(arr[:from], arr[to:]...)
		arg := &internal.ValueEventArgs{Target: it.Element.Target, Index: uint8(it.Element.AttributeIndex), Value: arr}
		_ = it.Element.Target.SetValue(settings, arg)
		return
	}
	e.Error = enums.ErrorCodeReadWriteDenied
}

func (g *GXDLMSArrayManager) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	switch e.Index {
	case 1:
		return helpers.LogicalNameToBytes(g.LogicalName())
	case 2:
		data := types.NewGXByteBuffer()
		_ = data.SetUint8(uint8(enums.DataTypeArray))
		types.SetObjectCount(len(g.Elements), data)
		for _, it := range g.Elements {
			if it.Element == nil || it.Element.Target == nil {
				continue
			}
			_ = data.SetUint8(uint8(enums.DataTypeStructure))
			_ = data.SetUint8(2)
			_ = internal.SetData(settings, data, enums.DataTypeUint8, it.ID)
			_ = data.SetUint8(uint8(enums.DataTypeStructure))
			_ = data.SetUint8(3)
			_ = internal.SetData(settings, data, enums.DataTypeUint16, uint16(it.Element.Target.Base().ObjectType()))
			ln, err := helpers.LogicalNameToBytes(it.Element.Target.Base().LogicalName())
			if err != nil {
				return nil, err
			}
			_ = internal.SetData(settings, data, enums.DataTypeOctetString, ln)
			_ = internal.SetData(settings, data, enums.DataTypeInt8, it.Element.AttributeIndex)
		}
		return data.Array(), nil
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
		return nil, nil
	}
}

func (g *GXDLMSArrayManager) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	switch e.Index {
	case 1:
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return err
		}
		return g.SetLogicalName(ln)
	case 2:
		g.Elements = g.Elements[:0]
		if e.Value != nil {
			for _, tmp := range e.Value.(types.GXArray) {
				item := tmp.(types.GXStructure)
				if len(item) < 2 {
					continue
				}
				id := item[0].(uint8)
				a := item[1].(types.GXStructure)
				if len(a) < 3 {
					continue
				}
				otv, err := toUint32(a[0])
				if err != nil {
					return err
				}
				ln, err := helpers.ToLogicalName(a[1])
				if err != nil {
					return err
				}
				attr32, err := toUint32(a[2])
				if err != nil {
					return err
				}
				ot := enums.ObjectType(otv)
				var obj IGXDLMSBase
				if settings != nil && settings.Objects != nil {
					switch objects := settings.Objects.(type) {
					case GXDLMSObjectCollection:
						obj = objects.FindByLN(ot, ln)
					case *GXDLMSObjectCollection:
						obj = objects.FindByLN(ot, ln)
					}
				}
				if obj == nil {
					obj, err = CreateObject(ot, ln, 0)
					if err != nil {
						return err
					}
				}
				g.Elements = append(g.Elements, GXDLMSArrayManagerItem{
					ID:      id,
					Element: NewGXDLMSTargetObject(obj, int8(attr32)),
				})
			}
		}
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil
}

func (g *GXDLMSArrayManager) Load(reader *GXXmlReader) error {
	g.Elements = g.Elements[:0]
	ok, err := reader.IsStartElementNamed("Elements", true)
	if err != nil {
		return err
	}
	if ok {
		for {
			ok, err = reader.IsStartElementNamed("Item", true)
			if err != nil {
				return err
			}
			if !ok {
				break
			}
			id, err := reader.ReadElementContentAsUInt8("Id", 0)
			if err != nil {
				return err
			}
			var element *GXDLMSTargetObject
			ok, err = reader.IsStartElementNamed("Target", true)
			if err != nil {
				return err
			}
			if ok {
				ot, err := reader.ReadElementContentAsInt("Type", 0)
				if err != nil {
					return err
				}
				ln, err := reader.ReadElementContentAsString("LN", "")
				if err != nil {
					return err
				}
				index, err := reader.ReadElementContentAsInt("Index", 0)
				if err != nil {
					return err
				}
				obj, err := CreateObject(enums.ObjectType(ot), ln, 0)
				if err != nil {
					return err
				}
				element = NewGXDLMSTargetObject(obj, int8(index))
				if err = reader.ReadEndElement("Target"); err != nil {
					return err
				}
			}
			if err = reader.ReadEndElement("Item"); err != nil {
				return err
			}
			g.Elements = append(g.Elements, GXDLMSArrayManagerItem{ID: id, Element: element})
		}
		if err = reader.ReadEndElement("Elements"); err != nil {
			return err
		}
	}
	return nil
}

func (g *GXDLMSArrayManager) Save(writer *GXXmlWriter) error {
	if err := writer.WriteStartElement("Elements"); err != nil {
		return err
	}
	for _, it := range g.Elements {
		if err := writer.WriteStartElement("Item"); err != nil {
			return err
		}
		if err := writer.WriteElementString("Id", it.ID); err != nil {
			return err
		}
		if it.Element != nil && it.Element.Target != nil {
			if err := writer.WriteStartElement("Target"); err != nil {
				return err
			}
			if err := writer.WriteElementString("Type", uint16(it.Element.Target.Base().ObjectType())); err != nil {
				return err
			}
			if err := writer.WriteElementString("LN", it.Element.Target.Base().LogicalName()); err != nil {
				return err
			}
			if err := writer.WriteElementString("Index", it.Element.AttributeIndex); err != nil {
				return err
			}
			if err := writer.WriteEndElement(); err != nil {
				return err
			}
		}
		if err := writer.WriteEndElement(); err != nil {
			return err
		}
	}
	return writer.WriteEndElement()
}

func (g *GXDLMSArrayManager) PostLoad(reader *GXXmlReader) error {
	if reader.Objects == nil {
		return nil
	}
	for pos := range g.Elements {
		it := &g.Elements[pos]
		if it.Element == nil || it.Element.Target == nil {
			continue
		}
		obj := reader.Objects.FindByLN(it.Element.Target.Base().ObjectType(), it.Element.Target.Base().LogicalName())
		if obj != nil {
			it.Element.Target = obj
		}
	}
	return nil
}

func reflectTypeOf(v any) reflect.Type {
	if v == nil {
		return nil
	}
	return reflect.TypeOf(v)
}

// NewGXDLMSArrayManager creates a new array manager object instance.
func NewGXDLMSArrayManager(ln string, sn int16) (*GXDLMSArrayManager, error) {
	if err := ValidateLogicalName(ln); err != nil {
		return nil, err
	}
	return &GXDLMSArrayManager{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeArrayManager,
			logicalName: ln,
			ShortName:   sn,
		},
		Elements: make([]GXDLMSArrayManagerItem, 0),
	}, nil
}

// NewGXDLMSTargetObject creates a new target object.
//
// The function validates `ln` before creating the object.
// `ln` is the Logical Name and `sn` is the Short Name of the object.
func NewGXDLMSTargetObject(target IGXDLMSBase, attributeIndex int8) *GXDLMSTargetObject {
	return &GXDLMSTargetObject{Target: target, AttributeIndex: attributeIndex}
}
