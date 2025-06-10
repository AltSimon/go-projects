package unitconverter

import (
	"fmt"
	"log"
	"math"
)

var TemperatureUnits = map[string][]string{
	"Kelvin":     {"K"},
	"Fahrenheit": {"F"},
	"Celsius":    {"C"},
}

var DistanceUnitsToSI = map[string]float64{
	"Kilometre":  1000,
	"Hectometre": 100,
	"Decametre":  10,
	"Metre":      1,
	"Decimetre":  1e-1,
	"Centimetre": 1e-2,
	"Millimetre": 1e-3,
	"Micrometre": 1e-6,
	"Nanometre":  1e-9,
	"Picometre":  1e-12,
	"Yard":       1143.0 / 1250.0,
	"Inch":       127.0 / 5000.0,
}

var VolumeUnitsToSI = map[string]float64{
	"Litre":            1.0 / 1000.0,
	"Cubic Foot":       1000.0 / 35315.0,
	"Gallon":           10.0 / 2642.0,
	"Quart":            1.0 / 1057.0,
	"Barrel":           100.0 / 629.0,
	"Pint":             1.0 / 2113.0,
	"Cubic Metre":      1,
	"Cubic Centimetre": 1e-6,
}

func aliasToUnit(alias string) string {
	return ""
}

func round(value float64, precision int) float64 {
	return math.Round(value*math.Pow10(precision)) / math.Pow10(precision)
}

func convertNonTemperatureUnit(value float64, unitFrom string, unitTo string) float64 {
	return 0
}

func convertTemperature(value float64, unitIn string, unitOut string) float64 {
	if unitIn == unitOut {
		fmt.Println("[WARNING] convertTemperature: same unit in and out")
		return value
	}

	switch unitIn {
	case "Kelvin":
		if unitOut == "Celsius" {
			return value + 273.15
		}
		if unitOut == "Fahrenheit" {
			return value*9/5 - 459.67
		}
	case "Celsius":
		if unitOut == "Kelvin" {
			return value - 273.15
		}
		if unitOut == "Fahrenheit" {
			return value*9/5 + 32
		}
	case "Fahrenheit":
		if unitOut == "Celsius" {
			return (value - 32) * 5 / 9
		}
		if unitOut == "Kelvin" {
			return (value + 459.67) * 5 / 9
		}
	default:
		log.Fatal("[ERROR] convertTemperature: Invalid units")
		return 0
	}
	return 0

}

func isTemperatureUnit(unit string) bool {
	if unit == "Fahrenheit" {
		return true
	}
	if unit == "Kelvin" {
		return true
	}
	if unit == "Celsius" {
		return true
	}
	return false
}

func ConvertUnit(value float64, unitFrom string, unitTo string) float64 {
	var result float64
	if isTemperatureUnit(unitFrom) {
		result = convertTemperature(value, unitFrom, unitTo)
	} else {
		result = convertNonTemperatureUnit(value, unitFrom, unitTo)
	}
	return round(result, 3)
}
