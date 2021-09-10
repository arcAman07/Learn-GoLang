package main
import "fmt"
var a  = 42
var b = "James Bond"
var c = true
func main() {
	s := fmt.Sprintf("%v\t%v\t%v",a,b,c)
	fmt.Println(s)



}
