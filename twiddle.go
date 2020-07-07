/*
Package choose returns all M-element subsets of an N-element set of various
types.
*/
package choose

//go:generate go run gen-choose.go string int float64

// A Container behaves like a slice.  It can be created; its length can be
// queried; and its elements, indexed by an integer, can be read and written.
type Container interface {
	New(n int) Container      // Create a new object of length n.
	Len() int                 // Return the length of an object.
	Get(i int) interface{}    // Return the element at a given index.
	Set(i int, v interface{}) // Set the element at a given index.
}

// state encapsulates the twiddle state.
type state struct {
	X int
	Y int
	Z int
	P []int
}

// newState initializes and returns new twiddle state.  The caller is
// responsible for properly initializing and maintaining the M-element array of
// the current combination.
func newState(n, m int) *state {
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

	// Create new state and return it.
	return &state{P: p}
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
