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
	"fmt"

	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/types"
)

type GXDLMSScriptAction struct {
	// Executed object
	Target IGXDLMSBase

	// Defines which action to be applied to the referenced object.
	Type enums.ScriptActionType

	// Defines which attribute of the selected object is affected
	// or which specific method is to be executed.
	Index int8

	// Parameter is service spesific.
	Parameter any

	// Parameter data type can be used to tell spesific data type.
	ParameterDataType enums.DataType
}

func (g *GXDLMSScriptAction) String() string {
	var tmp string
	if v, ok := g.Parameter.([]byte); ok {
		tmp = types.ToHex(v, false)
	} else {
		tmp = fmt.Sprint(g.Parameter)
	}
	if g.Target != nil {
		return g.Type.String() + " " + g.Target.Base().ObjectType().String() + " " + g.Target.Base().LogicalName() + " " +
			fmt.Sprint(g.Index) + " " + tmp
	}
	return g.Type.String() + " " + fmt.Sprint(g.Index) + " " + tmp
}
