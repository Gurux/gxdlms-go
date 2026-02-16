package settings

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
	"strings"

	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal/buffer"
)

type AesGcmParameter struct {
	tag byte
	// Enumerated security level.
	security enums.Security

	systemTitle []byte

	blockCipherKey []byte

	authenticationKey []byte

	recipientSystemTitle []byte

	InvocationCounter uint64

	// xml  This is used only on xml parser.
	Xml *GXDLMSTranslatorStructure

	Tag uint8

	// Used transaction ID.
	TransactionId uint64

	Settings *GXDLMSSettings
	// Is data send as a broadcast or unicast.
	Broacast bool

	// V.44 Compression is used.
	Compression bool

	Type CountType

	CountTag []byte

	// Date time.
	DateTime []byte

	// Other information.
	OtherInformation []byte

	// Key parameters.
	KeyParameters int

	// Key ciphered datg.
	KeyCipheredData []byte

	// Ciphered content.
	CipheredContent []byte

	// Signature.
	Signature []byte

	// Used security suite.
	SecuritySuite enums.SecuritySuite

	// System title is not send on pre-established connections.
	IgnoreSystemTitle bool
}

// Security returns the enumerated security level.
func (g *AesGcmParameter) Security() enums.Security {
	return g.security
}

// SetSecurity sets the enumerated security level.
func (g *AesGcmParameter) SetSetSecurity(value enums.Security) error {
	if value > enums.SecurityAuthenticationEncryption {
		value = enums.SecurityAuthenticationEncryption
	}
	g.security = value
	return nil
}

// SystemTitle returns the system title.
func (g *AesGcmParameter) SystemTitle() []byte {
	return g.systemTitle
}

// SetSystemTitle sets the system title.
func (g *AesGcmParameter) SetSystemTitle(value []byte) error {
	g.systemTitle = value
	return nil
}

// BlockCipherKey returns the block cipher key.
func (g *AesGcmParameter) BlockCipherKey() []byte {
	return g.blockCipherKey
}

// SetBlockCipherKey sets the block cipher key.
func (g *AesGcmParameter) SetBlockCipherKey(value []byte) error {
	g.blockCipherKey = value
	return nil
}

// AuthenticationKey returns the authentication key.
func (g *AesGcmParameter) AuthenticationKey() []byte {
	return g.authenticationKey
}

// SetAuthenticationKey sets the authentication key.
func (g *AesGcmParameter) SetAuthenticationKey(value []byte) error {
	g.authenticationKey = value
	return nil
}

// RecipientSystemTitle returns the recipient system title.
func (g *AesGcmParameter) RecipientSystemTitle() []byte {
	return g.recipientSystemTitle
}

// SetRecipientSystemTitle sets the recipient system title.
func (g *AesGcmParameter) SetRecipientSystemTitle(value []byte) error {
	if value != nil {
		if len(value) == 0 {
			value = nil
		} else if len(value) != 8 {
			return fmt.Errorf("Invalid recipient system title. Recipient system title size is 8 bytes.")
		}
	}
	g.recipientSystemTitle = value
	return nil
}

func NewAesGcmParameter(
	tag byte,
	settings *GXDLMSSettings,
	security enums.Security,
	securitySuite enums.SecuritySuite,
	invocationCounter uint64,
	systemTitle []byte,
	blockCipherKey []byte,
	authenticationKey []byte,
) *AesGcmParameter {
	return &AesGcmParameter{
		tag:               tag,
		Settings:          settings,
		security:          security,
		InvocationCounter: invocationCounter,
		SecuritySuite:     securitySuite,
		systemTitle:       systemTitle,
		blockCipherKey:    blockCipherKey,
		authenticationKey: authenticationKey,
	}
}

func NewAesGcmParameter2(
	tag byte,
	settings *GXDLMSSettings,
	security enums.Security,
	securitySuite enums.SecuritySuite,
	invocationCounter uint64,
	kdf []byte,
	authenticationKey []byte,
	originatorSystemTitle []byte,
	recipientSystemTitle []byte,
	dateTime []byte,
	otherInformation []byte,
) *AesGcmParameter {
	return &AesGcmParameter{
		tag:                  tag,
		Settings:             settings,
		security:             security,
		InvocationCounter:    invocationCounter,
		SecuritySuite:        securitySuite,
		blockCipherKey:       kdf,
		authenticationKey:    authenticationKey,
		systemTitle:          originatorSystemTitle,
		recipientSystemTitle: recipientSystemTitle,
		Type:                 CountTypePacket,
		DateTime:             dateTime,
		OtherInformation:     otherInformation,
	}
}

func NewAesGcmParameter3(
	settings *GXDLMSSettings,
	systemTitle []byte,
	blockCipherKey []byte,
	authenticationKey []byte) *AesGcmParameter {
	return &AesGcmParameter{
		security:          settings.Cipher.Security(),
		Settings:          settings,
		systemTitle:       systemTitle,
		blockCipherKey:    blockCipherKey,
		authenticationKey: authenticationKey,
		Type:              CountTypePacket,
		SecuritySuite:     settings.Cipher.SecuritySuite(),
	}
}

// String implements the fmt.Stringer interface.
func (g *AesGcmParameter) String() string {
	var sb strings.Builder
	sb.WriteString("Security: ")
	sb.WriteString(g.Security().String())
	sb.WriteString(" Invocation Counter: ")
	sb.WriteString(fmt.Sprintf("%d", g.InvocationCounter))
	sb.WriteString(" SystemTitle: ")
	sb.WriteString(buffer.ToHex(g.systemTitle, true))
	sb.WriteString(" AuthenticationKey: ")
	sb.WriteString(buffer.ToHex(g.authenticationKey, true))
	sb.WriteString(" BlockCipherKey: ")
	sb.WriteString(buffer.ToHex(g.blockCipherKey, true))
	return sb.String()
}
