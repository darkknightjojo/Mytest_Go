package foundation

import (
	"bytes"
	"fmt"
)

func Buffer() {
	var buffer = make([]byte, 0)                  // 初始化时设置其长度为0，防止出现空字节。因为长度不为0的[]byte会使用数值0来初始化元素列表
	buffer = append(buffer, []byte("一二三四五六七")...) // append第二个参数为可变参数，不可以直接传递slice，需要对slice解包，将里面的byte元素释放出来
	var bf = bytes.Buffer{}
	bf.Write(buffer)
	var data = bf.Bytes()
	fmt.Printf("缓冲区中的数据：%v\n", data)
}
