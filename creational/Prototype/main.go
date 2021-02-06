package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	//Deep Copying
	john := &Person{Name: "john", Address: &Address{"Jose Maria de Souze", "Sao Pedro", "Brasil"}}

	maria := &Person{Name: "maria", Address: &Address{john.Address.Street, john.Address.City, john.Address.Country}}

	fmt.Println("########## Deep Copying ##############")
	fmt.Println(john, maria)

	//Copy Method
	johny := &Person{Name: "johny", Address: &Address{"Jose Maria de Souza", "Sao Pedro", "Brasil"}}
	jane := johny.CopySerialization()
	jane.Name = "Jane"
	jane.Address.Street = "Vicente Luiz Grosso"

	fmt.Println("######### Copy Method ##############")
	fmt.Println(johny)
	fmt.Println(jane)

	fmt.Printf("%+V\n", johny)
	fmt.Printf("%+V\n", jane)
}

type Address struct {
	Street  string
	City    string
	Country string
}

//Copy Method
func (address *Address) Copy() *Address {
	return &Address{
		Street:  address.Street,
		City:    address.City,
		Country: address.Country,
	}
}

type Person struct {
	Name    string
	Address *Address
}

//Copy Method
func (person *Person) Copy() *Person {
	return &Person{
		Name:    person.Name,
		Address: person.Address.Copy(),
	}
}

//Copy Serialization
func (person *Person) CopySerialization() *Person {
	bytes, _ := json.Marshal(person)

	copy := &Person{}
	json.Unmarshal(bytes, copy)
	return copy
}
