package objects

import (
	"github.com/Gurux/gxdlms-go/dlmserrors"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
)

// GXDLMSPrimeNbOfdmPlcMacCounters represents PRIME NB OFDM PLC MAC counters.
type GXDLMSPrimeNbOfdmPlcMacCounters struct {
	GXDLMSObject
	TxDataPktCount  uint32
	RxDataPktCount  uint32
	TxCtrlPktCount  uint32
	RxCtrlPktCount  uint32
	CsmaFailCount   uint32
	CsmaChBusyCount uint32
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSPrimeNbOfdmPlcMacCounters) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

func (g *GXDLMSPrimeNbOfdmPlcMacCounters) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	if e.Index == 1 {
		g.TxDataPktCount, g.RxDataPktCount, g.TxCtrlPktCount, g.RxCtrlPktCount, g.CsmaFailCount, g.CsmaChBusyCount = 0, 0, 0, 0, 0, 0
	} else {
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil, nil
}
func (g *GXDLMSPrimeNbOfdmPlcMacCounters) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	for i := 2; i <= 7; i++ {
		if all || g.CanRead(i) {
			attributes = append(attributes, i)
		}
	}
	return attributes
}
func (g *GXDLMSPrimeNbOfdmPlcMacCounters) GetNames() []string {
	return []string{"Logical Name", "TxDataPktCount", "RxDataPktCount", "TxCtrlPktCount", "RxCtrlPktCount", "CsmaFailCount", "CsmaChBusyCount"}
}
func (g *GXDLMSPrimeNbOfdmPlcMacCounters) GetMethodNames() []string { return []string{"Reset"} }
func (g *GXDLMSPrimeNbOfdmPlcMacCounters) GetAttributeCount() int   { return 7 }
func (g *GXDLMSPrimeNbOfdmPlcMacCounters) GetMethodCount() int      { return 1 }
func (g *GXDLMSPrimeNbOfdmPlcMacCounters) GetDataType(index int) (enums.DataType, error) {
	if index == 1 {
		return enums.DataTypeOctetString, nil
	}
	if index >= 2 && index <= 7 {
		return enums.DataTypeUint32, nil
	}
	return 0, dlmserrors.ErrInvalidAttributeIndex
}
func (g *GXDLMSPrimeNbOfdmPlcMacCounters) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	switch e.Index {
	case 1:
		return helpers.LogicalNameToBytes(g.LogicalName())
	case 2:
		return g.TxDataPktCount, nil
	case 3:
		return g.RxDataPktCount, nil
	case 4:
		return g.TxCtrlPktCount, nil
	case 5:
		return g.RxCtrlPktCount, nil
	case 6:
		return g.CsmaFailCount, nil
	case 7:
		return g.CsmaChBusyCount, nil
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
		return nil, nil
	}
}
func (g *GXDLMSPrimeNbOfdmPlcMacCounters) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
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
		g.TxDataPktCount = v
	case 3:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.RxDataPktCount = v
	case 4:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.TxCtrlPktCount = v
	case 5:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.RxCtrlPktCount = v
	case 6:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.CsmaFailCount = v
	case 7:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.CsmaChBusyCount = v
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil
}
func (g *GXDLMSPrimeNbOfdmPlcMacCounters) Load(reader *GXXmlReader) error {
	var err error
	g.TxDataPktCount, err = reader.ReadElementContentAsUInt32("TxDataPktCount", 0)
	if err != nil {
		return err
	}
	g.RxDataPktCount, err = reader.ReadElementContentAsUInt32("RxDataPktCount", 0)
	if err != nil {
		return err
	}
	g.TxCtrlPktCount, err = reader.ReadElementContentAsUInt32("TxCtrlPktCount", 0)
	if err != nil {
		return err
	}
	g.RxCtrlPktCount, err = reader.ReadElementContentAsUInt32("RxCtrlPktCount", 0)
	if err != nil {
		return err
	}
	g.CsmaFailCount, err = reader.ReadElementContentAsUInt32("CsmaFailCount", 0)
	if err != nil {
		return err
	}
	g.CsmaChBusyCount, err = reader.ReadElementContentAsUInt32("CsmaChBusyCount", 0)
	return err
}
func (g *GXDLMSPrimeNbOfdmPlcMacCounters) Save(writer *GXXmlWriter) error {
	if err := writer.WriteElementString("TxDataPktCount", g.TxDataPktCount); err != nil {
		return err
	}
	if err := writer.WriteElementString("RxDataPktCount", g.RxDataPktCount); err != nil {
		return err
	}
	if err := writer.WriteElementString("TxCtrlPktCount", g.TxCtrlPktCount); err != nil {
		return err
	}
	if err := writer.WriteElementString("RxCtrlPktCount", g.RxCtrlPktCount); err != nil {
		return err
	}
	if err := writer.WriteElementString("CsmaFailCount", g.CsmaFailCount); err != nil {
		return err
	}
	return writer.WriteElementString("CsmaChBusyCount", g.CsmaChBusyCount)
}
func (g *GXDLMSPrimeNbOfdmPlcMacCounters) PostLoad(reader *GXXmlReader) error { return nil }
func (g *GXDLMSPrimeNbOfdmPlcMacCounters) GetValues() []any {
	return []any{g.LogicalName(), g.TxDataPktCount, g.RxDataPktCount, g.TxCtrlPktCount, g.RxCtrlPktCount, g.CsmaFailCount, g.CsmaChBusyCount}
}
func (g *GXDLMSPrimeNbOfdmPlcMacCounters) Reset(client IGXDLMSClient) ([][]byte, error) {
	return client.Method(g, 1, int8(0), enums.DataTypeInt8)
}
func NewGXDLMSPrimeNbOfdmPlcMacCounters(ln string, sn int16) (*GXDLMSPrimeNbOfdmPlcMacCounters, error) {
	if err := ValidateLogicalName(ln); err != nil {
		return nil, err
	}
	return &GXDLMSPrimeNbOfdmPlcMacCounters{GXDLMSObject: GXDLMSObject{objectType: enums.ObjectTypePrimeNbOfdmPlcMacCounters, logicalName: ln, ShortName: sn}}, nil
}
