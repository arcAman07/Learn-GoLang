package main
import (
	"fmt"
)
const(
	a1 int = 2022+iota
	b1 int = 2022+iota
	c1 int = 2022+iota
	d1 int = 2022+iota
)
func main() {
	fmt.Println("The next 4 years will be ")
	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(c1)
	fmt.Println(d1)

}
