package objects

import (
	"github.com/Gurux/gxdlms-go/dlmserrors"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
)

// GXDLMSSFSKMacSynchronizationTimeouts contains s-FSK MAC synchronization timeouts.
type GXDLMSSFSKMacSynchronizationTimeouts struct {
	GXDLMSObject
	SearchInitiatorTimeout             uint16
	SynchronizationConfirmationTimeout uint16
	TimeOutNotAddressed                uint16
	TimeOutFrameNotOK                  uint16
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSSFSKMacSynchronizationTimeouts) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

func (g *GXDLMSSFSKMacSynchronizationTimeouts) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, nil
}

func (g *GXDLMSSFSKMacSynchronizationTimeouts) GetAttributeIndexToRead(all bool) []int {
	var a []int
	if all || g.LogicalName() == "" {
		a = append(a, 1)
	}
	for i := 2; i <= 5; i++ {
		if all || g.CanRead(i) {
			a = append(a, i)
		}
	}
	return a
}

func (g *GXDLMSSFSKMacSynchronizationTimeouts) GetNames() []string {
	return []string{"Logical Name", "SearchInitiatorTimeout", "SynchronizationConfirmationTimeout", "TimeOutNotAddressed", "TimeOutFrameNotOK"}
}

func (g *GXDLMSSFSKMacSynchronizationTimeouts) GetMethodNames() []string { return []string{} }

func (g *GXDLMSSFSKMacSynchronizationTimeouts) GetAttributeCount() int { return 5 }

func (g *GXDLMSSFSKMacSynchronizationTimeouts) GetMethodCount() int { return 0 }

func (g *GXDLMSSFSKMacSynchronizationTimeouts) GetDataType(index int) (enums.DataType, error) {
	switch index {
	case 1:
		return enums.DataTypeOctetString, nil
	case 2, 3, 4, 5:
		return enums.DataTypeUint16, nil
	default:
		return 0, dlmserrors.ErrInvalidAttributeIndex
	}
}

func (g *GXDLMSSFSKMacSynchronizationTimeouts) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	switch e.Index {
	case 1:
		return helpers.LogicalNameToBytes(g.LogicalName())
	case 2:
		return g.SearchInitiatorTimeout, nil
	case 3:
		return g.SynchronizationConfirmationTimeout, nil
	case 4:
		return g.TimeOutNotAddressed, nil
	case 5:
		return g.TimeOutFrameNotOK, nil
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
		return nil, nil
	}
}

func (g *GXDLMSSFSKMacSynchronizationTimeouts) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
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
		g.SearchInitiatorTimeout = uint16(v)
	case 3:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.SynchronizationConfirmationTimeout = uint16(v)
	case 4:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.TimeOutNotAddressed = uint16(v)
	case 5:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.TimeOutFrameNotOK = uint16(v)
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil
}

func (g *GXDLMSSFSKMacSynchronizationTimeouts) Load(reader *GXXmlReader) error {
	var err error
	g.SearchInitiatorTimeout, err = reader.ReadElementContentAsUInt16("SearchInitiatorTimeout", 0)
	if err != nil {
		return err
	}
	g.SynchronizationConfirmationTimeout, err = reader.ReadElementContentAsUInt16("SynchronizationConfirmationTimeout", 0)
	if err != nil {
		return err
	}
	g.TimeOutNotAddressed, err = reader.ReadElementContentAsUInt16("TimeOutNotAddressed", 0)
	if err != nil {
		return err
	}
	g.TimeOutFrameNotOK, err = reader.ReadElementContentAsUInt16("TimeOutFrameNotOK", 0)
	return err
}

func (g *GXDLMSSFSKMacSynchronizationTimeouts) Save(writer *GXXmlWriter) error {
	if err := writer.WriteElementString("SearchInitiatorTimeout", g.SearchInitiatorTimeout); err != nil {
		return err
	}
	if err := writer.WriteElementString("SynchronizationConfirmationTimeout", g.SynchronizationConfirmationTimeout); err != nil {
		return err
	}
	if err := writer.WriteElementString("TimeOutNotAddressed", g.TimeOutNotAddressed); err != nil {
		return err
	}
	return writer.WriteElementString("TimeOutFrameNotOK", g.TimeOutFrameNotOK)
}

func (g *GXDLMSSFSKMacSynchronizationTimeouts) PostLoad(reader *GXXmlReader) error { return nil }

func (g *GXDLMSSFSKMacSynchronizationTimeouts) GetValues() []any {
	return []any{
		g.LogicalName(),
		g.SearchInitiatorTimeout,
		g.SynchronizationConfirmationTimeout,
		g.TimeOutNotAddressed,
		g.TimeOutFrameNotOK,
	}
}

// NewGXDLMSSFSKMacSynchronizationTimeouts creates a new s-FSK MAC synchronization timeouts object instance.
//
// The function validates `ln` before creating the object.
// `ln` is the Logical Name and `sn` is the Short Name of the object.
func NewGXDLMSSFSKMacSynchronizationTimeouts(ln string, sn int16) (*GXDLMSSFSKMacSynchronizationTimeouts, error) {
	if err := ValidateLogicalName(ln); err != nil {
		return nil, err
	}
	return &GXDLMSSFSKMacSynchronizationTimeouts{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeSFSKMacSynchronizationTimeouts,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
