package main
import (
	"fmt"
)
func main() {
	fmt.Println(Sum(1, 2, 3))
	fmt.Println(Sum(10,11,12,13))
	fmt.Println(Sum(3,4,5,6,7,8,9,10))
}

func Sum(x ...int) int {
	var sum int
	for _, v := range x {
		sum += v
	}
	return sum
}
 