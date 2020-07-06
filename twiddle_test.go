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

// TestStrings4C1 chooses one of four strings.
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
