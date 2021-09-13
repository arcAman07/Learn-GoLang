package main
import "fmt"
func main() {
	x:=[]int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for i,v:=range x {
		fmt.Println("Index is ",i)

		fmt.Println("The element is",v)
	}
	fmt.Println(x)
	fmt.Printf("%T",x)

}
