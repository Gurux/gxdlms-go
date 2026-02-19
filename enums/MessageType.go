package enums

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

type MessageType int

const (
	//MessageTypeCosemApdu defines that message type is COSEM APDU.
	MessageTypeCosemApdu MessageType = iota
	//MessageTypeCosemApduXML defines that message type is COSEM APDU in XML format.
	MessageTypeCosemApduXML
	//MessageTypeManufacturerSpesific defines that message type is manufacturer spesific.
	MessageTypeManufacturerSpesific MessageType = 128
)

// MessageTypeParse converts the given string into a MessageType value.
//
// It returns the corresponding MessageType constant if the string matches
// a known level name, or an error if the input is invalid.
func MessageTypeParse(value string) (MessageType, error) {
	var ret MessageType
	var err error
	switch {
	case strings.EqualFold(value, "CosemApdu"):
		ret = MessageTypeCosemApdu
	case strings.EqualFold(value, "CosemApduXml"):
		ret = MessageTypeCosemApduXML
	case strings.EqualFold(value, "ManufacturerSpesific"):
		ret = MessageTypeManufacturerSpesific
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the MessageType.
// It satisfies fmt.Stringer.
func (g MessageType) String() string {
	var ret string
	switch g {
	case MessageTypeCosemApdu:
		ret = "CosemApdu"
	case MessageTypeCosemApduXML:
		ret = "CosemApduXml"
	case MessageTypeManufacturerSpesific:
		ret = "ManufacturerSpesific"
	}
	return ret
}

// AllMessageType returns a slice containing all defined MessageType values.
func AllMessageType() []MessageType {
	return []MessageType{
		MessageTypeCosemApdu,
		MessageTypeCosemApduXML,
		MessageTypeManufacturerSpesific,
	}
}
