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

// LoadDataSet LoadDataSet describes load dataset errors.
type LoadDataSet int

const (
	// LoadDataSetOther defines that the other error has occurred.
	LoadDataSetOther LoadDataSet = iota
	// LoadDataSetPrimitiveOutOfSequence defines that the primitive out of sequence error has occurred.
	LoadDataSetPrimitiveOutOfSequence
	// LoadDataSetNotLoadable defines that the not loadable error has occurred.
	LoadDataSetNotLoadable
	// LoadDataSetDatasetSizeTooLarge defines that the dataset size is too large error has occurred.
	LoadDataSetDatasetSizeTooLarge
	// LoadDataSetNotAwaitedSegment defines that the not awaited segment error has occurred.
	LoadDataSetNotAwaitedSegment
	// LoadDataSetInterpretationFailure defines that the interpretation failure error has occurred.
	LoadDataSetInterpretationFailure
	// LoadDataSetStorageFailure defines that the storage failure error has occurred.
	LoadDataSetStorageFailure
	// LoadDataSetDatasetNotReady defines that the dataset not ready error has occurred.
	LoadDataSetDatasetNotReady
)

// LoadDataSetParse converts the given string into a LoadDataSet value.
//
// It returns the corresponding LoadDataSet constant if the string matches
// a known level name, or an error if the input is invalid.
func LoadDataSetParse(value string) (LoadDataSet, error) {
	var ret LoadDataSet
	var err error
	switch {
	case strings.EqualFold(value, "Other"):
		ret = LoadDataSetOther
	case strings.EqualFold(value, "PrimitiveOutOfSequence"):
		ret = LoadDataSetPrimitiveOutOfSequence
	case strings.EqualFold(value, "NotLoadable"):
		ret = LoadDataSetNotLoadable
	case strings.EqualFold(value, "DatasetSizeTooLarge"):
		ret = LoadDataSetDatasetSizeTooLarge
	case strings.EqualFold(value, "NotAwaitedSegment"):
		ret = LoadDataSetNotAwaitedSegment
	case strings.EqualFold(value, "InterpretationFailure"):
		ret = LoadDataSetInterpretationFailure
	case strings.EqualFold(value, "StorageFailure"):
		ret = LoadDataSetStorageFailure
	case strings.EqualFold(value, "DatasetNotReady"):
		ret = LoadDataSetDatasetNotReady
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the LoadDataSet.
// It satisfies fmt.Stringer.
func (g LoadDataSet) String() string {
	var ret string
	switch g {
	case LoadDataSetOther:
		ret = "Other"
	case LoadDataSetPrimitiveOutOfSequence:
		ret = "PrimitiveOutOfSequence"
	case LoadDataSetNotLoadable:
		ret = "NotLoadable"
	case LoadDataSetDatasetSizeTooLarge:
		ret = "DatasetSizeTooLarge"
	case LoadDataSetNotAwaitedSegment:
		ret = "NotAwaitedSegment"
	case LoadDataSetInterpretationFailure:
		ret = "InterpretationFailure"
	case LoadDataSetStorageFailure:
		ret = "StorageFailure"
	case LoadDataSetDatasetNotReady:
		ret = "DatasetNotReady"
	}
	return ret
}

// AllLoadDataSet returns a slice containing all defined LoadDataSet values.
func AllLoadDataSet() []LoadDataSet {
	return []LoadDataSet{
		LoadDataSetOther,
		LoadDataSetPrimitiveOutOfSequence,
		LoadDataSetNotLoadable,
		LoadDataSetDatasetSizeTooLarge,
		LoadDataSetNotAwaitedSegment,
		LoadDataSetInterpretationFailure,
		LoadDataSetStorageFailure,
		LoadDataSetDatasetNotReady,
	}
}
