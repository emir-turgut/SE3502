package main

import "fmt"

func main() {

	matrix := [][]int{
		{0, 1, 0},
		{1, 0, 1},
		{0, 1, 0},
	}

	fmt.Println(matrix)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][j] = 1
			} else {
				matrix[i][j] = 0
			}
		}
	}

	fmt.Println(matrix)
}
