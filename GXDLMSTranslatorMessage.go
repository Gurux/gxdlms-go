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
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/types"
)

// XML translator message detailed data.
type GXDLMSTranslatorMessage struct {
	// Is more data available. Return None if more data is not available or Frame or Block type.
	moreData enums.RequestTypes

	// Occurred exception.
	exception error

	// Message to convert to XML.
	Message types.GXByteBuffer

	// Converted XML.
	Xml string

	// Executed Command.
	Command enums.Command

	// System title from AARQ or AARE messages.
	SystemTitle []byte

	// Dedicated key from AARQ messages.
	DedicatedKey []byte

	// Interface type.
	InterfaceType enums.InterfaceType

	// Client address.
	SourceAddress int

	// Server address.
	TargetAddress int
}
