package classification

import (
	"log"

	"github.com/haya14busa/go-Intern-Machine-Learning/ml"
)

var _ Classifier = &Perceptron{}

type Perceptron struct {
	w   []float64 // weight vector
	eta float64   // $\eta$
}

func (p *Perceptron) Train(trainingSet ml.DataSet) {
	if len(trainingSet) == 0 {
		return
	}
	if p.eta == 0.0 {
		p.eta = 0.1
	}
	p.w = make([]float64, len(trainingSet[0].Features)+1)
	for i := range p.w {
		p.w[i] = 1
	}
	for _, data := range trainingSet {
		y := data.Label.ToInt()
		fx := p.Predict(data.Features).ToInt()
		if y*fx <= 0 {
			for i, _ := range p.w {
				x := 1.0
				if i > 0 {
					x = data.Features[i-1]
				}
				p.w[i] += p.eta * float64(y) * x
			}
		}
	}
}

func (p *Perceptron) Predict(features ml.Features) ml.Label {
	if len(features)+1 != len(p.w) {
		log.Printf("must be len(features) +1 == len(w): len(features) == %v, len(w) == %v",
			len(features), len(p.w))
		return ml.Negative
	}
	v := 0.0
	v += 1 * p.w[0]
	for j, xi := range features {
		wi := p.w[j+1]
		v += xi * wi
	}
	if v > 0 {
		return ml.Positive
	}
	return ml.Negative
}
