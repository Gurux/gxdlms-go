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

// GXDLMSSFSKMacCounters contains s-FSK MAC counters.
type GXDLMSSFSKMacCounters struct {
	GXDLMSObject
	SynchronizationRegister              []types.GXKeyValuePair[uint16, uint32]
	PhysicalLayerDesynchronization       uint32
	TimeOutNotAddressedDesynchronization uint32
	TimeOutFrameNotOkDesynchronization   uint32
	WriteRequestDesynchronization        uint32
	WrongInitiatorDesynchronization      uint32
	BroadcastFramesCounter               []types.GXKeyValuePair[uint16, uint32]
	RepetitionsCounter                   uint32
	TransmissionsCounter                 uint32
	CrcOkFramesCounter                   uint32
	CrcNOkFramesCounter                  uint32
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSSFSKMacCounters) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

func (g *GXDLMSSFSKMacCounters) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, nil
}

func (g *GXDLMSSFSKMacCounters) GetAttributeIndexToRead(all bool) []int {
	var a []int
	if all || g.LogicalName() == "" {
		a = append(a, 1)
	}
	for i := 2; i <= 8; i++ {
		if all || g.CanRead(i) {
			a = append(a, i)
		}
	}
	return a
}

func (g *GXDLMSSFSKMacCounters) GetNames() []string {
	return []string{
		"Logical Name",
		"SynchronizationRegister",
		"Desynchronization listing",
		"BroadcastFramesCounter",
		"RepetitionsCounter",
		"TransmissionsCounter",
		"CrcOkFramesCounter",
		"CrcNOkFramesCounter",
	}
}

func (g *GXDLMSSFSKMacCounters) GetMethodNames() []string { return []string{"Reset"} }

func (g *GXDLMSSFSKMacCounters) GetAttributeCount() int { return 8 }

func (g *GXDLMSSFSKMacCounters) GetMethodCount() int { return 1 }

func (g *GXDLMSSFSKMacCounters) GetDataType(index int) (enums.DataType, error) {
	switch index {
	case 1:
		return enums.DataTypeOctetString, nil
	case 2:
		return enums.DataTypeArray, nil
	case 3:
		return enums.DataTypeStructure, nil
	case 4:
		return enums.DataTypeArray, nil
	case 5, 6, 7, 8:
		return enums.DataTypeUint32, nil
	default:
		return 0, dlmserrors.ErrInvalidAttributeIndex
	}
}

func (g *GXDLMSSFSKMacCounters) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	switch e.Index {
	case 1:
		return helpers.LogicalNameToBytes(g.LogicalName())
	case 2:
		return g.encodeCounterPairs(settings, g.SynchronizationRegister)
	case 3:
		bb := types.NewGXByteBuffer()
		if err := bb.SetUint8(uint8(enums.DataTypeStructure)); err != nil {
			return nil, err
		}
		if err := bb.SetUint8(5); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeUint32, g.PhysicalLayerDesynchronization); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeUint32, g.TimeOutNotAddressedDesynchronization); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeUint32, g.TimeOutFrameNotOkDesynchronization); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeUint32, g.WriteRequestDesynchronization); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeUint32, g.WrongInitiatorDesynchronization); err != nil {
			return nil, err
		}
		return bb.Array(), nil
	case 4:
		return g.encodeCounterPairs(settings, g.BroadcastFramesCounter)
	case 5:
		return g.RepetitionsCounter, nil
	case 6:
		return g.TransmissionsCounter, nil
	case 7:
		return g.CrcOkFramesCounter, nil
	case 8:
		return g.CrcNOkFramesCounter, nil
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
		return nil, nil
	}
}

func (g *GXDLMSSFSKMacCounters) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	switch e.Index {
	case 1:
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return err
		}
		return g.SetLogicalName(ln)
	case 2:
		pairs, err := sfskParseCounterPairs(e.Value)
		if err != nil {
			return err
		}
		g.SynchronizationRegister = pairs
	case 3:
		if e.Value == nil {
			g.PhysicalLayerDesynchronization = 0
			g.TimeOutNotAddressedDesynchronization = 0
			g.TimeOutFrameNotOkDesynchronization = 0
			g.WriteRequestDesynchronization = 0
			g.WrongInitiatorDesynchronization = 0
			return nil
		}
		tmp, ok := e.Value.(types.GXStructure)
		if !ok || len(tmp) < 5 {
			return fmt.Errorf("invalid desynchronization listing: %T", e.Value)
		}
		var err error
		g.PhysicalLayerDesynchronization, err = toUint32(tmp[0])
		if err != nil {
			return err
		}
		g.TimeOutNotAddressedDesynchronization, err = toUint32(tmp[1])
		if err != nil {
			return err
		}
		g.TimeOutFrameNotOkDesynchronization, err = toUint32(tmp[2])
		if err != nil {
			return err
		}
		g.WriteRequestDesynchronization, err = toUint32(tmp[3])
		if err != nil {
			return err
		}
		g.WrongInitiatorDesynchronization, err = toUint32(tmp[4])
		if err != nil {
			return err
		}
	case 4:
		pairs, err := sfskParseCounterPairs(e.Value)
		if err != nil {
			return err
		}
		g.BroadcastFramesCounter = pairs
	case 5:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.RepetitionsCounter = v
	case 6:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.TransmissionsCounter = v
	case 7:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.CrcOkFramesCounter = v
	case 8:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.CrcNOkFramesCounter = v
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil
}

func (g *GXDLMSSFSKMacCounters) Load(reader *GXXmlReader) error {
	g.SynchronizationRegister = g.SynchronizationRegister[:0]
	ok, err := reader.IsStartElementNamed("SynchronizationRegisters", true)
	if err != nil {
		return err
	}
	if ok {
		for {
			ok, err = reader.IsStartElementNamed("Item", true)
			if err != nil {
				return err
			}
			if !ok {
				break
			}
			k, err := reader.ReadElementContentAsUInt16("Key", 0)
			if err != nil {
				return err
			}
			v, err := reader.ReadElementContentAsUInt32("Value", 0)
			if err != nil {
				return err
			}
			g.SynchronizationRegister = append(g.SynchronizationRegister, *types.NewGXKeyValuePair(k, v))
		}
		if err = reader.ReadEndElement("SynchronizationRegisters"); err != nil {
			return err
		}
	}
	g.PhysicalLayerDesynchronization, err = reader.ReadElementContentAsUInt32("PhysicalLayerDesynchronization", 0)
	if err != nil {
		return err
	}
	g.TimeOutNotAddressedDesynchronization, err = reader.ReadElementContentAsUInt32("TimeOutNotAddressedDesynchronization", 0)
	if err != nil {
		return err
	}
	g.TimeOutFrameNotOkDesynchronization, err = reader.ReadElementContentAsUInt32("TimeOutFrameNotOkDesynchronization", 0)
	if err != nil {
		return err
	}
	g.WriteRequestDesynchronization, err = reader.ReadElementContentAsUInt32("WriteRequestDesynchronization", 0)
	if err != nil {
		return err
	}
	g.WrongInitiatorDesynchronization, err = reader.ReadElementContentAsUInt32("WrongInitiatorDesynchronization", 0)
	if err != nil {
		return err
	}

	g.BroadcastFramesCounter = g.BroadcastFramesCounter[:0]
	ok, err = reader.IsStartElementNamed("BroadcastFramesCounters", true)
	if err != nil {
		return err
	}
	if ok {
		for {
			ok, err = reader.IsStartElementNamed("Item", true)
			if err != nil {
				return err
			}
			if !ok {
				break
			}
			k, err := reader.ReadElementContentAsUInt16("Key", 0)
			if err != nil {
				return err
			}
			v, err := reader.ReadElementContentAsUInt32("Value", 0)
			if err != nil {
				return err
			}
			g.BroadcastFramesCounter = append(g.BroadcastFramesCounter, *types.NewGXKeyValuePair(k, v))
		}
		if err = reader.ReadEndElement("BroadcastFramesCounters"); err != nil {
			return err
		}
	}
	g.RepetitionsCounter, err = reader.ReadElementContentAsUInt32("RepetitionsCounter", 0)
	if err != nil {
		return err
	}
	g.TransmissionsCounter, err = reader.ReadElementContentAsUInt32("TransmissionsCounter", 0)
	if err != nil {
		return err
	}
	g.CrcOkFramesCounter, err = reader.ReadElementContentAsUInt32("CrcOkFramesCounter", 0)
	if err != nil {
		return err
	}
	g.CrcNOkFramesCounter, err = reader.ReadElementContentAsUInt32("CrcNOkFramesCounter", 0)
	return err
}

func (g *GXDLMSSFSKMacCounters) Save(writer *GXXmlWriter) error {
	if len(g.SynchronizationRegister) != 0 {
		if err := writer.WriteStartElement("SynchronizationRegisters"); err != nil {
			return err
		}
		for _, it := range g.SynchronizationRegister {
			if err := writer.WriteStartElement("Item"); err != nil {
				return err
			}
			if err := writer.WriteElementString("Key", it.Key); err != nil {
				return err
			}
			if err := writer.WriteElementString("Value", it.Value); err != nil {
				return err
			}
			if err := writer.WriteEndElement(); err != nil {
				return err
			}
		}
		if err := writer.WriteEndElement(); err != nil {
			return err
		}
	}
	if err := writer.WriteElementString("PhysicalLayerDesynchronization", g.PhysicalLayerDesynchronization); err != nil {
		return err
	}
	if err := writer.WriteElementString("TimeOutNotAddressedDesynchronization", g.TimeOutNotAddressedDesynchronization); err != nil {
		return err
	}
	if err := writer.WriteElementString("TimeOutFrameNotOkDesynchronization", g.TimeOutFrameNotOkDesynchronization); err != nil {
		return err
	}
	if err := writer.WriteElementString("WriteRequestDesynchronization", g.WriteRequestDesynchronization); err != nil {
		return err
	}
	if err := writer.WriteElementString("WrongInitiatorDesynchronization", g.WrongInitiatorDesynchronization); err != nil {
		return err
	}
	if len(g.BroadcastFramesCounter) != 0 {
		if err := writer.WriteStartElement("BroadcastFramesCounters"); err != nil {
			return err
		}
		for _, it := range g.BroadcastFramesCounter {
			if err := writer.WriteStartElement("Item"); err != nil {
				return err
			}
			if err := writer.WriteElementString("Key", it.Key); err != nil {
				return err
			}
			if err := writer.WriteElementString("Value", it.Value); err != nil {
				return err
			}
			if err := writer.WriteEndElement(); err != nil {
				return err
			}
		}
		if err := writer.WriteEndElement(); err != nil {
			return err
		}
	}
	if err := writer.WriteElementString("RepetitionsCounter", g.RepetitionsCounter); err != nil {
		return err
	}
	if err := writer.WriteElementString("TransmissionsCounter", g.TransmissionsCounter); err != nil {
		return err
	}
	if err := writer.WriteElementString("CrcOkFramesCounter", g.CrcOkFramesCounter); err != nil {
		return err
	}
	return writer.WriteElementString("CrcNOkFramesCounter", g.CrcNOkFramesCounter)
}

func (g *GXDLMSSFSKMacCounters) PostLoad(reader *GXXmlReader) error { return nil }

func (g *GXDLMSSFSKMacCounters) GetValues() []any {
	return []any{
		g.LogicalName(),
		g.SynchronizationRegister,
		[]any{
			g.PhysicalLayerDesynchronization,
			g.TimeOutNotAddressedDesynchronization,
			g.TimeOutFrameNotOkDesynchronization,
			g.WriteRequestDesynchronization,
			g.WrongInitiatorDesynchronization,
		},
		g.BroadcastFramesCounter,
		g.RepetitionsCounter,
		g.TransmissionsCounter,
		g.CrcOkFramesCounter,
		g.CrcNOkFramesCounter,
	}
}

func (g *GXDLMSSFSKMacCounters) encodeCounterPairs(settings *settings.GXDLMSSettings, pairs []types.GXKeyValuePair[uint16, uint32]) ([]byte, error) {
	bb := types.NewGXByteBuffer()
	if err := bb.SetUint8(uint8(enums.DataTypeArray)); err != nil {
		return nil, err
	}
	types.SetObjectCount(len(pairs), bb)
	for _, it := range pairs {
		if err := bb.SetUint8(uint8(enums.DataTypeStructure)); err != nil {
			return nil, err
		}
		if err := bb.SetUint8(2); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeUint16, it.Key); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeUint32, it.Value); err != nil {
			return nil, err
		}
	}
	return bb.Array(), nil
}

func sfskParseCounterPairs(value any) ([]types.GXKeyValuePair[uint16, uint32], error) {
	if value == nil {
		return make([]types.GXKeyValuePair[uint16, uint32], 0), nil
	}
	rows, ok := value.(types.GXStructure)
	if !ok {
		return nil, fmt.Errorf("invalid counter pair array: %T", value)
	}
	ret := make([]types.GXKeyValuePair[uint16, uint32], 0, len(rows))
	for _, row := range rows {
		it, ok := row.(types.GXStructure)
		if !ok || len(it) < 2 {
			return nil, fmt.Errorf("invalid counter pair item: %T", row)
		}
		k32, err := toUint32(it[0])
		if err != nil {
			return nil, err
		}
		v, err := toUint32(it[1])
		if err != nil {
			return nil, err
		}
		ret = append(ret, *types.NewGXKeyValuePair(uint16(k32), v))
	}
	return ret, nil
}

// NewGXDLMSSFSKMacCounters creates a new s-FSK MAC counters object instance.
//
// The function validates `ln` before creating the object.
// `ln` is the Logical Name and `sn` is the Short Name of the object.
func NewGXDLMSSFSKMacCounters(ln string, sn int16) (*GXDLMSSFSKMacCounters, error) {
	if err := ValidateLogicalName(ln); err != nil {
		return nil, err
	}
	return &GXDLMSSFSKMacCounters{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeSFSKMacCounters,
			logicalName: ln,
			ShortName:   sn,
		},
		SynchronizationRegister: make([]types.GXKeyValuePair[uint16, uint32], 0),
		BroadcastFramesCounter:  make([]types.GXKeyValuePair[uint16, uint32], 0),
	}, nil
}
