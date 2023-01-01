// Path: day-4/4.go
// Solution for day 4 of Advent of Code
package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/fchsieh/AoC2015/aocutils"
)

func worker(data []string, zeros, start, end int, secret chan int) {
	zeros_string := "" + strings.Repeat("0", zeros)
	for i := start; i < end; i++ {
		key := data[0] + fmt.Sprint(i)
		hash := md5.Sum([]byte(key))
		hash_string := hex.EncodeToString(hash[:])
		if hash_string[:zeros] == zeros_string {
			secret <- i
			return
		}
	}
	secret <- -1 // not found
}

func part1(data []string, zeros int) int {
	for i := 0; i < 1e9; i += 100000 {
		secret := make(chan int)
		go worker(data, zeros, i, i+100000, secret)
		if result := <-secret; result != -1 {
			return result
		}
	}
	return 0
}

func part1_old(data []string, zeros int) int {
	zeros_string := "" + strings.Repeat("0", zeros)
	for i := 0; i < 1e9; i++ {
		key := data[0] + fmt.Sprint(i)
		hash := md5.Sum([]byte(key))
		hash_string := hex.EncodeToString(hash[:])
		if hash_string[:zeros] == zeros_string {
			return i
		}
	}
	return 0
}

func main() {
	data := aocutils.ReadPuzzle("input.txt")
	// without go routine
	start := time.Now()
	fmt.Println("Part 1:", part1_old(data, 5), "in", time.Since(start)/time.Millisecond)
	start = time.Now()
	fmt.Println("Part 2:", part1_old(data, 6), "in", time.Since(start)/time.Millisecond)

	// with go routine
	start = time.Now()
	fmt.Println("Part 1:", part1(data, 5), "in", time.Since(start)/time.Millisecond)
	start = time.Now()
	fmt.Println("Part 2:", part1(data, 6), "in", time.Since(start)/time.Millisecond)
}
