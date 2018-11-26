package main

import (
	"encoding/json"
	"io"
)

func loadWeightsFromJSON(input io.Reader) (weightVals, error) {
	type jsonValsType struct {
		Length []int64
		Axons  weightVals
	}

	var tempJsonVals jsonValsType

	decoder := json.NewDecoder(input)
	err := decoder.Decode(&tempJsonVals)

	return tempJsonVals.Axons, err
}
