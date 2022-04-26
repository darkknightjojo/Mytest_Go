package multi_process

import (
	"fmt"
	"sync"
)

var m sync.Map

func Map() {
	m.Store(1, "one")
	m.Store(2, "two")
	v, ok := m.LoadOrStore(3, "three")
	fmt.Println(v, ok)
	v, ok = m.LoadOrStore(1, "this one")
	fmt.Println(v, ok)

	v, ok = m.Load(1)

	if ok {
		fmt.Printf("value is %v\n", v)
	} else {
		fmt.Println(" no ")
	}

	f := func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	}

	m.Range(f)

	m.Delete(2)
	fmt.Println(m.Load(2))
}
