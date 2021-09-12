package main
import "fmt"
func main() {
	i:=2002
	for{
		if (i<2022) {
			fmt.Println(i)
		} else {
			break
		}
		i=i+1
	}
}
