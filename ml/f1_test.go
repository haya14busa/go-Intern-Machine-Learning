package ml

import (
	"math"
	"testing"
)

func almostEqual(a, b, tolerance float64) bool {
	return a == b || math.Abs(a-b) <= tolerance
}

func TestF1(t *testing.T) {
	dataset := []struct {
		Predicted  Label
		Supervised Label
	}{
		{Positive, Positive},
		{Negative, Positive},
		{Positive, Positive},
		{Negative, Positive},
		{Positive, Negative},
		{Positive, Positive},
		{Positive, Negative},
		{Negative, Negative},
		{Negative, Positive},
	}

	want := 0.545455

	f1 := &F1{}

	for _, data := range dataset {
		f1.Add(data.Predicted, data.Supervised)
	}

	if got := f1.Accuracy(); !almostEqual(got, want, 1e-6) {
		t.Errorf("F1 score: got %v, want %v", got, want)
	}
}
