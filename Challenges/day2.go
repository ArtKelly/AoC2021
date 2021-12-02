package main

import (
	"fmt"
	"strconv"
	"strings"

	util "github.com/ArtKelly/AoC2021/Util"
)

func main() {
	// words := strings.Fields(someString)
	part1 := a()
	fmt.Printf("Part 1: %d\n", part1)
	part2 := b()
	fmt.Printf("Part 2: %d\n", part2)
}

func a() int {
	inputs, err := util.ReadLines("../Inputs/day2.txt")
	util.Check(err)
	forward := 0
	down := 0
	for _, command := range inputs {
		commands := strings.Fields(command)
		value, err := strconv.Atoi(commands[1])
		util.Check(err)
		if commands[0] == "down" {
			down += value
		} else if commands[0] == "forward" {
			forward += value
		} else {
			down -= value
		}
	}
	return down * forward
}

func b() int {
	inputs, err := util.ReadLines("../Inputs/day2.txt")
	util.Check(err)
	horizontal := 0
	depth := 0
	aim := 0
	for _, command := range inputs {
		commands := strings.Fields(command)
		value, err := strconv.Atoi(commands[1])
		util.Check(err)
		if commands[0] == "down" {
			aim += value
		} else if commands[0] == "forward" {
			horizontal += value
			depth += aim * value
		} else {
			aim -= value
		}
	}
	return depth * horizontal
}
