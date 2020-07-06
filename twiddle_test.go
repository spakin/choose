/* This file tests our ability to produce combinations. */

package combinations_test

import (
	"testing"
	"github.com/spakin/combinations"
)

func TestStrings(t *testing.T) {
	a := []string{"foo", "bar", "baz", "quux"}
	_ = combinations.Strings(a, 2)
}
