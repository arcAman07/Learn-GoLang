package main
import "fmt"
func main() {
	var a = []int{1,2,3,4,5,6,7,8,9,12,13}
	fmt.Println("The sum of all numbers passed in the function is",sum(a...))
	var b = []int{1,2,3,4,5}
	fmt.Println("The sum of all numbers passed in the function is",bar(b))

}
//  func (r receiver) identifier(parameter(s)) (return(s)) { code} - Syntax of a general function body
func sum(x ...int) int {
	var total int = 0
	for i:=0;i<len(x);i++ {
		total+=x[i]
	}
	return total
}

func bar(x []int) int {
	var total int = 0
	for i:=0;i<len(x);i++ {
		total+=x[i]
	}
	return total
}

