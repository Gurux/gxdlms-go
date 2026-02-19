package internal

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/Gurux/gxcommon-go"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal/buffer"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// HDLCFrameStartEnd is the HDLC frame start and end byte.
const HDLCFrameStartEnd byte = 0x7E

// LLCSendBytes are the sent LLC bytes.
var LLCSendBytes = []byte{0xE6, 0xE6, 0x00}

// LLCReplyBytes are the received LLC bytes.
var LLCReplyBytes = []byte{0xE6, 0xE7, 0x00}

const CipheringHeaderSize = 7 + 12 + 3 // Security control + invocation counter + length bytes

type GXCommon struct {
}

func IsHexString(value string) bool {
	if value == "" {
		return false
	}
	for i := 0; i < len(value); i++ {
		c := value[i]
		if c != ' ' && ((c < '0' || c > '9') && (c < 'A' || c > 'F') && (c < 'a' || c > 'f')) {
			return false
		}
	}
	return true
}

func HexToBytes(value string) ([]byte, error) {
	if value == "" {
		return []byte{}, nil
	}
	if !IsHexString(value) {
		return nil, fmt.Errorf("invalid hex string: %q", value)
	}
	return buffer.HexToBytes(value), nil
}

func ToHex(bytes []byte, addSpace bool, index int, count int) string {
	if bytes == nil || count <= 0 || index < 0 || index >= len(bytes) {
		return ""
	}
	if index+count > len(bytes) {
		count = len(bytes) - index
	}
	return buffer.ToHexWithRange(bytes, addSpace, index, count)
}

func GetDataType(dt enums.DataType) (reflect.Type, error) {
	switch dt {
	case enums.DataTypeArray:
		return reflect.TypeOf(types.GXArray{}), nil
	case enums.DataTypeStructure:
		return reflect.TypeOf(types.GXStructure{}), nil
	case enums.DataTypeBoolean:
		return reflect.TypeOf(false), nil
	case enums.DataTypeBitString:
		return reflect.TypeOf(types.GXBitString{}), nil
	case enums.DataTypeInt8:
		return reflect.TypeOf(int8(0)), nil
	case enums.DataTypeInt16:
		return reflect.TypeOf(int16(0)), nil
	case enums.DataTypeInt32:
		return reflect.TypeOf(int32(0)), nil
	case enums.DataTypeInt64:
		return reflect.TypeOf(int64(0)), nil
	case enums.DataTypeUint8, enums.DataTypeEnum:
		return reflect.TypeOf(uint8(0)), nil
	case enums.DataTypeUint16:
		return reflect.TypeOf(uint16(0)), nil
	case enums.DataTypeUint32:
		return reflect.TypeOf(uint32(0)), nil
	case enums.DataTypeUint64:
		return reflect.TypeOf(uint64(0)), nil
	case enums.DataTypeFloat32:
		return reflect.TypeOf(float32(0)), nil
	case enums.DataTypeFloat64:
		return reflect.TypeOf(float64(0)), nil
	case enums.DataTypeString, enums.DataTypeStringUTF8:
		return reflect.TypeOf(""), nil
	case enums.DataTypeOctetString:
		return reflect.TypeOf([]byte{}), nil
	case enums.DataTypeDate, enums.DataTypeDateTime, enums.DataTypeTime:
		return reflect.TypeOf(time.Time{}), nil
	default:
		return nil, fmt.Errorf("unsupported data type: %v", dt)
	}
}

func GetDLMSDataType(dt reflect.Type) (enums.DataType, error) {
	if dt == nil {
		return enums.DataTypeNone, nil
	}
	if dt == reflect.TypeOf(types.GXArray{}) || dt == reflect.TypeOf([]any{}) {
		return enums.DataTypeArray, nil
	}
	if dt == reflect.TypeOf(types.GXStructure{}) {
		return enums.DataTypeStructure, nil
	}
	if dt == reflect.TypeOf(false) {
		return enums.DataTypeBoolean, nil
	}
	if dt == reflect.TypeOf(types.GXBitString{}) || dt == reflect.TypeOf(&types.GXBitString{}) {
		return enums.DataTypeBitString, nil
	}
	if dt == reflect.TypeOf(int8(0)) {
		return enums.DataTypeInt8, nil
	}
	if dt == reflect.TypeOf(int16(0)) {
		return enums.DataTypeInt16, nil
	}
	if dt == reflect.TypeOf(int32(0)) || dt == reflect.TypeOf(int(0)) {
		return enums.DataTypeInt32, nil
	}
	if dt == reflect.TypeOf(int64(0)) {
		return enums.DataTypeInt64, nil
	}
	if dt == reflect.TypeOf(uint8(0)) {
		return enums.DataTypeUint8, nil
	}
	if dt == reflect.TypeOf(uint16(0)) {
		return enums.DataTypeUint16, nil
	}
	if dt == reflect.TypeOf(uint32(0)) || dt == reflect.TypeOf(uint(0)) {
		return enums.DataTypeUint32, nil
	}
	if dt == reflect.TypeOf(uint64(0)) {
		return enums.DataTypeUint64, nil
	}
	if dt == reflect.TypeOf(float32(0)) {
		return enums.DataTypeFloat32, nil
	}
	if dt == reflect.TypeOf(float64(0)) {
		return enums.DataTypeFloat64, nil
	}
	if dt == reflect.TypeOf("") {
		return enums.DataTypeString, nil
	}
	if dt == reflect.TypeOf([]byte{}) {
		return enums.DataTypeOctetString, nil
	}
	if dt == reflect.TypeOf(time.Time{}) {
		return enums.DataTypeDateTime, nil
	}
	if dt == reflect.TypeOf(types.GXByteBuffer{}) {
		return enums.DataTypeOctetString, nil
	}
	return enums.DataTypeNone, fmt.Errorf("unknown DLMS data type for %v", dt)
}

func GetDataTypeSize(dt enums.DataType) int {
	switch dt {
	case enums.DataTypeNone, enums.DataTypeArray, enums.DataTypeStructure, enums.DataTypeOctetString,
		enums.DataTypeString, enums.DataTypeStringUTF8, enums.DataTypeCompactArray:
		return -1
	case enums.DataTypeBoolean, enums.DataTypeInt8, enums.DataTypeUint8, enums.DataTypeEnum,
		enums.DataTypeBcd, enums.DataTypeDeltaInt8, enums.DataTypeDeltaUint8:
		return 1
	case enums.DataTypeInt16, enums.DataTypeUint16, enums.DataTypeDeltaInt16, enums.DataTypeDeltaUint16:
		return 2
	case enums.DataTypeInt32, enums.DataTypeUint32, enums.DataTypeFloat32,
		enums.DataTypeDeltaInt32, enums.DataTypeDeltaUint32:
		return 4
	case enums.DataTypeInt64, enums.DataTypeUint64, enums.DataTypeFloat64:
		return 8
	case enums.DataTypeDate:
		return 5
	case enums.DataTypeTime:
		return 4
	case enums.DataTypeDateTime:
		return 12
	case enums.DataTypeBitString:
		return -1
	default:
		return -1
	}
}

func GetData(settings *settings.GXDLMSSettings, data *types.GXByteBuffer, info *GXDataInfo) (any, error) {
	var value any
	startIndex := data.Position()
	var err error
	if data.Available() == 0 {
		info.Complete = false
		return nil, nil
	}
	info.Complete = true
	knownType := info.Type != enums.DataTypeNone
	if !knownType {
		t, err := data.Uint8()
		if err != nil {
			return nil, err
		}
		info.Type = enums.DataType(t)
	}
	if info.Type == enums.DataTypeNone {
		if info.Xml != nil {
			info.Xml.AppendStringLine("<" + info.Xml.GetDataType(info.Type) + " />")
		}
		return value, nil
	}
	if data.Available() == 0 {
		info.Complete = false
		return nil, nil
	}
	switch info.Type {
	case enums.DataTypeArray, enums.DataTypeStructure:
		value, err = getArray(settings, data, info, startIndex)
		if err != nil {
			return nil, err
		}
	case enums.DataTypeBoolean:
		value = getBoolean(data, info)
	case enums.DataTypeBitString:
		value = getBitString(data, info)
	case enums.DataTypeInt32:
		value = getInt32(data, info)
	case enums.DataTypeInt8:
		value = getInt8(data, info)
	case enums.DataTypeInt16:
		value = getInt16(data, info)
	case enums.DataTypeInt64:
		value = getInt64(data, info)
	case enums.DataTypeUint8:
		value = getUint8(data, info)
	case enums.DataTypeUint16:
		value = getUint16(data, info)
	case enums.DataTypeUint32:
		value = getUint32(data, info)
	case enums.DataTypeUint64:
		value = getUint64(data, info)
	case enums.DataTypeEnum:
		value = getEnum(data, info)
	case enums.DataTypeFloat32:
		value = getFloat32(data, info)
	case enums.DataTypeFloat64:
		value = getFloat64(data, info)
	case enums.DataTypeString:
		value = getString(data, info, knownType)
	case enums.DataTypeStringUTF8:
		value = getUtf8String(data, info, knownType)
	case enums.DataTypeOctetString:
		value = getOctetString(settings, data, info, knownType)
	case enums.DataTypeBcd:
		value = getBcd(data, info)
	case enums.DataTypeTime:
		value, err = getTime(data, info)
	case enums.DataTypeDate:
		value, err = getDate(data, info)
	case enums.DataTypeDateTime:
		value, err = getDateTime(settings, data, info)
	case enums.DataTypeDeltaInt8:
		value = types.GXDeltaInt8{Value: getInt8(data, info)}
	case enums.DataTypeDeltaInt16:
		value = types.GXDeltaInt16{Value: getInt16(data, info)}
	case enums.DataTypeDeltaInt32:
		value = types.GXDeltaInt32{Value: getInt32(data, info)}
	case enums.DataTypeDeltaUint8:
		value = types.GXDeltaUInt8{Value: getUint8(data, info)}
	case enums.DataTypeDeltaUint16:
		value = types.GXDeltaUInt16{Value: getUint16(data, info)}
	case enums.DataTypeDeltaUint32:
		value = types.GXDeltaUInt32{Value: getUint32(data, info)}
	default:
		err = fmt.Errorf("unsupported DLMS data type: %v", info.Type)
	}
	return value, err
}

func getOctetString(settings *settings.GXDLMSSettings, buff *types.GXByteBuffer, info *GXDataInfo, knownType bool) any {
	len_ := 0
	if knownType {
		len_ = buff.Size()
	} else {
		len_, _ = types.GetObjectCount(buff)
		// If there is not enough data available.
		if buff.Available() < len_ {
			info.Complete = false
			return nil
		}
	}
	value := make([]byte, len_)
	buff.Get(value)
	if info.Xml != nil {
		if info.Xml.Comments && len(value) != 0 {
			// This might be logical name.
			if len(value) == 6 && value[5] == 0xFF {
				ret, _ := helpers.ToLogicalName(value)
				info.Xml.AppendComment(ret)
			} else {
				isString := true
				//Try to change octet string to DateTime, Date or time.
				if len(value) == 12 || len(value) == 5 || len(value) == 4 {
					var type_ enums.DataType
					if len(value) == 12 {
						type_ = enums.DataTypeDateTime
					} else if len(value) == 5 {
						type_ = enums.DataTypeDate
					} else {
						type_ = enums.DataTypeTime
					}
					ret, err := ChangeTypeFromByteArray(settings, value, type_)
					if err == nil && ret != nil {
						dt := ret.(types.GXDateTime)
						info.Xml.AppendComment(dt.ToFormatMeterString(nil))
						isString = false
					} else {
						isString = true
					}
				}
				if isString {
					for _, it := range value {
						if it < 32 || it > 126 {
							isString = false
							break
						}
					}
				}
				if isString {
					info.Xml.AppendComment(string(value))
				}
			}
		}
		info.Xml.AppendLine(info.Xml.GetDataType(info.Type), "Value", ToHex(value, false, 0, len(value)))
	}
	return value
}

func getString(buff *types.GXByteBuffer, info *GXDataInfo, knownType bool) any {
	var err error
	len_ := 0
	if knownType {
		len_ = buff.Size()
	} else {
		len_, err = types.GetObjectCount(buff)
		if err != nil {
			info.Complete = false
			return nil
		}
		// If there is not enough data available.
		if buff.Available() < len_ {
			info.Complete = false
			return nil
		}
	}
	var value string
	if len_ > 0 {
		tmp := make([]byte, len_)
		buff.Get(tmp)
		value = string(tmp)
	} else {
		value = ""
	}
	if info.Xml != nil {
		if info.Xml.ShowStringAsHex {
			info.Xml.AppendLine(info.Xml.GetDataType(info.Type), "Value", ToHex(buff.Array(), false, buff.Position()-len_, len_))
		} else {
			info.Xml.AppendLine(info.Xml.GetDataType(info.Type), "Value", strings.Replace(value, "\"", "'", -1))
		}
	}
	return value
}

func getUint32(buff *types.GXByteBuffer, info *GXDataInfo) uint32 {
	// If there is not enough data available.
	if buff.Available() < 4 {
		info.Complete = false
		return 0
	}
	value, _ := buff.Uint32()
	if info.Xml != nil {
		info.Xml.AppendLine(info.Xml.GetDataType(info.Type), "Value", value)
	}
	return value
}

// /<summary>
// /Get Int32 value from DLMS data.
// /</summary>
// /<param name="buff">
// /Received DLMS data.
// /</param>
// /<param name="info">
// /Data info.
// /</param>
// /<returns>
// /Parsed Int32 value.
// /</returns>
func getInt32(buff *types.GXByteBuffer, info *GXDataInfo) int32 {
	// If there is not enough data available.
	if buff.Available() < 4 {
		info.Complete = false
		return 0
	}
	value, _ := buff.Int32()
	if info.Xml != nil {
		info.Xml.AppendLine(info.Xml.GetDataType(info.Type), "Value", value)
	}
	return value
}

// /<summary>
// /Get bit string value from DLMS data.
// /</summary>
// /<param name="buff">
// /Received DLMS data.
// /</param>
// /<param name="info">
// /Data info.
// /</param>
// /<returns>
// /Parsed bit string value.
// /</returns>
func getBitString(buff *types.GXByteBuffer, info *GXDataInfo) types.GXBitString {
	cnt, _ := types.GetObjectCount(buff)
	t := float64(cnt)
	t /= 8
	if cnt%8 != 0 {
		t++
	}
	byteCnt := int(math.Floor(t))
	// If there is not enough data available.
	if buff.Available() < byteCnt {
		info.Complete = false
		return types.GXBitString{}
	}
	ret := make([]byte, byteCnt)
	buff.Get(ret)
	bs, _ := types.NewGXBitString(ret, 0)
	return *bs
}

// getBoolean retrieves a Boolean value from DLMS data.
func getBoolean(buff *types.GXByteBuffer, info *GXDataInfo) any {
	// If there is not enough data available.
	if buff.Available() < 1 {
		info.Complete = false
		return nil
	}
	value, _ := buff.Uint8()
	if info.Xml != nil {
		info.Xml.AppendLine(info.Xml.GetDataType(info.Type), "Value", strconv.FormatBool(value != 0))
	}
	return value != 0
}

// getUint8 retrieves a UInt8 value from DLMS data.
func getUint8(buff *types.GXByteBuffer, info *GXDataInfo) uint8 {
	// If there is not enough data available.
	if buff.Available() < 1 {
		info.Complete = false
		return 0
	}
	value, _ := buff.Uint8()
	if info.Xml != nil {
		info.Xml.AppendLine(info.Xml.GetDataType(info.Type), "Value", info.Xml.IntegerToHex(value, 2, false))
	}
	return value
}

// getInt16 retrieves an Int16 value from DLMS data.
func getInt16(buff *types.GXByteBuffer, info *GXDataInfo) int16 {
	// If there is not enough data available.
	if buff.Available() < 2 {
		info.Complete = false
		return 0
	}
	value, _ := buff.Int16()
	if info.Xml != nil {
		info.Xml.AppendLine(info.Xml.GetDataType(info.Type), "Value", info.Xml.IntegerToHex(value, 4, false))
	}
	return value
}

// getInt8 retrieves an Int8 value from DLMS data.
func getInt8(buff *types.GXByteBuffer, info *GXDataInfo) int8 {
	// If there is not enough data available.
	if buff.Available() < 1 {
		info.Complete = false
		return 0
	}
	value, _ := buff.Int8()
	if info.Xml != nil {
		info.Xml.AppendLine(info.Xml.GetDataType(info.Type), "Value", info.Xml.IntegerToHex(value, 2, false))
	}
	return value
}

// getBcd retrieves a BCD value from DLMS data.
func getBcd(buff *types.GXByteBuffer, info *GXDataInfo) any {
	// If there is not enough data available.
	if buff.Available() < 1 {
		info.Complete = false
		return nil
	}
	value, _ := buff.Uint8()
	if info.Xml != nil {
		info.Xml.AppendLine(info.Xml.GetDataType(info.Type), "Value", info.Xml.IntegerToHex(value, 2, false))
	}
	return value
}

// getUtf8String retrieves a UTF8 string from DLMS data.
func getUtf8String(buff *types.GXByteBuffer, info *GXDataInfo, knownType bool) any {
	var value any
	var len_ int
	if knownType {
		len_ = buff.Size()
	} else {
		len_, _ = types.GetObjectCount(buff)
		// If there is not enough data available.
		if buff.Available() < len_ {
			info.Complete = false
			return nil
		}
	}
	if len_ > 0 {
		tmp := make([]byte, len_)
		buff.Get(tmp)
		value = string(tmp)
	} else {
		value = ""
	}
	if info.Xml != nil {
		info.Xml.AppendLine(info.Xml.GetDataType(info.Type), "Value", value)
	}
	return value
}

func getArray(opts *settings.GXDLMSSettings, buff *types.GXByteBuffer, info *GXDataInfo, index int) (any, error) {
	if info.Count == 0 {
		info.Count, _ = types.GetObjectCount(buff)
	}
	if info.Xml != nil {
		info.Xml.AppendStartTag(settings.DataTypeOffset+int(info.Type), "Qty", info.Xml.IntegerToHex(info.Count, 2, false), true)
	}
	size := buff.Size() - buff.Position()
	if info.Count != 0 && size < 1 {
		info.Complete = false
		return nil, nil
	}
	startIndex := index
	var arr []any
	var pos int
	// Position where last row was found. Cache uses this info.
	for pos = info.Index; pos != info.Count; pos++ {
		info2 := &GXDataInfo{}
		info2.Xml = info.Xml
		tmp, err := GetData(opts, buff, info2)
		if err != nil {
			return nil, err
		}
		if !info2.Complete {
			if info.Xml != nil {
				info.Xml.AppendComment(fmt.Sprintf("Error: Not enough data. %d rows are missing.", info.Count-pos))
			}
			buff.SetPosition(startIndex)
			info.Complete = false
			break
		}
		if info2.Count == info2.Index {
			startIndex = buff.Position()
			arr = append(arr, tmp)
		}
	}
	if info.Xml != nil {
		info.Xml.AppendEndTag(settings.DataTypeOffset+int(info.Type), true)
	}
	info.Index = pos
	var ret any
	if info.Type == enums.DataTypeArray {
		ret = types.GXArray(arr)
	} else {
		ret = types.GXStructure(arr)
	}
	if opts != nil && opts.Version == 8 {
		return ret, nil
	}
	return ret, nil
}

// getTime retrieves a Time value from DLMS data.
func getTime(buff *types.GXByteBuffer, info *GXDataInfo) (any, error) {
	if buff.Available() < 4 {
		// If there is not enough data available.
		info.Complete = false
		return nil, nil
	}
	var str string
	if info.Xml != nil {
		str = ToHex(buff.Array(), false, buff.Position(), 4)
	}
	// Get time.
	hour, _ := buff.Uint8()
	minute, _ := buff.Uint8()
	second, _ := buff.Uint8()
	ms, _ := buff.Uint8()
	if ms != 0xFF {
		ms *= 10
	} else {
		ms = 0xff
	}
	value, err := types.NewGXTime(int(hour), int(minute), int(second), int(ms))
	if err != nil && info.Xml == nil {
		return nil, err
	}
	if info.Xml != nil {
		if value != nil {
			info.Xml.AppendComment(value.ToFormatString(nil, false))
		}
		info.Xml.AppendLine(info.Xml.GetDataType(info.Type), "Value", str)
	}
	return *value, nil
}

// getDate retrieves a Date value from DLMS data.
func getDate(buff *types.GXByteBuffer, info *GXDataInfo) (any, error) {
	if buff.Available() < 5 {
		// If there is not enough data available.
		info.Complete = false
		return nil, nil
	}
	var str string
	if info.Xml != nil {
		str = ToHex(buff.Array(), false, buff.Position(), 5)
	}
	// Get year.
	year, _ := buff.Uint16()
	// Get month
	month, _ := buff.Uint8()
	extra := enums.DateTimeExtraInfoNone
	skip := enums.DateTimeSkipsNone
	switch month {
	case 0, 0xFF:
		month = 1
		skip |= enums.DateTimeSkipsMonth
	case 0xFE:
		//Daylight savings begin.
		month = 1
		extra |= enums.DateTimeExtraInfoDstBegin
	case 0xFD:
		// Daylight savings end.
		month = 1
		extra |= enums.DateTimeExtraInfoDstEnd
	}
	// Get day
	day, _ := buff.Uint8()
	if day == 0xFD {
		// 2nd last day of month.
		day = 1
		extra |= enums.DateTimeExtraInfoLastDay2
	} else if day == 0xFE {
		//Last day of month
		day = 1
		extra |= enums.DateTimeExtraInfoLastDay
	} else if day < 1 || day == 0xFF {
		day = 1
		skip |= enums.DateTimeSkipsDay
	}
	value, err := types.NewGXDate(int(year), int(month), int(day))
	if err != nil && info.Xml == nil {
		return nil, err
	}
	value.Extra = extra
	value.Skip |= skip
	// Skip week day
	ret, _ := buff.Uint8()
	if ret == 0xFF {
		value.Skip |= enums.DateTimeSkipsDayOfWeek
	}
	if info.Xml != nil {
		if value != nil {
			info.Xml.AppendComment(value.ToFormatString(nil, false))
		}
		info.Xml.AppendLine(info.Xml.GetDataType(info.Type), "Value", str)
	}
	return *value, nil
}

// getDateTime retrieves a DateTime value from DLMS data.
func getDateTime(settings *settings.GXDLMSSettings, buff *types.GXByteBuffer, info *GXDataInfo) (any, error) {

	// If there is not enough data available.
	if buff.Available() < 12 {
		//If time.
		if buff.Available() < 5 {
			return getTime(buff, info)
		} else if buff.Available() < 6 {
			//If date.
			return getDate(buff, info)
		}
		info.Complete = false
		return nil, nil
	}
	var str string
	if info.Xml != nil {
		str = ToHex(buff.Array(), false, buff.Position(), 12)
	}
	dt := types.GXDateTime{}
	//Get year.
	year, _ := buff.Uint16()
	if year == 0xFFFF || year == 0 {
		year = uint16(time.Now().Year())
		dt.Skip |= enums.DateTimeSkipsYear
	}
	//Get month
	month, _ := buff.Uint8()
	switch month {
	case 0, 0xFF:
		month = 1
		dt.Skip |= enums.DateTimeSkipsMonth
	case 0xFE:
		//Daylight savings begin.
		month = 1
		dt.Extra |= enums.DateTimeExtraInfoDstBegin
	case 0xFD:
		// Daylight savings end.
		month = 1
		dt.Extra |= enums.DateTimeExtraInfoDstEnd
	}
	day, _ := buff.Uint8()
	if day == 0xFD {
		// 2nd last day of month.
		day = 1
		dt.Extra |= enums.DateTimeExtraInfoLastDay2
	} else if day == 0xFE {
		//Last day of month
		day = 1
		dt.Extra |= enums.DateTimeExtraInfoLastDay
	} else if day < 1 || day == 0xFF {
		day = 1
		dt.Skip |= enums.DateTimeSkipsDay
	}

	//Skip week day.
	wd, _ := buff.Uint8()
	if wd == 0xFF {
		dt.Skip |= enums.DateTimeSkipsDayOfWeek
	} else {
		dt.DayOfWeek = int(wd)
	}
	//Get time.
	hours, _ := buff.Uint8()
	if hours == 0xFF {
		hours = 0
		dt.Skip |= enums.DateTimeSkipsHour
	}
	minutes, _ := buff.Uint8()
	if minutes == 0xFF {
		minutes = 0
		dt.Skip |= enums.DateTimeSkipsMinute
	}
	seconds, _ := buff.Uint8()
	if seconds == 0xFF {
		seconds = 0
		dt.Skip |= enums.DateTimeSkipsSecond
	}
	milliseconds, _ := buff.Uint8()
	if milliseconds != 0xFF {
		milliseconds *= 10
	} else {
		milliseconds = 0
		dt.Skip |= enums.DateTimeSkipsMs
	}
	deviation, _ := buff.Int16()
	ret, _ := buff.Uint8()
	dt.Status = enums.ClockStatus(ret)
	if settings != nil && settings.UseUtc2NormalTime && deviation != -32768 {
		deviation = -deviation
	}
	//0x8000 == -32768
	//deviation = -1 if skipped.
	if deviation != -1 && deviation != -32768 && year != 1 && (dt.Skip&enums.DateTimeSkipsYear) == 0 {
		dt.Value = time.Date(int(year), time.Month(month), int(day),
			int(hours), int(minutes), int(seconds),
			int(milliseconds)*int(time.Millisecond), time.UTC).Add(time.Duration(deviation) * time.Minute)
	} else {
		//Use current time if deviation is not defined.
		dt.Skip |= enums.DateTimeSkipsDeviation
		dt.Value = time.Date(int(year), time.Month(month), int(day),
			int(hours), int(minutes), int(seconds),
			int(milliseconds)*int(time.Millisecond), time.Local)
	}
	if info.Xml != nil {
		info.Xml.AppendComment(dt.ToFormatMeterString(nil))
		info.Xml.AppendLine(info.Xml.GetDataType(info.Type), "Value", str)
	}
	return dt, nil
}

// getDouble retrieves a Double value from DLMS data.
func getFloat64(buff *types.GXByteBuffer, info *GXDataInfo) float64 {
	// If there is not enough data available.
	if buff.Available() < 8 {
		info.Complete = false
		return 0.0
	}
	value, _ := buff.Double()
	if info.Xml != nil {
		if info.Xml.Comments {
			info.Xml.AppendComment(strconv.FormatFloat(value, 'f', -1, 64))
		}
		tmp := types.NewGXByteBuffer()
		SetData(nil, tmp, enums.DataTypeFloat64, value)
		info.Xml.AppendLine(info.Xml.GetDataType(info.Type), "Value", ToHex(tmp.Array(), false, 1, tmp.Size()-1))
	}
	return value
}

// getFloat32 retrieves a Float32 value from DLMS data.
// /<param name="buff">
// /Received DLMS data.
// /</param>
// /<param name="info">
// /Data info.
// /</param>
// /<returns>
// /Parsed Float value.
// /</returns>
func getFloat32(buff *types.GXByteBuffer, info *GXDataInfo) float32 {
	// If there is not enough data available.
	if buff.Available() < 4 {
		info.Complete = false
		return 0.0
	}
	value, _ := buff.Float()
	if info.Xml != nil {
		if info.Xml.Comments {
			info.Xml.AppendComment(strconv.FormatFloat(float64(value), 'f', -1, 32))
		}
		tmp := types.NewGXByteBuffer()
		SetData(nil, tmp, enums.DataTypeFloat32, value)
		info.Xml.AppendLine(info.Xml.GetDataType(info.Type), "Value", ToHex(tmp.Array(), false, 1, tmp.Size()-1))
	}
	return value
}

// getEnum retrieves an Enum value from DLMS data.
// /<param name="buff">
// /Received DLMS data.
// /</param>
// /<param name="info">
// /Data info.
// /</param>
// /<returns>
// /Parsed Enum value.
// /</returns>
func getEnum(buff *types.GXByteBuffer, info *GXDataInfo) types.GXEnum {
	// If there is not enough data available.
	if buff.Available() < 1 {
		info.Complete = false
		return types.GXEnum{}
	}
	value, _ := buff.Uint8()
	if info.Xml != nil {
		info.Xml.AppendLine(info.Xml.GetDataType(info.Type), "Value", info.Xml.IntegerToHex(value, 2, false))
	}
	return types.GXEnum{Value: value}
}

// getUint64 retrieves a Uint64 value from DLMS data.
// /<param name="buff">
// /Received DLMS data.
// /</param>
// /<param name="info">
// /Data info.
// /</param>
// /<returns>
// /Parsed Uint64 value.
// /</returns>
func getUint64(buff *types.GXByteBuffer, info *GXDataInfo) uint64 {
	// If there is not enough data available.
	if buff.Available() < 8 {
		info.Complete = false
		return 0
	}
	value, _ := buff.Uint64()
	if info.Xml != nil {
		info.Xml.AppendLine(info.Xml.GetDataType(info.Type), "Value", info.Xml.IntegerToHex(value, 16, false))
	}
	return value
}

// getInt64 retrieves an Int64 value from DLMS data.
// /<param name="buff">
// /Received DLMS data.
// /</param>
// /<param name="info">
// /Data info.
// /</param>
// /<returns>
// /Parsed Int64 value.
// /</returns>
func getInt64(buff *types.GXByteBuffer, info *GXDataInfo) int64 {
	// If there is not enough data available.
	if buff.Available() < 8 {
		info.Complete = false
		return 0
	}
	value, _ := buff.Int64()
	if info.Xml != nil {
		info.Xml.AppendLine(info.Xml.GetDataType(info.Type), "Value", info.Xml.IntegerToHex(value, 16, false))
	}
	return value
}

// getUint16 retrieves a UInt16 value from DLMS data.
// /<param name="buff">
// /Received DLMS data.
// /</param>
// /<param name="info">
// /Data info.
// /</param>
// /<returns>
// /Parsed UInt16 value.
// /</returns>
func getUint16(buff *types.GXByteBuffer, info *GXDataInfo) uint16 {
	// If there is not enough data available.
	if buff.Available() < 2 {
		info.Complete = false
		return 0
	}
	value, _ := buff.Uint16()
	if info.Xml != nil {
		info.Xml.AppendLine(info.Xml.GetDataType(info.Type), "Value", info.Xml.IntegerToHex(value, 4, false))
	}
	return value
}

func SetData(s any, buff *types.GXByteBuffer, dt enums.DataType, value any) error {
	conf := s.(*settings.GXDLMSSettings)
	if err := buff.SetUint8(uint8(dt)); err != nil {
		return err
	}
	switch dt {
	case enums.DataTypeNone:
		return nil
	case enums.DataTypeArray, enums.DataTypeStructure:
		var items []any
		switch v := value.(type) {
		case []any:
			items = v
		case types.GXArray:
			items = []any(v)
		case types.GXStructure:
			items = []any(v)
		default:
			return fmt.Errorf("invalid value for %v: %T", dt, value)
		}
		types.SetObjectCount(len(items), buff)
		for _, it := range items {
			dt2, err := GetDLMSDataType(reflect.TypeOf(it))
			if err != nil {
				return err
			}
			if err = SetData(conf, buff, dt2, it); err != nil {
				return err
			}
		}
		return nil
	case enums.DataTypeBoolean:
		b, _ := value.(bool)
		if b {
			return buff.SetUint8(1)
		}
		return buff.SetUint8(0)
	case enums.DataTypeInt8:
		return buff.SetUint8(uint8(value.(int8)))
	case enums.DataTypeUint8, enums.DataTypeEnum:
		return buff.SetUint8(value.(uint8))
	case enums.DataTypeInt16:
		return buff.SetInt16(value.(int16))
	case enums.DataTypeUint16:
		return buff.SetUint16(value.(uint16))
	case enums.DataTypeInt32:
		return buff.SetInt32(value.(int32))
	case enums.DataTypeUint32:
		return buff.SetUint32(value.(uint32))
	case enums.DataTypeInt64:
		return buff.SetInt64(value.(int64))
	case enums.DataTypeUint64:
		return buff.SetInt64(int64(value.(uint64)))
	case enums.DataTypeFloat32:
		return buff.SetFloat(value.(float32))
	case enums.DataTypeFloat64:
		return buff.SetDouble(value.(float64))
	case enums.DataTypeString, enums.DataTypeStringUTF8:
		s := value.(string)
		types.SetObjectCount(len(s), buff)
		return buff.Set([]byte(s))
	case enums.DataTypeOctetString:
		var b []byte
		switch v := value.(type) {
		case []byte:
			b = v
		case string:
			tmp, err := HexToBytes(v)
			if err != nil {
				return err
			}
			b = tmp
		default:
			return fmt.Errorf("invalid octet string value: %T", value)
		}
		types.SetObjectCount(len(b), buff)
		return buff.Set(b)
	case enums.DataTypeBitString:
		bs, ok := value.(*types.GXBitString)
		if !ok {
			v, ok2 := value.(types.GXBitString)
			if !ok2 {
				return fmt.Errorf("invalid bit string value: %T", value)
			}
			bs = &v
		}
		types.SetObjectCount(len(bs.Value())+1, buff)
		if err := buff.SetUint8(uint8(bs.PadBits())); err != nil {
			return err
		}
		return buff.Set(bs.Value())
	case enums.DataTypeDate, enums.DataTypeTime, enums.DataTypeDateTime:
		return fmt.Errorf("%v serialization is not implemented", dt)
	default:
		return fmt.Errorf("unsupported DLMS data type: %v", dt)
	}
}

func ChangeTypeFromByteArray(settings *settings.GXDLMSSettings, value []byte, dt enums.DataType) (any, error) {
	bb := types.NewGXByteBufferWithData(value)
	return ChangeType(settings, bb, dt)
}

func ChangeType(settings *settings.GXDLMSSettings, value *types.GXByteBuffer, dt enums.DataType) (any, error) {
	if value == nil {
		return nil, errors.New("value is nil")
	}
	info := GXDataInfo{}
	info.Type = dt
	v, err := GetData(settings, value, &info)
	if err != nil {
		return nil, err
	}
	if dt == enums.DataTypeNone {
		return v, nil
	}
	if info.Type == dt {
		return v, nil
	}
	return v, nil
}

func GeneralizedTime(date *time.Time) string {
	t := date.UTC()
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%04d", t.Year()))
	sb.WriteString(fmt.Sprintf("%02d", int(t.Month())))
	sb.WriteString(fmt.Sprintf("%02d", t.Day()))
	sb.WriteString(fmt.Sprintf("%02d", t.Hour()))
	sb.WriteString(fmt.Sprintf("%02d", t.Minute()))
	sb.WriteString(fmt.Sprintf("%02d", t.Second()))
	sb.WriteByte('Z')
	return sb.String()
}

func EncryptManufacturer(flagName string) uint16 {
	if len(flagName) != 3 {
		return 0
	}
	f := strings.ToUpper(flagName)
	return uint16(f[0]-0x40)<<10 | uint16(f[1]-0x40)<<5 | uint16(f[2]-0x40)
}

func DecryptManufacturer(value uint16) string {
	return string([]byte{byte((value>>10)&0x1F + 0x40), byte((value>>5)&0x1F + 0x40), byte(value&0x1F + 0x40)})
}

func getSerialNumber(st []byte, isIdis bool) int64 {
	if len(st) < 8 {
		return 0
	}
	if isIdis {
		return int64(st[4])<<24 | int64(st[5])<<16 | int64(st[6])<<8 | int64(st[7])
	}
	return int64(st[0])<<24 | int64(st[1])<<16 | int64(st[2])<<8 | int64(st[3])
}

func SystemTitleToString(standard enums.Standard, st []byte, addComments bool) string {
	if len(st) != 8 {
		return buffer.ToHex(st, true)
	}
	sb := strings.Builder{}
	manufacturer := DecryptManufacturer(uint16(st[0])<<8 | uint16(st[1]))
	serial := getSerialNumber(st, standard == enums.StandardIdis)
	if addComments {
		sb.WriteString(manufacturer)
		sb.WriteString("\n")
		sb.WriteString(fmt.Sprintf("%d", serial))
	} else {
		sb.WriteString(manufacturer)
		sb.WriteString(fmt.Sprintf("%d", serial))
	}
	return sb.String()
}

func FromBase64(value string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(value)
}

func ToBase64(value []byte) string {
	return base64.StdEncoding.EncodeToString(value)
}

func Contains[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func InsertAt[T any](slice []T, index int, value T) ([]T, error) {
	if index < 0 || index > len(slice) {
		return nil, gxcommon.ErrArgumentOutOfRange
	}
	slice = append(
		slice[:index],
		append([]T{value}, slice[index:]...)...,
	)
	return slice, nil
}

// Remove removes first occurrence of value from slice and returns modified slice.
func Remove[T comparable](slice []T, value T) []T {
	for i, v := range slice {
		if v == value {
			slice = append(slice[:i], slice[i+1:]...)
			return slice
		}
	}
	return slice
}

func AnyToDouble(value any) float64 {
	switch v := value.(type) {
	case float64:
		return v
	case float32:
		return float64(v)
	case int64:
		return float64(v)
	case uint64:
		return float64(v)
	case int32:
		return float64(v)
	case uint32:
		return float64(v)
	case int16:
		return float64(v)
	case uint16:
		return float64(v)
	case int8:
		return float64(v)
	case uint8:
		return float64(v)
	default:
		panic(fmt.Sprintf("unsupported type: %T", value))
	}
}

var rfc3394IV = []byte{0xA6, 0xA6, 0xA6, 0xA6, 0xA6, 0xA6, 0xA6, 0xA6}

// Encrypt returns the encrypted data using AES key wrap algorithm defined in RFC 3394.
//
// Parameters:
//
//	kek: Key Encrypting Key.
//	data: Data to be encrypted.
//
// Returns:
//
//	Encrypted data.
func Encrypt(kek []byte, data []byte) ([]byte, error) {
	if len(kek) != 16 && len(kek) != 32 {
		return nil, errors.New("invalid key encrypting key")
	}
	if len(data) != 16 && len(data) != 32 {
		return nil, errors.New("invalid data")
	}

	n := len(data) / 8
	block := make([]byte, len(data)+len(rfc3394IV))
	copy(block, rfc3394IV)
	copy(block[len(rfc3394IV):], data)

	buf := make([]byte, 16)
	cipher, err := aes.NewCipher(kek)
	if err != nil {
		return nil, err
	}
	for j := 0; j < 6; j++ {
		for i := 1; i <= n; i++ {
			copy(buf[:8], block[:8])
			copy(buf[8:], block[8*i:8*i+8])
			cipher.Encrypt(buf, buf)
			t := n*j + i
			for k := 1; t != 0; k++ {
				buf[8-k] ^= byte(t)
				t >>= 8
			}
			copy(block[:8], buf[:8])
			copy(block[8*i:8*i+8], buf[8:])
		}
	}
	return block, nil
}

// Decrypt returns the decrypted data using AES key wrap algorithm defined in RFC 3394.
func Decrypt(kek []byte, input []byte) ([]byte, error) {
	if len(kek) != 16 && len(kek) != 32 {
		return nil, errors.New("invalid key encrypting key")
	}
	if len(input) < len(rfc3394IV)+8 {
		return nil, errors.New("invalid data")
	}

	block := make([]byte, len(input)-len(rfc3394IV))
	a := make([]byte, len(rfc3394IV))
	buf := make([]byte, 16)
	copy(a, input[:len(rfc3394IV)])
	copy(block, input[len(rfc3394IV):])

	n := len(input)/8 - 1
	if n == 0 {
		n = 1
	}
	cipher, err := aes.NewCipher(kek)
	if err != nil {
		return nil, err
	}
	for j := 5; j >= 0; j-- {
		for i := n; i >= 1; i-- {
			copy(buf[:8], a)
			copy(buf[8:], block[8*(i-1):8*(i-1)+8])
			t := n*j + i
			for k := 1; t != 0; k++ {
				buf[8-k] ^= byte(t)
				t >>= 8
			}
			cipher.Decrypt(buf, buf)
			copy(a, buf[:8])
			copy(block[8*(i-1):8*(i-1)+8], buf[8:])
		}
	}
	if !bytes.Equal(a, rfc3394IV) {
		return nil, errors.New("AES key wrapping failed.")
	}
	return block, nil
}
