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

// CertificateVersion .
type CertificateVersion int

const (
	// CertificateVersionVersion1 defines that the certificate version 1.
	CertificateVersionVersion1 CertificateVersion = iota
	// CertificateVersionVersion2 defines that the certificate version 2.
	CertificateVersionVersion2
	// CertificateVersionVersion3 defines that the certificate version 3.
	CertificateVersionVersion3
)

// CertificateVersionParse converts the given string into a CertificateVersion value.
//
// It returns the corresponding CertificateVersion constant if the string matches
// a known level name, or an error if the input is invalid.
func CertificateVersionParse(value string) (CertificateVersion, error) {
	var ret CertificateVersion
	var err error
	switch {
	case strings.EqualFold(value, "Version1"):
		ret = CertificateVersionVersion1
	case strings.EqualFold(value, "Version2"):
		ret = CertificateVersionVersion2
	case strings.EqualFold(value, "Version3"):
		ret = CertificateVersionVersion3
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the CertificateVersion.
// It satisfies fmt.Stringer.
func (g CertificateVersion) String() string {
	var ret string
	switch g {
	case CertificateVersionVersion1:
		ret = "Version1"
	case CertificateVersionVersion2:
		ret = "Version2"
	case CertificateVersionVersion3:
		ret = "Version3"
	}
	return ret
}

// AllCertificateVersion returns a slice containing all defined CertificateVersion values.
func AllCertificateVersion() []CertificateVersion {
	return []CertificateVersion{
	CertificateVersionVersion1,
	CertificateVersionVersion2,
	CertificateVersionVersion3,
	}
}
