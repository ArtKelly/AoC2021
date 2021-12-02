package challenges

import (
	"fmt"
	"math"

	util "github.com/ArtKelly/AoC2021/Util"
)

func Day1() {
	lines, currSonarDepth, part1Solution := getInputInts(), math.MaxInt, 0
	for _, line := range lines {
		if line >= currSonarDepth {
			part1Solution++
		}
		currSonarDepth = line
	}
	fmt.Printf("Part 1: %d\n", part1Solution)

	lines, currSonarDepth, part2Solution := getInputInts(), math.MaxInt, 0
	for i := 0; i < len(lines)-2; i++ {
		sum := lines[i] + lines[i+1] + lines[i+2]
		if sum > currSonarDepth {
			part2Solution++
		}
		currSonarDepth = sum
	}
	fmt.Printf("Part 2: %d\n", part2Solution)
}

func getInputInts() []int {
	inputs, err := util.ReadLines("Inputs/day1.txt")
	util.Check(err)
	lines, err := util.StringsArrayToInts(inputs)
	util.Check(err)
	return lines
}
