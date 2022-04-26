package foundation

import (
	"container/list"
	"fmt"
)

func List() {
	myList := list.New() // 返回指向list实例的指针
	myList.PushFront(101)
	myList.PushBack(102)
	myList.PushBack(200)
	printList(myList)
	clearList(myList)
	fmt.Printf("%d", myList.Len())
}

func printList(l *list.List) {
	for x := l.Front(); x != nil; x = x.Next() {
		fmt.Println(x.Value)
	}
}

func clearList(l *list.List) {
	for x := l.Front(); x != nil; x = l.Front() {
		l.Remove(x)
	}
}
