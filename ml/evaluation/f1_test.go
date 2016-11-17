package evaluation

import (
	"math"
	"testing"

	"github.com/haya14busa/go-Intern-Machine-Learning/ml"
)

func almostEqual(a, b, tolerance float64) bool {
	return a == b || math.Abs(a-b) <= tolerance
}

func TestF1(t *testing.T) {
	dataset := []struct {
		Predicted  ml.Label
		Supervised ml.Label
	}{
		{ml.Positive, ml.Positive},
		{ml.Negative, ml.Positive},
		{ml.Positive, ml.Positive},
		{ml.Negative, ml.Positive},
		{ml.Positive, ml.Negative},
		{ml.Positive, ml.Positive},
		{ml.Positive, ml.Negative},
		{ml.Negative, ml.Negative},
		{ml.Negative, ml.Positive},
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
