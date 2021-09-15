package main
import "fmt"
func main() {
	x:= []int{1,7,14,18,9,3,12,27}
	evenNo,evenIndex := even(x...)
	fmt.Println(evenNo,evenIndex)
	oddNo,oddIndex := odd(x...)
	fmt.Println(oddNo,oddIndex)
	fmt.Println("")


	evenNosum,confirmevenIndex := sum(even,x...)
	fmt.Println("The sum of even numbers is",evenNosum)
	fmt.Println("The even number index are",confirmevenIndex)

	fmt.Println("")

	oddNosum,confirmoddIndex := sum(odd,x...)
	fmt.Println("The sum of odd numbers is",oddNosum)
	fmt.Println("The odd number index are",confirmoddIndex)

	fmt.Println("")

	e:=totalSum(x...)
	fmt.Println("The total sum is",e)
	f:= evenNosum+oddNosum
	fmt.Println("Recheck total sum is",f)

}
func sum(x func(a ...int) ([]int,[]int),y ...int) (int,[]int){
	No,Index := x(y...)
	var total int
	for _,b := range No {
		total += b
	}
	return total,Index

}
// Function taking unlimited number of integers and returning the numbers which are even and their
// respective indexes in the slice of int ([] x ) datatype ( c holds the even numbers while d holds the respective indices)
// Remember after taking these unlimited number of integers the function puts all the values in x which becomes
// of the datatype [] int ( slice of int) and then we are using 'range' keyword to get the values
func even(x ...int) ([]int,[]int) {
	var c []int
	var d []int
	for a,b:=range x{
		if b%2==0 {
			c = append(c,b)
			d = append(d,a)
		}
	}
	return c,d

}
// Function taking unlimited number of integers and returning the numbers which are odd and their
// respective indexes in the slice of int ([] x ) datatype ( c holds the odd numbers while d holds the respective indices)
// Remember after taking these unlimited number of integers the function puts all the values in x which becomes
// of the datatype [] int ( slice of int) and then we are using the 'range' to get the values
func odd(x ...int) ([]int,[]int) {
	var c []int
	var d []int
	for a,b:=range x{
		if b%2==1 {
			c = append(c,b)
			d = append(d,a)
		}
	}
	return c,d

}
func totalSum(x ...int) int {
	var fullSum int
	for i:=0;i<len(x);i++ {
		fullSum += x[i]
	}
	return fullSum
}
