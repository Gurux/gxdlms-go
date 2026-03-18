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
	ErrorCodeReceiveNotReady ErrorCode = -3
	// ErrorCodeRejected defines that the connection is rejected.
	ErrorCodeRejected ErrorCode = -2
	// ErrorCodeUnacceptableFrame defines that the unacceptable frame.
	ErrorCodeUnacceptableFrame ErrorCode = -1
	// ErrorCodeOk defines that no error has occurred.
	ErrorCodeOk ErrorCode = 0
	// ErrorCodeHardwareFault defines that the device reports a hardware fault error has occurred.
	ErrorCodeHardwareFault ErrorCode = 1
	// ErrorCodeTemporaryFailure defines that the Device reports a temporary failure error has occurred.
	ErrorCodeTemporaryFailure ErrorCode = 2
	// ErrorCodeReadWriteDenied defines that the Device reports Read-Write denied error has occurred.
	ErrorCodeReadWriteDenied ErrorCode = 3
	// ErrorCodeUndefinedObject defines that the Device reports a undefined object error has occurred.
	ErrorCodeUndefinedObject ErrorCode = 4
	// ErrorCodeInconsistentClass defines that the device reports a inconsistent Class or object error has occurred.
	ErrorCodeInconsistentClass ErrorCode = 9
	// ErrorCodeUnavailableObject defines that the device reports a unavailable object error has occurred.
	ErrorCodeUnavailableObject ErrorCode = 11
	// ErrorCodeUnmatchedType defines that the device reports a unmatched type error has occurred.
	ErrorCodeUnmatchedType ErrorCode = 12
	// ErrorCodeAccessViolated defines that the device reports scope of access violated error has occurred.
	ErrorCodeAccessViolated ErrorCode = 13
	// ErrorCodeDataBlockUnavailable defines that the data Block Unavailable error has occurred.
	ErrorCodeDataBlockUnavailable ErrorCode = 14
	// ErrorCodeLongGetOrReadAborted defines that the long Get Or Read Aborted error has occurred.
	ErrorCodeLongGetOrReadAborted ErrorCode = 15
	// ErrorCodeNoLongGetOrReadInProgress defines that the no Long Get Or Read In Progress error has occurred.
	ErrorCodeNoLongGetOrReadInProgress ErrorCode = 16
	// ErrorCodeLongSetOrWriteAborted defines that the long Set Or Write Aborted error has occurred.
	ErrorCodeLongSetOrWriteAborted ErrorCode = 17
	// ErrorCodeNoLongSetOrWriteInProgress defines that the no Long Set Or Write In Progress error has occurred.
	ErrorCodeNoLongSetOrWriteInProgress ErrorCode = 18
	// ErrorCodeDataBlockNumberInvalid defines that the data Block Number Invalid error has occurred.
	ErrorCodeDataBlockNumberInvalid ErrorCode = 19
	// ErrorCodeOtherReason defines that the other Reason error has occurred.
	ErrorCodeOtherReason ErrorCode = 250
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

// Error returns the string representation of the ErrorCode.
func (e ErrorCode) Error() string {
	switch e {
	case ErrorCodeDisconnectMode:
		return "disconnect mode"
	case ErrorCodeReceiveNotReady:
		return "receive not ready"
	case ErrorCodeRejected:
		return "connection rejected"
	case ErrorCodeUnacceptableFrame:
		return "unacceptable frame"
	case ErrorCodeOk:
		return "ok"
	case ErrorCodeHardwareFault:
		return "hardware fault"
	case ErrorCodeTemporaryFailure:
		return "temporary failure"
	case ErrorCodeReadWriteDenied:
		return "read-write denied"
	case ErrorCodeUndefinedObject:
		return "undefined object"
	case ErrorCodeInconsistentClass:
		return "inconsistent class"
	case ErrorCodeUnavailableObject:
		return "unavailable object"
	case ErrorCodeUnmatchedType:
		return "unmatched type"
	case ErrorCodeAccessViolated:
		return "access violated"
	case ErrorCodeDataBlockUnavailable:
		return "data block unavailable"
	case ErrorCodeLongGetOrReadAborted:
		return "long get/read aborted"
	case ErrorCodeNoLongGetOrReadInProgress:
		return "no long get/read in progress"
	case ErrorCodeLongSetOrWriteAborted:
		return "long set/write aborted"
	case ErrorCodeNoLongSetOrWriteInProgress:
		return "no long set/write in progress"
	case ErrorCodeDataBlockNumberInvalid:
		return "data block number invalid"
	case ErrorCodeOtherReason:
		return "other reason"
	default:
		return fmt.Sprintf("unknown error code: %d", int(e))
	}
}
