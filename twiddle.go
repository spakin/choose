/*
Package choose returns selections of M elements from slices of length
N.  It was written by Scott Pakin <scott-sel@pakin.org>.  The code is
largely a Go implementation of Phillip J. Chase's twiddle algorithm,
based on Matthew Belmonte's C version
(http://www.netlib.no/netlib/toms/382).
*/
package choose

import (
	"reflect"
)

//go:generate go run gen-choose.go string int float64

// state encapsulates the twiddle state.
type state struct {
	X int
	Y int
	Z int
	P []int
	A reflect.Value // Original data (slice of length N)
	C reflect.Value // Current combination (slice of length M)
}

// newState initializes and returns new twiddle state.
func newState(a interface{}, m int) *state {
	// Determine properties of a.
	av := reflect.ValueOf(a)
	at := reflect.TypeOf(a)
	n := av.Len()

	// Initialize the p slice.
	p := make([]int, n+2)
	p[0] = n + 1
	for i := 0; i < m; i++ {
		p[n-m+i+1] = i + 1
	}
	p[n+1] = -2
	if m == 0 {
		p[1] = 1
	}

	// Initialize the c slice.
	cv := reflect.MakeSlice(at, m, m)
	for i := 0; i < m; i++ {
		cv.Index(i).Set(av.Index(n - m + i))
	}

	// Create new state and return it.
	return &state{
		P: p,
		A: av,
		C: cv,
	}
}

// nextCombination performs Chase's twiddle operation to advance to the next
// combination.  It returns false when there are no more combinations.
func (s *state) nextCombination() bool {
	var i, j, k int
	for j = 1; s.P[j] <= 0; j++ {
	}
	if s.P[j-1] == 0 {
		for i = j - 1; i != 1; i-- {
			s.P[i] = -1
		}
		s.P[j] = 0
		s.X = 0
		s.Z = 0
		s.P[1] = 1
		s.Y = j - 1
		return true
	}

	if j > 1 {
		s.P[j-1] = 0
	}
	for j++; s.P[j] > 0; j++ {
	}
	k = j - 1
	for i = j; s.P[i] == 0; i++ {
		s.P[i] = -1
	}
	if s.P[i] == -1 {
		s.P[i] = s.P[k]
		s.Z = s.P[k] - 1
		s.X = i - 1
		s.Y = k - 1
		s.P[k] = -1
	} else {
		if i == s.P[0] {
			return false
		}
		s.P[j] = s.P[i]
		s.Z = s.P[i] - 1
		s.P[i] = 0
		s.X = j - 1
		s.Y = i - 1
	}
	return true
}
