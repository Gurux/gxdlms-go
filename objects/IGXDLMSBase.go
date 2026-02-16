package objects

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
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/settings"
)

// IGXDLMSBase corresponds to Gurux.DLMS.Objects.IGXDLMSBase.
type IGXDLMSBase interface {
	// Base returns the base information from the COSEM object.
	Base() *GXDLMSObject

	// GetValues returns the returns attributes as an array.
	//
	//	Returns:
	//	    Collection of COSEM attribute values.
	GetValues() []any

	// GetAttributeIndexToRead returns collection of attributes to read.
	// If attribute is static and already read or device returned HW error it is not returned.
	// If all == true, all items are returned even if read already.
	GetAttributeIndexToRead(all bool) []int

	// GetAttributeCount returns amount of attributes.
	GetAttributeCount() int

	// GetMethodCount returns amount of methods.
	GetMethodCount() int

	// GetDataType returns data type of selected attribute index.
	GetDataType(index int) (enums.DataType, error)

	// GetUIDataType returns UI data type of selected attribute index.
	GetUIDataType(index int) enums.DataType

	// GetNames returns names of attribute indexes.
	GetNames() []string

	// GetMethodNames returns names of method indexes.
	GetMethodNames() []string

	// GetValue returns value of given attribute.
	// When raw parameter is not used example register multiplies value by scalar.
	GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error)

	// SetValue sets value of given attribute.
	SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error

	// Invoke invokes method.
	Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error)

	// Load loads object content from XML.
	Load(reader *GXXmlReader) error

	// Save saves object content to XML.
	Save(writer *GXXmlWriter) error

	// PostLoad handles actions after Load.
	PostLoad(reader *GXXmlReader) error
}
