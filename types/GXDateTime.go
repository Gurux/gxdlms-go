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
	"US": "01/02/2006 03:04:05.000 PM -0700",
	"DE": "02.01.2006 15:04:05.000 -0700",
	"GB": "02/01/2006 15:04:05.000 -0700",
	"ES": "02/01/2006 15:04:05.000 -0700",
	"ET": "02.01.2006 15:04:05.000 -0700",
	"FR": "02/01/2006 15:04:05.000 -0700",
	"IT": "02/01/2006 15:04:05.000 -0700",
	"HI": "02/01/2006 03:04:05.000 PM -0700",
	"SV": "02/01/2006 15:04:05.000 -0700",
}

// Returns the date time format for the given culture.
func getDateTimeFormat(language *language.Tag, skip enums.DateTimeSkips) string {
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
		pos := len(value) - 6
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
	return g.toString(g, language, useLocalTime)
}

func (g *GXDateTime) toString(target any, language *language.Tag, useLocalTime bool) string {
	//Get current date time format.
	format := getDateTimeFormat(language, g.Skip)
	if useLocalTime {
		//Remove time zone if year is removed.
		format = remove(format, "-0700", "")
		format = strings.TrimSpace(format)
	}
	if g.Skip != enums.DateTimeSkipsNone {
		var timeSeparator, dateSeparator string
		detectSeparators(format, &timeSeparator, &dateSeparator)
		// FormatException is thrown if length of format is 1.
		if !strings.Contains(format, dateSeparator) && !strings.Contains(format, timeSeparator) {
			if (g.Skip & enums.DateTimeSkipsYear) == 0 {
				return fmt.Sprint()
			}
			if (g.Skip & enums.DateTimeSkipsMonth) == 0 {
				return fmt.Sprint()
			}
			if (g.Skip & enums.DateTimeSkipsDay) == 0 {
				return fmt.Sprint()
			}
			if (g.Skip & enums.DateTimeSkipsHour) == 0 {
				return fmt.Sprint()
			}
			if (g.Skip & enums.DateTimeSkipsMinute) == 0 {
				return fmt.Sprint()
			}
			if (g.Skip & enums.DateTimeSkipsSecond) == 0 {
				return fmt.Sprint()
			}
			if (g.Skip & enums.DateTimeSkipsMs) == 0 {
				return fmt.Sprint()
			}
		}
		if useLocalTime {
			return g.Value.Local().Format(format)
		}
		ret := g.Value.Format(format)
		_, offset := g.Value.Zone()
		if offset == 0 {
			ret = ret[:len(ret)-6] + "Z"
		}
		return ret
	}
	if useLocalTime {
		return g.Value.Local().Format(format)
	}
	return g.Value.Format(format)
}

func (g *GXDateTime) ToFormatString(language *language.Tag, useLocalTime bool) string {
	//TODO:
	return g.toString(g, language, useLocalTime)
}

func (g *GXDateTime) ToFormatMeterString(language *language.Tag) string {
	//TODO:
	return g.toString(g, language, false)
}

// fromUnixTime returns the get date time from Epoch time.unixTime: Unix time.
//
//	Returns:
//	    Date and time.
func GXDateTimeFromUnixTime(unixTime int64) *GXDateTime {
	return &GXDateTime{
		Value: time.Unix(unixTime, 0),
	}
}

// fromHighResolutionTime returns the get date time from high resolution clock time.highResolution: High resolution clock time.
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

// Constructorvalue: Date time value as a string.
// culture: Used culture.
func (g *GXDateTime) parseInternal(value string, language *language.Tag) error {
	addTimeZone := true
	g.DayOfWeek = 0xFF
	if value != "" {
		format := getDateTimeFormat(language, g.Skip)
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
			g.Skip |= enums.DateTimeSkipsDayOfWeek
		}
		// If time zone is used.
		pos := timeZonePosition(value)
		if !addTimeZone || pos == -1 {
			//Remove time zone.
			var timeSeparator, dateSeparator string
			detectSeparators(format, &timeSeparator, &dateSeparator)
			format = remove(format, "-0700", timeSeparator)
			// Trim
			format = strings.TrimSpace(format)
			if language == nil {
				g.Skip |= enums.DateTimeSkipsDeviation
			}
		}
		var err error
		if language == nil {
			g.Value, err = time.Parse(v, format)
		} else {
			g.Value, err = time.ParseInLocation(format, v, time.Local)
		}
		if err != nil {
			//Remove ms.
			format = remove(format, ".000", "")
			// Trim
			format = strings.TrimSpace(format)
			g.Value, err = time.ParseInLocation(format, v, time.Local)
		}
		if err != nil {
			return err
		}
		g.Skip |= enums.DateTimeSkipsDayOfWeek
		g.Skip |= enums.DateTimeSkipsMs
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
	err := g.parseInternal(value, language)
	if err != nil {
		return nil, err
	}
	return g, nil
}

// Constructorvalue: Date time value.
func NewGXDateTimeFromTime(value time.Time) *GXDateTime {
	return &GXDateTime{Value: value}
}
