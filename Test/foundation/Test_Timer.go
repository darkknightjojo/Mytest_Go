package foundation

import (
	"fmt"
	"math/rand"
	"time"
)

func Timer() {
	timer := time.NewTimer(5 * time.Second)
	completed := make(chan bool) // 创建一个通道
	defer close(completed)       // 退出main函数时关闭通道

	// 匿名函数              **
	go func() {
		rand.Seed(time.Now().Unix())
		thelong := rand.Intn(10)
		// 暂停单前协程
		time.Sleep(time.Duration(thelong) * time.Second)
		completed <- true
	}()

	select {
	case <-completed:
		fmt.Println("Mission completed！")
	case <-timer.C:
		fmt.Println("Mission failed")
	}

}
