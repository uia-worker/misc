package conv

import (
	"reflect"
	"testing"
	//"fmt"
)

/**
	Mal for testfunksjoner
	Du skal skrive alle funksjonene basert på denne malen
	For alle konverteringsfunksjonene (tilsammen 6)
	kan du bruke malen som den er (du må selvsagt endre
	funksjonsnavn og testverdier)
	Korrekt skrivemåte: Fahrenheit
*/
// Celsius = (Farhrenheit - 32)*(5/9)
func TestFahrenheitToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 134, want: 56.66666666666667},
	}

	for _, tc := range tests {
		got := FahrenheitToCelsius(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

// Farhrenheit = Celsius*(9/5) + 32
func TestCelsiusToFahrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 56.67, want: 134.006},
	}

	for _, tc := range tests {
		got := CelsiusToFahrenheit(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

// Farhrenheit = (Kelvin - 273.15)*(9/5) + 32
func TestKelvinToFahrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 273.15, want: 32},
	}

	for _, tc := range tests {
		got := KelvinToFahrenheit(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

// Kelvin = (Farhenheit - 32) * (5/9) + 273.15
//
func TestFahrenheitToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 32, want: 273.15},
	}

	for _, tc := range tests {
		got := FahrenheitToKelvin(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

// Kelvin = Celsius + 273.15
func TestCelsiusToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 0, want: 273.15},
	}

	for _, tc := range tests {
		got := CelsiusToKelvin(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

// Celsius = Kelvin - 273.15
func TestKelvinToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 0, want: -273.15},
	}

	for _, tc := range tests {
		got := KelvinToCelsius(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestFormatBigInteger(t *testing.T) {
	type test struct {
		input float64
		want string
	}

	tests := []test {
		{input: 15000.5656, want: "15 001"},
		{input: 15000000, want: "15 000 000"},
		{input: 1000, want: "1 000"},
		{input: 150000000, want: "150 000 000"},
		{input: 100, want: "100"},
		{input: 10, want: "10"},
		{input: 1, want: "1"},
		{input: -10000, want: "-10 000"},
		{input: -150000000, want: "-150 000 000"},
	}

	for _, tc := range tests {
 	 got := FormatBigInteger(tc.input)
	 //fmt.Println("[]byte(got) = ", []byte(got), "[]byte(want) = ", []byte(tc.want))
 	 if !reflect.DeepEqual(tc.want, got) {
 		 t.Errorf("expected: %v, got: %v", tc.want, got)
 	 }
 }

}
