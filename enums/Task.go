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

// Task describes load task errors.
type Task int

const (
	// TaskOther defines that the other error.
	TaskOther Task = iota
	// TaskNoRemoteControl defines that the no remote control.
	TaskNoRemoteControl
	// TaskTiStopped defines that the ti is stopped.
	TaskTiStopped
	// TaskTiRunning defines that the ti is running.
	TaskTiRunning
	// TaskTiUnusable defines that the ti is unusable.
	TaskTiUnusable
)

// TaskParse converts the given string into a Task value.
//
// It returns the corresponding Task constant if the string matches
// a known level name, or an error if the input is invalid.
func TaskParse(value string) (Task, error) {
	var ret Task
	var err error
	switch strings.ToUpper(value) {
	case "OTHER":
		ret = TaskOther
	case "NOREMOTECONTROL":
		ret = TaskNoRemoteControl
	case "TISTOPPED":
		ret = TaskTiStopped
	case "TIRUNNING":
		ret = TaskTiRunning
	case "TIUNUSABLE":
		ret = TaskTiUnusable
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the Task.
// It satisfies fmt.Stringer.
func (g Task) String() string {
	var ret string
	switch g {
	case TaskOther:
		ret = "OTHER"
	case TaskNoRemoteControl:
		ret = "NOREMOTECONTROL"
	case TaskTiStopped:
		ret = "TISTOPPED"
	case TaskTiRunning:
		ret = "TIRUNNING"
	case TaskTiUnusable:
		ret = "TIUNUSABLE"
	}
	return ret
}
