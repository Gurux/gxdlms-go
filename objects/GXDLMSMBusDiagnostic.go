package objects

import (
	"fmt"
	"time"

	"github.com/Gurux/gxdlms-go/dlmserrors"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// Online help:
// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSMBusDiagnostic
type GXDLMSMBusDiagnostic struct {
	GXDLMSObject

	// Received signal strength in dBm.
	ReceivedSignalStrength uint8

	// Currently used channel ID.
	ChannelId uint8

	// Link status.
	LinkStatus enums.MBusLinkStatus

	// Broadcast frame counters.
	BroadcastFrames []GXBroadcastFrameCounter

	// Transmitted frames.
	Transmissions uint32

	// Received frames with a correct checksum.
	ReceivedFrames uint32

	// Received frames with an incorrect checksum.
	FailedReceivedFrames uint32

	// Last changed value.
	CaptureTime GXCaptureTime
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSMBusDiagnostic) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

func mbusDiagToUInt32(value any) (uint32, error) {
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

func mbusDiagAsArray(value any) (types.GXArray, error) {
	switch v := value.(type) {
	case types.GXArray:
		return v, nil
	case []any:
		return types.GXArray(v), nil
	default:
		return nil, fmt.Errorf("invalid array type: %T", value)
	}
}

func mbusDiagAsStructure(value any) (types.GXStructure, error) {
	switch v := value.(type) {
	case types.GXStructure:
		return v, nil
	case []any:
		return types.GXStructure(v), nil
	default:
		return nil, fmt.Errorf("invalid structure type: %T", value)
	}
}

func mbusDiagToDateTime(value any, dlmsSettings *settings.GXDLMSSettings) (types.GXDateTime, error) {
	if value == nil {
		return types.GXDateTime{}, nil
	}
	switch v := value.(type) {
	case types.GXDateTime:
		return v, nil
	case []byte:
		ret, err := internal.ChangeTypeFromByteArray(dlmsSettings, v, enums.DataTypeDateTime)
		if err != nil {
			return types.GXDateTime{}, err
		}
		return ret.(types.GXDateTime), nil
	case string:
		dt, err := types.NewGXDateTimeFromString(v, nil)
		if err != nil {
			return types.GXDateTime{}, err
		}
		return *dt, nil
	case time.Time:
		return *types.NewGXDateTimeFromTime(v), nil
	default:
		return types.GXDateTime{}, fmt.Errorf("invalid datetime type: %T", value)
	}
}

// Reset resets the diagnostic counters.
func (g *GXDLMSMBusDiagnostic) Reset(client IGXDLMSClient) ([][]byte, error) {
	return client.Method(g, 1, int8(0), enums.DataTypeInt8)
}

// Invoke invokes method.
func (g *GXDLMSMBusDiagnostic) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	if e.Index == 1 {
		g.ReceivedSignalStrength = 0
		g.Transmissions = 0
		g.ReceivedFrames = 0
		g.FailedReceivedFrames = 0
		g.CaptureTime = GXCaptureTime{}
		return nil, nil
	}
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, nil
}

// GetAttributeIndexToRead returns the collection of attributes to read.
func (g *GXDLMSMBusDiagnostic) GetAttributeIndexToRead(all bool) []int {
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
	return attributes
}

// GetNames returns the names of attribute indexes.
func (g *GXDLMSMBusDiagnostic) GetNames() []string {
	return []string{
		"Logical Name",
		"Received signal strength",
		"Channel Id",
		"Link status",
		"Broadcast frames",
		"Transmissions",
		"Received frames",
		"Failed received frames",
		"Capture time",
	}
}

// GetMethodNames returns the names of method indexes.
func (g *GXDLMSMBusDiagnostic) GetMethodNames() []string {
	return []string{"Reset"}
}

// GetAttributeCount returns the amount of attributes.
func (g *GXDLMSMBusDiagnostic) GetAttributeCount() int {
	return 9
}

// GetMethodCount returns the amount of methods.
func (g *GXDLMSMBusDiagnostic) GetMethodCount() int {
	return 1
}

// GetDataType returns the device data type of selected attribute index.
func (g *GXDLMSMBusDiagnostic) GetDataType(index int) (enums.DataType, error) {
	switch index {
	case 1:
		return enums.DataTypeOctetString, nil
	case 2:
		return enums.DataTypeUint8, nil
	case 3:
		return enums.DataTypeUint8, nil
	case 4:
		return enums.DataTypeEnum, nil
	case 5:
		return enums.DataTypeArray, nil
	case 6:
		return enums.DataTypeUint32, nil
	case 7:
		return enums.DataTypeUint32, nil
	case 8:
		return enums.DataTypeUint32, nil
	case 9:
		return enums.DataTypeStructure, nil
	default:
		return 0, dlmserrors.ErrInvalidAttributeIndex
	}
}

// GetValue returns the value of given attribute.
func (g *GXDLMSMBusDiagnostic) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	switch e.Index {
	case 1:
		return helpers.LogicalNameToBytes(g.LogicalName())
	case 2:
		return g.ReceivedSignalStrength, nil
	case 3:
		return g.ChannelId, nil
	case 4:
		return g.LinkStatus, nil
	case 5:
		data := types.NewGXByteBuffer()
		if err := data.SetUint8(uint8(enums.DataTypeArray)); err != nil {
			return nil, err
		}
		if err := types.SetObjectCount(len(g.BroadcastFrames), data); err != nil {
			return nil, err
		}
		for _, it := range g.BroadcastFrames {
			if err := data.SetUint8(uint8(enums.DataTypeStructure)); err != nil {
				return nil, err
			}
			if err := data.SetUint8(3); err != nil {
				return nil, err
			}
			if err := internal.SetData(settings, data, enums.DataTypeUint8, it.ClientID); err != nil {
				return nil, err
			}
			if err := internal.SetData(settings, data, enums.DataTypeUint32, it.Counter); err != nil {
				return nil, err
			}
			if err := internal.SetData(settings, data, enums.DataTypeDateTime, it.TimeStamp); err != nil {
				return nil, err
			}
		}
		return data.Array(), nil
	case 6:
		return g.Transmissions, nil
	case 7:
		return g.ReceivedFrames, nil
	case 8:
		return g.FailedReceivedFrames, nil
	case 9:
		data := types.NewGXByteBuffer()
		if err := data.SetUint8(uint8(enums.DataTypeStructure)); err != nil {
			return nil, err
		}
		if err := types.SetObjectCount(2, data); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, data, enums.DataTypeUint8, g.CaptureTime.AttributeID); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, data, enums.DataTypeDateTime, g.CaptureTime.TimeStamp); err != nil {
			return nil, err
		}
		return data.Array(), nil
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
		return nil, nil
	}
}

// SetValue sets value of given attribute.
func (g *GXDLMSMBusDiagnostic) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	switch e.Index {
	case 1:
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return err
		}
		return g.SetLogicalName(ln)
	case 2:
		v, err := mbusDiagToUInt32(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return err
		}
		g.ReceivedSignalStrength = uint8(v)
	case 3:
		v, err := mbusDiagToUInt32(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return err
		}
		g.ChannelId = uint8(v)
	case 4:
		v, err := mbusDiagToUInt32(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return err
		}
		g.LinkStatus = enums.MBusLinkStatus(v)
	case 5:
		g.BroadcastFrames = g.BroadcastFrames[:0]
		if e.Value == nil {
			return nil
		}
		arr, err := mbusDiagAsArray(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return err
		}
		for _, tmp := range arr {
			item, err := mbusDiagAsStructure(tmp)
			if err != nil {
				e.Error = enums.ErrorCodeReadWriteDenied
				return err
			}
			clientID, err := mbusDiagToUInt32(item[0])
			if err != nil {
				return err
			}
			counter, err := mbusDiagToUInt32(item[1])
			if err != nil {
				return err
			}
			ts, err := mbusDiagToDateTime(item[2], settings)
			if err != nil {
				return err
			}
			g.BroadcastFrames = append(g.BroadcastFrames, GXBroadcastFrameCounter{
				ClientID:  uint8(clientID),
				Counter:   counter,
				TimeStamp: ts,
			})
		}
	case 6:
		v, err := mbusDiagToUInt32(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return err
		}
		g.Transmissions = v
	case 7:
		v, err := mbusDiagToUInt32(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return err
		}
		g.ReceivedFrames = v
	case 8:
		v, err := mbusDiagToUInt32(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return err
		}
		g.FailedReceivedFrames = v
	case 9:
		if e.Value == nil {
			g.CaptureTime = GXCaptureTime{}
			return nil
		}
		item, err := mbusDiagAsStructure(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return err
		}
		attributeID, err := mbusDiagToUInt32(item[0])
		if err != nil {
			return err
		}
		ts, err := mbusDiagToDateTime(item[1], settings)
		if err != nil {
			return err
		}
		g.CaptureTime.AttributeID = uint8(attributeID)
		g.CaptureTime.TimeStamp = ts
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil
}

// Load loads object content from XML.
func (g *GXDLMSMBusDiagnostic) Load(reader *GXXmlReader) error {
	var err error
	g.ReceivedSignalStrength, err = reader.ReadElementContentAsUInt8("ReceivedSignalStrength", 0)
	if err != nil {
		return err
	}
	g.ChannelId, err = reader.ReadElementContentAsUInt8("ChannelId", 0)
	if err != nil {
		return err
	}
	ret, err := reader.ReadElementContentAsInt("LinkStatus", 0)
	if err != nil {
		return err
	}
	g.LinkStatus = enums.MBusLinkStatus(ret)
	g.BroadcastFrames = g.BroadcastFrames[:0]
	if ok, err := reader.IsStartElementNamed("BroadcastFrames", true); ok && err == nil {
		for {
			ok, err = reader.IsStartElementNamed("Item", true)
			if err != nil {
				return err
			}
			if !ok {
				break
			}
			clientID, err := reader.ReadElementContentAsUInt8("ClientId", 0)
			if err != nil {
				return err
			}
			counter, err := reader.ReadElementContentAsUInt32("Counter", 0)
			if err != nil {
				return err
			}
			timeStamp, err := reader.ReadElementContentAsDateTime("TimeStamp", nil)
			if err != nil {
				return err
			}
			g.BroadcastFrames = append(g.BroadcastFrames, GXBroadcastFrameCounter{
				ClientID:  clientID,
				Counter:   counter,
				TimeStamp: timeStamp,
			})
		}
		reader.ReadEndElement("BroadcastFrames")
	}
	g.Transmissions, err = reader.ReadElementContentAsUInt32("Transmissions", 0)
	if err != nil {
		return err
	}
	g.ReceivedFrames, err = reader.ReadElementContentAsUInt32("ReceivedFrames", 0)
	if err != nil {
		return err
	}
	g.FailedReceivedFrames, err = reader.ReadElementContentAsUInt32("FailedReceivedFrames", 0)
	if err != nil {
		return err
	}
	if ok, err := reader.IsStartElementNamed("CaptureTime", true); ok && err == nil {
		g.CaptureTime.AttributeID, err = reader.ReadElementContentAsUInt8("AttributeId", 0)
		if err != nil {
			return err
		}
		g.CaptureTime.TimeStamp, err = reader.ReadElementContentAsDateTime("TimeStamp", nil)
		if err != nil {
			return err
		}
		reader.ReadEndElement("CaptureTime")
	}
	return nil
}

// Save saves object content to XML.
func (g *GXDLMSMBusDiagnostic) Save(writer *GXXmlWriter) error {
	if err := writer.WriteElementString("ReceivedSignalStrength", g.ReceivedSignalStrength); err != nil {
		return err
	}
	if err := writer.WriteElementString("ChannelId", g.ChannelId); err != nil {
		return err
	}
	if err := writer.WriteElementString("LinkStatus", int(g.LinkStatus)); err != nil {
		return err
	}
	writer.WriteStartElement("BroadcastFrames")
	if g.BroadcastFrames != nil {
		for _, it := range g.BroadcastFrames {
			writer.WriteStartElement("Item")
			if err := writer.WriteElementString("ClientId", it.ClientID); err != nil {
				return err
			}
			if err := writer.WriteElementString("Counter", it.Counter); err != nil {
				return err
			}
			if err := writer.WriteElementString("TimeStamp", it.TimeStamp); err != nil {
				return err
			}
			writer.WriteEndElement()
		}
	}
	writer.WriteEndElement()
	if err := writer.WriteElementString("Transmissions", g.Transmissions); err != nil {
		return err
	}
	if err := writer.WriteElementString("ReceivedFrames", g.ReceivedFrames); err != nil {
		return err
	}
	if err := writer.WriteElementString("FailedReceivedFrames", g.FailedReceivedFrames); err != nil {
		return err
	}
	writer.WriteStartElement("CaptureTime")
	if err := writer.WriteElementString("AttributeId", g.CaptureTime.AttributeID); err != nil {
		return err
	}
	if err := writer.WriteElementString("TimeStamp", g.CaptureTime.TimeStamp); err != nil {
		return err
	}
	writer.WriteEndElement()
	return nil
}

// PostLoad handles actions after Load.
func (g *GXDLMSMBusDiagnostic) PostLoad(reader *GXXmlReader) error {
	return nil
}

// GetValues returns the object attribute values.
func (g *GXDLMSMBusDiagnostic) GetValues() []any {
	return []any{
		g.LogicalName(),
		g.ReceivedSignalStrength,
		g.ChannelId,
		g.LinkStatus,
		g.BroadcastFrames,
		g.Transmissions,
		g.ReceivedFrames,
		g.FailedReceivedFrames,
		g.CaptureTime,
	}
}

// NewGXDLMSMBusDiagnostic creates a new M-Bus Diagnostic object instance.
func NewGXDLMSMBusDiagnostic(ln string, sn int16) (*GXDLMSMBusDiagnostic, error) {
	if err := ValidateLogicalName(ln); err != nil {
		return nil, err
	}
	return &GXDLMSMBusDiagnostic{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeMBusDiagnostic,
			logicalName: ln,
			ShortName:   sn,
		},
		BroadcastFrames: make([]GXBroadcastFrameCounter, 0),
	}, nil
}
