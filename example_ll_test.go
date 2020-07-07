/* This file provides an example of choose.Generic. */

package choose_test

import (
	"container/list"
	"fmt"
	"math"
	"math/cmplx"

	"github.com/spakin/choose"
)

// A CompList is a linked list that implements the choose.Container interface.
type CompList struct {
	L *list.List
}

// New creates a new CompList of a given length.
func (cl CompList) New(n int) choose.Container {
	ll := list.New()
	for i := 0; i < n; i++ {
		ll.PushFront(nil)
	}
	return CompList{L: ll}
}

// Len returns the length of a CompList.
func (cl CompList) Len() int {
	return cl.L.Len()
}

// Get returns the ith element of a CompList.
func (cl CompList) Get(i int) interface{} {
	e := cl.L.Front()
	for ; i > 0; i-- {
		e = e.Next()
	}
	return e.Value
}

// Set sets the ith element of a CompList.
func (cl CompList) Set(i int, v interface{}) {
	e := cl.L.Front()
	for ; i > 0; i-- {
		e = e.Next()
	}
	e.Value = v
}

// Demonstrate taking subsets of a non-slice container.  Here, we use a linked
// list, which is horribly inefficient for random access but suffices for
// pedagogical purposes.
func ExampleGeneric() {
	// Store all sixth roots of unity in a CompList.
	const n = 6
	rou := CompList{}.New(n)
	for k := 0; k < n; k++ {
		kc := complex(float64(k), 0.0)
		rou.Set(k, cmplx.Exp(kc*2.0i*math.Pi/n))
	}

	// Output all 3-element subsets.
	for ri := range choose.Generic(rou, 3) {
		r := ri.(CompList)
		r0 := r.Get(0).(complex128)
		r1 := r.Get(1).(complex128)
		r2 := r.Get(2).(complex128)
		fmt.Printf("%7.4f + %7.4f + %7.4f = %7.4f\n", r0, r1, r2, r0+r1+r2)
	}
	// Output:
	// (-1.0000+0.0000i) + (-0.5000-0.8660i) + ( 0.5000-0.8660i) = (-1.0000-1.7321i)
	// ( 1.0000+0.0000i) + (-0.5000-0.8660i) + ( 0.5000-0.8660i) = ( 1.0000-1.7321i)
	// ( 0.5000+0.8660i) + (-0.5000-0.8660i) + ( 0.5000-0.8660i) = ( 0.5000-0.8660i)
	// (-0.5000+0.8660i) + (-0.5000-0.8660i) + ( 0.5000-0.8660i) = (-0.5000-0.8660i)
	// (-0.5000+0.8660i) + (-1.0000+0.0000i) + ( 0.5000-0.8660i) = (-1.0000+0.0000i)
	// ( 1.0000+0.0000i) + (-1.0000+0.0000i) + ( 0.5000-0.8660i) = ( 0.5000-0.8660i)
	// ( 0.5000+0.8660i) + (-1.0000+0.0000i) + ( 0.5000-0.8660i) = ( 0.0000+0.0000i)
	// ( 0.5000+0.8660i) + (-0.5000+0.8660i) + ( 0.5000-0.8660i) = ( 0.5000+0.8660i)
	// ( 1.0000+0.0000i) + (-0.5000+0.8660i) + ( 0.5000-0.8660i) = ( 1.0000+0.0000i)
	// ( 1.0000+0.0000i) + ( 0.5000+0.8660i) + ( 0.5000-0.8660i) = ( 2.0000+0.0000i)
	// ( 1.0000+0.0000i) + ( 0.5000+0.8660i) + (-0.5000+0.8660i) = ( 1.0000+1.7321i)
	// ( 1.0000+0.0000i) + ( 0.5000+0.8660i) + (-1.0000+0.0000i) = ( 0.5000+0.8660i)
	// ( 1.0000+0.0000i) + (-0.5000+0.8660i) + (-1.0000+0.0000i) = (-0.5000+0.8660i)
	// ( 0.5000+0.8660i) + (-0.5000+0.8660i) + (-1.0000+0.0000i) = (-1.0000+1.7321i)
	// ( 0.5000+0.8660i) + (-0.5000+0.8660i) + (-0.5000-0.8660i) = (-0.5000+0.8660i)
	// ( 1.0000+0.0000i) + (-0.5000+0.8660i) + (-0.5000-0.8660i) = (-0.0000+0.0000i)
	// ( 1.0000+0.0000i) + ( 0.5000+0.8660i) + (-0.5000-0.8660i) = ( 1.0000+0.0000i)
	// ( 1.0000+0.0000i) + (-1.0000+0.0000i) + (-0.5000-0.8660i) = (-0.5000-0.8660i)
	// ( 0.5000+0.8660i) + (-1.0000+0.0000i) + (-0.5000-0.8660i) = (-1.0000+0.0000i)
	// (-0.5000+0.8660i) + (-1.0000+0.0000i) + (-0.5000-0.8660i) = (-2.0000+0.0000i)
}
