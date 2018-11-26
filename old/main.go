package main

import (
	"fmt"
	"math"
)

type nn_size struct {
	outputLen          int
	inputLen           int
	hiddenLayersAmount int
	hiddenLayersLen    []int // This slice should be the same length as "hiddenLayersAmount"
}

var (
	NETWORK_SIZE nn_size = nn_size{
		outputLen:          3,
		inputLen:           4,
		hiddenLayersAmount: 2,
		hiddenLayersLen:    []int{5, 5},
	}
)

func main() {
	var (
		network [][]neuron
		input   []*float64
		output  []float64
	)
	network, input, output = createNeuralNetwork(NETWORK_SIZE)
	fmt.Println(input, "\n", output)
	printNN(network)
	fmt.Println("\n\n", NETWORK_SIZE)
}

func printNN(nn [][]neuron) {
	for i := range nn {
		for j := range nn[i] {
			fmt.Print(nn[i][j], "\t")
		}
		fmt.Println()
	}
}

func createNeuralNetwork(size nn_size) ([][]neuron, []*float64, []float64) {
	// Variables the pointers in output reference to
	var dereferencedOutput []float64 = make([]float64, size.outputLen)

	// Used as output the the neural network
	var output []*float64 = getPtrs(dereferencedOutput)

	// neural_network.network
	var network [][]neuron = make([][]neuron, size.hiddenLayersAmount+2)

	// Create the output layer
	network[len(network)-1] = createOutputLayer(output)

	// Create the hidden layers
	for i := size.hiddenLayersAmount; i > 0; i-- {
		var followingLayerInput []*float64 = getLayerInput(network[i+1])
		axonVals := make([][]float64, size.hiddenLayersLen[i-1])
		for j := range axonVals {
			axonVals[j] = make([]float64, len(followingLayerInput))
		}
		network[i] = createLayer(followingLayerInput, axonVals, size.hiddenLayersLen[i-1])
	}

	// Create the input layer
	inputAxonsVals := make([][]float64, size.inputLen)
	for i := range inputAxonsVals {
		inputAxonsVals[i] = make([]float64, len(network[1]))
	}
	network[0] = createLayer(getLayerInput(network[1]), inputAxonsVals, size.inputLen)
	var input []*float64 = getLayerInput(network[0])
	return network, input, dereferencedOutput
}

func clearArrDeref(a []*float64) {
	for i := range a {
		*a[i] = 0.0
	}
}

func getLayerInput(l []neuron) []*float64 {
	var lInput []*float64 = make([]*float64, len(l))
	for i := range l {
		lInput[i] = &l[i].value
	}
	return lInput
}

func getPtrs(a []float64) []*float64 {
	var b []*float64 = make([]*float64, len(a))
	for i := range a {
		b[i] = &a[i]
	}
	return b
}

func sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

type neuron struct {
	axons []axon
	value float64
}

type axon struct {
	mult float64
}

func (n *neuron) useAxons(output []*float64) {
	for i, v := range n.axons {
		*output[i] += sigmoid(n.value) * v.mult
	}
}

func createLayer(output []*float64, axonVals [][]float64, size int) []neuron {
	// The array inside axonVals[] must be the same length as output
	// The array of arrays axonVals must be the same lengts as "size"
	var layer []neuron = make([]neuron, size)
	for i := range layer {
		layer[i].axons = make([]axon, len(output))
		for j := range layer[i].axons {
			layer[i].axons[j] = axon{mult: axonVals[i][j]}
		}
	}
	return layer
}

func useLayer(layer []neuron, output []*float64) {
	for i := range layer {
		layer[i].useAxons(output)
	}
}

func createOutputLayer(output []*float64) []neuron {
	var layer []neuron = make([]neuron, len(output))
	for i := range layer {
		layer[i].axons = make([]axon, len(output))
		for j := range layer[i].axons {
			layer[i].axons[j] = axon{}
		}
		layer[i].axons[i].mult = 1.0
	}
	return layer
}

func useOutputLayer(layer []neuron, output []*float64) {
	for i := range layer {
		*output[i] = layer[i].value
	}
}
