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

// TestUint64Bits35C5 tests choose 5 of 35 bits in a Uint64.
func TestUint64Bits35C5(t *testing.T) {
	tally := 0
	for s := range choose.Uint64Bits(35, 5) {
		tally++
		num1s := bits.OnesCount64(s)
		if num1s != 5 {
			t.Fatalf("Expected 5 bits to be set in %064b, not %d", s, num1s)
		}
		if s >= 1<<35 {
			t.Fatalf("Value is out of range (%d > %d)", s, 1<<35)
		}
	}
	exp := 324632
	if tally != exp {
		t.Fatalf("Expected %d choices but received %d", exp, tally)
	}
}

// TestSlice8C1 tests choosing one of eight structs.
func TestSlice8C1(t *testing.T) {
	// Define a type to use.
	type Thing struct {
		N uint8
		C rune
	}

	// Iterate over 8 arbitrary things.
	a := []Thing{
		{1, 'A'},
		{2, 'B'},
		{3, 'C'},
		{5, 'D'},
		{9, 'E'},
		{32, 'F'},
		{56, 'G'},
		{144, 'H'},
	}
	var sum [2]int
	tally := 0
	for si := range choose.Slice(a, 1) {
		s := si.([]Thing)
		if len(s) != 1 {
			t.Fatalf("Expected 1 struct per selection but received %d (%v)", len(s), s)
		}
		sum[0] += int(s[0].N)
		sum[1] += int(s[0].C - 'A' + 1)
		tally++
	}

	// Ensure we received the correct number and value of items.
	if tally != 8 {
		t.Fatalf("Expected 8 unique structs but received %d", tally)
	}
	if sum[0] != 252 || sum[1] != 36 {
		t.Fatalf("Expected a sum of [252, 36] but received %v", sum)
	}
}

// A MyMap is a map that implements the choose.Container interface.
type MyMap map[int]int16

// New creates a new MyMap.
func (m MyMap) New(n int) choose.Container { return make(MyMap, n) }

// Len returns the length of a MyMap.
func (m MyMap) Len() int { return len(m) }

// Get returns an item from a MyMap.
func (m MyMap) Get(i int) interface{} { return m[i] }

// Set assigns a value to an item in a MyMap.
func (m MyMap) Set(i int, v interface{}) { m[i] = v.(int16) }

// TestMyMap8C1 tests choosing one of eight MyMaps.
func TestMyMap8C1(t *testing.T) {
	// Iterate over 8 arbitrary map entries.
	a := MyMap{
		0: 1,
		1: 2,
		2: 3,
		3: 5,
		4: 9,
		5: 32,
		6: 56,
		7: 144,
	}
	sum := 0
	tally := 0
	for si := range choose.ContainerElts(a, 1) {
		s := si.(MyMap)
		if s.Len() != 1 {
			t.Fatalf("Expected 1 item per selection but received %d (%v)", s.Len(), s)
		}
		for _, v := range s {
			sum += int(v)
		}
		tally++
	}

	// Ensure we received the correct number and value of items.
	if tally != 8 {
		t.Fatalf("Expected 8 unique items but received %d", tally)
	}
	if sum != 252 {
		t.Fatalf("Expected a sum of 252 but received %d", sum)
	}
}
