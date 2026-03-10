package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Year   string
}

type Stringer interface {
	String() string
}

func (b Book) String() string {
	return b.Title + " by " + b.Author + " (" + b.Year + ")"
}

func main() {
	book := Book{
		Title:  "The Book",
		Author: "The Writer",
		Year:   "2004",
	}

	fmt.Println(book)
	fmt.Println(book.String())
}
