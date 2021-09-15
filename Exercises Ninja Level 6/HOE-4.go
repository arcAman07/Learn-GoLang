package main
import "fmt"
type person struct {
	first string
	last string
	age int
}
func (s person) speak() int{
	fmt.Printf("My name is %v %v and my age is %v",s.first,s.last,s.age)
	fmt.Println()
	return 0
}
// used to store multiple methods of a struct(or number of structs)
//basically it is used to more multiple identifiers/structs/variables have multiple datatypes which is user defined
type showcase interface {
	speak()
}
func main() {
	p1 := person{
		first:"Aman",
		last:"Sharma",
		age:18,
	}
	p1.speak()




}
