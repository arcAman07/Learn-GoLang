package main
import (
	"fmt"
	"sync"
	"runtime"
)
func main() {
	fmt.Println("Go Routines:", runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("Hello")
		fmt.Println("Go Routines:", runtime.NumGoroutine())
		wg.Done()
	}()
	go func() {
		fmt.Println("World")
		fmt.Println("Go Routines:", runtime.NumGoroutine())
		wg.Done()
	}()
	wg.Wait()
}