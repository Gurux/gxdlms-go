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
// Gurux Device Framework is Open Source software you can redistribute it
// and/or modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation version 2 of the License.
// Gurux Device Framework is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY without even the implied warranty of
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
	"fmt"
	"strconv"
	"strings"

	"github.com/Gurux/gxcommon-go"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/constants"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// Online help:
// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSAssociationLogicalName
type GXDLMSAssociationLogicalName struct {
	GXDLMSObject
	accessRights map[IGXDLMSBase][]int

	methodAccessRights map[IGXDLMSBase][]int

	// Is this association including other association views.
	MultipleAssociationViews bool

	ObjectList GXDLMSObjectCollection

	// Contains the identifiers of the COSEM client APs within the physical devices hosting these APs,
	// which belong to the AA modelled by the Association LN object.
	ClientSAP int8

	// Contains the identifiers of the COSEM server (logical device) APs within the physical
	// devices hosting these APs, which belong to the AA modelled by the Association LN object.
	ServerSAP uint16

	ApplicationContextName GXApplicationContextName

	XDLMSContextInfo GXxDLMSContextType

	AuthenticationMechanismName GXAuthenticationMechanismName

	// Low Level Security secret.
	Secret []byte

	AssociationStatus enums.AssociationStatus

	SecuritySetupReference string

	UserList []*types.GXKeyValuePair[byte, string]

	CurrentUser *types.GXKeyValuePair[byte, string]
}

// base returns the base GXDLMSObject of the object.
func (g *GXDLMSAssociationLogicalName) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

// Invoke returns the invokes method.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Invoke parameters.
func (g *GXDLMSAssociationLogicalName) Invoke(settings *settings.GXDLMSSettings,
	e *internal.ValueEventArgs) ([]byte, error) {
	switch e.Index {
	case 1:
		return g.replyToHlsAuthentication(settings, e)
	case 2:
		g.changeHlsSecret(e)
		return nil, nil
	case 3:
		g.addObject(settings, e)
		return nil, nil
	case 4:
		g.removeObject(settings, e)
		return nil, nil
	case 5:
		g.addUser(e)
		return nil, nil
	case 6:
		g.removeUser(e)
		return nil, nil
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil, nil
}

func (g *GXDLMSAssociationLogicalName) replyToHlsAuthentication(s *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	ic := uint32(0)
	var err error
	var secret []byte
	equals := false
	clientChallenge := e.Parameters.([]byte)
	switch s.Authentication {
	case enums.AuthenticationHighGMAC:
		secret = s.SourceSystemTitle()
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
		err = tmp.Set(s.SourceSystemTitle())
		if err != nil {
			return nil, err
		}
		err = tmp.Set(s.Cipher.SystemTitle())
		if err != nil {
			return nil, err
		}
		err = tmp.Set(s.StoCChallenge())
		if err != nil {
			return nil, err
		}
		err = tmp.Set(s.CtoSChallenge())
		if err != nil {
			return nil, err
		}
		secret = tmp.Array()
	case enums.AuthenticationHighECDSA:
		secret = nil
		tmp := types.GXByteBuffer{}
		err = tmp.Set(g.Secret)
		if err != nil {
			return nil, err
		}
		err = tmp.Set(s.SourceSystemTitle())
		if err != nil {
			return nil, err
		}
		err = tmp.Set(s.Cipher.SystemTitle())
		if err != nil {
			return nil, err
		}
		err = tmp.Set(s.StoCChallenge())
		if err != nil {
			return nil, err
		}
		err = tmp.Set(s.CtoSChallenge())
		if err != nil {
			return nil, err
		}
		key := s.Cipher.SigningKeyPair().Value
		pub := s.Cipher.SigningKeyPair().Key
		if key == nil {
			key = s.GetKey(enums.CertificateTypeDigitalSignature, s.Cipher.SystemTitle(), true).(*types.GXPrivateKey)
			s.Cipher.SetKeyAgreementKeyPair(types.NewGXKeyValuePair(pub, key))
		}
		if pub == nil {
			pub = s.GetKey(enums.CertificateTypeDigitalSignature, s.SourceSystemTitle(), false).(*types.GXPublicKey)
			s.Cipher.SetKeyAgreementKeyPair(types.NewGXKeyValuePair(pub, key))
		}
		if pub == nil {
			pub = s.GetKey(enums.CertificateTypeDigitalSignature, s.SourceSystemTitle(), false).(*types.GXPublicKey)
			s.Cipher.SetKeyAgreementKeyPair(types.NewGXKeyValuePair(pub, key))
		}
		if key == nil {
			return nil, errors.New("Signing key is not set.")
		}
		sig, err := types.NewGXEcdsaFromPublicKey(pub)
		if err != nil {
			return nil, err
		}
		equals, err = sig.Verify(clientChallenge, tmp.Array())
		if err != nil {
			return nil, err
		}
	default:
		secret = g.Secret
	}
	if s.Authentication != enums.AuthenticationHighECDSA {
		serverChallenge, err := settings.Secure(s, s.Cipher, ic, s.StoCChallenge(), secret)
		if err != nil {
			return nil, err
		}
		equals = serverChallenge != nil && clientChallenge != nil && bytes.Equal(serverChallenge, clientChallenge)
	}
	if equals {
		if s.Authentication == enums.AuthenticationHighGMAC {
			secret = s.Cipher.SystemTitle()
			ic = s.Cipher.InvocationCounter()
			s.Cipher.SetInvocationCounter(s.Cipher.InvocationCounter() + 1)
		} else {
			secret = g.Secret
		}
		g.AssociationStatus = enums.AssociationStatusAssociated
		if s.Authentication == enums.AuthenticationHighSHA256 || s.Authentication == enums.AuthenticationHighECDSA {
			tmp := types.GXByteBuffer{}
			if s.Authentication == enums.AuthenticationHighSHA256 {
				err = tmp.Set(g.Secret)
				if err != nil {
					return nil, err
				}
			}
			err = tmp.Set(s.Cipher.SystemTitle())
			if err != nil {
				return nil, err
			}
			err = tmp.Set(s.SourceSystemTitle())
			if err != nil {
				return nil, err
			}
			err = tmp.Set(s.CtoSChallenge())
			if err != nil {
				return nil, err
			}
			err = tmp.Set(s.StoCChallenge())
			if err != nil {
				return nil, err
			}
			secret = tmp.Array()
		}
		return settings.Secure(s, s.Cipher, ic, s.CtoSChallenge(), secret)
	}
	g.AssociationStatus = enums.AssociationStatusNonAssociated
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, nil
}

func (g *GXDLMSAssociationLogicalName) removeUser(e *internal.ValueEventArgs) {
	tmp, _ := e.Parameters.([]any)
	if tmp == nil || len(tmp) != 2 {
		e.Error = enums.ErrorCodeReadWriteDenied
	} else {
		g.UserList = internal.Remove(g.UserList, types.NewGXKeyValuePair(tmp[0].(byte), tmp[1].(string)))
	}
}

func (g *GXDLMSAssociationLogicalName) addUser(e *internal.ValueEventArgs) {
	tmp, _ := e.Parameters.([]any)
	if tmp == nil || len(tmp) != 2 {
		e.Error = enums.ErrorCodeReadWriteDenied
	} else {
		g.UserList = append(g.UserList, types.NewGXKeyValuePair(tmp[0].(byte), tmp[1].(string)))
	}
}

func (g *GXDLMSAssociationLogicalName) removeObject(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) {
	// Remove COSEM object.
	obj := g.getObject(settings, e.Parameters.([]any), false)
	// Unknown objects are not removed.
	if obj != nil {
		t := g.ObjectList.FindByLN(obj.Base().ObjectType(), obj.Base().LogicalName())
		if t != nil {
			g.ObjectList = internal.Remove(g.ObjectList, t)
		}
	}
}

func (g *GXDLMSAssociationLogicalName) addObject(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) {
	// Add COSEM object.
	obj := g.getObject(settings, e.Parameters.([]any), true)
	// Unknown objects are not add.
	if obj != nil {
		exists := g.ObjectList.FindByLN(obj.Base().ObjectType(), obj.Base().LogicalName())
		// Add object to object list if it not exists yet.
		if exists == nil {
			g.ObjectList = append(g.ObjectList, obj)
		}
		if settings.IsServer() {
			// Object can be added only once for the association view.
			if exists != nil {
				e.Error = enums.ErrorCodeUndefinedObject
				return
			}
			if ln, ok := obj.(*GXDLMSAssociationLogicalName); ok {
				if ln.Base().LogicalName() == "0.0.40.0.0.255" {
					e.Error = enums.ErrorCodeUndefinedObject
					return
				}
				//All LN objects are using the same version.
				obj.Base().Version = g.Version
				ln.XDLMSContextInfo.Conformance = g.XDLMSContextInfo.Conformance
				ln.XDLMSContextInfo.MaxReceivePduSize = g.XDLMSContextInfo.MaxReceivePduSize
				ln.XDLMSContextInfo.MaxSendPduSize = g.XDLMSContextInfo.MaxSendPduSize
				if exists == nil {
					ln.ObjectList.Add(ln)
				}
			} else if ss, ok := obj.(*GXDLMSSecuritySetup); ok {
				// Update server system title and keys.
				ss.ServerSystemTitle = settings.Cipher.SystemTitle()
				ss.guek = settings.Cipher.BlockCipherKey()
				ss.gbek = settings.Cipher.BroadcastBlockCipherKey()
				ss.gak = settings.Cipher.AuthenticationKey()
				ss.Kek = settings.Kek
			}
			count := obj.GetAttributeCount()
			list := make([]int, count)
			for pos := 0; pos != count; pos++ {
				if g.Version == 3 {
					list[pos] = int(obj.Base().GetAccess3(1 + pos))
				} else {
					list[pos] = int(obj.Base().GetAccess(1 + pos))
				}
			}
			g.accessRights[obj] = list
			count = obj.GetMethodCount()
			list = make([]int, count)
			for pos := 0; pos != count; pos++ {
				if g.Version == 3 {
					list[pos] = int(obj.Base().GetMethodAccess3(1 + pos))
				} else {
					list[pos] = int(obj.Base().GetMethodAccess(1 + pos))
				}
			}
			g.methodAccessRights[obj] = list
		}
	}
}

func (g *GXDLMSAssociationLogicalName) changeHlsSecret(e *internal.ValueEventArgs) {
	tmp, _ := e.Parameters.([]byte)
	if len(tmp) == 0 {
		e.Error = enums.ErrorCodeReadWriteDenied
	} else {
		g.Secret = tmp
	}
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
func (g *GXDLMSAssociationLogicalName) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	// LN is static and read only once.
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	// ObjectList is static and read only once.
	if all || len(g.ObjectList) == 0 {
		attributes = append(attributes, 2)
	}
	// associated_partners_id is static and read only once.
	if all || !g.IsRead(3) {
		attributes = append(attributes, 3)
	}
	// Application Context Name is static and read only once.
	if all || !g.IsRead(4) {
		attributes = append(attributes, 4)
	}
	// xDLMS Context Info
	if all || !g.IsRead(5) {
		attributes = append(attributes, 5)
	}
	// Authentication Mechanism Name
	if all || !g.IsRead(6) {
		attributes = append(attributes, 6)
	}
	// LLS Secret
	if all || !g.IsRead(7) {
		attributes = append(attributes, 7)
	}
	// Association Status
	if all || !g.IsRead(8) {
		attributes = append(attributes, 8)
	}
	// Security Setup Reference is from version 1.
	if g.Version > 0 && (all || !g.IsRead(9)) {
		attributes = append(attributes, 9)
	}
	// User list and current user are in version 2.
	if g.Version > 1 {
		if all || !g.IsRead(10) {
			attributes = append(attributes, 10)
		}
		if all || !g.IsRead(11) {
			attributes = append(attributes, 11)
		}
	}
	return attributes
}

// GetNames returns the names of attribute indexes.
func (g *GXDLMSAssociationLogicalName) GetNames() []string {
	if g.Version == 0 {
		return []string{"Logical Name", "Object List", "Associated partners Id", "Application Context Name", "xDLMS Context Info", "Authentication Mechanism Name", "Secret", "Association Status"}
	}
	if g.Version == 1 {
		return []string{"Logical Name", "Object List", "Associated partners Id", "Application Context Name", "xDLMS Context Info", "Authentication Mechanism Name", "Secret", "Association Status", "Security Setup Reference"}
	}
	return []string{"Logical Name", "Object List", "Associated partners Id", "Application Context Name", "xDLMS Context Info", "Authentication Mechanism Name", "Secret", "Association Status", "Security Setup Reference", "UserList", "CurrentUser"}
}

// GetMethodNames returns the names of method indexes.
func (g *GXDLMSAssociationLogicalName) GetMethodNames() []string {
	if g.Version > 1 {
		return []string{"Reply to HLS authentication", "Change HLS secret", "Add object", "Remove object", "Add user", "Remove user"}
	}
	return []string{"Reply to HLS authentication", "Change HLS secret", "Add object", "Remove object"}
}

// GetAttributeCount returns the amount of attributes.
//
// Returns:
//
//	Count of attributes.
func (g *GXDLMSAssociationLogicalName) GetAttributeCount() int {
	if g.Version > 1 {
		return 11
	}
	// Security Setup Reference is from version 1.
	if g.Version > 0 {
		return 9
	}
	return 8
}

// GetMethodCount returns the amount of methods.
func (g *GXDLMSAssociationLogicalName) GetMethodCount() int {
	if g.Version > 1 {
		return 6
	}
	return 4
}

// getObjects returns the Association View.
func (g *GXDLMSAssociationLogicalName) getObjects(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (*types.GXByteBuffer, error) {
	found := false
	data := types.GXByteBuffer{}
	// Add count only for first time.
	if settings.Index == 0 {
		count := uint32(len(g.ObjectList))
		// Find current association and add it if's not found.
		if g.AssociationStatus == enums.AssociationStatusAssociated {
			for _, it := range g.ObjectList {
				if it.Base() != g.Base() && it.Base().ObjectType() == enums.ObjectTypeAssociationLogicalName {
					if it.Base().LogicalName() == "0.0.40.0.0.255" {
						found = true
					} else if !g.MultipleAssociationViews {
						count--
					}
				}
			}
			if !found {
				count++
			}
		} else {
			found = true
		}
		settings.Count = count
		err := data.SetUint8(uint8(enums.DataTypeArray))
		if err != nil {
			return nil, err
		}
		types.SetObjectCount(int(count), &data)
		// If default association view is not found.
		if !found {
			settings.Count--
			err = data.SetUint8(enums.DataTypeStructure)
			if err != nil {
				return nil, err
			}
			err = data.SetUint8(4)
			if err != nil {
				return nil, err
			}
			err := internal.SetData(settings, &data, enums.DataTypeUint16, g.ObjectType)
			if err != nil {
				return nil, err
			}
			err = internal.SetData(settings, &data, enums.DataTypeUint8, g.Version)
			if err != nil {
				return nil, err
			}
			ln, err := helpers.LogicalNameToBytes("0.0.40.0.0.255")
			if err != nil {
				return nil, err
			}
			err = internal.SetData(settings, &data, enums.DataTypeOctetString, ln)
			if err != nil {
				return nil, err
			}
			err = g.GetAccessRights(settings, g, e.Server, &data)
			if err != nil {
				return nil, err
			}
		}
	}
	pos := uint32(0)
	for _, it := range g.ObjectList {
		pos++
		if !(pos <= settings.Index) {
			if it.Base().ObjectType() == enums.ObjectTypeAssociationLogicalName {
				if !g.MultipleAssociationViews && !(it.Base() == g.Base() || it.Base().LogicalName() == "0.0.40.0.0.255") {
					settings.Index++
					continue
				}
			}
			err := data.SetUint8(uint8(enums.DataTypeStructure))
			if err != nil {
				return nil, err
			}
			err = data.SetUint8(uint8(4))
			if err != nil {
				return nil, err
			}
			err = internal.SetData(settings, &data, enums.DataTypeUint16, int(it.Base().ObjectType()))
			err = internal.SetData(settings, &data, enums.DataTypeUint8, it.Base().Version)
			ln, err := helpers.LogicalNameToBytes(it.Base().LogicalName())
			if err != nil {
				return nil, err
			}
			err = internal.SetData(settings, &data, enums.DataTypeOctetString, ln)
			err = g.GetAccessRights(settings, it, e.Server, &data)
			settings.Index++
			if settings.IsServer() {
				// If PDU is full.
				if (settings.NegotiatedConformance&enums.ConformanceGeneralBlockTransfer) == 0 && !e.SkipMaxPduSize && data.Size() >= int(settings.MaxPduSize()) {
					break
				}
			}
		}
	}
	// If all objects are read.
	if int(pos) == len(g.ObjectList) {
		settings.Count = 0
		settings.Index = 0

	}
	return &data, nil
}

func (g *GXDLMSAssociationLogicalName) GetAccessRights(settings *settings.GXDLMSSettings, item IGXDLMSBase, server internal.IGXDLMSServer, data *types.GXByteBuffer) error {
	err := data.SetUint8(uint8(enums.DataTypeStructure))
	if err != nil {
		return err
	}
	err = data.SetUint8(uint8(2))
	if err != nil {
		return err
	}
	err = data.SetUint8(uint8(enums.DataTypeArray))
	if err != nil {
		return err
	}
	cnt := item.GetAttributeCount()
	err = data.SetUint8(uint8(cnt))
	if err != nil {
		return err
	}
	var e *internal.ValueEventArgs
	var m uint8
	if server != nil {
		e = internal.NewValueEventArgs2(server, item, 0)
	} else {
		e = internal.NewValueEventArgs(settings, item, 0)
	}
	for pos := 0; pos != cnt; pos++ {
		e.Index = uint8(pos + 1)
		if server != nil {
			m = uint8(server.NotifyGetAttributeAccess(e))
		} else {
			if g.Version < 3 {
				m = uint8(item.Base().GetAccess(int(e.Index)))
			} else {
				m = uint8(item.Base().GetAccess3(int(e.Index)))
			}
		}
		err = data.SetUint8(uint8(enums.DataTypeStructure))
		if err != nil {
			return err
		}
		err = data.SetUint8(uint8(3))
		if err != nil {
			return err
		}
		internal.SetData(settings, data, enums.DataTypeInt8, e.Index)
		internal.SetData(settings, data, enums.DataTypeEnum, m)
		accessSelector := item.Base().GetAccessSelector(int(e.Index))
		if accessSelector == 0 {
			if _, ok := item.(*GXDLMSProfileGeneric); ok {
				accessSelector = 3
			}
		}
		if accessSelector != 0 {
			var list []any
			for index := 0; index != 8; index++ {
				if (accessSelector & (1 << index)) != 0 {
					list = append(list, index)
				}
			}
			internal.SetData(settings, data, enums.DataTypeArray, list)
		} else {
			internal.SetData(settings, data, enums.DataTypeNone, nil)
		}
	}
	err = data.SetUint8(uint8(enums.DataTypeArray))
	if err != nil {
		return err
	}
	cnt = item.GetMethodCount()
	err = data.SetUint8(uint8(cnt))
	if err != nil {
		return err
	}
	for pos := 0; pos != cnt; pos++ {
		e.Index = uint8(pos + 1)
		if server != nil {
			m = uint8(server.NotifyGetMethodAccess(e))
		} else {
			if g.Version < 3 {
				m = uint8(item.Base().GetMethodAccess(int(e.Index)))
			} else {
				m = uint8(item.Base().GetMethodAccess3(int(e.Index)))
			}
		}
		err = data.SetUint8(uint8(enums.DataTypeStructure))
		if err != nil {
			return err
		}
		err = data.SetUint8(uint8(2))
		if err != nil {
			return err
		}
		internal.SetData(settings, data, enums.DataTypeInt8, pos+1)
		internal.SetData(settings, data, enums.DataTypeEnum, m)
	}
	return err
}

func (g *GXDLMSAssociationLogicalName) UpdateAccessRights(obj IGXDLMSBase, buff types.GXStructure) {
	// Some meters return only supported access rights.All access rights are set to NoAccess.
	if g.Version < 3 {
		for pos := 0; pos != obj.GetAttributeCount(); pos++ {
			obj.Base().SetAccess(pos+1, enums.AccessModeNoAccess)
		}
		for pos := 0; pos != obj.GetMethodCount(); pos++ {
			obj.Base().SetMethodAccess(pos+1, enums.MethodAccessModeNoAccess)
		}
	} else {
		for pos := 0; pos != obj.GetAttributeCount(); pos++ {
			obj.Base().SetAccess3(pos+1, enums.AccessMode3NoAccess)
		}
		for pos := 0; pos != obj.GetMethodCount(); pos++ {
			obj.Base().SetMethodAccess3(pos+1, enums.MethodAccessMode3NoAccess)
		}
	}
	if len(buff) != 0 {
		arr := buff[0].(types.GXArray)
		for _, tmp := range arr {
			it := tmp.(types.GXStructure)
			id := int(it[0].(int8))
			mode := it[1].(types.GXEnum).Value
			if g.Version < 3 {
				obj.Base().SetAccess(id, enums.AccessMode(mode))
			} else {
				obj.Base().SetAccess3(id, enums.AccessMode3(mode))
			}
		}
		arr = buff[1].(types.GXArray)
		for _, tmp := range arr {
			it := tmp.(types.GXStructure)
			id := int(it[0].(int8))
			var tmp2 int
			// If version is 0.
			if v, ok := it[1].(bool); ok {
				if v {
					tmp2 = 1
				} else {
					tmp2 = 0
				}
			} else {
				tmp2 = int(it[1].(types.GXEnum).Value)
			}
			if g.Version < 3 {
				obj.Base().SetMethodAccess(id, enums.MethodAccessMode(tmp2))
			} else {
				obj.Base().SetMethodAccess3(id, enums.MethodAccessMode3(tmp2))
			}
		}
	}
}

// getUserList returns the User list.
func (g *GXDLMSAssociationLogicalName) getUserList(settings *settings.GXDLMSSettings) (*types.GXByteBuffer, error) {
	data := types.GXByteBuffer{}
	// Add count only for first time.
	if settings.Index == 0 {
		settings.Count = uint32(len(g.UserList))
		err := data.SetUint8(uint8(enums.DataTypeArray))
		if err != nil {
			return nil, err
		}
		types.SetObjectCount(len(g.UserList), &data)
	}
	pos := 0
	for k, v := range g.UserList {
		pos++
		if !(pos <= int(settings.Index)) {
			settings.Index++
			err := data.SetUint8(uint8(enums.DataTypeStructure))
			if err != nil {
				return nil, err
			}
			err = data.SetUint8(2)
			if err != nil {
				return nil, err
			}
			err = internal.SetData(settings, &data, enums.DataTypeUint8, k)
			if err != nil {
				return nil, err
			}
			err = internal.SetData(settings, &data, enums.DataTypeString, v)
			if err != nil {
				return nil, err
			}
		}
	}
	return &data, nil
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
func (g *GXDLMSAssociationLogicalName) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	if e.Index == 1 {
		v, err := helpers.LogicalNameToBytes(g.LogicalName())
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return v, err
	}
	if e.Index == 2 {
		e.ByteArray = true
		return g.getObjects(settings, e)
	}
	if e.Index == 3 {
		e.ByteArray = true
		data := types.GXByteBuffer{}
		err := data.SetUint8(uint8(enums.DataTypeStructure))
		if err != nil {
			return nil, err
		}
		err = data.SetUint8(2)
		if err != nil {
			return nil, err
		}
		err = data.SetUint8(uint8(enums.DataTypeInt8))
		if err != nil {
			return nil, err
		}
		err = data.SetUint8(uint8(g.ClientSAP))
		if err != nil {
			return nil, err
		}
		err = data.SetUint8(uint8(enums.DataTypeUint16))
		if err != nil {
			return nil, err
		}
		err = data.SetUint16(g.ServerSAP)
		if err != nil {
			return nil, err
		}
		return data, nil
	}
	if e.Index == 4 {
		e.ByteArray = true
		data := types.GXByteBuffer{}
		err := data.SetUint8(uint8(enums.DataTypeStructure))
		if err != nil {
			return nil, err
		}
		err = data.SetUint8(0x7)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, &data, enums.DataTypeUint8, g.ApplicationContextName.JointIsoCtt)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, &data, enums.DataTypeUint8, g.ApplicationContextName.Country)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, &data, enums.DataTypeUint16, g.ApplicationContextName.CountryName)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, &data, enums.DataTypeUint8, g.ApplicationContextName.IdentifiedOrganization)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, &data, enums.DataTypeUint8, g.ApplicationContextName.DlmsUA)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, &data, enums.DataTypeUint8, g.ApplicationContextName.ApplicationContext)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, &data, enums.DataTypeUint8, g.ApplicationContextName.ContextId)
		if err != nil {
			return nil, err
		}
		return data.Array(), nil
	}
	if e.Index == 5 {
		e.ByteArray = true
		data := types.GXByteBuffer{}
		err := data.SetUint8(uint8(enums.DataTypeStructure))
		if err != nil {
			return nil, err
		}
		err = data.SetUint8(6)
		if err != nil {
			return nil, err
		}
		bs, err := types.NewGXBitStringFromInteger(int(g.XDLMSContextInfo.Conformance), 24)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, &data, enums.DataTypeBitString, bs)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, &data, enums.DataTypeUint16, g.XDLMSContextInfo.MaxReceivePduSize)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, &data, enums.DataTypeUint16, g.XDLMSContextInfo.MaxSendPduSize)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, &data, enums.DataTypeUint8, g.XDLMSContextInfo.DlmsVersionNumber)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, &data, enums.DataTypeInt8, g.XDLMSContextInfo.QualityOfService)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, &data, enums.DataTypeOctetString, g.XDLMSContextInfo.CypheringInfo)
		if err != nil {
			return nil, err
		}
		return data.Array(), nil
	}
	if e.Index == 6 {
		e.ByteArray = true
		data := types.GXByteBuffer{}
		err := data.SetUint8(uint8(enums.DataTypeStructure))
		if err != nil {
			return nil, err
		}
		err = data.SetUint8(0x7)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, &data, enums.DataTypeUint8, g.AuthenticationMechanismName.JointIsoCtt)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, &data, enums.DataTypeUint8, g.AuthenticationMechanismName.Country)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, &data, enums.DataTypeUint16, g.AuthenticationMechanismName.CountryName)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, &data, enums.DataTypeUint8, g.AuthenticationMechanismName.IdentifiedOrganization)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, &data, enums.DataTypeUint8, g.AuthenticationMechanismName.DlmsUA)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, &data, enums.DataTypeUint8, g.AuthenticationMechanismName.AuthenticationMechanismName)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, &data, enums.DataTypeUint8, g.AuthenticationMechanismName.MechanismId)
		if err != nil {
			return nil, err
		}
		return data.Array(), nil
	}
	if e.Index == 7 {
		return g.Secret, nil
	}
	if e.Index == 8 {
		return g.AssociationStatus, nil
	}
	if e.Index == 9 {
		return helpers.LogicalNameToBytes(g.SecuritySetupReference)
	}
	if e.Index == 10 {
		return g.getUserList(settings)
	}
	if e.Index == 11 {
		e.ByteArray = true
		data := types.GXByteBuffer{}
		err := data.SetUint8(uint8(enums.DataTypeStructure))
		if err != nil {
			return nil, err
		}
		err = data.SetUint8(2)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, &data, enums.DataTypeUint8, g.CurrentUser.Key)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, &data, enums.DataTypeString, g.CurrentUser.Value)
		if err != nil {
			return nil, err
		}
		return data.Array(), nil
	}
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, nil
}

// getObject returns the get object.
//
// Parameters:
//
//	settings: DLMS settings.
//	item: Received data.
//	add: Is data added to settings object list.
func (g *GXDLMSAssociationLogicalName) getObject(settings *settings.GXDLMSSettings, item []any, add bool) IGXDLMSBase {

	type_ := enums.ObjectType(item[0].(uint16))
	version := item[1].(uint8)
	ln, err := helpers.ToLogicalName(item[2].([]byte))
	if err != nil {
		return nil
	}
	var obj IGXDLMSBase
	if settings != nil && g.AssociationStatus == enums.AssociationStatusAssociated {
		obj = getObjectCollection(settings.Objects).FindByLN(type_, ln)
	}
	if obj == nil {
		obj = CreateObject(type_)
		if obj != nil {
			obj.Base().SetLogicalName(ln)
			obj.Base().Version = version
			if add && settings.IsServer() {
				objects := getObjectCollection(settings.Objects)
				objects.Add(obj)
			}
		}
	}
	if obj != nil {
		arr := item[3].(types.GXStructure)
		g.UpdateAccessRights(obj, arr)
	}
	return obj
}

// SetValue returns the set value of given attribute.
// When raw parameter us not used example register multiplies value by scalar.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Set parameters.
func (g *GXDLMSAssociationLogicalName) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	switch e.Index {
	case 1:
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return g.Base().SetLogicalName(ln)
	case 2:
		g.UpdateObjectList(settings, e)
	case 3:
		if e.Value != nil {
			arr := e.Value.(types.GXStructure)
			g.ClientSAP = arr[0].(int8)
			g.ServerSAP = arr[1].(uint16)
		}
	case 4:
		if v, ok := e.Value.([]byte); ok {
			// Value of the object identifier encoded in BER
			arr := types.NewGXByteBufferWithData(v)
			ret, err := arr.Uint8At(0)
			if err != nil {
				return err
			}
			if ret == 0x60 {
				// BB 11.4.
				val, err := arr.Uint8()
				g.ApplicationContextName.JointIsoCtt = (byte)((val - 16) / 40)
				g.ApplicationContextName.Country = 16
				tmp, err := arr.Uint16()
				tmp = ((tmp & 0x7F00) >> 1) | (tmp & 0x7F)
				g.ApplicationContextName.CountryName = uint16(tmp)

				g.ApplicationContextName.CountryName = uint16(tmp)

				val, err = arr.Uint8()
				if err != nil {
					return err
				}
				g.ApplicationContextName.IdentifiedOrganization = val

				val, err = arr.Uint8()
				if err != nil {
					return err
				}
				g.ApplicationContextName.DlmsUA = val
				val, err = arr.Uint8()
				if err != nil {
					return err
				}
				g.ApplicationContextName.ApplicationContext = val
				val, err = arr.Uint8()
				if err != nil {
					return err
				}
				g.ApplicationContextName.ContextId = enums.ApplicationContextName(val)
				if g.ApplicationContextName.ContextId > enums.ApplicationContextNameShortNameWithCiphering {
					e.Error = enums.ErrorCodeReadWriteDenied
				}
			} else {
				// Get Tag and Len.
				ret, err := arr.Uint8()
				if err != nil {
					return err
				}
				ret2, err := arr.Uint8()
				if err != nil {
					return err
				}
				if ret != uint8(constants.BerTypeInteger) && ret2 != 7 {
					return gxcommon.ErrInvalidArgument
				}
				// Get tag
				ret, err = arr.Uint8()
				if err != nil {
					return err
				}
				if ret != 0x11 {
					return gxcommon.ErrInvalidArgument
				}
				ret, err = arr.Uint8()
				if err != nil {
					return err
				}
				g.ApplicationContextName.JointIsoCtt = ret
				// Get tag
				ret, err = arr.Uint8()
				if err != nil {
					return err
				}
				if ret != 0x11 {
					return gxcommon.ErrInvalidArgument
				}
				ret, err = arr.Uint8()
				if err != nil {
					return err
				}
				g.ApplicationContextName.Country = ret
				// Get tag
				ret, err = arr.Uint8()
				if err != nil {
					return err
				}
				if ret != 0x12 {
					return gxcommon.ErrInvalidArgument
				}
				ret3, err := arr.Uint16()
				if err != nil {
					return err
				}
				g.ApplicationContextName.CountryName = ret3
				// Get tag
				ret, err = arr.Uint8()
				if err != nil {
					return err
				}
				if ret != 0x11 {
					return gxcommon.ErrInvalidArgument
				}
				ret, err = arr.Uint8()
				if err != nil {
					return err
				}
				g.ApplicationContextName.IdentifiedOrganization = ret

				// Get tag
				ret, err = arr.Uint8()
				if err != nil {
					return err
				}
				if ret != 0x11 {
					return gxcommon.ErrInvalidArgument
				}
				ret, err = arr.Uint8()
				if err != nil {
					return err
				}
				g.ApplicationContextName.DlmsUA = ret

				// Get tag
				ret, err = arr.Uint8()
				if err != nil {
					return err
				}
				if ret != 0x11 {
					return gxcommon.ErrInvalidArgument
				}
				ret, err = arr.Uint8()
				if err != nil {
					return err
				}
				g.ApplicationContextName.ApplicationContext = ret
				// Get tag
				ret, err = arr.Uint8()
				if err != nil {
					return err
				}
				if ret != 0x11 {
					return gxcommon.ErrInvalidArgument
				}
				ret, err = arr.Uint8()
				if err != nil {
					return err
				}
				g.ApplicationContextName.ContextId = enums.ApplicationContextName(ret)
			}
		} else if e.Value != nil {
			arr := e.Value.(types.GXStructure)
			g.ApplicationContextName.JointIsoCtt = arr[0].(byte)
			g.ApplicationContextName.Country = arr[1].(byte)
			g.ApplicationContextName.CountryName = arr[2].(uint16)
			g.ApplicationContextName.IdentifiedOrganization = arr[3].(byte)
			g.ApplicationContextName.DlmsUA = arr[4].(byte)
			g.ApplicationContextName.ApplicationContext = arr[5].(byte)
			g.ApplicationContextName.ContextId = enums.ApplicationContextName(arr[6].(byte))
			if g.ApplicationContextName.ContextId > enums.ApplicationContextNameShortNameWithCiphering {
				return gxcommon.ErrInvalidArgument
			}
		}
	case 5:
		if e.Value != nil {
			arr := e.Value.(types.GXStructure)
			c := arr[0].(types.GXBitString)
			g.XDLMSContextInfo.Conformance = enums.Conformance(c.ToInteger())
			g.XDLMSContextInfo.MaxReceivePduSize = arr[1].(uint16)
			g.XDLMSContextInfo.MaxSendPduSize = arr[2].(uint16)
			g.XDLMSContextInfo.DlmsVersionNumber = arr[3].(uint8)
			g.XDLMSContextInfo.QualityOfService = arr[4].(int8)
			g.XDLMSContextInfo.CypheringInfo = arr[5].([]byte)
			if g.XDLMSContextInfo.Conformance == enums.ConformanceNone {
				return gxcommon.ErrInvalidArgument
			}
			if g.XDLMSContextInfo.MaxReceivePduSize < 64 || g.XDLMSContextInfo.MaxSendPduSize < 64 {
				return gxcommon.ErrInvalidArgument
			}
		}
	case 6:
		// Value of the object identifier encoded in BER
		if v, ok := e.Value.([]byte); ok {
			arr := types.NewGXByteBufferWithData(v)
			ret, err := arr.Uint8At(0)
			if err != nil {
				return err
			}
			if ret == 0x60 {
				// BB 11.4.
				val, err := arr.Uint8()
				if err != nil {
					return err
				}
				g.AuthenticationMechanismName.JointIsoCtt = uint8(((val - 16) / 40))
				g.AuthenticationMechanismName.Country = 16
				tmp, err := arr.Uint16()
				if err != nil {
					return err
				}
				tmp = ((tmp & 0x7F00) >> 1) | (tmp & 0x7F)
				g.AuthenticationMechanismName.CountryName = uint16(tmp)
				ret, err := arr.Uint8()
				if err != nil {
					return err
				}
				g.AuthenticationMechanismName.IdentifiedOrganization = ret
				ret, err = arr.Uint8()
				if err != nil {
					return err
				}
				g.AuthenticationMechanismName.DlmsUA = ret
				ret, err = arr.Uint8()
				if err != nil {
					return err
				}
				g.AuthenticationMechanismName.AuthenticationMechanismName = ret
				ret, err = arr.Uint8()
				if err != nil {
					return err
				}
				g.AuthenticationMechanismName.MechanismId = enums.Authentication(ret)
				if g.AuthenticationMechanismName.MechanismId > enums.AuthenticationHighECDSA {
					return gxcommon.ErrInvalidArgument
				}
			} else {
				// Get Tag and Len.
				tag, err := arr.Uint8()
				if err != nil {
					return err
				}
				len_, err := arr.Uint8()
				if err != nil {
					return err
				}
				if tag != uint8(constants.BerTypeInteger) && len_ != 7 {
					return gxcommon.ErrInvalidArgument
				}
				// Get tag
				tag, err = arr.Uint8()
				if err != nil {
					return err
				}
				if tag != 0x11 {
					return gxcommon.ErrInvalidArgument
				}
				tag, err = arr.Uint8()
				if err != nil {
					return err
				}

				g.AuthenticationMechanismName.JointIsoCtt = tag
				// Get tag
				tag, err = arr.Uint8()
				if err != nil {
					return err
				}

				if tag != 0x11 {
					return gxcommon.ErrInvalidArgument
				}
				tag, err = arr.Uint8()
				if err != nil {
					return err
				}

				g.AuthenticationMechanismName.Country = tag
				// Get tag
				tag, err = arr.Uint8()
				if err != nil {
					return err
				}
				if tag != 0x12 {
					return gxcommon.ErrInvalidArgument
				}
				val, err := arr.Uint16()
				if err != nil {
					return err
				}
				g.AuthenticationMechanismName.CountryName = val
				// Get tag
				tag, err = arr.Uint8()
				if err != nil {
					return err
				}

				if tag != 0x11 {
					return gxcommon.ErrInvalidArgument
				}
				tag, err = arr.Uint8()
				if err != nil {
					return err
				}

				g.AuthenticationMechanismName.IdentifiedOrganization = tag
				// Get tag
				tag, err = arr.Uint8()
				if err != nil {
					return err
				}
				if tag != 0x11 {
					return gxcommon.ErrInvalidArgument
				}
				tag, err = arr.Uint8()
				if err != nil {
					return err
				}
				g.AuthenticationMechanismName.DlmsUA = tag
				// Get tag
				tag, err = arr.Uint8()
				if err != nil {
					return err
				}

				if tag != 0x11 {
					return gxcommon.ErrInvalidArgument
				}
				tag, err = arr.Uint8()
				if err != nil {
					return err
				}
				g.AuthenticationMechanismName.AuthenticationMechanismName = tag
				// Get tag
				tag, err = arr.Uint8()
				if err != nil {
					return err
				}

				if tag != 0x11 {
					return gxcommon.ErrInvalidArgument
				}
				tag, err = arr.Uint8()
				if err != nil {
					return err
				}
				g.AuthenticationMechanismName.MechanismId = enums.Authentication(tag)
				if g.AuthenticationMechanismName.MechanismId > enums.AuthenticationHighECDSA {
					return gxcommon.ErrInvalidArgument
				}
			}
		} else if e.Value != nil {
			arr := e.Value.(types.GXStructure)
			g.AuthenticationMechanismName.JointIsoCtt = arr[0].(byte)
			g.AuthenticationMechanismName.Country = arr[1].(byte)
			g.AuthenticationMechanismName.CountryName = arr[2].(uint16)
			g.AuthenticationMechanismName.IdentifiedOrganization = arr[3].(byte)
			g.AuthenticationMechanismName.DlmsUA = arr[4].(byte)
			g.AuthenticationMechanismName.AuthenticationMechanismName = arr[5].(byte)
			g.AuthenticationMechanismName.MechanismId = enums.Authentication(arr[6].(byte))
			if g.AuthenticationMechanismName.MechanismId > enums.AuthenticationHighECDSA {
				return gxcommon.ErrInvalidArgument
			}
		}
	case 7:
		if e.Value != nil {
			g.Secret = e.Value.([]byte)
		} else {
			g.Secret = nil
		}
	case 8:
		if e.Value == nil {
			g.AssociationStatus = enums.AssociationStatusNonAssociated
		} else {
			g.AssociationStatus = enums.AssociationStatus(e.Value.(types.GXEnum).Value)
		}
	case 9:
		var err error
		g.SecuritySetupReference, err = helpers.ToLogicalName(e.Value)
		if err != nil {
			return err
		}
	case 10:
		g.UserList = g.UserList[:0]
		if e.Value != nil {
			for _, tmp := range e.Value.(types.GXArray) {
				item := tmp.(types.GXStructure)
				g.UserList = append(g.UserList, types.NewGXKeyValuePair(item[0].(byte), item[1].(string)))
			}
		}
	case 11:
		if e.Value != nil {
			arr := e.Value.(types.GXStructure)
			if len(arr) == 1 {
				g.CurrentUser = types.NewGXKeyValuePair(arr[0].(byte), "")
			} else {
				var user string
				// Some meters are sending current user as a string.
				if _, ok := arr[1].([]byte); ok {
					user = string(arr[1].([]byte))
				} else {
					user = arr[1].(string)
				}
				g.CurrentUser = types.NewGXKeyValuePair(arr[0].(byte), user)
			}
		} else {
			g.CurrentUser = types.NewGXKeyValuePair(byte(0), "")
		}
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil
}

func (g *GXDLMSAssociationLogicalName) UpdateObjectList(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) {
	g.ObjectList.Clear()
	if e.Value != nil {
		arr := e.Value.(types.GXArray)
		for _, tmp := range arr {
			item := tmp.(types.GXStructure)
			obj := g.getObject(settings, item, true)
			// Unknown objects are not shown.
			if obj != nil {
				g.ObjectList.Add(obj)
			}
		}
	}
}

// Load returns the load object content from XML.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSAssociationLogicalName) Load(reader *GXXmlReader) error {
	var str string
	var buff []int
	g.ObjectList.Clear()
	if ret, _ := reader.IsStartElementNamed("ObjectList", true); ret {
		var target string
		var obj IGXDLMSBase
		for !reader.EOF() {
			obj = nil
			if reader.IsStartElement() {
				target = reader.Name()
				if strings.HasPrefix(target, "GXDLMS") {
					str = target[6:]
					reader.Read()
					type_, err := enums.ObjectTypeParse(str)
					if err != nil {
						return err
					}
					ln, err := reader.ReadElementContentAsString("LN", "")
					if err != nil {
						return err
					}
					obj = reader.Objects.FindByLN(type_, ln)
					if obj == nil {
						obj = CreateObject(type_)
						obj.Base().Version = 0
						obj.Base().SetLogicalName(ln)
						reader.Objects.Add(obj)
					}
					if obj.Base() != g.Base() {
						g.ObjectList.Add(obj)
					}
					// methodAccessRights
					access, err := reader.ReadElementContentAsString("Access", "")
					if err != nil {
						return err
					}
					pos := 0
					if access != "" {
						buff = make([]int, len(access))
						for _, it := range access {
							buff[pos] = int(it - 0x30)
							pos++
						}
						g.accessRights[obj] = buff
						pos = 0
					}
					access, err = reader.ReadElementContentAsString("Access3", "")
					if err != nil {
						return err
					}
					if access != "" {
						buff = make([]int, len(access)/4)
						for pos := 0; pos != len(buff); pos++ {
							ret, err := strconv.ParseInt(access[4*pos:4+4*pos], 16, 16)
							if err != nil {
								return err
							}
							buff[pos] = int(ret & ^0x8000)
						}
						g.accessRights[obj] = buff
						pos = 0
					}
					access, err = reader.ReadElementContentAsString("MethodAccess", "")
					if err != nil {
						return err
					}
					if access != "" {
						buff = make([]int, len(access))
						for _, it := range access {
							buff[pos] = int(it - 0x30)
							pos++
						}
						g.methodAccessRights[obj] = buff
					}
					access, err = reader.ReadElementContentAsString("MethodAccess3", "")
					if err != nil {
						return err
					}
					if access != "" {
						buff = make([]int, len(access)/4)
						for pos = 0; pos != len(buff); pos++ {
							ret, err := strconv.ParseInt(access[4*pos:4+4*pos], 16, 16)
							if err != nil {
								return err
							}
							buff[pos] = int(ret & ^0x8000)
						}
						g.methodAccessRights[obj] = buff
					}
				}
			} else {
				if reader.Name() == "ObjectList" {
					break
				}
				reader.Read()
			}
		}
		reader.ReadEndElement("ObjectList")
	}
	if g.ObjectList.FindByLN(enums.ObjectTypeAssociationLogicalName, g.Base().LogicalName()) == nil {
		g.ObjectList.Add(g)
	}
	ret, err := reader.ReadElementContentAsInt("ClientSAP", 0)
	if err != nil {
		return err
	}

	g.ClientSAP = int8(ret)
	ret, err = reader.ReadElementContentAsInt("ServerSAP", 0)
	if err != nil {
		return err
	}
	g.ServerSAP = uint16(ret)
	if ret, _ := reader.IsStartElementNamed("ApplicationContextName", true); ret {
		ret, err := reader.ReadElementContentAsInt("JointIsoCtt", 0)
		if err != nil {
			return err
		}
		g.ApplicationContextName.JointIsoCtt = byte(ret)

		ret, err = reader.ReadElementContentAsInt("Country", 0)
		if err != nil {
			return err
		}
		g.ApplicationContextName.Country = byte(ret)

		ret, err = reader.ReadElementContentAsInt("CountryName", 0)
		if err != nil {
			return err
		}
		g.ApplicationContextName.CountryName = uint16(ret)

		ret, err = reader.ReadElementContentAsInt("IdentifiedOrganization", 0)
		if err != nil {
			return err
		}
		g.ApplicationContextName.IdentifiedOrganization = byte(ret)

		ret, err = reader.ReadElementContentAsInt("DlmsUA", 0)
		if err != nil {
			return err
		}
		g.ApplicationContextName.DlmsUA = byte(ret)

		ret, err = reader.ReadElementContentAsInt("ApplicationContext", 0)
		if err != nil {
			return err
		}
		g.ApplicationContextName.ApplicationContext = byte(ret)

		ret, err = reader.ReadElementContentAsInt("ContextId", 0)
		if err != nil {
			return err
		}
		g.ApplicationContextName.ContextId = enums.ApplicationContextName(ret)

		err = reader.ReadEndElement("ApplicationContextName")
		if err != nil {
			return err
		}
	}
	if ret, _ := reader.IsStartElementNamed("XDLMSContextInfo", true); ret {
		ret, err := reader.ReadElementContentAsInt("Conformance", 0)
		if err != nil {
			return err
		}
		g.XDLMSContextInfo.Conformance = enums.Conformance(ret)

		ret, err = reader.ReadElementContentAsInt("MaxReceivePduSize", 0)
		if err != nil {
			return err
		}
		g.XDLMSContextInfo.MaxReceivePduSize = uint16(ret)
		ret, err = reader.ReadElementContentAsInt("MaxSendPduSize", 0)
		if err != nil {
			return err
		}
		g.XDLMSContextInfo.MaxSendPduSize = uint16(ret)
		ret, err = reader.ReadElementContentAsInt("DlmsVersionNumber", 0)
		if err != nil {
			return err
		}
		g.XDLMSContextInfo.DlmsVersionNumber = uint8(ret)
		ret, err = reader.ReadElementContentAsInt("QualityOfService", 0)
		if err != nil {
			return err
		}
		g.XDLMSContextInfo.QualityOfService = int8(ret)

		str, err = reader.ReadElementContentAsString("CypheringInfo", "")
		if err != nil {
			return err
		}
		if str != "" {
			g.XDLMSContextInfo.CypheringInfo = types.HexToBytes(str)
		}
		err = reader.ReadEndElement("XDLMSContextInfo")
		if err != nil {
			return err
		}
	}
	ret1, err := reader.IsStartElementNamed("AuthenticationMechanismName", true)
	if err != nil {
		return err
	}
	ret2, err := reader.IsStartElementNamed("XDLMSContextInfo", true)
	if err != nil {
		return err
	}
	if ret1 || ret2 {
		ret, err = reader.ReadElementContentAsInt("JointIsoCtt", 0)
		if err != nil {
			return err
		}
		g.AuthenticationMechanismName.JointIsoCtt = uint8(ret)
		ret, err = reader.ReadElementContentAsInt("Country", 0)
		if err != nil {
			return err
		}
		g.AuthenticationMechanismName.Country = uint8(ret)

		ret, err = reader.ReadElementContentAsInt("CountryName", 0)
		if err != nil {
			return err
		}
		g.AuthenticationMechanismName.CountryName = uint16(ret)

		ret, err = reader.ReadElementContentAsInt("IdentifiedOrganization", 0)
		if err != nil {
			return err
		}
		g.AuthenticationMechanismName.IdentifiedOrganization = uint8(ret)

		ret, err = reader.ReadElementContentAsInt("DlmsUA", 0)
		if err != nil {
			return err
		}
		g.AuthenticationMechanismName.DlmsUA = uint8(ret)

		ret, err = reader.ReadElementContentAsInt("AuthenticationMechanismName", 0)
		if err != nil {
			return err
		}
		g.AuthenticationMechanismName.AuthenticationMechanismName = uint8(ret)

		ret, err = reader.ReadElementContentAsInt("MechanismId", 0)
		if err != nil {
			return err
		}

		g.AuthenticationMechanismName.MechanismId = enums.Authentication(ret)
		err = reader.ReadEndElement("AuthenticationMechanismName")
		if err != nil {
			return err
		}
		err = reader.ReadEndElement("XDLMSContextInfo")
		if err != nil {
			return err
		}
	}
	str, err = reader.ReadElementContentAsString("Secret", "")
	if err != nil {
		return err
	}
	if str == "" {
		g.Secret = nil
	} else {
		g.Secret = types.HexToBytes(str)
	}
	ret, err = reader.ReadElementContentAsInt("AssociationStatus", 0)
	if err != nil {
		return err
	}
	g.AssociationStatus = enums.AssociationStatus(ret)

	g.SecuritySetupReference, err = reader.ReadElementContentAsString("SecuritySetupReference", "")
	if err != nil {
		return err
	}
	g.UserList = g.UserList[:0]
	if ret, _ := reader.IsStartElementNamed("Users", true); ret {
		for true {
			ret, err := reader.IsStartElementNamed("Item", true)
			if err != nil {
				return err
			}
			if !ret {
				break
			}
			id, err := reader.ReadElementContentAsInt("Id", 0)
			if err != nil {
				return err
			}
			name, err := reader.ReadElementContentAsString("Name", "")
			if err != nil {
				return err
			}
			g.UserList = append(g.UserList, types.NewGXKeyValuePair(byte(id), name))
		}
		reader.ReadEndElement("Users")
	}
	ret, err = reader.ReadElementContentAsInt("MultipleAssociationViews", 0)
	if err != nil {
		return err
	}
	g.MultipleAssociationViews = ret != 0
	return nil
}

// Save returns the save object content to XML.
//
// Parameters:
//
//	writer: XML writer.
func (g *GXDLMSAssociationLogicalName) Save(writer *GXXmlWriter) error {
	// Save objects.
	if g.ObjectList != nil {
		writer.WriteStartElement("ObjectList")
		sb := strings.Builder{}
		if g.accessRights == nil {
			g.accessRights = make(map[IGXDLMSBase][]int)
		}
		if g.methodAccessRights == nil {
			g.methodAccessRights = make(map[IGXDLMSBase][]int)
		}
		for _, it := range g.ObjectList {
			// Default association view is not saved.
			if !(it.Base().ObjectType() == enums.ObjectTypeAssociationLogicalName && (it.Base().LogicalName() == "0.0.40.0.0.255")) {
				if g.MultipleAssociationViews || it.Base().ObjectType() != enums.ObjectTypeAssociationLogicalName {
					writer.WriteStartElement("GXDLMS" + it.Base().ObjectType().String())
					err := writer.WriteElementString("LN", it.Base().LogicalName())
					if err != nil {
						return err
					}
					if _, ok := g.accessRights[it]; !ok {
						count := it.GetAttributeCount()
						buff := make([]int, count)
						for pos := 1; pos <= count; pos++ {
							if g.Version < 3 {
								buff[pos-1] = int(it.Base().GetAccess(pos))
							} else {
								buff[pos-1] = int(it.Base().GetAccess3(pos))
							}
						}
						g.accessRights[it] = buff
					}
					// Add access rights if set.
					if buff, ok := g.accessRights[it]; ok && len(buff) != 0 {
						sb.Reset()
						if g.Version < 3 {
							for _, v := range buff {
								sb.WriteString(strconv.Itoa(v))
							}
							err = writer.WriteElementString("Access", sb.String())
						} else {
							for _, v := range buff {
								// Set highest bit so value is written with two bytes.
								sb.WriteString(fmt.Sprintf("%04X", 0x8000|v))
							}
							err = writer.WriteElementString("Access3", sb.String())
						}
						if err != nil {
							return err
						}
					}
					if _, ok := g.methodAccessRights[it]; !ok {
						count := it.GetMethodCount()
						buff := make([]int, count)
						for pos := 1; pos <= count; pos++ {
							if g.Version < 3 {
								buff[pos-1] = int(it.Base().GetMethodAccess(pos))
							} else {
								buff[pos-1] = int(it.Base().GetMethodAccess3(pos))
							}
						}
						g.methodAccessRights[it] = buff
					}
					if buff, ok := g.methodAccessRights[it]; ok && len(buff) != 0 {
						sb.Reset()
						if g.Version < 3 {
							for _, v := range buff {
								sb.WriteString(strconv.Itoa(v))
							}
							err = writer.WriteElementString("MethodAccess", sb.String())
						} else {
							for _, v := range buff {
								// Set highest bit so value is written with two bytes.
								sb.WriteString(fmt.Sprintf("%04X", 0x8000|v))
							}
							err = writer.WriteElementString("MethodAccess3", sb.String())
						}
						if err != nil {
							return err
						}
					}
					writer.WriteEndElement()
				}
			}
		}
	}
	writer.WriteEndElement()
	err := writer.WriteElementString("ClientSAP", g.ClientSAP)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("ServerSAP", g.ServerSAP)
	if err != nil {
		return err
	}
	writer.WriteStartElement("ApplicationContextName")
	err = writer.WriteElementString("JointIsoCtt", g.ApplicationContextName.JointIsoCtt)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("Country", g.ApplicationContextName.Country)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("CountryName", g.ApplicationContextName.CountryName)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("IdentifiedOrganization", g.ApplicationContextName.IdentifiedOrganization)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("DlmsUA", g.ApplicationContextName.DlmsUA)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("ApplicationContext", g.ApplicationContextName.ApplicationContext)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("ContextId", int(g.ApplicationContextName.ContextId))
	err = writer.WriteEndElement()
	if err != nil {
		return err
	}
	err = writer.WriteStartElement("XDLMSContextInfo")

	err = writer.WriteElementString("Conformance", int(g.XDLMSContextInfo.Conformance))
	if err != nil {
		return err
	}
	err = writer.WriteElementString("MaxReceivePduSize", g.XDLMSContextInfo.MaxReceivePduSize)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("MaxSendPduSize", g.XDLMSContextInfo.MaxSendPduSize)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("DlmsVersionNumber", g.XDLMSContextInfo.DlmsVersionNumber)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("QualityOfService", g.XDLMSContextInfo.QualityOfService)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("CypheringInfo", types.ToHex(g.XDLMSContextInfo.CypheringInfo, false))
	if err != nil {
		return err
	}
	err = writer.WriteEndElement()
	err = writer.WriteStartElement("AuthenticationMechanismName")
	err = writer.WriteElementString("JointIsoCtt", g.AuthenticationMechanismName.JointIsoCtt)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("Country", g.AuthenticationMechanismName.Country)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("CountryName", g.AuthenticationMechanismName.CountryName)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("IdentifiedOrganization", g.AuthenticationMechanismName.IdentifiedOrganization)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("DlmsUA", g.AuthenticationMechanismName.DlmsUA)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("AuthenticationMechanismName", g.AuthenticationMechanismName.AuthenticationMechanismName)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("MechanismId", int(g.AuthenticationMechanismName.MechanismId))
	if err != nil {
		return err
	}
	writer.WriteEndElement()
	ret := types.ToHex(g.Secret, false)
	err = writer.WriteElementString("Secret", ret)
	if err != nil {
		return err
	}
	err = writer.WriteElementString("AssociationStatus", int(g.AssociationStatus))
	if err != nil {
		return err
	}
	if g.SecuritySetupReference == "" {
		err = writer.WriteElementString("SecuritySetupReference", "0.0.0.0.0.0")
		if err != nil {
			return err
		}
	} else {
		err = writer.WriteElementString("SecuritySetupReference", g.SecuritySetupReference)
		if err != nil {
			return err
		}
	}
	// Save users.
	if g.UserList != nil {
		writer.WriteStartElement("Users")
		for k, v := range g.UserList {
			writer.WriteStartElement("User")
			err = writer.WriteElementString("Id", k)
			if err != nil {
				return err
			}
			err = writer.WriteElementString("Name", v)
			if err != nil {
				return err
			}
			writer.WriteEndElement()
		}
		writer.WriteEndElement()
	}
	err = writer.WriteElementString("MultipleAssociationViews", g.MultipleAssociationViews)
	if err != nil {
		return err
	}
	return err
}

// PostLoad returns the handle actions after Load.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSAssociationLogicalName) PostLoad(reader *GXXmlReader) error {
	return nil
}

// GetAttributeAccess returns the default attribute access mode for the selected object.
//
// Parameters:
//
//	target: target object.
//	attributeIndex: Attribute index.
//
// Returns:
//
//	Default access mode.
func (g *GXDLMSAssociationLogicalName) GetAttributeAccess(target IGXDLMSBase, attributeIndex int) int {
	if attributeIndex == 1 {
		return int(enums.AccessModeRead)
	}
	att := target.Base().Attributes.Find(attributeIndex)
	if att != nil {
		return int(att.Access)
	}
	switch target.Base().ObjectType() {
	case enums.ObjectTypeNone:
	case enums.ObjectTypeActionSchedule:
	case enums.ObjectTypeActivityCalendar:
	case enums.ObjectTypeAssociationLogicalName:
		// Association Status
		if attributeIndex == 8 {
			return int(enums.AccessModeRead)
		}
	case enums.ObjectTypeAssociationShortName:
	case enums.ObjectTypeAutoAnswer:
	case enums.ObjectTypeAutoConnect:
	case enums.ObjectTypeClock:
	case enums.ObjectTypeData:
	case enums.ObjectTypeDemandRegister:
	case enums.ObjectTypeMacAddressSetup:
	case enums.ObjectTypeExtendedRegister:
	case enums.ObjectTypeGprsSetup:
	case enums.ObjectTypeIecHdlcSetup:
	case enums.ObjectTypeIecLocalPortSetup:
	case enums.ObjectTypeIecTwistedPairSetup:
	case enums.ObjectTypeIP4Setup:
	case enums.ObjectTypeGSMDiagnostic:
	case enums.ObjectTypeIP6Setup:
	case enums.ObjectTypeMBusSlavePortSetup:
	case enums.ObjectTypeModemConfiguration:
	case enums.ObjectTypePushSetup:
	case enums.ObjectTypePppSetup:
	case enums.ObjectTypeProfileGeneric:
	case enums.ObjectTypeRegister:
	case enums.ObjectTypeRegisterActivation:
	case enums.ObjectTypeRegisterMonitor:
	case enums.ObjectTypeIec8802LlcType1Setup:
	case enums.ObjectTypeIec8802LlcType2Setup:
	case enums.ObjectTypeIec8802LlcType3Setup:
	case enums.ObjectTypeDisconnectControl:
	case enums.ObjectTypeLimiter:
	case enums.ObjectTypeMBusClient:
	case enums.ObjectTypeCompactData:
	case enums.ObjectTypeParameterMonitor:
	case enums.ObjectTypeWirelessModeQchannel:
	case enums.ObjectTypeMBusMasterPortSetup:
	case enums.ObjectTypeLlcSscsSetup:
	case enums.ObjectTypePrimeNbOfdmPlcPhysicalLayerCounters:
	case enums.ObjectTypePrimeNbOfdmPlcMacSetup:
	case enums.ObjectTypePrimeNbOfdmPlcMacFunctionalParameters:
	case enums.ObjectTypePrimeNbOfdmPlcMacCounters:
	case enums.ObjectTypePrimeNbOfdmPlcMacNetworkAdministrationData:
	case enums.ObjectTypePrimeNbOfdmPlcApplicationsIdentification:
	case enums.ObjectTypeRegisterTable:
	case enums.ObjectTypeZigBeeSasStartup:
	case enums.ObjectTypeZigBeeSasJoin:
	case enums.ObjectTypeZigBeeSasApsFragmentation:
	case enums.ObjectTypeZigBeeNetworkControl:
	case enums.ObjectTypeDataProtection:
	case enums.ObjectTypeAccount:
	case enums.ObjectTypeCredit:
	case enums.ObjectTypeCharge:
	case enums.ObjectTypeTokenGateway:
	case enums.ObjectTypeSapAssignment:
	case enums.ObjectTypeImageTransfer:
	case enums.ObjectTypeSchedule:
	case enums.ObjectTypeScriptTable:
	case enums.ObjectTypeSMTPSetup:
	case enums.ObjectTypeSpecialDaysTable:
	case enums.ObjectTypeStatusMapping:
	case enums.ObjectTypeSecuritySetup:
	case enums.ObjectTypeTCPUDPSetup:
	case enums.ObjectTypeUtilityTables:
	case enums.ObjectTypeSFSKPhyMacSetUp:
	case enums.ObjectTypeSFSKActiveInitiator:
	case enums.ObjectTypeSFSKMacSynchronizationTimeouts:
	case enums.ObjectTypeSFSKMacCounters:
	case enums.ObjectTypeIec61334_4_32LlcSetup:
	case enums.ObjectTypeSFSKReportingSystemList:
	case enums.ObjectTypeArbitrator:
	case enums.ObjectTypeG3PlcMacLayerCounters:
	case enums.ObjectTypeG3PlcMacSetup:
	case enums.ObjectTypeG3Plc6LoWPan:
	case enums.ObjectTypeTariffPlan:
	}
	return int(enums.AccessModeReadWrite)
}

// GetValues returns the an array containing the COSEM object's attribute values.
func (g *GXDLMSAssociationLogicalName) GetValues() []any {
	return []any{g.LogicalName, g.ObjectList, []any{g.ClientSAP, g.ServerSAP}, g.ApplicationContextName,
		g.XDLMSContextInfo, g.AuthenticationMechanismName, g.Secret, g.AssociationStatus, g.SecuritySetupReference, g.UserList, g.CurrentUser}
}

// UpdateSecret returns the updates secret.
//
// Parameters:
//
//	client: DLMS client.
//
// Returns:
//
//	Action bytes.
func (g *GXDLMSAssociationLogicalName) UpdateSecret(client IGXDLMSClient) ([][]byte, error) {
	if g.AuthenticationMechanismName.MechanismId == enums.AuthenticationNone {
		return nil, errors.New("Invalid authentication level in MechanismId.")
	}
	if g.AuthenticationMechanismName.MechanismId == enums.AuthenticationHighGMAC {
		return nil, errors.New("HighGMAC secret is updated using Security setup.")
	}
	if g.AuthenticationMechanismName.MechanismId == enums.AuthenticationLow {
		return client.Write(g, 7)
	}
	// Action is used to update High authentication password.
	return client.Method(g, 2, g.Secret, enums.DataTypeOctetString)
}

// AddObject returns the add object to object list.
//
// Parameters:
//
//	client: DLMS client.
//	obj: COSEM object.
//
// Returns:
//
//	Action bytes.
func (g *GXDLMSAssociationLogicalName) AddObject(client IGXDLMSClient, obj IGXDLMSBase) ([][]byte, error) {
	data := types.GXByteBuffer{}
	err := data.SetUint8(uint8(enums.DataTypeStructure))
	if err != nil {
		return nil, err
	}
	err = data.SetUint8(4)
	if err != nil {
		return nil, err
	}
	ln, err := helpers.LogicalNameToBytes(obj.Base().LogicalName())
	if err != nil {
		return nil, err
	}
	err = internal.SetData(nil, &data, enums.DataTypeUint16, int(obj.Base().ObjectType()))
	if err != nil {
		return nil, err
	}
	err = internal.SetData(nil, &data, enums.DataTypeUint8, obj.Base().Version)
	if err != nil {
		return nil, err
	}
	err = internal.SetData(nil, &data, enums.DataTypeOctetString, ln)
	if err != nil {
		return nil, err
	}
	g.GetAccessRights(nil, obj, nil, &data)
	return client.Method(g, 3, data.Array(), enums.DataTypeStructure)
}

// RemoveObject returns the remove object from object list.
//
// Parameters:
//
//	client: DLMS client.
//	obj: COSEM object.
//
// Returns:
//
//	Action bytes.
func (g *GXDLMSAssociationLogicalName) RemoveObject(client IGXDLMSClient, obj IGXDLMSBase) ([][]byte, error) {
	data := types.GXByteBuffer{}
	err := data.SetUint8(uint8(enums.DataTypeStructure))
	if err != nil {
		return nil, err
	}
	err = data.SetUint8(4)
	if err != nil {
		return nil, err
	}
	err = internal.SetData(nil, &data, enums.DataTypeUint16, obj.Base().ObjectType())
	if err != nil {
		return nil, err
	}
	err = internal.SetData(nil, &data, enums.DataTypeUint8, obj.Base().Version)
	if err != nil {
		return nil, err
	}
	ln, err := helpers.LogicalNameToBytes(obj.Base().LogicalName())
	if err != nil {
		return nil, err
	}
	err = internal.SetData(nil, &data, enums.DataTypeOctetString, ln)
	if err != nil {
		return nil, err
	}
	g.GetAccessRights(nil, obj, nil, &data)
	return client.Method(g, 4, data.Array(), enums.DataTypeStructure)
}

// AddUser returns the add user to user list.
//
// Parameters:
//
//	client: DLMS client.
//	id: User ID.
//	name: User name.
//
// Returns:
//
//	Action bytes.
func (g *GXDLMSAssociationLogicalName) AddUser(client IGXDLMSClient, id uint8, name string) ([][]byte, error) {
	data := types.GXByteBuffer{}
	err := data.SetUint8(uint8(enums.DataTypeStructure))
	if err != nil {
		return nil, err
	}
	err = data.SetUint8(2)
	if err != nil {
		return nil, err
	}
	err = internal.SetData(nil, &data, enums.DataTypeUint8, id)
	if err != nil {
		return nil, err
	}
	err = internal.SetData(nil, &data, enums.DataTypeString, name)
	if err != nil {
		return nil, err
	}
	return client.Method(g, 5, data.Array(), enums.DataTypeStructure)
}

// RemoveUser returns the remove user from user list.
//
// Parameters:
//
//	client: DLMS client.
//	id: User ID.
//	name: User name.
//
// Returns:
//
//	Action bytes.
func (g *GXDLMSAssociationLogicalName) RemoveUser(client IGXDLMSClient, id uint8, name string) ([][]byte, error) {
	data := types.GXByteBuffer{}
	err := data.SetUint8(uint8(enums.DataTypeStructure))
	if err != nil {
		return nil, err
	}
	err = data.SetUint8(2)
	if err != nil {
		return nil, err
	}
	err = internal.SetData(nil, &data, enums.DataTypeUint8, id)
	if err != nil {
		return nil, err
	}
	err = internal.SetData(nil, &data, enums.DataTypeString, name)
	if err != nil {
		return nil, err
	}
	return client.Method(g, 6, data.Array(), enums.DataTypeStructure)
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
func (g *GXDLMSAssociationLogicalName) GetDataType(index int) (enums.DataType, error) {
	ret := enums.DataTypeNone
	switch index {
	case 1:
		ret = enums.DataTypeOctetString
	case 2:
		ret = enums.DataTypeArray
	case 3:
		ret = enums.DataTypeStructure
	case 4:
		ret = enums.DataTypeStructure
	case 5:
		ret = enums.DataTypeStructure
	case 6:
		ret = enums.DataTypeStructure
	case 7:
		ret = enums.DataTypeOctetString
	case 8:
		ret = enums.DataTypeEnum
	}
	if g.Version > 0 {
		if index == 9 {
			ret = enums.DataTypeOctetString
		}
	}
	if g.Version > 1 {
		if index == 10 {
			ret = enums.DataTypeArray
		}
		if index == 11 {
			ret = enums.DataTypeStructure
		}
	}
	if ret != enums.DataTypeNone {
		return enums.DataType(ret), nil
	}
	return ret, errors.New("GetDataType failed. Invalid attribute index.")
}

// IsAccessRightSet returns the are access right sets for the given object.
//
// Parameters:
//
//	target: Target object.
func (g *GXDLMSAssociationLogicalName) IsAccessRightSet(target IGXDLMSBase) bool {
	return g.accessRights[target] != nil
}

// GetAccess returns the access mode for given object.
//
// Parameters:
//
//	target: COSEM object.
//	index: Attribute index.
//
// Returns:
//
//	Access mode.
func (g *GXDLMSAssociationLogicalName) GetObjectAccess(target IGXDLMSBase, index int) enums.AccessMode {
	if target == g {
		return g.GetAccess(index)
	}
	if target.Base().ObjectType() == enums.ObjectTypeAssociationLogicalName && target.Base().LogicalName() == "0.0.40.0.0.255" {
		return g.GetAccess(index)
	}
	if len(g.accessRights) == 0 {
		return enums.AccessModeRead | enums.AccessModeWrite
	}
	if tmp, ok := g.accessRights[target]; ok {
		return enums.AccessMode(tmp[index-1])
	}
	return enums.AccessModeNoAccess
}

// SetAccess returns the sets access mode for given object.
//
// Parameters:
//
//	target: COSEM object.
//	index: Attribute index.
//	access: Access mode.
func (g *GXDLMSAssociationLogicalName) SetAccess(target IGXDLMSBase, index int, access enums.AccessMode) {
	if _, ok := g.accessRights[target]; ok {
		g.accessRights[target][index-1] = int(access)
	} else {
		list := g.accessRights[target]
		list = append(list, int(access))
		g.accessRights[target] = list
	}
}

// SetAccessArray returns the sets access mode for given object.
//
// Parameters:
//
//	target: COSEM object.
//	access: Access modes.
func (g *GXDLMSAssociationLogicalName) SetAccessArray(target IGXDLMSBase, access []enums.AccessMode) error {
	count := target.GetAttributeCount()
	if count < len(access) {
		return errors.New("Invalid access buffer.")
	}
	buff := make([]int, count)
	for pos := 0; pos != len(access); pos++ {
		buff[pos] = int(access[pos])
	}
	g.accessRights[target] = buff
	return nil
}

// SetDefaultAccess returns the update default access mode for all objects in the association view.
// Server can use this to set default access mode for all the objects.
//
// Parameters:
//
//	mode: Defaule method access mode.
func (g *GXDLMSAssociationLogicalName) SetDefaultAccess(mode enums.AccessMode) error {
	if g.Version > 2 {
		return errors.New("Use SetDefaultMethodAccess3 to set default method access for logical name association version 3.")
	}
	for _, obj := range g.ObjectList {
		count := obj.GetAttributeCount()
		list := make([]int, count)
		for pos := 0; pos != count; pos++ {
			list[pos] = int(mode)
		}
		g.accessRights[obj] = list
	}
	return nil
}

// SetDefaultAccess3 returns the update default access mode for all objects in the association view.
// Server can use this to set default access mode for all the objects.
//
// Parameters:
//
//	mode: Defaule method access mode.
func (g *GXDLMSAssociationLogicalName) SetDefaultAccess3(mode enums.AccessMode3) error {
	if g.Version < 3 {
		return errors.New("Use SetDefaultMethodAccess to set default method access for logical name association version 3.")
	}
	for _, obj := range g.ObjectList {
		count := obj.GetAttributeCount()
		list := make([]int, count)
		for pos := 0; pos != count; pos++ {
			list[pos] = int(mode)
		}
		g.accessRights[obj] = list
	}
	return nil
}

// SetDefaultMethodAccess returns the update default method access mode for all objects.
// Server can use this to set default access mode for all the objects.
//
// Parameters:
//
//	mode: Defaule method access mode.
func (g *GXDLMSAssociationLogicalName) SetDefaultMethodAccess(mode enums.MethodAccessMode) error {
	if g.Version > 2 {
		return errors.New("Use SetDefaultMethodAccess3 to set default method access for logical name association version 3.")
	}
	for _, obj := range g.ObjectList {
		count := obj.GetMethodCount()
		list := make([]int, count)
		for pos := 0; pos != count; pos++ {
			list[pos] = int(mode)
		}
		g.methodAccessRights[obj] = list
	}
	return nil
}

// SetDefaultMethodAccess3 returns the update default method access mode for all objects.
// Server can use this to set default access mode for all the objects.
//
// Parameters:
//
//	mode: Defaule method access mode.
func (g *GXDLMSAssociationLogicalName) SetDefaultMethodAccess3(mode enums.MethodAccessMode3) error {
	if g.Version < 3 {
		return errors.New("Use SetDefaultMethodAccess to set default method access for logical name association version 1 or 2.")
	}
	for _, obj := range g.ObjectList {
		count := obj.GetMethodCount()
		list := make([]int, count)
		for pos := 0; pos != count; pos++ {
			list[pos] = int(mode)
		}
		g.methodAccessRights[obj] = list
	}
	return nil
}

// GetMethodAccess returns the method access mode for given object.
//
// Parameters:
//
//	target: COSEM object.
//	index: Attribute index.
//
// Returns:
//
//	Method access mode.
func (g *GXDLMSAssociationLogicalName) GetObjectMethodAccess(target IGXDLMSBase, index int) enums.MethodAccessMode {
	if _, ok := g.methodAccessRights[target]; !ok {
		return enums.MethodAccessModeNoAccess
	}

	if target.Base() == g.Base() {
		return g.GetMethodAccess(index)
	}
	if target.Base().ObjectType() == enums.ObjectTypeAssociationLogicalName && target.Base().LogicalName() == "0.0.40.0.0.255" {
		return g.GetMethodAccess(index)
	}
	if index <= len(g.methodAccessRights[target]) {
		return enums.MethodAccessMode(g.methodAccessRights[target][index-1])
	}
	return enums.MethodAccessModeNoAccess
}

// SetMethodAccess returns the sets method access mode for given object.
//
// Parameters:
//
//	target: COSEM object.
//	index: Attribute index.
//	access: Method access mode.
func (g *GXDLMSAssociationLogicalName) SetObjectMethodAccess(target IGXDLMSBase, index int, access enums.MethodAccessMode) {
	if target == nil {
		return
	}
	if g.methodAccessRights == nil {
		g.methodAccessRights = make(map[IGXDLMSBase][]int)
	}
	count := target.GetMethodCount()
	if index < 1 || index > count {
		return
	}
	if list, ok := g.methodAccessRights[target]; ok && len(list) >= count {
		list[index-1] = int(access)
		g.methodAccessRights[target] = list
		return
	}
	list := make([]int, count)
	for pos := range list {
		list[pos] = int(enums.MethodAccessModeAccess)
	}
	list[index-1] = int(access)
	g.methodAccessRights[target] = list
}

// SetMethodAccess returns the sets method access mode for given object.
//
// Parameters:
//
//	target: COSEM object.
//	access: Method access modes.
func (g *GXDLMSAssociationLogicalName) SetMethodAccessArray(target IGXDLMSBase, access []enums.MethodAccessMode) error {
	count := target.GetMethodCount()
	if count < len(access) {
		return errors.New("Invalid access buffer.")
	}
	buff := make([]int, count)
	for pos := 0; pos != len(access); pos++ {
		buff[pos] = int(access[pos])
	}
	g.methodAccessRights[target] = buff
	return nil
}

// String produces logical or Short Name of DLMS object.
func (g *GXDLMSAssociationLogicalName) String() string {
	str := g.Base().String()
	str = str + " " + g.AuthenticationMechanismName.MechanismId.String()
	return str
}

// GetAccess3 returns the access mode for given object.
//
// Parameters:
//
//	target: COSEM object.
//	index: Attribute index.
//
// Returns:
//
//	Access mode.
func (g *GXDLMSAssociationLogicalName) GetObjectAccess3(target IGXDLMSBase, index int) enums.AccessMode3 {
	if target == g {
		return g.GetAccess3(index)
	}
	if target.Base().ObjectType() == enums.ObjectTypeAssociationLogicalName && target.Base().LogicalName() == "0.0.40.0.0.255" {
		return g.GetAccess3(index)
	}
	if len(g.accessRights) == 0 {
		return enums.AccessMode3Read | enums.AccessMode3Write
	}
	if tmp, ok := g.accessRights[target]; ok {
		return enums.AccessMode3(tmp[index-1])
	}
	return enums.AccessMode3NoAccess
}

// SetAccess3 returns the sets access mode for given object.
//
// Parameters:
//
//	target: COSEM object.
//	index: Attribute index.
//	access: Access mode.
func (g *GXDLMSAssociationLogicalName) SetAccess3(target IGXDLMSBase, index int, access enums.AccessMode3) {
	if _, ok := g.accessRights[target]; ok {
		g.accessRights[target][index-1] = int(access)
	} else {
		list := g.accessRights[target]
		list = append(list, int(access))
		g.accessRights[target] = list
	}
}

// SetAccess3 returns the sets access mode for given object.
//
// Parameters:
//
//	target: COSEM object.
//	access: Access modes.
func (g *GXDLMSAssociationLogicalName) SetAccess3Array(target IGXDLMSBase, access []enums.AccessMode3) error {
	count := target.GetAttributeCount()
	if count < len(access) {
		return errors.New("Invalid access buffer.")
	}
	buff := make([]int, count)
	for pos := 0; pos != len(access); pos++ {
		buff[pos] = int(access[pos])
	}
	g.accessRights[target] = buff
	return nil
}

// GetMethodAccess3 returns the method access mode for given object.
//
// Parameters:
//
//	target: COSEM object.
//	index: Attribute index.
//
// Returns:
//
//	Method access mode.
func (g *GXDLMSAssociationLogicalName) GetObjectMethodAccess3(target IGXDLMSBase, index int) enums.MethodAccessMode3 {
	if len(g.methodAccessRights) == 0 {
		return enums.MethodAccessMode3Access
	}
	if target == g {
		return g.GetMethodAccess3(index)
	}
	if target.Base().ObjectType() == enums.ObjectTypeAssociationLogicalName && target.Base().LogicalName() == "0.0.40.0.0.255" {
		return g.GetMethodAccess3(index)
	}
	return enums.MethodAccessMode3(g.methodAccessRights[target][index-1])
}

// SetMethodAccess3 returns the sets method access mode for given object.
//
// Parameters:
//
//	target: COSEM object.
//	index: Attribute index.
//	access: Method access mode.
func (g *GXDLMSAssociationLogicalName) SetObjectMethodAccess3(target IGXDLMSBase, index int, access enums.MethodAccessMode3) error {
	if _, ok := g.methodAccessRights[target]; ok {
		g.methodAccessRights[target][index-1] = int(access)
	} else {
		list := g.methodAccessRights[target]
		list = append(list, int(access))
		g.methodAccessRights[target] = list
	}
	return nil
}

// SetMethodAccess3 returns the sets method access mode for given object.
//
// Parameters:
//
//	target: COSEM object.
//	access: Method access modes.
func (g *GXDLMSAssociationLogicalName) SetMethodAccess3(target IGXDLMSBase, access []enums.MethodAccessMode3) error {
	count := target.GetMethodCount()
	if count < len(access) {
		return errors.New("Invalid access buffer.")
	}
	buff := make([]int, count)
	for pos := 0; pos != len(access); pos++ {
		buff[pos] = int(access[pos])
	}
	g.methodAccessRights[target] = buff
	return nil
}

// Constructor.
// ln: Logical Name of the object.
// sn: Short Name of the object.
func NewGXDLMSAssociationLogicalName(ln string, sn int16) (*GXDLMSAssociationLogicalName, error) {
	err := ValidateLogicalName(ln)
	if err != nil {
		return nil, err
	}
	return &GXDLMSAssociationLogicalName{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeAssociationLogicalName,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
