package main

import (
	"fmt"
	"os"
)

func climb(dir string) int {
	if dir == "(" {
		return 1
	}
	return -1
}

func part1(data string) int {
	sum := 0
	for i := 0; i < len(data); i++ {
		sum += climb(string(data[i]))
	}
	return sum
}

func part2(data string) int {
	sum := 0
	firstBasement := 0
	for i := 0; i < len(data); i++ {
		sum += climb(string(data[i]))
		if sum < 0 {
			firstBasement = i + 1
			break
		}
	}
	return firstBasement
}

func main() {
	input, _ := os.ReadFile("./input.txt")
	data := string(input)
	fmt.Println("Part 1:", part1(data))
	fmt.Println("Part 2:", part2(data))
}
