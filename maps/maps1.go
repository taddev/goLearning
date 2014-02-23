package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Book struct {
	Title, Author string
}

var count int
var books = map[int]Book{}

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

func printBooks() {
	if len(books) == 0 {
		fmt.Println("Empty List")
	} else {
		fmt.Println(" ")
		fmt.Println("**********")
		for k, v := range books {
			fmt.Println("Book ID: ", k)
			fmt.Println("Book Author: ", v.Author)
			fmt.Println("Book Title: ", v.Title)
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
	//books := make(map[int]Book)
	//count = 0
	var menuChoice string

	for menuChoice != "q" {
		menuChoice = printMenu() //print menu and get input from user

		switch menuChoice {
		case "a":
			author := getString("Author")
			title := getString("Title")
			books[count] = Book{title, author}
			count++
		case "d":
			id := getInt("Id")
			delete(books, id)
		case "p":
			printBooks()
		case "q":
			fmt.Println("Goodbye")
		default:
			fmt.Printf("%s is not a valid choice\n", menuChoice)
		}
	}
}
