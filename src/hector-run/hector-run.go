package main

import (
	"fmt"

	"runner"
)

func main() {
	train, test, pred, method, params := runner.PrepareParams()

	action, _ := params["action"]

	classifier := runner.GetClassifier(method)

	if action == "" {
		auc, _, _ := runner.AlgorithmRun(classifier, train, test, pred, params)
		fmt.Println("AUC:")
		fmt.Println(auc)
	} else if action == "train" {
		runner.AlgorithmTrain(classifier, train, params)

	} else if action == "test" {
		auc, _, _ := runner.AlgorithmTest(classifier, test, pred, params)
		fmt.Println("AUC:")
		fmt.Println(auc)
	}
}
