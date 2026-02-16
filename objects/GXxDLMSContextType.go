package objects

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
	"github.com/Gurux/gxdlms-go/types"
)

type GXxDLMSContextType struct {
	// Conformance
	Conformance enums.Conformance

	MaxReceivePduSize uint16

	MaxSendPduSize uint16

	// Dlms Version Number.
	DlmsVersionNumber uint8

	QualityOfService int8

	CypheringInfo []byte
}

// String returns string presentation of GXxDLMSContextType struct.
func (g *GXxDLMSContextType) String() string {
	return g.Conformance.String() + " " + strconv.Itoa(int(g.MaxReceivePduSize)) + " " + strconv.Itoa(int(g.MaxSendPduSize)) + " " + strconv.Itoa(int(g.DlmsVersionNumber)) + " " + strconv.Itoa(int(g.QualityOfService)) + " " + types.ToHex(g.CypheringInfo, true)
}
