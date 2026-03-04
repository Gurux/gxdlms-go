package objects

import (
	"github.com/Gurux/gxdlms-go/dlmserrors"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
)

// GXDLMSPrimeNbOfdmPlcMacSetup represents PRIME NB OFDM PLC MAC setup.
type GXDLMSPrimeNbOfdmPlcMacSetup struct {
	GXDLMSObject
	MacMinSwitchSearchTime  uint8
	MacMaxPromotionPdu      uint8
	MacPromotionPduTxPeriod uint8
	MacBeaconsPerFrame      uint8
	MacScpMaxTxAttempts     uint8
	MacCtlReTxTimer         uint8
	MacMaxCtlReTx           uint8
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSPrimeNbOfdmPlcMacSetup) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

func (g *GXDLMSPrimeNbOfdmPlcMacSetup) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, nil
}
func (g *GXDLMSPrimeNbOfdmPlcMacSetup) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	for i := 2; i <= 8; i++ {
		if all || g.CanRead(i) {
			attributes = append(attributes, i)
		}
	}
	return attributes
}
func (g *GXDLMSPrimeNbOfdmPlcMacSetup) GetNames() []string {
	return []string{"Logical Name", "MacMinSwitchSearchTime", "MacMaxPromotionPdu", "MacPromotionPduTxPeriod", "MacBeaconsPerFrame", "MacScpMaxTxAttempts", "MacCtlReTxTimer", "MacMaxCtlReTx"}
}
func (g *GXDLMSPrimeNbOfdmPlcMacSetup) GetMethodNames() []string { return []string{} }
func (g *GXDLMSPrimeNbOfdmPlcMacSetup) GetAttributeCount() int   { return 8 }
func (g *GXDLMSPrimeNbOfdmPlcMacSetup) GetMethodCount() int      { return 0 }
func (g *GXDLMSPrimeNbOfdmPlcMacSetup) GetDataType(index int) (enums.DataType, error) {
	if index == 1 {
		return enums.DataTypeOctetString, nil
	}
	if index >= 2 && index <= 8 {
		return enums.DataTypeUint8, nil
	}
	return 0, dlmserrors.ErrInvalidAttributeIndex
}
func (g *GXDLMSPrimeNbOfdmPlcMacSetup) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	switch e.Index {
	case 1:
		return helpers.LogicalNameToBytes(g.LogicalName())
	case 2:
		return g.MacMinSwitchSearchTime, nil
	case 3:
		return g.MacMaxPromotionPdu, nil
	case 4:
		return g.MacPromotionPduTxPeriod, nil
	case 5:
		return g.MacBeaconsPerFrame, nil
	case 6:
		return g.MacScpMaxTxAttempts, nil
	case 7:
		return g.MacCtlReTxTimer, nil
	case 8:
		return g.MacMaxCtlReTx, nil
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
		return nil, nil
	}
}
func (g *GXDLMSPrimeNbOfdmPlcMacSetup) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
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
		g.MacMinSwitchSearchTime = v
	case 3:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.MacMaxPromotionPdu = v
	case 4:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.MacPromotionPduTxPeriod = v
	case 5:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.MacBeaconsPerFrame = v
	case 6:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.MacScpMaxTxAttempts = v
	case 7:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.MacCtlReTxTimer = v
	case 8:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.MacMaxCtlReTx = v
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil
}
func (g *GXDLMSPrimeNbOfdmPlcMacSetup) Load(reader *GXXmlReader) error {
	var err error
	g.MacMinSwitchSearchTime, err = reader.ReadElementContentAsUInt8("MacMinSwitchSearchTime", 0)
	if err != nil {
		return err
	}
	g.MacMaxPromotionPdu, err = reader.ReadElementContentAsUInt8("MacMaxPromotionPdu", 0)
	if err != nil {
		return err
	}
	g.MacPromotionPduTxPeriod, err = reader.ReadElementContentAsUInt8("MacPromotionPduTxPeriod", 0)
	if err != nil {
		return err
	}
	g.MacBeaconsPerFrame, err = reader.ReadElementContentAsUInt8("MacBeaconsPerFrame", 0)
	if err != nil {
		return err
	}
	g.MacScpMaxTxAttempts, err = reader.ReadElementContentAsUInt8("MacScpMaxTxAttempts", 0)
	if err != nil {
		return err
	}
	g.MacCtlReTxTimer, err = reader.ReadElementContentAsUInt8("MacCtlReTxTimer", 0)
	if err != nil {
		return err
	}
	g.MacMaxCtlReTx, err = reader.ReadElementContentAsUInt8("MacMaxCtlReTx", 0)
	return err
}
func (g *GXDLMSPrimeNbOfdmPlcMacSetup) Save(writer *GXXmlWriter) error {
	if err := writer.WriteElementString("MacMinSwitchSearchTime", g.MacMinSwitchSearchTime); err != nil {
		return err
	}
	if err := writer.WriteElementString("MacMaxPromotionPdu", g.MacMaxPromotionPdu); err != nil {
		return err
	}
	if err := writer.WriteElementString("MacPromotionPduTxPeriod", g.MacPromotionPduTxPeriod); err != nil {
		return err
	}
	if err := writer.WriteElementString("MacBeaconsPerFrame", g.MacBeaconsPerFrame); err != nil {
		return err
	}
	if err := writer.WriteElementString("MacScpMaxTxAttempts", g.MacScpMaxTxAttempts); err != nil {
		return err
	}
	if err := writer.WriteElementString("MacCtlReTxTimer", g.MacCtlReTxTimer); err != nil {
		return err
	}
	return writer.WriteElementString("MacMaxCtlReTx", g.MacMaxCtlReTx)
}
func (g *GXDLMSPrimeNbOfdmPlcMacSetup) PostLoad(reader *GXXmlReader) error { return nil }
func (g *GXDLMSPrimeNbOfdmPlcMacSetup) GetValues() []any {
	return []any{g.LogicalName(), g.MacMinSwitchSearchTime, g.MacMaxPromotionPdu, g.MacPromotionPduTxPeriod, g.MacBeaconsPerFrame, g.MacScpMaxTxAttempts, g.MacCtlReTxTimer, g.MacMaxCtlReTx}
}
func NewGXDLMSPrimeNbOfdmPlcMacSetup(ln string, sn int16) (*GXDLMSPrimeNbOfdmPlcMacSetup, error) {
	if err := ValidateLogicalName(ln); err != nil {
		return nil, err
	}
	return &GXDLMSPrimeNbOfdmPlcMacSetup{GXDLMSObject: GXDLMSObject{objectType: enums.ObjectTypePrimeNbOfdmPlcMacSetup, logicalName: ln, ShortName: sn}}, nil
}
