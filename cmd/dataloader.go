package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/haya14busa/go-Intern-Machine-Learning/ml"
)

func LoadCSVDataSet(fname, targetLabel string) (ml.DataSet, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, fmt.Errorf("fail to open %v: %v", fname, err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	dataset := make(ml.DataSet, len(records))

	for i, record := range records {
		label := labelOf(targetLabel, record[len(record)-1])
		var fs ml.Features = make(ml.Features, len(record)-1)
		for j, fstr := range record[:len(record)-1] {
			f, err := strconv.ParseFloat(fstr, 64)
			if err != nil {
				return nil, err
			}
			fs[j] = f
		}
		dataset[i] = ml.Data{
			Features: fs,
			Label:    label,
		}
	}
	return dataset, nil
}

func labelOf(targetLabel, label string) ml.Label {
	if targetLabel == label {
		return ml.Positive
	}
	return ml.Negative
}
