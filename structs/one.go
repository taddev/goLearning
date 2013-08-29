package main

import "fmt"

type Book struct {
	Index  int
	Title  string
	Author string
}

func main() {
	fmt.Println(Book{1, "Hello", "World"})
}
