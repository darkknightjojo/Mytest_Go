package foundation

import (
	"fmt"
	"unsafe"
)

func SizeOfNum() {
	var (
		n1 uint8     = 8
		n2 uint16    = 256
		n3 complex64 = 19 + 7i
	)

	fmt.Printf("8位无符号整型：%d\n", unsafe.Sizeof(n1))
	fmt.Printf("16位无符号整型：%d\n", unsafe.Sizeof(n2))
	fmt.Printf("64位复数：%d\n", unsafe.Sizeof(n3))
}
