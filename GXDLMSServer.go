package dlms

import (
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
)

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
// Gurux Device Framework is Open Source software you can redistribute it
// and/or modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation version 2 of the License.
// Gurux Device Framework is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
// See the GNU General Public License for more details.
//
// More information of Gurux products: https://www.gurux.org
//
// This code is licensed under the GNU General Public License v2.
// Full text may be retrieved at http://www.gnu.org/licenses/gpl-2.0.txt
//---------------------------------------------------------------------------

type GXDLMSServer struct {
	transaction *gxDLMSLongTransaction
}

func getServerTransaction(server internal.IGXDLMSServer) *gxDLMSLongTransaction {
	if server.Transaction() != nil {
		return server.Transaction().(*gxDLMSLongTransaction)
	}
	return nil
}

func (g *GXDLMSServer) NotifyGetMethodAccess(args *internal.ValueEventArgs) int {
	return 0
}

func (g *GXDLMSServer) NotifyGetAttributeAccess(args *internal.ValueEventArgs) int {
	return 0
}

func (g *GXDLMSServer) NotifyRead(args []*internal.ValueEventArgs) int {
	return 0
}

func (g *GXDLMSServer) NotifyPostRead(args []*internal.ValueEventArgs) int {
	return 0
}

func (g *GXDLMSServer) NotifyPostWrite(args []*internal.ValueEventArgs) int {
	return 0
}
func (g *GXDLMSServer) NotifyWrite(args []*internal.ValueEventArgs) int {
	return 0
}
func (g *GXDLMSServer) NotifyFindObject(objectType enums.ObjectType, sn int, ln string) interface{} {
	return nil
}
func (g *GXDLMSServer) Items() interface{} {
	return nil
}

func (g *GXDLMSServer) NotifyPreAction(args []*internal.ValueEventArgs) int {
	return 0
}

func (g *GXDLMSServer) NotifyPostAction(args []*internal.ValueEventArgs) int {
	return 0
}

func (g *GXDLMSServer) NotifyConnected(connectionInfo *GXDLMSConnectionEventArgs) {
}

func (g *GXDLMSServer) NotifyInvalidConnection(connectionInfo *GXDLMSConnectionEventArgs) {
}

func (g *GXDLMSServer) Transaction() interface{} {
	return g.transaction
}
func (g *GXDLMSServer) SetTransaction(value interface{}) {
	g.transaction = value.(*gxDLMSLongTransaction)
}

// GenerateConfirmedServiceError returns the generate confirmed service error.
//
// Parameters:
//
//	service: Confirmed service error.
//	type: Service error.
//	code: code
func GenerateConfirmedServiceError(service enums.ConfirmedServiceError, type_ enums.ServiceError, code uint8) []byte {
	return []byte{uint8(enums.CommandConfirmedServiceError), uint8(service), uint8(type_), code}
}
