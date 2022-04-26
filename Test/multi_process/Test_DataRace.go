package multi_process

// 通过通道没有数据读取时就阻塞的特性，实现多协程的互斥访问，从而避免了数据竞争
var realNum = make(chan int)
var delta = make(chan int)

func SetNumber(n int) {
	realNum <- n
}

func ChangeByDelta(d int) {
	delta <- d
}

func GetNumber() int {
	return <-realNum
}

func monitor() {
	var i int
	for {
		select {
		case i = <-realNum:
		case d := <-delta:
			i += d
		case realNum <- i:
		}
	}
}

func init() {
	go monitor()
}
