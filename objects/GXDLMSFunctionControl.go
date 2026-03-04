package objects

import (
	"fmt"

	"github.com/Gurux/gxdlms-go/dlmserrors"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// GXDLMSFunctionSpec maps a function name to a set of related COSEM objects.
type GXDLMSFunctionSpec struct {
	Name    string
	Objects []IGXDLMSBase
}

// GXDLMSFunctionControl models Function control object.
type GXDLMSFunctionControl struct {
	GXDLMSObject
	ActivationStatus []types.GXKeyValuePair[string, bool]
	FunctionList     []GXDLMSFunctionSpec
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSFunctionControl) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

func (g *GXDLMSFunctionControl) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	switch e.Index {
	case 1:
		updates, err := decodeFunctionStatus(e.Parameters)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return nil, err
		}
		for _, u := range updates {
			for i := range g.ActivationStatus {
				if g.ActivationStatus[i].Key == u.Key {
					g.ActivationStatus[i].Value = u.Value
					break
				}
			}
		}
	case 2:
		name, objects, err := decodeSingleFunctionSpec(e.Parameters)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return nil, err
		}
		g.removeFunction(name)
		g.FunctionList = append(g.FunctionList, GXDLMSFunctionSpec{Name: name, Objects: objects})
		g.ActivationStatus = append(g.ActivationStatus, *types.NewGXKeyValuePair(name, true))
	case 3:
		name, err := decodeOctetStringName(e.Parameters)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return nil, err
		}
		g.removeFunction(name)
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil, nil
}

func (g *GXDLMSFunctionControl) removeFunction(name string) {
	filteredF := g.FunctionList[:0]
	for _, it := range g.FunctionList {
		if it.Name != name {
			filteredF = append(filteredF, it)
		}
	}
	g.FunctionList = filteredF
	filteredA := g.ActivationStatus[:0]
	for _, it := range g.ActivationStatus {
		if it.Key != name {
			filteredA = append(filteredA, it)
		}
	}
	g.ActivationStatus = filteredA
}

func (g *GXDLMSFunctionControl) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	if all || g.CanRead(2) {
		attributes = append(attributes, 2)
	}
	if all || g.CanRead(3) {
		attributes = append(attributes, 3)
	}
	return attributes
}

func (g *GXDLMSFunctionControl) GetNames() []string {
	return []string{"Logical Name", "ActivationStatus", "FunctionList"}
}

func (g *GXDLMSFunctionControl) GetMethodNames() []string {
	return []string{"SetFunctionStatus", "AddFunction", "RemoveFunction"}
}

func (g *GXDLMSFunctionControl) GetAttributeCount() int { return 3 }
func (g *GXDLMSFunctionControl) GetMethodCount() int    { return 3 }

func (g *GXDLMSFunctionControl) GetDataType(index int) (enums.DataType, error) {
	switch index {
	case 1:
		return enums.DataTypeOctetString, nil
	case 2, 3:
		return enums.DataTypeArray, nil
	default:
		return enums.DataTypeNone, dlmserrors.ErrInvalidAttributeIndex
	}
}

func (g *GXDLMSFunctionControl) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	switch e.Index {
	case 1:
		return helpers.LogicalNameToBytes(g.LogicalName())
	case 2:
		return encodeFunctionStatus(g.ActivationStatus)
	case 3:
		return encodeFunctionList(settings, g.FunctionList)
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
		return nil, nil
	}
}

func (g *GXDLMSFunctionControl) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	switch e.Index {
	case 1:
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return err
		}
		return g.SetLogicalName(ln)
	case 2:
		v, err := decodeFunctionStatus(e.Value)
		if err != nil {
			return err
		}
		g.ActivationStatus = v
	case 3:
		v, err := decodeFunctionList(e.Value)
		if err != nil {
			return err
		}
		g.FunctionList = v
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil
}

func (g *GXDLMSFunctionControl) Load(reader *GXXmlReader) error {
	ok, err := reader.IsStartElementNamed("Activations", true)
	if err != nil {
		return err
	}
	if ok {
		g.ActivationStatus = g.ActivationStatus[:0]
		for {
			ok, err = reader.IsStartElementNamed("Item", true)
			if err != nil {
				return err
			}
			if !ok {
				break
			}
			name, err := reader.ReadElementContentAsString("Name", "")
			if err != nil {
				return err
			}
			status, err := reader.ReadElementContentAsInt("Status", 0)
			if err != nil {
				return err
			}
			g.ActivationStatus = append(g.ActivationStatus, *types.NewGXKeyValuePair(name, status != 0))
		}
		if err = reader.ReadEndElement("Activations"); err != nil {
			return err
		}
	}

	ok, err = reader.IsStartElementNamed("Functions", true)
	if err != nil {
		return err
	}
	if ok {
		g.FunctionList = g.FunctionList[:0]
		for {
			ok, err = reader.IsStartElementNamed("Item", true)
			if err != nil {
				return err
			}
			if !ok {
				break
			}
			name, err := reader.ReadElementContentAsString("Name", "")
			if err != nil {
				return err
			}
			spec := GXDLMSFunctionSpec{Name: name, Objects: make([]IGXDLMSBase, 0)}

			oo, err := reader.IsStartElementNamed("Objects", true)
			if err != nil {
				return err
			}
			if oo {
				for {
					oo, err = reader.IsStartElementNamed("Object", true)
					if err != nil {
						return err
					}
					if !oo {
						break
					}
					ot, err := reader.ReadElementContentAsInt("ObjectType", 0)
					if err != nil {
						return err
					}
					ln, err := reader.ReadElementContentAsString("LN", "")
					if err != nil {
						return err
					}
					obj, err := CreateObject(enums.ObjectType(ot), ln, 0)
					if err == nil && obj != nil {
						spec.Objects = append(spec.Objects, obj)
					}
					if err = reader.ReadEndElement("Object"); err != nil {
						return err
					}
				}
				if err = reader.ReadEndElement("Objects"); err != nil {
					return err
				}
			}
			g.FunctionList = append(g.FunctionList, spec)
		}
		if err = reader.ReadEndElement("Functions"); err != nil {
			return err
		}
	}
	return nil
}

func (g *GXDLMSFunctionControl) Save(writer *GXXmlWriter) error {
	if err := writer.WriteStartElement("Activations"); err != nil {
		return err
	}
	for _, it := range g.ActivationStatus {
		if err := writer.WriteStartElement("Item"); err != nil {
			return err
		}
		if err := writer.WriteElementString("Name", it.Key); err != nil {
			return err
		}
		if err := writer.WriteElementStringBool("Status", it.Value); err != nil {
			return err
		}
		if err := writer.WriteEndElement(); err != nil {
			return err
		}
	}
	if err := writer.WriteEndElement(); err != nil {
		return err
	}

	if err := writer.WriteStartElement("Functions"); err != nil {
		return err
	}
	for _, it := range g.FunctionList {
		if err := writer.WriteStartElement("Item"); err != nil {
			return err
		}
		if err := writer.WriteElementString("Name", it.Name); err != nil {
			return err
		}
		if err := writer.WriteStartElement("Objects"); err != nil {
			return err
		}
		for _, obj := range it.Objects {
			if obj == nil {
				continue
			}
			if err := writer.WriteStartElement("Object"); err != nil {
				return err
			}
			if err := writer.WriteElementString("ObjectType", int(obj.Base().ObjectType())); err != nil {
				return err
			}
			if err := writer.WriteElementString("LN", obj.Base().LogicalName()); err != nil {
				return err
			}
			if err := writer.WriteEndElement(); err != nil {
				return err
			}
		}
		if err := writer.WriteEndElement(); err != nil {
			return err
		}
		if err := writer.WriteEndElement(); err != nil {
			return err
		}
	}
	return writer.WriteEndElement()
}

func (g *GXDLMSFunctionControl) PostLoad(reader *GXXmlReader) error { return nil }

func (g *GXDLMSFunctionControl) GetValues() []any {
	return []any{g.LogicalName(), g.ActivationStatus, g.FunctionList}
}

// SetFunctionStatus sets activation status values for known function blocks.
func (g *GXDLMSFunctionControl) SetFunctionStatus(client IGXDLMSClient, functions []types.GXKeyValuePair[string, bool]) ([][]byte, error) {
	data, err := encodeFunctionStatus(functions)
	if err != nil {
		return nil, err
	}
	return client.Method(g, 1, data, enums.DataTypeArray)
}

// AddFunction adds or replaces a function definition.
func (g *GXDLMSFunctionControl) AddFunction(client IGXDLMSClient, name string, objects []IGXDLMSBase) ([][]byte, error) {
	bb := types.NewGXByteBuffer()
	if err := bb.SetUint8(uint8(enums.DataTypeStructure)); err != nil {
		return nil, err
	}
	if err := bb.SetUint8(2); err != nil {
		return nil, err
	}
	if err := internal.SetData(nil, bb, enums.DataTypeOctetString, []byte(name)); err != nil {
		return nil, err
	}
	if err := bb.SetUint8(uint8(enums.DataTypeArray)); err != nil {
		return nil, err
	}
	types.SetObjectCount(len(objects), bb)
	for _, it := range objects {
		if it == nil {
			continue
		}
		if err := bb.SetUint8(uint8(enums.DataTypeStructure)); err != nil {
			return nil, err
		}
		if err := bb.SetUint8(2); err != nil {
			return nil, err
		}
		if err := internal.SetData(nil, bb, enums.DataTypeUint16, uint16(it.Base().ObjectType())); err != nil {
			return nil, err
		}
		ln, err := helpers.LogicalNameToBytes(it.Base().LogicalName())
		if err != nil {
			return nil, err
		}
		if err = internal.SetData(nil, bb, enums.DataTypeOctetString, ln); err != nil {
			return nil, err
		}
	}
	return client.Method(g, 2, bb.Array(), enums.DataTypeStructure)
}

// RemoveFunction removes a function definition by name.
func (g *GXDLMSFunctionControl) RemoveFunction(client IGXDLMSClient, name string) ([][]byte, error) {
	bb := types.NewGXByteBuffer()
	if err := internal.SetData(nil, bb, enums.DataTypeOctetString, []byte(name)); err != nil {
		return nil, err
	}
	return client.Method(g, 3, bb.Array(), enums.DataTypeOctetString)
}

func encodeFunctionStatus(functions []types.GXKeyValuePair[string, bool]) ([]byte, error) {
	bb := types.NewGXByteBuffer()
	if err := bb.SetUint8(uint8(enums.DataTypeArray)); err != nil {
		return nil, err
	}
	types.SetObjectCount(len(functions), bb)
	for _, it := range functions {
		if err := bb.SetUint8(uint8(enums.DataTypeStructure)); err != nil {
			return nil, err
		}
		if err := bb.SetUint8(2); err != nil {
			return nil, err
		}
		if err := internal.SetData(nil, bb, enums.DataTypeOctetString, []byte(it.Key)); err != nil {
			return nil, err
		}
		if err := internal.SetData(nil, bb, enums.DataTypeBoolean, it.Value); err != nil {
			return nil, err
		}
	}
	return bb.Array(), nil
}

func encodeFunctionList(settings *settings.GXDLMSSettings, functions []GXDLMSFunctionSpec) ([]byte, error) {
	bb := types.NewGXByteBuffer()
	if err := bb.SetUint8(uint8(enums.DataTypeArray)); err != nil {
		return nil, err
	}
	types.SetObjectCount(len(functions), bb)
	for _, it := range functions {
		if err := bb.SetUint8(uint8(enums.DataTypeStructure)); err != nil {
			return nil, err
		}
		if err := bb.SetUint8(2); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeOctetString, []byte(it.Name)); err != nil {
			return nil, err
		}
		if err := bb.SetUint8(uint8(enums.DataTypeArray)); err != nil {
			return nil, err
		}
		types.SetObjectCount(len(it.Objects), bb)
		for _, obj := range it.Objects {
			if obj == nil {
				continue
			}
			if err := bb.SetUint8(uint8(enums.DataTypeStructure)); err != nil {
				return nil, err
			}
			if err := bb.SetUint8(2); err != nil {
				return nil, err
			}
			if err := internal.SetData(settings, bb, enums.DataTypeUint16, uint16(obj.Base().ObjectType())); err != nil {
				return nil, err
			}
			ln, err := helpers.LogicalNameToBytes(obj.Base().LogicalName())
			if err != nil {
				return nil, err
			}
			if err = internal.SetData(settings, bb, enums.DataTypeOctetString, ln); err != nil {
				return nil, err
			}
		}
	}
	return bb.Array(), nil
}

func decodeFunctionStatus(value any) ([]types.GXKeyValuePair[string, bool], error) {
	rows, ok := fcToAnySlice(value)
	if !ok {
		return nil, fmt.Errorf("invalid activation status: %T", value)
	}
	ret := make([]types.GXKeyValuePair[string, bool], 0, len(rows))
	for _, row := range rows {
		it, ok := fcToAnySlice(row)
		if !ok || len(it) < 2 {
			continue
		}
		name, err := decodeOctetStringName(it[0])
		if err != nil {
			return nil, err
		}
		status, err := toBool(it[1])
		if err != nil {
			return nil, err
		}
		ret = append(ret, *types.NewGXKeyValuePair(name, status))
	}
	return ret, nil
}

func decodeSingleFunctionSpec(value any) (string, []IGXDLMSBase, error) {
	s, ok := fcToAnySlice(value)
	if !ok || len(s) < 2 {
		return "", nil, fmt.Errorf("invalid function spec: %T", value)
	}
	name, err := decodeOctetStringName(s[0])
	if err != nil {
		return "", nil, err
	}
	objs, err := decodeFunctionObjects(s[1])
	return name, objs, err
}

func decodeFunctionList(value any) ([]GXDLMSFunctionSpec, error) {
	rows, ok := fcToAnySlice(value)
	if !ok {
		return nil, fmt.Errorf("invalid function list: %T", value)
	}
	ret := make([]GXDLMSFunctionSpec, 0, len(rows))
	for _, row := range rows {
		name, objs, err := decodeSingleFunctionSpec(row)
		if err != nil {
			return nil, err
		}
		ret = append(ret, GXDLMSFunctionSpec{Name: name, Objects: objs})
	}
	return ret, nil
}

func decodeFunctionObjects(value any) ([]IGXDLMSBase, error) {
	rows, ok := fcToAnySlice(value)
	if !ok {
		return nil, fmt.Errorf("invalid function object list: %T", value)
	}
	ret := make([]IGXDLMSBase, 0, len(rows))
	for _, row := range rows {
		it, ok := fcToAnySlice(row)
		if !ok || len(it) < 2 {
			continue
		}
		ot32, err := toUint32(it[0])
		if err != nil {
			return nil, err
		}
		ln, err := decodeOctetStringName(it[1])
		if err != nil {
			return nil, err
		}
		obj, err := CreateObject(enums.ObjectType(ot32), ln, 0)
		if err == nil && obj != nil {
			ret = append(ret, obj)
		}
	}
	return ret, nil
}

func decodeOctetStringName(value any) (string, error) {
	switch v := value.(type) {
	case string:
		return v, nil
	case []byte:
		return string(v), nil
	case types.GXBitString:
		return v.String(), nil
	default:
		return "", fmt.Errorf("invalid name type: %T", value)
	}
}

func fcToAnySlice(value any) ([]any, bool) {
	switch v := value.(type) {
	case []any:
		return v, true
	case types.GXArray:
		return []any(v), true
	case types.GXStructure:
		return []any(v), true
	default:
		return nil, false
	}
}

// NewGXDLMSFunctionControl creates a new Function control object.
func NewGXDLMSFunctionControl(ln string, sn int16) (*GXDLMSFunctionControl, error) {
	if err := ValidateLogicalName(ln); err != nil {
		return nil, err
	}
	return &GXDLMSFunctionControl{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeFunctionControl,
			logicalName: ln,
			ShortName:   sn,
		},
		ActivationStatus: make([]types.GXKeyValuePair[string, bool], 0),
		FunctionList:     make([]GXDLMSFunctionSpec, 0),
	}, nil
}
