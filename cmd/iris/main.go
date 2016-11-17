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
	Iteration = 1
)

var (
	training    string
	test        string
	iter        int
	targetLabel string
)

func init() {
	flag.StringVar(&training, "training", "", "training data set")
	flag.StringVar(&test, "test", "", "test data set")
	flag.IntVar(&iter, "i", -1, "iteration (max size)")
	flag.StringVar(&targetLabel, "label", "Iris-setosa", "target label (Iris-setosa/Iris-versicolor/Iris/virginica)")
}

func main() {
	flag.Parse()
	if err := Main(os.Stdout, training, test, iter, targetLabel); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Main(w io.Writer, training, test string, trainSize int, targetLabel string) error {
	if training == "" || test == "" {
		flag.Usage()
		return nil
	}
	tr, te, err := RunClassification(training, test, &classification.Perceptron{}, trainSize, targetLabel)
	if err != nil {
		return err
	}
	fmt.Fprintln(w, tr, te)
	return nil
}

func RunClassification(training, test string, c classification.Classifier, trainSize int, targetLabel string) (trainAccuracy, testAccuracy float64, err error) {
	trainAccuracy, err = train(training, c, Iteration, trainSize, targetLabel)
	if err != nil {
		return
	}
	e := &evaluation.F1{}
	testdata, err := cmd.LoadCSVDataSet(test, targetLabel)
	if err != nil {
		return
	}
	evaluate(c, testdata, e)
	testAccuracy = e.Accuracy()
	return trainAccuracy, testAccuracy, nil
}

func train(training string, c classification.Classifier, iteration int, trainSize int, targetLabel string) (float64, error) {
	dataset, err := cmd.LoadCSVDataSet(training, targetLabel)
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
