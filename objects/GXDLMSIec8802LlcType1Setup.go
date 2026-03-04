package objects

import (
	"github.com/Gurux/gxdlms-go/dlmserrors"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// GXDLMSIec8802LlcType1Setup represents ISO/IEC 8802-2 LLC type 1 setup.
type GXDLMSIec8802LlcType1Setup struct {
	GXDLMSObject
	MaximumOctetsUiPdu uint16
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSIec8802LlcType1Setup) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

func (g *GXDLMSIec8802LlcType1Setup) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, nil
}
func (g *GXDLMSIec8802LlcType1Setup) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	if all || g.CanRead(2) {
		attributes = append(attributes, 2)
	}
	return attributes
}
func (g *GXDLMSIec8802LlcType1Setup) GetNames() []string {
	return []string{"Logical Name", "MaximumOctetsUiPdu"}
}
func (g *GXDLMSIec8802LlcType1Setup) GetMethodNames() []string { return []string{} }
func (g *GXDLMSIec8802LlcType1Setup) GetAttributeCount() int   { return 2 }
func (g *GXDLMSIec8802LlcType1Setup) GetMethodCount() int      { return 0 }
func (g *GXDLMSIec8802LlcType1Setup) GetDataType(index int) (enums.DataType, error) {
	if index == 1 {
		return enums.DataTypeOctetString, nil
	}
	if index == 2 {
		return enums.DataTypeUint16, nil
	}
	return 0, dlmserrors.ErrInvalidAttributeIndex
}
func (g *GXDLMSIec8802LlcType1Setup) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	if e.Index == 1 {
		return helpers.LogicalNameToBytes(g.LogicalName())
	}
	if e.Index == 2 {
		return g.MaximumOctetsUiPdu, nil
	}
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, nil
}
func (g *GXDLMSIec8802LlcType1Setup) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	if e.Index == 1 {
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return err
		}
		return g.SetLogicalName(ln)
	}
	if e.Index == 2 {
		g.MaximumOctetsUiPdu, _ = toUint16(e.Value)
		return nil
	}
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil
}
func (g *GXDLMSIec8802LlcType1Setup) Load(reader *GXXmlReader) error {
	v, err := reader.ReadElementContentAsUInt16("MaximumOctetsUiPdu", 128)
	if err != nil {
		return err
	}
	g.MaximumOctetsUiPdu = v
	return nil
}
func (g *GXDLMSIec8802LlcType1Setup) Save(writer *GXXmlWriter) error {
	return writer.WriteElementString("MaximumOctetsUiPdu", g.MaximumOctetsUiPdu)
}
func (g *GXDLMSIec8802LlcType1Setup) PostLoad(reader *GXXmlReader) error { return nil }
func (g *GXDLMSIec8802LlcType1Setup) GetValues() []any {
	return []any{g.LogicalName(), g.MaximumOctetsUiPdu}
}
func NewGXDLMSIec8802LlcType1Setup(ln string, sn int16) (*GXDLMSIec8802LlcType1Setup, error) {
	if err := ValidateLogicalName(ln); err != nil {
		return nil, err
	}
	return &GXDLMSIec8802LlcType1Setup{GXDLMSObject: GXDLMSObject{objectType: enums.ObjectTypeIec8802LlcType1Setup, logicalName: ln, ShortName: sn}, MaximumOctetsUiPdu: 128}, nil
}

// keep import used when no fields are set by parser variants.
var _ = types.GXArray{}
