package main

import "fmt"

func swapStrings(str1 *string, str2 *string) {
	tmp := *str1
	*str1 = *str2
	*str2 = tmp
}

func main() {
	a, b := "Hello", "World"
	fmt.Println(a, b)
	swapStrings(&a, &b)
	fmt.Println(a, b)
}
