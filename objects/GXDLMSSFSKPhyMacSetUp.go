package objects

import (
	"fmt"

	"github.com/Gurux/gxdlms-go/dlmserrors"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// GXDLMSSFSKPhyMacSetup contains S-FSK PHY/MAC setup parameters.
type GXDLMSSFSKPhyMacSetup struct {
	GXDLMSObject
	InitiatorElectricalPhase enums.InitiatorElectricalPhase
	DeltaElectricalPhase     enums.DeltaElectricalPhase
	MaxReceivingGain         uint8
	MaxTransmittingGain      uint8
	SearchInitiatorThreshold uint8
	MarkFrequency            uint32
	SpaceFrequency           uint32
	MacAddress               uint16
	MacGroupAddresses        []uint16
	Repeater                 enums.Repeater
	RepeaterStatus           bool
	MinDeltaCredit           uint8
	InitiatorMacAddress      uint16
	SynchronizationLocked    bool
	TransmissionSpeed        enums.BaudRate
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSSFSKPhyMacSetup) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

func (g *GXDLMSSFSKPhyMacSetup) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, nil
}

func (g *GXDLMSSFSKPhyMacSetup) GetAttributeIndexToRead(all bool) []int {
	var a []int
	if all || g.LogicalName() == "" {
		a = append(a, 1)
	}
	if all || g.CanRead(2) {
		a = append(a, 2)
	}
	a = append(a, 3)
	if all || g.CanRead(4) {
		a = append(a, 4)
	}
	if all || g.CanRead(5) {
		a = append(a, 5)
	}
	if all || g.CanRead(6) {
		a = append(a, 6)
	}
	if all || g.CanRead(7) {
		a = append(a, 7)
	}
	a = append(a, 8)
	if all || g.CanRead(9) {
		a = append(a, 9)
	}
	if all || g.CanRead(10) {
		a = append(a, 10)
	}
	a = append(a, 11, 12, 13, 14)
	if all || g.CanRead(15) {
		a = append(a, 15)
	}
	return a
}

func (g *GXDLMSSFSKPhyMacSetup) GetNames() []string {
	return []string{
		"Logical Name",
		"InitiatorElectricalPhase",
		"DeltaElectricalPhase",
		"MaxReceivingGain",
		"MaxTransmittingGain",
		"SearchInitiatorThreshold",
		"Frequency",
		"MacAddress",
		"MacGroupAddresses",
		"Repeater",
		"RepeaterStatus",
		"MinDeltaCredit",
		"InitiatorMacAddress",
		"SynchronizationLocked",
		"TransmissionSpeed",
	}
}

func (g *GXDLMSSFSKPhyMacSetup) GetMethodNames() []string { return []string{} }

func (g *GXDLMSSFSKPhyMacSetup) GetAttributeCount() int { return 15 }

func (g *GXDLMSSFSKPhyMacSetup) GetMethodCount() int { return 0 }

func (g *GXDLMSSFSKPhyMacSetup) GetDataType(index int) (enums.DataType, error) {
	switch index {
	case 1:
		return enums.DataTypeOctetString, nil
	case 2, 3, 10, 15:
		return enums.DataTypeEnum, nil
	case 4, 5, 6, 12:
		return enums.DataTypeUint8, nil
	case 7:
		return enums.DataTypeStructure, nil
	case 8, 13:
		return enums.DataTypeUint16, nil
	case 9:
		return enums.DataTypeArray, nil
	case 11, 14:
		return enums.DataTypeBoolean, nil
	default:
		return 0, dlmserrors.ErrInvalidAttributeIndex
	}
}

func (g *GXDLMSSFSKPhyMacSetup) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	switch e.Index {
	case 1:
		return helpers.LogicalNameToBytes(g.LogicalName())
	case 2:
		return uint8(g.InitiatorElectricalPhase), nil
	case 3:
		return uint8(g.DeltaElectricalPhase), nil
	case 4:
		return g.MaxReceivingGain, nil
	case 5:
		return g.MaxTransmittingGain, nil
	case 6:
		return g.SearchInitiatorThreshold, nil
	case 7:
		bb := types.NewGXByteBuffer()
		if err := bb.SetUint8(uint8(enums.DataTypeStructure)); err != nil {
			return nil, err
		}
		if err := bb.SetUint8(2); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeUint32, g.MarkFrequency); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeUint32, g.SpaceFrequency); err != nil {
			return nil, err
		}
		return bb.Array(), nil
	case 8:
		return g.MacAddress, nil
	case 9:
		bb := types.NewGXByteBuffer()
		if err := bb.SetUint8(uint8(enums.DataTypeArray)); err != nil {
			return nil, err
		}
		types.SetObjectCount(len(g.MacGroupAddresses), bb)
		for _, it := range g.MacGroupAddresses {
			if err := internal.SetData(settings, bb, enums.DataTypeUint16, it); err != nil {
				return nil, err
			}
		}
		return bb.Array(), nil
	case 10:
		return uint8(g.Repeater), nil
	case 11:
		return g.RepeaterStatus, nil
	case 12:
		return g.MinDeltaCredit, nil
	case 13:
		return g.InitiatorMacAddress, nil
	case 14:
		return g.SynchronizationLocked, nil
	case 15:
		return uint8(g.TransmissionSpeed), nil
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
		return nil, nil
	}
}

func (g *GXDLMSSFSKPhyMacSetup) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
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
		g.InitiatorElectricalPhase = enums.InitiatorElectricalPhase(v)
	case 3:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.DeltaElectricalPhase = enums.DeltaElectricalPhase(v)
	case 4:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.MaxReceivingGain = v
	case 5:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.MaxTransmittingGain = v
	case 6:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.SearchInitiatorThreshold = v
	case 7:
		if e.Value == nil {
			g.MarkFrequency = 0
			g.SpaceFrequency = 0
			return nil
		}
		tmp := e.Value.(types.GXStructure)
		if len(tmp) < 2 {
			return fmt.Errorf("invalid frequency structure: %T", e.Value)
		}
		mark, err := toUint32(tmp[0])
		if err != nil {
			return err
		}
		space, err := toUint32(tmp[1])
		if err != nil {
			return err
		}
		g.MarkFrequency = mark
		g.SpaceFrequency = space
	case 8:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.MacAddress = uint16(v)
	case 9:
		if e.Value == nil {
			g.MacGroupAddresses = nil
			return nil
		}
		tmp, ok := e.Value.(types.GXStructure)
		if !ok {
			return fmt.Errorf("invalid mac group addresses: %T", e.Value)
		}
		g.MacGroupAddresses = make([]uint16, 0, len(tmp))
		for _, it := range tmp {
			v, err := toUint32(it)
			if err != nil {
				return err
			}
			g.MacGroupAddresses = append(g.MacGroupAddresses, uint16(v))
		}
	case 10:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.Repeater = enums.Repeater(v)
	case 11:
		v, err := toBool(e.Value)
		if err != nil {
			return err
		}
		g.RepeaterStatus = v
	case 12:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.MinDeltaCredit = v
	case 13:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.InitiatorMacAddress = uint16(v)
	case 14:
		v, err := toBool(e.Value)
		if err != nil {
			return err
		}
		g.SynchronizationLocked = v
	case 15:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.TransmissionSpeed = enums.BaudRate(v)
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil
}

func (g *GXDLMSSFSKPhyMacSetup) Load(reader *GXXmlReader) error {
	var ret int
	var err error
	ret, err = reader.ReadElementContentAsInt("InitiatorElectricalPhase", 0)
	if err != nil {
		return err
	}
	g.InitiatorElectricalPhase = enums.InitiatorElectricalPhase(ret)
	ret, err = reader.ReadElementContentAsInt("DeltaElectricalPhase", 0)
	if err != nil {
		return err
	}
	g.DeltaElectricalPhase = enums.DeltaElectricalPhase(ret)
	g.MaxReceivingGain, err = reader.ReadElementContentAsUInt8("MaxReceivingGain", 0)
	if err != nil {
		return err
	}
	g.MaxTransmittingGain, err = reader.ReadElementContentAsUInt8("MaxTransmittingGain", 0)
	if err != nil {
		return err
	}
	g.SearchInitiatorThreshold, err = reader.ReadElementContentAsUInt8("SearchInitiatorThreshold", 0)
	if err != nil {
		return err
	}
	g.MarkFrequency, err = reader.ReadElementContentAsUInt32("MarkFrequency", 0)
	if err != nil {
		return err
	}
	g.SpaceFrequency, err = reader.ReadElementContentAsUInt32("SpaceFrequency", 0)
	if err != nil {
		return err
	}
	g.MacAddress, err = reader.ReadElementContentAsUInt16("MacAddress", 0)
	if err != nil {
		return err
	}
	g.MacGroupAddresses = g.MacGroupAddresses[:0]
	ok, err := reader.IsStartElementNamed("MacGroupAddresses", true)
	if err != nil {
		return err
	}
	if ok {
		for {
			ok, err = reader.IsStartElementNamed("Value", false)
			if err != nil {
				return err
			}
			if !ok {
				break
			}
			v, err := reader.ReadElementContentAsUInt16("Value", 0)
			if err != nil {
				return err
			}
			g.MacGroupAddresses = append(g.MacGroupAddresses, v)
		}
		if err = reader.ReadEndElement("MacGroupAddresses"); err != nil {
			return err
		}
	}
	ret, err = reader.ReadElementContentAsInt("Repeater", 0)
	if err != nil {
		return err
	}
	g.Repeater = enums.Repeater(ret)
	g.RepeaterStatus, err = reader.ReadElementContentAsBool("RepeaterStatus", false)
	if err != nil {
		return err
	}
	g.MinDeltaCredit, err = reader.ReadElementContentAsUInt8("MinDeltaCredit", 0)
	if err != nil {
		return err
	}
	g.InitiatorMacAddress, err = reader.ReadElementContentAsUInt16("InitiatorMacAddress", 0)
	if err != nil {
		return err
	}
	g.SynchronizationLocked, err = reader.ReadElementContentAsBool("SynchronizationLocked", false)
	if err != nil {
		return err
	}
	ret, err = reader.ReadElementContentAsInt("TransmissionSpeed", 0)
	if err != nil {
		return err
	}
	g.TransmissionSpeed = enums.BaudRate(ret)
	return nil
}

func (g *GXDLMSSFSKPhyMacSetup) Save(writer *GXXmlWriter) error {
	if err := writer.WriteElementString("InitiatorElectricalPhase", int(g.InitiatorElectricalPhase)); err != nil {
		return err
	}
	if err := writer.WriteElementString("DeltaElectricalPhase", int(g.DeltaElectricalPhase)); err != nil {
		return err
	}
	if err := writer.WriteElementString("MaxReceivingGain", g.MaxReceivingGain); err != nil {
		return err
	}
	if err := writer.WriteElementString("MaxTransmittingGain", g.MaxTransmittingGain); err != nil {
		return err
	}
	if err := writer.WriteElementString("SearchInitiatorThreshold", g.SearchInitiatorThreshold); err != nil {
		return err
	}
	if err := writer.WriteElementString("MarkFrequency", g.MarkFrequency); err != nil {
		return err
	}
	if err := writer.WriteElementString("SpaceFrequency", g.SpaceFrequency); err != nil {
		return err
	}
	if err := writer.WriteElementString("MacAddress", g.MacAddress); err != nil {
		return err
	}
	if err := writer.WriteStartElement("MacGroupAddresses"); err != nil {
		return err
	}
	for _, it := range g.MacGroupAddresses {
		if err := writer.WriteElementString("Value", it); err != nil {
			return err
		}
	}
	if err := writer.WriteEndElement(); err != nil {
		return err
	}
	if err := writer.WriteElementString("Repeater", int(g.Repeater)); err != nil {
		return err
	}
	if err := writer.WriteElementStringBool("RepeaterStatus", g.RepeaterStatus); err != nil {
		return err
	}
	if err := writer.WriteElementString("MinDeltaCredit", g.MinDeltaCredit); err != nil {
		return err
	}
	if err := writer.WriteElementString("InitiatorMacAddress", g.InitiatorMacAddress); err != nil {
		return err
	}
	if err := writer.WriteElementStringBool("SynchronizationLocked", g.SynchronizationLocked); err != nil {
		return err
	}
	return writer.WriteElementString("TransmissionSpeed", int(g.TransmissionSpeed))
}

func (g *GXDLMSSFSKPhyMacSetup) PostLoad(reader *GXXmlReader) error { return nil }

func (g *GXDLMSSFSKPhyMacSetup) GetValues() []any {
	return []any{
		g.LogicalName(),
		g.InitiatorElectricalPhase,
		g.DeltaElectricalPhase,
		g.MaxReceivingGain,
		g.MaxTransmittingGain,
		g.SearchInitiatorThreshold,
		[]any{g.MarkFrequency, g.SpaceFrequency},
		g.MacAddress,
		g.MacGroupAddresses,
		g.Repeater,
		g.RepeaterStatus,
		g.MinDeltaCredit,
		g.InitiatorMacAddress,
		g.SynchronizationLocked,
		g.TransmissionSpeed,
	}
}

// NewGXDLMSSFSKPhyMacSetup creates a new s-FSK PHY/MAC setup object instance.
//
// The function validates `ln` before creating the object.
// `ln` is the Logical Name and `sn` is the Short Name of the object.
func NewGXDLMSSFSKPhyMacSetup(ln string, sn int16) (*GXDLMSSFSKPhyMacSetup, error) {
	if err := ValidateLogicalName(ln); err != nil {
		return nil, err
	}
	return &GXDLMSSFSKPhyMacSetup{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeSFSKPhyMacSetUp,
			logicalName: ln,
			ShortName:   sn,
		},
		MacGroupAddresses: make([]uint16, 0),
	}, nil
}
