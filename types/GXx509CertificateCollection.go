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

// GXx509CertificateCollection is a helper slice type for working with multiple X.509 certificates.
//
// It provides convenience lookup and filtering methods for common selection criteria,
// such as serial number, common name, key usage, or system title.
type GXx509CertificateCollection []*GXx509Certificate

// Clear removes all certificates from the collection.
func (g *GXx509CertificateCollection) Clear() {
	*g = (*g)[:0]
}

// Remove deletes the first certificate from the collection that is pointer-equal to the given value.
//
// The comparison is pointer-based (it checks for the same certificate instance).
func (g *GXx509CertificateCollection) Remove(value *GXx509Certificate) {
	for i, v := range *g {
		if v == value {
			*g = append((*g)[:i], (*g)[i+1:]...)
			break
		}
	}
}

// Find returns the first certificate whose public key matches the provided certificate.
//
// Parameters:
//
//	key: An X.509 certificate whose public key is used for matching.
//
// Returns:
//
//	The matching certificate, or nil if none is found.
func (g *GXx509CertificateCollection) Find(key *GXx509Certificate) *GXx509Certificate {
	for _, it := range *g {
		if it.PublicKey.Curve == key.PublicKey.Curve &&
			it.PublicKey.X.Cmp(key.PublicKey.X) == 0 &&
			it.PublicKey.Y.Cmp(key.PublicKey.Y) == 0 {
			return it
		}
	}
	return nil
}

// FindBySerial returns the first certificate matching the given serial number and issuer.
//
// Parameters:
//
//	serialNumber: X.509 certificate serial number to search for.
//	issuer: X.509 certificate issuer.
//
// Returns:
//
//	The matching certificate, or nil if none is found.
func (g *GXx509CertificateCollection) FindBySerial(serialNumber *big.Int, issuer string) *GXx509Certificate {
	for _, it := range *g {
		if it.SerialNumber.Cmp(serialNumber) == 0 && it.Issuer == issuer {
			return it
		}
	}
	return nil
}

// FindBySystemTitle returns the first certificate whose subject matches the provided
// system title (converted to a common name) and whose key usage matches the provided usage.
//
// Parameters:
//
//	systemTitle: ASN.1 system title.
//	usage: Key usage.
//
// Returns:
//
//	The matching certificate, or nil if none is found.
func (g *GXx509CertificateCollection) FindBySystemTitle(systemTitle []byte, usage enums.KeyUsage) *GXx509Certificate {
	var commonName string
	if len(systemTitle) != 0 {
		commonName = Asn1SystemTitleToSubject(systemTitle)
	}
	return g.FindByCommonName(commonName, usage)
}

// GetCertificates returns all certificates in the collection that match the specified key usage.
//
// Parameters:
//
//	usage: Key usage.
//
// Returns:
//
//	Slice of matching certificates (empty if none match).
func (g *GXx509CertificateCollection) GetCertificates(usage enums.KeyUsage) []*GXx509Certificate {
	certificates := []*GXx509Certificate{}
	for _, it := range *g {
		if it.KeyUsage == usage {
			certificates = append(certificates, it)
		}
	}
	return certificates
}

// FindByCommonName returns the first certificate whose subject contains the given
// common name and that matches the specified key usage.
//
// Parameters:
//
//	commonName: Common name substring to match.
//	usage: Key usage.
//
// Returns:
//
//	The matching certificate, or nil if none is found.
func (g *GXx509CertificateCollection) FindByCommonName(commonName string, usage enums.KeyUsage) *GXx509Certificate {
	for _, it := range *g {
		if (usage == enums.KeyUsageNone || (it.KeyUsage&usage) != 0) && strings.Contains(it.Subject, commonName) {
			return it
		}
	}
	return nil
}

// Importx509Collection reads X.509 certificates from the specified directory.
//
// Supported file extensions are .pem and .cer. Files that fail to load are skipped and logged.
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
