package main
import "fmt"
func main() {
	var x int = 100
	var y string = "Aman"
	var z = []int{1,2,3,4,5,6}

	fmt.Println("The value of variable x is",x)
	fmt.Println("The value of variable y is",y)
	fmt.Println("The value of variable z is",z)
	// & helps to reference a variable in golang and find its address in memory(like C/C++)
	// * helps to dereference a variable in golang from its address to get its value

	fmt.Println("The address of variable x in memory is",&x)
	fmt.Println("The address of variable y in memory is",&y)
	fmt.Println("The value of variable z in memory is",&z)
}

