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
// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSMBusPortSetup
type GXDLMSMBusPortSetup struct {
	GXDLMSObject

	// Reference to M-Bus communication port setup object.
	ProfileSelection string

	// Communication status of the M-Bus node.
	PortCommunicationStatus enums.MBusPortCommunicationState

	// M-Bus data header type.
	DataHeaderType enums.MBusDataHeaderType

	// The primary address of the M-Bus slave device.
	PrimaryAddress uint8

	// Identification Number element of the data header.
	IdentificationNumber uint32

	// Manufacturer Identification element.
	ManufacturerId uint16

	// M-Bus version.
	MBusVersion uint8

	// Device type.
	DeviceType enums.MBusDeviceType

	// Max PDU size.
	MaxPduSize uint16

	// Listening windows.
	ListeningWindow []types.GXKeyValuePair[types.GXDateTime, types.GXDateTime]
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSMBusPortSetup) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

func mbusPortToUInt32(value any) (uint32, error) {
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

func mbusPortAsArray(value any) (types.GXArray, error) {
	switch v := value.(type) {
	case types.GXArray:
		return v, nil
	case []any:
		return types.GXArray(v), nil
	default:
		return nil, fmt.Errorf("invalid array type: %T", value)
	}
}

func mbusPortAsStructure(value any) (types.GXStructure, error) {
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
func (g *GXDLMSMBusPortSetup) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, nil
}

// GetAttributeIndexToRead returns the collection of attributes to read.
func (g *GXDLMSMBusPortSetup) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	if all || g.CanRead(2) {
		attributes = append(attributes, 2)
	}
	if all || g.CanRead(3) {
		attributes = append(attributes, 3)
	}
	if all || g.CanRead(4) {
		attributes = append(attributes, 4)
	}
	if all || g.CanRead(5) {
		attributes = append(attributes, 5)
	}
	if all || g.CanRead(6) {
		attributes = append(attributes, 6)
	}
	if all || g.CanRead(7) {
		attributes = append(attributes, 7)
	}
	if all || g.CanRead(8) {
		attributes = append(attributes, 8)
	}
	if all || g.CanRead(9) {
		attributes = append(attributes, 9)
	}
	if all || g.CanRead(10) {
		attributes = append(attributes, 10)
	}
	if all || g.CanRead(11) {
		attributes = append(attributes, 11)
	}
	return attributes
}

// GetNames returns the names of attribute indexes.
func (g *GXDLMSMBusPortSetup) GetNames() []string {
	return []string{
		"Logical Name",
		"Profile selection",
		"Port communication status",
		"Data header type",
		"Primary address",
		"Identification number",
		"Manufacturer Id",
		"MBus version",
		"Device type",
		"Max PDU size",
		"Listening window",
	}
}

// GetMethodNames returns the names of method indexes.
func (g *GXDLMSMBusPortSetup) GetMethodNames() []string {
	return []string{}
}

// GetAttributeCount returns the amount of attributes.
func (g *GXDLMSMBusPortSetup) GetAttributeCount() int {
	return 11
}

// GetMethodCount returns the amount of methods.
func (g *GXDLMSMBusPortSetup) GetMethodCount() int {
	return 0
}

// GetDataType returns the device data type of selected attribute index.
func (g *GXDLMSMBusPortSetup) GetDataType(index int) (enums.DataType, error) {
	switch index {
	case 1:
		return enums.DataTypeOctetString, nil
	case 2:
		return enums.DataTypeOctetString, nil
	case 3:
		return enums.DataTypeEnum, nil
	case 4:
		return enums.DataTypeEnum, nil
	case 5:
		return enums.DataTypeUint8, nil
	case 6:
		return enums.DataTypeUint32, nil
	case 7:
		return enums.DataTypeUint16, nil
	case 8:
		return enums.DataTypeUint8, nil
	case 9:
		return enums.DataTypeUint8, nil
	case 10:
		return enums.DataTypeUint16, nil
	case 11:
		return enums.DataTypeArray, nil
	default:
		return 0, dlmserrors.ErrInvalidAttributeIndex
	}
}

// GetValue returns the value of given attribute.
func (g *GXDLMSMBusPortSetup) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	switch e.Index {
	case 1:
		return helpers.LogicalNameToBytes(g.LogicalName())
	case 2:
		return helpers.LogicalNameToBytes(g.ProfileSelection)
	case 3:
		return g.PortCommunicationStatus, nil
	case 4:
		return g.DataHeaderType, nil
	case 5:
		return g.PrimaryAddress, nil
	case 6:
		return g.IdentificationNumber, nil
	case 7:
		return g.ManufacturerId, nil
	case 8:
		return g.MBusVersion, nil
	case 9:
		return g.DeviceType, nil
	case 10:
		return g.MaxPduSize, nil
	case 11:
		data := types.NewGXByteBuffer()
		if err := data.SetUint8(uint8(enums.DataTypeArray)); err != nil {
			return nil, err
		}
		if err := types.SetObjectCount(len(g.ListeningWindow), data); err != nil {
			return nil, err
		}
		for _, it := range g.ListeningWindow {
			if err := data.SetUint8(uint8(enums.DataTypeStructure)); err != nil {
				return nil, err
			}
			if err := data.SetUint8(2); err != nil {
				return nil, err
			}
			if err := internal.SetData(settings, data, enums.DataTypeOctetString, it.Key); err != nil {
				return nil, err
			}
			if err := internal.SetData(settings, data, enums.DataTypeOctetString, it.Value); err != nil {
				return nil, err
			}
		}
		return data.Array(), nil
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
		return nil, nil
	}
}

// SetValue sets value of given attribute.
func (g *GXDLMSMBusPortSetup) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
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
			e.Error = enums.ErrorCodeReadWriteDenied
			return err
		}
		g.ProfileSelection = ln
	case 3:
		v, err := mbusPortToUInt32(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return err
		}
		g.PortCommunicationStatus = enums.MBusPortCommunicationState(v)
	case 4:
		v, err := mbusPortToUInt32(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return err
		}
		g.DataHeaderType = enums.MBusDataHeaderType(v)
	case 5:
		v, err := mbusPortToUInt32(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return err
		}
		g.PrimaryAddress = uint8(v)
	case 6:
		v, err := mbusPortToUInt32(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return err
		}
		g.IdentificationNumber = v
	case 7:
		v, err := mbusPortToUInt32(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return err
		}
		g.ManufacturerId = uint16(v)
	case 8:
		v, err := mbusPortToUInt32(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return err
		}
		g.MBusVersion = uint8(v)
	case 9:
		v, err := mbusPortToUInt32(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return err
		}
		g.DeviceType = enums.MBusDeviceType(v)
	case 10:
		v, err := mbusPortToUInt32(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return err
		}
		g.MaxPduSize = uint16(v)
	case 11:
		g.ListeningWindow = g.ListeningWindow[:0]
		if e.Value == nil {
			return nil
		}
		arr, err := mbusPortAsArray(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return err
		}
		for _, tmp := range arr {
			item, err := mbusPortAsStructure(tmp)
			if err != nil {
				e.Error = enums.ErrorCodeReadWriteDenied
				return err
			}
			startAny, err := internal.ChangeTypeFromByteArray(settings, item[0].([]byte), enums.DataTypeDateTime)
			if err != nil {
				return err
			}
			endAny, err := internal.ChangeTypeFromByteArray(settings, item[1].([]byte), enums.DataTypeDateTime)
			if err != nil {
				return err
			}
			g.ListeningWindow = append(g.ListeningWindow, *types.NewGXKeyValuePair(startAny.(types.GXDateTime), endAny.(types.GXDateTime)))
		}
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil
}

// Load loads object content from XML.
func (g *GXDLMSMBusPortSetup) Load(reader *GXXmlReader) error {
	var err error
	g.ProfileSelection, err = reader.ReadElementContentAsString("ProfileSelection", "")
	if err != nil {
		return err
	}
	ret, err := reader.ReadElementContentAsInt("Status", 0)
	if err != nil {
		return err
	}
	g.PortCommunicationStatus = enums.MBusPortCommunicationState(ret)
	ret, err = reader.ReadElementContentAsInt("DataHeaderType", 0)
	if err != nil {
		return err
	}
	g.DataHeaderType = enums.MBusDataHeaderType(ret)
	g.PrimaryAddress, err = reader.ReadElementContentAsUInt8("PrimaryAddress", 0)
	if err != nil {
		return err
	}
	g.IdentificationNumber, err = reader.ReadElementContentAsUInt32("IdentificationNumber", 0)
	if err != nil {
		return err
	}
	g.ManufacturerId, err = reader.ReadElementContentAsUInt16("ManufacturerId", 0)
	if err != nil {
		return err
	}
	g.MBusVersion, err = reader.ReadElementContentAsUInt8("Version", 0)
	if err != nil {
		return err
	}
	ret, err = reader.ReadElementContentAsInt("DeviceType", 0)
	if err != nil {
		return err
	}
	g.DeviceType = enums.MBusDeviceType(ret)
	g.MaxPduSize, err = reader.ReadElementContentAsUInt16("MaxPduSize", 0)
	if err != nil {
		return err
	}
	g.ListeningWindow = g.ListeningWindow[:0]
	if ok, err := reader.IsStartElementNamed("ListeningWindow", true); ok && err == nil {
		for {
			ok, err = reader.IsStartElementNamed("Item", true)
			if err != nil {
				return err
			}
			if !ok {
				break
			}
			start, err := reader.ReadElementContentAsDateTime("Start", nil)
			if err != nil {
				return err
			}
			end, err := reader.ReadElementContentAsDateTime("End", nil)
			if err != nil {
				return err
			}
			g.ListeningWindow = append(g.ListeningWindow, *types.NewGXKeyValuePair(start, end))
		}
		reader.ReadEndElement("ListeningWindow")
	}
	return nil
}

// Save saves object content to XML.
func (g *GXDLMSMBusPortSetup) Save(writer *GXXmlWriter) error {
	if err := writer.WriteElementString("ProfileSelection", g.ProfileSelection); err != nil {
		return err
	}
	if err := writer.WriteElementString("Status", int(g.PortCommunicationStatus)); err != nil {
		return err
	}
	if err := writer.WriteElementString("DataHeaderType", int(g.DataHeaderType)); err != nil {
		return err
	}
	if err := writer.WriteElementString("PrimaryAddress", g.PrimaryAddress); err != nil {
		return err
	}
	if err := writer.WriteElementString("IdentificationNumber", g.IdentificationNumber); err != nil {
		return err
	}
	if err := writer.WriteElementString("ManufacturerId", g.ManufacturerId); err != nil {
		return err
	}
	if err := writer.WriteElementString("Version", g.MBusVersion); err != nil {
		return err
	}
	if err := writer.WriteElementString("DeviceType", int(g.DeviceType)); err != nil {
		return err
	}
	if err := writer.WriteElementString("MaxPduSize", g.MaxPduSize); err != nil {
		return err
	}
	writer.WriteStartElement("ListeningWindow")
	if g.ListeningWindow != nil {
		for _, it := range g.ListeningWindow {
			writer.WriteStartElement("Item")
			if err := writer.WriteElementString("Start", it.Key); err != nil {
				return err
			}
			if err := writer.WriteElementString("End", it.Value); err != nil {
				return err
			}
			writer.WriteEndElement()
		}
	}
	writer.WriteEndElement()
	return nil
}

// PostLoad handles actions after Load.
func (g *GXDLMSMBusPortSetup) PostLoad(reader *GXXmlReader) error {
	return nil
}

// GetValues returns the object attribute values.
func (g *GXDLMSMBusPortSetup) GetValues() []any {
	return []any{
		g.LogicalName(),
		g.ProfileSelection,
		g.PortCommunicationStatus,
		g.DataHeaderType,
		g.PrimaryAddress,
		g.IdentificationNumber,
		g.ManufacturerId,
		g.MBusVersion,
		g.DeviceType,
		g.MaxPduSize,
		g.ListeningWindow,
	}
}

// NewGXDLMSMBusPortSetup creates a new M-Bus Port Setup object instance.
func NewGXDLMSMBusPortSetup(ln string, sn int16) (*GXDLMSMBusPortSetup, error) {
	if err := ValidateLogicalName(ln); err != nil {
		return nil, err
	}
	return &GXDLMSMBusPortSetup{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeMBusPortSetup,
			logicalName: ln,
			ShortName:   sn,
		},
		ListeningWindow: make([]types.GXKeyValuePair[types.GXDateTime, types.GXDateTime], 0),
	}, nil
}
