package classification

import "github.com/haya14busa/go-Intern-Machine-Learning/ml"

var _ Classifier = &Perceptron{}

type Perceptron struct {
	// NOTE: you can add fields
}

func (p *Perceptron) Train(trainingSet ml.DataSet) {
	// NOTE: implement learning algorithm here.
}

func (p *Perceptron) Predict(features ml.Features) ml.Label {
	// NOTE: implement prediction algorithm here.
	return ml.Positive
}
