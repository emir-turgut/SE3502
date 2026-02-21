package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	var number int = rand.IntN(100) + 1
	var guess int
	for {
		fmt.Print("Guess: ")
		fmt.Scan(&guess)
		if guess > number {
			fmt.Println("Too high!")
			continue
		}
		if guess < number {
			fmt.Println("Too low!")
			continue
		}
		if guess == number {
			fmt.Println("You got it!")
			break
		}
	}
}
