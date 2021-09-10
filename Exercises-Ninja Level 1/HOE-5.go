package main
import "fmt"
type hotdog int
var man hotdog
var t int
func main() {
	fmt.Println(man)
	fmt.Printf("%T\n",man)
	man = 42
	fmt.Println(man)
	fmt.Printf("%T\n",man)
	t = int(man)
	fmt.Println(t)
	fmt.Printf("%T\n",t)

}

