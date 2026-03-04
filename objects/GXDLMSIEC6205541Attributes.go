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

// GXDLMSMeterPrimaryAccountNumber contains meter PAN information.
type GXDLMSMeterPrimaryAccountNumber struct {
	IssuerId               uint32
	DecoderReferenceNumber uint64
	PanCheckDigit          uint8
}

// GXDLMSIec6205541Attributes contains IEC 62055-41 related attributes.
type GXDLMSIec6205541Attributes struct {
	GXDLMSObject
	MeterPan            GXDLMSMeterPrimaryAccountNumber
	Commodity           string
	TokenCarrierTypes   []byte
	EncryptionAlgorithm uint8
	SupplyGroupCode     uint32
	TariffIndex         uint8
	KeyRevisionNumber   uint8
	KeyType             uint8
	KeyExpiryNumber     uint8
	KctSupported        uint8
	StsCertificate      string
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSIec6205541Attributes) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

func (g *GXDLMSIec6205541Attributes) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, nil
}

func (g *GXDLMSIec6205541Attributes) GetAttributeIndexToRead(all bool) []int {
	var a []int
	if all || g.LogicalName() == "" {
		a = append(a, 1)
	}
	for i := 2; i <= 12; i++ {
		if all || g.CanRead(i) {
			a = append(a, i)
		}
	}
	return a
}

func (g *GXDLMSIec6205541Attributes) GetNames() []string {
	return []string{
		"Logical Name",
		"MeterPan",
		"Commodity",
		"TokenCarrierTypes",
		"EncryptionAlgorithm",
		"SupplyGroupCode",
		"TariffIndex",
		"KeyRevisionNumber",
		"KeyType",
		"KeyExpiryNumber",
		"KctSupported",
		"StsCertificate",
	}
}

func (g *GXDLMSIec6205541Attributes) GetMethodNames() []string { return []string{} }

func (g *GXDLMSIec6205541Attributes) GetAttributeCount() int { return 12 }

func (g *GXDLMSIec6205541Attributes) GetMethodCount() int { return 0 }

func (g *GXDLMSIec6205541Attributes) GetDataType(index int) (enums.DataType, error) {
	switch index {
	case 1:
		return enums.DataTypeOctetString, nil
	case 2:
		return enums.DataTypeStructure, nil
	case 3, 12:
		return enums.DataTypeString, nil
	case 4:
		return enums.DataTypeArray, nil
	case 5, 7, 8, 9, 10, 11:
		return enums.DataTypeUint8, nil
	case 6:
		return enums.DataTypeUint32, nil
	default:
		return 0, dlmserrors.ErrInvalidAttributeIndex
	}
}

func (g *GXDLMSIec6205541Attributes) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	switch e.Index {
	case 1:
		return helpers.LogicalNameToBytes(g.LogicalName())
	case 2:
		bb := types.NewGXByteBuffer()
		if err := bb.SetUint8(uint8(enums.DataTypeStructure)); err != nil {
			return nil, err
		}
		if err := bb.SetUint8(3); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeUint32, g.MeterPan.IssuerId); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeUint64, g.MeterPan.DecoderReferenceNumber); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeUint8, g.MeterPan.PanCheckDigit); err != nil {
			return nil, err
		}
		return bb.Array(), nil
	case 3:
		return g.Commodity, nil
	case 4:
		bb := types.NewGXByteBuffer()
		if err := bb.SetUint8(uint8(enums.DataTypeArray)); err != nil {
			return nil, err
		}
		types.SetObjectCount(len(g.TokenCarrierTypes), bb)
		for _, it := range g.TokenCarrierTypes {
			if err := internal.SetData(settings, bb, enums.DataTypeUint8, it); err != nil {
				return nil, err
			}
		}
		return bb.Array(), nil
	case 5:
		return g.EncryptionAlgorithm, nil
	case 6:
		return g.SupplyGroupCode, nil
	case 7:
		return g.TariffIndex, nil
	case 8:
		return g.KeyRevisionNumber, nil
	case 9:
		return g.KeyType, nil
	case 10:
		return g.KeyExpiryNumber, nil
	case 11:
		return g.KctSupported, nil
	case 12:
		return g.StsCertificate, nil
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
		return nil, nil
	}
}

func (g *GXDLMSIec6205541Attributes) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	switch e.Index {
	case 1:
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return err
		}
		return g.SetLogicalName(ln)
	case 2:
		if e.Value == nil {
			g.MeterPan = GXDLMSMeterPrimaryAccountNumber{}
			return nil
		}
		tmp, ok := e.Value.(types.GXStructure)
		if !ok || len(tmp) < 3 {
			return fmt.Errorf("invalid meter pan value: %T", e.Value)
		}
		issuerID, err := toUint32(tmp[0])
		if err != nil {
			return err
		}
		drn, err := toUint64(tmp[1])
		if err != nil {
			return err
		}
		check, err := toUint8(tmp[2])
		if err != nil {
			return err
		}
		g.MeterPan.IssuerId = issuerID
		g.MeterPan.DecoderReferenceNumber = drn
		g.MeterPan.PanCheckDigit = check
	case 3:
		if e.Value == nil {
			g.Commodity = ""
		} else {
			g.Commodity = fmt.Sprint(e.Value)
		}
	case 4:
		if e.Value == nil {
			g.TokenCarrierTypes = nil
			return nil
		}
		tmp, ok := e.Value.(types.GXArray)
		if !ok {
			return fmt.Errorf("invalid token carrier types: %T", e.Value)
		}
		g.TokenCarrierTypes = make([]byte, 0, len(tmp))
		for _, it := range tmp {
			v, err := toUint8(it)
			if err != nil {
				return err
			}
			g.TokenCarrierTypes = append(g.TokenCarrierTypes, v)
		}
	case 5:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.EncryptionAlgorithm = v
	case 6:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.SupplyGroupCode = v
	case 7:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.TariffIndex = v
	case 8:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.KeyRevisionNumber = v
	case 9:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.KeyType = v
	case 10:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.KeyExpiryNumber = v
	case 11:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.KctSupported = v
	case 12:
		if e.Value == nil {
			g.StsCertificate = ""
		} else {
			g.StsCertificate = fmt.Sprint(e.Value)
		}
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil
}

func (g *GXDLMSIec6205541Attributes) Load(reader *GXXmlReader) error {
	var err error
	g.MeterPan.IssuerId, err = reader.ReadElementContentAsUInt32("IssuerId", 0)
	if err != nil {
		return err
	}
	g.MeterPan.DecoderReferenceNumber, err = reader.ReadElementContentAsULong("DecoderReferenceNumber", 0)
	if err != nil {
		return err
	}
	g.MeterPan.PanCheckDigit, err = reader.ReadElementContentAsUInt8("PanCheckDigit", 0)
	if err != nil {
		return err
	}
	g.Commodity, err = reader.ReadElementContentAsString("Commodity", "")
	if err != nil {
		return err
	}
	g.TokenCarrierTypes = g.TokenCarrierTypes[:0]
	ok, err := reader.IsStartElementNamed("TokenCarrierTypes", true)
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
			v, err := reader.ReadElementContentAsUInt8("Value", 0)
			if err != nil {
				return err
			}
			g.TokenCarrierTypes = append(g.TokenCarrierTypes, v)
		}
		if err = reader.ReadEndElement("TokenCarrierTypes"); err != nil {
			return err
		}
	}
	g.EncryptionAlgorithm, err = reader.ReadElementContentAsUInt8("EncryptionAlgorithm", 0)
	if err != nil {
		return err
	}
	g.SupplyGroupCode, err = reader.ReadElementContentAsUInt32("SupplyGroupCode", 0)
	if err != nil {
		return err
	}
	g.TariffIndex, err = reader.ReadElementContentAsUInt8("TariffIndex", 0)
	if err != nil {
		return err
	}
	g.KeyRevisionNumber, err = reader.ReadElementContentAsUInt8("KeyRevisionNumber", 0)
	if err != nil {
		return err
	}
	g.KeyType, err = reader.ReadElementContentAsUInt8("KeyType", 0)
	if err != nil {
		return err
	}
	g.KeyExpiryNumber, err = reader.ReadElementContentAsUInt8("KeyExpiryNumber", 0)
	if err != nil {
		return err
	}
	g.KctSupported, err = reader.ReadElementContentAsUInt8("KctSupported", 0)
	if err != nil {
		return err
	}
	g.StsCertificate, err = reader.ReadElementContentAsString("StsCertificate", "")
	return err
}

func (g *GXDLMSIec6205541Attributes) Save(writer *GXXmlWriter) error {
	if err := writer.WriteElementString("IssuerId", g.MeterPan.IssuerId); err != nil {
		return err
	}
	if err := writer.WriteElementString("DecoderReferenceNumber", g.MeterPan.DecoderReferenceNumber); err != nil {
		return err
	}
	if err := writer.WriteElementString("PanCheckDigit", g.MeterPan.PanCheckDigit); err != nil {
		return err
	}
	if err := writer.WriteElementString("Commodity", g.Commodity); err != nil {
		return err
	}
	if err := writer.WriteStartElement("TokenCarrierTypes"); err != nil {
		return err
	}
	for _, it := range g.TokenCarrierTypes {
		if err := writer.WriteElementString("Value", it); err != nil {
			return err
		}
	}
	if err := writer.WriteEndElement(); err != nil {
		return err
	}
	if err := writer.WriteElementString("EncryptionAlgorithm", g.EncryptionAlgorithm); err != nil {
		return err
	}
	if err := writer.WriteElementString("SupplyGroupCode", g.SupplyGroupCode); err != nil {
		return err
	}
	if err := writer.WriteElementString("TariffIndex", g.TariffIndex); err != nil {
		return err
	}
	if err := writer.WriteElementString("KeyRevisionNumber", g.KeyRevisionNumber); err != nil {
		return err
	}
	if err := writer.WriteElementString("KeyType", g.KeyType); err != nil {
		return err
	}
	if err := writer.WriteElementString("KeyExpiryNumber", g.KeyExpiryNumber); err != nil {
		return err
	}
	if err := writer.WriteElementString("KctSupported", g.KctSupported); err != nil {
		return err
	}
	return writer.WriteElementString("StsCertificate", g.StsCertificate)
}

func (g *GXDLMSIec6205541Attributes) PostLoad(reader *GXXmlReader) error { return nil }

func (g *GXDLMSIec6205541Attributes) GetValues() []any {
	return []any{
		g.LogicalName(),
		g.MeterPan,
		g.Commodity,
		g.TokenCarrierTypes,
		g.EncryptionAlgorithm,
		g.SupplyGroupCode,
		g.TariffIndex,
		g.KeyRevisionNumber,
		g.KeyType,
		g.KeyExpiryNumber,
		g.KctSupported,
		g.StsCertificate,
	}
}

// NewGXDLMSIec6205541Attributes creates a new IEC 62055-41 attributes object instance.
func NewGXDLMSIec6205541Attributes(ln string, sn int16) (*GXDLMSIec6205541Attributes, error) {
	if err := ValidateLogicalName(ln); err != nil {
		return nil, err
	}
	return &GXDLMSIec6205541Attributes{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeIEC6205541Attributes,
			logicalName: ln,
			ShortName:   sn,
		},
		TokenCarrierTypes: make([]byte, 0),
	}, nil
}
