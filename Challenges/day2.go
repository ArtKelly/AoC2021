package challenges

import (
	"fmt"
	"strings"

	util "github.com/ArtKelly/AoC2021/Util"
)

func Day2() {
	inputs, horizontal, depth := getInputsStrs(), 0, 0
	for _, command := range inputs {
		command, value := parseCommandString(command)
		switch command {
		case "forward":
			horizontal += value
		case "down":
			depth += value
		case "up":
			depth -= value
		}
	}
	fmt.Printf("Part 1: %d\n", (depth * horizontal))

	inputs, horizontal, depth, aim := getInputsStrs(), 0, 0, 0
	for _, command := range inputs {
		command, value := parseCommandString(command)
		switch command {
		case "forward":
			horizontal += value
			depth += aim * value
		case "down":
			aim += value
		case "up":
			aim -= value
		}
	}
	fmt.Printf("Part 2: %d\n", (depth * horizontal))
}

func getInputsStrs() []string {
	inputs, err := util.ReadLines("Inputs/day2.txt")
	util.Check(err)
	return inputs
}

func parseCommandString(s string) (string, int) {
	commands := strings.Fields(s)
	return commands[0], util.StringtoInt(commands[1])
}
