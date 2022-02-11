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
	var mtx sync.Mutex
	counter := 0
	const gs = 100
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func(){
			mtx.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
			mtx.Unlock()
		}()

		fmt.Println("Count: ", counter)
		// fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}