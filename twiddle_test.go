/* This file tests our ability to produce combinations. */

package choose_test

import (
	"github.com/spakin/choose"
	"testing"
)

// stringKeys returns a string map's keys as a slice.
func stringKeys(m map[string]struct{}) []string {
	ks := make([]string, 0, len(m))
	for k := range m {
		ks = append(ks, k)
	}
	return ks
}

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
		t.Fatalf("Expected 4 unique strings but received %d (%v)", len(seen), stringKeys(seen))
	}
	for _, s := range a {
		if _, ok := seen[s]; !ok {
			t.Fatalf("Element %q does not appear in %q", s, stringKeys(seen))
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
