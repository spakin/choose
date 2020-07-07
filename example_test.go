/* This file provides some examples of the choose package. */

package choose_test

import (
	"fmt"
	"sort"

	"github.com/spakin/choose"
)

// Output all sets of three names selected from a pool of five.
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

// Output all 8-bit numbers with exactly two bits set.
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

// Take two items at a time from a user-defined type.
func ExampleSlice() {
	// Define an item on a menu.
	type MenuItem struct {
		Food  string
		Price float32
	}

	// Establish some menu choices.
	menu := []MenuItem{
		{"spaghetti", 12.50},
		{"steak", 22.30},
		{"salad", 4.75},
		{"shrimp", 19.70},
		{"swedish meatballs", 13.82},
		{"squid", 18.65},
	}

	// Select, store, then sort pairs of items.
	pairs := make([]MenuItem, 0, choose.NumChoices(len(menu), 2))
	for p := range choose.Slice(menu, 2) {
		mi := p.([]MenuItem)
		combo := MenuItem{
			Food:  mi[0].Food + " and " + mi[1].Food,
			Price: mi[0].Price + mi[1].Price,
		}
		pairs = append(pairs, combo)
	}
	sort.Slice(pairs, func(i, j int) bool {
		switch {
		case pairs[i].Price > pairs[j].Price:
			return true
		case pairs[i].Price < pairs[j].Price:
			return false
		default:
			return pairs[i].Food < pairs[j].Food
		}
	})

	// Output all pairs.
	for _, p := range pairs {
		food := p.Food + " ........................................"
		fmt.Printf("%-40.40s %5.2f\n", food, p.Price)
	}

	// Output:
	// steak and shrimp ....................... 42.00
	// steak and squid ........................ 40.95
	// shrimp and squid ....................... 38.35
	// steak and swedish meatballs ............ 36.12
	// spaghetti and steak .................... 34.80
	// shrimp and swedish meatballs ........... 33.52
	// swedish meatballs and squid ............ 32.47
	// spaghetti and shrimp ................... 32.20
	// spaghetti and squid .................... 31.15
	// steak and salad ........................ 27.05
	// spaghetti and swedish meatballs ........ 26.32
	// salad and shrimp ....................... 24.45
	// salad and squid ........................ 23.40
	// salad and swedish meatballs ............ 18.57
	// spaghetti and salad .................... 17.25
}
