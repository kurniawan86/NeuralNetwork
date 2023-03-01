package main

import (
	"fmt"
	"testing"
)

func TestNode(t *testing.T) {
	x1 := []float32{1, 2, 3}
	x2 := []float32{4, 5, 2}
	x3 := []float32{4, 5, 6}
	X := make([][]float32, 3)
	X[0] = x1
	X[1] = x2
	X[2] = x3
	fmt.Println(X)
	//weight := []float32{0.1, 0.1, 0.1}

	nn1 := Neuron{
		Weights:  nil,
		function: "",
	}

	InitialNeuron(&nn1, 3)
	fmt.Println(nn1.function)
	fmt.Println(nn1.Weights)

	tot := LinearEq(nn1.Weights, X[0])
	fmt.Println("output = ", tot)

	node := Node{
		NumberNeuron: 0,
		Neuron:       nil,
		Bias:         0,
	}

	AddNode(&node, 2)
	//fmt.Println("number of neurons : ", node.NumberNeuron)
	for i := 0; i < node.NumberNeuron; i++ {
		InitialNeuron(&node.Neuron[i], 3)
		fmt.Println("weights : ", node.Neuron[i].Weights)
	}
}

func TestGetOutputNode(t *testing.T) {
	x1 := []float32{1, 2, 3}
	x2 := []float32{4, 5, 2}
	x3 := []float32{4, 5, 6}
	X := make([][]float32, 3)
	X[0] = x1
	X[1] = x2
	X[2] = x3
	fmt.Println(X)
	//weight := []float32{0.1, 0.1, 0.1}

	nn1 := Neuron{
		Weights:  nil,
		function: "",
	}

	InitialNeuron(&nn1, 3)
	fmt.Println(nn1.function)
	fmt.Println(nn1.Weights)

	tot := LinearEq(nn1.Weights, X[0])
	fmt.Println("output = ", tot)

	node := Node{
		NumberNeuron: 0,
		Neuron:       nil,
		Bias:         0,
	}

	AddNode(&node, 2)
	//fmt.Println("number of neurons : ", node.NumberNeuron)
	for i := 0; i < node.NumberNeuron; i++ {
		InitialNeuron(&node.Neuron[i], 3)
		fmt.Println("weights : ", node.Neuron[i].Weights)
	}

	//get output
	for i := 0; i < node.NumberNeuron; i++ {
		GetOuput(&node, X[0], i)
	}

	//get information
	ViewInformationNode(&node)
}

func TestCreateFirstNode(t *testing.T) {
	x1 := []float32{1, 2, 3}
	x2 := []float32{4, 5, 2}
	x3 := []float32{4, 5, 6}
	X := make([][]float32, 3)
	X[0] = x1
	X[1] = x2
	X[2] = x3

	node := Node{
		NumberNeuron: 0,
		Neuron:       nil,
		Bias:         0,
	}
	CreateFirstNode(&node, X)
	for i := 0; i < node.NumberNeuron; i++ {
		GetOuput(&node, X[0], i)
	}
	ViewInformationNode(&node)
}

func TestAddFirstLayer(t *testing.T) {
	x1 := []float32{1, 2, 3}
	x2 := []float32{4, 5, 2}
	x3 := []float32{4, 5, 6}
	X := make([][]float32, 3)
	X[0] = x1
	X[1] = x2
	X[2] = x3

	layer := Layer{}
	AddFirstLayer(&layer, 2, X)
	fmt.Println(len(layer.Nodes))
	for i, _ := range layer.Nodes[0].Node {
		GetOuput(&layer.Nodes[0].Node[i], X[0], i)
		fmt.Println(layer.Nodes[0].Node[i])
	}
	AddLayer(&layer, 3)
	fmt.Println("layer nodes :", len(layer.Nodes))
	fmt.Println("layer nodes node :", len(layer.Nodes[0].Node))
	fmt.Println("layer nodes node :", len(layer.Nodes[1].Node))
	AddLayer(&layer, 5)
	fmt.Println("layer nodes node :", len(layer.Nodes[2].Node))
	StrukturNode(&layer)
}

func TestOutputStrukturNN(t *testing.T) {
	x1 := []float32{1, 2, 3}
	x2 := []float32{4, 5, 2}
	x3 := []float32{4, 5, 6}
	X := make([][]float32, 3)
	X[0] = x1
	X[1] = x2
	X[2] = x3

	layer := Layer{}
	AddFirstLayer(&layer, 2, X)
	fmt.Println(len(layer.Nodes))
	for i, _ := range layer.Nodes[0].Node {
		GetOuput(&layer.Nodes[0].Node[i], X[0], i)
		fmt.Println(layer.Nodes[0].Node[i])
	}
	AddLayer(&layer, 3)
	fmt.Println("layer nodes :", len(layer.Nodes))
	fmt.Println("layer nodes node :", len(layer.Nodes[0].Node))
	fmt.Println("layer nodes node :", len(layer.Nodes[1].Node))
	AddLayer(&layer, 5)
	fmt.Println("layer nodes node :", len(layer.Nodes[2].Node))
	StrukturNode(&layer)
}
