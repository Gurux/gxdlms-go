package secure

// --------------------------------------------------------------------------
//
//	Gurux Ltd
//
// Filename:        $HeadURL$
//
// Version:         $Revision$,
//
//	$Date$
//	$Author$
//
// # Copyright (c) Gurux Ltd
//
// ---------------------------------------------------------------------------
//
//	DESCRIPTION
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
// ---------------------------------------------------------------------------

import (
	"fmt"

	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// Gurux DLMS/COSEM Transport security (Ciphering) settings.
type GXCiphering struct {
	// Authentication key.
	authenticationKey []byte

	// System title.
	systemTitle []byte

	// Server System title.
	serverSystemTitle []byte

	// Block ciphering key.
	blockCipherKey []byte

	// Broadcast block ciphering key.
	broadcastBlockCipherKey []byte

	// Dedicated key.
	dedicatedKey []byte

	// Transaction Id.
	transactionId []byte

	// Used security level.
	security enums.Security

	// Security level can't be changed during the connection.
	securityChangeCheck bool

	// Used security policy.
	securityPolicy enums.SecurityPolicy

	// Used security suite.
	securitySuite enums.SecuritySuite

	// Invocation counter for sending.
	invocationCounter uint32

	// Available certificates.
	sertificates []types.GXx509Certificate

	// Public/private key signing key pair.
	signingKeyPair *types.GXKeyValuePair[*types.GXPublicKey, *types.GXPrivateKey]

	// Public/private key TLS key pair.
	tlsKeyPair *types.GXKeyValuePair[*types.GXPublicKey, *types.GXPrivateKey]

	// Ephemeral key pair.
	ephemeralKeyPair *types.GXKeyValuePair[*types.GXPublicKey, *types.GXPrivateKey]

	// Public/private key key agreement key pair.
	keyAgreementKeyPair *types.GXKeyValuePair[*types.GXPublicKey, *types.GXPrivateKey]

	// Ephemeral private key of the client.
	ClientEphemeralPrivateKey *types.GXPkcs8

	// Ephemeral private key of the server.
	ServerEphemeralPrivateKey *types.GXPkcs8

	// Used signing.
	signing enums.Signing

	// Used signing and ciphering order.
	signCipherOrder enums.SignCipherOrder

	// Are InitiateRequest and Response signed.
	signInitiateRequestResponse bool
}

// Signing returns the used signing.
func (g *GXCiphering) Signing() enums.Signing {
	return g.signing
}

// Signing returns the used signing.
func (g *GXCiphering) SetSigning(value enums.Signing) error {
	g.signing = value
	return nil
}

// SecurityPolicy returns the used security policy.
func (g *GXCiphering) SecurityPolicy() enums.SecurityPolicy {
	return g.securityPolicy
}

// SecurityPolicy returns the used security policy.
func (g *GXCiphering) SetSecurityPolicy(value enums.SecurityPolicy) error {
	g.securityPolicy = value
	return nil
}

// SecuritySuite returns the used security suite.
func (g *GXCiphering) SecuritySuite() enums.SecuritySuite {
	return g.securitySuite
}

// SecuritySuite returns the used security suite.
func (g *GXCiphering) SetSecuritySuite(value enums.SecuritySuite) error {
	g.securitySuite = value
	return nil
}

// TransactionId returns the transaction Id.
func (g *GXCiphering) TransactionId() []byte {
	return g.transactionId
}

// SetTransactionId sets the transaction Id.
func (g *GXCiphering) SetTransactionId(value []byte) error {
	if len(value) != 8 && len(value) != 0 {
		return fmt.Errorf("Invalid Transaction Id.")
	}
	g.transactionId = value
	return nil
}

// SystemTitle returns the the SystemTitle is a 8 bytes (64 bit) value that identifies a partner of the communication.
//
//	First 3 bytes contains the three letters manufacturer ID.
//	The remainder of the system title holds for example a serial number.
func (g *GXCiphering) SystemTitle() []byte {
	return g.systemTitle
}

// SetSystemTitle sets the the SystemTitle is a 8 bytes (64 bit) value that identifies a partner of the communication.
//
//	First 3 bytes contains the three letters manufacturer ID.
//	The remainder of the system title holds for example a serial number.
func (g *GXCiphering) SetSystemTitle(value []byte) error {
	if len(value) == 0 {
		value = nil
	}
	if len(value) != 8 {
		return fmt.Errorf("Invalid System Title.")
	}
	g.systemTitle = value
	return nil
}

// RecipientSystemTitle returns the recipient system Title.
func (g *GXCiphering) RecipientSystemTitle() []byte {
	return g.serverSystemTitle
}

// SetRecipientSystemTitle sets the recipient system Title.
func (g *GXCiphering) SetRecipientSystemTitle(value []byte) error {
	if len(value) == 0 {
		value = nil
	}
	if len(value) != 8 && len(value) != 0 {
		return fmt.Errorf("Invalid System Title.")
	}
	g.serverSystemTitle = value
	return nil
}

// BlockCipherKey returns the each block is ciphered with this key.
func (g *GXCiphering) BlockCipherKey() []byte {
	return g.blockCipherKey
}

// SetBlockCipherKey sets the each block is ciphered with this key.
func (g *GXCiphering) SetBlockCipherKey(value []byte) error {
	if len(value) == 0 {
		value = nil
	}
	if len(value) != 16 && len(value) != 32 {
		return fmt.Errorf("Invalid block cipher key. Block cipher key size is 16 or 32 bytes")
	}
	g.blockCipherKey = value
	return nil
}

// BroadcastBlockCipherKey returns the each broadcast block is ciphered with this key.
func (g *GXCiphering) BroadcastBlockCipherKey() []byte {
	return g.broadcastBlockCipherKey
}

// SetBroadcastBlockCipherKey sets the each broadcast block is ciphered with this key.
func (g *GXCiphering) SetBroadcastBlockCipherKey(value []byte) error {
	if len(value) == 0 {
		value = nil
	}
	if len(value) != 16 && len(value) != 32 {
		return fmt.Errorf("Invalid block cipher key. Block cipher key size is 16 or 32 bytes")
	}
	g.broadcastBlockCipherKey = value
	return nil
}

// AuthenticationKey returns the authentication Key is 16 bytes value.
func (g *GXCiphering) AuthenticationKey() []byte {
	return g.authenticationKey
}

// SetAuthenticationKey sets the authentication Key is 16 bytes value.
func (g *GXCiphering) SetAuthenticationKey(value []byte) error {
	if len(value) == 0 {
		value = nil
	}
	if len(value) != 16 && len(value) != 32 {
		return fmt.Errorf("Invalid authentication key. Authentication key size is 16 or 32 bytes.")
	}
	g.authenticationKey = value
	return nil
}

// DedicatedKey returns the dedicated Key is 16 bytes value.
func (g *GXCiphering) DedicatedKey() []byte {
	return g.dedicatedKey
}

// SetDedicatedKey sets the dedicated Key is 16 bytes value.
func (g *GXCiphering) SetDedicatedKey(value []byte) error {
	if len(value) == 0 {
		value = nil
	}
	if len(value) != 16 && len(value) != 32 {
		return fmt.Errorf("Invalid DedicatedKey Key.")
	}
	g.dedicatedKey = value
	return nil
}

// Encrypt returns the encrypt data.
func (g *GXCiphering) Encrypt(p *settings.AesGcmParameter, data []byte) ([]byte, error) {
	if p.Security() != enums.SecurityNone {
		return settings.EncryptAesGcm(p, data)
	}
	return data, nil
}

// Decrypt returns the decrypt data.
func (g *GXCiphering) Decrypt(p *settings.AesGcmParameter, data *types.GXByteBuffer) ([]byte, error) {
	var tmp, err = settings.DecryptAesGcm(p, data)
	if err != nil {
		return nil, err
	}
	data.Clear()
	data.Set(tmp)
	return tmp, nil
}

func (g *GXCiphering) CopyTo(target GXCiphering) {
	target.security = g.security
	target.securitySuite = g.securitySuite
	target.signing = g.signing
	target.invocationCounter = g.invocationCounter
	target.systemTitle = g.systemTitle
	target.blockCipherKey = g.blockCipherKey
	target.authenticationKey = g.authenticationKey
	target.transactionId = g.transactionId
}

// Reset returns the reset encrypt settings.
func (g *GXCiphering) Reset() {
	g.signing = enums.SigningNone
	g.security = enums.SecurityNone
	g.invocationCounter = 0
}

// GenerateGmacPassword returns the generate GMAC password from given challenge.challenge:
func (g *GXCiphering) GenerateGmacPassword(challenge []byte) ([]byte, error) {
	p := settings.NewAesGcmParameter(
		0x10,
		nil,
		enums.SecurityAuthentication,
		g.securitySuite,
		uint64(g.invocationCounter),
		g.systemTitle,
		g.blockCipherKey,
		g.authenticationKey,
	)
	p.Type = settings.CountTypeTag

	data, err := settings.EncryptAesGcm(p, challenge)
	if err != nil {
		return nil, err
	}
	bb := types.NewGXByteBufferWithCapacity(1 + 4 + len(data))
	if err = bb.SetUint8(0x10); err != nil {
		return nil, err
	}
	if err = bb.SetUint32(g.invocationCounter); err != nil {
		return nil, err
	}
	if err = bb.Set(data); err != nil {
		return nil, err
	}
	return bb.Array(), nil
}

// Available certificates.
func (g *GXCiphering) Certificates() []types.GXx509Certificate {
	return g.sertificates
}

// EphemeralKeyPair returns the ephemeral key pair.
func (g *GXCiphering) EphemeralKeyPair() *types.GXKeyValuePair[*types.GXPublicKey, *types.GXPrivateKey] {
	return g.ephemeralKeyPair
}

// Invocation counter for sending.
func (g *GXCiphering) InvocationCounter() uint32 {
	return g.invocationCounter
}

// Public/private key key agreement key pair.
func (g *GXCiphering) KeyAgreementKeyPair() *types.GXKeyValuePair[*types.GXPublicKey, *types.GXPrivateKey] {
	return g.keyAgreementKeyPair
}

// Used security level.
func (g *GXCiphering) Security() enums.Security {
	return g.security
}

// Ephemeral key pair.
func (g *GXCiphering) SetEphemeralKeyPair(value *types.GXKeyValuePair[*types.GXPublicKey, *types.GXPrivateKey]) error {
	g.ephemeralKeyPair = value
	return nil
}

// SetKeyAgreementKeyPair sets the public/private key key agreement key pair.
func (g *GXCiphering) SetKeyAgreementKeyPair(value *types.GXKeyValuePair[*types.GXPublicKey, *types.GXPrivateKey]) error {
	g.keyAgreementKeyPair = value
	return nil
}

// SetSecurity sets the used security level.
func (g *GXCiphering) SetSecurity(value enums.Security) error {
	g.security = value
	return nil
}

// Signing key pair.
func (g *GXCiphering) SetSigningKeyPair(value *types.GXKeyValuePair[*types.GXPublicKey, *types.GXPrivateKey]) error {
	g.signingKeyPair = value
	return nil
}

// SetTLSKeyPair sets the TLS signing key pair.
func (g *GXCiphering) SetTLSKeyPair(value *types.GXKeyValuePair[*types.GXPublicKey, *types.GXPrivateKey]) error {
	g.tlsKeyPair = value
	return nil
}

// SetSecurity sets the used security level.
func (g *GXCiphering) SecurityChangeCheck() bool {
	return g.securityChangeCheck
}

// SetSecurity sets the used security level.
func (g *GXCiphering) SetSecurityChangeCheck(value bool) error {
	g.securityChangeCheck = value
	return nil
}

// SetInvocationCounter sets the invocation counter value.
func (g *GXCiphering) SetInvocationCounter(value uint32) error {
	g.invocationCounter = value
	return nil
}

// Used signing and ciphering order.
func (g *GXCiphering) SignCipherOrder() enums.SignCipherOrder {
	return g.signCipherOrder
}

// Are InitiateRequest and Response signed.
func (g *GXCiphering) SignInitiateRequestResponse() bool {
	return g.signInitiateRequestResponse
}

// Public/private key signing key pair.
func (g *GXCiphering) SigningKeyPair() *types.GXKeyValuePair[*types.GXPublicKey, *types.GXPrivateKey] {
	return g.signingKeyPair
}

// Public/private key TLS key pair.
func (g *GXCiphering) TLSKeyPair() *types.GXKeyValuePair[*types.GXPublicKey, *types.GXPrivateKey] {
	return g.tlsKeyPair
}
