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

// Enumerates how data is access on read or write.
type VariableAccessSpecification int

const (
	// VariableAccessSpecificationVariableName defines that the Read data using SN.
	VariableAccessSpecificationVariableName VariableAccessSpecification = 2
	// VariableAccessSpecificationParameterisedAccess defines that the Get data using parameterised access.
	VariableAccessSpecificationParameterisedAccess VariableAccessSpecification = 4
	// VariableAccessSpecificationBlockNumberAccess defines that the Get next block.
	VariableAccessSpecificationBlockNumberAccess VariableAccessSpecification = 5
	// VariableAccessSpecificationReadDataBlockAccess defines that the Read data as blocks.
	VariableAccessSpecificationReadDataBlockAccess VariableAccessSpecification = 6
	// VariableAccessSpecificationWriteDataBlockAccess defines that the Write data as blocks.
	VariableAccessSpecificationWriteDataBlockAccess VariableAccessSpecification = 7
)

// VariableAccessSpecificationParse converts the given string into a VariableAccessSpecification value.
//
// It returns the corresponding VariableAccessSpecification constant if the string matches
// a known level name, or an error if the input is invalid.
func VariableAccessSpecificationParse(value string) (VariableAccessSpecification, error) {
	var ret VariableAccessSpecification
	var err error
	switch strings.ToUpper(value) {
	case "VARIABLENAME":
		ret = VariableAccessSpecificationVariableName
	case "PARAMETERISEDACCESS":
		ret = VariableAccessSpecificationParameterisedAccess
	case "BLOCKNUMBERACCESS":
		ret = VariableAccessSpecificationBlockNumberAccess
	case "READDATABLOCKACCESS":
		ret = VariableAccessSpecificationReadDataBlockAccess
	case "WRITEDATABLOCKACCESS":
		ret = VariableAccessSpecificationWriteDataBlockAccess
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the VariableAccessSpecification.
// It satisfies fmt.Stringer.
func (g VariableAccessSpecification) String() string {
	var ret string
	switch g {
	case VariableAccessSpecificationVariableName:
		ret = "VARIABLENAME"
	case VariableAccessSpecificationParameterisedAccess:
		ret = "PARAMETERISEDACCESS"
	case VariableAccessSpecificationBlockNumberAccess:
		ret = "BLOCKNUMBERACCESS"
	case VariableAccessSpecificationReadDataBlockAccess:
		ret = "READDATABLOCKACCESS"
	case VariableAccessSpecificationWriteDataBlockAccess:
		ret = "WRITEDATABLOCKACCESS"
	}
	return ret
}
