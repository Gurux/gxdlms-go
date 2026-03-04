package objects

import (
	"github.com/Gurux/gxdlms-go/dlmserrors"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
)

// GXDLMSIec8802LlcType3Setup represents ISO/IEC 8802-2 LLC type 3 setup.
type GXDLMSIec8802LlcType3Setup struct {
	GXDLMSObject
	MaximumOctetsACnPdu  uint16
	MaximumTransmissions uint8
	AcknowledgementTime  uint16
	ReceiveLifetime      uint16
	TransmitLifetime     uint16
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSIec8802LlcType3Setup) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

func (g *GXDLMSIec8802LlcType3Setup) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, nil
}
func (g *GXDLMSIec8802LlcType3Setup) GetAttributeIndexToRead(all bool) []int {
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
func (g *GXDLMSIec8802LlcType3Setup) GetNames() []string {
	return []string{"Logical Name", "MaximumOctetsACnPdu", "MaximumTransmissions", "AcknowledgementTime", "ReceiveLifetime", "TransmitLifetime"}
}
func (g *GXDLMSIec8802LlcType3Setup) GetMethodNames() []string { return []string{} }
func (g *GXDLMSIec8802LlcType3Setup) GetAttributeCount() int   { return 6 }
func (g *GXDLMSIec8802LlcType3Setup) GetMethodCount() int      { return 0 }
func (g *GXDLMSIec8802LlcType3Setup) GetDataType(index int) (enums.DataType, error) {
	switch index {
	case 1:
		return enums.DataTypeOctetString, nil
	case 3:
		return enums.DataTypeUint8, nil
	case 2, 4, 5, 6:
		return enums.DataTypeUint16, nil
	default:
		return 0, dlmserrors.ErrInvalidAttributeIndex
	}
}
func (g *GXDLMSIec8802LlcType3Setup) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	switch e.Index {
	case 1:
		return helpers.LogicalNameToBytes(g.LogicalName())
	case 2:
		return g.MaximumOctetsACnPdu, nil
	case 3:
		return g.MaximumTransmissions, nil
	case 4:
		return g.AcknowledgementTime, nil
	case 5:
		return g.ReceiveLifetime, nil
	case 6:
		return g.TransmitLifetime, nil
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
		return nil, nil
	}
}
func (g *GXDLMSIec8802LlcType3Setup) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
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
		g.MaximumOctetsACnPdu = uint16(v)
	case 3:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.MaximumTransmissions = v
	case 4:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.AcknowledgementTime = uint16(v)
	case 5:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.ReceiveLifetime = uint16(v)
	case 6:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.TransmitLifetime = uint16(v)
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil
}
func (g *GXDLMSIec8802LlcType3Setup) Load(reader *GXXmlReader) error {
	var err error
	g.MaximumOctetsACnPdu, err = reader.ReadElementContentAsUInt16("MaximumOctetsACnPdu", 0)
	if err != nil {
		return err
	}
	g.MaximumTransmissions, err = reader.ReadElementContentAsUInt8("MaximumTransmissions", 0)
	if err != nil {
		return err
	}
	g.AcknowledgementTime, err = reader.ReadElementContentAsUInt16("AcknowledgementTime", 0)
	if err != nil {
		return err
	}
	g.ReceiveLifetime, err = reader.ReadElementContentAsUInt16("ReceiveLifetime", 0)
	if err != nil {
		return err
	}
	g.TransmitLifetime, err = reader.ReadElementContentAsUInt16("TransmitLifetime", 0)
	return err
}
func (g *GXDLMSIec8802LlcType3Setup) Save(writer *GXXmlWriter) error {
	if err := writer.WriteElementString("MaximumOctetsACnPdu", g.MaximumOctetsACnPdu); err != nil {
		return err
	}
	if err := writer.WriteElementString("MaximumTransmissions", g.MaximumTransmissions); err != nil {
		return err
	}
	if err := writer.WriteElementString("AcknowledgementTime", g.AcknowledgementTime); err != nil {
		return err
	}
	if err := writer.WriteElementString("ReceiveLifetime", g.ReceiveLifetime); err != nil {
		return err
	}
	return writer.WriteElementString("TransmitLifetime", g.TransmitLifetime)
}
func (g *GXDLMSIec8802LlcType3Setup) PostLoad(reader *GXXmlReader) error { return nil }
func (g *GXDLMSIec8802LlcType3Setup) GetValues() []any {
	return []any{g.LogicalName(), g.MaximumOctetsACnPdu, g.MaximumTransmissions, g.AcknowledgementTime, g.ReceiveLifetime, g.TransmitLifetime}
}
func NewGXDLMSIec8802LlcType3Setup(ln string, sn int16) (*GXDLMSIec8802LlcType3Setup, error) {
	if err := ValidateLogicalName(ln); err != nil {
		return nil, err
	}
	return &GXDLMSIec8802LlcType3Setup{GXDLMSObject: GXDLMSObject{objectType: enums.ObjectTypeIec8802LlcType3Setup, logicalName: ln, ShortName: sn}}, nil
}
