// Path: day-4/4.go
// Solution for day 4 of Advent of Code
package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/fchsieh/AoC2015/aocutils"
)

func part1(data []string, zeros int) int {
	secret := 0
	zeros_string := "" + strings.Repeat("0", zeros)
	for {
		key := data[0] + fmt.Sprint(secret)
		hash := md5.Sum([]byte(key))
		hash_string := hex.EncodeToString(hash[:])
		if hash_string[:zeros] == zeros_string {
			return secret
		}
		secret++
	}
}

func main() {
	data := aocutils.ReadPuzzle("input.txt")
	//fmt.Println("Part 1:", part1([]string{"abcdef"}))
	fmt.Println("Part 1:", part1(data, 5))
	fmt.Println("Part 2:", part1(data, 6))
}
