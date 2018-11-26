package nnold

import (
	"math"
)

const (
	LAYER_INPUT = iota
	LAYER_HIDDEN = iota
	LAYER_OUTPUT = iota
)

type nn_size struct {
	outputLen          int
	inputLen           int
	hiddenLayersAmount int
	hiddenLayersLen    []int // the lenght of the slice should be equal to hiddenLayersLen
}

type neuron struct {
	axons []axon
	in    float64
}

type layer []neuron

func (l *layer)prepare(len, type, nextOutLen int) []*float64 {
	l = make([]neuron, len)
	switch type {
	case LAYER_INPUT, Layer_HIDDEN:
		for i := range l {
			l[i].axons = make([]axon, nextOutLen)
		}
	case LAYER_OUTPUT:
		for i := range l {
			l[i].axons = nil
		}
	}
	var lInput []*float64 = make([]*foat64, len(l))
	for i := range l {
		lInput[i] = &l[i].in
	}
	return lInput
}

func (l *layer)run(output []*float64, size *nn_size) {
	for 
}

type axon struct {
	mult float64
}

type neuralNetwork struct {
	layers []layer
	size   *nn_size
	inputs [][]*float64
	output []float64
}

func (nn *neuralNetwork)clear() {
	for i := 0; i < nn.size.inputLen; i++ {
		*nn.inputs[0][i] = 0
	}

	for i := 0; i < nn.size.hiddenLayersAmount; i++ {
		for j := 0; j < nn.size.hiddenLayersLen[i]; j++ {
			*nn.inputs[i+1][j] = 0
		}
	}

	for i := 0; i < nn.size.outputLen; i++ {
		*nn.inputs[nn.size.hiddenLayersAmount+1][i] = 0
	}
}

func (nn *neuralNetwork)run() {
	for i := 0; i < nn.size.hiddenLayersAmount + 1; i++ {
		
	}
	/*
	Output layer
	*/
}

// Non method functions

func createNN(size *nn_size) (retNN neuralNetwork) {
	retNN = neuralNetwork{
		size: size,
		layers: make([]layer, size.hiddenLayersAmount+2),
		inputs: make([][]*float64, size.hiddenLayersAmount+2)
		output: make([]float64, size.outputLen)
	}

	for i := range retNN.inputs {
		var inputSize int
		switch i {
		case 0:
			inputSize = size.inputLen
		case size.hiddenLayersAmount + 2:
			inputSize = size.outputLen
		default:
			inputSize = size.hiddenLayersLen[i-1]
		}
		retNN.inputs[i] = make([]*float64, inputSize)
	}

	var nextLayerLen int
	if size.hiddenLayersAmount != 0 {
		nextLayerLen = size.hiddenLayersLen[0]
	} else {
		nextLayerLen = size.outputLen
	}
	retNN.inputs[0] = retNN.layers[0].prepare(size.inputLen, LAYER_INPUT, nextLayerLen)

	for i := 0; i < size.hiddenLayersAmount; i++ {
		if i == size.hiddenLayersAmount - 1 {
			nextLayerLen = size.outputLen
		} else {
			nextLayerLen = size.hiddenLayersLen[i+1]
		}
		retNN.inputs[i+1] = retNN.layers[i+1].prepare(size.hiddenLayersLen[i], LAYER_HIDDEN, nextLayerLen)
	}

	retNN.inputs[size.hiddenLayersAmount + 1] = retNN.layers[size.hiddenLayersAmount + 1].prepare(size.OutputLen, LAYER_OUTPUT, 0)

	return
}

func sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}
