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
	"errors"

	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// Online help:
// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSActivityCalendar
type GXDLMSActivityCalendar struct {
	GXDLMSObject
	CalendarNameActive string

	SeasonProfileActive []GXDLMSSeasonProfile

	WeekProfileTableActive []GXDLMSWeekProfile

	DayProfileTableActive []GXDLMSDayProfile

	CalendarNamePassive string

	SeasonProfilePassive []GXDLMSSeasonProfile

	WeekProfileTablePassive []GXDLMSWeekProfile

	DayProfileTablePassive []GXDLMSDayProfile

	// Activate Passive Calendar Time.
	Time types.GXDateTime
}

// base returns the base GXDLMSObject of the object.
func (g *GXDLMSActivityCalendar) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

// GetAttributeIndexToRead returns the collection of attributes to read.
// If attribute is static and already read or device is returned HW error it is not returned.
//
// Parameters:
//
//	all: All items are returned even if they are read already.
//
// Returns:
//
//	Collection of attributes to read.
func (g *GXDLMSActivityCalendar) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	// LN is static and read only once.
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	// CalendarNameActive
	if all || g.CanRead(2) {
		attributes = append(attributes, 2)
	}
	// SeasonProfileActive
	if all || g.CanRead(3) {
		attributes = append(attributes, 3)
	}
	// WeekProfileTableActive
	if all || g.CanRead(4) {
		attributes = append(attributes, 4)
	}
	// DayProfileTableActive
	if all || g.CanRead(5) {
		attributes = append(attributes, 5)
	}
	// CalendarNamePassive
	if all || g.CanRead(6) {
		attributes = append(attributes, 6)
	}
	// SeasonProfilePassive
	if all || g.CanRead(7) {
		attributes = append(attributes, 7)
	}
	// WeekProfileTablePassive
	if all || g.CanRead(8) {
		attributes = append(attributes, 8)
	}
	// DayProfileTablePassive
	if all || g.CanRead(9) {
		attributes = append(attributes, 9)
	}
	// Time.
	if all || g.CanRead(10) {
		attributes = append(attributes, 10)
	}
	return attributes
}

// GetNames returns the names of attribute indexes.
func (g *GXDLMSActivityCalendar) GetNames() []string {
	return []string{"Logical Name", "Active calendar name", "Active season profile", "Active week profile table", "Active day profile table", "Passive calendar name", "Passive season profile", "Passive week profile table", "Passive day profile table", "Time"}
}

// GetMethodNames returns the names of method indexes.
func (g *GXDLMSActivityCalendar) GetMethodNames() []string {
	return []string{"Activate passive calendar"}
}

// GetAttributeCount returns the amount of attributes.
//
// Returns:
//
//	Count of attributes.
func (g *GXDLMSActivityCalendar) GetAttributeCount() int {
	return 10
}

// GetMethodCount returns the amount of methods.
func (g *GXDLMSActivityCalendar) GetMethodCount() int {
	return 1
}

func (g *GXDLMSActivityCalendar) isSec(settings *settings.GXDLMSSettings) bool {
	/* TODO:
	if c, ok := g.Parent.parent.(GXDLMSClient); ok {
		return c.Standard() == enums.StandardSaudiArabia
	}
	*/
	return false
}

// GetSeasonProfile returns the get season profile bytes.
//
// Parameters:
//
//	settings: DLMS settings.
//	target: Season profile array.
//	useOctetString: Is date time send as octet string.
func (g *GXDLMSActivityCalendar) GetSeasonProfile(settings *settings.GXDLMSSettings, target []GXDLMSSeasonProfile, useOctetString bool) (any, error) {
	data := types.NewGXByteBuffer()
	err := data.SetUint8(uint8(enums.DataTypeArray))
	if err != nil {
		return nil, err
	}
	if target == nil {
		err = types.SetObjectCount(0, data)
	} else {
		cnt := len(target)
		err = types.SetObjectCount(cnt, data)
		for _, it := range target {
			err = data.SetUint8(uint8(enums.DataTypeStructure))
			if err != nil {
				return nil, err
			}
			err = data.SetUint8(3)
			if err != nil {
				return nil, err
			}
			err = internal.SetData(settings, data, enums.DataTypeOctetString, it.Name)
			if useOctetString {
				err = internal.SetData(settings, data, enums.DataTypeOctetString, it.Start)
			} else {
				err = internal.SetData(settings, data, enums.DataTypeDateTime, it.Start)
			}
			err = internal.SetData(settings, data, enums.DataTypeOctetString, it.WeekName)
		}
	}
	return data, nil
}

func (g *GXDLMSActivityCalendar) GetWeekProfileTable(settings *settings.GXDLMSSettings, target []GXDLMSWeekProfile) ([]byte, error) {
	data := types.NewGXByteBuffer()
	err := data.SetUint8(uint8(enums.DataTypeArray))
	if err != nil {
		return nil, err
	}
	cnt := len(target)
	err = types.SetObjectCount(cnt, data)
	for _, it := range target {
		err = data.SetUint8(uint8(enums.DataTypeStructure))
		if err != nil {
			return nil, err
		}
		err = data.SetUint8(8)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, data, enums.DataTypeOctetString, it.Name)
		err = internal.SetData(settings, data, enums.DataTypeUint8, it.Monday)
		err = internal.SetData(settings, data, enums.DataTypeUint8, it.Tuesday)
		err = internal.SetData(settings, data, enums.DataTypeUint8, it.Wednesday)
		err = internal.SetData(settings, data, enums.DataTypeUint8, it.Thursday)
		err = internal.SetData(settings, data, enums.DataTypeUint8, it.Friday)
		err = internal.SetData(settings, data, enums.DataTypeUint8, it.Saturday)
		err = internal.SetData(settings, data, enums.DataTypeUint8, it.Sunday)
	}
	return data.Array(), nil
}

func (g *GXDLMSActivityCalendar) GetDayProfileTable(settings *settings.GXDLMSSettings, target []GXDLMSDayProfile) ([]byte, error) {
	data := types.NewGXByteBuffer()
	err := data.SetUint8(uint8(enums.DataTypeArray))
	if err != nil {
		return nil, err
	}
	cnt := len(target)
	//Add count
	err = types.SetObjectCount(cnt, data)
	for _, it := range target {
		err = data.SetUint8(uint8(enums.DataTypeStructure))
		if err != nil {
			return nil, err
		}
		err = data.SetUint8(2)
		if err != nil {
			return nil, err
		}
		err = internal.SetData(settings, data, enums.DataTypeUint8, it.DayId)
		if err != nil {
			return nil, err
		}
		err = data.SetUint8(uint8(enums.DataTypeArray))
		if err != nil {
			return nil, err
		}
		err = types.SetObjectCount(len(it.DaySchedules), data)
		if err != nil {
			return nil, err
		}
		for _, action := range it.DaySchedules {
			err = data.SetUint8(uint8(enums.DataTypeStructure))
			if err != nil {
				return nil, err
			}
			err = data.SetUint8(3)
			if err != nil {
				return nil, err
			}
			ln, err := helpers.LogicalNameToBytes(action.ScriptLogicalName)
			if err != nil {
				return nil, err
			}
			err = internal.SetData(settings, data, enums.DataTypeOctetString, action.StartTime)
			if err != nil {
				return nil, err
			}
			err = internal.SetData(settings, data, enums.DataTypeOctetString, ln)
			if err != nil {
				return nil, err
			}
			err = internal.SetData(settings, data, enums.DataTypeUint16, action.ScriptSelector)
			if err != nil {
				return nil, err
			}
		}
	}
	return data.Array(), nil
}

// GetValue returns the value of given attribute.
// When raw parameter us not used example register multiplies value by scalar.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Get parameters.
//
// Returns:
//
//	Value of the attribute index.
func (g *GXDLMSActivityCalendar) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	if e.Index == 1 {
		v, err := helpers.LogicalNameToBytes(g.LogicalName())
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return v, err
	}
	if e.Index == 2 {
		if g.isSec(settings) {
			return types.HexToBytes(g.CalendarNameActive), nil
		}
		return []byte(g.CalendarNameActive), nil
	}
	if e.Index == 3 {
		e.ByteArray = true
		useOctetString := settings.Standard != enums.StandardSaudiArabia
		return g.GetSeasonProfile(settings, g.SeasonProfileActive, useOctetString)
	}
	if e.Index == 4 {
		e.ByteArray = true
		return g.GetWeekProfileTable(settings, g.WeekProfileTableActive)
	}
	if e.Index == 5 {
		e.ByteArray = true
		return g.GetDayProfileTable(settings, g.DayProfileTableActive)
	}
	if e.Index == 6 {
		if g.isSec(settings) {
			return types.HexToBytes(g.CalendarNamePassive), nil
		}
		return []byte(g.CalendarNamePassive), nil
	}
	if e.Index == 7 {
		e.ByteArray = true
		useOctetString := settings.Standard != enums.StandardSaudiArabia
		return g.GetSeasonProfile(settings, g.SeasonProfilePassive, useOctetString)
	}
	if e.Index == 8 {
		e.ByteArray = true
		return g.GetWeekProfileTable(settings, g.WeekProfileTablePassive)
	}
	if e.Index == 9 {
		e.ByteArray = true
		return g.GetDayProfileTable(settings, g.DayProfileTablePassive)
	}
	if e.Index == 10 {
		return g.Time, nil
	}
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, nil
}

func (g *GXDLMSActivityCalendar) SetSeasonProfile(settings *settings.GXDLMSSettings, value any) ([]GXDLMSSeasonProfile, error) {
	if value != nil {
		items := []GXDLMSSeasonProfile{}
		for _, tmp := range value.(types.GXArray) {
			item := tmp.(types.GXStructure)
			it := GXDLMSSeasonProfile{}
			it.Name = item[0].([]byte)
			if v, ok := item[1].([]byte); ok {
				ret, err := internal.ChangeTypeFromByteArray(settings, v, enums.DataTypeDateTime)
				if err != nil {
					return nil, err
				}
				it.Start = ret.(types.GXDateTime)
			} else if v, ok := item[1].(types.GXDateTime); ok {
				it.Start = v
			} else {
				return nil, errors.New("Invalid date time.")
			}
			it.WeekName = item[2].([]byte)
			items = append(items, it)
		}
		return items, nil
	}
	return nil, nil
}

func (g *GXDLMSActivityCalendar) SetWeekProfileTable(settings *settings.GXDLMSSettings, value any) ([]GXDLMSWeekProfile, error) {
	items := []GXDLMSWeekProfile{}
	if value != nil {
		for _, tmp := range value.(types.GXArray) {
			item := tmp.(types.GXStructure)
			it := GXDLMSWeekProfile{}
			it.Name = item[0].([]byte)
			it.Monday = item[1].(uint8)
			it.Tuesday = item[2].(uint8)
			it.Wednesday = item[3].(uint8)
			it.Thursday = item[4].(uint8)
			it.Friday = item[5].(uint8)
			it.Saturday = item[6].(uint8)
			it.Sunday = item[7].(uint8)
			items = append(items, it)
		}
	}
	return items, nil
}

func (g *GXDLMSActivityCalendar) SetDayProfileTable(settings *settings.GXDLMSSettings, value any) ([]GXDLMSDayProfile, error) {
	var err error
	items := []GXDLMSDayProfile{}
	if value != nil {
		for _, tmp := range value.(types.GXArray) {
			item := tmp.(types.GXStructure)
			it := GXDLMSDayProfile{}
			it.DayId = item[0].(uint8)
			var actions []GXDLMSDayProfileAction
			for _, tmp2 := range item[1].(types.GXArray) {
				it2 := tmp2.(types.GXStructure)
				ac := GXDLMSDayProfileAction{}
				if t, ok := it2[0].(types.GXTime); ok {
					ac.StartTime = t
				} else if dt, ok := it2[0].(types.GXDateTime); ok {
					ac.StartTime = *types.NewGXTimeFromDateTime(&dt)
				} else {
					ret, err := internal.ChangeTypeFromByteArray(settings, it2[0].([]byte), enums.DataTypeTime)
					if err != nil {
						return nil, err
					}
					ac.StartTime = ret.(types.GXTime)
				}
				ac.ScriptLogicalName, err = helpers.ToLogicalName(it2[1].([]byte))
				if err != nil {
					return nil, err
				}
				ac.ScriptSelector = it2[2].(uint16)
				actions = append(actions, ac)
			}
			it.DaySchedules = actions
			items = append(items, it)
		}
	}
	return items, nil
}

// SetValue returns the set value of given attribute.
// When raw parameter us not used example register multiplies value by scalar.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Set parameters.
func (g *GXDLMSActivityCalendar) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	var err error
	if e.Index == 1 {
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return g.SetLogicalName(ln)
	} else if e.Index == 2 {
		if v, ok := e.Value.([]byte); ok {
			if g.isSec(settings) || !types.IsAsciiString(v) {
				g.CalendarNameActive = types.ToHex(v, false)
			} else {
				g.CalendarNameActive = string(v)
			}
		} else {
			g.CalendarNameActive = e.Value.(string)
		}
	} else if e.Index == 3 {
		g.SeasonProfileActive, err = g.SetSeasonProfile(settings, e.Value)
	} else if e.Index == 4 {
		g.WeekProfileTableActive, err = g.SetWeekProfileTable(settings, e.Value)
	} else if e.Index == 5 {
		g.DayProfileTableActive, err = g.SetDayProfileTable(settings, e.Value)
	} else if e.Index == 6 {
		if v, ok := e.Value.([]byte); ok {
			if g.isSec(settings) || !types.IsAsciiString(v) {
				g.CalendarNamePassive = types.ToHex(v, false)
			} else {
				g.CalendarNamePassive = string(v)
			}
		} else {
			g.CalendarNamePassive = e.Value.(string)
		}
	} else if e.Index == 7 {
		g.SeasonProfilePassive, err = g.SetSeasonProfile(settings, e.Value)
	} else if e.Index == 8 {
		g.WeekProfileTablePassive, err = g.SetWeekProfileTable(settings, e.Value)
	} else if e.Index == 9 {
		g.DayProfileTablePassive, err = g.SetDayProfileTable(settings, e.Value)
	} else if e.Index == 10 {
		if v, ok := e.Value.([]byte); ok {
			ret, err := internal.ChangeTypeFromByteArray(settings, v, enums.DataTypeDateTime)
			if err != nil {
				return err
			}
			g.Time = ret.(types.GXDateTime)
		} else if v, ok := e.Value.(types.GXDateTime); ok {
			g.Time = v
		}
	} else {
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return err
}

// Invoke returns the invokes method.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Invoke parameters.
func (g *GXDLMSActivityCalendar) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, nil
}

func (g *GXDLMSActivityCalendar) LoadSeasonProfile(reader *GXXmlReader, name string) ([]GXDLMSSeasonProfile, error) {
	var err error
	list := []GXDLMSSeasonProfile{}
	if reader.isStartElementNamed2(name, true) {
		for reader.isStartElementNamed2("Item", true) {
			it := GXDLMSSeasonProfile{}
			ret, err := reader.ReadElementContentAsString("Name", "")
			if err != nil {
				return nil, err
			}
			it.Name = types.HexToBytes(ret)
			it.Start, err = reader.ReadElementContentAsDateTime("Start", nil)
			if err != nil {
				return nil, err
			}
			ret, err = reader.ReadElementContentAsString("WeekName", "")
			if err != nil {
				return nil, err
			}
			it.WeekName = types.HexToBytes(ret)
			list = append(list, it)
		}
		reader.ReadEndElement(name)
	}
	return list, err
}

func (g *GXDLMSActivityCalendar) LoadWeekProfileTable(reader *GXXmlReader, name string) ([]GXDLMSWeekProfile, error) {
	var err error
	list := []GXDLMSWeekProfile{}
	if reader.isStartElementNamed2(name, true) {
		for reader.isStartElementNamed2("Item", true) {
			it := GXDLMSWeekProfile{}
			ret1, err := reader.ReadElementContentAsString("Name", "")
			if err != nil {
				return nil, err
			}
			it.Name = types.HexToBytes(ret1)
			ret, err := reader.ReadElementContentAsInt("Monday", 0)
			if err != nil {
				return nil, err
			}
			it.Monday = uint8(ret)
			ret, err = reader.ReadElementContentAsInt("Tuesday", 0)
			if err != nil {
				return nil, err
			}
			it.Tuesday = uint8(ret)
			ret, err = reader.ReadElementContentAsInt("Wednesday", 0)
			if err != nil {
				return nil, err
			}
			it.Wednesday = uint8(ret)
			ret, err = reader.ReadElementContentAsInt("Thursday", 0)
			if err != nil {
				return nil, err
			}
			it.Thursday = uint8(ret)
			ret, err = reader.ReadElementContentAsInt("Friday", 0)
			if err != nil {
				return nil, err
			}
			it.Friday = uint8(ret)
			ret, err = reader.ReadElementContentAsInt("Saturday", 0)
			if err != nil {
				return nil, err
			}
			it.Saturday = uint8(ret)
			ret, err = reader.ReadElementContentAsInt("Sunday", 0)
			if err != nil {
				return nil, err
			}
			it.Sunday = uint8(ret)
			list = append(list, it)
		}
		reader.ReadEndElement(name)
	}
	return list, err
}

func (g *GXDLMSActivityCalendar) LoadDayProfileTable(reader *GXXmlReader, name string) ([]GXDLMSDayProfile, error) {
	list := []GXDLMSDayProfile{}
	if reader.isStartElementNamed2(name, true) {
		for reader.isStartElementNamed2("Item", true) {
			it := GXDLMSDayProfile{}
			ret, err := reader.ReadElementContentAsInt("DayId", 0)
			if err != nil {
				return nil, err
			}
			it.DayId = uint8(ret)
			list = append(list, it)
			var actions []GXDLMSDayProfileAction
			if reader.isStartElementNamed2("Actions", true) {
				for reader.isStartElementNamed2("Action", true) {
					d := GXDLMSDayProfileAction{}
					actions = append(actions, d)
					d.StartTime, err = reader.ReadElementContentAsTime("Start")
					if err != nil {
						return nil, err
					}
					d.ScriptLogicalName, err = reader.ReadElementContentAsString("LN", "")
					if err != nil {
						return nil, err
					}
					ret, err := reader.ReadElementContentAsInt("Selector", 0)
					if err != nil {
						return nil, err
					}
					d.ScriptSelector = uint16(ret)
				}
				reader.ReadEndElement("Actions")
			}
			it.DaySchedules = actions
		}
		reader.ReadEndElement(name)
	}
	return list, nil
}

// Load returns the load object content from XML.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSActivityCalendar) Load(reader *GXXmlReader) error {
	var err error
	g.CalendarNameActive, err = reader.ReadElementContentAsString("CalendarNameActive", "")
	if err != nil {
		return err
	}
	g.SeasonProfileActive, err = g.LoadSeasonProfile(reader, "SeasonProfileActive")
	if err != nil {
		return err
	}
	g.WeekProfileTableActive, err = g.LoadWeekProfileTable(reader, "WeekProfileTableActive")
	if err != nil {
		return err
	}
	g.DayProfileTableActive, err = g.LoadDayProfileTable(reader, "DayProfileTableActive")
	if err != nil {
		return err
	}
	g.CalendarNamePassive, err = reader.ReadElementContentAsString("CalendarNamePassive", "")
	if err != nil {
		return err
	}
	g.SeasonProfilePassive, err = g.LoadSeasonProfile(reader, "SeasonProfilePassive")
	if err != nil {
		return err
	}
	g.WeekProfileTablePassive, err = g.LoadWeekProfileTable(reader, "WeekProfileTablePassive")
	if err != nil {
		return err
	}
	g.DayProfileTablePassive, err = g.LoadDayProfileTable(reader, "DayProfileTablePassive")
	if err != nil {
		return err
	}
	g.Time, err = reader.ReadElementContentAsGXDateTime("Time")
	return err
}

func (g *GXDLMSActivityCalendar) SaveSeasonProfile(writer *GXXmlWriter, list []GXDLMSSeasonProfile, name string) error {
	var err error
	if list != nil {
		writer.WriteStartElement(name)
		for _, it := range list {
			writer.WriteStartElement("Item")
			err = writer.WriteElementString("Name", types.ToHex(it.Name, false))
			if err != nil {
				return err
			}
			writer.WriteElementString("Start", it.Start)
			err = writer.WriteElementString("WeekName", types.ToHex(it.WeekName, false))
			if err != nil {
				return err
			}
			writer.WriteEndElement()
		}
		writer.WriteEndElement()
	}
	return err
}

func (g *GXDLMSActivityCalendar) SaveWeekProfileTable(writer *GXXmlWriter, list []GXDLMSWeekProfile, name string) error {
	var err error
	if list != nil {
		writer.WriteStartElement(name)
		for _, it := range list {
			writer.WriteStartElement("Item")
			err = writer.WriteElementString("Name", types.ToHex(it.Name, false))
			if err != nil {
				return err
			}
			err = writer.WriteElementString("Monday", it.Monday)
			if err != nil {
				return err
			}
			err = writer.WriteElementString("Tuesday", it.Tuesday)
			if err != nil {
				return err
			}
			err = writer.WriteElementString("Wednesday", it.Wednesday)
			if err != nil {
				return err
			}
			err = writer.WriteElementString("Thursday", it.Thursday)
			if err != nil {
				return err
			}
			err = writer.WriteElementString("Friday", it.Friday)
			if err != nil {
				return err
			}
			err = writer.WriteElementString("Saturday", it.Saturday)
			if err != nil {
				return err
			}
			err = writer.WriteElementString("Sunday", it.Sunday)
			if err != nil {
				return err
			}
			writer.WriteEndElement()
		}
		writer.WriteEndElement()
	}
	return err
}

func (g *GXDLMSActivityCalendar) SaveDayProfileTable(writer *GXXmlWriter, list []GXDLMSDayProfile, name string) error {
	var err error
	if list != nil {
		writer.WriteStartElement(name)
		for _, it := range list {
			writer.WriteStartElement("Item")
			err = writer.WriteElementString("DayId", it.DayId)
			if err != nil {
				return err
			}
			writer.WriteStartElement("Actions")
			for _, d := range it.DaySchedules {
				writer.WriteStartElement("Action")
				writer.WriteElementString("Start", d.StartTime)
				writer.WriteElementString("LN", d.ScriptLogicalName)
				writer.WriteElementString("Selector", d.ScriptSelector)
				writer.WriteEndElement()
			}
			writer.WriteEndElement()
			writer.WriteEndElement()
		}
		writer.WriteEndElement()
	}
	return err
}

// Save returns the save object content to XML.
//
// Parameters:
//
//	writer: XML writer.
func (g *GXDLMSActivityCalendar) Save(writer *GXXmlWriter) error {
	err := writer.WriteElementString("CalendarNameActive", g.CalendarNameActive)
	if err != nil {
		return err
	}
	g.SaveSeasonProfile(writer, g.SeasonProfileActive, "SeasonProfileActive")
	g.SaveWeekProfileTable(writer, g.WeekProfileTableActive, "WeekProfileTableActive")
	g.SaveDayProfileTable(writer, g.DayProfileTableActive, "DayProfileTableActive")
	err = writer.WriteElementString("CalendarNamePassive", g.CalendarNamePassive)
	if err != nil {
		return err
	}
	g.SaveSeasonProfile(writer, g.SeasonProfilePassive, "SeasonProfilePassive")
	g.SaveWeekProfileTable(writer, g.WeekProfileTablePassive, "WeekProfileTablePassive")
	g.SaveDayProfileTable(writer, g.DayProfileTablePassive, "DayProfileTablePassive")
	err = writer.WriteElementString("Time", g.Time)
	if err != nil {
		return err
	}
	return err
}

// PostLoad returns the handle actions after Load.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSActivityCalendar) PostLoad(reader *GXXmlReader) error {
	return nil
}

// GetValues returns the an array containing the COSEM object's attribute values.
func (g *GXDLMSActivityCalendar) GetValues() []any {
	return []any{g.LogicalName, g.CalendarNameActive, g.SeasonProfileActive, g.WeekProfileTableActive, g.DayProfileTableActive, g.CalendarNamePassive, g.SeasonProfilePassive, g.WeekProfileTablePassive, g.DayProfileTablePassive, g.Time}
}

// ActivatePassiveCalendar returns the this method copies all passive parameter to the active parameter.
//
// Parameters:
//
//	client: DLMS client.
//
// Returns:
//
//	Action bytes.
func (g *GXDLMSActivityCalendar) ActivatePassiveCalendar(client IGXDLMSClient) ([][]byte, error) {
	return client.Method(g, 1, int8(0), enums.DataTypeInt8)
}

// GetUIDataType returns the UI data type of selected index.
//
// Parameters:
//
//	index: Attribute index of the object.
//
// Returns:
//
//	UI data type of the object.
func (g *GXDLMSActivityCalendar) GetUIDataType(index int) enums.DataType {
	if index == 2 || index == 6 {
		return enums.DataTypeString
	}
	if index == 10 {
		return enums.DataTypeDateTime
	}
	return g.Base().GetUIDataType(index)
}

// GetDataType returns the device data type of selected attribute index.
//
// Parameters:
//
//	index: Attribute index of the object.
//
// Returns:
//
//	Device data type of the object.
func (g *GXDLMSActivityCalendar) GetDataType(index int) (enums.DataType, error) {
	ret := enums.DataTypeNone
	if index == 1 {
		ret = enums.DataTypeOctetString
	} else if index == 2 {
		ret = enums.DataTypeOctetString
	} else if index == 3 {
		ret = enums.DataTypeArray
	} else if index == 4 {
		ret = enums.DataTypeArray
	} else if index == 5 {
		ret = enums.DataTypeArray
	} else if index == 6 {
		ret = enums.DataTypeOctetString
	} else if index == 7 {
		ret = enums.DataTypeArray
	} else if index == 8 {
		ret = enums.DataTypeArray
	} else if index == 9 {
		ret = enums.DataTypeArray
	} else if index == 10 {
		if g.isSec(nil) {
			ret = enums.DataTypeDateTime
		} else {
			ret = enums.DataTypeOctetString
		}
	} else {
		return enums.DataTypeNone, errors.New("GetDataType failed. Invalid attribute index.")
	}
	return ret, nil
}

// Constructor.
// ln: Logical Name of the object.
// sn: Short Name of the object.
func NewGXDLMSActivityCalendar(ln string, sn int16) (*GXDLMSActivityCalendar, error) {
	err := ValidateLogicalName(ln)
	if err != nil {
		return nil, err
	}
	return &GXDLMSActivityCalendar{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeActivityCalendar,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
