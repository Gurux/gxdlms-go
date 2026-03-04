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

// GXDLMSNeighbourTable holds one entry from MAC neighbour table.
type GXDLMSNeighbourTable struct {
	ShortAddress       uint16
	Enabled            bool
	ToneMap            string
	Modulation         enums.Modulation
	TxGain             int8
	TxRes              enums.GainResolution
	TxCoeff            string
	Lqi                uint8
	PhaseDifferential  int8
	TMRValidTime       uint8
	NeighbourValidTime uint8
}

// GXDLMSMacPosTable holds one entry from MAC POS table.
type GXDLMSMacPosTable struct {
	ShortAddress uint16
	LQI          uint8
	ValidTime    uint8
}

// GXDLMSG3PlcMacSetup models G3-PLC MAC setup.
type GXDLMSG3PlcMacSetup struct {
	GXDLMSObject
	ShortAddress                 uint16
	RcCoord                      uint16
	PANId                        uint16
	KeyTable                     []types.GXKeyValuePair[byte, []byte]
	FrameCounter                 uint32
	ToneMask                     string
	TmrTtl                       uint8
	MaxFrameRetries              uint8
	NeighbourTableEntryTtl       uint8
	NeighbourTable               []GXDLMSNeighbourTable
	HighPriorityWindowSize       uint8
	CscmFairnessLimit            uint8
	BeaconRandomizationWindowLen uint8
	A                            uint8
	K                            uint8
	MinCwAttempts                uint8
	CenelecLegacyMode            uint8
	FccLegacyMode                uint8
	MaxBe                        uint8
	MaxCsmaBackoffs              uint8
	MinBe                        uint8
	MacBroadcastMaxCwEnabled     bool
	MacTransmitAtten             uint8
	MacPosTable                  []GXDLMSMacPosTable
	MacDuplicateDetectionTtl     uint8
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSG3PlcMacSetup) Base() *GXDLMSObject { return &g.GXDLMSObject }

func (g *GXDLMSG3PlcMacSetup) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	switch e.Index {
	case 1:
		address, err := toUint32(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return nil, err
		}
		filtered := make([]GXDLMSNeighbourTable, 0)
		for _, it := range g.NeighbourTable {
			if it.ShortAddress == uint16(address) {
				filtered = append(filtered, it)
			}
		}
		return getNeighbourTables(settings, filtered)
	case 2:
		address, err := toUint32(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return nil, err
		}
		filtered := make([]GXDLMSMacPosTable, 0)
		for _, it := range g.MacPosTable {
			if it.ShortAddress == uint16(address) {
				filtered = append(filtered, it)
			}
		}
		return getPosTables(settings, filtered)
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
		return nil, nil
	}
}

func (g *GXDLMSG3PlcMacSetup) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	for i := 2; i <= 25; i++ {
		if all || g.CanRead(i) {
			attributes = append(attributes, i)
		}
	}
	if g.Version > 2 && (all || g.CanRead(26)) {
		attributes = append(attributes, 26)
	}
	return attributes
}

func (g *GXDLMSG3PlcMacSetup) GetNames() []string {
	return []string{
		"Logical Name", "MacShortAddress", "MacRcCoord", "MacPANId", "MacKeyTable", "MacFrameCounter",
		"MacToneMask", "MacTmrTtl", "MacMaxFrameRetries", "MacNeighbourTableEntryTtl", "MacNeighbourTable",
		"MacHighPriorityWindowSize", "MacCscmFairnessLimit", "MacBeaconRandomizationWindowLength", "MacA",
		"MacK", "MacMinCwAttempts", "MacCenelecLegacyMode", "MacFCCLegacyMode", "MacMaxBe", "MacMaxCsmaBackoffs",
		"MacMinBe", "MacBroadcastMaxCwEnabled", "MacTransmitAtten", "MacPosTable", "MacDuplicateDetectionTtl",
	}
}

func (g *GXDLMSG3PlcMacSetup) GetMethodNames() []string {
	return []string{"MAC get neighbour table entry", "MAC get POS tableentry"}
}

func (g *GXDLMSG3PlcMacSetup) GetAttributeCount() int {
	if g.Version == 3 {
		return 26
	}
	if g.Version == 2 {
		return 25
	}
	return 22
}

func (g *GXDLMSG3PlcMacSetup) GetMethodCount() int {
	if g.Version == 3 {
		return 2
	}
	return 1
}

func (g *GXDLMSG3PlcMacSetup) GetDataType(index int) (enums.DataType, error) {
	switch index {
	case 1:
		return enums.DataTypeOctetString, nil
	case 2, 3, 4:
		return enums.DataTypeUint16, nil
	case 5:
		return enums.DataTypeArray, nil
	case 6:
		return enums.DataTypeUint32, nil
	case 7:
		return enums.DataTypeBitString, nil
	case 8, 9, 10:
		return enums.DataTypeUint8, nil
	case 11:
		return enums.DataTypeArray, nil
	case 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22:
		return enums.DataTypeUint8, nil
	case 23:
		return enums.DataTypeBoolean, nil
	case 24:
		return enums.DataTypeUint8, nil
	case 25:
		return enums.DataTypeArray, nil
	case 26:
		return enums.DataTypeUint8, nil
	default:
		return enums.DataTypeNone, dlmserrors.ErrInvalidAttributeIndex
	}
}

func (g *GXDLMSG3PlcMacSetup) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	switch e.Index {
	case 1:
		return helpers.LogicalNameToBytes(g.LogicalName())
	case 2:
		return g.ShortAddress, nil
	case 3:
		return g.RcCoord, nil
	case 4:
		return g.PANId, nil
	case 5:
		return encodeMacKeyTable(settings, g.KeyTable)
	case 6:
		return g.FrameCounter, nil
	case 7:
		return g.ToneMask, nil
	case 8:
		return g.TmrTtl, nil
	case 9:
		return g.MaxFrameRetries, nil
	case 10:
		return g.NeighbourTableEntryTtl, nil
	case 11:
		return getNeighbourTables(settings, g.NeighbourTable)
	case 12:
		return g.HighPriorityWindowSize, nil
	case 13:
		return g.CscmFairnessLimit, nil
	case 14:
		return g.BeaconRandomizationWindowLen, nil
	case 15:
		return g.A, nil
	case 16:
		return g.K, nil
	case 17:
		return g.MinCwAttempts, nil
	case 18:
		return g.CenelecLegacyMode, nil
	case 19:
		return g.FccLegacyMode, nil
	case 20:
		return g.MaxBe, nil
	case 21:
		return g.MaxCsmaBackoffs, nil
	case 22:
		return g.MinBe, nil
	case 23:
		return g.MacBroadcastMaxCwEnabled, nil
	case 24:
		return g.MacTransmitAtten, nil
	case 25:
		return getPosTables(settings, g.MacPosTable)
	case 26:
		return g.MacDuplicateDetectionTtl, nil
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
		return nil, nil
	}
}

func (g *GXDLMSG3PlcMacSetup) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	switch e.Index {
	case 1:
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return err
		}
		return g.SetLogicalName(ln)
	case 2:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.ShortAddress = uint16(v)
	case 3:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.RcCoord = uint16(v)
	case 4:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.PANId = uint16(v)
	case 5:
		keys, err := decodeMacKeyTable(e.Value)
		if err != nil {
			return err
		}
		g.KeyTable = keys
	case 6:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.FrameCounter = v
	case 7:
		g.ToneMask = toBitStringValue(e.Value)
	case 8:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.TmrTtl = v
	case 9:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.MaxFrameRetries = v
	case 10:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.NeighbourTableEntryTtl = v
	case 11:
		table, err := parseNeighbourTableEntry(e.Value)
		if err != nil {
			return err
		}
		g.NeighbourTable = table
	case 12:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.HighPriorityWindowSize = v
	case 13:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.CscmFairnessLimit = v
	case 14:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.BeaconRandomizationWindowLen = v
	case 15:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.A = v
	case 16:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.K = v
	case 17:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.MinCwAttempts = v
	case 18:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.CenelecLegacyMode = v
	case 19:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.FccLegacyMode = v
	case 20:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.MaxBe = v
	case 21:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.MaxCsmaBackoffs = v
	case 22:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.MinBe = v
	case 23:
		v, err := toBool(e.Value)
		if err != nil {
			return err
		}
		g.MacBroadcastMaxCwEnabled = v
	case 24:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.MacTransmitAtten = v
	case 25:
		table, err := parsePosTableEntry(e.Value)
		if err != nil {
			return err
		}
		g.MacPosTable = table
	case 26:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.MacDuplicateDetectionTtl = v
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil
}

func (g *GXDLMSG3PlcMacSetup) Load(reader *GXXmlReader) error {
	var err error
	g.ShortAddress, err = reader.ReadElementContentAsUInt16("ShortAddress", 0)
	if err != nil {
		return err
	}
	g.RcCoord, err = reader.ReadElementContentAsUInt16("RcCoord", 0)
	if err != nil {
		return err
	}
	g.PANId, err = reader.ReadElementContentAsUInt16("PANId", 0)
	if err != nil {
		return err
	}
	if err = g.loadKeyTable(reader); err != nil {
		return err
	}
	g.FrameCounter, err = reader.ReadElementContentAsUInt32("FrameCounter", 0)
	if err != nil {
		return err
	}
	g.ToneMask, err = reader.ReadElementContentAsString("ToneMask", "")
	if err != nil {
		return err
	}
	g.TmrTtl, err = reader.ReadElementContentAsUInt8("TmrTtl", 0)
	if err != nil {
		return err
	}
	g.MaxFrameRetries, err = reader.ReadElementContentAsUInt8("MaxFrameRetries", 0)
	if err != nil {
		return err
	}
	g.NeighbourTableEntryTtl, err = reader.ReadElementContentAsUInt8("NeighbourTableEntryTtl", 0)
	if err != nil {
		return err
	}
	if err = g.loadNeighbourTable(reader); err != nil {
		return err
	}
	g.HighPriorityWindowSize, err = reader.ReadElementContentAsUInt8("HighPriorityWindowSize", 0)
	if err != nil {
		return err
	}
	g.CscmFairnessLimit, err = reader.ReadElementContentAsUInt8("CscmFairnessLimit", 0)
	if err != nil {
		return err
	}
	g.BeaconRandomizationWindowLen, err = reader.ReadElementContentAsUInt8("BeaconRandomizationWindowLength", 0)
	if err != nil {
		return err
	}
	g.A, err = reader.ReadElementContentAsUInt8("A", 0)
	if err != nil {
		return err
	}
	g.K, err = reader.ReadElementContentAsUInt8("K", 0)
	if err != nil {
		return err
	}
	g.MinCwAttempts, err = reader.ReadElementContentAsUInt8("MinCwAttempts", 0)
	if err != nil {
		return err
	}
	g.CenelecLegacyMode, err = reader.ReadElementContentAsUInt8("CenelecLegacyMode", 0)
	if err != nil {
		return err
	}
	g.FccLegacyMode, err = reader.ReadElementContentAsUInt8("FccLegacyMode", 0)
	if err != nil {
		return err
	}
	g.MaxBe, err = reader.ReadElementContentAsUInt8("MaxBe", 0)
	if err != nil {
		return err
	}
	g.MaxCsmaBackoffs, err = reader.ReadElementContentAsUInt8("MaxCsmaBackoffs", 0)
	if err != nil {
		return err
	}
	g.MinBe, err = reader.ReadElementContentAsUInt8("MinBe", 0)
	if err != nil {
		return err
	}
	mb, err := reader.ReadElementContentAsInt("MacBroadcastMaxCwEnabled", 0)
	if err != nil {
		return err
	}
	g.MacBroadcastMaxCwEnabled = mb != 0
	g.MacTransmitAtten, err = reader.ReadElementContentAsUInt8("MacTransmitAtten", 0)
	if err != nil {
		return err
	}
	if err = g.loadMacPosTable(reader); err != nil {
		return err
	}
	g.MacDuplicateDetectionTtl, err = reader.ReadElementContentAsUInt8("MacDuplicateDetectionTtl", 0)
	return err
}

func (g *GXDLMSG3PlcMacSetup) Save(writer *GXXmlWriter) error {
	if err := writer.WriteElementString("ShortAddress", g.ShortAddress); err != nil {
		return err
	}
	if err := writer.WriteElementString("RcCoord", g.RcCoord); err != nil {
		return err
	}
	if err := writer.WriteElementString("PANId", g.PANId); err != nil {
		return err
	}
	if err := g.saveKeyTable(writer); err != nil {
		return err
	}
	if err := writer.WriteElementString("FrameCounter", g.FrameCounter); err != nil {
		return err
	}
	if err := writer.WriteElementString("ToneMask", g.ToneMask); err != nil {
		return err
	}
	if err := writer.WriteElementString("TmrTtl", g.TmrTtl); err != nil {
		return err
	}
	if err := writer.WriteElementString("MaxFrameRetries", g.MaxFrameRetries); err != nil {
		return err
	}
	if err := writer.WriteElementString("NeighbourTableEntryTtl", g.NeighbourTableEntryTtl); err != nil {
		return err
	}
	if err := g.saveNeighbourTable(writer); err != nil {
		return err
	}
	if err := writer.WriteElementString("HighPriorityWindowSize", g.HighPriorityWindowSize); err != nil {
		return err
	}
	if err := writer.WriteElementString("CscmFairnessLimit", g.CscmFairnessLimit); err != nil {
		return err
	}
	if err := writer.WriteElementString("BeaconRandomizationWindowLength", g.BeaconRandomizationWindowLen); err != nil {
		return err
	}
	if err := writer.WriteElementString("A", g.A); err != nil {
		return err
	}
	if err := writer.WriteElementString("K", g.K); err != nil {
		return err
	}
	if err := writer.WriteElementString("MinCwAttempts", g.MinCwAttempts); err != nil {
		return err
	}
	if err := writer.WriteElementString("CenelecLegacyMode", g.CenelecLegacyMode); err != nil {
		return err
	}
	if err := writer.WriteElementString("FccLegacyMode", g.FccLegacyMode); err != nil {
		return err
	}
	if err := writer.WriteElementString("MaxBe", g.MaxBe); err != nil {
		return err
	}
	if err := writer.WriteElementString("MaxCsmaBackoffs", g.MaxCsmaBackoffs); err != nil {
		return err
	}
	if err := writer.WriteElementString("MinBe", g.MinBe); err != nil {
		return err
	}
	if err := writer.WriteElementStringBool("MacBroadcastMaxCwEnabled", g.MacBroadcastMaxCwEnabled); err != nil {
		return err
	}
	if err := writer.WriteElementString("MacTransmitAtten", g.MacTransmitAtten); err != nil {
		return err
	}
	if err := g.saveMacPosTable(writer); err != nil {
		return err
	}
	return writer.WriteElementString("MacDuplicateDetectionTtl", g.MacDuplicateDetectionTtl)
}

func (g *GXDLMSG3PlcMacSetup) PostLoad(reader *GXXmlReader) error { return nil }

func (g *GXDLMSG3PlcMacSetup) GetValues() []any {
	return []any{
		g.LogicalName(),
		g.ShortAddress,
		g.RcCoord,
		g.PANId,
		g.KeyTable,
		g.FrameCounter,
		g.ToneMask,
		g.TmrTtl,
		g.MaxFrameRetries,
		g.NeighbourTableEntryTtl,
		g.NeighbourTable,
		g.HighPriorityWindowSize,
		g.CscmFairnessLimit,
		g.BeaconRandomizationWindowLen,
		g.A,
		g.K,
		g.MinCwAttempts,
		g.CenelecLegacyMode,
		g.FccLegacyMode,
		g.MaxBe,
		g.MaxCsmaBackoffs,
		g.MinBe,
		g.MacBroadcastMaxCwEnabled,
		g.MacTransmitAtten,
		g.MacPosTable,
		g.MacDuplicateDetectionTtl,
	}
}

// GetNeighbourTableEntry requests a neighbour table entry by short address.
func (g *GXDLMSG3PlcMacSetup) GetNeighbourTableEntry(client IGXDLMSClient, address uint16) ([][]byte, error) {
	return client.Method(g, 1, address, enums.DataTypeUint16)
}

// ParseNeighbourTableEntry parses a raw method reply to neighbour table entries.
func (g *GXDLMSG3PlcMacSetup) ParseNeighbourTableEntry(reply *types.GXByteBuffer) ([]GXDLMSNeighbourTable, error) {
	info := internal.GXDataInfo{}
	value, err := internal.GetData(nil, reply, &info)
	if err != nil {
		return nil, err
	}
	return parseNeighbourTableEntry(value)
}

// GetPosTableEntry requests a POS table entry by short address.
func (g *GXDLMSG3PlcMacSetup) GetPosTableEntry(client IGXDLMSClient, address uint16) ([][]byte, error) {
	return client.Method(g, 2, address, enums.DataTypeUint16)
}

// ParsePosTableEntry parses a raw method reply to POS table entries.
func (g *GXDLMSG3PlcMacSetup) ParsePosTableEntry(reply *types.GXByteBuffer) ([]GXDLMSMacPosTable, error) {
	info := internal.GXDataInfo{}
	value, err := internal.GetData(nil, reply, &info)
	if err != nil {
		return nil, err
	}
	return parsePosTableEntry(value)
}

func encodeMacKeyTable(settings *settings.GXDLMSSettings, keyTable []types.GXKeyValuePair[byte, []byte]) ([]byte, error) {
	bb := types.NewGXByteBuffer()
	if err := bb.SetUint8(uint8(enums.DataTypeArray)); err != nil {
		return nil, err
	}
	types.SetObjectCount(len(keyTable), bb)
	for _, it := range keyTable {
		if err := bb.SetUint8(uint8(enums.DataTypeStructure)); err != nil {
			return nil, err
		}
		if err := bb.SetUint8(2); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeUint8, it.Key); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeOctetString, it.Value); err != nil {
			return nil, err
		}
	}
	return bb.Array(), nil
}

func decodeMacKeyTable(value any) ([]types.GXKeyValuePair[byte, []byte], error) {
	if value == nil {
		return make([]types.GXKeyValuePair[byte, []byte], 0), nil
	}
	rows, ok := macToAnySlice(value)
	if !ok {
		return nil, fmt.Errorf("invalid key table type: %T", value)
	}
	ret := make([]types.GXKeyValuePair[byte, []byte], 0, len(rows))
	for _, row := range rows {
		it, ok := macToAnySlice(row)
		if !ok || len(it) < 2 {
			continue
		}
		k, err := toUint8(it[0])
		if err != nil {
			return nil, err
		}
		data, _ := it[1].([]byte)
		ret = append(ret, *types.NewGXKeyValuePair(byte(k), data))
	}
	return ret, nil
}

func getNeighbourTables(settings *settings.GXDLMSSettings, tables []GXDLMSNeighbourTable) ([]byte, error) {
	bb := types.NewGXByteBuffer()
	if err := bb.SetUint8(uint8(enums.DataTypeArray)); err != nil {
		return nil, err
	}
	types.SetObjectCount(len(tables), bb)
	for _, it := range tables {
		if err := bb.SetUint8(uint8(enums.DataTypeStructure)); err != nil {
			return nil, err
		}
		if err := bb.SetUint8(11); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeUint16, it.ShortAddress); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeBoolean, it.Enabled); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeBitString, it.ToneMap); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeEnum, uint8(it.Modulation)); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeInt8, it.TxGain); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeEnum, uint8(it.TxRes)); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeBitString, it.TxCoeff); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeUint8, it.Lqi); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeInt8, it.PhaseDifferential); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeUint8, it.TMRValidTime); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeUint8, it.NeighbourValidTime); err != nil {
			return nil, err
		}
	}
	return bb.Array(), nil
}

func getPosTables(settings *settings.GXDLMSSettings, tables []GXDLMSMacPosTable) ([]byte, error) {
	bb := types.NewGXByteBuffer()
	if err := bb.SetUint8(uint8(enums.DataTypeArray)); err != nil {
		return nil, err
	}
	types.SetObjectCount(len(tables), bb)
	for _, it := range tables {
		if err := bb.SetUint8(uint8(enums.DataTypeStructure)); err != nil {
			return nil, err
		}
		if err := bb.SetUint8(3); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeUint16, it.ShortAddress); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeUint8, it.LQI); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeUint8, it.ValidTime); err != nil {
			return nil, err
		}
	}
	return bb.Array(), nil
}

func parseNeighbourTableEntry(value any) ([]GXDLMSNeighbourTable, error) {
	if value == nil {
		return make([]GXDLMSNeighbourTable, 0), nil
	}
	rows, ok := macToAnySlice(value)
	if !ok {
		return nil, fmt.Errorf("invalid neighbour table: %T", value)
	}
	ret := make([]GXDLMSNeighbourTable, 0, len(rows))
	for _, row := range rows {
		it, ok := macToAnySlice(row)
		if !ok || len(it) < 11 {
			continue
		}
		sa, err := toUint32(it[0])
		if err != nil {
			return nil, err
		}
		mod, err := toUint32(it[3])
		if err != nil {
			return nil, err
		}
		txr, err := toUint32(it[5])
		if err != nil {
			return nil, err
		}
		lqi, err := toUint8(it[7])
		if err != nil {
			return nil, err
		}
		tmr, err := toUint8(it[9])
		if err != nil {
			return nil, err
		}
		nvt, err := toUint8(it[10])
		if err != nil {
			return nil, err
		}
		en, err := toBool(it[1])
		if err != nil {
			return nil, err
		}
		txg, err := toInt8Value(it[4])
		if err != nil {
			return nil, err
		}
		pd, err := toInt8Value(it[8])
		if err != nil {
			return nil, err
		}
		ret = append(ret, GXDLMSNeighbourTable{
			ShortAddress:       uint16(sa),
			Enabled:            en,
			ToneMap:            toBitStringValue(it[2]),
			Modulation:         enums.Modulation(mod),
			TxGain:             txg,
			TxRes:              enums.GainResolution(txr),
			TxCoeff:            toBitStringValue(it[6]),
			Lqi:                lqi,
			PhaseDifferential:  pd,
			TMRValidTime:       tmr,
			NeighbourValidTime: nvt,
		})
	}
	return ret, nil
}

func parsePosTableEntry(value any) ([]GXDLMSMacPosTable, error) {
	if value == nil {
		return make([]GXDLMSMacPosTable, 0), nil
	}
	rows, ok := macToAnySlice(value)
	if !ok {
		return nil, fmt.Errorf("invalid POS table: %T", value)
	}
	ret := make([]GXDLMSMacPosTable, 0, len(rows))
	for _, row := range rows {
		it, ok := macToAnySlice(row)
		if !ok || len(it) < 3 {
			continue
		}
		sa, err := toUint32(it[0])
		if err != nil {
			return nil, err
		}
		lqi, err := toUint8(it[1])
		if err != nil {
			return nil, err
		}
		vt, err := toUint8(it[2])
		if err != nil {
			return nil, err
		}
		ret = append(ret, GXDLMSMacPosTable{ShortAddress: uint16(sa), LQI: lqi, ValidTime: vt})
	}
	return ret, nil
}

func toBitStringValue(value any) string {
	switch v := value.(type) {
	case nil:
		return ""
	case string:
		return v
	case []byte:
		return string(v)
	case types.GXBitString:
		return v.String()
	case *types.GXBitString:
		if v == nil {
			return ""
		}
		return v.String()
	default:
		return fmt.Sprint(v)
	}
}

func (g *GXDLMSG3PlcMacSetup) loadKeyTable(reader *GXXmlReader) error {
	g.KeyTable = make([]types.GXKeyValuePair[byte, []byte], 0)
	ok, err := reader.IsStartElementNamed("KeyTable", true)
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}
	for {
		ok, err = reader.IsStartElementNamed("Item", true)
		if err != nil {
			return err
		}
		if !ok {
			break
		}
		k, err := reader.ReadElementContentAsUInt8("Key", 0)
		if err != nil {
			return err
		}
		d, err := reader.ReadElementContentAsString("Data", "")
		if err != nil {
			return err
		}
		g.KeyTable = append(g.KeyTable, *types.NewGXKeyValuePair(byte(k), types.HexToBytes(d)))
	}
	return reader.ReadEndElement("KeyTable")
}

func (g *GXDLMSG3PlcMacSetup) loadNeighbourTable(reader *GXXmlReader) error {
	g.NeighbourTable = make([]GXDLMSNeighbourTable, 0)
	ok, err := reader.IsStartElementNamed("NeighbourTable", true)
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}
	for {
		ok, err = reader.IsStartElementNamed("Item", true)
		if err != nil {
			return err
		}
		if !ok {
			break
		}
		sa, err := reader.ReadElementContentAsUInt16("ShortAddress", 0)
		if err != nil {
			return err
		}
		en, err := reader.ReadElementContentAsInt("Enabled", 0)
		if err != nil {
			return err
		}
		tm, err := reader.ReadElementContentAsString("ToneMap", "")
		if err != nil {
			return err
		}
		mod, err := reader.ReadElementContentAsInt("Modulation", 0)
		if err != nil {
			return err
		}
		txg, err := reader.ReadElementContentAsInt8("TxGain", 0)
		if err != nil {
			return err
		}
		txr, err := reader.ReadElementContentAsInt("TxRes", 0)
		if err != nil {
			return err
		}
		txc, err := reader.ReadElementContentAsString("TxCoeff", "")
		if err != nil {
			return err
		}
		lqi, err := reader.ReadElementContentAsUInt8("Lqi", 0)
		if err != nil {
			return err
		}
		pd, err := reader.ReadElementContentAsInt8("PhaseDifferential", 0)
		if err != nil {
			return err
		}
		tmr, err := reader.ReadElementContentAsUInt8("TMRValidTime", 0)
		if err != nil {
			return err
		}
		nvt, err := reader.ReadElementContentAsUInt8("NeighbourValidTime", 0)
		if err != nil {
			return err
		}
		g.NeighbourTable = append(g.NeighbourTable, GXDLMSNeighbourTable{
			ShortAddress:       sa,
			Enabled:            en != 0,
			ToneMap:            tm,
			Modulation:         enums.Modulation(mod),
			TxGain:             txg,
			TxRes:              enums.GainResolution(txr),
			TxCoeff:            txc,
			Lqi:                lqi,
			PhaseDifferential:  pd,
			TMRValidTime:       tmr,
			NeighbourValidTime: nvt,
		})
	}
	return reader.ReadEndElement("NeighbourTable")
}

func (g *GXDLMSG3PlcMacSetup) loadMacPosTable(reader *GXXmlReader) error {
	g.MacPosTable = make([]GXDLMSMacPosTable, 0)
	ok, err := reader.IsStartElementNamed("MacPosTable", true)
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}
	for {
		ok, err = reader.IsStartElementNamed("Item", true)
		if err != nil {
			return err
		}
		if !ok {
			break
		}
		sa, err := reader.ReadElementContentAsUInt16("ShortAddress", 0)
		if err != nil {
			return err
		}
		lqi, err := reader.ReadElementContentAsUInt8("LQI", 0)
		if err != nil {
			return err
		}
		vt, err := reader.ReadElementContentAsUInt8("ValidTime", 0)
		if err != nil {
			return err
		}
		g.MacPosTable = append(g.MacPosTable, GXDLMSMacPosTable{ShortAddress: sa, LQI: lqi, ValidTime: vt})
	}
	return reader.ReadEndElement("MacPosTable")
}

func (g *GXDLMSG3PlcMacSetup) saveKeyTable(writer *GXXmlWriter) error {
	if err := writer.WriteStartElement("KeyTable"); err != nil {
		return err
	}
	for _, it := range g.KeyTable {
		if err := writer.WriteStartElement("Item"); err != nil {
			return err
		}
		if err := writer.WriteElementString("Key", it.Key); err != nil {
			return err
		}
		if err := writer.WriteElementString("Data", types.ToHex(it.Value, false)); err != nil {
			return err
		}
		if err := writer.WriteEndElement(); err != nil {
			return err
		}
	}
	return writer.WriteEndElement()
}

func (g *GXDLMSG3PlcMacSetup) saveNeighbourTable(writer *GXXmlWriter) error {
	if err := writer.WriteStartElement("NeighbourTable"); err != nil {
		return err
	}
	for _, it := range g.NeighbourTable {
		if err := writer.WriteStartElement("Item"); err != nil {
			return err
		}
		if err := writer.WriteElementString("ShortAddress", it.ShortAddress); err != nil {
			return err
		}
		if err := writer.WriteElementStringBool("Enabled", it.Enabled); err != nil {
			return err
		}
		if err := writer.WriteElementString("ToneMap", it.ToneMap); err != nil {
			return err
		}
		if err := writer.WriteElementString("Modulation", int(it.Modulation)); err != nil {
			return err
		}
		if err := writer.WriteElementString("TxGain", it.TxGain); err != nil {
			return err
		}
		if err := writer.WriteElementString("TxRes", int(it.TxRes)); err != nil {
			return err
		}
		if err := writer.WriteElementString("TxCoeff", it.TxCoeff); err != nil {
			return err
		}
		if err := writer.WriteElementString("Lqi", it.Lqi); err != nil {
			return err
		}
		if err := writer.WriteElementString("PhaseDifferential", it.PhaseDifferential); err != nil {
			return err
		}
		if err := writer.WriteElementString("TMRValidTime", it.TMRValidTime); err != nil {
			return err
		}
		if err := writer.WriteElementString("NeighbourValidTime", it.NeighbourValidTime); err != nil {
			return err
		}
		if err := writer.WriteEndElement(); err != nil {
			return err
		}
	}
	return writer.WriteEndElement()
}

func (g *GXDLMSG3PlcMacSetup) saveMacPosTable(writer *GXXmlWriter) error {
	if err := writer.WriteStartElement("MacPosTable"); err != nil {
		return err
	}
	for _, it := range g.MacPosTable {
		if err := writer.WriteStartElement("Item"); err != nil {
			return err
		}
		if err := writer.WriteElementString("ShortAddress", it.ShortAddress); err != nil {
			return err
		}
		if err := writer.WriteElementString("LQI", it.LQI); err != nil {
			return err
		}
		if err := writer.WriteElementString("ValidTime", it.ValidTime); err != nil {
			return err
		}
		if err := writer.WriteEndElement(); err != nil {
			return err
		}
	}
	return writer.WriteEndElement()
}

func macToAnySlice(value any) ([]any, bool) {
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

// NewGXDLMSG3PlcMacSetup creates a new G3-PLC MAC setup object.
func NewGXDLMSG3PlcMacSetup(ln string, sn int16) (*GXDLMSG3PlcMacSetup, error) {
	if err := ValidateLogicalName(ln); err != nil {
		return nil, err
	}
	return &GXDLMSG3PlcMacSetup{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeG3PlcMacSetup,
			logicalName: ln,
			ShortName:   sn,
			Version:     3,
		},
		KeyTable:                     make([]types.GXKeyValuePair[byte, []byte], 0),
		ShortAddress:                 0xFFFF,
		RcCoord:                      0xFFFF,
		PANId:                        0xFFFF,
		TmrTtl:                       2,
		MaxFrameRetries:              5,
		NeighbourTableEntryTtl:       255,
		HighPriorityWindowSize:       7,
		CscmFairnessLimit:            25,
		BeaconRandomizationWindowLen: 12,
		A:                            8,
		K:                            5,
		MinCwAttempts:                10,
		CenelecLegacyMode:            1,
		FccLegacyMode:                1,
		MaxBe:                        8,
		MaxCsmaBackoffs:              50,
		MinBe:                        3,
		NeighbourTable:               make([]GXDLMSNeighbourTable, 0),
		MacPosTable:                  make([]GXDLMSMacPosTable, 0),
	}, nil
}
