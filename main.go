package main

import (
	"fmt"
	"os"
)

func main() {
	filePtr, err := os.Open("network1.json")
	if err != nil {
		panic(err)
	}

	weights, err := loadWeightsFromJSON(filePtr)
	if err != nil {
		panic(err)
	}
	if valid, reason := weights.validate(); !valid {
		panic("Neural network in \"network1.json\" is invalid, reason: " + reason)
	}

	fmt.Println(weights)

	/*network1 := createNNFromWeights(weights)

	result := network1.use([]float64{1, 0})
	fmt.Println(result)
	fmt.Println(network1)

	weights.completeRandomizer()
	fmt.Println(weights)
	weights.completeRandomizer()
	fmt.Println(weights)*/
	/*weights.mutate(1)
	fmt.Println(weights)*/
	weights.mutate(1)
	fmt.Println("\n", weights)
}

var exampleWeights weightVals = weightVals{
	[][]float64{
		[]float64{
			5.0, 6.0,
		},
		[]float64{
			3, 5,
		},
	},
	[][]float64{
		[]float64{
			1, 7,
		},
		[]float64{
			2, 3,
		},
	},
}
