package main

import (
	"fmt"
	//"math/rand"
	"time"
)

func myDate() (one, two, three, four int) {
	return 1, 2, 3, 4
}

func intToSlice(numbers ...int) []int {
	return numbers
}

func main() {
	/*
		s := make([]int, 10)
		r := rand.New(rand.NewSource(99))

		for i := 0; i < len(s); i++ {
			s[i] = r.Intn(5)
		}

		b := s[1:4]

		fmt.Println("s == ", s)
		fmt.Println("b == ", b)

		s = append(s, 11, 12, 13)
		b = append(b, 14, 15, 16)

		fmt.Println("s == ", s)
		fmt.Println("b == ", b)
		fmt.Println("s == ", s)
	*/

	t := time.Now()
	//r := intToSlice(t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second())
	r := []int{t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second()}


	fmt.Println(r)
}
