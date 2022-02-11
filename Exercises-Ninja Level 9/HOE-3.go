package main
import (
	"fmt"
	"sync"
	"runtime"
)
func main() {
	fmt.Println("Go Routines:", runtime.NumGoroutine())
	fmt.Println("NumCPU:", runtime.NumCPU())
	var wg sync.WaitGroup
	counter:=0
	gs := 20000
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("Go Routines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}