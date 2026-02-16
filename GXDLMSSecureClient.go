package dlms

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
	"errors"

	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

type GXDLMSSecureClient struct {
	GXDLMSClient
}

// securitySuite returns the used security level.
func (g *GXDLMSSecureClient) Security() enums.Security {
	return g.GXDLMSClient.settings.Cipher.Security()
}

// setSecuritySuite sets the used security level.
func (g *GXDLMSSecureClient) SetSecurity(value enums.Security) error {
	return g.GXDLMSClient.settings.Cipher.SetSecurity(value)
}

// securitySuite returns the used security suite.
func (g *GXDLMSSecureClient) SecuritySuite() enums.SecuritySuite {
	return g.GXDLMSClient.settings.Cipher.SecuritySuite()
}

// setSecuritySuite sets the used security suite.
func (g *GXDLMSSecureClient) SetSecuritySuite(value enums.SecuritySuite) error {
	return g.GXDLMSClient.settings.Cipher.SetSecuritySuite(value)
}

// Ciphering settings.
func (g *GXDLMSSecureClient) Ciphering() settings.GXICipher {
	return g.GXDLMSClient.settings.Cipher
}

// ServerSystemTitle returns the server system title.
// Server system title is optional and it's used when Pre-established Application Associations is used.
func (g *GXDLMSSecureClient) ServerSystemTitle() []byte {
	return g.GXDLMSClient.settings.PreEstablishedSystemTitle
}

// SetServerSystemTitle sets the server system title.
// Server system title is optional and it's used when Pre-established Application Associations is used.
func (g *GXDLMSSecureClient) SetServerSystemTitle(value []byte) error {
	if value != nil && len(value) == 0 {
		value = nil
	}
	if value != nil && len(value) != 8 {
		return errors.New("Invalid System Title.")
	}
	g.GXDLMSClient.settings.PreEstablishedSystemTitle = value
	return nil
}

// ClientPublicKeyCertificate returns the optional ECDSA public key certificate that is send in part of AARE.
func (g *GXDLMSSecureClient) ClientPublicKeyCertificate() *types.GXx509Certificate {
	return g.GXDLMSClient.settings.ClientPublicKeyCertificate
}

// SetClientPublicKeyCertificate sets the optional ECDSA public key certificate that is send in part of AARE.
func (g *GXDLMSSecureClient) SetClientPublicKeyCertificate(value *types.GXx509Certificate) error {
	g.GXDLMSClient.settings.ClientPublicKeyCertificate = value
	return nil
}

// ServerPublicKeyCertificate returns the optional ECDSA public key certificate that is send in part of AARE.
func (g *GXDLMSSecureClient) ServerPublicKeyCertificate() *types.GXx509Certificate {
	return g.GXDLMSClient.settings.ServerPublicKeyCertificate
}

// SetServerPublicKeyCertificate sets the optional ECDSA public key certificate that is send in part of AARE.
func (g *GXDLMSSecureClient) SetServerPublicKeyCertificate(value *types.GXx509Certificate) error {
	g.GXDLMSClient.settings.ServerPublicKeyCertificate = value
	return nil
}

// Encrypt returns the encrypt data using AES RFC3394 key-wrapping.
//
// Parameters:
//
//	kek: Key Encrypting Key, also known as Master key.
//	data: Data to encrypt.
//
// Returns:
//
//	Encrypt data.
func (g *GXDLMSSecureClient) Encrypt(kek []byte, data []byte) ([]byte, error) {
	return internal.Encrypt(kek, data)
}

// Decrypt returns the decrypt data using AES RFC3394 key-wrapping.
//
// Parameters:
//
//	kek: Key Encrypting Key, also known as Master key.
//	data: Data to decrypt.
//
// Returns:
//
//	Decrypted data.
func (g *GXDLMSSecureClient) Decrypt(kek []byte, input []byte) ([]byte, error) {
	return internal.Decrypt(kek, input)
}

func NewGXDLMSSecureClient(useLogicalNameReferencing bool, clientAddress int, serverAddress int, authentication enums.Authentication,
	password []byte, interfaceType enums.InterfaceType) (*GXDLMSSecureClient, error) {
	ret := &GXDLMSSecureClient{}

	cl, err := NewGXDLMSClient(useLogicalNameReferencing, clientAddress, serverAddress, authentication, password, interfaceType)
	if err != nil {
		return nil, err
	}
	ret.GXDLMSClient = *cl
	return ret, nil
}
