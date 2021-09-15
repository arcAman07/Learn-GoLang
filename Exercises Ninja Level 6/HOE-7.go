package main
import "fmt"
func main() {
	f:=  func() {
		fmt.Println("Hello Everyone")
	}
	f()
	g:= func(x int) {
		fmt.Println("The year is",x)
	}
	g(2021)
}
