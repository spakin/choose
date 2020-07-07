/*
This file defines a few wrappers for the twiddle algorithm.  It
comprises manually written functions while wrappers2.go comprises
generated code.
*/

package choose

import (
	"fmt"
	"reflect"
)

// Slice returns all length-M combinations of an arbitrary slice one at a
// time on a channel.
func Slice(a interface{}, m int) <-chan interface{} {
	// Initialize our state.
	av := reflect.ValueOf(a)
	if av.Kind() != reflect.Slice {
		panic(fmt.Sprintf("expected slice but received %s", av.Kind()))
	}
	ch := make(chan interface{}, 100)
	n := av.Len()
	st := newState(n, m)
	at := reflect.TypeOf(a)
	cv := reflect.MakeSlice(at, m, m)
	for i := 0; i < m; i++ {
		cv.Index(i).Set(av.Index(n - m + i))
	}

	// Spawn a goroutine to write all combinations into the channel.  We
	// always return a copy of the combination rather than the original
	// because the combination itself is modified in place.
	go func() {
		cvCopy := reflect.MakeSlice(at, m, m)
		reflect.Copy(cvCopy, cv)
		ch <- cvCopy.Interface()
		for st.nextCombination() {
			cv.Index(st.Z).Set(av.Index(st.X))
			cvCopy := reflect.MakeSlice(at, m, m)
			reflect.Copy(cvCopy, cv)
			ch <- cvCopy.Interface()
		}
		close(ch)
	}()
	return ch
}

// ContainerElts returns all length-M combinations of the elements of an
// arbitrary Container one at a time on a channel.
func ContainerElts(a Container, m int) <-chan Container {
	// Initialize our state.
	ch := make(chan Container, 100)
	n := a.Len()
	st := newState(n, m)
	c := a.New(m)
	for i := 0; i < m; i++ {
		c.Set(i, a.Get(n-m+i))
	}
	copyC := func() Container {
		cCopy := c.New(m)
		for i := 0; i < m; i++ {
			cCopy.Set(i, c.Get(i))
		}
		return cCopy
	}

	// Spawn a goroutine to write all combinations into the channel.  We
	// always return a copy of the combination rather than the original
	// because the combination itself is modified in place.
	go func() {
		ch <- copyC()
		for st.nextCombination() {
			c.Set(st.Z, a.Get(st.X))
			ch <- copyC()
		}
		close(ch)
	}()
	return ch
}

// Uint64Bits returns all uint64 values with M of the lower N bits set to 1.
func Uint64Bits(n, m int) <-chan uint64 {
	ch := make(chan uint64, 100)
	if n < 0 || n > 64 {
		// Instead of an error code, simply return a closed channel.
		close(ch)
		return ch
	}
	st := newState(n, m)
	var bits uint64 // Packed array of bits
	for i := n - m; i < n; i++ {
		bits |= 1 << i
	}
	go func() {
		ch <- bits
		for st.nextCombination() {
			bits |= 1 << st.X
			bits &= ^(1 << st.Y)
			ch <- bits
		}
		close(ch)
	}()
	return ch
}
