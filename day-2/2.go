package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/fchsieh/AoC2015/aocutils"
)

func read_file(file_name string) []string {
	raw_data, _ := os.ReadFile(file_name)
	data := strings.Fields(string(raw_data))
	return data
}

func min(arr []int) int {
	min := arr[0]
	for _, x := range arr {
		if x < min {
			min = x
		}
	}
	return min
}

func part1(data []string) int {
	total_num := 0
	for _, line := range data {
		// separate line by 'x'
		dimensions := strings.Split(line, "x")
		// calculate total surface area
		l, _ := strconv.Atoi(dimensions[0])
		w, _ := strconv.Atoi(dimensions[1])
		h, _ := strconv.Atoi(dimensions[2])
		total_num += 2*l*w + 2*w*h + 2*h*l + aocutils.MinArray([]int{l * w, w * h, h * l})
	}
	return total_num
}

func part2(data []string) int {
	total_num := 0
	for _, line := range data {
		// separate line by 'x'
		dimensions := strings.Split(line, "x")
		// find smallest perimeter
		l, _ := strconv.Atoi(dimensions[0])
		w, _ := strconv.Atoi(dimensions[1])
		h, _ := strconv.Atoi(dimensions[2])
		sorted_dimensions := []int{l, w, h}
		sort.Ints(sorted_dimensions)
		total_num += 2*sorted_dimensions[0] + 2*sorted_dimensions[1] + l*w*h
	}
	return total_num
}

func main() {
	data := read_file("input.txt")
	fmt.Println("Part 1:", part1(data))
	fmt.Println("Part 2:", part2(data))
}
