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

// CertificateIdentificationType Certificate is identified with entity identification
//
//	or the serial number of the certificate.
type CertificateIdentificationType int

const (
	// CertificateIdentificationTypeEntity defines that the certificate is identified with entity identification.
	CertificateIdentificationTypeEntity CertificateIdentificationType = iota
	// CertificateIdentificationTypeSerial defines that the certificate is identified with serial number of the certificate.
	CertificateIdentificationTypeSerial
)

// CertificateIdentificationTypeParse converts the given string into a CertificateIdentificationType value.
//
// It returns the corresponding CertificateIdentificationType constant if the string matches
// a known level name, or an error if the input is invalid.
func CertificateIdentificationTypeParse(value string) (CertificateIdentificationType, error) {
	var ret CertificateIdentificationType
	var err error
	switch {
	case strings.EqualFold(value, "Entity"):
		ret = CertificateIdentificationTypeEntity
	case strings.EqualFold(value, "Serial"):
		ret = CertificateIdentificationTypeSerial
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the CertificateIdentificationType.
// It satisfies fmt.Stringer.
func (g CertificateIdentificationType) String() string {
	var ret string
	switch g {
	case CertificateIdentificationTypeEntity:
		ret = "Entity"
	case CertificateIdentificationTypeSerial:
		ret = "Serial"
	}
	return ret
}

// AllCertificateIdentificationType returns a slice containing all defined CertificateIdentificationType values.
func AllCertificateIdentificationType() []CertificateIdentificationType {
	return []CertificateIdentificationType{
	CertificateIdentificationTypeEntity,
	CertificateIdentificationTypeSerial,
	}
}
