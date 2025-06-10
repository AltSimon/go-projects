package unitconverter

import "testing"

func TestFahrenheitToCelsius(t *testing.T) {
	result := ConvertUnit(32.0, "F", "C")
	var wants float32 = 0.0
	if result != wants {
		t.Errorf("ConvertUnit(32.0, %q, %q) = %f, should be equal to %f", "F", "C", result, wants)
	}
}

func TestFahrenheitToKelvin(t *testing.T) {
	result := ConvertUnit(44.0, "F", "K")
	var wants float32 = 279.817
	if result != wants {
		t.Errorf("ConvertUnit(32.0, %q, %q) = %f, should be equal to %f", "F", "C", result, wants)
	}
}
