/* This file tests our ability to produce combinations. */

package combinations_test

import (
	"github.com/spakin/combinations"
	"testing"
)

func TestStrings(t *testing.T) {
	a := []string{"foo", "bar", "baz", "quux"}
	for s := range combinations.Strings(a, 1) {
		t.Logf("Combination = %q", s) // Temporary debug code
	}
}
