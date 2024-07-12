package util

import (
	"fmt"
	"math"
)

func PearsonScore(values1 []float64, values2 []float64) (float64, error) {
	if len(values1) != len(values2) {
		return 0.0, fmt.Errorf("Length of both slices must be the same")
	}

  amountOfValues := float64(len(values1))

  var sum1 float64
  var sum1sq float64
  var sum2 float64
  var sum2sq float64
  var productSum float64

  for i, val := range values1 {
    sum1 += val
    sum1sq += math.Pow(val, 2)
    productSum += val * values2[i]
  }

  for _, val := range values2 {
    sum2 += val
    sum2sq += math.Pow(val, 2)
  }

  numerator := productSum - (sum1 * sum2 / amountOfValues)
  denumerator := math.Sqrt((sum1sq-math.Pow(sum1, 2) / amountOfValues) * (sum2sq - math.Pow(sum2, 2) / amountOfValues))
  
  if denumerator == 0 {
    return 0.0, nil
  }

  result := numerator / denumerator

  return result, nil
}
