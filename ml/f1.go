package ml

var _ Evaluator = &F1{}

// F1 represents F1 score evaluator.
type F1 struct {
	// NOTE: you can add fields
}

// Add aggregates true/false and positive/negative label.
func (f *F1) Add(predicted, supervised Label) {
	// NOTE: aggregate true/false and positive/negative here.
}

// Accuracy caluculates F1 score.
func (f *F1) Accuracy() float64 {
	// NOTE: caluculate F1 score here.
	return 0.0
}
