package main

import (
	"fmt"

	util "github.com/ArtKelly/AoC2021/Util"
)

func main() {
	part1 := part1()
	fmt.Printf("Part 1: %d\n", part1)
	part2 := part2()
	fmt.Printf("Part 2: %d\n", part2)
}

func part1() int {
	inputs, err := util.ReadLines("../Inputs/day1.txt")
	util.Check(err)

	lines, err := util.StringsArrayToInts(inputs)
	util.Check(err)
	currSonarDepth := 99999999
	solution := 0
	for _, line := range lines {
		util.Check(err)
		if line >= currSonarDepth {
			solution++
		}
		currSonarDepth = line
	}
	return solution
}

func part2() int {
	// sliding window
	inputs, err := util.ReadLines("../Inputs/day1.txt")
	util.Check(err)

	lines, err := util.StringsArrayToInts(inputs)
	util.Check(err)

	currSonarDepth := 99999999
	solution := 0

	for i := 0; i < len(lines)-2; i++ {
		sum := lines[i] + lines[i+1] + lines[i+2]
		if sum > currSonarDepth {
			solution++
		}
		currSonarDepth = sum
	}
	return solution
}
