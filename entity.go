package main

type Neuron struct {
	Weights  []float32
	function string
}

type Node struct {
	NumberNeuron int
	Neuron       []Neuron
	Bias         float32
}

type Nodes struct {
	Node []Node
}
