package main

import "fmt"

type Stats struct {
	Level  int
	Health int
}

type Character struct {
	Stats
	Name string
}

func main() {
	hero := Character{
		Name: "Emir",
		Stats: Stats{
			Level:  1,
			Health: 100,
		},
	}

	fmt.Println("Name", hero.Name)
	fmt.Println("Level", hero.Level)
	fmt.Println("Health", hero.Health)

}
