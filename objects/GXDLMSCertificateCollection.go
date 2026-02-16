package objects

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
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/types"
)

// Certificate info.
type GXDLMSCertificateCollection []*GXDLMSCertificateInfo

// Clear removes all certificates from the collection.
func (g *GXDLMSCertificateCollection) Clear() {
	*g = (*g)[:0]
}

// Find returns the find certificate with given parameters.
//
// Parameters:
//
//	entity: Certificate entity.
//	type: Certificate type.
//	systemtitle: System title.
func (g *GXDLMSCertificateCollection) Find(entity enums.CertificateEntity, type_ enums.CertificateType, systemtitle []byte) *GXDLMSCertificateInfo {
	subject := types.Asn1SystemTitleToSubject(systemtitle)
	for _, it := range *g {
		if (it.Entity == enums.CertificateEntityServer && entity == enums.CertificateEntityServer) ||
			(it.Entity == enums.CertificateEntityClient && entity == enums.CertificateEntityClient) && it.Subject == subject {
			return it
		}
	}
	return nil
}
