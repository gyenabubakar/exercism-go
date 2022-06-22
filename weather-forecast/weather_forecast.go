// Package weather names a group of files under a namespace, in this case, all source code relating to weather.
package weather

// CurrentCondition stores the current weather condition.
var CurrentCondition string
// CurrentLocation stores the current location for which the weather will be forecast.
var CurrentLocation string

// Forecast provides weather information based on the location and weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
