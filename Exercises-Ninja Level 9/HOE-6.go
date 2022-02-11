package main
import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go Routines:", runtime.NumGoroutine())
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Arch:", runtime.GOARCH)
	fmt.Println("Version:", runtime.Version())
	fmt.Println("NumCPU:", runtime.NumCPU())	

}