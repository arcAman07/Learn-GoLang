package main
import (
	"fmt"
)
func main() {
	for i:=65;i<=90;i++ {
		fmt.Printf("%v\t",i)
		fmt.Println("")
		for j:=0;j<=2;j++ {
			fmt.Printf("\t%#U",i)
			fmt.Println()

		}
	}

}

