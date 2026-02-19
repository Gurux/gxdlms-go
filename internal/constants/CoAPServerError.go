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

// CoAPServerError defines server error values.
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
