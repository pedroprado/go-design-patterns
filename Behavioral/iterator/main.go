package main

import "fmt"

func main() {

	p := &Person{"Luke", "Sky", "Walker"}
	fmt.Println("### Iteração sobre arrays ###")
	for _, name := range p.Names() {
		fmt.Println(name)
	}

	fmt.Println("### Iteração sobre canal ###")
	for name := range p.NamesGenerator() {
		fmt.Println(name)
	}

	fmt.Println("### Iteração usando um Iterator ###")
	personIterator := NewPersonIterator(p)
	for personIterator.MoveNext() {
		fmt.Println(personIterator.Value())
	}

	//TREE TRAVERSAL

	//     1
	//    / \
	//   2   3
	//        \
	//         4
	// inorder   [2,1,3,4]  left/parent/right  -> start at left most
	// preorder  [1,2,3,4]  parent/left/right  -> start at root
	// postorder [4,3,2,1]  left/right/parent  -> start at right branch, letf or right most

	fmt.Println("####Tree Traversal")
	root := NewNode(1, NewTerminalNode(2), NewNode(3, nil, NewTerminalNode(4)))
	iterator := NewInOrderIterator(root)
	for iterator.MoveNext() {
		fmt.Printf("Elemento: %d\n", iterator.GetValue())
	}
	// fmt.Println(iterator)

}

type Person struct {
	FirstName, MiddleName, LastName string
}

func (p *Person) Names() [3]string {
	return [3]string{p.FirstName, p.MiddleName, p.LastName}
}

func (p *Person) NamesGenerator() <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		if len(p.MiddleName) > 0 {
			out <- p.MiddleName
		}
		out <- p.LastName
	}()

	return out
}

//------------------- ITERATOR----------------------------
type PersonIterator struct {
	person  *Person
	current int
}

func NewPersonIterator(person *Person) *PersonIterator {
	return &PersonIterator{
		person:  person,
		current: -1,
	}
}

func (personIterator *PersonIterator) MoveNext() bool {
	personIterator.current++
	return personIterator.current < 3
}

func (personIterator *PersonIterator) Value() string {
	switch personIterator.current {
	case 0:
		return personIterator.person.FirstName
	case 1:
		return personIterator.person.MiddleName
	case 2:
		return personIterator.person.LastName
	default:
		panic("Should be between 0 and 2")
	}
}

//------------------ TREE TRAVERSAL -------------------------------

type Node struct {
	value               int
	left, right, parent *Node
}

func NewTerminalNode(value int) *Node {
	return &Node{value: value}
}

func NewNode(value int, left, right *Node) *Node {
	node := &Node{
		value: value,
		left:  left,
		right: right,
	}
	if left != nil {
		node.left.parent = node
	}
	if right != nil {
		node.right.parent = node
	}
	return node
}

type InOrderIterator struct {
	current *Node
	root    *Node
}

func NewInOrderIterator(root *Node) *InOrderIterator {
	iterator := &InOrderIterator{
		current: root,
		root:    root,
	}

	for iterator.current.left != nil {
		iterator.current = iterator.current.left
	}

	return iterator
}

func (iterator *InOrderIterator) MoveNext() bool {
	if iterator.current == nil {
		return false
	}

	if iterator.current.right != nil {
		iterator.current = iterator.current.right
		for iterator.current.left != nil {
			iterator.current = iterator.current.left
		}
		return true
	} else {
		p := iterator.current.parent
		for p != nil && iterator.current == p.right {
			iterator.current = p
			p = p.parent
		}
		iterator.current = p
		return iterator.current != nil
	}
}

func (iterator *InOrderIterator) GetValue() int {
	return iterator.current.value
}
