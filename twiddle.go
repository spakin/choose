/*
combinations is a Go implementation of Phillip J. Chase's twiddle
algorithm, based on Matthew Belmonte's C version
(http://www.netlib.no/netlib/toms/382).  It was written by Scott Pakin
<scott-comb@pakin.org>.
*/
package combinations

// state encapsulates the twiddle state.
type state struct {
	X int
	Y int
	Z int
	P []int
	A []string // Original data (length N)
	C []string // Current combination (length M)
}

// newState initializes and returns new twiddle state.
func newState(a []string, m int) *state {
	// Initialize the p array.
	n := len(a)
	p := make([]int, n+2)
	p[0] = n + 1
	for i := 0; i < m; i++ {
		p[n-m+i+1] = i
	}
	p[n+1] = -2
	if m == 0 {
		p[1] = 1
	}

	// Initialize the c array.
	c := make([]string, m)
	for i := range c {
		c[i] = a[n-m+i]
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
