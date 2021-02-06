package main

import "fmt"

func main() {

	//functional factory generator
	developerFactory := NewEmployeeFactory("developer", 60000)
	developer := developerFactory("Peter")

	managerFactory := NewEmployeeFactory("manager", 120000)
	manager := managerFactory("Elias")

	//structural factory generator
	bossFactory := NewEmployeeFactory2("boss", 240000)
	boss := bossFactory.Create("John")

	fmt.Println(developer, manager, boss)

	//prototype factory: preconfigured objects
	bossSam := NewEmployee("Sam", 2)
	fmt.Println(bossSam)
}

//-------------------SIMPLE FACTORY: FUNCTION FACTORY----------------------
type Person struct {
	name string
	age  int
}

//The Function Factory
func NewPerson(name string, age int) Person {
	return Person{
		name: name,
		age:  age,
	}
}

//--------------------FACTORY GENERATOR-------------------
type Employee struct {
	name, position string
	annualIncome   int
}

//Functional Factory Generator: returns a function that creates an instance of an object (high order functions)
//Advantage: factories can be passed as arguments
func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name: name, position: position, annualIncome: annualIncome}
	}
}

//Structural Factory Generator
type EmployeeFactory struct {
	position     string
	annualIncome int
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.position, f.annualIncome}
}

func NewEmployeeFactory2(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{position, annualIncome}
}

//---------------------------------PROTOTYPE FACTORY--------------------
//Factory that creates preconfigure objects (similar objects)
type Employee2 struct {
	name, position string
	annualIncome   int
}

const (
	Developer = iota
	Manager
	Boss
)

func NewEmployee(name string, role int) *Employee2 {
	switch role {
	case Developer:
		return &Employee2{name, "developer", 60000}
	case Manager:
		return &Employee2{name, "manager", 120000}
	case Boss:
		return &Employee2{name, "boss", 240000}
	default:
		panic("unsuported role")
	}
}
