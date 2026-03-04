package objects

import (
	"bytes"
	"errors"
	"strings"

	"github.com/Gurux/gxcommon-go"
	"github.com/Gurux/gxdlms-go/dlmserrors"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// Online help:
// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSImageTransfer
type GXDLMSImageTransfer struct {
	GXDLMSObject

	ImageBlockSize                   uint32
	ImageTransferredBlocksStatus     string
	ImageFirstNotTransferredBlockNum uint32
	ImageTransferEnabled             bool
	ImageTransferStatus              enums.ImageTransferStatus
	ImageActivateInfo                []GXDLMSImageActivateInfo
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSImageTransfer) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

// Invoke invokes object methods.
func (g *GXDLMSImageTransfer) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	switch e.Index {
	case 1:
		params := e.Parameters.(types.GXStructure)
		if len(params) < 2 {
			e.Error = enums.ErrorCodeReadWriteDenied
			return nil, nil
		}
		identifier := params[0].([]byte)
		size := params[1].(uint32)
		g.ImageFirstNotTransferredBlockNum = 0
		g.ImageTransferStatus = enums.ImageTransferStatusTransferInitiated
		g.setOrUpdateActivateInfo(identifier, size)
		g.ImageTransferredBlocksStatus = g.newBlockStatus(size)
		return nil, nil
	case 2:
		params := e.Parameters.(types.GXStructure)
		if len(params) < 1 {
			e.Error = enums.ErrorCodeReadWriteDenied
			return nil, nil
		}
		index := params[0].(uint32)
		status := []byte(g.ImageTransferredBlocksStatus)
		if int(index) < len(status) {
			status[index] = '1'
			g.ImageTransferredBlocksStatus = string(status)
		}
		g.ImageFirstNotTransferredBlockNum = index + 1
		return nil, nil
	case 3:
		return nil, nil
	case 4:
		return nil, nil
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
		return nil, nil
	}
}

// GetAttributeIndexToRead returns collection of attributes to read.
func (g *GXDLMSImageTransfer) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	if all || !g.IsRead(2) {
		attributes = append(attributes, 2)
	}
	attributes = append(attributes, 3)
	attributes = append(attributes, 4)
	if all || !g.IsRead(5) {
		attributes = append(attributes, 5)
	}
	attributes = append(attributes, 6)
	attributes = append(attributes, 7)
	return attributes
}

// GetNames returns names of attribute indexes.
func (g *GXDLMSImageTransfer) GetNames() []string {
	return []string{
		"Logical Name",
		"Image Block Size",
		"Image Transferred Blocks Status",
		"Image FirstNot Transferred Block Number",
		"Image Transfer Enabled",
		"Image Transfer Status",
		"Image Activate Info",
	}
}

// GetMethodNames returns names of method indexes.
func (g *GXDLMSImageTransfer) GetMethodNames() []string {
	return []string{"Image transfer initiate", "Image block transfer", "Image verify", "Image activate"}
}

// GetAttributeCount returns amount of attributes.
func (g *GXDLMSImageTransfer) GetAttributeCount() int {
	return 7
}

// GetMethodCount returns amount of methods.
func (g *GXDLMSImageTransfer) GetMethodCount() int {
	return 4
}

// GetDataType returns data type of selected attribute index.
func (g *GXDLMSImageTransfer) GetDataType(index int) (enums.DataType, error) {
	switch index {
	case 1:
		return enums.DataTypeOctetString, nil
	case 2:
		return enums.DataTypeUint32, nil
	case 3:
		return enums.DataTypeBitString, nil
	case 4:
		return enums.DataTypeUint32, nil
	case 5:
		return enums.DataTypeBoolean, nil
	case 6:
		return enums.DataTypeEnum, nil
	case 7:
		return enums.DataTypeArray, nil
	default:
		return 0, dlmserrors.ErrInvalidAttributeIndex
	}
}

// GetValue returns value of given attribute.
func (g *GXDLMSImageTransfer) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	switch e.Index {
	case 1:
		return helpers.LogicalNameToBytes(g.LogicalName())
	case 2:
		return g.ImageBlockSize, nil
	case 3:
		return g.ImageTransferredBlocksStatus, nil
	case 4:
		return g.ImageFirstNotTransferredBlockNum, nil
	case 5:
		return g.ImageTransferEnabled, nil
	case 6:
		return uint8(g.ImageTransferStatus), nil
	case 7:
		return g.getImageActivateInfo(settings)
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
		return nil, nil
	}
}

// SetValue sets value of given attribute.
func (g *GXDLMSImageTransfer) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	var err error
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
		g.ImageBlockSize = v
	case 3:
		bs := e.Value.(types.GXBitString)
		g.ImageTransferredBlocksStatus = bs.String()
	case 4:
		v, err := toUint32(e.Value)
		if err != nil {
			return err
		}
		g.ImageFirstNotTransferredBlockNum = v
	case 5:
		v, err := toBool(e.Value)
		if err != nil {
			return err
		}
		g.ImageTransferEnabled = v
	case 6:
		v, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.ImageTransferStatus = enums.ImageTransferStatus(v)
	case 7:
		g.ImageActivateInfo, err = parseImageActivateInfo(e.Value.(types.GXArray))
		if err != nil {
			return err
		}
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil
}

// Load loads object content from XML.
func (g *GXDLMSImageTransfer) Load(reader *GXXmlReader) error {
	var err error
	if g.ImageBlockSize, err = reader.ReadElementContentAsUInt32("ImageBlockSize", 0); err != nil {
		return err
	}
	if g.ImageTransferredBlocksStatus, err = reader.ReadElementContentAsString("ImageTransferredBlocksStatus", ""); err != nil {
		return err
	}
	if g.ImageFirstNotTransferredBlockNum, err = reader.ReadElementContentAsUInt32("ImageFirstNotTransferredBlockNumber", 0); err != nil {
		return err
	}
	if v, err := reader.ReadElementContentAsInt("ImageTransferEnabled", 0); err != nil {
		return err
	} else {
		g.ImageTransferEnabled = v != 0
	}
	if v, err := reader.ReadElementContentAsInt("ImageTransferStatus", 0); err != nil {
		return err
	} else {
		g.ImageTransferStatus = enums.ImageTransferStatus(v)
	}

	items := make([]GXDLMSImageActivateInfo, 0)
	ok, err := reader.IsStartElementNamed("ImageActivateInfo", true)
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
			item := GXDLMSImageActivateInfo{}
			if item.Size, err = reader.ReadElementContentAsUInt32("Size", 0); err != nil {
				return err
			}
			if v, err := reader.ReadElementContentAsString("Identification", ""); err != nil {
				return err
			} else {
				item.Identification = types.HexToBytes(v)
			}
			if v, err := reader.ReadElementContentAsString("Signature", ""); err != nil {
				return err
			} else {
				item.Signature = types.HexToBytes(v)
			}
			items = append(items, item)
		}
		if err := reader.ReadEndElement("ImageActivateInfo"); err != nil {
			return err
		}
	}
	g.ImageActivateInfo = items
	return nil
}

// Save saves object content to XML.
func (g *GXDLMSImageTransfer) Save(writer *GXXmlWriter) error {
	if err := writer.WriteElementString("ImageBlockSize", g.ImageBlockSize); err != nil {
		return err
	}
	if err := writer.WriteElementString("ImageTransferredBlocksStatus", g.ImageTransferredBlocksStatus); err != nil {
		return err
	}
	if err := writer.WriteElementString("ImageFirstNotTransferredBlockNumber", g.ImageFirstNotTransferredBlockNum); err != nil {
		return err
	}
	if err := writer.WriteElementString("ImageTransferEnabled", g.ImageTransferEnabled); err != nil {
		return err
	}
	if err := writer.WriteElementString("ImageTransferStatus", int(g.ImageTransferStatus)); err != nil {
		return err
	}

	writer.WriteStartElement("ImageActivateInfo")
	for _, it := range g.ImageActivateInfo {
		writer.WriteStartElement("Item")
		if err := writer.WriteElementString("Size", it.Size); err != nil {
			return err
		}
		if err := writer.WriteElementString("Identification", types.ToHex(it.Identification, false)); err != nil {
			return err
		}
		if err := writer.WriteElementString("Signature", types.ToHex(it.Signature, false)); err != nil {
			return err
		}
		writer.WriteEndElement()
	}
	writer.WriteEndElement()
	return nil
}

// PostLoad handles actions after Load.
func (g *GXDLMSImageTransfer) PostLoad(reader *GXXmlReader) error {
	return nil
}

// GetValues returns attributes as an array.
func (g *GXDLMSImageTransfer) GetValues() []any {
	return []any{
		g.LogicalName(),
		g.ImageBlockSize,
		g.ImageTransferredBlocksStatus,
		g.ImageFirstNotTransferredBlockNum,
		g.ImageTransferEnabled,
		g.ImageTransferStatus,
		g.ImageActivateInfo,
	}
}

// ImageTransferInitiate starts image transfer procedure.
func (g *GXDLMSImageTransfer) ImageTransferInitiate(client IGXDLMSClient, imageIdentifier []byte, imageSize uint32) ([][]byte, error) {
	if g.ImageBlockSize == 0 {
		return nil, errors.New("invalid image block size")
	}
	if s, ok := client.Settings().(*settings.GXDLMSSettings); ok {
		if g.ImageBlockSize > uint32(s.MaxPduSize()) {
			return nil, errors.New("image block size is bigger than max PDU size")
		}
	}
	params := types.GXStructure{imageIdentifier, imageSize}
	return client.Method(g, 1, params, enums.DataTypeStructure)
}

// ImageTransferInitiateByString starts image transfer using an ASCII identifier.
func (g *GXDLMSImageTransfer) ImageTransferInitiateByString(client IGXDLMSClient, imageIdentifier string, imageSize uint32) ([][]byte, error) {
	return g.ImageTransferInitiate(client, []byte(imageIdentifier), imageSize)
}

// GetImageBlocks splits image to transfer blocks using ImageBlockSize.
func (g *GXDLMSImageTransfer) GetImageBlocks(image []byte) ([][]byte, error) {
	if g.ImageBlockSize == 0 {
		return nil, errors.New("invalid image block size")
	}
	size := int(g.ImageBlockSize)
	cnt := len(image) / size
	if len(image)%size != 0 {
		cnt++
	}
	blocks := make([][]byte, 0, cnt)
	for pos := 0; pos < cnt; pos++ {
		start := pos * size
		end := start + size
		if end > len(image) {
			end = len(image)
		}
		part := make([]byte, end-start)
		copy(part, image[start:end])
		blocks = append(blocks, part)
	}
	return blocks, nil
}

// ImageBlockTransfer transfers image blocks from start index.
func (g *GXDLMSImageTransfer) ImageBlockTransfer(client IGXDLMSClient, image []byte) ([][]byte, int, error) {
	return g.ImageBlockTransferFromIndex(client, image, 0)
}

// ImageBlockTransferFromIndex transfers image blocks from a block index.
func (g *GXDLMSImageTransfer) ImageBlockTransferFromIndex(client IGXDLMSClient, image []byte, index int) ([][]byte, int, error) {
	blocks, err := g.GetImageBlocks(image)
	if err != nil {
		return nil, 0, err
	}
	imageBlockCount := len(blocks)
	if index < 0 || index >= imageBlockCount {
		return nil, imageBlockCount, errors.New("image start index is higher than image block count")
	}
	packets := make([][]byte, 0)
	for pos, block := range blocks {
		if pos < index {
			continue
		}
		params := types.GXStructure{uint32(pos), block}
		frames, err := client.Method(g, 2, params, enums.DataTypeStructure)
		if err != nil {
			return nil, imageBlockCount, err
		}
		packets = append(packets, frames...)
	}
	return packets, imageBlockCount, nil
}

// ImageBlockTransferByStatus transfers only missing blocks from a block-status bitstring.
func (g *GXDLMSImageTransfer) ImageBlockTransferByStatus(client IGXDLMSClient, image []byte, blocksStatus string) ([][]byte, int, error) {
	blocks, err := g.GetImageBlocks(image)
	if err != nil {
		return nil, 0, err
	}
	imageBlockCount := len(blocks)
	if len(blocksStatus) < imageBlockCount {
		return nil, imageBlockCount, errors.New("image block status is shorter than image block count")
	}
	packets := make([][]byte, 0)
	for pos, block := range blocks {
		if blocksStatus[pos] == '1' {
			continue
		}
		params := types.GXStructure{uint32(pos), block}
		frames, err := client.Method(g, 2, params, enums.DataTypeStructure)
		if err != nil {
			return nil, imageBlockCount, err
		}
		packets = append(packets, frames...)
	}
	return packets, imageBlockCount, nil
}

// ImageVerify verifies image in the meter.
func (g *GXDLMSImageTransfer) ImageVerify(client IGXDLMSClient) ([][]byte, error) {
	return client.Method(g, 3, int8(0), enums.DataTypeInt8)
}

// ImageActivate activates verified image in the meter.
func (g *GXDLMSImageTransfer) ImageActivate(client IGXDLMSClient) ([][]byte, error) {
	return client.Method(g, 4, int8(0), enums.DataTypeInt8)
}

func (g *GXDLMSImageTransfer) getImageActivateInfo(settings *settings.GXDLMSSettings) ([]byte, error) {
	data := types.NewGXByteBuffer()
	if err := data.SetUint8(uint8(enums.DataTypeArray)); err != nil {
		return nil, err
	}
	if g.ImageTransferStatus != enums.ImageTransferStatusVerificationSuccessful || g.ImageActivateInfo == nil {
		if err := types.SetObjectCount(0, data); err != nil {
			return nil, err
		}
		return data.Array(), nil
	}
	if err := types.SetObjectCount(len(g.ImageActivateInfo), data); err != nil {
		return nil, err
	}
	for _, it := range g.ImageActivateInfo {
		if err := data.SetUint8(uint8(enums.DataTypeStructure)); err != nil {
			return nil, err
		}
		if err := data.SetUint8(3); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, data, enums.DataTypeUint32, it.Size); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, data, enums.DataTypeOctetString, it.Identification); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, data, enums.DataTypeOctetString, it.Signature); err != nil {
			return nil, err
		}
	}
	return data.Array(), nil
}

func (g *GXDLMSImageTransfer) setOrUpdateActivateInfo(identifier []byte, size uint32) {
	idx := -1
	for pos, it := range g.ImageActivateInfo {
		if bytes.Equal(it.Identification, identifier) {
			idx = pos
			break
		}
	}
	if idx == -1 {
		g.ImageActivateInfo = append(g.ImageActivateInfo, GXDLMSImageActivateInfo{Size: size, Identification: identifier})
		return
	}
	g.ImageActivateInfo[idx].Size = size
	g.ImageActivateInfo[idx].Identification = identifier
}

func (g *GXDLMSImageTransfer) newBlockStatus(imageSize uint32) string {
	if g.ImageBlockSize == 0 {
		return ""
	}
	cnt := int((imageSize + g.ImageBlockSize - 1) / g.ImageBlockSize)
	if cnt < 0 {
		cnt = 0
	}
	return strings.Repeat("0", cnt)
}

func parseImageActivateInfo(value types.GXArray) ([]GXDLMSImageActivateInfo, error) {
	list := make([]GXDLMSImageActivateInfo, 0)
	for _, tmp := range value {
		it := tmp.(types.GXStructure)
		if len(it) < 3 {
			return nil, gxcommon.ErrInvalidArgument
		}
		size := it[0].(uint32)
		id, _ := it[1].([]byte)
		sig, _ := it[2].([]byte)
		list = append(list, GXDLMSImageActivateInfo{
			Size:           size,
			Identification: id,
			Signature:      sig,
		})
	}
	return list, nil
}

// NewGXDLMSImageTransfer creates a new Image transfer object instance.
func NewGXDLMSImageTransfer(ln string, sn int16) (*GXDLMSImageTransfer, error) {
	if err := ValidateLogicalName(ln); err != nil {
		return nil, err
	}
	return &GXDLMSImageTransfer{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeImageTransfer,
			logicalName: ln,
			ShortName:   sn,
		},
		ImageBlockSize:                   200,
		ImageFirstNotTransferredBlockNum: 0,
		ImageTransferEnabled:             true,
		ImageTransferStatus:              enums.ImageTransferStatusNotInitiated,
	}, nil
}
