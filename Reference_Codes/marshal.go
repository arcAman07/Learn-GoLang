package main
import (
	"fmt"
	"encoding/json"
)
type Person struct {
	Name string
	Age int
	Number int
	Coms []int
}
func main() {
	p1 := Person{Name:"James", Age:20, Number:12345,Coms:[]int{1,2,3,4,5}}
	fmt.Println(p1)
	fmt.Printf("%T\n",p1)
	bs, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(bs))
	fmt.Printf("%T\n",bs)
}