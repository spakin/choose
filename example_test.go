/* This file provides some examples of the choose package. */

package choose_test

import (
	"fmt"

	"github.com/spakin/choose"
)

// ExampleStrings demonstrates sample usage of the Strings function.  It
// outputs all sets of three names selected from a pool of five.
func ExampleStrings() {
	// Define a slice of strings from which to select sub-slices.
	stooges := []string{
		"Curly",
		"Larry",
		"Moe",
		"Shemp",
		"Joe",
	}

	// Select three strings at a time.
	for ss := range choose.Strings(stooges, 3) {
		fmt.Printf("%s, %s, and %s\n", ss[0], ss[1], ss[2])
	}

	// Output:
	// Moe, Shemp, and Joe
	// Curly, Shemp, and Joe
	// Larry, Shemp, and Joe
	// Larry, Moe, and Joe
	// Curly, Moe, and Joe
	// Curly, Larry, and Joe
	// Curly, Larry, and Moe
	// Curly, Larry, and Shemp
	// Curly, Moe, and Shemp
	// Larry, Moe, and Shemp
}

// ExampleUint64Bits demonstrates sample usage of the Uint64Bits function.  It
// outputs all 8-bit numbers with exactly two bits set.
func ExampleUint64Bits() {
	for n := range choose.Uint64Bits(8, 2) {
		fmt.Printf("%#08b (%3d)\n", n, n)
	}
	// Output:
	// 0b11000000 (192)
	// 0b10000001 (129)
	// 0b10000010 (130)
	// 0b10000100 (132)
	// 0b10001000 (136)
	// 0b10010000 (144)
	// 0b10100000 (160)
	// 0b01100000 ( 96)
	// 0b01000001 ( 65)
	// 0b01000010 ( 66)
	// 0b01000100 ( 68)
	// 0b01001000 ( 72)
	// 0b01010000 ( 80)
	// 0b00110000 ( 48)
	// 0b00100001 ( 33)
	// 0b00100010 ( 34)
	// 0b00100100 ( 36)
	// 0b00101000 ( 40)
	// 0b00011000 ( 24)
	// 0b00010001 ( 17)
	// 0b00010010 ( 18)
	// 0b00010100 ( 20)
	// 0b00001100 ( 12)
	// 0b00001001 (  9)
	// 0b00001010 ( 10)
	// 0b00000110 (  6)
	// 0b00000101 (  5)
	// 0b00000011 (  3)
}
