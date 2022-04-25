package client

import (
	"github.com/darkknightjojo/Mytest_Go/chatsever/protocol"
	"io"
	"log"
	"net"
	"time"
)

type TcpClient struct {
	conn      net.Conn
	name      string
	cmdReader *protocol.Reader
	cmdWriter *protocol.Writer
	incoming  chan protocol.MessCmd
}

func NewClient() *TcpClient {
	return &TcpClient{incoming: make(chan protocol.MessCmd)}
}

func (c *TcpClient) Dial(address string) error {
	log.Printf("server:%v\n", address)
	conn, err := net.Dial("tcp", address)

	if err == nil {
		c.conn = conn
	} else {
		log.Println("dial fail")
		return err
	}

	c.cmdReader = protocol.NewReader(conn)
	c.cmdWriter = protocol.NewWriter(conn)

	return nil
}

func (c *TcpClient) Start() {
	log.Printf("Starting client\n")
	//	等待界面启动
	time.Sleep(4 * time.Second)

	for {
		cmd, err := c.cmdReader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Printf("Read error %v", err)
		}

		if cmd != nil {
			switch v := cmd.(type) {
			case protocol.MessCmd:
				c.incoming <- v
			default:
				log.Printf("Unknown command!")
			}
		}
	}
}

func (c *TcpClient) Close() {
	_ = c.conn.Close()
}

func (c *TcpClient) Send(command interface{}) error {
	return c.cmdWriter.Write(command)
}

func (c *TcpClient) SetName(name string) error {
	err := c.Send(protocol.NameCmd{Name: name})
	if err == nil {
		c.name = name
	}
	return err
}

func (c *TcpClient) SendMess(message string) error {
	return c.Send(protocol.SendCmd{
		//Name: c.name,
		Message: message,
	})
}

func (c *TcpClient) InComing() chan protocol.MessCmd {
	return c.incoming
}
