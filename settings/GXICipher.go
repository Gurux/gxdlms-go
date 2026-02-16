package settings

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
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/types"
)

type GXICipher interface {
	// Used security level.
	Security() enums.Security

	// Set security level.
	SetSecurity(value enums.Security) error

	SetSecurityChangeCheck(value bool) error

	// Used security policy.
	SecurityPolicy() enums.SecurityPolicy

	// Set security policy.
	SetSecurityPolicy(value enums.SecurityPolicy) error

	// Used security suite.
	SecuritySuite() enums.SecuritySuite

	// System title.
	SystemTitle() []byte

	// Recipient system title.
	RecipientSystemTitle() []byte

	// Block cipher key.
	BlockCipherKey() []byte

	// Broadcast block cipher key.
	BroadcastBlockCipherKey() []byte

	// Authentication key.
	AuthenticationKey() []byte

	// Dedicated key.
	DedicatedKey() []byte

	// InvocationCounter returns the invocation counter value.
	InvocationCounter() uint32

	// SetInvocationCounter sets the invocation counter value.
	SetInvocationCounter(value uint32) error

	// Transaction Id.
	TransactionId() []byte

	// Ephemeral key pair.
	EphemeralKeyPair() *types.GXKeyValuePair[*types.GXPublicKey, *types.GXPrivateKey]

	// Client's key agreement key pair.
	KeyAgreementKeyPair() *types.GXKeyValuePair[*types.GXPublicKey, *types.GXPrivateKey]

	// Available certificates.
	Certificates() []types.GXx509Certificate

	// Signing key pair.
	SigningKeyPair() *types.GXKeyValuePair[*types.GXPublicKey, *types.GXPrivateKey]

	// TLS signing key pair.
	TLSKeyPair() *types.GXKeyValuePair[*types.GXPublicKey, *types.GXPrivateKey]

	// Used signing.
	Signing() enums.Signing

	// Used signing and ciphering order.
	SignCipherOrder() enums.SignCipherOrder

	// Are Initiate Request and Response signed.
	SignInitiateRequestResponse() bool

	//Reset returns the reset encrypt settings.
	Reset()

	// SetSystemTitle sets the system title.
	SetSystemTitle(value []byte) error

	// SetSecuritySuite sets the used security suite.
	SetSecuritySuite(value enums.SecuritySuite) error

	// Ephemeral key pair.
	SetEphemeralKeyPair(value *types.GXKeyValuePair[*types.GXPublicKey, *types.GXPrivateKey]) error

	// Client's key agreement key pair.
	SetKeyAgreementKeyPair(value *types.GXKeyValuePair[*types.GXPublicKey, *types.GXPrivateKey]) error

	// Signing key pair.
	SetSigningKeyPair(value *types.GXKeyValuePair[*types.GXPublicKey, *types.GXPrivateKey]) error

	// TLS signing key pair.
	SetTLSKeyPair(value *types.GXKeyValuePair[*types.GXPublicKey, *types.GXPrivateKey]) error

	// SetBlockCipherKey sets the block cipher key.
	SetBlockCipherKey(value []byte) error

	// SetBroadcastBlockCipherKey sets the broadcast block cipher key.
	SetBroadcastBlockCipherKey(value []byte) error

	// SetAuthenticationKey sets the authentication key.
	SetAuthenticationKey(value []byte) error

	// SetDedicatedKey sets the dedicated key.
	SetDedicatedKey(value []byte) error

	// SetRecipientSystemTitle sets the recipient system title.
	SetRecipientSystemTitle(value []byte) error

	// SetSigning sets the used signing.
	SetSigning(value enums.Signing) error
}
