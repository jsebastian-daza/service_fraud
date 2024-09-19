package utils

import (
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetEstimatedDistance(t *testing.T) {
	lat1, lon1 := 34.6037, -58.3816
	lat2, lon2 := 40.4168, -3.7038

	distance := GetEstimatedDistance(lat1, lon1, lat2, lon2)
	distanceNumber, err := strconv.ParseInt(distance, 0, 64)
	require.NoError(t, err)
	if int(distanceNumber) >= 6700 && distanceNumber < 6900 {
		t.Error("expected distance to be greater than 6700 and lower than 6900")
	}
}

func TestToRadians(t *testing.T) {
	tests := []struct {
		degrees float64
		radians float64
	}{
		{0, 0},
		{180, math.Pi},
		{90, math.Pi / 2},
		{360, 2 * math.Pi},
		{45, math.Pi / 4},
	}

	for _, value := range tests {
		result := ToRadians(value.degrees)
		if math.Abs(result-value.radians) > 1e-9 {
			t.Errorf("ToRadians(%f) = %f; want %f", value.degrees, result, value.radians)
		}
	}
}
