package main

func main() {
	shortSlice := make([]int, 0)
	// Allocations	Memory copies	Speed	Memory overhead
	//		~11			idk			slower		1024
	longSlice := make([]int, 0, 1000)
	// Allocations	Memory copies	Speed	Memory overhead
	//		1			0			faster		1000

}
