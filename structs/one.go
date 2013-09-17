package main

import (
	"fmt"
	"strings"
	//"runtime"
	"bufio"
	"os"
)

type Book struct {
	Id            int
	Title, Author string
	Previous      *Book
	Next          *Book
}

var Head *Book
var Tail *Book
var count int

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
	//runtime.GC() //force garbage collector, for testing
}

func findBookById(id int) *Book {
	for b := Head; b != nil; b = b.Next {
		if b.Id == id {
			return b
		}
	}
	return nil
}

func getString(inputMessage string) string {
	in := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter %s: ", inputMessage)
	inputString, _ := in.ReadString('\n')
	inputString = inputString[:len(inputString)-1] //remove the newline character from the end
	return inputString
}

func getInt(inputMessage string) int {
	var inputInt int
	
	fmt.Printf("Enter %s: ", inputMessage)
	fmt.Scanf("%d\n", &inputInt)
	fmt.Println(inputInt)
	
	return inputInt
}

func printList() {
	if Head == nil {
		fmt.Println("Empty List")
	} else {
		for b := Head; b != nil; b = b.Next {
			if( b == Head ) {
				fmt.Println(" ")
				fmt.Println("**********")
			}
			fmt.Println("Book ID: ", b.Id)
			fmt.Println("Book Author: ", b.Author)
			fmt.Println("Book Title: ", b.Title)
			fmt.Println("**********")
		}
		fmt.Println(" ")
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
	count = 1
	var menuChoice string

	for menuChoice != "q" {
		menuChoice = printMenu() //print menu and get input from user

		switch menuChoice {
		case "a":
			author := getString("Author")
			title := getString("Title")
			addBook(count, title, author)
			count++
		case "d":
			id := getInt("Id")
			b := findBookById(id)
			deleteBook(b)
		case "p":
			printList()
		case "q":
			fmt.Println("Goodbye")
		default:
			fmt.Printf("%s is not a valid choice\n", menuChoice)
		}
	}
}
