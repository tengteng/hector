package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
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
		f, err := os.Create(profile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	average_accuracy := 0.0
	for part := 0; part < total; part++ {
		train, test := SplitFile(dataset, total, part)
		classifier := runner.GetMutliClassClassifier(method)
		classifier.Init(params)
		accuracy := runner.MultiClassRunOnDataSet(classifier, train, test, "", params)
		fmt.Println("accuracy : ", accuracy)
		average_accuracy += accuracy
		classifier = nil
		train = nil
		test = nil
		runtime.GC()
	}
	fmt.Println(average_accuracy / float64(total))
}
