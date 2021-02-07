package main

import "fmt"

//Agregação múltipla
type Bird struct {
	age int
}

func NewBird(age int) *Bird {
	return &Bird{age: age}
}

func (ref *Bird) Fly() {
	fmt.Println("Flying")
}

type Lizard struct {
	age int
}

func NewLizard(age int) *Lizard {
	return &Lizard{age: age}
}

func (ref *Lizard) Crawl() {
	fmt.Println("Crawling")
}

//Mantém a referência das outras implementações, e adiciona novo comportamento (método Roar)
type Dragon struct {
	Bird   *Bird
	Lizard *Lizard
}

func NewDragon(age int) *Dragon {
	return &Dragon{
		Bird:   NewBird(age),
		Lizard: NewLizard(age),
	}
}

func (ref *Dragon) Fly() {
	ref.Bird.Fly()
}

func (ref *Dragon) Crawl() {
	ref.Lizard.Crawl()
}

func (ref *Dragon) Roar() {
	fmt.Println("roaring")
}

//Decorator usando agregação
