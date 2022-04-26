package multi_process

import (
	"fmt"
	"sync"
	"time"
)

var (
	ready     = false
	singerNum = 5
)

func Sing(singerId int, c *sync.Cond) {
	//fmt.Printf("Singer (%d) is ready\n", singerId)
	c.L.Lock()
	for !ready {
		fmt.Printf("Singer (%d) is waiting\n", singerId)
		c.Wait()
	}
	fmt.Printf("Singer (%d) sing a song\n", singerId)
	ready = false
	c.L.Unlock()
}

func Cond() {
	cond := sync.NewCond(&sync.Mutex{})
	for i := 0; i < singerNum; i++ {
		go Sing(i, cond)
	}
	time.Sleep(3 * time.Second)

	for i := 0; i < singerNum; i++ {
		ready = true
		//cond.Signal()
		cond.Broadcast()
		time.Sleep(3 * time.Second)
	}
}
