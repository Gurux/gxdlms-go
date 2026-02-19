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

// CoAPClientError defines CoAP client error codes.
type CoAPClientError int

const (
	// CoAPClientErrorBadRequest defines Bad Request.
	CoAPClientErrorBadRequest CoAPClientError = iota
	// CoAPClientErrorUnauthorized defines Unauthorized.
	CoAPClientErrorUnauthorized
	// CoAPClientErrorBadOption defines Bad Option.
	CoAPClientErrorBadOption
	// CoAPClientErrorForbidden defines Forbidden.
	CoAPClientErrorForbidden
	// CoAPClientErrorNotFound defines  Not Found.
	CoAPClientErrorNotFound
	// CoAPClientErrorMethodNotAllowed defines Method Not Allowed.
	CoAPClientErrorMethodNotAllowed
	// CoAPClientErrorNotAcceptable defines Not Acceptable.
	CoAPClientErrorNotAcceptable
	// CoAPClientErrorRequestEntityIncomplete defines Request Entity Incomplete.
	CoAPClientErrorRequestEntityIncomplete CoAPClientError = 8
	// CoAPClientErrorConflict defines Conflict.
	CoAPClientErrorConflict CoAPClientError = 9
	// CoAPClientErrorPreconditionFailed defines Precondition Failed.
	CoAPClientErrorPreconditionFailed CoAPClientError = 12
	// CoAPClientErrorRequestEntityTooLarge defines Request Entity Too Large.
	CoAPClientErrorRequestEntityTooLarge CoAPClientError = 13
	// CoAPClientErrorUnsupportedContentFormat defines Unsupported Content-Format.
	CoAPClientErrorUnsupportedContentFormat CoAPClientError = 15
)
