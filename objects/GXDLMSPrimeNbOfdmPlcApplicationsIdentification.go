package objects

import (
	"github.com/Gurux/gxdlms-go/dlmserrors"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
)

// GXDLMSPrimeNbOfdmPlcApplicationsIdentification stores PRIME application identification.
type GXDLMSPrimeNbOfdmPlcApplicationsIdentification struct {
	GXDLMSObject
	FirmwareVersion string
	VendorID        uint16
	ProductID       uint16
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSPrimeNbOfdmPlcApplicationsIdentification) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

func (g *GXDLMSPrimeNbOfdmPlcApplicationsIdentification) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, nil
}
func (g *GXDLMSPrimeNbOfdmPlcApplicationsIdentification) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	for i := 2; i <= 4; i++ {
		if all || g.CanRead(i) {
			attributes = append(attributes, i)
		}
	}
	return attributes
}
func (g *GXDLMSPrimeNbOfdmPlcApplicationsIdentification) GetNames() []string {
	return []string{"Logical Name", "FirmwareVersion", "VendorId", "ProductId"}
}
func (g *GXDLMSPrimeNbOfdmPlcApplicationsIdentification) GetMethodNames() []string { return []string{} }
func (g *GXDLMSPrimeNbOfdmPlcApplicationsIdentification) GetAttributeCount() int   { return 4 }
func (g *GXDLMSPrimeNbOfdmPlcApplicationsIdentification) GetMethodCount() int      { return 0 }
func (g *GXDLMSPrimeNbOfdmPlcApplicationsIdentification) GetDataType(index int) (enums.DataType, error) {
	switch index {
	case 1, 2:
		return enums.DataTypeOctetString, nil
	case 3, 4:
		return enums.DataTypeUint16, nil
	default:
		return 0, dlmserrors.ErrInvalidAttributeIndex
	}
}
func (g *GXDLMSPrimeNbOfdmPlcApplicationsIdentification) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	switch e.Index {
	case 1:
		return helpers.LogicalNameToBytes(g.LogicalName())
	case 2:
		if g.FirmwareVersion == "" {
			return nil, nil
		}
		return []byte(g.FirmwareVersion), nil
	case 3:
		return g.VendorID, nil
	case 4:
		return g.ProductID, nil
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
		return nil, nil
	}
}
func (g *GXDLMSPrimeNbOfdmPlcApplicationsIdentification) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
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
		if b, ok := e.Value.([]byte); ok {
			g.FirmwareVersion = string(b)
		} else {
			g.FirmwareVersion = ""
		}
	case 3:
		g.VendorID, err = toUint16(e.Value)
	case 4:
		g.ProductID, err = toUint16(e.Value)
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return err
}
func (g *GXDLMSPrimeNbOfdmPlcApplicationsIdentification) Load(reader *GXXmlReader) error {
	var err error
	g.FirmwareVersion, err = reader.ReadElementContentAsString("FirmwareVersion", "")
	if err != nil {
		return err
	}
	g.VendorID, err = reader.ReadElementContentAsUInt16("VendorId", 0)
	if err != nil {
		return err
	}
	g.ProductID, err = reader.ReadElementContentAsUInt16("ProductId", 0)
	return err
}
func (g *GXDLMSPrimeNbOfdmPlcApplicationsIdentification) Save(writer *GXXmlWriter) error {
	if err := writer.WriteElementString("FirmwareVersion", g.FirmwareVersion); err != nil {
		return err
	}
	if err := writer.WriteElementString("VendorId", g.VendorID); err != nil {
		return err
	}
	return writer.WriteElementString("ProductId", g.ProductID)
}
func (g *GXDLMSPrimeNbOfdmPlcApplicationsIdentification) PostLoad(reader *GXXmlReader) error {
	return nil
}
func (g *GXDLMSPrimeNbOfdmPlcApplicationsIdentification) GetValues() []any {
	return []any{g.LogicalName(), g.FirmwareVersion, g.VendorID, g.ProductID}
}
func NewGXDLMSPrimeNbOfdmPlcApplicationsIdentification(ln string, sn int16) (*GXDLMSPrimeNbOfdmPlcApplicationsIdentification, error) {
	if err := ValidateLogicalName(ln); err != nil {
		return nil, err
	}
	return &GXDLMSPrimeNbOfdmPlcApplicationsIdentification{GXDLMSObject: GXDLMSObject{objectType: enums.ObjectTypePrimeNbOfdmPlcApplicationsIdentification, logicalName: ln, ShortName: sn}}, nil
}
