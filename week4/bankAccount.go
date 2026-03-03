package main

import "fmt"

type BankAccount struct {
	balance int
}

type BankMethods interface {
	Deposit(amount *int) error
	Withdraw(amount *int) error
	GetBalance() error
}

func (account *BankAccount) Deposit(amount int) error {
	account.balance += amount
	fmt.Println("Deposited", amount)
	return nil
}

func (account *BankAccount) Withdraw(amount int) error {
	account.balance -= amount
	fmt.Println("Withdrawed", amount)
	return nil
}

func (account BankAccount) GetBalance() error {
	fmt.Println("Balance:", account.balance)
	return nil
}

func main() {
	acc := BankAccount{balance: 100}

	acc.GetBalance()

	acc.Deposit(100)

	acc.GetBalance()

	acc.Withdraw(50)

	acc.GetBalance()
}
