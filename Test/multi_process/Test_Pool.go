package multi_process

import (
	"fmt"
	"sync"
	"time"
)

var byteSlicePool = sync.Pool{
	New: func() interface{} {
		b := make([]byte, 1024)
		return &b
	},
}

func Pool() {
	t1 := time.Now().Unix()

	for i := 0; i < 1000000000; i++ {
		bytes := make([]byte, 1024)
		_ = bytes
	}

	t2 := time.Now().Unix()

	for i := 0; i < 1000000000; i++ {
		bytes := byteSlicePool.Get().(*[]byte)
		_ = bytes
		byteSlicePool.Put(bytes)
	}

	t3 := time.Now().Unix()
	fmt.Printf("不使用Pool：%d s\n", t2-t1)
	fmt.Printf("使用Pool：%d s\n", t3-t2)
}
