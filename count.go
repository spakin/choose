/*
This file provides routines for counting numbers of results without
actually returning them.
*/

package choose

// Choose returns the number of ways to select M out of N items (i.e., the
// binomial coefficient "N choose M").
func Choose(n, m int) int {
	if m > n {
		return 0 // Bad input
	}
	if m > n/2 {
		m = n - m // Simplify the problem and continue.
	}
	bc := 1
	for i := 1; i <= m; i++ {
		bc = ((n - m + i) * bc) / i
	}
	return bc
}
