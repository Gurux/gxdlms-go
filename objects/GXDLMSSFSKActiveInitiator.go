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

// GXDLMSSFSKActiveInitiator contains active initiator information.
type GXDLMSSFSKActiveInitiator struct {
	GXDLMSObject
	SystemTitle  []byte
	MacAddress   uint16
	LSapSelector uint8
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSSFSKActiveInitiator) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

func (g *GXDLMSSFSKActiveInitiator) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, nil
}

func (g *GXDLMSSFSKActiveInitiator) GetAttributeIndexToRead(all bool) []int {
	var a []int
	if all || g.LogicalName() == "" {
		a = append(a, 1)
	}
	// Active initiator value is dynamic and should always be read.
	a = append(a, 2)
	return a
}

func (g *GXDLMSSFSKActiveInitiator) GetNames() []string {
	return []string{"Logical Name", "Active Initiator"}
}

func (g *GXDLMSSFSKActiveInitiator) GetMethodNames() []string {
	return []string{"Reset NEW not synchronized"}
}

func (g *GXDLMSSFSKActiveInitiator) GetAttributeCount() int { return 2 }

func (g *GXDLMSSFSKActiveInitiator) GetMethodCount() int { return 1 }

func (g *GXDLMSSFSKActiveInitiator) GetDataType(index int) (enums.DataType, error) {
	if index == 1 {
		return enums.DataTypeOctetString, nil
	}
	if index == 2 {
		return enums.DataTypeStructure, nil
	}
	return 0, dlmserrors.ErrInvalidAttributeIndex
}

func (g *GXDLMSSFSKActiveInitiator) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
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
		if err := internal.SetData(settings, bb, enums.DataTypeOctetString, g.SystemTitle); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeUint16, g.MacAddress); err != nil {
			return nil, err
		}
		if err := internal.SetData(settings, bb, enums.DataTypeUint8, g.LSapSelector); err != nil {
			return nil, err
		}
		return bb.Array(), nil
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
		return nil, nil
	}
}

func (g *GXDLMSSFSKActiveInitiator) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
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
			g.SystemTitle = nil
			g.MacAddress = 0
			g.LSapSelector = 0
			return nil
		}
		tmp, ok := e.Value.(types.GXStructure)
		if !ok || len(tmp) < 3 {
			return fmt.Errorf("invalid active initiator value: %T", e.Value)
		}
		switch v := tmp[0].(type) {
		case nil:
			g.SystemTitle = nil
		case []byte:
			g.SystemTitle = v
		default:
			return fmt.Errorf("invalid system title type: %T", tmp[0])
		}
		mac, err := toUint32(tmp[1])
		if err != nil {
			return err
		}
		g.MacAddress = uint16(mac)
		lsap, err := toUint8(tmp[2])
		if err != nil {
			return err
		}
		g.LSapSelector = lsap
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil
}

func (g *GXDLMSSFSKActiveInitiator) Load(reader *GXXmlReader) error {
	v, err := reader.ReadElementContentAsString("SystemTitle", "")
	if err != nil {
		return err
	}
	g.SystemTitle = types.HexToBytes(v)
	g.MacAddress, err = reader.ReadElementContentAsUInt16("MacAddress", 0)
	if err != nil {
		return err
	}
	g.LSapSelector, err = reader.ReadElementContentAsUInt8("LSapSelector", 0)
	return err
}

func (g *GXDLMSSFSKActiveInitiator) Save(writer *GXXmlWriter) error {
	if err := writer.WriteElementString("SystemTitle", types.ToHex(g.SystemTitle, false)); err != nil {
		return err
	}
	if err := writer.WriteElementString("MacAddress", g.MacAddress); err != nil {
		return err
	}
	return writer.WriteElementString("LSapSelector", g.LSapSelector)
}

func (g *GXDLMSSFSKActiveInitiator) PostLoad(reader *GXXmlReader) error { return nil }

func (g *GXDLMSSFSKActiveInitiator) GetValues() []any {
	return []any{g.LogicalName(), []any{g.SystemTitle, g.MacAddress, g.LSapSelector}}
}

// NewGXDLMSSFSKActiveInitiator creates a new S-FSK active initiator object instance.
//
// The function validates `ln` before creating the object.
// `ln` is the Logical Name and `sn` is the Short Name of the object.
func NewGXDLMSSFSKActiveInitiator(ln string, sn int16) (*GXDLMSSFSKActiveInitiator, error) {
	if err := ValidateLogicalName(ln); err != nil {
		return nil, err
	}
	return &GXDLMSSFSKActiveInitiator{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeSFSKActiveInitiator,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
