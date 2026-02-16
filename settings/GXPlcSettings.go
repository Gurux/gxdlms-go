package settings

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

// PLC communication settings.
type GXPlcSettings struct {
	_systemTitle []byte

	settings GXDLMSSettings

	// Initial credit (IC) tells how many times the frame must be repeated. Maximum value is 7.
	InitialCredit uint8

	// The current credit (CC) initial value equal to IC and automatically decremented by the MAC layer after each repetition.
	//          Maximum value is 7.
	CurrentCredit uint8

	// Delta credit (DC) is used by the system management application entity
	//          (SMAE) of the Client for credit management, while it has no meaning for a Server or a REPEATER.
	//          It represents the difference(IC-CC) of the last communication originated by the system identified by the DA address to the system identified by the SA address.
	//           Maximum value is 3.
	DeltaCredit uint8

	// Source MAC address.
	MacSourceAddress uint16

	// Destination MAC address.
	MacDestinationAddress uint16

	// Response probability.
	ResponseProbability uint8

	// Allowed time slots.
	AllowedTimeSlots uint16

	// Server saves client system title.
	ClientSystemTitle []byte
}
