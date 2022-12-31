package main

import (
	"fmt"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/fchsieh/AoC2015/aocutils"
)

type coord struct {
	x int
	y int
}

func move(pos coord, line string) coord {
	switch line {
	case "^":
		pos.y++
	case "v":
		pos.y--
	case ">":
		pos.x++
	case "<":
		pos.x--
	}
	return pos
}

func part1(data []string) int {
	sep_data := strings.Split(data[0], "")

	delivered := mapset.NewSet[coord]()
	pos := coord{0, 0}

	for _, line := range sep_data {
		delivered.Add(pos)
		pos = move(pos, line)
	}
	return delivered.Cardinality()
}

func part2(data []string) int {
	sep_data := strings.Split(data[0], "")
	santa := coord{0, 0}
	robo := coord{0, 0}
	santa_delivered := mapset.NewSet[coord]()
	robo_delivered := mapset.NewSet[coord]()
	santa_delivered.Add(santa) // init
	robo_delivered.Add(robo)   // init
	for i, line := range sep_data {
		if i%2 == 0 {
			santa = move(santa, line)
			santa_delivered.Add(santa)
		} else {
			robo = move(robo, line)
			robo_delivered.Add(robo)
		}
	}
	return santa_delivered.Union(robo_delivered).Cardinality()
}

func main() {
	data := aocutils.ReadPuzzle("input.txt")
	fmt.Println("Part 1:", part1(data))
	fmt.Println("Part 2:", part2(data))
}
