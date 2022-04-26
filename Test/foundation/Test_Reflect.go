package foundation

import (
	"container/list"
	"fmt"
	"reflect"
)

func Reflect() {
	//testType()
	//enumMethod()
	//enumField()
	editObject()
}

// 获取数据类型
func testType() {
	var a uint16 = 1314

	var tp reflect.Type = reflect.TypeOf(a)
	fmt.Printf("变量类型是：%s\n", tp.Name())
	fmt.Printf("变量占用的内存大小是：%d\n", tp.Size())
}

// 检查数据类型
func checkType(o interface{}) {
	var tp = reflect.TypeOf(o)
	switch tp.Kind() {
	case reflect.Bool:
		fmt.Printf("布尔类型\n")
	case reflect.Int:
		fmt.Printf("有符号整数\n")
	case reflect.Uint:
		fmt.Printf("无符号整数\n")
	}
}

type song struct {
	songName string
	singer   string
}
type player struct {
	Name string
	list list.List
	song
}

func (x player) Start() {

}
func (x player) Stop(isClosing bool) int {
	return 0
}

// 枚举结构体方法
func enumMethod() {
	var obj = player{}
	ty := reflect.TypeOf(obj)
	mn := ty.NumMethod()

	for x := 0; x < mn; x++ {
		mt := ty.Method(x)
		fmt.Printf("方法名称是%s\n", mt.Name)
		fmt.Printf("方法的函数类型是%s\n", mt.Type)
	}
}

// 枚举结构体字段
func enumField() {
	var obj = player{}
	ty := reflect.TypeOf(obj)
	//a := ty.NumIn()
	//fmt.Printf("输入参数个数%d", a)
	fn := ty.NumField()

	for x := 0; x < fn; x++ {
		mt := ty.Field(x)
		fmt.Printf("字段名称是%s\n", mt.Name)
		fmt.Printf("字段类型是%s\n", mt.Type)
		fmt.Printf("包路径是%s\n", mt.PkgPath) // 只会打印非公共成员的路径
	}
}

// 修改对象的值
func editObject() {
	var name string = "小王"

	// 此时获取的是string类型的副本，不能对值进行修改
	//objValue := reflect.ValueOf(name)
	// 此时获取的数据类型是* string，是指向name变量的地址
	objValue := reflect.ValueOf(&name)
	// 获取变量的实际的值
	objValue = objValue.Elem()
	if objValue.Kind() == reflect.String {
		if objValue.CanSet() {
			objValue.Set(reflect.ValueOf("小迪"))
		} else {
			fmt.Printf("不能修改变量\n")
		}
	}
	fmt.Printf("变量的值是%s", name)
}
