package  main
import "fmt"

type person struct {
	fullName string
}

func main() {
	p1:= person{
	}
	fmt.Println("Initial Name",p1)

	changeMe(&p1)

	fmt.Println("Changed Name",p1)


}
func changeMe(x *person) {
	x.fullName = "Kanye West"
}
