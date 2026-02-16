package dlms

// --------------------------------------------------------------------------
//
//	Gurux Ltd
//
// Filename:        $HeadURL$
//
// Version:         $Revision$,
//
//	$Date$
//	$Author$
//
// # Copyright (c) Gurux Ltd
//
// ---------------------------------------------------------------------------
//
//	DESCRIPTION
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
// ---------------------------------------------------------------------------

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/types"
)

type GXStandardObisCodeCollection []GXStandardObisCode

// GetBytes converts logical name "A.B.C.D.E.F" to 6 bytes.
// If not dot-separated with 6 parts, tries to treat input as hex.
func getBytes(ln string) ([]byte, error) {
	if ln == "" {
		return nil, nil
	}
	tmp := strings.Split(ln, ".")
	if len(tmp) != 6 {
		// If value is given as hex.
		tmp2 := types.HexToBytes(ln)
		if len(tmp2) == 6 {
			return []byte{tmp2[0], tmp2[1], tmp2[2], tmp2[3], tmp2[4], tmp2[5]}, nil
		}
		return nil, errors.New("invalid OBIS code")
	}
	code := make([]byte, 6)
	for i := 0; i < 6; i++ {
		v, err := strconv.ParseUint(strings.TrimSpace(tmp[i]), 10, 8)
		if err != nil {
			return nil, fmt.Errorf("invalid OBIS part %q: %w", tmp[i], err)
		}
		code[i] = byte(v)
	}
	return code, nil
}

func equalsInterface(it GXStandardObisCode, ic int) bool {
	// If all interfaces are allowed.
	if ic == 0 || it.Interfaces == "*" {
		return true
	}
	for _, p := range strings.Split(it.Interfaces, ",") {
		if strings.TrimSpace(p) == strconv.Itoa(ic) {
			return true
		}
	}
	return false
}

// equalsMask checks if obisMask matches one byte ic.
// obisMask can be "1", "1,2,3", "1-10", or "&" special.
func equalsByteMask(obisMask string, ic byte) bool {
	obisMask = strings.TrimSpace(obisMask)

	// list "1,2,3" (can include ranges)
	if strings.Contains(obisMask, ",") {
		for _, part := range strings.Split(obisMask, ",") {
			part = strings.TrimSpace(part)
			if strings.Contains(part, "-") {
				if equalsByteMask(part, ic) {
					return true
				}
			} else {
				v, err := strconv.ParseUint(part, 10, 8)
				if err == nil && byte(v) == ic {
					return true
				}
			}
		}
		return false
	}

	// range "a-b"
	if strings.Contains(obisMask, "-") {
		tmp := strings.SplitN(obisMask, "-", 2)
		a, err1 := strconv.ParseUint(strings.TrimSpace(tmp[0]), 10, 8)
		b, err2 := strconv.ParseUint(strings.TrimSpace(tmp[1]), 10, 8)
		if err1 != nil || err2 != nil {
			return false
		}
		return ic >= byte(a) && ic <= byte(b)
	}

	// special "&"
	if obisMask == "&" {
		return ic == 0 || ic == 1 || ic == 7
	}

	v, err := strconv.ParseUint(obisMask, 10, 8)
	if err != nil {
		return false
	}
	return byte(v) == ic
}

func EqualsMaskLN(obisMask, ln string) (bool, error) {
	ic, err := getBytes(ln)
	if err != nil {
		return false, err
	}
	return equalsObisCode(strings.Split(obisMask, "."), ic)
}

func equalsMask(obisMask string, ln string) bool {
	ret, err := getBytes(ln)
	if err != nil {
		return false
	}
	ret2, err := equalsObisCode(strings.Split(obisMask, "."), ret)
	if err != nil {
		return false
	}
	return ret2
}

// equalsObisCode checks six parts.
func equalsObisCode(obisMaskParts []string, ic []byte) (bool, error) {
	if ic == nil {
		return true, nil
	}
	if len(obisMaskParts) != 6 {
		return false, errors.New("invalid OBIS mask")
	}
	for i := 0; i < 6; i++ {
		if !equalsByteMask(obisMaskParts[i], ic[i]) {
			return false, nil
		}
	}
	return true, nil
}

func getN1CDescription(str string) string {
	if str == "" || str[0] != '$' {
		return ""
	}
	value, err := strconv.Atoi(str[1:])
	if err != nil {
		return ""
	}
	switch value {
	case 41:
		return "Absolute temperature"
	case 42:
		return "Absolute pressure"
	case 44:
		return "Velocity of sound"
	case 45:
		return "Density(of gas)"
	case 46:
		return "Relative density"
	case 47:
		return "Gauge pressure"
	case 48:
		return "Differential pressure"
	case 49:
		return "Density of air"
	default:
		return ""
	}
}

func getDescription(str string) string {
	if str == "" || str[0] != '$' {
		return ""
	}
	value, err := strconv.Atoi(str[1:])
	if err != nil {
		return ""
	}
	switch value {
	case 1:
		return "Sum Li Active power+ (QI+QIV)"
	case 2:
		return "Sum Li Active power- (QII+QIII)"
	case 3:
		return "Sum Li Reactive power+ (QI+QII)"
	case 4:
		return "Sum Li Reactive power- (QIII+QIV)"
	case 5:
		return "Sum Li Reactive power QI"
	case 6:
		return "Sum Li Reactive power QII"
	case 7:
		return "Sum Li Reactive power QIII"
	case 8:
		return "Sum Li Reactive power QIV"
	case 9:
		return "Sum Li Apparent power+ (QI+QIV)"
	case 10:
		return "Sum Li Apparent power- (QII+QIII)"
	case 11:
		return "Current: any phase"
	case 12:
		return "Voltage: any phase"
	case 13:
		return "Sum Li Power factor"
	case 14:
		return "Supply frequency"
	case 15:
		return "Sum Li Active power (abs(QI+QIV)+abs(QII+QIII))"
	case 16:
		return "Sum Li Active power (abs(QI+QIV)-abs(QII+QIII))"
	case 17:
		return "Sum Li Active power QI"
	case 18:
		return "Sum Li Active power QII"
	case 19:
		return "Sum Li Active power QIII"
	case 20:
		return "Sum Li Active power QIV"
	case 21:
		return "L1 Active power+ (QI+QIV)"
	case 22:
		return "L1 Active power- (QII+QIII)"
	case 23:
		return "L1 Reactive power+ (QI+QII)"
	case 24:
		return "L1 Reactive power- (QIII+QIV)"
	case 25:
		return "L1 Reactive power QI"
	case 26:
		return "L1 Reactive power QII"
	case 27:
		return "L1 Reactive power QIII"
	case 28:
		return "L1 Reactive power QIV"
	case 29:
		return "L1 Apparent power+ (QI+QIV)"
	case 30:
		return "L1 Apparent power- (QII+QIII)"
	case 31:
		return "L1 Current"
	case 32:
		return "L1 Voltage"
	case 33:
		return "L1 Power factor"
	case 34:
		return "L1 Supply frequency"
	case 35:
		return "L1 Active power (abs(QI+QIV)+abs(QII+QIII))"
	case 36:
		return "L1 Active power (abs(QI+QIV)-abs(QII+QIII))"
	case 37:
		return "L1 Active power QI"
	case 38:
		return "L1 Active power QII"
	case 39:
		return "L1 Active power QIII"
	case 40:
		return "L1 Active power QIV"
	case 41:
		return "L2 Active power+ (QI+QIV)"
	case 42:
		return "L2 Active power- (QII+QIII)"
	case 43:
		return "L2 Reactive power+ (QI+QII)"
	case 44:
		return "L2 Reactive power- (QIII+QIV)"
	case 45:
		return "L2 Reactive power QI"
	case 46:
		return "L2 Reactive power QII"
	case 47:
		return "L2 Reactive power QIII"
	case 48:
		return "L2 Reactive power QIV"
	case 49:
		return "L2 Apparent power+ (QI+QIV)"
	case 50:
		return "L2 Apparent power- (QII+QIII)"
	case 51:
		return "L2 Current"
	case 52:
		return "L2 Voltage"
	case 53:
		return "L2 Power factor"
	case 54:
		return "L2 Supply frequency"
	case 55:
		return "L2 Active power (abs(QI+QIV)+abs(QII+QIII))"
	case 56:
		return "L2 Active power (abs(QI+QIV)-abs(QI+QIII))"
	case 57:
		return "L2 Active power QI"
	case 58:
		return "L2 Active power QII"
	case 59:
		return "L2 Active power QIII"
	case 60:
		return "L2 Active power QIV"
	case 61:
		return "L3 Active power+ (QI+QIV)"
	case 62:
		return "L3 Active power- (QII+QIII)"
	case 63:
		return "L3 Reactive power+ (QI+QII)"
	case 64:
		return "L3 Reactive power- (QIII+QIV)"
	case 65:
		return "L3 Reactive power QI"
	case 66:
		return "L3 Reactive power QII"
	case 67:
		return "L3 Reactive power QIII"
	case 68:
		return "L3 Reactive power QIV"
	case 69:
		return "L3 Apparent power+ (QI+QIV)"
	case 70:
		return "L3 Apparent power- (QII+QIII)"
	case 71:
		return "L3 Current"
	case 72:
		return "L3 Voltage"
	case 73:
		return "L3 Power factor"
	case 74:
		return "L3 Supply frequency"
	case 75:
		return "L3 Active power (abs(QI+QIV)+abs(QII+QIII))"
	case 76:
		return "L3 Active power (abs(QI+QIV)-abs(QI+QIII))"
	case 77:
		return "L3 Active power QI"
	case 78:
		return "L3 Active power QII"
	case 79:
		return "L3 Active power QIII"
	case 80:
		return "L3 Active power QIV"
	case 82:
		return "Unitless quantities (pulses or pieces)"
	case 84:
		return "Sum Li Power factor-"
	case 85:
		return "L1 Power factor-"
	case 86:
		return "L2 Power factor-"
	case 87:
		return "L3 Power factor-"
	case 88:
		return "Sum Li A2h QI+QII+QIII+QIV"
	case 89:
		return "Sum Li V2h QI+QII+QIII+QIV"
	case 90:
		return "SLi current (algebraic sum of the - unsigned - value of the currents in all phases)"
	case 91:
		return "Lo Current (neutral)"
	case 92:
		return "Lo Voltage (neutral)"
	default:
		return ""
	}
}

func getObisValue(formula string, value int) (string, error) {
	if len(formula) == 1 {
		return strconv.Itoa(value), nil
	}
	add, err := strconv.Atoi(formula[1:])
	if err != nil {
		return "", err
	}
	return strconv.Itoa(value + add), nil
}

// Find by string obisCode.
func (g *GXStandardObisCodeCollection) Find(obisCode string, objectType enums.ObjectType, standard enums.Standard) ([]GXStandardObisCode, error) {
	b, err := getBytes(obisCode)
	if err != nil {
		return nil, err
	}
	return g.FindBytes(b, int(objectType), standard)
}

// FindBytes is the main find (C# Find(byte[] obisCode, int IC, Standard standard)).
func (g *GXStandardObisCodeCollection) FindBytes(obisCode []byte, IC int, standard enums.Standard) ([]GXStandardObisCode, error) {
	_ = standard // not used in the pasted snippet; keep for signature compatibility.

	var list []GXStandardObisCode

	for _, it := range *g {
		// Interface is tested first because it's faster.
		okIntf := equalsInterface(it, IC)
		okObis, err := equalsObisCode(it.OBIS, obisCode)
		if err != nil {
			return nil, err
		}
		if okIntf && okObis {
			tmp := *NewGXStandardObisCode(it.OBIS, it.Description, it.Interfaces, it.DataType)
			tmp.UIDataType = it.UIDataType
			list = append(list, tmp)

			// Special description replacement after ';'
			tmp2 := strings.Split(it.Description, ";")
			if len(tmp2) > 1 {
				desc := ""
				if obisCode != nil && strings.TrimSpace(tmp2[1]) == "$1" {
					if obisCode[0] == 7 {
						desc = getN1CDescription("$" + strconv.Itoa(int(obisCode[2])))
					} else {
						desc = getDescription("$" + strconv.Itoa(int(obisCode[2])))
					}
				}
				if desc != "" {
					tmp2[1] = desc
					list[len(list)-1].Description = strings.Join(tmp2, ";")
				}
			}

			// Replace $A..$F and fill actual OBIS bytes.
			if obisCode != nil {
				last := &list[len(list)-1]
				for i := 0; i < 6; i++ {
					last.OBIS[i] = strconv.Itoa(int(obisCode[i]))
				}

				last.Description = strings.ReplaceAll(last.Description, "$A", last.OBIS[0])
				last.Description = strings.ReplaceAll(last.Description, "$B", last.OBIS[1])
				last.Description = strings.ReplaceAll(last.Description, "$C", last.OBIS[2])
				last.Description = strings.ReplaceAll(last.Description, "$D", last.OBIS[3])
				last.Description = strings.ReplaceAll(last.Description, "$E", last.OBIS[4])
				last.Description = strings.ReplaceAll(last.Description, "$F", last.OBIS[5])

				// Increase value: parse "$( ... )" block like in C#
				begin := strings.Index(last.Description, "$(")
				if begin != -1 {
					rest := last.Description[begin+2:]
					// emulate: Split(new[]{'(',')','$'}, RemoveEmptyEntries)
					parts := splitByAny(rest, "()", "$")
					last.Description = last.Description[:begin]

					for _, v := range parts {
						if v == "" {
							continue
						}
						switch v[0] {
						case 'A':
							s, err := getObisValue(v, int(obisCode[0]))
							if err != nil {
								return nil, err
							}
							last.Description += s
						case 'B':
							s, err := getObisValue(v, int(obisCode[1]))
							if err != nil {
								return nil, err
							}
							last.Description += s
						case 'C':
							s, err := getObisValue(v, int(obisCode[2]))
							if err != nil {
								return nil, err
							}
							last.Description += s
						case 'D':
							s, err := getObisValue(v, int(obisCode[3]))
							if err != nil {
								return nil, err
							}
							last.Description += s
						case 'E':
							s, err := getObisValue(v, int(obisCode[4]))
							if err != nil {
								return nil, err
							}
							last.Description += s
						case 'F':
							s, err := getObisValue(v, int(obisCode[5]))
							if err != nil {
								return nil, err
							}
							last.Description += s
						default:
							last.Description += v
						}
					}
				}

				// Replace ';' with space, collapse double spaces
				last.Description = strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(last.Description, ";", " "), "  ", " "))
			}
		}
	}

	// If invalid OBIS code.
	if len(list) == 0 {
		tmp := *NewGXStandardObisCode(nil, "Invalid", strconv.Itoa(IC), "")
		if len(obisCode) == 6 {
			tmp.OBIS = make([]string, 6)
			for i := 0; i < 6; i++ {
				tmp.OBIS[i] = strconv.Itoa(int(obisCode[i]))
			}
		}
		list = append(list, tmp)
	}

	return list, nil
}

func splitByAny(s string, seps1 string, seps2 string) []string {
	// Splits by any rune in seps1 or seps2 (we pass "()" and "$").
	isSep := func(r rune) bool {
		return strings.ContainsRune(seps1, r) || strings.ContainsRune(seps2, r)
	}
	var out []string
	var b strings.Builder
	for _, r := range s {
		if isSep(r) {
			if b.Len() > 0 {
				out = append(out, b.String())
				b.Reset()
			}
			continue
		}
		b.WriteRune(r)
	}
	if b.Len() > 0 {
		out = append(out, b.String())
	}
	return out
}
