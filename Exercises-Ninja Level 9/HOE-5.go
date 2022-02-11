package main
import (
	"fmt"
	"sync/atomic"
	"runtime"
	"sync"
)
func main() {
	fmt.Println("Go Routines:", runtime.NumGoroutine())
	fmt.Println("NumCPU:", runtime.NumCPU())
	var wg sync.WaitGroup
	var counter int64 = 0
	gs := 20000
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Count: ", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		// fmt.Println("Go Routines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}