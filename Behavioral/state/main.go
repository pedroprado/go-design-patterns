package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	//State as Interfaces
	sw := &Switch{State: &BaseState{}}
	sw.Off()

	sw.On()
	sw.On()

	sw.Off()
	sw.Off()

	//State as Constants and State Machine
	currentState, exitState := OffHook, OnHook

	for currentState != exitState {
		fmt.Println("The Phone is currently at: ", currentState)
		fmt.Println("Select the trigger: ")

		for i, triggerResult := range rules[currentState] {
			fmt.Println(strconv.Itoa(i), ".", triggerResult.Trigger)
		}

		input, _, _ := bufio.NewReader(os.Stdin).ReadLine()
		i, _ := strconv.Atoi(string(input))

		trigger := rules[currentState][i]
		currentState = trigger.NewState
	}

	fmt.Println("We are done using the Phone")
}

//------------------------------State using Interfaces and Structs----------------------
type Switch struct {
	State State
}

func (sw *Switch) On() {
	sw.State.On(sw)
}

func (sw *Switch) Off() {
	sw.State.Off(sw)
}

type State interface {
	On(sw *Switch)
	Off(sw *Switch)
}

type BaseState struct{}

func (bs *BaseState) On(sw *Switch) {
	fmt.Println("Switch is on")
	sw.State = &OnState{}
}

func (bs *BaseState) Off(sw *Switch) {
	fmt.Println("Switch already off")
}

type OnState struct{}

func (onState *OnState) On(sw *Switch) {
	fmt.Println("Switch is already on")
}

func (onState *OnState) Off(sw *Switch) {
	fmt.Println("Switch is off")
	sw.State = &OffState{}
}

type OffState struct{}

func (offState *OffState) On(sw *Switch) {
	fmt.Println("Switch is on")
	sw.State = &OnState{}
}

func (offState *OffState) Off(sw *Switch) {
	fmt.Println("Switch is already off")
}

//--------------------------State using Constants and State Machine-------------------------
type PhoneState int

const (
	OffHook PhoneState = iota
	Connecting
	Connected
	OnHold
	OnHook
)

func (phoneState PhoneState) String() string {
	switch phoneState {
	case OffHook:
		return "OffHook"
	case Connecting:
		return "Connecting"
	case Connected:
		return "Connected"
	case OnHold:
		return "OnHold"
	case OnHook:
		return "OnHook"
	}
	return "Unknown"
}

type Trigger int

const (
	CallDialed Trigger = iota
	HungUp
	CallConnected
	PlacedOnHold
	TakenOfHold
	LeftMessage
)

func (trigger Trigger) String() string {
	switch trigger {
	case CallDialed:
		return "CallDialed"
	case HungUp:
		return "HungUp"
	case CallConnected:
		return "CallConnected"
	case PlacedOnHold:
		return "PlacedOnHold"
	case TakenOfHold:
		return "TakenOfHold"
	case LeftMessage:
		return "LeftMessage"
	}
	return "Unknown"
}

//The Trigger and the Next State
type TriggerResult struct {
	Trigger  Trigger
	NewState PhoneState
}

//This is the state machine!
var rules = map[PhoneState][]TriggerResult{
	OffHook: {
		{CallDialed, Connecting},
	},
	Connecting: {
		{HungUp, OnHook},
		{CallConnected, Connected},
	},
	Connected: {
		{LeftMessage, OnHook},
		{HungUp, OnHook},
		{PlacedOnHold, OnHold},
	},
	OnHold: {
		{TakenOfHold, Connected},
		{HungUp, OnHook},
	},
}
