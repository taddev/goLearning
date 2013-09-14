package main

import (
	"fmt"
	"strings"
	"runtime"
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

func deleteBook(b *Book) {
	if b == nil {
		fmt.Println("No Book Referenced for Deletion")
		return
	}

	if Head == Tail {
		Head = nil
		Tail = nil
	} else if b == Head {
		Head = Head.Next
		Head.Previous = nil
	} else if b == Tail {
		Tail = Tail.Previous
		Tail.Next = nil
	} else {
		b.Next.Previous = b.Previous
		b.Previous.Next = b.Next
	}
	
	b = nil
	runtime.GC() //force garbage collector, for testing
}

func findBookById(id int) *Book {
	for b := Head; b != nil; b = b.Next {
		if b.Id == id {
			return b
		}
	}
	
	return nil
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

func printMenu() string {
	var c string

	fmt.Println("Menu")
	fmt.Println("A: Add Book")
	fmt.Println("D: Delete Book")
	fmt.Println("P: Print Book List")
	fmt.Println("Q: Quit")
	fmt.Printf("Make your choice: ")
	fmt.Scanf("%s\n", &c)
	c = strings.ToLower(c)

	return c
}

func main() {
	//books := make(map[*Book]int)
	//books[Book{"Mike Chriton", "Timeline"}] = 1

	addBook(1, "Tom Hanks", "Cast Away")
	addBook(2, "Bill Hicks", "Smoke if you Got'em")

	var menuChoice string

	for menuChoice != "q" {
		menuChoice = printMenu() //print menu and get input from user

		switch menuChoice {
		case "a":
			fmt.Println("You selected A")
		case "d":
			//fmt.Println("You selected C")
			b := findBookById(1)
			deleteBook(b)
		case "p":
			printList()
		case "q":
			break
		default:
			fmt.Printf("%s is not a valid choice\n", menuChoice)
		}
	}
}
