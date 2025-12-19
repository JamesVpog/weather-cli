// Convert from one temperature to another temperature
package weather

func ConvertFtoC(f float64) float64 {
	// returns the equivalent fahrenheit in celsius
	return (f - 32) * 5 / 9
}

func ConvertCToF(c float64) float64 {
	// returns the equivalent celsius in fahrenheit
	return (c * 9 / 5) + 32
}
