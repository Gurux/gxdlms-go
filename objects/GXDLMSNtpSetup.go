package objects

import (
	"reflect"
	"sort"

	"github.com/Gurux/gxcommon-go"
	"github.com/Gurux/gxdlms-go/dlmserrors"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// GXDLMSNtpSetup contains NTP setup parameters used for time synchronization.
type GXDLMSNtpSetup struct {
	GXDLMSObject
	Activated      bool
	ServerAddress  string
	Port           uint16
	Authentication enums.NtpAuthenticationMethod
	Keys           map[uint32][]byte
	ClientKey      []byte
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSNtpSetup) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

func (g *GXDLMSNtpSetup) GetUIDataType(index int) enums.DataType {
	if index == 3 {
		return enums.DataTypeString
	}
	return g.Base().GetUIDataType(index)
}

func (g *GXDLMSNtpSetup) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	switch e.Index {
	case 1:
		// Server synchronizes time.
	case 2:
		tmp, ok := e.Parameters.(types.GXStructure)
		if !ok || len(tmp) < 2 {
			e.Error = enums.ErrorCodeInconsistentClass
			return nil, nil
		}
		id, err := toUint32(tmp[0])
		if err != nil {
			e.Error = enums.ErrorCodeInconsistentClass
			return nil, nil
		}
		key, ok := tmp[1].([]byte)
		if !ok {
			e.Error = enums.ErrorCodeInconsistentClass
			return nil, nil
		}
		g.Keys[id] = key
	case 3:
		id, err := toUint32(e.Parameters)
		if err != nil {
			e.Error = enums.ErrorCodeInconsistentClass
			return nil, nil
		}
		delete(g.Keys, id)
	default:
		e.Error = enums.ErrorCodeInconsistentClass
	}
	return nil, nil
}

func (g *GXDLMSNtpSetup) GetAttributeIndexToRead(all bool) []int {
	var a []int
	if all || g.LogicalName() == "" {
		a = append(a, 1)
	}
	for i := 2; i <= 7; i++ {
		if all || g.CanRead(i) {
			a = append(a, i)
		}
	}
	return a
}

func (g *GXDLMSNtpSetup) GetNames() []string {
	return []string{"Logical Name", "Activated", "ServerAddress", "Port", "Authentication", "Keys", "ClientKey"}
}

func (g *GXDLMSNtpSetup) GetMethodNames() []string {
	return []string{"Synchronize", "Add authentication key", "Delete authentication key"}
}

func (g *GXDLMSNtpSetup) GetAttributeCount() int { return 7 }

func (g *GXDLMSNtpSetup) GetMethodCount() int { return 3 }

func (g *GXDLMSNtpSetup) GetDataType(index int) (enums.DataType, error) {
	switch index {
	case 1:
		return enums.DataTypeOctetString, nil
	case 2:
		return enums.DataTypeBoolean, nil
	case 3:
		return enums.DataTypeOctetString, nil
	case 4:
		return enums.DataTypeUint16, nil
	case 5:
		return enums.DataTypeEnum, nil
	case 6:
		return enums.DataTypeArray, nil
	case 7:
		return enums.DataTypeOctetString, nil
	default:
		return enums.DataTypeNone, dlmserrors.ErrInvalidAttributeIndex
	}
}

func (g *GXDLMSNtpSetup) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	switch e.Index {
	case 1:
		return helpers.LogicalNameToBytes(g.LogicalName())
	case 2:
		return g.Activated, nil
	case 3:
		return g.ServerAddress, nil
	case 4:
		return g.Port, nil
	case 5:
		return uint8(g.Authentication), nil
	case 6:
		bb := types.NewGXByteBuffer()
		if err := bb.SetUint8(uint8(enums.DataTypeArray)); err != nil {
			return nil, err
		}
		keys := make([]uint32, 0, len(g.Keys))
		for k := range g.Keys {
			keys = append(keys, k)
		}
		sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
		types.SetObjectCount(len(keys), bb)
		for _, id := range keys {
			key := g.Keys[id]
			if err := bb.SetUint8(uint8(enums.DataTypeStructure)); err != nil {
				return nil, err
			}
			if err := bb.SetUint8(2); err != nil {
				return nil, err
			}
			if err := bb.SetUint8(uint8(enums.DataTypeUint32)); err != nil {
				return nil, err
			}
			if err := bb.SetUint32(id); err != nil {
				return nil, err
			}
			if err := bb.SetUint8(uint8(enums.DataTypeOctetString)); err != nil {
				return nil, err
			}
			types.SetObjectCount(len(key), bb)
			if err := bb.Set(key); err != nil {
				return nil, err
			}
		}
		return bb.Array(), nil
	case 7:
		return g.ClientKey, nil
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
		return nil, nil
	}
}

func (g *GXDLMSNtpSetup) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
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
		g.Activated, err = toBool(e.Value)
	case 3:
		switch v := e.Value.(type) {
		case nil:
			g.ServerAddress = ""
		case []byte:
			g.ServerAddress = string(v)
		case string:
			g.ServerAddress = v
		default:
			g.ServerAddress = ""
		}
	case 4:
		g.Port, err = toUint16(e.Value)
	case 5:
		g.Authentication = enums.NtpAuthenticationMethod(e.Value.(types.GXEnum).Value)
	case 6:
		g.Keys = map[uint32][]byte{}
		if e.Value != nil {
			tmp, ok := e.Value.(types.GXArray)
			if !ok {
				return gxcommon.ErrInvalidArgument
			}
			for _, row := range tmp {
				s, ok := row.(types.GXStructure)
				if !ok || len(s) < 2 {
					continue
				}
				id, err := toUint32(s[0])
				if err != nil {
					return err
				}
				if key, ok := s[1].([]byte); ok {
					g.Keys[id] = key
				}
			}
		}
	case 7:
		if e.Value == nil {
			g.ClientKey = nil
		} else if v, ok := e.Value.([]byte); ok {
			g.ClientKey = v
		}
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return err
}

func (g *GXDLMSNtpSetup) Load(reader *GXXmlReader) error {
	var err error
	activated, err := reader.ReadElementContentAsInt("Activated", 1)
	if err != nil {
		return err
	}
	g.Activated = activated != 0
	g.ServerAddress, err = reader.ReadElementContentAsString("ServerAddress", "")
	if err != nil {
		return err
	}
	g.Port, err = reader.ReadElementContentAsUInt16("Port", 0)
	if err != nil {
		return err
	}
	auth, err := reader.ReadElementContentAsInt("Authentication", 0)
	if err != nil {
		return err
	}
	g.Authentication = enums.NtpAuthenticationMethod(auth)
	g.Keys = map[uint32][]byte{}
	ok, err := reader.IsStartElementNamed("Keys", true)
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
			id, err := reader.ReadElementContentAsUInt32("ID", 0)
			if err != nil {
				return err
			}
			keyHex, err := reader.ReadElementContentAsString("Key", "")
			if err != nil {
				return err
			}
			g.Keys[id] = types.HexToBytes(keyHex)
		}
		if err = reader.ReadEndElement("Keys"); err != nil {
			return err
		}
	}
	clientKeyHex, err := reader.ReadElementContentAsString("ClientKey", "")
	if err != nil {
		return err
	}
	g.ClientKey = types.HexToBytes(clientKeyHex)
	return nil
}

func (g *GXDLMSNtpSetup) Save(writer *GXXmlWriter) error {
	if err := writer.WriteElementStringBool("Activated", g.Activated); err != nil {
		return err
	}
	if err := writer.WriteElementString("ServerAddress", g.ServerAddress); err != nil {
		return err
	}
	if err := writer.WriteElementString("Port", g.Port); err != nil {
		return err
	}
	if err := writer.WriteElementString("Authentication", int(g.Authentication)); err != nil {
		return err
	}
	if err := writer.WriteStartElement("Keys"); err != nil {
		return err
	}
	keys := make([]uint32, 0, len(g.Keys))
	for k := range g.Keys {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	for _, id := range keys {
		if err := writer.WriteStartElement("Item"); err != nil {
			return err
		}
		if err := writer.WriteElementString("ID", id); err != nil {
			return err
		}
		if err := writer.WriteElementString("Key", types.ToHex(g.Keys[id], false)); err != nil {
			return err
		}
		if err := writer.WriteEndElement(); err != nil {
			return err
		}
	}
	if err := writer.WriteEndElement(); err != nil {
		return err
	}
	return writer.WriteElementString("ClientKey", types.ToHex(g.ClientKey, false))
}

func (g *GXDLMSNtpSetup) PostLoad(reader *GXXmlReader) error { return nil }

func (g *GXDLMSNtpSetup) GetValues() []any {
	return []any{
		g.LogicalName(),
		g.Activated,
		g.ServerAddress,
		g.Port,
		g.Authentication,
		g.Keys,
		g.ClientKey,
	}
}

// Synchronize synchronizes time of the server with NTP server.
func (g *GXDLMSNtpSetup) Synchronize(client IGXDLMSClient) ([][]byte, error) {
	return client.Method(g, 1, int8(0), enums.DataTypeInt8)
}

// AddAuthenticationKey adds a symmetric authentication key.
func (g *GXDLMSNtpSetup) AddAuthenticationKey(client IGXDLMSClient, id uint32, key []byte) ([][]byte, error) {
	bb := types.NewGXByteBuffer()
	if err := bb.SetUint8(uint8(enums.DataTypeStructure)); err != nil {
		return nil, err
	}
	if err := bb.SetUint8(2); err != nil {
		return nil, err
	}
	if err := bb.SetUint8(uint8(enums.DataTypeUint32)); err != nil {
		return nil, err
	}
	if err := bb.SetUint32(id); err != nil {
		return nil, err
	}
	if err := bb.SetUint8(uint8(enums.DataTypeOctetString)); err != nil {
		return nil, err
	}
	types.SetObjectCount(len(key), bb)
	if err := bb.Set(key); err != nil {
		return nil, err
	}
	return client.Method(g, 2, bb.Array(), enums.DataTypeStructure)
}

// DeleteAuthenticationKey removes a symmetric authentication key.
func (g *GXDLMSNtpSetup) DeleteAuthenticationKey(client IGXDLMSClient, id uint32) ([][]byte, error) {
	return client.Method(g, 3, id, enums.DataTypeUint32)
}

func (g *GXDLMSNtpSetup) GetDataTypeByValue(value any) (enums.DataType, error) {
	if value == nil {
		return enums.DataTypeNone, nil
	}
	return internal.GetDLMSDataType(reflect.TypeOf(value))
}

// NewGXDLMSNtpSetup creates a new NTP setup object instance.
func NewGXDLMSNtpSetup(ln string, sn int16) (*GXDLMSNtpSetup, error) {
	if err := ValidateLogicalName(ln); err != nil {
		return nil, err
	}
	return &GXDLMSNtpSetup{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeNtpSetup,
			logicalName: ln,
			ShortName:   sn,
		},
		Port: 123,
		Keys: map[uint32][]byte{},
	}, nil
}
