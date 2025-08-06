// Package weather contains tools that help forcast weather for Goblinocus.
package weather

// CurrentCondition is a string variable that holds the current weather conditions.
var CurrentCondition string

// CurrentLocation is a string variable that holds the current location.
var CurrentLocation string

// Forecast is a function that takes city and condition as paramerters, and outputs the current weather conditions.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
