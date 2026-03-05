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
	"fmt"
	"strings"
	"time"

	"github.com/Gurux/gxcommon-go"
	"github.com/Gurux/gxdlms-go/enums"
	"golang.org/x/text/language"
)

// GXDateTime represents a COSEM date-time value where selected fields can be skipped.
type GXDateTime struct {
	// Used date time value.
	Value time.Time

	// Skip selected date time fields.
	Skip enums.DateTimeSkips

	// Date time extra information.
	Extra enums.DateTimeExtraInfo

	// Day of week.
	DayOfWeek int

	// Status of the clock.
	Status enums.ClockStatus
}

func isNumeric(value byte) bool {
	return value >= '0' && value <= '9'
}

// Available date time formats.
var dateTimeFormats = map[string]string{
	"FI": "02.01.2006 15.04.05.000 -0700",
	"SE": "2006-01-02 15:04:05.000 -0700",
	"US": "01/02/2006 03:04:05.000 AM -0700",
	"DE": "02.01.2006 15:04:05.000 -0700",
	"GB": "02/01/2006 15:04:05.000 -0700",
	"ES": "02/01/2006 15:04:05.000 -0700",
	"ET": "02.01.2006 15:04:05.000 -0700",
	"FR": "02/01/2006 15:04:05.000 -0700",
	"IT": "02/01/2006 15:04:05.000 -0700",
	"HI": "02/01/2006 03:04:05.000 AM -0700",
	"SV": "02/01/2006 15:04:05.000 -0700",
}

// getDateTimeFormat returns a date-time layout for the given language and skipped fields.
func getDateTimeFormat(language *language.Tag, skip enums.DateTimeSkips) string {
	if language == nil {
		l := gxcommon.CurrentLanguage()
		language = &l
	}
	region, _ := language.Region()
	format, ok := dateTimeFormats[region.String()]
	if !ok {
		format = dateTimeFormats["US"]
	}

	var timeSeparator, dateSeparator string
	detectSeparators(format, &timeSeparator, &dateSeparator)
	if (skip & enums.DateTimeSkipsYear) != 0 {
		//Remove year.
		format = remove(format, "2006", dateSeparator)
		//Remove time zone if year is removed.
		format = remove(format, "-0700", "")
	}
	if (skip & enums.DateTimeSkipsMonth) != 0 {
		//Remove month.
		format = remove(format, "01", dateSeparator)
		format = remove(format, "-0700", "")
	}
	if (skip & enums.DateTimeSkipsDay) != 0 {
		//remove day.
		format = remove(format, "02", dateSeparator)
		format = remove(format, "-0700", "")
	}
	if (skip & enums.DateTimeSkipsHour) != 0 {
		//Hour can be 15 or 03. Remove both.
		format = remove(format, "15", timeSeparator)
		format = remove(format, "03", timeSeparator)
		format = remove(format, "tt", timeSeparator)
		format = remove(format, "-0700", "")
	}
	if (skip & enums.DateTimeSkipsMs) != 0 {
		format = strings.ReplaceAll(format, ".000", "")
	} else if !strings.Contains(format, ".000") {
		format = strings.ReplaceAll(format, "05", "05.000")
	}
	if (skip & enums.DateTimeSkipsSecond) != 0 {
		format = remove(format, "05", timeSeparator)
	} else if !strings.Contains(format, "05") {
		format = strings.ReplaceAll(format, "05", "04"+timeSeparator+"05")
	}
	if (skip & enums.DateTimeSkipsMinute) != 0 {
		format = remove(format, "04", timeSeparator)
		format = remove(format, "-0700", "")
	}
	format = strings.TrimSpace(format)
	return format
}

// timeZonePosition returns the start index of a trailing timezone offset (+/-HHMM), or -1.
func timeZonePosition(value string) int {
	if len(value) > 5 {
		pos := len(value) - 5
		sep := value[pos]
		if sep == '-' || sep == '+' {
			return pos
		}
	}
	return -1
}

func remove(value string, tag string, sep string) string {
	if sep != "" {
		if strings.Contains(value, tag+sep) {
			value = strings.ReplaceAll(value, tag+sep, "")
			return value
		} else if strings.Contains(value, sep+tag) {
			value = strings.ReplaceAll(value, sep+tag, "")
			return value
		}
	}
	value = strings.ReplaceAll(value, tag, "")
	return value
}

// String returns the date-time as a localized string using local time.
func (g *GXDateTime) String() string {
	return g.ToString(nil, true)
}

func isDateSeparator(b byte) bool {
	return b == '.' || b == '/' || b == '-'
}

// detectSeparators detects date and time separators from the given layout and sets them to the provided pointers.
func detectSeparators(layout string, timeSeparator *string, dateSeparator *string) {
	// Detect date separator
	datePatterns := []string{"2006", "02", "01"}
	for _, p := range datePatterns {
		if idx := strings.Index(layout, p); idx != -1 {
			// look right after pattern
			end := idx + len(p)
			if end < len(layout) {
				c := layout[end]
				if isDateSeparator(c) {
					*dateSeparator = string(c)
					break
				}
			}
		}
	}

	// Detect time separator
	timePatterns := []string{"15", "03", "04", "05"}
	for _, p := range timePatterns {
		if idx := strings.Index(layout, p); idx != -1 {
			end := idx + len(p)
			if end < len(layout) {
				c := layout[end]
				if c == ':' || c == '.' {
					*timeSeparator = string(c)
					break
				}
			}
		}
	}
}

// ToString returns the date-time formatted for the given language.
// If useLocalTime is true, the value is formatted using the local timezone.
func (g *GXDateTime) ToString(language *language.Tag, useLocalTime bool) string {
	return toString(g, language, useLocalTime, false)
}

func toString(target any, language *language.Tag, useLocalTime bool, useFormat bool) string {
	//Get current date time format.
	var skip, formatSkip enums.DateTimeSkips
	var value time.Time
	if v, ok := target.(*GXDateTime); ok {
		value = v.Value
		skip = v.Skip
		//Don't show ms if it's not used.
		if (skip&enums.DateTimeSkipsMs) != 0 || value.Nanosecond()/int(time.Millisecond) == 0 {
			formatSkip |= enums.DateTimeSkipsMs
		}
	} else if v, ok := target.(*GXDate); ok {
		value = v.Value
		skip = v.Skip
		if useFormat {
			formatSkip = enums.DateTimeSkipsHour | enums.DateTimeSkipsMinute | enums.DateTimeSkipsSecond | enums.DateTimeSkipsDeviation | enums.DateTimeSkipsMs
		} else {
			formatSkip = skip
		}
	} else if v, ok := target.(*GXTime); ok {
		value = v.Value
		skip = v.Skip
		if useFormat {
			formatSkip = enums.DateTimeSkipsYear | enums.DateTimeSkipsMonth | enums.DateTimeSkipsDay | enums.DateTimeSkipsDayOfWeek
		} else {
			formatSkip = skip
		}
		//Don't show ms if it's not used.
		if (skip&enums.DateTimeSkipsMs) != 0 || value.Nanosecond()/int(time.Millisecond) == 0 {
			formatSkip |= enums.DateTimeSkipsMs
		}
	}
	format := getDateTimeFormat(language, formatSkip)
	if value.Hour() > 11 {
		format = strings.ReplaceAll(format, "AM", "PM")
	}
	if useLocalTime {
		//Remove time zone if year is removed.
		format = remove(format, "-0700", "")
		format = strings.TrimSpace(format)
	}
	if skip != enums.DateTimeSkipsNone {
		if (skip & enums.DateTimeSkipsYear) != 0 {
			format = strings.ReplaceAll(format, "2006", "*")
		}
		if (skip & enums.DateTimeSkipsMonth) != 0 {
			format = strings.ReplaceAll(format, "01", "*")
		}
		if (skip & enums.DateTimeSkipsDay) != 0 {
			format = strings.ReplaceAll(format, "02", "*")
		}
		if (skip & enums.DateTimeSkipsHour) != 0 {
			format = strings.ReplaceAll(format, "15", "*")
		}
		if (skip & enums.DateTimeSkipsMinute) != 0 {
			format = strings.ReplaceAll(format, "04", "*")
		}
		if (skip & enums.DateTimeSkipsSecond) != 0 {
			format = strings.ReplaceAll(format, "05", "*")
		}
		if useLocalTime {
			return value.Local().Format(format)
		}
		ret := value.Format(format)
		_, offset := value.Zone()
		if offset == 0 {
			ret = ret[:len(ret)-6] + "Z"
		}
		return ret
	}
	if useLocalTime {
		return value.Local().Format(format)
	}
	return value.Format(format)
}

// ToFormatString returns the date-time using a full fixed output format.
// If useLocalTime is true, the value is formatted using the local timezone.
func (g *GXDateTime) ToFormatString(language *language.Tag, useLocalTime bool) string {
	return toString(g, language, useLocalTime, true)
}

// ToFormatMeterString returns the date-time using meter timezone formatting.
func (g *GXDateTime) ToFormatMeterString(language *language.Tag) string {
	return toString(g, language, false, true)
}

// GXDateTimeFromUnixTime creates a GXDateTime from Unix epoch seconds.
func GXDateTimeFromUnixTime(unixTime int64) *GXDateTime {
	return &GXDateTime{
		Value: time.Unix(unixTime, 0),
	}
}

// GXDateTimeFromHighResolutionTime creates a GXDateTime from high-resolution clock seconds.
func GXDateTimeFromHighResolutionTime(highResolution int64) *GXDateTime {
	return &GXDateTime{
		Value: time.Unix(int64(highResolution), 0),
	}
}

// ToUnixTime converts the date-time to Unix epoch seconds.
func (g *GXDateTime) ToUnixTime() int64 {
	return g.Value.Unix()
}

// ToHighResolutionTime converts the date-time to high-resolution clock seconds.
func (g *GXDateTime) ToHighResolutionTime() uint64 {
	if g.Value.IsZero() {
		return 0
	}
	return uint64(g.Value.Unix())
}

// ToHex returns the date-time as a hexadecimal string.
// If addSpace is true, spaces are added between bytes.
// If useMeterTimeZone is true, meter timezone is used instead of local PC timezone.
func (g *GXDateTime) ToHex(addSpace bool, useMeterTimeZone bool) string {
	buff := GXByteBuffer{}
	return buff.ToHexByIndex(addSpace, 0)
}

func getDateTimeToken(format string, index int) (string, enums.DateTimeSkips) {
	if index < 0 || index >= len(format) {
		return "", enums.DateTimeSkipsNone
	}
	tokens := []struct {
		token string
		field enums.DateTimeSkips
	}{
		{"-0700", enums.DateTimeSkipsDeviation},
		{"2006", enums.DateTimeSkipsYear},
		{".000", enums.DateTimeSkipsMs},
		{"15", enums.DateTimeSkipsHour},
		{"03", enums.DateTimeSkipsHour},
		{"04", enums.DateTimeSkipsMinute},
		{"05", enums.DateTimeSkipsSecond},
		{"01", enums.DateTimeSkipsMonth},
		{"02", enums.DateTimeSkipsDay},
	}
	for _, it := range tokens {
		if strings.HasPrefix(format[index:], it.token) {
			return it.token, it.field
		}
	}
	return "", enums.DateTimeSkipsNone
}

func updateSingleDigitParts(value string, layout string) string {
	v := strings.FieldsFunc(value, func(r rune) bool {
		return r == '/' || r == '.' || r == '-' || r == ':' || r == ' '
	})

	l := strings.FieldsFunc(layout, func(r rune) bool {
		return r == '/' || r == '.' || r == '-' || r == ':' || r == ' '
	})

	for i := 0; i < len(v) && i < len(l); i++ {
		part := l[i]
		switch part {
		case "02", "01", "03", "04", "05":
			if len(v[i]) != len(part) {
				layout = strings.ReplaceAll(layout, l[i], l[i][1:])
			}
		}
	}
	return layout
}

// parseInternal parses a date-time string into GXDateTime, GXDate, or GXTime targets.
func parseInternal(target any, value string, language *language.Tag) error {
	addTimeZone := true
	var g *GXDateTime
	//Get current date time format.
	var skip enums.DateTimeSkips
	if v, ok := target.(*GXDateTime); ok {
		g = v
	} else if v, ok := target.(*GXDate); ok {
		g = &v.GXDateTime
		skip = v.Skip
	} else if v, ok := target.(*GXTime); ok {
		g = &v.GXDateTime
		skip = v.Skip
	}
	g.DayOfWeek = 0xFF
	if value != "" {
		format := getDateTimeFormat(language, skip)
		if strings.Contains(value, "PM") {
			format = strings.ReplaceAll(format, "AM", "PM")
		}
		var timeSeparator, dateSeparator string
		detectSeparators(format, &timeSeparator, &dateSeparator)
		if strings.Contains(value, "BEGIN") {
			g.Extra |= enums.DateTimeExtraInfoDstBegin
			value = strings.ReplaceAll(value, "BEGIN", "01")
		}
		if strings.Contains(value, "END") {
			g.Extra |= enums.DateTimeExtraInfoDstEnd
			value = strings.ReplaceAll(value, "END", "01")
		}
		if strings.Contains(value, "LASTDAY2") {
			g.Extra |= enums.DateTimeExtraInfoLastDay2
			value = strings.ReplaceAll(value, "LASTDAY2", "01")
		}
		if strings.Contains(value, "LASTDAY") {
			g.Extra |= enums.DateTimeExtraInfoLastDay
			value = strings.ReplaceAll(value, "LASTDAY", "01")
		}

		v := value
		if strings.IndexByte(value, '*') != -1 {
			//Day of week is not supported when date time is give as a string.
			g.Skip |= enums.DateTimeSkipsDayOfWeek
			lastFormatIndex := -1
			offset := 0
			for pos := 0; pos < len(value); pos++ {
				c := value[pos]
				if !isNumeric(c) {
					if c == '*' {
						token, field := getDateTimeToken(format, lastFormatIndex+1)
						if token == "" {
							return fmt.Errorf("invalid date time format")
						}
						val := "1" + strings.Repeat("0", len(token)-1)
						switch field {
						case enums.DateTimeSkipsYear:
							val = "2006"
							addTimeZone = false
							g.Skip |= enums.DateTimeSkipsYear
						case enums.DateTimeSkipsMonth:
							addTimeZone = false
							g.Skip |= enums.DateTimeSkipsMonth
						case enums.DateTimeSkipsDay:
							addTimeZone = false
							g.Skip |= enums.DateTimeSkipsDay
						case enums.DateTimeSkipsHour:
							addTimeZone = false
							g.Skip |= enums.DateTimeSkipsHour
						case enums.DateTimeSkipsMinute:
							addTimeZone = false
							g.Skip |= enums.DateTimeSkipsMinute
						case enums.DateTimeSkipsSecond:
							g.Skip |= enums.DateTimeSkipsSecond
						default:
							return fmt.Errorf("invalid date time format")
						}
						v = v[:pos+offset] + val + value[pos+1:]
						offset += len(val) - 1
					} else {
						start := lastFormatIndex + 1
						tmp := strings.IndexByte(format[start:], c)
						if tmp != -1 {
							lastFormatIndex = start + tmp
						} else {
							lastFormatIndex = -1
						}
						//Dot is used time separator in some countries.
						if lastFormatIndex == -1 && c == byte(timeSeparator[0]) {
							tmp = strings.IndexByte(format[start:], '.')
							if tmp != -1 {
								lastFormatIndex = start + tmp
							}
						}
					}
				}
			}
		}
		format = updateSingleDigitParts(value, format)
		// If time zone is used.
		pos := timeZonePosition(value)
		if !addTimeZone || pos == -1 {
			g.Skip |= enums.DateTimeSkipsDeviation
			format = remove(format, "-0700", timeSeparator)
			// Trim
			format = strings.TrimSpace(format)
			if language == nil {
				g.Skip |= enums.DateTimeSkipsDeviation
			}
		}
		parseWithLayout := func(layout string) (time.Time, error) {
			if language == nil {
				return time.Parse(layout, v)
			}
			return time.ParseInLocation(layout, v, time.Local)
		}
		var err error
		g.Value, err = parseWithLayout(format)
		if err != nil && g.Skip&enums.DateTimeSkipsMs == 0 {
			//Remove ms.
			g.Skip |= enums.DateTimeSkipsMs
			format = remove(format, ".000", "")
			// Trim
			format = strings.TrimSpace(format)
			g.Value, err = parseWithLayout(format)
			if err != nil && g.Skip&enums.DateTimeSkipsSecond == 0 {
				//Remove seconds.
				format = remove(format, "05", timeSeparator)
				// Trim
				format = strings.TrimSpace(format)
				g.Value, err = parseWithLayout(format)
				if err == nil {
					g.Skip |= enums.DateTimeSkipsMs | enums.DateTimeSkipsSecond
				}
			} else {
				g.Skip |= enums.DateTimeSkipsMs
			}
		}
		if err != nil {
			g.Value = time.Date(2006, 1, 2, 3, 4, 5, 0, time.Local)
			expected := g.ToFormatString(language, true)
			expected = strings.ReplaceAll(expected, "2006", "yyyy")
			expected = strings.ReplaceAll(expected, "01", "MM")
			expected = strings.ReplaceAll(expected, "02", "dd")
			expected = strings.ReplaceAll(expected, "03", "hh")
			expected = strings.ReplaceAll(expected, "04", "mm")
			expected = strings.ReplaceAll(expected, "05", "ss")
			return fmt.Errorf("parsing '%s' failed. Expected format: '%s'", value, expected)
		}
		if strings.Contains(value, "AM") {
			g.Value = g.Value.Add(-12 * time.Hour)
		}
		g.Skip |= enums.DateTimeSkipsDayOfWeek
		if (g.Skip & (enums.DateTimeSkipsYear | enums.DateTimeSkipsMonth | enums.DateTimeSkipsDay | enums.DateTimeSkipsHour | enums.DateTimeSkipsMinute)) == 0 {
			if g.Value.IsDST() {
				g.Status |= enums.ClockStatusDaylightSavingActive
			}
		}
	}
	return nil
}

// NewGXDateTimeFromString creates a GXDateTime by parsing the given string.
func NewGXDateTimeFromString(value string, language *language.Tag) (*GXDateTime, error) {
	g := &GXDateTime{}
	err := parseInternal(g, value, language)
	if err != nil {
		return nil, err
	}
	return g, nil
}

// NewGXDateTimeFromTime creates a GXDateTime from time.Time.
func NewGXDateTimeFromTime(value time.Time) *GXDateTime {
	return &GXDateTime{Value: value}
}
