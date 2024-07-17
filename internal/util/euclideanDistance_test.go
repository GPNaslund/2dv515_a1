package util

import "testing"

func TestEuclideanDistance_ShouldThrowErrorOnInvalidInputLengths(t *testing.T) {
	slice1 := []float64{1.0, 2.0}
	slice2 := []float64{1.0}

	_, err := EuclideanDistance(slice1, slice2)
	if err == nil {
		t.Fatalf("Did not throw error on two slices with different lengths")
	}
}

func TestEuclideanDistance_ShouldNotThrowErrorOnValidInputLengths(t *testing.T) {
	slice1 := []float64{1.0, 2.0}
	slice2 := []float64{1.0, 2.0}

	_, err := EuclideanDistance(slice1, slice2)
	if err != nil {
		t.Fatalf("Threw error on valid input lengths")
	}
}

func TestEuclideanDistance_ShouldReturnCorrectValue_SingleComparison(t *testing.T) {
	p1 := []float64{4.5, 4.0, 1.0}
	p2 := []float64{3.5, 3.5, 2.5}
	expectedScore := 0.22

	result, err := EuclideanDistance(p1, p2)
	if err != nil {
		t.Fatalf("Threw error when trying to calculate: %s", err)
	}

	if result != expectedScore {
		t.Fatalf("Expected: %v, but got: %v", expectedScore, result)
	}
}

func TestEuclideanDistance_ShouldReturnCorrectValue_MultipleComparsion(t *testing.T) {
	base1 := []float64{4.5, 4.0, 1.0}

	comp1 := []float64{3.5, 3.5, 2.5}
	expected1 := 0.22

	comp2 := []float64{3.5, 5.0, 3.5}
	expected2 := 0.11

	comp3 := []float64{3.5, 4.0, 2.5}
	expected3 := 0.24

	comp4 := []float64{4.0, 3.0, 2.0}
	expected4 := 0.31

	comp5 := []float64{4.0, 5.0, 3.5}
	expected5 := 0.12

	comparsions := [][]float64{comp1, comp2, comp3, comp4, comp5}
	expectedValues := []float64{expected1, expected2, expected3, expected4, expected5}

	for i, val := range comparsions {
		result, err := EuclideanDistance(base1, val)
		if err != nil {
			t.Fatalf("Threw an error: %s", err)
		}
		if result != expectedValues[i] {
			t.Fatalf("Expected: %v but got: %v", expectedValues[i], result)
		}
	}
}
