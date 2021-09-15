package main
import "fmt"
func main() {
	a:=foo()
	fmt.Printf("My favourite number is %v",a)
	fmt.Println()
	b,c:=bar()
	fmt.Printf("%v's favourite number is %v",c,b)
	fmt.Println()

}
func foo() int {
	return 7
}

func bar() (int,string) {
	return 7,"Aman"
}

