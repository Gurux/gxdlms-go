package objects

import (
	"github.com/Gurux/gxdlms-go/dlmserrors"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// GXCoapMessagesCounter contains CoAP message counters.
type GXCoapMessagesCounter struct {
	Tx               uint32
	Rx               uint32
	TxResend         uint32
	TxReset          uint32
	RxReset          uint32
	TxAck            uint32
	RxAck            uint32
	RxDrop           uint32
	TxNonPiggybacked uint32
	MaxRtxExceeded   uint32
}

// GXCoapRequestResponseCounter contains CoAP request/response counters.
type GXCoapRequestResponseCounter struct {
	RxRequests    uint32
	TxRequests    uint32
	RxResponse    uint32
	TxResponse    uint32
	TxClientError uint32
	RxClientError uint32
	TxServerError uint32
	RxServerError uint32
}

// GXCoapBtCounter contains block transfer counters.
type GXCoapBtCounter struct {
	BlockWiseTransferStarted   uint32
	BlockWiseTransferCompleted uint32
	BlockWiseTransferTimeout   uint32
}

// GXCoapCaptureTime contains capture timestamp and source attribute ID.
type GXCoapCaptureTime struct {
	AttributeID uint8
	TimeStamp   types.GXDateTime
}

// GXDLMSCoAPDiagnostic contains CoAP diagnostic counters.
type GXDLMSCoAPDiagnostic struct {
	GXDLMSObject
	MessagesCounter        GXCoapMessagesCounter
	RequestResponseCounter GXCoapRequestResponseCounter
	BtCounter              GXCoapBtCounter
	CaptureTime            GXCoapCaptureTime
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSCoAPDiagnostic) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

func (g *GXDLMSCoAPDiagnostic) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, nil
}

// Reset resets diagnostic values in the server.
func (g *GXDLMSCoAPDiagnostic) Reset(client IGXDLMSClient) ([][]byte, error) {
	return client.Method(g, 1, int8(0), enums.DataTypeInt8)
}

func (g *GXDLMSCoAPDiagnostic) GetAttributeIndexToRead(all bool) []int {
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

func (g *GXDLMSCoAPDiagnostic) GetNames() []string {
	return []string{"Logical Name", "Messages counter", "Request response counter", "BT counter", "Capture time"}
}

func (g *GXDLMSCoAPDiagnostic) GetMethodNames() []string { return []string{"Reset"} }

func (g *GXDLMSCoAPDiagnostic) GetAttributeCount() int { return 5 }

func (g *GXDLMSCoAPDiagnostic) GetMethodCount() int { return 1 }

func (g *GXDLMSCoAPDiagnostic) GetDataType(index int) (enums.DataType, error) {
	switch index {
	case 1:
		return enums.DataTypeOctetString, nil
	case 2, 3, 4, 5:
		return enums.DataTypeStructure, nil
	default:
		return enums.DataTypeNone, dlmserrors.ErrInvalidAttributeIndex
	}
}

func (g *GXDLMSCoAPDiagnostic) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	buff := types.NewGXByteBuffer()
	switch e.Index {
	case 1:
		return helpers.LogicalNameToBytes(g.LogicalName())
	case 2:
		_ = buff.SetUint8(uint8(enums.DataTypeStructure))
		types.SetObjectCount(10, buff)
		_ = internal.SetData(settings, buff, enums.DataTypeUint32, g.MessagesCounter.Tx)
		_ = internal.SetData(settings, buff, enums.DataTypeUint32, g.MessagesCounter.Rx)
		_ = internal.SetData(settings, buff, enums.DataTypeUint32, g.MessagesCounter.TxResend)
		_ = internal.SetData(settings, buff, enums.DataTypeUint32, g.MessagesCounter.TxReset)
		_ = internal.SetData(settings, buff, enums.DataTypeUint32, g.MessagesCounter.RxReset)
		_ = internal.SetData(settings, buff, enums.DataTypeUint32, g.MessagesCounter.TxAck)
		_ = internal.SetData(settings, buff, enums.DataTypeUint32, g.MessagesCounter.RxAck)
		_ = internal.SetData(settings, buff, enums.DataTypeUint32, g.MessagesCounter.RxDrop)
		_ = internal.SetData(settings, buff, enums.DataTypeUint32, g.MessagesCounter.TxNonPiggybacked)
		_ = internal.SetData(settings, buff, enums.DataTypeUint32, g.MessagesCounter.MaxRtxExceeded)
		return buff.Array(), nil
	case 3:
		_ = buff.SetUint8(uint8(enums.DataTypeStructure))
		types.SetObjectCount(8, buff)
		_ = internal.SetData(settings, buff, enums.DataTypeUint32, g.RequestResponseCounter.RxRequests)
		_ = internal.SetData(settings, buff, enums.DataTypeUint32, g.RequestResponseCounter.TxRequests)
		_ = internal.SetData(settings, buff, enums.DataTypeUint32, g.RequestResponseCounter.RxResponse)
		_ = internal.SetData(settings, buff, enums.DataTypeUint32, g.RequestResponseCounter.TxResponse)
		_ = internal.SetData(settings, buff, enums.DataTypeUint32, g.RequestResponseCounter.TxClientError)
		_ = internal.SetData(settings, buff, enums.DataTypeUint32, g.RequestResponseCounter.RxClientError)
		_ = internal.SetData(settings, buff, enums.DataTypeUint32, g.RequestResponseCounter.TxServerError)
		_ = internal.SetData(settings, buff, enums.DataTypeUint32, g.RequestResponseCounter.RxServerError)
		return buff.Array(), nil
	case 4:
		_ = buff.SetUint8(uint8(enums.DataTypeStructure))
		types.SetObjectCount(3, buff)
		_ = internal.SetData(settings, buff, enums.DataTypeUint32, g.BtCounter.BlockWiseTransferStarted)
		_ = internal.SetData(settings, buff, enums.DataTypeUint32, g.BtCounter.BlockWiseTransferCompleted)
		_ = internal.SetData(settings, buff, enums.DataTypeUint32, g.BtCounter.BlockWiseTransferTimeout)
		return buff.Array(), nil
	case 5:
		_ = buff.SetUint8(uint8(enums.DataTypeStructure))
		types.SetObjectCount(2, buff)
		_ = internal.SetData(settings, buff, enums.DataTypeUint8, g.CaptureTime.AttributeID)
		_ = internal.SetData(settings, buff, enums.DataTypeDateTime, g.CaptureTime.TimeStamp)
		return buff.Array(), nil
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
		return nil, nil
	}
}

func (g *GXDLMSCoAPDiagnostic) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	switch e.Index {
	case 1:
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return err
		}
		return g.SetLogicalName(ln)
	case 2:
		s := e.Value.(types.GXStructure)
		if len(s) < 10 {
			return nil
		}
		g.MessagesCounter.Tx, _ = toUint32(s[0])
		g.MessagesCounter.Rx, _ = toUint32(s[1])
		g.MessagesCounter.TxResend, _ = toUint32(s[2])
		g.MessagesCounter.TxReset, _ = toUint32(s[3])
		g.MessagesCounter.RxReset, _ = toUint32(s[4])
		g.MessagesCounter.TxAck, _ = toUint32(s[5])
		g.MessagesCounter.RxAck, _ = toUint32(s[6])
		g.MessagesCounter.RxDrop, _ = toUint32(s[7])
		g.MessagesCounter.TxNonPiggybacked, _ = toUint32(s[8])
		g.MessagesCounter.MaxRtxExceeded, _ = toUint32(s[9])
	case 3:
		s := e.Value.(types.GXStructure)
		if len(s) < 8 {
			return nil
		}
		g.RequestResponseCounter.RxRequests, _ = toUint32(s[0])
		g.RequestResponseCounter.TxRequests, _ = toUint32(s[1])
		g.RequestResponseCounter.RxResponse, _ = toUint32(s[2])
		g.RequestResponseCounter.TxResponse, _ = toUint32(s[3])
		g.RequestResponseCounter.TxClientError, _ = toUint32(s[4])
		g.RequestResponseCounter.RxClientError, _ = toUint32(s[5])
		g.RequestResponseCounter.TxServerError, _ = toUint32(s[6])
		g.RequestResponseCounter.RxServerError, _ = toUint32(s[7])
	case 4:
		s := e.Value.(types.GXStructure)
		if len(s) < 3 {
			return nil
		}
		g.BtCounter.BlockWiseTransferStarted, _ = toUint32(s[0])
		g.BtCounter.BlockWiseTransferCompleted, _ = toUint32(s[1])
		g.BtCounter.BlockWiseTransferTimeout, _ = toUint32(s[2])
	case 5:
		s := e.Value.(types.GXStructure)
		if len(s) < 2 {
			return nil
		}
		g.CaptureTime.AttributeID = s[0].(uint8)
		if ts, ok := s[1].(types.GXDateTime); ok {
			g.CaptureTime.TimeStamp = ts
		}
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil
}

func (g *GXDLMSCoAPDiagnostic) Load(reader *GXXmlReader) error {
	var err error
	g.MessagesCounter.Tx, err = reader.ReadElementContentAsUInt32("Tx", 0)
	if err != nil {
		return err
	}
	g.MessagesCounter.Rx, err = reader.ReadElementContentAsUInt32("Rx", 0)
	if err != nil {
		return err
	}
	g.MessagesCounter.TxResend, err = reader.ReadElementContentAsUInt32("TxResend", 0)
	if err != nil {
		return err
	}
	g.MessagesCounter.TxReset, err = reader.ReadElementContentAsUInt32("TxReset", 0)
	if err != nil {
		return err
	}
	g.MessagesCounter.RxReset, err = reader.ReadElementContentAsUInt32("RxReset", 0)
	if err != nil {
		return err
	}
	g.MessagesCounter.TxAck, err = reader.ReadElementContentAsUInt32("TxAck", 0)
	if err != nil {
		return err
	}
	g.MessagesCounter.RxAck, err = reader.ReadElementContentAsUInt32("RxAck", 0)
	if err != nil {
		return err
	}
	g.MessagesCounter.RxDrop, err = reader.ReadElementContentAsUInt32("RxDrop", 0)
	if err != nil {
		return err
	}
	g.MessagesCounter.TxNonPiggybacked, err = reader.ReadElementContentAsUInt32("TxNonPiggybacked", 0)
	if err != nil {
		return err
	}
	g.MessagesCounter.MaxRtxExceeded, err = reader.ReadElementContentAsUInt32("MaxRtxExceeded", 0)
	if err != nil {
		return err
	}
	g.RequestResponseCounter.RxRequests, err = reader.ReadElementContentAsUInt32("RxRequests", 0)
	if err != nil {
		return err
	}
	g.RequestResponseCounter.TxRequests, err = reader.ReadElementContentAsUInt32("TxRequests", 0)
	if err != nil {
		return err
	}
	g.RequestResponseCounter.RxResponse, err = reader.ReadElementContentAsUInt32("RxResponse", 0)
	if err != nil {
		return err
	}
	g.RequestResponseCounter.TxResponse, err = reader.ReadElementContentAsUInt32("TxResponse", 0)
	if err != nil {
		return err
	}
	g.RequestResponseCounter.TxClientError, err = reader.ReadElementContentAsUInt32("TxClientError", 0)
	if err != nil {
		return err
	}
	g.RequestResponseCounter.RxClientError, err = reader.ReadElementContentAsUInt32("RxClientError", 0)
	if err != nil {
		return err
	}
	g.RequestResponseCounter.TxServerError, err = reader.ReadElementContentAsUInt32("TxServerError", 0)
	if err != nil {
		return err
	}
	g.RequestResponseCounter.RxServerError, err = reader.ReadElementContentAsUInt32("RxServerError", 0)
	if err != nil {
		return err
	}
	g.BtCounter.BlockWiseTransferStarted, err = reader.ReadElementContentAsUInt32("TransferStarted", 0)
	if err != nil {
		return err
	}
	g.BtCounter.BlockWiseTransferCompleted, err = reader.ReadElementContentAsUInt32("TransferCompleted", 0)
	if err != nil {
		return err
	}
	g.BtCounter.BlockWiseTransferTimeout, err = reader.ReadElementContentAsUInt32("TransferTimeout", 0)
	if err != nil {
		return err
	}
	g.CaptureTime.AttributeID, err = reader.ReadElementContentAsUInt8("AttributeId", 0)
	if err != nil {
		return err
	}
	g.CaptureTime.TimeStamp, err = reader.ReadElementContentAsDateTime("TimeStamp", nil)
	return err
}

func (g *GXDLMSCoAPDiagnostic) Save(writer *GXXmlWriter) error {
	if err := writer.WriteElementString("Tx", g.MessagesCounter.Tx); err != nil {
		return err
	}
	if err := writer.WriteElementString("Rx", g.MessagesCounter.Rx); err != nil {
		return err
	}
	if err := writer.WriteElementString("TxResend", g.MessagesCounter.TxResend); err != nil {
		return err
	}
	if err := writer.WriteElementString("TxReset", g.MessagesCounter.TxReset); err != nil {
		return err
	}
	if err := writer.WriteElementString("RxReset", g.MessagesCounter.RxReset); err != nil {
		return err
	}
	if err := writer.WriteElementString("TxAck", g.MessagesCounter.TxAck); err != nil {
		return err
	}
	if err := writer.WriteElementString("RxAck", g.MessagesCounter.RxAck); err != nil {
		return err
	}
	if err := writer.WriteElementString("RxDrop", g.MessagesCounter.RxDrop); err != nil {
		return err
	}
	if err := writer.WriteElementString("TxNonPiggybacked", g.MessagesCounter.TxNonPiggybacked); err != nil {
		return err
	}
	if err := writer.WriteElementString("MaxRtxExceeded", g.MessagesCounter.MaxRtxExceeded); err != nil {
		return err
	}
	if err := writer.WriteElementString("RxRequests", g.RequestResponseCounter.RxRequests); err != nil {
		return err
	}
	if err := writer.WriteElementString("TxRequests", g.RequestResponseCounter.TxRequests); err != nil {
		return err
	}
	if err := writer.WriteElementString("RxResponse", g.RequestResponseCounter.RxResponse); err != nil {
		return err
	}
	if err := writer.WriteElementString("TxResponse", g.RequestResponseCounter.TxResponse); err != nil {
		return err
	}
	if err := writer.WriteElementString("TxClientError", g.RequestResponseCounter.TxClientError); err != nil {
		return err
	}
	if err := writer.WriteElementString("RxClientError", g.RequestResponseCounter.RxClientError); err != nil {
		return err
	}
	if err := writer.WriteElementString("TxServerError", g.RequestResponseCounter.TxServerError); err != nil {
		return err
	}
	if err := writer.WriteElementString("RxServerError", g.RequestResponseCounter.RxServerError); err != nil {
		return err
	}
	if err := writer.WriteElementString("TransferStarted", g.BtCounter.BlockWiseTransferStarted); err != nil {
		return err
	}
	if err := writer.WriteElementString("TransferCompleted", g.BtCounter.BlockWiseTransferCompleted); err != nil {
		return err
	}
	if err := writer.WriteElementString("TransferTimeout", g.BtCounter.BlockWiseTransferTimeout); err != nil {
		return err
	}
	if err := writer.WriteElementString("AttributeId", g.CaptureTime.AttributeID); err != nil {
		return err
	}
	return writer.WriteElementString("TimeStamp", g.CaptureTime.TimeStamp)
}

func (g *GXDLMSCoAPDiagnostic) PostLoad(reader *GXXmlReader) error { return nil }

func (g *GXDLMSCoAPDiagnostic) GetValues() []any {
	return []any{
		g.LogicalName(),
		g.MessagesCounter,
		g.RequestResponseCounter,
		g.BtCounter,
		g.CaptureTime,
	}
}

// NewGXDLMSCoAPDiagnostic creates a new CoAP diagnostic object instance.
func NewGXDLMSCoAPDiagnostic(ln string, sn int16) (*GXDLMSCoAPDiagnostic, error) {
	if err := ValidateLogicalName(ln); err != nil {
		return nil, err
	}
	return &GXDLMSCoAPDiagnostic{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeCoAPDiagnostic,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
