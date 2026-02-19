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

// CertificateType : .
type CertificateType int

const (
	// CertificateTypeDigitalSignature defines that the certificate type is digital signature.
	CertificateTypeDigitalSignature CertificateType = iota
	// CertificateTypeKeyAgreement defines that the certificate type is key agreement.
	CertificateTypeKeyAgreement
	// CertificateTypeTLS defines that the certificate type is TLS.
	CertificateTypeTLS
	// CertificateTypeOther defines that the certificate type is other.
	CertificateTypeOther
)

// CertificateTypeParse converts the given string into a CertificateType value.
//
// It returns the corresponding CertificateType constant if the string matches
// a known level name, or an error if the input is invalid.
func CertificateTypeParse(value string) (CertificateType, error) {
	var ret CertificateType
	var err error
	switch {
	case strings.EqualFold(value, "DigitalSignature"):
		ret = CertificateTypeDigitalSignature
	case strings.EqualFold(value, "KeyAgreement"):
		ret = CertificateTypeKeyAgreement
	case strings.EqualFold(value, "TLS"):
		ret = CertificateTypeTLS
	case strings.EqualFold(value, "Other"):
		ret = CertificateTypeOther
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the CertificateType.
// It satisfies fmt.Stringer.
func (g CertificateType) String() string {
	var ret string
	switch g {
	case CertificateTypeDigitalSignature:
		ret = "DigitalSignature"
	case CertificateTypeKeyAgreement:
		ret = "KeyAgreement"
	case CertificateTypeTLS:
		ret = "TLS"
	case CertificateTypeOther:
		ret = "Other"
	}
	return ret
}

// AllCertificateType returns a slice containing all defined CertificateType values.
func AllCertificateType() []CertificateType {
	return []CertificateType{
	CertificateTypeDigitalSignature,
	CertificateTypeKeyAgreement,
	CertificateTypeTLS,
	CertificateTypeOther,
	}
}
