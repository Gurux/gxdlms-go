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
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/types"
)

// Executed scripts.
type GXScheduleEntry struct {
	// Schedule entry index.
	Index uint16

	// Is Schedule entry enabled.
	Enable bool

	// Executed Script.
	Script *GXDLMSScriptTable

	// Script identifier of the script to be executed.
	ScriptSelector uint16

	// Switch time.
	SwitchTime types.GXTime

	// Defines a period in minutes, in which an entry shall be processed after power fail.
	ValidityWindow uint16

	// Days of the week on which the entry is valid.
	ExecWeekdays enums.Weekdays

	// Perform the link to the IC "Special days table", day_id.
	ExecSpecDays string

	// Date starting period in which the entry is valid.
	BeginDate types.GXDate

	// Date ending period in which the entry is valid.
	EndDate types.GXDate
}
