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

// Online help:
// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSLteMonitoring
type GXDLMSLteMonitoring struct {
	GXDLMSObject

	// Network parameters for the LTE network.
	NetworkParameters GXLteNetworkParameters

	// Quality of service of the LTE network.
	QualityOfService GXLteQualityOfService
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSLteMonitoring) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

func lteToUint32(value any) (uint32, error) {
	switch v := value.(type) {
	case uint8:
		return uint32(v), nil
	case uint16:
		return uint32(v), nil
	case uint32:
		return v, nil
	case uint64:
		return uint32(v), nil
	case int8:
		return uint32(v), nil
	case int16:
		return uint32(v), nil
	case int32:
		return uint32(v), nil
	case int:
		return uint32(v), nil
	case types.GXEnum:
		return uint32(v.Value), nil
	default:
		return 0, fmt.Errorf("invalid integer type: %T", value)
	}
}

func lteToInt8(value any) (int8, error) {
	switch v := value.(type) {
	case int8:
		return v, nil
	case uint8:
		return int8(v), nil
	case int16:
		return int8(v), nil
	case uint16:
		return int8(v), nil
	case int32:
		return int8(v), nil
	case uint32:
		return int8(v), nil
	case int:
		return int8(v), nil
	case types.GXEnum:
		return int8(v.Value), nil
	default:
		return 0, fmt.Errorf("invalid int8 type: %T", value)
	}
}

func lteAsStructure(value any) (types.GXStructure, error) {
	switch v := value.(type) {
	case types.GXStructure:
		return v, nil
	case []any:
		return types.GXStructure(v), nil
	default:
		return nil, fmt.Errorf("invalid structure type: %T", value)
	}
}

// Invoke invokes method.
func (g *GXDLMSLteMonitoring) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, nil
}

// GetAttributeIndexToRead returns the collection of attributes to read.
func (g *GXDLMSLteMonitoring) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	if all || g.CanRead(2) {
		attributes = append(attributes, 2)
	}
	if g.Version > 0 {
		if all || g.CanRead(3) {
			attributes = append(attributes, 3)
		}
	}
	return attributes
}

// GetNames returns the names of attribute indexes.
func (g *GXDLMSLteMonitoring) GetNames() []string {
	return []string{"Logical Name", "Network parameters", "Quality of service"}
}

// GetMethodNames returns the names of method indexes.
func (g *GXDLMSLteMonitoring) GetMethodNames() []string {
	return []string{}
}

// GetAttributeCount returns the amount of attributes.
func (g *GXDLMSLteMonitoring) GetAttributeCount() int {
	if g.Version == 0 {
		return 2
	}
	return 3
}

// GetMethodCount returns the amount of methods.
func (g *GXDLMSLteMonitoring) GetMethodCount() int {
	return 0
}

// GetDataType returns the device data type of selected attribute index.
func (g *GXDLMSLteMonitoring) GetDataType(index int) (enums.DataType, error) {
	switch index {
	case 1:
		return enums.DataTypeOctetString, nil
	case 2:
		return enums.DataTypeStructure, nil
	case 3:
		return enums.DataTypeStructure, nil
	default:
		return 0, dlmserrors.ErrInvalidAttributeIndex
	}
}

// GetValue returns the value of given attribute.
func (g *GXDLMSLteMonitoring) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	switch e.Index {
	case 1:
		return helpers.LogicalNameToBytes(g.LogicalName())
	case 2:
		buff := types.NewGXByteBuffer()
		if err := buff.SetUint8(uint8(enums.DataTypeStructure)); err != nil {
			return nil, err
		}
		if g.Version == 0 {
			if err := types.SetObjectCount(5, buff); err != nil {
				return nil, err
			}
			if err := internal.SetData(settings, buff, enums.DataTypeUint16, g.NetworkParameters.T3402); err != nil {
				return nil, err
			}
			if err := internal.SetData(settings, buff, enums.DataTypeUint16, g.NetworkParameters.T3412); err != nil {
				return nil, err
			}
			if err := internal.SetData(settings, buff, enums.DataTypeUint8, g.QualityOfService.SignalQuality); err != nil {
				return nil, err
			}
			if err := internal.SetData(settings, buff, enums.DataTypeUint8, g.QualityOfService.SignalLevel); err != nil {
				return nil, err
			}
			if err := internal.SetData(settings, buff, enums.DataTypeInt8, g.NetworkParameters.QRxlevMinCE); err != nil {
				return nil, err
			}
		} else {
			if err := types.SetObjectCount(9, buff); err != nil {
				return nil, err
			}
			if err := internal.SetData(settings, buff, enums.DataTypeUint16, g.NetworkParameters.T3402); err != nil {
				return nil, err
			}
			if err := internal.SetData(settings, buff, enums.DataTypeUint16, g.NetworkParameters.T3412); err != nil {
				return nil, err
			}
			if err := internal.SetData(settings, buff, enums.DataTypeUint32, g.NetworkParameters.T3412ext2); err != nil {
				return nil, err
			}
			if err := internal.SetData(settings, buff, enums.DataTypeUint16, g.NetworkParameters.T3324); err != nil {
				return nil, err
			}
			if err := internal.SetData(settings, buff, enums.DataTypeUint32, g.NetworkParameters.TeDRX); err != nil {
				return nil, err
			}
			if err := internal.SetData(settings, buff, enums.DataTypeUint16, g.NetworkParameters.TPTW); err != nil {
				return nil, err
			}
			if err := internal.SetData(settings, buff, enums.DataTypeInt8, g.NetworkParameters.QRxlevMin); err != nil {
				return nil, err
			}
			if err := internal.SetData(settings, buff, enums.DataTypeInt8, g.NetworkParameters.QRxlevMinCE); err != nil {
				return nil, err
			}
			if err := internal.SetData(settings, buff, enums.DataTypeInt8, g.NetworkParameters.QRxLevMinCE1); err != nil {
				return nil, err
			}
		}
		return buff.Array(), nil
	case 3:
		if g.Version == 0 {
			e.Error = enums.ErrorCodeReadWriteDenied
			return nil, nil
		}
		buff := types.NewGXByteBuffer()
		if err := buff.SetUint8(uint8(enums.DataTypeStructure)); err != nil {
			return nil, err
		}
		if err := types.SetObjectCount(4, buff); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, buff, enums.DataTypeInt8, g.QualityOfService.SignalQuality); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, buff, enums.DataTypeInt8, g.QualityOfService.SignalLevel); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, buff, enums.DataTypeInt8, g.QualityOfService.SignalToNoiseRatio); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, buff, enums.DataTypeEnum, uint8(g.QualityOfService.CoverageEnhancement)); err != nil {
			return nil, err
		}
		return buff.Array(), nil
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
		return nil, nil
	}
}

// SetValue sets value of given attribute.
func (g *GXDLMSLteMonitoring) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	switch e.Index {
	case 1:
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return err
		}
		return g.SetLogicalName(ln)
	case 2:
		s, err := lteAsStructure(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return err
		}
		if g.Version == 0 {
			v, err := lteToUint32(s[0])
			if err != nil {
				return err
			}
			g.NetworkParameters.T3402 = uint16(v)
			v, err = lteToUint32(s[1])
			if err != nil {
				return err
			}
			g.NetworkParameters.T3412 = uint16(v)
			i, err := lteToInt8(s[2])
			if err != nil {
				return err
			}
			g.QualityOfService.SignalQuality = i
			i, err = lteToInt8(s[3])
			if err != nil {
				return err
			}
			g.QualityOfService.SignalLevel = i
			i, err = lteToInt8(s[4])
			if err != nil {
				return err
			}
			g.QualityOfService.SignalToNoiseRatio = i
		} else {
			v, err := lteToUint32(s[0])
			if err != nil {
				return err
			}
			g.NetworkParameters.T3402 = uint16(v)
			v, err = lteToUint32(s[1])
			if err != nil {
				return err
			}
			g.NetworkParameters.T3412 = uint16(v)
			v, err = lteToUint32(s[2])
			if err != nil {
				return err
			}
			g.NetworkParameters.T3412ext2 = v
			v, err = lteToUint32(s[3])
			if err != nil {
				return err
			}
			g.NetworkParameters.T3324 = uint16(v)
			v, err = lteToUint32(s[4])
			if err != nil {
				return err
			}
			g.NetworkParameters.TeDRX = v
			v, err = lteToUint32(s[5])
			if err != nil {
				return err
			}
			g.NetworkParameters.TPTW = uint16(v)
			i, err := lteToInt8(s[6])
			if err != nil {
				return err
			}
			g.NetworkParameters.QRxlevMin = i
			i, err = lteToInt8(s[7])
			if err != nil {
				return err
			}
			g.NetworkParameters.QRxlevMinCE = i
			i, err = lteToInt8(s[8])
			if err != nil {
				return err
			}
			g.NetworkParameters.QRxLevMinCE1 = i
		}
	case 3:
		if g.Version == 0 {
			e.Error = enums.ErrorCodeReadWriteDenied
			return nil
		}
		s, err := lteAsStructure(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return err
		}
		i, err := lteToInt8(s[0])
		if err != nil {
			return err
		}
		g.QualityOfService.SignalQuality = i
		i, err = lteToInt8(s[1])
		if err != nil {
			return err
		}
		g.QualityOfService.SignalLevel = i
		i, err = lteToInt8(s[2])
		if err != nil {
			return err
		}
		g.QualityOfService.SignalToNoiseRatio = i
		v, err := lteToUint32(s[3])
		if err != nil {
			return err
		}
		g.QualityOfService.CoverageEnhancement = enums.LteCoverageEnhancement(v)
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil
}

// Load loads object content from XML.
func (g *GXDLMSLteMonitoring) Load(reader *GXXmlReader) error {
	var err error
	g.NetworkParameters.T3402, err = reader.ReadElementContentAsUInt16("T3402", 0)
	if err != nil {
		return err
	}
	g.NetworkParameters.T3412, err = reader.ReadElementContentAsUInt16("T3412", 0)
	if err != nil {
		return err
	}
	g.NetworkParameters.T3412ext2, err = reader.ReadElementContentAsUInt32("T3412ext2", 0)
	if err != nil {
		return err
	}
	g.NetworkParameters.T3324, err = reader.ReadElementContentAsUInt16("T3324", 0)
	if err != nil {
		return err
	}
	g.NetworkParameters.TeDRX, err = reader.ReadElementContentAsUInt32("TeDRX", 0)
	if err != nil {
		return err
	}
	g.NetworkParameters.TPTW, err = reader.ReadElementContentAsUInt16("TPTW", 0)
	if err != nil {
		return err
	}
	v, err := reader.ReadElementContentAsInt("QRxlevMin", 0)
	if err != nil {
		return err
	}
	g.NetworkParameters.QRxlevMin = int8(v)
	v, err = reader.ReadElementContentAsInt("QRxlevMinCE", 0)
	if err != nil {
		return err
	}
	g.NetworkParameters.QRxlevMinCE = int8(v)
	v, err = reader.ReadElementContentAsInt("QRxLevMinCE1", 0)
	if err != nil {
		return err
	}
	g.NetworkParameters.QRxLevMinCE1 = int8(v)
	v, err = reader.ReadElementContentAsInt("SignalQuality", 0)
	if err != nil {
		return err
	}
	g.QualityOfService.SignalQuality = int8(v)
	v, err = reader.ReadElementContentAsInt("SignalLevel", 0)
	if err != nil {
		return err
	}
	g.QualityOfService.SignalLevel = int8(v)
	v, err = reader.ReadElementContentAsInt("SignalToNoiseRatio", 0)
	if err != nil {
		return err
	}
	g.QualityOfService.SignalToNoiseRatio = int8(v)
	v, err = reader.ReadElementContentAsInt("CoverageEnhancement", 0)
	if err != nil {
		return err
	}
	g.QualityOfService.CoverageEnhancement = enums.LteCoverageEnhancement(v)
	return nil
}

// Save saves object content to XML.
func (g *GXDLMSLteMonitoring) Save(writer *GXXmlWriter) error {
	if err := writer.WriteElementString("T3402", g.NetworkParameters.T3402); err != nil {
		return err
	}
	if err := writer.WriteElementString("T3412", g.NetworkParameters.T3412); err != nil {
		return err
	}
	if err := writer.WriteElementString("T3412ext2", g.NetworkParameters.T3412ext2); err != nil {
		return err
	}
	if err := writer.WriteElementString("T3324", g.NetworkParameters.T3324); err != nil {
		return err
	}
	if err := writer.WriteElementString("TeDRX", g.NetworkParameters.TeDRX); err != nil {
		return err
	}
	if err := writer.WriteElementString("TPTW", g.NetworkParameters.TPTW); err != nil {
		return err
	}
	if err := writer.WriteElementString("QRxlevMin", g.NetworkParameters.QRxlevMin); err != nil {
		return err
	}
	if err := writer.WriteElementString("QRxlevMinCE", g.NetworkParameters.QRxlevMinCE); err != nil {
		return err
	}
	if err := writer.WriteElementString("QRxLevMinCE1", g.NetworkParameters.QRxLevMinCE1); err != nil {
		return err
	}
	if err := writer.WriteElementString("SignalQuality", g.QualityOfService.SignalQuality); err != nil {
		return err
	}
	if err := writer.WriteElementString("SignalLevel", g.QualityOfService.SignalLevel); err != nil {
		return err
	}
	if err := writer.WriteElementString("SignalToNoiseRatio", g.QualityOfService.SignalToNoiseRatio); err != nil {
		return err
	}
	return writer.WriteElementString("CoverageEnhancement", int(g.QualityOfService.CoverageEnhancement))
}

// PostLoad handles actions after Load.
func (g *GXDLMSLteMonitoring) PostLoad(reader *GXXmlReader) error {
	return nil
}

// GetValues returns the object attribute values.
func (g *GXDLMSLteMonitoring) GetValues() []any {
	return []any{g.LogicalName(), g.NetworkParameters, g.QualityOfService}
}

// NewGXDLMSLteMonitoring creates a new LTE Monitoring object instance.
func NewGXDLMSLteMonitoring(ln string, sn int16) (*GXDLMSLteMonitoring, error) {
	if err := ValidateLogicalName(ln); err != nil {
		return nil, err
	}
	return &GXDLMSLteMonitoring{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeLteMonitoring,
			logicalName: ln,
			ShortName:   sn,
			Version:     1,
		},
	}, nil
}
