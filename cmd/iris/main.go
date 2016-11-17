package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/haya14busa/go-Intern-Machine-Learning/cmd"
	"github.com/haya14busa/go-Intern-Machine-Learning/ml"
	"github.com/haya14busa/go-Intern-Machine-Learning/ml/classification"
	"github.com/haya14busa/go-Intern-Machine-Learning/ml/evaluation"
)

const (
	TargetLabel = "Iris-setosa"
	Iteration   = 1
)

var (
	training string
	test     string
	iter     int
)

func init() {
	flag.StringVar(&training, "training", "", "training data set")
	flag.StringVar(&test, "test", "", "test data set")
	flag.IntVar(&iter, "i", -1, "iteration (max size)")
}

func main() {
	flag.Parse()
	if err := Main(os.Stdout, training, test, iter); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Main(w io.Writer, training, test string, trainSize int) error {
	if training == "" || test == "" {
		flag.Usage()
		return nil
	}
	tr, te, err := RunClassification(w, training, test, &classification.Perceptron{}, trainSize)
	if err != nil {
		return err
	}
	fmt.Fprintln(w, tr, te)
	return nil
}

func RunClassification(w io.Writer, training, test string, c classification.Classifier, trainSize int) (trainAccuracy, testAccuracy float64, err error) {
	trainAccuracy, err = train(w, training, c, Iteration, trainSize)
	if err != nil {
		return
	}
	e := &evaluation.F1{}
	testdata, err := cmd.LoadCSVDataSet(test, TargetLabel)
	if err != nil {
		return
	}
	evaluate(c, testdata, e)
	testAccuracy = e.Accuracy()
	return trainAccuracy, testAccuracy, nil
}

func train(w io.Writer, training string, c classification.Classifier, iteration int, trainSize int) (float64, error) {
	dataset, err := cmd.LoadCSVDataSet(training, TargetLabel)
	if err != nil {
		return 0, err
	}
	if trainSize > 0 && len(dataset) > trainSize {
		dataset = dataset[:trainSize]
	}
	for i := 0; i < iteration; i++ {
		c.Train(dataset)
	}
	e := &evaluation.F1{}
	evaluate(c, dataset, e)
	return e.Accuracy(), nil
}

func evaluate(c classification.Classifier, dataset ml.DataSet, e evaluation.Evaluator) {
	for _, data := range dataset {
		e.Add(c.Predict(data.Features), data.Label)
	}
}
