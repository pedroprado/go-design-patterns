package main

import "fmt"

func main() {

	bankAccount := BankAccount{balance: 10}

	//The mementos represent the history
	memento0 := &Memento{balance: bankAccount.balance}
	memento1 := bankAccount.Deposit(100)
	memento2 := bankAccount.Deposit(25)

	fmt.Println(memento0)
	fmt.Println(memento1)
	fmt.Println(memento2)

	bankAccountHistory := NewBankAccountHistory(100)
	bankAccountHistory.Deposit(1)
	bankAccountHistory.Deposit(2)
	bankAccountHistory.Deposit(3)

	fmt.Println("##### Bank Account with it history of changes")
	fmt.Printf("%+V\n", bankAccountHistory.history)
	fmt.Println(bankAccountHistory.currentState)
	bankAccountHistory.Undo()
	fmt.Println("#### Undoing, the current state changes the pointer to the history changes")
	fmt.Println(bankAccountHistory.currentState)
}

type Memento struct {
	balance int
}

type BankAccount struct {
	balance int
}

func (bankAccount *BankAccount) Deposit(amount int) *Memento {
	bankAccount.balance += amount
	return &Memento{balance: bankAccount.balance}
}

func (bankAccount *BankAccount) Restore(memento *Memento) {
	bankAccount.balance = memento.balance
}

//Bank Account with history of deposits (saved in Mementos)
type BankAccountHistory struct {
	balance      int
	history      []*Memento
	currentState int
}

func NewBankAccountHistory(balance int) *BankAccountHistory {
	return &BankAccountHistory{
		balance:      balance,
		history:      []*Memento{&Memento{balance}},
		currentState: 0,
	}
}

func (bankAccount *BankAccountHistory) Deposit(amount int) *Memento {
	bankAccount.balance += amount
	memento := &Memento{balance: bankAccount.balance}
	bankAccount.history = append(bankAccount.history, memento)
	bankAccount.currentState++
	return memento
}

func (bankAccount *BankAccountHistory) Restore(memento *Memento) {
	if memento != nil {
		bankAccount.balance = memento.balance
		bankAccount.history = append(bankAccount.history, memento)
		bankAccount.currentState = len(bankAccount.history) - 1
	}
}

func (bankAccount *BankAccountHistory) Undo() *Memento {
	if bankAccount.currentState > 0 {
		bankAccount.currentState--
		memento := bankAccount.history[bankAccount.currentState]
		bankAccount.balance = memento.balance
		return memento
	}
	return nil
}

func (bankAccount *BankAccountHistory) Redo() *Memento {
	if bankAccount.currentState+1 < len(bankAccount.history) {
		bankAccount.currentState++
		memento := bankAccount.history[bankAccount.currentState]
		bankAccount.balance = memento.balance
		return memento
	}
	return nil
}
