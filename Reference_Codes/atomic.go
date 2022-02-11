package main
import (
	"fmt"
	"runtime"
	"sync/atomic"
	"sync"
)
func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Arch:", runtime.GOARCH)
	var wg sync.WaitGroup
	var counter int64 = 0
	const gs = 100
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func(){
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Count: ", atomic.LoadInt64(&counter))
			wg.Done()
		}()
	
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Final Count:", counter)
	// Code that checks for race Conditions
}