package main

import "fmt"

func main() {

	circle := NewCircle("Blue")
	square := NewSquare("Red")

	group := GraphicObject{Name: "Group 1"}

	group.Children = append(group.Children, *circle)
	group.Children = append(group.Children, *square)

	group.Print()

	neuronLeft := &Neuron{}

	neuronLayer := &NeuronLayer{Neurons: []*Neuron{&Neuron{}, &Neuron{}, &Neuron{}}}

	Connect(neuronLeft, neuronLayer)

	//Imprimindo: O neuronLeft tem que ter no seu Out todo os neuronios da layer
	//Cade neuronio da Layer tem que ter no seu In o neuronLeft
	fmt.Printf("Neuronio left tem na saida (out) %v\n", neuronLeft.Out)
	for i, neuron := range neuronLayer.Collect() {
		fmt.Printf("Neuronio %d da layer tem na entrada (In) %v\n", i, neuron.In)
	}
}

type GraphicObject struct {
	Name, Color string
	Children    []GraphicObject
}

func (ref *GraphicObject) Print() {
	msg := fmt.Sprintf("%v %v", ref.Color, ref.Name)
	fmt.Println(msg)
	for _, child := range ref.Children {
		child.Print()
	}
}

func NewCircle(color string) *GraphicObject {
	return &GraphicObject{"Circle", color, nil}
}

func NewSquare(color string) *GraphicObject {
	return &GraphicObject{"Square", color, nil}
}

//Sabemos conectar neurônios
type Neuron struct {
	In, Out []*Neuron
}

func (n *Neuron) ConnectTo(other *Neuron) {
	n.Out = append(n.Out, other)
	other.In = append(other.In, n)
}

//Queremos conectar layers
type NeuronLayer struct {
	Neurons []*Neuron
}

//Pelo padrão Composite, precisamos de uma interface única capaz de conectar Neuronios-Neuronios e Layers-Layers
type NeuronInterface interface {
	Collect() []*Neuron //Collect all the neuros of the interface
}

//Implementamos a interface para ambos os elementos, escalar e coleção
func (n *Neuron) Collect() []*Neuron {
	return []*Neuron{n}
}

func (nl *NeuronLayer) Collect() []*Neuron {
	return nl.Neurons
}

//Com essa inteface comum agora podemos ter uma única função capaz de conectar ambos, Neurons-Neurouns e Layers-Layers
func Connect(left NeuronInterface, right NeuronInterface) {
	for _, leftNeuron := range left.Collect() {

		for _, rightNeuron := range right.Collect() {

			leftNeuron.ConnectTo(rightNeuron)
		}
	}
}
