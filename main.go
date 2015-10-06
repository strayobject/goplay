package main

import (
	"fmt"
	"sort"
)

func main() {
	lnHeader("slice")
	slice()
	lnHeader("sort")
	sortInt()
}

func sortInt() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	var lowest int

	for i := 0; i < len(x); i++ {
		if lowest == 0 || lowest > x[i] {
			lowest = x[i]
		}
	}
	fmt.Println(lowest)
	fmt.Println("------")

	sort.Ints(x)

	fmt.Println(x[0])
}

func slice() {
	// slice containing 3 items
	slice1 := []int{1, 2, 3}
	// make an empty slice
	slice2 := make([]int, 2, 5)
	slice3 := append(slice2, 4)
	// we get [0 0 4] as expected
	fmt.Println(slice3)
	copy(slice2, slice1)
	// slice3 holds a reference to slice2 so it ends up being [1 2 4]
	// how do we get a copy of slice2 ?
	// if we declare slice2 := make([]int, 2) it will work,
	// but it's not a solution
	fmt.Println(slice1, slice2, slice3)
	// we can grow them as much as we want
	slice1 = append(slice1, 22, 33, 43)
	slice2 = append(slice2, 3, 4, 5, 6, 7)
	slice3 = append(slice3, 4, 5)
	fmt.Println(slice1, slice2, slice3)
}

func lnHeader(name string) {
	fmt.Println("\n############### " + name + " ###############\n")
}
