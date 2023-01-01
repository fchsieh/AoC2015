// Path: day-7/7.go
// Solution for day 7 of Advent of Code
package main

import (
	"fmt"
	"math"

	"github.com/fchsieh/AoC2015/aocutils"
)

func parse_line(line string) (string, string, string, func(int, int) int) {
	left := "0"
	right := "0"
	name := ""
	op := func(a, b int) int { return 0 }

	if _, err := fmt.Sscanf(line, "%s -> %s", &left, &name); err == nil {
		op = func(a, b int) int { return a }
	} else if _, err := fmt.Sscanf(line, "%s AND %s -> %s", &left, &right, &name); err == nil {
		op = func(a, b int) int { return a & b }
	} else if _, err := fmt.Sscanf(line, "%s OR %s -> %s", &left, &right, &name); err == nil {
		op = func(a, b int) int { return a | b }
	} else if _, err := fmt.Sscanf(line, "%s LSHIFT %s -> %s", &left, &right, &name); err == nil {
		op = func(a, b int) int { return a << b }
	} else if _, err := fmt.Sscanf(line, "%s RSHIFT %s -> %s", &left, &right, &name); err == nil {
		op = func(a, b int) int { return a >> b }
	} else if _, err := fmt.Sscanf(line, "NOT %s -> %s", &left, &name); err == nil {
		op = func(a, b int) int { return (math.MaxUint16) ^ a }
	}
	return name, left, right, op
}

type Wire struct {
	name  string
	left  string
	right string
	op    func(int, int) int
}

func build_wire_map(w Wire, val_map map[string]int, wire_map map[string]Wire) int {
	if wire_val, ok := val_map[w.name]; ok {
		return wire_val
	}
	// check if left or right is a number
	left := 0
	right := 0
	if aocutils.IsNum(w.left) {
		left = aocutils.Atoi(w.left)
	} else {
		left = build_wire_map(wire_map[w.left], val_map, wire_map)
	}
	if aocutils.IsNum(w.right) {
		right = aocutils.Atoi(w.right)
	} else {
		right = build_wire_map(wire_map[w.right], val_map, wire_map)
	}
	val_map[w.name] = w.op(left, right)
	return val_map[w.name]

}

func part1(data []string) int {
	wire_map := make(map[string]Wire)
	val_map := make(map[string]int) // stores results

	for _, line := range data {
		name, left, right, op := parse_line(line)
		wire_map[name] = Wire{name, left, right, op}
		// if left or right is a number, store it in val_map
		if aocutils.IsNum(left) {
			val_map[left] = aocutils.Atoi(left)
		}
		if aocutils.IsNum(right) {
			val_map[right] = aocutils.Atoi(right)
		}
	}

	for _, wire := range wire_map {
		build_wire_map(wire, val_map, wire_map)
	}

	return val_map["a"]
}

func part2(data []string) int {
	wire_map := make(map[string]Wire)
	val_map := make(map[string]int) // stores results

	for _, line := range data {
		name, left, right, op := parse_line(line)
		wire_map[name] = Wire{name, left, right, op}
		// if left or right is a number, store it in val_map
		if aocutils.IsNum(left) {
			val_map[left] = aocutils.Atoi(left)
		}
		if aocutils.IsNum(right) {
			val_map[right] = aocutils.Atoi(right)
		}
	}

	for _, wire := range wire_map {
		build_wire_map(wire, val_map, wire_map)
	}
	// set b to the value of a
	to_set := val_map["a"]
	// delete everything in val_map
	for _, wire := range wire_map {
		delete(val_map, wire.name)
	}
	// set b to the value of a
	val_map["b"] = to_set
	// rebuild the map
	for _, wire := range wire_map {
		build_wire_map(wire, val_map, wire_map)
	}
	return val_map["a"]
}

func main() {
	data := aocutils.ReadPuzzle("input.txt")
	fmt.Println("Part 1:", part1(data))
	fmt.Println("Part 2:", part2(data))
}
