package main
import (
	"fmt"
	"encoding/json"
)
type Person struct {
	Name string `json:"Name"`
	Age int `json:"Age"`
	Number int `json:"Number"`
}

func main() {
	p := ([]byte)(`[{"Name":"James","Age":20,"Number":12345},{"Name":"Aman","Age":18,"Number":54321}]`)
	var person []Person
	err := json.Unmarshal(p, &person)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(person)

}
