package objects

import (
	"reflect"

	"github.com/Gurux/gxdlms-go/dlmserrors"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// GXDLMSCoAPSetup contains CoAP setup parameters.
type GXDLMSCoAPSetup struct {
	GXDLMSObject
	UdpReference       *GXDLMSTcpUdpSetup
	AckTimeout         uint16
	AckRandomFactor    uint16
	MaxRetransmit      uint16
	NStart             uint16
	DelayAckTimeout    uint16
	ExponentialBackOff uint16
	ProbingRate        uint16
	CoAPUriPath        string
	TransportMode      enums.TransportMode
	WrapperVersion     any
	TokenLength        uint8
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSCoAPSetup) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

func (g *GXDLMSCoAPSetup) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, nil
}

func (g *GXDLMSCoAPSetup) GetAttributeIndexToRead(all bool) []int {
	var a []int
	if all || g.LogicalName() == "" {
		a = append(a, 1)
	}
	for i := 2; i <= 13; i++ {
		if all || g.CanRead(i) {
			a = append(a, i)
		}
	}
	return a
}

func (g *GXDLMSCoAPSetup) GetNames() []string {
	return []string{
		"Logical Name",
		"UdpReference",
		"AckTimeout",
		"AckRandomFactor",
		"MaxRetransmit",
		"NStart",
		"DelayAckTimeout",
		"ExponentialBackOff",
		"ProbingRate",
		"CoAPUriPath",
		"TransportMode",
		"WrapperVersion",
		"TokenLength",
	}
}

func (g *GXDLMSCoAPSetup) GetMethodNames() []string { return []string{} }

func (g *GXDLMSCoAPSetup) GetAttributeCount() int { return 13 }

func (g *GXDLMSCoAPSetup) GetMethodCount() int { return 0 }

func (g *GXDLMSCoAPSetup) GetDataType(index int) (enums.DataType, error) {
	switch index {
	case 1, 2, 10:
		return enums.DataTypeOctetString, nil
	case 3, 4, 5, 6, 7, 8, 9:
		return enums.DataTypeUint16, nil
	case 11:
		return enums.DataTypeEnum, nil
	case 12:
		dt, err := g.GXDLMSObject.GetDataType(index)
		if err == nil && dt == enums.DataTypeNone && g.WrapperVersion != nil {
			dt, err = internal.GetDLMSDataType(reflect.TypeOf(g.WrapperVersion))
			if err != nil {
				return enums.DataTypeNone, err
			}
		}
		return dt, err
	case 13:
		return enums.DataTypeUint8, nil
	default:
		return enums.DataTypeNone, dlmserrors.ErrInvalidAttributeIndex
	}
}

func (g *GXDLMSCoAPSetup) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	switch e.Index {
	case 1:
		return helpers.LogicalNameToBytes(g.LogicalName())
	case 2:
		if g.UdpReference == nil {
			return nil, nil
		}
		return helpers.LogicalNameToBytes(g.UdpReference.LogicalName())
	case 3:
		return g.AckTimeout, nil
	case 4:
		return g.AckRandomFactor, nil
	case 5:
		return g.MaxRetransmit, nil
	case 6:
		return g.NStart, nil
	case 7:
		return g.DelayAckTimeout, nil
	case 8:
		return g.ExponentialBackOff, nil
	case 9:
		return g.ProbingRate, nil
	case 10:
		return []byte(g.CoAPUriPath), nil
	case 11:
		return uint8(g.TransportMode), nil
	case 12:
		return g.WrapperVersion, nil
	case 13:
		return g.TokenLength, nil
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
		return nil, nil
	}
}

func (g *GXDLMSCoAPSetup) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	switch e.Index {
	case 1:
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return err
		}
		return g.SetLogicalName(ln)
	case 2:
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			return err
		}
		g.UdpReference = nil
		if settings != nil && settings.Objects != nil {
			switch objects := settings.Objects.(type) {
			case GXDLMSObjectCollection:
				if tmp := objects.FindByLN(enums.ObjectTypeTCPUDPSetup, ln); tmp != nil {
					if udp, ok := tmp.(*GXDLMSTcpUdpSetup); ok {
						g.UdpReference = udp
					}
				}
			case *GXDLMSObjectCollection:
				if tmp := objects.FindByLN(enums.ObjectTypeTCPUDPSetup, ln); tmp != nil {
					if udp, ok := tmp.(*GXDLMSTcpUdpSetup); ok {
						g.UdpReference = udp
					}
				}
			}
		}
		if g.UdpReference == nil {
			udp, err := NewGXDLMSTcpUdpSetup(ln, 0)
			if err != nil {
				return err
			}
			g.UdpReference = udp
		}
	case 3:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.AckTimeout = uint16(v)
	case 4:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.AckRandomFactor = uint16(v)
	case 5:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.MaxRetransmit = uint16(v)
	case 6:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.NStart = uint16(v)
	case 7:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.DelayAckTimeout = uint16(v)
	case 8:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.ExponentialBackOff = uint16(v)
	case 9:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.ProbingRate = uint16(v)
	case 10:
		if e.Value == nil {
			g.CoAPUriPath = ""
		} else if b, ok := e.Value.([]byte); ok {
			g.CoAPUriPath = string(b)
		} else {
			g.CoAPUriPath = ""
		}
	case 11:
		g.TransportMode = enums.TransportMode(e.Value.(types.GXEnum).Value)
	case 12:
		g.WrapperVersion = e.Value
	case 13:
		g.TokenLength = e.Value.(uint8)
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil
}

func (g *GXDLMSCoAPSetup) Load(reader *GXXmlReader) error {
	var err error
	ln, err := reader.ReadElementContentAsString("UdpReference", "")
	if err != nil {
		return err
	}
	if ln != "" && reader.Objects != nil {
		if tmp := reader.Objects.FindByLN(enums.ObjectTypeTCPUDPSetup, ln); tmp != nil {
			if udp, ok := tmp.(*GXDLMSTcpUdpSetup); ok {
				g.UdpReference = udp
			}
		}
	}
	if g.UdpReference == nil && ln != "" {
		udp, err := NewGXDLMSTcpUdpSetup(ln, 0)
		if err != nil {
			return err
		}
		g.UdpReference = udp
	}
	g.AckTimeout, err = reader.ReadElementContentAsUInt16("AckTimeout", 0)
	if err != nil {
		return err
	}
	g.AckRandomFactor, err = reader.ReadElementContentAsUInt16("AckRandomFactor", 0)
	if err != nil {
		return err
	}
	g.MaxRetransmit, err = reader.ReadElementContentAsUInt16("MaxRetransmit", 0)
	if err != nil {
		return err
	}
	g.NStart, err = reader.ReadElementContentAsUInt16("NStart", 0)
	if err != nil {
		return err
	}
	g.DelayAckTimeout, err = reader.ReadElementContentAsUInt16("DelayAckTimeout", 0)
	if err != nil {
		return err
	}
	g.ExponentialBackOff, err = reader.ReadElementContentAsUInt16("ExponentialBackOff", 0)
	if err != nil {
		return err
	}
	g.ProbingRate, err = reader.ReadElementContentAsUInt16("ProbingRate", 0)
	if err != nil {
		return err
	}
	g.CoAPUriPath, err = reader.ReadElementContentAsString("CoAPUriPath", "")
	if err != nil {
		return err
	}
	mode, err := reader.ReadElementContentAsInt("TransportMode", 0)
	if err != nil {
		return err
	}
	g.TransportMode = enums.TransportMode(mode)
	g.WrapperVersion, err = reader.ReadElementContentAsObject("WrapperVersion", nil, g, 12)
	if err != nil {
		return err
	}
	g.TokenLength, err = reader.ReadElementContentAsUInt8("TokenLength", 0)
	return err
}

func (g *GXDLMSCoAPSetup) Save(writer *GXXmlWriter) error {
	if g.UdpReference != nil {
		if err := writer.WriteElementString("UdpReference", g.UdpReference.LogicalName()); err != nil {
			return err
		}
	}
	if err := writer.WriteElementString("AckTimeout", g.AckTimeout); err != nil {
		return err
	}
	if err := writer.WriteElementString("AckRandomFactor", g.AckRandomFactor); err != nil {
		return err
	}
	if err := writer.WriteElementString("MaxRetransmit", g.MaxRetransmit); err != nil {
		return err
	}
	if err := writer.WriteElementString("NStart", g.NStart); err != nil {
		return err
	}
	if err := writer.WriteElementString("DelayAckTimeout", g.DelayAckTimeout); err != nil {
		return err
	}
	if err := writer.WriteElementString("ExponentialBackOff", g.ExponentialBackOff); err != nil {
		return err
	}
	if err := writer.WriteElementString("ProbingRate", g.ProbingRate); err != nil {
		return err
	}
	if err := writer.WriteElementString("CoAPUriPath", g.CoAPUriPath); err != nil {
		return err
	}
	if err := writer.WriteElementString("TransportMode", int(g.TransportMode)); err != nil {
		return err
	}
	dt, err := g.GetDataType(12)
	if err != nil {
		return err
	}
	if err = writer.WriteElementObject("WrapperVersion", g.WrapperVersion, dt, g.GetUIDataType(12)); err != nil {
		return err
	}
	return writer.WriteElementString("TokenLength", g.TokenLength)
}

func (g *GXDLMSCoAPSetup) PostLoad(reader *GXXmlReader) error {
	if g.UdpReference == nil || reader.Objects == nil {
		return nil
	}
	if target := reader.Objects.FindByLN(enums.ObjectTypeTCPUDPSetup, g.UdpReference.LogicalName()); target != nil {
		if udp, ok := target.(*GXDLMSTcpUdpSetup); ok {
			g.UdpReference = udp
		}
	}
	return nil
}

func (g *GXDLMSCoAPSetup) GetValues() []any {
	return []any{
		g.LogicalName(),
		g.UdpReference,
		g.AckTimeout,
		g.AckRandomFactor,
		g.MaxRetransmit,
		g.NStart,
		g.DelayAckTimeout,
		g.ExponentialBackOff,
		g.ProbingRate,
		g.CoAPUriPath,
		g.TransportMode,
		g.WrapperVersion,
		g.TokenLength,
	}
}

// NewGXDLMSCoAPSetup creates a new CoAP setup object instance.
func NewGXDLMSCoAPSetup(ln string, sn int16) (*GXDLMSCoAPSetup, error) {
	if err := ValidateLogicalName(ln); err != nil {
		return nil, err
	}
	return &GXDLMSCoAPSetup{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeCoAPSetup,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
