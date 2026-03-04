package objects

import (
	"github.com/Gurux/gxcommon-go"
	"github.com/Gurux/gxdlms-go/dlmserrors"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// GXDLMSPrimeNbOfdmPlcMacNetworkAdministrationData stores PRIME MAC network administration data.
type GXDLMSPrimeNbOfdmPlcMacNetworkAdministrationData struct {
	GXDLMSObject
	MulticastEntries  []GXMacMulticastEntry
	SwitchTable       []int16
	DirectTable       []GXMacDirectTable
	AvailableSwitches []GXMacAvailableSwitch
	Communications    []GXMacPhyCommunication
}

func (g *GXDLMSPrimeNbOfdmPlcMacNetworkAdministrationData) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}
func (g *GXDLMSPrimeNbOfdmPlcMacNetworkAdministrationData) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	if e.Index == 1 {
		g.MulticastEntries = nil
		g.SwitchTable = nil
		g.DirectTable = nil
		g.AvailableSwitches = nil
		g.Communications = nil
	} else {
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil, nil
}
func (g *GXDLMSPrimeNbOfdmPlcMacNetworkAdministrationData) GetAttributeIndexToRead(all bool) []int {
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
func (g *GXDLMSPrimeNbOfdmPlcMacNetworkAdministrationData) GetNames() []string {
	return []string{"Logical Name", "MulticastEntries", "SwitchTable", "DirectTable", "AvailableSwitches", "Communications"}
}
func (g *GXDLMSPrimeNbOfdmPlcMacNetworkAdministrationData) GetMethodNames() []string {
	return []string{"Reset"}
}
func (g *GXDLMSPrimeNbOfdmPlcMacNetworkAdministrationData) GetAttributeCount() int { return 6 }
func (g *GXDLMSPrimeNbOfdmPlcMacNetworkAdministrationData) GetMethodCount() int    { return 1 }
func (g *GXDLMSPrimeNbOfdmPlcMacNetworkAdministrationData) GetDataType(index int) (enums.DataType, error) {
	if index == 1 {
		return enums.DataTypeOctetString, nil
	}
	if index >= 2 && index <= 6 {
		return enums.DataTypeArray, nil
	}
	return 0, dlmserrors.ErrInvalidAttributeIndex
}

func (g *GXDLMSPrimeNbOfdmPlcMacNetworkAdministrationData) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	switch e.Index {
	case 1:
		return helpers.LogicalNameToBytes(g.LogicalName())
	case 2:
		return g.getMulticastEntries(settings)
	case 3:
		return g.getSwitchTable(settings)
	case 4:
		return g.getDirectTable(settings)
	case 5:
		return g.getAvailableSwitches(settings)
	case 6:
		return g.getCommunications(settings)
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
		return nil, nil
	}
}

func (g *GXDLMSPrimeNbOfdmPlcMacNetworkAdministrationData) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
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
		g.MulticastEntries = parseMulticastEntries(e.Value)
	case 3:
		g.SwitchTable = parseSwitchTable(e.Value)
	case 4:
		g.DirectTable = parseDirectTable(e.Value)
	case 5:
		g.AvailableSwitches, err = parseAvailableSwitches(e.Value)
	case 6:
		g.Communications, err = parseCommunications(e.Value)
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return err
}

func (g *GXDLMSPrimeNbOfdmPlcMacNetworkAdministrationData) Load(reader *GXXmlReader) error {
	var err error
	g.MulticastEntries, err = loadMulticastEntries(reader)
	if err != nil {
		return err
	}
	g.SwitchTable, err = loadSwitchTable(reader)
	if err != nil {
		return err
	}
	g.DirectTable, err = loadDirectTable(reader)
	if err != nil {
		return err
	}
	g.AvailableSwitches, err = loadAvailableSwitches(reader)
	if err != nil {
		return err
	}
	g.Communications, err = loadCommunications(reader)
	return err
}

func (g *GXDLMSPrimeNbOfdmPlcMacNetworkAdministrationData) Save(writer *GXXmlWriter) error {
	if err := saveMulticastEntries(writer, g.MulticastEntries); err != nil {
		return err
	}
	if err := saveSwitchTable(writer, g.SwitchTable); err != nil {
		return err
	}
	if err := saveDirectTable(writer, g.DirectTable); err != nil {
		return err
	}
	if err := saveAvailableSwitches(writer, g.AvailableSwitches); err != nil {
		return err
	}
	return saveCommunications(writer, g.Communications)
}

func (g *GXDLMSPrimeNbOfdmPlcMacNetworkAdministrationData) PostLoad(reader *GXXmlReader) error {
	return nil
}
func (g *GXDLMSPrimeNbOfdmPlcMacNetworkAdministrationData) GetValues() []any {
	return []any{g.LogicalName(), g.MulticastEntries, g.SwitchTable, g.DirectTable, g.AvailableSwitches, g.Communications}
}
func (g *GXDLMSPrimeNbOfdmPlcMacNetworkAdministrationData) Reset(client IGXDLMSClient) ([][]byte, error) {
	return client.Method(g, 1, int8(0), enums.DataTypeInt8)
}

func (g *GXDLMSPrimeNbOfdmPlcMacNetworkAdministrationData) getMulticastEntries(settings *settings.GXDLMSSettings) ([]byte, error) {
	bb := types.NewGXByteBuffer()
	_ = bb.SetUint8(uint8(enums.DataTypeArray))
	_ = types.SetObjectCount(len(g.MulticastEntries), bb)
	for _, it := range g.MulticastEntries {
		_ = bb.SetUint8(uint8(enums.DataTypeStructure))
		_ = bb.SetUint8(2)
		if err := internal.SetData(settings, bb, enums.DataTypeInt8, it.Id); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeInt16, it.Members); err != nil {
			return nil, err
		}
	}
	return bb.Array(), nil
}
func (g *GXDLMSPrimeNbOfdmPlcMacNetworkAdministrationData) getSwitchTable(settings *settings.GXDLMSSettings) ([]byte, error) {
	bb := types.NewGXByteBuffer()
	_ = bb.SetUint8(uint8(enums.DataTypeArray))
	_ = types.SetObjectCount(len(g.SwitchTable), bb)
	for _, it := range g.SwitchTable {
		if err := internal.SetData(settings, bb, enums.DataTypeInt16, it); err != nil {
			return nil, err
		}
	}
	return bb.Array(), nil
}
func (g *GXDLMSPrimeNbOfdmPlcMacNetworkAdministrationData) getDirectTable(settings *settings.GXDLMSSettings) ([]byte, error) {
	bb := types.NewGXByteBuffer()
	_ = bb.SetUint8(uint8(enums.DataTypeArray))
	_ = types.SetObjectCount(len(g.DirectTable), bb)
	for _, it := range g.DirectTable {
		_ = bb.SetUint8(uint8(enums.DataTypeStructure))
		_ = bb.SetUint8(7)
		if err := internal.SetData(settings, bb, enums.DataTypeInt16, it.SourceSId); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeInt16, it.SourceLnId); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeInt16, it.SourceLcId); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeInt16, it.DestinationSId); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeInt16, it.DestinationLnId); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeInt16, it.DestinationLcId); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeOctetString, it.Did); err != nil {
			return nil, err
		}
	}
	return bb.Array(), nil
}
func (g *GXDLMSPrimeNbOfdmPlcMacNetworkAdministrationData) getAvailableSwitches(settings *settings.GXDLMSSettings) ([]byte, error) {
	bb := types.NewGXByteBuffer()
	_ = bb.SetUint8(uint8(enums.DataTypeArray))
	_ = types.SetObjectCount(len(g.AvailableSwitches), bb)
	for _, it := range g.AvailableSwitches {
		_ = bb.SetUint8(uint8(enums.DataTypeStructure))
		_ = bb.SetUint8(5)
		if err := internal.SetData(settings, bb, enums.DataTypeOctetString, it.Sna); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeInt16, it.LsId); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeInt8, it.Level); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeInt8, it.RxLevel); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeInt8, it.RxSnr); err != nil {
			return nil, err
		}
	}
	return bb.Array(), nil
}
func (g *GXDLMSPrimeNbOfdmPlcMacNetworkAdministrationData) getCommunications(settings *settings.GXDLMSSettings) ([]byte, error) {
	bb := types.NewGXByteBuffer()
	_ = bb.SetUint8(uint8(enums.DataTypeArray))
	_ = types.SetObjectCount(len(g.Communications), bb)
	for _, it := range g.Communications {
		_ = bb.SetUint8(uint8(enums.DataTypeStructure))
		_ = bb.SetUint8(9)
		if err := internal.SetData(settings, bb, enums.DataTypeOctetString, it.Eui); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeInt8, it.TxPower); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeInt8, it.TxCoding); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeInt8, it.RxCoding); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeInt8, it.RxLvl); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeInt8, it.Snr); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeInt8, it.TxPowerModified); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeInt8, it.TxCodingModified); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeInt8, it.RxCodingModified); err != nil {
			return nil, err
		}
	}
	return bb.Array(), nil
}

func parseMulticastEntries(value any) []GXMacMulticastEntry {
	ret := make([]GXMacMulticastEntry, 0)
	for _, tmp := range value.(types.GXArray) {
		it := tmp.(types.GXStructure)
		if len(it) < 2 {
			continue
		}
		id, err1 := toInt8Value(it[0])
		members, err2 := toInt16Value(it[1])
		if err1 == nil && err2 == nil {
			ret = append(ret, GXMacMulticastEntry{Id: id, Members: members})
		}
	}
	return ret
}
func parseSwitchTable(value any) []int16 {
	ret := make([]int16, 0)
	for _, it := range value.(types.GXArray) {
		v, err := toInt16Value(it)
		if err == nil {
			ret = append(ret, v)
		}
	}
	return ret
}
func parseDirectTable(value any) []GXMacDirectTable {
	ret := make([]GXMacDirectTable, 0)
	for _, tmp := range value.(types.GXArray) {
		it := tmp.(types.GXStructure)
		if len(it) < 7 {
			continue
		}
		v := GXMacDirectTable{}
		v.SourceSId, _ = toInt16Value(it[0])
		v.SourceLnId, _ = toInt16Value(it[1])
		v.SourceLcId, _ = toInt16Value(it[2])
		v.DestinationSId, _ = toInt16Value(it[3])
		v.DestinationLnId, _ = toInt16Value(it[4])
		v.DestinationLcId, _ = toInt16Value(it[5])
		if b, ok := it[6].([]byte); ok {
			v.Did = b
		}
		ret = append(ret, v)
	}
	return ret
}
func parseAvailableSwitches(value any) ([]GXMacAvailableSwitch, error) {
	ret := make([]GXMacAvailableSwitch, 0)
	for _, tmp := range value.(types.GXArray) {
		it := tmp.(types.GXStructure)
		if len(it) < 5 {
			return nil, gxcommon.ErrArgumentOutOfRange
		}
		v := GXMacAvailableSwitch{}
		if b, ok := it[0].([]byte); ok {
			v.Sna = b
		}
		v.LsId, _ = toInt16Value(it[1])
		v.Level, _ = toInt8Value(it[2])
		v.RxLevel, _ = toInt8Value(it[3])
		v.RxSnr, _ = toInt8Value(it[4])
		ret = append(ret, v)
	}
	return ret, nil
}

func parseCommunications(value any) ([]GXMacPhyCommunication, error) {
	ret := make([]GXMacPhyCommunication, 0)
	for _, tmp := range value.(types.GXArray) {
		it := tmp.(types.GXStructure)
		if len(it) < 9 {
			return nil, gxcommon.ErrArgumentOutOfRange
		}
		v := GXMacPhyCommunication{}
		if b, ok := it[0].([]byte); ok {
			v.Eui = b
		}
		v.TxPower, _ = toInt8Value(it[1])
		v.TxCoding, _ = toInt8Value(it[2])
		v.RxCoding, _ = toInt8Value(it[3])
		v.RxLvl, _ = toInt8Value(it[4])
		v.Snr, _ = toInt8Value(it[5])
		v.TxPowerModified, _ = toInt8Value(it[6])
		v.TxCodingModified, _ = toInt8Value(it[7])
		v.RxCodingModified, _ = toInt8Value(it[8])
		ret = append(ret, v)
	}
	return ret, nil
}

func loadMulticastEntries(reader *GXXmlReader) ([]GXMacMulticastEntry, error) {
	list := make([]GXMacMulticastEntry, 0)
	ok, err := reader.IsStartElementNamed("MulticastEntries", true)
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
		it := GXMacMulticastEntry{}
		it.Id, err = reader.ReadElementContentAsInt8("Id", 0)
		if err != nil {
			return nil, err
		}
		it.Members, err = reader.ReadElementContentAsInt16("Members", 0)
		if err != nil {
			return nil, err
		}
		list = append(list, it)
	}
	if err := reader.ReadEndElement("MulticastEntries"); err != nil {
		return nil, err
	}
	return list, nil
}
func loadSwitchTable(reader *GXXmlReader) ([]int16, error) {
	list := make([]int16, 0)
	ok, err := reader.IsStartElementNamed("SwitchTable", true)
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
		v, err := reader.ReadElementContentAsInt16("Item", 0)
		if err != nil {
			return nil, err
		}
		list = append(list, v)
	}
	if err := reader.ReadEndElement("SwitchTable"); err != nil {
		return nil, err
	}
	return list, nil
}
func loadDirectTable(reader *GXXmlReader) ([]GXMacDirectTable, error) {
	list := make([]GXMacDirectTable, 0)
	ok, err := reader.IsStartElementNamed("DirectTable", true)
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
		it := GXMacDirectTable{}
		it.SourceSId, err = reader.ReadElementContentAsInt16("SourceSId", 0)
		if err != nil {
			return nil, err
		}
		it.SourceLnId, err = reader.ReadElementContentAsInt16("SourceLnId", 0)
		if err != nil {
			return nil, err
		}
		it.SourceLcId, err = reader.ReadElementContentAsInt16("SourceLcId", 0)
		if err != nil {
			return nil, err
		}
		it.DestinationSId, err = reader.ReadElementContentAsInt16("DestinationSId", 0)
		if err != nil {
			return nil, err
		}
		it.DestinationLnId, err = reader.ReadElementContentAsInt16("DestinationLnId", 0)
		if err != nil {
			return nil, err
		}
		it.DestinationLcId, err = reader.ReadElementContentAsInt16("DestinationLcId", 0)
		if err != nil {
			return nil, err
		}
		s, err := reader.ReadElementContentAsString("Did", "")
		if err != nil {
			return nil, err
		}
		it.Did = types.HexToBytes(s)
		list = append(list, it)
	}
	if err := reader.ReadEndElement("DirectTable"); err != nil {
		return nil, err
	}
	return list, nil
}
func loadAvailableSwitches(reader *GXXmlReader) ([]GXMacAvailableSwitch, error) {
	list := make([]GXMacAvailableSwitch, 0)
	ok, err := reader.IsStartElementNamed("AvailableSwitches", true)
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
		it := GXMacAvailableSwitch{}
		s, err := reader.ReadElementContentAsString("Sna", "")
		if err != nil {
			return nil, err
		}
		it.Sna = types.HexToBytes(s)
		it.LsId, err = reader.ReadElementContentAsInt16("LsId", 0)
		if err != nil {
			return nil, err
		}
		it.Level, err = reader.ReadElementContentAsInt8("Level", 0)
		if err != nil {
			return nil, err
		}
		it.RxLevel, err = reader.ReadElementContentAsInt8("RxLevel", 0)
		if err != nil {
			return nil, err
		}
		it.RxSnr, err = reader.ReadElementContentAsInt8("RxSnr", 0)
		if err != nil {
			return nil, err
		}
		list = append(list, it)
	}
	if err := reader.ReadEndElement("AvailableSwitches"); err != nil {
		return nil, err
	}
	return list, nil
}
func loadCommunications(reader *GXXmlReader) ([]GXMacPhyCommunication, error) {
	list := make([]GXMacPhyCommunication, 0)
	ok, err := reader.IsStartElementNamed("Communications", true)
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
		it := GXMacPhyCommunication{}
		s, err := reader.ReadElementContentAsString("Eui", "")
		if err != nil {
			return nil, err
		}
		it.Eui = types.HexToBytes(s)
		it.TxPower, err = reader.ReadElementContentAsInt8("TxPower", 0)
		if err != nil {
			return nil, err
		}
		it.TxCoding, err = reader.ReadElementContentAsInt8("TxCoding", 0)
		if err != nil {
			return nil, err
		}
		it.RxCoding, err = reader.ReadElementContentAsInt8("RxCoding", 0)
		if err != nil {
			return nil, err
		}
		it.RxLvl, err = reader.ReadElementContentAsInt8("RxLvl", 0)
		if err != nil {
			return nil, err
		}
		it.Snr, err = reader.ReadElementContentAsInt8("Snr", 0)
		if err != nil {
			return nil, err
		}
		it.TxPowerModified, err = reader.ReadElementContentAsInt8("TxPowerModified", 0)
		if err != nil {
			return nil, err
		}
		it.TxCodingModified, err = reader.ReadElementContentAsInt8("TxCodingModified", 0)
		if err != nil {
			return nil, err
		}
		it.RxCodingModified, err = reader.ReadElementContentAsInt8("RxCodingModified", 0)
		if err != nil {
			return nil, err
		}
		list = append(list, it)
	}
	if err := reader.ReadEndElement("Communications"); err != nil {
		return nil, err
	}
	return list, nil
}

func saveMulticastEntries(writer *GXXmlWriter, list []GXMacMulticastEntry) error {
	writer.WriteStartElement("MulticastEntries")
	for _, it := range list {
		writer.WriteStartElement("Item")
		if err := writer.WriteElementString("Id", it.Id); err != nil {
			return err
		}
		if err := writer.WriteElementString("Members", it.Members); err != nil {
			return err
		}
		writer.WriteEndElement()
	}
	writer.WriteEndElement()
	return nil
}
func saveSwitchTable(writer *GXXmlWriter, list []int16) error {
	writer.WriteStartElement("SwitchTable")
	for _, it := range list {
		if err := writer.WriteElementString("Item", it); err != nil {
			return err
		}
	}
	writer.WriteEndElement()
	return nil
}
func saveDirectTable(writer *GXXmlWriter, list []GXMacDirectTable) error {
	writer.WriteStartElement("DirectTable")
	for _, it := range list {
		writer.WriteStartElement("Item")
		if err := writer.WriteElementString("SourceSId", it.SourceSId); err != nil {
			return err
		}
		if err := writer.WriteElementString("SourceLnId", it.SourceLnId); err != nil {
			return err
		}
		if err := writer.WriteElementString("SourceLcId", it.SourceLcId); err != nil {
			return err
		}
		if err := writer.WriteElementString("DestinationSId", it.DestinationSId); err != nil {
			return err
		}
		if err := writer.WriteElementString("DestinationLnId", it.DestinationLnId); err != nil {
			return err
		}
		if err := writer.WriteElementString("DestinationLcId", it.DestinationLcId); err != nil {
			return err
		}
		if err := writer.WriteElementString("Did", types.ToHex(it.Did, false)); err != nil {
			return err
		}
		writer.WriteEndElement()
	}
	writer.WriteEndElement()
	return nil
}
func saveAvailableSwitches(writer *GXXmlWriter, list []GXMacAvailableSwitch) error {
	writer.WriteStartElement("AvailableSwitches")
	for _, it := range list {
		writer.WriteStartElement("Item")
		if err := writer.WriteElementString("Sna", types.ToHex(it.Sna, false)); err != nil {
			return err
		}
		if err := writer.WriteElementString("LsId", it.LsId); err != nil {
			return err
		}
		if err := writer.WriteElementString("Level", it.Level); err != nil {
			return err
		}
		if err := writer.WriteElementString("RxLevel", it.RxLevel); err != nil {
			return err
		}
		if err := writer.WriteElementString("RxSnr", it.RxSnr); err != nil {
			return err
		}
		writer.WriteEndElement()
	}
	writer.WriteEndElement()
	return nil
}
func saveCommunications(writer *GXXmlWriter, list []GXMacPhyCommunication) error {
	writer.WriteStartElement("Communications")
	for _, it := range list {
		writer.WriteStartElement("Item")
		if err := writer.WriteElementString("Eui", types.ToHex(it.Eui, false)); err != nil {
			return err
		}
		if err := writer.WriteElementString("TxPower", it.TxPower); err != nil {
			return err
		}
		if err := writer.WriteElementString("TxCoding", it.TxCoding); err != nil {
			return err
		}
		if err := writer.WriteElementString("RxCoding", it.RxCoding); err != nil {
			return err
		}
		if err := writer.WriteElementString("RxLvl", it.RxLvl); err != nil {
			return err
		}
		if err := writer.WriteElementString("Snr", it.Snr); err != nil {
			return err
		}
		if err := writer.WriteElementString("TxPowerModified", it.TxPowerModified); err != nil {
			return err
		}
		if err := writer.WriteElementString("TxCodingModified", it.TxCodingModified); err != nil {
			return err
		}
		if err := writer.WriteElementString("RxCodingModified", it.RxCodingModified); err != nil {
			return err
		}
		writer.WriteEndElement()
	}
	writer.WriteEndElement()
	return nil
}

func NewGXDLMSPrimeNbOfdmPlcMacNetworkAdministrationData(ln string, sn int16) (*GXDLMSPrimeNbOfdmPlcMacNetworkAdministrationData, error) {
	if err := ValidateLogicalName(ln); err != nil {
		return nil, err
	}
	return &GXDLMSPrimeNbOfdmPlcMacNetworkAdministrationData{GXDLMSObject: GXDLMSObject{objectType: enums.ObjectTypePrimeNbOfdmPlcMacNetworkAdministrationData, logicalName: ln, ShortName: sn}}, nil
}
