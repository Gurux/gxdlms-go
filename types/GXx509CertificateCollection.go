package types

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
	"log"
	"math/big"
	"os"
	"path/filepath"
	"strings"

	"github.com/Gurux/gxdlms-go/enums"
)

// List of x509 certificates.
type GXx509CertificateCollection []*GXx509Certificate

// Clear removes all certificates from the collection.
func (g *GXx509CertificateCollection) Clear() {
	*g = (*g)[:0]
}

// Remove removes first occurrence of the specified certificate from the collection.
//
// Parameters:
//
//	value: Certificate to remove.
func (g *GXx509CertificateCollection) Remove(value *GXx509Certificate) {
	for i, v := range *g {
		if v == value {
			*g = append((*g)[:i], (*g)[i+1:]...)
			break
		}
	}
}

// Find returns the find public key certificate by public key.
//
// Parameters:
//
//	key: X509 certificate to search for.
//
// Returns:
//
//	Certificate found or nil, if the certificate is not found.
func (g *GXx509CertificateCollection) Find(key *GXx509Certificate) *GXx509Certificate {
	for _, it := range *g {
		if it.PublicKey.Equals(key) {
			return it
		}
	}
	return nil
}

// FindBySerial returns the find public key certificate by serial number.
//
// Parameters:
//
//	serialNumber: X509 certificate serial number to search for.
//	issuer: X509 certificate issuer.
//
// Returns:
//
//	Found certificate or nil if certificate is not found.
func (g *GXx509CertificateCollection) FindBySerial(serialNumber *big.Int, issuer string) *GXx509Certificate {
	for _, it := range *g {
		if it.SerialNumber.Cmp(serialNumber) == 0 && it.Issuer == issuer {
			return it
		}
	}
	return nil
}

// FindBySystemTitle returns the find public key certificate by system title.
//
// Parameters:
//
//	systemTitle: System title.
//	usage: Key usage.
//
// Returns:
//
//	Found certificate or nil if certificate is not found.
func (g *GXx509CertificateCollection) FindBySystemTitle(systemTitle []byte, usage enums.KeyUsage) *GXx509Certificate {
	var commonName string
	if len(systemTitle) != 0 {
		commonName = Asn1SystemTitleToSubject(systemTitle)
	}
	return g.FindByCommonName(commonName, usage)
}

// GetCertificates returns the find certificates by key usage.
//
// Parameters:
//
//	usage: Key usage.
//
// Returns:
//
//	>Found certificates.
func (g *GXx509CertificateCollection) GetCertificates(usage enums.KeyUsage) []*GXx509Certificate {
	certificates := []*GXx509Certificate{}
	for _, it := range *g {
		if it.KeyUsage == usage {
			certificates = append(certificates, it)
		}
	}
	return certificates
}

// FindByCommonName returns the find public key certificate by common name (CN).
//
// Parameters:
//
//	commonName: Common name.
//	usage: Key usage.
//
// Returns:
//
//	Found certificate or nil if certificate is not found.
func (g *GXx509CertificateCollection) FindByCommonName(commonName string, usage enums.KeyUsage) *GXx509Certificate {
	for _, it := range *g {
		if (usage == enums.KeyUsageNone || (it.KeyUsage&usage) != 0) && strings.Contains(it.Subject, commonName) {
			return it
		}
	}
	return nil
}

// Importx509Collection returns the import certificates from the given folder.
func Importx509Collection(path string) (GXx509CertificateCollection, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	l := GXx509CertificateCollection{}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		full := filepath.Join(path, e.Name())
		ext := strings.ToLower(filepath.Ext(full))
		if ext == ".pem" || ext == ".cer" {
			cert, err := X509CertificateLoad(full)
			if err != nil {
				log.Printf("Failed to load PKCS #8 certificate. %s: %v", full, err)
				continue
			}
			l = append(l, cert)
		}
	}
	return l, nil
}
