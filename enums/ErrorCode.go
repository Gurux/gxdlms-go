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

// ErrorCode enumerates DLMS error codes.
// https://www.gurux.fi/Gurux.DLMS.ErrorCodes
type ErrorCode int

const (
	// ErrorCodeDisconnectMode defines that the disconnect Mode error has occurred.
	ErrorCodeDisconnectMode ErrorCode = -4
	// ErrorCodeReceiveNotReady defines that the Receive Not Ready.
	ErrorCodeReceiveNotReady = -3
	// ErrorCodeRejected defines that the connection is rejected.
	ErrorCodeRejected = -2
	// ErrorCodeUnacceptableFrame defines that the unacceptable frame.
	ErrorCodeUnacceptableFrame = -1
	// ErrorCodeOk defines that no error has occurred.
	ErrorCodeOk = 0
	// ErrorCodeHardwareFault defines that the device reports a hardware fault error has occurred.
	ErrorCodeHardwareFault = 1
	// ErrorCodeTemporaryFailure defines that the Device reports a temporary failure error has occurred.
	ErrorCodeTemporaryFailure = 2
	// ErrorCodeReadWriteDenied defines that the Device reports Read-Write denied error has occurred.
	ErrorCodeReadWriteDenied = 3
	// ErrorCodeUndefinedObject defines that the Device reports a undefined object error has occurred.
	ErrorCodeUndefinedObject = 4
	// ErrorCodeInconsistentClass defines that the device reports a inconsistent Class or object error has occurred.
	ErrorCodeInconsistentClass = 9
	// ErrorCodeUnavailableObject defines that the device reports a unavailable object error has occurred.
	ErrorCodeUnavailableObject = 11
	// ErrorCodeUnmatchedType defines that the device reports a unmatched type error has occurred.
	ErrorCodeUnmatchedType = 12
	// ErrorCodeAccessViolated defines that the device reports scope of access violated error has occurred.
	ErrorCodeAccessViolated = 13
	// ErrorCodeDataBlockUnavailable defines that the data Block Unavailable error has occurred.
	ErrorCodeDataBlockUnavailable = 14
	// ErrorCodeLongGetOrReadAborted defines that the long Get Or Read Aborted error has occurred.
	ErrorCodeLongGetOrReadAborted = 15
	// ErrorCodeNoLongGetOrReadInProgress defines that the no Long Get Or Read In Progress error has occurred.
	ErrorCodeNoLongGetOrReadInProgress = 16
	// ErrorCodeLongSetOrWriteAborted defines that the long Set Or Write Aborted error has occurred.
	ErrorCodeLongSetOrWriteAborted = 17
	// ErrorCodeNoLongSetOrWriteInProgress defines that the no Long Set Or Write In Progress error has occurred.
	ErrorCodeNoLongSetOrWriteInProgress = 18
	// ErrorCodeDataBlockNumberInvalid defines that the data Block Number Invalid error has occurred.
	ErrorCodeDataBlockNumberInvalid = 19
	// ErrorCodeOtherReason defines that the other Reason error has occurred.
	ErrorCodeOtherReason = 250
)

// ErrorCodeParse converts the given string into a ErrorCode value.
//
// It returns the corresponding ErrorCode constant if the string matches
// a known level name, or an error if the input is invalid.
func ErrorCodeParse(value string) (ErrorCode, error) {
	var ret ErrorCode
	var err error
	switch {
	case strings.EqualFold(value, "DisconnectMode"):
		ret = ErrorCodeDisconnectMode
	case strings.EqualFold(value, "ReceiveNotReady"):
		ret = ErrorCodeReceiveNotReady
	case strings.EqualFold(value, "Rejected"):
		ret = ErrorCodeRejected
	case strings.EqualFold(value, "UnacceptableFrame"):
		ret = ErrorCodeUnacceptableFrame
	case strings.EqualFold(value, "Ok"):
		ret = ErrorCodeOk
	case strings.EqualFold(value, "HardwareFault"):
		ret = ErrorCodeHardwareFault
	case strings.EqualFold(value, "TemporaryFailure"):
		ret = ErrorCodeTemporaryFailure
	case strings.EqualFold(value, "ReadWriteDenied"):
		ret = ErrorCodeReadWriteDenied
	case strings.EqualFold(value, "UndefinedObject"):
		ret = ErrorCodeUndefinedObject
	case strings.EqualFold(value, "InconsistentClass"):
		ret = ErrorCodeInconsistentClass
	case strings.EqualFold(value, "UnavailableObject"):
		ret = ErrorCodeUnavailableObject
	case strings.EqualFold(value, "UnmatchedType"):
		ret = ErrorCodeUnmatchedType
	case strings.EqualFold(value, "AccessViolated"):
		ret = ErrorCodeAccessViolated
	case strings.EqualFold(value, "DataBlockUnavailable"):
		ret = ErrorCodeDataBlockUnavailable
	case strings.EqualFold(value, "LongGetOrReadAborted"):
		ret = ErrorCodeLongGetOrReadAborted
	case strings.EqualFold(value, "NoLongGetOrReadInProgress"):
		ret = ErrorCodeNoLongGetOrReadInProgress
	case strings.EqualFold(value, "LongSetOrWriteAborted"):
		ret = ErrorCodeLongSetOrWriteAborted
	case strings.EqualFold(value, "NoLongSetOrWriteInProgress"):
		ret = ErrorCodeNoLongSetOrWriteInProgress
	case strings.EqualFold(value, "DataBlockNumberInvalid"):
		ret = ErrorCodeDataBlockNumberInvalid
	case strings.EqualFold(value, "OtherReason"):
		ret = ErrorCodeOtherReason
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the ErrorCode.
// It satisfies fmt.Stringer.
func (g ErrorCode) String() string {
	var ret string
	switch g {
	case ErrorCodeDisconnectMode:
		ret = "DisconnectMode"
	case ErrorCodeReceiveNotReady:
		ret = "ReceiveNotReady"
	case ErrorCodeRejected:
		ret = "Rejected"
	case ErrorCodeUnacceptableFrame:
		ret = "UnacceptableFrame"
	case ErrorCodeOk:
		ret = "Ok"
	case ErrorCodeHardwareFault:
		ret = "HardwareFault"
	case ErrorCodeTemporaryFailure:
		ret = "TemporaryFailure"
	case ErrorCodeReadWriteDenied:
		ret = "ReadWriteDenied"
	case ErrorCodeUndefinedObject:
		ret = "UndefinedObject"
	case ErrorCodeInconsistentClass:
		ret = "InconsistentClass"
	case ErrorCodeUnavailableObject:
		ret = "UnavailableObject"
	case ErrorCodeUnmatchedType:
		ret = "UnmatchedType"
	case ErrorCodeAccessViolated:
		ret = "AccessViolated"
	case ErrorCodeDataBlockUnavailable:
		ret = "DataBlockUnavailable"
	case ErrorCodeLongGetOrReadAborted:
		ret = "LongGetOrReadAborted"
	case ErrorCodeNoLongGetOrReadInProgress:
		ret = "NoLongGetOrReadInProgress"
	case ErrorCodeLongSetOrWriteAborted:
		ret = "LongSetOrWriteAborted"
	case ErrorCodeNoLongSetOrWriteInProgress:
		ret = "NoLongSetOrWriteInProgress"
	case ErrorCodeDataBlockNumberInvalid:
		ret = "DataBlockNumberInvalid"
	case ErrorCodeOtherReason:
		ret = "OtherReason"
	}
	return ret
}

// AllErrorCode returns a slice containing all defined ErrorCode values.
func AllErrorCode() []ErrorCode {
	return []ErrorCode{
		ErrorCodeDisconnectMode,
		ErrorCodeReceiveNotReady,
		ErrorCodeRejected,
		ErrorCodeUnacceptableFrame,
		ErrorCodeOk,
		ErrorCodeHardwareFault,
		ErrorCodeTemporaryFailure,
		ErrorCodeReadWriteDenied,
		ErrorCodeUndefinedObject,
		ErrorCodeInconsistentClass,
		ErrorCodeUnavailableObject,
		ErrorCodeUnmatchedType,
		ErrorCodeAccessViolated,
		ErrorCodeDataBlockUnavailable,
		ErrorCodeLongGetOrReadAborted,
		ErrorCodeNoLongGetOrReadInProgress,
		ErrorCodeLongSetOrWriteAborted,
		ErrorCodeNoLongSetOrWriteInProgress,
		ErrorCodeDataBlockNumberInvalid,
		ErrorCodeOtherReason,
	}
}
