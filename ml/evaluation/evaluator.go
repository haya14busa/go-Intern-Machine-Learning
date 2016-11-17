package evaluation

import "github.com/haya14busa/go-Intern-Machine-Learning/ml"

// Evaluator represents interface for evaluation of machine learning
// implementation.
type Evaluator interface {
	Add(predicted, supervised ml.Label)
	Accuracy() float64
}
