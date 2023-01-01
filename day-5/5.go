// Path: day-5/5.go
// Solution for day 5 of Advent of Code
package main

import (
	"fmt"
	"strings"

	"github.com/fchsieh/AoC2015/aocutils"
)

func part1(data []string) int {
	naughty := func(s string) bool {
		// check if it contains ab, cd, pq, or xy
		if strings.Contains(s, "ab") ||
			strings.Contains(s, "cd") ||
			strings.Contains(s, "pq") ||
			strings.Contains(s, "xy") {
			return true
		}
		// check vowels
		vowels := 0
		for _, c := range s {
			switch c {
			case 'a', 'e', 'i', 'o', 'u':
				vowels++
			}
		}
		if vowels < 3 { // not enough vowels
			return true
		}

		// check for double letters
		has_double := false
		for i := 0; i < len(s)-1; i++ {
			if s[i] == s[i+1] {
				has_double = true
				break
			}
		}
		if !has_double {
			return true
		}
		return false
	}
	count := 0
	for _, s := range data {
		if !naughty(s) {
			count++
		}
	}
	return count
}

func part2(data []string) int {
	naughty := func(s string) bool {
		// check for pairs that appear twice
		pairs := make(map[string]int) // map of pair to first appearance
		has_duplicate_pair := false
		for i := 0; i < len(s)-1; i++ {
			// check if pair has appeared before
			if j, ok := pairs[s[i:i+2]]; ok {
				if i-j > 1 { // not overlapping
					has_duplicate_pair = true
					break // has pair and not overlapping before
				}
			} else {
				pairs[s[i:i+2]] = i // first appearance of pair
			}
		}
		// check for repeats (xyx, aaa, etc.)
		has_repeats := false
		for i := 0; i < len(s)-2; i++ {
			if s[i] == s[i+2] {
				has_repeats = true
				break
			}
		}
		if has_duplicate_pair && has_repeats {
			return false // is a nice string
		}
		return true
	}
	count := 0
	for _, s := range data {
		if !naughty(s) {
			count++
		}
	}
	return count
}

func main() {
	data := aocutils.ReadPuzzle("input.txt")
	fmt.Println("Part 1:", part1(data))
	fmt.Println("Part 2:", part2(data))
}
