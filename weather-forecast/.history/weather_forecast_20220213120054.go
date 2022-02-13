// Package weather Describes a weather package utility.
package weather

// CurrentCondition represents current condition of the weather.
var CurrentCondition string

// CurrentLocation represents current location of the weather.
var CurrentLocation string

// Forcast returns a String describing the current weather at the current city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
