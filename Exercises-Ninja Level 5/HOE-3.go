package main
import "fmt"
type vehicle struct {
	doors int
	color string
}
type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main() {
	x:= truck {
		vehicle:vehicle{
			doors: 2,
			color:"red",
		},
		fourWheel: true,

	}
	y:= sedan{
		vehicle:vehicle{
			doors:4,
			color:"black",
		},
		luxury: false,
	}
	fmt.Println("The details of sedan are:")
	fmt.Printf("\tNumber of doors in the sedan are %v",y.doors)
	fmt.Println("")
	fmt.Printf("\tColor of sedan is %v",y.color)
	fmt.Println("")
	fmt.Printf("\tIs the car luxurious? %v",y.luxury)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("The details of truck are:")
	fmt.Printf("\tNumber of doors in the truck are %v",x.doors)
	fmt.Println("")
	fmt.Printf("\tColor of truck is %v",x.color)
	fmt.Println("")
	fmt.Printf("\tDoes the truck have 4 wheels? %v",x.fourWheel)
	fmt.Println("")

}
