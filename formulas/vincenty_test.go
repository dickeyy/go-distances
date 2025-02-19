package formulas

import (
	"math"
	"testing"
)

func TestVincentyZeroGivesZero(t *testing.T) {
	lon1, lat1, lon2, lat2, earthRadius := 0.00, 0.00, 0.00, 0.00, 6371.00
	distance := Vincenty(lat1, lon1, lat2, lon2, earthRadius)
	want := 0.00
	if distance != want {
		t.Errorf("got %f, want %f", distance, want)
	}
}

// based on kdickey.json
func TestVincenty(t *testing.T) {
	earthRadius := 6967404.0

	lat1, lon1 := 75.20479441439075, -87.42362995032933
	lat2, lon2 := -63.27864890563778, -9.284284051915705

	distance := int(math.Round(Vincenty(lat1, lon1, lat2, lon2, earthRadius)))
	want := 17892709

	if distance != want {
		t.Errorf("got %d, want %d", distance, want)
	}
}
