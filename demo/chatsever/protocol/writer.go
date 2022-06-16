package protocol

import (
	"fmt"
	"io"
)

type Writer struct {
	writer io.Writer
}

func NewWriter(writer io.Writer) *Writer {
	return &Writer{
		writer: writer,
	}
}

func (w *Writer) writeString(msg string) error {
	// io.Writer.Write方法返回写入字节长度与错误信息
	_, err := w.writer.Write([]byte(msg))
	return err
}

// 使用指针类型的接收器
func (w *Writer) Write(command interface{}) error {
	var err error

	// 进行类型断言，判断command的类型
	// 将类型断言的结果赋值给一个变量，方便后续调用
	switch v := command.(type) {
	case SendCmd:
		err = w.writeString(fmt.Sprintf("SEND %v\n", v.Message))
	case NameCmd:
		err = w.writeString(fmt.Sprintf("NAME %v\n", v.Name))
	case MessCmd:
		err = w.writeString(fmt.Sprintf("MESSAGE %v %v\n", v.Name, v.Message))
	default:
		err = UnknownCommand
	}
	return err
}
