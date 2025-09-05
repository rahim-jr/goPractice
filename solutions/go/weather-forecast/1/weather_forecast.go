// Package weather provides current weather information for cities in Goblinocus.
// It stores the current weather condition and location and returns a summary string.
package weather

// CurrentCondition is the current weather condition (e.g., sunny, rainy) for the location.
var CurrentCondition string

// CurrentLocation is the city or place for which the CurrentCondition applies.
var CurrentLocation string

// Forecast updates the current weather for a given city and condition,
// returning a summary string of the weather condition at that location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
