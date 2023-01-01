// Path: day-6/6.go
// Solution for day 6 of Advent of Code
package main

import (
	"fmt"

	"github.com/fchsieh/AoC2015/aocutils"
)

func part1(data []string) int {
	lights := make([][]bool, 1000)
	for i := range lights {
		lights[i] = make([]bool, 1000)
	}

	for _, line := range data {
		var x1, y1, x2, y2 int
		if _, err := fmt.Sscanf(line, "turn on %d,%d through %d,%d\n", &x1, &y1, &x2, &y2); err == nil {
			for i := x1; i <= x2; i++ {
				for j := y1; j <= y2; j++ {
					lights[i][j] = true
				}
			}
		} else if _, err := fmt.Sscanf(line, "turn off %d,%d through %d,%d\n", &x1, &y1, &x2, &y2); err == nil {
			for i := x1; i <= x2; i++ {
				for j := y1; j <= y2; j++ {
					lights[i][j] = false
				}
			}
		} else if _, err := fmt.Sscanf(line, "toggle %d,%d through %d,%d\n", &x1, &y1, &x2, &y2); err == nil {
			for i := x1; i <= x2; i++ {
				for j := y1; j <= y2; j++ {
					lights[i][j] = !lights[i][j]
				}
			}
		}
	}
	// count lights
	count := 0
	for i := range lights {
		for j := range lights[i] {
			if lights[i][j] {
				count++
			}
		}
	}
	return count
}

func part2(data []string) int {
	lights := make([][]int, 1000)
	for i := range lights {
		lights[i] = make([]int, 1000)
	}

	for _, line := range data {
		var x1, y1, x2, y2 int
		if _, err := fmt.Sscanf(line, "turn on %d,%d through %d,%d\n", &x1, &y1, &x2, &y2); err == nil {
			for i := x1; i <= x2; i++ {
				for j := y1; j <= y2; j++ {
					lights[i][j]++
				}
			}
		} else if _, err := fmt.Sscanf(line, "turn off %d,%d through %d,%d\n", &x1, &y1, &x2, &y2); err == nil {
			for i := x1; i <= x2; i++ {
				for j := y1; j <= y2; j++ {
					if lights[i][j] > 0 {
						lights[i][j]--
					}
				}
			}
		} else if _, err := fmt.Sscanf(line, "toggle %d,%d through %d,%d\n", &x1, &y1, &x2, &y2); err == nil {
			for i := x1; i <= x2; i++ {
				for j := y1; j <= y2; j++ {
					lights[i][j] += 2
				}
			}
		}
	}
	// count lights
	count := 0
	for i := range lights {
		for j := range lights[i] {
			count += lights[i][j]
		}
	}
	return count
}

func main() {
	data := aocutils.ReadPuzzle("input.txt")
	fmt.Println("Part 1:", part1(data))
	fmt.Println("Part 2:", part2(data))
}
