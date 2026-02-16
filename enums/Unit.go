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
package enums

import (
	"fmt"
	"strings"

	"github.com/Gurux/gxcommon-go"
)

// Unit :  describes available COSEM unit types.
type Unit int

const (
	// UnitNone defines that the no unit is used.
	UnitNone Unit = iota
	// UnitYear defines that the unit is year.
	UnitYear
	// UnitMonth defines that the unit is month.
	UnitMonth
	// UnitWeek defines that the unit is week.
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
	switch strings.ToUpper(value) {
	case "NONE":
		ret = UnitNone
	case "YEAR":
		ret = UnitYear
	case "MONTH":
		ret = UnitMonth
	case "WEEK":
		ret = UnitWeek
	case "DAY":
		ret = UnitDay
	case "HOUR":
		ret = UnitHour
	case "MINUTE":
		ret = UnitMinute
	case "SECOND":
		ret = UnitSecond
	case "PHASEANGLEDEGREE":
		ret = UnitPhaseAngleDegree
	case "TEMPERATURE":
		ret = UnitTemperature
	case "LOCALCURRENCY":
		ret = UnitLocalCurrency
	case "LENGTH":
		ret = UnitLength
	case "SPEED":
		ret = UnitSpeed
	case "VOLUMECUBICMETER":
		ret = UnitVolumeCubicMeter
	case "CORRECTEDVOLUME":
		ret = UnitCorrectedVolume
	case "VOLUMEFLUXHOUR":
		ret = UnitVolumeFluxHour
	case "CORRECTEDVOLUMEFLUXHOUR":
		ret = UnitCorrectedVolumeFluxHour
	case "VOLUMEFLUXDAY":
		ret = UnitVolumeFluxDay
	case "CORRECTEDVOLUMEFLUXDAY":
		ret = UnitCorrectedVolumeFluxDay
	case "VOLUMELITER":
		ret = UnitVolumeLiter
	case "MASSKG":
		ret = UnitMassKg
	case "FORCE":
		ret = UnitForce
	case "ENERGY":
		ret = UnitEnergy
	case "PRESSUREPASCAL":
		ret = UnitPressurePascal
	case "PRESSUREBAR":
		ret = UnitPressureBar
	case "ENERGYJOULE":
		ret = UnitEnergyJoule
	case "THERMALPOWER":
		ret = UnitThermalPower
	case "ACTIVEPOWER":
		ret = UnitActivePower
	case "APPARENTPOWER":
		ret = UnitApparentPower
	case "REACTIVEPOWER":
		ret = UnitReactivePower
	case "ACTIVEENERGY":
		ret = UnitActiveEnergy
	case "APPARENTENERGY":
		ret = UnitApparentEnergy
	case "REACTIVEENERGY":
		ret = UnitReactiveEnergy
	case "CURRENT":
		ret = UnitCurrent
	case "ELECTRICALCHARGE":
		ret = UnitElectricalCharge
	case "VOLTAGE":
		ret = UnitVoltage
	case "ELECTRICALFIELDSTRENGTH":
		ret = UnitElectricalFieldStrength
	case "CAPACITY":
		ret = UnitCapacity
	case "RESISTANCE":
		ret = UnitResistance
	case "RESISTIVITY":
		ret = UnitResistivity
	case "MAGNETICFLUX":
		ret = UnitMagneticFlux
	case "INDUCTION":
		ret = UnitInduction
	case "MAGNETIC":
		ret = UnitMagnetic
	case "INDUCTIVITY":
		ret = UnitInductivity
	case "FREQUENCY":
		ret = UnitFrequency
	case "ACTIVE":
		ret = UnitActive
	case "REACTIVE":
		ret = UnitReactive
	case "APPARENT":
		ret = UnitApparent
	case "V260":
		ret = UnitV260
	case "A260":
		ret = UnitA260
	case "MASSKGPERSECOND":
		ret = UnitMassKgPerSecond
	case "CONDUCTANCE":
		ret = UnitConductance
	case "KELVIN":
		ret = UnitKelvin
	case "RU2H":
		ret = UnitRU2h
	case "RI2H":
		ret = UnitRI2h
	case "CUBICMETERRV":
		ret = UnitCubicMeterRV
	case "PERCENTAGE":
		ret = UnitPercentage
	case "AMPEREHOUR":
		ret = UnitAmpereHour
	case "ENERGYPERVOLUME":
		ret = UnitEnergyPerVolume
	case "WOBBE":
		ret = UnitWobbe
	case "MOLEPERCENT":
		ret = UnitMolePercent
	case "MASSDENSITY":
		ret = UnitMassDensity
	case "PASCALSECOND":
		ret = UnitPascalSecond
	case "JOULEKILOGRAM":
		ret = UnitJouleKilogram
	case "PRESSUREGRAMPERSQUARECENTIMETER":
		ret = UnitPressureGramPerSquareCentimeter
	case "PRESSUREATMOSPHERE":
		ret = UnitPressureAtmosphere
	case "SIGNALSTRENGTHMILLIWATT":
		ret = UnitSignalStrengthMilliWatt
	case "SIGNALSTRENGTHMICROVOLT":
		ret = UnitSignalStrengthMicroVolt
	case "DB":
		ret = UnitdB
	case "INCH":
		ret = UnitInch
	case "FOOT":
		ret = UnitFoot
	case "POUND":
		ret = UnitPound
	case "FAHRENHEIT":
		ret = UnitFahrenheit
	case "RANKINE":
		ret = UnitRankine
	case "SQUAREINCH":
		ret = UnitSquareInch
	case "SQUAREFOOT":
		ret = UnitSquareFoot
	case "ACRE":
		ret = UnitAcre
	case "CUBICINCH":
		ret = UnitCubicInch
	case "CUBICFOOT":
		ret = UnitCubicFoot
	case "ACREFOOT":
		ret = UnitAcreFoot
	case "GALLONIMPERIAL":
		ret = UnitGallonImperial
	case "GALLONUS":
		ret = UnitGallonUS
	case "POUNDFORCE":
		ret = UnitPoundForce
	case "POUNDFORCEPERSQUAREINCH":
		ret = UnitPoundForcePerSquareInch
	case "POUNDPERCUBICFOOT":
		ret = UnitPoundPerCubicFoot
	case "POUNDPERFOOTSECOND":
		ret = UnitPoundPerFootSecond
	case "SQUAREFOOTPERSECOND":
		ret = UnitSquareFootPerSecond
	case "BRITISHTHERMALUNIT":
		ret = UnitBritishThermalUnit
	case "THERMEU":
		ret = UnitThermEU
	case "THERMUS":
		ret = UnitThermUS
	case "BRITISHTHERMALUNITPERPOUND":
		ret = UnitBritishThermalUnitPerPound
	case "BRITISHTHERMALUNITPERCUBICFOOT":
		ret = UnitBritishThermalUnitPerCubicFoot
	case "CUBICFEET":
		ret = UnitCubicFeet
	case "FOOTPERSECOND":
		ret = UnitFootPerSecond
	case "CUBICFOOTPERSECOND":
		ret = UnitCubicFootPerSecond
	case "CUBICFOOTPERMIN":
		ret = UnitCubicFootPerMin
	case "CUBICFOOTPERHOUR":
		ret = UnitCubicFootPerHour
	case "CUBICFOOTPERDAY":
		ret = UnitCubicFootPerDay
	case "ACREFOOTPERSECOND":
		ret = UnitAcreFootPerSecond
	case "ACREFOOTPERMIN":
		ret = UnitAcreFootPerMin
	case "ACREFOOTPERHOUR":
		ret = UnitAcreFootPerHour
	case "ACREFOOTPERDAY":
		ret = UnitAcreFootPerDay
	case "IMPERIALGALLON":
		ret = UnitImperialGallon
	case "IMPERIALGALLONPERSECOND":
		ret = UnitImperialGallonPerSecond
	case "IMPERIALGALLONPERMIN":
		ret = UnitImperialGallonPerMin
	case "IMPERIALGALLONPERHOUR":
		ret = UnitImperialGallonPerHour
	case "IMPERIALGALLONPERDAY":
		ret = UnitImperialGallonPerDay
	case "USGALLON":
		ret = UnitUSGallon
	case "USGALLONPERSECOND":
		ret = UnitUSGallonPerSecond
	case "USGALLONPERMIN":
		ret = UnitUSGallonPerMin
	case "USGALLONPERHOUR":
		ret = UnitUSGallonPerHour
	case "USGALLONPERDAY":
		ret = UnitUSGallonPerDay
	case "BRITISHTHERMALUNITPERSECOND":
		ret = UnitBritishThermalUnitPerSecond
	case "BRITISHTHERMALUNITPERMINUTE":
		ret = UnitBritishThermalUnitPerMinute
	case "BRITISHTHERMALUNITPERHOUR":
		ret = UnitBritishThermalUnitPerHour
	case "BRITISHTHERMALUNITPERDAY":
		ret = UnitBritishThermalUnitPerDay
	case "OTHERUNIT":
		ret = UnitOtherUnit
	case "NOUNIT":
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
		ret = "NONE"
	case UnitYear:
		ret = "YEAR"
	case UnitMonth:
		ret = "MONTH"
	case UnitWeek:
		ret = "WEEK"
	case UnitDay:
		ret = "DAY"
	case UnitHour:
		ret = "HOUR"
	case UnitMinute:
		ret = "MINUTE"
	case UnitSecond:
		ret = "SECOND"
	case UnitPhaseAngleDegree:
		ret = "PHASEANGLEDEGREE"
	case UnitTemperature:
		ret = "TEMPERATURE"
	case UnitLocalCurrency:
		ret = "LOCALCURRENCY"
	case UnitLength:
		ret = "LENGTH"
	case UnitSpeed:
		ret = "SPEED"
	case UnitVolumeCubicMeter:
		ret = "VOLUMECUBICMETER"
	case UnitCorrectedVolume:
		ret = "CORRECTEDVOLUME"
	case UnitVolumeFluxHour:
		ret = "VOLUMEFLUXHOUR"
	case UnitCorrectedVolumeFluxHour:
		ret = "CORRECTEDVOLUMEFLUXHOUR"
	case UnitVolumeFluxDay:
		ret = "VOLUMEFLUXDAY"
	case UnitCorrectedVolumeFluxDay:
		ret = "CORRECTEDVOLUMEFLUXDAY"
	case UnitVolumeLiter:
		ret = "VOLUMELITER"
	case UnitMassKg:
		ret = "MASSKG"
	case UnitForce:
		ret = "FORCE"
	case UnitEnergy:
		ret = "ENERGY"
	case UnitPressurePascal:
		ret = "PRESSUREPASCAL"
	case UnitPressureBar:
		ret = "PRESSUREBAR"
	case UnitEnergyJoule:
		ret = "ENERGYJOULE"
	case UnitThermalPower:
		ret = "THERMALPOWER"
	case UnitActivePower:
		ret = "ACTIVEPOWER"
	case UnitApparentPower:
		ret = "APPARENTPOWER"
	case UnitReactivePower:
		ret = "REACTIVEPOWER"
	case UnitActiveEnergy:
		ret = "ACTIVEENERGY"
	case UnitApparentEnergy:
		ret = "APPARENTENERGY"
	case UnitReactiveEnergy:
		ret = "REACTIVEENERGY"
	case UnitCurrent:
		ret = "CURRENT"
	case UnitElectricalCharge:
		ret = "ELECTRICALCHARGE"
	case UnitVoltage:
		ret = "VOLTAGE"
	case UnitElectricalFieldStrength:
		ret = "ELECTRICALFIELDSTRENGTH"
	case UnitCapacity:
		ret = "CAPACITY"
	case UnitResistance:
		ret = "RESISTANCE"
	case UnitResistivity:
		ret = "RESISTIVITY"
	case UnitMagneticFlux:
		ret = "MAGNETICFLUX"
	case UnitInduction:
		ret = "INDUCTION"
	case UnitMagnetic:
		ret = "MAGNETIC"
	case UnitInductivity:
		ret = "INDUCTIVITY"
	case UnitFrequency:
		ret = "FREQUENCY"
	case UnitActive:
		ret = "ACTIVE"
	case UnitReactive:
		ret = "REACTIVE"
	case UnitApparent:
		ret = "APPARENT"
	case UnitV260:
		ret = "V260"
	case UnitA260:
		ret = "A260"
	case UnitMassKgPerSecond:
		ret = "MASSKGPERSECOND"
	case UnitConductance:
		ret = "CONDUCTANCE"
	case UnitKelvin:
		ret = "KELVIN"
	case UnitRU2h:
		ret = "RU2H"
	case UnitRI2h:
		ret = "RI2H"
	case UnitCubicMeterRV:
		ret = "CUBICMETERRV"
	case UnitPercentage:
		ret = "PERCENTAGE"
	case UnitAmpereHour:
		ret = "AMPEREHOUR"
	case UnitEnergyPerVolume:
		ret = "ENERGYPERVOLUME"
	case UnitWobbe:
		ret = "WOBBE"
	case UnitMolePercent:
		ret = "MOLEPERCENT"
	case UnitMassDensity:
		ret = "MASSDENSITY"
	case UnitPascalSecond:
		ret = "PASCALSECOND"
	case UnitJouleKilogram:
		ret = "JOULEKILOGRAM"
	case UnitPressureGramPerSquareCentimeter:
		ret = "PRESSUREGRAMPERSQUARECENTIMETER"
	case UnitPressureAtmosphere:
		ret = "PRESSUREATMOSPHERE"
	case UnitSignalStrengthMilliWatt:
		ret = "SIGNALSTRENGTHMILLIWATT"
	case UnitSignalStrengthMicroVolt:
		ret = "SIGNALSTRENGTHMICROVOLT"
	case UnitdB:
		ret = "DB"
	case UnitInch:
		ret = "INCH"
	case UnitFoot:
		ret = "FOOT"
	case UnitPound:
		ret = "POUND"
	case UnitFahrenheit:
		ret = "FAHRENHEIT"
	case UnitRankine:
		ret = "RANKINE"
	case UnitSquareInch:
		ret = "SQUAREINCH"
	case UnitSquareFoot:
		ret = "SQUAREFOOT"
	case UnitAcre:
		ret = "ACRE"
	case UnitCubicInch:
		ret = "CUBICINCH"
	case UnitCubicFoot:
		ret = "CUBICFOOT"
	case UnitAcreFoot:
		ret = "ACREFOOT"
	case UnitGallonImperial:
		ret = "GALLONIMPERIAL"
	case UnitGallonUS:
		ret = "GALLONUS"
	case UnitPoundForce:
		ret = "POUNDFORCE"
	case UnitPoundForcePerSquareInch:
		ret = "POUNDFORCEPERSQUAREINCH"
	case UnitPoundPerCubicFoot:
		ret = "POUNDPERCUBICFOOT"
	case UnitPoundPerFootSecond:
		ret = "POUNDPERFOOTSECOND"
	case UnitSquareFootPerSecond:
		ret = "SQUAREFOOTPERSECOND"
	case UnitBritishThermalUnit:
		ret = "BRITISHTHERMALUNIT"
	case UnitThermEU:
		ret = "THERMEU"
	case UnitThermUS:
		ret = "THERMUS"
	case UnitBritishThermalUnitPerPound:
		ret = "BRITISHTHERMALUNITPERPOUND"
	case UnitBritishThermalUnitPerCubicFoot:
		ret = "BRITISHTHERMALUNITPERCUBICFOOT"
	case UnitCubicFeet:
		ret = "CUBICFEET"
	case UnitFootPerSecond:
		ret = "FOOTPERSECOND"
	case UnitCubicFootPerSecond:
		ret = "CUBICFOOTPERSECOND"
	case UnitCubicFootPerMin:
		ret = "CUBICFOOTPERMIN"
	case UnitCubicFootPerHour:
		ret = "CUBICFOOTPERHOUR"
	case UnitCubicFootPerDay:
		ret = "CUBICFOOTPERDAY"
	case UnitAcreFootPerSecond:
		ret = "ACREFOOTPERSECOND"
	case UnitAcreFootPerMin:
		ret = "ACREFOOTPERMIN"
	case UnitAcreFootPerHour:
		ret = "ACREFOOTPERHOUR"
	case UnitAcreFootPerDay:
		ret = "ACREFOOTPERDAY"
	case UnitImperialGallon:
		ret = "IMPERIALGALLON"
	case UnitImperialGallonPerSecond:
		ret = "IMPERIALGALLONPERSECOND"
	case UnitImperialGallonPerMin:
		ret = "IMPERIALGALLONPERMIN"
	case UnitImperialGallonPerHour:
		ret = "IMPERIALGALLONPERHOUR"
	case UnitImperialGallonPerDay:
		ret = "IMPERIALGALLONPERDAY"
	case UnitUSGallon:
		ret = "USGALLON"
	case UnitUSGallonPerSecond:
		ret = "USGALLONPERSECOND"
	case UnitUSGallonPerMin:
		ret = "USGALLONPERMIN"
	case UnitUSGallonPerHour:
		ret = "USGALLONPERHOUR"
	case UnitUSGallonPerDay:
		ret = "USGALLONPERDAY"
	case UnitBritishThermalUnitPerSecond:
		ret = "BRITISHTHERMALUNITPERSECOND"
	case UnitBritishThermalUnitPerMinute:
		ret = "BRITISHTHERMALUNITPERMINUTE"
	case UnitBritishThermalUnitPerHour:
		ret = "BRITISHTHERMALUNITPERHOUR"
	case UnitBritishThermalUnitPerDay:
		ret = "BRITISHTHERMALUNITPERDAY"
	case UnitOtherUnit:
		ret = "OTHERUNIT"
	case UnitNoUnit:
		ret = "NOUNIT"
	}
	return ret
}
