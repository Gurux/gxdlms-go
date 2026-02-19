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

// MBusCommand defines M-Bus commands.
type MBusCommand int

const (
	// MBusCommandRspUd defines that the // Access demand from Meter to Other Device. This message requests an access to
	// the Meter (contains no application data).
	MBusCommandRspUd MBusCommand = 0x8
	// MBusCommandSndNr defines that the // Send unsolicited/periodical application data without request (Send/No Reply)
	MBusCommandSndNr MBusCommand = 0x4
	// MBusCommandSndUd defines that the // Send a command (Send User Data).
	MBusCommandSndUd MBusCommand = 0x3
)

// String returns the canonical name of the MBusCommand.
// It satisfies fmt.Stringer.
func (g MBusCommand) String() string {
	var ret string
	switch g {
	case MBusCommandRspUd:
		ret = "RspUd"
	case MBusCommandSndNr:
		ret = "SndNr"
	case MBusCommandSndUd:
		ret = "SndUd"
	}
	return ret
}
