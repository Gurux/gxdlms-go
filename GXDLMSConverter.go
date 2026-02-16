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
	"embed"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/objects"
	"github.com/Gurux/gxdlms-go/types"
	"golang.org/x/text/language"
)

//go:embed resources/*
var resourcesFS embed.FS

// DLMS Converter is used to get string value for enumeration types.
type GXDLMSConverter struct {
	// Collection of standard OBIS codes.
	codes    GXStandardObisCodeCollection
	Standard enums.Standard
}

// NewGXDLMSConverter returns a new instance of GXDLMSConverter
// standard: Used standard.
func NewGXDLMSConverter(standard enums.Standard) *GXDLMSConverter {
	return &GXDLMSConverter{Standard: standard}
}

// Converts byte array to logical name string.
func ToLogicalName(value any) (string, error) {
	return helpers.ToLogicalName(value)
}

// Converts logical name string to byte array.
func LogicalNameToBytes(value string) ([]byte, error) {
	return helpers.LogicalNameToBytes(value)
}

// GetCode returns the get OBIS code part.value: OBIS code part.
func (g *GXDLMSConverter) getCode(value string) string {
	var index = strings.IndexAny(value, ".-,")
	if index != -1 {
		return value[0:index]
	}
	return value
}

// ReadStandardObisInfo returns the read standard OBIS code information from the file.codes: Collection of standard OBIS codes.
func (g *GXDLMSConverter) readStandardObisInfo() error {
	if g.Standard != enums.StandardDLMS {
		err := g.GetObjects()
		if err != nil {
			return err
		}
	}
	data, err := resourcesFS.ReadFile("resources/OBISCodes.txt")
	if err != nil {
		return err
	}
	rows := strings.Split(string(data), "\n")
	for _, it := range rows {
		// Skip comments and empty lines.
		if !strings.HasPrefix(it, "#") && it != "" {
			items := strings.Split(it, ";")
			obis := strings.Split(items[0], ".")
			desc := items[3] + "; " + items[4] + "; " + items[5] + "; " + items[6] + "; " + items[7]
			code := GXStandardObisCode{OBIS: obis, Description: desc,
				Interfaces: items[1],
				DataType:   items[2]}
			g.codes = append(g.codes, code)
		}
	}
	return nil
}

func (g *GXDLMSConverter) updateOBISCodeInfo(codes GXStandardObisCodeCollection,
	it objects.IGXDLMSBase,
	standard enums.Standard) error {
	base := it.Base()
	list, err := codes.Find(base.LogicalName(), base.ObjectType(), standard)
	if err != nil {
		return err
	}
	code := list[0]
	if base.Description == "" {
		base.Description = code.Description
		if standard == enums.StandardSaudiArabia {
			base.Description = strings.ReplaceAll(base.Description, "U(", "V(")
		}
	}
	// Update data type from DLMS standard.
	if standard != enums.StandardDLMS {
		d := list[len(list)-1]
		code.DataType = d.DataType
	}
	// If string is used
	datatype := code.DataType
	if code.UIDataType == "" {
		if strings.Contains(datatype, "10") {
			code.UIDataType = "10"
		} else if strings.Contains(datatype, "25") || strings.Contains(datatype, "26") {
			//If date time is used.
			code.UIDataType = "25"
		} else if strings.Contains(datatype, "9") {
			if equalsMask("0.0-64.96.7.10-14.255", base.LogicalName()) || equalsMask("0.0-64.0.1.5.0-99,255", base.LogicalName()) || equalsMask("0.0-64.0.1.2.0-99,255", base.LogicalName()) || equalsMask("1.0-64.0.1.2.0-99,255", base.LogicalName()) || equalsMask("1.0-64.0.1.5.0-99,255", base.LogicalName()) || equalsMask("1.0-64.0.9.0.255", base.LogicalName()) || equalsMask("1.0-64.0.9.6.255", base.LogicalName()) || equalsMask("1.0-64.0.9.7.255", base.LogicalName()) || equalsMask("1.0-64.0.9.13.255", base.LogicalName()) || equalsMask("1.0-64.0.9.14.255", base.LogicalName()) || equalsMask("1.0-64.0.9.15.255", base.LogicalName()) {
				//Time stamps of the billing periods objects (first scheme if there are two)
				code.UIDataType = "25"
			} else if equalsMask("1.0-64.0.9.1.255", base.LogicalName()) {
				//Local time
				code.UIDataType = "27"
			} else if equalsMask("1.0-64.0.9.2.255", base.LogicalName()) {
				//Local date
				code.UIDataType = "26"
			} else if equalsMask("1.0.0.2.0.255", base.LogicalName()) {
				//Active firmware identifier
				code.UIDataType = "10"
			}
		} else if base.ObjectType() == enums.ObjectTypeData && equalsMask("0.0.1.1.0.255", base.LogicalName()) {
			//Unix time
			code.UIDataType = "25"
		}
	}
	if code.DataType != "*" && code.DataType != "" && !strings.Contains(code.DataType, ",") {
		ret, err := strconv.Atoi(code.DataType)
		if err != nil {
			return err
		}
		dt := enums.DataType(ret)
		switch base.ObjectType() {
		case enums.ObjectTypeData:
		case enums.ObjectTypeRegister:
		case enums.ObjectTypeRegisterActivation:
		case enums.ObjectTypeExtendedRegister:
			base.SetDataType(2, dt)
		default:
		}
	}
	if code.UIDataType != "" {
		ret, err := strconv.Atoi(code.UIDataType)
		if err != nil {
			return err
		}
		uiType := enums.DataType(ret)
		switch base.ObjectType() {
		case enums.ObjectTypeData:
		case enums.ObjectTypeRegister:
		case enums.ObjectTypeRegisterActivation:
		case enums.ObjectTypeExtendedRegister:
			base.SetUIDataType(2, uiType)
		default:
		}
	}
	return nil
}

// GetDescription returns the get OBIS code description.logicalName: Logical name (OBIS code).
// ot: Object type.
// description: Description filter.
//
//	Returns:
//	    Array of descriptions that match given OBIS code.
func (g *GXDLMSConverter) GetDescription(logicalName string, ot enums.ObjectType, description string) ([]string, error) {
	if len(g.codes) == 0 {
		g.readStandardObisInfo()
	}
	tmp2, err := g.codes.Find(logicalName, ot, g.Standard)
	if err != nil {
		return nil, err
	}
	var list []string
	all := logicalName == ""
	for _, it := range tmp2 {
		if description != "" && !strings.Contains(strings.ToLower(it.Description), strings.ToLower(description)) {
			continue
		}
		if all {
			list = append(list, "A="+it.OBIS[0]+", B="+it.OBIS[1]+", C="+it.OBIS[2]+", D="+it.OBIS[3]+", E="+it.OBIS[4]+", F="+it.OBIS[5]+"\n"+it.Description)
		} else {
			if g.Standard == enums.StandardSaudiArabia {
				list = append(list, strings.Replace(it.Description, "U(", "V(", -1))
			} else {
				list = append(list, it.Description)
			}
		}
	}
	return list, nil
}

// GetAllowedDataTypes returns the get allowed data types for given OBIS code.
// logicalName: Logical name (OBIS code).
// type: Object type.
//
//	Returns:
//	    Array of data types that match given OBIS code.
func (g *GXDLMSConverter) GetAllowedDataTypes(logicalName string, ot enums.ObjectType) ([]enums.DataType, error) {
	if logicalName == "" {
		return nil, fmt.Errorf("Invalid logical name.")
	}
	codes, err := g.codes.Find(logicalName, ot, g.Standard)
	if err != nil {
		return nil, err
	}
	var list []enums.DataType
	for _, it := range codes {
		types := strings.Split(it.DataType, ",")
		for _, dt := range types {
			v, err := strconv.Atoi(dt)
			if err != nil {
				return nil, err
			}
			list = append(list, enums.DataType(v))
		}
	}
	return list, nil
}

// GetObisCodesByType returns the get example OBIS codes using object type as a filter.
// ot: Object type.
//
//	Returns:
//	    Array of OBIS codes and descriptions that match given type.
func (g *GXDLMSConverter) GetObisCodesByType(ot enums.ObjectType) ([]types.GXKeyValuePair[string, string], error) {
	var list []types.GXKeyValuePair[string, string]
	tmp, err := g.codes.Find("", ot, g.Standard)
	if err != nil {
		return nil, err
	}
	for _, it := range tmp {
		if it.Interfaces != "*" {
			obis := g.getCode(it.OBIS[0]) + "." + g.getCode(it.OBIS[1]) + "." + g.getCode(it.OBIS[2]) + "." + g.getCode(it.OBIS[3]) + "." + g.getCode(it.OBIS[4]) + "." + g.getCode(it.OBIS[5])
			tmp, err := g.codes.Find(obis, ot, g.Standard)
			if err != nil {
				return nil, err
			}
			list = append(list, types.GXKeyValuePair[string, string]{Key: obis, Value: tmp[0].Description})
		}
	}
	return list, nil
}

// UpdateOBISCodeInformation returns the update standard OBIS codes description and type if defined.target: COSEM object.
func (g *GXDLMSConverter) UpdateOBISCodeInformation(target objects.IGXDLMSBase) error {
	if len(g.codes) == 0 {
		err := g.readStandardObisInfo()
		if err != nil {
			return err
		}
	}
	return g.updateOBISCodeInfo(g.codes, target, g.Standard)
}

// UpdateOBISCodesInformation returns the update standard OBIS codes descriptions and type if defined.targets: Collection of COSEM objects to update.
func (g *GXDLMSConverter) UpdateOBISCodesInformation(targets objects.GXDLMSObjectCollection) error {
	var err error
	if len(g.codes) == 0 {
		err = g.readStandardObisInfo()
		if err != nil {
			return err
		}
	}
	for _, it := range targets {
		err = g.updateOBISCodeInfo(g.codes, it, g.Standard)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetUnit returns the returns unit text.value:
//
//	Returns:
func (g *GXDLMSConverter) GetUnit(value enums.Unit) string {
	switch value {
	case enums.UnitYear:
		return "Year"
	case enums.UnitMonth:
		return "Month"
	case enums.UnitWeek:
		return "Week"
	case enums.UnitDay:
		return "Day"
	case enums.UnitHour:
		return "Hour"
	case enums.UnitMinute:
		return "Minute"
	case enums.UnitSecond:
		return "Second"
	case enums.UnitPhaseAngleDegree:
		return "Phase angle degree rad*180/p"
	case enums.UnitTemperature:
		return "Temperature T degree centigrade"
	case enums.UnitLocalCurrency:
		return "Local currency"
	case enums.UnitLength:
		return "Length l meter m"
	case enums.UnitSpeed:
		return "Speed v m/s"
	case enums.UnitVolumeCubicMeter:
		return "Volume V m3"
	case enums.UnitCorrectedVolume:
		return "Corrected volume m3"
	case enums.UnitVolumeFluxHour:
		return "Volume flux m3/60*60s"
	case enums.UnitCorrectedVolumeFluxHour:
		return "Corrected volume flux m3/60*60s"
	case enums.UnitVolumeFluxDay:
		return "Volume flux m3/24*60*60s"
	case enums.UnitCorrectedVolumeFluxDay:
		return "Corrected volume flux m3/24*60*60s"
	case enums.UnitVolumeLiter:
		return "Volume 10-3 m3"
	case enums.UnitMassKg:
		return "Mass m kilogram kg"
	case enums.UnitForce:
		return "Force F newton N"
	case enums.UnitEnergy:
		return "Energy newtonmeter J = Nm = Ws"
	case enums.UnitPressurePascal:
		return "Pressure p pascal N/m2"
	case enums.UnitPressureBar:
		return "Pressure p bar 10-5 N/m2"
	case enums.UnitEnergyJoule:
		return "Energy joule J = Nm = Ws"
	case enums.UnitThermalPower:
		return "Thermal power J/60*60s"
	case enums.UnitActivePower:
		return "Active power P watt W = J/s"
	case enums.UnitApparentPower:
		return "Apparent power S"
	case enums.UnitReactivePower:
		return "Reactive power Q"
	case enums.UnitActiveEnergy:
		return "Active energy W*60*60s"
	case enums.UnitApparentEnergy:
		return "Apparent energy VA*60*60s"
	case enums.UnitReactiveEnergy:
		return "Reactive energy var*60*60s"
	case enums.UnitCurrent:
		return "Current I ampere A"
	case enums.UnitElectricalCharge:
		return "Electrical charge Q coulomb C = As"
	case enums.UnitVoltage:
		return "Voltage"
	case enums.UnitElectricalFieldStrength:
		return "Electrical field strength E V/m"
	case enums.UnitCapacity:
		return "Capacity C farad C/V = As/V"
	case enums.UnitResistance:
		return "Resistance R ohm = V/A"
	case enums.UnitResistivity:
		return "Resistivity"
	case enums.UnitMagneticFlux:
		return "Magnetic flux F weber Wb = Vs"
	case enums.UnitInduction:
		return "Induction T tesla Wb/m2"
	case enums.UnitMagnetic:
		return "Magnetic field strength H A/m"
	case enums.UnitInductivity:
		return "Inductivity L henry H = Wb/A"
	case enums.UnitFrequency:
		return "Frequency f"
	case enums.UnitActive:
		return "Active energy meter constant 1/Wh"
	case enums.UnitReactive:
		return "Reactive energy meter constant"
	case enums.UnitApparent:
		return "Apparent energy meter constant"
	case enums.UnitV260:
		return "V260*60s"
	case enums.UnitA260:
		return "A260*60s"
	case enums.UnitMassKgPerSecond:
		return "Mass flux kg/s"
	case enums.UnitConductance:
		return "Conductance siemens 1/ohm"
	case enums.UnitOtherUnit:
		return "Other Unit"
	case enums.UnitNoUnit:
		return "No Unit"
	case enums.UnitKelvin:
		return "Kelvin"
	case enums.UnitRU2h:
		return "RU2h"
	case enums.UnitRI2h:
		return "RI2h"
	case enums.UnitCubicMeterRV:
		return "Cubic meter RV"
	case enums.UnitPercentage:
		return "Percentage"
	case enums.UnitAmpereHour:
		return "Ampere hour"
	case enums.UnitEnergyPerVolume:
		return "Energy per volume"
	case enums.UnitWobbe:
		return "Wobbe"
	case enums.UnitMolePercent:
		return "Mole percent"
	case enums.UnitMassDensity:
		return "Mass density"
	case enums.UnitPascalSecond:
		return "Pascal second"
	case enums.UnitJouleKilogram:
		return "Joule kilogram"
	case enums.UnitPressureGramPerSquareCentimeter:
		return "Pressure gram per square centimeter"
	case enums.UnitPressureAtmosphere:
		return "Pressure atmosphere"
	case enums.UnitSignalStrengthMilliWatt:
		return "Signal strength milli watt"
	case enums.UnitSignalStrengthMicroVolt:
		return "Signal strength micro volt"
	case enums.UnitdB:
		return "dB"
	case enums.UnitInch:
		return "Inch"
	case enums.UnitFoot:
		return "Foot"
	case enums.UnitPound:
		return "Pound"
	case enums.UnitFahrenheit:
		return "Fahrenheit"
	case enums.UnitRankine:
		return "Rankine"
	case enums.UnitSquareInch:
		return "SquareInch"
	case enums.UnitSquareFoot:
		return "SquareFoot"
	case enums.UnitAcre:
		return "Acre"
	case enums.UnitCubicInch:
		return "CubicInch"
	case enums.UnitCubicFoot:
		return "CubicFoot"
	case enums.UnitAcreFoot:
		return "AcreFoot"
	case enums.UnitGallonImperial:
		return "Gallon imperial"
	case enums.UnitGallonUS:
		return "GallonUS"
	case enums.UnitPoundForce:
		return "Pound force"
	case enums.UnitPoundForcePerSquareInch:
		return "Pound force per square inch"
	case enums.UnitPoundPerCubicFoot:
		return "Pound per cubic foot"
	case enums.UnitPoundPerFootSecond:
		return "Pound per foot second"
	case enums.UnitSquareFootPerSecond:
		return "Square foot per second"
	case enums.UnitBritishThermalUnit:
		return "British thermal unit"
	case enums.UnitThermEU:
		return "Therm EU"
	case enums.UnitThermUS:
		return "Therm US"
	case enums.UnitBritishThermalUnitPerPound:
		return "British thermal unit per pound"
	case enums.UnitBritishThermalUnitPerCubicFoot:
		return "British thermal unit per cubic foot"
	case enums.UnitCubicFeet:
		return "Cubic feet"
	case enums.UnitFootPerSecond:
		return "Foot per second"
	case enums.UnitCubicFootPerSecond:
		return "Cubic foot per second"
	case enums.UnitCubicFootPerMin:
		return "Cubic foot per min"
	case enums.UnitCubicFootPerHour:
		return "Cubic foot per hour"
	case enums.UnitCubicFootPerDay:
		return "Cubic foot per day"
	case enums.UnitAcreFootPerSecond:
		return "Acre foot per second"
	case enums.UnitAcreFootPerMin:
		return "Acre foot per min"
	case enums.UnitAcreFootPerHour:
		return "Acre foot per hour"
	case enums.UnitAcreFootPerDay:
		return "Acre foot per day"
	case enums.UnitImperialGallon:
		return "Imperial gallon"
	case enums.UnitImperialGallonPerSecond:
		return "Imperial gallon per second"
	case enums.UnitImperialGallonPerMin:
		return "Imperial gallon per min"
	case enums.UnitImperialGallonPerHour:
		return "Imperial gallon per hour"
	case enums.UnitImperialGallonPerDay:
		return "Imperial gallon per day"
	case enums.UnitUSGallon:
		return "US gallon"
	case enums.UnitUSGallonPerSecond:
		return "US gallon per second"
	case enums.UnitUSGallonPerMin:
		return "US gallon per min"
	case enums.UnitUSGallonPerHour:
		return "US gallon per hour"
	case enums.UnitUSGallonPerDay:
		return "US gallon per day"
	case enums.UnitBritishThermalUnitPerSecond:
		return "British thermal unit per second"
	case enums.UnitBritishThermalUnitPerMinute:
		return "British thermal unit per minute"
	case enums.UnitBritishThermalUnitPerHour:
		return "British thermal unit per hour"
	case enums.UnitBritishThermalUnitPerDay:
		return "British thermal unit per day"
	}
	return ""
}

// GetDLMSDataType returns the get DLMS data type.value: Object
//
//	Returns:
//	    DLMS data type.
func (g *GXDLMSConverter) GetDLMSDataType(value any) (enums.DataType, error) {
	if value == "" {
		return enums.DataTypeNone, nil
	}
	return internal.GetDLMSDataType(reflect.TypeOf(value))
}

func (g *GXDLMSConverter) GetBytes(value any, dt enums.DataType) ([]byte, error) {
	bb := types.GXByteBuffer{}
	err := internal.SetData(nil, &bb, dt, value)
	return bb.Array(), err
}

// GetObjects returns the get country spesific OBIS codes.
// standard: Used standard.
//
//	Returns:
//	    Collection for special OBIC codes.
func (g *GXDLMSConverter) GetObjects() error {
	var rows []string
	var resourcesFS embed.FS
	var redsourcePath string
	if g.Standard == enums.StandardItaly {
	} else if g.Standard == enums.StandardIndia {
		redsourcePath = "resources/Italy.txt"
	} else if g.Standard == enums.StandardSaudiArabia {
		redsourcePath = "resources/SaudiArabia.txt"
	} else if g.Standard == enums.StandardSpain {
		redsourcePath = "resources/Spain.txt"
	} else {
		return nil
	}
	data, err := resourcesFS.ReadFile(redsourcePath)
	if err != nil {
		return err
	}
	var obis string
	var tmp []byte
	rows = strings.Split(string(data), "\n")
	for _, it := range rows {
		if !strings.HasPrefix(it, "#") && it != "" {
			items := strings.Split(it, ";")
			str, _ := strconv.Atoi(items[0])
			ot := enums.ObjectType(str)
			tmp, err = LogicalNameToBytes(items[1])
			obis, err = ToLogicalName(tmp)
			desc := items[3]
			code := GXStandardObisCode{OBIS: []string{obis}, Description: desc,
				Interfaces: ot.String()}
			if len(items) > 4 {
				code.UIDataType = items[4]
			}
			g.codes = append(g.codes, code)
		}
	}
	return nil
}

func ChangeValueType(value any, dt enums.DataType, language *language.Tag) (any, error) {
	var ret any
	var err error
	var ok bool
	switch dt {
	case enums.DataTypeOctetString:
		if _, ok = value.([]byte); ok {
			ret = value
		} else if value == "" {
			ret = ""
		} else {
			ret = types.HexToBytes(value.(string))
		}
	case enums.DataTypeDateTime:
		if _, ok := value.(types.GXDateTime); ok {
			ret = value
		} else {
			ret, err = types.NewGXDateTimeFromString(value.(string), language)
			if err != nil {
				return nil, err
			}
		}
	case enums.DataTypeDate:
		if _, ok := value.(types.GXDateTime); ok {
			ret = value
		} else {
			ret, err = types.NewGXDateFromString(value.(string), language)
			if err != nil {
				return nil, err
			}
		}
	case enums.DataTypeTime:
		if _, ok := value.(types.GXDateTime); ok {
			ret = value
		} else {
			ret, err = types.NewGXTimeFromString(value.(string), language)
			if err != nil {
				return nil, err
			}
		}
	case enums.DataTypeEnum:
		if _, ok := value.(types.GXEnum); ok {
			ret = value
		} else {
			ret = types.GXEnum{Value: value.(byte)}
		}
	case enums.DataTypeStructure, enums.DataTypeArray:
		//TODO: ret = GXDLMSTranslator.XmlToValue(string(value))
	case enums.DataTypeBitString:
		ret, err = types.NewGXBitStringFromString(value.(string))
		if err != nil {
			return nil, err
		}
	default:
		dst, err := internal.GetDataType(dt)
		if err != nil {
			return nil, err
		}
		v := reflect.ValueOf(value)
		src := v.Type()
		if src.AssignableTo(dst) {
			return v.Interface(), nil
		}
		if src.ConvertibleTo(dst) {
			return v.Convert(dst).Interface(), nil
		}
		switch s := value.(type) {
		case string:
			return coerceFromString(s, dst)
		case []byte:
			return coerceFromString(string(s), dst)
		}

		return nil, fmt.Errorf("cannot coerce %v to %v", src, dst)
	}
	return ret, nil
}

func coerceFromString(s string, dst reflect.Type) (any, error) {
	switch dst.Kind() {
	case reflect.String:
		return s, nil

	case reflect.Bool:
		b, err := strconv.ParseBool(s)
		if err != nil {
			return nil, err
		}
		return b, nil

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		bits := dst.Bits()
		n, err := strconv.ParseInt(s, 10, bits)
		if err != nil {
			return nil, err
		}
		out := reflect.New(dst).Elem()
		out.SetInt(n)
		return out.Interface(), nil

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		bits := dst.Bits()
		n, err := strconv.ParseUint(s, 10, bits)
		if err != nil {
			return nil, err
		}
		out := reflect.New(dst).Elem()
		out.SetUint(n)
		return out.Interface(), nil

	case reflect.Float32, reflect.Float64:
		bits := dst.Bits()
		f, err := strconv.ParseFloat(s, bits)
		if err != nil {
			return nil, err
		}
		out := reflect.New(dst).Elem()
		out.SetFloat(f)
		return out.Interface(), nil
	}

	return nil, fmt.Errorf("cannot parse %q into %v", s, dst)
}

// SystemTitleToString returns the convert system title to string.standard: Used standard.
// st: System title.
// addComments: Are comments added.
//
//	Returns:
//	    System title in string format.
func (g *GXDLMSConverter) SystemTitleToString(standard enums.Standard, st []byte, addComments bool) string {
	return internal.SystemTitleToString(standard, st, addComments)
}

// KeyUsageToCertificateType returns the convert key usage to certificate type.
//
//	Returns:
//	    Certificate type.
func (g *GXDLMSConverter) KeyUsageToCertificateType(value enums.KeyUsage) (enums.CertificateType, error) {
	switch value {
	case enums.KeyUsageDigitalSignature:
		return enums.CertificateTypeDigitalSignature, nil
	case enums.KeyUsageKeyAgreement:
		return enums.CertificateTypeKeyAgreement, nil
	case enums.KeyUsageDigitalSignature | enums.KeyUsageKeyAgreement:
		return enums.CertificateTypeTLS, nil
	default:
		return 0, fmt.Errorf("invalid key usage: %v", value)
	}
}

// CertificateTypeToKeyUsage returns the convert key usage to certificate type.value: Key usage
//
//	Returns:
//	    Certificate type.
func (g *GXDLMSConverter) CertificateTypeToKeyUsage(value enums.CertificateType) (enums.KeyUsage, error) {
	switch value {
	case enums.CertificateTypeDigitalSignature:
		return enums.KeyUsageDigitalSignature, nil
	case enums.CertificateTypeKeyAgreement:
		return enums.KeyUsageKeyAgreement, nil
	case enums.CertificateTypeTLS:
		return enums.KeyUsageDigitalSignature | enums.KeyUsageKeyAgreement, nil
	default:
		return 0, fmt.Errorf("invalid certificate type: %v", value)
	}
}
