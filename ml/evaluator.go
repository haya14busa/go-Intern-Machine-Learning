package ml

// Evaluator represents interface for evaluation of machine learning
// implementation.
type Evaluator interface {
	Add(predicted, supervised Label)
	Accuracy() float64
}
