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
	for i := 0; i < node.NumberNeuron; i++ {
		fmt.Println("output : ", node.output)
		fmt.Println("bias : ", node.Bias)
	}
}
