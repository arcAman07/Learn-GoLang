package main
import "fmt"
type person struct {
	first string
	last string
	iceCream []string
}
func main() {
	x := person{
		first:    "Aman",
		last:     "Sharma",
		iceCream: []string{"chocolate", "vanilla", "blueberry"},
	}
	y := person{
		first:    "Deepak",
		last:     "Sharma",
		iceCream: []string{"blackberry", "mango", "coconut"},
	}
	a := map[string] person{
		x.first:x,
		y.first:y,
	}
	fmt.Println(a)
	for i,v:= range a {
		fmt.Println(v)
		fmt.Println(i)

	}

}
