package weather

import "testing"

var convertFtoCTests = []struct{
	name string
	f 	 float64		//input
	want float64	//output
} {
	{"Freezing", 32.0, 0.0},
	{"Boiling", 212.0, 100.0},
	{"Negative", -40.0, -40.0},
	
}
func TestConvertFtoC(t *testing.T){
	for _, tt := range convertFtoCTests {
		t.Run(tt.name, func(t *testing.T){
			got := ConvertFtoC(tt.f)
			if got != tt.want {
				t.Errorf("got %f, want %f", got, tt.want)
			}
		})
	}
}