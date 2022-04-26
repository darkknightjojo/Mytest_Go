package foundation

import "fmt"

func Test() {
	// 测试指针
	var pt point
	pt.x = 10.5
	pt.y = 5.5
	fmt.Printf("调用前： % + v \n", pt)
	clean(&pt)
	fmt.Printf("调用后： % + v \n", pt)
}

type point struct {
	x float32
	y float32
}

func clean(p *point) {
	p.x = 0.0
	p.y = 0.0
}
