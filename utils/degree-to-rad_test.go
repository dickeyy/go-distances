package utils

import (
	"testing"
)

func TestDegreeToRad(t *testing.T) {
	deg := 90.0
	rad := DegreeToRad(deg)
	want := 1.5707963267948966
	if rad != want {
		t.Errorf("got %f, want %f", rad, want)
	}
}
