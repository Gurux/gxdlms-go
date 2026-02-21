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
	"os"
	"strings"
	"time"

	"github.com/Gurux/gxdlms-go/enums"
	"golang.org/x/text/language"
)

// This class is used because in COSEM object model some fields from date time can be ignored.
//
//	Default behavior of DateTime do not allow this.
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

// currentLanguage returns current user language.
func currentLanguage() *language.Tag {
	langEnv := os.Getenv("LANG")
	if langEnv == "" {
		return &language.AmericanEnglish
	}
	langEnv = strings.Split(langEnv, ".")[0]
	tag, err := language.Parse(langEnv)
	if err != nil {
		return &language.AmericanEnglish
	}
	return &tag
}

// Returns the date time format for the given culture.
func getDateTimeFormat(language *language.Tag, skip enums.DateTimeSkips) string {
	if language == nil {
		language = currentLanguage()
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

// timeZonePosition returns the check is time zone included and return index of time zone.value:
//
//	Returns:
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

func (g *GXDateTime) ToFormatString(language *language.Tag, useLocalTime bool) string {
	return toString(g, language, useLocalTime, true)
}

func (g *GXDateTime) ToFormatMeterString(language *language.Tag) string {
	return toString(g, language, false, true)
}

// GXDateTimeFromUnixTime creater the  date time from Epoch time.
// unixTime: Unix time.
//
//	Returns:
//	    Date and time.
func GXDateTimeFromUnixTime(unixTime int64) *GXDateTime {
	return &GXDateTime{
		Value: time.Unix(unixTime, 0),
	}
}

// GXDateTimeFromHighResolutionTime returns the get date time from high resolution clock time.highResolution: High resolution clock time.
//
//	Returns:
//	    Date and time.
func GXDateTimeFromHighResolutionTime(highResolution int64) *GXDateTime {
	return &GXDateTime{
		Value: time.Unix(int64(highResolution), 0),
	}
}

// ToUnixTime returns the convert date time to Epoch
//
//	Returns:
//	    Unix time.
func (g *GXDateTime) ToUnixTime() int64 {
	return g.Value.Unix()
}

// ToHighResolutionTime returns the convert date time to high resolution time.
//
//	Returns:
//	    High resolution time.
func (g *GXDateTime) ToHighResolutionTime() uint64 {
	if g.Value.IsZero() {
		return 0
	}
	return uint64(g.Value.Unix())
}

// ToHex returns the get date time as hex string.
// addSpace: Add space between bytes.
// useMeterTimeZone: Date-Time values are shown using meter's time zone and it's not localized to use PC time.
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

// Constructorvalue: Date time value as a string.
// culture: Used culture.
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
		// If time zone is used.
		pos := timeZonePosition(value)
		if !addTimeZone || pos == -1 {
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
			return err
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

// Constructorvalue: Date time value as a string.
// culture: Used culture.
func NewGXDateTimeFromString(value string, language *language.Tag) (*GXDateTime, error) {
	g := &GXDateTime{}
	err := parseInternal(g, value, language)
	if err != nil {
		return nil, err
	}
	return g, nil
}

// Constructorvalue: Date time value.
func NewGXDateTimeFromTime(value time.Time) *GXDateTime {
	return &GXDateTime{Value: value}
}
