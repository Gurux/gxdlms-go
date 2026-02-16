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
	"os"
	"path/filepath"
	"strings"
)

// List of PKCS #8 certificates.
type GXPkcs8Collection []GXPkcs8

// Find returns the find private key certificate by public name.
func (g *GXPkcs8Collection) Find(key *GXPublicKey) *GXPkcs8 {
	for _, it := range *g {
		if it.PublicKey().Equals(key) {
			return &it
		}
	}
	return nil
}

// Pkcs8CollectionImport returns the import private key certificates from the given folder.
func Pkcs8CollectionImport(path string) (GXPkcs8Collection, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	l := GXPkcs8Collection{}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		full := filepath.Join(path, e.Name())
		ext := strings.ToLower(filepath.Ext(full))
		if ext == ".pem" || ext == ".cer" {
			cert, err := Pkcs8Load(full)
			if err != nil {
				log.Printf("Failed to load PKCS #8 certificate. %s: %v", full, err)
				continue
			}
			l = append(l, *cert)
		}
	}
	return l, nil
}
