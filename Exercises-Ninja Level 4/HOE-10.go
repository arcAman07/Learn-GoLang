package main
import "fmt"
func main() {
	x := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}
	x["Aman"] = []string{`Coding`,`Listening to Music`,`Learning new things`}
	delete(x,`moneypenny_miss`)
	for i, v := range x {
		fmt.Printf("The things %v likes to do", i)
		fmt.Println("")
		for a, b := range v {
			fmt.Printf("\t")
			fmt.Println(a, b)
		}
	}
}


