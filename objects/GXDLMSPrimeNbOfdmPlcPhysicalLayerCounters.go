package objects

import (
	"github.com/Gurux/gxdlms-go/dlmserrors"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
)

// GXDLMSPrimeNbOfdmPlcPhysicalLayerCounters represents PRIME NB OFDM PLC physical layer counters.
type GXDLMSPrimeNbOfdmPlcPhysicalLayerCounters struct {
	GXDLMSObject
	CrcIncorrectCount uint16
	CrcFailedCount    uint16
	TxDropCount       uint16
	RxDropCount       uint16
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSPrimeNbOfdmPlcPhysicalLayerCounters) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

func (g *GXDLMSPrimeNbOfdmPlcPhysicalLayerCounters) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	if e.Index == 1 {
		g.CrcIncorrectCount, g.CrcFailedCount, g.TxDropCount, g.RxDropCount = 0, 0, 0, 0
	} else {
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil, nil
}

func (g *GXDLMSPrimeNbOfdmPlcPhysicalLayerCounters) GetAttributeIndexToRead(all bool) []int {
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
	if all || g.CanRead(4) {
		attributes = append(attributes, 4)
	}
	if all || g.CanRead(5) {
		attributes = append(attributes, 5)
	}
	return attributes
}

func (g *GXDLMSPrimeNbOfdmPlcPhysicalLayerCounters) GetNames() []string {
	return []string{"Logical Name", "CrcIncorrectCount", "CrcFailedCount", "TxDropCount", "RxDropCount"}
}

func (g *GXDLMSPrimeNbOfdmPlcPhysicalLayerCounters) GetMethodNames() []string {
	return []string{"Reset"}
}
func (g *GXDLMSPrimeNbOfdmPlcPhysicalLayerCounters) GetAttributeCount() int { return 5 }
func (g *GXDLMSPrimeNbOfdmPlcPhysicalLayerCounters) GetMethodCount() int    { return 1 }

func (g *GXDLMSPrimeNbOfdmPlcPhysicalLayerCounters) GetDataType(index int) (enums.DataType, error) {
	switch index {
	case 1:
		return enums.DataTypeOctetString, nil
	case 2, 3, 4, 5:
		return enums.DataTypeUint16, nil
	default:
		return 0, dlmserrors.ErrInvalidAttributeIndex
	}
}

func (g *GXDLMSPrimeNbOfdmPlcPhysicalLayerCounters) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	switch e.Index {
	case 1:
		return helpers.LogicalNameToBytes(g.LogicalName())
	case 2:
		return g.CrcIncorrectCount, nil
	case 3:
		return g.CrcFailedCount, nil
	case 4:
		return g.TxDropCount, nil
	case 5:
		return g.RxDropCount, nil
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
		return nil, nil
	}
}

func (g *GXDLMSPrimeNbOfdmPlcPhysicalLayerCounters) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
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
		g.CrcIncorrectCount, err = toUint16(e.Value)
	case 3:
		g.CrcFailedCount, err = toUint16(e.Value)
	case 4:
		g.TxDropCount, err = toUint16(e.Value)
	case 5:
		g.RxDropCount, err = toUint16(e.Value)
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return err
}

func (g *GXDLMSPrimeNbOfdmPlcPhysicalLayerCounters) Load(reader *GXXmlReader) error {
	var err error
	g.CrcIncorrectCount, err = reader.ReadElementContentAsUInt16("CrcIncorrectCount", 0)
	if err != nil {
		return err
	}
	g.CrcFailedCount, err = reader.ReadElementContentAsUInt16("CrcFailedCount", 0)
	if err != nil {
		return err
	}
	g.TxDropCount, err = reader.ReadElementContentAsUInt16("TxDropCount", 0)
	if err != nil {
		return err
	}
	g.RxDropCount, err = reader.ReadElementContentAsUInt16("RxDropCount", 0)
	return err
}

func (g *GXDLMSPrimeNbOfdmPlcPhysicalLayerCounters) Save(writer *GXXmlWriter) error {
	if err := writer.WriteElementString("CrcIncorrectCount", g.CrcIncorrectCount); err != nil {
		return err
	}
	if err := writer.WriteElementString("CrcFailedCount", g.CrcFailedCount); err != nil {
		return err
	}
	if err := writer.WriteElementString("TxDropCount", g.TxDropCount); err != nil {
		return err
	}
	return writer.WriteElementString("RxDropCount", g.RxDropCount)
}

func (g *GXDLMSPrimeNbOfdmPlcPhysicalLayerCounters) PostLoad(reader *GXXmlReader) error { return nil }

func (g *GXDLMSPrimeNbOfdmPlcPhysicalLayerCounters) GetValues() []any {
	return []any{g.LogicalName(), g.CrcIncorrectCount, g.CrcFailedCount, g.TxDropCount, g.RxDropCount}
}

func NewGXDLMSPrimeNbOfdmPlcPhysicalLayerCounters(ln string, sn int16) (*GXDLMSPrimeNbOfdmPlcPhysicalLayerCounters, error) {
	if err := ValidateLogicalName(ln); err != nil {
		return nil, err
	}
	return &GXDLMSPrimeNbOfdmPlcPhysicalLayerCounters{GXDLMSObject: GXDLMSObject{objectType: enums.ObjectTypePrimeNbOfdmPlcPhysicalLayerCounters, logicalName: ln, ShortName: sn}}, nil
}
