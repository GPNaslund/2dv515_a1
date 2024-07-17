package util

import "testing"

func TestPearsonScore_ShouldThrowErrorOnInvalidInput(t *testing.T) {
	input1 := []float64{1.0, 2.0}
	input2 := []float64{1.0}

	_, err := PearsonScore(input1, input2)
	if err == nil {
		t.Fatalf("Did not throw error on invalid input: %s", err)
	}
}

func TestPearsonScore_ShouldNotThrowErrorOnValidInput(t *testing.T) {
	input1 := []float64{1.0}
	input2 := []float64{1.0}

	_, err := PearsonScore(input1, input2)
	if err != nil {
		t.Fatalf("Threw error when not supposed to: %v", err)
	}
}

func TestPearsonScore_ShouldReturnCorrect_SingleComparsion(t *testing.T) {
	p1 := []float64{4.5, 4.0, 1.0}
	p2 := []float64{3.5, 3.5, 2.5}
	expectedScore := 0.99

	result, err := PearsonScore(p1, p2)
	if err != nil {
		t.Fatalf("Threw error when trying to calculate: %s", err)
	}

	if result != expectedScore {
		t.Fatalf("Expected: %v, but got: %v", expectedScore, result)
	}
}

func TestPearsonScore_ShouldReturnCorrect_MultipleComparsions(t *testing.T) {
	base1 := []float64{4.5, 4.0, 1.0}

	comp1 := []float64{3.5, 3.5, 2.5}
	expected1 := 0.99

	comp2 := []float64{3.5, 5.0, 3.5}
	expected2 := 0.38

	comp3 := []float64{3.5, 4.0, 2.5}
	expected3 := 0.89

	comp4 := []float64{4.0, 3.0, 2.0}
	expected4 := 0.92

	comp5 := []float64{4.0, 5.0, 3.5}
	expected5 := 0.66

	comparsions := [][]float64{comp1, comp2, comp3, comp4, comp5}
	expectedValues := []float64{expected1, expected2, expected3, expected4, expected5}

	for i, val := range comparsions {
		result, err := PearsonScore(base1, val)
		if err != nil {
			t.Fatalf("Threw an error: %s", err)
		}
		if result != expectedValues[i] {
			t.Fatalf("Expected: %v but got: %v", expectedValues[i], result)
		}
	}
}
