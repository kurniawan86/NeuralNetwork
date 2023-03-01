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

func AddFirstLayer(layer *Layer, nNode int, x [][]float32) {
	nodes := Nodes{}
	for i := 0; i < nNode; i++ {
		tempNode := Node{}
		CreateFirstNode(&tempNode, x)
		nodes.Node = append(nodes.Node, tempNode)
	}
	layer.Nodes = append(layer.Nodes, nodes)
}

func ViewInformationNeuron(neuron Neuron) {
	fmt.Println("weight", neuron.Weights)
}

func AddLayer(layer *Layer, nNode int) {
	//fmt.Println(len(layer.Nodes))
	n := len(layer.Nodes)
	//fmt.Println(len(layer.Nodes[n-1].Node))
	m := len(layer.Nodes[n-1].Node)
	var outputs []float32
	for i := 0; i < m; i++ {
		outputs = append(outputs, layer.Nodes[n-1].Node[i].output)
	}

	nodes := Nodes{}

	for j := 0; j < nNode; j++ {
		node := Node{}
		node.NumberNeuron = m
		for i := 0; i < m; i++ {
			tempNeuron := Neuron{}
			node.Neuron = append(node.Neuron, tempNeuron)
			InitialNeuron(&node.Neuron[i], m)
		}
		nodes.Node = append(nodes.Node, node)
	}
	layer.Nodes = append(layer.Nodes, nodes)
}

func StrukturNode(layer *Layer) {
	for _, nodes := range layer.Nodes {
		fmt.Println("Nodes : ", len(nodes.Node))
		for _, node := range nodes.Node {
			fmt.Println("node : ", node.Neuron)
		}
	}
}

func Foward(layer *Layer) {

}
