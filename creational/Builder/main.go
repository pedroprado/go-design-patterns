package main

import (
	"fmt"
	"strings"
)

func main() {

	//builder for html like strings
	b := NewHtmlBuilder("ul")
	b.AddChild("li", "hello")
	b.AddChild("li", "world")
	fmt.Println(b.String())

	//builder for diferent aspects of an object
	pb := NewPersonBuilder()
	person := pb.Lives().
		At("Rua Roberto Sandal").In("Santos").WithPostcode("11500300").
		Works().
		At("IBM").As("Salesman").WithSalary(120000).Build()

	fmt.Println(person)

	//functional builder
	builder2 := &Person2Builder{}
	person2 := builder2.Called("Dimitri").WorksAt("IBM").Build()
	fmt.Println(person2)

}

//--------Builder for contructing html like strings -------------
//Without the builder, the strings construction would be very "not user friendly"
const (
	indentSize = 2
)

type HtmlElement struct {
	name, text string
	elements   []HtmlElement
}

func (e *HtmlElement) String() string {
	return e.string(0)
}

func (e *HtmlElement) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentSize*indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n",
		i, e.name))
	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ",
			indentSize*(indent+1)))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}
	for _, el := range e.elements {
		sb.WriteString(el.string(indent + 1))
	}
	sb.WriteString(fmt.Sprintf("%s</%s>\n",
		i, e.name))
	return sb.String()
}

type HtmlBuilder struct {
	rootName string
	root     HtmlElement
}

func NewHtmlBuilder(rootName string) *HtmlBuilder {
	return &HtmlBuilder{
		rootName: rootName,
		root:     HtmlElement{name: rootName, elements: []HtmlElement{}},
	}
}
func (b *HtmlBuilder) String() string {
	return b.root.String()
}

func (b *HtmlBuilder) AddChild(childName, childText string) {
	e := HtmlElement{name: childName, text: childText, elements: []HtmlElement{}}
	b.root.elements = append(b.root.elements, e)
}

//-------------Builder for Facets-----------
//Separating the Building of diferent aspects of an object
//Aggregating builders
type Person struct {
	Street, Postcode, City string

	Company, Position string
	AnnualIncome      int
}

type PersonBuilder struct {
	person *Person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

//ways to transition among the builders
func (pb *PersonBuilder) Lives() *PersonAddresBuilder {
	return &PersonAddresBuilder{PersonBuilder: *pb}
}

//ways to transition among the builders
func (pb *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{PersonBuilder: *pb}
}

func (pb *PersonBuilder) Build() *Person {
	return pb.person
}

type PersonAddresBuilder struct {
	PersonBuilder
}

func (pb *PersonAddresBuilder) At(street string) *PersonAddresBuilder {
	pb.person.Street = street
	return pb
}

func (pb *PersonAddresBuilder) In(city string) *PersonAddresBuilder {
	pb.person.City = city
	return pb
}

func (pb *PersonAddresBuilder) WithPostcode(postCode string) *PersonAddresBuilder {
	pb.person.Postcode = postCode
	return pb
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (pb *PersonJobBuilder) At(company string) *PersonJobBuilder {
	pb.person.Company = company
	return pb
}

func (pb *PersonJobBuilder) As(position string) *PersonJobBuilder {
	pb.person.Position = position
	return pb
}

func (pb *PersonJobBuilder) WithSalary(annualIncome int) *PersonJobBuilder {
	pb.person.AnnualIncome = annualIncome
	return pb
}

//------------FUNCTIONAL BUILDER------------

type Person2 struct {
	name, position string
}

type personMod func(*Person2)
type Person2Builder struct {
	actions []personMod
}

func (b *Person2Builder) Called(name string) *Person2Builder {
	b.actions = append(b.actions, func(p *Person2) {
		p.name = name
	})
	return b
}

func (b *Person2Builder) WorksAt(position string) *Person2Builder {
	b.actions = append(b.actions, func(p *Person2) { p.position = position })
	return b
}

func (b *Person2Builder) Build() *Person2 {
	p := Person2{}
	for _, action := range b.actions {
		action(&p)
	}
	return &p
}
