package main
import "fmt"
func main() {
	x:= []string{"James", "Bond", "Shaken, not stirred"}
	y:=[]string{"Miss", "Moneypenny", "Helloooooo, James."}
	z:=[][]string{x,y}
	fmt.Println(z)
	var a []string

	a = append(x,y...)
	fmt.Println(a)


}
