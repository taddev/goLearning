package main

import (
	"fmt"
	"math/rand"
)

func main() {
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

}
