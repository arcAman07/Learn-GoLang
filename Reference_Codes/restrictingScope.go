package main

import (
	"fmt"
)

func main() {
	x := 24
	fmt.Println("The value of x is", x)
	{
		y := 42
		fmt.Println("The value of y is", y)
		y++
		fmt.Println("After incremention the new value of y is", y)
	}
	x++
	fmt.Println("The new value of x is", x)

}
