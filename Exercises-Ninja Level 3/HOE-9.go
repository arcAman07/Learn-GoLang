package main
import "fmt"
func main() {
	favArtist := "Weeknd"
	switch favArtist {
	case "kanye","weeknd":
		fmt.Println("The compiler guessed it wrong")
	case "Drake":
		fmt.Println("The compiler guessed it wrong")
	case "Weeknd":
		fmt.Println("The compiler finally guessed correct my fav Artist")
	default:
		fmt.Println("I don't have any fav artists")

	}
}
