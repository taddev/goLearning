package main

import (
	"fmt"
	"strings"
)

type Book struct {
	Id            int
	Title, Author string
	Previous      *Book
	Next          *Book
}

var Head *Book
var Tail *Book

func addBook(id int, title string, author string) {
	insertBook := Book{id, title, author, nil, nil}

	if Head == nil {
		Head = &insertBook
		Tail = &insertBook
	} else {
		Tail.Next = &insertBook
		Tail.Next.Previous = Tail
		Tail = Tail.Next
	}
}

func printMenu() string {
	var c string

	fmt.Println("Menu")
	fmt.Println("A: Option A")
	fmt.Println("B: Option B")
	fmt.Println("C: Option C")
	fmt.Println("Q: Quit")
	fmt.Printf("Make your choice: ")
	fmt.Scanf("%s\n", &c)
	c = strings.ToLower(c)

	return c
}

func printList() {
	if Head == nil {
		fmt.Println("Empty List")
	} else {
		for b := Head; b != nil; b = b.Next {
			fmt.Println("Book ID: ", b.Id)
			fmt.Println("Book Author: ", b.Author)
			fmt.Println("Book Title: ", b.Title)
			
			if b.Next != nil {
				fmt.Println("*****")
			}
		}
	}
}

func main() {
	//books := make(map[*Book]int)
	//books[Book{"Mike Chriton", "Timeline"}] = 1

	addBook(1, "Tom Hanks", "Cast Away")
	addBook(2, "Bill Hicks", "Smoke if you Got'em")

	printList()
}
