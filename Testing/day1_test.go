package testing

import (
	"testing"

	challenges "github.com/ArtKelly/AoC2021/Challenges"
)

func TestDay1(t *testing.T) {
	part1, part2 := challenges.Day1("../Inputs/day1_test.txt")

	if part1 != 7 {
		t.Errorf("Part1 failed. got %d, wanted %d", part1, 10)
	}

	if part2 != 5 {
		t.Errorf("Part2 failed. got %d, wanted %d", part1, 5)
	}

}
