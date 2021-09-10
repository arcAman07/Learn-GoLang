package main
import "fmt"
var x int
var y string
var z bool
func main() {

	// %T is used to find the datatype of the declared variables defined in the scope with the var keyword
	// So here we are printing the datatypes and also their default values
	// \n is used to add new line
	fmt.Printf("%T\n",x)
	fmt.Printf("%T\n",y)
	fmt.Printf("%T\n",z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

}