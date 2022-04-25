package protocol

import (
	"bufio"
	"io"
	"log"
)

type Reader struct {
	reader *bufio.Reader
}

func NewReader(r io.Reader) *Reader {
	return &Reader{
		reader: bufio.NewReader(r),
	}
}

func (r *Reader) Read() (interface{}, error) {
	cmd, err := r.reader.ReadString(' ')

	if err != nil {
		return nil, err
	}

	// 因为是在缓冲区读，会在上次读到的地方继续读
	switch cmd {
	case "SEND ":
		message, err := r.reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		// 去掉换行符
		return SendCmd{Message: message[:len(message)-1]}, nil
	case "NAME ":
		name, err := r.reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		return NameCmd{Name: name[:len(name)-1]}, nil
	case "MESSAGE ":
		// 先读取用户名
		user, err := r.reader.ReadString(' ')
		if err != nil {
			return nil, err
		}
		// 再读取消息内容
		message, err := r.reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		return MessCmd{
			Name:    user[:len(user)-1],
			Message: message[:len(message)-1]}, nil
	default:
		log.Printf("Unknown command: %v", cmd)
	}

	return nil, UnknownCommand
}

func (r *Reader) ReadAll() (interface{}, error) {
	var commands []interface{}
	for {
		command, err := r.Read()

		if command != nil {
			commands = append(commands, command)
		}

		if err == io.EOF {
			break
		} else if err != nil {
			return commands, err
		}
	}
	return commands, nil
}
