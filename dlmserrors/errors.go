package dlmserrors

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
	"errors"
)

// ErrInvalidGloCommand is returned when an invalid global ciphering command is received.
var ErrInvalidGloCommand = errors.New("invalid glo command")

// ErrUnknownPduTag is returned when a PDU tag is unknown or unsupported.
var ErrUnknownPduTag = errors.New("invalid pdu tag")

// ErrPduTooSmall is returned when a received PDU does not contain enough bytes.
var ErrPduTooSmall = errors.New("PDU is too small")

// ErrInvalidApplicationContextName is returned when the application context name is invalid.
var ErrInvalidApplicationContextName = errors.New("Invalid application context name")

// ErrInvalidDLMSVersionNumber is returned when the DLMS version number is invalid.
var ErrInvalidDLMSVersionNumber = errors.New("invalid DLMS version number")

// ErrInvalidVAA is returned when the VAA value is invalid.
var ErrInvalidVAA = errors.New("invalid VAA")

// ErrInvalidAttributeIndex is returned when an attribute index is invalid.
var ErrInvalidAttributeIndex = errors.New("invalid attribute index")

// ErrDataTooShort is returned when a message cannot be parsed because it does not contain enough data.
var ErrDataTooShort = errors.New("data too short")
