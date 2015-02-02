package main

import (
	"os"
	"strings"

	"combine"
	"core"
	"runner"
)

func main() {
	train, _, _, _, params := runner.PrepareParams()

	feature_combination := combine.CategoryFeatureCombination{}
	feature_combination.Init(params)

	dataset := core.NewRawDataSet()
	dataset.Load(train)

	combinations := feature_combination.FindCombination(dataset)

	output := params["output"]

	file, _ := os.Create(output)
	defer file.Close()

	for _, combination := range combinations {
		file.WriteString(strings.Join(combination, "\t") + "\n")
	}
}
