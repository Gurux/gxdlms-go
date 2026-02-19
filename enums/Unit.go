package enums

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

	"github.com/Gurux/gxcommon-go"
)

// Unit describes available COSEM unit types.
type Unit int

const (
	// UnitNone defines that the // No unit is used.
	UnitNone Unit = iota
	// UnitYear defines that the // Unit is year.
	UnitYear
	// UnitMonth defines that the // Unit is month.
	UnitMonth
	// UnitWeek defines that the // Unit is week.
	UnitWeek
	// UnitDay defines that the unit is day.
	UnitDay
	// UnitHour defines that the unit is hour.
	UnitHour
	// UnitMinute defines that the unit is minute.
	UnitMinute
	// UnitSecond defines that the unit is second.
	UnitSecond
	// UnitPhaseAngleDegree defines that the unit is phase angle degree rad*180/p
	UnitPhaseAngleDegree
	// UnitTemperature defines that the unit is temperature T degree centigrade
	UnitTemperature
	// UnitLocalCurrency defines that the local currency is used as unit.
	UnitLocalCurrency
	// UnitLength defines that the length l meter m is used as an unit.
	UnitLength
	// UnitSpeed defines that the unit is Speed v m/s.
	UnitSpeed
	// UnitVolumeCubicMeter defines that the unit is Volume V m3.
	UnitVolumeCubicMeter
	// UnitCorrectedVolume defines that the unit is Corrected volume m3.
	UnitCorrectedVolume
	// UnitVolumeFluxHour defines that the unit is Volume flux m3/60*60s.
	UnitVolumeFluxHour
	// UnitCorrectedVolumeFluxHour defines that the unit is Corrected volume flux m3/60*60s.
	UnitCorrectedVolumeFluxHour
	// UnitVolumeFluxDay defines that the unit is Volume flux m3/24*60*60s.
	UnitVolumeFluxDay
	// UnitCorrectedVolumeFluxDay defines that the unit is Corrected volume flux m3/24*60*60s.
	UnitCorrectedVolumeFluxDay
	// UnitVolumeLiter defines that the unit is Volume 10-3 m3.
	UnitVolumeLiter
	// UnitMassKg defines that the unit is Mass m kilogram kg.
	UnitMassKg
	// UnitForce defines that the unit is Force F newton N.
	UnitForce
	// UnitEnergy defines that the unit is Energy newtonmeter J = Nm = Ws.
	UnitEnergy
	// UnitPressurePascal defines that the unit is Pressure p pascal N/m2.
	UnitPressurePascal
	// UnitPressureBar defines that the unit is Pressure p bar 10-5 N/m2.
	UnitPressureBar
	// UnitEnergyJoule defines that the unit is Energy joule J = Nm = Ws.
	UnitEnergyJoule
	// UnitThermalPower defines that the unit is Thermal power J/60*60s.
	UnitThermalPower
	// UnitActivePower defines that the unit is Active power P watt W = J/s.
	UnitActivePower
	// UnitApparentPower defines that the unit is Apparent power S.
	UnitApparentPower
	// UnitReactivePower defines that the unit is Reactive power Q.
	UnitReactivePower
	// UnitActiveEnergy defines that the unit is Active energy W*60*60s.
	UnitActiveEnergy
	// UnitApparentEnergy defines that the unit is Apparent energy VA*60*60s.
	UnitApparentEnergy
	// UnitReactiveEnergy defines that the unit is Reactive energy var*60*60s.
	UnitReactiveEnergy
	// UnitCurrent defines that the unit is Current I ampere A.
	UnitCurrent
	// UnitElectricalCharge defines that the unit is Electrical charge Q coulomb C = As.
	UnitElectricalCharge
	// UnitVoltage defines that the unit is Voltage.
	UnitVoltage
	// UnitElectricalFieldStrength defines that the unit is Electrical field strength E V/m.
	UnitElectricalFieldStrength
	// UnitCapacity defines that the unit is Capacity C farad C/V = As/V.
	UnitCapacity
	// UnitResistance defines that the unit is Resistance R ohm = V/A.
	UnitResistance
	// UnitResistivity defines that the unit is Resistivity.
	UnitResistivity
	// UnitMagneticFlux defines that the unit is Magnetic flux F weber Wb = Vs.
	UnitMagneticFlux
	// UnitInduction defines that the unit is Induction T tesla Wb/m2.
	UnitInduction
	// UnitMagnetic defines that the unit is Magnetic field strength H A/m.
	UnitMagnetic
	// UnitInductivity defines that the unit is Inductivity L henry H = Wb/A.
	UnitInductivity
	// UnitFrequency defines that the unit is Frequency f.
	UnitFrequency
	// UnitActive defines that the unit is Active energy meter constant 1/Wh.
	UnitActive
	// UnitReactive defines that the unit is Reactive energy meter constant.
	UnitReactive
	// UnitApparent defines that the unit is Apparent energy meter constant.
	UnitApparent
	// UnitV260 defines that the unit is V260*60s.
	UnitV260
	// UnitA260 defines that the unit is A260*60s.
	UnitA260
	// UnitMassKgPerSecond defines that the unit is Mass flux kg/s.
	UnitMassKgPerSecond
	// UnitConductance defines that the unit is Conductance siemens 1/ohm.
	UnitConductance
	// UnitKelvin defines that the temperature in Kelvin.
	UnitKelvin
	// UnitRU2h defines that the 1/(V2h) RU2h , volt-squared hour meter constant or pulse value.
	UnitRU2h
	// UnitRI2h defines that the 1/(A2h) RI2h , ampere-squared hour meter constant or pulse value.
	UnitRI2h
	// UnitCubicMeterRV defines that the 1/m3 RV , meter constant or pulse value (volume).
	UnitCubicMeterRV
	// UnitPercentage defines that the unit is percentage
	UnitPercentage
	// UnitAmpereHour defines that the unit is ampere hour.
	UnitAmpereHour
	// UnitEnergyPerVolume defines that the wh/m3 energy per volume 3,6*103 J/m3.
	UnitEnergyPerVolume Unit = 60
	// UnitWobbe defines that the j/m3 calorific value, wobbe.
	UnitWobbe Unit = 61
	// UnitMolePercent defines that the mol % molar fraction of gas composition mole percent (Basic gas composition unit)
	UnitMolePercent Unit = 62
	// UnitMassDensity defines that the g/m3 mass density, quantity of material.
	UnitMassDensity Unit = 63
	// UnitPascalSecond defines that the the dynamic viscosity pascal second (Characteristic of gas stream).
	UnitPascalSecond Unit = 64
	// UnitJouleKilogram defines that the j/kg Specific energy
	//  NOTE The amount of energy per unit of mass of a
	//  substance Joule / kilogram m2 . kg . s -2 / kg = m2 . s �2
	UnitJouleKilogram Unit = 65
	// UnitPressureGramPerSquareCentimeter defines that the pressure, gram per square centimeter.
	UnitPressureGramPerSquareCentimeter Unit = 66
	// UnitPressureAtmosphere defines that the pressure, atmosphere.
	UnitPressureAtmosphere Unit = 67
	// UnitSignalStrengthMilliWatt defines that the signal strength, dB milliwatt (e.g. of GSM radio systems).
	UnitSignalStrengthMilliWatt Unit = 70
	// UnitSignalStrengthMicroVolt defines that the signal strength, dB microvolt.
	UnitSignalStrengthMicroVolt Unit = 71
	// UnitdB defines that the logarithmic unit that expresses the ratio between two values of a physical quantity
	UnitdB Unit = 72
	// UnitInch defines that the length in inches.
	UnitInch Unit = 128
	// UnitFoot defines that the foot (Length).
	UnitFoot Unit = 129
	// UnitPound defines that the pound (mass).
	UnitPound Unit = 130
	// UnitFahrenheit defines that the fahrenheit
	UnitFahrenheit Unit = 131
	// UnitRankine defines that the rankine
	UnitRankine Unit = 132
	// UnitSquareInch defines that the square inch.
	UnitSquareInch Unit = 133
	// UnitSquareFoot defines that the square foot.
	UnitSquareFoot Unit = 134
	// UnitAcre defines that the acre
	UnitAcre Unit = 135
	// UnitCubicInch defines that the cubic inch.
	UnitCubicInch Unit = 136
	// UnitCubicFoot defines that the cubic foot.
	UnitCubicFoot Unit = 137
	// UnitAcreFoot defines that the acre-foot.
	UnitAcreFoot Unit = 138
	// UnitGallonImperial defines that the gallon (imperial).
	UnitGallonImperial Unit = 139
	// UnitGallonUS defines that the gallon (US).
	UnitGallonUS Unit = 140
	// UnitPoundForce defines that the pound force.
	UnitPoundForce Unit = 141
	// UnitPoundForcePerSquareInch defines that the pound force per square inch
	UnitPoundForcePerSquareInch Unit = 142
	// UnitPoundPerCubicFoot defines that the pound per cubic foot.
	UnitPoundPerCubicFoot Unit = 143
	// UnitPoundPerFootSecond defines that the pound per (foot second)
	UnitPoundPerFootSecond Unit = 144
	// UnitSquareFootPerSecond defines that the square foot per second.
	UnitSquareFootPerSecond Unit = 145
	// UnitBritishThermalUnit defines that the british thermal unit.
	UnitBritishThermalUnit Unit = 146
	// UnitThermEU defines that the therm EU.
	UnitThermEU Unit = 147
	// UnitThermUS defines that the therm US.
	UnitThermUS Unit = 148
	// UnitBritishThermalUnitPerPound defines that the british thermal unit per pound.
	UnitBritishThermalUnitPerPound Unit = 149
	// UnitBritishThermalUnitPerCubicFoot defines that the british thermal unit per cubic foot.
	UnitBritishThermalUnitPerCubicFoot Unit = 150
	// UnitCubicFeet defines that the cubic feet.
	UnitCubicFeet Unit = 151
	// UnitFootPerSecond defines that the foot per second.
	UnitFootPerSecond Unit = 152
	// UnitCubicFootPerSecond defines that the cubic foot per second.
	UnitCubicFootPerSecond Unit = 153
	// UnitCubicFootPerMin defines that the cubic foot per min.
	UnitCubicFootPerMin Unit = 154
	// UnitCubicFootPerHour defines that the cubic foot per hour.
	UnitCubicFootPerHour Unit = 155
	// UnitCubicFootPerDay defines that the cubic foot per day
	UnitCubicFootPerDay Unit = 156
	// UnitAcreFootPerSecond defines that the acre foot per second.
	UnitAcreFootPerSecond Unit = 157
	// UnitAcreFootPerMin defines that the acre foot per min.
	UnitAcreFootPerMin Unit = 158
	// UnitAcreFootPerHour defines that the acre foot per hour.
	UnitAcreFootPerHour Unit = 159
	// UnitAcreFootPerDay defines that the acre foot per day.
	UnitAcreFootPerDay Unit = 160
	// UnitImperialGallon defines that the imperial gallon.
	UnitImperialGallon Unit = 161
	// UnitImperialGallonPerSecond defines that the imperial gallon per second.
	UnitImperialGallonPerSecond Unit = 162
	// UnitImperialGallonPerMin defines that the imperial gallon per min.
	UnitImperialGallonPerMin Unit = 163
	// UnitImperialGallonPerHour defines that the imperial gallon per hour.
	UnitImperialGallonPerHour Unit = 164
	// UnitImperialGallonPerDay defines that the imperial gallon per day.
	UnitImperialGallonPerDay Unit = 165
	// UnitUSGallon defines that the uS gallon.
	UnitUSGallon Unit = 166
	// UnitUSGallonPerSecond defines that the uS gallon per second.
	UnitUSGallonPerSecond Unit = 167
	// UnitUSGallonPerMin defines that the uS gallon per min.
	UnitUSGallonPerMin Unit = 168
	// UnitUSGallonPerHour defines that the uS gallon per hour.
	UnitUSGallonPerHour Unit = 169
	// UnitUSGallonPerDay defines that the uS gallon per day.
	UnitUSGallonPerDay Unit = 170
	// UnitBritishThermalUnitPerSecond defines that the british thermal unit per second.
	UnitBritishThermalUnitPerSecond Unit = 171
	// UnitBritishThermalUnitPerMinute defines that the british thermal unit per minute.
	UnitBritishThermalUnitPerMinute Unit = 172
	// UnitBritishThermalUnitPerHour defines that the british thermal unit per hour.
	UnitBritishThermalUnitPerHour Unit = 173
	// UnitBritishThermalUnitPerDay defines that the british thermal unit per day.
	UnitBritishThermalUnitPerDay Unit = 174
	// UnitOtherUnit defines that the other Unit is used.
	UnitOtherUnit Unit = 254
	// UnitNoUnit defines that the no Unit is used.
	UnitNoUnit Unit = 255
)

// UnitParse converts the given string into a Unit value.
//
// It returns the corresponding Unit constant if the string matches
// a known level name, or an error if the input is invalid.
func UnitParse(value string) (Unit, error) {
	var ret Unit
	var err error
	switch {
	case strings.EqualFold(value, "None"):
		ret = UnitNone
	case strings.EqualFold(value, "Year"):
		ret = UnitYear
	case strings.EqualFold(value, "Month"):
		ret = UnitMonth
	case strings.EqualFold(value, "Week"):
		ret = UnitWeek
	case strings.EqualFold(value, "Day"):
		ret = UnitDay
	case strings.EqualFold(value, "Hour"):
		ret = UnitHour
	case strings.EqualFold(value, "Minute"):
		ret = UnitMinute
	case strings.EqualFold(value, "Second"):
		ret = UnitSecond
	case strings.EqualFold(value, "PhaseAngleDegree"):
		ret = UnitPhaseAngleDegree
	case strings.EqualFold(value, "Temperature"):
		ret = UnitTemperature
	case strings.EqualFold(value, "LocalCurrency"):
		ret = UnitLocalCurrency
	case strings.EqualFold(value, "Length"):
		ret = UnitLength
	case strings.EqualFold(value, "Speed"):
		ret = UnitSpeed
	case strings.EqualFold(value, "VolumeCubicMeter"):
		ret = UnitVolumeCubicMeter
	case strings.EqualFold(value, "CorrectedVolume"):
		ret = UnitCorrectedVolume
	case strings.EqualFold(value, "VolumeFluxHour"):
		ret = UnitVolumeFluxHour
	case strings.EqualFold(value, "CorrectedVolumeFluxHour"):
		ret = UnitCorrectedVolumeFluxHour
	case strings.EqualFold(value, "VolumeFluxDay"):
		ret = UnitVolumeFluxDay
	case strings.EqualFold(value, "CorrectedVolumeFluxDay"):
		ret = UnitCorrectedVolumeFluxDay
	case strings.EqualFold(value, "VolumeLiter"):
		ret = UnitVolumeLiter
	case strings.EqualFold(value, "MassKg"):
		ret = UnitMassKg
	case strings.EqualFold(value, "Force"):
		ret = UnitForce
	case strings.EqualFold(value, "Energy"):
		ret = UnitEnergy
	case strings.EqualFold(value, "PressurePascal"):
		ret = UnitPressurePascal
	case strings.EqualFold(value, "PressureBar"):
		ret = UnitPressureBar
	case strings.EqualFold(value, "EnergyJoule"):
		ret = UnitEnergyJoule
	case strings.EqualFold(value, "ThermalPower"):
		ret = UnitThermalPower
	case strings.EqualFold(value, "ActivePower"):
		ret = UnitActivePower
	case strings.EqualFold(value, "ApparentPower"):
		ret = UnitApparentPower
	case strings.EqualFold(value, "ReactivePower"):
		ret = UnitReactivePower
	case strings.EqualFold(value, "ActiveEnergy"):
		ret = UnitActiveEnergy
	case strings.EqualFold(value, "ApparentEnergy"):
		ret = UnitApparentEnergy
	case strings.EqualFold(value, "ReactiveEnergy"):
		ret = UnitReactiveEnergy
	case strings.EqualFold(value, "Current"):
		ret = UnitCurrent
	case strings.EqualFold(value, "ElectricalCharge"):
		ret = UnitElectricalCharge
	case strings.EqualFold(value, "Voltage"):
		ret = UnitVoltage
	case strings.EqualFold(value, "ElectricalFieldStrength"):
		ret = UnitElectricalFieldStrength
	case strings.EqualFold(value, "Capacity"):
		ret = UnitCapacity
	case strings.EqualFold(value, "Resistance"):
		ret = UnitResistance
	case strings.EqualFold(value, "Resistivity"):
		ret = UnitResistivity
	case strings.EqualFold(value, "MagneticFlux"):
		ret = UnitMagneticFlux
	case strings.EqualFold(value, "Induction"):
		ret = UnitInduction
	case strings.EqualFold(value, "Magnetic"):
		ret = UnitMagnetic
	case strings.EqualFold(value, "Inductivity"):
		ret = UnitInductivity
	case strings.EqualFold(value, "Frequency"):
		ret = UnitFrequency
	case strings.EqualFold(value, "Active"):
		ret = UnitActive
	case strings.EqualFold(value, "Reactive"):
		ret = UnitReactive
	case strings.EqualFold(value, "Apparent"):
		ret = UnitApparent
	case strings.EqualFold(value, "V260"):
		ret = UnitV260
	case strings.EqualFold(value, "A260"):
		ret = UnitA260
	case strings.EqualFold(value, "MassKgPerSecond"):
		ret = UnitMassKgPerSecond
	case strings.EqualFold(value, "Conductance"):
		ret = UnitConductance
	case strings.EqualFold(value, "Kelvin"):
		ret = UnitKelvin
	case strings.EqualFold(value, "RU2h"):
		ret = UnitRU2h
	case strings.EqualFold(value, "RI2h"):
		ret = UnitRI2h
	case strings.EqualFold(value, "CubicMeterRV"):
		ret = UnitCubicMeterRV
	case strings.EqualFold(value, "Percentage"):
		ret = UnitPercentage
	case strings.EqualFold(value, "AmpereHour"):
		ret = UnitAmpereHour
	case strings.EqualFold(value, "EnergyPerVolume"):
		ret = UnitEnergyPerVolume
	case strings.EqualFold(value, "Wobbe"):
		ret = UnitWobbe
	case strings.EqualFold(value, "MolePercent"):
		ret = UnitMolePercent
	case strings.EqualFold(value, "MassDensity"):
		ret = UnitMassDensity
	case strings.EqualFold(value, "PascalSecond"):
		ret = UnitPascalSecond
	case strings.EqualFold(value, "JouleKilogram"):
		ret = UnitJouleKilogram
	case strings.EqualFold(value, "PressureGramPerSquareCentimeter"):
		ret = UnitPressureGramPerSquareCentimeter
	case strings.EqualFold(value, "PressureAtmosphere"):
		ret = UnitPressureAtmosphere
	case strings.EqualFold(value, "SignalStrengthMilliWatt"):
		ret = UnitSignalStrengthMilliWatt
	case strings.EqualFold(value, "SignalStrengthMicroVolt"):
		ret = UnitSignalStrengthMicroVolt
	case strings.EqualFold(value, "dB"):
		ret = UnitdB
	case strings.EqualFold(value, "Inch"):
		ret = UnitInch
	case strings.EqualFold(value, "Foot"):
		ret = UnitFoot
	case strings.EqualFold(value, "Pound"):
		ret = UnitPound
	case strings.EqualFold(value, "Fahrenheit"):
		ret = UnitFahrenheit
	case strings.EqualFold(value, "Rankine"):
		ret = UnitRankine
	case strings.EqualFold(value, "SquareInch"):
		ret = UnitSquareInch
	case strings.EqualFold(value, "SquareFoot"):
		ret = UnitSquareFoot
	case strings.EqualFold(value, "Acre"):
		ret = UnitAcre
	case strings.EqualFold(value, "CubicInch"):
		ret = UnitCubicInch
	case strings.EqualFold(value, "CubicFoot"):
		ret = UnitCubicFoot
	case strings.EqualFold(value, "AcreFoot"):
		ret = UnitAcreFoot
	case strings.EqualFold(value, "GallonImperial"):
		ret = UnitGallonImperial
	case strings.EqualFold(value, "GallonUS"):
		ret = UnitGallonUS
	case strings.EqualFold(value, "PoundForce"):
		ret = UnitPoundForce
	case strings.EqualFold(value, "PoundForcePerSquareInch"):
		ret = UnitPoundForcePerSquareInch
	case strings.EqualFold(value, "PoundPerCubicFoot"):
		ret = UnitPoundPerCubicFoot
	case strings.EqualFold(value, "PoundPerFootSecond"):
		ret = UnitPoundPerFootSecond
	case strings.EqualFold(value, "SquareFootPerSecond"):
		ret = UnitSquareFootPerSecond
	case strings.EqualFold(value, "BritishThermalUnit"):
		ret = UnitBritishThermalUnit
	case strings.EqualFold(value, "ThermEU"):
		ret = UnitThermEU
	case strings.EqualFold(value, "ThermUS"):
		ret = UnitThermUS
	case strings.EqualFold(value, "BritishThermalUnitPerPound"):
		ret = UnitBritishThermalUnitPerPound
	case strings.EqualFold(value, "BritishThermalUnitPerCubicFoot"):
		ret = UnitBritishThermalUnitPerCubicFoot
	case strings.EqualFold(value, "CubicFeet"):
		ret = UnitCubicFeet
	case strings.EqualFold(value, "FootPerSecond"):
		ret = UnitFootPerSecond
	case strings.EqualFold(value, "CubicFootPerSecond"):
		ret = UnitCubicFootPerSecond
	case strings.EqualFold(value, "CubicFootPerMin"):
		ret = UnitCubicFootPerMin
	case strings.EqualFold(value, "CubicFootPerHour"):
		ret = UnitCubicFootPerHour
	case strings.EqualFold(value, "CubicFootPerDay"):
		ret = UnitCubicFootPerDay
	case strings.EqualFold(value, "AcreFootPerSecond"):
		ret = UnitAcreFootPerSecond
	case strings.EqualFold(value, "AcreFootPerMin"):
		ret = UnitAcreFootPerMin
	case strings.EqualFold(value, "AcreFootPerHour"):
		ret = UnitAcreFootPerHour
	case strings.EqualFold(value, "AcreFootPerDay"):
		ret = UnitAcreFootPerDay
	case strings.EqualFold(value, "ImperialGallon"):
		ret = UnitImperialGallon
	case strings.EqualFold(value, "ImperialGallonPerSecond"):
		ret = UnitImperialGallonPerSecond
	case strings.EqualFold(value, "ImperialGallonPerMin"):
		ret = UnitImperialGallonPerMin
	case strings.EqualFold(value, "ImperialGallonPerHour"):
		ret = UnitImperialGallonPerHour
	case strings.EqualFold(value, "ImperialGallonPerDay"):
		ret = UnitImperialGallonPerDay
	case strings.EqualFold(value, "USGallon"):
		ret = UnitUSGallon
	case strings.EqualFold(value, "USGallonPerSecond"):
		ret = UnitUSGallonPerSecond
	case strings.EqualFold(value, "USGallonPerMin"):
		ret = UnitUSGallonPerMin
	case strings.EqualFold(value, "USGallonPerHour"):
		ret = UnitUSGallonPerHour
	case strings.EqualFold(value, "USGallonPerDay"):
		ret = UnitUSGallonPerDay
	case strings.EqualFold(value, "BritishThermalUnitPerSecond"):
		ret = UnitBritishThermalUnitPerSecond
	case strings.EqualFold(value, "BritishThermalUnitPerMinute"):
		ret = UnitBritishThermalUnitPerMinute
	case strings.EqualFold(value, "BritishThermalUnitPerHour"):
		ret = UnitBritishThermalUnitPerHour
	case strings.EqualFold(value, "BritishThermalUnitPerDay"):
		ret = UnitBritishThermalUnitPerDay
	case strings.EqualFold(value, "OtherUnit"):
		ret = UnitOtherUnit
	case strings.EqualFold(value, "NoUnit"):
		ret = UnitNoUnit
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the Unit.
// It satisfies fmt.Stringer.
func (g Unit) String() string {
	var ret string
	switch g {
	case UnitNone:
		ret = "None"
	case UnitYear:
		ret = "Year"
	case UnitMonth:
		ret = "Month"
	case UnitWeek:
		ret = "Week"
	case UnitDay:
		ret = "Day"
	case UnitHour:
		ret = "Hour"
	case UnitMinute:
		ret = "Minute"
	case UnitSecond:
		ret = "Second"
	case UnitPhaseAngleDegree:
		ret = "PhaseAngleDegree"
	case UnitTemperature:
		ret = "Temperature"
	case UnitLocalCurrency:
		ret = "LocalCurrency"
	case UnitLength:
		ret = "Length"
	case UnitSpeed:
		ret = "Speed"
	case UnitVolumeCubicMeter:
		ret = "VolumeCubicMeter"
	case UnitCorrectedVolume:
		ret = "CorrectedVolume"
	case UnitVolumeFluxHour:
		ret = "VolumeFluxHour"
	case UnitCorrectedVolumeFluxHour:
		ret = "CorrectedVolumeFluxHour"
	case UnitVolumeFluxDay:
		ret = "VolumeFluxDay"
	case UnitCorrectedVolumeFluxDay:
		ret = "CorrectedVolumeFluxDay"
	case UnitVolumeLiter:
		ret = "VolumeLiter"
	case UnitMassKg:
		ret = "MassKg"
	case UnitForce:
		ret = "Force"
	case UnitEnergy:
		ret = "Energy"
	case UnitPressurePascal:
		ret = "PressurePascal"
	case UnitPressureBar:
		ret = "PressureBar"
	case UnitEnergyJoule:
		ret = "EnergyJoule"
	case UnitThermalPower:
		ret = "ThermalPower"
	case UnitActivePower:
		ret = "ActivePower"
	case UnitApparentPower:
		ret = "ApparentPower"
	case UnitReactivePower:
		ret = "ReactivePower"
	case UnitActiveEnergy:
		ret = "ActiveEnergy"
	case UnitApparentEnergy:
		ret = "ApparentEnergy"
	case UnitReactiveEnergy:
		ret = "ReactiveEnergy"
	case UnitCurrent:
		ret = "Current"
	case UnitElectricalCharge:
		ret = "ElectricalCharge"
	case UnitVoltage:
		ret = "Voltage"
	case UnitElectricalFieldStrength:
		ret = "ElectricalFieldStrength"
	case UnitCapacity:
		ret = "Capacity"
	case UnitResistance:
		ret = "Resistance"
	case UnitResistivity:
		ret = "Resistivity"
	case UnitMagneticFlux:
		ret = "MagneticFlux"
	case UnitInduction:
		ret = "Induction"
	case UnitMagnetic:
		ret = "Magnetic"
	case UnitInductivity:
		ret = "Inductivity"
	case UnitFrequency:
		ret = "Frequency"
	case UnitActive:
		ret = "Active"
	case UnitReactive:
		ret = "Reactive"
	case UnitApparent:
		ret = "Apparent"
	case UnitV260:
		ret = "V260"
	case UnitA260:
		ret = "A260"
	case UnitMassKgPerSecond:
		ret = "MassKgPerSecond"
	case UnitConductance:
		ret = "Conductance"
	case UnitKelvin:
		ret = "Kelvin"
	case UnitRU2h:
		ret = "RU2h"
	case UnitRI2h:
		ret = "RI2h"
	case UnitCubicMeterRV:
		ret = "Cubic meter RV"
	case UnitPercentage:
		ret = "Percentage"
	case UnitAmpereHour:
		ret = "AmpereHour"
	case UnitEnergyPerVolume:
		ret = "EnergyPerVolume"
	case UnitWobbe:
		ret = "Wobbe"
	case UnitMolePercent:
		ret = "MolePercent"
	case UnitMassDensity:
		ret = "MassDensity"
	case UnitPascalSecond:
		ret = "PascalSecond"
	case UnitJouleKilogram:
		ret = "JouleKilogram"
	case UnitPressureGramPerSquareCentimeter:
		ret = "PressureGramPerSquareCentimeter"
	case UnitPressureAtmosphere:
		ret = "PressureAtmosphere"
	case UnitSignalStrengthMilliWatt:
		ret = "SignalStrengthMilliWatt"
	case UnitSignalStrengthMicroVolt:
		ret = "SignalStrengthMicroVolt"
	case UnitdB:
		ret = "dB"
	case UnitInch:
		ret = "Inch"
	case UnitFoot:
		ret = "Foot"
	case UnitPound:
		ret = "Pound"
	case UnitFahrenheit:
		ret = "Fahrenheit"
	case UnitRankine:
		ret = "Rankine"
	case UnitSquareInch:
		ret = "SquareInch"
	case UnitSquareFoot:
		ret = "SquareFoot"
	case UnitAcre:
		ret = "Acre"
	case UnitCubicInch:
		ret = "CubicInch"
	case UnitCubicFoot:
		ret = "CubicFoot"
	case UnitAcreFoot:
		ret = "AcreFoot"
	case UnitGallonImperial:
		ret = "GallonImperial"
	case UnitGallonUS:
		ret = "Gallon US"
	case UnitPoundForce:
		ret = "PoundForce"
	case UnitPoundForcePerSquareInch:
		ret = "PoundForcePerSquareInch"
	case UnitPoundPerCubicFoot:
		ret = "PoundPerCubicFoot"
	case UnitPoundPerFootSecond:
		ret = "PoundPerFootSecond"
	case UnitSquareFootPerSecond:
		ret = "SquareFootPerSecond"
	case UnitBritishThermalUnit:
		ret = "BritishThermalUnit"
	case UnitThermEU:
		ret = "Therm EU"
	case UnitThermUS:
		ret = "Therm US"
	case UnitBritishThermalUnitPerPound:
		ret = "BritishThermalUnitPerPound"
	case UnitBritishThermalUnitPerCubicFoot:
		ret = "BritishThermalUnitPerCubicFoot"
	case UnitCubicFeet:
		ret = "CubicFeet"
	case UnitFootPerSecond:
		ret = "FootPerSecond"
	case UnitCubicFootPerSecond:
		ret = "CubicFootPerSecond"
	case UnitCubicFootPerMin:
		ret = "CubicFootPerMin"
	case UnitCubicFootPerHour:
		ret = "CubicFootPerHour"
	case UnitCubicFootPerDay:
		ret = "CubicFootPerDay"
	case UnitAcreFootPerSecond:
		ret = "AcreFootPerSecond"
	case UnitAcreFootPerMin:
		ret = "AcreFootPerMin"
	case UnitAcreFootPerHour:
		ret = "AcreFootPerHour"
	case UnitAcreFootPerDay:
		ret = "AcreFootPerDay"
	case UnitImperialGallon:
		ret = "ImperialGallon"
	case UnitImperialGallonPerSecond:
		ret = "ImperialGallonPerSecond"
	case UnitImperialGallonPerMin:
		ret = "ImperialGallonPerMin"
	case UnitImperialGallonPerHour:
		ret = "ImperialGallonPerHour"
	case UnitImperialGallonPerDay:
		ret = "ImperialGallonPerDay"
	case UnitUSGallon:
		ret = "US Gallon"
	case UnitUSGallonPerSecond:
		ret = "US gallon per second"
	case UnitUSGallonPerMin:
		ret = "US gallon per min"
	case UnitUSGallonPerHour:
		ret = "US gallon per hour"
	case UnitUSGallonPerDay:
		ret = "US gallon per day"
	case UnitBritishThermalUnitPerSecond:
		ret = "BritishThermalUnitPerSecond"
	case UnitBritishThermalUnitPerMinute:
		ret = "BritishThermalUnitPerMinute"
	case UnitBritishThermalUnitPerHour:
		ret = "BritishThermalUnitPerHour"
	case UnitBritishThermalUnitPerDay:
		ret = "BritishThermalUnitPerDay"
	case UnitOtherUnit:
		ret = "OtherUnit"
	case UnitNoUnit:
		ret = "NoUnit"
	}
	return ret
}

// AllUnit returns a slice containing all defined Unit values.
func AllUnit() []Unit {
	return []Unit{
		UnitNone,
		UnitYear,
		UnitMonth,
		UnitWeek,
		UnitDay,
		UnitHour,
		UnitMinute,
		UnitSecond,
		UnitPhaseAngleDegree,
		UnitTemperature,
		UnitLocalCurrency,
		UnitLength,
		UnitSpeed,
		UnitVolumeCubicMeter,
		UnitCorrectedVolume,
		UnitVolumeFluxHour,
		UnitCorrectedVolumeFluxHour,
		UnitVolumeFluxDay,
		UnitCorrectedVolumeFluxDay,
		UnitVolumeLiter,
		UnitMassKg,
		UnitForce,
		UnitEnergy,
		UnitPressurePascal,
		UnitPressureBar,
		UnitEnergyJoule,
		UnitThermalPower,
		UnitActivePower,
		UnitApparentPower,
		UnitReactivePower,
		UnitActiveEnergy,
		UnitApparentEnergy,
		UnitReactiveEnergy,
		UnitCurrent,
		UnitElectricalCharge,
		UnitVoltage,
		UnitElectricalFieldStrength,
		UnitCapacity,
		UnitResistance,
		UnitResistivity,
		UnitMagneticFlux,
		UnitInduction,
		UnitMagnetic,
		UnitInductivity,
		UnitFrequency,
		UnitActive,
		UnitReactive,
		UnitApparent,
		UnitV260,
		UnitA260,
		UnitMassKgPerSecond,
		UnitConductance,
		UnitKelvin,
		UnitRU2h,
		UnitRI2h,
		UnitCubicMeterRV,
		UnitPercentage,
		UnitAmpereHour,
		UnitEnergyPerVolume,
		UnitWobbe,
		UnitMolePercent,
		UnitMassDensity,
		UnitPascalSecond,
		UnitJouleKilogram,
		UnitPressureGramPerSquareCentimeter,
		UnitPressureAtmosphere,
		UnitSignalStrengthMilliWatt,
		UnitSignalStrengthMicroVolt,
		UnitdB,
		UnitInch,
		UnitFoot,
		UnitPound,
		UnitFahrenheit,
		UnitRankine,
		UnitSquareInch,
		UnitSquareFoot,
		UnitAcre,
		UnitCubicInch,
		UnitCubicFoot,
		UnitAcreFoot,
		UnitGallonImperial,
		UnitGallonUS,
		UnitPoundForce,
		UnitPoundForcePerSquareInch,
		UnitPoundPerCubicFoot,
		UnitPoundPerFootSecond,
		UnitSquareFootPerSecond,
		UnitBritishThermalUnit,
		UnitThermEU,
		UnitThermUS,
		UnitBritishThermalUnitPerPound,
		UnitBritishThermalUnitPerCubicFoot,
		UnitCubicFeet,
		UnitFootPerSecond,
		UnitCubicFootPerSecond,
		UnitCubicFootPerMin,
		UnitCubicFootPerHour,
		UnitCubicFootPerDay,
		UnitAcreFootPerSecond,
		UnitAcreFootPerMin,
		UnitAcreFootPerHour,
		UnitAcreFootPerDay,
		UnitImperialGallon,
		UnitImperialGallonPerSecond,
		UnitImperialGallonPerMin,
		UnitImperialGallonPerHour,
		UnitImperialGallonPerDay,
		UnitUSGallon,
		UnitUSGallonPerSecond,
		UnitUSGallonPerMin,
		UnitUSGallonPerHour,
		UnitUSGallonPerDay,
		UnitBritishThermalUnitPerSecond,
		UnitBritishThermalUnitPerMinute,
		UnitBritishThermalUnitPerHour,
		UnitBritishThermalUnitPerDay,
		UnitOtherUnit,
		UnitNoUnit,
	}
}
