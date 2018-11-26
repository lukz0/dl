package main

import (
	"fmt"
	"math"
	"strings"
)

type weightVals [][][]float64

func (wv *weightVals) validate() (bool, string) {
	// Check if there are layers at all
	if len(*wv) <= 0 {
		return false, "No layers"
	}

	// Check for empty layers
	for i := range *wv {
		if len((*wv)[i]) <= 0 {
			return false, "Empty layer"
		}
	}

	// Check if the amount of axons is the same as the amount of neurons in the next layer
	for i := 0; i < len(*wv)-1; i++ {
		for j := range (*wv)[i] {
			if len((*wv)[i+1]) != len((*wv)[i][j]) {
				return false, "Invalid amount of axons on a hidden neuron"
			}
		}
	}

	// Check if the first neuron in the last layer has axons
	var outputLen int = len((*wv)[len(*wv)-1][0])
	if outputLen <= 0 {
		return false, "No output axons"
	}

	// Check if the neurons in the last layer has the same amount of axons
	for i := range (*wv)[len(*wv)-1] {
		if len((*wv)[len(*wv)-1][i]) != outputLen {
			return false, "Amount of axons in the last layer varies"
		}
	}

	return true, ""
}

type axon struct {
	weight float64
}

type neuron struct {
	axons []axon
	value float64
}

type layer []neuron

type nn struct {
	layers []layer
	input  []float64
	output []float64
}

// Processes the values in the first layer and outputs it in output
// The sigmoid function isn't used on output
func (n *nn) run() {
	for i := 0; i < len(n.layers)-1; i++ {
		for j := range n.layers[i] {
			// n.layers[i][j] is a neutron
			var sigVal float64 = sigmoid(n.layers[i][j].value)
			for k := range n.layers[i][j].axons {
				// TODO
				n.layers[i+1][k].value += sigVal * n.layers[i][j].axons[k].weight

			}
		}
	}

	// run the last layer without sigmoid
	for i := range n.layers[len(n.layers)-1] {
		for j := range n.layers[len(n.layers)-1][i].axons {
			n.output[j] += n.layers[len(n.layers)-1][i].value * n.layers[len(n.layers)-1][i].axons[j].weight
		}
	}
}

func (n *nn) loadInput() {
	for i := range n.input {
		n.layers[0][i].value = n.input[i]
	}
}

func (n *nn) clearNeuronValues() {
	for i := range n.layers {
		for j := range n.layers[i] {
			n.layers[i][j].value = 0
		}
	}
}

func (n *nn) clearOutput() {
	for i := range n.output {
		n.output[i] = 0
	}
}

func (n *nn) use(input []float64) []float64 {
	n.clearNeuronValues()
	n.clearOutput()
	for i := range input {
		n.input[i] = input[i]
	}
	n.run()

	return n.output
}

func (n nn) String() string {
	var outputBuilder strings.Builder
	outputBuilder.WriteString("Layers:\n")
	for i := range n.layers {
		outputBuilder.WriteString(fmt.Sprintf("Layer %d:\n", i))
		for j := range n.layers[i] {
			outputBuilder.WriteString(fmt.Sprintf("\tValue: %v\n\tAxons: %v\n\n", n.layers[i][j].value, n.layers[i][j].axons))
		}
	}

	outputBuilder.WriteString(fmt.Sprintf("Input values %v\n", n.input))
	outputBuilder.WriteString(fmt.Sprintf("Output values %v\n", n.output))

	return outputBuilder.String()
}

// Squeezes the whole real number line between 0 and 1
func sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

/*
 * Length of weights[layer][neuron+1] must be equal to
 * the length of weights[layer][neuron][axon]
 */
func createNNFromWeights(weights weightVals) (retNN nn) {
	// Create layers
	retNN.layers = make([]layer, len(weights))

	// Create neurons
	for i := range retNN.layers {
		retNN.layers[i] = make(layer, len(weights[i]))

		// Create axons
		for j := range retNN.layers[i] {
			retNN.layers[i][j].axons = make([]axon, len(weights[i][j]))

			// Assign weights to axons
			for k := range retNN.layers[i][j].axons {
				retNN.layers[i][j].axons[k].weight = weights[i][j][k]
			}
		}
	}

	// Create input array
	retNN.input = make([]float64, len(weights[0]))

	// Create an output array
	retNN.output = make([]float64, len(weights[len(weights)-1][0]))
	return
}
