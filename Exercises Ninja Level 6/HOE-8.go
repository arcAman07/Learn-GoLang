package main
import "fmt"
func main() {
	fmt.Printf("Aman is %v years old in 2021",foo(18)())

}
func foo(a int) func() int {
	return func() int {
		return a
	}
}
