package main

import "fmt"

func main() {

	bankAccount := NewBankAccount(1000)

	commands := []BankAccountCommand{
		NewBankAccountCommand(bankAccount, "deposit", 100),
		NewBankAccountCommand(bankAccount, "withdraw", 1000),
		NewBankAccountCommand(bankAccount, "withdraw", 100),
		NewBankAccountCommand(bankAccount, "deposit", 95),
	}

	for _, command := range commands {
		command.Call()
	}

	fmt.Printf("Sobrou a quantia de : %v dólares\n", bankAccount.balance)
}

//-------------------------Conta--------------------------
type BankAccount struct {
	balance int
}

func NewBankAccount(balance int) *BankAccount {
	return &BankAccount{balance: balance}
}

func (bankAccount *BankAccount) Deposit(amount int) {
	bankAccount.balance = bankAccount.balance + amount
}

func (bankAccount *BankAccount) WithDraw(amount int) {
	if amount <= bankAccount.balance {
		bankAccount.balance = bankAccount.balance - amount
	}
}

func (bankAccount *BankAccount) GetBalance() {
	fmt.Println(bankAccount.balance)
}

//-------------------------Comando--------------------------
type Command interface {
	Call()
}

type BankAccountCommand struct {
	account *BankAccount
	action  string
	amount  int
}

func NewBankAccountCommand(account *BankAccount, action string, amount int) BankAccountCommand {
	return BankAccountCommand{
		account: account,
		action:  action,
		amount:  amount,
	}
}

func (bankAccountCommand BankAccountCommand) Call() {
	if bankAccountCommand.action == "deposit" {
		bankAccountCommand.account.Deposit(bankAccountCommand.amount)
	}
	if bankAccountCommand.action == "withdraw" {
		bankAccountCommand.account.WithDraw(bankAccountCommand.amount)
	}
}
