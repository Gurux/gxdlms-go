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

// CertificateEntity : .
type CertificateEntity int

const (
	// CertificateEntityServer defines that the certificate entity is server.
	CertificateEntityServer CertificateEntity = iota
	// CertificateEntityClient defines that the certificate entity is client.
	CertificateEntityClient
	// CertificateEntityCertificationAuthority defines that the certificate entity is certification authority.
	CertificateEntityCertificationAuthority
	// CertificateEntityOther defines that the certificate entity is other.
	CertificateEntityOther
)

// CertificateEntityParse converts the given string into a CertificateEntity value.
//
// It returns the corresponding CertificateEntity constant if the string matches
// a known level name, or an error if the input is invalid.
func CertificateEntityParse(value string) (CertificateEntity, error) {
	var ret CertificateEntity
	var err error
	switch {
	case strings.EqualFold(value, "Server"):
		ret = CertificateEntityServer
	case strings.EqualFold(value, "Client"):
		ret = CertificateEntityClient
	case strings.EqualFold(value, "CertificationAuthority"):
		ret = CertificateEntityCertificationAuthority
	case strings.EqualFold(value, "Other"):
		ret = CertificateEntityOther
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the CertificateEntity.
// It satisfies fmt.Stringer.
func (g CertificateEntity) String() string {
	var ret string
	switch g {
	case CertificateEntityServer:
		ret = "Server"
	case CertificateEntityClient:
		ret = "Client"
	case CertificateEntityCertificationAuthority:
		ret = "CertificationAuthority"
	case CertificateEntityOther:
		ret = "Other"
	}
	return ret
}

// AllCertificateEntity returns a slice containing all defined CertificateEntity values.
func AllCertificateEntity() []CertificateEntity {
	return []CertificateEntity{
	CertificateEntityServer,
	CertificateEntityClient,
	CertificateEntityCertificationAuthority,
	CertificateEntityOther,
	}
}
