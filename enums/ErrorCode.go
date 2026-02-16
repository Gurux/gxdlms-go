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

// ErrorCode Enumerated DLMS error codes.
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
	switch strings.ToUpper(value) {
	case "DISCONNECTMODE":
		ret = ErrorCodeDisconnectMode
	case "RECEIVENOTREADY":
		ret = ErrorCodeReceiveNotReady
	case "REJECTED":
		ret = ErrorCodeRejected
	case "UNACCEPTABLEFRAME":
		ret = ErrorCodeUnacceptableFrame
	case "OK":
		ret = ErrorCodeOk
	case "HARDWAREFAULT":
		ret = ErrorCodeHardwareFault
	case "TEMPORARYFAILURE":
		ret = ErrorCodeTemporaryFailure
	case "READWRITEDENIED":
		ret = ErrorCodeReadWriteDenied
	case "UNDEFINEDOBJECT":
		ret = ErrorCodeUndefinedObject
	case "INCONSISTENTCLASS":
		ret = ErrorCodeInconsistentClass
	case "UNAVAILABLEOBJECT":
		ret = ErrorCodeUnavailableObject
	case "UNMATCHEDTYPE":
		ret = ErrorCodeUnmatchedType
	case "ACCESSVIOLATED":
		ret = ErrorCodeAccessViolated
	case "DATABLOCKUNAVAILABLE":
		ret = ErrorCodeDataBlockUnavailable
	case "LONGGETORREADABORTED":
		ret = ErrorCodeLongGetOrReadAborted
	case "NOLONGGETORREADINPROGRESS":
		ret = ErrorCodeNoLongGetOrReadInProgress
	case "LONGSETORWRITEABORTED":
		ret = ErrorCodeLongSetOrWriteAborted
	case "NOLONGSETORWRITEINPROGRESS":
		ret = ErrorCodeNoLongSetOrWriteInProgress
	case "DATABLOCKNUMBERINVALID":
		ret = ErrorCodeDataBlockNumberInvalid
	case "OTHERREASON":
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
		ret = "DISCONNECTMODE"
	case ErrorCodeReceiveNotReady:
		ret = "RECEIVENOTREADY"
	case ErrorCodeRejected:
		ret = "REJECTED"
	case ErrorCodeUnacceptableFrame:
		ret = "UNACCEPTABLEFRAME"
	case ErrorCodeOk:
		ret = "OK"
	case ErrorCodeHardwareFault:
		ret = "HARDWAREFAULT"
	case ErrorCodeTemporaryFailure:
		ret = "TEMPORARYFAILURE"
	case ErrorCodeReadWriteDenied:
		ret = "READWRITEDENIED"
	case ErrorCodeUndefinedObject:
		ret = "UNDEFINEDOBJECT"
	case ErrorCodeInconsistentClass:
		ret = "INCONSISTENTCLASS"
	case ErrorCodeUnavailableObject:
		ret = "UNAVAILABLEOBJECT"
	case ErrorCodeUnmatchedType:
		ret = "UNMATCHEDTYPE"
	case ErrorCodeAccessViolated:
		ret = "ACCESSVIOLATED"
	case ErrorCodeDataBlockUnavailable:
		ret = "DATABLOCKUNAVAILABLE"
	case ErrorCodeLongGetOrReadAborted:
		ret = "LONGGETORREADABORTED"
	case ErrorCodeNoLongGetOrReadInProgress:
		ret = "NOLONGGETORREADINPROGRESS"
	case ErrorCodeLongSetOrWriteAborted:
		ret = "LONGSETORWRITEABORTED"
	case ErrorCodeNoLongSetOrWriteInProgress:
		ret = "NOLONGSETORWRITEINPROGRESS"
	case ErrorCodeDataBlockNumberInvalid:
		ret = "DATABLOCKNUMBERINVALID"
	case ErrorCodeOtherReason:
		ret = "OTHERREASON"
	}
	return ret
}
