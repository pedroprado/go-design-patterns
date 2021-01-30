package main

import "fmt"

func main() {

	//This is the Mediator: it avoids people need to know about each other (holding references to each other).
	//They all only need to hold reference to the Mediator
	chatRoom := NewChatRoom()

	john := NewPerson("John")
	kyle := NewPerson("Kyle")
	maria := NewPerson("Maria")
	people := []*Person{john, kyle, maria}

	for _, person := range people {
		chatRoom.Join(person)
	}

	john.Say("Hello")
	john.PrivateMessage("Hi", "Maria")

}

type Person struct {
	name    string
	room    *ChatRoom
	chatLog []string
}

func NewPerson(name string) *Person {
	return &Person{name: name}
}

func (person *Person) Receive(sender string, message string) {
	msg := fmt.Sprintf("[%s receives from %s]: %s", person.name, sender, message)
	fmt.Println(msg)
}

func (person *Person) Say(message string) {
	person.room.Broadcast(person.name, message)
}

func (person *Person) PrivateMessage(message string, receiver string) {
	person.room.Message(person.name, message, receiver)
}

type ChatRoom struct {
	people []*Person
}

func NewChatRoom() *ChatRoom {
	return &ChatRoom{}
}

func (chat *ChatRoom) Join(person *Person) {
	joinMsg := person.name + " joins the chat"
	chat.Broadcast("ROOM", joinMsg)

	person.room = chat
	chat.people = append(chat.people, person)
}

func (chat *ChatRoom) Broadcast(sourceName string, message string) {
	for _, person := range chat.people {
		if person.name != sourceName {
			person.Receive(sourceName, message)
		}
	}
}

func (chat *ChatRoom) Message(sender string, message string, receiver string) {
	person := chat.FindByName(receiver)
	if person != nil {
		person.Receive(sender, message)
	}

}

func (chat *ChatRoom) FindByName(name string) *Person {
	for _, person := range chat.people {
		if person.name == name {
			return person
		}
	}
	return nil
}
