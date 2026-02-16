package objects

import "github.com/Gurux/gxdlms-go/enums"

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

// IGXDLMSBase corresponds to Gurux.DLMS.Objects.IGXDLMSBase.
type IGXDLMSClient interface {
	Settings() any

	// GetDataType returns the device data type of selected attribute index.
	// Write returns the generates a write message.
	//
	// Parameters:
	//
	//	item: object to write.
	//	index: Attribute index where data is write.
	Write(item IGXDLMSBase, index int) ([][]byte, error)

	// Method returns the generate Method (Action) request.
	//
	// Parameters:
	//
	//	item: object to write.
	//	index: Attribute index where data is write.
	Method(item IGXDLMSBase, index int, data any, dataType enums.DataType) ([][]byte, error)
}
