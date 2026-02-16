package internal

import "github.com/Gurux/gxdlms-go/enums"

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

// GXDLMSServer is used to implement server side of DLMS/COSEM protocol. It is used to handle incoming requests and send responses back to client.
type IGXDLMSServer interface {
	NotifyGetAttributeAccess(args *ValueEventArgs) int

	NotifyGetMethodAccess(args *ValueEventArgs) int

	NotifyRead(args []*ValueEventArgs)
	NotifyWrite(args []*ValueEventArgs)
	NotifyPostRead(args []*ValueEventArgs)
	NotifyPostWrite(args []*ValueEventArgs)

	NotifyPreAction(args []*ValueEventArgs)
	NotifyPostAction(args []*ValueEventArgs)

	NotifyFindObject(objectType enums.ObjectType, sn int, ln string) interface{}
	Items() interface{}

	Transaction() interface{}
	SetTransaction(value interface{})
}
