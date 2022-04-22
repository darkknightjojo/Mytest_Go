package Test

import "fmt"

func Switch() {
	var x interface{} = "hello"
	switch x.(type) {
	case nil:
		fmt.Printf(" x 的类型 :nil")
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型")
	default:
		fmt.Printf("未知型")
	}
}

type tester interface {
	getDescription() string
}

type data1 struct{}

func (d data1) getDescription() string {
	return "Data v1"
}

type data2 struct{}

func (d data2) getDescription() string {
	return "Data v2"
}

type data3 struct{}

func (d data3) getDescription() string {
	return "Data v3"
}

func Switch_2() {
	var s tester = data3{}
	switch val := s.(type) {
	case data1:
		fmt.Println(val.getDescription())
	case data2:
		fmt.Println(val.getDescription())
	case data3:
		fmt.Println(val.getDescription())
	}
}

func Switch_3() {
	switch {
	case false:
		fmt.Println("1、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("2、case 条件语句为 true")
		fallthrough
	default:
		fmt.Println("6、默认 case")
	}
}
