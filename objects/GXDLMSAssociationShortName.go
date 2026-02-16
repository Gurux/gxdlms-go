package objects

//
// --------------------------------------------------------------------------
//  Gurux Ltd
//
//
//
// Filename:        $HeadURL$
//
// Version:         $Revision$,
//                  $Date$
//                  $Author$
//
// Copyright (c) Gurux Ltd
//
//---------------------------------------------------------------------------
//
//  DESCRIPTION
//
// This file is a part of Gurux Device Framework.
//
// Gurux Device Framework is Open Source software; you can redistribute it
// and/or modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation; version 2 of the License.
// Gurux Device Framework is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
// See the GNU General Public License for more details.
//
// More information of Gurux products: https://www.gurux.org
//
// This code is licensed under the GNU General Public License v2.
// Full text may be retrieved at http://www.gnu.org/licenses/gpl-2.0.txt
//---------------------------------------------------------------------------

import (
	"bytes"
	"errors"

	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// Online help:
// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSAssociationShortName
type GXDLMSAssociationShortName struct {
	GXDLMSObject
	// Secret used in Authentication
	Secret []byte

	// List of available objects in short name referencing.
	ObjectList GXDLMSObjectCollection

	// Security setup reference.
	SecuritySetupReference string
}

// base returns the base GXDLMSObject of the object.
func (g *GXDLMSAssociationShortName) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

// GetAttributeIndexToRead returns the collection of attributes to read.
// If attribute is static and already read or device is returned HW error it is not returned.
//
// Parameters:
//
//	all: All items are returned even if they are read already.
//
// Returns:
//
//	Collection of attributes to read.
func (g *GXDLMSAssociationShortName) GetAttributeIndexToRead(all bool) []int {
	attributes := []int{}
	// LN is static and read only once.
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	// ObjectList is static and read only once.
	if all || !g.IsRead(2) {
		attributes = append(attributes, 2)
	}
	if g.Version > 1 {
		// AccessRightsList is static and read only once.
		if all || !g.IsRead(3) {
			attributes = append(attributes, 3)
		}
		// SecuritySetupReference is static and read only once.
		if all || !g.IsRead(4) {
			attributes = append(attributes, 4)
		}
		if g.Version > 2 {
		}
	}
	return attributes
}

// GetNames returns the names of attribute indexes.
func (g *GXDLMSAssociationShortName) GetNames() []string {
	if g.Version < 2 {
		return []string{"Logical Name", "Object List"}
	}
	return []string{"Logical Name", "Object List", "Access Rights List", "Security Setup Reference"}
}

// GetMethodNames returns the names of method indexes.
func (g *GXDLMSAssociationShortName) GetMethodNames() []string {
	return []string{"Getlist by classid", "Getobj by logicalname", "Read by logicalname", "Get attributes&services", "Change LLS secret", "Change HLS secret", "Get HLS challenge", "Reply to HLS challenge", "Add user", "Remove user"}
}

// GetAttributeCount returns the amount of attributes.
//
// Returns:
//
//	Count of attributes.
func (g *GXDLMSAssociationShortName) GetAttributeCount() int {
	if g.Version < 2 {
		return 2
	}
	return 4
}

// GetMethodCount returns the amount of methods.
func (g *GXDLMSAssociationShortName) GetMethodCount() int {
	return 8
}

func (g *GXDLMSAssociationShortName) getAccessRights(settings *settings.GXDLMSSettings,
	item IGXDLMSBase,
	server internal.IGXDLMSServer,
	data *types.GXByteBuffer) error {
	err := data.SetUint8(uint8(enums.DataTypeStructure))
	if err != nil {
		return err
	}
	err = data.SetUint8(uint8(3))
	if err != nil {
		return err
	}
	err = internal.SetData(settings, data, enums.DataTypeInt16, item.Base().ShortName)
	if err != nil {
		return err
	}
	cnt := item.GetAttributeCount()
	err = data.SetUint8(uint8(enums.DataTypeArray))
	if err != nil {
		return err
	}
	err = data.SetUint8(uint8(cnt))
	if err != nil {
		return err
	}
	var e *internal.ValueEventArgs
	if server != nil {
		e = internal.NewValueEventArgs2(server, item, 0)
	} else {
		e = internal.NewValueEventArgs(settings, item, 0)
	}
	for pos := 0; pos != cnt; pos++ {
		e.Index = uint8(pos + 1)
		var m int
		if server != nil {
			m = server.NotifyGetAttributeAccess(e)
		} else {
			m = int(enums.AccessModeReadWrite)
		}
		err = data.SetUint8(uint8(enums.DataTypeStructure))
		if err != nil {
			return err
		}
		err = data.SetUint8(uint8(2))
		if err != nil {
			return err
		}
		err = internal.SetData(settings, data, enums.DataTypeInt8, e.Index)
		if err != nil {
			return err
		}
		err = internal.SetData(settings, data, enums.DataTypeEnum, m)
		if err != nil {
			return err
		}
	}
	cnt = item.GetMethodCount()
	err = data.SetUint8(uint8(enums.DataTypeArray))
	if err != nil {
		return err
	}
	err = data.SetUint8(uint8(cnt))
	if err != nil {
		return err
	}
	for pos := 0; pos != cnt; pos++ {
		e.Index = uint8(pos + 1)
		var m int
		if server != nil {
			m = server.NotifyGetMethodAccess(e)
		} else {
			m = int(enums.MethodAccessModeAccess)
		}
		err = data.SetUint8(uint8(enums.DataTypeStructure))
		if err != nil {
			return err
		}
		err = data.SetUint8(uint8(2))
		if err != nil {
			return err
		}
		err = internal.SetData(settings, data, enums.DataTypeInt8, e.Index)
		if err != nil {
			return err
		}
		err = internal.SetData(settings, data, enums.DataTypeEnum, m)
		if err != nil {
			return err
		}
	}
	return err
}

// getObjects returns the Association View.
func (g *GXDLMSAssociationShortName) getObjects(settings *settings.GXDLMSSettings,
	e *internal.ValueEventArgs) (*types.GXByteBuffer, error) {
	var err error
	cnt := len(g.ObjectList)
	data := types.NewGXByteBuffer()
	// Add count only for first time.
	if settings.Index == 0 {
		settings.Count = uint32(cnt)
		err = data.SetUint8(uint8(enums.DataTypeArray))
		if err != nil {
			return nil, err
		}
		err := types.SetObjectCount(cnt, data)
		if err != nil {
			return nil, err
		}
	}
	pos := uint32(0)
	for _, it := range g.ObjectList {
		pos++
		if !(pos <= settings.Index) {
			err = data.SetUint8(uint8(enums.DataTypeStructure))
			if err != nil {
				return nil, err
			}
			err = data.SetUint8(uint8(4))
			if err != nil {
				return nil, err
			}
			err = internal.SetData(settings, data, enums.DataTypeInt16, it.Base().ShortName)
			if err != nil {
				return nil, err
			}
			err = internal.SetData(settings, data, enums.DataTypeUint16, it.Base().ObjectType)
			if err != nil {
				return nil, err
			}
			err = internal.SetData(settings, data, enums.DataTypeUint8, it.Base().Version)
			if err != nil {
				return nil, err
			}
			ln, err := helpers.LogicalNameToBytes(it.Base().LogicalName())
			if err != nil {
				return nil, err
			}
			err = internal.SetData(settings, data, enums.DataTypeOctetString,
				ln)
			if err != nil {
				return nil, err
			}
			settings.Index++
			if settings.IsServer() {
				// If PDU is full.
				if !e.SkipMaxPduSize && uint16(data.Size()) >= settings.MaxPduSize() {
					break
				}
			}
		}
	}
	return data, nil
}

// GetValue returns the value of given attribute.
// When raw parameter us not used example register multiplies value by scalar.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Get parameters.
//
// Returns:
//
//	Value of the attribute index.
func (g *GXDLMSAssociationShortName) GetValue(settings *settings.GXDLMSSettings,
	e *internal.ValueEventArgs) (any, error) {
	var err error
	switch e.Index {
	case 1:
		if e.Index == 1 {
			v, err := helpers.LogicalNameToBytes(g.LogicalName())
			if err != nil {
				e.Error = enums.ErrorCodeReadWriteDenied
			}
			return v, err
		}
	case 2:
		ret, err := g.getObjects(settings, e)
		if err != nil {
			return nil, err
		}
		return ret.Array(), nil
	case 3:
		lnExists := g.ObjectList.FindBySN(uint16(g.ShortName)) != nil
		// Add count
		cnt := len(g.ObjectList)
		if !lnExists {
			cnt++
		}
		data := types.NewGXByteBuffer()
		err = data.SetUint8(uint8(enums.DataTypeArray))
		if err != nil {
			return nil, err
		}
		err = types.SetObjectCount(cnt, data)
		if err != nil {
			return nil, err
		}
		for _, it := range g.ObjectList {
			err = g.getAccessRights(settings, it, e.Server, data)
			if err != nil {
				return nil, err
			}
		}
		if !lnExists {
			err = g.getAccessRights(settings, g, e.Server, data)
			if err != nil {
				return nil, err
			}
		}
		return data.Array(), nil
	case 4:
		ln, err := helpers.LogicalNameToBytes(g.SecuritySetupReference)
		if err != nil {
			return nil, err
		}
		return ln, nil
	}
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, nil
}

func (g *GXDLMSAssociationShortName) updateAccessRights(buff []any) error {
	for _, it := range buff {
		access := it.([]any)
		sn := access[0].(uint16)
		obj := g.ObjectList.FindBySN(sn)
		if obj != nil {
			for _, it2 := range access[1].([]any) {
				attributeAccess := it2.([]any)
				id := attributeAccess[0].(int)
				mode := attributeAccess[1].(uint8)
				obj.Base().SetAccess(id, enums.AccessMode(mode))
			}
			for _, it2 := range access[2].([]any) {
				methodAccess := it2.([]any)
				id := methodAccess[0].(int)
				mode := methodAccess[1].(uint8)
				obj.Base().SetMethodAccess(id, enums.MethodAccessMode(mode))
			}
		}
	}
	return nil
}

// SetValue returns the set value of given attribute.
// When raw parameter us not used example register multiplies value by scalar.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Set parameters.
func (g *GXDLMSAssociationShortName) SetValue(settings *settings.GXDLMSSettings,
	e *internal.ValueEventArgs) error {
	var err error
	switch e.Index {
	case 1:
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return g.SetLogicalName(ln)
	case 2:
		g.ObjectList.Clear()
		if e.Value != nil {
			for _, it := range e.Value.([]any) {
				item := it.([]any)
				sn := item[0].(int16)
				type_ := (item[1]).(enums.ObjectType)
				version := item[2].(uint8)
				ln, err := helpers.ToLogicalName(item[3].([]byte))
				if err != nil {
					e.Error = enums.ErrorCodeReadWriteDenied
					return err
				}
				var obj IGXDLMSBase
				if settings.Objects != nil {
					obj = settings.Objects.(*GXDLMSObjectCollection).FindBySN(uint16(sn))
				}
				if obj == nil {
					obj = CreateObject(type_)
					if obj != nil {
						obj.Base().SetLogicalName(ln)
						obj.Base().ShortName = sn
						obj.Base().Version = version
					}
				}
				// Unknown objects are not shown.
				if obj != nil {
					g.ObjectList.Add(obj)
				}
			}
		}
	case 3:
		if e.Value == nil {
			for _, it := range g.ObjectList {
				for pos := 1; pos != it.GetAttributeCount(); pos++ {
					it.Base().SetAccess(pos, enums.AccessModeNoAccess)
				}
			}
		} else {
			err = g.updateAccessRights((e.Value).([]any))
		}
	case 4:
		g.SecuritySetupReference, err = helpers.ToLogicalName(e.Value)
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return err
}

// Load returns the load object content from XML.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSAssociationShortName) Load(reader *GXXmlReader) error {
	var err error
	str, err := reader.ReadElementContentAsString("Secret", "")
	if err != nil {
		return err
	}
	if str == "" {
		g.Secret = nil
	} else {
		g.Secret = types.HexToBytes(str)
	}
	g.SecuritySetupReference, err = reader.ReadElementContentAsString("SecuritySetupReference", "")
	if err != nil {
		return err
	}
	return err
}

// Save returns the save object content to XML.
//
// Parameters:
//
//	writer: XML writer.
func (g *GXDLMSAssociationShortName) Save(writer *GXXmlWriter) error {
	var err error
	err = writer.WriteElementString("Secret", types.ToHex(g.Secret, false))
	if err != nil {
		return err
	}
	err = writer.WriteElementString("SecuritySetupReference", g.SecuritySetupReference)
	if err != nil {
		return err
	}
	return err
}
func (g *GXDLMSAssociationShortName) Invoke(conf *settings.GXDLMSSettings,
	e *internal.ValueEventArgs) ([]byte, error) {
	// Check reply_to_HLS_authentication
	var err error
	if e.Index == 8 {
		var ic uint32
		var secret []byte
		switch conf.Authentication {
		case enums.AuthenticationHighGMAC:
			secret = conf.SourceSystemTitle()
			bb := types.NewGXByteBufferWithData(e.Parameters.([]byte))
			_, err := bb.Uint8()
			if err != nil {
				return nil, err
			}
			ic, err = bb.Uint32()
			if err != nil {
				return nil, err
			}
		case enums.AuthenticationHighSHA256:
			tmp := types.GXByteBuffer{}
			err = tmp.Set(g.Secret)
			if err != nil {
				return nil, err
			}
			err = tmp.Set(conf.SourceSystemTitle())
			if err != nil {
				return nil, err
			}
			err = tmp.Set(conf.Cipher.SystemTitle())
			if err != nil {
				return nil, err
			}
			err = tmp.Set(conf.StoCChallenge())
			if err != nil {
				return nil, err
			}
			err = tmp.Set(conf.CtoSChallenge())
			if err != nil {
				return nil, err
			}
			secret = tmp.Array()
		default:
			secret = g.Secret
		}
		serverChallenge, err := settings.Secure(conf, conf.Cipher, ic, conf.StoCChallenge(), secret)
		if err != nil {
			return nil, err
		}
		clientChallenge := e.Parameters.([]byte)
		if !bytes.Equal(serverChallenge, clientChallenge) {
			if conf.Authentication == enums.AuthenticationHighGMAC {
				secret = conf.Cipher.SystemTitle()
				ic = conf.Cipher.InvocationCounter()
				conf.Cipher.SetInvocationCounter(conf.Cipher.InvocationCounter() + 1)
			} else {
				secret = g.Secret
			}
			conf.Connected |= enums.ConnectionStateDlms
			if conf.Authentication == enums.AuthenticationHighSHA256 {
				tmp := types.GXByteBuffer{}
				err := tmp.Set(g.Secret)
				if err != nil {
					return nil, err
				}
				err = tmp.Set(conf.Cipher.SystemTitle())
				if err != nil {
					return nil, err
				}
				err = tmp.Set(conf.SourceSystemTitle())
				if err != nil {
					return nil, err
				}
				err = tmp.Set(conf.CtoSChallenge())
				if err != nil {
					return nil, err
				}
				err = tmp.Set(conf.StoCChallenge())
				if err != nil {
					return nil, err
				}
				secret = tmp.Array()
			}
			return settings.Secure(conf, conf.Cipher, ic, conf.CtoSChallenge(), secret)
		} else {
			conf.Connected &= ^enums.ConnectionStateDlms
			return nil, nil
		}
	}
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, nil
}

// PostLoad returns the handle actions after Load.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSAssociationShortName) PostLoad(reader *GXXmlReader) error {
	return nil
}

// GetValues returns the an array containing the COSEM object's attribute values.
func (g *GXDLMSAssociationShortName) GetValues() []any {
	return []any{g.LogicalName, g.ObjectList, nil, g.SecuritySetupReference}
}

// GetDataType returns the device data type_ of selected attribute index.
//
// Parameters:
//
//	index: Attribute index of the object.
//
// Returns:
//
//	Device data type_ of the object.
func (g *GXDLMSAssociationShortName) GetDataType(index int) (enums.DataType, error) {
	switch index {
	case 1:
		return enums.DataTypeOctetString, nil
	case 2:
		return enums.DataTypeArray, nil
	case 3:
		return enums.DataTypeArray, nil
	case 4:
		return enums.DataTypeOctetString, nil
	}
	return 0, errors.New("GetDataType failed. Invalid attribute index.")
}

// Constructor.
// ln: Logical Name of the object.
// sn: Short Name of the object.
func NewGXDLMSAssociationShortName(ln string, sn int16) (*GXDLMSAssociationShortName, error) {
	err := ValidateLogicalName(ln)
	if err != nil {
		return nil, err
	}
	return &GXDLMSAssociationShortName{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeAssociationShortName,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
