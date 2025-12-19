package weather

import "testing"

var convertFtoCTests = []struct{
	name string
	f 	 float64		//input
	c float64	//output
} {
	{"Freezing", 32.0, 0.0},
	{"Boiling", 212.0, 100.0},
	{"Negative", -40.0, -40.0},
	
}
func TestConvertFtoC(t *testing.T){
	for _, tt := range convertFtoCTests {
		t.Run(tt.name, func(t *testing.T){
			got := ConvertFtoC(tt.f)
			if got != tt.c {
				t.Errorf("got %f, want %f", got, tt.c)
			}
		})
	}
}

func TestConvertCToF(t *testing.T){
	for _, tt := range convertFtoCTests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvertCToF(tt.c)
			if got != tt.f {
				t.Errorf("got %f, want %f", got, tt.f)
			}
		})
	}
}