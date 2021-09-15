package main
import "fmt"
func main() {
	defer bat()
	foo()
}
func foo() {

	defer func(){
		fmt.Println("How are you doing in life?")
	}()
	fmt.Println("Hello Everyone")

}
func bat() {
	fmt.Println("Bye Everyone")
}
