package main
import (
	"fmt"
)
type Person struct {
	first string
	last string
}
func (p Person) speak () {
	fmt.Println("My name is", p.first, p.last)
}
type human interface {
	speak()
}
func saySomething(h human) {
	h.speak()
}
func main() {
	p1 := Person{first: "Aman", last: "Sharma"}
	saySomething(&p1)
}