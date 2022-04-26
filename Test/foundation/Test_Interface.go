package foundation

import "fmt"

type iFile1 interface {
	getFileName() string
}

type iFile2 interface {
	iFile1
	getTypeExt() string
}

type file struct {
	name, ext string
}

func (f *file) getFileName() string {
	return f.name
}

func (f *file) getTypeExt() string {
	return f.ext
}

func Interface() {
	var f iFile2 = &file{
		name: "wyd",
		ext:  "ut",
	}

	fmt.Printf("文件名是 %s,\n 文件扩展名是 %s\n", f.getFileName(), f.getTypeExt())
}

type iPrint interface {
	MyPrint()
}

type IS1 struct {
	A, B int
	S    string
}

type IS2 struct {
	S string
}

func (i IS1) MyPrint() {
	fmt.Println(i.S)
}

func (i IS2) MyPrint() {
	fmt.Println(i.S)
}

func Main() {
	var is1 iPrint
	s1 := IS1{0, 1, "Hello, world!1"}
	is1 = s1
	is1.MyPrint()

}
