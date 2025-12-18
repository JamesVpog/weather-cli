// Convert from one temperature to another temperature
package weather

func ConvertFtoC(f float64) float64 {
	// returns the equivalent farenheit in celsius
	return (f - 32) * 5/9
}