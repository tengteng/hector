package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strconv"

	"core"
	"runner"
)

func SplitFile(dataset *core.DataSet, total, part int) (*core.DataSet, *core.DataSet) {

	train := core.NewDataSet()
	test := core.NewDataSet()

	for i, sample := range dataset.Samples {
		if i%total == part {
			test.AddSample(sample)
		} else {
			train.AddSample(sample)
		}
	}
	return train, test
}

func main() {
	train_path, _, _, method, params := runner.PrepareParams()
	global, _ := strconv.ParseInt(params["global"], 10, 64)
	profile, _ := params["profile"]
	dataset := core.NewDataSet()
	dataset.Load(train_path, global)

	cv, _ := strconv.ParseInt(params["cv"], 10, 32)
	total := int(cv)

	if profile != "" {
		fmt.Println(profile)
		f, err := os.Create(profile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	average_auc := 0.0
	for part := 0; part < total; part++ {
		train, test := SplitFile(dataset, total, part)
		classifier := runner.GetClassifier(method)
		classifier.Init(params)
		auc, _ := runner.AlgorithmRunOnDataSet(classifier, train, test, "", params)
		fmt.Println("AUC:")
		fmt.Println(auc)
		average_auc += auc
		classifier = nil
	}
	fmt.Println(average_auc / float64(total))
}
