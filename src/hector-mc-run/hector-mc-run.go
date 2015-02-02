package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	"runner"
)

func main() {
	train, test, pred, method, params := runner.PrepareParams()

	action, _ := params["action"]

	classifier := runner.GetMutliClassClassifier(method)

	profile, _ := params["profile"]
	if profile != "" {
		fmt.Printf("Profile data => %s\n", profile)
		f, err := os.Create(profile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	if action == "" {
		accuracy, _ := runner.MultiClassRun(classifier, train, test, pred, params)
		fmt.Println("accuracy : ", accuracy)
	} else if action == "train" {
		runner.MultiClassTrain(classifier, train, params)

	} else if action == "test" {
		accuracy, _ := runner.MultiClassTest(classifier, test, pred, params)
		fmt.Println("accuracy", accuracy)
	}
}
