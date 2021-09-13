package main
import "fmt"
func main() {
	x:=map[string][]string{
		`bond_james`: {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`:{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:{`Being evil`, `Ice cream`, `Sunsets`},
	}
	// b refers to whether that key exists in the map(returns true/false)
	a,b := x[`no_dr`]
	fmt.Println(a,b)

	fmt.Println(x[`Aman`]) // I get default []string because aman as a key is not stored in the map
	// Real syntax to print or delete a key in a map; first check the key is present and then make the necessary changes in the map
	if c,ok:=x[`no_dr`]; ok {   // Here the if loop runs as ok is true and then the key is deleted
		fmt.Println(c,ok)
		delete(x,`no_dr`)

	}
	for i, v := range x {
		fmt.Printf("The things %v likes to do", i)
		fmt.Println("")
		for a, b := range v {
			fmt.Printf("\t")
			fmt.Println(a, b)
		}
	}

}
