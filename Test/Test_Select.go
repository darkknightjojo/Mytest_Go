package Test

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// goroutine超时控制
func Select() {
	var wg sync.WaitGroup
	wg.Add(1)
	//初始化
	rand.Seed(time.Now().UnixNano())
	no := rand.Intn(6)
	no *= 1000
	//	超时时间
	duration := time.Duration(int32(no)) * time.Millisecond
	fmt.Println("timeout duration is :", duration)
	wg.Done()
	if isTimeOUt(&wg, duration) {
		fmt.Println("Time out!")
	} else {
		fmt.Println("Not time out")
	}
}

func isTimeOUt(wg *sync.WaitGroup, du time.Duration) bool {
	ch1 := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		defer close(ch1)
		wg.Wait()
	}()
	select {
	case <-ch1:
		return false
	case <-time.After(du):
		return true
	}
}
