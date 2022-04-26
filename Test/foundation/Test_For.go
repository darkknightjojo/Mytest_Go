package foundation

import "fmt"

func For() {
	//var str = "wangyingndi"
	//for i, x := range str{
	//	fmt.Printf("%2d --> %c\n", i, x)
	//}
	var ch = make(chan int)

	go func() {
		defer close(ch)
		ch <- 1
		ch <- 2
		ch <- 3
		ch <- 4
		ch <- 5
		ch <- 6
		ch <- 7
	}()

	for v := range ch {
		fmt.Printf("read from channel: %d\n", v)
	}
}
