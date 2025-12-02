// Package weather contains a forecast and current conditions and location.
package weather

var (
	// CurrentCondition is the current weather condition.
	CurrentCondition string
	// CurrentLocation is the current location.
	CurrentLocation  string
)

// Forecast returns the location and weather condition at the current location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
