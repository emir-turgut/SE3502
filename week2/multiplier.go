package main

import "fmt"

func doubleValue(x int) {
	x = x * 2
}

func doublePointer(x *int) {
	*x = *x * 2
}

func main() {
	var count int = 10
	doubleValue(count)
	fmt.Println("doubleValue:", count)
	doublePointer(&count)
	fmt.Println("doublePointer:", count)

}
