package main
import "fmt"
type person struct {
	first string
	last string
	iceCream []string
}
func main() {
	x:=person{
		first:"Aman",
		last:"Sharma",
		iceCream:[]string{"chocolate","vanilla","blueberry"},
	}
	y:=person{
		first:"Deepak",
		last:"Sharma",
		iceCream:[]string{"blackberry","mango","coconut"},
	}
	fmt.Printf("The favourite Ice-Cream flavours of %v %v are ",x.first,y.last)
	fmt.Println(" ")

	for i,v:=range x.iceCream {
		fmt.Printf("%v)%v",i+1,v)
		fmt.Println(" ")

	}
	fmt.Printf("The favourite Ice-Cream flavours of %v %v are ",y.first,x.last)
	fmt.Println(" ")
	for a,b:=range y.iceCream {
		fmt.Printf("%v)%v",a+1,b)
		fmt.Println(" ")

	}


}
