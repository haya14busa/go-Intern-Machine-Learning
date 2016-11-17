package classification

import "github.com/haya14busa/go-Intern-Machine-Learning/ml"

type Classifier interface {
	Train(trainingSet ml.DataSet)
	Predict(features ml.Features) ml.Label
}
