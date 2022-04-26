package multi_process

import (
	"fmt"
	"sync"
	"time"
)

var (
	mutex sync.Mutex
	v1    int
)

func Set(i int) {
	mutex.Lock()
	fmt.Printf("set %v\n", i)
	time.Sleep(2 * time.Second)
	v1 = i
	mutex.Unlock()
}

func Read() int {
	mutex.Lock()
	fmt.Printf("read\n")
	a := v1
	mutex.Unlock()
	return a
}

func Mutex() {
	numG := 5

	var wg sync.WaitGroup
	fmt.Printf("-> %d\n", Read())
	for i := 1; i <= numG; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			Set(i)
			fmt.Printf("-> %d\n", Read())
		}(i)
	}

	wg.Wait()
}
