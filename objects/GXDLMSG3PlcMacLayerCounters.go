package objects

import (
	"github.com/Gurux/gxdlms-go/dlmserrors"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
)

// GXDLMSG3PlcMacLayerCounters contains G3-PLC MAC layer statistic counters.
type GXDLMSG3PlcMacLayerCounters struct {
	GXDLMSObject
	TxDataPacketCount    uint32
	RxDataPacketCount    uint32
	TxCmdPacketCount     uint32
	RxCmdPacketCount     uint32
	CSMAFailCount        uint32
	CSMANoAckCount       uint32
	BadCrcCount          uint32
	TxDataBroadcastCount uint32
	RxDataBroadcastCount uint32
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSG3PlcMacLayerCounters) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

func (g *GXDLMSG3PlcMacLayerCounters) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	if e.Index == 1 {
		g.TxDataPacketCount = 0
		g.RxDataPacketCount = 0
		g.TxCmdPacketCount = 0
		g.RxCmdPacketCount = 0
		g.CSMAFailCount = 0
		g.CSMANoAckCount = 0
		g.BadCrcCount = 0
		g.TxDataBroadcastCount = 0
		g.RxDataBroadcastCount = 0
	} else {
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil, nil
}

func (g *GXDLMSG3PlcMacLayerCounters) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	for i := 2; i <= 10; i++ {
		if all || g.CanRead(i) {
			attributes = append(attributes, i)
		}
	}
	return attributes
}

func (g *GXDLMSG3PlcMacLayerCounters) GetNames() []string {
	return []string{
		"Logical Name",
		"TxDataPacketCount",
		"RxDataPacketCount",
		"TxCmdPacketCount",
		"RxCmdPacketCount",
		"CSMAFailCount",
		"CSMANoAckCount",
		"BadCrcCount",
		"TxDataBroadcastCount",
		"RxDataBroadcastCount",
	}
}

func (g *GXDLMSG3PlcMacLayerCounters) GetMethodNames() []string { return []string{"Reset"} }
func (g *GXDLMSG3PlcMacLayerCounters) GetAttributeCount() int   { return 10 }
func (g *GXDLMSG3PlcMacLayerCounters) GetMethodCount() int      { return 1 }

func (g *GXDLMSG3PlcMacLayerCounters) GetDataType(index int) (enums.DataType, error) {
	if index == 1 {
		return enums.DataTypeOctetString, nil
	}
	if index >= 2 && index <= 10 {
		return enums.DataTypeUint32, nil
	}
	return enums.DataTypeNone, dlmserrors.ErrInvalidAttributeIndex
}

func (g *GXDLMSG3PlcMacLayerCounters) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	switch e.Index {
	case 1:
		return helpers.LogicalNameToBytes(g.LogicalName())
	case 2:
		return g.TxDataPacketCount, nil
	case 3:
		return g.RxDataPacketCount, nil
	case 4:
		return g.TxCmdPacketCount, nil
	case 5:
		return g.RxCmdPacketCount, nil
	case 6:
		return g.CSMAFailCount, nil
	case 7:
		return g.CSMANoAckCount, nil
	case 8:
		return g.BadCrcCount, nil
	case 9:
		return g.TxDataBroadcastCount, nil
	case 10:
		return g.RxDataBroadcastCount, nil
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
		return nil, nil
	}
}

func (g *GXDLMSG3PlcMacLayerCounters) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
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
		g.TxDataPacketCount = v
	case 3:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.RxDataPacketCount = v
	case 4:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.TxCmdPacketCount = v
	case 5:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.RxCmdPacketCount = v
	case 6:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.CSMAFailCount = v
	case 7:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.CSMANoAckCount = v
	case 8:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.BadCrcCount = v
	case 9:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.TxDataBroadcastCount = v
	case 10:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.RxDataBroadcastCount = v
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil
}

func (g *GXDLMSG3PlcMacLayerCounters) Load(reader *GXXmlReader) error {
	var err error
	g.TxDataPacketCount, err = reader.ReadElementContentAsUInt32("TxDataPacketCount", 0)
	if err != nil {
		return err
	}
	g.RxDataPacketCount, err = reader.ReadElementContentAsUInt32("RxDataPacketCount", 0)
	if err != nil {
		return err
	}
	g.TxCmdPacketCount, err = reader.ReadElementContentAsUInt32("TxCmdPacketCount", 0)
	if err != nil {
		return err
	}
	g.RxCmdPacketCount, err = reader.ReadElementContentAsUInt32("RxCmdPacketCount", 0)
	if err != nil {
		return err
	}
	g.CSMAFailCount, err = reader.ReadElementContentAsUInt32("CSMAFailCount", 0)
	if err != nil {
		return err
	}
	g.CSMANoAckCount, err = reader.ReadElementContentAsUInt32("CSMANoAckCount", 0)
	if err != nil {
		return err
	}
	g.BadCrcCount, err = reader.ReadElementContentAsUInt32("BadCrcCount", 0)
	if err != nil {
		return err
	}
	g.TxDataBroadcastCount, err = reader.ReadElementContentAsUInt32("TxDataBroadcastCount", 0)
	if err != nil {
		return err
	}
	g.RxDataBroadcastCount, err = reader.ReadElementContentAsUInt32("RxDataBroadcastCount", 0)
	return err
}

func (g *GXDLMSG3PlcMacLayerCounters) Save(writer *GXXmlWriter) error {
	if err := writer.WriteElementString("TxDataPacketCount", g.TxDataPacketCount); err != nil {
		return err
	}
	if err := writer.WriteElementString("RxDataPacketCount", g.RxDataPacketCount); err != nil {
		return err
	}
	if err := writer.WriteElementString("TxCmdPacketCount", g.TxCmdPacketCount); err != nil {
		return err
	}
	if err := writer.WriteElementString("RxCmdPacketCount", g.RxCmdPacketCount); err != nil {
		return err
	}
	if err := writer.WriteElementString("CSMAFailCount", g.CSMAFailCount); err != nil {
		return err
	}
	if err := writer.WriteElementString("CSMANoAckCount", g.CSMANoAckCount); err != nil {
		return err
	}
	if err := writer.WriteElementString("BadCrcCount", g.BadCrcCount); err != nil {
		return err
	}
	if err := writer.WriteElementString("TxDataBroadcastCount", g.TxDataBroadcastCount); err != nil {
		return err
	}
	return writer.WriteElementString("RxDataBroadcastCount", g.RxDataBroadcastCount)
}

func (g *GXDLMSG3PlcMacLayerCounters) PostLoad(reader *GXXmlReader) error { return nil }

func (g *GXDLMSG3PlcMacLayerCounters) GetValues() []any {
	return []any{
		g.LogicalName(),
		g.TxDataPacketCount,
		g.RxDataPacketCount,
		g.TxCmdPacketCount,
		g.RxCmdPacketCount,
		g.CSMAFailCount,
		g.CSMANoAckCount,
		g.BadCrcCount,
		g.TxDataBroadcastCount,
		g.RxDataBroadcastCount,
	}
}

// Reset resets all MAC layer counters.
func (g *GXDLMSG3PlcMacLayerCounters) Reset(client IGXDLMSClient) ([][]byte, error) {
	return client.Method(g, 1, int8(0), enums.DataTypeInt8)
}

// NewGXDLMSG3PlcMacLayerCounters creates a new G3-PLC MAC layer counters object.
func NewGXDLMSG3PlcMacLayerCounters(ln string, sn int16) (*GXDLMSG3PlcMacLayerCounters, error) {
	if err := ValidateLogicalName(ln); err != nil {
		return nil, err
	}
	return &GXDLMSG3PlcMacLayerCounters{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeG3PlcMacLayerCounters,
			logicalName: ln,
			ShortName:   sn,
			Version:     1,
		},
	}, nil
}
