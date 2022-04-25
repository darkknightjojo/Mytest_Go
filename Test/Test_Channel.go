package Test

import (
	"fmt"
	"sync"
	"time"
)

func Channel() {
	c := make(chan int)
	go writeChan(c, 666)
	time.Sleep(1 * time.Second)
	readChan(c)
	if _, ok := <-c; ok {
		fmt.Println("Channel is open")
	} else {
		fmt.Println("Channel is closed")
	}
}

func writeChan(c chan<- int, x int) {
	fmt.Println(x)
	c <- x
	close(c)
	fmt.Println(x)
}

func readChan(c <-chan int) {
	fmt.Println("Read:", <-c)
}

const (
	noGoroutine = 5
	noTask      = 10
)

var wg sync.WaitGroup

func Channel2() {
	task := make(chan int, noTask)

	for i := 1; i < noGoroutine; i++ {
		wg.Add(1)
		go TaskProcess(task, i)
	}

	for taskNo := 1; taskNo < noTask; taskNo++ {
		task <- taskNo
	}
	close(task)
	wg.Wait()
}

func TaskProcess(task chan int, workerNo int) {
	defer wg.Done()

	for t := range task {
		fmt.Printf("Worker %d is processing Task no %d\n", workerNo, t)
	}
	fmt.Printf("Worker %d got off work.\n", workerNo)
}
