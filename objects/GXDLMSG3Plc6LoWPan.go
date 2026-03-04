package objects

import (
	"fmt"

	"github.com/Gurux/gxcommon-go"
	"github.com/Gurux/gxdlms-go/dlmserrors"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

type GXDLMSRoutingConfiguration struct {
	NetTraversalTime       uint8
	RoutingTableEntryTtl   uint16
	Kr, Km, Kc, Kq, Kh     uint8
	Krt, RreqRetries       uint8
	RreqReqWait            uint8
	BlacklistTableEntryTtl uint16
	UnicastRreqGenEnable   bool
	RlcEnabled             bool
	AddRevLinkCost         uint8
}

type GXDLMSRoutingTable struct {
	DestinationAddress uint16
	NextHopAddress     uint16
	RouteCost          uint16
	HopCount           uint8
	WeakLinkCount      uint8
	ValidTime          uint16
}

type GXDLMSContextInformationTable struct {
	CID           string
	Context       []byte
	Compression   bool
	ValidLifetime uint16
}

type GXDLMSBroadcastLogTable struct {
	SourceAddress  uint16
	SequenceNumber uint8
	ValidTime      uint16
}

// GXDLMSG3Plc6LoWPan models G3-PLC 6LoWPAN adaptation layer setup.
type GXDLMSG3Plc6LoWPan struct {
	GXDLMSObject
	MaxHops, WeakLqiValue, SecurityLevel uint8
	PrefixTable                          []uint8
	RoutingConfiguration                 []GXDLMSRoutingConfiguration
	BroadcastLogTableTtl                 uint16
	RoutingTable                         []GXDLMSRoutingTable
	ContextInformation                   []GXDLMSContextInformationTable
	BlacklistTable                       []types.GXKeyValuePair[uint16, uint16]
	BroadcastLogTable                    []GXDLMSBroadcastLogTable
	GroupTable                           []uint16
	MaxJoinWaitTime                      uint16
	PathDiscoveryTime, ActiveKeyIndex    uint8
	MetricType                           uint8
	CoordShortAddress                    uint16
	DisableDefaultRouting                bool
	DeviceType                           enums.DeviceType
	DefaultCoordRoute                    bool
	DestinationAddress                   []uint16
	LowLQI, HighLQI                      uint8
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSG3Plc6LoWPan) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

func (g *GXDLMSG3Plc6LoWPan) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, nil
}
func (g *GXDLMSG3Plc6LoWPan) GetMethodNames() []string { return []string{} }
func (g *GXDLMSG3Plc6LoWPan) GetMethodCount() int      { return 0 }

func (g *GXDLMSG3Plc6LoWPan) GetAttributeCount() int {
	switch g.Version {
	case 0:
		return 16
	case 1:
		return 19
	case 2:
		return 21
	default:
		return 23
	}
}

func (g *GXDLMSG3Plc6LoWPan) GetAttributeIndexToRead(all bool) []int {
	var a []int
	if all || g.LogicalName() == "" {
		a = append(a, 1)
	}
	for i := 2; i <= g.GetAttributeCount(); i++ {
		if all || g.CanRead(i) {
			a = append(a, i)
		}
	}
	return a
}

func (g *GXDLMSG3Plc6LoWPan) GetNames() []string {
	return []string{
		"Logical Name", "MaxHops", "WeakLqiValue", "SecurityLevel", "PrefixTable", "RoutingConfiguration",
		"BroadcastLogTableEntryTtl", "RoutingTable", "ContextInformationTable", "BlacklistTable", "BroadcastLogTable",
		"GroupTable", "MaxJoinWaitTime", "PathDiscoveryTime", "ActiveKeyIndex", "MetricType", "CoordShortAddress",
		"DisableDefaultRouting", "DeviceType", "Default coord route enabled", "Destination address", "Low LQI", "High LQI",
	}
}

func (g *GXDLMSG3Plc6LoWPan) GetDataType(index int) (enums.DataType, error) {
	switch index {
	case 1:
		return enums.DataTypeOctetString, nil
	case 2, 3, 4, 14, 15, 16, 22, 23:
		return enums.DataTypeUint8, nil
	case 5, 6, 8, 9, 10, 11, 12, 21:
		return enums.DataTypeArray, nil
	case 7, 13, 17:
		return enums.DataTypeUint16, nil
	case 18, 20:
		return enums.DataTypeBoolean, nil
	case 19:
		return enums.DataTypeEnum, nil
	default:
		return enums.DataTypeNone, dlmserrors.ErrInvalidAttributeIndex
	}
}

func (g *GXDLMSG3Plc6LoWPan) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	switch e.Index {
	case 1:
		return helpers.LogicalNameToBytes(g.LogicalName())
	case 2:
		return g.MaxHops, nil
	case 3:
		return g.WeakLqiValue, nil
	case 4:
		return g.SecurityLevel, nil
	case 5:
		return lowpanEncodeU8(settings, g.PrefixTable)
	case 6:
		return lowpanEncodeRoutingConfiguration(settings, g.RoutingConfiguration)
	case 7:
		return g.BroadcastLogTableTtl, nil
	case 8:
		return lowpanEncodeRoutingTable(settings, g.RoutingTable)
	case 9:
		return lowpanEncodeContextInformation(settings, g.ContextInformation)
	case 10:
		return lowpanEncodeBlacklist(settings, g.BlacklistTable)
	case 11:
		return lowpanEncodeBroadcastLog(settings, g.BroadcastLogTable)
	case 12:
		return lowpanEncodeU16(settings, g.GroupTable)
	case 13:
		return g.MaxJoinWaitTime, nil
	case 14:
		return g.PathDiscoveryTime, nil
	case 15:
		return g.ActiveKeyIndex, nil
	case 16:
		return g.MetricType, nil
	case 17:
		return g.CoordShortAddress, nil
	case 18:
		return g.DisableDefaultRouting, nil
	case 19:
		return uint8(g.DeviceType), nil
	case 20:
		return g.DefaultCoordRoute, nil
	case 21:
		return lowpanEncodeU16(settings, g.DestinationAddress)
	case 22:
		return g.LowLQI, nil
	case 23:
		return g.HighLQI, nil
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
		return nil, nil
	}
}

func (g *GXDLMSG3Plc6LoWPan) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	switch e.Index {
	case 1:
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			return err
		}
		return g.SetLogicalName(ln)
	case 2:
		g.MaxHops = e.Value.(uint8)
	case 3:
		g.WeakLqiValue = e.Value.(uint8)
	case 4:
		g.SecurityLevel = e.Value.(uint8)
	case 5:
		v, err := lowpanParseU8Array(e.Value)
		if err != nil {
			return err
		}
		g.PrefixTable = v
	case 6:
		v, err := lowpanParseRoutingConfiguration(e.Value)
		if err != nil {
			return err
		}
		g.RoutingConfiguration = v
	case 7:
		v, _ := toUint32(e.Value)
		g.BroadcastLogTableTtl = uint16(v)
	case 8:
		v, err := lowpanParseRoutingTable(e.Value)
		if err != nil {
			return err
		}
		g.RoutingTable = v
	case 9:
		v, err := lowpanParseContextInfo(e.Value)
		if err != nil {
			return err
		}
		g.ContextInformation = v
	case 10:
		v, err := lowpanParseBlacklist(e.Value)
		if err != nil {
			return err
		}
		g.BlacklistTable = v
	case 11:
		v, err := lowpanParseBroadcastLog(e.Value)
		if err != nil {
			return err
		}
		g.BroadcastLogTable = v
	case 12:
		v, err := lowpanParseU16Array(e.Value)
		if err != nil {
			return err
		}
		g.GroupTable = v
	case 13:
		v, _ := toUint32(e.Value)
		g.MaxJoinWaitTime = uint16(v)
	case 14:
		g.PathDiscoveryTime = e.Value.(uint8)
	case 15:
		g.ActiveKeyIndex = e.Value.(uint8)
	case 16:
		g.MetricType = e.Value.(uint8)
	case 17:
		v, _ := toUint32(e.Value)
		g.CoordShortAddress = uint16(v)
	case 18:
		v, _ := toBool(e.Value)
		g.DisableDefaultRouting = v
	case 19:
		v, _ := toUint32(e.Value)
		g.DeviceType = enums.DeviceType(v)
	case 20:
		v, _ := toBool(e.Value)
		g.DefaultCoordRoute = v
	case 21:
		v, err := lowpanParseU16Array(e.Value)
		if err != nil {
			return err
		}
		g.DestinationAddress = v
	case 22:
		g.LowLQI = e.Value.(uint8)
	case 23:
		g.HighLQI = e.Value.(uint8)
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil
}

func (g *GXDLMSG3Plc6LoWPan) Load(reader *GXXmlReader) error {
	var err error
	g.MaxHops, err = reader.ReadElementContentAsUInt8("MaxHops", 0)
	if err != nil {
		return err
	}
	g.WeakLqiValue, err = reader.ReadElementContentAsUInt8("WeakLqiValue", 0)
	if err != nil {
		return err
	}
	g.SecurityLevel, err = reader.ReadElementContentAsUInt8("SecurityLevel", 0)
	if err != nil {
		return err
	}
	g.BroadcastLogTableTtl, _ = reader.ReadElementContentAsUInt16("BroadcastLogTableEntryTtl", 0)
	g.MaxJoinWaitTime, _ = reader.ReadElementContentAsUInt16("MaxJoinWaitTime", 0)
	g.PathDiscoveryTime, _ = reader.ReadElementContentAsUInt8("PathDiscoveryTime", 0)
	g.ActiveKeyIndex, _ = reader.ReadElementContentAsUInt8("ActiveKeyIndex", 0)
	g.MetricType, _ = reader.ReadElementContentAsUInt8("MetricType", 0)
	g.CoordShortAddress, _ = reader.ReadElementContentAsUInt16("CoordShortAddress", 0)
	ddr, _ := reader.ReadElementContentAsInt("DisableDefaultRouting", 0)
	g.DisableDefaultRouting = ddr != 0
	dt, _ := reader.ReadElementContentAsInt("DeviceType", int(enums.DeviceTypeNotDefined))
	g.DeviceType = enums.DeviceType(dt)
	dcr, _ := reader.ReadElementContentAsInt("DefaultCoordRouteEnabled", 0)
	g.DefaultCoordRoute = dcr != 0
	g.LowLQI, _ = reader.ReadElementContentAsUInt8("LowLQI", 0)
	g.HighLQI, _ = reader.ReadElementContentAsUInt8("HighLQI", 0)
	return nil
}

func (g *GXDLMSG3Plc6LoWPan) Save(writer *GXXmlWriter) error {
	if err := writer.WriteElementString("MaxHops", g.MaxHops); err != nil {
		return err
	}
	if err := writer.WriteElementString("WeakLqiValue", g.WeakLqiValue); err != nil {
		return err
	}
	if err := writer.WriteElementString("SecurityLevel", g.SecurityLevel); err != nil {
		return err
	}
	if err := writer.WriteElementString("BroadcastLogTableEntryTtl", g.BroadcastLogTableTtl); err != nil {
		return err
	}
	if err := writer.WriteElementString("MaxJoinWaitTime", g.MaxJoinWaitTime); err != nil {
		return err
	}
	if err := writer.WriteElementString("PathDiscoveryTime", g.PathDiscoveryTime); err != nil {
		return err
	}
	if err := writer.WriteElementString("ActiveKeyIndex", g.ActiveKeyIndex); err != nil {
		return err
	}
	if err := writer.WriteElementString("MetricType", g.MetricType); err != nil {
		return err
	}
	if err := writer.WriteElementString("CoordShortAddress", g.CoordShortAddress); err != nil {
		return err
	}
	if err := writer.WriteElementStringBool("DisableDefaultRouting", g.DisableDefaultRouting); err != nil {
		return err
	}
	if err := writer.WriteElementString("DeviceType", int(g.DeviceType)); err != nil {
		return err
	}
	if err := writer.WriteElementStringBool("DefaultCoordRouteEnabled", g.DefaultCoordRoute); err != nil {
		return err
	}
	if err := writer.WriteElementString("LowLQI", g.LowLQI); err != nil {
		return err
	}
	return writer.WriteElementString("HighLQI", g.HighLQI)
}

func (g *GXDLMSG3Plc6LoWPan) PostLoad(reader *GXXmlReader) error { return nil }

func (g *GXDLMSG3Plc6LoWPan) GetValues() []any {
	return []any{g.LogicalName(), g.MaxHops, g.WeakLqiValue, g.SecurityLevel, g.PrefixTable, g.RoutingConfiguration, g.BroadcastLogTableTtl, g.RoutingTable, g.ContextInformation, g.BlacklistTable, g.BroadcastLogTable, g.GroupTable, g.MaxJoinWaitTime, g.PathDiscoveryTime, g.ActiveKeyIndex, g.MetricType, g.CoordShortAddress, g.DisableDefaultRouting, g.DeviceType, g.DefaultCoordRoute, g.DestinationAddress, g.LowLQI, g.HighLQI}
}

func lowpanToAnySlice(value any) ([]any, bool) {
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
func lowpanEncodeU8(settings *settings.GXDLMSSettings, v []uint8) ([]byte, error) {
	bb := types.NewGXByteBuffer()
	_ = bb.SetUint8(uint8(enums.DataTypeArray))
	types.SetObjectCount(len(v), bb)
	for _, it := range v {
		_ = internal.SetData(settings, bb, enums.DataTypeUint8, it)
	}
	return bb.Array(), nil
}
func lowpanEncodeU16(settings *settings.GXDLMSSettings, v []uint16) ([]byte, error) {
	bb := types.NewGXByteBuffer()
	_ = bb.SetUint8(uint8(enums.DataTypeArray))
	types.SetObjectCount(len(v), bb)
	for _, it := range v {
		_ = internal.SetData(settings, bb, enums.DataTypeUint16, it)
	}
	return bb.Array(), nil
}
func lowpanParseU8Array(value any) ([]uint8, error) {
	rows, ok := lowpanToAnySlice(value)
	if !ok {
		return nil, fmt.Errorf("invalid uint8 array: %T", value)
	}
	ret := make([]uint8, 0, len(rows))
	for _, it := range rows {
		ret = append(ret, it.(uint8))
	}
	return ret, nil
}
func lowpanParseU16Array(value any) ([]uint16, error) {
	rows, ok := lowpanToAnySlice(value)
	if !ok {
		return nil, fmt.Errorf("invalid uint16 array: %T", value)
	}
	ret := make([]uint16, 0, len(rows))
	for _, it := range rows {
		v, err := toUint32(it)
		if err != nil {
			return nil, err
		}
		ret = append(ret, uint16(v))
	}
	return ret, nil
}

func lowpanEncodeRoutingConfiguration(settings *settings.GXDLMSSettings, value []GXDLMSRoutingConfiguration) ([]byte, error) {
	bb := types.NewGXByteBuffer()
	_ = bb.SetUint8(uint8(enums.DataTypeArray))
	types.SetObjectCount(len(value), bb)
	for _, it := range value {
		_ = bb.SetUint8(uint8(enums.DataTypeStructure))
		_ = bb.SetUint8(14)
		_ = internal.SetData(settings, bb, enums.DataTypeUint8, it.NetTraversalTime)
		_ = internal.SetData(settings, bb, enums.DataTypeUint16, it.RoutingTableEntryTtl)
		_ = internal.SetData(settings, bb, enums.DataTypeUint8, it.Kr)
		_ = internal.SetData(settings, bb, enums.DataTypeUint8, it.Km)
		_ = internal.SetData(settings, bb, enums.DataTypeUint8, it.Kc)
		_ = internal.SetData(settings, bb, enums.DataTypeUint8, it.Kq)
		_ = internal.SetData(settings, bb, enums.DataTypeUint8, it.Kh)
		_ = internal.SetData(settings, bb, enums.DataTypeUint8, it.Krt)
		_ = internal.SetData(settings, bb, enums.DataTypeUint8, it.RreqRetries)
		_ = internal.SetData(settings, bb, enums.DataTypeUint8, it.RreqReqWait)
		_ = internal.SetData(settings, bb, enums.DataTypeUint16, it.BlacklistTableEntryTtl)
		_ = internal.SetData(settings, bb, enums.DataTypeBoolean, it.UnicastRreqGenEnable)
		_ = internal.SetData(settings, bb, enums.DataTypeBoolean, it.RlcEnabled)
		_ = internal.SetData(settings, bb, enums.DataTypeUint8, it.AddRevLinkCost)
	}
	return bb.Array(), nil
}
func lowpanParseRoutingConfiguration(value any) ([]GXDLMSRoutingConfiguration, error) {
	rows, ok := lowpanToAnySlice(value)
	if !ok {
		return nil, fmt.Errorf("invalid routing configuration: %T", value)
	}
	ret := make([]GXDLMSRoutingConfiguration, 0, len(rows))
	for _, row := range rows {
		it, ok := lowpanToAnySlice(row)
		if !ok || len(it) < 14 {
			continue
		}
		n0, _ := it[0].(uint8)
		n1, _ := it[1].(uint32)
		n2, _ := it[2].(uint8)
		n3, _ := it[3].(uint8)
		n4, _ := it[4].(uint8)
		n5, _ := it[5].(uint8)
		n6, _ := it[6].(uint8)
		n7, _ := it[7].(uint8)
		n8, _ := it[8].(uint8)
		n9, _ := it[9].(uint8)
		n10, _ := it[10].(uint32)
		n11, _ := it[11].(bool)
		n12, _ := it[12].(bool)
		n13, _ := it[13].(uint8)
		ret = append(ret, GXDLMSRoutingConfiguration{NetTraversalTime: n0, RoutingTableEntryTtl: uint16(n1), Kr: n2, Km: n3, Kc: n4, Kq: n5, Kh: n6, Krt: n7, RreqRetries: n8, RreqReqWait: n9, BlacklistTableEntryTtl: uint16(n10), UnicastRreqGenEnable: n11, RlcEnabled: n12, AddRevLinkCost: n13})
	}
	return ret, nil
}

func lowpanEncodeRoutingTable(settings *settings.GXDLMSSettings, value []GXDLMSRoutingTable) ([]byte, error) {
	bb := types.NewGXByteBuffer()
	_ = bb.SetUint8(uint8(enums.DataTypeArray))
	types.SetObjectCount(len(value), bb)
	for _, it := range value {
		_ = bb.SetUint8(uint8(enums.DataTypeStructure))
		_ = bb.SetUint8(6)
		_ = internal.SetData(settings, bb, enums.DataTypeUint16, it.DestinationAddress)
		_ = internal.SetData(settings, bb, enums.DataTypeUint16, it.NextHopAddress)
		_ = internal.SetData(settings, bb, enums.DataTypeUint16, it.RouteCost)
		_ = internal.SetData(settings, bb, enums.DataTypeUint8, it.HopCount)
		_ = internal.SetData(settings, bb, enums.DataTypeUint8, it.WeakLinkCount)
		_ = internal.SetData(settings, bb, enums.DataTypeUint16, it.ValidTime)
	}
	return bb.Array(), nil
}
func lowpanParseRoutingTable(value any) ([]GXDLMSRoutingTable, error) {
	rows, ok := lowpanToAnySlice(value)
	if !ok {
		return nil, fmt.Errorf("invalid routing table: %T", value)
	}
	ret := make([]GXDLMSRoutingTable, 0, len(rows))
	for _, row := range rows {
		it, ok := lowpanToAnySlice(row)
		if !ok || len(it) < 6 {
			return nil, gxcommon.ErrArgumentOutOfRange
		}
		d, err := toUint32(it[0])
		if err != nil {
			return nil, err
		}
		n, err := toUint32(it[1])
		if err != nil {
			return nil, err
		}
		c, err := toUint32(it[2])
		if err != nil {
			return nil, err
		}
		h, err := toUint8(it[3])
		if err != nil {
			return nil, err
		}
		w, err := toUint8(it[4])
		if err != nil {
			return nil, err
		}
		v, err := toUint32(it[5])
		if err != nil {
			return nil, err
		}
		ret = append(ret, GXDLMSRoutingTable{DestinationAddress: uint16(d), NextHopAddress: uint16(n),
			RouteCost: uint16(c), HopCount: h, WeakLinkCount: w, ValidTime: uint16(v)})
	}
	return ret, nil
}

func lowpanEncodeContextInformation(settings *settings.GXDLMSSettings, value []GXDLMSContextInformationTable) ([]byte, error) {
	var err error
	bb := types.NewGXByteBuffer()
	_ = bb.SetUint8(uint8(enums.DataTypeArray))
	types.SetObjectCount(len(value), bb)
	for _, it := range value {
		err = bb.SetUint8(uint8(enums.DataTypeStructure))
		if err != nil {
			return nil, err
		}
		err = bb.SetUint8(5)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, bb, enums.DataTypeOctetString, []byte(it.CID))
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, bb, enums.DataTypeUint8, uint8(len(it.Context)))
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, bb, enums.DataTypeOctetString, it.Context)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, bb, enums.DataTypeBoolean, it.Compression)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, bb, enums.DataTypeUint16, it.ValidLifetime)
		if err != nil {
			return nil, err
		}
	}
	return bb.Array(), nil
}
func lowpanParseContextInfo(value any) ([]GXDLMSContextInformationTable, error) {
	rows, ok := lowpanToAnySlice(value)
	if !ok {
		return nil, fmt.Errorf("invalid context information table: %T", value)
	}
	ret := make([]GXDLMSContextInformationTable, 0, len(rows))
	for _, row := range rows {
		it, ok := lowpanToAnySlice(row)
		if !ok || len(it) < 5 {
			continue
		}
		cid := ""
		if b, ok := it[0].([]byte); ok {
			cid = string(b)
		} else {
			cid = fmt.Sprint(it[0])
		}
		ctx, _ := it[2].([]byte)
		comp, _ := toBool(it[3])
		vl, _ := toUint32(it[4])
		ret = append(ret, GXDLMSContextInformationTable{CID: cid, Context: ctx, Compression: comp, ValidLifetime: uint16(vl)})
	}
	return ret, nil
}

func lowpanEncodeBlacklist(settings *settings.GXDLMSSettings, value []types.GXKeyValuePair[uint16, uint16]) ([]byte, error) {
	bb := types.NewGXByteBuffer()
	_ = bb.SetUint8(uint8(enums.DataTypeArray))
	types.SetObjectCount(len(value), bb)
	for _, it := range value {
		_ = bb.SetUint8(uint8(enums.DataTypeStructure))
		_ = bb.SetUint8(2)
		_ = internal.SetData(settings, bb, enums.DataTypeUint16, it.Key)
		_ = internal.SetData(settings, bb, enums.DataTypeUint16, it.Value)
	}
	return bb.Array(), nil
}
func lowpanParseBlacklist(value any) ([]types.GXKeyValuePair[uint16, uint16], error) {
	rows, ok := lowpanToAnySlice(value)
	if !ok {
		return nil, fmt.Errorf("invalid blacklist table: %T", value)
	}
	ret := make([]types.GXKeyValuePair[uint16, uint16], 0, len(rows))
	for _, row := range rows {
		it, ok := lowpanToAnySlice(row)
		if !ok || len(it) < 2 {
			continue
		}
		k, _ := toUint32(it[0])
		v, _ := toUint32(it[1])
		ret = append(ret, *types.NewGXKeyValuePair(uint16(k), uint16(v)))
	}
	return ret, nil
}

func lowpanEncodeBroadcastLog(settings *settings.GXDLMSSettings, value []GXDLMSBroadcastLogTable) ([]byte, error) {
	bb := types.NewGXByteBuffer()
	_ = bb.SetUint8(uint8(enums.DataTypeArray))
	types.SetObjectCount(len(value), bb)
	for _, it := range value {
		_ = bb.SetUint8(uint8(enums.DataTypeStructure))
		_ = bb.SetUint8(3)
		_ = internal.SetData(settings, bb, enums.DataTypeUint16, it.SourceAddress)
		_ = internal.SetData(settings, bb, enums.DataTypeUint8, it.SequenceNumber)
		_ = internal.SetData(settings, bb, enums.DataTypeUint16, it.ValidTime)
	}
	return bb.Array(), nil
}
func lowpanParseBroadcastLog(value any) ([]GXDLMSBroadcastLogTable, error) {
	rows, ok := lowpanToAnySlice(value)
	if !ok {
		return nil, fmt.Errorf("invalid broadcast log table: %T", value)
	}
	ret := make([]GXDLMSBroadcastLogTable, 0, len(rows))
	for _, row := range rows {
		it, ok := lowpanToAnySlice(row)
		if !ok || len(it) < 3 {
			continue
		}
		s, _ := toUint32(it[0])
		q, _ := toUint8(it[1])
		v, _ := toUint32(it[2])
		ret = append(ret, GXDLMSBroadcastLogTable{SourceAddress: uint16(s), SequenceNumber: q, ValidTime: uint16(v)})
	}
	return ret, nil
}

// NewGXDLMSG3Plc6LoWPan creates a new G3-PLC 6LoWPAN setup object.
func NewGXDLMSG3Plc6LoWPan(ln string, sn int16) (*GXDLMSG3Plc6LoWPan, error) {
	if err := ValidateLogicalName(ln); err != nil {
		return nil, err
	}
	return &GXDLMSG3Plc6LoWPan{
		GXDLMSObject: GXDLMSObject{objectType: enums.ObjectTypeG3Plc6LoWPan, logicalName: ln, ShortName: sn, Version: 3},
		MaxHops:      8, WeakLqiValue: 52, SecurityLevel: 5, PrefixTable: []uint8{}, RoutingConfiguration: []GXDLMSRoutingConfiguration{},
		RoutingTable: []GXDLMSRoutingTable{}, ContextInformation: []GXDLMSContextInformationTable{}, BlacklistTable: []types.GXKeyValuePair[uint16, uint16]{},
		BroadcastLogTable: []GXDLMSBroadcastLogTable{}, BroadcastLogTableTtl: 2, GroupTable: []uint16{}, MaxJoinWaitTime: 20, PathDiscoveryTime: 40,
		ActiveKeyIndex: 0, MetricType: 0x0F, CoordShortAddress: 0, DisableDefaultRouting: false, DeviceType: enums.DeviceTypeNotDefined,
		DefaultCoordRoute: false, DestinationAddress: []uint16{},
	}, nil
}
