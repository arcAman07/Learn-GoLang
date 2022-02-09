package main

import (
	"fmt"
	"sort"
)

func main() {
	p := []int{32,7,63,12,56,23,19}
	a := []string{"James","Aman","Amit","Ajay","Amit","Ajay","Amit"}
	fmt.Printf("The Initial Array is: %v\n", p)
	sort.Ints(p)
	fmt.Printf("The Sorted Array is: %v\n", p)
	fmt.Printf("The Initial Array is: %v\n", a)
	sort.Strings(a)
	fmt.Printf("The Sorted Array is: %v\n", a)
}
