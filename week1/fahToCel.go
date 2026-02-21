package main

import "fmt"

func main() {
	var fahrenheit int
	fmt.Print("Enter the fahrenheit: ")
	fmt.Scan(&fahrenheit)
	var celsius int = (fahrenheit - 32) * 5 / 9
	fmt.Println("Celsius: ", celsius)

}
