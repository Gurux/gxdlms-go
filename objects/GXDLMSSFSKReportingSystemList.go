package objects

import (
	"github.com/Gurux/gxdlms-go/dlmserrors"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// GXDLMSSFSKReportingSystemList contains registered reporting systems.
type GXDLMSSFSKReportingSystemList struct {
	GXDLMSObject
	ReportingSystemList [][]byte
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSSFSKReportingSystemList) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

func (g *GXDLMSSFSKReportingSystemList) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, nil
}

func (g *GXDLMSSFSKReportingSystemList) GetAttributeIndexToRead(all bool) []int {
	var a []int
	if all || g.LogicalName() == "" {
		a = append(a, 1)
	}
	if all || g.CanRead(2) {
		a = append(a, 2)
	}
	return a
}

func (g *GXDLMSSFSKReportingSystemList) GetNames() []string {
	return []string{"Logical Name", "ReportingSystemList"}
}

func (g *GXDLMSSFSKReportingSystemList) GetMethodNames() []string { return []string{} }

func (g *GXDLMSSFSKReportingSystemList) GetAttributeCount() int { return 2 }

func (g *GXDLMSSFSKReportingSystemList) GetMethodCount() int { return 0 }

func (g *GXDLMSSFSKReportingSystemList) GetDataType(index int) (enums.DataType, error) {
	if index == 1 {
		return enums.DataTypeOctetString, nil
	}
	if index == 2 {
		return enums.DataTypeArray, nil
	}
	return 0, dlmserrors.ErrInvalidAttributeIndex
}

func (g *GXDLMSSFSKReportingSystemList) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	switch e.Index {
	case 1:
		return helpers.LogicalNameToBytes(g.LogicalName())
	case 2:
		data := types.NewGXByteBuffer()
		_ = data.SetUint8(uint8(enums.DataTypeArray))
		_ = types.SetObjectCount(len(g.ReportingSystemList), data)
		for _, it := range g.ReportingSystemList {
			if err := internal.SetData(settings, data, enums.DataTypeOctetString, it); err != nil {
				return nil, err
			}
		}
		return data.Array(), nil
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
		return nil, nil
	}
}

func (g *GXDLMSSFSKReportingSystemList) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	switch e.Index {
	case 1:
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return err
		}
		return g.SetLogicalName(ln)
	case 2:
		g.ReportingSystemList = g.ReportingSystemList[:0]
		for _, tmp := range e.Value.(types.GXArray) {
			if b, ok := tmp.([]byte); ok {
				g.ReportingSystemList = append(g.ReportingSystemList, b)
			}
		}
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil
}

func (g *GXDLMSSFSKReportingSystemList) Load(reader *GXXmlReader) error {
	g.ReportingSystemList = g.ReportingSystemList[:0]
	ok, err := reader.IsStartElementNamed("ReportingSystems", true)
	if err != nil {
		return err
	}
	if ok {
		for {
			ok, err = reader.IsStartElementNamed("Item", false)
			if err != nil {
				return err
			}
			if !ok {
				break
			}
			v, err := reader.ReadElementContentAsString("Item", "")
			if err != nil {
				return err
			}
			g.ReportingSystemList = append(g.ReportingSystemList, types.HexToBytes(v))
		}
		_ = reader.ReadEndElement("ReportingSystems")
	}
	return nil
}

func (g *GXDLMSSFSKReportingSystemList) Save(writer *GXXmlWriter) error {
	writer.WriteStartElement("ReportingSystems")
	for _, it := range g.ReportingSystemList {
		if err := writer.WriteElementString("Item", types.ToHex(it, false)); err != nil {
			return err
		}
	}
	writer.WriteEndElement()
	return nil
}

func (g *GXDLMSSFSKReportingSystemList) PostLoad(reader *GXXmlReader) error { return nil }

func (g *GXDLMSSFSKReportingSystemList) GetValues() []any {
	return []any{g.LogicalName(), g.ReportingSystemList}
}

// NewGXDLMSSFSKReportingSystemList creates a new S-FSK reporting system list object instance.
//
// The function validates `ln` before creating the object.
// `ln` is the Logical Name and `sn` is the Short Name of the object.
func NewGXDLMSSFSKReportingSystemList(ln string, sn int16) (*GXDLMSSFSKReportingSystemList, error) {
	if err := ValidateLogicalName(ln); err != nil {
		return nil, err
	}
	return &GXDLMSSFSKReportingSystemList{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeSFSKReportingSystemList,
			logicalName: ln,
			ShortName:   sn},
		ReportingSystemList: make([][]byte, 0)}, nil
}
