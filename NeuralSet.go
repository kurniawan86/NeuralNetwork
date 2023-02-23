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
		fmt.Println("hello")
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
