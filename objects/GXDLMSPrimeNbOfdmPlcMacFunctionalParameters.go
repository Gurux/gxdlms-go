package objects

import (
	"github.com/Gurux/gxdlms-go/dlmserrors"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// GXDLMSPrimeNbOfdmPlcMacFunctionalParameters stores PRIME MAC functional parameters.
type GXDLMSPrimeNbOfdmPlcMacFunctionalParameters struct {
	GXDLMSObject
	LnId               int16
	LsId               uint8
	SId                uint8
	Sna                []byte
	State              enums.MacState
	ScpLength          int16
	NodeHierarchyLevel uint8
	BeaconSlotCount    uint8
	BeaconRxSlot       uint8
	BeaconTxSlot       uint8
	BeaconRxFrequency  uint8
	BeaconTxFrequency  uint8
	Capabilities       enums.MacCapabilities
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSPrimeNbOfdmPlcMacFunctionalParameters) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

func (g *GXDLMSPrimeNbOfdmPlcMacFunctionalParameters) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, nil
}
func (g *GXDLMSPrimeNbOfdmPlcMacFunctionalParameters) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	for i := 2; i <= 14; i++ {
		if all || g.CanRead(i) {
			attributes = append(attributes, i)
		}
	}
	return attributes
}
func (g *GXDLMSPrimeNbOfdmPlcMacFunctionalParameters) GetNames() []string {
	return []string{"Logical Name", "LnId", "LsId", "SId", "SNa", "State", "ScpLength", "NodeHierarchyLevel", "BeaconSlotCount", "BeaconRxSlot", "BeaconTxSlot", "BeaconRxFrequency", "BeaconTxFrequency", "Capabilities"}
}
func (g *GXDLMSPrimeNbOfdmPlcMacFunctionalParameters) GetMethodNames() []string { return []string{} }
func (g *GXDLMSPrimeNbOfdmPlcMacFunctionalParameters) GetAttributeCount() int   { return 14 }
func (g *GXDLMSPrimeNbOfdmPlcMacFunctionalParameters) GetMethodCount() int      { return 0 }
func (g *GXDLMSPrimeNbOfdmPlcMacFunctionalParameters) GetDataType(index int) (enums.DataType, error) {
	switch index {
	case 1:
		return enums.DataTypeOctetString, nil
	case 2, 7:
		return enums.DataTypeInt16, nil
	case 3, 4, 8, 9, 10, 11, 12, 13:
		return enums.DataTypeUint8, nil
	case 5:
		return enums.DataTypeOctetString, nil
	case 6:
		return enums.DataTypeEnum, nil
	case 14:
		return enums.DataTypeUint16, nil
	default:
		return 0, dlmserrors.ErrInvalidAttributeIndex
	}
}
func (g *GXDLMSPrimeNbOfdmPlcMacFunctionalParameters) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	switch e.Index {
	case 1:
		return helpers.LogicalNameToBytes(g.LogicalName())
	case 2:
		return g.LnId, nil
	case 3:
		return g.LsId, nil
	case 4:
		return g.SId, nil
	case 5:
		return g.Sna, nil
	case 6:
		return uint8(g.State), nil
	case 7:
		return g.ScpLength, nil
	case 8:
		return g.NodeHierarchyLevel, nil
	case 9:
		return g.BeaconSlotCount, nil
	case 10:
		return g.BeaconRxSlot, nil
	case 11:
		return g.BeaconTxSlot, nil
	case 12:
		return g.BeaconRxFrequency, nil
	case 13:
		return g.BeaconTxFrequency, nil
	case 14:
		return uint16(g.Capabilities), nil
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
		return nil, nil
	}
}
func (g *GXDLMSPrimeNbOfdmPlcMacFunctionalParameters) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	switch e.Index {
	case 1:
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return err
		}
		return g.SetLogicalName(ln)
	case 2:
		v, err := toInt16Value(e.Value)
		if err != nil {
			return err
		}
		g.LnId = v
	case 3:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.LsId = v
	case 4:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.SId = v
	case 5:
		if b, ok := e.Value.([]byte); ok {
			g.Sna = b
		} else if s, ok := e.Value.(string); ok {
			g.Sna = types.HexToBytes(s)
		}
	case 6:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.State = enums.MacState(v)
	case 7:
		v, err := toInt16Value(e.Value)
		if err != nil {
			return err
		}
		g.ScpLength = v
	case 8:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.NodeHierarchyLevel = v
	case 9:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.BeaconSlotCount = v
	case 10:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.BeaconRxSlot = v
	case 11:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.BeaconTxSlot = v
	case 12:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.BeaconRxFrequency = v
	case 13:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.BeaconTxFrequency = v
	case 14:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.Capabilities = enums.MacCapabilities(v)
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil
}
func (g *GXDLMSPrimeNbOfdmPlcMacFunctionalParameters) Load(reader *GXXmlReader) error {
	var err error
	g.LnId, err = reader.ReadElementContentAsInt16("LnId", 0)
	if err != nil {
		return err
	}
	g.LsId, err = reader.ReadElementContentAsUInt8("LsId", 0)
	if err != nil {
		return err
	}
	g.SId, err = reader.ReadElementContentAsUInt8("SId", 0)
	if err != nil {
		return err
	}
	if v, err := reader.ReadElementContentAsString("SNa", ""); err != nil {
		return err
	} else {
		g.Sna = types.HexToBytes(v)
	}
	if v, err := reader.ReadElementContentAsInt("State", 0); err != nil {
		return err
	} else {
		g.State = enums.MacState(v)
	}
	g.ScpLength, err = reader.ReadElementContentAsInt16("ScpLength", 0)
	if err != nil {
		return err
	}
	g.NodeHierarchyLevel, err = reader.ReadElementContentAsUInt8("NodeHierarchyLevel", 0)
	if err != nil {
		return err
	}
	g.BeaconSlotCount, err = reader.ReadElementContentAsUInt8("BeaconSlotCount", 0)
	if err != nil {
		return err
	}
	g.BeaconRxSlot, err = reader.ReadElementContentAsUInt8("BeaconRxSlot", 0)
	if err != nil {
		return err
	}
	g.BeaconTxSlot, err = reader.ReadElementContentAsUInt8("BeaconTxSlot", 0)
	if err != nil {
		return err
	}
	g.BeaconRxFrequency, err = reader.ReadElementContentAsUInt8("BeaconRxFrequency", 0)
	if err != nil {
		return err
	}
	g.BeaconTxFrequency, err = reader.ReadElementContentAsUInt8("BeaconTxFrequency", 0)
	if err != nil {
		return err
	}
	if v, err := reader.ReadElementContentAsInt("Capabilities", 0); err != nil {
		return err
	} else {
		g.Capabilities = enums.MacCapabilities(v)
	}
	return nil
}
func (g *GXDLMSPrimeNbOfdmPlcMacFunctionalParameters) Save(writer *GXXmlWriter) error {
	if err := writer.WriteElementString("LnId", g.LnId); err != nil {
		return err
	}
	if err := writer.WriteElementString("LsId", g.LsId); err != nil {
		return err
	}
	if err := writer.WriteElementString("SId", g.SId); err != nil {
		return err
	}
	if err := writer.WriteElementString("SNa", types.ToHex(g.Sna, false)); err != nil {
		return err
	}
	if err := writer.WriteElementString("State", int(g.State)); err != nil {
		return err
	}
	if err := writer.WriteElementString("ScpLength", g.ScpLength); err != nil {
		return err
	}
	if err := writer.WriteElementString("NodeHierarchyLevel", g.NodeHierarchyLevel); err != nil {
		return err
	}
	if err := writer.WriteElementString("BeaconSlotCount", g.BeaconSlotCount); err != nil {
		return err
	}
	if err := writer.WriteElementString("BeaconRxSlot", g.BeaconRxSlot); err != nil {
		return err
	}
	if err := writer.WriteElementString("BeaconTxSlot", g.BeaconTxSlot); err != nil {
		return err
	}
	if err := writer.WriteElementString("BeaconRxFrequency", g.BeaconRxFrequency); err != nil {
		return err
	}
	if err := writer.WriteElementString("BeaconTxFrequency", g.BeaconTxFrequency); err != nil {
		return err
	}
	return writer.WriteElementString("Capabilities", int(g.Capabilities))
}
func (g *GXDLMSPrimeNbOfdmPlcMacFunctionalParameters) PostLoad(reader *GXXmlReader) error { return nil }
func (g *GXDLMSPrimeNbOfdmPlcMacFunctionalParameters) GetValues() []any {
	return []any{g.LogicalName(), g.LnId, g.LsId, g.SId, g.Sna, g.State, g.ScpLength, g.NodeHierarchyLevel, g.BeaconSlotCount, g.BeaconRxSlot, g.BeaconTxSlot, g.BeaconRxFrequency, g.BeaconTxFrequency, g.Capabilities}
}
func NewGXDLMSPrimeNbOfdmPlcMacFunctionalParameters(ln string, sn int16) (*GXDLMSPrimeNbOfdmPlcMacFunctionalParameters, error) {
	if err := ValidateLogicalName(ln); err != nil {
		return nil, err
	}
	return &GXDLMSPrimeNbOfdmPlcMacFunctionalParameters{GXDLMSObject: GXDLMSObject{objectType: enums.ObjectTypePrimeNbOfdmPlcMacFunctionalParameters, logicalName: ln, ShortName: sn}}, nil
}
