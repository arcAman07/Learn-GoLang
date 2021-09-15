package main
import "fmt"
func main() {
	p:= struct {
		first string
		last string
	} {
		first:"Aman",
		last:"Sharma",
	}
	fmt.Printf("My name is %v %v",p.first,p.last)
}