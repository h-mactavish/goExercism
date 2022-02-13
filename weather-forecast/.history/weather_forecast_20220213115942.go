//Package weather Describes a weather package utility.
package weather

//Respresents current condition of the weather.
var CurrentCondition string

//Represents current location of the weather.
var CurrentLocation string

//Forcasts returns a String describing the current weather at the current city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
