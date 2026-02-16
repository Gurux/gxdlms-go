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
	"github.com/Gurux/gxdlms-go/internal/constants"
)

// CoAP settings contains CoAP settings.
type GXCoAPSettings struct {
	// CoAP block number.
	blockNumber uint8

	// CoAP version.
	Version uint8

	// CoAP type.
	Type constants.CoAPType

	// CoAP class code.
	Class constants.CoAPClass

	// CoAP Method.
	Method constants.CoAPMethod

	// CoAP Success.
	Success constants.CoAPSuccess

	// CoAP client error.
	ClientError constants.CoAPClientError

	// CoAP server error.
	ServerError constants.CoAPServerError

	// CoAP signaling.
	Signaling constants.CoAPSignaling

	// CoAP message Id.
	MessageId uint16

	// CoAP token.
	Token uint64

	// Uri host.
	Host string

	// Uri Path.
	Path string

	// Uri port.
	Port uint16

	// If none match.
	IfNoneMatch enums.CoAPContentType

	// Content format.
	ContentFormat enums.CoAPContentType

	// Max age.
	MaxAge uint16

	// Unknown options.
	Options map[uint16]any
}

// Reset returns the reset all values.
func (g *GXCoAPSettings) Reset() {
	g.Version = 0
	g.Type = constants.CoAPTypeConfirmable
	g.Class = constants.CoAPClassMethod
	g.Method = constants.CoAPMethodNone
	g.Success = constants.CoAPSuccessNone
	g.ClientError = constants.CoAPClientErrorBadRequest
	g.ServerError = constants.CoAPServerErrorInternal
	g.Signaling = constants.CoAPSignalingUnassigned
	g.MessageId = 0
	g.Token = 0
	g.Host = ""
	g.Path = ""
	g.Port = 0
	g.IfNoneMatch = enums.CoAPContentTypeNone
	g.ContentFormat = enums.CoAPContentTypeNone
	g.MaxAge = 0
	g.blockNumber = 0
	g.Options = make(map[uint16]any)
}
