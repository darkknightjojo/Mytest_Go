package image

import "fmt"

type image struct {
	Height int
	Weight int
	Name   string
}

type MyImage struct {
	myImage image
}

func init() {
	fmt.Println("引入image模块")
}
