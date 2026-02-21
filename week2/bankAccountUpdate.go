package main

import (
	"fmt"
)

func withdraw(balance *float64, amount float64) bool {
	if *balance < amount {
		return false
	}
	if *balance > amount {
		*balance = *balance - amount
		return true
	}
	return false
}

func main() {
	balance := 100.0
	fmt.Println("Withdraw Status:", withdraw(&balance, 150.0), "| Balance:", balance)
	fmt.Println("Withdraw Status:", withdraw(&balance, 50.0), "| Balance:", balance)

}
