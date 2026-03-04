package objects

import (
	"github.com/Gurux/gxdlms-go/dlmserrors"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
)

// GXDLMSIec8802LlcType2Setup represents ISO/IEC 8802-2 LLC type 2 setup.
type GXDLMSIec8802LlcType2Setup struct {
	GXDLMSObject
	TransmitWindowSizeK        uint8
	TransmitWindowSizeRW       uint8
	MaximumOctetsPdu           uint16
	MaximumNumberTransmissions uint8
	AcknowledgementTimer       uint16
	BitTimer                   uint16
	RejectTimer                uint16
	BusyStateTimer             uint16
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSIec8802LlcType2Setup) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

func (g *GXDLMSIec8802LlcType2Setup) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, nil
}
func (g *GXDLMSIec8802LlcType2Setup) GetAttributeIndexToRead(all bool) []int {
	var a []int
	if all || g.LogicalName() == "" {
		a = append(a, 1)
	}
	for i := 2; i <= 9; i++ {
		if all || g.CanRead(i) {
			a = append(a, i)
		}
	}
	return a
}
func (g *GXDLMSIec8802LlcType2Setup) GetNames() []string {
	return []string{"Logical Name", "TransmitWindowSizeK", "TransmitWindowSizeRW", "MaximumOctetsPdu", "MaximumNumberTransmissions", "AcknowledgementTimer", "BitTimer", "RejectTimer", "BusyStateTimer"}
}
func (g *GXDLMSIec8802LlcType2Setup) GetMethodNames() []string { return []string{} }
func (g *GXDLMSIec8802LlcType2Setup) GetAttributeCount() int   { return 9 }
func (g *GXDLMSIec8802LlcType2Setup) GetMethodCount() int      { return 0 }
func (g *GXDLMSIec8802LlcType2Setup) GetDataType(index int) (enums.DataType, error) {
	switch index {
	case 1:
		return enums.DataTypeOctetString, nil
	case 2, 3, 5:
		return enums.DataTypeUint8, nil
	case 4, 6, 7, 8, 9:
		return enums.DataTypeUint16, nil
	default:
		return 0, dlmserrors.ErrInvalidAttributeIndex
	}
}
func (g *GXDLMSIec8802LlcType2Setup) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	switch e.Index {
	case 1:
		return helpers.LogicalNameToBytes(g.LogicalName())
	case 2:
		return g.TransmitWindowSizeK, nil
	case 3:
		return g.TransmitWindowSizeRW, nil
	case 4:
		return g.MaximumOctetsPdu, nil
	case 5:
		return g.MaximumNumberTransmissions, nil
	case 6:
		return g.AcknowledgementTimer, nil
	case 7:
		return g.BitTimer, nil
	case 8:
		return g.RejectTimer, nil
	case 9:
		return g.BusyStateTimer, nil
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
		return nil, nil
	}
}
func (g *GXDLMSIec8802LlcType2Setup) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	switch e.Index {
	case 1:
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return err
		}
		return g.SetLogicalName(ln)
	case 2:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.TransmitWindowSizeK = v
	case 3:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.TransmitWindowSizeRW = v
	case 4:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.MaximumOctetsPdu = uint16(v)
	case 5:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.MaximumNumberTransmissions = v
	case 6:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.AcknowledgementTimer = uint16(v)
	case 7:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.BitTimer = uint16(v)
	case 8:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.RejectTimer = uint16(v)
	case 9:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.BusyStateTimer = uint16(v)
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil
}
func (g *GXDLMSIec8802LlcType2Setup) Load(reader *GXXmlReader) error {
	var err error
	g.TransmitWindowSizeK, err = reader.ReadElementContentAsUInt8("TransmitWindowSizeK", 1)
	if err != nil {
		return err
	}
	g.TransmitWindowSizeRW, err = reader.ReadElementContentAsUInt8("TransmitWindowSizeRW", 1)
	if err != nil {
		return err
	}
	g.MaximumOctetsPdu, err = reader.ReadElementContentAsUInt16("MaximumOctetsPdu", 128)
	if err != nil {
		return err
	}
	g.MaximumNumberTransmissions, err = reader.ReadElementContentAsUInt8("MaximumNumberTransmissions", 0)
	if err != nil {
		return err
	}
	g.AcknowledgementTimer, err = reader.ReadElementContentAsUInt16("AcknowledgementTimer", 0)
	if err != nil {
		return err
	}
	g.BitTimer, err = reader.ReadElementContentAsUInt16("BitTimer", 0)
	if err != nil {
		return err
	}
	g.RejectTimer, err = reader.ReadElementContentAsUInt16("RejectTimer", 0)
	if err != nil {
		return err
	}
	g.BusyStateTimer, err = reader.ReadElementContentAsUInt16("BusyStateTimer", 0)
	return err
}
func (g *GXDLMSIec8802LlcType2Setup) Save(writer *GXXmlWriter) error {
	if err := writer.WriteElementString("TransmitWindowSizeK", g.TransmitWindowSizeK); err != nil {
		return err
	}
	if err := writer.WriteElementString("TransmitWindowSizeRW", g.TransmitWindowSizeRW); err != nil {
		return err
	}
	if err := writer.WriteElementString("MaximumOctetsPdu", g.MaximumOctetsPdu); err != nil {
		return err
	}
	if err := writer.WriteElementString("MaximumNumberTransmissions", g.MaximumNumberTransmissions); err != nil {
		return err
	}
	if err := writer.WriteElementString("AcknowledgementTimer", g.AcknowledgementTimer); err != nil {
		return err
	}
	if err := writer.WriteElementString("BitTimer", g.BitTimer); err != nil {
		return err
	}
	if err := writer.WriteElementString("RejectTimer", g.RejectTimer); err != nil {
		return err
	}
	return writer.WriteElementString("BusyStateTimer", g.BusyStateTimer)
}
func (g *GXDLMSIec8802LlcType2Setup) PostLoad(reader *GXXmlReader) error { return nil }
func (g *GXDLMSIec8802LlcType2Setup) GetValues() []any {
	return []any{g.LogicalName(), g.TransmitWindowSizeK, g.TransmitWindowSizeRW, g.MaximumOctetsPdu, g.MaximumNumberTransmissions, g.AcknowledgementTimer, g.BitTimer, g.RejectTimer, g.BusyStateTimer}
}
func NewGXDLMSIec8802LlcType2Setup(ln string, sn int16) (*GXDLMSIec8802LlcType2Setup, error) {
	if err := ValidateLogicalName(ln); err != nil {
		return nil, err
	}
	return &GXDLMSIec8802LlcType2Setup{GXDLMSObject: GXDLMSObject{objectType: enums.ObjectTypeIec8802LlcType2Setup, logicalName: ln, ShortName: sn}, TransmitWindowSizeK: 1, TransmitWindowSizeRW: 1, MaximumOctetsPdu: 128}, nil
}
