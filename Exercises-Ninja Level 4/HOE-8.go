package main
import "fmt"
func main() {
	x:=map[string][]string{
		`bond_james`: {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`:{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:{`Being evil`, `Ice cream`, `Sunsets`},
	}
	for i,v:=range x{
		fmt.Printf("The things %v likes to do",i)
		fmt.Println("")
		for a,b:=range v{
			fmt.Printf("\t")
			fmt.Println(a+1,b)
		}
	}

}
