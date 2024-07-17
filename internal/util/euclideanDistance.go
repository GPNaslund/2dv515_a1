package util

import (
	"fmt"
	"math"
)

func EuclideanDistance(values1 []float64, values2 []float64) (float64, error) {
	if len(values1) != len(values2) {
		return 0.0, fmt.Errorf("Length of both slices must be the same")
	}
	sumDiff := 0.0

	for i := 0; i < len(values1); i++ {
		diff := math.Pow(values1[i]-values2[i], 2)
		sumDiff += diff
	}

	result := 1 / (1 + sumDiff)
	rounded := math.Round(result*100) / 100

	return rounded, nil
}
