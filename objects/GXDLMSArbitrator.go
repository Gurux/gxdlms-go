package objects

import (
	"strings"

	"github.com/Gurux/gxcommon-go"
	"github.com/Gurux/gxdlms-go/dlmserrors"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// GXDLMSArbitrator represents DLMS arbitrator object.
type GXDLMSArbitrator struct {
	GXDLMSObject
	Actions                 []GXDLMSActionItem
	PermissionsTable        []string
	WeightingsTable         [][]uint16
	MostRecentRequestsTable []string
	LastOutcome             uint8
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSArbitrator) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

func (g *GXDLMSArbitrator) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	switch e.Index {
	case 1:
		args := e.Parameters.(types.GXStructure)
		if len(args) > 0 {
			g.LastOutcome = args[0].(uint8)
		}
	case 2:
		for i := range g.PermissionsTable {
			g.PermissionsTable[i] = ""
		}
		for i := range g.WeightingsTable {
			g.PermissionsTable[i] = strings.Repeat("0", len(g.PermissionsTable[i]))
		}
		for i := range g.MostRecentRequestsTable {
			g.MostRecentRequestsTable[i] = ""
		}
		g.LastOutcome = 0
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil, nil
}
func (g *GXDLMSArbitrator) GetAttributeIndexToRead(all bool) []int {
	var a []int
	if all || g.LogicalName() == "" {
		a = append(a, 1)
	}
	for i := 2; i <= 6; i++ {
		if all || g.CanRead(i) {
			a = append(a, i)
		}
	}
	return a
}
func (g *GXDLMSArbitrator) GetNames() []string {
	return []string{"Logical Name", "Actions", "Permissions table", "Weightings table", "Most recent requests table", "Last outcome"}
}
func (g *GXDLMSArbitrator) GetMethodNames() []string { return []string{"Request Action", "Reset"} }
func (g *GXDLMSArbitrator) GetAttributeCount() int   { return 6 }
func (g *GXDLMSArbitrator) GetMethodCount() int      { return 2 }
func (g *GXDLMSArbitrator) GetDataType(index int) (enums.DataType, error) {
	switch index {
	case 1:
		return enums.DataTypeOctetString, nil
	case 2, 3, 4, 5:
		return enums.DataTypeArray, nil
	case 6:
		return enums.DataTypeUint8, nil
	default:
		return 0, dlmserrors.ErrInvalidAttributeIndex
	}
}
func (g *GXDLMSArbitrator) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	switch e.Index {
	case 1:
		return helpers.LogicalNameToBytes(g.LogicalName())
	case 2:
		return g.encodeActions(settings)
	case 3:
		return g.encodeBitStrings(settings, g.PermissionsTable)
	case 4:
		return g.encodeWeights(settings)
	case 5:
		return g.encodeBitStrings(settings, g.MostRecentRequestsTable)
	case 6:
		return g.LastOutcome, nil
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
		return nil, nil
	}
}
func (g *GXDLMSArbitrator) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	var err error
	switch e.Index {
	case 1:
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return err
		}
		return g.SetLogicalName(ln)
	case 2:
		g.Actions, err = parseActions(e.Value)
	case 3:
		g.PermissionsTable, err = parseStringArray(e.Value)
	case 4:
		g.WeightingsTable, err = parseWeights(e.Value)
	case 5:
		g.MostRecentRequestsTable, err = parseStringArray(e.Value)
	case 6:
		g.LastOutcome = e.Value.(uint8)
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return err
}

func (g *GXDLMSArbitrator) Load(reader *GXXmlReader) error {
	var err error
	g.Actions, err = loadArbActions(reader)
	if err != nil {
		return err
	}
	g.PermissionsTable, err = loadStringItems(reader, "PermissionTable")
	if err != nil {
		return err
	}
	g.WeightingsTable, err = loadWeights(reader)
	if err != nil {
		return err
	}
	g.MostRecentRequestsTable, err = loadStringItems(reader, "MostRecentRequestsTable")
	if err != nil {
		return err
	}
	g.LastOutcome, err = reader.ReadElementContentAsUInt8("LastOutcome", 0)
	return err
}
func (g *GXDLMSArbitrator) Save(writer *GXXmlWriter) error {
	if err := saveArbActions(writer, g.Actions); err != nil {
		return err
	}
	if err := saveStringItems(writer, "PermissionTable", g.PermissionsTable); err != nil {
		return err
	}
	if err := saveWeights(writer, g.WeightingsTable); err != nil {
		return err
	}
	if err := saveStringItems(writer, "MostRecentRequestsTable", g.MostRecentRequestsTable); err != nil {
		return err
	}
	return writer.WriteElementString("LastOutcome", g.LastOutcome)
}
func (g *GXDLMSArbitrator) PostLoad(reader *GXXmlReader) error { return nil }
func (g *GXDLMSArbitrator) GetValues() []any {
	return []any{g.LogicalName(), g.Actions, g.PermissionsTable, g.WeightingsTable, g.MostRecentRequestsTable, g.LastOutcome}
}

func (g *GXDLMSArbitrator) RequestAction(client IGXDLMSClient, actor uint8, actions string) ([][]byte, error) {
	bb := types.NewGXByteBuffer()
	_ = bb.SetUint8(uint8(enums.DataTypeStructure))
	_ = bb.SetUint8(2)
	if err := internal.SetData(nil, bb, enums.DataTypeUint8, actor); err != nil {
		return nil, err
	}
	if err := internal.SetData(nil, bb, enums.DataTypeBitString, actions); err != nil {
		return nil, err
	}
	return client.Method(g, 1, bb.Array(), enums.DataTypeStructure)
}
func (g *GXDLMSArbitrator) Reset(client IGXDLMSClient) ([][]byte, error) {
	return client.Method(g, 2, int8(0), enums.DataTypeInt8)
}

func (g *GXDLMSArbitrator) encodeActions(settings *settings.GXDLMSSettings) ([]byte, error) {
	data := types.NewGXByteBuffer()
	_ = data.SetUint8(uint8(enums.DataTypeArray))
	_ = types.SetObjectCount(len(g.Actions), data)
	for _, it := range g.Actions {
		_ = data.SetUint8(uint8(enums.DataTypeStructure))
		_ = data.SetUint8(2)
		ln, err := helpers.LogicalNameToBytes(it.LogicalName)
		if err != nil {
			return nil, err
		}
		if err = internal.SetData(settings, data, enums.DataTypeOctetString, ln); err != nil {
			return nil, err
		}
		if err = internal.SetData(settings, data, enums.DataTypeUint16, it.ScriptSelector); err != nil {
			return nil, err
		}
	}
	return data.Array(), nil
}
func (g *GXDLMSArbitrator) encodeBitStrings(settings *settings.GXDLMSSettings, arr []string) ([]byte, error) {
	data := types.NewGXByteBuffer()
	_ = data.SetUint8(uint8(enums.DataTypeArray))
	_ = types.SetObjectCount(len(arr), data)
	for _, it := range arr {
		if err := internal.SetData(settings, data, enums.DataTypeBitString, it); err != nil {
			return nil, err
		}
	}
	return data.Array(), nil
}
func (g *GXDLMSArbitrator) encodeWeights(settings *settings.GXDLMSSettings) ([]byte, error) {
	data := types.NewGXByteBuffer()
	_ = data.SetUint8(uint8(enums.DataTypeArray))
	_ = types.SetObjectCount(len(g.WeightingsTable), data)
	for _, row := range g.WeightingsTable {
		_ = data.SetUint8(uint8(enums.DataTypeArray))
		_ = types.SetObjectCount(len(row), data)
		for _, it := range row {
			if err := internal.SetData(settings, data, enums.DataTypeUint16, it); err != nil {
				return nil, err
			}
		}
	}
	return data.Array(), nil
}
func parseActions(value any) ([]GXDLMSActionItem, error) {
	list := make([]GXDLMSActionItem, 0)
	for _, row := range value.(types.GXArray) {
		it := row.(types.GXStructure)
		if len(it) < 2 {
			return nil, gxcommon.ErrArgumentOutOfRange
		}
		ln, _ := helpers.ToLogicalName(it[0])
		ss, err := toUint32(it[1])
		if err != nil {
			continue
		}
		list = append(list, GXDLMSActionItem{LogicalName: ln, ScriptSelector: uint16(ss)})
	}
	return list, nil
}
func parseStringArray(value any) ([]string, error) {
	ret := make([]string, 0)
	for _, it := range value.(types.GXArray) {
		tmp := it.(types.GXBitString)
		ret = append(ret, tmp.String())
	}
	return ret, nil
}
func parseWeights(value any) ([][]uint16, error) {
	var val uint16
	var err error
	ret := make([][]uint16, 0)
	for _, row := range value.(types.GXArray) {
		tmp := make([]uint16, 0)
		for _, it := range row.(types.GXArray) {
			val, err = toUint16(it)
			if err != nil {
				return nil, err
			}
			tmp = append(tmp, val)
		}
		ret = append(ret, tmp)
	}
	return ret, nil
}

func loadArbActions(reader *GXXmlReader) ([]GXDLMSActionItem, error) {
	list := make([]GXDLMSActionItem, 0)
	ok, err := reader.IsStartElementNamed("Actions", true)
	if err != nil {
		return nil, err
	}
	if !ok {
		return list, nil
	}
	for {
		ok, err = reader.IsStartElementNamed("Item", true)
		if err != nil {
			return nil, err
		}
		if !ok {
			break
		}
		ln, err := reader.ReadElementContentAsString("LN", "")
		if err != nil {
			return nil, err
		}
		ss, err := reader.ReadElementContentAsUInt16("ScriptSelector", 0)
		if err != nil {
			return nil, err
		}
		list = append(list, GXDLMSActionItem{LogicalName: ln, ScriptSelector: ss})
	}
	if err := reader.ReadEndElement("Actions"); err != nil {
		return nil, err
	}
	return list, nil
}
func loadStringItems(reader *GXXmlReader, name string) ([]string, error) {
	list := make([]string, 0)
	ok, err := reader.IsStartElementNamed(name, true)
	if err != nil {
		return nil, err
	}
	if !ok {
		return list, nil
	}
	for {
		ok, err = reader.IsStartElementNamed("Item", false)
		if err != nil {
			return nil, err
		}
		if !ok {
			break
		}
		v, err := reader.ReadElementContentAsString("Item", "")
		if err != nil {
			return nil, err
		}
		list = append(list, v)
	}
	if err := reader.ReadEndElement(name); err != nil {
		return nil, err
	}
	return list, nil
}
func loadWeights(reader *GXXmlReader) ([][]uint16, error) {
	list := make([][]uint16, 0)
	ok, err := reader.IsStartElementNamed("WeightingTable", true)
	if err != nil {
		return nil, err
	}
	if !ok {
		return list, nil
	}
	for {
		ok, err = reader.IsStartElementNamed("Weightings", true)
		if err != nil {
			return nil, err
		}
		if !ok {
			break
		}
		row := make([]uint16, 0)
		for {
			ok, err = reader.IsStartElementNamed("Item", false)
			if err != nil {
				return nil, err
			}
			if !ok {
				break
			}
			v, err := reader.ReadElementContentAsUInt16("Item", 0)
			if err != nil {
				return nil, err
			}
			row = append(row, v)
		}
		list = append(list, row)
	}
	if err := reader.ReadEndElement("WeightingTable"); err != nil {
		return nil, err
	}
	return list, nil
}
func saveArbActions(writer *GXXmlWriter, list []GXDLMSActionItem) error {
	if len(list) == 0 {
		return nil
	}
	writer.WriteStartElement("Actions")
	for _, it := range list {
		writer.WriteStartElement("Item")
		if err := writer.WriteElementString("LN", it.LogicalName); err != nil {
			return err
		}
		if err := writer.WriteElementString("ScriptSelector", it.ScriptSelector); err != nil {
			return err
		}
		writer.WriteEndElement()
	}
	writer.WriteEndElement()
	return nil
}
func saveStringItems(writer *GXXmlWriter, name string, list []string) error {
	if len(list) == 0 {
		return nil
	}
	writer.WriteStartElement(name)
	for _, it := range list {
		if err := writer.WriteElementString("Item", it); err != nil {
			return err
		}
	}
	writer.WriteEndElement()
	return nil
}
func saveWeights(writer *GXXmlWriter, list [][]uint16) error {
	if len(list) == 0 {
		return nil
	}
	writer.WriteStartElement("WeightingTable")
	for _, row := range list {
		writer.WriteStartElement("Weightings")
		for _, it := range row {
			if err := writer.WriteElementString("Item", it); err != nil {
				return err
			}
		}
		writer.WriteEndElement()
	}
	writer.WriteEndElement()
	return nil
}

// NewGXDLMSArbitrator creates a new arbitrator object instance.
//
// The function validates `ln` before creating the object.
// `ln` is the Logical Name and `sn` is the Short Name of the object.
func NewGXDLMSArbitrator(ln string, sn int16) (*GXDLMSArbitrator, error) {
	if err := ValidateLogicalName(ln); err != nil {
		return nil, err
	}
	return &GXDLMSArbitrator{
			GXDLMSObject: GXDLMSObject{
				objectType:  enums.ObjectTypeArbitrator,
				logicalName: ln,
				ShortName:   sn}},
		nil
}
