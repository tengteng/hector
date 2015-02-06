package main

import (
	"fmt"

	"preprocessor"
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
	} else if action == "preprocess" {
		if execution_plan_path, ok := params["execution_plan_path"]; ok {
			fmt.Println("Preprocess ", execution_plan_path)
			preprocessor.Run(execution_plan_path)
		} else {
			fmt.Println("No execution_plan_path found in config.")
		}
	}
}
