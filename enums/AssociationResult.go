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

// The AssociationResult enumerates the answers, which the server can give client's association request.
type AssociationResult int

const (
	// AssociationResultAccepted defines that the association request is accepted.
	AssociationResultAccepted AssociationResult = iota
	// AssociationResultPermanentRejected defines that the association request is permanently rejected.
	AssociationResultPermanentRejected
	// AssociationResultTransientRejected defines that the association request is transiently rejected.
	AssociationResultTransientRejected
)

// AssociationResultParse converts the given string into a AssociationResult value.
//
// It returns the corresponding AssociationResult constant if the string matches
// a known level name, or an error if the input is invalid.
func AssociationResultParse(value string) (AssociationResult, error) {
	var ret AssociationResult
	var err error
	switch {
	case strings.EqualFold(value, "Accepted"):
		ret = AssociationResultAccepted
	case strings.EqualFold(value, "PermanentRejected"):
		ret = AssociationResultPermanentRejected
	case strings.EqualFold(value, "TransientRejected"):
		ret = AssociationResultTransientRejected
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the association result.
// It satisfies fmt.Stringer.
func (g AssociationResult) String() string {
	var ret string
	switch g {
	case AssociationResultAccepted:
		ret = "Accepted"
	case AssociationResultPermanentRejected:
		ret = "PermanentRejected"
	case AssociationResultTransientRejected:
		ret = "TransientRejected"
	}
	return ret
}

// AllAssociationResult returns a slice containing all defined association result values.
func AllAssociationResult() []AssociationResult {
	return []AssociationResult{
		AssociationResultAccepted,
		AssociationResultPermanentRejected,
		AssociationResultTransientRejected,
	}
}
