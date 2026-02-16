package constants

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
	"fmt"
	"strings"

	"github.com/Gurux/gxcommon-go"
)

// CoAP server error codes.
type CoAPServerError int

const (
	// CoAPServerErrorInternal defines that the // Internal Server error.
	CoAPServerErrorInternal CoAPServerError = iota
	// CoAPServerErrorNotImplemented defines that the // Not implemented.
	CoAPServerErrorNotImplemented
	// CoAPServerErrorBadGateway defines that the // Bad gateway.
	CoAPServerErrorBadGateway
	// CoAPServerErrorServiceUnavailable defines that the // Service unavailable.
	CoAPServerErrorServiceUnavailable
	// CoAPServerErrorGatewayTimeout defines that the // Gateway timeout.
	CoAPServerErrorGatewayTimeout
	// CoAPServerErrorProxyingNotSupported defines that the // Proxying not supported.
	CoAPServerErrorProxyingNotSupported
)

// CoAPServerErrorParse converts the given string into a CoAPServerError value.
//
// It returns the corresponding CoAPServerError constant if the string matches
// a known level name, or an error if the input is invalid.
func CoAPServerErrorParse(value string) (CoAPServerError, error) {
	var ret CoAPServerError
	var err error
	switch strings.ToUpper(value) {
	case "INTERNAL":
		ret = CoAPServerErrorInternal
	case "NOTIMPLEMENTED":
		ret = CoAPServerErrorNotImplemented
	case "BADGATEWAY":
		ret = CoAPServerErrorBadGateway
	case "SERVICEUNAVAILABLE":
		ret = CoAPServerErrorServiceUnavailable
	case "GATEWAYTIMEOUT":
		ret = CoAPServerErrorGatewayTimeout
	case "PROXYINGNOTSUPPORTED":
		ret = CoAPServerErrorProxyingNotSupported
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the CoAPServerError.
// It satisfies fmt.Stringer.
func (g CoAPServerError) String() string {
	var ret string
	switch g {
	case CoAPServerErrorInternal:
		ret = "INTERNAL"
	case CoAPServerErrorNotImplemented:
		ret = "NOTIMPLEMENTED"
	case CoAPServerErrorBadGateway:
		ret = "BADGATEWAY"
	case CoAPServerErrorServiceUnavailable:
		ret = "SERVICEUNAVAILABLE"
	case CoAPServerErrorGatewayTimeout:
		ret = "GATEWAYTIMEOUT"
	case CoAPServerErrorProxyingNotSupported:
		ret = "PROXYINGNOTSUPPORTED"
	}
	return ret
}
