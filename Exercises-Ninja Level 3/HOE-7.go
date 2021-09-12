package main
import "fmt"
var profit int = 60
func main() {
	if profit<50 {
		fmt.Println("You need to work harder!")
	} else if profit>=50 && profit<100 {
		fmt.Println("You are doing just fine. ")
	} else {
		fmt.Println("Keep up the amazing work!")
	}
}
