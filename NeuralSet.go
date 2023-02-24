package main

import (
	"fmt"
	"math/rand"
)

func InitialNeuron(neuron *Neuron, n int) {
	if "" == neuron.function {
		neuron.function = "default"
	}
	for i := 0; i < n; i++ {
		a := rand.Float32()
		neuron.Weights = append(neuron.Weights, a)
	}
}

func AddNode(node *Node, nNeuron int) {
	for i := 0; i < nNeuron; i++ {
		tempNeuron := Neuron{}
		node.NumberNeuron = nNeuron
		node.Neuron = append(node.Neuron, tempNeuron)
	}
}

func GetOuput(node *Node, x []float32, index int) {
	m := node.Neuron[index].Weights
	eq := LinearEq(m, x)
	node.output = eq + node.Bias
}

func ViewInformationNode(node *Node) {
	fmt.Println("output : ", node.output)
	fmt.Println("bias : ", node.Bias)
	//nNeuron := len(node.Neuron)
	for _, neu := range node.Neuron {
		fmt.Println(neu.Weights)
	}
}

func CreateFirstNode(node *Node, x [][]float32) {
	if node.Neuron == nil {
		nWeight := len(x[0])
		node.NumberNeuron = nWeight
		for i := 0; i < nWeight; i++ {
			tempNeuron := Neuron{}
			node.Neuron = append(node.Neuron, tempNeuron)
			InitialNeuron(&node.Neuron[i], nWeight)
		}
	}
}
