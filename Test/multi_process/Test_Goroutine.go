package multi_process

import (
	"fmt"
	"sync"
	"time"
)

// 因为主goroutine和自启动goroutine之间是并发执行，打印结果可能会混在一起。
func Goroutine() {
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Print(" ", i)
		}
	}()
	fmt.Println()
	for i := 10; i < 20; i++ {
		fmt.Print(" ", i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("End!")
}

// 使用WaitGroup代替sleep
func Goroutine2() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			fmt.Print(" ", x)
		}(i)
	}
	wg.Wait()
	fmt.Print("End!")
}
