package main

import "fmt"

func divMod(numerator int, denominator int) (int, int) {
	var quotient int = numerator / denominator
	var remainder int = numerator % denominator
	return quotient, remainder
}
func main() {
	q, r := divMod(13, 4)
	fmt.Println("13 divided by 4 is", q, "with a remainder of", r)
}
