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

// CoAPMethod defines CoAP method values.
type CoAPMethod int

const (
	// CoAPMethodNone defines empty command.
	CoAPMethodNone CoAPMethod = iota
	// CoAPMethodGet defines get command.
	CoAPMethodGet
	// CoAPMethodPost defines post command.
	CoAPMethodPost
	// CoAPMethodPut defines Put command.
	CoAPMethodPut
	// CoAPMethodDelete defines delete command.
	CoAPMethodDelete
	// CoAPMethodFetch defines fetch command.
	CoAPMethodFetch
	// CoAPMethodPatch defines patch command.
	CoAPMethodPatch
	// CoAPMethodIPatch defines IPatch command.
	CoAPMethodIPatch
)
