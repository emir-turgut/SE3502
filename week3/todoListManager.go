package main

import "fmt"

func deleteTask(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	todo := make([]string, 4, 10)

	todo[0] = "Wash the dishes"
	todo[1] = "Vacuum the house"
	todo[2] = "Do the daily assignments"
	todo[3] = "Practice calculus"

	fmt.Println(todo)
	todo = deleteTask(todo, 2)
	fmt.Println(todo)
}
