/*
Package combinations is a Go implementation of Phillip J. Chase's
twiddle algorithm, based on Matthew Belmonte's C version
(http://www.netlib.no/netlib/toms/382).  It was written by Scott Pakin
<scott-comb@pakin.org>.
*/
package combinations

import (
	"fmt"
	"reflect"
)

// state encapsulates the twiddle state.
type state struct {
	X int
	Y int
	Z int
	P []int
	A interface{} // Original data (slice of length N)
	C interface{} // Current combination (slice of length M)
}

// newState initializes and returns new twiddle state.
func newState(a interface{}, m int) *state {
	// Determine properties of a.
	av := reflect.ValueOf(a)
	at := reflect.TypeOf(a)
	if av.Kind() != reflect.Slice {
		panic(fmt.Sprintf("expected slice but received %s", av.Kind()))
	}
	n := av.Len()

	// Initialize the p slice.
	p := make([]int, n+2)
	p[0] = n + 1
	for i := 0; i < m; i++ {
		p[n-m+i+1] = i
	}
	p[n+1] = -2
	if m == 0 {
		p[1] = 1
	}

	// Initialize the c slice.
	c := reflect.MakeSlice(at, m, m)
	for i := 0; i < m; i++ {
		c.Index(i).Set(av.Index(n - m + i))
	}

	// Create new state and return it.
	return &state{
		P: p,
		A: a,
		C: c,
	}
}

// Strings returns all length-M combinations of a slice of strings.
func Strings(a []string, m int) <-chan string {
	ch := make(chan string, 100)
	return ch
}
