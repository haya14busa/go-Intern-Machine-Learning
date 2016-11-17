package ml

type Feature float64
type Features []float64

type Data struct {
	Features Features
	Label    Label
}

type DataSet []Data
