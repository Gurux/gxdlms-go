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
	"strconv"

	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

type gxDLMSXmlSettings struct {
	// Are numeric values shows as hex.
	showNumericsAsHex bool

	outputType enums.TranslatorOutputType

	Result enums.AssociationResult

	Diagnostic enums.SourceDiagnostic

	ServiceError int

	Reason uint8

	Command enums.Command

	Count int

	RequestType uint8

	// GW network ID.
	GwCommand enums.Command

	// GW network ID.
	NetworkId uint8

	// GW physical device address.
	PhysicalDeviceAddress []byte

	AttributeDescriptor types.GXByteBuffer

	Data types.GXByteBuffer

	Settings settings.GXDLMSSettings
	Tags     map[string]int

	Time types.GXDateTime

	// Is xml used as a reply template.
	Template bool

	ShowStringAsHex bool
}

func (g *gxDLMSXmlSettings) OutputType() enums.TranslatorOutputType {
	return g.outputType
}

func (g *gxDLMSXmlSettings) ParseInt(value string) (int, error) {
	if g.showNumericsAsHex {
		v, err := strconv.ParseInt(value, 16, 32)
		if err != nil {
			return 0, err
		}
		return int(v), nil
	}
	return strconv.Atoi(value)
}

func (g *gxDLMSXmlSettings) ParseShort(value string) (int16, error) {
	if g.showNumericsAsHex {
		v, err := strconv.ParseInt(value, 16, 16)
		return int16(v), err
	}
	v, err := strconv.Atoi(value)
	return int16(v), err
}

func (g *gxDLMSXmlSettings) ParseLong(value string) (int64, error) {
	if g.showNumericsAsHex {
		v, err := strconv.ParseInt(value, 16, 64)
		return v, err
	}
	return strconv.ParseInt(value, 10, 64)
}

func (g *gxDLMSXmlSettings) ParseULong(value string) (uint64, error) {
	if g.showNumericsAsHex {
		v, err := strconv.ParseUint(value, 16, 64)
		return v, err
	}
	return strconv.ParseUint(value, 10, 64)
}
