package main

import (
	"fmt"
	"math/rand"
	"os"
	//"text/tabwriter"
	"strconv"
	"time"
	"sort"
)
/*
func myQuicksort (list []int) []int {
	if len(list) <= 1 {
		return list
	}


}

func findPivot(list []int) int {
	listLen = len(list)
	if listLen < 3 {
		return list[0]
	}

	first := list[0]
	middle := list[listLen/2]
	last := list[listLen]

	if first > middle && first < last {
		return first
	}

	if middle > first && middle < last {
		return middle
	}

	if last > first && last < 
} 
*/

type ByNumb []int

func (a ByNumb) Len() int {
	return len(a)
}

func (a ByNumb) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByNumb) Less(i, j int) bool {
	return a[i] < a[j]
}

func main() {
	count, _ := strconv.Atoi(os.Args[1])
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	list := make([]int, count)

	for i := 0; i < count; i++ {
		list[i] = r.Intn(100)
	}

	sort.Sort(ByNumb(list))

	fmt.Println(list)
}
