package evaluation

import "github.com/haya14busa/go-Intern-Machine-Learning/ml"

var _ Evaluator = &F1{}

// F1 represents F1 score evaluator.
type F1 struct {
	tp int // the number of true positive
	fp int // the number of false positive
	fn int // the number of false negative
}

// Add aggregates true/false and positive/negative label.
func (f *F1) Add(predicted, supervised ml.Label) {
	switch predicted {
	case ml.Positive:
		switch supervised {
		case ml.Positive:
			f.tp++
		case ml.Negative:
			f.fp++
		}
	case ml.Negative:
		switch supervised {
		case ml.Positive:
			f.fn++
		case ml.Negative:
			// true negative
		}
	}
}

// Accuracy caluculates F1 score.
func (f *F1) Accuracy() float64 {
	p := f.Precision()
	r := f.Recall()
	if p+r == 0 {
		return 0
	}
	return 2 * ((p * r) / (p + r))
}

func (f *F1) Precision() float64 {
	denom := float64(f.tp + f.fp)
	if denom == 0 {
		return 0
	}
	return float64(f.tp) / denom
}

func (f *F1) Recall() float64 {
	denom := float64(f.tp + f.fn)
	if denom == 0 {
		return 0
	}
	return float64(f.tp) / denom
}
