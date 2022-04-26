package foundation

import (
	"container/ring"
	"fmt"
)

func Ring() {
	myring := ring.New(7)
	n := myring.Len()
	pointer := myring
	v := 'A'

	for x := 0; x < n; x++ {
		pointer.Value = v
		pointer = pointer.Next()
		v++
	}
	printRing(myring)
	new_pt := myring.Move(3)
	fmt.Printf("\n向前滚动3个位置之后\n")
	printRing(new_pt)
}

func printRing(r *ring.Ring) {
	n := r.Len()
	pt := r

	for x := 0; x < n; x++ {
		fmt.Printf("%c  ", pt.Value)
		pt = pt.Next()
	}

}
