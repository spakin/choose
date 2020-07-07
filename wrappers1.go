/*
This file defines a few wrappers for the twiddle algorithm.  It
comprises manually written functions while wrappers2.go comprises
generated code.
*/

package choose

import "reflect"

// Slice returns all length-M combinations of an arbitrary slice one at a
// time on a channel.
func Slice(a interface{}, m int) <-chan interface{} {
	ch := make(chan interface{}, 100)
	st := newState(a, m)
	at := reflect.TypeOf(a)
	go func() {
		ch <- st.C.Interface()
		for st.nextCombination() {
			cv := reflect.MakeSlice(at, m, m)
			reflect.Copy(cv, st.C)
			cv.Index(st.Z).Set(st.A.Index(st.X))
			st.C = cv
			ch <- st.C.Interface()
		}
		close(ch)
	}()
	return ch
}

// Uint64Bits returns all uint64 values with M of the first N bits set to 1.
func Uint64Bits(n, m int) <-chan uint64 {
	ch := make(chan uint64, 100)
	if n < 0 || n > 64 {
		// Instead of an error code, simply return a closed channel.
		close(ch)
		return ch
	}
	a := make([]struct{}, n)
	st := newState(a, m)
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
