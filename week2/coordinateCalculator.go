package main

import "fmt"

func rectProps(length float64, width float64) (area float64, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return
}
func main() {
	a, p := rectProps(12.3, 4.5)
	fmt.Println("Length: 12.3", "Width: 4.5")
	fmt.Println("Area:", a, "Perimeter:", p)
}
