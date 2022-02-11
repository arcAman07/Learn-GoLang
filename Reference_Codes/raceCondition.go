package main
import (
	"fmt"
	"runtime"
	"sync"
)
func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Arch:", runtime.GOARCH)
	var wg sync.WaitGroup
	counter := 0
	const gs = 100
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func(){
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
	
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
	// Code that checks for race Conditions
}