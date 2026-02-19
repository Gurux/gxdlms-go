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

type ImageTransferStatus int

const (
	ImageTransferStatusNotInitiated ImageTransferStatus = iota
	ImageTransferStatusTransferInitiated
	ImageTransferStatusVerificationInitiated
	ImageTransferStatusVerificationSuccessful
	ImageTransferStatusVerificationFailed
	ImageTransferStatusActivationInitiated
	ImageTransferStatusActivationSuccessful
	ImageTransferStatusActivationFailed
)

// ImageTransferStatusParse converts the given string into a ImageTransferStatus value.
//
// It returns the corresponding ImageTransferStatus constant if the string matches
// a known level name, or an error if the input is invalid.
func ImageTransferStatusParse(value string) (ImageTransferStatus, error) {
	var ret ImageTransferStatus
	var err error
	switch {
	case strings.EqualFold(value, "NotInitiated"):
		ret = ImageTransferStatusNotInitiated
	case strings.EqualFold(value, "TransferInitiated"):
		ret = ImageTransferStatusTransferInitiated
	case strings.EqualFold(value, "VerificationInitiated"):
		ret = ImageTransferStatusVerificationInitiated
	case strings.EqualFold(value, "VerificationSuccessful"):
		ret = ImageTransferStatusVerificationSuccessful
	case strings.EqualFold(value, "VerificationFailed"):
		ret = ImageTransferStatusVerificationFailed
	case strings.EqualFold(value, "ActivationInitiated"):
		ret = ImageTransferStatusActivationInitiated
	case strings.EqualFold(value, "ActivationSuccessful"):
		ret = ImageTransferStatusActivationSuccessful
	case strings.EqualFold(value, "ActivationFailed"):
		ret = ImageTransferStatusActivationFailed
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the ImageTransferStatus.
// It satisfies fmt.Stringer.
func (g ImageTransferStatus) String() string {
	var ret string
	switch g {
	case ImageTransferStatusNotInitiated:
		ret = "NotInitiated"
	case ImageTransferStatusTransferInitiated:
		ret = "TransferInitiated"
	case ImageTransferStatusVerificationInitiated:
		ret = "VerificationInitiated"
	case ImageTransferStatusVerificationSuccessful:
		ret = "VerificationSuccessful"
	case ImageTransferStatusVerificationFailed:
		ret = "VerificationFailed"
	case ImageTransferStatusActivationInitiated:
		ret = "ActivationInitiated"
	case ImageTransferStatusActivationSuccessful:
		ret = "ActivationSuccessful"
	case ImageTransferStatusActivationFailed:
		ret = "ActivationFailed"
	}
	return ret
}

// AllImageTransferStatus returns a slice containing all defined ImageTransferStatus values.
func AllImageTransferStatus() []ImageTransferStatus {
	return []ImageTransferStatus{
	ImageTransferStatusNotInitiated,
	ImageTransferStatusTransferInitiated,
	ImageTransferStatusVerificationInitiated,
	ImageTransferStatusVerificationSuccessful,
	ImageTransferStatusVerificationFailed,
	ImageTransferStatusActivationInitiated,
	ImageTransferStatusActivationSuccessful,
	ImageTransferStatusActivationFailed,
	}
}
