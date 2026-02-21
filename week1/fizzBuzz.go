package main

import "fmt"

func main() {
	for i := 1; i < 51; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, " FizzBuzz")
			continue
		}
		if i%3 == 0 && i%5 != 0 {
			fmt.Println(i, " Fizz")
			continue
		}
		if i%3 != 0 && i%5 == 0 {
			fmt.Println(i, " Buzz")
			continue
		}
		if i%3 != 0 && i%5 != 0 {
			fmt.Println(i)
			continue
		}
	}
}
