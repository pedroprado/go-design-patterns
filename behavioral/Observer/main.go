package main

import "fmt"

func main() {

	doctor := &Doctor{}
	supporter := &Supporter{}

	observable := &Observable{}
	observable.Subscribe(doctor)
	observable.Subscribe(supporter)

	person := &Person{observable: observable, name: "Patrick"}
	person.CatchCold()

	fmt.Println("###### Observers")
	fmt.Printf("%+V\n", observable.subscribers)

	observable.Unsubscribe(doctor)
	fmt.Println("###### Observers updated")
	fmt.Printf("%+V\n", observable.subscribers)
}

type Observable struct {
	subscribers []Observer
}

func (observable *Observable) Subscribe(observer Observer) {
	observable.subscribers = append(observable.subscribers, observer)
}

func (observable *Observable) Unsubscribe(observer Observer) {
	updated := []Observer{}
	for _, subscriber := range observable.subscribers {
		if subscriber.GetName() != observer.GetName() {
			updated = append(updated, subscriber)
		}
	}
	observable.subscribers = updated
}

func (observable *Observable) Notify(data interface{}) {
	for _, subscriber := range observable.subscribers {
		subscriber.ReceiveNotification(data)
	}
}

type Person struct {
	observable *Observable
	name       string
}

func (person *Person) CatchCold() {
	msg := fmt.Sprintf("%v pegou um resfriado!", person.name)
	person.observable.Notify(msg)
}

type Observer interface {
	GetName() string
	ReceiveNotification(data interface{})
}

type Doctor struct{}

func (doctor *Doctor) GetName() string {
	return "Doctor"
}

func (doctor *Doctor) ReceiveNotification(data interface{}) {
	fmt.Printf("Doutor recebeu notificação: %v\n", data)
}

type Supporter struct{}

func (supporter *Supporter) GetName() string {
	return "Suporte"
}

func (supporter *Supporter) ReceiveNotification(data interface{}) {
	fmt.Printf("Suporte técnico recebeu a notificação: %v\n", data)
}
