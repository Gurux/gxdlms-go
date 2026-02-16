package internal

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
	"errors"
	"strconv"
	"strings"
	"unicode"
)

// GetValues returns the get values to count together.
func getValues(expressions string) []string {
	var values []string
	last := 0
	for i, ch := range expressions {
		switch ch {
		case '%', '+', '-', '*', '/':
			values = append(values, expressions[last:i])
			values = append(values, string(ch))
			last = i + 1
		}
	}
	if last != len(expressions) {
		values = append(values, expressions[last:])
	}
	return values
}

func getValue(value string, sn uint64) (uint64, error) {
	if value == "sn" {
		return sn, nil
	}
	return strconv.ParseUint(value, 10, 64)
}

// FormatString returns the produce formatted string by the given math expression.
//
// Parameters:
//
//	expression: Unformatted math expression.
//
// Returns:
//
//	Formatted math expression.
func formatString(expression string) (string, error) {
	if expression == "" {
		return "", errors.New("expression is nil or empty")
	}
	var sb strings.Builder
	for _, ch := range expression {
		if ch == '(' || ch == ')' {
			return "", errors.New("invalid serial number formula")
		}
		if unicode.IsSpace(ch) {
			continue
		}
		if unicode.IsUpper(ch) {
			sb.WriteRune(unicode.ToLower(ch))
		} else {
			sb.WriteRune(ch)
		}
	}
	return sb.String(), nil
}

// SerialnumberCounterCount returns the count of the given serial number and formula.
//
// Parameters:
//
//	sn: Serial number.
//	formula: Formula to count.
func SerialnumberCounterCount(sn uint64, formula string) (uint64, error) {
	ret, err := formatString(formula)
	if err != nil {
		return 0, err
	}
	values := getValues(ret)
	if len(values)%2 == 0 {
		return 0, errors.New("invalid serial number formula")
	}

	var total uint64
	value, err := getValue(values[0], sn)
	if err != nil {
		return 0, err
	}

	for i := 1; i != len(values); i += 2 {
		nextToken := values[i+1]
		if nextToken == "sn" {
			total += value
			value = 0
		}
		rhs, err := getValue(nextToken, sn)
		if err != nil {
			return 0, err
		}

		switch values[i] {
		case "%":
			if rhs == 0 {
				return 0, errors.New("division by zero in formula (mod)")
			}
			value = value % rhs
		case "+":
			value = value + rhs
		case "-":
			value = value - rhs
		case "*":
			value = value * rhs
		case "/":
			if rhs == 0 {
				return 0, errors.New("division by zero in formula (div)")
			}
			value = value / rhs
		default:
			return 0, errors.New("invalid serial number formula")
		}
	}
	total += value
	return total, nil
}
