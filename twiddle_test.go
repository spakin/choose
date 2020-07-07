/* This file tests our ability to produce combinations. */

package choose_test

import (
	"math/bits"
	"testing"

	"github.com/spakin/choose"
)

// TestStrings4C1 tests choosing one of four strings.
func TestStrings4C1(t *testing.T) {
	// Generate and store all combinations.
	a := []string{"foo", "bar", "baz", "quux"}
	seen := make(map[string]struct{}, len(a))
	for s := range choose.Strings(a, 1) {
		if len(s) != 1 {
			t.Fatalf("Expected 1 string per selection but received %d (%v)", len(s), s)
		}
		seen[s[0]] = struct{}{}
	}

	// Ensure we received the correct number and value of items.
	if len(seen) != 4 {
		t.Fatalf("Expected 4 unique strings but received %d (%v)", len(seen), seen)
	}
	for _, s := range a {
		if _, ok := seen[s]; !ok {
			t.Fatalf("Element %q does not appear in %v", s, seen)
		}
	}
}

// TestStrings4C2 tests choosing two of four strings.
func TestStrings4C2(t *testing.T) {
	// Generate and store all combinations.
	a := []string{"foo", "bar", "baz", "quux"}
	seen := make(map[[2]string]struct{}, 6)
	for s := range choose.Strings(a, 2) {
		if len(s) != 2 {
			t.Fatalf("Expected 2 strings per selection but received %d (%v)", len(s), s)
		}
		s1, s2 := s[0], s[1]
		if s1 > s2 {
			s1, s2 = s2, s1
		}
		seen[[2]string{s1, s2}] = struct{}{}
	}

	// Ensure we received the correct number and value of items.
	if len(seen) != 6 {
		t.Fatalf("Expected 6 unique string pairs but received %d (%v)", len(seen), seen)
	}
	exp := [][2]string{
		{"bar", "foo"},
		{"baz", "foo"},
		{"foo", "quux"},
		{"bar", "baz"},
		{"bar", "quux"},
		{"baz", "quux"},
	}
	for _, s := range exp {
		if _, ok := seen[s]; !ok {
			t.Fatalf("Element %#v does not appear in %#v", s, exp)
		}
	}
}

// TestInts10C1 tests choosing one of ten ints.
func TestInts10C1(t *testing.T) {
	// Generate and store all combinations.
	a := []int{53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	seen := make(map[int]struct{}, len(a))
	for s := range choose.Ints(a, 1) {
		if len(s) != 1 {
			t.Fatalf("Expected 1 int per selection but received %d (%v)", len(s), s)
		}
		seen[s[0]] = struct{}{}
	}

	// Ensure we received the correct number and value of items.
	if len(seen) != 10 {
		t.Fatalf("Expected 10 unique ints but received %d (%v)", len(seen), seen)
	}
	for _, s := range a {
		if _, ok := seen[s]; !ok {
			t.Fatalf("Element %d does not appear in %v", s, seen)
		}
	}
}

// TestInts5C3 tests choosing three of five ints.
func TestInts5C3(t *testing.T) {
	// Consider the first five powers of 2.
	a := []int{1, 2, 4, 8, 16}
	for s := range choose.Ints(a, 3) {
		if len(s) != 3 {
			t.Fatalf("Expected 3 ints per selection but received %d (%v)", len(s), s)
		}
		sum := s[0] + s[1] + s[2]
		if sum < 1 || sum > 28 || bits.OnesCount(uint(sum)) != 3 {
			t.Fatalf("Encountered invalid selection %v", s)
		}
	}
}

// TestFloat64s10C1 tests choosing one of ten float64s.
func TestFloat64s10C1(t *testing.T) {
	// Generate and store all combinations.
	a := []float64{53.1, 59.2, 61.3, 67.4, 71.5, 73.6, 79.7, 83.8, 89.9, 97.0}
	seen := make(map[float64]struct{}, len(a))
	for s := range choose.Float64s(a, 1) {
		if len(s) != 1 {
			t.Fatalf("Expected 1 float64 per selection but received %d (%v)", len(s), s)
		}
		seen[s[0]] = struct{}{}
	}

	// Ensure we received the correct number and value of items.
	if len(seen) != 10 {
		t.Fatalf("Expected 10 unique float64s but received %d (%v)", len(seen), seen)
	}
	for _, s := range a {
		if _, ok := seen[s]; !ok {
			t.Fatalf("Element %.3g does not appear in %v", s, seen)
		}
	}
}

// TestFloat64s5C3 tests choosing three of five float64s.
func TestFloat64s5C3(t *testing.T) {
	// Consider the first five powers of 2.
	a := []float64{1.0, 2.0, 4.0, 8.0, 16.0}
	for s := range choose.Float64s(a, 3) {
		if len(s) != 3 {
			t.Fatalf("Expected 3 float64s per selection but received %d (%v)", len(s), s)
		}
		sum := int(s[0] + s[1] + s[2])
		if sum < 1 || sum > 28 || bits.OnesCount(uint(sum)) != 3 {
			t.Fatalf("Encountered invalid selection %v", s)
		}
	}
}
