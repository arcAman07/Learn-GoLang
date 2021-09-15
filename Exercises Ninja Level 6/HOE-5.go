package main
import (
	"fmt"
	"math"
)
type circle struct {
	radius float64
}
type rectangle struct {
	length float64
	breadth float64
}
type square struct {
	side float64
}
func(s circle) area() float64 {
	return math.Pi*s.radius*s.radius
}
func(s square) area() float64 {
	return s.side*s.side
}
func(s rectangle) area() float64 {
	return s.length*s.breadth
}
type shape interface {
	area() float64
}
func info(s shape){
	fmt.Println(s.area())
}
func main() {
	p1 := circle{
		radius:7.0,
	}
	p2 := square{
		side:10.0,
	}
	p3:=rectangle{
		length: 6.0,
		breadth: 4.0,
	}
	a:=p1.area()
	fmt.Println("The area of the circle is",a)

	b:=p2.area()
	fmt.Println("The area of the square is",b)

	c:=p3.area()
	fmt.Println("The area of the rectangle is",c)

	info(p1)
	info(p2)
	info(p3)

}