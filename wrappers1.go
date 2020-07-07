/*
This file defines a few wrappers for the twiddle algorithm.
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
